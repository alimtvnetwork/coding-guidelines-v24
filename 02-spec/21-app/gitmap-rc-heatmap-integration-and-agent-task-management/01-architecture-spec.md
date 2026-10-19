# Specification: GitMap RC Heatmap Integration & Agent Task Management

> **Target:** `02-spec/21-app/gitmap-rc-heatmap-integration-and-agent-task-management/01-architecture-spec.md`  
> **Status:** Authoritative Spec  
> **Version:** 1.0.0  
> **Date:** 2026-10-19  

---

## 1. Executive Summary

This specification defines the architectural enhancements for GitMap RC Heatmap (`gitmap heatmap rc` / `gitmap rc heatmap`), prompt-hash task deduplication and resumption, multi-agent task roadmap linkage, and cross-prompt execution synchronization across the repository.

---

## 2. Core Architectural Pillars

### 2.1 GitMap RC Heatmap Suite & Automatic Initialization
- **Automatic Initialization:** If `repo-cache` (`rc`) or `.ai-memory/heatmap.db` does not exist in the working environment, GitMap automatically provisions the required directories and tables linked to the active mission and workspace folder.
- **Unified Tri-Alias Command Surface:**
  - `gitmap heatmap rc <subcommand>` ↔ `gitmap rc heatmap <subcommand>` ↔ `gitmap heatmap rc-heatmap <subcommand>`.
- **Command Matrix:**
  - `gitmap heatmap rc status`: Cognitive action metrics, total touched files, and active lease counts.
  - `gitmap heatmap rc tables`: Inspect active tables in the RC split-database.
  - `gitmap heatmap rc tree`: Hierarchical tree rendering of cognitive actions grouped by task slug and file path.
  - `gitmap heatmap rc live`: Continuous real-time stream of read/edit events and rationale.
  - `gitmap heatmap rc query "<sql>"`: Non-blocking SQL queries over cognitive telemetry.
  - `gitmap heatmap rc view <table>`: Paginated inspection of specific heatmap tables.
  - `gitmap heatmap rc list-ai`: Enumerate registered AI namespaces and active agent workers.
  - `gitmap heatmap rc commit -m "<msg>"`: Checkpoint and persist heatmap state into `repo-cache`.
  - `gitmap heatmap rc sync`: Bidirectional reconciliation between local `.ai-memory/heatmap.db` and `repo-cache`.

### 2.2 Task Creation, Deduplication & Prompt Hash Verification
- **Deduplication Gate:** When an agent initializes a task via `gitmap agent task init --name "<slug>"`, GitMap detects if the task has already been created.
- **Existence Report:** If the task already exists, GitMap returns structured feedback:
  - `"status": "ALREADY_CREATED"` or `"RESUME_FOUND"`
  - List of active/assigned sub-agents.
  - List of subtasks and their current states (`PENDING`, `CLAIMED`, `DONE`, `FAILED`).
  - List of files already touched or modified (derived from Heatmap telemetry).
- **Prompt Hash Verification:**
  - The agent computes or passes the MD5 hash of the incoming prompt (`--prompt-hash <md5>`).
  - If the prompt hash matches the stored hash: GitMap confirms exact task identity and resumes seamlessly from the first pending subtask.
  - If the prompt hash differs: GitMap flags that the prompt was updated, allowing the orchestrator to ingest new requirements while retaining the prior file touch graph and roadmap progress.

### 2.3 Roadmap-Linked Heatmap for Dynamic Join & Continuation
- **Roadmap Telemetry:** The task roadmap (`.ai-memory/plans/<slug>.md` and subtasks) is indexed alongside Heatmap records.
- **Fast-Path Agent Joining:** When subagents spawn or join (`gitmap agent subtask join`), they inspect the Heatmap to immediately learn:
  1. What files were previously read and why (`READ` events + rationale).
  2. What files were modified and what decisions were made (`EDIT` events + rationale).
  3. Which files currently hold active leases (preventing lock contention).
  4. What subtasks remain in the roadmap.
- This eliminates duplicate reading steps and enables incoming agents to immediately accelerate execution.

### 2.4 Prompt & Skill Synchronization Invariants
- All execution prompts (`01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`, `01-prompts/22-letterly/09-execute-with-verification-letterly.md`, `01-prompts/27-muse-prompts/02-muse-generic.md`, and `cg-execute` families) must explicitly mandate:
  1. Pull first in target repository (`gitmap pull` / `git pull`).
  2. RC Heatmap preflight and status inspection (`gitmap rc heatmap status`).
  3. Prompt hash verification and existence detection upon `gitmap agent task init`.
  4. Logging all read/edit decisions with rationale via `gitmap heatmap log`.
  5. Contention checking via `gitmap agent heatmap` to enforce "One Owner Per File".
  6. Atomic commit with hyphen syntax (`gitmap cpf "<module> - <summary>"`).
