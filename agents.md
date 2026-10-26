<!-- LOVABLE:BEGIN -->
> [!IMPORTANT]
> This project is connected to [Lovable](https://lovable.dev). Avoid rewriting
> published git history — force pushing, or rebasing/amending/squashing commits
> that are already pushed — as it rewrites history on Lovable's side and the
> user will likely lose their project history.
>
> Commits you push to the connected branch sync back to Lovable and show up in
> the editor, so keep the branch in a working state.
<!-- LOVABLE:END -->

# Prompt Architect: Global AI Guidelines

The following rules apply to all AI agents operating within the Prompt Architect meta-repository and any codebase it manages.

## 1. Boolean Principles (Cross-Language)

- **No Explicit True Checks (TOTAL BAN):** NEVER evaluate a boolean explicitly against `true` (e.g., `if isReady == true`). Positive booleans MUST ALWAYS be evaluated implicitly: `if isReady { ... }`.
- **No Mixed Polarity:** NEVER combine a positive check and a negative check in the same `if` condition (e.g., `if isA && !isB`).

## 2. STRICT AVOIDANCE: Never Disable CI/CD

- **NEVER disable any CI/CD checks, GitHub Actions, or validation workflows.**
- Strictly avoid commenting out, bypassing, or deleting CI/CD steps to force a pipeline to pass. Your job is to fix the underlying code so that the CI/CD pipeline passes legitimately. Disabling CI/CD is an auto-reject failure.

## 3. Anti-Hallucination, Micro-Tasking, & Self-Looping

To survive massive checklists and complex codebases, you MUST operate using these three principles:

1. **Phase 1: Read & Understand (Isolated Loop):** Your very first action must be purely exploratory. Do NOT write code. Break down the task, read the specific files, trace the dependencies, and understand the architectural boundary. Once you understand the scope, end your turn and self-loop to begin execution.
2. **Phase 2: Bounded Micro-Tasking (Sequential Self-Looping):** Never attempt to execute the entire checklist in one response. Treat each checklist section or file as a strict, isolated boundary. Execute *only* the first small portion, verify it, end your turn, and self-loop to process the next portion.
3. **Phase 3: Multi-Agent Parallelization:** If tasks are independent, you MUST spawn dedicated sub-agents to handle them concurrently. Give each sub-agent an extremely small, strictly defined bounding box (e.g., "Only edit File X"). Never give a sub-agent a generic or multi-file task.

## 4. Lowercase File Naming Convention

- **Strict Lowercase (No Exceptions):** All files, scripts, documentation, and system files generated or modified by the AI MUST use strictly lowercase naming (e.g., `readme.md`, `01-file-manipulator.py`, `agents.md`, `skill.md`). There are absolutely no exceptions for uppercase letters in filenames.

## 5. Strict Relative Git Paths Mandate (TOTAL BAN on Absolute Paths / `file:///` URIs)

- **Strict Relative Git Paths:** All file paths, markdown links, citations, subtask paths (`.ai-memory/plans/subtasks/`), release pages, release notes, changelogs, and memory logs MUST be strictly relative paths starting from the git repository root (e.g., `02-spec/02-coding-guidelines/04-error-handling.md`, `.ai-memory/plans/subtasks/01-task.md`, `cmd/main.go`). Only add the relative paths, never add the absolute path during your work. This MUST be respected on the release page and in release notes as well.
- **TOTAL BAN:** NEVER write absolute filesystem paths (e.g., `/absolute/path/to/...`, `C:\Users\...`, `/home/...`) or absolute file URIs (`file:///absolute/path/to/...`, `file:///absolute/path/to/`) inside ANY repository files, plans, specs, comments, release pages, release notes, or documentation.
  - ❌ **BAD:** `[SSH Commands](file:///absolute/path/to/.ai-memory/02-spec/commands/01-ssh-commands.md)`
  - ❌ **BAD:** `Target File: /absolute/path/to/...\cmd\login.go`
  - ✅ **GOOD:** `[SSH Commands](.ai-memory/02-spec/commands/01-ssh-commands.md)`
  - ✅ **GOOD:** `Target File: cmd/login.go`

## 6. Go Error Return Type: `*appfault.AppError` (Standard)

- **Structured Go AppError Type:** In all Go packages, functions returning structured failure metadata MUST use `*appfault.AppError` as their return type (e.g. `func validate() *appfault.AppError`).
- **Package Naming Convention:** Use package `appfault` (`04-code/golang/pkg/appfault`) with struct `AppError` to eliminate redundant package-stutter (e.g. `appfault.AppError` instead of `apperror.AppError`).
- **Result Containers:** `Result[T]`, `ResultSlice[T]`, and `ResultMap[K, V]` provide `.AppError()` and `.Fault()` returning `*appfault.AppError`.
- **AI Migration Rule:** When encountering legacy code or specs referencing `*apperror.AppError` or `*apperror.Fault`, AI agents MUST update the import to `pkg/appfault` and type to `*appfault.AppError`. Package `appfault` provides `type Fault = AppError`, and package `apperror` provides alias forwarders for non-breaking compatibility.

## 7. Zero-Storage GitHub Actions Mandate (Total Ban on CI Artifact Uploads)

- **Total Ban on `actions/upload-artifact`:** GitHub Actions workflows MUST NOT upload routine build artifacts, test results, Playwright reports, coverage files, drift reports, or logs to GitHub Actions storage. Free-tier accounts have a strict 0.5 GB quota across all account repositories; uploading artifacts exhausts this quota rapidly and blocks repository workflows.
- **Zero-Storage Diagnostic Reporting:**
  - All test reports, drift summaries, and lint outputs MUST be written directly to `$GITHUB_STEP_SUMMARY` (renders natively with 0 storage cost) or console standard output (`cat file.log`).
  - Diagnostic failures MUST use GitHub Actions annotations (`::error::` / `::warning::`).
  - Pull request summaries MUST use sticky PR comments.
- **Release Assets Exemption:** Distribution binaries and release archives attached directly to GitHub Releases via `gh release create` / `gh release upload` are exempt from this ban because GitHub Release assets do not consume the monthly Actions storage quota.

## 8. Essential AI Coding Constraints

- **Strict Boolean Standard:** is and has only (can, should, was, etc. are banned).
- **No Bare Void in Go:** Functions must return Result[T] or *appfault.AppError.
- **Parameter Structs:** Banned loose >2-3 parameters; use *Params structs.
- **Vertical Line Gaps:** Mandatory blank lines before if, after }, before return, and around multiline struct calls.
- **5-8 Files Micro-Batching:** All refactors broken into bounded subtasks.

