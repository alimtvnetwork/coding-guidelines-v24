#!/usr/bin/env node
// ci-config.mjs — Zero-dep loader for ci-guards.yaml (or .json).
//
// Validates the unified config schema documented in
//   spec/12-cicd-pipeline-workflows/03-reusable-ci-guards/08-config-schema.md
// and emits one of three output formats:
//   --emit env      shell `export VAR=...` lines (default)
//   --emit json     resolved config as JSON
//   --emit summary  human-readable per-guard table
//
// Designed to be sourced by ci-runner.sh:
//   eval "$(node scripts/ci-config.mjs --config ci-guards.yaml --emit env)"

import { readFileSync, existsSync } from "node:fs";
import { extname, resolve } from "node:path";
import process from "node:process";

const EXIT_OK = 0;
const EXIT_VIOLATION = 1;
const EXIT_TOOL_ERROR = 2;
const EXIT_USAGE_ERROR = 64;

const SUPPORTED_LANGUAGES = ["go", "node", "python", "rust", "polyglot"];
const KNOWN_GUARDS = [
  "forbidden_names",
  "naming_baseline",
  "collisions",
  "lint_diff",
  "lint_suggest",
  "test_summary",
];

function logError(msg) {
  process.stderr.write(`::error::[ci-config] ${msg}\n`);
}

function printUsage() {
  process.stdout.write(`Usage: ci-config.mjs --config <path> [--emit env|json|summary]

Loads ci-guards.yaml (or .json) and emits resolved settings.

Options:
  --config <path>   Path to ci-guards.yaml or .json (required)
  --emit <format>   Output format: env (default) | json | summary
  --help            Show this message

Exit codes:
   0  Config valid, output emitted.
   1  Config invalid (validation errors printed).
   2  Tool error (file unreadable, parse failure).
  64  Usage error (bad flags).
`);
}

function parseFlags(argv) {
  const flags = { config: "", emit: "env" };
  for (let i = 2; i < argv.length; i++) {
    const arg = argv[i];
    if (arg === "--config") flags.config = argv[++i] || "";
    else if (arg === "--emit") flags.emit = argv[++i] || "";
    else if (arg === "--help" || arg === "-h") {
      printUsage();
      process.exit(EXIT_OK);
    } else {
      logError(`unknown flag: ${arg}`);
      printUsage();
      process.exit(EXIT_USAGE_ERROR);
    }
  }
  return flags;
}

// ---------------------------------------------------------------------------
// Minimal YAML subset parser
//
// Supports exactly what the schema needs:
//   - top-level keys (key: value)
//   - nested maps (2-space indent)
//   - scalar values: strings (quoted or bare), integers, booleans
//   - sequences ("- item" lines under a key)
//   - "#" line comments
//
// Does NOT support: anchors, aliases, multi-line scalars, flow style,
// tags, or merge keys. The schema is intentionally simple to keep the
// parser auditable.
// ---------------------------------------------------------------------------

function parseYamlSubset(text) {
  const lines = text.split(/\r?\n/);
  const root = {};
  // Stack frames: { indent, container, parent, parentKey }
  //   container       = the map currently being filled
  //   parent + parentKey = where to swap container to an array if a list
  //                        item is encountered as the first child
  const stack = [{ indent: -1, container: root, parent: null, parentKey: null }];

  for (let lineNo = 0; lineNo < lines.length; lineNo++) {
    const raw = lines[lineNo];
    const stripped = stripComment(raw);
    if (!stripped.trim()) continue;

    const indent = countIndent(stripped);
    const body = stripped.slice(indent);

    while (stack.length > 1 && stack[stack.length - 1].indent >= indent) {
      stack.pop();
    }

    if (body.startsWith("- ")) {
      handleListItem(stack, body, lineNo, indent);
      continue;
    }

    handleMapEntry(stack, body, lineNo, indent);
  }

  return root;
}

