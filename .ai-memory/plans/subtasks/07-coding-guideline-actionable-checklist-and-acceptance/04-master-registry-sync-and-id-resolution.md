# Subtask 04: Master Registry Harmonization, Verification Commands & ID Collision Resolution

> **/goal** Harmonize the master acceptance criteria registry, verify runnable verification commands across all domains, and resolve root style ID collisions.
> **/learn** Master the criteria taxonomy (`AC-CG-ROOT-*` vs `AC-CG-STYLE-*`), verify 1:1 traceability to authoritative guideline specs, and enforce zero-drift quality gates.

**Parent Plan:** `.ai-memory/plans/completed/06-coding-guidelines-actionable-checklists-and-acceptance.md`  
**Parent Spec:** `02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/02-component-spec.md`  
**Traceability ID:** Task-04  
**Target Files:**
- `02-spec/02-coding-guidelines/97-acceptance-criteria.md`
- `02-spec/02-coding-guidelines/02-canonical-size-tier.md`
- `02-spec/02-coding-guidelines/03-coding-style-checklist.md`
- `02-spec/02-coding-guidelines/04-consolidated-review-guide-condensed.md`
- `02-spec/02-coding-guidelines/05-consolidated-review-guide.md`

**Assigned Role:** Worker 02  
**Status:** COMPLETED  
**Updated:** 2026-10-03  

---

## 🎯 Actionable CI/CD & Agent Checklist

- [x] `/goal` Verify ID collision resolution: Ensure criteria in files 2..5 use `AC-CG-ROOT-002`, `AC-CG-ROOT-003`, `AC-CG-ROOT-004`, `AC-CG-ROOT-005` (not `AC-CG-STYLE-`).
- [x] `/learn` Verify Master Registry (`02-spec/02-coding-guidelines/97-acceptance-criteria.md`) Section `## AC-00` reflects `AC-CG-ROOT-002` through `AC-CG-ROOT-005`.
- [x] `/goal` Verify all criteria tables across all domains have explicit runnable `Verification Command` column values.
- [x] `/learn` Verify sections for `## AC-09: PowerShell Integration Registry (AC-CG-PWSH-000)` and `## AC-10: Research Standards Registry (AC-CG-RES-000)` are present and populated.
- [x] `/goal` Verify all stale links pointing to non-existent per-folder `98-changelog.md` files are completely removed.
- [x] `/learn` Run targeted linter checks (`03-ai-scripts/05-guideline-autofixer.py` and `linter-scripts/check-markdown-headings.py`) with zero violations.

---

## 1. Resolution Summary & Deliverables

### 1.1 ID Collision Resolution

- `02-spec/02-coding-guidelines/02-canonical-size-tier.md`: Confirmed criterion ID `AC-CG-ROOT-002` (Canonical Size Tier Enforcement).
- `02-spec/02-coding-guidelines/03-coding-style-checklist.md`: Confirmed criterion ID `AC-CG-ROOT-003` (Core Coding Style and Parameter Rules).
- `02-spec/02-coding-guidelines/04-consolidated-review-guide-condensed.md`: Confirmed criterion ID `AC-CG-ROOT-004` (Condensed Review Guide Rules).
- `02-spec/02-coding-guidelines/05-consolidated-review-guide.md`: Confirmed criterion ID `AC-CG-ROOT-005` (Consolidated Master Review Rules).
- Collision with `AC-CG-STYLE-` (reserved exclusively for `01-cross-language/04-code-style/`) completely resolved.

### 1.2 Master Registry Harmonization (`97-acceptance-criteria.md`)

- `## AC-00`: Maps `AC-CG-ROOT-002` through `AC-CG-ROOT-005` with valid relative links and runnable verification commands.
- `## AC-01` through `## AC-14`: All 16 criteria tables define explicit `Verification Command` values invoking `python 03-ai-scripts/05-guideline-autofixer.py ... --check-only`.
- `## AC-09`: PowerShell Integration Registry (`AC-CG-PWSH-000`) integrated with link to `09-powershell-integration/readme.md`.
- `## AC-10`: Research Standards Registry (`AC-CG-RES-000`) integrated with link to `10-research/readme.md`.
- Cross-references footer updated with all domain readmes and registries.
- Zero broken or stale links across all 133 external file references.

---

## 2. Verification Evidence

1. **Guideline Autofixer Validation:**
   ```bash
   python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
   ```
   - **Result:** Exit code 0.
   - **Output:**
     - All 198 files in `02-spec/02-coding-guidelines` have clean newlines.
     - All 0 code files in `02-spec/02-coding-guidelines` conform to implicit boolean rules.

2. **Markdown Heading Linter:**
   ```bash
   python linter-scripts/check-markdown-headings.py 02-spec/02-coding-guidelines/97-acceptance-criteria.md
   ```
   - **Result:** Exit code 0.
   - **Output:** `Markdown heading linter: all files OK.`

3. **Link Integrity Audit:**
   - Evaluated 133 file links in `02-spec/02-coding-guidelines/97-acceptance-criteria.md`.
   - Result: 133/133 links valid, 0 broken links, 0 stale changelog references.

4. **Database Task Manager Status:**
   ```bash
   python 03-ai-scripts/46-agent-sqlite-task-manager.py status --db .ai-memory/temp-agents/07-coding-guideline-actionable-checklist-and-acceptance/agent-task.db
   ```
   - **Result:** SubtaskId 6 / Task-04 recorded as `DONE`.
