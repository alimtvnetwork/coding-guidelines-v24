# Subtask Plan: Letterly & Cursor Prompts Refactoring, IDE Syntax Standards & Skills Synchronization

> **Subtask Plan ID:** `.ai-memory/plans/subtasks/12-multi-repo-sync-engine-and-prompt-upgrades/02-letterly-and-cursor-subtask.md`  
> **Parent Specification:** `02-spec/21-app/09-multi-repo-sync-engine-and-prompt-upgrades/02-letterly-and-cursor-spec.md`  
> **Parent Architecture Spec:** `02-spec/21-app/09-multi-repo-sync-engine-and-prompt-upgrades/01-architecture-spec.md`  
> **Status:** READY  
> **Target Subsystems:** `01-prompts/22-letterly/`, `01-prompts/23-cursor-prompts/`, `.agents/skills/letterly-*/`, `.cursor/skills/letterly-*/`

---

## 1. Overview & Scope

This subtask plan details the exact, step-by-step execution roadmap to refactor the entire Letterly and Cursor prompt ecosystem. The work is decomposed into 6 discrete, bounded phases:

1. **Phase 1:** Directory restructuring and migration of Cursor prompts from `01-prompts/22-letterly/cursor/` to dedicated `01-prompts/23-cursor-prompts/`.
2. **Phase 2:** Refactoring all 10 Letterly prompt formatters under `01-prompts/22-letterly/`.
3. **Phase 3:** Refactoring all 10 Cursor prompt formatters under `01-prompts/23-cursor-prompts/`.
4. **Phase 4:** Updating and synchronizing all 10 companion skills under `.agents/skills/letterly-*/`.
5. **Phase 5:** Updating and synchronizing all 10 companion skills under `.cursor/skills/letterly-*/`.
6. **Phase 6:** Automated validation, path linting, and quality verification.

---

## 2. Phase 1: Directory Restructuring & Legacy Deprecation

### Task 1.1: Scaffolding `01-prompts/23-cursor-prompts/`
- Create target directory: `01-prompts/23-cursor-prompts/`.
- Ensure directory adheres strictly to lowercase naming conventions.

### Task 1.2: Relocate Prompt Files to `01-prompts/23-cursor-prompts/`
Migrate files from `01-prompts/22-letterly/cursor/` to `01-prompts/23-cursor-prompts/` following standardized naming:
- `01-mobile-letterly-cursor.md` -> `01-mobile-cursor.md`
- `02-desktop-letterly-cursor.md` -> `02-desktop-cursor.md`
- `03-execute-n-steps-letterly-cursor.md` -> `03-execute-n-steps-cursor.md`
- `04-plan-letterly-cursor.md` -> `04-plan-cursor.md`
- `05-release-letterly-cursor.md` -> `05-release-cursor.md`
- `06-cicd-fix-release-letterly-cursor.md` -> `06-cicd-fix-release-cursor.md`
- `07-mobile-cicd-fix-letterly-cursor.md` -> `07-mobile-cicd-fix-cursor.md`
- `08-run-letterly-cursor.md` -> `08-run-cursor.md`
- `09-execute-with-verification-letterly-cursor.md` -> `09-execute-with-verification-cursor.md`
- `10-execute-with-release-letterly-cursor.md` -> `10-execute-with-release-cursor.md`

### Task 1.3: Deprecate & Remove Legacy Directory
- Delete obsolete nested directory: `01-prompts/22-letterly/cursor/`.
- Verify no orphaned references remain in `01-prompts/22-letterly/`.

### Task 1.4: Author Directory Documentation
- Author `01-prompts/23-cursor-prompts/readme.md` indexing the 10 Cursor prompts with `.cursor/skills/` references.
- Update `01-prompts/22-letterly/readme.md` documenting the separation of Letterly and Cursor suites.

---

## 3. Phase 2: Refactoring All 10 Letterly Prompts (`01-prompts/22-letterly/`)

Each Letterly prompt must be refactored to enforce IDE-detectable syntax, standard spec/plan paths, and proper skill links.

### Task 2.1: `01-prompts/22-letterly/01-mobile-letterly.md`
- Maintain single continuous line output starting with `# High Priority Instruction: `.
- Update suffix link to: ` - must follow the skill [execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`.

