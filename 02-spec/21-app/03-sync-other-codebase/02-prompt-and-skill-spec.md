# Specification: Multi-Repository Synchronization Prompt & Native Antigravity Skill (V6 Architecture)

> **/goal** Establish the formal architectural design, execution contracts, and verification protocols for Prompt `01-prompts/23-sync/01-sync-other-codebase.md` and native Antigravity skill `.agents/skills/sync-other-codebase/skill.md`, enforcing parameter-driven, zero-hardcoded-path multi-repository synchronization across downstream codebases while safeguarding the 5 Non-Negotiable Boundaries.  
> **/learn** Master the V6 execution parameters (`N = 300, A = 2, H = 2, C = 30`), dynamic parameter ingestion (`SOURCE_REPO`, `TARGET_REPOS`), semicolon slash commands (`[/goal]`, `[/learn]`, `[/plan]`), mandatory subagent spawning gate (`invoke_subagent`), the 5 Non-Negotiable Boundaries (Spec 21 Exclusion, Bump Script Protection, Additive-Only AI Scripts, Memory & Plans Protection, Zero Secret Keys Leakage), AI script diff inspection, pre-flight pull and backup branch ceremonies, and strict R1–R16 operational rule citations.

**Version:** 1.0.0  
**Updated:** 2026-10-02  
**Status:** Active  
**AI Confidence:** Production-Ready  
**Ambiguity:** None  
**Target Scope:** Prompt Category `01-prompts/23-sync/` & Native Skill `.agents/skills/sync-other-codebase/`  
**Reference Parent Spec:** `02-spec/21-app/03-sync-other-codebase/01-architecture-spec.md`  
**Subtask Plan:** `.ai-memory/plans/subtasks/03-sync-other-codebase/02-sync-other-codebase-skill.md`  

---

## 1. Executive Summary & Architectural Motivation

The Prompt Architect meta-repository (`coding-guidelines`) serves as the central authority for AI prompt infrastructure (`01-prompts/`), agent skills (`.agents/skills/`), automation scripts (`03-ai-scripts/`, `.agents/scripts/`), and cross-language coding guidelines (`02-spec/02-coding-guidelines/`).

In real-world multi-repository workflows, organizations synchronize governance assets from this central meta-repository to numerous downstream codebases. However, historical synchronization suffered from critical pain points:

1. **Hardcoded Repository Paths:** Legacy prompts and scripts hardcoded static lists of connected repositories (such as 42 fixed repository paths in parent directories). This prevented engineers from running ad-hoc synchronizations against a single repository, synchronizing new projects dynamically, or utilizing custom workspace locations.
2. **Lack of User-Provided Parameter Blocks:** Prompts lacked dynamic parameter interfaces where users or caller agents could supply `SOURCE_REPO` and `TARGET_REPOS` directly in the prompt header or preamble.
3. **Boundary Violations & Overwrite Risks:** Blind directory copying risked overwriting downstream-specific application specifications (`02-spec/21-*`), destroying custom package version bump scripts, wiping out target repository agent execution memory (`.ai-memory/plans/` and `.ai-memory/memory/`), or clobbering modified automation scripts without inspecting diffs.
4. **Discipline Gaps:** Previous synchronization prompts did not strictly follow the **V6 Autonomous Execution Engine** (`N = 300, A = 2, H = 2, C = 30`), semicolon slash commands, mandatory `invoke_subagent` spawning gates, pre-flight pull and backup branch ceremonies, and explicit R1–R16 rule citations.

### The Solution: Prompt `01-sync-other-codebase.md` & Native Skill `sync-other-codebase`

To solve these challenges with production rigor, this specification defines:
- **`01-prompts/23-sync/01-sync-other-codebase.md`:** A canonical V6 autonomous execution prompt that consumes source and target repository paths dynamically, enforces pre-flight pull and safety backup branches, allocates disjoint target repositories to spawned subagents (`A = 2, H = 2`), rigorously applies the 5 Non-Negotiable Boundaries, requires diff inspection for AI scripts, and commits atomically per repository.
- **`.agents/skills/sync-other-codebase/skill.md`:** A native Antigravity skill enabling users and agents to invoke cross-repository synchronization via slash commands (`/sync-other-codebase`) with runtime arguments and clear step-by-step verification commands.

---

## 2. Specification for Prompt `01-sync-other-codebase.md`

### 2.1 File Placement & Naming Architecture

