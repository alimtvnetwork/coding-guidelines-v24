# Master Plan: Execute Parent Task and Prompt Fixes (Understood Gate & Azure AI Integration)

## User Request (Verbatim)
> From the execute prompts in V6, execute parent task in V6, this is the prompt. Previously, it mentioned the Azure AI, did it understood or not? Currently, it does not mention that. Fix it in the execute prompt V6 and all other similar prompts in CG execute and all others that have the similar pattern, and also fix these skills and deploy these to all other repositories as a skill and prompts. Follow the process.
> 
> ## task: Execute Parent Task and Prompt Fixes
> ## slug: execute-parent-task-and-prompt-fixes

---

## 1. Executive Summary & Root Cause Analysis

### A. The Understood / AI Comprehension Omission
- **Issue:** The prompt instructions for Phase 1A mandate an "Extreme Understood Turn-1 Showcase". However, the copy-paste markdown schema under `### 📋 Confirmed Task Breakdown & Requirement Ingestion` previously omitted the explicit field:
  `- **Understood / AI Comprehension:** [YES] — [1–2 concise sentences proving deep understanding of intent, scope, and verified constraints; answering did the AI understand the requirements or not]`
- **Consequence:** Language models strictly reproducing the markdown template omitted the explicit comprehension declaration, resulting in tasks being listed with only `State`, `Intent & Technical Deliverables`, etc. (as verified in the user's screenshot).

### B. The Azure AI & Model Runtime Context Omission
- **Issue:** Modern execution workflows run across diverse environments (Azure AI Foundry, OpenAI, Google Antigravity, Cursor IDE, local models). Prompts lacked an explicit field in the task breakdown declaring the execution runtime, model endpoint, or cloud platform:
  `- **Azure AI / Model Runtime Context:** [Explicitly document Azure AI, Azure OpenAI, or local model endpoint context if applicable, or state 'N/A — Local CLI/GitMap']`
- **Remediation:** Both fields will be integrated into the 6-part task verification schema across `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`, `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`, all 38 prompts in `01-prompts/15-cg-execute/` and `01-prompts/v4/15-cg-execute/`, and companion skills.

---

## 2. Decomposed Subtasks Ledger

| Subtask ID | Title | Owned Files (Disjoint) | Status | Evidence |
| :--- | :--- | :--- | :--- | :--- |
| `01-subtask` | V6 Master Prompts & Skills | `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`, `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`, `.agents/skills/execute-parent-task-with-n-steps-v6/*`, `.cursor/skills/execute-parent-task-with-n-steps-v6/*` | `DONE` | Phase 1A 6-part schema & task templates updated |
| `02-subtask` | CG Execute Prompts Upgrade | `01-prompts/15-cg-execute/*.md`, `01-prompts/v4/15-cg-execute/*.md` (38 prompts, 76 files) | `DONE` | 38 prompts upgraded & mirrored byte-for-byte |
| `03-subtask` | Companion Skills Parity | `.agents/skills/*/skill.md`, `.cursor/skills/*/skill.md` (priority execution skills) | `DONE` | 8 skill files updated with 1:1 schema parity |
| `04-subtask` | Multi-Repo Fleet Sync | 43 target repositories via `03-ai-scripts/38-sync-prompts-skills-scripts.py` | `DONE` | 41/43 repos released & pushed upstream |
| `05-subtask` | Compliance Audit & Commit | Relative paths check, heatmap 100% closure, `gitmap cpf` push | `DONE` | 0 path violations, atomic push upstream |

---

## 3. Negative Boundaries & Invariants
- TOTAL BAN on raw shell searches (`rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `findstr`).
- TOTAL BAN on `Get-ChildItem` (all variants), `dir`, and `ls` in shell.
- TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`) and `file:///` URIs.
- Worker subagents never run git commands.
- Commit message format: `gitmap cpf "<module> - <summary>"` with hyphen separator and no colons in the argument.
