#!/usr/bin/env node
/**
 * inject-verification-sections.mjs
 *
 * Batch generator that injects a topic-appropriate ``## Verification``
 * section into every prose markdown file under ``spec/``. Each section
 * is derived from a *folder profile* (see ``profiles.mjs``) so the body
 * stays minimal and topically relevant — no copy-paste boilerplate.
 *
 * Behaviour
 * ---------
 *  - Idempotent: if a ``## Verification`` heading already exists, the
 *    block is replaced (everything from that heading down to the next
 *    H2 or EOF). Files unchanged after replacement are left untouched.
 *  - Skips: ``00-overview.md``, ``97-acceptance-criteria.md``,
 *    ``99-consistency-report.md``, ``readme.md``, ``changelog.md``.
 *  - Each section deterministically derives its AC tag from
 *    ``<folder-prefix>-<file-prefix>`` so cross-references stay stable
 *    across runs.
 *
 * Flags
 * -----
 *   --root <dir>   Root to scan (default ``spec``)
 *   --dry-run      Print what would change, write nothing
 *   --only <glob>  Restrict to a folder substring (e.g. ``04-database``)
 *   --json         Emit JSON summary
 *
 * Exit codes
 * ----------
 *   0  Success
 *   1  At least one file failed to process (IO error)
 *   2  Invocation error
 */
import { readFileSync, writeFileSync, readdirSync, statSync } from "node:fs";
import { join, relative, basename, dirname, sep } from "node:path";
import { fileURLToPath } from "node:url";
import { FOLDER_PROFILES, DEFAULT_PROFILE, resolveCommand } from "./profiles.mjs";

const HERE = dirname(fileURLToPath(import.meta.url));
const REPO_ROOT = join(HERE, "..", "..");

const SKIP_BASENAMES = new Set([
  "97-acceptance-criteria.md",
  "99-consistency-report.md",
  "readme.md",
  "changelog.md",
]);

const VERIFICATION_HEADING = "## Verification";
const TODAY = new Date().toISOString().slice(0, 10);

function parseArgs(argv) {
  const args = { root: "spec", dryRun: false, only: null, json: false, mode: "overview-only", strip: false };
  for (let i = 2; i < argv.length; i++) {
    const a = argv[i];
    if (a === "--dry-run") args.dryRun = true;
    else if (a === "--json") args.json = true;
    else if (a === "--root") args.root = argv[++i];
    else if (a === "--only") args.only = argv[++i];
    else if (a === "--mode") args.mode = argv[++i];
    else if (a === "--strip") args.strip = true;
    else if (a === "-h" || a === "--help") { printHelp(); process.exit(0); }
    else { console.error(`Unknown flag: ${a}`); process.exit(2); }
  }
  if (!["overview-only", "all-files"].includes(args.mode)) {
    console.error(`--mode must be 'overview-only' or 'all-files', got '${args.mode}'`);
    process.exit(2);
  }
  return args;
}

function printHelp() {
  console.log(`Usage: node scripts/spec-verification/inject-verification-sections.mjs [flags]
  --root <dir>   Spec root (default 'spec')
  --only <substr> Restrict to folders whose path includes this substring
  --mode <m>      'overview-only' (default) — inject only into each folder's
                  00-overview.md. 'all-files' — inject into every prose file.
  --strip         Remove existing ## Verification blocks instead of writing
                  new ones. Useful when downsizing scope.
  --dry-run       Show plan without writing
  --json          Emit JSON summary`);
}

function walk(dir, out = []) {
  for (const entry of readdirSync(dir)) {
    const full = join(dir, entry);
    const st = statSync(full);
    if (st.isDirectory()) walk(full, out);
    else if (entry.endsWith(".md")) out.push(full);
  }
  return out;
}

function topLevelFolder(relPath) {
  const parts = relPath.split(sep);
  return parts[0] || "";
}

function pickProfile(relPath) {
  const folder = topLevelFolder(relPath);
  return FOLDER_PROFILES[folder] || DEFAULT_PROFILE;
}

function deriveTag(profile, fullPath, relPath) {
  const file = basename(fullPath, ".md");
  // file like "02-schema-design" -> 002 / "01-foo/02-bar" -> use last numeric
  const match = file.match(/^(\d+)/);
  const num = match ? match[1].padStart(3, "0") : "001";
  // Subfolder depth adds a suffix letter so siblings under different
  // subfolders don't collide (a, b, c…).
  const depth = relPath.split(sep).length;
  const suffix = depth > 2 ? String.fromCharCode(96 + Math.min(depth - 2, 26)) : "";
  return `${profile.tag}-${num}${suffix}`;
}

function buildSection(profile, fullPath, relPath) {
  const tag = deriveTag(profile, fullPath, relPath);
  const folder = topLevelFolder(relPath);
  const acRef = `\`spec/${folder}/97-acceptance-criteria.md\``;
  const fileSlug = basename(fullPath, ".md").replace(/^\d+-/, "").replace(/-/g, " ");
  const given = profile.given(fileSlug);
  const when = profile.when;
  const then = profile.then(fileSlug);
  const cmd = resolveCommand(profile);
  return [
    "",
    "---",
    "",
    VERIFICATION_HEADING,
    "",
    `_Auto-generated section — see ${acRef} for the full criteria index._`,
    "",
    `### ${tag}: ${profile.title(fileSlug)}`,
    "",
    `**Given** ${given}  `,
    `**When** ${when}  `,
    `**Then** ${then}`,
    "",
    "**Verification command:**",
    "",
    "```bash",
    cmd,
    "```",
    "",
    "**Expected:** exit 0. Any non-zero exit is a hard fail and blocks merge.",
    "",
    `_Verification section last updated: ${TODAY}_`,
    "",
  ].join("\n");
}

