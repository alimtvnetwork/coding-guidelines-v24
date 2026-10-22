# Specification: Execution Prompt Standards & Plan Motto Modernization

## 1. System Overview & Core Rationale

### 1.1 Background & Context
Across previous iterations of prompt templates, skills, and desktop Letterly formatters, the standard directive for invoking the planning slash command was phrased around credit minimization:
```markdown
- [/plan](slashCommand;plan) first before doing work to reduce credits.
```
or variations such as:
```markdown
- [/plan](slashCommand;plan) first before doing the work to reduce the credits.
```

While the original intention was to deter AI agents from spinning in rapid unguided execution loops that exhaust user tokens and billing credits, phrasing the mandate as "reduce credits" introduces subtle yet detrimental operational defects:
1. **Misaligned Incentive (Scarcity Mindset):** Autonomous agents primed with "reduce credits" tend to truncate architectural analysis, generate shallow specifications, avoid necessary subagent spawning, and rush into superficial edits under the false assumption that less output equates to success.
2. **Contradiction with High-Authority Engineering:** The repository's foundational engineering rule dictates: *"Going deep IS the job. If you are not going deep, you are not doing the job."* Cost-skimping language contradicts the mandate for exhaustive 400-step execution, comprehensive test coverage, and bounded tree swarm hierarchy.
3. **The Real Purpose of Planning:** Planning in autonomous agent workflows exists to provide **clarity, architectural bounding, dependency resolution, and friction elimination**. A well-planned task prevents rework, eliminates merge collisions, ensures discrete file ownership, and produces clean atomic deliverables on the first pass.

### 1.2 Canonical Motto Re-Anchoring
The execution prompt standard is officially modernized and re-anchored from credit minimization to engineering excellence and developer friction reduction:

```markdown
- [/plan](slashCommand;plan) first before doing work for improved work or less friction.
```

### 1.3 Core Principles of Modernized Execution
- **Clarity Over Cost:** The primary metric of success for upfront planning is eliminating execution friction and rework.
- **Strict IDE Skill Link Syntax:** All slash commands retain the standardized IDE-detectable semicolon syntax: `[/plan](slashCommand;plan)` and `[/learn](slashCommand;learn)`.
- **Slot Placement Invariant:** Slash command suggestions belong strictly under `## Additional Instructions` at the bottom of execution prompts. They remain strictly forbidden in retrospective verification audit prompts.
- **De-Coupled Execution Pipeline:** Upfront planning is an isolated Phase 1 milestone. It is never an excuse to halt execution prematurely or request user permission before completing implementation.

---

## 2. De-Coupling of Planning from Execution Steps

### 2.1 The Two-Phase Execution Boundary
A recurrent failure mode in autonomous agents is conflating *planning* with *turn closure*: an agent authors a plan in `.ai-memory/plans/` and then halts, asking the user "Should I proceed with the plan?". This behavior violates Section 3 of `AGENTS.md` and Rule R6 of the V6 execution engine.

To maintain strict operational momentum:
1. **Phase 1: Planning, Research & Specification (Steps 1–100):**
   - Autonomous lead agent spawns read-only discovery subagents (`TypeName: "research"`).
   - Ingests repository context, reads `what-to-read.md`, and confirms clean git tree.
   - Authors canonical spec under `02-spec/21-app/<slug>/` and task plan in `.ai-memory/plans/<slug>.md`.
   - Populates subtask files under `.ai-memory/plans/subtasks/<slug>/`.
   - Registers tasks in GitMap agent engine (`gitmap agent task init`, `gitmap agent subtask add`).
2. **Phase 2 & Phase 3: Bounded Implementation & Verification (Steps 101–400):**
   - Lead agent unconditionally transitions to implementation without asking permission.
   - Spawns worker subagents with strictly disjoint single-file bounding boxes.
   - Validates code changes with targeted local linters and unit tests (strict no-build / no-test during intermediate phases, run targeted checks at verification gate).
   - Consolidates ledger and subtasks.
   - Performs retrospective AI verification audit.
   - Executes single atomic GitMap commit and remote push (`gitmap cpf "<module> - <summary>"`).

### 2.2 Operational Distinction: Slash Command vs. Execution Steps
- `[/plan](slashCommand;plan)` is a prompt decorator designed to instruct IDE chat interfaces to configure reasoning depth and present the roadmap cleanly to the human user.
- It is decoupled from the internal step count of the executing agent. An agent receiving `[/plan](slashCommand;plan)` must treat the planning phase as the foundation for the upcoming execution, immediately flowing into code modification within the same turn or via reactive subagent dispatch.

---

## 3. Comprehensive Audit Matrix (28 Occurrences)

The codebase audit identified 28 occurrences of credit reduction language across prompt templates, IDE formatters, generic Muse templates, test suites, and historical planning artifacts.

