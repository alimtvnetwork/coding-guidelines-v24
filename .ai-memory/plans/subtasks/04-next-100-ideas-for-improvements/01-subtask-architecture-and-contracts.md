# Subtask 01: Architecture Specification & System Contracts Design

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/04-next-100-ideas-for-improvements/01-subtask-architecture-and-contracts.md`  
> **Parent Plan:** `.ai-memory/plans/pending/04-next-100-ideas-for-improvements.md`  
> **Specification File:** `02-spec/21-app/04-next-100-ideas-for-improvements/01-architecture-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Spec Author 01 (Architecture & System Design Lead)  

---

## 1. Objective & Scope

Design and author the authoritative system architecture and technical contracts specification for the **Next 100 Ideas for Improvements** autonomous innovation flywheel:
1. Establish the overarching architectural vision, problem statement, and solution design for `05-next-100-ideas-for-improvements.md` and its companion skills in `.agents/skills/` and `.cursor/skills/`.
2. Define the synergistic relationship and continuous improvement lifecycle across all 5 pillars of the `30-improvements` category:
   - `01-next-100-ui-improvements.md` (UI/UX Perceptual Polish & Tokens)
   - `02-next-100-code-improvements.md` (Codebase Guidelines & Monadic Result Architecture)
   - `03-next-100-whole-system-improvements.md` (Whole-System UX Friction & Operational Flow)
   - `04-next-detail-out-100-improvements.md` (Multi-Step Detailing & Review Docket Assembly)
   - `05-next-100-ideas-for-improvements.md` (Generative Innovation & Feature Superpowers)
3. Formalize the **5-Layer Diagnostic Inspection Engine** using high-speed GitMap streaming commands (`gitmap aum search`, `gitmap find`, `gitmap cat`, `gitmap lf`), completely eliminating slow shell search tools.
4. Establish the **10 Ideation Taxonomy Clusters** (10 ideas each = 100 concrete ideas).
5. Specify the standardized **8-Field Idea Specification Schema** to ensure 100% contract compliance and seamless handoff to Prompt 04.
6. Enforce the **Strict Review-First Principle** and **Zero Immediate Execution Mandate**, ensuring discovered ideas are presented in review dockets rather than executed prematurely.
7. Model the entire architecture using an expressive, end-to-end Mermaid diagram (`flowchart TD`).

---

## 2. Context & Guardrail Invariants

This deliverable strictly obeys all non-negotiable rules defined in `AGENTS.md` and task instructions:
- **Strict Relative Git Paths:** Every file reference, specification link, and subtask pointer uses strictly relative paths starting from the repository root (e.g. `02-spec/21-app/...`, `.ai-memory/plans/...`, `01-prompts/30-improvements/...`). TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`) and `file:///` URIs.
- **Strict Lowercase File Naming:** All created directories and files use strictly lowercase characters without exceptions (`01-architecture-spec.md`, `01-subtask-architecture-and-contracts.md`).
- **Zero Git Command Execution:** As a subagent worker, NEVER run git commands (`git add`, `git commit`, `git push`, `git status`) to avoid index lock collisions.
- **No-Build / No-Test Invariant (Rule R1):** Never trigger compiler or test commands (`go build`, `npm run build`, `go test`, `pytest`). Pure specification and contract authoring only.
- **Boolean Discipline:** Positive boolean conventions (`is` / `has`), implicit evaluation, and total ban on explicit true checks or mixed polarity conditions.
- **Error Handling Standards:** Structured failure metadata mandates `*appfault.AppError` and monadic `Result[T]` wrappers.

---

## 3. Action Checklist

- [x] Analyze existing architecture specifications in `02-spec/21-app/` and prompt catalogs in `01-prompts/30-improvements/readme.md`.
- [x] Design the 5-pillar improvement continuum connecting discovery prompts (01, 02, 03, 05) with the detailing engine (04) and execution swarms.
- [x] Author the comprehensive architecture specification `02-spec/21-app/04-next-100-ideas-for-improvements/01-architecture-spec.md`.
- [x] Include comprehensive problem statement contrasting defensive maintenance with generative ideation.
- [x] Construct an end-to-end Mermaid architecture diagram (`flowchart TD`) visualizing the continuous innovation flywheel.
- [x] Specify the 5-layer diagnostic inspection methodology utilizing GitMap zero-disk streaming.
- [x] Detail the 10 balanced ideation clusters (`IDEA-001` through `IDEA-100`).
- [x] Formalize the 8-Field Idea Specification Schema ensuring deterministic downstream handoffs.
- [x] Define plan persistence contracts in `.ai-memory/plans/pending/` and bounded subtasks in `.ai-memory/plans/subtasks/`.
- [x] Codify the multi-agent swarm hierarchy ($A = 3$, $H = 2$, Capacity = 6) with dynamic joining logic.
- [x] Author this subtask record `01-subtask-architecture-and-contracts.md` with complete execution trace and verification evidence.
- [x] Validate 100% relative path hygiene and strict lowercase naming.

