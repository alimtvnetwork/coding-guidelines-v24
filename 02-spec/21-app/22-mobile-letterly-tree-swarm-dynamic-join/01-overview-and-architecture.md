# Specification: Mobile Letterly Formatter Modernization, Tree Swarm Hierarchy & Dynamic Task Join

## 1. Overview & Objectives

This specification establishes the architectural patterns, contract formats, and operational protocols for:
1. **Professional Mobile Letterly Formatter**:
   - Transitioning mobile prompt formatters from rigid unstructured strings into professionally chained, single-line directives.
   - Chaining standard: `# <task title> - high priority instruction, non-negotiable task (task: <task title>, slug: <task-slug>): ` (hyphen preceding instruction, colon terminating header where instructions commence).
   - Seamlessly blending slash command prefixes (`[/goal]`, `[/learn] from [gitmap]`, `[/plan]`).
   - Embedding the pull-first mandate (`gitmap pull` fallback `git pull`), spec authoring, GitMap search primacy, and strictly relative git paths.
   - Incorporating an extracted list of discrete understandable tasks directly into the single-line output.
   - Enforcing strictly ONE continuous single-line paragraph with ZERO newlines and zero conversational filler across all mobile prompt formatters and companion skills.
2. **Tree-Structured Multi-Agent Swarm in V6**:
   - Transitioning the multi-agent execution model from a flat pool into a coordinated **Tree Hierarchy** (Lead Orchestrator -> Branch Coordinators -> Leaf Workers & Dynamic Assistants).
   - Mandating an up-front **Brain & Comprehensive Roadmap** during Phase 1 planning that maps target files, modification sites, type definitions, and dependency flow before dispatch.
   - Introducing the **Dynamic Agent Join / Collaborative Assistance Protocol**: whenever an agent finishes its assigned subtask while other subtasks remain in flight and capacity (`A`) is not maxed out, newly idle agents dynamically join ongoing subtasks for parallel reading, caller analysis, memory ingestion, and companion verification without file write collisions.
3. **GitMap Agent Swarm Collaboration & Tree Inspection**:
   - Integrating `gitmap agent subtask join` and `gitmap agent task tree` into the GitMap skill suites and V6 execution engine.

---

## 2. Architectural Components

### 2.1 Mobile Letterly Formatter Chaining Architecture

The mobile Letterly output format is engineered for ultra-fast, mobile-initiated autonomous execution. The output must strictly reside on **ONE single continuous line** with zero line breaks.

#### Output Shape Contract:
```text
[/goal](slashCommand;goal) [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) [/plan](slashCommand;plan) # <task title> - high priority instruction, non-negotiable task (task: <task title>, slug: <task-slug>): do gitmap pull <repo path> (fallback git pull) and resolve conflicts before starting; write spec under 02-spec/21-app/<slug>/ and enqueue plan in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/); search codebase exclusively via GitMap; strictly use relative git paths only; ${Input Text Verbatim} - discrete understandable tasks to execute: 1. [task-1], 2. [task-2], 3. [task-3] - strictly use relative git paths only; only add the relative paths, never add the absolute path during your work (this must be respected on the release page and in release notes as well); must follow the skill [execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)
```

#### Key Invariants:
1. **Slash Command Prefix Chain**: Begins with `[/goal](slashCommand;goal) [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) [/plan](slashCommand;plan)` to prime autonomous persistence, GitMap toolchain caching, and planning before execution.
2. **Header Chaining**:
   - Uses `# <task title> - high priority instruction, non-negotiable task (task: <task title>, slug: <task-slug>): `
   - Hyphen separates `<task title>` and `high priority instruction, non-negotiable task`.
   - Colon concludes the parenthetical identification block and marks the immediate beginning of actionable directives.
3. **Core Actionable Directives**:
   - `do gitmap pull <repo path> (fallback git pull) and resolve conflicts before starting;`
   - `write spec under 02-spec/21-app/<slug>/ and enqueue plan in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/);`
   - `search codebase exclusively via GitMap;`
   - `strictly use relative git paths only;`
4. **Verbatim Content & Task List Extraction**:
   - Cleaned `${Input Text Verbatim}` without filler words.
   - Suffix with discrete understandable tasks: `- discrete understandable tasks to execute: 1. [task-1], 2. [task-2]...`
5. **Mandatory Suffix Link**:
   - Antigravity: `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`
   - Cursor: `[execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)`
6. **Mobile CI/CD Fix Variant**:
   - Targets live pipeline diagnosis (`gitmap pe -t`), 4-part RCA, local verification, atomic commit (`gitmap cpf`), and minor release ceremony via `[ci-cd-fix-gitmap-release]`.

---

### 2.2 Tree-Structured Multi-Agent Swarm Hierarchy

