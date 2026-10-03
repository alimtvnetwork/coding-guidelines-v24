# Subtask 04: Master Registry Harmonization, Verification Commands & ID Collision Resolution

> **Parent Spec:** [`02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/02-component-spec.md`](../../../../02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/02-component-spec.md)  
> **Status:** `PENDING`  
> **Traceability ID:** `Task-04`  
> **Subtask DB ID:** `6`  
> **Assigned Agent:** `Worker 02`  
> **Target Files (Disjoint Boundary):**  
> 1. `02-spec/02-coding-guidelines/97-acceptance-criteria.md`  
> 2. `02-spec/02-coding-guidelines/02-canonical-size-tier.md`  
> 3. `02-spec/02-coding-guidelines/03-coding-style-checklist.md`  
> 4. `02-spec/02-coding-guidelines/04-consolidated-review-guide-condensed.md`  
> 5. `02-spec/02-coding-guidelines/05-consolidated-review-guide.md`  

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Add runnable `Verification Command` column across all tables in `02-spec/02-coding-guidelines/97-acceptance-criteria.md`.
- [ ] `/learn` Re-tag root style criteria in files 2..5 from `AC-CG-STYLE-002..005` to `AC-CG-ROOT-002..005` to eliminate direct ID collisions with `01-cross-language/04-code-style/`.
- [ ] `/goal` Add missing sections in master registry for `09-powershell-integration` (`AC-CG-PWSH-000`) and `10-research` (`AC-CG-RES-000`).
- [ ] `/learn` Remove stale links pointing to non-existent per-folder `98-changelog.md` files from master registry.
- [ ] `/goal` Verify compliance using `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only` with `Expected: exit 0`.

. **CRITICAL AI INSTRUCTION:** Workers modifying these 5 files must operate within this isolated bounding box and observe the strict worker git ban.

---

## Work Instructions

1. **ID Collision Resolution:**
   - In `02-canonical-size-tier.md`, update `AC-CG-STYLE-002` -> `AC-CG-ROOT-002`.
   - In `03-coding-style-checklist.md`, update `AC-CG-STYLE-003` -> `AC-CG-ROOT-003`.
   - In `04-consolidated-review-guide-condensed.md`, update `AC-CG-STYLE-004` -> `AC-CG-ROOT-004`.
   - In `05-consolidated-review-guide.md`, update `AC-CG-STYLE-005` -> `AC-CG-ROOT-005`.
2. **Master Registry Harmonization (`97-acceptance-criteria.md`):**
   - Update Section `## AC-00: Root Style and Sizing Guidelines` to reflect `AC-CG-ROOT-002..005`.
   - Add `Verification Command` column to all tables.
   - Add sections for `## AC-09: PowerShell Integration Registry` and `## AC-10: Research Standards Registry`.
   - Remove stale links to deleted per-folder `98-changelog.md` files; point to root `../changelog.md` or omit dead links.
