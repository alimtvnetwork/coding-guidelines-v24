# Subtask 03: Quality Verification, Linters & Local Atomic Commit

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/letterly-prompts-heatmap-one-liner-integration/03-verification-and-commit.md`  
> **Parent Plan:** `.ai-memory/plans/letterly-prompts-heatmap-one-liner-integration.md`  
> **Specification:** `02-spec/21-app/letterly-prompts-heatmap-one-liner-integration/01-architecture-spec.md`  
> **Status:** PENDING  
> **Assigned Role:** LeadOrchestrator

---

## 1. Objective

Execute comprehensive local verification across modified Letterly and Cursor prompts:
1. Run `python -m unittest linter-scripts/tests/test_check_letterly_format.py` (ensure golden tests pass).
2. Run `python linter-scripts/check-relative-paths.py` (ensure 0 absolute paths).
3. Run `python linter-scripts/check-forbidden-strings.py` (ensure 0 secrets or forbidden strings).
4. Run `python linter-scripts/check-prompts-loaded.py` (ensure all prompts on disk are loaded).
5. Inspect `git status --porcelain` to verify blast radius containment.
6. Commit atomically and push upstream via `gitmap cpf "letterly - include heatmap telemetry one liner across all prompts and cursor mirrors"`.
7. **Strict Invariant:** NO external fleet sync.

---

## 2. Action Checklist

- [ ] Execute letterly format linter tests
- [ ] Execute relative paths linter
- [ ] Execute forbidden strings linter
- [ ] Execute prompt index linter
- [ ] Inspect git working tree cleanliness
- [ ] Stage, commit, and push via `gitmap cpf`