## 9. Special Default Repositories (`repo-secrets` / `rs` & `repo-cache` / `rc`)

- **Zero Secrets in Standard Repositories (`repo-secrets` / `rs`):** NEVER store `.env` files, API keys, passwords, authentication email/password pairs, tokens, or private credentials inside standard or public repositories.
  - **Context & Location Mandate:** Never specify or provide repository URLs, git remote URLs, or absolute folder paths. State strictly that if the `repo-secrets` folder exists in the default work directory, that is the context where secrets MUST be stored.
  - If the `repo-secrets` folder exists in the default work directory, offload and store secrets in `repo-secrets` using `gitmap rs`:
    - `gitmap cd rs` — Navigate directly to `repo-secrets`.
    - `gitmap rs file <filepath> [--repo <repo-name>]` — Copies `<filepath>` into `repo-secrets/<XX-repo-name>/<01-filename>` and automatically commits & pushes.
    - `gitmap rs folder <folderpath> [--repo <repo-name>]` — Copies `<folderpath>` into `repo-secrets/<XX-repo-name>/<01-foldername>` and automatically commits & pushes.
    - `gitmap rs text "<secret-or-password>" [--slug <slug>]` — Writes `<secret-or-password>` into `repo-secrets/<XX-repo-name>/<01-slug>.txt` and automatically commits & pushes.
