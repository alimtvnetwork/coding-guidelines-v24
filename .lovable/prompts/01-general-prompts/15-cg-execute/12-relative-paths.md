# Relative Git Paths & Absolute Path Elimination — Coding Guideline (must follow)

Trigger Keywords & Aliases: `cg-relative-paths`, `cg-execute relative-paths`, `audit relative paths`, `fix absolute paths`, `fix file paths`, `fix full paths`, `fix paths`, `relative paths audit`


> **Prompt Version:** 2.1.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

```text
N = 200
```

N = total self-loop steps budget that the agents will perform.

/goal Autonomously scan, plan, refactor, and fix all absolute filesystem paths (`D:\...`, `C:\...`, `/home/...`) and `file:///` URIs across the codebase, directly updating markdown documents, specifications, plans, subtasks, citations, and code comments to use strictly relative Git repository paths until 100% green without stopping.

### Master Task Checklist (Atomic Numbered Steps)

1. [ ] /goal Phase 1 (Step A): Deeply scan the target codebase to inventory all architectural violations and anti-patterns.
2. [ ] /goal Phase 1 (Step B): Write the master audit specification in `.lovable/plans/pending/` with an exhaustive Violation Ledger.
3. [ ] /goal Phase 1 (Step C): Decompose the master plan into granular, atomic subtasks in `.lovable/plans/subtasks/`.
4. [ ] /goal Phase 1 (Step D): Verify or create the automated quality linter and register in `.lovable/ai-fix-scripts/index.md`.
5. [ ] /goal Phase 2 (Step A): Open each target file and perform surgical refactoring following authoritative guidelines.
6. [ ] /goal Phase 2 (Step B): Enforce <= 8–15 line function decomposition, single return types, and clean formatting.
7. [ ] /goal Phase 2 (Step C): Execute local linters to verify 0 remaining violations across all modified files.
8. [ ] /goal Phase 2 (Step D): Execute local CI quality gates via `python .lovable/ai-fix-scripts/03-cicd-local-runner.py` with exit code 0 (`exit 0`).
9. [ ] /learn Ingest `.lovable/memory/00-index.md` for project memory index and past learnings.
10. [ ] /learn Ingest `.lovable/strictly-avoid.md` for banned anti-patterns and strict constraints.
11. [ ] /learn Ingest `spec/02-coding-guidelines/00-canonical-size-tier.md` for canonical file and function size tiers.
12. [ ] /learn Ingest `spec/02-coding-guidelines/06-ai-optimization/01-anti-hallucination-rules.md` for hallucination prevention and micro-tasking.
13. [ ] /learn Ingest `spec/02-coding-guidelines/06-ai-optimization/05-citation-requirement.md` for strict relative path citation requirements.
14. [ ] /learn Ingest `spec/02-coding-guidelines/` for domain-specific architectural specifications.
15. [ ] /learn Ingest `.lovable/coding-guidelines/coding-guidelines.md` for master consolidated coding guidelines.
16. [ ] /goal Create or update agent rules in the repository if missing from agent memory.


```text
PHASE_1_STEPS = N / 2   (Steps 1 .. N/2: Scan Codebase for Absolute Paths, Write .lovable/plans/pending/ Spec, Create .lovable/plans/subtasks/, Verify/Create Linter Hook)
PHASE_2_STEPS = N / 2   (Steps N/2+1 .. N: Execute Relative Path Fixer, Verify Relative Links, Run Linter, Verify Local CI Quality Gate)
```

N, PHASE_1_STEPS, and PHASE_2_STEPS are read-only after initialization. Never modify them mid-execution.

---

## Dedicated Section: Strict Relative Git Paths Architecture (Zero Tolerance)

Absolute filesystem paths (e.g., `D:\work\project\...`, `C:\Users\...`, `/home/...`) and absolute file URI schemes (`file:///d:/...`, `file:///C:/...`) are **strictly forbidden** inside committed repository files, specifications, markdown plans, subtask files, code comments, and citations.

### Why Absolute Paths Are Strictly Banned

