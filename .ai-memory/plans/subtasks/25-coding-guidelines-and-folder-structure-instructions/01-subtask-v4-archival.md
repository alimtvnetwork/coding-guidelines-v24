# Subtask 01: v4 Prompts Archival & Skill Purge

## Task Details
- **Parent Task:** `25-coding-guidelines-and-folder-structure-instructions`
- **Subtask Code:** `Subtask-01`
- **Slug:** `01-v4-archival`
- **Status:** PENDING
- **Plan Motto:** Improved work and less friction

## Owned Files & Scope
- `01-prompts/v4/` (to be removed)
- `06-archive/v4/` (to receive the 5 missing files)
- `.ai-memory/prompts.md` (prune lines 225-280)
- `03-ai-scripts/38-sync-prompts-skills-scripts.py` (add `"v4"` to exclusions)
- `03-ai-scripts/47-upgrade-cg-execute-to-v6.py` (remove `V4_PROMPTS_DIR`)
- `03-ai-scripts/48-standardize-execution-prompts.py` (remove `V4_DIR`)

## Actionable Steps
1. [ ] Copy 5 missing files from `01-prompts/v4/` to `06-archive/v4/`:
   - `01-prompts/v4/14-execute/08-execute-with-e2e.md` -> `06-archive/v4/14-execute/08-execute-with-e2e.md`
   - `01-prompts/v4/15-cg-execute/38-map-dispatch-and-handler-lookup.md` -> `06-archive/v4/15-cg-execute/38-map-dispatch-and-handler-lookup.md`
   - `01-prompts/v4/24-sync/01-sync.md` -> `06-archive/v4/24-sync/01-sync.md`
   - `01-prompts/v4/24-sync/02-sync-other-codebase.md` -> `06-archive/v4/24-sync/02-sync-other-codebase.md`
   - `01-prompts/v4/24-sync/readme.md` -> `06-archive/v4/24-sync/readme.md`
2. [ ] Verify `06-archive/v4/` contains all 56 archived prompt files.
3. [ ] Delete `01-prompts/v4/` directory from `coding-guidelines`.
4. [ ] Prune lines 225–280 from `.ai-memory/prompts.md`.
5. [ ] Update `03-ai-scripts/38-sync-prompts-skills-scripts.py` to add `"v4"` to `EXCLUDE_NAMES`.
6. [ ] Update scripts 47 & 48 to remove `V4_PROMPTS_DIR` / `V4_DIR` references.
7. [ ] Run `python linter-scripts/check-prompts-loaded.py` to confirm zero prompt drift.

## Acceptance Criteria
- `gitmap find "*.*" 01-prompts/v4` returns empty.
- `gitmap find "*.*" 06-archive/v4` contains 56 files.
- Linter passes with 0 errors.
