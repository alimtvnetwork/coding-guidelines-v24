# Subtask 01: Architecture Specification & System Contracts Design

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/06-cg-execute-golang-pointers-and-mutation-reduce/01-subtask-architecture-and-contracts.md`  
> **Parent Plan:** `.ai-memory/plans/pending/06-cg-execute-golang-pointers-and-mutation-reduce.md`  
> **Specification File:** `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/01-architecture-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Spec Author 01 (Architecture & System Design Lead)  

---

## 1. Objective & Scope

Design and author the authoritative system architecture and technical contracts specification for adding **Prompt 39** (`39-golang-pointers-and-mutation-reduction.md`) to Category 15 (`15-cg-execute`) and companion skills:
1. Establish the overarching architectural vision, problem statement, and dual-pillar solution design for Golang pointer reduction and mutation avoidance across the repository fleet.
2. Formalize the dual architectural challenges:
   - **Excessive Pointer Usage (`*T`):** Heap escapes, garbage collection pressure, cache line misses, pointer chasing, and nil-pointer dereference hazards.
   - **In-Place Mutation:** Branch state corruption, non-deterministic dataflow, partial object initialization, and concurrency race conditions.
3. Codify the combined dual solution:
   - **Pillar 1 (Pointer Reduction & Value Semantics):** 64-byte stack threshold, value receivers for queries/read-only operations, direct slice/map value passing, and a strict 5-item pointer justification whitelist.
   - **Pillar 2 (Mutation Avoidance & Immutability):** Single-assignment immutability, pure constructors, zero intermediate mutation in `if`/`switch` branch ladders, copy-on-write functional transformations, and mutex-synchronized thread-safe caches.
4. Define the execution pattern architecture adhering strictly to `execute-parent-task-with-n-steps-v6`:
   - Concurrency parameters: $N = 400$, $A = 3$, $H = 2$, Capacity = 6.
   - Tree Swarm Hierarchy: Root Orchestrator -> Branch Coordinators -> Leaf Workers & Dynamic Assistants.
   - Fast-forwarding micro-batches (waves of 5–8 files) with strict NO-BUILD and NO-TEST during active execution turns.
   - Decoupled terminal verification gate for compilation, race detection, and linters.
   - Cognitive Heatmap Telemetry (`gitmap heatmap log`, `gitmap heatmap update`, `gitmap heatmap status`).
5. Render a comprehensive, syntactically safe Mermaid architecture diagram (`flowchart TD`) visualizing the end-to-end execution lifecycle.

---

## 2. Context & Guardrail Invariants

This subtask strictly complies with all core invariants established in `AGENTS.md` and subagent operational rules:
- **Strict Relative Git Paths:** All file paths, markdown links, citations, and subtask references use strictly relative paths starting from the git repository root (e.g. `02-spec/21-app/...`, `.ai-memory/plans/...`, `01-prompts/15-cg-execute/...`). TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`) and `file:///` URIs.
- **Strict Lowercase File Naming:** All paths, directories, and files use strictly lowercase characters without exception (`01-architecture-spec.md`, `01-subtask-architecture-and-contracts.md`).
- **Zero Git Command Rule for Workers:** Subagents never execute git commands (`git add`, `git commit`, `git push`, `git status`). All git operations are reserved exclusively for the lead orchestrator.
- **No-Build / No-Test Rule (Rule R1):** Never trigger compiler or test commands (`go build`, `go test`, `golangci-lint`) during specification and plan authoring.
- **Error Handling Standards:** All Go failure metadata structures adhere to `*appfault.AppError` and monadic `Result[T]` conventions.
- **Boolean Discipline:** Positive boolean conventions (`is`, `has`), implicit evaluation, and total ban on explicit true comparisons or mixed polarity expressions.

---

## 3. Action Checklist

- [x] Review parent execution plan `.ai-memory/plans/pending/06-cg-execute-golang-pointers-and-mutation-reduce.md`.
- [x] Inspect existing Category 15 prompts (`33-branch-immutability-and-clean-construction.md`, `34-golang-pointer-reduction-and-value-semantics.md`).
- [x] Author the primary architecture specification `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/01-architecture-spec.md`.
- [x] Document the architectural problem: heap escapes, pointer chasing, GC thrashing, nil panics, and in-place mutation races.
- [x] Formalize Pillar 1: Pointer Reduction & Value Semantics (64-byte threshold, value receivers, direct slice/map passing, pointer whitelist).
- [x] Formalize Pillar 2: Mutation Avoidance & Immutability (single-assignment, pure constructors, zero branch mutation, copy-on-write, mutex caches).
- [x] Codify the V6 Tree Swarm execution pattern ($N = 400$, $A = 3$, $H = 2$, Capacity = 6, fast-forwarding micro-batches, decoupled verification gate).
- [x] Construct an end-to-end Mermaid architecture diagram (`flowchart TD`) with quote-safe node labels.
- [x] Log Cognitive Heatmap actions in SQLite via `gitmap heatmap log` and `gitmap heatmap update`.
- [x] Author this subtask record `01-subtask-architecture-and-contracts.md` with complete metadata and verification criteria.
- [x] Verify 100% relative path hygiene and strict lowercase naming.

