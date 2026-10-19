# Plan: Skills and Prompts Synchronization and Configuration

- **Task Slug:** `skills-and-prompts-synchronization-and-configuration`
- **Spec:** `02-spec/21-app/skills-and-prompts-synchronization-and-configuration/01-architecture-spec.md`
- **Goal:** Verify prompts-to-skills migration coverage across the repository, upgrade Muse master and generic prompts to default to 400 steps with aggressive write-and-execute power, synchronize across all 43 repositories via `gitmap sync`, and perform a minor release bump to `v6.82.0`.
- **Status:** COMPLETED (5/5 subtasks done)

## Execution Subtasks

1. [x] **Subtask 01:** Architecture spec and plan creation in `02-spec/21-app/skills-and-prompts-synchronization-and-configuration/` and `.ai-memory/plans/`.
2. [x] **Subtask 02:** Prompts-to-skills migration audit and coverage verification across `01-prompts/`, `.agents/skills/`, and `.cursor/skills/`.
3. [x] **Subtask 03:** Muse prompt 400-step upgrade (`N = 400`, Phase 1 = 100, Phase 2 = 150, Phase 3 = 150) with powerful write-execute directives across `01-muse-master-prompt.md`, `02-muse-generic.md`, and skill mirrors.
4. [x] **Subtask 04:** Repository verification, internal sync (`npm run sync`, `npm run sync:check`), package installers, and minor bump release ceremony (`v6.82.0`).
5. [x] **Subtask 05:** Multi-repository fleet synchronization across all 43 connected repos via `python 03-ai-scripts/38-sync-prompts-skills-scripts.py`.

## Final Outcomes
- Confirmed 100% prompt-to-skill parity (128 Antigravity skills and 128 Cursor skills with valid YAML frontmatter).
- Upgraded Muse master and generic prompts to `N = 400` default budget with phase budgets (100 / 150 / 150) and enforced the powerful write-and-execute invariant.
- Released `v6.82.0` with tagged release assets on GitHub.
- Synchronized across the fleet with 41 repositories committed, tagged, and pushed.