| # | File Path | Line | Current Verbatim Text | Proposed Canonical Replacement | Scope Category |
|---|---|---|---|---|---|
| 1 | `01-prompts/22-letterly/02-desktop-letterly.md` | 52 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Active Prompt |
| 2 | `01-prompts/22-letterly/03-execute-n-steps-letterly.md` | 41 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Active Prompt |
| 3 | `01-prompts/22-letterly/04-plan-letterly.md` | 50 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Active Prompt |
| 4 | `01-prompts/22-letterly/05-release-letterly.md` | 52 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Active Prompt |
| 5 | `01-prompts/22-letterly/06-cicd-fix-release-letterly.md` | 54 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Active Prompt |
| 6 | `01-prompts/22-letterly/08-run-letterly.md` | 48 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Active Prompt |
| 7 | `01-prompts/22-letterly/09-execute-with-verification-letterly.md` | 48 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Active Prompt |
| 8 | `01-prompts/22-letterly/10-execute-with-release-letterly.md` | 50 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Active Prompt |
| 9 | `01-prompts/22-letterly/11-git-pull-first-letterly.md` | 46 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Active Prompt |
| 10 | `01-prompts/23-cursor-prompts/02-desktop-letterly-cursor.md` | 52 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Cursor Mirror |
| 11 | `01-prompts/23-cursor-prompts/03-execute-n-steps-letterly-cursor.md` | 41 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Cursor Mirror |
| 12 | `01-prompts/23-cursor-prompts/04-plan-letterly-cursor.md` | 50 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Cursor Mirror |
| 13 | `01-prompts/23-cursor-prompts/05-release-letterly-cursor.md` | 52 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Cursor Mirror |
| 14 | `01-prompts/23-cursor-prompts/06-cicd-fix-release-letterly-cursor.md` | 54 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Cursor Mirror |
| 15 | `01-prompts/23-cursor-prompts/08-run-letterly-cursor.md` | 48 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Cursor Mirror |
| 16 | `01-prompts/23-cursor-prompts/09-execute-with-verification-letterly-cursor.md` | 48 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Cursor Mirror |
| 17 | `01-prompts/23-cursor-prompts/10-execute-with-release-letterly-cursor.md` | 50 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Cursor Mirror |
| 18 | `01-prompts/23-cursor-prompts/11-git-pull-first-letterly-cursor.md` | 46 | `- [/plan](slashCommand;plan) first before doing work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Cursor Mirror |
| 19 | `01-prompts/27-muse-prompts/02-muse-generic.md` | 57 | `- [/plan](slashCommand;plan) first before doing the work to reduce credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Active Muse Template |
| 20 | `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md` | 667 | `...strictly wait or sleep based on etaSeconds to eliminate credit waste.` | `...strictly wait or sleep based on etaSeconds to prevent unneeded API thrashing.` | Core V6 Prompt |
| 21 | `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` | 695 | `...strictly wait or sleep based on etaSeconds to eliminate credit waste.` | `...strictly wait or sleep based on etaSeconds to prevent unneeded API thrashing.` | Core V6 Skill |
| 22 | `linter-scripts/tests/test_check_letterly_format.py` | 52 | `- [/plan](slashCommand;plan) first before doing the work to reduce the credits.` | `- [/plan](slashCommand;plan) first before doing work for improved work or less friction.` | Test Fixture |
| 23 | `02-spec/21-app/letterly-prompt-cleanup-and-fleet-sync/01-architecture-spec.md` | 51 | `- [/plan](slashCommand;plan) first before doing the work to reduce the credits.` | Referenced historical spec (Record in audit matrix) | Historical Spec |
| 24 | `02-spec/21-app/15-muse-prompts-execute-in-a-step/02-new-prompt-and-template-changes.md` | 91 | `- [/plan](slashCommand;plan) first before doing the work to reduce the credits.` | Referenced historical spec (Record in audit matrix) | Historical Spec |
| 25 | `02-spec/21-app/gitmap-update-and-cicd-check/readme.md` | 37 | `- [/plan](slashCommand;plan) first before doing the work to reduce the credits.` | Referenced historical readme (Record in audit matrix) | Historical Spec |
| 26 | `02-spec/21-app/11-gitmap-training-and-letterly-golden-prompt-sync/01-architecture-spec.md` | 180, 213, 235 | `[/plan](slashCommand;plan) first before doing the work to reduce the credits.` | Referenced historical spec (Record in audit matrix) | Historical Spec |
| 27 | `.ai-memory/plans/sequence-and-backup-testing.md` | 40 | `- /plan first before doing the work to reduce the credits.` | Referenced historical plan (Record in audit matrix) | Historical Plan |
| 28 | `.ai-memory/plans/completed/18-gitmap-search-subagent-enforcement-and-sqlite-schema.md` | 35 | `/plan first before doing the work to reduce the credits.` | Referenced historical plan (Record in audit matrix) | Historical Plan |

---

## 4. Exact Before / After Diffs

