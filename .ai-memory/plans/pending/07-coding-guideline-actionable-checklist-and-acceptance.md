# Execution Plan: Coding Guidelines Actionable Checklists & Testable Acceptance Criteria

> **Status:** ACTIVE  
> **Traceability ID:** `07-coding-guideline-actionable-checklist-and-acceptance`  
> **Database:** `.ai-memory/temp-agents/07-coding-guideline-actionable-checklist-and-acceptance/agent-task.db`  
> **Parent Spec:** `02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/01-architecture-spec.md`  

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

## 2. Research & Discovery Synthesis (A = 2 Subagents)

From Research 01 (Checklist & Style Audit) and Research 02 (Acceptance Criteria Parity):
- **198 Active Guideline Files Audited:** 100% feature AI Execution Prompt headers and `## 🎯 Actionable CI/CD & Agent Checklist`.
- **98.5% Acceptance Criteria Compliance:** 195/198 files contain full Gherkin `Given / When / Then` and runnable verification commands.
- **Specific Remediations Identified:**
  1. **Heading Blank Line Spacing:** Fix MD-H001 violations in `01-cross-language/34-string-normalization-and-equalfoldany.md` and `06-ai-optimization/97-acceptance-criteria.md`.
  2. **Duplicate Heading Numbering:** Renumber duplicate Principle 9 to Principle 12 in `01-cross-language/02-boolean-principles/04-parameters-and-conditions.md`.
  3. **ID Collisions:** Retag root style criteria in `02-spec/02-coding-guidelines/` from `AC-CG-STYLE-` to `AC-CG-ROOT-` to eliminate collision with `01-cross-language/04-code-style/`.
  4. **Master Registry Parity:** Synchronize `02-spec/02-coding-guidelines/97-acceptance-criteria.md` with runnable verification commands, add missing sections for `09-powershell-integration` and `10-research`, and fix stale links.

---

## 3. Subtask Decomposition (Worker Waves)

| Subtask ID | Title | Owner | Target Files | Status |
|:---|:---|:---|:---|:---:|
| `Task-03-Spacing-and-Headings` | Fix Heading Spacing, Duplicate Principle Numbering & Stale Links | Worker 01 | `01-cross-language/34-string-normalization-and-equalfoldany.md`, `01-cross-language/02-boolean-principles/04-parameters-and-conditions.md`, `06-ai-optimization/97-acceptance-criteria.md`, `01-cross-language/97-acceptance-criteria.md` | READY |
| `Task-04-Master-Registry-Sync` | Master Registry Harmonization, Verification Commands & ID Collision Resolution | Worker 02 | `02-spec/02-coding-guidelines/97-acceptance-criteria.md`, `02-spec/02-coding-guidelines/02-canonical-size-tier.md`, `02-spec/02-coding-guidelines/03-coding-style-checklist.md`, `02-spec/02-coding-guidelines/04-consolidated-review-guide-condensed.md`, `02-spec/02-coding-guidelines/05-consolidated-review-guide.md` | READY |

---

## 4. Verification Gates

1. `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only` -> exit 0
2. `python linter-scripts/check-relative-paths.py` -> exit 0
3. `python linter-scripts/check-sequence-integrity.py` -> exit 0
4. `python linter-scripts/check-markdown-headings.py` -> exit 0