- **Path:** `01-prompts/23-sync/01-sync-other-codebase.md`
- **Category:** `23-sync` (new canonical prompt category dedicated to cross-repository synchronization)
- **Naming Convention:** Strictly lowercase, hyphen-separated. Zero uppercase characters.

### 2.2 Top Header & Dynamic Parameter Architecture (Strict V6 Format)

The prompt must open with the standardized V6 parameter header augmented with a dynamic, user-provided parameter block:

```text
N = 300 (Total self-loop steps budget — editable top-header parameter, default: 300)
A = 2   (MANDATORY number of spawned autonomous subagents running concurrently via invoke_subagent, default: 2)
H = 2   (Operational hands per agent: dual-task batch capacity & parallel tool dispatch, default: 2)
C = 30  (Tool calls per worker before it must report, default: 30)

System Concurrency Capacity = A × H = 2 agents × 2 hands = 4 concurrent subtask operations
PHASE_1_BUDGET = N / 2   (Steps 1 .. 150: Dynamic Target Ingestion, Pre-Flight Git Checks & Safety Backup Branches)
PHASE_2_BUDGET = N / 2   (Steps 151 .. 300: Guarded Multi-Agent Sync Execution, Diff Inspection, Atomic Commits)
WAVES = ceil(target_repos / (A x H))
```

#### Dynamic Parameter Block (Zero Hardcoded Paths Mandate)

Directly following the V6 header parameters, the prompt MUST define the dynamic parameter block. No static filesystem paths or hardcoded repository lists are permitted:

```text
SOURCE_REPO   = "<user_provided_source_path>"   # Source repository root (default: "." / current repository root)
TARGET_REPOS  = [                               # User-provided target repository path(s) or relative paths
    "<user_provided_target_repo_path_1>",
    "<user_provided_target_repo_path_2>",
]
WORKERS       = 4                               # Parallel worker subagents / threads (default: 4)
DRY_RUN       = false                           # Preview file copy/delete actions without disk mutation (default: false)
NO_PUSH       = false                           # Commit and tag locally without pushing to git remote (default: false)
```

### 2.3 Top-Instruction Priority Mandate (Above Precedence)

The prompt MUST prominently display the Top-Instruction Priority Mandate:

> [!IMPORTANT]
> **Top-Instruction Priority Mandate (Above Precedence / Preamble Precedence):**  
> Whatever directives, constraints, custom repository paths (`SOURCE_REPO`, `TARGET_REPOS`), checklists, or user instructions are provided ABOVE this prompt (including user preamble, header blocks, or incoming message arguments) are HIGHEST PRIORITY and MUST BE FOLLOWED as strictly NON-NEGOTIABLE. They supersede and strictly override any conflicting general advice, default conventions, or lower-level guidelines below. The agent MUST inspect and follow the instructions above with absolute precedence.

### 2.4 Semicolon Slash Commands

Prompt `01-sync-other-codebase.md` must implement three standardized semicolon slash commands:

1. `[/goal](slashCommand;goal)`:
   > Autonomously orchestrate and execute multi-repository synchronization from `SOURCE_REPO` into all specified `TARGET_REPOS`: FIRST showcase and list out the given target repositories and synchronization parameters in visible chat during Turn 1, capture the user request verbatim, plan the synchronization wave in the repository, spawn autonomous subagents via `invoke_subagent` (A = 2, H = 2; solo execution without calling `invoke_subagent` is an auto-reject failure) allocating disjoint target repository sets, enforce the 5 Non-Negotiable Boundaries (Spec 21 Exclusion, Bump Script Protection, Additive-Only AI Scripts with Diff Inspection, Memory & Plans Protection, Zero Secret Keys Leakage), execute mandatory pre-flight pull and safety backup branch ceremonies per target repo, prove every synchronization step with concrete diff evidence, and finish with atomic GitMap/git commits per target repository holding strictly synced assets without modifying protected files.

2. `[/learn](slashCommand;learn)`:
   > Enforce the Top-Instruction Priority Mandate: whatever directives, custom repository paths, or target lists are provided ABOVE this prompt outrank everything below. Turn 1 MUST showcase the target repository matrix in visible chat before any background execution. Master the 5 Non-Negotiable Boundaries: `02-spec/21-app/03-sync-other-codebase/01-architecture-spec.md` and `02-spec/21-app/03-sync-other-codebase/02-prompt-and-skill-spec.md`. Each rule is stated once (R1 to R16) and cited by ID. Progress lives in the ledger and in `.ai-memory/plans/`, never only in chat.

