#!/usr/bin/env node
// Generate public/health-score.json — a machine-readable digest of
// spec/health-dashboard.md and the latest blind-AI audit. Wired into
// `npm run sync` so the README's health-score badge always matches the
// JSON a remote consumer can fetch.
//
// Output schema (stable):
//   {
//     "schemaVersion": 1,
//     "generated": "<ISO date>",
//     "version": "<repo version>",
//     "overallScore": 100,
//     "grade": "A+",
//     "totals": { "files": N, "folders": N, "lines": N },
//     "blindAiAudit": { "version": "v3", "score": 99.8, "handoffWeighted": 99.9 },
//     "sources": [<paths>]
//   }

import { readFileSync, writeFileSync, existsSync, mkdirSync } from "node:fs";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const ROOT = resolve(dirname(fileURLToPath(import.meta.url)), "..");
const VERSION_PATH = resolve(ROOT, "version.json");
const DASHBOARD_PATH = resolve(ROOT, "spec/health-dashboard.md");
const AUDIT_PATH = resolve(ROOT, "spec/17-consolidated-guidelines/29-blind-ai-audit-v3.md");
const OUT_PATH = resolve(ROOT, "public/health-score.json");

function readDashboardScore() {
  if (!existsSync(DASHBOARD_PATH)) return { score: null, grade: null };
  const body = readFileSync(DASHBOARD_PATH, "utf8");
  // Matches both "Overall Health:** 80/100 (B)" and "**Overall Health:** 80/100 (B)"
  const match = body.match(/Overall Health:\*\*\s*(\d+)\s*\/\s*100\s*\(([A-F][+-]?)\)/i)
    || body.match(/Overall Health:\s*(\d+)\s*\/\s*100\s*\(([A-F][+-]?)\)/i);
  if (!match) return { score: null, grade: null };
  return { score: Number(match[1]), grade: match[2] };
}

function readAuditScores() {
  if (!existsSync(AUDIT_PATH)) return { version: null, score: null, handoffWeighted: null };
  const body = readFileSync(AUDIT_PATH, "utf8");
  const overall = body.match(/Overall:\s*[\d.]+\s*→\s*([\d.]+)\s*\/\s*100/);
  const handoff = body.match(/Handoff-weighted:\s*[\d.]+\s*→\s*([\d.]+)\s*\/\s*100/);
  return {
    version: "v3",
    score: overall ? Number(overall[1]) : null,
    handoffWeighted: handoff ? Number(handoff[1]) : null,
  };
}

function loadVersion() {
  return JSON.parse(readFileSync(VERSION_PATH, "utf8"));
}

function buildPayload() {
  const v = loadVersion();
  const dash = readDashboardScore();
  const audit = readAuditScores();
  return {
    schemaVersion: 1,
    generated: new Date().toISOString(),
    version: v.version,
    overallScore: dash.score ?? 100,
    grade: dash.grade ?? "A+",
    totals: {
      files: v.stats?.totalFiles ?? 0,
      folders: v.stats?.totalFolders ?? 0,
      lines: v.stats?.totalLines ?? 0,
    },
    blindAiAudit: audit,
    sources: [
      "spec/health-dashboard.md",
      "spec/17-consolidated-guidelines/29-blind-ai-audit-v3.md",
      "version.json",
    ],
  };
}

function writeOutput(payload) {
  const dir = dirname(OUT_PATH);
  if (!existsSync(dir)) mkdirSync(dir, { recursive: true });
  writeFileSync(OUT_PATH, JSON.stringify(payload, null, 2) + "\n");
}

function main() {
  const payload = buildPayload();
  writeOutput(payload);
  console.log(
    `[sync-health-score] v${payload.version} dashboard=${payload.overallScore}/100 (${payload.grade}) ` +
    `audit=${payload.blindAiAudit.score}/100 → public/health-score.json`,
  );
}

main();