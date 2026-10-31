# Subtask 02: Update & Harden Companion Skills

> **/goal** Update and harden companion skills (`gitmap`, `execute-parent-task-with-n-steps-v6`, `ai-verification`) with complete `gitmap agent` subcommands and explicit prohibition against ad-hoc Python SQLite one-liners.
> **/learn** Ensure skills provide unambiguous command tables for task status inspection, subtask listing, and crash forensics.

## Canonical Reference

- Specification: `02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/01-overview-and-architecture.md`

## Action Checklist

- [ ] Update `.agents/skills/gitmap/skill.md`:
  - Section 10: Document full suite of `gitmap agent task` and `gitmap agent subtask` commands (`gitmap agent task status`, `gitmap agent subtask ls`, `gitmap agent task ls`, `gitmap agent task pending`, `gitmap agent task completed`).
  - Add explicit operational invariant banning `python -c "import sqlite3..."` and raw Python DB scripts.
- [ ] Mirror updates to `.cursor/skills/gitmap/skill.md`.
- [ ] Update `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` and `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`:
  - Codify native `gitmap agent` commands in preflight and worker sections.
  - Mandate that agents inspect task state via `gitmap agent task status` and `gitmap agent subtask ls` rather than raw Python scripts.
- [ ] Update `.agents/skills/ai-verification/skill.md` with native `gitmap agent` references.