---

## 4. Execution Trace & System Design Evidence

### 4.1 Deliverable Creation: Architecture Specification
The primary deliverable `02-spec/21-app/04-next-100-ideas-for-improvements/01-architecture-spec.md` was authored with comprehensive structural depth:
- **Section 1 (Executive Summary & Problem Statement):** Explains why projects suffer from innovation blindness when trapped in routine defect remediation. Contrasts reactive bug fixing with generative feature discovery grounded in codebase context.
- **Section 2 (Mermaid Architecture Diagram):** A clean `flowchart TD` tracking the lifecycle from user invocation, across the 5-layer survey engine, through the 10 clusters, plan enqueueing, detailing handoff into Prompt 04, human review gate, and downstream execution swarms.
- **Section 3 (Synergistic Architecture of Category 30):** Full comparative matrix of all 5 prompts in `30-improvements`, explaining how raw ideas from 05 feed into 04 for exhaustive AST refactoring blueprints and caller blast-radius maps.
- **Section 4 (5-Layer Diagnostic Inspection Engine):** Concrete technical criteria for scanning project identity, user journeys, developer ergonomics, architectural bottlenecks, and ecosystem bridges using GitMap streaming commands.
- **Section 5 (10 Ideation Clusters):** Exhaustive taxonomy grouping 100 ideas into 10 domains (Core Features, Autonomous Workflows, DX/CLI, UX Polish, Data Architecture, Extensibility/MCP, Observability/Heatmaps, Performance, Resilience, and Spec Governance).
- **Section 6 (8-Field Idea Specification Schema):** Deterministic schema mandating Strategic Value, Grounded Anchors, Architectural Pattern, Interface Contracts (`*appfault.AppError`, `Result[T]`), Blast Radius, Risk Tier, Binary Acceptance Criteria, and Detailing Handoff Directives.
- **Section 7 (Plan Enqueueing & Continuity):** Specifies directory structures for `.ai-memory/plans/pending/XX-next-100-ideas-for-improvements.md` and bounded 5–8 file micro-batch subtasks.
- **Section 8 (Multi-Agent Swarm Concurrency):** Dynamic joining protocol maintaining pinned concurrency capacity of 6 across Lead, Branch Coordinators, and Leaf Workers.
- **Section 9 (Architectural Invariants & Compliance Matrix):** Clear governance table mapping rules to enforcement mechanisms.

### 4.2 Clean Boundary Separation
This specification strictly isolates architectural systems design, leaving component-level catalog listings and script migration mechanics to companion specifications:
- `01-architecture-spec.md` (Spec Author 01 — Owned): Overarching architecture, flywheel mechanics, 5-pillar synergy, 10 clusters, and 8-field schema.
- `02-component-spec.md` (Spec Author 02): Deep prompt component decomposition, frontmatter templates, and IDE skill link formats.
- `03-migration-and-verification-spec.md` (Spec Author 03): Multi-repo fleet distribution, backup safety protocols, and verification check suites.

---

## 5. Verification & Acceptance Criteria

| Check Description | Target Value | Verified Result | Status |
| :--- | :--- | :--- | :---: |
| **Primary Deliverable Path** | `02-spec/21-app/04-next-100-ideas-for-improvements/01-architecture-spec.md` | File exists on disk, complete content authored | `PASS` |
| **Subtask Deliverable Path** | `.ai-memory/plans/subtasks/04-next-100-ideas-for-improvements/01-subtask-architecture-and-contracts.md` | File exists on disk, complete content authored | `PASS` |
| **Strict Relative Paths** | 100% relative Git paths, zero `file:///` URIs, zero absolute paths | Verified across all authored lines | `PASS` |
| **Strict Lowercase Naming** | 100% lowercase filenames and paths | Verified: all paths strictly lowercase | `PASS` |
| **No-Build / No-Test Rule** | Zero builds or tests executed | Verified: zero build/test calls executed | `PASS` |
| **Zero Git Command Rule** | Zero git commands run by worker | Verified: zero git commands run | `PASS` |
| **Mermaid Syntax Fidelity** | Valid `flowchart TD` syntax without unquoted special characters | Verified: valid Mermaid syntax, syntax-safe node labels | `PASS` |
| **10 Ideation Clusters Defined** | Exactly 10 clusters covering 100 ideas | Verified: Section 5 defines all 10 clusters | `PASS` |
| **8-Field Schema Specified** | 8 fields covering contracts, blast radius, ROI | Verified: Section 6 defines full schema | `PASS` |
| **Flywheel Synergy Documented**| Explains 05 discover -> 04 detail -> V6 execute loop | Verified: Section 3 provides detailed analysis | `PASS` |

---

## 6. Conclusion & Handoff

Subtask 01 is fully complete. The architectural foundation, technical contracts, and ideation flywheel specifications for `04-next-100-ideas-for-improvements` are codified and ready for downstream prompt and skill implementation.
