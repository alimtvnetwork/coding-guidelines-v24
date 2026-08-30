# AI Fix Scripts Index

This directory contains automated, reusable utility scripts designed for AI agents and developers to enforce repository conventions, normalize file encodings, validate relative paths, and execute quality gates.

## Available Scripts

### `02-newline-fixer.py`
* **Purpose:** Scans repository text and code files to ensure every file ends with exactly one UNIX-style newline (`\n`) and strips extraneous trailing whitespace.
* **Supported Extensions:** `.md`, `.txt`, `.go`, `.ts`, `.js`, `.mjs`, `.cjs`, `.jsx`, `.cs`, `.vb`, `.rs`, `.json`, `.yml`, `.yaml`, `.sh`, `.ps1`
* **Usage:** `python .lovable/ai-fix-scripts/02-newline-fixer.py`
* **Tags:** `#newline`, `#formatting`, `#whitespace`, `#normalization`, `#eol`

### `03-cicd-local-runner.py`
* **Purpose:** Local CI/CD pipeline orchestrator. Executes all local quality gates, linters, and smoke checks concurrently in parallel using Python `ThreadPoolExecutor`.
* **Usage:** `python .lovable/ai-fix-scripts/03-cicd-local-runner.py`
* **Tags:** `#cicd`, `#runner`, `#local-ci`, `#linter`, `#quality-gate`, `#parallel`

### `04-relative-path-fixer.py`
* **Purpose:** Scans markdown documents, plans, specs, and code files to resolve absolute filesystem paths (`/absolute/path/to/...`, `/absolute/path/to/...`, `/home/...`) and `file:///` URIs into clean, strictly relative Git paths.
* **Usage:** `python .lovable/ai-fix-scripts/04-relative-path-fixer.py`
* **Tags:** `#relative-paths`, `#path-fixer`, `#absolute-path-ban`, `#uri-fix`

### `05-naming-autofixer.py`
* **Purpose:** Audits and refactors naming conventions across TypeScript, Go, and Python files to enforce positive boolean prefixes (`is`, `has`), eliminate `Type` suffix violations, and standardize PascalCase/camelCase identifiers.
* **Usage:** `python .lovable/ai-fix-scripts/05-naming-autofixer.py`
* **Tags:** `#naming`, `#booleans`, `#enums`, `#refactoring`, `#conventions`

### `06-cli-help-auditor.py`
* **Purpose:** Validates and audits CLI help text, argument parsing flags, and usage documentation for consistency across commands.
* **Usage:** `python .lovable/ai-fix-scripts/06-cli-help-auditor.py`
* **Tags:** `#cli`, `#help`, `#flags`, `#argparse`, `#documentation`

### `07-encoding-normalizer.py`
* **Purpose:** Removes corrupt binary/control characters (such as `NUL`, `BEL`, `BS`, `VT`, `FF`, `ESC`) and normalizes all repository files to clean UTF-8 without Byte Order Marks (BOM).
* **Usage:** `python .lovable/ai-fix-scripts/07-encoding-normalizer.py`
* **Tags:** `#encoding`, `#utf8`, `#bom-remover`, `#control-characters`, `#text-cleanup`
