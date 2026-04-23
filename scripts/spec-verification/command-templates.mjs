/**
 * command-templates.mjs — Reusable verification command registry.
 *
 * Single source of truth for the shell commands that appear in every
 * injected `## Verification` block. Keying commands by name (instead of
 * inlining strings in profiles) gives us:
 *
 *   1. Consistency — every profile that needs "spec link check" emits the
 *      same command, so a future flag/path change is one edit.
 *   2. Discoverability — the registry is the canonical inventory of
 *      verification entry points; missing tooling is obvious.
 *   3. Composability — `composeCommands(['lint:spec-links', 'lint:forbidden'])`
 *      yields a single `&&`-joined command string.
 *
 * Adding a new template: append an entry, give it a stable key, document
 * what it asserts. Then reference the key from `profiles.mjs`.
 */

/**
 * @typedef {Object} CommandTemplate
 * @property {string} key       Stable identifier used by profiles.
 * @property {('lint'|'test'|'schema'|'build'|'security'|'meta')} category
 * @property {string} command   The exact shell command to run.
 * @property {string} asserts   Plain-language description of what success means.
 * @property {string} [requires] Optional precondition (e.g. installed tool).
 */

/** @type {Record<string, CommandTemplate>} */
export const COMMAND_TEMPLATES = {
  // ─── Lint family ───────────────────────────────────────────────────
  "lint:spec-links": {
    key: "lint:spec-links",
    category: "lint",
    command: "python3 linter-scripts/check-spec-cross-links.py --root spec --repo-root .",
    asserts: "Every internal markdown link in spec/ resolves to an existing file and heading anchor.",
  },
  "lint:spec-folders": {
    key: "lint:spec-folders",
    category: "lint",
    command: "python3 linter-scripts/check-spec-folder-refs.py",
    asserts: "Every spec folder has the required 00-overview.md and follows kebab-case numeric prefixes.",
  },
  "lint:forbidden-strings": {
    key: "lint:forbidden-strings",
    category: "lint",
    command: "python3 linter-scripts/check-forbidden-strings.py",
    asserts: "No forbidden tokens (createdAt, UUID, snake_case columns, legacy paths) appear anywhere in the repo.",
  },
  "lint:forbidden-paths": {
    key: "lint:forbidden-paths",
    category: "lint",
    command: "bash linter-scripts/check-forbidden-spec-paths.sh",
    asserts: "No re-split or merge-proposal placeholder paths leak into spec/.",
  },
  "lint:axios-version": {
    key: "lint:axios-version",
    category: "lint",
    command: "bash linter-scripts/check-axios-version.sh",
    asserts: "Axios is pinned to an exact safe version (1.14.0 or 0.30.3); blocked versions are rejected.",
  },
  "lint:guidelines-go": {
    key: "lint:guidelines-go",
    category: "lint",
    command: "go run linter-scripts/validate-guidelines.go --path spec --max-lines 15",
    asserts: "Zero CODE-RED violations: functions ≤ 15 lines, files ≤ 300 lines, no nested ifs, max 2 boolean operands.",
    requires: "go 1.22+",
  },
  "lint:guidelines-py": {
    key: "lint:guidelines-py",
    category: "lint",
    command: "python3 linter-scripts/validate-guidelines.py spec",
    asserts: "Python-side guideline validator agrees with the Go validator on every rule.",
  },
  "lint:eslint": {
    key: "lint:eslint",
    category: "lint",
    command: "npm run lint",
    asserts: "ESLint passes with zero errors against the project's TypeScript and React surface.",
  },
  "lint:powershell": {
    key: "lint:powershell",
    category: "lint",
    command: "pwsh -NoProfile -Command \"Invoke-ScriptAnalyzer -Path scripts -Recurse -Severity Warning\"",
    asserts: "PowerShell scripts pass PSScriptAnalyzer at Warning severity or higher.",
    requires: "pwsh + PSScriptAnalyzer",
  },

  // ─── Test family ───────────────────────────────────────────────────
  "test:unit": {
    key: "test:unit",
    category: "test",
    command: "npm run test",
    asserts: "Vitest unit suite passes with zero failures.",
  },
  "test:build": {
    key: "test:build",
    category: "test",
    command: "npm run build",
    asserts: "Production Vite build completes without errors or type failures.",
  },

  // ─── Schema family ─────────────────────────────────────────────────
  "schema:db": {
    key: "schema:db",
    category: "schema",
    command: "python3 linter-scripts/check-forbidden-strings.py",
    asserts: "DDL conforms to PascalCase singular table names, integer PKs, no forbidden tokens (createdAt/UUID/snake_case).",
  },
  "schema:spec-tree": {
    key: "schema:spec-tree",
    category: "schema",
    command: "node scripts/sync-spec-tree.mjs && git diff --exit-code -- src/data/specTree.json",
    asserts: "src/data/specTree.json is in sync with the on-disk spec/ tree.",
  },
  "schema:version": {
    key: "schema:version",
    category: "schema",
    command: "node scripts/sync-version.mjs",
    asserts: "version.json structural content matches package.json (volatile git/updated fields ignored).",
  },

  // ─── Meta — high-level composites ──────────────────────────────────
  "meta:spec-health": {
    key: "meta:spec-health",
    category: "meta",
    command: "python3 linter-scripts/check-spec-folder-refs.py && python3 linter-scripts/check-spec-cross-links.py --root spec --repo-root .",
    asserts: "Spec tree is structurally valid AND every internal cross-reference resolves.",
  },
  "meta:full-ci": {
    key: "meta:full-ci",
    category: "meta",
    command: "npm run sync && npm run lint && npm run test",
    asserts: "Sync drift, lint, and unit tests all pass — equivalent to the merge-blocking CI gate.",
  },
};

/**
 * Resolve a template key to its command string.
 * Throws if the key is unknown — fail loud so typos surface immediately.
 */
export function commandFor(key) {
  const tpl = COMMAND_TEMPLATES[key];
  if (!tpl) {
    throw new Error(`Unknown command template key: '${key}'. Known keys: ${Object.keys(COMMAND_TEMPLATES).join(", ")}`);
  }
  return tpl.command;
}

/**
 * Compose multiple templates into a single shell line joined by `&&`.
 * Useful when a profile needs both a lint and a schema check.
 */
export function composeCommands(keys) {
  return keys.map(commandFor).join(" && ");
}

/**
 * Group templates by category for documentation tables.
 */
export function listByCategory() {
  const out = {};
  for (const tpl of Object.values(COMMAND_TEMPLATES)) {
    if (!out[tpl.category]) out[tpl.category] = [];
    out[tpl.category].push(tpl);
  }
  return out;
}

/**
 * Verification-need → recommended template key map. Documents the
 * "common needs" ↔ "templates" mapping the user asked for.
 */
export const VERIFICATION_NEED_MAP = {
  "Internal links resolve":           "lint:spec-links",
  "Spec folder structure valid":      "lint:spec-folders",
  "No forbidden tokens / paths":      "lint:forbidden-strings",
  "DB schema conforms":               "schema:db",
  "Code-style guidelines (Go/Py)":    "lint:guidelines-go",
  "TypeScript/ESLint clean":          "lint:eslint",
  "PowerShell scripts clean":         "lint:powershell",
  "Axios pinned safely":              "lint:axios-version",
  "Unit tests pass":                  "test:unit",
  "Production build succeeds":        "test:build",
  "specTree.json in sync":            "schema:spec-tree",
  "version.json in sync":             "schema:version",
  "Spec health (structure + links)":  "meta:spec-health",
  "Full CI gate":                     "meta:full-ci",
};