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
> **User Directive (Verbatim):**  
> *"When you synchronize with other repositories, few things you have to keep in mind. First of all, the spec 21 folder, you should not sync. Okay? If that has updates, you should leave it as it is. AI scripts. If we have the new AI scripts, it should put into there. That's fine. But if the existing script is modified by the repository itself, do not touch this. Okay? So also the bump script, do not modify bump script because the bump script can be updated using its repo's own stuff. So keep this in mind, and based on that, if you think the script needs to be updated, follow that. Follow that, and also update in your memory so that this never happens. Whenever I say sync, you should follow this through."*

---

## 3. The Four Non-Negotiable Sync Principles

### 3.1. Spec 21 Exclusion (`02-spec/21-*` / `21-app`)
- **Rule:** Never synchronize `02-spec/21-*` (`02-spec/21-app`, `02-spec/21-app-issues`, etc.) to target repositories.
- **Rationale:** Application specifications in `02-spec/21-*` are exclusive to specific client repositories and local domain implementations. If target repositories have updates, they must be left completely intact. Upstream must never overwrite or mirror this directory.

### 3.2. Additive-Only AI Scripts (`03-ai-scripts/` and `.agents/scripts/`)
- **Rule:** When syncing AI automation scripts:
  - **New scripts:** If a script exists in the source meta-repo but is missing in the target repo (`not dst_file.exists()`), it may be copied over.
  - **Existing scripts:** If a script already exists in the target repository (`dst_file.exists()`), **DO NOT TOUCH OR OVERWRITE IT**.
  - **No deletions:** Never delete scripts in the target repository that are absent from upstream.
- **Rationale:** Target repositories often adapt or optimize shared AI scripts for their specific technology stack or local environment. Overwriting them clobbers local custom behavior.

### 3.3. Bump Script Protection (`bump*`)
- **Rule:** NEVER modify or overwrite version bump scripts (`bump-version.mjs`, `bump_versions.py`, `37-bump-version.py`, etc.) in target repositories during synchronization.
- **Rationale:** Each repository owns its specific versioning mechanics (e.g. Node vs Python vs Go release flows). Overwriting the bump script breaks repo-level release ceremonies.

### 3.4. Memory & Plans Protection (`.ai-memory/memory/*`, `.ai-memory/plans/*`)
- **Rule:** If target repositories contain operational memory, execution plans, pending tasks, or agent tracking files (`.ai-memory/memory/`, `.ai-memory/plans/`, `.ai-memory/temp-agents/`, `.ai-memory/cicd-issues/`, `.ai-memory/ambiguous-questions/`), they must **NEVER BE MODIFIED, OVERWRITTEN, OR DELETED** during synchronization.
- **Rationale:** Target repositories own their own operational history, execution tasks, active plans, and learned memories. Overwriting or mirroring upstream memory files onto a child repository wipes out that repository's local context and active agent state.

---

## 4. Enforcement in Automation Toolchain

1. `03-ai-scripts/38-sync-prompts-skills-scripts.py`:
   - `SYNC_DIRS` marks script directories with `is_additive = True`.
   - `copy_single_file` and `mirror_directory` enforce `is_spec_21`, `is_bump_script`, `is_additive_only`, and `is_protected_memory_or_plan` guards.
   - `EXCLUDE_NAMES` includes all `21-app*` variants, as well as `plans`, `temp-agents`, `cicd-issues`, and `ambiguous-questions`.
2. `.agents/skills/prompts-and-skills-sync/skill.md`:
   - Codified as top-priority Non-Negotiable Quality Gates.
3. `.ai-memory/strictly-avoid.md`:
   - Registered under **Cross-Repository Synchronization Hard Prohibitions — TOTAL BAN**.
