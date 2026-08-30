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

- **Strict Relative Git Paths:** All file paths, markdown links, citations, subtask paths (`.lovable/plans/subtasks/`), and memory logs MUST be strictly relative paths starting from the git repository root (e.g., `spec/02-coding-guidelines/04-error-handling.md`, `.lovable/plans/subtasks/01-task.md`, `cmd/main.go`).
- **TOTAL BAN:** NEVER write absolute filesystem paths (e.g., `D:\work\...`, `C:\Users\...`, `/home/...`) or absolute file URIs (`file:///d:/...`, `file:///C:/...`) inside ANY repository files, plans, specs, comments, or documentation.
  - ❌ **BAD:** `[SSH Commands](file:///d:/.../.lovable/spec/commands/01-ssh-commands.md)`
  - ❌ **BAD:** `Target File: D:\...\cmd\login.go`
  - ✅ **GOOD:** `[SSH Commands](.lovable/spec/commands/01-ssh-commands.md)`
  - ✅ **GOOD:** `Target File: cmd/login.go`

