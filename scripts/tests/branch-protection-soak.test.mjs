#!/usr/bin/env node
/**
 * scripts/tests/branch-protection-soak.test.mjs
 *
 * Self-test for scripts/branch-protection-soak.mjs `verdict(passRate, sampleSize)`.
 * Guards the threshold logic that gates promotion of desired-but-not-yet-required
 * checks to REQUIRED on branch protection for `main`.
 *
 * Threshold contract (locked, documented in the script header):
 *   sampleSize === 0                     => "NO DATA"
 *   passRate >= 0.95                     => "READY"
 *   0.80 <= passRate < 0.95              => "WATCH"
 *   passRate < 0.80                      => "NOT READY"
 *
 * Regression classes covered:
 *   1. NO DATA when sample is empty (never claim READY on zero runs).
 *   2. Exact boundary at 0.95 counts as READY.
 *   3. Just below 0.95 (e.g. 0.9499) counts as WATCH, not READY.
 *   4. Exact boundary at 0.80 counts as WATCH.
 *   5. Just below 0.80 counts as NOT READY.
 *   6. 1.0 => READY.  0.0 => NOT READY.
 *
 * Runs standalone. Exit code = number of failures.
 */
import { verdict } from "../branch-protection-soak.mjs";

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

console.log("branch-protection-soak.test.mjs");

assert("sampleSize 0 -> NO DATA (never fake READY)", verdict(1.0, 0) === "NO DATA");
assert("sampleSize 0 with zero pass -> NO DATA", verdict(0.0, 0) === "NO DATA");
assert("1.0 -> READY", verdict(1.0, 20) === "READY");
assert("0.95 boundary -> READY", verdict(0.95, 20) === "READY");
assert("0.9499 -> WATCH (below READY threshold)", verdict(0.9499, 20) === "WATCH");
assert("0.90 -> WATCH", verdict(0.9, 20) === "WATCH");
assert("0.80 boundary -> WATCH", verdict(0.8, 20) === "WATCH");
assert("0.7999 -> NOT READY (below WATCH threshold)", verdict(0.7999, 20) === "NOT READY");
assert("0.5 -> NOT READY", verdict(0.5, 20) === "NOT READY");
assert("0.0 -> NOT READY", verdict(0.0, 20) === "NOT READY");

if (failed > 0) {
  console.log(`FAIL ${failed} assertion(s)`);
  process.exit(1);
}
console.log("ok, all verdict() assertions passed");
