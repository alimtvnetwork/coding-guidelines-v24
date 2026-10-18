# Master Execution Plan: Coding Guidelines V6 Upgrade, Data Schema Verification & GitMap Heatmap Integration

> **Plan Identifier:** `.ai-memory/plans/cg-execute-folder-synchronization-and-heatmap-integration.md`  
> **Status:** COMPLETED  
> **Specification:** `02-spec/21-app/cg-execute-folder-synchronization-and-heatmap-integration/01-architecture-spec.md`  
> **Total Steps Budget:** N = 400 (Phase 1: 100, Phase 2: 150, Phase 3: 150)  
> **Concurrency Capacity:** A = 3, H = 2, Capacity = 6 (Dynamic Immediate Respawn, 100% Duty Cycle)  
> **Local Completion Boundary:** STRICT LOCAL ONLY (Total ban on fleet multi-repo sync)

---

## User Request (Verbatim)

```text
You are Spec Author 01 for task cg-execute-folder-synchronization-and-heatmap-integration. You have no prior chat context; this brief is your complete specification.

### Boundaries & Crash Prevention:
- You are an authoring subagent (TypeName: 'self').
- Edit and create ONLY your assigned disjoint files:
  1. 02-spec/21-app/cg-execute-folder-synchronization-and-heatmap-integration/01-architecture-spec.md
  2. .ai-memory/plans/cg-execute-folder-synchronization-and-heatmap-integration.md
- TOTAL BAN ON RAW SEARCH TOOLS: NEVER run rg, ripgrep, grep, git grep, Select-String, Get-ChildItem -Recurse, or findstr.
- TOTAL BAN ON GIT COMMANDS: NEVER run git add, git commit, git push, git status, git diff, git checkout.
- Strict Relative Git Paths: All paths cited in files MUST be strictly relative to the git repository root. TOTAL BAN on absolute filesystem paths (C:\..., /home/...) and file:/// URIs.
- Strict Lowercase Filenames: All files, directory paths, and references must be strictly lowercase.
- Boolean Standards: Use positive booleans (is, has), never compare explicit == true.

### Required Spec Content for 01-architecture-spec.md:
1. Executive Summary: Upgrading the entire Coding Guidelines execution suite (01-prompts/15-cg-execute/ and companion cg-* skills) to V6 architecture with Data Schema Verification, GitMap Heatmap telemetry, and RC Agent database location transparency.
2. The 3-Tier Database Architecture & Relational Topology:
   - Tier 1: Machine-Local Agent Root Database (~/.gitmap/agent-root.db). Global catalog of active workspaces, sessions, and contention leases.
   - Tier 2: Repo-Cache Companion Database (repo-cache/.heatmap/<ai-slug>/<repo-slug>/heatmap.db and root-heatmap.db). Persistent, commit-tracked companion storage for cognitive decisions and subtask action logs without dirtying git worktrees. Synchronized via gitmap heatmap rc sync.
   - Tier 3: In-Flight Local Run Database (.ai-memory/temp-agents/<slug>/agent-task.db and .ai-memory/heatmap.db). Microsecond (<1ms) compiled Go SQLite with WAL mode (PRAGMA journal_mode=WAL; PRAGMA busy_timeout=5000;).
3. Data Schema Verification Protocol:
   - Step 0 preflight schema inspection: gitmap agent task status --ddl / --json.
   - Positive boolean constraints (Is*, Has*, non-null defaults).
   - Singular PascalCase table naming (User, Transaction), PascalCase columns, {TableName}Id keys.
   - Physical FK constraints within single SQLite databases; relational repository enforcement across Split-DB boundaries.
4. GitMap Heatmap & RC Agent Command Parity:
   - Command suites: gitmap heatmap log, gitmap heatmap discuss, gitmap heatmap live, gitmap heatmap rc query, gitmap agent task status.
   - Reporting contract: When returning heatmap data, agents must display explicit database filepaths and executable inspection commands.
5. 400-Step Budget & Concurrency Invariants:
   - N = 400 (Phase 1 = 100, Phase 2 = 150, Phase 3 = 150).
   - A = 3, H = 2, Capacity = 6. Dynamic immediate respawn (100% duty cycle).
   - Strict Non-Negotiable: DO NOT sync with other repos (local completion and verification first).

### Required Content for .ai-memory/plans/cg-execute-folder-synchronization-and-heatmap-integration.md:
- Complete Master Execution Plan following canonical format:
  - Header parameters and task metadata.
  - User Request (Verbatim).
  - 4 Discrete Subtasks linking to .ai-memory/plans/subtasks/cg-execute-folder-synchronization-and-heatmap-integration/01-data-schema-and-v6-parity.md through 04-verification-and-commit.md.
  - Non-Negotiable Boundaries & Safety Invariants (including explicit local-only rule: No Fleet Sync).
```

