#!/usr/bin/env node
/**
 * generate-coverage-report.mjs
 *
 * Walks the spec tree, classifies each prose markdown file, and reports
 * which files carry a `## Verification` block, which are expected to but
 * don't, and which are intentionally skipped (readme/changelog/sacred
 * files / spec root).
 *
 * Two scope modes mirror the injector:
 *   - overview-only (default): only each top-level folder's
 *     ``00-overview.md`` (depth 2) is *expected* to carry a block.
 *   - all-files: every prose file outside the skip set is expected
 *     to carry a block.
 *
 * Output formats: Markdown (default) or JSON (``--json``). Reports are
 * written to ``--out <path>`` (parent dirs created on demand) and also
 * mirrored to stdout so CI logs capture the summary.
 *
 * Exit codes
 * ----------
 *   0  Success — all expected files carry a ``## Verification`` block,
 *      OR ``--strict`` was not requested.
 *   1  IO error while scanning or writing.
 *   2  Invocation error (bad flag).
 *   3  ``--strict`` (or ``--fail-on-missing``) was set AND at least
 *      one expected file is missing a ``## Verification`` block.
 */
import { readFileSync, writeFileSync, mkdirSync, readdirSync, statSync } from "node:fs";
import { join, relative, basename, dirname, sep } from "node:path";
import { fileURLToPath } from "node:url";

const HERE = dirname(fileURLToPath(import.meta.url));
const REPO_ROOT = join(HERE, "..", "..");

const VERIFICATION_HEADING = "## Verification";
const SACRED_BASENAMES = new Set([
  "97-acceptance-criteria.md",
  "99-consistency-report.md",
  "readme.md",
  "changelog.md",
]);

const VALID_MODES = ["overview-only", "all-files"];

function parseArgs(argv) {
  const args = {
    root: "spec",
    mode: "overview-only",
    json: false,
    out: null,
    strict: false,
  };
  for (let i = 2; i < argv.length; i++) {
    const a = argv[i];
    if (a === "--json") args.json = true;
    else if (a === "--root") args.root = argv[++i];
    else if (a === "--mode") args.mode = argv[++i];
    else if (a === "--out") args.out = argv[++i];
    else if (a === "--strict" || a === "--fail-on-missing") args.strict = true;
    else if (a === "-h" || a === "--help") { printHelp(); process.exit(0); }
    else { console.error(`Unknown flag: ${a}`); process.exit(2); }
  }
  if (!VALID_MODES.includes(args.mode)) {
    console.error(`--mode must be one of ${VALID_MODES.join("|")}, got '${args.mode}'`);
    process.exit(2);
  }
  return args;
}

