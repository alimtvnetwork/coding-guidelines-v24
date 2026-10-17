# Subtask 03: Heat Map Action Decision Logging & CLI Visualization

## Objective
Embed granular in-flight action logging (`gitmap agent log`) with explicit reasoning into `execute-parent-task-with-n-steps-v6` so that file reads, modifications, and rationale are indexed in the SQLite task database and rendered via `gitmap heatmap`.

## Targets
- `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`
- `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`
- `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`
- `D:\work\gitmap` (CLI command `gitmap heatmap`)

## Acceptance Criteria
- Mandatory file access action logging: `gitmap agent log --subtask <id> --agent "<name>" --action "<read|modify|plan>" --file "<relative-path>" --details "<reasoning>"`.
- `gitmap heatmap` command in `D:\work\gitmap` displays active file touches and reasoning.