- **Reusable Temporary Scripts & Test Harnesses (`repo-cache` / `rc`):** Whenever temporary scripts (such as PowerShell `.ps1` scripts, diagnostic harnesses, or reusable test fixtures) are created during development or debugging, store them in `repo-cache` (`repo-storage`) via `gitmap rc` so any repository can reuse them cleanly without polluting standard git worktrees:
  - `gitmap cd rc` — Navigate directly to `repo-cache`.
  - `gitmap rc file <script.ps1> [--repo <repo-name>]` — Copies `<script.ps1>` into `repo-cache/<XX-repo-name>/<01-script.ps1>` and automatically commits & pushes.
  - `gitmap rc folder <folderpath> [--repo <repo-name>]` — Copies `<folderpath>` into `repo-cache/<XX-repo-name>/<01-foldername>` and automatically commits & pushes.
  - `gitmap rc text "<script-content>" --slug <slug> --ext .ps1` — Writes `<script-content>` into `repo-cache/<XX-repo-name>/<01-slug>.ps1` and automatically commits & pushes.

## 10. IDE Skill Link Syntax & Prompt Formatter Invariants

- **Mandatory IDE Skill Link Syntax:** All skill links within prompt templates, instructions, and documentation MUST strictly use the IDE-detectable syntax:
  - Antigravity Agents: `[<skill-name>](file;.agents/skills/<skill-name>)`
  - Cursor IDE: `[<skill-name>](file;.cursor/skills/<skill-name>)`
  - Slash Commands: `[/plan](slashCommand;plan)` and `[/learn](slashCommand;learn)`
  - NEVER append `/skill.md` or use raw filesystem paths that break IDE interactive detection.

