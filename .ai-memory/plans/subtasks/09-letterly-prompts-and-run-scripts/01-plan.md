# Execution Plan: Letterly Prompts Modernization, CI/CD Fix GitMap Release & Run Script Architecture

> **Plan Slug:** `09-letterly-prompts-and-run-scripts`  
> **Architecture Spec:** `02-spec/21-app/05-letterly-prompts-and-run-scripts/01-architecture-spec.md`  
> **Database:** `.ai-memory/temp-agents/10-letterly-prompts-and-run-scripts/agent-task.db`

---

## Subtask Breakdown

### Subtask 1: Spec & Plan Authoring (DONE)
- Author `02-spec/21-app/05-letterly-prompts-and-run-scripts/01-architecture-spec.md`.
- Author `.ai-memory/plans/subtasks/09-letterly-prompts-and-run-scripts/01-plan.md`.
- Initialize SQLite task database.

### Subtask 2: Letterly Prompts (`-letterly.md`) & Skills Modernization
- Rename/update all 7 prompts in `01-prompts/22-letterly/` to include `-letterly.md` suffix:
  - `01-mobile-letterly.md`: Strict single-line paragraph starting with `# High Priority Instruction: `, no `[/goal]` or `[/learn]` prefix, ending with `- must follow the skill [execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`.
  - `02-desktop-letterly.md`: Strict 5-block `execute-n-steps` format, starting with `# High Priority Instruction`, action item 1 as `1. Write spec and plan first`, ending with `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`.
  - `03-execute-n-steps-letterly.md`: Refined, articulate introduction while keeping user voice and format.
  - `04-plan-letterly.md`: Follows `execute-n-steps` format, action item 1 `Write the plan and architectural spec first`, action item 2 `Write the first step on task decomposition`, ending with `[plan-spec-steps-v2](file;.agents/skills/plan-spec-steps-v2)`.
  - `05-release-letterly.md`: Follows `execute-n-steps` format, action item 1 `Write spec and plan first`, action item 2 `Execute minor bump script 37-bump-version.py -t minor`, ending with `[minor-bump](file;.agents/skills/minor-bump)`.
  - `06-cicd-fix-release-letterly.md`: Follows `execute-n-steps` format, ending with `[ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release)`.
  - `07-mobile-cicd-fix-letterly.md`: Preserves mobile one-liner format, ending with `[ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release)`.
  - Update `01-prompts/22-letterly/readme.md`.
- Purge older skills (`desktop`, `mobile`, `execute-n-steps`) from `.agents/skills/` and `.cursor/skills/`.
- Synchronize Letterly skills in `.agents/skills/` and `.cursor/skills/`.

### Subtask 3: CI/CD Fix GitMap Release (`11-ci-cd-fix-gitmap-release.md`) & Skills
- Rename `01-prompts/16-ci-cd/11-ci-cd-fix-gitmap.md` to `11-ci-cd-fix-gitmap-release.md`.
- Update prompt content to prominently detail the minor bump script ceremony (`python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>"`), release branch option, atomic GitMap commit, and companion skill linkages (`ci-cd-fix-with-release`, `minor-bump`).
- Create `.agents/skills/ci-cd-fix-gitmap-release/skill.md` and `.cursor/skills/ci-cd-fix-gitmap-release/skill.md`.
- Remove obsolete `ci-cd-fix-gitmap` skills or provide aliases.
- Update `01-prompts/16-ci-cd/readme.md`.

### Subtask 4: Execute Run Prompt (`14-run.md`), Create Run PS1 Prompt (`04-create-run-ps1-file.md`), and Companion Skills
- Create `01-prompts/14-execute/14-run.md`:
  - Dedicated prompt to execute `.\run.ps1` (or `./run.sh`).
  - Pre-flight checks and fallback to installer.
  - Create companion skills `.agents/skills/run/skill.md` and `.cursor/skills/run/skill.md`.
- Update `01-prompts/16-ci-cd/04-cicd-run-ps1.md` -> `04-create-run-ps1-file.md`:
  - Upgrade to V6 architecture (`N = 300, A = 2, H = 2, C = 30`).
  - Specify dynamic `run.ps1` with fallback to `local-install.ps1` / `local-install.sh` via GitMap or standalone installers.
  - Create companion skills `.agents/skills/create-run-ps1-file/skill.md` and `.cursor/skills/create-run-ps1-file/skill.md`.
- Update `01-prompts/14-execute/readme.md` and `01-prompts/16-ci-cd/readme.md`.

### Subtask 5: Verification, Linters, Cursor Sync & Atomic GitMap Commit
- Run `python linter-scripts/check-prompts-loaded.py --fix`.
- Run `python linter-scripts/check-relative-paths.py`.
- Run `python 03-ai-scripts/05-guideline-autofixer.py` on touched files.
- Verify complete symmetry between `.agents/skills/` and `.cursor/skills/`.
- Commit atomically via `gitmap cpf`.