### Task 2.2: `01-prompts/22-letterly/02-desktop-letterly.md`
- Under `# Actionable Items Must Follow Non-Negotiable`, set Item 1:  
  `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)`.
- Update suffix link to: `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`.
- Enforce slash commands: `[/learn](slashCommand;learn)` and `[/plan](slashCommand;plan)`.

### Task 2.3: `01-prompts/22-letterly/03-execute-n-steps-letterly.md`
- Output block line 1 MUST start with `[/plan](slashCommand;plan)`.
- Actionable Item 1 MUST be:  
  `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)`.
- Final Actionable Item MUST be:  
  `N. Run retrospective AI verification prompt/script (01-retrospective-ai-verification.md / 03-ai-scripts/47-retrospective-ai-verification.py) or skill [ai-verification](file;.agents/skills/ai-verification) to audit specs, touched files, code quality, and CI/CD status upon task completion`.
- Suffix link: `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`.
- Additional instructions: `learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.`.

### Task 2.4: `01-prompts/22-letterly/04-plan-letterly.md`
- Actionable Item 1: `1. Write the plan and architectural spec first under 02-spec/21-app/<slug>/`.
- Actionable Item 2: `2. Write the task decomposition and bounded subtasks under .ai-memory/plans/subtasks/<slug>/`.
- Actionable Item 3: `3. Enforce strict no-build and no-test rules throughout the planning phase`.
- Suffix link: `[plan-spec-steps-v2](file;.agents/skills/plan-spec-steps-v2)`.
- Additional instructions: `learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.`.

### Task 2.5: `01-prompts/22-letterly/05-release-letterly.md`
- Actionable Item 1: `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)`.
- Release ceremony items: verify CI/CD via `gitmap pe -t`, execute bump via `python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>"`, atomic commit `gitmap cpf`.
- Suffix link: `[minor-bump](file;.agents/skills/minor-bump)`.
- Additional instructions: `learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.`.

### Task 2.6: `01-prompts/22-letterly/06-cicd-fix-release-letterly.md`
- Actionable Item 1: `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)`.
- 4-part RCA and surgical fix directives without disabling CI checks.
- Suffix link: `[ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release)`.
- Additional instructions: `learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.`.

### Task 2.7: `01-prompts/22-letterly/07-mobile-cicd-fix-letterly.md`
- Single continuous line format starting with `# High Priority Instruction: `.
- Suffix link: ` - must follow the skill [ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release)`.

### Task 2.8: `01-prompts/22-letterly/08-run-letterly.md`
- Actionable Item 1: `1. Inspect run.config.json for target service configuration and port mappings`.
- Actionable Item 2: `2. Execute run script (./run.ps1 or ./run.sh) with automatic dependency verification`.
- Actionable Item 3: `3. If dependencies or compilers are missing, fall back to gitmap aum install or ./local-install.ps1, then auto-rerun`.
- Suffix link: `[run](file;.agents/skills/run)`.
- Additional instructions: `learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.`.

### Task 2.9: `01-prompts/22-letterly/09-execute-with-verification-letterly.md`
- Output block line 1 MUST start with `[/plan](slashCommand;plan)`.
- Actionable Item 1: `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)`.
- Final Actionable Item: `Run retrospective AI verification prompt/script (01-retrospective-ai-verification.md / 03-ai-scripts/47-retrospective-ai-verification.py) or skill [ai-verification](file;.agents/skills/ai-verification) to audit specs, touched files, code quality, and CI/CD status upon task completion`.
- Invocation suffix MUST mandate both:  
  `Must follow and spawn agent using [execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6) and [ai-verification](file;.agents/skills/ai-verification)`.
- Additional instructions: `learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.`.

### Task 2.10: `01-prompts/22-letterly/10-execute-with-release-letterly.md`
- Output block line 1 MUST start with `[/plan](slashCommand;plan)`.
- Actionable Item 1: `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)`.
- Actionable items: CI/CD verification via `gitmap pe -t` and minor bump via `python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>"`.
- Suffix link: `[minor-bump](file;.agents/skills/minor-bump)`.
- Additional instructions: `learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.`.

---

## 4. Phase 3: Refactoring All 10 Cursor Prompts (`01-prompts/23-cursor-prompts/`)

Refactor all 10 Cursor prompts with strict parity to Phase 2, but referencing `.cursor/skills/` using `file;.cursor/skills/<skill-name>` syntax:

