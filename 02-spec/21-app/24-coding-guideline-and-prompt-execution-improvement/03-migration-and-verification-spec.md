# Specification 03: GitMap Skills Enrichment, Real-World Examples & Companion Skill Alignment

## 1. System Overview & Executive Summary

### 1.1 Mission & Architectural Purpose
The primary objective of this specification is to eliminate command ambiguity, prevent flag hallucination, and enforce strict execution standards across all GitMap skills and prompts within the Prompt Architect meta-repository and connected fleet.

Historically, AI agents operating across complex execution tasks frequently fell back to slow shell commands (`Get-ChildItem`, `rg`, `Select-String`, `dir`, `cat`), hallucinated non-existent flags (such as `--max-concurrency` instead of `--concurrency` or `--workers`), invoked forbidden `--help` probes during mission runs, or referenced stale command signatures in companion skills (such as `gitmap agy clean-cache` instead of `gitmap agy cache-clear`, and `gitmap chrome picker` instead of `gitmap chrome list` / `gitmap chrome observe`).

This specification defines:
1. **The Comprehensive 12-Category Example Catalog:** A rich, production-grounded collection of concrete, copy-paste real-world examples to be embedded directly into `.agents/skills/gitmap/skill.md` and mirrored to `.cursor/skills/gitmap/skill.md`.
2. **Companion Skills Audit & Realignment:** Detailed surgical remediations across all companion skills (`.agents/skills/gitmap-*/`), excising all `--help` invocations, correcting invalid flags, aligning Antigravity IDE commands, and standardizing Chrome profile management.
3. **Multi-Phase Verification Strategy & Completion Gate:** Systematic validation protocol using relative path linters, skill parity verifiers, and cognitive heatmap closure.

---

## 2. Comprehensive 12-Category Example Catalog

The following 12 categories provide concrete, production-grade command patterns with exact flags, arguments, and expected output behavior to enrich `.agents/skills/gitmap/skill.md`.

### Category 1: Authentication & GitHub Credential Management
Managing credentials safely without interactive prompts, plaintext leakage, or browser popups during automated turns.

```bash
# 1. Non-interactive token ingestion (pre-validated against GitHub API):
gitmap login --token "ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

# 2. Offline token ingestion without network verification (for air-gapped or CI runners):
gitmap login --token "ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" --no-verify

# 3. Interactive OAuth / browser flow (manual developer sessions only):
gitmap login --web

# 4. Inspect current authenticated user and masked credential origin:
gitmap login --status

# 5. List resolved token metadata and scopes from git configuration:
gitmap token list

# 6. Securely purge stored credentials from global git configuration:
gitmap logout
```

### Category 2: Workspace & Multi-Repository Status
High-speed workspace auditing across 40+ repositories in a single invocation without shell loops.

```bash
# 1. Multi-repository status matrix filtered strictly to dirty working trees:
gitmap status --dirty

# 2. Check repositories with unpushed local commits waiting for remote delivery:
gitmap status --ahead

# 3. Check repositories with incoming remote commits waiting to be pulled:
gitmap status --behind

# 4. Structured machine-readable JSON status for automated agents and scripts:
gitmap status --json

# 5. Check if active repository remote tracking branch has any incoming updates:
gitmap has-any-updates

# 6. Query default remote branch name (main vs master) across repositories:
gitmap branch default

# 7. Discovers the most recently committed or updated remote branch:
gitmap latest-branch

# 8. Single-line compact status summary across all tracked repositories:
gitmap status --compact
```

### Category 3: Script Execution & Runner Engines
Cross-platform polyglot execution with automatic interpreter resolution, `-NoProfile` safety, and zero shell quoting breakage.

```bash
# 1. Run Python script with verified interpreter from installation.db:
gitmap py "03-ai-scripts/06-cicd-local-runner.py"

# 2. Inline Python expression evaluation with sub-millisecond execution:
gitmap py -c "import sys; print(sys.version_info[:2])"

# 3. Cross-platform PowerShell execution with deterministic -NoProfile and UTF-8:
gitmap pwsh "Get-Process | Where-Object WorkingSet -gt 500MB"

# 4. Direct inline PowerShell command evaluation:
gitmap ps -c "Test-Path .ai-memory/heatmap.db"

# 5. Cross-platform POSIX Bash execution across Windows, Linux, and macOS:
gitmap bash "chmod +x scripts/*.sh && ./scripts/check.sh"

# 6. Direct inline Bash command evaluation:
gitmap sh -c "uname -sm"

# 7. Runner dry-run preview with JSON telemetry output:
gitmap py "03-ai-scripts/14-version-sync-checker.py" --dry-run --json
```

