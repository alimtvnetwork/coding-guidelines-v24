# Subtask 01: Standardize Boolean Principles Index (AI Execution Prompt)

> **/goal** Clean up Document Inventory table in `02-spec/02-coding-guidelines/01-cross-language/02-boolean-principles/readme.md`, ensuring all links point to existing files and removing duplicate rows.
> **/learn** Master relative path integrity and table hygiene across modular specification subfolders.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Inspect `02-spec/02-coding-guidelines/01-cross-language/02-boolean-principles/readme.md` Document Inventory table.
- [ ] `/learn` Fix file link target for Principle 1 to `[02-naming-prefixes.md](./02-naming-prefixes.md)` instead of `01-naming-prefixes.md`.
- [ ] `/goal` Remove the duplicated row `| — | 99-consistency-report.md | — | — |`.
- [ ] `/learn` Verify that `02-spec/02-coding-guidelines/01-cross-language/02-boolean-principles/readme.md` passes autofixer.

. **CRITICAL AI INSTRUCTION:** Worker 01 owns `02-spec/02-coding-guidelines/01-cross-language/02-boolean-principles/readme.md`. Do not modify any other file.

---

## Verification & Acceptance Criteria

### AC-CG-SUB-001: Boolean Principles Index Table Integrity

**Given** `02-spec/02-coding-guidelines/01-cross-language/02-boolean-principles/readme.md`.
**When** Audited by guideline linter and link verifiers.
**Then** All links in Document Inventory resolve to valid files and no duplicate rows exist.

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/01-cross-language/02-boolean-principles --check-only
```
**Expected:** exit 0. Zero violations.