### 4.1 Letterly Formatters (`01-prompts/22-letterly/`)
Applied across files 02, 03, 04, 05, 06, 08, 09, 10, 11:
```diff
 ## Additional Instructions

-- [/plan](slashCommand;plan) first before doing work to reduce credits.
+- [/plan](slashCommand;plan) first before doing work for improved work or less friction.
 - [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
```

### 4.2 Cursor Mirror Formatters (`01-prompts/23-cursor-prompts/`)
Applied across files 02, 03, 04, 05, 06, 08, 09, 10, 11:
```diff
 ## Additional Instructions

-- [/plan](slashCommand;plan) first before doing work to reduce credits.
+- [/plan](slashCommand;plan) first before doing work for improved work or less friction.
 - [/learn](slashCommand;learn) from [gitmap](file;.cursor/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
```

### 4.3 Muse Generic Template (`01-prompts/27-muse-prompts/02-muse-generic.md`)
```diff
 ## Additional Instructions

 - Execute within the 400-step lifecycle budget: Phase 1 (Steps 1..100) planning and spec authoring, Phase 2 (Steps 101..250) sanitization, archival, and prerequisites, Phase 3 (Steps 251..400) worker swarm implementation, targeted checks, and retrospective verification.
-- [/plan](slashCommand;plan) first before doing the work to reduce credits.
+- [/plan](slashCommand;plan) first before doing work for improved work or less friction.
 - [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
```

### 4.4 Letterly Linter Test Fixture (`linter-scripts/tests/test_check_letterly_format.py`)
```diff
 ## Additional Instructions

-- [/plan](slashCommand;plan) first before doing the work to reduce the credits.
+- [/plan](slashCommand;plan) first before doing work for improved work or less friction.
 - [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
```

### 4.5 V6 Execution Engine & Skill Preamble
In `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md` and `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`:
```diff
-- [ ] NO UNREQUESTED CI/CD VERIFICATION & NO RAPID POLLING (TOTAL BAN): Do CI/CD verification only if asked; if not asked explicitly, don't do it (never monitor, wait for, or verify remote pipelines via `gitmap pe`, `gitmap pe -t`, or GitHub Actions after push unless explicitly commanded). If CI/CD verification is explicitly requested, never query or loop rapidly (`gh run view` in tight loops); agents must query pipeline state using GitMap Pipeline-AI (`gitmap pipeline-ai status --json` or `gitmap pl-ai status -t <sec>`) and strictly wait or sleep based on `etaSeconds` to eliminate credit waste.
+- [ ] NO UNREQUESTED CI/CD VERIFICATION & NO RAPID POLLING (TOTAL BAN): Do CI/CD verification only if asked; if not asked explicitly, don't do it (never monitor, wait for, or verify remote pipelines via `gitmap pe`, `gitmap pe -t`, or GitHub Actions after push unless explicitly commanded). If CI/CD verification is explicitly requested, never query or loop rapidly (`gh run view` in tight loops); agents must query pipeline state using GitMap Pipeline-AI (`gitmap pipeline-ai status --json` or `gitmap pl-ai status -t <sec>`) and strictly wait or sleep based on `etaSeconds` to prevent unneeded API thrashing and eliminate execution friction.
```

---

## 5. Verification Contracts & Quality Gates

### 5.1 Mechanical Linter Compatibility
The mechanical linter `linter-scripts/check-letterly-format.py` validates the presence of the `/plan` directive via substring inspection:
```python
if "[/plan](slashCommand;plan)" not in text:
    violations.append("`## Additional Instructions` is missing the /plan directive")
```
Because the check targets `[/plan](slashCommand;plan)` rather than the trailing phrase, updating the suffix to `first before doing work for improved work or less friction.` maintains 100% backward and forward compatibility with the validation script.

### 5.2 Test Suite Execution Gate
Run the pytest suite locally to verify the updated golden fixture:
```bash
gitmap py -m pytest linter-scripts/tests/test_check_letterly_format.py -v
```
All unit tests in `test_check_letterly_format.py` must pass with zero failures.

### 5.3 Relative Git Paths Mandate
In strict accordance with Section 5 of `AGENTS.md`:
- All file links and paths must remain strictly relative to the repository root.
- Zero absolute filesystem paths (e.g. `C:\...`, `/home/...`).
- Zero `file:///` URIs anywhere in markdown, specs, plans, comments, or commit messages.
- Verification command:
```bash
gitmap py linter-scripts/check-relative-paths.py .
```

### 5.4 Multi-Repository Synchronization Boundary
Under Section 11 of `AGENTS.md`:
- Changes to `01-prompts/22-letterly/` and `01-prompts/23-cursor-prompts/` are canonical core assets subject to synchronization across the connected repository fleet.
- Application-specific specs under `02-spec/21-app/` are strictly excluded from fleet synchronization (Spec 21 Exclusion Mandate).
- Synchronization script `03-ai-scripts/38-sync-prompts-skills-scripts.py` handles distribution to connected worktrees while respecting all 5 non-negotiable boundaries.