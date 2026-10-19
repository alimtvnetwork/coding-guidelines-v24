# Plan: Skills and Prompts Synchronization and Configuration

- **Task Slug:** `skills-and-prompts-synchronization-and-configuration`
- **Spec:** `02-spec/21-app/skills-and-prompts-synchronization-and-configuration/01-architecture-spec.md`
- **Goal:** Verify prompts-to-skills migration coverage across the repository, upgrade Muse master and generic prompts to default to 400 steps with aggressive write-and-execute power, synchronize across all 43 repositories via `gitmap sync`, and perform a minor release bump to `v6.82.0`.

## Execution Subtasks

1. [x] **Subtask 01:** Architecture spec and plan creation in `02-spec/21-app/skills-and-prompts-synchronization-and-configuration/` and `.ai-memory/plans/`.
2. [ ] **Subtask 02:** Prompts-to-skills migration audit and coverage verification across `01-prompts/`, `.agents/skills/`, and `.cursor/skills/`.
3. [ ] **Subtask 03:** Muse prompt 400-step upgrade (`N = 400`, Phase 1 = 100, Phase 2 = 150, Phase 3 = 150) with powerful write-execute directives across `01-muse-master-prompt.md`, `02-muse-generic.md`, and skill mirrors.
4. [ ] **Subtask 04:** Repository verification, internal sync (`npm run sync`, `npm run sync:check`), package installers, and minor bump release ceremony (`v6.82.0`).
5. [ ] **Subtask 05:** Multi-repository fleet synchronization across all 43 connected repos via `gitmap sync --workers 8`.