function stripComment(line) {
  // Strip "# ..." comments not inside quotes. Schema has no '#' inside values.
  let inSingle = false;
  let inDouble = false;
  for (let i = 0; i < line.length; i++) {
    const ch = line[i];
    if (ch === "'" && !inDouble) inSingle = !inSingle;
    else if (ch === '"' && !inSingle) inDouble = !inDouble;
    else if (ch === "#" && !inSingle && !inDouble) return line.slice(0, i);
  }
  return line;
}

function countIndent(line) {
  let i = 0;
  while (i < line.length && line[i] === " ") i++;
  return i;
}

function handleListItem(stack, body, lineNo, _indent) {
  const top = stack[stack.length - 1];
  if (!top.parent || top.parentKey === null) {
    throw new Error(`line ${lineNo + 1}: list item without parent key`);
  }
  // Promote container to array on first list item.
  if (!Array.isArray(top.parent[top.parentKey])) {
    top.parent[top.parentKey] = [];
    top.container = top.parent[top.parentKey];
  }
  top.container.push(parseScalar(body.slice(2).trim()));
}

function handleMapEntry(stack, body, lineNo, indent) {
  const colonIdx = findKeyColon(body);
  if (colonIdx === -1) {
    throw new Error(`line ${lineNo + 1}: expected "key: value" got "${body}"`);
  }
  const key = body.slice(0, colonIdx).trim();
  const rest = body.slice(colonIdx + 1).trim();
  const top = stack[stack.length - 1];

  if (Array.isArray(top.container)) {
    throw new Error(`line ${lineNo + 1}: map entry inside list not supported`);
  }

  if (rest === "") {
    // Could become a map or an array depending on next line; start as map.
    top.container[key] = {};
    stack.push({
      indent,
      container: top.container[key],
      parent: top.container,
      parentKey: key,
    });
    return;
  }

  top.container[key] = parseScalar(rest);
}

function findKeyColon(body) {
  // Find the first ':' that is followed by space or end-of-line (YAML rule).
  for (let i = 0; i < body.length; i++) {
    if (body[i] === ":" && (i + 1 === body.length || body[i + 1] === " ")) {
      return i;
    }
  }
  return -1;
}

