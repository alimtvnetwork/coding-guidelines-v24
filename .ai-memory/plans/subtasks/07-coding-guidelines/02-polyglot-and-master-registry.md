# Subtask 02: Polyglot Language Registries & Master Registry Harmonization (AI Execution Prompt)

> **/goal** Harmonize criteria IDs across Python and C++ specs, add missing `*-REG-001` criteria across 7 domain registries, and update the Master Registry in `02-spec/02-coding-guidelines/97-acceptance-criteria.md`.
> **/learn** Enforce strictly relative paths, zero absolute paths, 4-column table layouts, and run guideline autofixer verification.

## 🎯 Actionable CI/CD & Agent Checklist

- [x] `/goal` Update `12-python/readme.md` to `AC-CG-PY-001`, and append Gherkin blocks for `AC-CG-PY-003` and `004` to `12-python/02-standards.md`.
- [x] `/learn` Update `13-cpp/readme.md` to `AC-CG-CPP-001`, and append Gherkin blocks for `AC-CG-CPP-003` and `004` to `13-cpp/02-standards.md`.
- [x] `/goal` Add missing `*-REG-001` criteria to domain registries (`03-golang`, `04-php`, `05-rust`, `06-ai-optimization`, `07-csharp`, `08-file-folder-naming`, `11-security`).
- [x] `/learn` Update Master Registry `97-acceptance-criteria.md` with `AC-CG-STYLE-009`, Python/C++ aligned IDs, and new section `AC-15` indexing `02-spec/21-app/07-.../`.
- [x] `/goal` Verify 4-column table layout across all Master Registry tables.
- [x] `/learn` Run `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only`.

. **CRITICAL AI INSTRUCTION:** Do NOT run git commands. Report findings and edits back to the orchestrator.

**Assigned Worker:** Worker 2
**Status:** Completed
