# Execution Plan: Mobile Prompt Input Placement & Letterly Format Modernization

> Plan ID: `28-mobile-prompt-input-placement`
> Specification: `02-spec/21-app/23-mobile-prompt-input-placement/`
> Status: Completed
> Concurrency Engine: Parameter-Driven V6 Engine (`N=300`, `A=2`, `H=2`, `C=30`)

---

## 1. Objectives & Scope

1. **Reposition User Input in Mobile Formatters**:
   - Repositioned `${Input Text Verbatim}` in all mobile Letterly formatters and companion skills so that it goes **IMMEDIATELY AFTER** the header colon:
     `# <task title> - high priority instruction, non-negotiable task (task: <task title>, slug: <task-slug>): ${Input Text Verbatim}; `
   - Guaranteed the user's core prompt input is never buried or wasted beneath repetitive boilerplate clauses.
2. **Sequential Checklist Following Input**:
   - Structured the execution checklist to follow the input: pull first (`gitmap pull` fallback `git pull`), write spec under `02-spec/21-app/<slug>/`, search codebase exclusively via GitMap, relative paths mandate, extracted discrete tasks, and mandatory skill trigger.
3. **Strict Single-Line Invariant**:
   - Preserved strictly ONE continuous single-line paragraph with ZERO newlines, zero line gaps, and zero conversational filler across all mobile prompt formatters and companion skills.
4. **Target Files Modernized**:
   - `01-prompts/22-letterly/01-mobile-letterly.md`
   - `01-prompts/22-letterly/07-mobile-cicd-fix-letterly.md`
   - `01-prompts/23-cursor-prompts/01-mobile-letterly-cursor.md`
   - `01-prompts/23-cursor-prompts/07-mobile-cicd-fix-letterly-cursor.md`
   - `.agents/skills/letterly-mobile/skill.md`
   - `.agents/skills/letterly-mobile-cicd-fix/skill.md`
   - `.cursor/skills/letterly-mobile/skill.md`
   - `.cursor/skills/letterly-mobile-cicd-fix/skill.md`
5. **Validation, Verification & Release Ceremony**:
   - Verified formatters against `test_check_letterly_format.py` (8/8 passed).
   - Verified relative path hygiene across 3,719 tracked files with zero violations.
   - Verified prompt indexing with `check-prompts-loaded.py`.
   - Staged and committed atomically via GitMap (`gitmap cpf`).
   - Minor release ceremony to `v6.78.0`.

---

## 2. Completed Subtasks

- [x] Subtask 1: Spec and Architecture authoring in `02-spec/21-app/23-mobile-prompt-input-placement/` (`01-architecture-spec.md`, `readme.md`) and registration in `02-spec/21-app/readme.md`.
- [x] Subtask 2: Mobile Letterly formatters and companion skills upgrade (`01-mobile-letterly.md`, `07-mobile-cicd-fix-letterly.md`, `.agents/skills/letterly-mobile/`, `.agents/skills/letterly-mobile-cicd-fix/`).
- [x] Subtask 3: Cursor mirrors upgrade (`01-mobile-letterly-cursor.md`, `07-mobile-cicd-fix-letterly-cursor.md`, `.cursor/skills/letterly-mobile/`, `.cursor/skills/letterly-mobile-cicd-fix/`).
- [x] Subtask 4: Linter validation, relative path checks, plan consolidation, atomic commit via GitMap, and minor version release ceremony.
