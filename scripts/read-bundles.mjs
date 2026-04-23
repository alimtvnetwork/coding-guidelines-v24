#!/usr/bin/env node
// =====================================================================
// read-bundles.mjs — small CLI that emits shell-friendly bundle data.
//
// Used by .github/workflows/release.yml so the workflow doesn't have
// to duplicate the bundle list. Driven by bundles.json.
//
// Usage:
//   node scripts/read-bundles.mjs names
//       → newline-separated bundle names
//   node scripts/read-bundles.mjs folders <name>
//       → space-separated src folders for the given bundle
//   node scripts/read-bundles.mjs stable-name <name>
//       → archive.stableName for the given bundle
//   node scripts/read-bundles.mjs versioned-prefix <name>
//       → archive.versionedPrefix for the given bundle
//   node scripts/read-bundles.mjs prebuilt-src <name>
//       → prebuilt.src (or empty when bundle has no prebuilt)
//   node scripts/read-bundles.mjs prebuilt-dest <name>
//       → prebuilt.dest (or empty)
//   node scripts/read-bundles.mjs prebuilt-build <name>
//       → prebuilt.build shell command (or empty)
//   node scripts/read-bundles.mjs auto-open <name>
//       → autoOpen.entry path (or empty)
// =====================================================================

import { readFileSync } from "node:fs";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const REPO_ROOT = resolve(dirname(fileURLToPath(import.meta.url)), "..");
const MANIFEST = JSON.parse(readFileSync(resolve(REPO_ROOT, "bundles.json"), "utf8"));

const [, , command, name] = process.argv;

function findBundle(n) {
  const b = MANIFEST.bundles.find((x) => x.name === n);
  if (!b) {
    console.error(`Unknown bundle: ${n}`);
    process.exit(1);
  }
  return b;
}

switch (command) {
  case "names":
    console.log(MANIFEST.bundles.map((b) => b.name).join("\n"));
    break;
  case "folders":
    console.log(findBundle(name).folders.map((f) => f.src).join(" "));
    break;
  case "stable-name":
    console.log(findBundle(name).archive.stableName);
    break;
  case "versioned-prefix":
    console.log(findBundle(name).archive.versionedPrefix);
    break;
  case "prebuilt-src":
    console.log(findBundle(name).prebuilt?.src ?? "");
    break;
  case "prebuilt-dest":
    console.log(findBundle(name).prebuilt?.dest ?? "");
    break;
  case "prebuilt-build":
    console.log(findBundle(name).prebuilt?.build ?? "");
    break;
  case "auto-open":
    console.log(findBundle(name).autoOpen?.entry ?? "");
    break;
  default:
    console.error("Usage: read-bundles.mjs <names|folders|stable-name|versioned-prefix|prebuilt-src|prebuilt-dest|prebuilt-build|auto-open> [name]");
    process.exit(2);
}