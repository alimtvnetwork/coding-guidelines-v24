# GitMap Fleet Sync & Native Agent Task Engine — Architecture Specification

> **Spec Identifier:** `02-spec/21-app/gitmap-sync-and-agent-task-engine/01-architecture-spec.md`  
> **Status:** APPROVED & BINDING  
> **Version:** 1.0.0  
> **Target Subsystems:** `gitmap` CLI, `coding-guidelines` Prompts & Skills, Multi-Repo Fleet Sync

---

## 1. Executive Summary & Problem Statement

### 1.1 Context
In distributed multi-agent workflows and monorepo/meta-repo architectures, 43+ connected Git repositories must stay synchronized with canonical prompts (`01-prompts/`), IDE skills (`.agents/skills/`, `.cursor/skills/`), shared specifications (`02-spec/01-*` to `02-spec/20-*`), and additive AI automation scripts (`03-ai-scripts/`).

### 1.2 Identified Bottlenecks
1. **Python Script Overhead:** The current synchronization process relies on `03-ai-scripts/38-sync-prompts-skills-scripts.py`. While functional, running a Python interpreter with file I/O walks and sequential git subprocess spawning across 43 repositories introduces latency and requires Python runtime dependencies.
2. **Configuration Inflexibility:** Target repositories are hardcoded in the Python script. Users cannot easily pass an ad-hoc JSON manifest of project URLs, folder paths, or dynamic workspace modes without modifying Python code.
3. **Agent Task Manager Latency:** Agent state tracking and SQLite task coordination rely on `03-ai-scripts/46-agent-sqlite-task-manager.py`. Invoking `python` repeatedly for every claim, add, complete, and status check incurs Python VM bootstrap overhead (~200ms per invocation).
4. **Tool Disunity:** AI agents currently alternate between native `gitmap` CLI commands and multiple Python helper scripts. Bringing fleet synchronization and SQLite agent task tracking natively into `gitmap` unifies the autonomous toolchain into a single compiled Go binary.

---

## 2. Architectural Design: GitMap Native Fleet Sync (`gitmap sync`)

### 2.1 CLI Interface & Syntax
`gitmap sync` provides high-speed, parallel synchronization of canonical assets across repositories:

```bash
# 1. Sync all 43 default fleet repositories using built-in canonical registry
gitmap sync

# 2. Sync specific repositories by name or pattern
gitmap sync --repo cat-my --repo gitmap
gitmap sync -r "wp-*"

# 3. Supply custom projects via JSON file or JSON string
gitmap sync --projects "projects.json"
gitmap sync --projects '[{"name": "custom-app", "path": "d:/work/custom-app"}]'
gitmap sync --projects '[{"url": "git@github.com:org/repo.git", "folder": "repo"}]'

# 4. Preview changes safely (dry run)
gitmap sync --dry-run

# 5. Disable automated post-sync release tagging / push
gitmap sync --no-push
gitmap sync --no-release

# 6. Configure goroutine worker concurrency (default: 8)
gitmap sync --workers 12
```

### 2.2 JSON Input Schema
The `--projects` flag accepts either a path to a `.json` file or an inline JSON string:
```json
[
  {
    "name": "project-slug",
    "path": "d:/work/project-slug",
    "url": "git@github.com:alimtvnetwork/project-slug.git",
    "mode": "all"
  }
]
```
- `name`: Human-readable identifier.
- `path`: Local absolute or relative workspace folder path.
- `url`: Optional remote Git URL for clone/pull verification.
- `mode`: `"all"` (default), `"prompts-only"`, `"skills-only"`, or `"specs-only"`.

