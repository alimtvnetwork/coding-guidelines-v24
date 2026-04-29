#!/usr/bin/env node
// ============================================================
// sync-check.mjs — drift detector for generated files
// ============================================================
//
// Verifies that every file produced by the `npm run sync` pipeline
// matches what `npm run sync` would currently regenerate. When drift
// is found, prints the exact command operators must run, lists the
// drifted files, and exits non-zero (suitable for CI).
//
// Pipeline (kept in lock-step with package.json `scripts.sync`):
//
//   1. node scripts/sync-version.mjs         → version.json, package.json
//   2. node scripts/sync-spec-tree.mjs       → src/data/specTree.json
//   3. node scripts/sync-health-score.mjs    → public/health-score.json
//   4. node scripts/sync-readme-stats.mjs    → readme.md, docs/architecture.md,
//                                              docs/principles.md, docs/author.md
//
// Volatile fields (e.g. version.json `git`, `updated`; readme `UPDATED`
// stamp) are normalised before comparison so a fresh CI checkout on a
// different SHA / day does not produce false drift.
//
// USAGE
//   node scripts/sync-check.mjs                # check, exit non-zero on drift
//   node scripts/sync-check.mjs --verbose      # print diff for each drifted file
//   node scripts/sync-check.mjs --fix          # run `npm run sync` and re-check
//
// EXIT CODES
//   0  no drift
//   1  drift detected (or, with --fix, drift remained after sync ran)
//   2  generation error (a sync script failed)
// ============================================================

import { readFileSync, writeFileSync, existsSync, mkdirSync, copyFileSync } from "node:fs";
import { resolve, dirname, relative } from "node:path";
import { fileURLToPath } from "node:url";
import { spawnSync } from "node:child_process";
import { tmpdir } from "node:os";
import { mkdtempSync } from "node:fs";

const __dirname = dirname(fileURLToPath(import.meta.url));
const ROOT = resolve(__dirname, "..");

// ----- Argument parsing ----------------------------------------------------

const argv = process.argv.slice(2);
const VERBOSE = argv.includes("--verbose") || argv.includes("-v");
const FIX = argv.includes("--fix");

// ----- Files watched for drift --------------------------------------------
//
// Each entry declares the on-disk path (relative to repo root) and an
// optional `normalise(content)` callback that strips volatile fields
// before comparison. Keep this list aligned with what `npm run sync`
// actually writes — adding a new sync output without listing it here
// means CI will silently miss the drift.

const TRACKED = [
  {
    path: "version.json",
    normalise: stripVolatileVersionFields,
    note: "structural fields only — commit/date provenance ignored (legitimately differs per checkout / day).",
  },
  {
    path: "src/data/specTree.json",
  },
  {
    path: "public/health-score.json",
    normalise: stripVolatileHealthScoreFields,
    note: "the `generated` ISO timestamp is ignored.",
  },
  {
    path: "readme.md",
    normalise: stripVolatileReadmeStamps,
    note: "the `<!-- UPDATED -->` date stamp is ignored.",
  },
  {
    path: "docs/architecture.md",
    normalise: stripVolatileReadmeStamps,
    note: "the `<!-- UPDATED -->` date stamp is ignored.",
  },
  {
    path: "docs/principles.md",
    normalise: stripVolatileReadmeStamps,
    note: "the `<!-- UPDATED -->` date stamp is ignored.",
  },
  {
    path: "docs/author.md",
    normalise: stripVolatileReadmeStamps,
    note: "the `<!-- UPDATED -->` date stamp is ignored.",
  },
];

// ----- Normalisers ---------------------------------------------------------

function stripVolatileVersionFields(content) {
  // version.json is JSON; preserve formatting by re-serialising after
  // dropping the two known-volatile fields.
  let parsed;
  try {
    parsed = JSON.parse(content);
  } catch {
    return content;
  }
  delete parsed.LastCommitSha;
  delete parsed.git;
  delete parsed.updated;
  return JSON.stringify(parsed, null, 2) + "\n";
}

function stripVolatileHealthScoreFields(content) {
  // sync-health-score writes a fresh `generated` ISO timestamp on every
  // run. Drop it before comparing so a no-op regeneration doesn't look
  // like drift.
  let parsed;
  try {
    parsed = JSON.parse(content);
  } catch {
    return content;
  }
  delete parsed.generated;
  return JSON.stringify(parsed, null, 2) + "\n";
}

function stripVolatileReadmeStamps(content) {
  // sync-readme-stats stamps the current date between
  // ``<!-- UPDATED:start -->`` and ``<!-- UPDATED:end -->`` markers.
  // The exact date drifts every day; replace the body with a constant
  // sentinel before diffing.
  return content.replace(
    /<!--\s*UPDATED:start\s*-->[\s\S]*?<!--\s*UPDATED:end\s*-->/g,
    "<!-- UPDATED:start -->__NORMALISED__<!-- UPDATED:end -->",
  );
}

// ----- Helpers -------------------------------------------------------------

function readIfExists(absPath) {
  if (!existsSync(absPath)) return null;
  return readFileSync(absPath, "utf8");
}

function snapshot(file) {
  const abs = resolve(ROOT, file.path);
  const raw = readIfExists(abs);
  if (raw === null) return { exists: false, content: null };
  const content = file.normalise ? file.normalise(raw) : raw;
  return { exists: true, content };
}

function snapshotAll() {
  return TRACKED.map((file) => ({
    file,
    snap: snapshot(file),
  }));
}

function runSyncPipeline() {
  const result = spawnSync("npm", ["run", "sync"], {
    cwd: ROOT,
    stdio: VERBOSE ? "inherit" : "pipe",
    encoding: "utf8",
  });
  if (result.status !== 0) {
    process.stderr.write(result.stdout || "");
    process.stderr.write(result.stderr || "");
    return false;
  }
  return true;
}

