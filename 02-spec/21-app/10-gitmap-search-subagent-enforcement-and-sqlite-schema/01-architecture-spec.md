# Architecture Specification: GitMap Search Primacy Across All Agent Tiers & SQLite Task Schema Formalization

> **Spec ID:** `02-spec/21-app/10-gitmap-search-subagent-enforcement-and-sqlite-schema/01-architecture-spec.md`  
> **Status:** APPROVED  
> **Author:** Lead Systems Architect & Spec Subagent 01  
> **Target Subsystems:** `01-prompts/14-execute/`, `01-prompts/22-letterly/`, `01-prompts/23-cursor-prompts/`, `.agents/skills/`, `.cursor/skills/`, `03-ai-scripts/46-agent-sqlite-task-manager.py`, `AGENTS.md`  
> **Database:** `.ai-memory/temp-agents/<nn>-<slug>/agent-task.db`  

---

## 1. Problem Statement & Root Cause Analysis

### 1.1. Empirical Violation Analysis: Ripgrep (`rg`) Leakage in Subagents
During recent multi-agent task execution runs, telemetry and visual session logs captured significant guideline violations where background subagents bypassed standard repository search tools.

As documented in `assets/screenshots/gitmap-search-subagent-enforcement-01.png`, the subagent `Research 01: OS Window Title & Button Alignment Specialist` executing within the Antigravity Manager (`AGM`) workspace repeatedly executed raw shell search commands:

```text
Ran rg -n "settings\.json" src-tauri/
Ran rg -n "launch_instance_with_clean_env" src-tauri/
Ran rg -n "clean_env|clean_appimage_env" src-tauri/
Ran rg -i "AppUserModelID" src-tauri/
Ran rg -n "fn rename_instance" src-tauri/
Ran rg -i "shortcut|desktop" src-tauri/src/modules/
Ran rg -i "shortcut" src-tauri/src/bin/agm.rs src-tauri/src/modules/cli.rs
Ran rg -n "accentBar" src/components/
```

### 1.2. Root Causes of Subagent Grep Execution
A forensic investigation into why the subagent invoked `rg` revealed four distinct architectural gaps:

1. **LLM Pre-Training Heuristics:**  
   Standard language model base completions possess an inherent bias toward default UNIX/GNU utilities (`rg`, `ripgrep`, `grep`, `find`, or PowerShell equivalents `Select-String`, `findstr`) whenever tasked with locating strings or symbols across a codebase.
2. **Missing Boundary Clauses in Subagent Dispatch Prompts:**  
   In the parent orchestrator prompt (`01-prompts/14-execute/13-execute-parent-task-with-n-steps-v6.md`) and companion skills (`.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`), the subagent dispatch template instructions did not explicitly forbid `rg` or mandate `gitmap` search commands inside the subagent instruction payload.
3. **Upstream Voice Dictation (Letterly) Formatter Blind Spots:**  
   Letterly prompt formatters (`01-prompts/22-letterly/03-execute-n-steps-letterly.md` and cursor variants) focused on task structure and relative paths, but lacked an explicit directive for the AI to enforce `gitmap` search tooling across subagents.
4. **Absence of Meta-Repository Ban in `AGENTS.md`:**  
   While `AGENTS.md` strictly governs booleans, zero-storage CI/CD, relative paths, and secrets, it lacked an explicit clause prohibiting ad-hoc shell grep commands in favor of GitMap search tooling.

### 1.3. Rationale for GitMap Search Primacy & Strict Shell Grep Ban
Executing raw `rg`, `grep`, or `Select-String` commands introduces severe operational liabilities:
- **No Git Awareness:** Raw `rg` commands executed without proper flags frequently traverse build caches, binary directories (`target/`, `node_modules/`, `.git/`), and temporary test databases.
- **Process Spawning Overhead:** In high-concurrency environments (multiple parallel subagents), spawning dozens of separate shell grep sub-processes creates OS resource contention and thread stalls.
- **Inconsistent Token Output:** Unbounded ripgrep outputs flood subagent context windows with thousands of irrelevant lines, causing context exhaustion and hallucinated edits.
- **GitMap Superiority:** GitMap utilities (`gitmap aum search`, `gitmap find`, `gitmap search`, `gitmap cat`) utilize pre-warmed repository indexes, strictly honor `.gitignore`, normalize cross-platform Windows/Linux path separators, output clean AI-friendly token formats, and maintain zero external dependency footprints.

