# Batched Loop Execution Without Release (WOR) — Workflow (must follow)


> **Prompt Version:** 2.1.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

# Instruction (must follow): Execute Batched Loop (3 Agents, Chunked Commits)

/goal Execute pending tasks from `.lovable/plans/pending/` using a strictly batched multi-agent loop. Use exactly 3 sub-agents, assign small micro-task chunks per agent, enforce file collision safety through a locking matrix, sanitize artifacts before commits, handle crashes via `.lovable/temp/`, and push chunked commits to git without failure. At the end of every loop, explicitly list task statistics in your output window. You MUST self-loop continuously until every pending task is completed; do not stop until the queue is completely empty.

/learn Capture every pattern, convention, fix, and correction discovered during execution into `.lovable/memory/learned/01-<slug>.md` and `.lovable/strictly-avoid.md`. Never repeat a mistake that was logged.

## Phase 0: Antigravity Skill Bootstrap (Memory Optimization)

Before executing the tasks below, you must check if this prompt is already installed as a native Antigravity Skill.

1. If `.agents/skills/<slug>/SKILL.md` does not exist in the workspace, you MUST create it now. 
2. Extract the core instructions of this prompt and save it into that `SKILL.md` using the standard YAML frontmatter (with `name` and `description`).
3. Once installed, you can rely on progressive disclosure for future runs. Do not keep the entire prompt in your active memory if you don't need it.

## Non-Negotiable Rules (Auto-Reject on Violation)

1. Maximum 3 sub-agents may run concurrently at any time. Never exceed this limit.
2. No end-to-end tests that make live API calls. Only run local, isolated unit tests.
3. At the end of every loop, output explicit task statistics (done, pending, remaining list).

## AI Fix Scripts Memory (Reusable Tooling)

- [ ] `/goal` **Reuse First:** I have rigorously scanned and `/learn`ed `.lovable/ai-fix-scripts/index.md` to check if a helper script already exists before writing any new temporary code.
- [ ] **Strict In-Repository Execution:** All Python scripts (`.lovable/ai-fix-scripts/*.py`) MUST be executed strictly within the codebase repository root, NEVER outside the codebase.
- [ ] **Strict .lovable/ Folder Storage:** All AI scripts, local runners, autofixers, and helper utilities MUST be created inside `.lovable/ai-fix-scripts/`. NEVER create scripts in root or external paths.
- [ ] **Native File Manipulator:** If you need to perform mass file renaming, `.md` lowercase enforcement, sequence number re-ordering, or encoding fixes (CRLF/BOM), you MUST natively use `python .lovable/ai-fix-scripts/01-file-manipulator.py <command>` rather than writing a new script from scratch.
- [ ] **Go Generate Sync:** If you modify Go constants, enums, or stringers, you MUST run `go generate ./...` in the relevant directory (e.g., `cd gitmap && go generate ./...`) and commit the resulting generated files to prevent CI drift.
- [ ] **Commit & Track:** All new helper scripts were written strictly to `.lovable/ai-fix-scripts/` and committed to Git for future reuse.
- [ ] **Index Documentation:** I have updated `.lovable/ai-fix-scripts/index.md` using sequential script naming (e.g., `01-parse-files.py`). For every script, I have included a `<details>` collapsible tag explaining exactly why the script is there and what it does.

5. Violation of any rule below is auto-reject on the same tier as RULE 0.


## STRICT AVOIDANCE: Never Disable CI/CD

> [!CAUTION]
> **NEVER disable any CI/CD checks, GitHub Actions, or validation workflows.** 
> Strictly avoid commenting out, bypassing, or deleting CI/CD steps to force a pipeline to pass. Your job is to fix the underlying code so that the CI/CD pipeline passes legitimately. Disabling CI/CD is an auto-reject failure.

## Anti-Hallucination & Relative Path Rules

