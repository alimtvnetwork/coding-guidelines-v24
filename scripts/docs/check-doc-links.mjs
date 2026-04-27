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

function collectLinks(filePath) {
  const text = readFileSync(filePath, "utf8");
  const fileDir = dirname(filePath);
  const results = [];
  for (const match of text.matchAll(LINK_RE)) {
    const raw = match[1];
    if (isExternal(raw)) continue;
    const path = stripAnchor(raw);
    if (path === "" || path === ".") continue;
    const base = isAbsolute(path) ? REPO_ROOT : fileDir;
    const cleaned = isAbsolute(path) ? path.replace(/^\/+/, "") : path;
    const resolved = resolve(base, cleaned);
    results.push({ raw, resolved });
  }
  return results;
}

function checkFile(filePath) {
  const links = collectLinks(filePath);
  const broken = links.filter((l) => !existsSync(l.resolved));
  return { filePath, total: links.length, broken };
}

function main() {
  const targets = process.argv.slice(2);
  if (targets.length === 0) {
    console.error("usage: check-doc-links.mjs <file> [<file> …]");
    process.exit(2);
  }
  let totalBroken = 0;
  for (const t of targets) {
    const abs = resolve(REPO_ROOT, t);
    if (!existsSync(abs)) {
      console.error(`❌ ${t}: file not found`);
      totalBroken++;
      continue;
    }
    const { total, broken } = checkFile(abs);
    if (broken.length === 0) {
      console.log(`✅ ${t}: ${total} link(s) OK`);
      continue;
    }
    console.log(`❌ ${t}: ${broken.length} of ${total} link(s) broken`);
    for (const b of broken) {
      console.log(`     ↳ ${b.raw}  →  ${b.resolved}`);
    }
    totalBroken += broken.length;
  }
  console.log("");
  console.log(totalBroken === 0 ? "✅ all links resolve" : `❌ ${totalBroken} broken link(s)`);
  process.exit(totalBroken === 0 ? 0 : 1);
}

main();
