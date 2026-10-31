# Archived Prompt Tiers & Deprecated Collections

This directory stores historical versioned prompt tiers (`execute/`, `v1/`, `v2/`, `v3/`, `v4/`) and deprecated prompt collections (`12-old-plan-prompts/`, `18-insults/`, `19-old-execute-prompts/`) for archival and reference inside the `coding-guidelines` meta-repository.

## Archived Collections

- **`execute/`**: Archived legacy parent task execution prompt versions (`v2`, `v3`, `v4`, `v5`, and legacy `excute-parent-old.md`).
- **`v1/`**: Legacy V1 prompt tier.
- **`v2/`**: Legacy V2 prompt tier.
- **`v3/`**: Legacy V3 prompt tier.
- **`v4/`**: Legacy V4 prompt tier.
- **`12-old-plan-prompts/`**: Legacy planning protocols (`01-plan-maximum-enforcement.md`, `02-plan-steps.md`, `03-plan-steps-by-groups.md`, `04-plan-maximum-enforcement-v5.md`, `05-plan-spec-steps.md`, `06-audit-app-spec-old.md`, `07-fix-spec-from-audit-old.md`) superseded by V2 spec planning and V6 execution.
- **`18-insults/`**: Legacy insult prompts (`01-raw-insults.md`, `02-consolidated-insults-v2.md`).
- **`19-old-execute-prompts/`**: Legacy execute prompt variants and backups (`01-execute-robust-loop.md`, `02-fix-subtask-naming-convention.md`, `03-execute-parent-task-with-n-steps.md`, `04-parent-task-in-below-steps.md`, and `backup-v6.46.0/`) superseded by `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`.

## Active Prompts & Canonical Rotation

The active, canonical prompts live directly under [`01-prompts/`](01-prompts/) (including canonical V6 execute prompts in [`01-prompts/14-execute/`](01-prompts/14-execute/)).

## Sync Exclusion Policy

`06-archive/` is **NEVER** synchronized to downstream repositories. Connected repositories strictly receive active prompts from [`01-prompts/`](01-prompts/), and `03-ai-scripts/38-sync-prompts-skills-scripts.py` strictly excludes `06-archive`, `12-old-plan-prompts`, `18-insults`, `19-old-execute-prompts`, `05-changes-history`, and `skills-archived` from synchronization.
