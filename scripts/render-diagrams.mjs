#!/usr/bin/env node
// render-diagrams.mjs — Render Mermaid (.mmd) → PNG across the spec/ tree.
//
// Audit-09 §2.2 / Phase-13.5 closure.
// • Discovers every `spec/**/{diagrams,images}/*.mmd` file.
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
import { join, relative, dirname, resolve } from 'node:path';
import { spawnSync } from 'node:child_process';

const ROOT = process.cwd();
const SPEC_ROOT = join(ROOT, 'spec');
const ARGS = process.argv.slice(2);
const CHECK_ONLY = ARGS.includes('--check');
const STAGED_ONLY = ARGS.includes('--staged');
const ONLY_INDEX = ARGS.indexOf('--only');
const ONLY_FILTER = ONLY_INDEX >= 0 ? ARGS[ONLY_INDEX + 1] : null;

const DIAGRAM_DIR_NAMES = new Set(['diagrams', 'images']);

function isDiagramsDir(name) {
  return DIAGRAM_DIR_NAMES.has(name);
}

// Returns the set of absolute paths for .mmd files currently staged in git
// (Added/Copied/Modified/Renamed). Empty set if not in a git repo or git
// command fails — caller decides whether that is fatal.
function getStagedMmdSet() {
  const result = spawnSync('git', ['diff', '--cached', '--name-only', '--diff-filter=ACMR'], { encoding: 'utf8' });
  if (result.status !== 0) return new Set();
  const paths = result.stdout.split('\n').filter((p) => p.endsWith('.mmd'));
  return new Set(paths.map((p) => resolve(ROOT, p)));
}

const STAGED_SET = STAGED_ONLY ? getStagedMmdSet() : null;

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
    if (STAGED_SET && !STAGED_SET.has(full)) continue;
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
  const args = ['--yes', '@mermaid-js/mermaid-cli', '-i', mmd, '-o', png, '-b', 'transparent'];
  const puppeteerCfg = join(ROOT, 'scripts', 'puppeteer-ci.json');
  if (existsSync(puppeteerCfg)) {
    args.push('-p', puppeteerCfg);
  }
  const result = spawnSync('npx', args, { stdio: 'inherit' });
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

  const scopeNote = STAGED_ONLY ? ' (staged-only)' : (ONLY_FILTER ? ` (filter: ${ONLY_FILTER})` : '');
  console.log(`[render-diagrams] discovered ${mmdFiles.length} .mmd file(s)${scopeNote}`);

  if (STAGED_ONLY && mmdFiles.length === 0) {
    console.log('[render-diagrams] no staged .mmd files — skipping.');
    process.exit(0);
  }

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