- **Execution Formatter Standards (Letterly & Cursor):**
  - Line 1 of generated execution prompts must begin directly with `# High Priority Instruction`.
  - Header slug in Letterly execution prompt templates must be formatted as `## slug: <task-slug>`.
  - Slash command suggestions such as `[/plan](slashCommand;plan)` belong strictly under `## Additional Instructions` at the bottom of execution prompts, and are strictly prohibited on verification audit prompts.
  - The plan directive motto under `## Additional Instructions` must explicitly state: `[/plan](slashCommand;plan) first before doing work for improved work or less friction, faster work and more pure and effective work.`
  - Item 1/2 of Actionable Items must explicitly mandate writing specs in `02-spec/21-app/XX-<slug>/` and enqueueing plan tasks in `.ai-memory/plans/pending/XX-<slug>.md` (with subtasks in `.ai-memory/plans/subtasks/XX-<slug>/`), where `XX` is the zero-padded sequential number of that folder (e.g. `xx is sequence 01, 02`).
  - When writing about planning (Actionable Item 2 and Phase 1), there is NO NEED to include execute steps. Planning is strictly dedicated to architecture, specification, boundaries, and task decomposition (NO-BUILD, NO-TEST, NO-CODE-CHANGES). Execution steps belong strictly to subsequent phases.
  - Item 2 of Actionable Items must explicitly mandate searching codebase exclusively via GitMap (`gitmap aum search`, `gitmap find`, `gitmap cat`, `gitmap ps`) with a total ban on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`.
  - Item 3 of Actionable Items must explicitly mandate strictly relative Git paths (`02-spec/...`, `.ai-memory/...`, `cmd/...`): only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well.
  - Retrospective AI verification is an audit and verification gate, NOT a planning step. Only `09-execute-with-verification-letterly.md` (and its mirrors) contains retrospective AI verification; standard `execute-n-steps` formatters MUST NOT mandate retrospective AI verification.
  - Execution with verification must embed both `execute-parent-task-with-n-steps-v6` and `ai-verification`.

- **GitMap Commit Formatting:**
  - In all GitMap commit commands (`gitmap cpf`, `gitmap cpb`, `gitmap cpc`, `gitmap cpr`), NEVER include colons inside the message argument. GitMap automatically provides `Feature: `, `Bug: `, or chore prefixes. Format messages strictly with hyphens:
    - Feature: `gitmap cpf "<module> - <summary>"`
    - Bug Fix: `gitmap cpb "<module> - <summary>"`
    - Chore / Maintenance: `gitmap cpc "<module> - <summary>"`
    - Release: `gitmap cpr "<module> - <summary>"`

  - **GitMap Semantic & Flat Commit Suite:**
    - **Flat Commit Syntax:** `gitmap commit "<message>"` (alias: `gitmap cm "<message>"`).
    - **Mechanics:** 1-step auto-stage (`git add -A`) and local commit without semantic prefixing ('Feature: ', 'Bug: ') and without remote push.
    - **Chore Commit:** `gitmap cpc "<module> - <summary>"` stages all files, commits with chore prefix, and pushes upstream.
    - **Operational Distinction:**
      - Flat local commits (`gitmap commit`, `gitmap cm`) create instant local checkpoints during active iteration and micro-batches without altering commit message prefixes and without touching network remotes or triggering CI/CD webhooks.
      - Semantic auto-pushing commits (`gitmap cpf`, `gitmap cpb`, `gitmap cpc`, `gitmap cpr`) enforce structured prefixes, validate format, and push directly to GitHub tracking branches.
    - **Mandatory Completion Invariant Reminder:** Flat commits are strictly for intermediate local iteration. Under Section 13 ('No Push = Not Done'), every task completion mandates a semantic auto-pushing commit (`gitmap cpf "<module> - <summary>"`) or an explicit remote push (`gitmap push`). Leaving unpushed local flat commits upon task completion violates the completion invariant.

## 11. Multi-Repository Synchronization & Boundary Invariants

When synchronizing canonical assets (prompts, skills, shared specs, scripts) across connected repositories:

- **Pre-Change Safety Protocol (Pull & Backup):**
  1. Detect the base branch and pull latest changes: `git pull origin <base_branch> --no-rebase`.
  2. Create and push a safety backup branch before making changes: `git checkout -b backup/sync-<timestamp>` and `git push origin backup/sync-<timestamp>`.
  3. Return to the base branch: `git checkout <base_branch>` before applying asset updates.

- **The 5 Non-Negotiable Boundaries:**
  1. **Spec 21 Exclusion (TOTAL BAN):** NEVER synchronize, copy, or touch `02-spec/21-*` through `02-spec/25-*` (private application domain specs, issues, db, and UI designs). Only shared specifications `02-spec/01-*` through `02-spec/20-*` are synchronized.
  2. **Additive-Only AI Scripts:** Brand-new AI scripts (`03-ai-scripts/`, `.agents/scripts/`) that do not exist in the target repository are copied cleanly. Existing scripts modified by the target repository must NEVER be overwritten.
  3. **Bump Script Protection (IMMUTABLE):** NEVER overwrite version bump scripts or manifests (`bump*`, `bump_versions.py`, `bump-version.mjs`, `version.json`). Each repository maintains custom SemVer targets and release logic.
  4. **Memory & Plans Protection (TOTAL ISOLATION):** NEVER modify, sync, or mirror `.ai-memory/memory/` or `.ai-memory/plans/` in target repositories.
  5. **Zero Secrets Leakage:** NEVER synchronize `.env` files, tokens, or credentials across repositories. Secrets reside strictly in `repo-secrets` via `gitmap rs`.

## 12. GitMap High-Speed Search Primacy & Strict Shell Search Ban (Total Ban on rg, ripgrep, Select-String)

- **Search Primacy:** All AI agents, discovery subagents, and worker subagents MUST strictly use GitMap commands (`gitmap aum search`, `gitmap find`, `gitmap search`, `gitmap cat`, `gitmap lf`, `gitmap ff`, `gitmap ffa`) for all code searches, symbol discoveries, and directory listings.
- **TOTAL BAN on Shell Search & File Traversal Tools:** TOTAL BAN on `Get-ChildItem` (all variants: bare `Get-ChildItem`, `-Path`, `-Filter`, `-Recurse`, etc.), `dir`, `ls` in shell, `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, and `findstr`.
- **Ban on Deceptive Tool Summaries:** Explicitly banned from using deceptive tool summaries or action titles (e.g., naming a step 'GitMap find...' or 'Search files' while executing `Get-ChildItem`, `dir`, or `Select-String` underneath). The underlying command executed MUST strictly be a GitMap command.
- **Canonical GitMap Command Suite:**
  - Directory file inventory: `gitmap lf <dir>` (e.g. `gitmap lf src-tauri/src/modules/repo_db`) or `gitmap find "*.*" <dir>`
  - Live streaming code/symbol search: `gitmap aum search "<pattern>" [dir] [-e <.ext>] [-r] [-i]` (e.g. `gitmap aum search "conversation_id" src-tauri/src/modules/repo_db -e .rs`, `gitmap aum search -r "func.*AppError" pkg -e .go`)
  - Substring and glob file finder: `gitmap find "<pattern>" [dir]`, `gitmap ff <name>`, `gitmap ffa <substr>`
  - Zero-disk file streaming: `gitmap cat <filepath>`
  - PowerShell command runner: `gitmap pwsh "<command>"` or `gitmap ps "<command>"`

