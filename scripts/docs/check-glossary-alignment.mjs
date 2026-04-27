#!/usr/bin/env node
// Asserts the §1a glossary table in docs/installer-fix-repo-flags.md keeps
// the **Env var** and **Precedence** columns width-aligned across every row.
// Zero deps. Exit 0 = aligned, 1 = misaligned, 2 = parse error.

import { readFileSync } from 'node:fs';

const DOC = 'docs/installer-fix-repo-flags.md';
const SECTION_RE = /^##\s+1a\.\s+Glossary/;
const TARGET_HEADERS = ['Env var', 'Precedence (highest → lowest)'];

function fail(msg) { console.error(`❌ ${msg}`); process.exit(1); }
function bail(msg) { console.error(`⚠️  ${msg}`); process.exit(2); }

const lines = readFileSync(DOC, 'utf8').split('\n');

// Locate the first markdown table after the §1a header.
const sectionStart = lines.findIndex((l) => SECTION_RE.test(l));
if (sectionStart < 0) bail(`Could not find "## 1a. Glossary" in ${DOC}`);

const headerIdx = lines.findIndex(
  (l, i) => i > sectionStart && l.trimStart().startsWith('|') && l.includes('Env var'),
);
if (headerIdx < 0) bail('Could not find glossary header row');

// Collect contiguous table rows: header, separator, then data rows.
const rows = [];
for (let i = headerIdx; i < lines.length; i++) {
  const line = lines[i];
  const isTableRow = line.trimStart().startsWith('|');
  if (!isTableRow) break;
  rows.push(line);
}
if (rows.length < 3) bail(`Glossary table too short (${rows.length} rows)`);

// Split a pipe row into raw cells (preserving padding/whitespace).
function splitCells(row) {
  const trimmed = row.trim().replace(/^\|/, '').replace(/\|$/, '');
  return trimmed.split('|');
}

const headerCells = splitCells(rows[0]).map((c) => c.trim());
const targetIndices = TARGET_HEADERS.map((h) => {
  const idx = headerCells.indexOf(h);
  if (idx < 0) bail(`Header "${h}" missing from glossary table`);
  return idx;
});

// Skip header (0) and separator (1); audit data rows only.
const dataRows = rows.slice(2);
const widths = new Map(); // headerName -> Set of observed cell widths

for (const headerName of TARGET_HEADERS) widths.set(headerName, new Set());

for (const row of dataRows) {
  const cells = splitCells(row);
  for (let i = 0; i < TARGET_HEADERS.length; i++) {
    const headerName = TARGET_HEADERS[i];
    const cellIdx = targetIndices[i];
    const cell = cells[cellIdx];
    if (cell === undefined) fail(`Row missing column "${headerName}": ${row}`);
    widths.get(headerName).add(cell.length);
  }
}

const misaligned = [];
for (const [headerName, widthSet] of widths) {
  if (widthSet.size === 1) continue;
  misaligned.push(`  • "${headerName}" column has ${widthSet.size} distinct widths: ${[...widthSet].sort((a, b) => a - b).join(', ')}`);
}

if (misaligned.length > 0) {
  console.error('❌ Glossary table columns are not width-aligned:');
  for (const m of misaligned) console.error(m);
  console.error(`\nFix: re-pad the affected cells in ${DOC} §1a so every row matches.`);
  process.exit(1);
}

console.log(`✅ ${DOC} §1a glossary: "${TARGET_HEADERS.join('" and "')}" columns are width-aligned across ${dataRows.length} data rows.`);