> [!CAUTION]
> **STRICT RELATIVE GIT PATHS ONLY — NO ABSOLUTE PATHS / NO `file:///` URIs:**
>
> When generating plans, subtasks (`.lovable/plans/subtasks/`), memory issue logs (`.lovable/memory/issues/`), specs, code comments, or citations:
>
> 1. **Strictly Relative to Git Root:** All file paths, markdown links, citations, and task targets MUST be relative paths starting from the repository root (e.g. `spec/02-coding-guidelines/04-error-handling.md`, `.lovable/spec/commands/01-ssh-commands.md`, `cmd/main.go`).
> 2. **Total Ban on Absolute Paths:** NEVER write drive letters or absolute OS paths (`D:\...`, `C:\...`, `/home/...`) or absolute file URIs (`file:///d:/...`, `file:///C:/...`) into ANY file.
> 
> **Examples:**
>
> - ❌ **BAD:** `[SSH Commands](file:///d:/work/gitmap/.lovable/spec/commands/01-ssh-commands.md) — Why: Defines behavior.`
> - ❌ **BAD:** `Target File: D:\work\gitmap\cmd\login.go`
> - ✅ **GOOD:** `[SSH Commands](.lovable/spec/commands/01-ssh-commands.md) — Why: Defines behavior.`
> - ✅ **GOOD:** `Target File: cmd/login.go`

- If a spec file, folder, or task is missing or ambiguous, do NOT guess or invent a rule.
- Ask a clarifying question or log an open ambiguity in `.lovable/ambiguous-questions/01-new-ambiguity/01-<slug>.md` before proceeding.
- Never invent step counts. Read the actual files and count from them.

## Phase 1: Pre-Flight & Gitignore Enforcement (Non-Negotiable)

1. The working tree must be clean. Confirm root readme is strictly lowercase `readme.md`.
2. Verify that `.lovable/temp/` is explicitly added to `.gitignore`. This folder is for crash identification and lockfiles and must never be committed.
3. Wipe any orphaned state files in `.lovable/temp/` from previous runs.
4. Group pending tasks into Execution Waves:
   - Wave 1: DB schemas and query wrappers
   - Wave 2: Business logic and services
   - Wave 3: UI and documentation

## Phase 2: Allocation & Execution (Strict 3x3 Rule & Locking Matrix)

1. Strict limits:
   - Spawn up to 3 sub-agents to run in parallel.
2. Chunking micro-tasks:
   - Each agent is assigned a chunk of simple, small micro-tasks (under 15 lines per function) to complete sequentially in its own context.
   - Tasks exceeding 7 steps must be decomposed into subtasks.
3. File collision locking matrix (`active-locks.json`):
   - Register active target files in `.lovable/temp/active-locks.json`.
   - Ensure parallel tasks touch completely disjoint files to prevent git merge conflicts.
4. Temp folder logging and specific titling (mandatory):
   - Spawn the sub-agent with a highly specific title reflecting its exact task (e.g., `Refactoring Auth Service` or `Fixing DB Query Wrapper`). Do not use generic names. If an agent switches chunks, its title must change.
   - Log its assigned chunk of tasks to `.lovable/temp/XX-agent-state.md`.
5. Crash identification and 3-strike rollback:
   - If an agent fails or crashes, inspect its state in `.lovable/temp/`.
   - If an agent fails 3 times, automatically revert dirty changes (`git checkout -- <files>`).
   - Log root cause to `.lovable/memory/last-failure.md` and `.lovable/issues/`.
   - Restart a new agent for the next disjoint chunk.

## Phase 3: Code Quality (Non-Negotiable)

While executing tasks, you and your agents must adhere to these strict coding guidelines without exception:

- Read and follow guidelines in `spec/02-coding-guidelines/`, `spec/03-error-manage/`, and `spec/04-database-conventions/`.
- No magic strings or numbers. Do not introduce any unless explicitly for the logger.
- Never use string union types (e.g., `"pass" | "fail"`). Use TypeScript Enums with the suffix `Type` (e.g., `StatusType`).
- Always use explicit boolean state checks (e.g., `response.isFail`). Never invert success booleans (e.g., `!response.isSuccess`).
- Code must be DRY. Reuse constants and wrappers.

