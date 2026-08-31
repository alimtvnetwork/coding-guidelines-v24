# AI Fix Scripts Index & Tooling Guide

> /goal Master, discover, and execute the repository's suite of ultra-fast Python scripts for linting, path resolution, naming enforcement, version synchronization, and local CI verification.
> /learn Read the script specifications below and run scripts via `python .lovable/ai-fix-scripts/<script-name>.py`.

---

## 📋 AI Agent Pre-Flight Checklist

Follow this sequence before and during any repository modification task:

- [ ] **/learn** Inspect `02-shared-engine.py` to understand shared enums (`RegexPatternType`, `ScanModeType`, `SeverityType`, `ExitCodeType`), regex cache, and dual-platform locks.
- [ ] **/goal** Run rapid repo-wide file discovery using `11-fast-file-scanner.py` or instant cache lookup `<1ms`.
- [ ] **/goal** Rapidly read target files or explore folder contents using `17-fast-file-reader.py`.
- [ ] **/goal** Search multi-threaded regex patterns across files using `12-fast-cached-grep.py`.
- [ ] **/learn** Auto-fix whitespace, line endings, and boolean checks using `05-guideline-autofixer.py`.
- [ ] **/goal** Sanitize absolute filesystem paths and `file:///` URIs using `07-relative-path-fixer.py`.
- [ ] **/goal** Validate all 18 quality gates in parallel before submitting using `06-cicd-local-runner.py`.

---

## 🛠️ Master Script Catalog & Search Tags

| # | Script | Primary Purpose | Speed | Discovery Tags |
|:---:|---|---|:---:|---|
| **01** | `01-index.md` | Master index, script catalog, AI instructions, and tag registry | — | `docs`, `index`, `ai-instructions`, `catalog` |
| **02** | `02-shared-engine.py` | Shared engine: regex registry, centralized constants, dual-platform locks | ~2ms | `core`, `engine`, `regex`, `locking`, `cache`, `enums` |
| **03** | `03-file-manipulator.py` | Mass lowercasing, sequence fixing, and UTF-8 LF normalization CLI | ~15ms | `rename`, `lowercase`, `sequence`, `encoding`, `cli` |
| **04** | `04-newline-fixer.py` | Fixes trailing whitespace and missing final newlines across folders | ~15ms | `newlines`, `whitespace`, `crlf`, `lf`, `formatting` |
| **05** | `05-guideline-autofixer.py` | Composite runner combining newline fixing and boolean naming checks | ~25ms | `autofix`, `composite`, `guidelines`, `booleans` |
| **06** | `06-cicd-local-runner.py` | Runs all 18 CI quality checks locally via `ThreadPoolExecutor` | ~35ms | `ci-cd`, `runner`, `parallel`, `quality-gates`, `test` |
| **07** | `07-relative-path-fixer.py` | Detects and fixes absolute paths / `file:///` URIs across folders | ~30ms | `paths`, `relative-paths`, `absolute-paths`, `sanitizer` |
| **08** | `08-naming-autofixer.py` | Enforces lowercase filenames, boolean conventions, and condition rules | ~20ms | `naming`, `booleans`, `is-prefix`, `has-prefix`, `linter` |
| **09** | `09-cli-help-auditor.py` | Validates CLI `--help` examples against actual implementations | ~25ms | `cli`, `help`, `cobra`, `commander`, `docstrings` |
| **10** | `10-encoding-normalizer.py` | Normalizes all files to strict UTF-8 with UNIX LF line endings | ~35ms | `encoding`, `utf-8`, `bom-stripping`, `unix-lf` |
| **11** | `11-fast-file-scanner.py` | High-speed repo file scanner (<15ms full scan, <1ms cache query) | ~14ms | `scanner`, `cache`, `indexing`, `file-list`, `discovery` |
| **12** | `12-fast-cached-grep.py` | Parallel regex matcher leveraging pre-warmed file cache | ~12ms | `grep`, `search`, `regex`, `parallel`, `content-search` |
| **13** | `13-file-size-guard.py` | Audits repository files for oversized binary blobs (>2MB) | ~10ms | `file-size`, `blob-guard`, `security`, `binary-check` |
| **14** | `14-version-sync-checker.py` | Verifies synchronization of `version.json`, `package.json`, `changelog.md` | ~5ms | `version`, `sync`, `changelog`, `package-json`, `release` |
| **15** | `15-sequence-and-title-auditor.py` | Audits and aligns numeric file sequence prefixes and `# H1` titles | ~20ms | `sequence`, `title`, `h1-headers`, `markdown-audit` |
| **16** | `16-installer-smoke-tester.py` | Generic installer smoke test validating script placeholders & hashes | ~8ms | `installer`, `smoke-test`, `install-sh`, `install-ps1` |
| **17** | `17-fast-file-reader.py` | AI agent fast file reader and folder explorer using `tmp/cache/` | <1ms | `reader`, `explorer`, `instant-read`, `ai-tool` |

