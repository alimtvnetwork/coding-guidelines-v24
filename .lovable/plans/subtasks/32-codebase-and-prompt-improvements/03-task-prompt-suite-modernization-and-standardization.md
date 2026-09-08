# Subtask 03: Prompt Suite Modernization and Standardization

**Plan:** [32-codebase-and-prompt-improvements.md](.lovable/plans/completed/32-codebase-and-prompt-improvements.md)  
**Status:** Completed  
**Disjoint File Scope:**
- `01-prompts/15-cg-execute/18-function-argument-reduction-and-params.md`
- `01-prompts/15-cg-execute/16-multi-language-enums-and-traits.md`
- `01-prompts/03-read-write/02-read-memory-enhanced.md`
- `01-prompts/13-plan-audit/02-plan-spec-steps-v2.md`
- `01-prompts/05-coding-guidelines/02-execute-coding-guideline-fix.md`
- `01-prompts/05-coding-guidelines/03-cg-audit-gap-n-steps.md`
- `01-prompts/14-execute/04-execute-ai-instruction-writer.md`
- `01-prompts/00-folder-structure/01-canonical-folder-structure.md`
- `01-prompts/08-dry-code/01-python-dry-architecture-and-caching.md`
- `01-prompts/22-ai-fix-script-prompts/01-python-file-manipulator.md`

---

## Acceptance Criteria

- [x] 1. Enforce Repo Rule #6: Replace all occurrences of `*apperror.AppError`, `apperror.AppError`, `gitmap/apperror`, and `apperror.WrapSimple` with `*appfault.AppError`, `appfault.AppError`, and `appfault` in target prompt files.
- [x] 2. Modernize Go enum guidelines in `01-prompts/15-cg-execute/16-multi-language-enums-and-traits.md` to reflect `03-ai-scripts/30-enum-generator.py` and the multi-file architecture (`variant.go`, `vars.go`, `variant_test.go`, `readme.md`, `byte` backing, `iota` starting at `Invalid`, `BasicEnum` integration).
- [x] 3. Update legacy header `- [ ] **Strict .lovable/ Folder Storage:**` to `- [ ] **Strict 03-ai-scripts/ Tooling Storage:**` in affected prompt files.
- [x] 4. Replace broken anchor `[AI Fix Scripts Memory](#ai-fix-scripts-memory)` with valid relative markdown link `[AI Fix Scripts Catalog](03-ai-scripts/01-index.md)`.
- [x] 5. Fix inverted link syntax `(AI Fix Scripts Memory)[#ai-fix-scripts-memory]` in `01-prompts/14-execute/04-execute-ai-instruction-writer.md`.
- [x] 6. Replace stale `spec/` references with `02-spec/` in `01-prompts/00-folder-structure/01-canonical-folder-structure.md` and `01-prompts/03-read-write/02-read-memory-enhanced.md`.
- [x] 7. Harmonize prompt versions to `Prompt Version: 2.1.0` in `01-canonical-folder-structure.md`, `01-python-dry-architecture-and-caching.md`, and `01-python-file-manipulator.md`.

---

## Verification Commands

```powershell
python 03-ai-scripts/21-sequence-integrity-linter.py
python 03-ai-scripts/22-doc-path-linter.py
```
