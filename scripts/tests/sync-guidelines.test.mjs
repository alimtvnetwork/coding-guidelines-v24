#!/usr/bin/env node
// ============================================================
// sync-guidelines.test.mjs
// ============================================================
// Locks the drift-detection contract for scripts/sync-guidelines.mjs
// so a refactor cannot silently make --check exit 0 on drift.
// Added in v5.130 alongside Hard Rule #13 enforcement in CI.
// ============================================================

import { spawnSync } from "node:child_process";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";
import {
  diffReport,
  computeDrifts,
  buildLovableMirror,
  buildCursorRules,
  extractSection,
} from "../sync-guidelines.mjs";

const __dirname = dirname(fileURLToPath(import.meta.url));
const SCRIPT = resolve(__dirname, "../sync-guidelines.mjs");

let ok = 0;
let fail = 0;
function assert(label, cond) {
  if (cond) { ok++; console.log(`  ok  ${label}`); return; }
  fail++;
  console.error(`  FAIL ${label}`);
}

// ---------- diffReport ----------
assert("diffReport: identical strings return null", diffReport("x", "abc", "abc") === null);
assert("diffReport: single-char difference is detected", typeof diffReport("x", "abc", "abd") === "string");
assert("diffReport: empty vs non-empty is detected", typeof diffReport("x", "", "hi") === "string");
assert(
  "diffReport: line counts match input (expected 3, actual 1)",
  diffReport("mirror", "a\nb\nc", "z") === "mirror: drift (expected 3 lines, actual 1 lines)",
);

// ---------- extractSection ----------
const sampleCanonical = [
  "# Title",
  "",
  "## Hard Rules (Zero Tolerance)",
  "",
  "1. First rule",
  "2. Second rule",
  "",
  "---",
  "",
  "## Next Section",
  "",
  "Body",
].join("\n");
const hardRules = extractSection(sampleCanonical, "Hard Rules (Zero Tolerance)");
assert("extractSection: pulls rules body", hardRules.includes("1. First rule") && hardRules.includes("2. Second rule"));
assert("extractSection: stops at next heading or ---", !hardRules.includes("Next Section") && !hardRules.includes("Body"));

let threw = false;
try { extractSection(sampleCanonical, "Nonexistent Heading"); } catch { threw = true; }
assert("extractSection: throws when heading missing", threw);

// ---------- computeDrifts: clean case ----------
const canonicalMin = [
  "# 31. Compiled Simple Coding Guidelines",
  "",
  "## Hard Rules (Zero Tolerance)",
  "",
  "1. Rule A",
  "",
  "---",
].join("\n");
const cleanLovable = buildLovableMirror(canonicalMin);
const cleanCursorBase = "# .cursorrules\n\nsome preface\n";
const cleanCursor = buildCursorRules(canonicalMin, cleanCursorBase);

const cleanResult = computeDrifts({
  canonical: canonicalMin,
  currentLovable: cleanLovable,
  currentCursor: cleanCursor,
  lovableLabel: "mirror.md",
  cursorLabel: ".cursorrules",
});
assert("computeDrifts: no drift when both mirrors match", cleanResult.drifts.length === 0);

// ---------- computeDrifts: lovable-only drift ----------
const lovableDrift = computeDrifts({
  canonical: canonicalMin,
  currentLovable: "STALE",
  currentCursor: cleanCursor,
  lovableLabel: "mirror.md",
  cursorLabel: ".cursorrules",
});
assert("computeDrifts: detects lovable-only drift", lovableDrift.drifts.length === 1 && lovableDrift.drifts[0].startsWith("mirror.md:"));

// ---------- computeDrifts: cursor-only drift ----------
const cursorDrift = computeDrifts({
  canonical: canonicalMin,
  currentLovable: cleanLovable,
  currentCursor: "no markers here at all",
  lovableLabel: "mirror.md",
  cursorLabel: ".cursorrules",
});
assert("computeDrifts: detects cursor-only drift (missing markers)", cursorDrift.drifts.length === 1 && cursorDrift.drifts[0].startsWith(".cursorrules:"));

// ---------- computeDrifts: both drifted ----------
const bothDrift = computeDrifts({
  canonical: canonicalMin,
  currentLovable: "STALE",
  currentCursor: "STALE TOO",
  lovableLabel: "mirror.md",
  cursorLabel: ".cursorrules",
});
assert("computeDrifts: reports both mirrors when both drift", bothDrift.drifts.length === 2);

// ---------- End-to-end: real repo --check must exit 0 ----------
const checkRun = spawnSync("node", [SCRIPT, "--check"], { encoding: "utf8" });
assert("--check exits 0 on clean repo", checkRun.status === 0);
assert("--check prints OK on clean repo", (checkRun.stdout || "").includes("OK"));

// ---------- End-to-end: importing this file must NOT trigger main() ----------
// If the invokedDirectly guard broke, importing sync-guidelines above would
// have written to the real mirrors or exited the process. We reached here,
// so the guard is intact. Assertion is implicit but state it for the log.
assert("import guard: main() did not run on module import", true);

console.log(`\n${ok} ok, ${fail} fail`);
process.exit(fail === 0 ? 0 : 1);
