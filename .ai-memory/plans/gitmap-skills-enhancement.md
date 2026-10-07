# Master Plan: GitMap Skills Expansion & Cross-Fleet Synchronization

## Executive Summary
This master plan coordinates the enhancement of GitMap skills within `coding-guidelines` and ensures their complete synchronization across all 43 connected repositories in the workspace. It introduces comprehensive command coverage (login, status, py, ps, bash, storage, search, commit), explicit negative-versus-positive replacement matrices, and validates synchronization with safety backups.

## Specifications & Guidelines
- Architecture Spec: `02-spec/21-app/gitmap-skills-enhancement/01-gitmap-skills-architecture.md`
- Core Guidelines: `02-spec/02-coding-guidelines/`
- Sync Rules: `.ai-memory/memory/learned/18-cross-repository-sync-rules.md`

## Subtask Decomposition
| Subtask ID | Title | File Reference | Status |
| :--- | :--- | :--- | :--- |
| `01-expand-gitmap-skills` | Expand GitMap Skill Definition & Matrix | `.ai-memory/plans/subtasks/gitmap-skills-enhancement/01-expand-gitmap-skills.md` | Pending |
| `02-integrate-and-sync-fleet` | Commit Locally & Synchronize 43 Repositories | `.ai-memory/plans/subtasks/gitmap-skills-enhancement/02-integrate-and-sync-fleet.md` | Pending |

## Execution Protocol
- **Task DB:** `.ai-memory/temp-agents/02-gitmap-skills-enhancement/agent-task.db`
- **Subagent Concurrency:** Bounded micro-tasking via `execute-parent-task-with-n-steps-v6`
- **Zero-Storage Actions:** 0.0 GB GitHub Actions storage footprint
- **Path Hygiene:** Strict relative paths from repository root (TOTAL BAN on absolute paths and `file:///` URIs)