---

## 4. Execution Trace & System Design Evidence

### 4.1 Authoring Deliverable 1: Architecture Specification
The primary deliverable `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/01-architecture-spec.md` was authored with full structural rigor:
- **Section 1 (Executive Summary & Problem Statement):** Establishes the operational rationale for Prompt 39 within Category 15. Details the technical liabilities of excessive `*T` pointers (escape analysis failures, GC overhead, cache invalidation, tri-state nilability) and in-place variable mutation (data races, branch accumulation bugs, incomplete initialization).
- **Section 2 (Mermaid Architecture Diagram):** A clean `flowchart TD` diagram mapping the end-to-end pipeline from operator trigger, across the dual-pillar core, through the V6 Tree Swarm concurrent engine, fast-forwarding micro-batches, SQLite heatmap telemetry, and the decoupled terminal verification gate.
- **Section 3 (Pillar 1: Pointer Reduction & Value Semantics):** Establishes the 64-byte struct stack copying rule, value receivers on queries/read-only methods, direct passing of slice and map headers, and the 5-item pointer justification whitelist (Concurrency Primitives, State Mutation Receivers, Tri-State Nilability, AppError Envelopes, Large Memory Structs).
- **Section 4 (Pillar 2: Mutation Avoidance & Immutability):** Specifies single-assignment immutability, fully-formed pure constructors, elimination of intermediate branch re-assignment, copy-on-write functional transformations, and thread-safe mutex guards for unavoidable state.
- **Section 5 (V6 Tree Swarm Execution Pattern):** Codifies $N = 400$, $A = 3$, $H = 2$, Capacity = 6, Tree Swarm Hierarchy, 5–8 file fast-forwarding micro-batches with strict NO-BUILD and NO-TEST during turns, decoupled terminal verification gate, and Cognitive Heatmap telemetry.
- **Section 6 (Synergistic Continuum in Category 15):** Situates Prompt 39 alongside Prompts 33 and 34, showing how Prompt 39 serves as the comprehensive Go master prompt.
- **Section 7 (Non-Negotiable Invariants):** Summarizes INV-01 through INV-07 governing path hygiene, naming, concurrency, and completion standards.

### 4.2 Cognitive Heatmap Telemetry Evidence
Actions were logged to the native SQLite heatmap engine:
- Action Log ID 63: `CREATE` for `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/01-architecture-spec.md` (`SPECIFICATION` state) -> Status: `MET`.
- Action Log ID 64: `CREATE` for `.ai-memory/plans/subtasks/06-cg-execute-golang-pointers-and-mutation-reduce/01-subtask-architecture-and-contracts.md` (`SPECIFICATION` state) -> Active.

---

## 5. Verification & Acceptance Criteria

| Check Description | Target Value | Verified Result | Status |
| :--- | :--- | :--- | :---: |
| **Primary Deliverable Path** | `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/01-architecture-spec.md` | File exists on disk, fully authored | `PASS` |
| **Subtask Deliverable Path** | `.ai-memory/plans/subtasks/06-cg-execute-golang-pointers-and-mutation-reduce/01-subtask-architecture-and-contracts.md` | File exists on disk, fully authored | `PASS` |
| **Strict Relative Paths** | 100% relative Git paths, zero `file:///` URIs, zero absolute paths | Verified across all lines in both deliverables | `PASS` |
| **Strict Lowercase Naming** | 100% lowercase filenames and directory paths | Verified: all paths strictly lowercase | `PASS` |
| **No-Build / No-Test Rule** | Zero builds or tests executed during authoring | Verified: zero build/test calls executed | `PASS` |
| **Zero Git Command Rule** | Zero git commands run by worker subagent | Verified: zero git commands run | `PASS` |
| **Mermaid Syntax Fidelity** | Valid `flowchart TD` syntax without unquoted special characters | Verified: valid Mermaid syntax with safe labels | `PASS` |
| **Dual-Pillar Framework** | Both Pillar 1 (Pointers) and Pillar 2 (Mutation) fully specified | Verified: Sections 3 & 4 provide detailed specs | `PASS` |
| **V6 Tree Swarm Codified** | $N=400$, $A=3$, $H=2$, Capacity=6, micro-batches, decoupled gate | Verified: Section 5 provides complete parameters | `PASS` |
| **Heatmap Action Logging** | Pre-action intent logged and updated in SQLite | Verified: Action IDs 63 & 64 logged | `PASS` |

---

## 6. Conclusion & Handoff

Subtask 01 is fully complete. The architectural foundation, systems rationale, dual-pillar rules, and V6 Tree Swarm execution contracts for `06-cg-execute-golang-pointers-and-mutation-reduce` are codified and ready for Spec Author 02 and Spec Author 03 to proceed with component decomposition, migration planning, and fleet synchronization.
