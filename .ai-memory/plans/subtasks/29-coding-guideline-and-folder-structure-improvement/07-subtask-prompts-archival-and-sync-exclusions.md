# Subtask 07: Archive Old Files/Prompts & Sync Boundary Isolation

- **Task Code:** `Task-07`
- **Subtask ID:** `sub-161919-task-07`
- **Assigned Owner:** Worker 03
- **Status:** `PENDING`
- **Owned Files:**
  - `01-prompts/`
  - `06-archive/`
  - `03-ai-scripts/38-sync-prompts-skills-scripts.py`
  - `.agents/scripts/38-sync-prompts-skills-scripts.py`

## Objectives
1. Move deprecated prompt folders (`01-prompts/12-old-plan-prompts/`, `01-prompts/18-insults/`, `01-prompts/19-old-execute-prompts/`) into `06-archive/`.
2. Update `03-ai-scripts/38-sync-prompts-skills-scripts.py` and `.agents/scripts/38-sync-prompts-skills-scripts.py` to add `"05-changes-history"`, `"skills-archived"`, `"12-old-plan-prompts"`, `"18-insults"`, `"19-old-execute-prompts"` to `EXCLUDE_NAMES`.
3. Verify that `05-changes-history/` and `06-archive/` are never synced to other repositories.
