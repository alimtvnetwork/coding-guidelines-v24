#!/usr/bin/env node
// ============================================================
// release.mjs — one-shot release ceremony
// ============================================================
//
// Wraps the full release flow into a single command:
//   1. Compute next SemVer from --tier patch|minor|major (or explicit --version)
//   2. Delegate to scripts/bump-version.mjs (which bumps package.json,
//      runs `npm run sync`, and stamps changelogs)
//   3. Optionally run aggregate-prompts if the script exists
//   4. Verify `npm run sync:check` is GREEN post-release
//
// Usage:
//   node scripts/release.mjs --tier patch
//   node scripts/release.mjs --tier minor --scope "Release ceremony wrap-up"
//   node scripts/release.mjs --tier major --target root
//   node scripts/release.mjs --version 5.50.0 --scope "Hot fix"
//   node scripts/release.mjs --tier patch --dry-run
//
// Exit codes:
//   0 success | 1 bad args | 2 bump failed | 3 post-sync drift
// ============================================================

import { readFileSync, existsSync } from "node:fs";
import { spawnSync } from "node:child_process";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const ROOT = resolve(__dirname, "..");

function parseArgs(argv) {
  const out = {
    tier: null,
    version: null,
    scope: null,
    target: "both",
    dryRun: false,
  };
  for (let i = 0; i < argv.length; i++) {
    const a = argv[i];
    if (a === "--tier" || a === "-t") out.tier = argv[++i];
    else if (a === "--version" || a === "-v") out.version = argv[++i];
    else if (a === "--scope" || a === "-s") out.scope = argv[++i];
    else if (a === "--target") out.target = argv[++i];
    else if (a === "--dry-run") out.dryRun = true;
    else if (a === "-h" || a === "--help") {
      console.log("Usage: node scripts/release.mjs --tier patch|minor|major [--scope \"...\"] [--target root|spec19|both] [--dry-run]");
      process.exit(0);
    }
  }
  return out;
}

function fail(code, msg) {
  console.error(`✗ ${msg}`);
  process.exit(code);
}

function readCurrentVersion() {
  const pkg = JSON.parse(readFileSync(resolve(ROOT, "package.json"), "utf8"));
  return pkg.version;
}

function bumpSemver(version, tier) {
  const m = /^(\d+)\.(\d+)\.(\d+)$/.exec(version);
  if (!m) fail(1, `Cannot parse current version '${version}'`);
  let [, major, minor, patch] = m.map(Number);
  if (tier === "patch") patch += 1;
  else if (tier === "minor") { minor += 1; patch = 0; }
  else if (tier === "major") { major += 1; minor = 0; patch = 0; }
  else fail(1, `Invalid --tier '${tier}' (expected patch|minor|major)`);
  return `${major}.${minor}.${patch}`;
}

function resolveTargetVersion(args) {
  if (args.version) {
    if (!/^\d+\.\d+\.\d+$/.test(args.version)) fail(1, "Invalid --version (expected X.Y.Z)");
    return args.version;
  }
  if (!args.tier) fail(1, "Must supply --tier patch|minor|major (or explicit --version)");
  return bumpSemver(readCurrentVersion(), args.tier);
}

function defaultScope(tier) {
  return `Routine ${tier ?? "release"} ceremony (one-shot)`;
}

function runBump(args, nextVersion) {
  const bumpArgs = [
    "scripts/bump-version.mjs",
    "--version", nextVersion,
    "--scope", args.scope ?? defaultScope(args.tier),
    "--target", args.target,
  ];
  if (args.dryRun) bumpArgs.push("--dry-run");
  console.log(`→ node ${bumpArgs.join(" ")}`);
  const res = spawnSync("node", bumpArgs, { cwd: ROOT, stdio: "inherit" });
  if (res.status !== 0) fail(2, "bump-version step failed");
}

function maybeAggregatePrompts() {
  const p = resolve(ROOT, "scripts/aggregate-prompts.mjs");
  if (!existsSync(p)) return;
  console.log("→ node scripts/aggregate-prompts.mjs");
  const res = spawnSync("node", ["scripts/aggregate-prompts.mjs"], { cwd: ROOT, stdio: "inherit" });
  if (res.status !== 0) fail(2, "aggregate-prompts failed");
}

function verifySync(dryRun) {
  if (dryRun) return;
  console.log("→ npm run sync:check");
  const res = spawnSync("npm", ["run", "--silent", "sync:check"], { cwd: ROOT, stdio: "inherit" });
  if (res.status !== 0) fail(3, "post-release sync:check drift — run `npm run sync` and inspect");
}

function main() {
  const args = parseArgs(process.argv.slice(2));
  const current = readCurrentVersion();
  const next = resolveTargetVersion(args);
  console.log(`\n=== Release ceremony: ${current} → ${next} (${args.tier ?? "explicit"}) ===\n`);
  runBump(args, next);
  maybeAggregatePrompts();
  verifySync(args.dryRun);
  console.log(`\n✓ Release ${args.dryRun ? "(dry-run) " : ""}complete: v${next}`);
}

main();
