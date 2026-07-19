#!/usr/bin/env node
/**
 * scripts/check-lint-ci-drift.mjs
 *
 * Drift check: every runnable lint step in `.github/workflows/ci.yml`
 * (the `lint` job) MUST be mirrored by an entry in the STEPS registry
 * inside `scripts/lint-ci.sh`. Without this guard, someone can add a
 * new CI step and local `bun run lint:ci` silently falls behind, so
 * pre-push stops matching CI. This is the follow-up called out at
 * scripts/lint-ci.sh:9-11.
 *
 * Matching rule: a lint-ci STEPS label is a "match" for a ci.yml step
 * name if the ci.yml name starts with the lint-ci label (case-sensitive).
 * That absorbs suffixes like "(diff mode)" that ci.yml appends when a
 * step is parameterised for the workflow.
 *
 * Steps that are infrastructure (Checkout / Setup / Cache / Restore /
 * Save / Resolve / Upload / Skipped / Detect / on-failure summary /
 * installer-only checks that don't belong in the local placeholder
 * pipeline) are listed in IGNORED_CI_STEPS with a one-line reason.
 */
import { readFileSync } from "node:fs";

const CI_YAML = ".github/workflows/ci.yml";
const LINT_CI = "scripts/lint-ci.sh";

// CI step names that intentionally don't need a mirror in lint-ci.sh.
// Reason must be a real reason, not "we don't feel like running it".
const IGNORED_CI_STEPS = new Map([
  ["Checkout", "actions/checkout — env setup, not a lint"],
  ["Setup Go", "toolchain provisioning"],
  ["Cache Go build and module cache", "actions/cache — infra"],
  ["Setup Python", "toolchain provisioning"],
  ["Cache pip", "actions/cache — infra"],
  ["Setup Node", "toolchain provisioning"],
  ["Install dependencies (no scripts)", "bun install — infra"],
  ["Install Mermaid CLI", "toolchain provisioning"],
  ["Discover .mmd files", "infra shell for later diagram render"],
  ["Render Mermaid → PNG (all spec/**/{diagrams,images}/*.mmd)", "render side effect; drift is caught by the follow-up check step"],
  ["Verify no drift after render", "duplicated by 'Check Mermaid diagram drift' step which IS mirrored"],
  ["Upload rendered PNGs as artifact", "artifact upload — infra"],
  ["Forbidden-strings summary report (on failure)", "on-failure summary; lint-ci.sh runs it inline when the mirrored step fails"],
  ["Restore placeholder-linter PASS cache", "cache infra for the placeholder step"],
  ["Resolve placeholder-linter diff base", "computes DIFF_FLAG; consumed by the mirrored placeholder step"],
  ["Upload JSON report on failure", "artifact upload — infra"],
  ["Suggest auto-fixes for broken links (advisory)", "advisory-only; not a gating lint"],
  ["Upload suggestions JSON", "artifact upload — infra"],
  ["Detect codegen-relevant changes", "path-filter gate for a separate codegen job"],
  ["Verify codegen determinism", "codegen job, not a linter"],
  ["Regen-diff guard (no uncommitted codegen output)", "codegen job, not a linter"],
  ["Skipped — no codegen-relevant files changed", "path-filter skip marker"],
  ["Verify checker present", "presence-check for the cross-links job; the actual check IS mirrored"],
  ["Check sync drift (all generated files)", "runs in the sync-drift job; lint-ci.sh scope is placeholder-linters"],
  ["Guard release-install.* against /releases/latest calls", "installer-tests job, mirrored in scripts/installer-*.sh not lint-ci.sh"],
  ["release-install.sh exit-code acceptance (spec §F + §AC)", "installer-tests job"],
  ["release.sh bake-output integration (spec §Release-Time Build Step)", "installer-tests job"],
  ["Bundle installer conformance (14 files, spec §3/§7/§8)", "installer-tests job"],
  ["Root installer §27 conformance (install.sh / install.ps1 / linters-cicd)", "installer-tests job"],
  ["BOOL-NEG-001 unit tests (linters-cicd)", "installer-tests job"],
  ["BOOL-NEG-001 pipeline smoke test (run-all.sh end-to-end)", "installer-tests job"],
  ["Orchestrator flags smoke test (--strict / --total-timeout / --split-by)", "installer-tests job"],
  ["Run guard fixture suite (positive + negative)", "runner-guard job"],
  ["Run guard against real run.sh / run.ps1 (emit report)", "runner-guard job"],
  ["Publish report to GitHub Step Summary", "runner-guard reporting"],
  ["Upload runner-guard-report.md", "artifact upload — infra"],
  ["Check tunable constants (T1–T4 — single-source-of-truth, seed parity)", "runs in the tunable-constants job scope"],
]);

function ciStepNames() {
  const src = readFileSync(CI_YAML, "utf8");
  const names = [];
  const re = /^\s+-\s+name:\s+(.+?)\s*$/gm;
  let m;
  while ((m = re.exec(src)) !== null) names.push(m[1]);
  return names;
}

function lintCiLabels() {
  const src = readFileSync(LINT_CI, "utf8");
  const start = src.indexOf("STEPS=(");
  const end = src.indexOf(")", start);
  if (start < 0 || end < 0) throw new Error("STEPS=(...) block not found in lint-ci.sh");
  const block = src.slice(start, end);
  const labels = [];
  for (const line of block.split("\n")) {
    const t = line.trim();
    if (!t.startsWith('"') || !t.includes("|")) continue;
    const inner = t.slice(1, t.lastIndexOf('"'));
    labels.push(inner.split("|", 1)[0]);
  }
  return labels;
}

function main() {
  const ciNames = ciStepNames();
  const labels = lintCiLabels();
  const missing = [];
  for (const ciName of ciNames) {
    if (IGNORED_CI_STEPS.has(ciName)) continue;
    const matched = labels.some((label) => ciName.startsWith(label));
    if (!matched) missing.push(ciName);
  }
  const unusedIgnores = [...IGNORED_CI_STEPS.keys()].filter((k) => !ciNames.includes(k));

  console.log(`[check-lint-ci-drift] ci.yml steps: ${ciNames.length}`);
  console.log(`[check-lint-ci-drift] lint-ci.sh labels: ${labels.length}`);
  console.log(`[check-lint-ci-drift] ignored (with reason): ${IGNORED_CI_STEPS.size}`);
  if (unusedIgnores.length) {
    console.log("");
    console.log("STALE IGNORES (present in IGNORED_CI_STEPS but no longer in ci.yml):");
    for (const n of unusedIgnores) console.log(`  - ${n}`);
  }
  if (missing.length) {
    console.log("");
    console.log("DRIFT: ci.yml step(s) with no mirror in scripts/lint-ci.sh STEPS[]:");
    for (const n of missing) console.log(`  - ${n}`);
    console.log("");
    console.log("Fix: either add a matching STEPS entry in scripts/lint-ci.sh, or,");
    console.log("if the step is legitimately infra-only, add it to IGNORED_CI_STEPS");
    console.log("in scripts/check-lint-ci-drift.mjs with a one-line reason.");
    process.exit(1);
  }
  if (unusedIgnores.length) process.exit(1);
  console.log("OK — lint-ci.sh mirrors ci.yml (no drift).");
}

main();
