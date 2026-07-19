#!/usr/bin/env node
/**
 * scripts/tests/print-required-checks.test.mjs
 *
 * Self-test for scripts/print-required-checks.mjs. Guards the two
 * known parsing traps in parseWorkflowText():
 *   1. Must locate `jobs:` even when preceded by other top-level keys
 *      (name, on, permissions, concurrency, env, defaults).
 *   2. Must terminate at the next column-0 top-level key so trailing
 *      top-level sections (env, concurrency at file end) do not leak
 *      into the last job.
 *   3. Must prefer explicit `name:` over the job id for the check
 *      context name (that string is what GitHub branch-protection
 *      matches against, not the job id).
 *   4. checkStale() must flag expected required-context names that
 *      no workflow job produces.
 *
 * Runs standalone (no test runner). Exit code = number of failures.
 * Wired as a step in scripts/lint-ci.sh and in the sync-drift CI job.
 */
import { parseWorkflowText, checkStale } from "../print-required-checks.mjs";

let failed = 0;
function assert(name, cond, detail) {
  if (cond) {
    console.log(`  ok  ${name}`);
    return;
  }
  failed += 1;
  console.log(`  FAIL ${name}`);
  if (detail !== undefined) console.log(`       ${JSON.stringify(detail)}`);
}

// --- Fixture 1: minimal workflow, id-only job names ------------------
const minimal = [
  "name: CI",
  "on: [push]",
  "",
  "jobs:",
  "  lint:",
  "    runs-on: ubuntu-latest",
  "    steps:",
  "      - run: echo hi",
  "  test:",
  "    runs-on: ubuntu-latest",
  "    steps:",
  "      - run: echo hi",
  "",
].join("\n");

const p1 = parseWorkflowText(minimal, "fallback");
assert("minimal: workflow name from top-level", p1.workflowName === "CI", p1);
assert("minimal: 2 jobs parsed", p1.jobs.length === 2, p1.jobs);
assert(
  "minimal: job id used as name when no explicit name",
  p1.jobs[0].name === "lint" && p1.jobs[1].name === "test",
  p1.jobs,
);

// --- Fixture 2: explicit job names (branch-protection matches these) -
const withNames = [
  "name: Visual Regression",
  "on: push",
  "permissions:",
  "  contents: read",
  "",
  "jobs:",
  "  visual:",
  "    name: Slides Visual Regression",
  "    runs-on: ubuntu-latest",
  "    steps:",
  "      - run: echo hi",
  "  smoke:",
  "    name: Slides Smoke",
  "    runs-on: ubuntu-latest",
  "    steps:",
  "      - run: echo hi",
  "",
  "concurrency:",
  "  group: visual",
].join("\n");

const p2 = parseWorkflowText(withNames, "fallback");
assert("withNames: 2 jobs parsed", p2.jobs.length === 2, p2.jobs);
assert(
  "withNames: explicit name preferred over id",
  p2.jobs[0].name === "Slides Visual Regression" &&
    p2.jobs[1].name === "Slides Smoke",
  p2.jobs,
);
assert(
  "withNames: trailing top-level `concurrency:` does not leak as job",
  p2.jobs.every((j) => j.id !== "group"),
  p2.jobs,
);

// --- Fixture 3: no jobs section (malformed workflow) -----------------
const noJobs = ["name: Empty", "on: push", ""].join("\n");
const p3 = parseWorkflowText(noJobs, "fallback");
assert("noJobs: empty jobs array", p3.jobs.length === 0, p3);
assert("noJobs: fallback name", p3.workflowName === "Empty", p3);

// --- Fixture 4: jobs: not at column 0 (indented, unusual) ------------
// parseWorkflowText requires the `\njobs:\n` sentinel. Anything else
// is treated as "no jobs found", which is the safe fallback.
const indentedJobs = ["name: Odd", "  jobs:", "    lint:", ""].join("\n");
const p4 = parseWorkflowText(indentedJobs, "fallback");
assert("indentedJobs: falls back to zero jobs (safe)", p4.jobs.length === 0, p4);

// --- Fixture 5: checkStale() ----------------------------------------
const workflows = [
  { workflowName: "CI", file: "ci.yml", jobs: [{ id: "lint", name: "Lint" }] },
  {
    workflowName: "Visual",
    file: "visual.yml",
    jobs: [{ id: "visual", name: "Slides Visual Regression" }],
  },
];
const stale1 = checkStale(workflows, { required: ["Lint", "Slides Visual Regression"] });
assert("checkStale: zero stale when every required matches a job", stale1.length === 0, stale1);

const stale2 = checkStale(workflows, {
  required: ["Lint", "Slides Visual Regression", "Deleted Job"],
});
assert(
  "checkStale: flags required names with no matching job",
  stale2.length === 1 && stale2[0] === "Deleted Job",
  stale2,
);

const stale3 = checkStale(workflows, { required: [] });
assert("checkStale: empty required list is not stale", stale3.length === 0, stale3);

// --- Fixture 6: real .github/workflows/*.yml parse without crashing --
import { readdirSync, readFileSync } from "node:fs";
import { join } from "node:path";
const realDir = ".github/workflows";
try {
  const files = readdirSync(realDir).filter((f) => f.endsWith(".yml") || f.endsWith(".yaml"));
  let totalJobs = 0;
  for (const f of files) {
    const parsed = parseWorkflowText(readFileSync(join(realDir, f), "utf8"), f);
    totalJobs += parsed.jobs.length;
  }
  assert(
    "real workflows: at least one job discovered across all files",
    files.length > 0 && totalJobs > 0,
    { files: files.length, totalJobs },
  );
} catch (err) {
  assert("real workflows: parse without throwing", false, { error: String(err) });
}

if (failed > 0) {
  console.log(`\n${failed} test(s) failed.`);
  process.exit(failed);
}
console.log(`\nAll checks passed.`);
