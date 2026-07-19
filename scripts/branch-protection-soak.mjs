#!/usr/bin/env node
/**
 * scripts/branch-protection-soak.mjs
 *
 * Read-only soak report for every entry in
 * `.github/branch-protection.expected.json` under
 * `desired-but-not-yet-required[]`. For each workflow, queries the last
 * N runs on `main` via `gh run list` and prints pass rate + verdict.
 *
 * Root cause it closes: promotion of `visual` + `smoke` to REQUIRED has
 * been deferred every release with the same reason: "wait for one full
 * release cycle without spurious diffs". Nobody has a cheap way to
 * check "has it actually soaked?" today, so the deferral is perpetual.
 *
 * Usage:
 *   npm run branch-protection:soak                        # last 20 runs, main
 *   npm run branch-protection:soak -- --runs 50           # override count
 *   npm run branch-protection:soak -- --repo owner/name   # override repo
 *
 * Verdict thresholds (mediocre-AI-safe, tune only with reason):
 *   pass rate >= 0.95 over the window => READY
 *   pass rate >= 0.80                 => WATCH
 *   otherwise                         => NOT READY
 *
 * Exit codes:
 *   0 = every desired entry is READY
 *   1 = one or more entries WATCH or NOT READY (still prints full report)
 *   2 = gh CLI missing / not authenticated
 *   3 = repo auto-detect failed and --repo not provided
 *
 * Never mutates anything.
 */
import { readFileSync, existsSync } from "node:fs";
import { spawnSync } from "node:child_process";

const EXPECTED_FILE = ".github/branch-protection.expected.json";

function parseArgs(argv) {
  const args = { runs: 20, repo: null };
  for (let i = 0; i < argv.length; i += 1) {
    if (argv[i] === "--runs" && argv[i + 1]) {
      args.runs = Number(argv[i + 1]);
      i += 1;
    } else if (argv[i] === "--repo" && argv[i + 1]) {
      args.repo = argv[i + 1];
      i += 1;
    }
  }
  if (!Number.isFinite(args.runs) || args.runs < 1) args.runs = 20;
  return args;
}

function ghAvailable() {
  return spawnSync("gh", ["auth", "status"], { encoding: "utf8" }).status === 0;
}

function detectRepo() {
  const r = spawnSync("gh", ["repo", "view", "--json", "nameWithOwner", "-q", ".nameWithOwner"], {
    encoding: "utf8",
  });
  if (r.status !== 0) return null;
  return r.stdout.trim() || null;
}

function loadExpected() {
  if (!existsSync(EXPECTED_FILE)) {
    console.error(`Missing ${EXPECTED_FILE}`);
    process.exit(1);
  }
  return JSON.parse(readFileSync(EXPECTED_FILE, "utf8"));
}

function fetchRuns(repo, workflowFile, runs) {
  const r = spawnSync(
    "gh",
    [
      "run",
      "list",
      "--repo",
      repo,
      "--workflow",
      workflowFile,
      "--branch",
      "main",
      "--limit",
      String(runs),
      "--json",
      "conclusion,status,createdAt",
    ],
    { encoding: "utf8" },
  );
  if (r.status !== 0) {
    console.error(`gh run list failed for ${workflowFile}:`);
    console.error(r.stderr);
    return null;
  }
  return JSON.parse(r.stdout);
}

export function verdict(passRate, sampleSize) {
  if (sampleSize === 0) return "NO DATA";
  if (passRate >= 0.95) return "READY";
  if (passRate >= 0.8) return "WATCH";
  return "NOT READY";
}

function reportOne(repo, entry, runs) {
  const wfFile = entry.workflow.replace(/^\.github\/workflows\//, "");
  const list = fetchRuns(repo, wfFile, runs);
  if (list === null) return { entry, verdict: "ERROR" };
  const completed = list.filter((r) => r.status === "completed");
  const passed = completed.filter((r) => r.conclusion === "success").length;
  const rate = completed.length ? passed / completed.length : 0;
  const v = verdict(rate, completed.length);
  console.log("");
  console.log(`  [${v}] ${entry.check}  (${entry.workflow})`);
  console.log(`    window: ${completed.length}/${list.length} completed of last ${runs}`);
  console.log(`    pass:   ${passed}/${completed.length}  rate=${(rate * 100).toFixed(1)}%`);
  console.log(`    reason: ${entry.reason}`);
  return { entry, verdict: v };
}

function main() {
  if (!ghAvailable()) {
    console.error("gh CLI not found or not authenticated. Run `gh auth login`.");
    process.exit(2);
  }
  const args = parseArgs(process.argv.slice(2));
  const repo = args.repo || detectRepo();
  if (!repo) {
    console.error("Repo not detected. Pass --repo owner/name.");
    process.exit(3);
  }
  const expected = loadExpected();
  const desired = expected["desired-but-not-yet-required"] || [];
  console.log(`Repo: ${repo}`);
  console.log(`Window: last ${args.runs} runs on main`);
  console.log(`Desired-but-not-yet-required entries: ${desired.length}`);
  if (!desired.length) {
    console.log("(no soak candidates; every desired check is already REQUIRED)");
    return;
  }
  const results = desired.map((e) => reportOne(repo, e, args.runs));
  const notReady = results.filter((r) => r.verdict !== "READY");
  console.log("");
  if (notReady.length === 0) {
    console.log("VERDICT: every desired check is READY for promotion.");
    console.log("Next: run `.lovable/procedures/branch-protection.md` to flip them to REQUIRED.");
    return;
  }
  console.log(`VERDICT: ${notReady.length}/${results.length} entr${notReady.length === 1 ? "y" : "ies"} not READY.`);
  process.exit(1);
}

import { fileURLToPath } from "node:url";
if (process.argv[1] === fileURLToPath(import.meta.url)) main();
