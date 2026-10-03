# Execution Plan: Letterly Skill Path Normalization & Cursor Skill Format Architecture

> **Plan Slug:** `10-letterly-skill-paths-and-cursor-prompts`  
> **Architecture Spec:** `02-spec/21-app/06-letterly-skill-paths-and-cursor-prompts/01-architecture-spec.md`  
> **Database:** `.ai-memory/temp-agents/12-letterly-skill-paths-and-cursor-prompts/agent-task.db`

---

## Subtask Breakdown

### Subtask 1: Spec & Plan Authoring (DONE)
- Author `02-spec/21-app/06-letterly-skill-paths-and-cursor-prompts/01-architecture-spec.md`.
- Author `.ai-memory/plans/subtasks/10-letterly-skill-paths-and-cursor-prompts/01-plan.md`.
- Populate SQLite task database.

### Subtask 2: Correct Relative Paths in Agent Letterly Prompts (`01-prompts/22-letterly/`)
- Update all 7 prompts in `01-prompts/22-letterly/` to replace `(file;.agents/skills/<name>)` with explicit folder and file paths:
  `[.agents/skills/<name>/skill.md](.agents/skills/<name>/skill.md)`
- Synchronize `.agents/skills/letterly-*/skill.md` with explicit folder and file paths.

### Subtask 3: Create Cursor Letterly Prompts Suite (`01-prompts/22-letterly/cursor/`)
- Create `01-prompts/22-letterly/cursor/` directory.
- Author all 7 Cursor prompts referencing `.cursor/skills/<name>/skill.md`:
  - `01-mobile-letterly-cursor.md`
  - `02-desktop-letterly-cursor.md`
  - `03-execute-n-steps-letterly-cursor.md`
  - `04-plan-letterly-cursor.md`
  - `05-release-letterly-cursor.md`
  - `06-cicd-fix-release-letterly-cursor.md`
  - `07-mobile-cicd-fix-letterly-cursor.md`
  - `readme.md` (Cursor directory index)
- Synchronize `.cursor/skills/letterly-*/skill.md` to reference `.cursor/skills/<name>/skill.md`.

### Subtask 4: Verification, Linters & Atomic GitMap Commit
- Run `python linter-scripts/check-prompts-loaded.py --fix`.
- Run `python linter-scripts/check-relative-paths.py`.
- Run `python 03-ai-scripts/05-guideline-autofixer.py` on touched files.
- Commit atomically via `gitmap cpf`.
