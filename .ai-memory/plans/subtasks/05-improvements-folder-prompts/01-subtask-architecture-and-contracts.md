# Subtask 01: Architecture Specification & System Contracts Design

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/05-improvements-folder-prompts/01-subtask-architecture-and-contracts.md`  
> **Parent Plan:** `.ai-memory/plans/pending/05-improvements-folder-prompts.md`  
> **Specification File:** `02-spec/21-app/05-improvements-folder-prompts/01-architecture-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Spec Author 01 (Architecture & System Design Lead)  

---

## 1. Objective & Scope

Design and author the authoritative system architecture and technical contracts specification for the **Improvements Folder Prompts: Check Pending & Implement Next 10 Improvements** system:
1. Establish the overarching architectural vision, problem statement, and solution design for the two new prompts in Category 30:
   - `01-prompts/30-improvements/06-check-pending-improvements.md` (Check Pending Improvements)
   - `01-prompts/30-improvements/07-implement-next-10-improvements.md` (Implement Next 10 Improvements)
2. Formulate the comprehensive 7-prompt continuous improvement continuum for `30-improvements`:
   - `01-next-100-ui-improvements.md` (UI Materiality & Token Polish)
   - `02-next-100-code-improvements.md` (Codebase Guidelines & Monadic Typing)
   - `03-next-100-whole-system-improvements.md` (System Friction & Operational Flow)
   - `04-next-detail-out-100-improvements.md` (Multi-Step Detailing & Contracts)
   - `05-next-100-ideas-for-improvements.md` (Generative Innovation & Superpowers)
   - `06-check-pending-improvements.md` (Backlog Telemetry & Readiness Inspection)
   - `07-implement-next-10-improvements.md` (Fast-Forwarding Bounded Batch Implementation)
3. Formalize the **Telemetry Ingestion Engine** of Prompt 06 to scan `.ai-memory/plans/pending/`, `.ai-memory/plans/subtasks/`, and active improvement roadmaps, rendering exact pending counts, category splits, completion percentages, priority rankings (P0–P3), and readiness dockets.
4. Codify the **Fast-Forwarding Protocol (`NO-BUILD`, `NO-TEST`)** of Prompt 07, allowing high-velocity micro-batch execution without intermediate compile or test cycle overhead.
5. Define the **Repeatable "next" Workflow**, enabling operators to type "next" or run the prompt repeatedly to process consecutive batches of 10 items.
6. Codify the **Strict Coding Guidelines Enforcement Matrix**:
   - Error management: Zero orphaned errors (`_ = err` banned), structured `*appfault.AppError`.
   - Monadic Result envelopes: Returns `Result[T]` or `*appfault.AppError`.
   - Parameter structs: Signatures with >2-3 parameters encapsulated into `*Params` DTOs.
   - String constants and Enums: Zero raw magic strings or inline literals.
   - Strict function sizing: Functions bounded between 8 and 15 lines.
   - Strict file sizing: Files bounded between 150 and 300 lines (components <100 lines).
   - Vertical spacing rhythm: Clean line gaps around `if`, after `}`, before `return`, and multiline structs.
   - Positive boolean evaluation: Affirmative `is*`/`has*`, implicit checks, total ban on explicit `== true` and mixed polarity.
7. Detail the **Decoupled End-to-End Verification Gate** to run full builds and test suites separately after batch execution.
8. Specify checkpoint enhancements to be integrated into `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`.
9. Construct an end-to-end Mermaid architecture diagram (`flowchart TD`).

---

## 2. Context & Guardrail Invariants

