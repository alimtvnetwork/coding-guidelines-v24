# Architecture Specification: Skills and Prompts Synchronization & Configuration

- **Module / Task Slug:** `skills-and-prompts-synchronization-and-configuration`
- **Specification ID:** `02-spec/21-app/skills-and-prompts-synchronization-and-configuration/01-architecture-spec.md`
- **Plan File:** `.ai-memory/plans/skills-and-prompts-synchronization-and-configuration.md`
- **Status:** APPROVED & ACTIVE
- **Target Release:** `v6.82.0`

---

## 1. Executive Summary & Purpose

This specification governs two core directives:
1. **Prompts-to-Skills Parity & Migration Audit:** Perform a comprehensive audit confirming that every active prompt in `01-prompts/` has a corresponding, properly structured skill in `.agents/skills/` and `.cursor/skills/`.
2. **Muse Prompt 400-Step Lifecycle & High-Impact Execution Directives:** Update `01-prompts/27-muse-prompts/01-muse-master-prompt.md`, `02-muse-generic.md`, `.agents/skills/muse-master-prompt/skill.md`, and `.cursor/skills/muse-master-prompt/skill.md` to default to an aggressive 400-step budget (`N = 400`). Every step must enforce powerful, proactive action: authoring files, implementing code, and executing validations immediately without conversational delays or permission-seeking.
3. **Fleet Synchronization Ceremony:** Deploy all synchronized prompts, skills, and updated specifications across the 43 connected repositories using the compiled Go engine (`gitmap sync --workers 8`) and execute the minor release bump to `v6.82.0`.

---

## 2. 400-Step Budget Partitioning for Muse AI

The Muse AI execution engine must allocate the 400-step budget systematically:
- **`PHASE_1_BUDGET = 100` (Steps 1..100):** Discovery, prompt capture, architecture spec authoring (`02-spec/21-app/<slug>/`), and lean subtasks setup (`.ai-memory/plans/subtasks/<slug>/`).
- **`PHASE_2_BUDGET = 150` (Steps 101..250):** Intermediate archival, repo-cache setup, prerequisite refactoring, and file staging.
- **`PHASE_3_BUDGET = 150` (Steps 251..400):** Worker swarm parallel execution, write-and-execute implementation, targeted linter checks, and atomic commit push.

### High-Impact "Write and Execute" Directives
Each turn and step executed by Muse must be forceful and concrete:
- **Write Then Execute:** Never announce intentions in chat without tool execution. In the same turn, write the file (`write_to_file` / `replace_file_content`) and invoke verification tools.
- **Zero Permission Seeking:** Do not ask "Should I proceed?". Execute the plan autonomously to completion.
- **Evidence Contract:** Every claim of completion must cite tangible diffs, exit codes (`exit 0`), and verified paths.

---

## 3. Prompts-to-Skills Mapping Matrix

All canonical prompts under `01-prompts/` correspond 1:1 with Antigravity skills (`.agents/skills/`) and Cursor skills (`.cursor/skills/`):
- `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` ↔ `execute-parent-task-with-n-steps-v6`
- `01-prompts/15-cg-execute/01-execute-coding-guideline-fix.md` ↔ `execute-coding-guideline-fix`
- `01-prompts/24-sync/01-sync.md` ↔ `sync`
- `01-prompts/25-ai-verification/01-retrospective-ai-verification.md` ↔ `ai-verification`
- `01-prompts/26-gitmap/01-gitmap-core-engine.md` ↔ `gitmap`
- `01-prompts/27-muse-prompts/01-muse-master-prompt.md` ↔ `muse-master-prompt`

---

## 4. Non-Negotiable Boundaries for Fleet Sync

1. **Spec 21 Exclusion:** Never sync `02-spec/21-*` through `02-spec/25-*` to downstream repositories.
2. **Bump Script Protection:** Never overwrite downstream version bump scripts (`bump*`).
3. **Additive-Only AI Scripts:** Unmodified existing scripts updated; new scripts added cleanly.
4. **Memory & Plans Isolated:** Downstream `.ai-memory/plans/` and `.ai-memory/memory/` remain isolated.
5. **Zero Secrets Leakage:** No environment credentials or API keys transferred.
