#!/usr/bin/env node
// Generates src/constants/roleEnum.ts from the canonical Role enum
// defined in spec/01-spec-authoring-guide/17-version-schema.md §6.
//
// The spec is the source of truth. This generator parses the §6 table,
// emits a TypeScript file with the Role enum, type, and helpers, and
// fails loudly if the spec table is missing or malformed.
//
// Run: node scripts/generate-role-enum.mjs
// Wired: package.json "gen:role-enum" + sync pipeline (optional).

import { readFileSync, writeFileSync, mkdirSync } from "node:fs";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const ROOT = resolve(__dirname, "..");
const SPEC = resolve(ROOT, "spec/01-spec-authoring-guide/17-version-schema.md");
const OUT = resolve(ROOT, "src/constants/roleEnum.ts");

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

function renderFile(roles) {
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

function writeOutput(content) {
  mkdirSync(dirname(OUT), { recursive: true });
  writeFileSync(OUT, content);
}

function main() {
  const text = readSpec();
  const section = extractSection(text);
  const roles = parseRoles(section);
  writeOutput(renderFile(roles));
  const names = roles.map((r) => r.name).join(", ");
  console.log(`  OK roleEnum.ts -> ${roles.length} roles: ${names}`);
}

main();
