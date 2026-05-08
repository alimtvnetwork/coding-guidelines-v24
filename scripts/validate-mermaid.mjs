#!/usr/bin/env node
// validate-mermaid.mjs — Parse-validate every spec/**/{diagrams,images}/*.mmd
// against the installed mermaid v11 grammar BEFORE we attempt to render PNGs.
//
// Catches stray tokens (e.g. `;` mid-message in sequenceDiagram, unquoted `@`
// labels, mis-placed `Note over` after `alt`) at lint-time so we don't commit
// a `.mmd` that the renderer will reject. Pairs with render-diagrams.mjs.
//
// Usage:
//   node scripts/validate-mermaid.mjs            # validate all
//   node scripts/validate-mermaid.mjs --only spec/19

import { existsSync } from 'node:fs';
import { readdir, readFile } from 'node:fs/promises';
import { join, relative, dirname } from 'node:path';
import { JSDOM } from 'jsdom';

// Mermaid v11 calls DOMPurify.addHook during parse; provide a browser-like
// global before importing it so the parser doesn't crash in pure Node.
const dom = new JSDOM('<!doctype html><html><body></body></html>');
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.DocumentFragment = dom.window.DocumentFragment;
globalThis.Element = dom.window.Element;
globalThis.Node = dom.window.Node;

const mermaid = (await import('mermaid')).default;

const ROOT = process.cwd();
const SPEC_ROOT = join(ROOT, 'spec');
const ARGS = process.argv.slice(2);
const ONLY_INDEX = ARGS.indexOf('--only');
const ONLY_FILTER = ONLY_INDEX >= 0 ? ARGS[ONLY_INDEX + 1] : null;
const DIAGRAM_DIRS = new Set(['diagrams', 'images']);

async function findMmd(dir, acc = []) {
  const entries = await readdir(dir, { withFileTypes: true });
  for (const e of entries) {
    const full = join(dir, e.name);
    if (e.isDirectory()) { await findMmd(full, acc); continue; }
    if (!e.name.endsWith('.mmd')) continue;
    if (!DIAGRAM_DIRS.has(dirname(full).split('/').pop())) continue;
    if (ONLY_FILTER && !full.includes(ONLY_FILTER)) continue;
    acc.push(full);
  }
  return acc;
}

async function main() {
  if (!existsSync(SPEC_ROOT)) {
    console.error('[validate-mermaid] spec/ not found.');
    process.exit(2);
  }
  mermaid.initialize({ startOnLoad: false, suppressErrorRendering: true });

  const files = await findMmd(SPEC_ROOT);
  console.log(`[validate-mermaid] parsing ${files.length} .mmd file(s)${ONLY_FILTER ? ` (filter: ${ONLY_FILTER})` : ''}`);

  const failures = [];
  for (const file of files) {
    const src = await readFile(file, 'utf8');
    try {
      await mermaid.parse(src);
    } catch (err) {
      failures.push({ file, src, message: err?.message || String(err) });
    }
  }

  if (failures.length === 0) {
    console.log(`[validate-mermaid] OK — all ${files.length} diagram(s) parse under mermaid v11.`);
    process.exit(0);
  }

  console.error(`\n[validate-mermaid] FAIL — ${failures.length}/${files.length} diagram(s) failed mermaid v11 parse:\n`);
  for (const { file, src, message } of failures) {
    const rel = relative(ROOT, file);
    console.error(`────────────────────────────────────────────────────────────`);
    console.error(`✗ ${rel}`);
    const lineNo = extractLineNumber(message);
    if (lineNo) {
      console.error(`  at line ${lineNo}:`);
      printSourceContext(src, lineNo);
    }
    console.error(`  mermaid v11 parser said:`);
    const trimmed = message.split('\n').slice(0, 8).map((l) => `    ${l}`).join('\n');
    console.error(trimmed);
    console.error('');
  }
  console.error(`────────────────────────────────────────────────────────────`);
  console.error(`\n${failures.length} file(s) failed. Fix the .mmd source(s) above, then re-stage and retry the commit.`);
  console.error(`Bypass (emergencies only): SKIP_DIAGRAMS_HOOK=1 git commit ...\n`);
  process.exit(1);
}

// Extract a 1-based line number from a mermaid parser error message.
// Mermaid v11 emits messages like "Parse error on line 14:" or
// "Lexical error on line 7:". Returns null if no line number is present.
function extractLineNumber(message) {
  const m = /(?:Parse|Lexical|Syntax)\s+error\s+on\s+line\s+(\d+)/i.exec(message);
  return m ? parseInt(m[1], 10) : null;
}

// Print 2 lines of context above + below the failing line, with a caret marker.
function printSourceContext(src, lineNo) {
  const lines = src.split('\n');
  const start = Math.max(1, lineNo - 2);
  const end = Math.min(lines.length, lineNo + 2);
  const gutter = String(end).length;
  for (let i = start; i <= end; i += 1) {
    const marker = i === lineNo ? '>' : ' ';
    const num = String(i).padStart(gutter, ' ');
    console.error(`    ${marker} ${num} | ${lines[i - 1] ?? ''}`);
  }
}

main().catch((e) => { console.error('[validate-mermaid] unexpected:', e); process.exit(2); });