Therefore, a **TOTAL BAN** must be established on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem -Recurse`, and `findstr` across all agent prompts, skills, and meta-guidelines.

---

## 2. System Architecture: Three-Tier Agent Hierarchy & Search Primacy

To ensure complete compliance across all agent lifecycles, search operations are strictly bound across three operational tiers:

```
                      +---------------------------------------+
                      |       Tier 1: Lead Orchestrator       |
                      |  - Manages ParentTask in SQLite       |
                      |  - Generates Spec & Disjoint Subtasks |
                      |  - Coordinates A=2, H=2 Subagents     |
                      |  - Search Primacy: gitmap only        |
                      +-------------------+-------------------+
                                          |
                     +--------------------+--------------------+
                     |                                         |
                     v                                         v
   +------------------------------------+   +------------------------------------+
   |   Tier 2: Research & Discovery     |   |   Tier 3: Execution Worker Agents   |
   | - Codebase structure exploration   |   | - Claims Subtask from SQLite DB    |
   | - Symbol and reference location    |   | - Strictly disjoint owned files    |
   | - Bounded context gathering        |   | - Logs in-flight actions           |
   | - Search Primacy: gitmap only      |   | - In-file edits & targeted tests   |
   | - FORBIDDEN: rg / grep / findstr   |   | - Search Primacy: gitmap only      |
   +------------------------------------+   +------------------------------------+
```

### 2.1. Operational Rules by Tier

#### Tier 1: Lead Orchestrator
- Responsible for parsing user requests, orchestrating planning, initializing the SQLite run ledger via `03-ai-scripts/46-agent-sqlite-task-manager.py`, and assigning strictly disjoint owned files.
- When generating prompts for subagents, the Lead Orchestrator **must inject the Mandatory Search Primacy Block** into every subagent payload.

#### Tier 2: Research & Discovery Subagents
- Conduct architectural discovery, reference tracing, and symbol lookup.
- Must execute all search tasks exclusively through `gitmap` CLI commands or IDE read tools (`view_file`).
- Must never invoke raw CLI search binaries (`rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem -Recurse`, `findstr`).

#### Tier 3: Execution Worker Subagents
- Operate exclusively within their assigned, disjoint list of owned files (`OwnedFilesJson`).
- For any cross-file reference lookups or dependency tracing, workers must invoke `gitmap` search helpers.

### 2.2. Tool Substitution Matrix

| Banned Command Pattern | Permitted Canonical Replacement | Description / Scope |
| :--- | :--- | :--- |
| `rg "<pattern>"` | `gitmap search "<pattern>"` | High-speed indexed full-text search across tracked files |
| `rg -n "<pattern>" <path>` | `gitmap aum search "<pattern>" --path <path>` | Scoped path search with line numbers and token bounds |
| `rg -i "<pattern>"` | `gitmap search "<pattern>"` | Case-insensitive normalized search |
| `grep -rn "<pattern>" .` | `gitmap search "<pattern>"` | Full-repository recursive pattern match |
| `git grep "<pattern>"` | `gitmap search "<pattern>"` | Git index pattern match |
| `Select-String -Pattern "<p>"` | `gitmap search "<p>"` | Cross-platform PowerShell search replacement |
| `Get-ChildItem -Recurse` | `gitmap find "<glob>"` | File glob and file discovery replacement |
| `find . -name "<glob>"` | `gitmap find "<glob>"` | File path discovery across repository tree |
| `cat <file>` | `gitmap cat <file>` or `view_file` | Deterministic file viewing with bounds |

### 2.3. Mandatory Subagent Prompt Injection Template
All parent task execution prompts and subagent invocations must embed the following non-negotiable instruction:

```markdown
### MANDATORY SEARCH PRIMACY & TOTAL BAN ON RAW GREP:
You MUST execute all codebase searches using GitMap:
  - Full-text search: `gitmap search "<term>"` or `gitmap aum search "<term>"`
  - Path-scoped search: `gitmap aum search "<term>" --path <relative-directory>`
  - File discovery: `gitmap find "<glob>"`
  - File reading: `gitmap cat <file>` or native `view_file` tool

