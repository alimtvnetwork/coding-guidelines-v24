# Subtask 03: Fix Heading Spacing, Duplicate Principle Numbering & Stale Links

> **Parent Spec:** [`02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/01-architecture-spec.md`](../../../../02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/01-architecture-spec.md)  
> **Status:** `PENDING`  
> **Traceability ID:** `Task-03`  
> **Subtask DB ID:** `5`  
> **Assigned Agent:** `Worker 01`  
> **Target Files (Disjoint Boundary):**  
> 1. `02-spec/02-coding-guidelines/01-cross-language/34-string-normalization-and-equalfoldany.md`  
> 2. `02-spec/02-coding-guidelines/01-cross-language/02-boolean-principles/04-parameters-and-conditions.md`  
> 3. `02-spec/02-coding-guidelines/06-ai-optimization/97-acceptance-criteria.md`  
> 4. `02-spec/02-coding-guidelines/01-cross-language/97-acceptance-criteria.md`  

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Fix MD-H001 heading blank line spacing in `01-cross-language/34-string-normalization-and-equalfoldany.md`.
- [ ] `/learn` Renumber duplicate "Principle 9" to "Principle 12" in `01-cross-language/02-boolean-principles/04-parameters-and-conditions.md`.
- [ ] `/goal` Fix heading blank line spacing in `06-ai-optimization/97-acceptance-criteria.md`.
- [ ] `/learn` Remove stale links to deleted `98-changelog.md` from `01-cross-language/97-acceptance-criteria.md`.
- [ ] `/goal` Verify changes using `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/01-cross-language --check-only` with `Expected: exit 0`.

. **CRITICAL AI INSTRUCTION:** Workers modifying these 4 files must operate within this isolated bounding box and observe the strict worker git ban.

---

## Work Instructions

1. **Heading Blank Lines in `34-string-normalization-and-equalfoldany.md`:** Ensure all markdown headers (`##`, `###`, `####`) are preceded and succeeded by a blank line.
2. **Duplicate Heading Numbering in `04-parameters-and-conditions.md`:** Locate the second `## Principle 9: No Explicit True Checks (TOTAL BAN)` (around line 342) and renumber it to `## Principle 12: No Explicit True Checks (TOTAL BAN)`.
3. **Heading Spacing in `06-ai-optimization/97-acceptance-criteria.md`:** Ensure blank lines around headings on lines 42 and 49.
4. **Stale Links in `01-cross-language/97-acceptance-criteria.md`:** Remove stale link references to deleted per-folder `98-changelog.md` and replace with valid relative references to root `../../changelog.md` where appropriate.
