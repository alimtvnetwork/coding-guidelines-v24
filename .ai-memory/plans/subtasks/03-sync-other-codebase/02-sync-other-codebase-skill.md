# Subtask 02: Author Native Antigravity Skill `sync-other-codebase` & Register Prompt `01-sync-other-codebase.md` in Catalogs

> **Parent Spec:** [`02-spec/21-app/03-sync-other-codebase/02-prompt-and-skill-spec.md`](../../../../02-spec/21-app/03-sync-other-codebase/02-prompt-and-skill-spec.md)  
> **Status:** `PENDING`  
> **Traceability ID:** `Task-03-Subtask-02`  
> **Target Files:**
> - `.agents/skills/sync-other-codebase/skill.md` (to create)
> - `01-prompts/readme.md` (to update)
> - `.ai-memory/prompts.md` (to update)

---

## 1. Subtask Objective & Context

This subtask governs authoring the native Antigravity skill `.agents/skills/sync-other-codebase/skill.md` and registering prompt `01-sync-other-codebase.md` across repository-wide prompt catalogs and indices.

In the multi-agent task execution model:
- **Worker 01 (Subtask 01):** Authors the canonical prompt `01-prompts/23-sync/01-sync-other-codebase.md` and category catalog `01-prompts/23-sync/readme.md`.
- **Worker 02 (Subtask 02):** Authors the native IDE/CLI skill `.agents/skills/sync-other-codebase/skill.md` and integrates prompt `01-sync-other-codebase.md` into the master indexes `01-prompts/readme.md` and `.ai-memory/prompts.md`.

This separation of concerns ensures that the skill provides an immediate, user-facing slash command interface (`/sync-other-codebase`) that can dynamically ingest source and target repository paths without hardcoded assumptions, while updating all registry references to point to the new canonical sync prompt.

---

## 2. Deliverable Specifications

### 2.1 Deliverable A: Native Antigravity Skill (`.agents/skills/sync-other-codebase/skill.md`)

- **File Path:** `.agents/skills/sync-other-codebase/skill.md`
- **Naming:** Strictly lowercase, hyphen-separated.
- **YAML Frontmatter:**
  ```yaml
  ---
  name: sync-other-codebase
  description: Autonomously synchronizes canonical prompts, skills, AI scripts, and coding guidelines from a source repository into user-specified target repositories with pre-pull safety, backup branches, and 5 non-negotiable boundary protections.
  ---
  ```

#### Required Skill Sections & Content Architecture

1. **Title & Semicolon Slash Command Headers:**
   - `# Multi-Repository Synchronization & Downstream Codebase Propagation`
   - `> **[/goal](slashCommand;goal)** Propagate canonical prompts, agent skills, shared scripts, and coding guidelines into specified target repositories without clobbering downstream specifications, memory, or custom scripts.`
   - `> **[/learn](slashCommand;learn)** Enforce the 5 Non-Negotiable Boundaries, execute pre-flight pulls and safety backup branch ceremonies, and dynamically ingest repository paths without hardcoding.`

2. **When to Use & Invocation Triggers:**
   - User requests syncing or updating prompts/skills to another repository.
   - User executes `/sync-other-codebase --source <path> --targets <paths>`.
   - Propagating new coding guidelines or Antigravity skills to downstream projects.

3. **Dynamic Parameter Ingestion (Zero Hardcoded Paths):**
   - Must document that target repository paths MUST be supplied by the user or dynamically queried from context.
   - Example parameter block:
     ```text
     SOURCE_REPO  = "."                        # Source repository root (default: current workspace)
     TARGET_REPOS = ["../repo-a", "../repo-b"] # Target repositories to synchronize
     WORKERS      = 4                          # Concurrency threads
     DRY_RUN      = false                      # Simulate file operations without modifying disk
     NO_PUSH      = false                      # Commit locally without pushing remote branches
     ```

4. **Step-by-Step Shell & Git Protocol:**
   - **Step 1: Pre-Flight Cleanliness & Fast-Forward Pull:**
     ```bash
     cd <target_repo_path>
     git status --porcelain
     git checkout main
     git pull origin main --no-rebase
     ```
   - **Step 2: Safety Backup Branch Creation & Push:**
     ```bash
     git branch backup/sync-$(date +%Y%m%d-%H%M%S)
     git push origin backup/sync-$(date +%Y%m%d-%H%M%S)
     git checkout -b feat/sync-guidelines-$(date +%Y%m%d-%H%M%S)
     ```
   - **Step 3: Bounded Asset Synchronization:**
     - Detail commands for mirroring `01-prompts/`, `.agents/skills/`, `02-spec/02-coding-guidelines/`, `.cursor/skills/`.
     - Detail additive-only copying for `03-ai-scripts/` and `.agents/scripts/`.
   - **Step 4: AI Script Diff Inspection & Boundary Verification:**
     ```bash
     git diff 03-ai-scripts/
     git diff .agents/scripts/
     ```
     - Verification commands ensuring `02-spec/21-*`, `*bump*`, and `.ai-memory/plans/` remain untouched.
   - **Step 5: Atomic Commit & Remote Push:**
     ```bash
     git add 01-prompts/ .agents/skills/ 03-ai-scripts/ .agents/scripts/ 02-spec/02-coding-guidelines/
     git commit -m "chore(sync): synchronize canonical prompts, skills, and coding guidelines"
     git push origin feat/sync-guidelines-<timestamp>
     ```

5. **The 5 Non-Negotiable Boundaries Table:**
   - Reference table summarizing Boundary 1 (Spec 21 Exclusion), Boundary 2 (Additive-Only AI Scripts), Boundary 3 (Bump Script Protection), Boundary 4 (Memory & Plans Protection), and Boundary 5 (Zero Secret Keys Leakage).