This deliverable strictly complies with all core rules defined in `AGENTS.md` and task instructions:
- **Strict Relative Git Paths:** Every file reference, specification link, and subtask pointer uses strictly relative paths starting from the repository root (e.g. `02-spec/21-app/...`, `.ai-memory/plans/...`, `01-prompts/30-improvements/...`). TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`) and `file:///` URIs.
- **Strict Lowercase File Naming:** All created directories and files use strictly lowercase characters without exceptions (`01-architecture-spec.md`, `01-subtask-architecture-and-contracts.md`).
- **Zero Git Command Execution:** As a subagent worker, NEVER run git commands (`git add`, `git commit`, `git push`, `git status`) to avoid index lock collisions.
- **No-Build / No-Test Invariant (Rule R1):** Never trigger compiler or test commands (`go build`, `npm run build`, `go test`, `pytest`). Pure specification and contract authoring only.
- **Boolean Discipline:** Positive boolean conventions (`is` / `has`), implicit evaluation, and total ban on explicit true checks or mixed polarity conditions.
- **Error Handling Standards:** Structured failure metadata mandates `*appfault.AppError` and monadic `Result[T]` wrappers.

---

## 3. Action Checklist

- [x] Analyze existing architecture specifications in `02-spec/21-app/` and prompt catalogs in `01-prompts/30-improvements/readme.md`.
- [x] Ingest parent plan `.ai-memory/plans/pending/05-improvements-folder-prompts.md` and verify verbatim user requirements.
- [x] Author the comprehensive architecture specification `02-spec/21-app/05-improvements-folder-prompts/01-architecture-spec.md`.
- [x] Document the architectural problem statement contrasting backlog invisibility and compile-test friction with high-speed telemetry and fast-forwarding batch implementation.
- [x] Construct an end-to-end Mermaid architecture diagram (`flowchart TD`) visualizing the 5-layer continuum of Category 30.
- [x] Detail the complete 7-prompt improvements lifecycle table covering Prompts 01 through 07.
- [x] Define Prompt 06 backlog telemetry scanning engine, metrics calculation, and dashboard schema.
- [x] Formalize Prompt 07 fast-forwarding batch execution engine, "next" repeatable workflow, and decoupled verification handoff.
- [x] Formulate the strict coding guideline enforcement matrix covering 8 core standards (`*appfault.AppError`, `Result[T]`, parameter structs, enums, function sizing, file sizing, vertical spacing, positive booleans).
- [x] Specify V6 parent task prompt checkpoint enhancements (`01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`).
- [x] Specify plan persistence schemas and multi-agent swarm concurrency topologies ($A = 3$, $H = 2$, Capacity = 6).
- [x] Author this subtask record `01-subtask-architecture-and-contracts.md` with complete execution trace and verification evidence.
- [x] Validate 100% relative path hygiene and strict lowercase naming.

---

## 4. Execution Trace & System Design Evidence

### 4.1 Deliverable Creation: Architecture Specification
The primary deliverable `02-spec/21-app/05-improvements-folder-prompts/01-architecture-spec.md` was authored with comprehensive structural depth:
- **Section 1 (Executive Summary & Problem Statement):** Explains how discovery and detailing engines generate backlogs but stumble into visibility and compile-test friction bottlenecks during downstream execution. Contrasts heavy build-test churn with decoupled fast-forwarding micro-batches.
- **Section 2 (Mermaid Architecture Diagram):** A clean `flowchart TD` tracking the lifecycle from discovery prompts (01, 02, 03, 05), through detailing (04), backlog inspection (06), fast-forward batch implementation (07), repeatable "next" loops, decoupled E2E verification, and atomic commit push.
- **Section 3 (Synergistic Architecture of Category 30):** Full comparative matrix of all 7 prompts in `30-improvements`, defining the complete loop from discovery to execution.
- **Section 4 (Architectural Specification of Prompt 06):** Technical specification of `check-pending-improvements`, including telemetry ingestion from `.ai-memory/plans/pending/` and `.ai-memory/plans/subtasks/`, output dashboard schema, and readiness queue generation.
- **Section 5 (Architectural Specification of Prompt 07):** Technical specification of `implement-next-10-improvements`, including fast-forwarding mode (`NO-BUILD`, `NO-TEST`), bounded 10-item queue locking, the 8-rule coding guideline enforcement matrix, and handoff to decoupled E2E testing.
- **Section 6 (V6 Parent Task Checkpoint & Governance Enhancements):** Specifies integration with `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` for fast-forwarding checkpoints and decoupled verification gates.
- **Section 7 (Data Contracts & Plan Persistence Structures):** Specifies schemas for improvement cards and subtask manifests.
- **Section 8 (Multi-Agent Swarm Concurrency):** Dynamic joining protocol maintaining pinned concurrency capacity of 6 across Lead, Branch Coordinators, and Leaf Workers.
- **Section 9 (Architectural Invariants & Compliance Matrix):** Governance table verifying compliance with R1 through R10.