3. `[/plan](slashCommand;plan)`:
   > Execute thorough step-by-step target repository validation before any file copy operations. Verify clean git working trees, resolve base branches, execute fast-forward pulls, create and push safety backup branches, inventory existing bump scripts and local scripts, and confirm boundaries before dispatching worker waves.

### 2.5 Mandatory Subagent Spawning Gate (`invoke_subagent`, `A = 2, H = 2`)

- **Actual Tool Call Enforcement:** The lead orchestrator MUST invoke the actual `invoke_subagent` tool via its JSON tool-calling API. Outputting textual claims of dispatch without calling `invoke_subagent` is an auto-reject failure.
- **3-Stage Dispatch Pattern:**
  1. *Discovery & Pre-Flight Stage (`TypeName: "research"`, A = 2):* Subagents inspect target repositories, verify clean git status, detect base branches (`main` or `master`), scan for existing bump scripts and local scripts, and report findings to the lead.
  2. *Spec & Allocation Stage (`TypeName: "self"` or Lead):* Subagents partition target repositories into disjoint batches (e.g. Worker 1 handles targets 1..K, Worker 2 handles targets K+1..M).
  3. *Execution Stage (`TypeName: "self"`, A = 2 worker subagents):* Spawn 2 concurrent worker subagents, each processing up to `H = 2` target repositories per wave in strict isolation.
- **Solo Execution Ban:** The lead orchestrator is strictly prohibited from executing target repository file writes solo.

### 2.6 The 3-Phase Multi-Repository Synchronization Pipeline

Prompt `01-sync-other-codebase.md` must execute synchronization via a rigorous 3-Phase pipeline:

#### Phase 1: Pre-Flight Pull, Backup Branch & Target Validation (Steps 1 .. 150)

For every target repository specified in `TARGET_REPOS`:

1. **Working Tree Cleanliness Verification:**
   - Execute `git status --porcelain` in the target repository directory.
   - If uncommitted changes exist, halt synchronization for that repository, log failure in the ledger, and notify the user. Never overwrite a dirty working tree.
2. **Base Branch Detection:**
   - Query current branch and verify whether `main` or `master` is the default upstream branch.
3. **Pre-Flight Fast-Forward Pull:**
   - Switch to base branch: `git checkout <base_branch>`.
   - Pull latest remote changes: `git pull origin <base_branch> --no-rebase`.
   - Prevent non-fast-forward push rejections and ensure the local repository matches remote state before branching.
4. **Safety Backup Branch Creation & Push:**
   - Generate timestamped backup branch name: `backup/sync-<YYYYMMDD-HHMMSS>`.
   - Create and immediately push backup branch:
     ```bash
     git branch backup/sync-<timestamp>
     git push origin backup/sync-<timestamp>
     ```
   - This provides an immutable restore point in case of any synchronization anomalies.
5. **Dedicated Work Branch Checkout:**
   - Create and switch to an isolated feature branch for synchronization:
     ```bash
     git checkout -b feat/sync-guidelines-<timestamp>
     ```

#### Phase 2: Bounded Mirroring with 5 Boundary Protections & Diff Inspection (Steps 151 .. 250)

Synchronize canonical assets from `SOURCE_REPO` to each target repository while enforcing the **5 Non-Negotiable Boundaries**:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                       SOURCE REPOSITORY (SOURCE_REPO)                       │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
                      BOUNDED MIRRORING PIPELINE (V6 ENGINE)
                                       │
          ┌────────────────────────────┼────────────────────────────┐
          ▼                            ▼                            ▼
┌──────────────────┐         ┌──────────────────┐         ┌──────────────────┐
│   01-prompts/    │         │  .agents/skills/ │         │ 02-coding-guide/ │
│  (Clean Mirror)  │         │  (Clean Mirror)  │         │  (Clean Mirror)  │
└──────────────────┘         └──────────────────┘         └──────────────────┘
                                       │
                        STRICT BOUNDARY GUARDS (TOTAL BANS)
                                       │
   ┌───────────────────────┬───────────┴───────────┬───────────────────────┐
   ▼                       ▼                       ▼                       ▼