### Category 4: Storage, Cache & Secrets Offloading
Preventing repository pollution by routing scratch scripts, diagnostic fixtures, and secrets to dedicated vaults.

```bash
# 1. Offload sensitive environment file into repo-secrets and auto-push:
gitmap rs file .env --repo coding-guidelines

# 2. Offload private directory into repo-secrets:
gitmap rs folder private/keys --repo coding-guidelines

# 3. Write secret string directly to repo-secrets vault:
gitmap rs text "ghp_secret_pat_value" --slug github-token

# 4. Offload temporary PowerShell diagnostic script to repo-cache (rc):
gitmap rc file debug-runner.ps1 --repo coding-guidelines

# 5. Offload fixture directory to repo-cache:
gitmap rc folder test/fixtures --repo coding-guidelines

# 6. Write reusable temporary test harness directly to repo-cache:
gitmap rc text "$env:DEBUG=1; go test -v ./..." --slug test-debug --ext .ps1

# 7. Fast navigation shortcuts to vault locations:
gitmap cd rs
gitmap cd rc
```

### Category 5: High-Performance Automation (AUM) & File Discovery
Replacing slow PowerShell, grep, and filesystem crawlers with multi-core streaming, lazy regex, and SQLite hot caches.

```bash
# 1. Multi-core streaming live code search with extension and directory scoping:
gitmap aum search "RunFleetPASCommand" cli -e .go

# 2. Case-insensitive regex search with boundary patterns:
gitmap aum search -r "(\"pas\"|\"pa\")" cli/cmd -e .go -i

# 3. Path-scoped live search restricted to specifications:
gitmap aum search "Subagents" 01-prompts -e .md

# 4. Sub-millisecond directory listing (instant replacement for Get-ChildItem):
gitmap lf "02-spec/21-app"

# 5. Fast glob file discovery across 10,000+ files in <10ms:
gitmap find "*.md" 02-spec/21-app

# 6. Substring file discovery across entire codebase:
gitmap ffa "heatmap"

# 7. Find files by filename prefix:
gitmap ffs "01-architecture"

# 8. Find files by filename suffix:
gitmap ffe "_test.go"

# 9. Direct zero-disk file content streaming into stdout:
gitmap cat "version.json"

# 10. Markdown sequence gap detector and header autofixer:
gitmap aum sequence

# 11. Normalize CRLF to LF line endings and strip trailing whitespace:
gitmap aum newlines --fix

# 12. Literal string replacement with recorded audit trail:
gitmap replace "old_function_name" "new_function_name"
```

### Category 6: Autonomous CI/CD Self-Healing (Pipeline AI)
Non-polling CI/CD status inspection, bounded failure extraction, and zero-storage artifact cleanup.

```bash
# 1. Dynamic non-polling ETA wait for running GitHub Actions workflow:
gitmap pipeline-ai status -t 120

# 2. Check pipeline workflow execution state with JSON output:
gitmap pipeline-ai status --json

# 3. Extract concise failing step error summary for 4-part RCA:
gitmap pe -t

# 4. Extract bounded failure frame lines (5 before, 20 after) from latest failed run:
gitmap pe -1

# 5. Clear accumulated pipeline error logs non-interactively:
gitmap pe clear -y

# 6. Analyze CI/CD pipeline historical trend across recent runs:
gitmap pe history-ai

# 7. Purge Actions workflow run artifacts to maintain strict 0.0 GB storage quota:
gitmap pipeline purge
```

### Category 7: Semantic & Flat Commit Suite
Disciplined atomic commit workflow enforcing hyphen-separated messages and remote delivery invariants.

```bash
# 1. Intermediate local checkpoint commit (auto-stages, no remote push, verbatim message):
gitmap commit "wip - draft initial types and interfaces"
gitmap cm "wip - draft initial types and interfaces"

# 2. Feature commit (auto-prefixes 'Feature: ', stages all, and pushes to GitHub):
gitmap cpf "skills - enrich gitmap skill with 12 real world categories"

# 3. Bugfix commit (auto-prefixes 'Bug: ', stages all, and pushes to GitHub):
gitmap cpb "cloner - correct concurrency flag parsing and worker limits"

# 4. Chore/Maintenance commit (auto-prefixes 'Chore: ', stages all, and pushes):
gitmap cpc "docs - realign companion skills and remove help flags"

# 5. Release chore commit (auto-prefixes release chore, stages, and pushes):
gitmap cpr "release - bump minor version and update changelog"

# 6. Pull latest with rebase, stage all, commit, and push in one verified flow:
gitmap pcp "skills - sync cross-platform runner templates"

# 7. Apply automated repository remediation (stash, discard, or auto-fix):
gitmap fix all -y

# 8. Safe 2-step git mv file case normalization for lowercase hygiene:
gitmap lowercase
gitmap lowercase-readme
```

