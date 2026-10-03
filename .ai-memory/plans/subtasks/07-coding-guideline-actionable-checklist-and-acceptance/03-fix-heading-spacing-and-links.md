# Subtask 03: Fix Heading Spacing, Duplicate Principle Numbering & Stale Links

> **Parent Spec:** [`02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/01-architecture-spec.md`](../../../../02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/01-architecture-spec.md)  
> **Status:** `COMPLETED`  
> **Traceability ID:** `Task-03`  
> **Target Files (Strict Disjoint Bounding Box):**
> - `02-spec/02-coding-guidelines/01-cross-language/34-string-normalization-and-equalfoldany.md`
> - `02-spec/02-coding-guidelines/01-cross-language/02-boolean-principles/04-parameters-and-conditions.md`
> - `02-spec/02-coding-guidelines/06-ai-optimization/97-acceptance-criteria.md`
> - `02-spec/02-coding-guidelines/01-cross-language/97-acceptance-criteria.md`

---

## 1. Subtask Objective & Context

Ensure strict markdown heading spacing (MD-H001 compliance) across cross-language and AI optimization guideline documents, ensure that duplicate principle numbering in `04-parameters-and-conditions.md` is cleanly resolved to Principle 12: No Explicit True Checks (TOTAL BAN), and verify that acceptance criteria registries contain 100% valid relative links without stale references.

---

## 2. Actionable Verification & Execution Checklist

- [x] Verify and ensure that all markdown headings in `02-spec/02-coding-guidelines/01-cross-language/34-string-normalization-and-equalfoldany.md` are preceded and followed by a blank line (MD-H001 compliance).
- [x] Verify that in `02-spec/02-coding-guidelines/01-cross-language/02-boolean-principles/04-parameters-and-conditions.md`, the duplicate "Principle 9" (around line 342) is cleanly renumbered to "Principle 12: No Explicit True Checks (TOTAL BAN)".
- [x] Verify that headings in `02-spec/02-coding-guidelines/06-ai-optimization/97-acceptance-criteria.md` have proper blank lines around them.
- [x] Verify that in `02-spec/02-coding-guidelines/01-cross-language/97-acceptance-criteria.md`, stale links to non-existent per-folder `98-changelog.md` are removed or point cleanly to valid relative paths (points to root `changelog.md`).
- [x] Run targeted linter checks:
  - `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/01-cross-language --check-only`
  - `python linter-scripts/check-markdown-headings.py 02-spec/02-coding-guidelines/01-cross-language`
  - `python linter-scripts/check-markdown-headings.py 02-spec/02-coding-guidelines/06-ai-optimization`
- [x] Update subtask status to COMPLETED and report verification evidence back to lead orchestrator.

---

## 3. Verification Evidence & Quality Gates

1. **Heading Linter (`check-markdown-headings.py`):**
   - `02-spec/02-coding-guidelines/01-cross-language/34-string-normalization-and-equalfoldany.md`: PASS (all headings preceded and followed by blank lines)
   - `02-spec/02-coding-guidelines/01-cross-language`: PASS (all 68 files OK, 0 violations)
   - `02-spec/02-coding-guidelines/06-ai-optimization`: PASS (all 10 files OK, 0 violations)

2. **Principle Sequence Check:**
   - In `02-spec/02-coding-guidelines/01-cross-language/02-boolean-principles/04-parameters-and-conditions.md`:
     - Line 342 is cleanly `## Principle 12: No Explicit True Checks (TOTAL BAN)` following Principle 11 (`IsDefined` replacement).

3. **Link Validation:**
   - All 38 relative markdown links in `02-spec/02-coding-guidelines/01-cross-language/97-acceptance-criteria.md` resolve to valid existing files.
   - Changelog reference cleanly links to root [`changelog.md`](../../../changelog.md) with exit 0.
   - All 9 relative markdown links in `02-spec/02-coding-guidelines/06-ai-optimization/97-acceptance-criteria.md` resolve cleanly.

4. **Guideline Autofixer (`03-ai-scripts/05-guideline-autofixer.py`):**
   - `02-spec/02-coding-guidelines/01-cross-language`: 68 files scanned, 0 newline violations, 0 boolean violations, exit 0.
   - `02-spec/02-coding-guidelines/06-ai-optimization`: 10 files scanned, 0 newline violations, 0 boolean violations, exit 0.
