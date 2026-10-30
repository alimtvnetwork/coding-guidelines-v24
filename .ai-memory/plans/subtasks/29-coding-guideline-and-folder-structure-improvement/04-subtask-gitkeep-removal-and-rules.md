# Subtask 04: Remove .gitkeep Files from Populated Folders & Codify Rule

- **Task Code:** `Task-04`
- **Subtask ID:** `sub-161914-task-04`
- **Assigned Owner:** Worker 03
- **Status:** `DONE`
- **Owned Files:**
  - Redundant `.gitkeep` files in populated directories
  - `.ai-memory/strictly-avoid.md`
  - `AGENTS.md`

## Objectives
1. Remove all 23 redundant `.gitkeep` files identified across `.agents/rules/`, `.ai-memory/`, `01-prompts/`, `02-spec/`, `06-archive/`, and `assets/`.
2. Strictly retain the 5 legitimate `.gitkeep` files in truly empty placeholder directories.
3. Formally codify the non-negotiable `.gitkeep` rule in `.ai-memory/strictly-avoid.md`: "Whenever a folder contains any tracked file, any .gitkeep file inside that folder MUST be removed immediately."