### Category 8: Multi-Repo, Cluster & Toolchain Operations
Managing remote cluster topologies, SSH node pooling, and developer toolchains.

```bash
# 1. Inspect cluster node status and heartbeat connectivity:
gitmap cluster status

# 2. Execute command across all online cluster nodes in parallel:
gitmap cluster exec "df -h / && free -m"

# 3. Servers-clients topology service list:
gitmap sc list

# 4. Run background daemon task via server-client engine:
gitmap sc bash "python -m http.server 8080"

# 5. Check remote SSH fleet health and port 22 connectivity:
gitmap ssh health --all

# 6. Register remote node with authentication vault:
gitmap ssh join user@192.168.1.50 --alias worker-node-1

# 7. Inspect Rust and Cargo toolchain status:
gitmap cargo status

# 8. Install missing toolchain or system runtime:
gitmap install cargo
```

### Category 9: Native Fleet Synchronization & SQLite Agent Task Engine
Compiled native Go task tracking (<1ms WAL mode) and 43-repo synchronization.

```bash
# 1. Synchronize canonical prompts, skills, and scripts across all 43 repositories:
gitmap sync --workers 8

# 2. Synchronize a single targeted repository:
gitmap sync --repo aukgit/core

# 3. Preview fleet synchronization without applying changes:
gitmap sync --dry-run

# 4. Initialize agent task tracking in .ai-memory/temp-agents/<slug>/agent-task.db:
gitmap agent task init --name "skills-enrichment" --budget 400

# 5. Add decomposed subtask bounding box with exclusive file ownership:
gitmap agent subtask add --code "ST-01" --slug "enrich-gitmap-skill" --title "Enrich gitmap skill with examples"

# 6. List pending subtasks in current task ledger:
gitmap agent task pending

# 7. Claim next available subtask atomically:
gitmap agent subtask claim --agent "worker-01"

# 8. Mark claimed subtask completed with verification evidence:
gitmap agent subtask complete 1 --agent "worker-01" --evidence "12 categories added and verified"

# 9. Perform post-mortem autopsy and crash diagnosis:
gitmap agent diagnose
gitmap agent crashed

# 10. Clean ephemeral agent task database after verified completion:
gitmap agent temp-clear
```

### Category 10: File Contention Heatmap, Cognitive Telemetry & RC Sync
Sub-millisecond SQLite WAL tracking of agent decisions, active leases, and contention hotspots.

```bash
# 1. Step 0 Pre-Flight schema and status verification:
gitmap heatmap status
gitmap heatmap tables

# 2. Pre-Action Intent Logging (mandatory prior to reading or modifying files):
gitmap heatmap log --task "skills-enrichment" --action "READ" --file ".agents/skills/gitmap/skill.md" --state "EXPLORATION" --expected "Analyze existing command matrix"

# 3. Post-Action Outcome Closure (mandatory immediately after operation):
gitmap heatmap update --id 1 --outcome "Identified missing real-world examples" --met true

# 4. Render hierarchical causal tree of all actions taken during the task:
gitmap heatmap tree --task "skills-enrichment"

# 5. Multi-agent file contention analysis and collision blast radius check:
gitmap agent heatmap

# 6. Safe read-only SQL query against cognitive heatmap database:
gitmap heatmap query "SELECT HeatmapActionLogId, ActionType, FilePath, IsExpectationMet FROM HeatmapActionLog ORDER BY HeatmapActionLogId DESC LIMIT 5" --json

# 7. Companion repo-cache (rc) heatmap synchronization:
gitmap rc heatmap status
gitmap rc heatmap sync
```

### Category 11: Native SQLite Database Exploration, REST API Server & Lifecycle Suite
High-speed native Go SQLite exploration, mutation, join views, and dynamic REST gateways.

