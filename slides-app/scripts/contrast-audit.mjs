#!/usr/bin/env node
/**
 * Dark-mode + light-mode contrast audit for slide design tokens.
 *
 * Root cause of the gap: token values in slides-app/src/styles/tokens.css were
 * never verified against WCAG 2.1 AA, so a token tweak could silently drop
 * text pairs below 4.5:1 in either theme. This script parses the token file
 * for both themes and asserts contrast ratios for every foreground/background
 * pair the deck actually renders.
 *
 * Run: node slides-app/scripts/contrast-audit.mjs
 * npm: npm run slides:contrast-audit
 */
import { readFileSync } from "node:fs";
import { fileURLToPath } from "node:url";
import { dirname, resolve } from "node:path";

const HERE = dirname(fileURLToPath(import.meta.url));
const TOKENS_FILE = resolve(HERE, "..", "src", "styles", "tokens.css");

// Foreground -> Background pairs with WCAG AA minimum ratios.
// 4.5 = normal text (WCAG 1.4.3 AA)
// 3.0 = large text / UI components + graphics (WCAG 1.4.11 AA)
const PAIRS = [
  { fg: "fg",         bg: "bg",        min: 4.5, label: "body text on canvas" },
  { fg: "fg",         bg: "bg-raised", min: 4.5, label: "body text on card" },
  { fg: "muted-fg",   bg: "bg",        min: 4.5, label: "muted text on canvas" },
  { fg: "muted-fg",   bg: "bg-raised", min: 4.5, label: "muted text on card" },
  { fg: "primary",    bg: "bg",        min: 3.0, label: "primary UI on canvas" },
  { fg: "primary-fg", bg: "primary",   min: 4.5, label: "text on primary button" },
  { fg: "accent",     bg: "bg",        min: 3.0, label: "accent UI on canvas" },
  { fg: "destructive",bg: "bg",        min: 3.0, label: "destructive UI on canvas" },
  { fg: "fg",         bg: "code-bg",   min: 4.5, label: "code text on code bg" },
];

function parseTokens() {
  const src = readFileSync(TOKENS_FILE, "utf8");
  const dark = {};
  const light = {};
  const rootBlock = src.match(/:root\s*\{([\s\S]*?)\}/);
  const lightBlock = src.match(/\[data-theme="light"\]\s*\{([\s\S]*?)\}/);
  if (!rootBlock || !lightBlock) throw new Error("Could not parse tokens.css blocks");
  extractHsl(rootBlock[1], dark);
  Object.assign(light, dark);
  extractHsl(lightBlock[1], light);
  return { dark, light };
}

function extractHsl(block, target) {
  const re = /--([a-z0-9-]+):\s*([0-9.]+)\s+([0-9.]+)%\s+([0-9.]+)%\s*(?:\/\s*[0-9.]+)?\s*;/g;
  let m;
  while ((m = re.exec(block)) !== null) {
    target[m[1]] = { h: +m[2], s: +m[3], l: +m[4] };
  }
}

function hslToRgb({ h, s, l }) {
  s /= 100; l /= 100;
  const k = (n) => (n + h / 30) % 12;
  const a = s * Math.min(l, 1 - l);
  const f = (n) => l - a * Math.max(-1, Math.min(k(n) - 3, Math.min(9 - k(n), 1)));
  return [f(0), f(8), f(4)];
}

function relLuminance([r, g, b]) {
  const lin = (c) => (c <= 0.03928 ? c / 12.92 : Math.pow((c + 0.055) / 1.055, 2.4));
  return 0.2126 * lin(r) + 0.7152 * lin(g) + 0.0722 * lin(b);
}

function contrast(fgHsl, bgHsl) {
  const l1 = relLuminance(hslToRgb(fgHsl));
  const l2 = relLuminance(hslToRgb(bgHsl));
  const [hi, lo] = l1 > l2 ? [l1, l2] : [l2, l1];
  return (hi + 0.05) / (lo + 0.05);
}

function auditTheme(name, tokens) {
  const rows = [];
  const failures = [];
  for (const pair of PAIRS) {
    const fg = tokens[pair.fg];
    const bg = tokens[pair.bg];
    if (!fg || !bg) {
      failures.push(`[${name}] missing token in pair ${pair.fg}/${pair.bg}`);
      continue;
    }
    const ratio = contrast(fg, bg);
    const ok = ratio >= pair.min;
    rows.push({ pair, ratio, ok });
    if (!ok) {
      failures.push(
        `[${name}] ${pair.fg} on ${pair.bg} (${pair.label}): ${ratio.toFixed(2)}:1 < required ${pair.min}:1`,
      );
    }
  }
  return { rows, failures };
}

function printReport(name, rows) {
  console.log(`\n=== ${name} theme ===`);
  for (const r of rows) {
    const mark = r.ok ? "PASS" : "FAIL";
    console.log(
      `  ${mark}  ${r.pair.fg.padEnd(11)} on ${r.pair.bg.padEnd(11)}  ` +
      `${r.ratio.toFixed(2).padStart(5)}:1 (min ${r.pair.min}:1)  ${r.pair.label}`,
    );
  }
}

function main() {
  const { dark, light } = parseTokens();
  const darkR = auditTheme("dark", dark);
  const lightR = auditTheme("light", light);
  printReport("dark", darkR.rows);
  printReport("light", lightR.rows);
  const failures = [...darkR.failures, ...lightR.failures];
  if (failures.length > 0) {
    console.error(`\nContrast audit FAILED with ${failures.length} violation(s):`);
    for (const f of failures) console.error(`  - ${f}`);
    process.exit(1);
  }
  console.log(`\nContrast audit PASSED. ${darkR.rows.length + lightR.rows.length} pairs meet WCAG AA.`);
}

main();
