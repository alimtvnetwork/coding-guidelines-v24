# AI Fix Scripts Index

> **/goal** Master and execute the repository's suite of ultra-fast Python scripts for linting, path resolution, naming enforcement, version synchronization, and local CI verification.
> **/learn** Read the script specifications below and run scripts via `python .lovable/ai-fix-scripts/<script-name>.py`.

## 🎯 Actionable CI/CD & Agent Checklist

1. [ ] `/goal` Verify local code and markdown compliance before submitting changes.
2. [ ] `/learn` Run all local verification linters via `python .lovable/ai-fix-scripts/06-cicd-local-runner.py`.
3. [ ] `/goal` Ensure all files use strict relative paths with zero absolute filesystem references.
4. [ ] `/learn` Scan repository files rapidly via `python .lovable/ai-fix-scripts/11-fast-file-scanner.py`.

---

**Version:** 3.3.0  
**Updated:** 2026-08-31  
**AI Confidence:** Production-Ready  
**Ambiguity:** None

---

## 🛠️ Script Catalog

| # | Script | Purpose | Execution Time |
|---|--------|---------|----------------|
| 01 | `01-index.md` | Master index, script catalog, and benchmark registry | — |
| 02 | `02-shared-engine.py` | Shared engine: thread-safe lazy regex registry, centralized configuration maps, dual-platform locks, pluggable `tmp/cache/`, two-phase streaming | ~2ms |
| 03 | `03-file-manipulator.py` | Mass lowercasing, sequence fixing, and UTF-8 LF normalization CLI | ~15ms |
| 04 | `04-newline-fixer.py` | Fixes trailing whitespace and missing final newlines across folders | ~15ms |
| 05 | `05-guideline-autofixer.py` | Composite runner combining newline fixing and boolean naming checks | ~25ms |
| 06 | `06-cicd-local-runner.py` | Runs all 18 CI quality checks locally via `ThreadPoolExecutor` | ~35ms |
| 07 | `07-relative-path-fixer.py` | Detects and fixes absolute paths / `file:///` URIs across folders | ~30ms |
| 08 | `08-naming-autofixer.py` | Enforces lowercase filenames, boolean conventions, and condition rules | ~20ms |
| 09 | `09-cli-help-auditor.py` | Validates CLI `--help` examples against actual implementations | ~25ms |
| 10 | `10-encoding-normalizer.py` | Normalizes all files to strict UTF-8 with UNIX LF line endings | ~35ms |
| 11 | `11-fast-file-scanner.py` | High-speed repo file scanner (<15ms full scan, <1ms cache query) | ~14ms |
| 12 | `12-fast-cached-grep.py` | Parallel regex matcher leveraging pre-warmed file cache | ~12ms |
| 13 | `13-file-size-guard.py` | Audits repository files for oversized binary blobs (>2MB) | ~10ms |
| 14 | `14-version-sync-checker.py` | Verifies synchronization of `version.json`, `package.json`, and `changelog.md` | ~5ms |
| 15 | `15-sequence-and-title-auditor.py` | Audits and aligns numeric file sequence prefixes and `# H1` titles | ~20ms |
| 16 | `16-installer-smoke-tester.py` | Generic installer smoke test validating script placeholders & hashes | ~8ms |
| 17 | `17-fast-file-reader.py` | AI agent fast file reader and folder explorer using `tmp/cache/` | <1ms |

---

## 📖 Detailed Script Specifications

<details>
<summary><b>02-shared-engine.py — Shared Core Engine & Lazy Regex Registry</b></summary>

- **Capabilities:**
  - `RegexPatternType` enum & `RegexRegistry.get(pattern_type)` thread-safe double-checked lock memoization.
  - Centralized global configurations: `EXCLUDE_DIRS`, `BINARY_EXTENSIONS`, `DEFAULT_TEXT_EXTENSIONS`, `DEFAULT_CODE_EXTENSIONS`, `DEFAULT_CLI_EXTENSIONS`, `LANG_EXT_MAP`, `DEFAULT_MAX_FILE_KB`, `ALLOWED_LARGE_FILES`.
  - Dual-platform locking: `fcntl.flock` on Unix and `os.O_CREAT | os.O_EXCL` on Windows with stale lock recovery.
  - Inode cycle tracking on Unix and executable bit (`st_mode`) preservation.
- **Ignore Pruning:** Recursively prunes `.git`, `.gitmap`, `node_modules`, `dist`, `build`, `.venv`, `.gemini`, `tmp`, `.system_generated`, and `release-artifacts`.
</details>

<details>
<summary><b>03-file-manipulator.py — Standalone File Manipulator CLI</b></summary>

- **Subcommands:**
  - `lowercase <path>`: Recursively converts filenames and directory names to lowercase.
  - `fix-seq-files <path>`: Sequentially numbers files with options for `--order-by-time`, `--order-by-az`, `--keep-old-order`, and `--pin`.
  - `fix-encoding <path>`: Normalizes files to UTF-8 without BOM and strict UNIX LF.
</details>

<details>
<summary><b>06-cicd-local-runner.py — Parallel Quality Gate Matrix Runner</b></summary>

- **Usage:** `python .lovable/ai-fix-scripts/06-cicd-local-runner.py`
- **Features:** Dispatches all 18 quality gates in parallel using 4 worker threads via `ThreadPoolExecutor`.
</details>

<details>
<summary><b>11-fast-file-scanner.py — Ultra-Fast File Scanner & Cache Indexer</b></summary>

- **Usage:** `python .lovable/ai-fix-scripts/11-fast-file-scanner.py --path <dir> --lang <languages> --ext <extensions>`
- **Cache Persistence:** Writes dual caches to `tmp/cache/repo-file-cache.json` and `tmp/repo-file-cache.json`.
- **Instant Query:** `python .lovable/ai-fix-scripts/11-fast-file-scanner.py --query-cache "<term>"` executes in <1ms without disk scanning.
</details>

<details>
<summary><b>12-fast-cached-grep.py — Parallel Cached Grepper</b></summary>

- **Usage:** `python .lovable/ai-fix-scripts/12-fast-cached-grep.py --pattern "<text>" --path <dir> --lang ts,go`
- **Thread-Safe Matcher:** Runs 8 parallel worker threads using pre-compiled regex patterns.
</details>

<details>
<summary><b>17-fast-file-reader.py — AI Agent Fast File Reader</b></summary>

- **Usage:**
  - `python .lovable/ai-fix-scripts/17-fast-file-reader.py --list-folder <dir> [--ext .md,.ts]`
  - `python .lovable/ai-fix-scripts/17-fast-file-reader.py --read-file <path>`
  - `python .lovable/ai-fix-scripts/17-fast-file-reader.py --search-pattern <text> [--path <dir>]`
</details>

---

## Usage Guidelines

- **AI Fast File/Folder Reading:** Run `python .lovable/ai-fix-scripts/17-fast-file-reader.py --list-folder <path>` or `--read-file <path>`.
- **Fast Content Search:** Run `python .lovable/ai-fix-scripts/12-fast-cached-grep.py --pattern "<text>" --lang ts,go`.
- **Local CI Testing:** Run `python .lovable/ai-fix-scripts/06-cicd-local-runner.py` before finalizing any PR or major edit batch.
