# Subtask 01: Specs Authoring and Non-Destructive Engine Hardening

## Overview
Author detailed specifications in `02-spec/21-app/17-multi-repo-skills-sync/` and update `03-ai-scripts/38-sync-prompts-skills-scripts.py` to enforce the non-destructive rule preventing any deletions of repo-owned files.

## Actions
1. Author specifications:
   - `02-spec/21-app/17-multi-repo-skills-sync/01-overview-and-architecture.md`
   - `02-spec/21-app/17-multi-repo-skills-sync/02-acceptance-criteria.md`
   - `02-spec/21-app/17-multi-repo-skills-sync/03-non-destructive-sync-spec.md`
   - `02-spec/21-app/17-multi-repo-skills-sync/readme.md`
2. Update `03-ai-scripts/38-sync-prompts-skills-scripts.py`:
   - Disable or guard stale file deletion in `mirror_directory` so that no files created by target repositories are deleted.
   - Ensure `TARGET_REPOS` accurately reflects all 43 repositories provided in the user instruction.
3. Verify compliance with strict relative paths and affirmative boolean principles.
