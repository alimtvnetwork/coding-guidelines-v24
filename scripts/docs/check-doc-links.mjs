#!/usr/bin/env node
// =====================================================================
// check-doc-links.mjs
//
// Lightweight link checker for documentation files. Verifies that every
// relative markdown link (and image src) in the given files resolves to
// an existing path in the repo. Anchors (#…), absolute URLs (http(s)://,
// mailto:, etc.) and same-file anchors are skipped.
//
// Usage:
//   node scripts/docs/check-doc-links.mjs <file> [<file> …]
//
// Exits 0 if every link resolves; exits 1 and prints a per-file report
// of broken links otherwise. Designed to run cheaply in CI (no network,
// no deps beyond Node stdlib).
// =====================================================================
import { readFileSync, existsSync } from "node:fs";
import { resolve, dirname, isAbsolute } from "node:path";
import { fileURLToPath } from "node:url";

const REPO_ROOT = resolve(dirname(fileURLToPath(import.meta.url)), "..", "..");

// Match [text](target) and ![alt](target). Captures the target only.
const LINK_RE = /(?<!\\)!?\[[^\]]*\]\(([^)\s]+)(?:\s+"[^"]*")?\)/g;

function isExternal(target) {
  return /^[a-z][a-z0-9+.-]*:/i.test(target) || target.startsWith("//");
}

function stripAnchor(target) {
  const hashAt = target.indexOf("#");
  return hashAt === -1 ? target : target.slice(0, hashAt);
}

function stripCodeBlocks(text) {
  // Remove fenced code blocks (``` … ```` etc.) so links inside code
  // samples are not validated as filesystem paths.
  return text.replace(/(`{3,})[\s\S]*?\1/g, "");
}

function collectLinks(filePath) {
  const raw = readFileSync(filePath, "utf8");
  const text = stripCodeBlocks(raw);
  const fileDir = dirname(filePath);
  const results = [];
  for (const match of text.matchAll(LINK_RE)) {
    const target = match[1];
    if (isExternal(target)) continue;
    const path = stripAnchor(target);
    if (path === "" || path === ".") continue;
    const base = isAbsolute(path) ? REPO_ROOT : fileDir;
    const cleaned = isAbsolute(path) ? path.replace(/^\/+/, "") : path;
    const resolved = resolve(base, cleaned);
    results.push({ raw: target, resolved });
  }
  return results;
}

function checkFile(filePath) {
  const links = collectLinks(filePath);
  const broken = links.filter((l) => !existsSync(l.resolved));
  return { filePath, total: links.length, broken };
}

function loadAllowlist(allowlistPath) {
  // Returns Map<sourceFile, Set<rawTarget>>. Empty map when file missing.
  const out = new Map();
  if (!existsSync(allowlistPath)) return out;
  const lines = readFileSync(allowlistPath, "utf8").split(/\r?\n/);
  for (const line of lines) {
    const trimmed = line.trim();
    if (trimmed === "" || trimmed.startsWith("#")) continue;
    const [src, target] = trimmed.split(/\t+/);
    if (!src || !target) continue;
    if (!out.has(src)) out.set(src, new Set());
    out.get(src).add(target);
  }
  return out;
}

function reportFile(t, allowedForFile, broken, total) {
  const allowed = broken.filter((b) => allowedForFile.has(b.raw));
  const real = broken.filter((b) => !allowedForFile.has(b.raw));
  const allowedResolved = [];
  for (const target of allowedForFile) {
    const stillBroken = broken.some((b) => b.raw === target);
    if (!stillBroken) allowedResolved.push(target);
  }
  if (real.length === 0 && allowedResolved.length === 0) {
    const note = allowed.length > 0 ? ` (${allowed.length} allow-listed)` : "";
    console.log(`✅ ${t}: ${total} link(s) OK${note}`);
    return 0;
  }
  if (real.length > 0) {
    console.log(`❌ ${t}: ${real.length} broken link(s)`);
    for (const b of real) console.log(`     ↳ ${b.raw}  →  ${b.resolved}`);
  }
  if (allowedResolved.length > 0) {
    console.log(`❌ ${t}: ${allowedResolved.length} allow-listed link(s) now resolve — remove from allowlist:`);
    for (const a of allowedResolved) console.log(`     ↳ ${a}`);
  }
  return real.length + allowedResolved.length;
}

function main() {
  const targets = process.argv.slice(2);
  if (targets.length === 0) {
    console.error("usage: check-doc-links.mjs <file> [<file> …]");
    process.exit(2);
  }
  const allowlist = loadAllowlist(resolve(REPO_ROOT, "scripts/docs/doc-links.allowlist"));
  let totalFailures = 0;
  for (const t of targets) {
    const abs = resolve(REPO_ROOT, t);
    if (!existsSync(abs)) {
      console.error(`❌ ${t}: file not found`);
      totalFailures++;
      continue;
    }
    const { total, broken } = checkFile(abs);
    const allowedForFile = allowlist.get(t) ?? new Set();
    totalFailures += reportFile(t, allowedForFile, broken, total);
  }
  console.log("");
  console.log(totalFailures === 0 ? "✅ all links resolve (allowlist respected)" : `❌ ${totalFailures} link issue(s)`);
  process.exit(totalFailures === 0 ? 0 : 1);
}

main();
