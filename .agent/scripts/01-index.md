# AI Fix Scripts Index

> **/goal** Master and execute the repository's suite of ultra-fast Python scripts for linting, path resolution, naming enforcement, version synchronization, and local CI verification.
> **/learn** Read the script specifications below and run scripts via `python .lovable/ai-fix-scripts/<script-name>.py`.

## 🎯 Actionable CI/CD & Agent Checklist

1. [ ] `/goal` Verify local code and markdown compliance before submitting changes.
2. [ ] `/learn` Run all local verification linters via `python .lovable/ai-fix-scripts/03-cicd-local-runner.py`.
3. [ ] `/goal` Ensure all files use strict relative paths with zero absolute filesystem references.
4. [ ] `/learn` Scan repository files rapidly via `python .lovable/ai-fix-scripts/08-fast-file-scanner.py`.

---

**Version:** 3.0.0  
**Updated:** 2026-08-30  
**AI Confidence:** Production-Ready  
**Ambiguity:** None

---

## 🛠️ Script Catalog

| # | Script | Purpose | Execution Time |
|---|--------|---------|----------------|
| 02 | `02-newline-fixer.py` | Fixes trailing whitespace and missing final newlines | ~15ms |
| 03 | `03-cicd-local-runner.py` | Runs all 6 CI quality checks locally via `ThreadPoolExecutor` | ~40ms |
| 04 | `04-relative-path-fixer.py` | Detects and fixes absolute paths / `file:///` URIs | ~30ms |
| 05 | `05-naming-autofixer.py` | Enforces lowercase filenames and directory naming rules | ~20ms |
| 06 | `06-cli-help-auditor.py` | Validates CLI `--help` examples against actual implementations | ~25ms |
| 07 | `07-encoding-normalizer.py` | Normalizes all files to strict UTF-8 with UNIX LF line endings | ~35ms |
| 08 | `08-fast-file-scanner.py` | High-speed repo file scanner (<15ms full scan, <1ms cache query) | ~14ms |
| 09 | `09-fast-cached-grep.py` | Parallel regex matcher leveraging pre-warmed file cache | ~12ms |
| 10 | `10-file-size-guard.py` | Audits repository files for oversized binary blobs (>1MB) | ~10ms |
| 11 | `11-version-sync-checker.py` | Verifies synchronization of `version.json`, `package.json`, and `changelog.md` | ~5ms |
| 12 | `12-sequence-and-title-auditor.py` | Audits and aligns numeric file sequence prefixes and `# H1` titles | ~20ms |
| 13 | `13-installer-smoke-tester.py` | Generic installer smoke test validating script placeholders & hashes | ~8ms |

---

## Usage Guidelines

- **Fast Indexing:** Run `python .lovable/ai-fix-scripts/08-fast-file-scanner.py --lang go,ts --path spec/` for filtered file lookups.
- **Local CI Testing:** Run `python .lovable/ai-fix-scripts/03-cicd-local-runner.py` before finalizing any PR or major edit batch.
- **Line Ending & Sequence Normalization:** Run `python .lovable/ai-fix-scripts/07-encoding-normalizer.py` and `python .lovable/ai-fix-scripts/12-sequence-and-title-auditor.py --fix`.
