# Subtask 03: Multi-Repository Leak Detection & Cleanup

## Task Details
- **Parent Task:** `25-coding-guidelines-and-folder-structure-instructions`
- **Subtask Code:** `Subtask-03`
- **Slug:** `03-multi-repo-leak-cleanup`
- **Status:** DONE
- **Plan Motto:** Improved work and less friction

## Owned Files & Scope
- `cat-my: 02-spec/21-app/` (surgically delete 3 leaked folders, preserve 80+ native app specs)
- `slack-ai-agent: guidelines/02-spec/21-app/` (surgically delete 7 leaked files)
- 42 downstream repositories containing `01-prompts/v4/`

## Actionable Steps
1. [x] In `cat-my`:
   - Surgically delete `02-spec/21-app/01-variadic-spread-params-and-multi-repo`.
   - Surgically delete `02-spec/21-app/02-string-normalization-and-equalfoldany`.
   - Surgically delete `02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance`.
   - Delete `01-prompts/v4/`.
   - Ensure all other 80+ native photo editor specs are completely untouched.
   - Commit via `gitmap cpc "spec - surgically purge leaked coding guidelines folders"`.
2. [x] In `slack-ai-agent`:
   - Surgically delete the 7 leaked files in `guidelines/02-spec/21-app/`.
   - Delete `01-prompts/v4/`.
   - Commit via `gitmap cpc "spec - surgically purge leaked coding guidelines files"`.
3. [x] Across remaining downstream repositories in workspace:
   - Delete `01-prompts/v4/` directory wherever present.
   - Commit via `gitmap cpc "chore - purge deprecated v4 prompts"`.
4. [x] In `coding-guidelines/03-ai-scripts/38-sync-prompts-skills-scripts.py`:
   - Verify `"v4"` is in `EXCLUDE_NAMES` so future synchronization never replicates `v4`.
   - Verify Rule 11 boundary enforcement (Spec 21 exclusion).

## Acceptance Criteria
- Zero leaked files in `cat-my` and `slack-ai-agent`.
- Zero `v4` prompt directories across all repositories in workspace.
- All native application specs in downstream repos are 100% preserved.
