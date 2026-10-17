# GitMap Agent Task Engine & Python SQLite Purge Specification

> **/goal** Master and enforce the native compiled GitMap agent task engine, eliminating all legacy references to Python SQLite task manager scripts and ad-hoc Python SQLite one-liners across all prompts, skills, and connected repositories.
> **/learn** Read the numbered specification files in this directory, apply mandatory GitMap agent CLI command mappings, eliminate legacy script references, update companion skills, and synchronize cleanly across the fleet.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Read and understand `01-overview-and-architecture.md` and `02-acceptance-criteria.md`.
- [ ] `/learn` Total ban on `python 03-ai-scripts/46-agent-sqlite-task-manager.py` and `python -c "import sqlite3..."`.
- [ ] `/goal` Use native compiled `gitmap agent` subcommands (`task init`, `task status`, `task ls`, `subtask add`, `subtask claim`, `subtask complete`, `subtask fail`, `log`, `diagnose`, `crashed`, `temp-clear`).
- [ ] `/learn` Verify zero explicit `true` boolean evaluations and strict relative Git paths.
- [ ] `/learn` Deploy synchronized assets non-destructively across all 43 connected repositories via `03-ai-scripts/38-sync-prompts-skills-scripts.py`.

---

## Documents

| Document | Title | Description |
|:---|:---|:---|
| [`01-overview-and-architecture.md`](01-overview-and-architecture.md) | Architecture & Command Mapping | Complete architectural transition from legacy Python task scripts to native Go compiled `gitmap agent` CLI. |
| [`02-acceptance-criteria.md`](02-acceptance-criteria.md) | Verification & Quality Gates | Strict quality gates, ban on ad-hoc Python sqlite3, and non-destructive fleet sync acceptance criteria. |