In `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` and companion skills, multi-agent orchestration transitions from a flat dispatch pool into a **Tree Swarm Hierarchy**:

```
                  ┌────────────────────────────────────┐
                  │    Lead Orchestrator (Root Node)   │
                  │   Brain & Full Roadmap Custodian   │
                  │  Capacity Controller (A=N, H=2)    │
                  └─────────────────┬──────────────────┘
                                    │
         ┌──────────────────────────┴──────────────────────────┐
         ▼                                                     ▼
┌───────────────────────────────┐             ┌───────────────────────────────┐
│  Branch Coordinator / Worker 01│             │  Branch Coordinator / Worker 02│
│  Subtask A Owner (Owned Files)│             │  Subtask B Owner (Owned Files)│
└───────────────┬───────────────┘             └───────────────┬───────────────┘
                │                                             │
         ┌──────┴──────┐                               ┌──────┴──────┐
         ▼             ▼                               ▼             ▼
   [File Edit]   [Join Assistant]                [File Edit]   [Join Assistant]
   (Disjoint)    (Read/Analyze/Lint)             (Disjoint)    (Read/Analyze/Lint)
```

#### Node Roles & Responsibilities:
1. **Lead Orchestrator (Root Node)**:
   - Evaluates incoming directives, maintains Top-Instruction Priority.
   - Constructs the up-front **Brain & Comprehensive Roadmap** during Phase 1.
   - Manages global state, SQLite task database (`gitmap agent task`), and audit ledger.
   - Enforces capacity `A` (maintains active workers up to `A` concurrent slots).
   - Executes the final atomic GitMap commit and push (`gitmap cpf "<module> - <summary>"`).
2. **Branch Coordinators / Primary Subtask Workers (Internal Nodes)**:
   - Claim subtask ownership atomically (`gitmap agent subtask claim`).
   - Own designated disjoint file boxes (Rule R6: One Owner Per File).
   - Log cognitive heat map traces before reading or editing (`gitmap agent log` / `gitmap heatmap log`).
3. **Leaf Workers & Dynamic Joining Assistants**:
   - Spawned or re-assigned when an agent finishes early to prevent idle capacity.
   - Join in-flight subtasks to accelerate execution through parallel reading, dependency analysis, memory generation, and companion linters.

---

### 2.3 Phase 1: Brain & Comprehensive Roadmap Creation

Before dispatching parallel execution workers in Phase 2, the Lead Orchestrator must author the unified architectural "Brain" within the specification and parent plan:
1. **Target File Identification**: Exhaustive manifest of all files to be read, created, or refactored.
2. **Modification Locations & Call Signatures**: Specific functions, structs, interfaces, and line blocks targeted for refactoring.
3. **Data & Type Flow Contracts**: Type safety boundaries, struct conversions, and error handling wrapping (`*appfault.AppError`).
4. **Disjoint Partitioning Boundaries**: Explicit mapping of files to subtask codes (`Task-01`, `Task-02`, etc.) ensuring zero overlap across workers.

---

### 2.4 Phase 2: Dynamic Agent Join & Collaborative Assistance Protocol

To maximize speed and achieve true continuous parallelism:
1. **Capacity Saturation Invariant**: The total number of active concurrent workers must match `A` whenever uncompleted work remains.
2. **Early Completion Trigger**: If Worker 01 completes its assigned subtask (`gitmap agent subtask complete`) while Worker 02 is still executing, and no unassigned subtasks remain in the queue:
   - Worker 01 does NOT terminate or enter idle sleep.
   - The orchestrator dispatches Worker 01 (or spawns an assistant) to **DYNAMICALLY JOIN** the in-flight subtask.
3. **Assistance Modes (Zero File Collision Mandate - R6)**:
   - **Mode A: Parallel Reading & Blast Radius Analysis**: The joining agent reads related modules, callers, and tests to confirm no breaking changes or type mismatches exist, reporting insights to the primary worker.
   - **Mode B: Cognitive Memory & Spec Documentation**: The joining agent records architectural findings, updates memory registers, and generates companion audit specs.
   - **Mode C: Companion Verification & Lint Auditing**: The joining agent executes targeted linters (`check-relative-paths.py`, `check-forbidden-strings.py`, guideline autofixer check) across related workspaces.
   - **Mode D: Disjoint File Partitioning**: If the ongoing subtask encompasses multiple files, the subtask is subdivided dynamically so the joining agent takes ownership of File 2 while Worker 02 focuses exclusively on File 1.
4. **GitMap Telemetry Tracking**:
   - `gitmap agent subtask join --subtask <id> --agent "Worker <NN>" --role "Assistant: <focus>"`
   - `gitmap agent task tree` to inspect the real-time swarm topology.
