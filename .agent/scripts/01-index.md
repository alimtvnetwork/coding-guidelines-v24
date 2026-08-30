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

### `08-fast-file-scanner.py`
* **Purpose:** High-performance repository file scanner and cache indexer. Scans thousands of files in milliseconds with language filters (`--lang go,ts,py`), path filters (`--path spec/`), and substring search (`--search`). Automatically persists results to `tmp/repo-file-cache.json` and `tmp/repo-file-list.txt` for instant cached access in subsequent steps without ad-hoc filesystem queries.
* **Usage:** `python .lovable/ai-fix-scripts/08-fast-file-scanner.py [--lang <lang>] [--path <dir>] [--search <term>] [--stats]`
* **Tags:** `#file-scanner`, `#file-indexer`, `#fast-search`, `#cache`, `#files`, `#performance`

### `09-fast-cached-grep.py`
* **Purpose:** High-speed parallel content matcher and regex grepper. Reads the pre-computed file cache from `tmp/repo-file-cache.json` and searches repository contents in parallel via ThreadPoolExecutor, outputting structured JSON results to `tmp/grep-results.json`.
* **Usage:** `python .lovable/ai-fix-scripts/09-fast-cached-grep.py --pattern <pattern> [--regex] [--lang <lang>] [--path <dir>] [--limit <N>]`
* **Tags:** `#grep`, `#pattern-search`, `#regex`, `#parallel-grep`, `#cached-search`, `#fast-query`


