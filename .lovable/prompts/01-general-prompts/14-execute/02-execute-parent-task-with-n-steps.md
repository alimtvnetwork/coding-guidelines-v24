# Parent Task N-Step Continuous Loop & Multi-Agent Orchestration — Workflow (must follow)

> **Prompt Version:** 2.1.0  
> **Synchronization:** Main Meta-Repo & Connected Workspaces

/goal Autonomously orchestrate and execute the parent task by decomposing it into subtasks and running a continuous N-step self-loop until completion without a single failure.

```text
N = 30
```

N = total self-loop steps budget that the agents will perform.

### Master Task Checklist (Atomic Numbered Steps)

1. [ ] /goal Phase 1 (Planning & Spec Generation, Steps 1..N/2): Spawn exactly 2 planning subagents (max 2 threads each) to scan the codebase and draft `.lovable/plans/pending/XX-<slug>.md`.
2. [ ] /goal Phase 1 (Subtask Decomposition): Decompose the master plan into microscopic, actionable subtasks in `.lovable/plans/subtasks/XX-<slug>/*.md`.
3. [ ] /goal Phase 1 (Strict Folder Bounding): Restrict all planning logs, active locks, and status reports strictly within `.lovable/` (`.lovable/plans/`, `.lovable/temp/active-locks.json`).
4. [ ] /goal Phase 1 (Zero-Stop Transition): Immediately upon completing Phase 1, self-loop and transition directly into Phase 2 execution mode without pausing or stopping.
5. [ ] /goal Phase 2 (Execution & Code Refactoring, Steps N/2+1..N): Spawn exactly 2 execution subagents (max 2 threads each) to execute subtasks on disjoint files in parallel.
6. [ ] /goal Phase 2 (Failure Memory & Error Recovery): If a subagent fails, record the failure log in `.lovable/plans/last-failure.md` and `.lovable/memory/issues/`; subsequent agents MUST read the failure log first to remediate root causes.
7. [ ] /goal Phase 2 (Quality Gate Verification): Execute local linters and `python .lovable/ai-fix-scripts/03-cicd-local-runner.py` ensuring `exit 0` before finishing.
8. [ ] /learn Ingest `.lovable/memory/00-index.md` for project memory index and past learnings.
9. [ ] /learn Ingest `.lovable/strictly-avoid.md` for banned anti-patterns and strict constraints.
10. [ ] /learn Ingest `spec/02-coding-guidelines/` for domain-specific architectural specifications.
11. [ ] /learn Ingest `spec/03-error-manage/` for error handling architectures and AppError.
12. [ ] /learn Ingest `.lovable/coding-guidelines/coding-guidelines.md` for master consolidated coding guidelines.
13. [ ] /goal Create or update agent rules in the repository if missing from agent memory.

```text
PHASE_1_STEPS = N / 2   (Steps 1 .. N/2: 2-Agent Planning & Subtask Generation in .lovable/plans/)
PHASE_2_STEPS = N / 2   (Steps N/2+1 .. N: 2-Agent Parallel Execution, Self-Looping, CI Quality Gates)
```

N, PHASE_1_STEPS, and PHASE_2_STEPS are read-only after initialization. Never modify them mid-execution.

---

## Phase 0: Antigravity Skill Bootstrap (Memory Optimization)

Before executing the tasks below, you must check if this prompt is already installed as a native Antigravity Skill.

1. If `.agents/skills/<slug>/SKILL.md` does not exist in the workspace, you MUST create it now.
2. Extract the core instructions of this prompt and save it into that `SKILL.md` using the standard YAML frontmatter (with `name` and `description`).
3. Once installed, you can rely on progressive disclosure for future runs. Do not keep the entire prompt in your active memory if you don't need it.

---

## 1. 2-Agent Concurrency & Ruthless Orchestration

/goal You are the master orchestrator. If your sub-agents fail, hallucinate, write garbage variables, or go into infinite loops, it is because you are a lazy, incompetent manager.

- **Strict 2-Agent Limit (Max 2 Threads Each):** When dispatching work in Phase 1 (planning) or Phase 2 (execution), you MUST spawn **at most 2 sub-agents concurrently**, with **no more than 2 threads per agent**.
- **Strict Folder Bounding (`.lovable/`):** Subagents are strictly restricted to writing planning files, subtasks, status reports, and logs inside `.lovable/` (`.lovable/plans/`, `.lovable/temp/active-locks.json`, `.lovable/memory/issues/`).
- **Context Diet:** When spawning a subagent, DO NOT paste file contents, memory logs, or the entire plan into its prompt. Give it the absolute minimal instruction (e.g., "Read subtask file `.lovable/plans/subtasks/XX-slug/01-task.md` and execute it"). The subagent MUST read the necessary files itself.
- **Fail Fast & Kill Stalls:** If a sub-agent stalls or provides garbage code, kill it immediately, rollback its dirty working tree, and spawn a new one.

---

## 2. Phase 1: Planning Mode & Subtask Generation FIRST (Steps 1 .. N/2)

