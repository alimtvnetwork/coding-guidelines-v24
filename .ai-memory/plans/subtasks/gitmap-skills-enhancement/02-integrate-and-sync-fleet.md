# Subtask: Commit Locally & Synchronize 43 Repositories

## Metadata
- **Subtask ID:** `02-integrate-and-sync-fleet`
- **Parent Plan:** `.ai-memory/plans/gitmap-skills-enhancement.md`
- **Spec Reference:** `02-spec/21-app/gitmap-skills-enhancement/01-gitmap-skills-architecture.md`
- **Status:** Pending

## Objective
Commit the newly authored specs, plans, and enhanced GitMap skills into `coding-guidelines`, then run the multi-repo synchronization script (`03-ai-scripts/38-sync-prompts-skills-scripts.py`) to propagate the updated skills and prompts across all 43 repositories.

## Execution Sequence
1. Check repository status using `gitmap status --dirty` and verify clean diffs.
2. Commit in `coding-guidelines` using hyphen format:
   `gitmap cpf "skills - expand gitmap skill suite with authentication status and runners"`
3. Execute dry-run multi-repo sync:
   `python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run`
4. Execute live multi-repo sync across all 43 repositories:
   `python 03-ai-scripts/38-sync-prompts-skills-scripts.py`
5. Verify that each repository was pulled, backed up with `backup/sync-<timestamp>`, updated, and that the 5 Non-Negotiable Boundaries were strictly honored.

## Acceptance Criteria
- [ ] Working tree in `coding-guidelines` committed cleanly.
- [ ] All 43 repositories updated with enhanced `gitmap` skills.
- [ ] No private specs (`02-spec/21-*`), modified AI scripts, bump scripts, or memory folders overwritten.
