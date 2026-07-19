#!/usr/bin/env node
/**
 * scripts/tests/branch-protection-diff.test.mjs
 *
 * Self-test for scripts/branch-protection-diff.mjs `diff(live, expected)`.
 * Guards the classification logic that decides whether live GitHub
 * branch-protection required_status_checks match .github/branch-protection.expected.json.
 *
 * Regression classes covered:
 *   1. Exact match (both empty).
 *   2. Exact match (identical order).
 *   3. Exact match (different order; set-equal).
 *   4. Missing on live (expected has entries live lacks).
 *   5. Extra on live (live has entries not in expected).
 *   6. Mixed drift (both missing and extra).
 *   7. Duplicate entries on live do not create phantom "missing" entries.
 *
 * Runs standalone. Exit code = number of failures.
 */
import { diff } from "../branch-protection-diff.mjs";

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

console.log("branch-protection-diff.test.mjs");

// 1
{
  const r = diff([], []);
  assert("empty match: no missing", r.missingOnLive.length === 0, r);
  assert("empty match: no extra", r.extraOnLive.length === 0, r);
}
// 2
{
  const r = diff(["a", "b"], ["a", "b"]);
  assert("identical order: no drift", r.missingOnLive.length === 0 && r.extraOnLive.length === 0, r);
}
// 3
{
  const r = diff(["b", "a"], ["a", "b"]);
  assert("different order: still set-equal", r.missingOnLive.length === 0 && r.extraOnLive.length === 0, r);
}
// 4
{
  const r = diff(["a"], ["a", "b", "c"]);
  assert("missing on live: flags b and c", JSON.stringify(r.missingOnLive) === JSON.stringify(["b", "c"]), r);
  assert("missing on live: no false extras", r.extraOnLive.length === 0, r);
}
// 5
{
  const r = diff(["a", "b", "c"], ["a"]);
  assert("extra on live: flags b and c", JSON.stringify(r.extraOnLive) === JSON.stringify(["b", "c"]), r);
  assert("extra on live: no false missing", r.missingOnLive.length === 0, r);
}
// 6
{
  const r = diff(["a", "x"], ["a", "b"]);
  assert("mixed drift: missing b", JSON.stringify(r.missingOnLive) === JSON.stringify(["b"]), r);
  assert("mixed drift: extra x", JSON.stringify(r.extraOnLive) === JSON.stringify(["x"]), r);
}
// 7
{
  const r = diff(["a", "a", "b"], ["a", "b"]);
  assert("duplicate on live: no phantom missing", r.missingOnLive.length === 0, r);
  // Note: current impl reports the duplicate as extra=[]; both "a" are in expSet.
  assert("duplicate on live: dedup via set membership", r.extraOnLive.length === 0, r);
}

if (failed > 0) {
  console.log(`FAIL ${failed} assertion(s)`);
  process.exit(1);
}
console.log("ok, all diff() assertions passed");
