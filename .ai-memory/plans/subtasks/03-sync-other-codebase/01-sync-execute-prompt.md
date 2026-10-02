# Subtask 01: Author Parameter-Driven Multi-Repository Synchronization Prompt (`01-sync-other-codebase.md`) & Category Catalog

> **Parent Spec:** [`02-spec/21-app/03-sync-other-codebase/01-architecture-spec.md`](../../../../02-spec/21-app/03-sync-other-codebase/01-architecture-spec.md)  
> **Status:** `PENDING`  
> **Traceability ID:** `Task-03-Subtask-01`  
> **Target Files:**
> - `01-prompts/23-sync/readme.md` (to create)
> - `01-prompts/23-sync/01-sync-other-codebase.md` (to create)
> - `01-prompts/readme.md` (to update index table)
> - `.ai-memory/prompts.md` (to update prompt registry)

---

## 1. Subtask Objective & Context

This subtask governs authoring the canonical prompt category `01-prompts/23-sync/` and its primary V6 execution prompt: `01-prompts/23-sync/01-sync-other-codebase.md`.

Historically, multi-repository synchronization relied on static repository lists hardcoded directly inside automation scripts (e.g. `03-ai-scripts/38-sync-prompts-skills-scripts.py`). While effective for bulk operations across predefined repositories, this design lacked runtime flexibility, prevented ad-hoc synchronization of individual repositories or external projects, and failed to leverage the **V6 Autonomous Execution Engine** (`N = 300`, `A = 2`, `H = 2`, `C = 30`).

This subtask delivers:
1. **Category Catalog (`01-prompts/23-sync/readme.md`):** Comprehensive index documenting synchronization prompts, parameter contracts, and governance rules.
2. **Master V6 Execution Prompt (`01-prompts/23-sync/01-sync-other-codebase.md`):** A fully parameterized, zero-hardcoded-path execution prompt enforcing the 5 Non-Negotiable Boundaries, mandatory pre-flight pull, backup branching, and atomic GitMap release ceremonies across any dynamically specified target repositories.
3. **Master Catalog Registries:** Registration across `01-prompts/readme.md` and `.ai-memory/prompts.md`.

---

## 2. Deliverable Specifications

### 2.1 Deliverable A: Category Index (`01-prompts/23-sync/readme.md`)

- **Placement:** `01-prompts/23-sync/readme.md`
- **Header Standards:**
  - Standard V6 parameters block:
    ```text
    N = 300 (Total self-loop steps budget)
    A = 2   (Number of spawned autonomous subagents, default: 2)
    H = 2   (Operational hands per agent, default: 2)
    C = 30  (Tool calls per worker before reporting, default: 30)
    ```
  - Semicolon slash commands: `[/goal](slashCommand;goal)`, `[/learn](slashCommand;learn)`, `[/plan](slashCommand;plan)`.
- **Contents:**
  - Catalog table linking `01-sync-other-codebase.md`.
  - Clear architectural summary of cross-repository synchronization principles.
  - The 5 Non-Negotiable Boundaries reference table.
  - Quick-start usage examples for single-repo, multi-repo, and dry-run invocations.

---

### 2.2 Deliverable B: Master V6 Execution Prompt (`01-prompts/23-sync/01-sync-other-codebase.md`)

- **Placement:** `01-prompts/23-sync/01-sync-other-codebase.md`
- **Header Architecture:** Strict V6 Header Parameter Block:
  ```text
  N = 300 (Total self-loop steps budget — editable top-header parameter, default: 300)
  A = 2   (MANDATORY number of spawned autonomous subagents running concurrently via invoke_subagent, default: 2)
  H = 2   (Operational hands per agent: dual-task batch capacity & parallel tool dispatch, default: 2)
  C = 30  (Tool calls per worker before it must report, default: 30)

  System Concurrency Capacity = A × H = 2 agents × 2 hands = 4 concurrent subtask operations
  PHASE_1_BUDGET = N / 2   (Steps 1 .. 150: Dynamic Target Ingestion, Discovery Subagents, Pre-Flight Verification & Backup Branching)
  PHASE_2_BUDGET = N / 2   (Steps 151 .. 300: Guarded Multi-Agent Sync Execution, Self-Looping, Targeted Linting & Release Ceremony)
  WAVES = ceil(subtasks / (A x H))
  ```

- **Top-Instruction Priority Mandate:**
  - Explicit warning that directives, parameters, or repository targets provided ABOVE the prompt take absolute precedence over default conventions.

