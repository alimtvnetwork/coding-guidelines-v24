# Subtask 01: Purge Prompts & Sync Manifest

> **/goal** Purge all references to `46-agent-sqlite-task-manager.py` and ad-hoc `python -c "import sqlite3..."` across active prompts, scripts, and sync manifests.
> **/learn** Replace all legacy calls with native `gitmap agent` subcommands.

## Canonical Reference

- Specification: `02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/01-overview-and-architecture.md`

## Action Checklist

- [ ] Update `01-prompts/14-execute/01-execute-pending-tasks.md`:
  - Replace `python 03-ai-scripts/46-agent-sqlite-task-manager.py init` with `gitmap agent task init --name "<slug>" --budget <N>`.
  - Replace `python 03-ai-scripts/46-agent-sqlite-task-manager.py schema` with `gitmap agent task status`.
  - Replace `python 03-ai-scripts/46-agent-sqlite-task-manager.py add-subtasks` with `gitmap agent subtask add --json '<json>'`.
  - Add explicit ban on `python -c "import sqlite3..."`.
- [ ] Update `01-prompts/25-ai-verification/01-retrospective-ai-verification.md`:
  - Replace `46-agent-sqlite-task-manager.py init` with `gitmap agent task init`.
  - Replace `46-agent-sqlite-task-manager.py` action logging with `gitmap agent log`.
- [ ] Update `01-prompts/19-old-execute-prompts/03-execute-parent-task-with-n-steps.md` & `04-parent-task-in-below-steps.md`.
- [ ] Update `01-prompts/v4/14-execute/01-execute-pending-tasks.md`.
- [ ] Update `03-ai-scripts/38-sync-prompts-skills-scripts.py` to remove `46-agent-sqlite-task-manager.py` from sync targets.
- [ ] Update `03-ai-scripts/47-upgrade-cg-execute-to-v6.py`, `.agents/scripts/47-upgrade-cg-execute-to-v6.py`, and `03-ai-scripts/readme.md`.
