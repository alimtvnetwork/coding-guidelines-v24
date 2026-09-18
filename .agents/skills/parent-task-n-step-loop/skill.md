---
name: parent-task-n-step-loop
description: Autonomously orchestrate and execute the parent task by decomposing it into subtasks and running a continuous N-step self-loop until completion.
---
# Parent Task N-Step Continuous Loop & Multi-Agent Orchestration

## Phase 1: Planning Mode
1. Verbatim Prompt Capture: Write user prompt into `.ai-memory/plans/pending/xx-<slug>.md`. Extract tasks.
2. Scan & Discover: Spawn 2 planning subagents to scan codebase.
3. Master Spec Generation: Save architectural plan. Add custom rules.
4. Lean Subtask Decomposition: Break down into focused subtasks in `.ai-memory/plans/subtasks/xx-<slug>/`.
5. Strict Relative Paths: Zero absolute paths.
6. Mandatory Auto-Loop: Transition directly to Phase 2 without stopping.

## Phase 2: Execution Mode
1. Parallel Dispatch: Spawn 2 execution subagents (max 2 threads each).
2. Coding Guidelines: Enforce max 15-line functions, single return types, Unix LF.
3. Failure Memory: On crash, rollback and log to `.ai-memory/memory/issues/`.
4. Targeted Linting: Run file-level linters ONLY. NO `06-cicd-local-runner.py`.
5. Atomic Change Tracking: Use `33-test-inventory-generator.py`.
6. Consolidate & Final Commit: Group all completed subtasks into a single `.ai-memory/plans/completed/` file. Single atomic git commit.