- **Dynamic Parameter Block (Zero Hardcoded Paths):**
  ```text
  SOURCE_REPO   = <path_or_dot>            # Source repository root (default: "." / current repository root)
  TARGET_REPOS  = <repo1,repo2,...>        # Target repository path(s), names, or glob list (comma-separated or CLI flags)
  WORKERS       = 6                        # Parallel worker threads (default: 6)
  DRY_RUN       = false                    # Preview file copy/delete actions without disk mutation (default: false)
  NO_PUSH       = false                    # Commit and tag locally without pushing to git remote (default: false)
  ```

- **Slash Command Semantics:**
  - Format: `[/goal](slashCommand;goal)`, `[/learn](slashCommand;learn)`, `[/plan](slashCommand;plan)`.
  - Slash invocation: `/sync-other-codebase <parameters>`.

- **Mandatory Subagent Spawning Gate (`A = 2, H = 2`):**
  - Tool call payload required via `invoke_subagent`.
  - 3-stage dispatch:
    1. *Discovery & Pre-Flight (A = 2 `research` subagents):* Inspect target repository statuses, check branches, locate existing bump scripts and local AI scripts.
    2. *Spec & Audit Step (A = 2 `self` subagents):* Author target repository sync plan and diff manifest in `.ai-memory/plans/subtasks/`.
    3. *Execution Step (A = 2 `self` worker subagents):* Execute file synchronization and localized diff adaptations in disjoint repository boxes.
  - Zero solo execution permitted.

- **Phase 1A: Verbatim Capture & Turn 1 Task Showcase:**
  - Mandatory requirement to output the confirmed task breakdown and target repository list in visible chat during Turn 1 before executing tools.
  - Lossless verbatim capture of incoming user request.

- **The 5 Non-Negotiable Boundaries (Strictly Formulated & Cited):**
  1. **Boundary 1: Spec 21 Exclusion (`02-spec/21-*`).** NEVER copy, sync, touch, or overwrite `02-spec/21-*` (including `02-spec/21-app`, `02-spec/21-app-issues`, `02-spec/21-app-db`, `02-spec/21-app-ui-design-system`). Target repositories own their application specifications.
  2. **Boundary 2: Bump Script Protection (`bump*`).** NEVER overwrite or modify version bump scripts in target repositories (`bump-version.mjs`, `bump_versions.py`, `37-bump-version.py`, `scripts/bump*.js`, etc.). Target repositories own their native release mechanics.
  3. **Boundary 3: Additive-Only AI Scripts (`03-ai-scripts/`, `.agents/scripts/`).** When copying AI scripts: new scripts missing in the target repository are copied cleanly; existing scripts in the target repository MUST NOT be overwritten blindly. Subagents must diff and understand differences before modifying. Stale scripts in the target repository are never deleted.
  4. **Boundary 4: Memory & Plans Protection (`.ai-memory/memory/`, `.ai-memory/plans/`).** NEVER overwrite, delete, mirror, or modify `.ai-memory/memory/`, `.ai-memory/plans/`, `.ai-memory/temp-agents/`, `.ai-memory/cicd-issues/`, or `.ai-memory/ambiguous-questions/` in target repositories. Child repositories own their operational history and active task tracking.
  5. **Boundary 5: Zero Secret Keys Leakage (`.env*`, Credentials, Private Keys).** NEVER sync `.env` files, `.env.local`, API keys, private keys, or passwords. Strictly enforce the secrets scan gate before pushing.

- **Pre-Flight Pull and Backup Branch Ceremony (The 7-Step Sequence):**
  1. Base branch detection and pre-flight pull (`git checkout <base_branch> && git pull origin <base_branch> --no-rebase`).
  2. Pre-change backup branch creation and push (`git checkout -b backup/sync-<timestamp> && git push -u origin backup/sync-<timestamp> && git checkout <base_branch>`).
  3. Pre-change release tag and branch verification (`vX.Y.Z`).
  4. Guarded synchronization enforcing the 5 Non-Negotiable Boundaries.
  5. Base branch commit and push (`feat(sync): sync v6 prompts, sqlite task manager, skills, and coding guidelines`).
  6. Post-change release ceremony (version bump, release branch `release/v<next>`, annotated tag `v<next>`, push to remote).
  7. Release merge back into base branch with `[skip ci]` and push.