### Task 3.1: `01-prompts/23-cursor-prompts/01-mobile-cursor.md`
- Single continuous line format.
- Suffix: ` - must follow the skill [execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)`.

### Task 3.2: `01-prompts/23-cursor-prompts/02-desktop-cursor.md`
- Item 1: `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)`.
- Suffix: `[execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)`.

### Task 3.3: `01-prompts/23-cursor-prompts/03-execute-n-steps-cursor.md`
- Output block starts with `[/plan](slashCommand;plan)`.
- Item 1: `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)`.
- Final Item: `Run retrospective AI verification prompt/script (01-retrospective-ai-verification.md / 03-ai-scripts/47-retrospective-ai-verification.py) or skill [ai-verification](file;.cursor/skills/ai-verification) to audit specs, touched files, code quality, and CI/CD status upon task completion`.
- Suffix: `[execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)`.

### Task 3.4: `01-prompts/23-cursor-prompts/04-plan-cursor.md`
- Item 1: `1. Write the plan and architectural spec first under 02-spec/21-app/<slug>/`.
- Item 2: `2. Write the task decomposition and bounded subtasks under .ai-memory/plans/subtasks/<slug>/`.
- Suffix: `[plan-spec-steps-v2](file;.cursor/skills/plan-spec-steps-v2)`.

### Task 3.5: `01-prompts/23-cursor-prompts/05-release-cursor.md`
- Item 1: `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)`.
- Suffix: `[minor-bump](file;.cursor/skills/minor-bump)`.

### Task 3.6: `01-prompts/23-cursor-prompts/06-cicd-fix-release-cursor.md`
- Item 1: `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)`.
- Suffix: `[ci-cd-fix-gitmap-release](file;.cursor/skills/ci-cd-fix-gitmap-release)`.

### Task 3.7: `01-prompts/23-cursor-prompts/07-mobile-cicd-fix-cursor.md`
- Single continuous line format.
- Suffix: ` - must follow the skill [ci-cd-fix-gitmap-release](file;.cursor/skills/ci-cd-fix-gitmap-release)`.

### Task 3.8: `01-prompts/23-cursor-prompts/08-run-cursor.md`
- Inspect `run.config.json` and execute runner.
- Suffix: `[run](file;.cursor/skills/run)`.

### Task 3.9: `01-prompts/23-cursor-prompts/09-execute-with-verification-cursor.md`
- Output block starts with `[/plan](slashCommand;plan)`.
- Item 1: `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)`.
- Final Item: `Run retrospective AI verification prompt/script (01-retrospective-ai-verification.md / 03-ai-scripts/47-retrospective-ai-verification.py) or skill [ai-verification](file;.cursor/skills/ai-verification) to audit specs, touched files, code quality, and CI/CD status upon task completion`.
- Suffix: `Must follow and spawn agent using [execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6) and [ai-verification](file;.cursor/skills/ai-verification)`.

### Task 3.10: `01-prompts/23-cursor-prompts/10-execute-with-release-cursor.md`
- Output block starts with `[/plan](slashCommand;plan)`.
- Item 1: `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)`.
- Suffix: `[minor-bump](file;.cursor/skills/minor-bump)`.

---

## 5. Phase 4: Updating Companion Skills in `.agents/skills/letterly-*`

Update all 10 `skill.md` files in `.agents/skills/` to embed the exact prompt formatter instructions from `01-prompts/22-letterly/`:

| Subtask | Skill Folder | Target Prompt | Key Skill Link |
| :---: | :--- | :--- | :--- |
| **4.1** | `.agents/skills/letterly-mobile/skill.md` | `01-mobile-letterly.md` | `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)` |
| **4.2** | `.agents/skills/letterly-desktop/skill.md` | `02-desktop-letterly.md` | `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)` |
| **4.3** | `.agents/skills/letterly-execute-n-steps/skill.md` | `03-execute-n-steps-letterly.md` | `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)` + verification |
| **4.4** | `.agents/skills/letterly-plan/skill.md` | `04-plan-letterly.md` | `[plan-spec-steps-v2](file;.agents/skills/plan-spec-steps-v2)` |
| **4.5** | `.agents/skills/letterly-release/skill.md` | `05-release-letterly.md` | `[minor-bump](file;.agents/skills/minor-bump)` |
| **4.6** | `.agents/skills/letterly-cicd-fix-release/skill.md` | `06-cicd-fix-release-letterly.md` | `[ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release)` |
| **4.7** | `.agents/skills/letterly-mobile-cicd-fix/skill.md` | `07-mobile-cicd-fix-letterly.md` | `[ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release)` |
| **4.8** | `.agents/skills/letterly-run/skill.md` | `08-run-letterly.md` | `[run](file;.agents/skills/run)` |
| **4.9** | `.agents/skills/letterly-execute-with-verification/skill.md` | `09-execute-with-verification-letterly.md` | `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)` AND `[ai-verification](file;.agents/skills/ai-verification)` |
| **4.10** | `.agents/skills/letterly-execute-with-release/skill.md` | `10-execute-with-release-letterly.md` | `[minor-bump](file;.agents/skills/minor-bump)` |

