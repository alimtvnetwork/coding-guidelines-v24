# Plan 30 — Commit Pseudo-Module Root Cause Analysis & Remediation

> **Status:** Completed  
> **Sequence ID:** 30  
> **Slug:** `root-cause-analysis-and-fix`  
> **Specification:** `02-spec/21-app/30-root-cause-analysis-and-fix/`  
> **Subtasks Directory:** `.ai-memory/plans/subtasks/30-root-cause-analysis-and-fix/`  
> **Date:** 2026-10-27  

---

## 1. Executive Summary & Root Cause Summary
The user pointed out that recent Git commit messages prominently featured the word `commit` in the module slot (`Chore: commit - lowercase module names...` and `Feature: COMMIT - author attribution...`).
The user explicitly asked:
1. "Why this commit part is here? There should be never commit. Commit is what? It is already committing, why is it writing at the beginning commits?"
2. "Title case is absolutely fine, or lower case is also fine."
3. "Find the root cause in the instruction, in the execute prompt in the v6, in the coding guideline, map the skills in the coding guideline and also in the code. Tell me the root cause, fix the root cause, and make a release with minor bump."

---

## 2. Subtasks Roadmap

| Subtask ID | File | Description | Target | Status |
| :--- | :--- | :--- | :--- | :--- |
| **Subtask-01** | `01-spec-and-plans.md` | Author canonical specifications and pending subtask roadmap in `02-spec/21-app/30-root-cause-analysis-and-fix/` and `.ai-memory/plans/pending/30-root-cause-analysis-and-fix.md`. | `02-spec/`, `.ai-memory/` | Completed |
| **Subtask-02** | `02-prompts-and-v6-audit.md` | Audit execute prompts, V6 templates, and skills across `coding-guidelines` and eliminate all directives/examples prompting `commit` as a module name. | `01-prompts/`, `.agents/skills/` | Completed |
| **Subtask-03** | `03-coding-guidelines-and-rules.md` | Update `AGENTS.md` and `.agents/rules/` to ban `commit` / `commits` as module names and document Title Case / lowercase valid module rules. | `AGENTS.md`, `.agents/rules/` | Completed |
| **Subtask-04** | `04-gitmap-commit-formatter-engine.md` | Patch GitMap `cli/cmdcommit/commit_formatter.go` to treat `COMMIT`, `COMMITS`, `GITCOMMIT` as redundant type tokens, stripping them automatically and allowing Title Case / lowercase modules. | `d:\work\gitmap\cli\cmdcommit\` | Completed |
| **Subtask-05** | `05-verification-and-minor-release.md` | Verify all test suites, compile/deploy `gitmap`, verify clean dry-run formatting, and execute minor release ceremony. | Multi-repo | Completed |

---

## 3. Verification & Acceptance Criteria
- [x] Zero instances of `commit - ` or `(commit)` as module recommendations in V6 prompts or guidelines.
- [x] `AGENTS.md` Section 10 explicitly mandates that module names reflect functional subsystems, never `commit`, and permits Title Case (`Ui`, `Auth`) or lowercase (`ui`, `auth`).
- [x] `gitmap cpc -n "commit - update readme"` emits `Chore: update readme` (pseudo-module stripped).
- [x] `gitmap cpf -n "commit - cli - add flag"` emits `Feature: cli - add flag` (nested module promoted).
- [x] `gitmap cpf -n "Auth - support jwt"` emits `Feature: Auth - support jwt` (Title Case preserved).
- [x] `gitmap` test suite in `cli/cmdcommit` passes 100%.
- [x] Release ceremony with minor version bump completed (`coding-guidelines` v6.85.0 & `gitmap` v6.619.0).