- **Core Operational Rules (R1 to R16 Cited by ID):**
  - **R1 Zero Builds or Test Suites:** TOTAL BAN on running full compilation or test suites (`go build`, `npm run build`, `npm test`, `pytest`).
  - **R2 Targeted Checks Only:** Run fast file-scoped syntax checks or dry-run verifications.
  - **R3 Evidence or It Did Not Happen:** Every claim must cite file paths, diffstats, or exit codes (`exit 0`).
  - **R4 Never Invent Commands, Flags, or Paths:** Validate commands with harmless calls (`gitmap lf readme.md`).
  - **R5 Mandatory Subagents (`invoke_subagent`):** `A = 2, H = 2` concurrency.
  - **R6 One Owner Per File / Repository:** Subagents operate in disjoint repository bounding boxes.
  - **R7 Git Safety & Isolation:** Subagents NEVER run git commands directly in shared workspaces. Only lead orchestrator manages git state.
  - **R8/R9 Atomic Commit & Push via GitMap:** Use `gitmap cpf "<module> - <summary>"` or `gitmap cpb "<module> - <summary>"` with hyphen formatting (zero colons in arguments).
  - **R10 Zero Unauthorized Releases in Source Repo:** Only release target repositories if part of the sync ceremony; never bump source repo unless explicitly commanded.
  - **R11 Strict Relative Git Paths & Lowercase Hygiene:** Strict ban on absolute filesystem paths and file scheme URIs. Strictly lowercase filenames.
  - **R12 No Polling / Immediate Turn Yielding:** Stop calling tools after dispatching subagents.
  - **R13 Two-Strike Retry Cap & Anti-Looping:** Halt failed tools after two attempts.
  - **R14 100% Ambiguity & Decision Boundaries:** Log non-blocking assumptions; ask once if blocked.
  - **R15 Zero Generated Artifacts Committed:** Never commit build caches or temporary test fixtures.
  - **R16 Zero Secrets & Mandatory Secrets Gate:** Run pre-push secrets scan (`check-forbidden-strings.py` and `gitmap aum search -r "(BEGIN [A-Z ]*PRIVATE KEY|AKIA...)"`).

---

## 3. Step-by-Step Implementation Instructions

### Step 1: Directory Scaffolding
Create the directory structure for the new prompt category:
- `01-prompts/23-sync/`

### Step 2: Author `01-prompts/23-sync/readme.md`
Author the category catalog following the standard `01-prompts/` category format:
- Include version metadata, goal/learn alert blocks.
- Document parameter header (`N=300, A=2, H=2, C=30`).
- Provide catalog table indexing `01-sync-other-codebase.md`.
- Detail the 5 Non-Negotiable Boundaries and dynamic parameter contracts.

### Step 3: Author `01-prompts/23-sync/01-sync-other-codebase.md`
Author the master V6 execution prompt:
- Embed complete V6 parameter header and concurrency formula.
- Document Phase 1A Turn 1 task breakdown and verbatim capture requirements.
- Codify dynamic parameter parsing (`SOURCE_REPO`, `TARGET_REPOS`, `WORKERS`, `DRY_RUN`, `NO_PUSH`).
- Detail the 3-stage subagent dispatch workflow (`A = 2, H = 2`).
- Codify the 5 Non-Negotiable Boundaries with explicit rules and forbidden patterns.
- Detail the 7-step pre-flight pull and backup branch ceremony.
- Enforce rules R1 through R16 cited by ID.
- Define end-of-task consolidation and atomic GitMap commit standards.

### Step 4: Update Global Registries
1. Register `23-sync` in `01-prompts/readme.md` in the category directory table.
2. Register `01-sync-other-codebase.md` in `.ai-memory/prompts.md` in the prompt index table.

---

## 4. Quality Gates & Acceptance Verification

| Gate ID | Verification Check | Expected Outcome |
|:---|:---|:---|
| **GATE-01** | Relative Git Paths Check | 0 instances of absolute filesystem paths or file scheme URIs in authored files. |
| **GATE-02** | Lowercase Filename Check | All created paths strictly lowercase (`01-prompts/23-sync/readme.md`, `01-sync-other-codebase.md`). |
| **GATE-03** | V6 Parameter Header Conformance | Header specifies `N = 300`, `A = 2`, `H = 2`, `C = 30`, `PHASE_1_BUDGET = 150`, `PHASE_2_BUDGET = 150`. |
| **GATE-04** | 5 Non-Negotiable Boundaries Coverage | Spec 21, Bump Scripts, Additive AI Scripts, Memory/Plans, and Zero Secrets explicitly detailed. |
| **GATE-05** | Pre-Flight Pull & Backup Ceremony Coverage | 7-step git ceremony (`pull --no-rebase`, `backup/sync-<timestamp>`, tags, release) fully detailed. |
| **GATE-06** | Zero Hardcoded Paths Verification | Prompt operates entirely on dynamic `SOURCE_REPO` and `TARGET_REPOS` parameters. |
