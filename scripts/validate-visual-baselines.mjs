#!/usr/bin/env node
/**
 * validate-visual-baselines.mjs
 *
 * Guards against the "silent partial baseline" class of bug: the visual
 * regression suite (`slides-app/tests/visual.spec.ts`) iterates
 * `DECK.length` slides, but the suite only actually protects a slide if a
 * committed baseline PNG exists under `slides-app/tests/visual.spec.ts-snapshots/`.
 * Missing PNGs cause Playwright to auto-create a baseline on first run
 * and pass, which looks green in CI while offering zero protection.
 *
 * Root cause of the class: baseline count and deck length can drift apart
 * silently (deck grew from 17 to 70 while baselines lagged). This script
 * fails fast when they disagree.
 *
 * Usage:
 *   node scripts/validate-visual-baselines.mjs            # verify
 *   node scripts/validate-visual-baselines.mjs --list     # list missing
 *
 * Exit codes:
 *   0  count matches, every expected PNG present
 *   1  drift detected (count mismatch or missing files)
 *   2  invocation error (deck unreadable, snapshots dir missing structurally)
 */
import { readdirSync, existsSync, readFileSync } from "node:fs";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const HERE = dirname(fileURLToPath(import.meta.url));
const REPO = resolve(HERE, "..");
const REGISTRY = resolve(REPO, "slides-app/src/deck/registry.ts");
const SNAPSHOT_DIR = resolve(
  REPO,
  "slides-app/tests/visual.spec.ts-snapshots",
);

function isMissingRegistry() {
  return !existsSync(REGISTRY);
}

function isMissingSnapshotDir() {
  return !existsSync(SNAPSHOT_DIR);
}

/**
 * Count DECK entries by parsing top-level object literals in the exported
 * array. We avoid importing the TS module so the check has zero build
 * deps and runs in <50 ms.
 */
function countDeckEntries() {
  const src = readFileSync(REGISTRY, "utf8");
  const marker = "export const DECK";
  const start = src.indexOf(marker);
  if (start < 0) return 0;
  const open = src.indexOf("[", start);
  if (open < 0) return 0;
  let depth = 0;
  let count = 0;
  for (let i = open; i < src.length; i += 1) {
    const ch = src[i];
    if (ch === "{") {
      if (depth === 0) count += 1;
      depth += 1;
    } else if (ch === "}") {
      depth -= 1;
    } else if (ch === "]" && depth === 0) {
      break;
    }
  }
  return count;
}

function listBaselinePngs() {
  return readdirSync(SNAPSHOT_DIR)
    .filter((name) => name.endsWith(".png"))
    .sort();
}

function expectedBaselineNames(deckCount) {
  const names = [];
  for (let i = 0; i < deckCount; i += 1) {
    names.push(`slide-${String(i).padStart(2, "0")}-chromium-linux.png`);
  }
  return names;
}

function main() {
  const wantList = process.argv.includes("--list");

  if (isMissingRegistry()) {
    process.stderr.write(
      `[validate-visual-baselines] registry not found: ${REGISTRY}\n`,
    );
    process.exit(2);
  }
  if (isMissingSnapshotDir()) {
    process.stderr.write(
      `[validate-visual-baselines] snapshots dir missing: ${SNAPSHOT_DIR}\n`,
    );
    process.exit(2);
  }

  const deckCount = countDeckEntries();
  const pngs = listBaselinePngs();
  const expected = expectedBaselineNames(deckCount);
  const present = new Set(pngs);
  const missing = expected.filter((name) => !present.has(name));

  const line = `[validate-visual-baselines] deck=${deckCount} baselines=${pngs.length} missing=${missing.length}`;
  process.stdout.write(`${line}\n`);

  const hasDrift = missing.length > 0 || pngs.length < deckCount;
  if (!hasDrift) {
    process.exit(0);
  }

  process.stderr.write(
    `[validate-visual-baselines] DRIFT: ${missing.length} baselines missing for ${deckCount} deck slides.\n`,
  );
  if (wantList) {
    for (const name of missing) process.stderr.write(`  - ${name}\n`);
  }
  process.stderr.write(
    "[validate-visual-baselines] Fix: run the slides-visual workflow via workflow_dispatch with update_baselines=true, then commit the artifact.\n",
  );
  process.exit(1);
}

main();
