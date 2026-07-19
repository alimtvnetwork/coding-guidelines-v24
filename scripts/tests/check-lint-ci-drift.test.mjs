#!/usr/bin/env node
/**
 * scripts/tests/check-lint-ci-drift.test.mjs
 *
 * Self-test for scripts/check-lint-ci-drift.mjs. Guards the two known
 * parsing traps:
 *   1. parseLintCiLabels() must stop at the closing `)` of STEPS=(...)
 *      and must ignore comment / blank lines inside the block.
 *   2. analyzeDrift() must (a) flag ci steps with no mirror,
 *      (b) accept parenthetical-qualifier drift ("(spec)" vs
 *      "(spec, advisory)") via stem match, (c) flag STALE IGNORES,
 *      (d) return zero-missing when every ci step is mirrored or
 *      explicitly ignored.
 *
 * Runs standalone (no test runner) so it works in every environment
 * where CI already runs `node`. Exit code = number of failed cases.
 * Wired as step 17 of scripts/lint-ci.sh.
 */
import {
  parseCiStepNames,
  parseLintCiLabels,
  analyzeDrift,
} from "../check-lint-ci-drift.mjs";

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

// --- Fixtures ---------------------------------------------------------
const ciYaml = [
  "jobs:",
  "  lint:",
  "    steps:",
  "      - name: Checkout",
  "      - name: Run alpha check",
  "      - name: Run beta check (spec)",
  "      - name: Run gamma check",
  "      - name: Ghost step that no longer exists elsewhere",
].join("\n");

const lintCi = [
  "#!/usr/bin/env bash",
  "STEPS=(",
  "  # inline comment must be skipped",
  '  "Run alpha check|node scripts/alpha.mjs"',
  "",
  '  "Run beta check (spec, advisory)|node scripts/beta.mjs"',
  ")",
  "echo done",
  '"Run zeta|should-not-parse-outside-block"',
].join("\n");

const ignored = new Map([
  ["Checkout", "infra"],
  ["Ghost step that no longer exists elsewhere", "kept for drift-check test"],
  ["Long-gone step", "STALE — no longer in ci.yml"],
]);

// --- Cases ------------------------------------------------------------
const ciNames = parseCiStepNames(ciYaml);
assert("parseCiStepNames extracts 4 names", ciNames.length === 4, ciNames);
assert("parseCiStepNames preserves order",
  ciNames[1] === "Run alpha check" && ciNames[2] === "Run beta check (spec)");

const labels = parseLintCiLabels(lintCi);
assert("parseLintCiLabels stops at closing paren", labels.length === 2, labels);
assert("parseLintCiLabels ignores comments/blank lines",
  labels[0] === "Run alpha check" && labels[1] === "Run beta check (spec, advisory)");

const result = analyzeDrift(ciNames, labels, ignored);
assert("analyzeDrift flags un-mirrored step",
  result.missing.length === 1 && result.missing[0] === "Run gamma check",
  result.missing);
assert("analyzeDrift accepts parenthetical stem match (no false positive for beta)",
  !result.missing.includes("Run beta check (spec)"));
assert("analyzeDrift respects ignored map",
  !result.missing.includes("Checkout") &&
  !result.missing.includes("Ghost step that no longer exists elsewhere"));
assert("analyzeDrift surfaces STALE ignore",
  result.unusedIgnores.length === 1 && result.unusedIgnores[0] === "Long-gone step",
  result.unusedIgnores);
assert("analyzeDrift counts are accurate",
  result.counts.ci === 4 && result.counts.labels === 2 && result.counts.ignored === 3,
  result.counts);

// Zero-drift case: mirror every ci step, remove stale ignore.
const cleanIgnored = new Map([["Checkout", "infra"], ["Ghost step that no longer exists elsewhere", "infra"]]);
const cleanLabels = [...labels, "Run gamma check"];
const clean = analyzeDrift(ciNames, cleanLabels, cleanIgnored);
assert("analyzeDrift returns zero drift when fully mirrored",
  clean.missing.length === 0 && clean.unusedIgnores.length === 0,
  clean);

if (failed) {
  console.log(`\n${failed} case(s) failed`);
  process.exit(1);
}
console.log("\nOK — check-lint-ci-drift self-test passed.");