## 13. Mandatory Completion Invariant ('No Push = Not Done') & Non-Owned Repository Exclusion

- **Mandatory Completion Invariant ('No Push = Not Done'):** If the code is not committed to Git and pushed upstream to GitHub (main/master/tracking branch), the task is strictly considered INCOMPLETE and NOT DONE. Never conclude a turn, claim success, or mark tasks as completed while leaving uncommitted changes or unpushed local commits.
- **Non-Owned Repository Exclusion (TOTAL BAN):** NEVER stage, commit, mutate, or push repositories that are not owned by us, including `omis`, `oh-my-zsh`, `ohmyzsh`, `zsh`, `oh-my-posh`, `dotfiles`, and system runtime packages (`homebrew`, `brew`, `.cargo`, `.rustup`, `.nvm`, `.asdf`, `.pyenv`).

## 14. GitMap Agent Task Engine Primacy & Strict Python SQLite Ban

- **Native GitMap Agent Commands:** All multi-agent task tracking, claiming, status checks, in-flight action logging, and crash autopsy MUST use native compiled `gitmap agent` subcommands (`gitmap agent task init`, `gitmap agent task status`, `gitmap agent task ls`, `gitmap agent task pending`, `gitmap agent subtask add`, `gitmap agent subtask ls`, `gitmap agent subtask claim`, `gitmap agent subtask complete`, `gitmap agent subtask fail`, `gitmap agent log`, `gitmap agent diagnose`, `gitmap agent crashed`, `gitmap agent temp-clear`).
- **TOTAL BAN on Python SQLite Scripts & Inline Queries:** TOTAL BAN on executing `python 03-ai-scripts/46-agent-sqlite-task-manager.py` and TOTAL BAN on raw `python -c "import sqlite3..."` or ad-hoc Python DB queries. All database operations are natively handled by compiled GitMap Go binaries with WAL mode in <1ms.

## 15. GitMap Cognitive Heatmap Telemetry Primacy & Multi-Agent Activity Observation

- **Mandatory Cognitive Action Logging:** Enforce mandatory cognitive action logging across all AI agent sessions:
  - **Step 0 Preflight Inspection:** `gitmap heatmap status` and `gitmap heatmap tables`
  - **Pre-Action Intent Logging:** Before reading or editing files, agents MUST run:
    `gitmap heatmap log --task <task> --action <READ|EDIT|CREATE|DELETE|ANALYZE> --file <path> --state <EXPLORATION|SPECIFICATION|EXECUTION|VERIFICATION> --expected "<goal>"`
  - **Post-Action Closure:** Upon completing an action or receiving subagent findings, run:
    `gitmap heatmap update --id <id> --outcome "<outcome>" --met true`
  - **Phase 3 Verification Audit:** Run `gitmap heatmap tree --task <task>` and `gitmap heatmap status --task <task>` to verify 100% action closure.
  - **Companion Repo-Cache Telemetry:** `gitmap rc heatmap <status|tables|query|view|tree|live|sync>` for multi-AI synchronization.

## 16. Redundant .gitkeep Prohibition

- **Whenever a folder contains any tracked file, any `.gitkeep` file inside that folder MUST be removed immediately.**
- Retain `.gitkeep` ONLY in truly empty placeholder directories that require Git tracking presence.