function printHelp() {
  console.log(`Usage: node scripts/spec-verification/generate-coverage-report.mjs [flags]
  --root <dir>          Spec root (default 'spec')
  --mode <m>            'overview-only' (default) | 'all-files'
  --out <path>          Write report to file (also printed to stdout)
  --json                Emit JSON instead of Markdown
  --strict              Exit code 3 if any expected file is missing a
                        '## Verification' block (alias: --fail-on-missing)`);
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

function isSpecRootFile(rel) {
  return rel.split(sep).length === 1;
}

function isTopLevelOverview(rel) {
  const parts = rel.split(sep);
  return parts.length === 2 && parts[1].toLowerCase() === "00-overview.md";
}

/**
 * Classify a file as one of: expected, sacred-skip, root-skip.
 */
function classifyExpectation(rel, mode) {
  if (isSpecRootFile(rel)) return "root-skip";
  const base = basename(rel).toLowerCase();
  if (SACRED_BASENAMES.has(base)) return "sacred-skip";
  if (mode === "overview-only") {
    if (!isTopLevelOverview(rel)) return "out-of-scope";
  }
  return "expected";
}

function hasVerificationBlock(content) {
  return content.includes(`\n${VERIFICATION_HEADING}`);
}

function buildReport(rootAbs, mode) {
  const files = walk(rootAbs);
  const buckets = {
    updated: [],     // expected AND has block
    missing: [],     // expected AND no block
    sacredSkip: [],  // readme / changelog / 97 / 99
    rootSkip: [],    // spec/<root-file>.md
    outOfScope: [],  // not expected under current mode but exists
    stray: [],       // out-of-scope yet still carries a block (informational)
    errors: [],
  };
  for (const full of files) {
    const rel = relative(rootAbs, full);
    let content;
    try { content = readFileSync(full, "utf8"); }
    catch (e) { buckets.errors.push({ file: rel, error: e.message }); continue; }
    const klass = classifyExpectation(rel, mode);
    const present = hasVerificationBlock(content);
    if (klass === "expected" && present) buckets.updated.push(rel);
    else if (klass === "expected" && !present) buckets.missing.push(rel);
    else if (klass === "sacred-skip") buckets.sacredSkip.push(rel);
    else if (klass === "root-skip") buckets.rootSkip.push(rel);
    else if (klass === "out-of-scope") {
      buckets.outOfScope.push(rel);
      if (present) buckets.stray.push(rel);
    }
  }
  for (const k of Object.keys(buckets)) {
    if (Array.isArray(buckets[k])) buckets[k].sort();
  }
  return buckets;
}

function pct(num, denom) {
  if (!denom) return "0.0";
  return ((num / denom) * 100).toFixed(1);
}

function buildSummary(buckets, mode) {
  const expectedTotal = buckets.updated.length + buckets.missing.length;
  return {
    mode,
    generatedAt: new Date().toISOString(),
    totals: {
      scanned: buckets.updated.length + buckets.missing.length
        + buckets.sacredSkip.length + buckets.rootSkip.length + buckets.outOfScope.length,
      expected: expectedTotal,
      updated: buckets.updated.length,
      missing: buckets.missing.length,
      sacredSkip: buckets.sacredSkip.length,
      rootSkip: buckets.rootSkip.length,
      outOfScope: buckets.outOfScope.length,
      stray: buckets.stray.length,
      errors: buckets.errors.length,
    },
    coveragePct: Number(pct(buckets.updated.length, expectedTotal)),
  };
}

function renderMarkdown(buckets, summary) {
  const lines = [];
  lines.push("# Spec Verification Coverage Report");
  lines.push("");
  lines.push(`- **Mode:** \`${summary.mode}\``);
  lines.push(`- **Generated:** ${summary.generatedAt}`);
  lines.push(`- **Coverage:** ${summary.totals.updated} / ${summary.totals.expected} expected files (${summary.coveragePct}%)`);
  lines.push("");
  lines.push("## Totals");
  lines.push("");
  lines.push("| Bucket | Count |");
  lines.push("|---|---:|");
  lines.push(`| Scanned | ${summary.totals.scanned} |`);
  lines.push(`| Expected | ${summary.totals.expected} |`);
  lines.push(`| Updated (has \`## Verification\`) | ${summary.totals.updated} |`);
  lines.push(`| **Missing (still need \`## Verification\`)** | **${summary.totals.missing}** |`);
  lines.push(`| Skipped — sacred (readme/changelog/97/99) | ${summary.totals.sacredSkip} |`);
  lines.push(`| Skipped — spec root files | ${summary.totals.rootSkip} |`);
  lines.push(`| Out of scope under current mode | ${summary.totals.outOfScope} |`);
  lines.push(`| Stray blocks (out-of-scope yet present) | ${summary.totals.stray} |`);
  lines.push(`| IO errors | ${summary.totals.errors} |`);
  lines.push("");
  if (buckets.missing.length) {
    lines.push("## Missing — files still needing `## Verification`");
    lines.push("");
    for (const f of buckets.missing) lines.push(`- \`spec/${f}\``);
    lines.push("");
  }
  if (buckets.stray.length) {
    lines.push("## Stray blocks");
    lines.push("");
    lines.push("These files carry a `## Verification` block but are out of scope under the current mode. Run with `--mode all-files` or strip them.");
    lines.push("");
    for (const f of buckets.stray) lines.push(`- \`spec/${f}\``);
    lines.push("");
  }
  if (buckets.updated.length) {
    lines.push("## Updated");
    lines.push("");
    for (const f of buckets.updated) lines.push(`- \`spec/${f}\``);
    lines.push("");
  }
  if (buckets.errors.length) {
    lines.push("## Errors");
    lines.push("");
    for (const e of buckets.errors) lines.push(`- \`spec/${e.file}\` — ${e.error}`);
    lines.push("");
  }
  return lines.join("\n");
}

function ensureParent(filePath) {
  mkdirSync(dirname(filePath), { recursive: true });
}

function main() {
  const args = parseArgs(process.argv);
  const rootAbs = join(REPO_ROOT, args.root);
  let buckets;
  try { buckets = buildReport(rootAbs, args.mode); }
  catch (e) {
    console.error(`::error::cannot scan root: ${rootAbs} (${e.message})`);
    process.exit(1);
  }
  const summary = buildSummary(buckets, args.mode);
  const payload = args.json
    ? JSON.stringify({ summary, buckets }, null, 2)
    : renderMarkdown(buckets, summary);

  if (args.out) {
    const outAbs = join(REPO_ROOT, args.out);
    try { ensureParent(outAbs); writeFileSync(outAbs, payload + "\n", "utf8"); }
    catch (e) {
      console.error(`::error::cannot write report: ${outAbs} (${e.message})`);
      process.exit(1);
    }
    console.log(`Wrote ${args.out} (mode=${args.mode}, missing=${summary.totals.missing})`);
  } else {
    console.log(payload);
  }

  if (args.strict && summary.totals.missing > 0) {
    console.error(`::error::strict mode: ${summary.totals.missing} expected file(s) missing '## Verification'.`);
    for (const f of buckets.missing) console.error(`  - spec/${f}`);
    process.exit(3);
  }
  process.exit(0);
}

main();