# Cross-Repository Synchronization Rules & Safety Architecture

> **Type:** Institutional Knowledge & Learned Architecture  
> **Status:** Active & Canonical  
> **Date:** 2026-10-02  
> **Target Scripts:** `03-ai-scripts/38-sync-prompts-skills-scripts.py`, `.agents/skills/prompts-and-skills-sync/skill.md`

---

## 1. Executive Overview

This document codifies non-negotiable cross-repository synchronization policies. Whenever prompts, skills, or scripts are synchronized across connected repositories (via `python 03-ai-scripts/38-sync-prompts-skills-scripts.py`, `gitmap`, or any manual sync invocation), four strict boundaries must be preserved to protect repository autonomy and custom implementations.

---

## 2. Mandatory Directives & User Intent

> [!IMPORTANT]
> **User Directives (Verbatim):**  
> *"When you synchronize with other repositories, few things you have to keep in mind. First of all, the spec 21 folder, you should not sync. Okay? If that has updates, you should leave it as it is. AI scripts. If we have the new AI scripts, it should put into there. That's fine. But if the existing script is modified by the repository itself, do not touch this. Okay? So also the bump script, do not modify bump script because the bump script can be updated using its repo's own stuff. So keep this in mind, and based on that, if you think the script needs to be updated, follow that. Follow that, and also update in your memory so that this never happens. Whenever I say sync, you should follow this through."*  
>  
> *"Create a prompt sync folder with the sequence, and its job is to sync the repositories. Only the repository names will be there. It will just say as sync, zero one sync dot MD file, and create a skill for this. For the agent cursor, deploy all these things to all these repositories. First, pull these repositories, ensure you have the latest, and then take a backup of those repositories because the prompts are going to change. Then return to the previous branch, sync the prompts, and ensure not to touch the above script and other changes in the AI script folder for that repository. If there is no touch on that file, then yes, our one will be prioritized and pushed. The same goes for spec. One to 2120 folder, synchronize from the coding guideline. Do not touch the rest of the things. Be very careful when modifying things that are modified. Remember that."*

---

## 3. The Non-Negotiable Sync Principles

### 3.1. Pre-Flight Pull & Safety Backup Branches
- **Rule:** Before applying any changes to target repositories:
  1. Detect the active base branch (`main` / `master`) and pull remote: `git pull origin <base_branch> --no-rebase`.
  2. Create and push a safety backup branch before changes: `git checkout -b backup/sync-<timestamp>` -> `git push origin backup/sync-<timestamp>`.
  3. Return to the base branch: `git checkout <base_branch>`.

### 3.2. Shared Specifications Synchronization (`02-spec/01-*` to `02-spec/20-*`)
- **Rule:** Shared specifications in `02-spec/` matching prefixes `01-*` through `20-*` (such as `01-spec-authoring-guide`, `02-coding-guidelines`, `03-error-manage`, `04-database-conventions`, etc.) synchronize from `coding-guidelines` across target repositories.

### 3.3. Spec 21 & Application Specs Total Exclusion (`02-spec/21-*` to `02-spec/25-*`)
- **Rule:** Never synchronize `02-spec/21-*` (`21-app`), `22-app-issues`, `23-app-db`, `24-app-ui-design-system`, or `25-spec-audits` to target repositories.
- **Rationale:** Application specifications are exclusive to specific client repositories and local domain implementations. If target repositories have updates, they must be left completely intact. Upstream must never overwrite or mirror this directory.

### 3.4. Additive & Unmodified AI Scripts Synchronization (`03-ai-scripts/` and `.agents/scripts/`)
- **Rule:** When syncing AI automation scripts:
  - **New scripts:** If a script exists in the source meta-repo but is missing in the target repo (`not dst_file.exists()`), add it.
  - **Target-Modified scripts:** If a script exists and has been modified by the target repository itself (either uncommitted changes in `git status` or the last commit message does NOT start with/contain `sync`), **DO NOT TOUCH OR OVERWRITE IT**.
  - **Unmodified scripts:** If a script exists and has NOT been modified by the target repository (its last commit is an automated sync commit), upstream is prioritized and updated.
  - **No deletions:** Never delete scripts in the target repository that are absent from upstream.

### 3.5. Bump Script Protection (`bump*`)
- **Rule:** NEVER modify or overwrite version bump scripts (`bump-version.mjs`, `bump_versions.py`, `37-bump-version.py`, etc.) in target repositories during synchronization.
- **Rationale:** Each repository owns its specific versioning mechanics (e.g. Node vs Python vs Go release flows). Overwriting the bump script breaks repo-level release ceremonies.

### 3.6. Memory & Plans Protection (`.ai-memory/memory/*`, `.ai-memory/plans/*`)
- **Rule:** If target repositories contain operational memory, execution plans, pending tasks, or agent tracking files (`.ai-memory/memory/`, `.ai-memory/plans/`, `.ai-memory/temp-agents/`, `.ai-memory/cicd-issues/`, `.ai-memory/ambiguous-questions/`), they must **NEVER BE MODIFIED, OVERWRITTEN, OR DELETED** during synchronization.
- **Rationale:** Target repositories own their own operational history, execution tasks, active plans, and learned memories. Overwriting or mirroring upstream memory files onto a child repository wipes out that repository's local context and active agent state.

---

## 4. Enforcement in Automation Toolchain

1. `03-ai-scripts/38-sync-prompts-skills-scripts.py` and `.agents/scripts/38-sync-prompts-skills-scripts.py`:
   - `SYNC_DIRS` dynamically discovers and syncs `02-spec/01-*` through `02-spec/20-*`.
   - `copy_single_file` and `mirror_directory` enforce `is_spec_21`, `is_bump_script`, target modification checks, and `is_protected_memory_or_plan` guards.
   - `EXCLUDE_NAMES` includes all `06-archive`, `21-app*` variants, as well as `plans`, `temp-agents`, `cicd-issues`, and `ambiguous-questions`.
2. `01-prompts/24-sync/01-sync.md`: Canonical synchronization prompt detailing the 43 repository sequence and boundary rules.
3. `.agents/skills/sync/skill.md` & `.cursor/skills/sync/skill.md`: Autonomous skills for running multi-repo synchronization.
4. `.ai-memory/strictly-avoid.md`: Registered under **Cross-Repository Synchronization Hard Prohibitions — TOTAL BAN**.
