#!/usr/bin/env node
// ============================================================
// sync-check.mjs — drift detector for generated files
// ============================================================
//
// Verifies that every file produced by `npm run sync` matches what
// `npm run sync` would currently regenerate. When drift is found,
// prints the fix command, lists the drifted files, and exits non-zero.
//
// USAGE
//   node scripts/sync-check.mjs                # check, exit non-zero on drift
//   node scripts/sync-check.mjs --verbose      # print unified diff per file
//   node scripts/sync-check.mjs --fix          # run `npm run sync` and keep result
//   node scripts/sync-check.mjs --report <dir> # write drift.json + drift.md
//
// EXIT CODES: 0 clean, 1 drift, 2 sync-pipeline failure.
// ============================================================

import { readFileSync, writeFileSync, existsSync, mkdirSync, copyFileSync, mkdtempSync } from "node:fs";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";
import { spawnSync } from "node:child_process";
import { tmpdir } from "node:os";
import { buildTrackedList } from "./lib/sync-check-normalisers.mjs";
import { writeReport, diffSummary } from "./lib/sync-check-report.mjs";

const __dirname = dirname(fileURLToPath(import.meta.url));
const ROOT = resolve(__dirname, "..");
const TRACKED = buildTrackedList();

const argv = process.argv.slice(2);
const VERBOSE = argv.includes("--verbose") || argv.includes("-v");
const FIX = argv.includes("--fix");
const REPORT_DIR = (() => {
  const idx = argv.indexOf("--report");
  if (idx === -1 || !argv[idx + 1]) return null;
  return resolve(ROOT, argv[idx + 1]);
})();

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
  return TRACKED.map((file) => ({ file, snap: snapshot(file) }));
}

function runSyncPipeline() {
  const result = spawnSync("npm", ["run", "sync"], {
    cwd: ROOT,
    stdio: VERBOSE ? "inherit" : "pipe", shell: process.platform === "win32",
    encoding: "utf8",
  });
  if (result.status !== 0) {
    process.stderr.write(result.stdout || "");
    process.stderr.write(result.stderr || "");
    return false;
  }
  return true;
}

function backupTracked() {
  const tmpDir = mkdtempSync(resolve(tmpdir(), "sync-check-backup-"));
  const backups = [];
  for (const file of TRACKED) {
    const abs = resolve(ROOT, file.path);
    if (!existsSync(abs)) continue;
    const bk = resolve(tmpDir, file.path.replace(/[\\/]/g, "__"));
    mkdirSync(dirname(bk), { recursive: true });
    copyFileSync(abs, bk);
    backups.push({ abs, bk });
  }
  return backups;
}

function computeDrift(before, after) {
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
  return drifted;
}

function detectDrift() {
  const before = snapshotAll();
  const backups = backupTracked();
  if (!runSyncPipeline()) {
    for (const { abs, bk } of backups) copyFileSync(bk, abs);
    return { genError: true };
  }
  const after = snapshotAll();
  const drifted = computeDrift(before, after);
  if (!FIX) {
    for (const { abs, bk } of backups) copyFileSync(bk, abs);
  }
  return { genError: false, drifted };
}

function printVerboseDiff(file, beforeText, afterText) {
  const tmp = mkdtempSync(resolve(tmpdir(), "sync-check-"));
  const a = resolve(tmp, "before");
  const b = resolve(tmp, "after");
  writeFileSync(a, beforeText);
  writeFileSync(b, afterText);
  const result = spawnSync("diff", ["-u", a, b], { encoding: "utf8" });
  process.stdout.write(`\n--- diff for ${file.path} ---\n`);
  process.stdout.write(result.stdout || "(diff binary not available)\n");
}

function reportOptions(mode, drifted) {
  return { reportDir: REPORT_DIR, root: ROOT, drifted, mode, tracked: TRACKED.length };
}

function handleClean() {
  writeReport(reportOptions(FIX ? "fix" : "check", []));
  process.stdout.write(`OK All ${TRACKED.length} sync-managed file(s) are up to date.\n`);
  process.exit(0);
}

function handleFix(drifted) {
  writeReport(reportOptions("fix", drifted));
  process.stdout.write(`\nFixed ${drifted.length} of ${TRACKED.length} sync-managed file(s):\n\n`);
  for (const item of drifted) process.stdout.write(`  • ${item.file.path}  (regenerated)\n`);
  process.stdout.write("\nNext: review with `git diff`, then `git add` + commit.\n");
  process.exit(0);
}

function driftTag(item) {
  if (!item.existedBefore) return "new file would be created";
  if (!item.existsAfter) return "file would be removed";
  const s = diffSummary(item.beforeText, item.afterText);
  return `+${s.added} / -${s.removed} line(s)`;
}

function printDriftList(drifted) {
  for (const item of drifted) {
    process.stdout.write(`  • ${item.file.path}  (${driftTag(item)})\n`);
    if (item.file.note) process.stdout.write(`      note: ${item.file.note}\n`);
    if (VERBOSE && item.existedBefore && item.existsAfter) {
      printVerboseDiff(item.file, item.beforeText, item.afterText);
    }
  }
}

function printFixInstructions(drifted) {
  const paths = drifted.map((d) => d.file.path).join(" ");
  process.stdout.write([
    "", "How to fix:",
    "  1. Run:    npm run sync",
    `  2. Stage: git add ${paths}`,
    "  3. Commit and push.",
    "", "Or, in one shot from this checkout:",
    "  node scripts/sync-check.mjs --fix && git add -A && git commit -m 'chore: sync generated files'",
    "",
  ].join("\n"));
  if (process.env.GITHUB_ACTIONS === "true") {
    process.stderr.write(
      `::error title=Sync drift detected::Run \`npm run sync\` locally and commit the regenerated files. ` +
        `Drifted: ${drifted.map((d) => d.file.path).join(", ")}\n`,
    );
  }
}

function main() {
  const result = detectDrift();
  if (result.genError) {
    process.stderr.write("\n::error::Sync pipeline failed — `npm run sync` did not exit cleanly.\n");
    process.exit(2);
  }
  if (result.drifted.length === 0) handleClean();
  if (FIX) handleFix(result.drifted);
  writeReport(reportOptions("check", result.drifted));
  process.stdout.write(
    `\nDrift detected in ${result.drifted.length} of ${TRACKED.length} sync-managed file(s):\n\n`,
  );
  printDriftList(result.drifted);
  printFixInstructions(result.drifted);
  process.exit(1);
}

main();