1. **Portability Destruction:** Absolute paths tie repository files to one specific machine, drive letter, or user directory, instantly breaking links on macOS, Linux, Windows, and CI/CD pipelines.
2. **Reviewer Friction:** Broken absolute file URIs prevent developers and other AI agents from navigating cross-references and citations.
3. **CI/CD Flakiness:** Tests or scripts relying on local machine paths fail deterministically on GitHub Actions and containerized runners.

### Mandatory Relative Git Path Conversion Standards

All markdown links, citations, subtask paths, and file references MUST start from the repository root as clean relative paths without drive letters or `file:///` prefixes:

#### ❌ INVALID (Absolute Path / File URI):

```markdown
- [SSH Commands](file:///d:/work/gitmap/.lovable/spec/commands/01-ssh-commands.md) — Why: Defines required behavior.
- [App Error Docs](file:///d:/work/gitmap/spec/05-coding-guidelines/04-error-handling.md) — Why: Standards for returning results.
- [gitmap/cmd/ssh_login_install_cmd.go](file:///d:/work/gitmap/gitmap/cmd/ssh_login_install_cmd.go) — Why: Target file.
```

#### ✅ VALID (Strict Relative Git Path):

```markdown
- [SSH Commands](.lovable/spec/commands/01-ssh-commands.md) — Why: Defines required behavior.
- [App Error Docs](spec/05-coding-guidelines/04-error-handling.md) — Why: Standards for returning results.
- [gitmap/cmd/ssh_login_install_cmd.go](gitmap/cmd/ssh_login_install_cmd.go) — Why: Target file.
```

---

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

## Strict In-Repository Execution & `.lovable/` Bounding Mandate

> [!IMPORTANT]
> **STRICT IN-REPOSITORY EXECUTION & `.lovable/` STORAGE CONTRACT:**
>
> 1. **In-Codebase Execution Only:** Whenever a Python script (runner, autofixer, linter, test aggregator) is executed or created, it MUST be executed **strictly within the repository root** (current working directory), NEVER outside the codebase or against external arbitrary directories.
> 2. **Strict Folder Bounding (`.lovable/`):** All AI scripts, local runners, autofixers, helper utilities, memory issue logs, and planning files MUST be created inside the `.lovable/` folder:
>    - Python AI Scripts: `.lovable/ai-fix-scripts/` (e.g. `01-file-manipulator.py`, `02-guideline-autofixer.py`, `03-cicd-local-runner.py`, `04-relative-path-fixer.py`).
>    - RCA & Issue Logs: `.lovable/memory/issues/` and `.lovable/cicd-issues/`.
>    - Execution Plans & Subtasks: `.lovable/plans/pending/`, `.lovable/plans/subtasks/`.
>    - Coding Guidelines Mirror: `.lovable/coding-guidelines/`.
> 3. **Worker Pool & Log Aggregation Architecture:** All local runners and test orchestrators must use a concurrent worker pool (2–3 workers via `ThreadPoolExecutor`), announce enqueued tasks upfront, show real-time progress, handle failures gracefully without cancelling sibling workers, and print a consolidated final summary with full stdout/stderr error logs for failed jobs.
> 4. **`force` Keyword Support:** If the user wrote `force`, `force rebuild`, or `force create` on top of the prompt or trigger: **ALWAYS recreate/regenerate the Python runner script from scratch**, regardless of whether the file already exists on disk.
> 5. **No External or Random File Creation:** NEVER write scripts, temporary test scripts, or scratch files to root, `/tmp`, global system paths, or outside the repository boundary.

---

## AI Fix Scripts Memory (Reusable Tooling)

