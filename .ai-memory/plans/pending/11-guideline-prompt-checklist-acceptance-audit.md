# Plan: Coding Guidelines Prompt, Actionable Checklists & Acceptance Audit (AI Execution Prompt)

> **/goal** Execute an end-to-end verification and hardening of coding guidelines in `02-spec/02-coding-guidelines/`, ensuring 100% adherence to active AI prompt headers, actionable agent checklists, and testable acceptance criteria across Style, Boolean, and Naming specifications.
> **/learn** Master the 4-part architectural anatomy, multi-agent wave partitioning (A = 2, H = 2), worker git ban, and automated zero-regression quality gates.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify all 199 files in `02-spec/02-coding-guidelines/` have the active AI prompt header `(AI Execution Prompt)` and `> **/goal**` / `> **/learn**`.
- [ ] `/learn` Verify every file includes `## 🎯 Actionable CI/CD & Agent Checklist` with interactive checkboxes and `. **CRITICAL AI INSTRUCTION:**`.
- [ ] `/goal` Verify all style, boolean, and naming specs conclude with testable `## Verification & Acceptance Criteria` (`AC-CG-*`) in Given/When/Then format.
- [ ] `/learn` Execute targeted linters (`03-ai-scripts/05-guideline-autofixer.py`, `linter-scripts/check-relative-paths.py`) ensuring exit code 0.

. **CRITICAL AI INSTRUCTION:** This plan governs the audit, hardening, and verification of coding guidelines.

**Version:** 1.0.0  
**Updated:** 2026-10-03  
**Status:** Active  
**AI Confidence:** Production-Ready  
**Ambiguity:** None  

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

## 2. Subtasks & Wave Distribution

| Task-ID | Subtask | Owner | Owned Files | Status |
|:---|:---|:---|:---|:---|
| Task-01 | Standardize Boolean Principles index | Worker 01 | `02-spec/02-coding-guidelines/01-cross-language/02-boolean-principles/readme.md` | PENDING |
| Task-02 | Harden style and naming guidelines | Worker 02 | `02-spec/02-coding-guidelines/03-coding-style-checklist.md`, `02-spec/02-coding-guidelines/01-cross-language/22-variable-naming-conventions.md` | PENDING |
| Task-03 | Author parent plan and subtasks | Lead | `.ai-memory/plans/pending/11-guideline-prompt-checklist-acceptance-audit.md` | IN PROGRESS |
| Task-04 | Run verification and finalize | Lead | `.ai-memory/plans/readme.md` | PENDING |

---

## 3. Verification & Acceptance Criteria

### AC-CG-PLAN-011: Coding Guidelines Audit Plan Conformance

**Given** The audit plan in `.ai-memory/plans/pending/11-guideline-prompt-checklist-acceptance-audit.md`.  
**When** Audited by repository linters and guideline autofixers.  
**Then** Plan fulfills 4-part anatomy, links all subtasks, and exits with 0.

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
```
**Expected:** exit 0. Zero violations.
