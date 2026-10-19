# Specification 03: GitMap Cognitive Heatmap Operational Integration & Multi-Repo Fleet Sync

## 1. Executive Overview & Problem Statement

### 1.1 The Heatmap Telemetry Vacuum
In preceding execution cycles, autonomous AI agents completely ignored GitMap cognitive heatmap telemetry commands (`gitmap heatmap log`, `gitmap heatmap update`, `gitmap heatmap status`, `gitmap heatmap tree`). While the GitMap SQLite-backed heatmap engine was fully implemented, compiled, and capable of microsecond logging, agents operated in an unmonitored blind spot.

### 1.2 Root Cause Analysis
1. **Passive Reference Documentation vs. Operational Gates:** Heatmap commands were documented solely in reference tables and capability indices. Agents treat reference documentation as informational unless explicit execution gates demand invocation.
2. **Absence in Core Prompt Loops:** Neither `execute-parent-task-with-n-steps-v6` nor `AGENTS.md` contained mandatory checklist items or pre-action/post-action requirements for heatmap logging.
3. **Missing Copy-Paste Command Templates:** Prompts lacked explicit, drop-in syntax showing how to structure `--task`, `--action`, `--state`, `--file`, and `--expected` arguments.
4. **Disconnection from Repo-Cache Architecture:** Multi-AI telemetry routing via `gitmap heatmap rc` (and `gitmap rc heatmap`) was not specified as the default persistence layer for multi-agent swarms.

---

## 2. GitMap Cognitive Heatmap Command Suite

The GitMap heatmap architecture provides sub-millisecond SQLite WAL-mode recording of agent cognitive state, planned file actions, and post-action verification.

### 2.1 Primary Command Matrix

| Command | Action Type | Lifecycle Stage | Primary Purpose |
| :--- | :--- | :--- | :--- |
| `gitmap heatmap log` | Write / Stage | Pre-Action (Explore, Spec, Edit) | Log intent, target file, action verb, and expected outcome |
| `gitmap heatmap update` | Write / Close | Post-Action (Verify) | Update action record with actual outcome and boolean `met` flag |
| `gitmap heatmap status` | Read / Telemetry | Step 0, Gate Checks | Display action counts, open items, and state distribution |
| `gitmap heatmap tree` | Read / Graph | Verification / Audit | Render hierarchical causal execution tree of all agent actions |
| `gitmap heatmap tables` | Read / Schema | Diagnostics / Init | List SQLite tables and record counts in heatmap database |
| `gitmap heatmap query` | Read / SQL | Custom Telemetry | Run safe, read-only SQL queries against cognitive heatmap database |
| `gitmap heatmap rc <subcommand>` | Multi-AI / Storage | Cross-Agent Synchronization | Manage `repo-cache` isolated multi-AI heatmap telemetry |

### 2.2 Concrete Command Syntax & Parameter Standards

#### A. Pre-Action Logging: `gitmap heatmap log`
Every agent MUST invoke `gitmap heatmap log` BEFORE reading, analyzing, creating, editing, or deleting any file:

```bash
gitmap heatmap log \
  --task "<task-slug>" \
  --action "<READ|EDIT|CREATE|DELETE|ANALYZE>" \
  --file "<relative/git/path>" \
  --state "<EXPLORATION|SPECIFICATION|EXECUTION|VERIFICATION>" \
  --expected "<discrete goal description>"
```

**Parameters:**
- `--task`: Slug matching the parent task or subtask ledger (e.g. `gitmap-commands-agent-spawning-heatmap-sync`).
- `--action`: Strict uppercase action verb: `READ`, `EDIT`, `CREATE`, `DELETE`, `ANALYZE`.
- `--file`: Strict relative repository path from git root (e.g. `02-spec/21-app/...`).
- `--state`: Current swarm lifecycle phase: `EXPLORATION`, `SPECIFICATION`, `EXECUTION`, `VERIFICATION`.
- `--expected`: Concrete expected condition (e.g. "Author heatmap specification and execution gates").

#### B. Post-Action Verification: `gitmap heatmap update`
Immediately after performing the operation (file edit, test run, or spec generation), the agent MUST update the action log:

```bash
gitmap heatmap update \
  --id <action-id> \
  --outcome "<discrete outcome description>" \
  --met true
```

**Parameters:**
- `--id`: Numeric ID returned by `gitmap heatmap log` or discovered via `gitmap heatmap status`.
- `--outcome`: Summary of actual result achieved.
- `--met`: Boolean flag indicating whether the expectation was satisfied (`true` or `false`).

