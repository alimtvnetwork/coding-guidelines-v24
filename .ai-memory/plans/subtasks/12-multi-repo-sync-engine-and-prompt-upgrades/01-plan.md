# Subtask Plan: Multi-Repository Synchronization Engine, V4 Archive & Prompt Upgrades

> **Plan ID:** `.ai-memory/plans/subtasks/12-multi-repo-sync-engine-and-prompt-upgrades/01-plan.md`  
> **Parent Architecture Spec:** `02-spec/21-app/09-multi-repo-sync-engine-and-prompt-upgrades/01-architecture-spec.md`  
> **Status:** IN_PROGRESS  
> **Concurrency Capacity:** A = 2, H = 2

---

## 1. Objectives & Deliverables

1. **Subtask 1: Planning, SQLite Ledger & Memory Update**
   - Initialize SQLite task database: `.ai-memory/temp-agents/14-multi-repo-sync-engine-and-prompt-upgrades/agent-task.db`.
   - Update `.ai-memory/memory/learned/18-cross-repository-sync-rules.md` documenting non-negotiable sync rules: pull base branch, backup branch creation, prompt/skill sync, spec `01-*` to `20-*` sync, spec `21-*` exclusion, additive-only AI scripts with repo modification checks, bump script protection, and memory/plans protection.

2. **Subtask 2: V4 Archive Tier & `01-sync.md` Prompt with Companion Skills**
   - Create `06-archive/v4/readme.md` establishing the V4 prompt archive tier.
   - Author `01-prompts/23-sync/01-sync.md` containing the 43 repository inventory sequence and synchronization ceremony instructions following V6 execution architecture.
   - Update `01-prompts/23-sync/readme.md`.
   - Create `.agents/skills/sync/skill.md` and `.cursor/skills/sync/skill.md`.

3. **Subtask 3: Upgrade Synchronization Engine (`38-sync-prompts-skills-scripts.py`)**
   - Update `03-ai-scripts/38-sync-prompts-skills-scripts.py` and `.agents/scripts/38-sync-prompts-skills-scripts.py`:
     - Dynamically mirror `02-spec/01-*` through `02-spec/20-*` specifications across connected repositories.
     - Protect `02-spec/21-*` through `25-*` from synchronization.
     - Implement repo-modification checks for AI scripts: if a script was modified by the repo (last commit does not contain `sync`), preserve it untouched; if unmodified, prioritize upstream.
     - Validate with `python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run`.

4. **Subtask 4: Multi-Repository Synchronization Execution**
   - Execute `python 03-ai-scripts/38-sync-prompts-skills-scripts.py` across all 43 connected repositories.
   - Verify backup branches and remote pushes.

5. **Subtask 5: Verification & Clean Atomic Commit**
   - Run prompt and relative path linters.
   - Run retrospective AI verification.
   - Stage and commit cleanly via GitMap.
