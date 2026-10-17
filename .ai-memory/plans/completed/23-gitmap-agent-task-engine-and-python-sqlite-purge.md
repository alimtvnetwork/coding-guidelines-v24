# Plan: GitMap Agent Task Engine & Python SQLite Purge

> **/goal** Autonomously purge all references to `46-agent-sqlite-task-manager.py` and ad-hoc `python -c "import sqlite3..."` across all prompts, skills, and scripts, replace them with native compiled `gitmap agent` CLI commands, and deploy non-destructively across all 43 connected repositories.
> **/learn** Enforce the native Go `gitmap agent` engine, verify relative paths, and execute fleet-wide synchronization.

## Canonical Reference

- Specification: [`02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/01-overview-and-architecture.md`](../../../02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/01-overview-and-architecture.md)
- Acceptance Criteria: [`02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/02-acceptance-criteria.md`](../../../02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/02-acceptance-criteria.md)

## Subtasks

1. [x] **Subtask 01: Purge Prompts & Sync Manifest** ([`01-purge-prompts-and-sync-manifest.md`](../subtasks/23-gitmap-agent-task-engine-and-python-sqlite-purge/01-purge-prompts-and-sync-manifest.md)) — Purged `46-agent-sqlite-task-manager.py` and ad-hoc `python -c "import sqlite3..."` across active prompts (`01-prompts/14-execute/01-execute-pending-tasks.md`, `02-execute-parent-task-with-n-steps-v6.md`, `01-prompts/25-ai-verification/01-retrospective-ai-verification.md`, `01-prompts/19-old-execute-prompts/`, `01-prompts/v4/`), specs, and sync scripts.
2. [x] **Subtask 02: Update & Harden Companion Skills** ([`02-update-and-harden-skills.md`](../subtasks/23-gitmap-agent-task-engine-and-python-sqlite-purge/02-update-and-harden-skills.md)) — Hardened `.agents/skills/gitmap/skill.md`, `.cursor/skills/gitmap/skill.md`, `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`, `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`, and `AGENTS.md` (Section 14) with native `gitmap agent` subcommands and total ban on python sqlite one-liners.
3. [x] **Subtask 03: Verify, Commit & Deploy Fleet Synchronization** ([`03-verify-commit-and-fleet-sync.md`](../subtasks/23-gitmap-agent-task-engine-and-python-sqlite-purge/03-verify-commit-and-fleet-sync.md)) — Validated relative paths (3686 files clean), committed to `coding-guidelines` via `gitmap cpf` (`9603c708`), and deployed non-destructively across 43 repositories via `38-sync-prompts-skills-scripts.py` (+11/-0, zero repo-owned file deletions).

## Quality Gates

- [x] Zero occurrences of `46-agent-sqlite-task-manager.py` in active prompts.
- [x] Total ban on `python -c "import sqlite3..."` explicitly codified.
- [x] Relative paths linter passes with exit code 0.
- [x] All 43 connected repositories synchronized cleanly with 0 repo-owned file deletions.
- [x] Upstream completion invariant ('No Push = Not Done') strictly satisfied.