BOUNDARY 1              BOUNDARY 2              BOUNDARY 3              BOUNDARY 4 & 5
Spec 21 Exclusion       Additive AI Scripts     Bump Script Guard       Memory, Plans & Secrets
`02-spec/21-*`          `03-ai-scripts/`        `*bump*`                `.ai-memory/plans/`
TOTAL BAN on sync.      Diff & understand.      NEVER overwrite         `.ai-memory/memory/`
Target remains intact.  Never delete scripts.   target bump script.     NEVER touch or leak.
```

1. **Boundary 1: Spec 21 Exclusion (`02-spec/21-*`):**
   - NEVER copy, mirror, touch, or overwrite `02-spec/21-*` (`02-spec/21-app`, `02-spec/21-app-issues`, `02-spec/21-app-db`, `02-spec/21-app-ui-design-system`). Target repositories own their unique application specifications.
2. **Boundary 2: Additive-Only AI Scripts & Diff Inspection (`03-ai-scripts/`, `.agents/scripts/`):**
   - New scripts in `SOURCE_REPO` missing from the target repository are copied cleanly.
   - Existing scripts in the target repository must NOT be overwritten blindly. Subagents must run `git diff <script>` to understand modifications. If the target repository has specialized logic (e.g. local container paths or environment flags), preserve those customizations.
   - Stale scripts existing in the target repository are NEVER deleted.
3. **Boundary 3: Version Bump Script Protection (`*bump*`):**
   - NEVER overwrite or modify any version bump script in target repositories (`bump-version.mjs`, `bump_versions.py`, `37-bump-version.py`, `bump.js`, etc.). Target repositories own their ecosystem-specific versioning mechanisms.
4. **Boundary 4: Memory & Plans Protection (`.ai-memory/memory/`, `.ai-memory/plans/`):**
   - NEVER overwrite, delete, or mirror files into `.ai-memory/memory/`, `.ai-memory/plans/`, `.ai-memory/temp-agents/`, `.ai-memory/cicd-issues/`, or `.ai-memory/ambiguous-questions/`. Target repositories own their operational history and active task tracking.
   - Top-level index references (`.ai-memory/coding-guidelines.md` and `.ai-memory/prompts.md`) may be updated cleanly to reflect canonical standards.
5. **Boundary 5: Zero Secret Keys Leakage & Archival/Temp Exclusions:**
   - NEVER synchronize `.env` files, `.env.local`, API keys, private tokens, or credentials.
   - Strictly exclude `06-old-prompts/`, `.git/`, `__pycache__/`, `.pytest_cache/`, `.mypy_cache/`, `.DS_Store`, and temporary test fixtures.

#### Phase 3: Verification, Atomic Commits & Multi-Repo Summary (Steps 251 .. 300)

1. **Post-Sync Verification & Boundary Audit:**
   - Inspect `git status --porcelain` in each target repository.
   - Verify that zero files under `02-spec/21-*`, `*bump*`, or `.ai-memory/plans/` appear as added, modified, or deleted.
   - Run targeted relative path linters (`python linter-scripts/check-relative-paths.py` if present).
2. **Atomic Commit per Target Repository:**
   - Stage strictly synchronized directories:
     ```bash
     git add 01-prompts/ .agents/skills/ 03-ai-scripts/ .agents/scripts/ 02-spec/02-coding-guidelines/ .cursor/skills/
     git commit -m "chore(sync): synchronize canonical prompts, skills, and coding guidelines"
     ```
   - If `NO_PUSH` is false, push the feature branch:
     ```bash
     git push origin feat/sync-guidelines-<timestamp>
     ```
3. **Multi-Repo Execution Summary Table:**
   - Emit a structured summary table in chat and persist to the ledger:
     | Target Repository | Branch | Backup Branch | Synced Folders | Boundary Violations | Commit SHA | Status |
     | :--- | :--- | :--- | :--- | :--- | :--- | :--- |

---

### 2.7 Core Operational Rules (Cited R1 to R16 by ID)

Prompt `01-sync-other-codebase.md` must explicitly cite and enforce all 16 core rules:

| Rule ID | Rule Name | Operational Enforcement in Multi-Repository Synchronization |
| :--- | :--- | :--- |
| **R1** | Zero Builds or Test Suites | TOTAL BAN on executing `go test`, `npm test`, `pytest`, `cargo build`, or repository build suites during synchronization. |
| **R2** | Targeted Checks Only | Execute only fast, targeted relative path and prompt linters. Scanning 0 files is an instant FAIL. |
| **R3** | Evidence or It Did Not Happen | Every synchronization claim must cite concrete file diffs, git commit hashes, or clean status outputs. |
| **R4** | Never Invent Commands, Flags, or Paths | Use only documented git, GitMap, and Python CLI flags. Validate path existence before file operations. |
| **R5** | Mandatory Subagents (`A = 2, H = 2`) | Spawning subagents via `invoke_subagent` is mandatory across all phases. Solo execution is an auto-reject failure. |
| **R6** | One Owner Per File / Repo | Partition target repositories into disjoint sets. No two worker subagents may operate on the same repository concurrently. |
| **R7** | Git Safety & Isolation | Pre-flight pull and backup branches mandatory before modifying target repositories. Workers operate in assigned directories. |
| **R8** | Explicit-Path Staging & Clean Gitignore | Stage strictly synchronized paths (`git add <dirs>`). Never use `git add .` or `git add -A`. Verify `.gitignore` hygiene. |
| **R9** | Atomic Commit per Repository | Generate one focused, atomic commit per target repository capturing only synchronized governance assets. |
| **R10** | Zero Unauthorized Releases | Never bump versions, update changelogs, or cut release tags in target repositories unless explicitly commanded. |
| **R11** | Strict Relative Git Paths & Lowercase Hygiene | TOTAL BAN on absolute filesystem paths or `file:///` URIs in any synced file. All filenames must be strictly lowercase. |
| **R12** | No Polling / Immediate Turn Yielding | Yield turns immediately after dispatching worker subagents via `invoke_subagent`. Never busy-wait or poll in loops. |
| **R13** | Two-Strike Retry Cap & Anti-Looping | If a repository synchronization encounters repeated errors (e.g. merge conflicts), halt that repo after 2 attempts and log. |
| **R14** | 100% Ambiguity & Decision Boundaries | If target repository paths are ambiguous or unprovided, ask the user or fail safe rather than guessing random directories. |
| **R15** | Zero Generated Artifacts Committed | Temporary diff manifests, sync logs, and scratch files must reside in gitignored temp paths or `.ai-memory/temp-agents/`. |
| **R16** | Zero Secrets in Standard Repos | Never copy `.env` files, API keys, or private tokens. Offload private configuration to `repo-secrets` via `gitmap rs`. |