---

## 📖 Individual Script Usage & Specifications

### 02-shared-engine.py — Shared Core Engine & Lazy Regex Registry
- **Tags:** `core`, `engine`, `regex`, `locking`, `cache`, `enums`
- **Description:** Central foundation module providing centralized constants, `RegexPatternType` lazy compilation with thread-safe double-checked locking, dual-platform process locking (`fcntl.flock` on POSIX and `O_CREAT | O_EXCL` on Windows), Unix inode recursion guards, and two-phase cached file streaming.
- **Import Pattern:**
  ```python
  from importlib import import_module
  engine = import_module("02-shared-engine")
  process_repository_files = engine.process_repository_files
  read_file_safe = engine.read_file_safe
  write_file_lf = engine.write_file_lf
  ```

### 03-file-manipulator.py — Standalone File Manipulator CLI
- **Tags:** `rename`, `lowercase`, `sequence`, `encoding`, `cli`
- **Description:** Multi-purpose CLI tool for mass lowercasing of files/folders, sequential re-numbering of files, and UTF-8 LF normalization.
- **Commands:**
  - `python .lovable/ai-fix-scripts/03-file-manipulator.py lowercase <path> [--except <patterns>]`
  - `python .lovable/ai-fix-scripts/03-file-manipulator.py fix-seq-files <path> [--order-by-time|--order-by-az] [--pin "readme=00,intro=01"]`
  - `python .lovable/ai-fix-scripts/03-file-manipulator.py fix-encoding <path> [--ext .md,.ts,.py]`

### 04-newline-fixer.py — Fast Newline & Whitespace Fixer
- **Tags:** `newlines`, `whitespace`, `crlf`, `lf`, `formatting`
- **Description:** Audits and sanitizes trailing whitespace and ensures a single trailing newline across text files.
- **Commands:**
  - `python .lovable/ai-fix-scripts/04-newline-fixer.py <path> --fix`
  - `python .lovable/ai-fix-scripts/04-newline-fixer.py <path> --ext .md,.py`

### 05-guideline-autofixer.py — Composite Guideline Autofixer
- **Tags:** `autofix`, `composite`, `guidelines`, `booleans`
- **Description:** Forwarder script running newline normalization (`04-newline-fixer.py`) and implicit boolean convention auditing (`08-naming-autofixer.py`) in a single command.
- **Commands:**
  - `python .lovable/ai-fix-scripts/05-guideline-autofixer.py <path>`
  - `python .lovable/ai-fix-scripts/05-guideline-autofixer.py <path> --check-only`

### 06-cicd-local-runner.py — Parallel Quality Gate Matrix Runner
- **Tags:** `ci-cd`, `runner`, `parallel`, `quality-gates`, `test`
- **Description:** Dispatches all 18 repository quality checks concurrently across 4 worker threads. Must pass 100% with exit code 0.
- **Commands:**
  - `python .lovable/ai-fix-scripts/06-cicd-local-runner.py`

### 07-relative-path-fixer.py — Relative Path Fixer & Absolute URI Auditor
- **Tags:** `paths`, `relative-paths`, `absolute-paths`, `sanitizer`
- **Description:** Detects and auto-sanitizes absolute filesystem paths (`C:\...`, `D:\...`, `/home/...`) and `file:///` URIs in Markdown and configuration files.
- **Commands:**
  - `python .lovable/ai-fix-scripts/07-relative-path-fixer.py <path>`
  - `python .lovable/ai-fix-scripts/07-relative-path-fixer.py <path> --fix`

### 08-naming-autofixer.py — Boolean Naming & Code Convention Guard
- **Tags:** `naming`, `booleans`, `is-prefix`, `has-prefix`, `linter`
- **Description:** Flags explicit boolean comparisons (`== True`, `=== true`) and verifies positive prefix naming conventions (`is_`, `has_`).
- **Commands:**
  - `python .lovable/ai-fix-scripts/08-naming-autofixer.py <path> [--ext .ts,.go,.py]`

### 09-cli-help-auditor.py — CLI Command Discovery & Help Text Parity Auditor
- **Tags:** `cli`, `help`, `cobra`, `commander`, `docstrings`
- **Description:** Audits Go Cobra, TypeScript Commander, and Python CLI entry points to ensure `Short` descriptions and `Example` usage strings exist.
- **Commands:**
  - `python .lovable/ai-fix-scripts/09-cli-help-auditor.py <path> [--strict]`

