# Plan: Multi-Repo Sync Date Format and Branch Pruning

## Goal
Implement `dd-MMM-yy` 12-hour am/pm date representation and automated 10-day lifecycle branch pruning in `03-ai-scripts/38-sync-prompts-skills-scripts.py`, mirror to `.agents/scripts/38-sync-prompts-skills-scripts.py`, verify across connected repositories, and commit/push upstream.

## Context & Specifications
- Specification: `02-spec/21-app/34-sync-date-format-and-branch-pruning/01-architecture-spec.md`
- Master Audit Ledger: `02-spec/21-app/34-sync-date-format-and-branch-pruning/00-master-audit-ledger.md`

## Subtasks
- [ ] `01-implement-date-helpers-and-pruning-engine.md`: Add `format_sync_date`, `format_sync_branch_name`, and `prune_older_sync_branches` to `03-ai-scripts/38-sync-prompts-skills-scripts.py`.
- [ ] `02-integrate-pruning-and-mirror-script.md`: Wire branch pruning and new timestamped backup branches into `sync_repo` flow, mirror to `.agents/scripts/38-sync-prompts-skills-scripts.py`.
- [ ] `03-verification-and-upstream-push.md`: Test dry-run and live executions, verify date format and pruning, commit with GitMap and push upstream.
