# Subtask 30.4 — GitMap Commit Formatter Engine Hardening

> **Parent Plan:** `.ai-memory/plans/pending/30-root-cause-analysis-and-fix.md`  
> **Target Package:** `d:\work\gitmap\cli\cmdcommit\`  
> **Status:** Pending  

---

## 1. Objectives
1. Inspect `cli/cmdcommit/commit_formatter.go`:
   - Add `COMMIT`, `COMMITS`, `GITCOMMIT` to `isCommitTypeToken()`.
   - Update `resolveHyphenModuleAndItem()` to strip redundant `commit` / `commits` module tokens automatically.
   - Support Title Case (`Auth`, `Scanner`, `Ui`) and lowercase (`auth`, `scanner`, `ui`) modules without forcing everything to lowercase if user provided Title Case.
   - Down-case ALL-CAPS words so shout-casing is eliminated.
2. Update unit tests in `cli/cmdcommit/commit_formatter_test.go` to cover `commit` stripping and Title Case preservation.
3. Recompile and install binary.
