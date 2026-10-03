# Subtask 02: Harden Style & Naming Guidelines (AI Execution Prompt)

> **/goal** Verify and harden actionable checklists and acceptance criteria in `02-spec/02-coding-guidelines/03-coding-style-checklist.md` and `02-spec/02-coding-guidelines/01-cross-language/22-variable-naming-conventions.md`.
> **/learn** Master 4-part anatomy, affirmative boolean rules, and parameter limits.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Inspect `02-spec/02-coding-guidelines/03-coding-style-checklist.md` and ensure checklist and acceptance criteria are complete and clear.
- [ ] `/learn` Verify `02-spec/02-coding-guidelines/01-cross-language/22-variable-naming-conventions.md` checklist and acceptance criteria enforce semantic naming and ban generic garbage names.
- [ ] `/goal` Ensure vertical whitespace before return and after `}` are strictly observed.
- [ ] `/learn` Verify with guideline autofixer returning exit code 0.

. **CRITICAL AI INSTRUCTION:** Worker 02 owns `02-spec/02-coding-guidelines/03-coding-style-checklist.md` and `02-spec/02-coding-guidelines/01-cross-language/22-variable-naming-conventions.md`. Do not modify any other file.

---

## Verification & Acceptance Criteria

### AC-CG-SUB-002: Style and Naming Guidelines Compliance

**Given** `02-spec/02-coding-guidelines/03-coding-style-checklist.md` and `01-cross-language/22-variable-naming-conventions.md`.
**When** Audited by guideline autofixer.
**Then** Both files strictly conform to the 4-part prompt anatomy and have exit code 0.

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
```
**Expected:** exit 0. Zero violations.