---

## 3. Specification for Native Antigravity Skill (`sync-other-codebase`)

### 3.1 Skill Placement & Frontmatter Architecture

- **Path:** `.agents/skills/sync-other-codebase/skill.md`
- **YAML Frontmatter:**

```yaml
---
name: sync-other-codebase
description: Autonomously synchronizes canonical prompts, skills, AI scripts, and coding guidelines from a source repository into user-specified target repositories with pre-pull safety, backup branches, and 5 non-negotiable boundary protections.
---
```

### 3.2 Dynamic Invocation Patterns

The skill must be usable both via slash commands and natural language prompts without hardcoding paths:

```bash
# Pattern 1: Slash Command with Target Flags
/sync-other-codebase --source . --targets "../target-repo-1, ../target-repo-2"

# Pattern 2: Single Target Synchronize
/sync-other-codebase --source . --targets "../movie-cli"

# Pattern 3: Dry-Run Mode
/sync-other-codebase --source . --targets "../target-repo" --dry-run

# Pattern 4: Natural Language Chat Invocation
"Please sync prompts, skills, and coding guidelines from this repository into ../lara-licensing and ../wp-exam"
```

#### Zero-Hardcoded Fallback Protocol
If the user invokes `/sync-other-codebase` without specifying target paths:
1. The skill must inspect the conversation context and prompt the user:
   *"Please specify the target repository path(s) to synchronize (e.g. `../target-repo` or comma-separated paths)."*
2. Under NO circumstance should the skill assume or hardcode a static array of repository paths.

### 3.3 Step-by-Step Shell & Git Commands

The skill documentation must detail the exact shell commands executed during each phase:

#### Step 1: Pre-Flight Pull & Safety Backup Branch
```bash
# Navigate to target repository
cd <target_repo_path>

# Verify clean status
git status --porcelain

# Checkout base branch and pull latest remote commits
git checkout main
git pull origin main --no-rebase

# Create and push safety backup branch
git branch backup/sync-$(date +%Y%m%d-%H%M%S)
git push origin backup/sync-$(date +%Y%m%d-%H%M%S)

# Switch to dedicated work branch
git checkout -b feat/sync-guidelines-$(date +%Y%m%d-%H%M%S)
```

#### Step 2: Bounded Asset Synchronization
Execute via the parameter-driven Python synchronizer or manual bounded copy:

