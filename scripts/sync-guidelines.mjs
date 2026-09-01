#!/usr/bin/env node
// ============================================================
// sync-guidelines.mjs
// ============================================================
// Single source of truth: spec/17-consolidated-guidelines/34-compiled-simple-coding-guidelines.md
//
// Mirrors:
//   1. .lovable/coding-guidelines.md   (exact body copy + auto-gen banner)
//   2. .cursorrules                                       (Hard Rules block injected between markers)
//
// Usage:
//   node scripts/sync-guidelines.mjs           # write mirrors
//   node scripts/sync-guidelines.mjs --check   # exit 1 on drift, no writes
// ============================================================

import { readFileSync, writeFileSync } from "node:fs";
import { resolve, dirname, relative } from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const ROOT = resolve(__dirname, "..");

const CANONICAL = resolve(ROOT, "spec/17-consolidated-guidelines/34-compiled-simple-coding-guidelines.md");
const LOVABLE_MIRROR = resolve(ROOT, ".lovable/coding-guidelines.md");
const LOVABLE_NESTED_MIRROR = resolve(ROOT, ".lovable/coding-guidelines.md");
const CURSORRULES = resolve(ROOT, ".cursorrules");

const BANNER = [
  "<!-- AUTO-GENERATED FILE. DO NOT EDIT DIRECTLY. -->",
  "<!-- Source: spec/17-consolidated-guidelines/34-compiled-simple-coding-guidelines.md -->",
  "<!-- Regenerate with: npm run sync:guidelines -->",
  "",
  "",
].join("\n");

const CURSOR_BEGIN = "<!-- BEGIN:SYNC-HARD-RULES -->";
const CURSOR_END = "<!-- END:SYNC-HARD-RULES -->";

function readCanonical() {
  return readFileSync(CANONICAL, "utf8");
}

function extractSection(source, heading) {
  const lines = source.split("\n");
  const startIdx = lines.findIndex((line) => line.trim() === `## ${heading}`);
  if (startIdx === -1) throw new Error(`Cannot find '## ${heading}' in canonical file`);
  let endIdx = lines.length;
  for (let i = startIdx + 1; i < lines.length; i++) {
    if (lines[i].startsWith("## ") || lines[i].startsWith("---")) { endIdx = i; break; }
  }
  return lines.slice(startIdx + 1, endIdx).join("\n").trim();
}

function buildLovableMirror(canonical) {
  return `${BANNER}${canonical}`;
}

function buildCursorRules(canonical, currentCursor) {
  const hardRules = extractSection(canonical, "Hard Rules (Zero Tolerance)");
  const block = [
    CURSOR_BEGIN,
    "<!-- Auto-generated from spec/17-consolidated-guidelines/34-compiled-simple-coding-guidelines.md -->",
    "<!-- Edit the canonical file then run `npm run sync:guidelines`. -->",
    "",
    "## Hard Rules (Zero Tolerance, canonical)",
    "",
    hardRules,
    "",
    CURSOR_END,
  ].join("\n");

  if (currentCursor.includes(CURSOR_BEGIN) && currentCursor.includes(CURSOR_END)) {
    const pattern = new RegExp(`${CURSOR_BEGIN}[\\s\\S]*?${CURSOR_END}`, "m");
    return currentCursor.replace(pattern, block);
  }
  return `${currentCursor.trimEnd()}\n\n${block}\n`;
}

export function diffReport(label, expected, actual) {
  if (expected === actual) return null;
  const expLines = expected.split("\n").length;
  const actLines = actual.split("\n").length;
  return `${label}: drift (expected ${expLines} lines, actual ${actLines} lines)`;
}

export function computeDrifts({ canonical, currentLovable, currentCursor, lovableLabel, cursorLabel }) {
  const nextLovable = buildLovableMirror(canonical);
  const nextCursor = buildCursorRules(canonical, currentCursor);
  return {
    nextLovable,
    nextCursor,
    drifts: [
      diffReport(lovableLabel, nextLovable, currentLovable),
      diffReport(cursorLabel, nextCursor, currentCursor),
    ].filter(Boolean),
  };
}

export { buildLovableMirror, buildCursorRules, extractSection };

function main() {
  const check = process.argv.includes("--check");
  const canonical = readCanonical();
  const currentCursor = readFileSync(CURSORRULES, "utf8");
  const currentLovable = readFileSync(LOVABLE_MIRROR, "utf8");
  const { nextLovable, nextCursor, drifts } = computeDrifts({
    canonical,
    currentLovable,
    currentCursor,
    lovableLabel: relative(ROOT, LOVABLE_MIRROR),
    cursorLabel: relative(ROOT, CURSORRULES),
  });

  if (check) {
    if (drifts.length === 0) { console.log("OK guideline mirrors are up to date"); return; }
    console.error("guideline mirror drift detected:");
    for (const d of drifts) console.error(`  - ${d}`);
    console.error("\nFix: npm run sync:guidelines");
    process.exit(1);
  }

  if (drifts.length === 0) { console.log("OK guideline mirrors already in sync"); return; }
  writeFileSync(LOVABLE_MIRROR, nextLovable);
  writeFileSync(LOVABLE_NESTED_MIRROR, nextLovable);
  writeFileSync(CURSORRULES, nextCursor);
  console.log(`✓ synced ${drifts.length} guideline mirror(s):`);
  for (const d of drifts) console.log(`  - ${d}`);
}

const invokedDirectly = process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url);
if (invokedDirectly) main();
