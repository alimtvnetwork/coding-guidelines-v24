#!/usr/bin/env node
/**
 * scripts/print-required-checks.mjs
 *
 * Enumerates workflow -> job pairs from .github/workflows/*.yml and
 * prints the exact `gh api` command to promote them to REQUIRED status
 * on branch protection for `main`. Also flags any workflow that is
 * already listed in `.github/branch-protection.expected.json` as
 * required so operators can see current-vs-desired without opening the
 * GitHub UI.
 *
 * Root cause it addresses: backlog item "promote slides-visual.yml to
 * required check" has been carried forward for two releases because
 * the exact `gh api` payload was never captured. Manual UI steps that
 * live only in someone's head do not get done.
 *
 * Usage:
 *   node scripts/print-required-checks.mjs           # human report
 *   node scripts/print-required-checks.mjs --json    # machine payload
 *   node scripts/print-required-checks.mjs --check   # exit non-zero
 *                                                    # if expected.json
 *                                                    # references a
 *                                                    # workflow that
 *                                                    # no longer exists
 *
 * The generated payload matches
 *   PUT /repos/{owner}/{repo}/branches/main/protection/required_status_checks
 * See https://docs.github.com/en/rest/branches/branch-protection.
 */
import { readFileSync, readdirSync, existsSync } from "node:fs";
import { join, basename } from "node:path";

const WORKFLOWS_DIR = ".github/workflows";
const EXPECTED_FILE = ".github/branch-protection.expected.json";

export function parseWorkflowText(src, fallbackName) {
  // Top-level `name:` (workflow display name shown in the Checks tab).
  const wfNameMatch = src.match(/^name:\s*(.+?)\s*$/m);
  const workflowName = wfNameMatch ? wfNameMatch[1] : fallbackName;
  // Job entries under `jobs:`. Job id = 2-space indented key; job's
  // `name:` (if present) is what GitHub shows as the check name.
  const jobs = [];
  const jobsIdx = src.indexOf("\njobs:\n");
  if (jobsIdx < 0) return { workflowName, jobs };
  const after = src.slice(jobsIdx + 7);
  const lines = after.split("\n");
  let currentJob = null;
  for (const line of lines) {
    // Stop scanning when a new top-level yaml key begins (column 0).
    // Check BEFORE the job-id match so a top-level key that happens to
    // match the id regex (none today, but future-proof) still terminates.
    if (/^[a-zA-Z]/.test(line) && !line.startsWith(" ")) {
      if (currentJob) jobs.push(currentJob);
      currentJob = null;
      break;
    }
    const jobIdMatch = line.match(/^ {2}([a-zA-Z0-9_-]+):\s*$/);
    if (jobIdMatch) {
      if (currentJob) jobs.push(currentJob);
      currentJob = { id: jobIdMatch[1], name: jobIdMatch[1] };
      continue;
    }
    if (!currentJob) continue;
    const jobNameMatch = line.match(/^ {4}name:\s*(.+?)\s*$/);
    if (jobNameMatch) currentJob.name = jobNameMatch[1];
  }
  if (currentJob) jobs.push(currentJob);
  return { workflowName, jobs };
}

function parseWorkflow(path) {
  const src = readFileSync(path, "utf8");
  const parsed = parseWorkflowText(src, basename(path, ".yml"));
  return { ...parsed, file: path };
}

function loadWorkflows() {
  return readdirSync(WORKFLOWS_DIR)
    .filter((f) => f.endsWith(".yml") || f.endsWith(".yaml"))
    .map((f) => parseWorkflow(join(WORKFLOWS_DIR, f)));
}

function loadExpected() {
  if (!existsSync(EXPECTED_FILE)) return { required: [] };
  return JSON.parse(readFileSync(EXPECTED_FILE, "utf8"));
}

function checkStale(workflows, expected) {
  const contexts = new Set();
  for (const wf of workflows) for (const j of wf.jobs) contexts.add(j.name);
  return (expected.required || []).filter((n) => !contexts.has(n));
}

function humanReport(workflows, expected) {
  console.log(`Workflows scanned: ${workflows.length}`);
  console.log(`Expected required checks: ${(expected.required || []).length}`);
  console.log("");
  for (const wf of workflows) {
    console.log(`  ${wf.workflowName}  [${wf.file}]`);
    for (const j of wf.jobs) {
      const required = (expected.required || []).includes(j.name);
      const mark = required ? "REQUIRED" : "advisory";
      console.log(`    - ${j.name}  (${mark})`);
    }
  }
  console.log("");
  console.log("To promote a job to REQUIRED (repo admin, gh CLI logged in):");
  console.log("");
  console.log('  gh api -X PATCH "repos/{owner}/{repo}/branches/main/protection/required_status_checks" \\');
  console.log('    -f strict=true \\');
  const ctxList = (expected.required || []).map((n) => `      "${n}"`).join(",\n");
  console.log("    -F contexts='[");
  console.log(ctxList || '      "<paste job name here>"');
  console.log("    ]'");
  console.log("");
  console.log("Then add the newly-required job name to");
  console.log(`  ${EXPECTED_FILE}  under "required": [...]`);
  console.log("so this report stays in sync with reality.");
}

function jsonPayload(expected) {
  return {
    endpoint: "PATCH /repos/{owner}/{repo}/branches/main/protection/required_status_checks",
    body: { strict: true, contexts: expected.required || [] },
  };
}

function main() {
  const args = new Set(process.argv.slice(2));
  const workflows = loadWorkflows();
  const expected = loadExpected();
  const stale = checkStale(workflows, expected);

  if (args.has("--check")) {
    if (stale.length) {
      console.error("STALE required-check names in " + EXPECTED_FILE + ":");
      for (const n of stale) console.error("  - " + n);
      console.error("Fix: remove them, or restore the workflow job that produced them.");
      process.exit(1);
    }
    console.log("OK, no stale required-check names.");
    return;
  }
  if (args.has("--json")) {
    console.log(JSON.stringify(jsonPayload(expected), null, 2));
    return;
  }
  humanReport(workflows, expected);
  if (stale.length) {
    console.log("");
    console.log("WARNING, stale entries in " + EXPECTED_FILE + ":");
    for (const n of stale) console.log("  - " + n);
    process.exit(1);
  }
}

main();
