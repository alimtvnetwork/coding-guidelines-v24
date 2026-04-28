#!/usr/bin/env node
// Validates version.json against the PascalCase schema defined in
// spec/01-spec-authoring-guide/17-version-schema.md.
//
// Exit codes:
//   0 — valid
//   1 — file missing or unparseable
//   2 — schema violations
//
// Usage:
//   node scripts/validate-version-json.mjs [path/to/version.json]

import { readFileSync, existsSync } from "node:fs";
import { resolve } from "node:path";

const REQUIRED_TOP_KEYS = [
  "Version",
  "Title",
  "RepoSlug",
  "RepoUrl",
  "LastCommitSha",
  "Description",
  "Authors",
];

const REQUIRED_AUTHOR_KEYS = ["Name", "Urls", "Role", "Background"];

const ROLE_ENUM = new Set([
  "PrimaryAuthor",
  "Contributor",
  "Maintainer",
  "Reviewer",
  "Sponsor",
]);

const SEMVER_RE = /^\d+\.\d+\.\d+(?:[-+][\w.-]+)?$/;
const SHA_RE = /^[0-9a-f]{40}$/;
const URL_RE = /^https?:\/\/\S+$/;

const isNonEmptyString = (value) =>
  typeof value === "string" && value.trim().length > 0;

const isStringArray = (value) =>
  Array.isArray(value) && value.every((entry) => typeof entry === "string");

function loadManifest(path) {
  const exists = existsSync(path);
  if (exists === false) {
    return { ok: false, error: `version.json not found at ${path}` };
  }
  const raw = readFileSync(path, "utf8");
  try {
    return { ok: true, data: JSON.parse(raw) };
  } catch (err) {
    return { ok: false, error: `version.json is not valid JSON: ${err.message}` };
  }
}

function checkTopLevelKeys(data, errors) {
  for (const key of REQUIRED_TOP_KEYS) {
    const present = Object.prototype.hasOwnProperty.call(data, key);
    if (present === false) errors.push(`Missing required key: ${key}`);
  }
}

function checkScalarFields(data, errors) {
  if (isNonEmptyString(data.Version) === false) errors.push("Version must be a non-empty string");
  else if (SEMVER_RE.test(data.Version) === false) errors.push(`Version is not SemVer: ${data.Version}`);

  if (isNonEmptyString(data.Title) === false) errors.push("Title must be a non-empty string");
  if (isNonEmptyString(data.RepoSlug) === false) errors.push("RepoSlug must be a non-empty string");
  if (isNonEmptyString(data.Description) === false) errors.push("Description must be a non-empty string");

  if (isNonEmptyString(data.RepoUrl) === false) errors.push("RepoUrl must be a non-empty string");
  else if (URL_RE.test(data.RepoUrl) === false) errors.push(`RepoUrl is not an http(s) URL: ${data.RepoUrl}`);

  if (isNonEmptyString(data.LastCommitSha) === false) errors.push("LastCommitSha must be a non-empty string");
  else if (SHA_RE.test(data.LastCommitSha) === false) errors.push(`LastCommitSha must be a 40-char hex SHA: ${data.LastCommitSha}`);
}

function checkAuthorShape(author, index, errors) {
  for (const key of REQUIRED_AUTHOR_KEYS) {
    const present = Object.prototype.hasOwnProperty.call(author, key);
    if (present === false) errors.push(`Authors[${index}] missing key: ${key}`);
  }
  if (isNonEmptyString(author.Name) === false) errors.push(`Authors[${index}].Name must be a non-empty string`);
  if (isNonEmptyString(author.Background) === false) errors.push(`Authors[${index}].Background must be a non-empty string`);
  if (isStringArray(author.Urls) === false) errors.push(`Authors[${index}].Urls must be an array of strings`);
  const roleValid = ROLE_ENUM.has(author.Role);
  if (roleValid === false) errors.push(`Authors[${index}].Role invalid: "${author.Role}" (allowed: ${[...ROLE_ENUM].join(", ")})`);
}

function checkAuthors(data, errors) {
  const authors = data.Authors;
  if (Array.isArray(authors) === false) {
    errors.push("Authors must be an array");
    return;
  }
  if (authors.length === 0) {
    errors.push("Authors must contain at least one entry");
    return;
  }
  authors.forEach((author, index) => checkAuthorShape(author, index, errors));
  const primaryCount = authors.filter((a) => a.Role === "PrimaryAuthor").length;
  if (primaryCount !== 1) errors.push(`Authors must contain exactly one PrimaryAuthor (found ${primaryCount})`);
}

function validate(data) {
  const errors = [];
  checkTopLevelKeys(data, errors);
  checkScalarFields(data, errors);
  checkAuthors(data, errors);
  return errors;
}

function reportAndExit(errors, path) {
  if (errors.length === 0) {
    console.log(`✓ version.json is valid: ${path}`);
    process.exit(0);
  }
  console.error(`✗ version.json failed schema validation: ${path}`);
  for (const msg of errors) console.error(`  - ${msg}`);
  console.error(`\n${errors.length} violation(s). See spec/01-spec-authoring-guide/17-version-schema.md`);
  process.exit(2);
}

function main() {
  const argPath = process.argv[2];
  const path = resolve(argPath ?? "version.json");
  const loaded = loadManifest(path);
  if (loaded.ok === false) {
    console.error(`✗ ${loaded.error}`);
    process.exit(1);
  }
  const errors = validate(loaded.data);
  reportAndExit(errors, path);
}

main();