### 10-encoding-normalizer.py — Fast UTF-8 & UNIX LF Normalizer
- **Tags:** `encoding`, `utf-8`, `bom-stripping`, `unix-lf`
- **Description:** Normalizes text files to UTF-8 without BOM and converts Windows CRLF to UNIX LF (`\n`).
- **Commands:**
  - `python .lovable/ai-fix-scripts/10-encoding-normalizer.py <path> --fix`

### 11-fast-file-scanner.py — High-Speed File Scanner & Cache Indexer
- **Tags:** `scanner`, `cache`, `indexing`, `file-list`, `discovery`
- **Description:** Scans repository files in <15ms, writes pluggable index to `tmp/cache/repo-file-cache.json`, and supports instant `<1ms` cache queries.
- **Commands:**
  - `python .lovable/ai-fix-scripts/11-fast-file-scanner.py --path <dir> --lang ts,go`
  - `python .lovable/ai-fix-scripts/11-fast-file-scanner.py --query-cache "<term>"`
  - `python .lovable/ai-fix-scripts/11-fast-file-scanner.py --check`

### 12-fast-cached-grep.py — Parallel Cached Grepper
- **Tags:** `grep`, `search`, `regex`, `parallel`, `content-search`
- **Description:** Multi-threaded parallel regex grep across repository files using pre-compiled regex objects and 8 worker threads.
- **Commands:**
  - `python .lovable/ai-fix-scripts/12-fast-cached-grep.py --pattern "<search-term>" --path <dir> --lang ts,go`
  - `python .lovable/ai-fix-scripts/12-fast-cached-grep.py --pattern "<regex>" --regex --path spec/`

### 13-file-size-guard.py — Fast Repository File Size & Blob Guard
- **Tags:** `file-size`, `blob-guard`, `security`, `binary-check`
- **Description:** Scans tracked files and alerts if any unapproved binary file exceeds the 2MB threshold.
- **Commands:**
  - `python .lovable/ai-fix-scripts/13-file-size-guard.py [--max-kb 2048] [--path <dir>]`

### 14-version-sync-checker.py — Version Synchronization & Changelog Guard
- **Tags:** `version`, `sync`, `changelog`, `package-json`, `release`
- **Description:** Verifies 100% version parity across `version.json`, `package.json`, and `changelog.md`.
- **Commands:**
  - `python .lovable/ai-fix-scripts/14-version-sync-checker.py <path>`

### 15-sequence-and-title-auditor.py — Sequence, Numbering & Title Header Auditor
- **Tags:** `sequence`, `title`, `h1-headers`, `markdown-audit`
- **Description:** Checks for sequence gaps and synchronizes Markdown `# H1` titles with numeric filename prefixes (e.g. `02-file.md` -> `# 02 - Title`).
- **Commands:**
  - `python .lovable/ai-fix-scripts/15-sequence-and-title-auditor.py <path> --fix`

### 16-installer-smoke-tester.py — Generic Installer Smoke Tester
- **Tags:** `installer`, `smoke-test`, `install-sh`, `install-ps1`
- **Description:** Smoke tests shell and PowerShell installer scripts to guarantee no unresolved placeholder tokens, verified SHA256 checksums, and non-destructive rename-first updates.
- **Commands:**
  - `python .lovable/ai-fix-scripts/16-installer-smoke-tester.py <path>`

### 17-fast-file-reader.py — AI Fast File Reader & Directory Explorer
- **Tags:** `reader`, `explorer`, `instant-read`, `ai-tool`
- **Description:** Sub-millisecond file reader, folder explorer, and pattern searcher designed specifically for AI agents.
- **Commands:**
  - `python .lovable/ai-fix-scripts/17-fast-file-reader.py --list-folder <path> [--ext .md,.ts]`
  - `python .lovable/ai-fix-scripts/17-fast-file-reader.py --read-file <file_path>`
  - `python .lovable/ai-fix-scripts/17-fast-file-reader.py --search-pattern "<term>" [--path <dir>]`

---

## 🎯 AI Operational Rules

1. **No Absolute Paths:** Always use relative paths from the git root. Never output `C:\...` or `file:///...`.
2. **Implicit Booleans:** Always evaluate positive booleans implicitly (`if is_valid:`, never `if is_valid == True:`).
3. **Prefix Boolean Variables & Functions:** Use `is_` or `has_` prefix for all boolean variables and return functions (`is_ready`, `has_match`, `is_success`, `has_failures`).
4. **Enums Format:** Python enums MUST use `PascalCase` class name ending in `Type`, `UPPER_CASE` members, and string values mirroring the member names.
5. **Quality Gates:** Before completing any work session, execute `python .lovable/ai-fix-scripts/06-cicd-local-runner.py` and verify all 18 checks pass.
