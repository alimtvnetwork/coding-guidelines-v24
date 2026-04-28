#!/usr/bin/env node
// Sync version.json from package.json + spec/ frontmatter.
//
// Output schema (extended manifest, backwards-compatible — `version`,
// `updated`, `name`, `description` remain at the root for the installer
// script that reads them with `cat version.json | python3 -c ...`):
//
//   {
//     version, updated, name, description,        // root summary
//     git: { sha, shortSha, branch },             // build provenance
//     stats: { totalFiles, totalLines, totalFolders, totalBytes },
//     folders: [
//       {
//         path: "01-spec-authoring-guide",
//         name: "Spec Authoring Guide",
//         title: "Spec Authoring Guide",          // from H1 in 00-overview
//         version: "3.1.0" | null,                // from frontmatter
//         updated: "2026-04-16" | null,
//         status: "Active" | null,
//         aiConfidence: "Production-Ready" | null,
//         ambiguity: "None" | null,
//         fileCount, lineCount, byteCount         // recursive
//       }
//     ]
//   }
//
// Run: node scripts/sync-version.mjs
// Wired into: `sync`, `prebuild`, husky pre-commit.

import { execSync } from "node:child_process";
import { readFileSync, writeFileSync, readdirSync, statSync } from "node:fs";
import { resolve, join, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const ROOT = resolve(__dirname, "..");
const SPEC_ROOT = resolve(ROOT, "spec");
const VERSION_PATH = resolve(ROOT, "version.json");
const PKG_PATH = resolve(ROOT, "package.json");
const AUTHORS_PATH = resolve(ROOT, "version-authors.json");

// ---------- canonical Role enum (spec/01-spec-authoring-guide/17-version-schema.md §6) ----------

const VALID_ROLES = new Set([
  "PrimaryAuthor",
  "Contributor",
  "Maintainer",
  "Reviewer",
  "Sponsor",
]);

// ---------- file IO ----------

function readJson(path) {
  return JSON.parse(readFileSync(path, "utf8"));
}

function writeJson(path, data) {
  writeFileSync(path, JSON.stringify(data, null, 2) + "\n");
}

function readText(path) {
  return readFileSync(path, "utf8");
}

// ---------- date (UTC+8 per user prefs) ----------

function todayUtc8() {
  const now = new Date(Date.now() + 8 * 60 * 60 * 1000);
  return now.toISOString().slice(0, 10);
}

// ---------- git provenance ----------

function tryGit(args) {
  try {
    return execSync(`git ${args}`, { cwd: ROOT, stdio: ["ignore", "pipe", "ignore"] })
      .toString()
      .trim();
  } catch {
    return null;
  }
}

function readGitInfo() {
  const sha = tryGit("rev-parse HEAD");
  return {
    sha: sha,
    shortSha: sha ? sha.slice(0, 7) : null,
    branch: tryGit("rev-parse --abbrev-ref HEAD"),
  };
}

// ---------- repo identity (PascalCase schema) ----------

// Strip embedded credentials (https://user:token@host/...) — never write
// secrets into version.json.
function stripUrlCredentials(url) {
  if (!url) return null;
  return url.replace(/^(https?:\/\/)[^@/]+@/, "$1");
}

function deriveRepoSlug(repoUrl, fallbackName) {
  if (!repoUrl) return fallbackName;
  const m = repoUrl.match(/\/([^/]+?)(?:\.git)?\/?$/);
  if (!m) return fallbackName;
  return m[1];
}

function loadAuthorsConfig() {
  try {
    return readJson(AUTHORS_PATH);
  } catch {
    return {};
  }
}

function validateAuthors(authors) {
  if (!Array.isArray(authors) || authors.length === 0) return [];
  const primaryCount = authors.filter((a) => a.Role === "PrimaryAuthor").length;
  if (primaryCount !== 1) {
    console.warn(
      `  WARN version-authors.json: expected exactly one PrimaryAuthor, found ${primaryCount}`,
    );
  }
  for (const a of authors) {
    if (!VALID_ROLES.has(a.Role)) {
      console.warn(`  WARN version-authors.json: invalid Role '${a.Role}' for '${a.Name}'`);
    }
  }
  return authors;
}

function buildPascalIdentity(pkg, gitInfo) {
  const cfg = loadAuthorsConfig();
  const remoteUrlRaw = tryGit("config --get remote.origin.url");
  const repoUrl = cfg.RepoUrl || stripUrlCredentials(remoteUrlRaw) || "";
  const slug = deriveRepoSlug(repoUrl, pkg.name || "unknown-repo");
  return {
    Version:       pkg.version,
    Title:         cfg.Title || pkg.name || slug,
    RepoSlug:      slug,
    RepoUrl:       repoUrl,
    LastCommitSha: gitInfo.sha || "",
    Description:   cfg.Description || pkg.description || "",
    Authors:       validateAuthors(cfg.Authors || []),
  };
}

// ---------- frontmatter parsing ----------

// Match `**Key:** Value` or `> **Key:** Value` (blockquote variant).
// Trailing two-space line break is tolerated.
function buildKeyMatcher(key) {
  return new RegExp(
    `^\\s*>?\\s*\\*\\*${key}:\\*\\*\\s*(.+?)\\s*$`,
    "im",
  );
}

function extractField(text, ...keyAliases) {
  for (const key of keyAliases) {
    const m = text.match(buildKeyMatcher(key));
    if (m) return m[1].replace(/\s+$/, "").trim();
  }
  return null;
}

function extractTitle(text) {
  // First H1 — strip "# 08 — " style numeric prefix.
  const h1 = text.match(/^#\s+(.+?)\s*$/m);
  if (!h1) return null;
  return h1[1]
    .replace(/^\d+\s*[-—–]\s*/, "")
    .replace(/\s+—\s+Overview\s*$/i, "")
    .trim();
}

function parseOverview(absPath) {
  let text;
  try { text = readText(absPath); } catch { return {}; }

  return {
    title:        extractTitle(text),
    version:      extractField(text, "Version", "Spec Version"),
    updated:      extractField(text, "Updated"),
    status:       extractField(text, "Status"),
    aiConfidence: extractField(text, "AI Confidence"),
    ambiguity:    extractField(text, "Ambiguity"),
  };
}

// ---------- folder stats ----------

function isVisible(name) {
  if (name.startsWith(".")) return false;
  if (name === "node_modules") return false;
  return true;
}

function walkStats(absPath) {
  let files = 0;
  let lines = 0;
  let bytes = 0;

  const stack = [absPath];
  while (stack.length > 0) {
    const cur = stack.pop();
    let entries;
    try { entries = readdirSync(cur); } catch { continue; }

    for (const entry of entries) {
      if (!isVisible(entry)) continue;
      const child = join(cur, entry);
      let st;
      try { st = statSync(child); } catch { continue; }

      if (st.isDirectory()) { stack.push(child); continue; }
      if (!entry.endsWith(".md")) continue;

      files += 1;
      bytes += st.size;
      try {
        const content = readText(child);
        lines += content.split("\n").length;
      } catch { /* skip unreadable */ }
    }
  }
  return { fileCount: files, lineCount: lines, byteCount: bytes };
}

// ---------- folder name derivation ----------

function specialAcronyms(s) {
  return s
    .replace(/\bCicd\b/g, "CICD")
    .replace(/\bWp\b/g,   "WP")
    .replace(/\bUi\b/g,   "UI")
    .replace(/\bDb\b/g,   "DB")
    .replace(/\bApi\b/g,  "API")
    .replace(/\bCli\b/g,  "CLI");
}

function deriveDisplayName(slug) {
  const stripped = slug.replace(/^\d+[-_]/, "");
  const titled = stripped
    .split(/[-_]/)
    .filter(Boolean)
    .map((w) => w.charAt(0).toUpperCase() + w.slice(1))
    .join(" ");
  return specialAcronyms(titled);
}

// ---------- folder discovery ----------

function listSpecFolders() {
  let entries;
  try { entries = readdirSync(SPEC_ROOT); } catch { return []; }

  const folders = [];
  for (const name of entries) {
    if (!isVisible(name)) continue;
    if (!/^\d+[-_]/.test(name)) continue;        // only NN-* folders
    // Exclude archived sibling-repo snapshots — these are not first-class
    // spec folders and should not appear in version.json or stats.
    if (name === "26-spec-outsides") continue;
    const abs = join(SPEC_ROOT, name);
    let st;
    try { st = statSync(abs); } catch { continue; }
    if (!st.isDirectory()) continue;
    folders.push({ slug: name, abs });
  }

  // Numeric sort: 02 < 09 < 10.
  folders.sort((a, b) =>
    a.slug.localeCompare(b.slug, undefined, { numeric: true }),
  );
  return folders;
}

function buildFolderEntry({ slug, abs }) {
  const overview = parseOverview(join(abs, "00-overview.md"));
  const stats = walkStats(abs);
  return {
    path:         slug,
    name:         deriveDisplayName(slug),
    title:        overview.title || null,
    version:      overview.version || null,
    updated:      overview.updated || null,
    status:       overview.status || null,
    aiConfidence: overview.aiConfidence || null,
    ambiguity:    overview.ambiguity || null,
    fileCount:    stats.fileCount,
    lineCount:    stats.lineCount,
    byteCount:    stats.byteCount,
  };
}

function buildFolderEntries() {
  return listSpecFolders().map(buildFolderEntry);
}

// ---------- aggregate stats ----------

function rootFileStats() {
  // Top-level .md files in spec/ that don't belong to any NN-* folder
  // (e.g. spec/00-overview.md, spec/spec-index.md, spec/health-dashboard.md).
  let entries;
  try { entries = readdirSync(SPEC_ROOT); } catch { return { fileCount: 0, lineCount: 0, byteCount: 0 }; }

  let files = 0, lines = 0, bytes = 0;
  for (const name of entries) {
    if (!isVisible(name)) continue;
    if (!name.endsWith(".md")) continue;
    const abs = join(SPEC_ROOT, name);
    let st;
    try { st = statSync(abs); } catch { continue; }
    if (!st.isFile()) continue;

    files += 1;
    bytes += st.size;
    try { lines += readText(abs).split("\n").length; } catch { /* skip */ }
  }
  return { fileCount: files, lineCount: lines, byteCount: bytes };
}

function sumStats(folders, rootStats) {
  const folderFiles = folders.reduce((a, f) => a + f.fileCount, 0);
  const folderLines = folders.reduce((a, f) => a + f.lineCount, 0);
  const folderBytes = folders.reduce((a, f) => a + f.byteCount, 0);
  return {
    totalFiles:   folderFiles + rootStats.fileCount,
    totalLines:   folderLines + rootStats.lineCount,
    totalFolders: folders.length,
    totalBytes:   folderBytes + rootStats.byteCount,
    rootFiles:    rootStats.fileCount,
  };
}

// ---------- build & write manifest ----------

function loadExistingMeta() {
  try {
    const v = readJson(VERSION_PATH);
    return { name: v.name, description: v.description };
  } catch {
    return {};
  }
}

function shouldEmitLegacy() {
  const flag = process.argv.includes("--no-legacy");
  const env = process.env.SYNC_VERSION_NO_LEGACY === "1";
  return flag === false && env === false;
}

function buildPascalSection(identity) {
  // §4 PascalCase identity (canonical, always emitted).
  return {
    Version:       identity.Version,
    Title:         identity.Title,
    RepoSlug:      identity.RepoSlug,
    RepoUrl:       identity.RepoUrl,
    LastCommitSha: identity.LastCommitSha,
    Description:   identity.Description,
    Authors:       identity.Authors,
  };
}

function buildLegacySection(pkg, existing, gitInfo, folders, rootStats) {
  // §10 transitional camelCase block — emitted only while legacy
  // readers (installer, sync-check, dashboard) still consume it.
  return {
    version:     pkg.version,
    updated:     todayUtc8(),
    name:        existing.name || "coding-guidelines",
    description:
      existing.description ||
      "Cross-language coding standards, error handling, CI/CD, and self-update specifications.",
    git:         gitInfo,
    stats:       sumStats(folders, rootStats),
    folders:     folders,
  };
}

function buildManifest() {
  const pkg = readJson(PKG_PATH);
  const existing = loadExistingMeta();
  const folders = buildFolderEntries();
  const rootStats = rootFileStats();
  const gitInfo = readGitInfo();
  const identity = buildPascalIdentity(pkg, gitInfo);

  const pascal = buildPascalSection(identity);
  if (shouldEmitLegacy() === false) return pascal;

  const legacy = buildLegacySection(pkg, existing, gitInfo, folders, rootStats);
  return { ...pascal, ...legacy };
}

function logSummary(manifest) {
  const v = manifest.Version;
  const sha = manifest.LastCommitSha ? manifest.LastCommitSha.slice(0, 7) : "no-git";
  const legacy = manifest.folders ? "with-legacy" : "pascal-only";
  console.log(`  OK version.json synced -> v${v} (git ${sha}, ${legacy})`);
  if (manifest.stats) {
    const s = manifest.stats;
    console.log(
      `     ${s.totalFolders} folders, ${s.totalFiles} files, ${s.totalLines.toLocaleString()} lines, ${(s.totalBytes / 1024).toFixed(1)} KB`,
    );
  }
  if (manifest.folders) {
    const versioned = manifest.folders.filter((f) => f.version).length;
    const missing = manifest.folders.length - versioned;
    if (missing > 0) {
      console.log(`     !! ${missing} folder(s) missing **Version:** in 00-overview.md`);
    }
  }
}

function main() {
  const manifest = buildManifest();
  writeJson(VERSION_PATH, manifest);
  logSummary(manifest);
}

main();
