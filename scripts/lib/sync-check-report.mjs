// Drift report writer for sync-check.mjs (JSON + Markdown).

import { writeFileSync, mkdirSync, mkdtempSync } from "node:fs";
import { resolve, relative } from "node:path";
import { tmpdir } from "node:os";
import { spawnSync } from "node:child_process";

export function diffSummary(beforeText, afterText) {
  const before = beforeText.split("\n");
  const after = afterText.split("\n");
  let added = 0, removed = 0;
  const max = Math.max(before.length, after.length);
  for (let i = 0; i < max; i++) {
    if (before[i] === undefined) added++;
    else if (after[i] === undefined) removed++;
    else if (before[i] !== after[i]) { added++; removed++; }
  }
  return { added, removed, beforeLines: before.length, afterLines: after.length };
}

function unifiedDiff(a, b) {
  const r = spawnSync("diff", ["-u", a, b], { encoding: "utf8" });
  return r.stdout || "";
}

function buildFileEntries(drifted, tmp) {
  return drifted.map((item) => {
    const a = resolve(tmp, "before_" + item.file.path.replace(/[\\/]/g, "__"));
    const b = resolve(tmp, "after_" + item.file.path.replace(/[\\/]/g, "__"));
    writeFileSync(a, item.beforeText || "");
    writeFileSync(b, item.afterText || "");
    const summary = item.existedBefore && item.existsAfter
      ? diffSummary(item.beforeText, item.afterText)
      : { added: 0, removed: 0 };
    return {
      path: item.file.path,
      status: !item.existedBefore ? "created" : !item.existsAfter ? "removed" : "modified",
      added: summary.added,
      removed: summary.removed,
      note: item.file.note || null,
      unifiedDiff: item.existedBefore && item.existsAfter ? unifiedDiff(a, b) : null,
    };
  });
}

function renderMarkdown(json, files) {
  const md = [`# Sync drift report`, "", `- Mode: \`${json.mode}\``,
    `- Generated: ${json.generatedAt}`,
    `- Drifted: **${files.length}** of ${json.tracked} tracked file(s)`, ""];
  if (files.length === 0) { md.push("No drift detected."); return md.join("\n"); }
  md.push("| File | Status | +Lines | -Lines |", "| --- | --- | ---: | ---: |");
  for (const f of files) md.push(`| \`${f.path}\` | ${f.status} | ${f.added} | ${f.removed} |`);
  md.push("");
  for (const f of files) {
    md.push(`## \`${f.path}\``);
    if (f.note) md.push(`> ${f.note}`);
    md.push("");
    if (f.unifiedDiff) md.push("```diff", f.unifiedDiff.trimEnd(), "```");
    else md.push(`_${f.status === "created" ? "New file would be created" : "File would be removed"}._`);
    md.push("");
  }
  return md.join("\n");
}

export function writeReport({ reportDir, root, drifted, mode, tracked }) {
  if (!reportDir) return;
  mkdirSync(reportDir, { recursive: true });
  const tmp = mkdtempSync(resolve(tmpdir(), "sync-check-report-"));
  const files = buildFileEntries(drifted, tmp);
  const json = {
    mode, generatedAt: new Date().toISOString(),
    tracked, driftedCount: files.length, files,
  };
  writeFileSync(resolve(reportDir, "drift.json"), JSON.stringify(json, null, 2) + "\n");
  writeFileSync(resolve(reportDir, "drift.md"), renderMarkdown(json, files));
  process.stdout.write(`\nReport written to ${relative(root, reportDir)}/drift.{json,md}\n`);
}
