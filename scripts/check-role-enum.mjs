#!/usr/bin/env node
// Verifies src/constants/roleEnum.ts is in sync with the spec.
// Regenerates the file content in-memory and compares to disk.
// Exits 1 with a clear message if drift is detected.
//
// Run: node scripts/check-role-enum.mjs
// Wired: package.json "check:role-enum"

import { readFileSync, writeFileSync, mkdtempSync, rmSync } from "node:fs";
import { execFileSync } from "node:child_process";
import { resolve, dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { tmpdir } from "node:os";

const __dirname = dirname(fileURLToPath(import.meta.url));
const ROOT = resolve(__dirname, "..");
const GENERATOR = resolve(ROOT, "scripts/generate-role-enum.mjs");
const COMMITTED = resolve(ROOT, "src/constants/roleEnum.ts");

function readCommitted() {
  return readFileSync(COMMITTED, "utf8");
}

function generateExpected() {
  const tmp = mkdtempSync(join(tmpdir(), "role-enum-"));
  const shimOut = join(tmp, "roleEnum.ts");
  const wrapper = join(tmp, "run.mjs");
  const wrapperSrc = [
    "import { readFileSync, writeFileSync, mkdirSync } from 'node:fs';",
    "import { dirname } from 'node:path';",
    `const SPEC = ${JSON.stringify(resolve(ROOT, "spec/01-spec-authoring-guide/17-version-schema.md"))};`,
    `const OUT = ${JSON.stringify(shimOut)};`,
    `const SECTION_RE = /## §6 — \`Role\` Enum([\\s\\S]*?)(?=\\n## §|\\n---)/;`,
    `const ROW_RE = /^\\|\\s*\`([A-Za-z]+)\`\\s*\\|\\s*([^|]+?)\\s*\\|/gm;`,
    "const text = readFileSync(SPEC, 'utf8');",
    "const m = text.match(SECTION_RE); if (!m) throw new Error('§6 not found');",
    "const roles = []; let row; while ((row = ROW_RE.exec(m[1])) !== null) roles.push({ name: row[1], description: row[2].trim() });",
    "const lines = [",
    "  '// AUTO-GENERATED — do not edit by hand.',",
    "  '// Source: spec/01-spec-authoring-guide/17-version-schema.md §6',",
    "  '// Regenerate: npm run gen:role-enum',",
    "  '',",
    "  'export const ROLE_VALUES = [',",
    "  ...roles.map(r => `  \"${r.name}\",`),",
    "  '] as const;',",
    "  '',",
    "  'export type Role = (typeof ROLE_VALUES)[number];',",
    "  '',",
    "  'export const ROLE_DESCRIPTIONS: Record<Role, string> = {',",
    "  ...roles.map(r => `  ${r.name}: ${JSON.stringify(r.description)},`),",
    "  '};',",
    "  '',",
    "  'const ROLE_SET: ReadonlySet<string> = new Set(ROLE_VALUES);',",
    "  '',",
    "  'export function isRole(value: unknown): value is Role {',",
    "  '  return typeof value === \"string\" && ROLE_SET.has(value);',",
    "  '}',",
    "  '',",
    "].join('\\n');",
    "mkdirSync(dirname(OUT), { recursive: true });",
    "writeFileSync(OUT, lines);",
  ].join("\n");
  writeFileSync(wrapper, wrapperSrc);
  execFileSync(process.execPath, [wrapper], { stdio: "pipe" });
  const expected = readFileSync(shimOut, "utf8");
  rmSync(tmp, { recursive: true, force: true });
  return expected;
}

function reportDrift(expected, actual) {
  console.error("✗ src/constants/roleEnum.ts is out of sync with the spec.");
  console.error("  Spec source: spec/01-spec-authoring-guide/17-version-schema.md §6");
  console.error("  Fix: run `npm run gen:role-enum` and commit the result.");
  console.error("");
  console.error(`  Expected ${expected.length} bytes, found ${actual.length} bytes.`);
}

function main() {
  const expected = generateExpected();
  const actual = readCommitted();
  if (expected === actual) {
    console.log("  OK roleEnum.ts matches spec");
    return;
  }
  reportDrift(expected, actual);
  process.exit(1);
}

main();
