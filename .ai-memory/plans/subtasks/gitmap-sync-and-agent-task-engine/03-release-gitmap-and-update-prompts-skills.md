# Subtask 03: Release GitMap, Update Prompts/Skills & Fleet Sync

> **Subtask ID:** `Task-03`  
> **Parent Plan:** `.ai-memory/plans/gitmap-sync-and-agent-task-engine.md`  
> **Status:** DONE (100%)  
> **Assigned Agent Role:** `FleetSyncWorker`  
> **Completed At:** 2026-10-08T09:42:00+08:00

---

## Evidence of Completion
- Released GitMap `v6.507.2`, tagged, and pushed to upstream GitHub.
- Updated prompts in `01-prompts/` and `v4/` for `gitmap sync` and `gitmap task`.
- Updated skills in `.agents/skills/` and `.cursor/skills/` (`gitmap`, `sync`, `sync-codebase`, `sync-other-codebase`, `execute-parent-task-with-n-steps-v6`).
- Committed `coding-guidelines` via `gitmap cpf` and pushed to `main`.
- Executed full fleet synchronization across 43 repositories using `gitmap sync --workers 8`.

## 1. Objectives
1. Commit, bump version, tag, and push `gitmap` repository.
2. Update prompts in `01-prompts/` and `v4/` to document `gitmap sync` and `gitmap task`.
3. Update skills in `.agents/skills/` and `.cursor/skills/` (`gitmap`, `execute-parent-task-with-n-steps-v6`, `sync`).
4. Commit and push `coding-guidelines` via `gitmap cpf`.
5. Execute full fleet synchronization across all 43 repositories using the new `gitmap sync` command.

---

## 2. Deliverables
- Released `gitmap` repository with updated tag.
- Updated prompts and skills across `coding-guidelines`.
- Successful fleet synchronization using `gitmap sync`.