### 2.3 Strict Boundary Invariants (Enforced in Go)
The Go engine MUST enforce all 5 non-negotiable synchronization boundaries:
1. **Spec 21-25 Exclusion (TOTAL BAN):** NEVER copy, sync, or mutate `02-spec/21-*` through `02-spec/25-*` (private application specs, issues, databases).
2. **Additive-Only AI Scripts:** Copy new scripts into `03-ai-scripts/` and `.agents/scripts/`. NEVER overwrite existing scripts modified by the target repository.
3. **Bump Script Protection (IMMUTABLE):** NEVER overwrite version bump scripts (`bump*`, `bump-version.mjs`, `bump_versions.py`).
4. **Memory & Plans Protection (TOTAL ISOLATION):** NEVER modify, sync, or mirror `.ai-memory/memory/` or `.ai-memory/plans/`.
5. **Zero Secrets Leakage:** NEVER synchronize `.env` files, tokens, or credentials.

### 2.4 Pre-Change Safety Protocol (Automatic in Go)
For each target repository before any file modifications:
1. `git pull origin <base_branch> --no-rebase`
2. Create and push safety backup branch: `git checkout -b backup/sync-<timestamp>` and `git push origin backup/sync-<timestamp>`
3. Return to base branch: `git checkout <base_branch>`
4. Apply asset diffs in-memory / on-disk using fast Go file hashing (SHA-256)
5. Atomically commit: `gitmap cpf "sync - update canonical prompts, skills, and shared specs"`
6. Minor version bump release ceremony and tag push (unless `--no-release`).

---

## 3. Architectural Design: GitMap Native SQLite Task Engine (`gitmap task`)

### 3.1 CLI Interface & Syntax
Native Go SQLite task manager eliminating Python startup latency:

```bash
# 1. Initialize task ledger in SQLite
gitmap task init --db <databasePath> --slug <slug> --title <title> --budget 300

# 2. Add decomposed subtasks from JSON manifest
gitmap task add --db <databasePath> --tasks-json '[{"code": "Task-01", "title": "...", "owned_files": ["..."], "agent_role": "..."}]'

# 3. Worker subagent claims next available task atomically
gitmap task claim --db <databasePath> --agent "Worker 01" [--subtask-id <id>]

# 4. Worker subagent completes task with evidence
gitmap task complete --db <databasePath> --subtask-id <id> --evidence "<summary>"

# 5. Check real-time progress and summary (terminal or JSON)
gitmap task status --db <databasePath>
gitmap task status --db <databasePath> --json

# 6. Inspect schema
gitmap task schema
```

### 3.2 Database Schema (Compatible with Existing SQLite Engine)
The Go implementation uses SQLite via modern Go drivers (`modernc.org/sqlite` or `mattn/go-sqlite3`):
- `task_metadata`: `task_slug`, `task_name`, `total_steps`, `completed_steps`, `status`, `created_at`, `updated_at`
- `subtasks`: `subtask_id`, `task_code`, `title`, `owned_files_json`, `assigned_agent_role`, `status`, `evidence`, `started_at`, `completed_at`
- `step_ledger`: `step_index`, `action_type`, `description`, `actor`, `recorded_at`

---

## 4. Acceptance Criteria & Verification Protocol

1. **AC-1 (Default Registry Sync):** `gitmap sync` seamlessly processes the default 43 connected repositories with pre-pull, safety backup branch creation, asset mirroring, and boundary enforcement.
2. **AC-2 (Custom Projects via JSON):** `gitmap sync --projects <json>` accepts external project configurations and synchronizes them without errors.
3. **AC-3 (Go Concurrency Performance):** Go-native `gitmap sync` completes synchronization across 43 repositories in <30 seconds (over 5x faster than Python).
4. **AC-4 (Native Task Operations):** `gitmap task init`, `add`, `claim`, `complete`, and `status` operate directly against SQLite with sub-10ms response times.
5. **AC-5 (Release & Upstream Push):** `gitmap` binary compiled, released with updated version tag, and pushed to GitHub.
6. **AC-6 (Prompt & Skill Alignment):** `01-prompts/` and `.agents/skills/` updated to instruct agents on using `gitmap sync` and `gitmap task`.
