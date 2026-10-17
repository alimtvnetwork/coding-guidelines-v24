# GitMap Agent Task Engine & Python SQLite Purge: Overview & Architecture

> **/goal** Architect the total replacement of legacy Python SQLite task scripts and ad-hoc Python SQLite one-liners with native compiled `gitmap agent` CLI commands across all prompts, skills, and fleet repositories.
> **/learn** Understand the native Go `gitmap agent` engine, command mappings, and crash forensic telemetry.

## 1. User Request (Verbatim)

```text
We should be removing this script from everywhere adn correct this wtih gitmap agent commands please check and also update all skills regardign this and also sync with other repo, clear?/
```

---

## 2. Root Cause & Problem Analysis

In recent autonomous runs across projects (such as `Antigravity-Manager`), agents were observed falling back to cumbersome, error-prone Python one-liners:

```bash
python -c "import sqlite3; conn = sqlite3.connect('.ai-memory/temp-agents/12-prompt-tree-subtree-and-import-export/agent-task.db'); c = conn.cursor(); [print(row) for row in c.execute('SELECT ...')]
```

### Root Cause Identification:
1. **Ghost Script References:** Legacy prompt files (`01-prompts/14-execute/01-execute-pending-tasks.md`, `01-prompts/25-ai-verification/01-retrospective-ai-verification.md`, `01-prompts/19-old-execute-prompts/`, and `01-prompts/v4/`) still instructed agents to call `python 03-ai-scripts/46-agent-sqlite-task-manager.py`.
2. **Missing Script:** The script `46-agent-sqlite-task-manager.py` did not exist in target repositories, causing agents to attempt manual SQLite inspection via inline Python scripts.
3. **Lack of Query Command Awareness:** Agents lacked clear documentation on how to inspect, list, and interrogate task and subtask status using native GitMap commands.
4. **Performance Penalty:** Spawning Python runtimes (`python.exe` / `python3`) introduces a 150–300ms startup penalty per tool invocation, whereas native compiled Go (`gitmap agent`) executes in under 1ms.

---

## 3. The Two Absolute Bans

### 3.1 Total Ban on Python SQLite Task Manager Scripts
- **BANNED:** `python 03-ai-scripts/46-agent-sqlite-task-manager.py ...`
- **BANNED:** `.agents/scripts/46-agent-sqlite-task-manager.py`
- All references to `46-agent-sqlite-task-manager.py` are strictly eradicated from all active prompts, skills, scripts, and documentation.

### 3.2 Total Ban on Ad-Hoc Python SQLite One-Liners
- **BANNED:** `python -c "import sqlite3; ..."`
- Agents MUST NOT execute raw Python snippets or custom scripts to query or mutate `.ai-memory/temp-agents/<slug>/agent-task.db`.
- All database operations, task tracking, claiming, completion, action logging, and diagnostics MUST use compiled `gitmap agent` commands.

---

## 4. Native `gitmap agent` Command Mappings

| Operational Phase | Legacy Python Call (❌ FORBIDDEN) | Native GitMap Command (✅ MANDATORY) | Description |
|:---|:---|:---|:---|
| **Initialize Parent Task** | `python 46-agent...py init --name "<slug>" --budget <N>` | `gitmap agent task init --name "<slug>" --budget <N>` | Initializes run dir in `.ai-memory/temp-agents/<slug>/agent-task.db` with WAL mode and schema. |
| **Inspect Status / Schema** | `python 46-agent...py schema` | `gitmap agent task status` | Emits task status, subtask progress, and schema validation. Use `--json` for structured output. |
| **List Parent Tasks** | `python -c "SELECT * FROM ParentTask"` | `gitmap agent task ls` | Lists all active and completed parent tasks across the workspace. |
| **Check Pending Subtasks** | `python -c "SELECT * FROM Subtask WHERE HasCompleted=0"` | `gitmap agent task pending` | Lists or counts pending subtasks across agent runs. |
| **List Completed Tasks** | `python -c "SELECT * FROM ParentTask WHERE HasCompleted=1"` | `gitmap agent task completed` | Lists completed root-level tasks with verification proof. |
| **Enqueue Subtasks** | `python 46-agent...py add-subtasks --tasks-json '[...]'` | `gitmap agent subtask add --json '[...]'` | Enqueues decomposed subtasks into the parent task database. |
| **List Subtasks** | `python -c "SELECT * FROM Subtask"` | `gitmap agent subtask ls` | Lists all subtasks, their assignees, statuses, and file boxes. |
| **Claim Subtask** | `python 46-agent...py claim --agent "Worker NN"` | `gitmap agent subtask claim --agent "Worker NN"` | Atomically claims next available pending subtask for worker. |
| **Declare File Box** | *(None)* | `gitmap agent subtask claim-files` | Pre-declares file leases for collision and contention tracking. |
| **Log In-Flight Action** | `python 46-agent...py log-action ...` | `gitmap agent log --subtask <id> --agent "Worker NN" --action "<action>" --file "<file>" --details "<desc>"` | Writes crash forensics record BEFORE modifying any file. |
| **Complete Subtask** | `python 46-agent...py complete --subtask-id <id> ...` | `gitmap agent subtask complete <id> --agent "Worker NN" --evidence "<proof>"` | Marks subtask done with concrete verification evidence. |
| **Fail Subtask** | `python 46-agent...py fail --subtask-id <id> ...` | `gitmap agent subtask fail <id> --agent "Worker NN" --reason "<reason>"` | Marks subtask failed and records root-cause failure. |
| **Crash Autopsy** | *(None)* | `gitmap agent crashed` | Detects crashed or abandoned agents and outputs autopsy report. |
| **Forensic Diagnosis** | `python 46-agent...py diagnose` | `gitmap agent diagnose` | Performs deep forensic inspection of agent tasks and action logs. |
| **Clean Ephemeral State** | `rm -rf .ai-memory/temp-agents/` | `gitmap agent temp-clear` | Safely purges ephemeral task run directory. |

---

## 5. Skills Hardening

The following skills are updated to mandate native `gitmap agent` commands and explicitly prohibit Python SQLite one-liners:

1. **`gitmap` (`.agents/skills/gitmap/skill.md` & `.cursor/skills/gitmap/skill.md`):**
   - Section 10 ("Native Fleet Synchronization & SQLite Agent Task Engine") expanded with complete subtask inspection commands (`gitmap agent task status`, `gitmap agent subtask ls`, `gitmap agent task ls`, `gitmap agent task pending`, `gitmap agent task completed`).
   - Section 10 updated with explicit ban warning against `python -c "import sqlite3..."`.
2. **`execute-parent-task-with-n-steps-v6` (`.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` & `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`):**
   - Step 0 Preflight & Section 7 Worker Waves updated to reference `gitmap agent` commands exclusively.
   - Ban on raw Python SQLite queries.
3. **`ai-verification` (`.agents/skills/ai-verification/skill.md`):**
   - Retrospective ledger initialization updated to `gitmap agent task init`.

---

## 6. Non-Destructive Fleet Synchronization

The canonical assets are mirrored across all 43 connected repositories using `03-ai-scripts/38-sync-prompts-skills-scripts.py`.
- **Zero Deletions Rule:** In accordance with the non-destructive sync invariant, only canonical prompts, skills, shared specs, and additive scripts are updated. No repo-owned custom files are deleted (`-0`).
- **Script Removal from Sync Manifest:** References to `46-agent-sqlite-task-manager.py` are removed from `03-ai-scripts/38-sync-prompts-skills-scripts.py`.
- **Upstream Push Invariant:** Every target repository is staged, committed via GitMap (`gitmap cpf`), and pushed to GitHub.
