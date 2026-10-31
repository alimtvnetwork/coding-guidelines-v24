# Subtask 01: Implement Date Helpers and Pruning Engine

## Scope
File: `03-ai-scripts/38-sync-prompts-skills-scripts.py`

## Instructions
1. Implement `format_sync_date(dt: datetime | None = None) -> str`:
   - Returns date formatted as `dd-MMM-yy hh:mm a` (e.g. `28-Oct-26 10:11 am`).
2. Implement `format_sync_branch_name(dt: datetime | None = None) -> str`:
   - Returns branch-safe string `backup/sync-dd-MMM-yy-hh-mma` (e.g. `backup/sync-28-Oct-26-10-11am`).
3. Implement `prune_older_sync_branches(target: Path, max_age_days: int = 10, is_dry_run: bool = False, is_push_disabled: bool = False) -> tuple[list[str], list[str]]`:
   - Queries `git for-each-ref --format="%(refname)|%(committerdate:unix)" refs/heads/ refs/remotes/origin/`.
   - Filters refs containing `sync` under `backup/` or `feat/`.
   - Compares commit timestamp against `now - max_age_days * 86400`.
   - Safely deletes local branches via `git branch -D` and remote branches via `git push origin --delete`.
   - Returns `(pruned_local_branches, pruned_remote_branches)`.

## Status: COMPLETED
- `format_sync_date` implemented and returns `dd-MMM-yy hh:mm a`.
- `format_sync_branch_name` implemented and returns `backup/sync-dd-MMM-yy-hh-mma`.
- `prune_older_sync_branches` queries `committerdate:unix` and deletes older branches in batches.