```bash
# 1. List all registered and system SQLite databases with sizes and query metrics:
gitmap db-listing
gitmap dbl

# 2. List all tables and views in database with column and row counts:
gitmap db-read ".ai-memory/heatmap.db" ls

# 3. Execute read-only SQL query with timing telemetry:
gitmap db-read ".ai-memory/heatmap.db" "SELECT COUNT(*) FROM HeatmapActionLog"

# 4. Inspect table schema, data types, and primary key constraints:
gitmap db-read ".ai-memory/heatmap.db" table HeatmapActionLog --columns --json

# 5. Insert new record from CSV values:
gitmap db-add my-db HeatmapActionLog csv "READ","02-spec/spec.md","EXPLORATION"

# 6. Update existing record by primary key:
gitmap db-edit my-db -table HeatmapActionLog -id 42 csv "EDIT","02-spec/spec.md","EXECUTION"

# 7. Search database across multiple tables using full-text wildcard:
gitmap db-find my-db -table HeatmapActionLog,HeatmapParentTask --find "enrichment" --json

# 8. Create persistent joined view across tables:
gitmap db-view my-db -t "HeatmapActionLog as a, HeatmapParentTask as p" --viewname ViewTaskActions --find "a.TaskSlug = p.TaskSlug"

# 9. Launch REST API daemon with auto-generated Swagger UI (/swagger) and Visualizer (/ui):
gitmap db-rest my-db 8080,8081

# 10. Delete a single record by primary key:
gitmap db-record-del my-db -table HeatmapActionLog -id 42 -y
```

### Category 12: Screenshot Extraction & Lightshot Downloading
Native compiled parallel screenshot asset extraction, CDN resolution, and Antigravity Brain sync.

```bash
# 1. Download single Lightshot/prnt.sc screenshot and auto-sync to Antigravity Brain:
gitmap prnt "https://prnt.sc/abcdef"

# 2. Batch download multiple screenshot URLs in parallel:
gitmap prnt "https://prnt.sc/url1" "https://prnt.sc/url2" "https://prnt.sc/url3"

# 3. Resolve direct image URLs without downloading (dry-run probe):
gitmap prnt "https://prnt.sc/abcdef" --extract-only

# 4. Emits structured Universal JSON response envelope:
gitmap prnt "https://prnt.sc/abcdef" --json

# 5. Save downloaded screenshots to custom target directory:
gitmap prnt "https://prnt.sc/abcdef" -o "assets/screenshots/ui-audit"
```

---

## 3. Companion Skills Audit & Realignment Specification

### 3.1 Prohibited `--help` Instruction Elimination
Under Rule R4 and Coding Guidelines, AI agents must never run `--help` commands during automated task execution. Harmless probe commands or active management subcommands must replace all occurrences of `--help`.

| File Path | Prohibited Legacy Instruction | Corrected Active Command | Rationale |
| :--- | :--- | :--- | :--- |
| `.agents/skills/gitmap/skill.md` (L177) | `gitmap cluster --help` | `gitmap cluster status` | Active status inspection without dumping help text. |
| `.agents/skills/gitmap/skill.md` (L178) | `gitmap sc --help` | `gitmap sc list` | Lists active and registered services directly. |
| `.agents/skills/gitmap/skill.md` (L179) | `gitmap ssh --help` | `gitmap ssh health --all` | Verifies port 22 liveness across known nodes. |
| `.cursor/skills/gitmap/skill.md` | `gitmap cluster/sc/ssh --help` | `gitmap cluster status`, `sc list`, `ssh health` | Mirrors Antigravity skill corrections. |

### 3.2 Concurrency Flag Rectification
AI models occasionally hallucinate `--max-concurrency` or `--threads`. The canonical flags across GitMap commands are strictly:
1. `gitmap clone`: Uses `--concurrency <N>` (e.g. `gitmap clone manifest.json --concurrency 8`).
2. `gitmap sync`: Uses `--workers <N>` (e.g. `gitmap sync --workers 8`).
3. `gitmap aum search`: Uses `-w, --workers <N>` (e.g. `gitmap aum search "pattern" -w 8`).

All skill files must cite these verified flags explicitly and add cautionary warnings against hallucinated variants.

### 3.3 Antigravity Skill Rectification (`gitmap-developer-hygiene-and-agy`)
The following command corrections must be applied to `.agents/skills/gitmap-developer-hygiene-and-agy/skill.md` and `.cursor/skills/gitmap-developer-hygiene-and-agy/skill.md`:

