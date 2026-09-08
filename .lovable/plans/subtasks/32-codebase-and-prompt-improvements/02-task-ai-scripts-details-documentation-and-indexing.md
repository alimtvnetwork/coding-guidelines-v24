# Subtask 02: AI Scripts Details Documentation and Indexing

**Plan:** [32-codebase-and-prompt-improvements.md](.lovable/plans/completed/32-codebase-and-prompt-improvements.md)  
**Status:** Completed  
**Disjoint File Scope:**
- `03-ai-scripts/01-index.md`

---

## Acceptance Criteria

- [x] 1. Satisfy repository rule: *"For every script, I have included a `<details>` collapsible tag explaining exactly why the script is there and what it does."*
- [x] 2. Implement comprehensive `<details>` blocks for all 31 scripts (`01-index.md` through `31-md-gap-fixer.py`):
  - `01-index.md`: Script catalog and AI pre-flight guide.
  - `02-shared-engine.py`: Centralized constants, regex registry, dual-platform locks, caching.
  - `03-file-manipulator.py`: Mass renaming, lowercasing, and sequence normalization CLI.
  - `04-newline-fixer.py`: Trailing whitespace and POSIX final newline normalizer.
  - `05-guideline-autofixer.py`: Composite autofixer for newlines, boolean naming, and guidelines.
  - `06-cicd-local-runner.py`: High-concurrency local test suite runner with selective failure logging.
  - `07-relative-path-fixer.py`: Absolute path and `file:///` URI detector and sanitizer.
  - `08-naming-autofixer.py`: Lowercase naming and boolean prefix enforcement linter.
  - `09-cli-help-auditor.py`: Command-line documentation and help verification tool.
  - `10-encoding-normalizer.py`: UTF-8 and LF line ending normalization utility.
  - `11-fast-file-scanner.py`: High-performance filesystem scanner with memory cache.
  - `12-fast-cached-grep.py`: Multi-threaded pattern matcher over cached repository files.
  - `13-file-size-guard.py`: Binary blob and oversized file gatekeeper.
  - `14-version-sync-checker.py`: Semantic version synchronization auditor.
  - `15-sequence-and-title-auditor.py`: Numeric file sequence and H1 title consistency checker.
  - `16-installer-smoke-tester.py`: Bash and PowerShell installer sandbox smoke test runner.
  - `17-fast-file-reader.py`: Instantaneous file and directory reader.
  - `18-codebase-topology-discoverer.py`: Polyglot architecture mapper with TTL caching.
  - `19-artifact-remover.py`: Safe test artifact and cache cleaner with git index hygiene.
  - `20-plan-consolidator.py`: Lovable plan index synchronizer and archiver.
  - `21-sequence-integrity-linter.py`: Plan and subtask sequence integrity validator.
  - `22-doc-path-linter.py`: Documentation relative path reference checker.
  - `23-coding-guideline-path-consolidator.py`: Coding guideline cross-link canonicalizer.
  - `24-spec-path-migrator.py`: Specification directory path migrator.
  - `25-repo-migrator.py`: Transactional repository asset migrator with rollback.
  - `26-go-code-formatter.py`: Cross-platform `gofmt` code formatter with staged file support.
  - `27-misspell-auditor.py`: American English spelling linter and autofixer.
  - `28-go-preflight-ci.py`: Standalone Go unit test and linter runner.
  - `29-release-orchestrator.py`: Semantic release automation orchestrator.
  - `30-enum-generator.py`: Multi-file Go enum generator with `BasicEnum` scaffolding.
  - `31-md-gap-fixer.py`: Markdown consecutive empty line normalizer.
- [x] 3. Each details block must include: (a) Why it exists, (b) What it does, (c) Command line usage examples for check and fix modes.
- [x] 4. All markdown links and paths in `03-ai-scripts/01-index.md` must be strictly relative Git paths.

---

## Verification Commands

```powershell
python 03-ai-scripts/21-sequence-integrity-linter.py
python 03-ai-scripts/31-md-gap-fixer.py 03-ai-scripts/01-index.md
```