---

## Executive Summary & Objectives

This master execution plan governs the repository-wide upgrade of the Coding Guidelines execution suite (`01-prompts/15-cg-execute/` and companion `cg-*` skills under `.agents/skills/` and `.cursor/skills/`) to V6 architectural parity:
1. **Data Schema Verification & V6 Concurrency Parity:** Standardizes prompts and skills on N = 400, A = 3, H = 2, Capacity = 6, and mandates Step 0 schema verification (`gitmap agent task status --ddl` / `--json`, singular PascalCase tables, `{TableName}Id` keys, positive boolean constraints `Is*`/`Has*`).
2. **GitMap Heatmap Telemetry & Cognitive Action Logging:** Embeds pre-edit file contention inspection (`gitmap heatmap discuss`), granular cognitive decision logging (`gitmap heatmap log`, `gitmap agent log`), and live telemetry streaming (`gitmap heatmap live`).
3. **RC Agent Database Location Transparency & Split-DB Queries:** Establishes the 3-tier database architecture (Tier 1 `~/.gitmap/agent-root.db`, Tier 2 `repo-cache/.heatmap/...`, Tier 3 `.ai-memory/temp-agents/...`), explicit database path reporting, and zero-lock read-only querying (`gitmap heatmap rc query`).
4. **Targeted Verification, Path Hygiene & Atomic Local Commit:** Executes thorough file-scoped verification, validates relative path hygiene, confirms lowercase naming, enforces strict local completion without fleet sync, and seals the release with a single atomic GitMap commit.

---

## Subtask Decomposition

| Task ID | Title | Owned Subtask File | Target Subsystems | Owner Role | Status |
|---|---|---|---|---|---|
| `Task-01` | Data Schema Verification & V6 Concurrency Parity | `.ai-memory/plans/subtasks/cg-execute-folder-synchronization-and-heatmap-integration/01-data-schema-and-v6-parity.md` | `01-prompts/15-cg-execute/`, `.agents/skills/cg-*`, `.cursor/skills/cg-*` | `V6PromptWorker` | DONE |
| `Task-02` | CG Execute Prompts Modernization across 01-prompts and v4 | `.ai-memory/plans/subtasks/cg-execute-folder-synchronization-and-heatmap-integration/02-cg-prompts-modernization.md` | `01-prompts/15-cg-execute/`, `01-prompts/v4/15-cg-execute/` | `Worker 02` | DONE |
| `Task-03` | Companion CG Skills Modernization & Parity | `.ai-memory/plans/subtasks/cg-execute-folder-synchronization-and-heatmap-integration/03-skills-and-rc-agent.md` | `.agents/skills/cg-*`, `.cursor/skills/cg-*`, `app-db-architecture` | `Worker 03` | DONE |
| `Task-04` | Verification, Path Hygiene & Local Commit | `.ai-memory/plans/subtasks/cg-execute-folder-synchronization-and-heatmap-integration/04-verification-and-commit.md` | Workspace Root, `.ai-memory/`, GitMap Commit Engine | `LeadOrchestrator` | DONE |

---

## Subtask Execution Details

### Task-01: Data Schema Verification & V6 Concurrency Parity
- **Owned Subtask File:** `.ai-memory/plans/subtasks/cg-execute-folder-synchronization-and-heatmap-integration/01-data-schema-and-v6-parity.md`
- **Scope:** Upgrade top-header execution parameters across all 37 files in `01-prompts/15-cg-execute/` and companion `cg-*` skills to `N = 400`, `A = 3`, `H = 2`, `Capacity = 6` with dynamic immediate respawn.
- **Key Invariants:**
  - Embed Step 0 preflight DDL inspection: `gitmap agent task status --ddl` and `gitmap agent task status --json`.
  - Enforce singular PascalCase table naming (`ParentTask`, `Subtask`, `AgentActionLog`, `FileLease`).
  - Enforce `{TableName}Id` keys and positive boolean columns (`IsActive`, `HasCompleted`, `IsBlocked`).
  - Total ban on raw Python SQLite queries (`python -c "import sqlite3..."`); all task tracking via compiled Go binary `gitmap agent`.