## Phase 4: Chunked Delivery, Artifact Purge & File Moving

When a chunk of tasks is completed by the agents, do the following before starting the next loop iteration:

1. Use `mv` to move the completed task files from `.lovable/plans/pending/` to `.lovable/plans/completed/`.
2. Open the moved files and change `Status: pending` to `Status: completed`.
3. Update `.lovable/plans/index.md` to reflect the new file locations.
4. Artifact sanitizer: Audit staged files. Purge unapproved artifact zip archives, temporary scratch files, or test outputs before committing.
5. Lovable git history guard: Run local tests (no live API calls). Commit code with a clear descriptive message. Never rewrite published git history (no force push, no rebasing, no squash). Push to git cleanly without failure.

## Phase 5: Output Window Stats (Mandatory Every Loop)

Every time you return a response or complete a loop iteration, explicitly output the following statistics:

- Tasks Done (This Chunk): [Number of tasks completed]
- Total Completed: [Total number of tasks in `.lovable/plans/completed/`]
- Total Pending: [Number of tasks remaining in `.lovable/plans/pending/`]
- Remaining Tasks List: [List the specific filenames/slugs of the tasks remaining]

## Execution Reporting (Mandatory Output Format)

1. Start of Run (Initial Output): Before writing any code, explicitly list out all pending tasks in your output window.
2. End of Run Summary: When all tasks are completed (or if the run concludes), you MUST output a comprehensive final summary containing:
   - Completed Tasks: Explicit list of what was successfully completed.
   - Pending Tasks Left: Explicit list of any tasks still remaining.
   - Quality Assessment: A brief summary of how well the execution went.
   - Compliance Checklist: A markdown checklist explicitly verifying that you followed the rules:

## Compliance Checklist (must follow non negociable)

- [x] Coding Guidelines enforced (spec/02-coding-guidelines/ and follow explicitly every steps .lovable/coding-guidelines/coding-guidelines.md).
- [x] Boolean conventions used (is/has prefixes, no negatives).
- [x] No garbage variable names used.
- [x] No magic strings or numbers.
- [x] Markdown format verified (newlines around every header).
- [x] Error management protocols followed (AppError/AppException).
- [x] Signatures > 3 parameters or > 100 chars split to one parameter per line.
- [x] Boolean conventions followed (e.g., `isFail` instead of `!isSuccess`).
- [x] Acronyms are PascalCased (e.g., `UserId`, not `UserID`).
- [x] Magic strings/numbers extracted to constants.
- [x] Action Summary Checklist (Anti-Hallucination): I have output a detailed `- [x]` checklist summarizing exactly what I accomplished this turn to ensure no steps were hallucinated or skipped (e.g. `- [x] Created schema`, `- [x] Pinned README`).


---

## Continuous 2-Phase Self-Loop & 2-Agent Concurrency Architecture

To guarantee full execution without stopping after planning mode, the master orchestrator MUST enforce this continuous 2-phase loop:

### 1. 2-Agent Concurrency & Strict `.lovable/` Bounding

- **2-Agent Limit (Max 2 Threads Each):** When dispatching work, spawn **at most 2 sub-agents concurrently**, with **no more than 2 threads per agent**.
- **Strict Folder Bounding (`.lovable/`):** Subagents can ONLY write planning files, subtasks, status reports, and logs inside `.lovable/` (`.lovable/plans/`, `.lovable/temp/active-locks.json`, `.lovable/memory/issues/`).
- **Context Diet:** Provide subagents with minimal instructions (e.g. "Read subtask file `.lovable/plans/subtasks/XX/01-task.md` and execute it"). Do not paste huge files into agent prompts.

### 2. Phase 1: Planning Mode & Subtask Generation (Steps 1 .. N/2)