---

## 6. Phase 5: Updating Companion Skills in `.cursor/skills/letterly-*`

Update all 10 `skill.md` files in `.cursor/skills/` to embed the exact prompt formatter instructions from `01-prompts/23-cursor-prompts/`:

| Subtask | Skill Folder | Target Prompt | Key Skill Link |
| :---: | :--- | :--- | :--- |
| **5.1** | `.cursor/skills/letterly-mobile/skill.md` | `01-mobile-cursor.md` | `[execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)` |
| **5.2** | `.cursor/skills/letterly-desktop/skill.md` | `02-desktop-cursor.md` | `[execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)` |
| **5.3** | `.cursor/skills/letterly-execute-n-steps/skill.md` | `03-execute-n-steps-cursor.md` | `[execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)` + verification |
| **5.4** | `.cursor/skills/letterly-plan/skill.md` | `04-plan-cursor.md` | `[plan-spec-steps-v2](file;.cursor/skills/plan-spec-steps-v2)` |
| **5.5** | `.cursor/skills/letterly-release/skill.md` | `05-release-cursor.md` | `[minor-bump](file;.cursor/skills/minor-bump)` |
| **5.6** | `.cursor/skills/letterly-cicd-fix-release/skill.md` | `06-cicd-fix-release-cursor.md` | `[ci-cd-fix-gitmap-release](file;.cursor/skills/ci-cd-fix-gitmap-release)` |
| **5.7** | `.cursor/skills/letterly-mobile-cicd-fix/skill.md` | `07-mobile-cicd-fix-cursor.md` | `[ci-cd-fix-gitmap-release](file;.cursor/skills/ci-cd-fix-gitmap-release)` |
| **5.8** | `.cursor/skills/letterly-run/skill.md` | `08-run-cursor.md` | `[run](file;.cursor/skills/run)` |
| **5.9** | `.cursor/skills/letterly-execute-with-verification/skill.md` | `09-execute-with-verification-cursor.md` | `[execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)` AND `[ai-verification](file;.cursor/skills/ai-verification)` |
| **5.10** | `.cursor/skills/letterly-execute-with-release/skill.md` | `10-execute-with-release-cursor.md` | `[minor-bump](file;.cursor/skills/minor-bump)` |

---

## 7. Phase 6: Automated Validation, Path Linting & Verification

### Task 6.1: Run Relative Path Linter
- Execute:
  ```bash
  python 03-ai-scripts/07-relative-path-linter.py
  ```
- Verify zero absolute filesystem paths or file triple-slash URIs across all touched files.

### Task 6.2: Validate Lowercase Naming Hygiene
- Ensure every newly created and modified file strictly adheres to lowercase naming:
  - `01-prompts/23-cursor-prompts/*.md`
  - `01-prompts/22-letterly/*.md`
  - `.agents/skills/letterly-*/skill.md`
  - `.cursor/skills/letterly-*/skill.md`

### Task 6.3: Validate IDE Link Syntax
- Verify all skill links match `[<skill>](file;.agents/skills/<skill>)` or `[<skill>](file;.cursor/skills/<skill>)`.
- Verify all slash command links match `[/plan](slashCommand;plan)` or `[/learn](slashCommand;learn)`.
- Verify `03`, `09`, and `10` output templates start with `[/plan](slashCommand;plan)`.

### Task 6.4: Retrospective AI Verification Gate
- Run retrospective audit:
  ```bash
  python 03-ai-scripts/47-retrospective-ai-verification.py
  ```
- Confirm all acceptance criteria pass with 0 failures.
