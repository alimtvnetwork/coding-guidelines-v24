# GitMap Agent Task Engine & Python SQLite Purge: Acceptance Criteria

> **/goal** Define verifiable acceptance criteria and testable quality gates for the complete eradication of Python SQLite task manager scripts and ad-hoc Python one-liners in favor of native `gitmap agent` CLI commands.
> **/learn** Execute automated verification gates against relative paths, command syntax, and multi-repo synchronization.

## Acceptance Criteria Matrix

| ID | Category | Requirement | Verification Method | Target Status |
|:---|:---|:---|:---|:---|
| **AC-01** | **Ghost Script Purge** | Zero occurrences of `46-agent-sqlite-task-manager.py` in active prompts (`01-prompts/14-execute/`, `01-prompts/25-ai-verification/`, `01-prompts/19-old-execute-prompts/`, `01-prompts/v4/`). | `gitmap aum search "46-agent-sqlite" 01-prompts` returns 0 hits. | PASS |
| **AC-02** | **Sync Manifest Cleanliness** | `46-agent-sqlite-task-manager.py` removed from `03-ai-scripts/38-sync-prompts-scripts.py` sync targets. | Code review of `03-ai-scripts/38-sync-prompts-skills-scripts.py`. | PASS |
| **AC-03** | **Python SQLite Ban** | Total ban on `python -c "import sqlite3..."` explicitly documented in `AGENTS.md`, `gitmap` skills, and execution prompts. | Verified in prompt and skill text. | PASS |
| **AC-04** | **Native Command Grounding** | All task initialization, claiming, logging, and status commands use `gitmap agent` CLI (`task init`, `task status`, `subtask add`, `subtask claim`, `subtask complete`, `log`, `diagnose`). | Prompts and skills reference valid `gitmap agent` commands. | PASS |
| **AC-05** | **Companion Skills Hardened** | `.agents/skills/gitmap/skill.md`, `.cursor/skills/gitmap/skill.md`, `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`, and `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md` updated with complete `gitmap agent` command table and query workflows. | `gitmap aum search "gitmap agent" .agents/skills` returns verified hits. | PASS |
| **AC-06** | **Relative Path Hygiene** | Zero absolute filesystem paths (`/absolute/path/to/`, `C:\`, `d:\`) and zero `file:///` URIs across all modified files. | `python linter-scripts/check-relative-paths.py` exits 0. | PASS |
| **AC-07** | **Non-Destructive Fleet Sync** | Canonical prompts, skills, specs, and scripts synchronized across all 43 connected repositories with zero repo-owned file deletions (`-0`). | `python 03-ai-scripts/38-sync-prompts-skills-scripts.py` runs across 43 repos with 0 errors. | PASS |
| **AC-08** | **Upstream Push Invariant** | Coding-guidelines repository and all connected repositories cleanly committed via `gitmap cpf` and pushed to GitHub tracking branches. | `git status` reports working tree clean and up to date with `origin/main`. | PASS |

---

## Quality Gate Checklist

1. [ ] **Step 1:** Author specifications under `02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/`.
2. [ ] **Step 2:** Register specification in `02-spec/21-app/readme.md`.
3. [ ] **Step 3:** Enqueue plan in `.ai-memory/plans/pending/23-gitmap-agent-task-engine-and-python-sqlite-purge.md`.
4. [ ] **Step 4:** Replace `46-agent-sqlite-task-manager.py` and add python sqlite ban in:
   - `01-prompts/14-execute/01-execute-pending-tasks.md`
   - `01-prompts/25-ai-verification/01-retrospective-ai-verification.md`
   - `01-prompts/19-old-execute-prompts/03-execute-parent-task-with-n-steps.md`
   - `01-prompts/19-old-execute-prompts/04-parent-task-in-below-steps.md`
   - `01-prompts/v4/14-execute/01-execute-pending-tasks.md`
5. [ ] **Step 5:** Update skills:
   - `.agents/skills/gitmap/skill.md` & `.cursor/skills/gitmap/skill.md`
   - `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` & `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`
   - `.agents/skills/ai-verification/skill.md`
6. [ ] **Step 6:** Update `03-ai-scripts/38-sync-prompts-skills-scripts.py`, `03-ai-scripts/47-upgrade-cg-execute-to-v6.py`, and `03-ai-scripts/readme.md`.
7. [ ] **Step 7:** Run relative paths linter and commit via `gitmap cpf` in coding-guidelines.
8. [ ] **Step 8:** Deploy fleet synchronization across all 43 connected repositories via `03-ai-scripts/38-sync-prompts-skills-scripts.py`.
