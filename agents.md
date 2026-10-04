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

- **Strict Relative Git Paths:** All file paths, markdown links, citations, subtask paths (`.ai-memory/plans/subtasks/`), and memory logs MUST be strictly relative paths starting from the git repository root (e.g., `02-spec/02-coding-guidelines/04-error-handling.md`, `.ai-memory/plans/subtasks/01-task.md`, `cmd/main.go`).
- **TOTAL BAN:** NEVER write absolute filesystem paths (e.g., `/absolute/path/to/...`, `C:\Users\...`, `/home/...`) or absolute file URIs (`file:///absolute/path/to/...`, `file:///absolute/path/to/`) inside ANY repository files, plans, specs, comments, or documentation.
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
  - Slash command suggestions such as `[/plan](slashCommand;plan)` belong strictly under `## Additional Instructions` at the bottom of execution prompts, and are strictly prohibited on verification audit prompts.
  - Item 1 of Actionable Items must explicitly mandate writing specs in `02-spec/21-app/<slug>/` and enqueueing plan tasks in `.ai-memory/plans/<slug>.md` (with subtasks in `.ai-memory/plans/subtasks/<slug>/`).
  - Item 2 of Actionable Items must explicitly mandate searching codebase exclusively via GitMap (`gitmap aum search`, `gitmap find`, `gitmap cat`, `gitmap ps`) with a total ban on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`.
  - Item 3 of Actionable Items must explicitly mandate strictly relative Git paths (`02-spec/...`, `.ai-memory/...`, `cmd/...`) with a TOTAL BAN on absolute filesystem paths and `file:///` URIs across all work, code, and release notes.
  - Retrospective AI verification is an audit and verification gate, NOT a planning step. Only `09-execute-with-verification-letterly.md` (and its mirrors) contains retrospective AI verification; standard `execute-n-steps` formatters MUST NOT mandate retrospective AI verification.
  - Execution with verification must embed both `execute-parent-task-with-n-steps-v6` and `ai-verification`.

- **GitMap Commit Formatting:**
  - In all GitMap commit commands (`gitmap cpf`, `gitmap cpb`, `gitmap cpr`), NEVER include colons inside the message argument. GitMap automatically provides `Feature: ` or `Bug: `. Format messages strictly with hyphens:
    - Feature: `gitmap cpf "<module> - <summary>"`
    - Bug Fix: `gitmap cpb "<module> - <summary>"`

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

- **Search Primacy:** All AI agents, discovery subagents, and worker subagents MUST strictly use GitMap commands (`gitmap aum search`, `gitmap find`, `gitmap search`, `gitmap cat`) for all code searches and symbol discoveries.
- **TOTAL BAN on Shell Search Tools:** TOTAL BAN on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem -Recurse`, and `findstr`.