### Task-02: CG Execute Prompts Modernization across 01-prompts and v4
- **Owned Subtask File:** `.ai-memory/plans/subtasks/cg-execute-folder-synchronization-and-heatmap-integration/02-cg-prompts-modernization.md`
- **Scope:** Modernize all 37 prompts in `01-prompts/15-cg-execute/` and `01-prompts/v4/15-cg-execute/` with full V6 multi-agent orchestration, Data Schema Verification, and GitMap Heatmap telemetry suite.
- **Key Invariants:**
  - Mandate pre-edit blast radius and write-collision check: `gitmap heatmap discuss [--window 15m] [--threshold 2]`.
  - Mandate in-flight cognitive decision logging before reading or modifying files:
    - Exploration / Read: `gitmap heatmap log -t <slug> -a "READ" -f "<path>" -r "Reasoning: <rationale>" -s "EXPLORATION"`
    - Modification / Edit: `gitmap heatmap log -t <slug> -a "EDIT" -f "<path>" -r "Decision: <rationale>" -s "EXECUTION"`
  - Support real-time continuous activity stream: `gitmap heatmap live` and hierarchical view `gitmap heatmap tree`.

### Task-03: Companion CG Skills Modernization & Parity
- **Owned Subtask File:** `.ai-memory/plans/subtasks/cg-execute-folder-synchronization-and-heatmap-integration/03-skills-and-rc-agent.md`
- **Scope:** Standardize all 18 companion skills in `.agents/skills/` and `.cursor/skills/` with PascalCase column naming, 3-tier database architecture documentation, and location transparency reporting requirements.
- **Key Invariants:**
  - Formalize Tier 1 (`~/.gitmap/agent-root.db`), Tier 2 (`repo-cache/.heatmap/<ai-slug>/<repo-slug>/heatmap.db`), and Tier 3 (`.ai-memory/temp-agents/<slug>/agent-task.db`, `.ai-memory/heatmap.db`).
  - Mandate Reporting Contract: Telemetry outputs must state explicit SQLite database paths and executable read-only inspection commands (`gitmap heatmap rc query "<sql>"`).
  - Include multi-AI namespace isolation and companion database reconciliation via `gitmap heatmap rc sync` and `gitmap heatmap rc commit`.

### Task-04: Verification, Path Hygiene & Local Commit
- **Owned Subtask File:** `.ai-memory/plans/subtasks/cg-execute-folder-synchronization-and-heatmap-integration/04-verification-and-commit.md`
- **Scope:** Perform final quality audits, relative path verification, and local atomic commit.
- **Key Invariants:**
  - Execute file-scoped linting: `python 03-ai-scripts/11-relative-path-linter.py` across `01-prompts/15-cg-execute/` and skills.
  - Verify zero absolute filesystem paths (`D:\...`, `C:\...`) and zero `file:///` URIs.
  - Enforce strict local-only completion invariant: **TOTAL BAN ON FLEET SYNC** (`gitmap sync` is strictly excluded).
  - Stage all files and create atomic GitMap commit: `gitmap cpf "cg-execute - v6 data schema verification and heatmap telemetry integration"` (using hyphen format without colons in argument).

---

## Non-Negotiable Invariants & Safety Guardrails

1. **Strict Local-Only Execution (TOTAL BAN ON FLEET SYNC):**
   - All changes must be developed, tested, and verified strictly within the local `coding-guidelines` repository.
   - NEVER trigger multi-repository synchronization (`gitmap sync`, `03-ai-scripts/38-sync-prompts-skills-scripts.py`). Multi-repo deployment is deferred until full local validation is completed.
2. **GitMap High-Speed Search Primacy (TOTAL BAN ON RAW GREP):**
   - TOTAL BAN on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem -Recurse`, and `findstr`.
   - All code searches and symbol lookups must use GitMap commands: `gitmap aum search`, `gitmap search`, `gitmap find`, and `gitmap cat`.
3. **Subagent Git Command Ban (Lock Collision Prevention):**
   - Subagents NEVER execute git commands (`git add`, `git commit`, `git push`, `git status`, `git diff`, `git checkout`).
   - Git operations are exclusively executed by the Lead Orchestrator upon task completion to prevent `.git/index.lock` collisions.
4. **Strict Relative Git Paths:**
   - TOTAL BAN on absolute filesystem paths (`D:\...`, `C:\...`, `/home/...`) and `file:///` URIs.
   - All markdown links, file paths, citations, and subtask paths must be strictly relative to the repository root.
5. **Strict Lowercase Naming:**
   - All generated or modified files, directories, and path references must strictly adhere to lowercase naming with hyphens.
6. **Positive Boolean Principles:**
   - Positive booleans only (`is`, `has`). NEVER evaluate against explicit `== true`. NEVER combine positive and negative checks in the same condition.
7. **Zero Builds or Tests:**
   - NEVER run full builds or tests (`go build`, `npm run build`, `go test`, `pytest`). Rely exclusively on fast file-scoped linters and schema inspection tools.