function parseScalar(raw) {
  if (raw === "" || raw === "~" || raw === "null") return null;
  if (raw === "true") return true;
  if (raw === "false") return false;
  if (/^-?\d+$/.test(raw)) return Number(raw);
  if (raw.startsWith('"') && raw.endsWith('"')) {
    return raw.slice(1, -1).replace(/\\"/g, '"').replace(/\\\\/g, "\\");
  }
  if (raw.startsWith("'") && raw.endsWith("'")) {
    return raw.slice(1, -1);
  }
  return raw;
}

// ---------------------------------------------------------------------------
// Loader
// ---------------------------------------------------------------------------

function loadConfig(path) {
  if (!existsSync(path)) {
    throw new Error(`config file not found: ${path}`);
  }
  const text = readFileSync(path, "utf-8");
  const ext = extname(path).toLowerCase();
  if (ext === ".json") return JSON.parse(text);
  if (ext === ".yaml" || ext === ".yml") return parseYamlSubset(text);
  throw new Error(`unsupported config extension: ${ext} (use .yaml/.yml/.json)`);
}

// ---------------------------------------------------------------------------
// Validation
// ---------------------------------------------------------------------------

function validateConfig(cfg) {
  const errors = [];
  if (typeof cfg.version !== "number") {
    errors.push("top-level 'version' must be an integer");
  }
  if (cfg.language && !SUPPORTED_LANGUAGES.includes(cfg.language)) {
    errors.push(
      `'language' must be one of ${SUPPORTED_LANGUAGES.join(", ")} (got '${cfg.language}')`
    );
  }
  for (const guard of KNOWN_GUARDS) {
    const section = cfg[guard];
    if (section === undefined) continue;
    if (typeof section !== "object" || Array.isArray(section)) {
      errors.push(`'${guard}' must be a map`);
      continue;
    }
    if (section.enabled !== undefined && typeof section.enabled !== "boolean") {
      errors.push(`'${guard}.enabled' must be true or false`);
    }
  }
  return errors;
}

// ---------------------------------------------------------------------------
// Emitters
// ---------------------------------------------------------------------------

function flattenForEnv(cfg) {
  const out = {};
  out.CI_GUARDS_VERSION = String(cfg.version ?? "");
  out.CI_GUARDS_LANGUAGE = cfg.language ?? "";
  out.CI_GUARDS_SCRIPTS_DIR = cfg.scripts_dir ?? ".github/scripts";

  for (const guard of KNOWN_GUARDS) {
    const section = cfg[guard];
    if (!section) continue;
    const prefix = `CI_GUARDS_${guard.toUpperCase()}`;
    for (const [key, value] of Object.entries(section)) {
      const envKey = `${prefix}_${key.toUpperCase()}`;
      out[envKey] = serializeForEnv(value);
    }
  }
  return out;
}

function serializeForEnv(value) {
  if (value === null || value === undefined) return "";
  if (typeof value === "boolean") return value ? "1" : "0";
  if (Array.isArray(value)) return value.join("\n");
  return String(value);
}

function shellQuote(value) {
  return `'${String(value).replace(/'/g, "'\\''")}'`;
}

function emitEnv(cfg) {
  const flat = flattenForEnv(cfg);
  return Object.entries(flat)
    .map(([k, v]) => `export ${k}=${shellQuote(v)}`)
    .join("\n");
}

function emitJson(cfg) {
  return JSON.stringify(cfg, null, 2);
}

function emitSummary(cfg) {
  const rows = ["Guard               Enabled  Notes"];
  rows.push("------------------- -------- -----------------------------------");
  for (const guard of KNOWN_GUARDS) {
    const section = cfg[guard];
    const enabled = section?.enabled === true ? "yes" : "no ";
    const notes = describeGuard(guard, section);
    rows.push(`${guard.padEnd(20)}${enabled.padEnd(9)}${notes}`);
  }
  return rows.join("\n");
}

function describeGuard(guard, section) {
  if (!section) return "(not configured)";
  if (guard === "forbidden_names") return `dir=${section.source_dir ?? "?"} patterns=${(section.patterns ?? []).length}`;
  if (guard === "naming_baseline") return `dir=${section.source_dir ?? "?"} baseline=${section.baseline_file ?? "?"}`;
  if (guard === "collisions") return `glob=${section.file_glob ?? "?"}`;
  if (guard === "lint_diff") return `normalizer=${section.normalizer ?? "?"}`;
  if (guard === "lint_suggest") return `table=${section.suggester_table ?? "?"}`;
  if (guard === "test_summary") return `dir=${section.results_dir ?? "?"}`;
  return "";
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

function main() {
  const flags = parseFlags(process.argv);
  if (!flags.config) {
    logError("--config <path> is required");
    printUsage();
    process.exit(EXIT_USAGE_ERROR);
  }

  let cfg;
  try {
    cfg = loadConfig(resolve(flags.config));
  } catch (err) {
    logError(err.message);
    process.exit(EXIT_TOOL_ERROR);
  }

  const errors = validateConfig(cfg);
  if (errors.length > 0) {
    for (const e of errors) logError(e);
    process.exit(EXIT_VIOLATION);
  }

  let output;
  if (flags.emit === "json") output = emitJson(cfg);
  else if (flags.emit === "summary") output = emitSummary(cfg);
  else if (flags.emit === "env") output = emitEnv(cfg);
  else {
    logError(`unknown --emit format: ${flags.emit}`);
    process.exit(EXIT_USAGE_ERROR);
  }

  process.stdout.write(output + "\n");
  process.exit(EXIT_OK);
}

main();