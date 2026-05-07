#!/usr/bin/env node
// auto-update-readme.mjs — Auto-bump version + regenerate readme.md
// "What's new in vX.Y.Z" subsection when diagram or spec changes land.
//
// What it does (single command):
//   1. Detects changed files vs a baseline (default: last git tag, fallback
//      HEAD~1) scoped to `spec/**` and `spec/**/{diagrams,images}/*.{mmd,png}`.
//   2. If no relevant changes → exits 0 with "no-op" message.
//   3. Otherwise bumps package.json patch (or --minor) version.
//   4. Replaces the existing "### What's new in vX.Y.Z" block in readme.md
//      with a freshly generated one summarizing the detected change set
//      (diagrams updated, spec chapters touched, audit docs added).
//   5. Runs `npm run sync` so all sync-managed artifacts pick up the bump.
//
// Usage:
//   node scripts/auto-update-readme.mjs            # patch bump
//   node scripts/auto-update-readme.mjs --minor    # minor bump
//   node scripts/auto-update-readme.mjs --base HEAD~3
//   node scripts/auto-update-readme.mjs --dry-run  # plan only, no writes
//
// CODE RED: small functions, zero nesting, explicit failure surfacing.

import { readFileSync, writeFileSync } from 'node:fs';
import { spawnSync } from 'node:child_process';
import { resolve, dirname } from 'node:path';
import { fileURLToPath } from 'node:url';

const ROOT = resolve(dirname(fileURLToPath(import.meta.url)), '..');
const README = resolve(ROOT, 'readme.md');
const PKG = resolve(ROOT, 'package.json');

const ARGS = process.argv.slice(2);
const DRY = ARGS.includes('--dry-run');
const MINOR = ARGS.includes('--minor');
const BASE = readFlag('--base') ?? defaultBase();

function readFlag(name) {
  const i = ARGS.indexOf(name);
  return i >= 0 ? ARGS[i + 1] : null;
}

function git(args) {
  const r = spawnSync('git', args, { cwd: ROOT, encoding: 'utf8' });
  return { ok: r.status === 0, out: (r.stdout || '').trim() };
}

function defaultBase() {
  const tag = git(['describe', '--tags', '--abbrev=0']);
  return tag.ok && tag.out ? tag.out : 'HEAD~1';
}

function changedFiles(base) {
  const r = git(['diff', '--name-only', `${base}...HEAD`]);
  if (!r.ok) return [];
  return r.out.split('\n').filter(Boolean);
}

function isDiagram(p) {
  return /^spec\/.+\/(diagrams|images)\/.+\.(mmd|png)$/.test(p);
}

function isSpecMarkdown(p) {
  return p.startsWith('spec/') && p.endsWith('.md');
}

function isAuditDoc(p) {
  return /^spec\/[^/]+\/audit\/.+\.md$/.test(p);
}

function classify(files) {
  const diagrams = files.filter(isDiagram);
  const audits = files.filter(isAuditDoc);
  const specMd = files.filter((p) => isSpecMarkdown(p) && !isAuditDoc(p));
  return { diagrams, audits, specMd };
}

function bumpSemver(version, kind) {
  const [maj, min, pat] = version.split('.').map(Number);
  if (kind === 'minor') return `${maj}.${min + 1}.0`;
  return `${maj}.${min}.${pat + 1}`;
}

function uniqueSpecFolders(files) {
  const set = new Set();
  for (const f of files) {
    const m = f.match(/^spec\/([^/]+)\//);
    if (m) set.add(m[1]);
  }
  return [...set].sort();
}

function todayIso() {
  return new Date().toISOString().slice(0, 10);
}

function buildBullets({ diagrams, audits, specMd }, version) {
  const bullets = [];
  bullets.push(`- **Auto-bumped to v${version}** on ${todayIso()} via \`scripts/auto-update-readme.mjs\`.`);
  if (diagrams.length > 0) {
    const folders = uniqueSpecFolders(diagrams).map((f) => `\`spec/${f}\``).join(', ');
    bullets.push(`- **Diagrams updated (${diagrams.length} file${diagrams.length === 1 ? '' : 's'})** across ${folders}. Re-rendered via \`npm run diagrams:rebaseline\`.`);
  }
  if (audits.length > 0) {
    const links = audits.map((a) => `[\`${a}\`](${a})`).join(', ');
    bullets.push(`- **New audit doc${audits.length === 1 ? '' : 's'}:** ${links}.`);
  }
  if (specMd.length > 0) {
    const folders = uniqueSpecFolders(specMd).map((f) => `\`spec/${f}\``).join(', ');
    bullets.push(`- **Spec markdown changes** in ${folders} (${specMd.length} file${specMd.length === 1 ? '' : 's'}). See \`CHANGELOG.md\` for detail.`);
  }
  return bullets.join('\n');
}

function buildBlock(version, bullets) {
  return `### What's new in v${version}\n\n${bullets}\n`;
}

function replaceWhatsNew(readme, block) {
  const re = /### What's new in v[\d.]+\n[\s\S]*?(?=\n---|\n## |$)/;
  if (re.test(readme)) return readme.replace(re, block + '\n');
  // Insert before the first `---` after the spec-tree block as a safe fallback.
  return readme.replace(/(\nLive spec tree:[^\n]+\n)/, `$1\n${block}\n`);
}

function bumpPackageJson(version) {
  const pkg = JSON.parse(readFileSync(PKG, 'utf8'));
  pkg.version = version;
  writeFileSync(PKG, JSON.stringify(pkg, null, 2) + '\n');
}

function runSync() {
  const r = spawnSync('npm', ['run', 'sync'], { cwd: ROOT, stdio: 'inherit' });
  return r.status === 0;
}

function main() {
  const files = changedFiles(BASE);
  const groups = classify(files);
  const total = groups.diagrams.length + groups.audits.length + groups.specMd.length;
  console.log(`[auto-update-readme] base=${BASE} changed=${files.length} relevant=${total}`);
  console.log(`  diagrams=${groups.diagrams.length} audits=${groups.audits.length} specMd=${groups.specMd.length}`);

  if (total === 0) {
    console.log('[auto-update-readme] No diagram/spec changes detected — nothing to do.');
    return 0;
  }

  const pkg = JSON.parse(readFileSync(PKG, 'utf8'));
  const next = bumpSemver(pkg.version, MINOR ? 'minor' : 'patch');
  const block = buildBlock(next, buildBullets(groups, next));
  const readme = readFileSync(README, 'utf8');
  const updated = replaceWhatsNew(readme, block);

  if (DRY) {
    console.log(`\n[dry-run] would bump ${pkg.version} → ${next}`);
    console.log(`[dry-run] would write block:\n\n${block}`);
    return 0;
  }

  bumpPackageJson(next);
  writeFileSync(README, updated);
  console.log(`✔ bumped ${pkg.version} → ${next}; readme "What's new" regenerated.`);

  const ok = runSync();
  if (!ok) {
    console.error('✘ npm run sync failed — review and re-run before committing.');
    return 2;
  }
  console.log('\nNext: review readme.md + add a CHANGELOG entry, then commit.');
  return 0;
}

process.exit(main());