function replaceOrAppend(content, newSection) {
  const idx = content.indexOf(`\n${VERIFICATION_HEADING}`);
  if (idx === -1) {
    const trimmed = content.replace(/\s+$/, "");
    return trimmed + "\n" + newSection;
  }
  // Replace from the preceding `---\n` (if present) through end-of-file.
  let cutStart = idx;
  const beforeIdx = content.lastIndexOf("\n---\n", idx);
  if (beforeIdx !== -1 && beforeIdx > idx - 20) cutStart = beforeIdx;
  const head = content.slice(0, cutStart).replace(/\s+$/, "");
  return head + "\n" + newSection;
}

function shouldSkip(fullPath, relPath, mode) {
  const base = basename(fullPath).toLowerCase();
  if (SKIP_BASENAMES.has(base)) return true;
  // Spec-root files (e.g. spec/00-overview.md) own their own bespoke
  // verification logic — never touch them.
  if (relPath.split(sep).length === 1) return true;
  // overview-only mode: only touch each folder's 00-overview.md.
  if (mode === "overview-only") {
    if (base !== "00-overview.md") return true;
    // Only the top-level overview (e.g. spec/04-database-conventions/00-overview.md
    // — depth 2). Skip nested 00-overview.md files in subfolders.
    if (relPath.split(sep).length !== 2) return true;
  }
  return false;
}

function stripVerificationBlock(content) {
  const idx = content.indexOf(`\n${VERIFICATION_HEADING}`);
  if (idx === -1) return content;
  let cutStart = idx;
  const beforeIdx = content.lastIndexOf("\n---\n", idx);
  if (beforeIdx !== -1 && beforeIdx > idx - 20) cutStart = beforeIdx;
  return content.slice(0, cutStart).replace(/\s+$/, "") + "\n";
}

function main() {
  const args = parseArgs(process.argv);
  const rootAbs = join(REPO_ROOT, args.root);
  let files;
  try {
    files = walk(rootAbs);
  } catch (e) {
    console.error(`::error::cannot scan root: ${rootAbs} (${e.message})`);
    process.exit(2);
  }

  const summary = {
    scanned: 0, skipped: 0, written: 0, unchanged: 0, dryRun: args.dryRun,
    perProfile: {}, errors: [],
  };

  for (const full of files) {
    summary.scanned++;
    const rel = relative(rootAbs, full);
    if (args.only && !rel.includes(args.only)) { summary.skipped++; continue; }
    // For --strip, we visit every prose file regardless of mode so we can
    // remove blocks that earlier runs may have planted.
    if (!args.strip && shouldSkip(full, rel, args.mode)) { summary.skipped++; continue; }
    if (args.strip) {
      const base = basename(full).toLowerCase();
      // For --strip we always visit readme.md / changelog.md too, so we
      // can clean up blocks planted by earlier "all-files" runs. Only
      // 97-acceptance-criteria.md and 99-consistency-report.md are
      // sacred — those own their own verification narratives.
      const stripSacred = new Set(["97-acceptance-criteria.md", "99-consistency-report.md"]);
      if (stripSacred.has(base) || rel.split(sep).length === 1) {
        summary.skipped++; continue;
      }
      // In overview-only strip mode, do NOT strip the overview itself.
      // (Only the top-level overview — depth 2 — is preserved; nested
      // 00-overview.md files lost their per-file block earlier and are
      // now treated like any other prose file.)
      if (args.mode === "overview-only" && base === "00-overview.md" && rel.split(sep).length === 2) {
        summary.skipped++; continue;
      }
    }
    const profile = pickProfile(rel);
    summary.perProfile[profile.tag] = (summary.perProfile[profile.tag] || 0) + 1;
    let original;
    try { original = readFileSync(full, "utf8"); }
    catch (e) { summary.errors.push({ file: rel, error: e.message }); continue; }
    let next;
    if (args.strip) {
      next = stripVerificationBlock(original);
    } else {
      const section = buildSection(profile, full, rel);
      next = replaceOrAppend(original, section);
    }
    if (next === original) { summary.unchanged++; continue; }
    if (!args.dryRun) {
      try { writeFileSync(full, next, "utf8"); }
      catch (e) { summary.errors.push({ file: rel, error: e.message }); continue; }
    }
    summary.written++;
  }

  if (args.json) {
    console.log(JSON.stringify(summary, null, 2));
  } else {
    const verb = args.dryRun ? "would write" : "wrote";
    console.log(`OK Verification injector — scanned=${summary.scanned} ${verb}=${summary.written} unchanged=${summary.unchanged} skipped=${summary.skipped}`);
    const profileLines = Object.entries(summary.perProfile)
      .sort((a, b) => b[1] - a[1])
      .map(([k, v]) => `  ${k.padEnd(8)} ${v}`).join("\n");
    if (profileLines) console.log("Per-profile coverage:\n" + profileLines);
    if (summary.errors.length) {
      console.error(`FAIL ${summary.errors.length} error(s):`);
      for (const e of summary.errors) console.error(`  ${e.file}: ${e.error}`);
    }
  }
  process.exit(summary.errors.length ? 1 : 0);
}

main();