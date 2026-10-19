# Master Plan: GitMap Command Examples, A=3 Subagent Spawning Enforcement, Heatmap Telemetry, & Fleet Sync

## User Request (Verbatim)
> check promtps why not using gitmap commands add as an exaple also new the prompt sare not swpaing 3 agents properly check and correctlymention in the execute v6 prompt and also make sure to use heatmap commands , also not using at all fix and sync with skills and other repos please

---

## 1. Executive Summary & Root Cause Analysis

### A. Root Cause 1: Fallback to Shell Listing (`Get-ChildItem`)
- **Evidence:** The user screenshot showed an AI agent executing `Get-ChildItem -Path "src-tauri/src/modules/repo_db"` while labeling the tool call `GitMap find conversation_id in repo_db >`.
- **Cause:**
  1. The ban in `AGENTS.md` and prompt templates previously stated `Get-ChildItem -Recurse`. This led language models to assume that `Get-ChildItem` without `-Recurse` or `dir` or `ls` was permitted.
  2. Prompts and skills lacked prominent, explicit copy-paste examples showing how to list files in a folder using GitMap (`gitmap lf <dir>` or `gitmap find "*.*" <dir>`) or search code within a folder (`gitmap aum search "<pattern>" [dir] [-e <.ext>]`).
- **Remediation:**
  1. Enact a TOTAL BAN on ALL forms of shell listing and searching: `Get-ChildItem` (all variants), `dir`, `ls` in shell, `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, and `findstr`.
  2. Provide comprehensive, unmistakable command examples for directory inventory, code search, file globbing, and file streaming.

### B. Root Cause 2: Failure to Spawn 3 Agents Concurrently ($A = 3$)
- **Cause:**
  1. In `execute-parent-task-with-n-steps-v6`, line 40 stated: `2. Spec Step (A = 3 self subagents or lead)`. This "or lead" created an escape hatch allowing agents to skip spawning subagents and author specs solo.
  2. The prompt lacked a complete 3-agent JSON payload example for Phase 1 Spec authoring.
  3. Prompts did not explicitly mandate that every call to `invoke_subagent` must contain an array of EXACTLY 3 subagent objects: `"Subagents": [agent1, agent2, agent3]`.
- **Remediation:**
  1. Eliminate all "or lead" solo authoring escape hatches.
  2. Enforce structural validation: every `invoke_subagent` tool call MUST contain an array of exactly 3 subagents.
  3. Provide concrete JSON tool call templates for all three phases (Planning discovery with 3 `research` subagents, Spec authoring with 3 `self` subagents, Execution with 3 `self` subagents).

### C. Root Cause 3: Heatmap Telemetry Neglect
- **Cause:**
  1. Heatmap commands were documented only in reference tables without an active operational requirement.
  2. Agents were not instructed to execute `gitmap heatmap log` before/after reading or editing files, `gitmap heatmap status` at Step 0, or `gitmap heatmap tree` at verification.
- **Remediation:**
  1. Add Section 15 to `AGENTS.md`: "GitMap Cognitive Heatmap Telemetry Primacy".
  2. Mandate operational heatmap logging in `execute-parent-task-with-n-steps-v6` across Step 0, Phase 1, Phase 2, and Phase 3.
  3. Provide concrete copy-paste command examples for `gitmap heatmap log`, `gitmap heatmap update`, `gitmap heatmap status`, and `gitmap heatmap tree`.

---

## 2. Decomposed Subtasks Ledger

| Subtask ID | Title | Owned Files (Disjoint) | Status | Evidence |
| :--- | :--- | :--- | :--- | :--- |
| `01-subtask` | Git Sync & Preflight | `.ai-memory/plans/subtasks/gitmap-commands-agent-spawning-heatmap-sync/01-subtask-git-sync-and-preflight.md` | `DONE` | `gitmap pull` clean at `ae83ad9` |
| `02-subtask` | Architecture & Component Specs | `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/*.md` | `DONE` | Authored 01-architecture, 02-component, 03-heatmap specs |
| `03-subtask` | AGENTS.md Updates | `AGENTS.md` | `DONE` | Section 12 shell ban expansion & Section 15 Heatmap Primacy |
| `04-subtask` | Execute V6 Prompts & Skills | `01-prompts/14-execute/*`, `.agents/skills/execute-parent-task-with-n-steps-v6/*`, `.cursor/skills/execute-parent-task-with-n-steps-v6/*` | `DONE` | A=3 enforcement, 3-agent JSON templates & heatmap gates |
| `05-subtask` | GitMap Skills Updates | `.agents/skills/gitmap/*`, `.cursor/skills/gitmap/*` | `DONE` | Command matrix update & 4-gate heatmap docs |
| `06-subtask` | Multi-Repo Fleet Sync & Commit | Fleet repos via `03-ai-scripts/38-sync-prompts-skills-scripts.py` | `DONE` | 41/43 repos released & pushed upstream |

---

## 3. Negative Boundaries & Invariants
- TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`) and `file:///` URIs.
- All created and modified files must use strictly lowercase naming.
- Zero secrets in standard repositories; offload to `repo-secrets` via `gitmap rs`.
- Worker subagents never run git commands.
- Commit message format: `gitmap cpf "<module> - <summary>"` with hyphen separator and no colons in the message argument.
