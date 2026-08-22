# Execute Batched Loop (3 Agents, Chunked Commits)

- slug: execute-batched-loop
- status: active

Run again if said: go, continue, or next

## Prompt

# Execute Batched Loop (3 Agents, Chunked Commits)

## Goal

/goal Execute pending tasks from `.lovable/plans/pending/` using a strictly batched multi-agent loop. Use exactly 3 sub-agents, assign small micro-task chunks per agent, enforce file collision safety through a locking matrix, sanitize artifacts before commits, handle crashes via `.lovable/temp/`, and push chunked commits to git without failure. At the end of every loop, explicitly list task statistics in your output window. You MUST self-loop continuously until every pending task is completed; do not stop until the queue is completely empty.

/learn Capture every pattern, convention, fix, and correction discovered during execution into `.lovable/memory/learned/01-<slug>.md` and `.lovable/strictly-avoid.md`. Never repeat a mistake that was logged.

## Non-Negotiable Rules (Auto-Reject on Violation)

1. Maximum 3 sub-agents may run concurrently at any time. Never exceed this limit.
2. No end-to-end tests that make live API calls. Only run local, isolated unit tests.
3. At the end of every loop, output explicit task statistics (done, pending, remaining list).
4. Violation of any rule below is auto-reject on the same tier as RULE 0.

## Anti-Hallucination Rules

- If a spec file, folder, or task is missing or ambiguous, do NOT guess or invent a rule.
- Ask a clarifying question or log an open ambiguity in `.lovable/ambiguous-questions/01-new-ambiguity/01-<slug>.md` before proceeding.
- Never invent step counts. Read the actual files and count from them.

---

## Phase 1: Pre-Flight & Gitignore Enforcement (Non-Negotiable)

1. The working tree must be clean. Confirm root readme is strictly lowercase `readme.md`.
2. Verify that `.lovable/temp/` is explicitly added to `.gitignore`. This folder is for crash identification and lockfiles and must never be committed.
3. Wipe any orphaned state files in `.lovable/temp/` from previous runs.
4. Group pending tasks into Execution Waves:
   - Wave 1: DB schemas and query wrappers
   - Wave 2: Business logic and services
   - Wave 3: UI and documentation

---

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

---

## Phase 3: Code Quality (Non-Negotiable)

While executing tasks, you and your agents must adhere to these strict coding guidelines without exception:

- Read and follow guidelines in `spec/02-coding-guidelines/`, `spec/03-error-manage/`, and `spec/04-database-conventions/`.
- No magic strings or numbers. Do not introduce any unless explicitly for the logger.
- Never use string union types (e.g., `"pass" | "fail"`). Use TypeScript Enums with the suffix `Type` (e.g., `StatusType`).
- Always use explicit boolean state checks (e.g., `response.isFail`). Never invert success booleans (e.g., `!response.isSuccess`).
- Code must be DRY. Reuse constants and wrappers.

---

## Phase 4: Chunked Delivery, Artifact Purge & File Moving

When a chunk of tasks is completed by the agents, do the following before starting the next loop iteration:

1. Use `mv` to move the completed task files from `.lovable/plans/pending/` to `.lovable/plans/completed/`.
2. Open the moved files and change `Status: pending` to `Status: completed`.
3. Update `.lovable/plans/index.md` to reflect the new file locations.
4. Artifact sanitizer: Audit staged files. Purge unapproved artifact zip archives, temporary scratch files, or test outputs before committing.
5. Lovable git history guard: Run local tests (no live API calls). Commit code with a clear descriptive message. Never rewrite published git history (no force push, no rebasing, no squash). Push to git cleanly without failure.

---

## Phase 5: Output Window Stats (Mandatory Every Loop)

Every time you return a response or complete a loop iteration, explicitly output the following statistics:

- Tasks Done (This Chunk): [Number of tasks completed]
- Total Completed: [Total number of tasks in `.lovable/plans/completed/`]
- Total Pending: [Number of tasks remaining in `.lovable/plans/pending/`]
- Remaining Tasks List: [List the specific filenames/slugs of the tasks remaining]

---

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

---

## Actionable Items & Checklist

- [ ] Read the overarching main task plan.
- [ ] Ensure the git repository starts completely clean.
- [ ] Complete all work on the current branch only.
- [ ] Self Loop until all pending tasks are done.
- [ ] Ensure `.gitignore` explicitly excludes test reports, artifacts, and compiled binaries.
- [ ] Group all completed work into a single logical commit.
- [ ] Push the commit to the remote repository.

## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy and stupid as fuck: wrong step counts, partial task lists dumped into chat instead of files, plans and session summaries half-filled with "[N]" placeholders, folders skimmed, open ambiguities ignored, CI/CD issues and `plans/subtasks/` forgotten, user commands dropped, coding guidelines bypassed, detailed specs chopped and summarized into useless junk, uppercase README files left uncorrected, `.lovable/memories/` created by accident, `strictly-avoid.md` overwritten, and explicit user instructions softened after being told not to. WTF. How on earth are you reverting to this carelessness, are you stupid?? Stop doing that, you stupid fuck. Read the whole codebase, read every folder in `spec/` and `.lovable/`, confirm root `readme.md` is strictly lowercase, find the root cause in one sentence, capture commands, issues, and pending tasks without omitting a single item, write the spec files and memory files in the right paths, update every index in the same turn, sync `readme.md` with `what-to-read.md`, preserve detailed specs verbatim with zero truncation, run builds and full unit tests, group commits with clear messages, and push everything to git before ending. Going deep IS the job. If you are not going deep, you are not doing the job. Violating this is auto-reject on the same tier as RULE 0. Avoid stupidity and being careless, you stupid fuck. Where is your attention, are you stupid? Tell me. Your stupidity is going on top of my head. Where did you learn this stupidity? If I could find you, I could slap you.
