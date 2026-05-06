#!/usr/bin/env node
// render-diagrams.mjs — Render Mermaid (.mmd) → PNG across the spec/ tree.
//
// Audit-09 §2.2 / Phase-13.5 closure.
// • Discovers every `spec/**/diagrams/*.mmd` file.
// • Renders to a sibling `*.png` using @mermaid-js/mermaid-cli (`mmdc`).
// • Honours `--check` mode: exits non-zero if any PNG is missing or older
//   than its `.mmd` source (CI drift guard, no rendering attempted).
// • Honours `--only <glob-substring>` to scope to one spec folder.
//
// Local render usage (requires `npx @mermaid-js/mermaid-cli` available):
//   node scripts/render-diagrams.mjs
//   node scripts/render-diagrams.mjs --only spec/19
//
// CI drift-check usage (no Mermaid runtime needed; passes when no PNGs exist
// yet — opt-in adoption):
//   node scripts/render-diagrams.mjs --check

import { existsSync, statSync } from 'node:fs';
import { readdir } from 'node:fs/promises';
import { join, relative, dirname } from 'node:path';
import { spawnSync } from 'node:child_process';

const ROOT = process.cwd();
const SPEC_ROOT = join(ROOT, 'spec');
const ARGS = process.argv.slice(2);
const CHECK_ONLY = ARGS.includes('--check');
const ONLY_INDEX = ARGS.indexOf('--only');
const ONLY_FILTER = ONLY_INDEX >= 0 ? ARGS[ONLY_INDEX + 1] : null;

function isDiagramsDir(name) {
  return name === 'diagrams';
}

async function findMmdFiles(dir, acc = []) {
  const entries = await readdir(dir, { withFileTypes: true });
  for (const entry of entries) {
    const full = join(dir, entry.name);
    if (entry.isDirectory()) {
      await findMmdFiles(full, acc);
      continue;
    }
    const isMmd = entry.isFile() && entry.name.endsWith('.mmd');
    if (!isMmd) continue;
    if (!isDiagramsDir(dirname(full).split('/').pop())) continue;
    if (ONLY_FILTER && !full.includes(ONLY_FILTER)) continue;
    acc.push(full);
  }
  return acc;
}

function pngPathFor(mmd) {
  return mmd.replace(/\.mmd$/, '.png');
}

function isPngFresh(mmd, png) {
  if (!existsSync(png)) return false;
  return statSync(png).mtimeMs >= statSync(mmd).mtimeMs;
}

function renderOne(mmd, png) {
  const result = spawnSync(
    'npx',
    ['--yes', '@mermaid-js/mermaid-cli', '-i', mmd, '-o', png, '-b', 'transparent'],
    { stdio: 'inherit' },
  );
  return result.status === 0;
}

function reportDrift(stale) {
  if (stale.length === 0) {
    console.log('[render-diagrams] OK — no diagram drift detected.');
    return 0;
  }
  console.error(`[render-diagrams] FAIL — ${stale.length} stale or missing PNG(s):`);
  for (const { mmd, png, reason } of stale) {
    console.error(`  • ${relative(ROOT, mmd)} → ${relative(ROOT, png)}  (${reason})`);
  }
  console.error('\nFix: run `node scripts/render-diagrams.mjs` locally and commit the regenerated PNGs.');
  return 1;
}

async function main() {
  if (!existsSync(SPEC_ROOT)) {
    console.error('[render-diagrams] spec/ not found at repo root.');
    process.exit(2);
  }

  const mmdFiles = await findMmdFiles(SPEC_ROOT);
  console.log(`[render-diagrams] discovered ${mmdFiles.length} .mmd file(s)${ONLY_FILTER ? ` (filter: ${ONLY_FILTER})` : ''}`);

  if (CHECK_ONLY) {
    // Drift-check mode: pass if no PNGs exist yet (adoption is opt-in);
    // fail only when a PNG is present but older than its .mmd source.
    const stale = [];
    for (const mmd of mmdFiles) {
      const png = pngPathFor(mmd);
      if (!existsSync(png)) continue; // adoption pending; not a failure
      if (!isPngFresh(mmd, png)) {
        stale.push({ mmd, png, reason: 'PNG older than .mmd' });
      }
    }
    process.exit(reportDrift(stale));
  }

  // Render mode: regenerate every PNG that is missing or stale.
  let rendered = 0;
  let skipped = 0;
  let failed = 0;
  for (const mmd of mmdFiles) {
    const png = pngPathFor(mmd);
    if (isPngFresh(mmd, png)) {
      skipped += 1;
      continue;
    }
    console.log(`[render-diagrams] rendering ${relative(ROOT, mmd)}`);
    const ok = renderOne(mmd, png);
    if (ok) rendered += 1;
    else failed += 1;
  }
  console.log(`[render-diagrams] rendered=${rendered} skipped=${skipped} failed=${failed}`);
  process.exit(failed === 0 ? 0 : 1);
}

main().catch((err) => {
  console.error('[render-diagrams] unexpected error:', err);
  process.exit(2);
});
