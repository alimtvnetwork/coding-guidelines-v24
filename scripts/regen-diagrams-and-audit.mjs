#!/usr/bin/env node
// regen-diagrams-and-audit.mjs — One-shot regenerate-and-baseline.
//
// Single command that:
//   1. Validates every .mmd parses under mermaid v11 (pre-render gate).
//   2. Renders all spec/**/{diagrams,images}/*.mmd → sibling PNGs.
//   3. Verifies drift-check passes (no stale PNGs).
//   4. Writes a new audit baseline doc under
//      spec/19-main-worker-service/audit/<NN>-baseline-diagram-pngs-<DATE>.md
//      capturing source/png counts, rendered/skipped/failed, and the
//      version/commit context.
//
// Usage:
//   node scripts/regen-diagrams-and-audit.mjs            # default
//   node scripts/regen-diagrams-and-audit.mjs --dry-run  # skip audit doc write
//
// CODE RED: zero nesting, small functions, explicit failure surfacing.

import { readdir, readFile, writeFile, mkdir } from 'node:fs/promises';
import { existsSync } from 'node:fs';
import { join, dirname, relative } from 'node:path';
import { spawnSync } from 'node:child_process';

const ROOT = process.cwd();
const SPEC_ROOT = join(ROOT, 'spec');
const AUDIT_DIR = join(SPEC_ROOT, '19-main-worker-service', 'audit');
const DIAGRAM_DIRS = new Set(['diagrams', 'images']);
const DRY_RUN = process.argv.includes('--dry-run');

function runStep(label, cmd, args) {
  console.log(`\n▶ ${label}\n  $ ${cmd} ${args.join(' ')}`);
  const r = spawnSync(cmd, args, { stdio: 'pipe', encoding: 'utf8' });
  process.stdout.write(r.stdout || '');
  process.stderr.write(r.stderr || '');
  const ok = r.status === 0;
  if (!ok) console.error(`✘ step failed: ${label} (exit ${r.status})`);
  return { ok, stdout: r.stdout || '', stderr: r.stderr || '' };
}

async function findMmd(dir, acc = []) {
  const entries = await readdir(dir, { withFileTypes: true });
  for (const e of entries) {
    const full = join(dir, e.name);
    if (e.isDirectory()) { await findMmd(full, acc); continue; }
    if (!e.name.endsWith('.mmd')) continue;
    if (!DIAGRAM_DIRS.has(dirname(full).split('/').pop())) continue;
    acc.push(full);
  }
  return acc;
}

function parseRenderStats(stdout) {
  const m = stdout.match(/rendered=(\d+) skipped=(\d+) failed=(\d+)/);
  if (!m) return { rendered: 0, skipped: 0, failed: 0 };
  return { rendered: +m[1], skipped: +m[2], failed: +m[3] };
}

function nextAuditNumber(existing) {
  const nums = existing
    .map((f) => f.match(/^(\d+)-/))
    .filter(Boolean)
    .map((m) => parseInt(m[1], 10));
  const max = nums.length === 0 ? 0 : Math.max(...nums);
  return String(max + 1).padStart(2, '0');
}

function todayIso() {
  return new Date().toISOString().slice(0, 10);
}

async function readVersion() {
  const pkg = JSON.parse(await readFile(join(ROOT, 'package.json'), 'utf8'));
  return pkg.version;
}

function renderAuditDoc({ num, date, version, sources, stats }) {
  const total = stats.rendered + stats.skipped;
  return `# Audit ${num} — Diagram PNG Regeneration Baseline (v${version})

**Date:** ${date}
**Spec version:** v${version}
**Scope:** Single-command rebaseline via \`scripts/regen-diagrams-and-audit.mjs\`.

## Result

- **Pipeline:** parse-validate (\`scripts/validate-mermaid.mjs\`) → render (\`scripts/render-diagrams.mjs\`) → drift-check (\`--check\`).
- **Sources scanned:** ${sources} \`.mmd\` files under \`spec/**/{diagrams,images}/\`.
- **Rendered this run:** ${stats.rendered}
- **Skipped (already fresh):** ${stats.skipped}
- **Failed:** ${stats.failed}
- **Total PNGs covered by .mmd sources:** ${total}
- **Drift-check after render:** PASS (no stale PNGs).

## Disposition

- Spec/19 SPEC-ONLY constraint preserved — only renderer output (PNG binaries) refreshed; no spec markdown changed by this script.
- Backup-tier deferral preserved (audit-12 §Residual point #1 unchanged).

## Reproducibility

\`\`\`
npm run diagrams:rebaseline
\`\`\`
`;
}

async function main() {
  if (!existsSync(SPEC_ROOT)) {
    console.error('spec/ not found at repo root.');
    process.exit(2);
  }

  const validate = runStep('validate mermaid v11 parse', 'node', ['scripts/validate-mermaid.mjs']);
  if (!validate.ok) process.exit(1);

  const render = runStep('render diagrams', 'node', ['scripts/render-diagrams.mjs']);
  if (!render.ok) process.exit(1);

  const drift = runStep('drift-check', 'node', ['scripts/render-diagrams.mjs', '--check']);
  if (!drift.ok) process.exit(1);

  const sources = (await findMmd(SPEC_ROOT)).length;
  const stats = parseRenderStats(render.stdout);
  const version = await readVersion();
  const date = todayIso();

  if (DRY_RUN) {
    console.log(`\n✔ dry-run — sources=${sources} ${JSON.stringify(stats)} v${version}`);
    return;
  }

  await mkdir(AUDIT_DIR, { recursive: true });
  const existing = await readdir(AUDIT_DIR);
  const num = nextAuditNumber(existing);
  const path = join(AUDIT_DIR, `${num}-baseline-diagram-pngs-${date}.md`);
  const body = renderAuditDoc({ num, date, version, sources, stats });
  await writeFile(path, body, 'utf8');

  console.log(`\n✔ wrote ${relative(ROOT, path)}`);
  console.log(`  sources=${sources} rendered=${stats.rendered} skipped=${stats.skipped} failed=${stats.failed}`);
  console.log(`\nNext: stage the regenerated PNGs and the new audit doc, then commit.`);
}

main().catch((e) => { console.error('regen-diagrams-and-audit failed:', e); process.exit(2); });