#### C. Status Overview: `gitmap heatmap status`
Inspect global or task-scoped telemetry volume and completion states:

```bash
# Task-scoped status summary
gitmap heatmap status --task "gitmap-commands-agent-spawning-heatmap-sync"

# Global repository status summary
gitmap heatmap status
```

#### D. Causal Hierarchy Tree: `gitmap heatmap tree`
Render the complete chronological tree of actions taken during the task:

```bash
gitmap heatmap tree --task "gitmap-commands-agent-spawning-heatmap-sync"
```

#### E. Table & Schema Verification: `gitmap heatmap tables`
Verify database connection, active tables, and row counts:

```bash
gitmap heatmap tables
```

#### F. Safe Read-Only SQL Inspection: `gitmap heatmap query`
Query cognitive actions directly via read-only SQL:

```bash
gitmap heatmap query "SELECT id, action, file_path, state, is_met FROM heatmap_actions WHERE task_slug = 'gitmap-commands-agent-spawning-heatmap-sync' ORDER BY id DESC LIMIT 10"
```

#### G. Multi-AI Telemetry Routing via Repo-Cache: `gitmap rc heatmap`
For multi-agent environments sharing `repo-cache` storage:

```bash
# Check repo-cache heatmap status
gitmap rc heatmap status --task "<task-slug>"

# View schema tables in repo-cache
gitmap rc heatmap tables

# Query repo-cache heatmap records
gitmap rc heatmap query "SELECT * FROM rc_heatmap_actions ORDER BY timestamp DESC LIMIT 5"

# Sync local telemetry to repo-cache
gitmap rc heatmap sync
```

---

## 3. Mandatory Heatmap Execution Gates

To eliminate telemetry omissions, four strict operational gates are integrated into the autonomous execution lifecycle:

```
+-----------------------------------------------------------------------------------+
| GATE 0: Pre-Flight & Discovery (Step 0)                                           |
| -> gitmap heatmap status                                                          |
| -> gitmap heatmap tables                                                          |
+-----------------------------------------------------------------------------------+
                                         |
                                         v
+-----------------------------------------------------------------------------------+
| GATE 1: Pre-Action Intent Logging (Prior to ANY Read, Edit, Create, Delete)       |
| -> gitmap heatmap log --task <slug> --action <VERB> --file <path> --expected "..."|
+-----------------------------------------------------------------------------------+
                                         |
                                         v
                                [ Execute File Action ]
                                         |
                                         v
+-----------------------------------------------------------------------------------+
| GATE 2: Post-Action Outcome Closure (Immediately Following File Action)           |
| -> gitmap heatmap update --id <id> --outcome "..." --met true                     |
+-----------------------------------------------------------------------------------+
                                         |
                                         v
+-----------------------------------------------------------------------------------+
| GATE 3: Retrospective Audit & Verification (Phase 3)                              |
| -> gitmap heatmap tree --task <slug>                                              |
| -> gitmap heatmap status --task <slug>                                            |
| Verify 100% action outcome closure prior to atomic commit                         |
+-----------------------------------------------------------------------------------+
```

### 3.1 Gate 0: Pre-Flight Initialization
- Executed during Step 0 before any planning or file reading.
- Verifies that the SQLite database engine is responsive and prints baseline row counts.
- Command: `gitmap heatmap status` and `gitmap heatmap tables`.

### 3.2 Gate 1: Pre-Action Declaration
- Executed before inspecting a file, modifying code, or authoring a document.
- Declares intent, target relative path, and expected outcome.
- Captures the returned `--id` for subsequent closure.

### 3.3 Gate 2: Post-Action Closure
- Executed immediately after tool execution.
- Records actual outcome, errors encountered, and boolean satisfaction flag.
- Prevents hanging or unresolved action states.

### 3.4 Gate 3: Phase 3 Telemetry Verification
- Executed at the conclusion of all subtasks before git commit.
- Inspects the causal tree with `gitmap heatmap tree --task <slug>` to verify zero unclosed actions.
- Outputs status report into `$GITHUB_STEP_SUMMARY` or chat transcript.

---

## 4. Multi-Repo Fleet Synchronization Protocol

Once canonical assets (prompts, skills, specs, scripts) are verified in `coding-guidelines`, they must be synchronized across all 43 connected repositories in the workspace.

### 4.1 Canonical Fleet Target Repositories

The canonical fleet comprises 43 repositories managed by `03-ai-scripts/38-sync-prompts-skills-scripts.py`:

1. `02-prompts/ai-empathy-prompt-tuner`
2. `alim-cv`
3. `alim-karim-profile`
4. `aukgit/alim.karim.profile`
5. `antigravity-manager`
6. `cat-my`
7. `03-aukgo/core`
8. `digital-name-card`
9. `presentations-repos/flat-slide-show`
10. `gitlogger-new`
11. `gitmap`
12. `presentations-repos/global-ppt-v1`
13. `presentations-repos/hiltrax`
14. `icon-coding-guidelines`
15. `img-pdf`
16. `presentations-repos/ki-health-ppt`
17. `aukgit/kubernetes-training`
18. `lara-licensing`
19. `lara-publishing`
20. `laravel-automation`
21. `letsmarknow-ui`
22. `letsmarknow`
23. `macro-ahk`
24. `presentations-repos/maid-app-spec-presentation`
25. `movie-cli`
26. `03-aukgo/pathhelper`
27. `presentations-repos/presentation-aug-2026-plans-alim`
28. `02-prompts/prompts-connect`
29. `punam-case-studies-v1`
30. `presentations-repos/rasia-logo`
31. `scripts-fixer`
32. `presentations-repos/slides-spec`
33. `spec-builder`
34. `web-system/sweet-digs-finder`
35. `ui-prompts-cat`
36. `presentations-repos/white-presentation-v1`
37. `workflowy-ui`
38. `workflowy`
39. `wp-exam`
40. `wp-git-log`
41. `wp-html-automate`
42. `wp-link-manager`
43. `wp-onboarding`

### 4.2 The 5 Non-Negotiable Boundaries

All synchronization actions MUST enforce the five inviolable boundaries defined in `AGENTS.md` Section 11:

1. **Spec 21 Exclusion (TOTAL BAN):** NEVER synchronize, copy, or touch `02-spec/21-*` through `02-spec/25-*` (private application domain specs, issues, db, and UI designs). Only shared specifications `02-spec/01-*` through `02-spec/20-*` are mirrored.
2. **Additive-Only AI Scripts:** Brand-new AI scripts (`03-ai-scripts/`, `.agents/scripts/`) that do not exist in the target repository are copied cleanly. Existing scripts modified by the target repository must NEVER be overwritten.
3. **Bump Script Protection (IMMUTABLE):** NEVER overwrite version bump scripts or manifests (`bump*`, `bump_versions.py`, `bump-version.mjs`, `version.json`). Each repository maintains custom SemVer targets and release logic.
4. **Memory & Plans Protection (TOTAL ISOLATION):** NEVER modify, sync, or mirror `.ai-memory/memory/` or `.ai-memory/plans/` in target repositories.
5. **Zero Secrets Leakage:** NEVER synchronize `.env` files, tokens, or credentials across repositories. Secrets reside strictly in `repo-secrets` via `gitmap rs`.

### 4.3 Pre-Change Safety & Multi-Branch Protocol

Before mirroring assets, the automated synchronizer executes:
1. Detect base branch and pull latest changes: `git pull origin <base_branch> --no-rebase`.
2. Create and push safety backup branch: `git checkout -b backup/sync-<timestamp>` and `git push origin backup/sync-<timestamp>`.
3. Return to base branch: `git checkout <base_branch>` before applying asset updates.
4. Ensure pre-change release tag exists.
5. Mirror canonical assets with boundary protections.
6. Commit and push changes on base branch.
7. Post-change release ceremony with version bump and tag push.

---

## 5. Compliance, Linters & Atomic Push Workflow

### 5.1 Strict Relative Git Paths Verification
- Every generated specification, plan, subtask, and skill file must be scanned for absolute paths (`C:\`, `/home/`, `file:///`).
- Validation tool: `python 03-ai-scripts/07-relative-paths-linter.py` (or GitMap code search).
- Zero violations tolerated prior to commit.

### 5.2 Strict Lowercase Filename Verification
- Filenames must match `^[0-9a-z_.-]+$` without uppercase letters.
- Enforced on all files in `02-spec/`, `.ai-memory/plans/`, `01-prompts/`, and `.agents/skills/`.

### 5.3 Mandatory Completion Invariant ('No Push = Not Done')
Under `AGENTS.md` Section 13:
- Code changes without commit and upstream push are strictly considered INCOMPLETE.
- Commit command: `gitmap cpf "<module> - <summary>"` with hyphen-separated format and no colons in the summary string.
- Remote push verification via `gitmap push` or verified upstream tracking status.
