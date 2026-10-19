# Plan: Letterly Prompt Cleanup and Fleet Sync

- **Task Slug:** `letterly-prompt-cleanup-and-fleet-sync`
- **Spec:** `02-spec/21-app/letterly-prompt-cleanup-and-fleet-sync/01-architecture-spec.md`
- **Goal:** Decouple step budgets and Muse mentions from Letterly prompts, compact instructions by removing filler words and reducing line count, mirror into Cursor prompts, verify linters, release v6.83.0, and sync fleet across 43 repositories.
- **Status:** IN PROGRESS (Subtask 01 active)

## Execution Subtasks

1. [x] **Subtask 01:** Architecture spec and plan creation in `02-spec/21-app/letterly-prompt-cleanup-and-fleet-sync/` and `.ai-memory/plans/`.
2. [ ] **Subtask 02:** Letterly prompts cleanup in `01-prompts/22-letterly/` (specifically `03-execute-n-steps-letterly.md` and related formatters), removing step budgets, removing Muse references, and compacting text.
3. [ ] **Subtask 03:** Cursor letterly prompts cleanup in `01-prompts/23-cursor-prompts/` (specifically `03-execute-n-steps-letterly-cursor.md` and related formatters), ensuring mirror parity with `.cursor/skills` links.
4. [ ] **Subtask 04:** Repository verification, internal sync (`npm run sync`, `npm run sync:check`), package installers, minor bump release ceremony (`v6.83.0`), and GitHub release.
5. [ ] **Subtask 05:** Multi-repository fleet synchronization across all 43 connected repos via `python 03-ai-scripts/38-sync-prompts-skills-scripts.py`.