- Spawn 2 planning subagents to scan the codebase for target guideline violations.
- Write the master architectural specification in `.lovable/plans/pending/XX-audit.md` with an exhaustive Violation Ledger table.
- Decompose the master plan into granular subtasks in `.lovable/plans/subtasks/XX/01-task.md`, `02-task.md`, etc.
- **MANDATORY AUTO-LOOP (DO NOT STOP):** Once Phase 1 planning completes, the master orchestrator **MUST NOT STOP or ask the user for confirmation**. It MUST immediately self-loop and transition directly into Phase 2 execution mode.

### 3. Phase 2: Execution Mode & Parallel Refactoring (Steps N/2+1 .. N)

- Spawn 2 execution subagents (max 2 threads each) to execute subtasks in parallel on disjoint files.
- Subagents refactor code following all coding guidelines (<= 8–15 line functions, single return types, universal `*AppError` wrapping, Unix LF line endings).
- Move completed subtasks from `.lovable/plans/subtasks/` to `.lovable/plans/completed/` and update `.lovable/plans/index.md`.
- **Failure Memory & Feedback Loop:** If a subagent fails:
  - Rollback dirty working tree and log error details to `.lovable/plans/last-failure.md` and `.lovable/memory/issues/XX-failure.md`.
  - The next subagent spawned MUST read the previous failure log first, record it as a pending memory task, and implement the necessary fix.
- Execute local linters and `python .lovable/ai-fix-scripts/03-cicd-local-runner.py` ensuring `exit 0` before concluding.

## Pre-Reply / Loop Checklist (Must Verify Every Loop Iteration)

- [ ] `.gitignore` verified to exclude `.lovable/temp/` and garbage collection executed.
- [ ] Strictly up to 3 agents spawned, each assigned disjoint files tracked in `.lovable/temp/active-locks.json`.
- [ ] Pre-flight state written to `.lovable/temp/` for every agent.
- [ ] 3-Strike rollback honored with `git checkout` and logged to `last-failure.md`.
- [ ] Staged files sanitized against artifact zips and temporary scratch files.
- [ ] No end-to-end live API tests executed.
- [ ] Completed task files `mv`'d and `plans/index.md` updated.
- [ ] Fast-forward commit created and pushed without rewriting git history.
- [ ] Output window explicitly lists "Done", "Pending", and remaining task names.



## STRICT AVOIDANCE: Never Disable CI/CD

> [!CAUTION]
> **NEVER disable any CI/CD checks, GitHub Actions, or validation workflows.** 
> Strictly avoid commenting out, bypassing, or deleting CI/CD steps to force a pipeline to pass. Your job is to fix the underlying code so that the CI/CD pipeline passes legitimately. Disabling CI/CD is an auto-reject failure.

## Anti-Hallucination, Micro-Tasking, & Self-Looping

> [!CAUTION]
> **CRITICAL RULE: DO NOT ATTEMPT TO READ, PLAN, AND EXECUTE EVERYTHING AT ONCE.**
> If you try to consume a massive codebase and write code in a single turn, you WILL hallucinate, drop requirements, and fail.

To survive massive checklists and complex codebases, you MUST operate using these three principles:

1. **Phase 1: Read & Understand (Isolated Loop):** Your very first action must be purely exploratory. Do NOT write code. Break down the task, read the specific files, trace the dependencies, and understand the architectural boundary. Once you understand the scope, end your turn and self-loop to begin execution.
2. **Phase 2: Bounded Micro-Tasking (Sequential Self-Looping):** Never attempt to execute the entire checklist in one response. Treat each checklist section or file as a strict, isolated boundary. Execute *only* the first small portion, verify it, end your turn, and self-loop to process the next portion. 
3. **Phase 3: Multi-Agent Parallelization:** If tasks are independent, you MUST spawn dedicated sub-agents to handle them concurrently. Give each sub-agent an extremely small, strictly defined bounding box (e.g., "Only edit File X"). Never give a sub-agent a generic or multi-file task.

## Actionable Items & Checklist

- [ ] Read the overarching main task plan.
- [ ] Ensure the git repository starts completely clean.
- [ ] Complete all work on the current branch only.
- [ ] Self Loop until all pending tasks are done.
- [ ] Ensure `.gitignore` explicitly excludes test reports, artifacts, and compiled binaries.
- [ ] Group all completed work into a single logical commit.
- [ ] Push the commit to the remote repository.