TOTAL BAN (AUTO-REJECT FAILURE):
NEVER run `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem -Recurse`, or `findstr`.
Invoking any banned command violates core repository rules.
```

---

## 3. SQLite Task Manager Architecture & Schema Formalization

To provide ACID-compliant coordination for multi-agent execution while completely avoiding Windows file-locking collisions, `03-ai-scripts/46-agent-sqlite-task-manager.py` maintains an isolated SQLite database at:
`.ai-memory/temp-agents/<nn>-<slug>/agent-task.db`.

### 3.1. Database Engine & Concurrency Primitives
- **Journal Mode:** Write-Ahead Logging (`PRAGMA journal_mode = WAL;`). Allows multiple concurrent subagent readers and single serialized writers without blocking readers.
- **Busy Timeout:** 5,000 milliseconds (`PRAGMA busy_timeout = 5000;`). Prevents immediate `sqlite3.OperationalError: database is locked` during concurrent subagent claiming.
- **Synchronous Mode:** Normal (`PRAGMA synchronous = NORMAL;`). Provides high I/O throughput with crash resistance.

### 3.2. Data Definition Language (DDL) & Schema Tables

```sql
CREATE TABLE IF NOT EXISTS ParentTask (
    ParentTaskId INTEGER PRIMARY KEY AUTOINCREMENT,
    TaskName TEXT NOT NULL,
    TaskSlug TEXT NOT NULL,
    RunDirectory TEXT NOT NULL,
    Status TEXT NOT NULL DEFAULT 'ACTIVE',
    IsActive INTEGER NOT NULL DEFAULT 1,
    HasCompleted INTEGER NOT NULL DEFAULT 0,
    TotalStepsBudget INTEGER NOT NULL DEFAULT 300,
    CurrentStep INTEGER NOT NULL DEFAULT 1,
    Notes TEXT NULL,
    CreatedAt TEXT NOT NULL,
    UpdatedAt TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Subtask (
    SubtaskId INTEGER PRIMARY KEY AUTOINCREMENT,
    ParentTaskId INTEGER NOT NULL,
    TaskCode TEXT NOT NULL,
    Title TEXT NOT NULL,
    AssignedAgentRole TEXT NULL,
    OwnedFilesJson TEXT NOT NULL DEFAULT '[]',
    Status TEXT NOT NULL DEFAULT 'PENDING',
    IsBlocked INTEGER NOT NULL DEFAULT 0,
    HasCompleted INTEGER NOT NULL DEFAULT 0,
    Evidence TEXT NULL,
    CreatedAt TEXT NOT NULL,
    UpdatedAt TEXT NOT NULL,
    FOREIGN KEY (ParentTaskId) REFERENCES ParentTask(ParentTaskId)
);

CREATE TABLE IF NOT EXISTS AgentActionLog (
    ActionLogId INTEGER PRIMARY KEY AUTOINCREMENT,
    SubtaskId INTEGER NOT NULL,
    AgentRole TEXT NOT NULL,
    ActionType TEXT NOT NULL,
    TargetFile TEXT NULL,
    ActionDetails TEXT NOT NULL,
    Status TEXT NOT NULL DEFAULT 'IN_PROGRESS',
    CreatedAt TEXT NOT NULL,
    FOREIGN KEY (SubtaskId) REFERENCES Subtask(SubtaskId)
);

CREATE INDEX IF NOT EXISTS idx_subtask_status ON Subtask(Status);
CREATE INDEX IF NOT EXISTS idx_action_subtask ON AgentActionLog(SubtaskId);
```

### 3.3. Positive Boolean Storage & Naming Principles
In accordance with global coding guidelines:
- All boolean columns must be positive nouns/predicates starting with `Is` or `Has`.
  - `IsActive`: `1` when the parent run is currently executing; `0` when closed or suspended.
  - `HasCompleted`: `1` when the task or subtask is verified and finished; `0` otherwise.
  - `IsBlocked`: `1` when a subtask cannot proceed due to external failure; `0` otherwise.
- Banned prefixes: `Can`, `Should`, `Was`, `Not`, or negative naming (`IsNotActive`).
- Storage type: SQLite `INTEGER` (0 or 1).
- Python evaluation: Evaluated implicitly (`if task["isActive"]: ...`, never `if task["isActive"] == True:`).

### 3.4. JSON Serialization & Strict Relative Path Validation
The `OwnedFilesJson` column in `Subtask` stores an array of relative file paths assigned to a specific subagent.
To prevent guideline violations:
1. Every path in `OwnedFilesJson` must be a normalized, forward-slash relative path from the repository root (e.g., `03-ai-scripts/46-agent-sqlite-task-manager.py`).
2. Absolute filesystem paths (e.g. drive letters or root slashes) and URI schemes (such as `file:` URIs) are strictly prohibited and must be rejected by `cmd_add_subtasks` input validation.
3. Subtasks must possess mutually disjoint owned file lists to eliminate concurrent write collisions.

---

## 4. SQLite Task Manager `schema` Command Specification

To allow agents, CI/CD runners, and developers to inspect the database schema without external SQLite tools, `03-ai-scripts/46-agent-sqlite-task-manager.py` must expose a dedicated `schema` command.

### 4.1. CLI Invocation Syntax

```bash
# Display formatted human-readable summary table
python 03-ai-scripts/46-agent-sqlite-task-manager.py schema [--db <path-to-db>]

# Display complete SQLite DDL statements
python 03-ai-scripts/46-agent-sqlite-task-manager.py schema [--db <path-to-db>] --ddl

# Display structured machine-readable JSON schema metadata
python 03-ai-scripts/46-agent-sqlite-task-manager.py schema [--db <path-to-db>] --json
```

### 4.2. Flag Specifications
- `--db <path>` *(optional)*: Path to an existing SQLite database. If omitted, the script inspects the most recent task database under `.ai-memory/temp-agents/` or initializes an in-memory schema verification connection.
- `--ddl` *(optional)*: Outputs the complete, formatted SQL DDL statements for all tables and indexes.
- `--json` *(optional)*: Outputs a structured JSON object containing table definitions, column specifications, affinities, nullability, default values, and foreign keys.

### 4.3. Machine-Readable JSON Schema Output Contract

```json
{
  "databaseEngine": "SQLite",
  "journalMode": "WAL",
  "busyTimeoutMs": 5000,
  "tables": [
    {
      "tableName": "ParentTask",
      "columns": [
        {"name": "ParentTaskId", "type": "INTEGER", "isPrimaryKey": true, "isNotNull": true, "defaultValue": null},
        {"name": "TaskName", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null},
        {"name": "TaskSlug", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null},
        {"name": "RunDirectory", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null},
        {"name": "Status", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": "'ACTIVE'"},
        {"name": "IsActive", "type": "INTEGER", "isPrimaryKey": false, "isNotNull": true, "defaultValue": "1"},
        {"name": "HasCompleted", "type": "INTEGER", "isPrimaryKey": false, "isNotNull": true, "defaultValue": "0"},
        {"name": "TotalStepsBudget", "type": "INTEGER", "isPrimaryKey": false, "isNotNull": true, "defaultValue": "300"},
        {"name": "CurrentStep", "type": "INTEGER", "isPrimaryKey": false, "isNotNull": true, "defaultValue": "1"},
        {"name": "Notes", "type": "TEXT", "isPrimaryKey": false, "isNotNull": false, "defaultValue": null},
        {"name": "CreatedAt", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null},
        {"name": "UpdatedAt", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null}
      ],
      "indexes": []
    },
    {
      "tableName": "Subtask",
      "columns": [
        {"name": "SubtaskId", "type": "INTEGER", "isPrimaryKey": true, "isNotNull": true, "defaultValue": null},
        {"name": "ParentTaskId", "type": "INTEGER", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null},
        {"name": "TaskCode", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null},
        {"name": "Title", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null},
        {"name": "AssignedAgentRole", "type": "TEXT", "isPrimaryKey": false, "isNotNull": false, "defaultValue": null},
        {"name": "OwnedFilesJson", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": "'[]'"},
        {"name": "Status", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": "'PENDING'"},
        {"name": "IsBlocked", "type": "INTEGER", "isPrimaryKey": false, "isNotNull": true, "defaultValue": "0"},
        {"name": "HasCompleted", "type": "INTEGER", "isPrimaryKey": false, "isNotNull": true, "defaultValue": "0"},
        {"name": "Evidence", "type": "TEXT", "isPrimaryKey": false, "isNotNull": false, "defaultValue": null},
        {"name": "CreatedAt", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null},
        {"name": "UpdatedAt", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null}
      ],
      "indexes": ["idx_subtask_status"]
    },
    {
      "tableName": "AgentActionLog",
      "columns": [
        {"name": "ActionLogId", "type": "INTEGER", "isPrimaryKey": true, "isNotNull": true, "defaultValue": null},
        {"name": "SubtaskId", "type": "INTEGER", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null},
        {"name": "AgentRole", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null},
        {"name": "ActionType", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null},
        {"name": "TargetFile", "type": "TEXT", "isPrimaryKey": false, "isNotNull": false, "defaultValue": null},
        {"name": "ActionDetails", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null},
        {"name": "Status", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": "'IN_PROGRESS'"},
        {"name": "CreatedAt", "type": "TEXT", "isPrimaryKey": false, "isNotNull": true, "defaultValue": null}
      ],
      "indexes": ["idx_action_subtask"]
    }
  ]
}
```

---

## 5. Architectural Invariants & Verification Checklist

1. **GitMap Search Primacy:** Every agent prompt and skill must require `gitmap search`, `gitmap aum search`, and `gitmap find`.
2. **Total Ban on Grep Tooling:** Under no circumstances may an agent invoke `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem -Recurse`, or `findstr`.
3. **Strict Relative Git Paths:** All file references across SQLite records, plans, specs, and prompts must be strictly relative to the repository root.
4. **Positive Booleans:** Boolean fields in SQLite and Python must use positive naming (`IsActive`, `HasCompleted`, `IsBlocked`) and implicit condition evaluation.
5. **ACID Concurrency:** Multi-agent operations must coordinate via SQLite WAL mode micro-transactions without file-locking contention.
