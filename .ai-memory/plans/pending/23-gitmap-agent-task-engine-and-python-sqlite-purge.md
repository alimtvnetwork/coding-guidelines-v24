# Plan: GitMap Agent Task Engine & Python SQLite Purge

> **/goal** Autonomously purge all references to `46-agent-sqlite-task-manager.py` and ad-hoc `python -c "import sqlite3..."` across all prompts, skills, and scripts, replace them with native compiled `gitmap agent` CLI commands, and deploy non-destructively across all 43 connected repositories.
> **/learn** Enforce the native Go `gitmap agent` engine, verify relative paths, and execute fleet-wide synchronization.

## Canonical Reference

- Specification: [`02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/01-overview-and-architecture.md`](../../../02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/01-overview-and-architecture.md)
- Acceptance Criteria: [`02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/02-acceptance-criteria.md`](../../../02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/02-acceptance-criteria.md)

## Subtasks

1. [ ] **Subtask 01: Purge Prompts & Sync Manifest** ([`01-purge-prompts-and-sync-manifest.md`](../subtasks/23-gitmap-agent-task-engine-and-python-sqlite-purge/01-purge-prompts-and-sync-manifest.md))
2. [ ] **Subtask 02: Update & Harden Companion Skills** ([`02-update-and-harden-skills.md`](../subtasks/23-gitmap-agent-task-engine-and-python-sqlite-purge/02-update-and-harden-skills.md))
3. [ ] **Subtask 03: Verify, Commit & Deploy Fleet Synchronization** ([`03-verify-commit-and-fleet-sync.md`](../subtasks/23-gitmap-agent-task-engine-and-python-sqlite-purge/03-verify-commit-and-fleet-sync.md))

## Quality Gates

- Zero occurrences of `46-agent-sqlite-task-manager.py` in active prompts.
- Total ban on `python -c "import sqlite3..."` explicitly codified.
- Relative paths linter passes with exit code 0.
- All 43 connected repositories synchronized cleanly with 0 repo-owned file deletions.
- Upstream completion invariant ('No Push = Not Done') strictly satisfied.
