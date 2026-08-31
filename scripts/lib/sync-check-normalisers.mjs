// Normalisers for sync-check.mjs — strip volatile fields before diffing.

export function stripVolatileVersionFields(content) {
  let parsed;
  try { parsed = JSON.parse(content); } catch { return content; }
  delete parsed.LastCommitSha;
  delete parsed.git;
  delete parsed.updated;
  return JSON.stringify(parsed, null, 2) + "\n";
}

export function stripVolatileHealthScoreFields(content) {
  let parsed;
  try { parsed = JSON.parse(content); } catch { return content; }
  delete parsed.generated;
  return JSON.stringify(parsed, null, 2) + "\n";
}

export function stripVolatileReadmeStamps(content) {
  return content.replace(
    /<!--\s*UPDATED:start\s*-->[\s\S]*?<!--\s*UPDATED:end\s*-->/g,
    "<!-- UPDATED:start -->__NORMALISED__<!-- UPDATED:end -->",
  );
}

export function buildTrackedList() {
  return [
    { path: "version.json", normalise: stripVolatileVersionFields,
      note: "structural fields only — commit/date provenance ignored (legitimately differs per checkout / day)." },
    { path: "src/data/specTree.json" },
    { path: "public/health-score.json", normalise: stripVolatileHealthScoreFields,
      note: "the `generated` ISO timestamp is ignored." },
    { path: "readme.md", normalise: stripVolatileReadmeStamps,
      note: "the `<!-- UPDATED -->` date stamp is ignored." },
    { path: "docs/principles.md", normalise: stripVolatileReadmeStamps,
      note: "the `<!-- UPDATED -->` date stamp is ignored." },
    { path: "docs/author.md", normalise: stripVolatileReadmeStamps,
      note: "the `<!-- UPDATED -->` date stamp is ignored." },
  ];
}
