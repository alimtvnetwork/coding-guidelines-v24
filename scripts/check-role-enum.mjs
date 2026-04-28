#!/usr/bin/env node
// Verifies src/constants/roleEnum.ts is in sync with the spec.
// Re-parses the spec §6 Role Enum table, renders the expected file
// content in-memory, and compares it byte-for-byte to the committed
// file. Exits 1 with a clear message if drift is detected.
//
// Run: node scripts/check-role-enum.mjs
// Wired: package.json "check:role-enum"

import { readFileSync } from "node:fs";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const ROOT = resolve(__dirname, "..");
const SPEC = resolve(ROOT, "spec/01-spec-authoring-guide/17-version-schema.md");
const COMMITTED = resolve(ROOT, "src/constants/roleEnum.ts");

const SECTION_RE = /## §6 — `Role` Enum([\s\S]*?)(?=\n## §|\n---)/;
const ROW_RE = /^\|\s*`([A-Za-z]+)`\s*\|\s*([^|]+?)\s*\|/gm;

function readSpec() {
  return readFileSync(SPEC, "utf8");
}

function extractSection(text) {
  const match = text.match(SECTION_RE);
  if (match === null) throw new Error(`§6 Role Enum section not found in ${SPEC}`);
  return match[1];
}

function parseRoles(section) {
  const roles = [];
  let row;
  while ((row = ROW_RE.exec(section)) !== null) {
    roles.push({ name: row[1], description: row[2].trim() });
  }
  if (roles.length === 0) throw new Error("No Role rows parsed from §6 table");
  return roles;
}

function renderExpected(roles) {
  const lines = [
    "// AUTO-GENERATED — do not edit by hand.",
    "// Source: spec/01-spec-authoring-guide/17-version-schema.md §6",
    "// Regenerate: npm run gen:role-enum",
    "",
    "export const ROLE_VALUES = [",
    ...roles.map((r) => `  "${r.name}",`),
    "] as const;",
    "",
    "export type Role = (typeof ROLE_VALUES)[number];",
    "",
    "export const ROLE_DESCRIPTIONS: Record<Role, string> = {",
    ...roles.map((r) => `  ${r.name}: ${JSON.stringify(r.description)},`),
    "};",
    "",
    "const ROLE_SET: ReadonlySet<string> = new Set(ROLE_VALUES);",
    "",
    "export function isRole(value: unknown): value is Role {",
    "  return typeof value === \"string\" && ROLE_SET.has(value);",
    "}",
    "",
  ];
  return lines.join("\n");
}

function readCommitted() {
  return readFileSync(COMMITTED, "utf8");
}

function reportDrift(expected, actual) {
  console.error("✗ src/constants/roleEnum.ts is out of sync with the spec.");
  console.error("  Spec source: spec/01-spec-authoring-guide/17-version-schema.md §6");
  console.error("  Fix: run `npm run gen:role-enum` and commit the result.");
  console.error(`  Expected ${expected.length} bytes, found ${actual.length} bytes.`);
}

function main() {
  const roles = parseRoles(extractSection(readSpec()));
  const expected = renderExpected(roles);
  const actual = readCommitted();
  if (expected === actual) {
    console.log(`  OK roleEnum.ts matches spec (${roles.length} roles)`);
    return;
  }
  reportDrift(expected, actual);
  process.exit(1);
}

main();