| Legacy / Hallucinated Command | Corrected GitMap AGY Command | Operational Behavior |
| :--- | :--- | :--- |
| `gitmap agy clean-cache` | `gitmap agy cache-clear` (alias `cc`) | Cleans runtime cache and prunes old conversations (`--keep 10`). |
| `gitmap agy prune --except <id>` | `gitmap agy ccko` / `gitmap agy cckf` | Cleans cache keeping only 1 (`ccko`) or top 5 (`cckf`) active conversations. |
| (None / Missing) | `gitmap agy remove-empty-convs` | Purges orphaned projects with 0 conversation steps. |
| (None / Missing) | `gitmap agy running-projects ls` (alias `rp`) | Lists active Antigravity projects and open conversation IDs. |

### 3.4 Chrome Profile Skill Rectification
The following corrections must be applied to Chrome profile management sections:

| Legacy / Hallucinated Command | Corrected GitMap Chrome Command | Operational Behavior |
| :--- | :--- | :--- |
| `gitmap chrome profiles --reconcile` | `gitmap chrome reconcile` (alias `repair`, `sync-state`) | Synchronizes on-disk directories with Chromium `Local State`. |
| `gitmap chrome picker` | `gitmap chrome list` (alias `ls`, `cpl`) | Lists all Chrome profiles known to GitMap. |
| (None / Missing) | `gitmap chrome observe` (alias `tabs`, `status`) | Inspects running Chrome processes and open tabs. |
| (None / Missing) | `gitmap chrome open [urls]` (alias `launch`, `tab`) | Opens URLs in specified or active Chrome profile. |
| (None / Missing) | `gitmap chrome reset [prof] [flags]` | Safely purges caches, cookies, or history without profile loss. |

---

## 4. Multi-Phase Verification Strategy & Completion Gate

To guarantee flawless execution, the implementation must pass the following 5-phase verification protocol:

```
+-----------------------------------------------------------------------------------+
| PHASE 1: Static Syntax & Schema Conformance                                      |
| -> Verify strictly lowercase filenames                                            |
| -> Validate markdown frontmatter in .agents/skills/ and .cursor/skills/           |
+-----------------------------------------------------------------------------------+
                                         |
                                         v
+-----------------------------------------------------------------------------------+
| PHASE 2: Relative Git Paths & Link Hygiene                                        |
| -> python linter-scripts/check-relative-paths.py                                  |
| -> Verify zero absolute paths (C:\, /home) and zero file:/// URIs                 |
+-----------------------------------------------------------------------------------+
                                         |
                                         v
+-----------------------------------------------------------------------------------+
| PHASE 3: Dual Skill Tree Parity Validation                                        |
| -> Verify 1:1 text parity between .agents/skills/ and .cursor/skills/             |
| -> Confirm zero remaining --help flags or hallucinated arguments                 |
+-----------------------------------------------------------------------------------+
                                         |
                                         v
+-----------------------------------------------------------------------------------+
| PHASE 4: Cognitive Heatmap Closure Audit                                          |
| -> gitmap heatmap status --task coding-guideline-and-prompt-execution-improvement |
| -> gitmap heatmap tree --task coding-guideline-and-prompt-execution-improvement   |
| -> Confirm 100% action closure (isExpectationMet = true)                          |
+-----------------------------------------------------------------------------------+
                                         |
                                         v
+-----------------------------------------------------------------------------------+
| PHASE 5: Completion Gate & Atomic Push Protocol                                   |
| -> Orchestrator commit via:                                                       |
|    gitmap cpf "skills - enrich gitmap skill and realign companion skills"         |
| -> Remote push to origin tracking branch                                          |
+-----------------------------------------------------------------------------------+
```

### 4.1 Automated Linter Commands

```bash
# 1. Check relative path hygiene across updated skills and specs:
python 03-ai-scripts/07-relative-path-fixer.py --check-only

# 2. Check prompt and skill sequence integrity:
python 03-ai-scripts/21-sequence-integrity-linter.py

# 3. Check documentation paths and markdown links:
python 03-ai-scripts/22-doc-path-linter.py

# 4. Heatmap verification audit:
gitmap heatmap status --task "coding-guideline-and-prompt-execution-improvement"
gitmap heatmap tree --task "coding-guideline-and-prompt-execution-improvement"
```

### 4.2 Non-Negotiable Invariants for Completion
1. **Zero Git Commands by Subagents:** Subagents must NEVER execute `git` commands (`git add`, `git commit`, `git push`, `git status`). All version control staging and pushing is strictly handled by the root orchestrator.
2. **Zero Storage Ban:** GitHub Actions workflows must NEVER upload build artifacts or logs. All summaries must write to `$GITHUB_STEP_SUMMARY`.
3. **No Push = Not Done:** Under Section 13 of AGENTS.md, work is only complete when committed and pushed upstream to GitHub.
