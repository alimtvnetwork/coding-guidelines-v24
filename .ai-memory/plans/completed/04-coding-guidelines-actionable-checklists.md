# Completed Plan: Coding Guidelines Actionable Checklists & Acceptance Criteria Transformation

> **/goal** Autonomously transform 32 cross-language coding guidelines into actionable prompt-driven specifications with agent checklists and testable acceptance criteria using execute-parent-task-with-n-steps-v6.
> **/learn** Enforce disjoint file boundaries, bounded 5–8 file micro-batches, parallel subagent dispatch (A = 2, H = 2), and zero-lock execution under execute-parent-task-with-n-steps-v6.

**Version:** 1.0.0  
**Completed:** 2026-10-02  
**Status:** COMPLETED  
**Parent Task ID:** 1 (SQLite DB: `.ai-memory/temp-agents/04-coding-guidelines-actionable-checklists/agent-task.db`)  
**Canonical Spec:** [`02-spec/21-app/04-coding-guidelines-actionable-checklists-and-acceptance-criteria/01-architecture-spec.md`](../../02-spec/21-app/04-coding-guidelines-actionable-checklists-and-acceptance-criteria/01-architecture-spec.md)  

---

## 1. User Request (Verbatim)

```text
# High Priority Instruction

Can you please go into the coding guideline and see the coding guideline, how it is written? It's not written as an action or prompt on every file. Try to have from the README.md file inside the coding guideline folder, like a spec and then coding guideline. Try to have every one of the file try to mention as actionable items that the AI agent must follow. Okay. As a checklist. And what is acceptance criteria at the end that needs to be mentioned on each one of the specs for coding guideline inside, like the style guideline, Boolean guideline, naming convention, things like that. Okay. Do you understand? Can you please follow through? Do you have any question and confusion?

# Actionable Items Must Follow Non-Negotiable

1. Review the coding guidelines in the README.md file within the coding guideline folder.
2. Create a specification and coding guideline document.
3. Ensure each file includes actionable items for AI agents as a checklist.
4. Define acceptance criteria for each spec, including style guidelines, Boolean guidelines, and naming conventions.

Must follow and spawn agent using 

@[.agents/skills/execute-parent-task-with-n-steps-v6]

## Additional Instructions

learn /learn if you have to learn something and /plan stuff before working please./plan
```

---

## 2. Completed Subtasks & Evidence Matrix

| Wave | Task-ID | Worker | Owned Files | Status | Evidence |
|:---|:---|:---|:---|:---|:---|
| **Phase 1** | Task-01 | Lead | `02-spec/21-app/04-coding-guidelines-actionable-checklists-and-acceptance-criteria/*` | DONE | Created architecture spec, rollout spec, registered in 21-app readme |
| **Wave 1** | Task-02A | Worker 01 | `01-cross-language/02-boolean-principles/*` (6 files) | DONE | PASS exit 0, 6 files upgraded with AI checklists and AC-CG-BOOL-001..006 |
| **Wave 1** | Task-02B | Worker 02 | `01-cross-language/12-no-negatives.md`, `24-boolean-flag-methods.md` | DONE | PASS exit 0, 2 files upgraded with AI checklists and AC-CG-BOOL-012/024 |
| **Wave 2** | Task-03A | Worker 01 | `01-cross-language/04-code-style/02-05` (4 files) + readme | DONE | PASS exit 0, 5 files upgraded with AI checklists and AC-CG-STYLE-001..005 |
| **Wave 2** | Task-03B | Worker 02 | `01-cross-language/04-code-style/06-08`, `21-newline...` | DONE | PASS exit 0, 4 files upgraded with AI checklists and AC-CG-STYLE-005..008 |
| **Wave 3** | Task-04A | Worker 01 | `01-cross-language/10`, `11`, `22` | DONE | PASS exit 0, 3 files upgraded with AI checklists and AC-CG-NAME-010..022 |
| **Wave 3** | Task-04B | Worker 02 | `01-cross-language/07`, `28` | DONE | PASS exit 0, 2 files upgraded with AI checklists and AC-CG-NAME-007/028 |
| **Wave 4** | Task-05A | Worker 01 | `01-cross-language/13`, `18`, `32` | DONE | PASS exit 0, 3 files upgraded with AI checklists and AC-CG-TYPE-013..032 |
| **Wave 4** | Task-05B | Worker 02 | `01-cross-language/33`, `34` | DONE | PASS exit 0, 2 files upgraded with AI checklists and AC-CG-TYPE-033/034 |
| **Wave 5** | Task-06A | Worker 01 | `01-cross-language/03`, `06`, `08` | DONE | PASS exit 0, 3 files upgraded with AI checklists and AC-CG-ARCH-006..008 |
| **Wave 5** | Task-06B | Worker 02 | `01-cross-language/16`, `19`, `20`, `23` | DONE | PASS exit 0, 4 files upgraded with AI checklists and AC-CG-ARCH-016..023 |
| **Wave 6** | Task-07 | Lead | `01-cross-language/97-acceptance-criteria.md`, `97-acceptance-criteria.md`, readmes | DONE | PASS exit 0, consolidated Acceptance Criteria registry created |
| **Wave 7A** | Task-08A | Worker 01 | `03-golang/readme.md`, `02-boolean-standards.md`, `03-httpmethod-enum.md`, `05-defer-rules.md` | DONE | PASS exit 0, 4 files upgraded with AI checklists and AC-CG-GO-000..005 |
| **Wave 7B** | Task-08B | Worker 02 | `03-golang/06-string-slice-internals.md`, `07-code-severity-taxonomy.md`, `08-pathutil-fileutil-spec.md`, `09-wrapped-boolean-results.md` | DONE | PASS exit 0, 4 files upgraded with AI checklists and AC-CG-GO-006..009 |
| **Wave 7C** | Task-08C | Lead | `03-golang/97-acceptance-criteria.md`, `02-coding-guidelines/97-acceptance-criteria.md` | DONE | PASS exit 0, Go Acceptance Criteria registry and root criteria synchronized |

---

## 3. Verification & Compliance Evidence

1. **Relative Paths:** `python linter-scripts/check-relative-paths.py` -> exit 0 (0 absolute paths or `file:///` URIs across repository).
2. **Forbidden Strings & Secrets:** `python linter-scripts/check-forbidden-strings.py` -> exit 0 (zero secrets, zero forbidden patterns).
3. **Guideline Formatting (Cross-Language):** `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/01-cross-language --check-only` -> exit 0.
4. **Guideline Formatting (Golang):** `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/03-golang --check-only` -> exit 0.