- [ ] `/goal` **Reuse First:** I have rigorously scanned and `/learn`ed `.lovable/ai-fix-scripts/index.md` to check if a helper script already exists before writing any new temporary code.
- [ ] **Strict In-Repository Execution:** All Python scripts (`.lovable/ai-fix-scripts/*.py`) MUST be executed strictly within the codebase repository root, NEVER outside the codebase.
- [ ] **Strict .lovable/ Folder Storage:** All AI scripts, local runners, autofixers, and helper utilities MUST be created inside `.lovable/ai-fix-scripts/`. NEVER create scripts in root or external paths.
- [ ] **Native Relative Path Fixer:** If you need to scan and fix absolute paths or `file:///` URIs, you MUST natively use `python .lovable/ai-fix-scripts/04-relative-path-fixer.py [path]` rather than writing a new script from scratch.
- [ ] **Go Generate Sync:** If you modify Go constants, enums, or stringers, you MUST run `go generate ./...` in the relevant directory (e.g., `cd gitmap && go generate ./...`) and commit the resulting generated files to prevent CI drift.
- [ ] **Commit & Track:** All new helper scripts were written strictly to `.lovable/ai-fix-scripts/` and committed to Git for future reuse.
- [ ] **Index Documentation:** I have updated `.lovable/ai-fix-scripts/index.md` using sequential script naming (e.g., `04-relative-path-fixer.py`). For every script, I have included a `<details>` collapsible tag explaining exactly why the script is there and what it does.

---

## Pre-Reply / Loop Checklist (Must Verify Every Loop Iteration)

- [ ] Git working tree is clean before new code changes.
- [ ] Sub-agents are actively assigned disjoint files verified against `.lovable/temp/active-locks.json`.
- [ ] Completed tasks were `mv`'d to `plans/completed/` and `plans/index.md` was updated.
- [ ] 3-strike rule respected: failed tasks cleanly rolled back and logged to `last-failure.md`.
- [ ] **Strict Relative Git Paths:** All file paths, markdown links, citations, and subtask references in plans, specs, and memory logs are strictly relative to the git repository root. Zero absolute paths (`D:\...`, `C:\...`) or `file:///` URIs.
- [ ] Coding Guidelines & Master Consolidated File: I have fully read, checked, and strictly enforced every file in `spec/02-coding-guidelines/`, as well as the master consolidated coding guideline file at `.lovable/coding-guidelines/coding-guidelines.md`.


1. [ ] /learn and apply as a /goal `.lovable/coding-guidelines/coding-guidelines.md` and also make sure the agent rules are created in the repo to read in the future quickly.

- [ ] `python linter-scripts/check-relative-paths.py` exited with code 0.
- [ ] Local CI runner `python .lovable/ai-fix-scripts/03-cicd-local-runner.py` exited with code 0.

---

## Non-Negotiable Coding Guidelines Checklist (Auto-Reject on Violation)

/goal You MUST verify every item on this checklist before committing any code. If a subagent violated one of these rules, you must reject their work.

- [ ] Strict Relative Git Paths: All file paths, markdown links, citations, and subtask references in plans, specs, and memory logs are strictly relative to the git repository root. Zero absolute paths or `file:///` URIs.
- [ ] Master Guidelines: I have fully read and strictly enforced `spec/02-coding-guidelines/06-ai-optimization/05-citation-requirement.md` and `.lovable/coding-guidelines/coding-guidelines.md`.


1. [ ] /learn the section as a /goal [AI Fix Scripts Memory](#ai-fix-scripts-memory)

- [ ] Action Summary: I have output a detailed `- [x]` checklist summarizing exactly what I accomplished this turn to prove I did not hallucinate.

---

## Mandatory Linter & CI/CD Integration

1. **Linter Script:** `linter-scripts/check-relative-paths.py`
2. **Local Run Command:** `python linter-scripts/check-relative-paths.py`
3. **Autofixer Command:** `python .lovable/ai-fix-scripts/04-relative-path-fixer.py`
4. **CI/CD Integration (`.github/workflows/ci.yml`):**
   ```yaml
   - name: Validate Strict Relative Git Paths
     run: python linter-scripts/check-relative-paths.py
   ```
5. **Runner Registration (`.lovable/ai-fix-scripts/03-cicd-local-runner.py`):**
   ```python
   JOBS = {
       "Relative Path Check": [sys.executable, "linter-scripts/check-relative-paths.py"],
   }
   ```