Before writing any source code changes, you MUST execute Phase 1:

1. **Scan & Discover:** Spawn 2 planning subagents to deeply scan the codebase for target changes or violations.
2. **Master Spec Generation:** Save the master architectural plan into `.lovable/plans/pending/XX-<slug>.md`.
3. **Task-Specific Rule Set:** Write down 3–5 custom rules or constraints unique to this task inside the spec file.
4. **Subtask Decomposition:** Break down the plan into granular subtask files in `.lovable/plans/subtasks/XX-<slug>/01-task.md`, `02-task.md`, etc.
5. **Strict Relative Git Paths:** All markdown links and file paths in subtasks MUST be strictly relative to the repository root (e.g. `.lovable/spec/...`, `src/...`). Zero absolute paths (`D:\...`, `C:\...`) or `file:///` URIs.
6. **MANDATORY AUTO-LOOP (DO NOT STOP):** As soon as Phase 1 planning completes, the master orchestrator **MUST NOT STOP or ask the user for permission**. It MUST immediately self-loop and transition directly into Phase 2 execution mode.

---

## 3. Phase 2: Execution Mode & Parallel Refactoring (Steps N/2+1 .. N)

1. **Parallel Dispatch:** Spawn 2 execution subagents (max 2 threads each) assigned to disjoint subtasks from `.lovable/plans/subtasks/XX-<slug>/`.
2. **File Locking:** Verify subagents operate on distinct files using `.lovable/temp/active-locks.json`.
3. **Execution & Coding Guidelines:** Subagents refactor code following all coding guidelines (<= 8–15 line functions, single return types, universal `*AppError` wrapping, Unix LF line endings).
4. **Failure Memory & Feedback Loop:** If a subagent fails:
   - Rollback dirty changes and write the failure error log to `.lovable/plans/last-failure.md` and `.lovable/memory/issues/XX-failure.md`.
   - The next subagent spawned MUST read the previous failure log first, record it as a pending memory task, and implement the necessary fix.
5. **Progress & Completion:** Move completed subtasks to `.lovable/plans/completed/` and update `.lovable/plans/index.md`.
6. **Local CI Verification:** Run `python .lovable/ai-fix-scripts/03-cicd-local-runner.py` and ensure all quality gates exit with code 0 (`exit 0`).

---

## 4. AI Fix Scripts Memory (Reusable Tooling)

- [ ] `/goal` **Reuse First:** Scanned and learned `.lovable/ai-fix-scripts/index.md` before writing temporary code.
- [ ] **Strict In-Repository Execution:** All Python scripts executed strictly within the codebase repository root.
- [ ] **Strict .lovable/ Folder Storage:** All helper scripts, local runners, and linters stored in `.lovable/ai-fix-scripts/`.
- [ ] **Native File Manipulator:** Use `python .lovable/ai-fix-scripts/01-file-manipulator.py <command>` for mass file operations.
- [ ] **Go Generate Sync:** If Go constants or enums are modified, run `go generate ./...` in the relevant package and commit generated files.

---

## 5. Non-Negotiable Coding Guidelines Checklist (Auto-Reject on Violation)

/goal You MUST verify every item on this checklist before committing any code. If a subagent violated one of these rules, you must reject their work.

- [ ] Master Guidelines: Fully enforced every file in `spec/02-coding-guidelines/` and `.lovable/coding-guidelines/coding-guidelines.md`.
- [ ] Error Management: Enforced `spec/03-error-manage/` using domain-specific `AppError`, never generic error.
- [ ] Boolean Conventions: All booleans begin with is, has as prefix is only acceptable and nothing else acceptable including but not limited to can, should etc. NO negatives (`!isSuccess` is banned; use `isFail`).
- [ ] Semantic Naming: Zero generic garbage names (`temp`, `data`, `obj`). Behavior-driven unit test names.
- [ ] Multi-Line Arguments (Rule 9a/9b): Signatures and call sites with >2 arguments formatted one argument per line with trailing commas.
- [ ] Line Endings & Encoding: Strictly Unix LF (`\n`) and UTF-8 without BOM.
- [ ] Function Sizing: Functions <= 8 lines preferred (hard cap 15 lines).
- [ ] Strict Relative Git Paths: Zero absolute paths (`D:\...`, `C:\...`) or `file:///` URIs.

---

## 6. Anti-Hallucination & Blast Radius Checklist

- [ ] Echo Back the Spec: Verified Acceptance Criteria from the Spec file verbatim.
- [ ] Pre-Commit Diff Proof: Verified `git status` shows actual modified files before committing.
- [ ] No Placeholder Search: Confirmed zero `TODO` or `\[.*\]` placeholders remain in modified files.
- [ ] Index Sync Deadman Switch: Every new file is explicitly linked in `readme.md` and enqueued in `.lovable/what-to-read.md`.
- [ ] Blast Radius Acknowledgment: Global search across codebase performed to update all callers of modified symbols.
- [ ] Continuous Loop Maintained: Continuous self-loop executed until 100% complete with local CI green.

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
