# Master Plan: Fleet Commit, Push & Pooling Invariant

## Executive Summary
This master plan enforces the mandatory workspace-wide commit, push, and pooling invariant: "If code is not committed to Git and pushed upstream to GitHub, the task is strictly considered INCOMPLETE and NOT DONE." It implements non-owned repository filtering (`oh-my-zsh`, `zsh`, `omis`, dotfiles), pre-commit pooling/pulling, embeds the completion agreement into all V6 prompts and skills, and synchronizes the fleet.

## Specifications & Rules
- Spec: `02-spec/21-app/fleet-commit-and-push/01-architecture-spec.md`
- Core Invariant: Rule 2 & 10 in `agents.md`
- Sync Rules: `.ai-memory/memory/learned/18-cross-repository-sync-rules.md`

## Subtask Decomposition
| Subtask ID | Title | Target Files | Status |
| :--- | :--- | :--- | :--- |
| `01-script-and-pooling-enhancements` | Enhance 49-commit-and-push script with non-owned exclusion and pooling | `03-ai-scripts/49-commit-and-push-all-repos.py` | Pending |
| `02-v6-prompt-and-skill-invariants` | Embed 'No Push = Not Done' into V6 prompts and skills | `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`, `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`, `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md` | Pending |
| `03-fleet-sync-and-verification` | Update commit-and-push prompts/skills, commit, and sync 43 repositories | `01-prompts/09-commit-and-multi-agent-code-fix/09-commit-and-push-all-repos.md`, `.agents/skills/commit-and-push-all-repos/skill.md`, `.cursor/skills/commit-and-push-all-repos/skill.md` | Pending |

## Execution Protocol
- **Task DB:** `.ai-memory/temp-agents/03-fleet-commit-push-and-pooling-invariant/agent-task.db`
- **Subagent Concurrency:** A = 2, H = 2
- **Path Hygiene:** Strict relative paths from git root (zero absolute paths and zero `file:///` URIs)