function diffSummary(beforeText, afterText) {
  const before = beforeText.split("\n");
  const after = afterText.split("\n");
  let added = 0;
  let removed = 0;
  const max = Math.max(before.length, after.length);
  for (let i = 0; i < max; i++) {
    if (before[i] === undefined) added++;
    else if (after[i] === undefined) removed++;
    else if (before[i] !== after[i]) {
      added++;
      removed++;
    }
  }
  return { added, removed, beforeLines: before.length, afterLines: after.length };
}

function printVerboseDiff(file, beforeText, afterText) {
  // We avoid pulling in a diff library — write the two snapshots to
  // /tmp and shell out to plain `diff -u` for a familiar unified diff.
  const tmp = mkdtempSync(resolve(tmpdir(), "sync-check-"));
  const a = resolve(tmp, "before");
  const b = resolve(tmp, "after");
  writeFileSync(a, beforeText);
  writeFileSync(b, afterText);
  const result = spawnSync("diff", ["-u", a, b], { encoding: "utf8" });
  process.stdout.write(`\n--- diff for ${file.path} ---\n`);
  process.stdout.write(result.stdout || "(diff binary not available)\n");
}

// ----- Main ----------------------------------------------------------------

function detectDrift() {
  // Snapshot the committed (pre-sync) state with normalisation applied.
  const before = snapshotAll();

  // Backup originals so we can restore them after running sync.
  const tmpDir = mkdtempSync(resolve(tmpdir(), "sync-check-backup-"));
  const backups = [];
  for (const { file } of before) {
    const abs = resolve(ROOT, file.path);
    if (existsSync(abs)) {
      const bk = resolve(tmpDir, file.path.replace(/[\\/]/g, "__"));
      mkdirSync(dirname(bk), { recursive: true });
      copyFileSync(abs, bk);
      backups.push({ abs, bk });
    }
  }

  // Regenerate.
  const ok = runSyncPipeline();
  if (!ok) {
    // Restore originals so a failed run doesn't leave the working tree dirty.
    for (const { abs, bk } of backups) copyFileSync(bk, abs);
    return { genError: true };
  }

  const after = snapshotAll();

  // Compare.
  const drifted = [];
  for (let i = 0; i < TRACKED.length; i++) {
    const { file } = before[i];
    const a = before[i].snap;
    const b = after[i].snap;
    if (!a.exists && !b.exists) continue;
    if (!a.exists || !b.exists || a.content !== b.content) {
      drifted.push({
        file,
        beforeText: a.content ?? "",
        afterText: b.content ?? "",
        existedBefore: a.exists,
        existsAfter: b.exists,
      });
    }
  }

  // Restore committed state — sync-check is read-only by default.
  // (--fix mode skips the restore so the regenerated files remain.)
  if (!FIX) {
    for (const { abs, bk } of backups) copyFileSync(bk, abs);
  }

  return { genError: false, drifted };
}

function main() {
  const result = detectDrift();
  if (result.genError) {
    process.stderr.write(
      "\n::error::Sync pipeline failed — `npm run sync` did not exit cleanly.\n",
    );
    process.exit(2);
  }

  if (result.drifted.length === 0) {
    process.stdout.write(
      `OK All ${TRACKED.length} sync-managed file(s) are up to date.\n`,
    );
    process.exit(0);
  }

  // --fix mode: the regenerated files are already on disk (we skipped
  // the restore step), so the drift IS the fix. Report what changed
  // and exit 0 so callers can chain `git add` / commit.
  if (FIX) {
    process.stdout.write(
      `\nFixed ${result.drifted.length} of ${TRACKED.length} sync-managed file(s):\n\n`,
    );
    for (const item of result.drifted) {
      process.stdout.write(`  • ${item.file.path}  (regenerated)\n`);
    }
    process.stdout.write(
      "\nNext: review with `git diff`, then `git add` + commit.\n",
    );
    process.exit(0);
  }

  process.stdout.write(
    `\nDrift detected in ${result.drifted.length} of ${TRACKED.length} sync-managed file(s):\n\n`,
  );
  for (const item of result.drifted) {
    const summary = item.existedBefore && item.existsAfter
      ? diffSummary(item.beforeText, item.afterText)
      : null;
    const tag = !item.existedBefore
      ? "new file would be created"
      : !item.existsAfter
        ? "file would be removed"
        : `+${summary.added} / -${summary.removed} line(s)`;
    process.stdout.write(`  • ${item.file.path}  (${tag})\n`);
    if (item.file.note) {
      process.stdout.write(`      note: ${item.file.note}\n`);
    }
    if (VERBOSE && item.existedBefore && item.existsAfter) {
      printVerboseDiff(item.file, item.beforeText, item.afterText);
    }
  }

  const githubAnnotation = process.env.GITHUB_ACTIONS === "true";
  const fixCmd = "npm run sync";
  const lines = [
    "",
    "How to fix:",
    `  1. Run:    ${fixCmd}`,
    "  2. Stage: git add " + result.drifted.map((d) => d.file.path).join(" "),
    "  3. Commit and push.",
    "",
    "Or, in one shot from this checkout:",
    "  node scripts/sync-check.mjs --fix && git add -A && git commit -m 'chore: sync generated files'",
    "",
  ];
  process.stdout.write(lines.join("\n"));

  if (githubAnnotation) {
    process.stderr.write(
      `::error title=Sync drift detected::Run \`${fixCmd}\` locally and commit the regenerated files. ` +
        `Drifted: ${result.drifted.map((d) => d.file.path).join(", ")}\n`,
    );
  }

  process.exit(1);
}

main();
