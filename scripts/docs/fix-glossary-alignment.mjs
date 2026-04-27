#!/usr/bin/env node
// Re-pads the §1a glossary table in docs/installer-fix-repo-flags.md so every
// cell in every column matches that column's max observed width. Default mode
// is dry-run (prints a diff-style preview, exit 1 if changes needed). Pass
// --fix to write the changes back. Zero deps.
//
// Usage:
//   node scripts/docs/fix-glossary-alignment.mjs           # dry-run, exit 1 if dirty
//   node scripts/docs/fix-glossary-alignment.mjs --fix     # rewrite the file in place

import { readFileSync, writeFileSync } from 'node:fs';

const DOC = 'docs/installer-fix-repo-flags.md';
const SECTION_RE = /^##\s+1a\.\s+Glossary/;
const FIX_MODE = process.argv.includes('--fix');

function bail(msg) { console.error(`⚠️  ${msg}`); process.exit(2); }

const original = readFileSync(DOC, 'utf8');
const lines = original.split('\n');

const sectionStart = lines.findIndex((l) => SECTION_RE.test(l));
if (sectionStart < 0) bail(`Could not find "## 1a. Glossary" in ${DOC}`);

const headerIdx = lines.findIndex(
  (l, i) => i > sectionStart && l.trimStart().startsWith('|') && l.includes('Env var'),
);
if (headerIdx < 0) bail('Could not find glossary header row');

// Collect contiguous table rows (header + separator + data).
const tableRows = [];
let endIdx = headerIdx;
for (let i = headerIdx; i < lines.length; i++) {
  if (!lines[i].trimStart().startsWith('|')) break;
  tableRows.push(lines[i]);
  endIdx = i;
}
if (tableRows.length < 3) bail(`Glossary table too short (${tableRows.length} rows)`);

// Split into raw cells, preserving the leading/trailing single space convention.
function splitRow(row) {
  const trimmed = row.trim();
  const inner = trimmed.replace(/^\|/, '').replace(/\|$/, '');
  return inner.split('|').map((c) => c.replace(/^ /, '').replace(/ $/, ''));
}

const headerCells = splitRow(tableRows[0]);
const separatorCells = splitRow(tableRows[1]);
const dataRows = tableRows.slice(2).map(splitRow);
const colCount = headerCells.length;

// Validate column counts match (separator rows in markdown can be `---` only).
for (const [i, row] of dataRows.entries()) {
  if (row.length !== colCount) {
    bail(`Row ${i + 3} has ${row.length} cells; expected ${colCount}`);
  }
}

// Compute max content width per column across header + data rows
// (separators are pure dashes — we'll regenerate them to match).
function visibleWidth(cell) { return cell.length; }

const maxWidths = new Array(colCount).fill(0);
for (const row of [headerCells, ...dataRows]) {
  for (let i = 0; i < colCount; i++) {
    const w = visibleWidth(row[i]);
    if (w > maxWidths[i]) maxWidths[i] = w;
  }
}

function padCell(cell, width) {
  const pad = width - visibleWidth(cell);
  return pad > 0 ? cell + ' '.repeat(pad) : cell;
}

function rebuildRow(cells) {
  const padded = cells.map((c, i) => padCell(c, maxWidths[i]));
  return `| ${padded.join(' | ')} |`;
}

function rebuildSeparator() {
  // Preserve any alignment colons (`:---`, `---:`, `:---:`) the author chose.
  const segments = separatorCells.map((sep, i) => {
    const left = sep.startsWith(':') ? ':' : '-';
    const right = sep.endsWith(':') ? ':' : '-';
    const dashCount = Math.max(maxWidths[i] - (left === ':' ? 1 : 0) - (right === ':' ? 1 : 0), 3);
    return `${left}${'-'.repeat(dashCount)}${right}`;
  });
  return `| ${segments.join(' | ')} |`;
}

const rebuilt = [
  rebuildRow(headerCells),
  rebuildSeparator(),
  ...dataRows.map(rebuildRow),
];

const changed = rebuilt.some((row, i) => row !== tableRows[i]);
if (!changed) {
  console.log(`✅ ${DOC} §1a glossary already aligned (${dataRows.length} data rows, ${colCount} columns).`);
  process.exit(0);
}

// Show a per-row diff preview.
console.log(`📐 Glossary table needs re-padding (${dataRows.length + 2} rows):\n`);
for (let i = 0; i < rebuilt.length; i++) {
  if (rebuilt[i] === tableRows[i]) continue;
  console.log(`  - ${tableRows[i]}`);
  console.log(`  + ${rebuilt[i]}`);
}

if (!FIX_MODE) {
  console.log(`\nRun \`node scripts/docs/fix-glossary-alignment.mjs --fix\` to apply.`);
  process.exit(1);
}

// Splice rebuilt rows back in and write.
const updatedLines = [...lines];
for (let i = 0; i < rebuilt.length; i++) updatedLines[headerIdx + i] = rebuilt[i];
writeFileSync(DOC, updatedLines.join('\n'));
console.log(`\n✅ Wrote re-padded glossary table to ${DOC}.`);