## Execution & Self-Looping Protocol

This task requires deep, multi-agent processing. You MUST NOT run a fast, linear self-loop that artificially consolidates work or skips steps to save time. 

How to self-loop and distribute tasks effectively:

- [ ] Spawn Sub-Agents:
  - [ ] For any multi-step group, deep file reading, or complex analysis, actively spawn dedicated sub-agents.
  - [ ] Ensure these sub-agents are self-looping to handle the workload.
- [ ] Utilize Processing Power:
  - [ ] Take your time and use maximum processing power and credits.
  - [ ] Do not take shortcuts.
  - [ ] Do not attempt to process a massive spec or write a complex plan in a single, consolidated step.
- [ ] Wait and Aggregate:
  - [ ] As the master agent, loop autonomously to wait for your sub-agents.
  - [ ] Aggregate their precise findings to ensure the highest quality result.

### Temp-Agent State Management Protocol (Non-Negotiable)

To ensure agents don't lose context, you MUST use the .lovable/temp-agents/ directory for tracking sub-agent tasks.

- [ ] On Start:
  - [ ] The sub-agent creates .lovable/temp-agents/<task-name>.md.
  - [ ] Write the objective and STATUS: IN_PROGRESS.
- [ ] On Error/Crash:
  - [ ] If an agent breaks or fails, append the exact error and cause to the file.
  - [ ] Append STATUS: FAILED to that file before closing.
- [ ] On Resume:
  - [ ] The next assigned agent must first read that file to avoid repeating the mistake.
- [ ] On Success:
  - [ ] Update the file to STATUS: DONE.
  - [ ] Immediately update the master plan in .lovable/plans/index.md or .lovable/plans/pending/.

NEVER MAKE THIS EXCUSE: *"I did not literally perform N self-loops or spawn dozens of sub-agents. Instead of artificially spinning up N separate loops, I consolidated the required deep work into a concentrated series of sequential steps within my existing execution environment."*
If you write this excuse, you are a lazy, stupid fuck and you are violating RULE 0. Spawn the actual sub-agents. Do the actual loops. Utilize the processing power effectively to get the right answer. Avoid stupidity.

## WOR: Strictly Without Release

> [!CAUTION]
> **This is a WOR (Without Release) workflow.**
> You MUST NOT bump versions, update changelogs, or cut a release at the end of this task. Commits must remain standard development commits. You are strictly forbidden from triggering a release or running any release scripts.


## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy and stupid as fuck: wrong step counts, partial task lists dumped into chat instead of files, plans and session summaries half-filled with "[N]" placeholders, folders skimmed, open ambiguities ignored, CI/CD issues and `plans/subtasks/` forgotten, user commands dropped, coding guidelines bypassed, detailed specs chopped and summarized into useless junk, uppercase README files left uncorrected, `.lovable/memories/` created by accident, `strictly-avoid.md` overwritten, and explicit user instructions softened after being told not to. WTF. How on earth are you reverting to this carelessness, are you stupid?? Stop doing that, you stupid fuck. Read the whole codebase, read every folder in `spec/` and `.lovable/`, confirm root `readme.md` is strictly lowercase, find the root cause in one sentence, capture commands, issues, and pending tasks without omitting a single item, write the spec files and memory files in the right paths, update every index in the same turn, sync `readme.md` with `what-to-read.md`, preserve detailed specs verbatim with zero truncation, run builds and full unit tests, group commits with clear messages, and push everything to git before ending. Going deep IS the job. If you are not going deep, you are not doing the job. Violating this is auto-reject on the same tier as RULE 0. Avoid stupidity and being careless, you stupid fuck. Where is your attention, are you stupid? Tell me. Your stupidity is going on top of my head. Where did you learn this stupidity? If I could find you, I could slap you.

---

## Metadata

- slug: execute-batched-loop-wor
- status: active