### 4.2 Clean Boundary Separation
This specification strictly isolates architectural systems design, leaving prompt component definitions and migration scripts to companion specifications:
- `01-architecture-spec.md` (Spec Author 01 — Owned): Overarching architecture, Category 30 continuum, Prompt 06 & 07 engines, fast-forwarding protocol, coding guideline matrix, and V6 checkpoint design.
- `02-component-spec.md` (Spec Author 02): Deep prompt component decomposition, prompt frontmatter templates, and IDE skill link formats.
- `03-migration-and-verification-spec.md` (Spec Author 03): Multi-repo fleet distribution, backup safety protocols, and verification check suites.

---

## 5. Verification & Acceptance Criteria

| Check Description | Target Value | Verified Result | Status |
| :--- | :--- | :--- | :---: |
| **Primary Deliverable Path** | `02-spec/21-app/05-improvements-folder-prompts/01-architecture-spec.md` | File exists on disk, complete content authored | `PASS` |
| **Subtask Deliverable Path** | `.ai-memory/plans/subtasks/05-improvements-folder-prompts/01-subtask-architecture-and-contracts.md` | File exists on disk, complete content authored | `PASS` |
| **Strict Relative Paths** | 100% relative Git paths, zero `file:///` URIs, zero absolute paths | Verified across all authored lines | `PASS` |
| **Strict Lowercase Naming** | 100% lowercase filenames and paths | Verified: all paths strictly lowercase | `PASS` |
| **No-Build / No-Test Rule** | Zero builds or tests executed | Verified: zero build/test calls executed | `PASS` |
| **Zero Git Command Rule** | Zero git commands run by worker | Verified: zero git commands run | `PASS` |
| **Mermaid Syntax Fidelity** | Valid `flowchart TD` syntax without unquoted special characters | Verified: valid Mermaid syntax, syntax-safe node labels | `PASS` |
| **Backlog Telemetry Specified** | Prompt 06 scans plans/subtasks, shows exact pending count & categories | Verified: Section 4 defines complete telemetry engine | `PASS` |
| **Fast-Forwarding Specified** | Prompt 07 implements batches of 10 with NO-BUILD / NO-TEST | Verified: Section 5 defines fast-forwarding protocol | `PASS` |
| **Repeatable Workflow Specified**| User can simply type "next" or run prompt again to continue | Verified: Section 5.1 defines repeatable flow | `PASS` |
| **Coding Guidelines Matrix** | 8 strict guidelines (*appfault.AppError, Result[T], params, sizing, spacing, booleans) | Verified: Section 5.3 defines comprehensive matrix | `PASS` |
| **Decoupled E2E Verification** | Separate end-to-end testing phase at conclusion | Verified: Section 5.5 defines decoupled verification | `PASS` |
| **V6 Parent Task Checkpoint** | Checkpoint enhancements for execute-parent-task-with-n-steps-v6 | Verified: Section 6 defines V6 enhancements | `PASS` |

---

## 6. Conclusion & Handoff

Subtask 01 is fully complete. The architectural foundation, technical contracts, and implementation specifications for `05-improvements-folder-prompts` are codified and ready for downstream prompt and skill implementation.
