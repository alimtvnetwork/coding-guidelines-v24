#!/usr/bin/env node
// =====================================================================
// check-bundle-installers-drift.mjs
//
// Pre-test guard: re-runs the bundle-installer generator IN-MEMORY and
// compares the output to the committed `<bundle>-install.{sh,ps1}`
// files at the repo root. Fails fast with a unified diff if any file
// has drifted.
//
// Why: the installer test suite asserts properties of the *generated*
// installers. If the committed files are stale, downstream tests
// silently exercise old behavior. This check makes the drift loud.
//
// Usage:
//   node scripts/check-bundle-installers-drift.mjs
//
// Exit codes:
//   0 — all 14 files match generator output
//   1 — at least one file has drifted (diff printed to stdout)
//   2 — internal error (manifest load / generator import failed)
// =====================================================================

import { readFileSync, existsSync } from "node:fs";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const HERE = dirname(fileURLToPath(import.meta.url));
const REPO_ROOT = resolve(HERE, "..");

let BUNDLES, bashScript, powershellScript;
try {
  ({ BUNDLES, bashScript, powershellScript } = await import(
    resolve(HERE, "generate-bundle-installers.mjs")
  ));
} catch (err) {
  console.error(`drift-check: failed to import generator: ${err.message}`);
  process.exit(2);
}

function unifiedDiff(expected, actual, path) {
  const exp = expected.split("\n");
  const act = actual.split("\n");
  const out = [`--- ${path} (committed)`, `+++ ${path} (generator)`];
  const max = Math.max(exp.length, act.length);
  let shown = 0;
  for (let i = 0; i < max && shown < 40; i++) {
    if (exp[i] === act[i]) continue;
    out.push(`@@ line ${i + 1} @@`);
    if (exp[i] !== undefined) out.push(`- ${exp[i]}`);
    if (act[i] !== undefined) out.push(`+ ${act[i]}`);
    shown++;
  }
  if (shown >= 40) out.push("... (diff truncated at 40 hunks)");
  return out.join("\n");
}

function checkOne(relPath, expected) {
  const abs = resolve(REPO_ROOT, relPath);
  if (!existsSync(abs)) {
    return { path: relPath, status: "missing" };
  }
  const actual = readFileSync(abs, "utf8");
  if (actual === expected) return { path: relPath, status: "ok" };
  return { path: relPath, status: "drift", diff: unifiedDiff(actual, expected, relPath) };
}

const results = [];
for (const bundle of BUNDLES) {
  results.push(checkOne(`${bundle.name}-install.sh`, bashScript(bundle)));
  results.push(checkOne(`${bundle.name}-install.ps1`, powershellScript(bundle)));
}

const drifted = results.filter((r) => r.status !== "ok");
const ok      = results.length - drifted.length;

if (drifted.length === 0) {
  console.log(`✅ bundle installer drift check: ${ok}/${results.length} files match generator output`);
  process.exit(0);
}

console.error(`❌ bundle installer drift detected: ${drifted.length}/${results.length} files differ from generator output\n`);
for (const r of drifted) {
  if (r.status === "missing") {
    console.error(`  • ${r.path}: MISSING (generator would create it)`);
    continue;
  }
  console.error(`  • ${r.path}: DRIFTED`);
  console.error(r.diff);
  console.error("");
}
console.error("Fix: run `node scripts/generate-bundle-installers.mjs` and commit the regenerated files.");
process.exit(1);