```bash
# Option A: Parameter-driven Python script from source repository
python 03-ai-scripts/38-sync-prompts-skills-scripts.py --source <source_repo> --target <target_repo>

# Option B: Manual bounded rsync / copy adhering to 5 boundaries
# Mirror prompts
rsync -av --delete --exclude '06-old-prompts' <source>/01-prompts/ <target>/01-prompts/

# Mirror skills
rsync -av --delete <source>/.agents/skills/ <target>/.agents/skills/

# Mirror guidelines
rsync -av --delete <source>/02-spec/02-coding-guidelines/ <target>/02-spec/02-coding-guidelines/

# Additive copy of new AI scripts only (never overwrite existing)
rsync -av --ignore-existing <source>/03-ai-scripts/ <target>/03-ai-scripts/
rsync -av --ignore-existing <source>/.agents/scripts/ <target>/.agents/scripts/
```

#### Step 3: Script Diff Inspection & Boundary Verification
```bash
# Inspect any modified scripts to verify no specialized logic was overwritten
git diff 03-ai-scripts/
git diff .agents/scripts/

# Verify 5 boundaries are strictly respected
# Ensure NO changes in 02-spec/21-*
git status --porcelain | grep "02-spec/21-" && echo "FAIL: Spec 21 violated" || echo "OK"

# Ensure NO changes in bump scripts
git status --porcelain | grep -i "bump" && echo "FAIL: Bump script modified" || echo "OK"

# Ensure NO changes in child memory/plans
git status --porcelain | grep ".ai-memory/plans" && echo "FAIL: Memory/plans touched" || echo "OK"
```

#### Step 4: Atomic Commit & Push
```bash
# Stage strictly synced directories
git add 01-prompts/ .agents/skills/ 03-ai-scripts/ .agents/scripts/ 02-spec/02-coding-guidelines/

# Commit atomically
git commit -m "chore(sync): synchronize canonical prompts, skills, and coding guidelines"

# Push feature branch
git push origin feat/sync-guidelines-<timestamp>
```

### 3.4 Verification Checklist

The skill file must provide an actionable verification checklist for agents:
- [ ] Working tree in target repository confirmed clean before synchronization.
- [ ] Remote `origin/<base_branch>` pulled cleanly without fast-forward conflicts.
- [ ] Safety backup branch `backup/sync-<timestamp>` pushed to remote.
- [ ] `02-spec/21-*` completely untouched in target repository (0 files modified/added).
- [ ] Target version bump scripts (`*bump*`) completely untouched (0 files modified).
- [ ] Existing modified AI scripts inspected via `git diff`; local configs preserved.
- [ ] `.ai-memory/memory/` and `.ai-memory/plans/` in target repository untouched.
- [ ] Zero secret keys, `.env` files, or API credentials present in staged diff.
- [ ] Strictly relative git paths verified; zero absolute paths or `file:///` URIs.
- [ ] Atomic commit created containing strictly synced assets.

---

## 4. Registries & Catalog Integrations

### 4.1 Master Prompts Library (`01-prompts/readme.md`)
- Update the Directory Index to include `23-sync/`:
  ```text
  ├── 22-letterly/
  └── 23-sync/
  ```
- Add Section 10 to Core Architecture describing the Parameter-Driven Multi-Repository Synchronization Engine.

### 4.2 Category Readme (`01-prompts/23-sync/readme.md`)
- Create dedicated category readme indexing:
  - Sequence: `01`
  - Filename: `01-sync-other-codebase.md`
  - Purpose: `Parameter-Driven Multi-Repository Synchronization & Boundary Protection`

### 4.3 Prompts Matrix Registry (`.ai-memory/prompts.md`)
- Register prompt `01-sync-other-codebase.md` under category `23-sync` in the master prompts matrix table:
  ```markdown
  | `23-sync` | [`23-sync/01-sync-other-codebase.md`](../01-prompts/23-sync/01-sync-other-codebase.md) | Multi-Repository Synchronization Engine — Parameter-Driven Workflow (must follow) |
  ```

---

## 5. Traceability & Dependency Matrix

- **Parent Spec:** `02-spec/21-app/03-sync-other-codebase/01-architecture-spec.md`
- **Subtask Plan 01 (Prompt Authoring):** `.ai-memory/plans/subtasks/03-sync-other-codebase/01-sync-execute-prompt.md`
- **Subtask Plan 02 (Skill & Catalogs):** `.ai-memory/plans/subtasks/03-sync-other-codebase/02-sync-other-codebase-skill.md`
- **Master Plan Register:** `.ai-memory/plans/readme.md`
- **Master Prompt Catalog:** `01-prompts/readme.md`
- **Memory Prompts Registry:** `.ai-memory/prompts.md`
- **Reference Script:** `03-ai-scripts/38-sync-prompts-skills-scripts.py`
