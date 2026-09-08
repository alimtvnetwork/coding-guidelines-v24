# Subtask 06: Verify Linters and Quality Gates

> **Plan:** `16-completed-plans-consolidation`  
> **Status:** Completed  
> **Target File:** Repository-wide  

---

## Intent

1. Run relative path checks to verify all internal links in consolidated milestones: `python linter-scripts/check-relative-paths.py`.
2. Run markdown spacing linter: `python linter-scripts/check-markdown-header-spacing.py`.
3. Run newline styling linter: `node linter-scripts/check-newline-styling.mjs`.
4. Run Go unit test suite: `go test ./pkg/... ./examples/... -count=1`.
5. Run full CI/CD local runner: `python 03-ai-scripts/06-cicd-local-runner.py` with exit code 0.

## Verification

- All linters and 36 CI/CD gates pass 100% green.