6. **Actionable Verification Checklist:**
   - Comprehensive checklist covering pre-flight pull, backup branch, boundary enforcement, diff inspection, relative path compliance, and clean git status.

---

### 2.2 Deliverable B: Master Prompt Catalog Registration (`01-prompts/readme.md`)

- **File Path:** `01-prompts/readme.md`
- **Changes Required:**
  1. **Directory Index Update:**
     - Add `23-sync/` to the directory tree list under `01-prompts/`:
       ```text
       ├── 21-temp-end-to-end-tests/
       ├── 22-letterly/
       └── 23-sync/
       ```
  2. **Core Architecture Section Update:**
     - Add section item summarizing the **Parameter-Driven Multi-Repository Synchronization Engine (`01-prompts/23-sync/01-sync-other-codebase.md`)**:
       - Dynamic parameter ingestion (`SOURCE_REPO`, `TARGET_REPOS`).
       - Zero hardcoded repository paths.
       - Mandatory pre-flight pull and backup branch ceremony.
       - 5 Non-Negotiable Boundary Protections.
       - Diff inspection for AI scripts.

---

### 2.3 Deliverable C: Memory Prompts Matrix Registration (`.ai-memory/prompts.md`)

- **File Path:** `.ai-memory/prompts.md`
- **Changes Required:**
  - In the `## Prompts Matrix` table, add the entry for `23-sync`:
    ```markdown
    | `23-sync` | [`23-sync/01-sync-other-codebase.md`](../01-prompts/23-sync/01-sync-other-codebase.md) | Multi-Repository Synchronization Engine — Parameter-Driven Workflow (must follow) |
    ```
  - Place row immediately following `22-letterly` entries to preserve alphabetical and sequence order.

---

## 3. Step-by-Step Implementation Instructions

Follow these exact steps to complete this subtask:

### Step 1: Create `.agents/skills/sync-other-codebase/skill.md`
1. Create directory `.agents/skills/sync-other-codebase/` if it does not already exist.
2. Author `skill.md` with:
   - Valid YAML frontmatter (`name: sync-other-codebase`, `description: ...`).
   - Semicolon slash commands (`[/goal]`, `[/learn]`).
   - Dynamic parameter ingestion documentation (explaining how users provide target paths).
   - Complete 5-step shell command sequence (pre-flight pull, backup branch, bounded sync, diff inspection, atomic commit).
   - The 5 Non-Negotiable Boundaries table.
   - Verification checklist.
3. Ensure strictly relative git paths and strictly lowercase file naming.

### Step 2: Update `01-prompts/readme.md`
1. Locate the `## Directory Index` block at the bottom of `01-prompts/readme.md`.
2. Update the ASCII directory tree to add `└── 23-sync/` after `22-letterly/`.
3. Add a numbered entry in `## Core Architecture & Capabilities` summarizing the Parameter-Driven Multi-Repository Synchronization Engine.
4. Verify all relative links remain intact.

### Step 3: Update `.ai-memory/prompts.md`
1. Locate the `## Prompts Matrix` table in `.ai-memory/prompts.md`.
2. Find the row for `22-letterly/03-execute-n-steps.md`.
3. Insert the new row for `23-sync/01-sync-other-codebase.md`:
   ```markdown
   | `23-sync` | [`23-sync/01-sync-other-codebase.md`](../01-prompts/23-sync/01-sync-other-codebase.md) | Multi-Repository Synchronization Engine — Parameter-Driven Workflow (must follow) |
   ```
4. Save file and ensure no broken formatting.

### Step 4: Execute Quality Linters & Verification
1. Run relative path validation:
   ```bash
   python linter-scripts/check-relative-paths.py
   ```
2. Run prompt loading linter:
   ```bash
   python linter-scripts/check-prompts-loaded.py
   ```
3. Run git status check:
   ```bash
   git status --porcelain
   ```
4. Confirm exit code 0 across all verification linters.

---

## 4. Verification & Acceptance Criteria

- **AC-01:** File `.agents/skills/sync-other-codebase/skill.md` exists and contains valid YAML frontmatter with `name: sync-other-codebase`.
- **AC-02:** Skill `.agents/skills/sync-other-codebase/skill.md` defines zero hardcoded paths, requiring dynamic user parameters for `SOURCE_REPO` and `TARGET_REPOS`.
- **AC-03:** Skill `.agents/skills/sync-other-codebase/skill.md` fully documents the pre-flight pull, backup branch creation, 5 boundary protections, AI script diff inspection, and atomic commit commands.
- **AC-04:** Master prompt index `01-prompts/readme.md` includes category `23-sync/` in the directory index and capability matrix.
- **AC-05:** Registry `.ai-memory/prompts.md` contains the matrix row for `23-sync/01-sync-other-codebase.md`.
- **AC-06:** Strict relative git paths mandate is 100% satisfied; zero absolute filesystem paths or `file:///` URIs exist in authored or edited files.
- **AC-07:** Strict lowercase file naming is enforced; zero uppercase characters in file paths.

---

## 5. Traceability & Dependency Matrix

- **Parent Spec:** `02-spec/21-app/03-sync-other-codebase/02-prompt-and-skill-spec.md`
- **Architecture Spec:** `02-spec/21-app/03-sync-other-codebase/01-architecture-spec.md`
- **Sibling Subtask Plan:** `.ai-memory/plans/subtasks/03-sync-other-codebase/01-sync-execute-prompt.md`
- **Master Plan Index:** `.ai-memory/plans/readme.md`
- **Prompts Library Index:** `01-prompts/readme.md`
- **Memory Prompts Registry:** `.ai-memory/prompts.md`
