# Subtask 03: Muse Prompt 400-Step Upgrade & Powerful Write-Execute Directives

- **Parent Task:** `skills-and-prompts-synchronization-and-configuration`
- **Scope:** Upgrade Muse prompts (`01-muse-master-prompt.md`, `02-muse-generic.md`) and skills (`.agents/skills/muse-master-prompt/skill.md`, `.cursor/skills/muse-master-prompt/skill.md`) to default to 400 steps with aggressive "write-and-execute" directives.
- **Status:** DONE

## Execution Summary & Evidence

1. **Muse Master Prompt (`01-prompts/27-muse-prompts/01-muse-master-prompt.md`):**
   - Configured `N = 400` default budget with phase partitioning:
     - `PHASE_1_BUDGET = 100` (Steps 1 .. 100: Discovery, Architecture Spec, and Lean Subtask Generation)
     - `PHASE_2_BUDGET = 150` (Steps 101 .. 250: Intermediate Archival, Repo-Cache Setup, Prerequisite Refactoring, File Staging)
     - `PHASE_3_BUDGET = 150` (Steps 251 .. 400: Worker Swarm Parallel Execution, Powerful Write-and-Execute Implementation, Targeted Quality Linting, Atomic Push)
   - Updated V6 Engine parameters report to `N=400`.
   - Updated task initialization preflight command: `gitmap agent task init --name "<task name>" --budget 400`.
   - Injected **Powerful Step Execution (Write-and-Execute Invariant)**: Each step and turn must write code/files immediately to disk and execute tools in the same turn without asking for permission.

2. **Muse Generic Prompt (`01-prompts/27-muse-prompts/02-muse-generic.md`):**
   - Verified `BUDGET = 400 (Steps: 100 Plan & Spec, 150 Subtasks & Prerequisites, 150 Execution & Verification)`.
   - Injected **Powerful Step Execution (Write-and-Execute Invariant)** into Section 2 and Section 8.

3. **Skills Mirror Synchronization:**
   - Synchronized `.agents/skills/muse-master-prompt/skill.md` with `N = 400`, `PHASE_1_BUDGET = 100`, `PHASE_2_BUDGET = 150`, `PHASE_3_BUDGET = 150`, `--budget 400`, and the write-and-execute invariant.
   - Synchronized `.cursor/skills/muse-master-prompt/skill.md` with identical 400-step directives and parameters.
