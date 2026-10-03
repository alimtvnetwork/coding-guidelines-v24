# Execution Plan: AI Verification Engine & Letterly Formatter Extensions

> **Plan Slug:** `11-ai-verification-and-letterly-extensions`  
> **Architecture Spec:** `02-spec/21-app/08-ai-verification-and-letterly-extensions/01-architecture-spec.md`  
> **Database:** `.ai-memory/temp-agents/13-ai-verification-and-letterly-extensions/agent-task.db`

---

## Subtask Breakdown

### Subtask 1: Spec Authoring & SQLite Setup (DONE)
- Author `02-spec/21-app/08-ai-verification-and-letterly-extensions/01-architecture-spec.md`.
- Author `.ai-memory/plans/subtasks/11-ai-verification-and-letterly-extensions/01-plan.md`.
- Initialize `.ai-memory/temp-agents/13-ai-verification-and-letterly-extensions/agent-task.db`.

### Subtask 2: Normalize Existing Letterly Prompts & Skills to Native Trigger Syntax
- Update all 7 prompts in `01-prompts/22-letterly/` to use `@[.agents/skills/<name>]`.
- Update all 7 prompts in `01-prompts/22-letterly/cursor/` to use `@[.cursor/skills/<name>]`.
- Update all 7 companion skills in `.agents/skills/letterly-*` and `.cursor/skills/letterly-*`.

### Subtask 3: Implement Letterly Lifecycle Extensions (`run`, `execute-with-verification`, `execute-with-release`)
- Author prompts 08, 09, 10 in `01-prompts/22-letterly/`.
- Author prompts 08, 09, 10 in `01-prompts/22-letterly/cursor/`.
- Author companion skills:
  - `letterly-run` in `.agents/skills/` and `.cursor/skills/`
  - `letterly-execute-with-verification` in `.agents/skills/` and `.cursor/skills/`
  - `letterly-execute-with-release` in `.agents/skills/` and `.cursor/skills/`
- Update `01-prompts/22-letterly/readme.md` and `01-prompts/22-letterly/cursor/readme.md`.

### Subtask 4: Retrospective AI Verification Engine, Script & Skills
- Author `03-ai-scripts/47-retrospective-ai-verification.py`.
- Author `01-prompts/24-ai-verification/readme.md`.
- Author `01-prompts/24-ai-verification/01-retrospective-ai-verification.md` (V6 continuous loop architecture).
- Author companion skills `.agents/skills/ai-verification/skill.md` and `.cursor/skills/ai-verification/skill.md`.
- Update `01-prompts/readme.md`.

### Subtask 5: Verification, Linters & Atomic GitMap Commit
- Run `python linter-scripts/check-prompts-loaded.py --fix`.
- Run `python linter-scripts/check-relative-paths.py`.
- Run `python 03-ai-scripts/47-retrospective-ai-verification.py`.
- Commit atomically via `git add -A; gitmap cpf`.
