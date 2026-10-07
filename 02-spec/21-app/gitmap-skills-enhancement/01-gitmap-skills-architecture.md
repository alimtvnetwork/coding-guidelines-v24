# GitMap Skills Expansion & Cross-Fleet Synchronization Architecture

## 1. Executive Summary & Objective

This specification defines the comprehensive expansion of GitMap skills within `coding-guidelines` and their deterministic synchronization across all 43 connected repositories in the workspace.

GitMap (`gitmap-cli`) is the primary developer companion CLI tool powering ultra-fast repository search, workspace telemetry, cross-platform script execution, zero-storage CI/CD diagnostics, and atomic hyphenated commits. Previously, the bundled skill was minimal and lacked critical operational commands such as `gitmap login`, workspace status inspection, direct execution runners for Python, PowerShell, and Bash, and an explicit enforcement matrix replacing legacy shell tools.

This architecture formalizes:
1. **Authentication & Credential Management:** `gitmap login`, `gitmap login --web`, `gitmap login --token`, `gitmap login --status`, and `gitmap logout`.
2. **Repository & Workspace Status:** `gitmap status`, `gitmap st`, `--dirty`, `--ahead`, `--behind`, `--json`, `gitmap hau` (`has-any-updates`), and `gitmap lb` (`latest-branch`).
3. **High-Speed Execution Engines:** `gitmap py`, `gitmap pwsh` / `gitmap ps`, and `gitmap bash` / `gitmap sh` eliminating slow ambient shell invocations.
4. **Storage & Secret Offloading:** `gitmap rs` (zero secrets in repos) and `gitmap rc` (reusable test harness and script storage).
5. **Strict Substitution Matrix:** Explicit forbidden shell tools (`rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, raw `python`, raw `powershell`) with required GitMap drop-in equivalents.
6. **Cross-Fleet Synchronization:** Mirroring the updated `.agents/skills/gitmap/` and `.cursor/skills/gitmap/` to all 43 repositories following the 5 Non-Negotiable Boundaries.

## 2. Architectural Boundaries & Non-Negotiable Invariants

- **Strict Relative Git Paths:** All documentation, plans, subtasks, and skill files MUST exclusively use relative paths from the repository root (e.g., `02-spec/...`, `.agents/skills/...`). Absolute paths and `file:///` URIs are strictly banned.
- **Hyphen-Separated Commits:** All GitMap commit commands (`gitmap cpf`, `gitmap cpb`, `gitmap cpr`) MUST format commit messages with hyphens and NEVER include colons (`gitmap cpf "skills - expand gitmap skill suite"`).
- **Tool Primacy:** Codebase exploration across all repositories MUST use GitMap commands (`gitmap aum search`, `gitmap find`, `gitmap cat`, `gitmap ps`, `gitmap py`).
- **5 Synchronization Boundaries:**
  1. Spec 21 Exclusion (`02-spec/21-*` through `25-*` are strictly local to each repository and never synchronized).
  2. Additive-Only AI Scripts (new scripts are copied; existing repo-modified scripts are preserved).
  3. Bump Script Protection (version bump scripts are never overwritten).
  4. Memory & Plans Protection (`.ai-memory/memory/` and `.ai-memory/plans/` are isolated).
  5. Zero Secrets Leakage (credentials belong strictly in `repo-secrets` via `gitmap rs`).

## 3. Skill Content Architecture

The expanded GitMap skill (`.agents/skills/gitmap/skill.md` and `.cursor/skills/gitmap/skill.md`) is structured into dedicated operational sections:

### Section 1: Authentication & GitHub Access
- `gitmap login` — Interactive GitHub authentication picker.
- `gitmap login --web` (alias `--browser`) — Browser-assisted authentication via GitHub CLI / OAuth flow.
- `gitmap login --token <PAT>` — Secure token ingestion validated against GitHub API before writing.
- `gitmap login --status` — Displays current masked credential status.
- `gitmap token list` — Resolves active token source.
- `gitmap logout` — Purges stored credentials from Git global configuration.

### Section 2: Workspace & Repository Status
- `gitmap status` (alias `gitmap st`) — Comprehensive multi-repo status table across tracked workspace.
- `gitmap status --dirty` (`-d`) — Filters solely to repositories with uncommitted working changes.
- `gitmap status --ahead` / `--behind` — Pinpoints repositories out of sync with upstream remotes.
- `gitmap status --json` (`-j`) — Machine-readable status telemetry for AI scripts.
- `gitmap has-any-updates` (alias `gitmap hau`, `gitmap hac`) — Checks remote tracking branch for incoming commits.
- `gitmap latest-branch` (alias `gitmap lb`) — Discovers the most recently committed remote branch.
- `gitmap watch` (alias `gitmap w`) — Live dashboard monitoring working tree state.

### Section 3: High-Speed Script & Shell Runners
- `gitmap py "<code-or-script>"` — High-performance Python execution without virtualenv or PATH friction.
- `gitmap py -c "<expression>"` — Direct inline Python evaluation.
- `gitmap pwsh "<cmd>"` / `gitmap ps "<cmd>"` — Cross-platform PowerShell execution with `-NoProfile`.
- `gitmap ps -c "<cmd>"` — Direct inline PowerShell execution.
- `gitmap bash "<cmd>"` / `gitmap sh "<cmd>"` — Cross-platform Bash execution.
- `gitmap bash -c "<cmd>"` — Direct inline Bash command evaluation.

### Section 4: Reusable Storage & Secrets Offloading
- `gitmap rs file <path> [--repo <name>]` — Offloads sensitive credentials to `repo-secrets`.
- `gitmap rs text "<secret>" [--slug <slug>]` — Saves secret token into encrypted `repo-secrets`.
- `gitmap rc file <path> [--repo <name>]` — Persists test scripts and diagnostic tools in `repo-cache`.
- `gitmap rc text "<content>" --slug <slug> --ext .ps1` — Generates reusable cross-repo scripts.

### Section 5: High-Speed File Discovery & In-Memory Cat
- `gitmap aum search "<pattern>" [dir] [--ext <ext>]` — Multi-threaded regex/symbol search (<15ms).
- `gitmap search "<query>"` — SQLite DH2D hot-cache symbol search.
- `gitmap find "<pattern>" [-ext <ext>]` — Glob file locator.
- `gitmap find-files <name>` (`ff`), `find-files-any <str>` (`ffa`).
- `gitmap list-files [dir]` (`lf`).
- `gitmap cat <filepath>` — Zero-disk direct stream of file content to terminal stdout.

### Section 6: Semantic Git & Branch Operations
- `gitmap cpf "<module> - <summary>"` — Stage, commit, and push feature branch.
- `gitmap cpb "<module> - <summary>"` — Stage, commit, and push bugfix branch.
- `gitmap cpr "<module> - <summary>"` — Stage, commit, and push chore/refactor.
- `gitmap pcp "<module> - <summary>"` — Preflight pull, atomic commit, and push.
- `gitmap pull <name>` / `gitmap pull-all` (`pa`) — Pull tracked repositories.
- `gitmap fix [repo]` — Remediate merge or working tree locks.
- `gitmap lowercase` (`lcf`) — Batch case normalization for file systems.

### Section 7: Zero-Storage Pipeline AI
- `gitmap pipeline-ai status -t <eta>` — Dynamic timeout sleep waiting for CI/CD runs.
- `gitmap pipeline error-logs` (`gitmap pe`, `gitmap pe -t`) — Structured error log extraction for 4-part RCA.
- `gitmap pipeline purge` — GitHub Actions artifact purge maintaining 0.0 GB quota.

### Section 8: Strict Prohibition & Positive Replacement Matrix
An authoritative lookup table mapping legacy shell commands to mandatory GitMap alternatives:
| Prohibited Legacy Command | GitMap Mandatory Alternative | Justification |
| :--- | :--- | :--- |
| `rg`, `ripgrep`, `grep`, `git grep` | `gitmap aum search "<pattern>" [dir]` | Multi-threaded streaming with binary null-byte probe & 500 KB guard |
| `Select-String -Pattern "..."` | `gitmap aum search "<pattern>" [dir]` | Avoids PowerShell pipeline boxing & memory bloat |
| `Get-ChildItem -Recurse`, `find .` | `gitmap find "<pattern>"`, `gitmap lf` | Sub-millisecond glob resolution via GitMap index |
| Raw unbuffered `cat`, `type` | `gitmap cat <filepath>` | Zero-disk streamed stdout without terminal freeze |
| Raw `python <script>` | `gitmap py "<script>"` | Auto-resolves cached Python interpreter via `installation.db` |
| `powershell -Command "..."` | `gitmap ps -c "..."` / `gitmap pwsh` | Deterministic `-NoProfile` execution across platforms |
| `bash -c "..."` / `sh -c "..."` | `gitmap bash -c "..."` / `gitmap sh` | Uniform POSIX execution with cross-platform environment isolation |
| `git status` per-repo loops | `gitmap status --dirty` / `--json` | Single-shot multi-repository audit matrix |
| `gh auth login`, plaintext `.env` | `gitmap login --web`, `gitmap rs` | Validated credential storage & zero repo secrets leakage |
| Commits with colons (`git commit -m "feat: .."`) | `gitmap cpf "<module> - <summary>"` | Automatic conventional prefixing with strict hyphen formatting |
| Raw script pollution in repo | `gitmap rc text "..." --slug <slug>` | Centralized cache reuse without git tree pollution |
| Tight CI polling loops (`gh run watch`) | `gitmap pipeline-ai status -t <eta>` | Dynamic non-blocking ETA wait |

## 4. Rollout & Fleet Synchronization Plan

1. **Step 1:** Expand `.agents/skills/gitmap/skill.md` and `.cursor/skills/gitmap/skill.md` with complete documentation, command schemas, and the substitution matrix.
2. **Step 2:** Commit changes locally in `coding-guidelines` using `gitmap cpf "skills - expand gitmap skill suite with authentication status and runners"`.
3. **Step 3:** Execute fleet synchronization across all 43 connected repositories using `03-ai-scripts/38-sync-prompts-skills-scripts.py`. Each repository is pulled, backed up with `backup/sync-<timestamp>`, updated, and cleanly synced.
