# Subtask 02: Component Specification, Schemas & Concrete Examples

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/02-detail-plans-and-improvements-for-review/02-subtask-component-and-examples.md`  
> **Parent Plan:** `.ai-memory/plans/pending/02-detail-plans-and-improvements-for-review.md`  
> **Specification:** `02-spec/21-app/02-detail-plans-and-improvements-for-review/02-component-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Spec Author 02  

---

## 1. Objective & Scope

Author the exhaustive, grounded component specifications, multi-step detailing schemas, review docket templates, and concrete operational plan cards for the **Detail Plans and Improvements for Review** engine.

This subtask delivers `02-spec/21-app/02-detail-plans-and-improvements-for-review/02-component-spec.md` covering:
1. **Executive Summary & Trigger Keywords:** Activation triggers (`detail-plans-and-improvements`, `detail-improvements-for-review`, `Can you please detail out all the plans that you have done and improvements so that it can be improved by any AI easily?`), grounding in the review-first invariant, and decoupled execution handoff.
2. **The 8-Field Detailed Plan Specification Schema:** Complete architectural definitions for:
   - Field 1: Plan/Improvement Identifier & Canonical Slug (`CB-DET-XXX`, `SYS-DET-XXX`, `UI-DET-XXX`)
   - Field 2: High-Level Strategic Intent & Systems Reasoning
   - Field 3: Current State / Friction / Defect Analysis (Verbatim source excerpt, line numbers, RCA)
   - Field 4: Step-by-Step Multi-Stage Execution Blueprint (Discrete subtasks, file bounding boxes, exact target symbols)
   - Field 5: Structural Code Contracts & Type Definitions (`*appfault.AppError`, `result.Result[T]`, `*Params` structs, affirmative booleans)
   - Field 6: Blast Radius, Downstream Callers & Concurrency Boundaries (Disjoint files, lock-free, zero-collision)
   - Field 7: Zero Harm & Regression Safety Verification (No breaking changes, backward compatibility shims, rollbacks)
   - Field 8: Review Docket & Go/No-Go Decision Gate (Interactive checklist presented to user before any execution)
3. **The Review Docket & Human-in-the-Loop Protocol:** Standardized Markdown review docket template, stop-turn protocol awaiting human approval, and decoupled execution handoff to execution swarms.
4. **Concrete Production-Grade Plan Cards:** Three fully elaborated, ready-to-execute cards representing each core domain:
   - `CB-DET-001`: Code Base Refactoring (Monolithic Session Token Validator to structured `AppError`, parameter struct, and monadic result).
   - `SYS-DET-001`: System Architecture (High-concurrency SQLite Split-DB Agent Task Engine with WAL and CAS locking).
   - `UI-DET-001`: User Interface (72px Glassmorphic Sticky Header, 14px safe hover corridor, and WCAG 2.2 AA/AAA tokens).
5. **Symmetrical Antigravity & Cursor Skills Architecture:** Full directory layouts, YAML frontmatter, operational instructions, and platform-specific skill link syntax invariants.
6. **Verification Protocol & Acceptance Criteria:** Deterministic verification matrix and GitMap command validations.

---

## 2. Context & Guardrail Invariants

- **Strict Relative Git Paths:** All markdown links, file paths, citations, and subtask paths must strictly use relative paths (`02-spec/...`, `.ai-memory/...`, `01-prompts/...`). TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`) and `file:///` URIs.
- **Strict Lowercase Naming:** All created and modified files, directories, and paths must adhere to lowercase naming conventions.
- **Zero Builds & Tests (Rule R1):** No builds, compilation, or test suites may be executed during specification authoring.
- **No Git Commands:** Subagents must never run `git add`, `git commit`, `git push`, or `git status` to prevent working tree collisions and index locks.
- **Review-First & No Immediate Execution:** Improvements and plans MUST NOT be executed or implemented right away before reviewing them.
- **Zero Placeholders:** Specifications and subtasks must contain fully resolved file paths, exact code samples, and deterministic criteria. Placeholders like `[N]`, `TODO`, or `...` are strictly prohibited.

---

## 3. Action Checklist

- [x] Analyze master plan requirements in `.ai-memory/plans/pending/02-detail-plans-and-improvements-for-review.md`.
- [x] Define exact trigger keywords and activation phrases for plan detailing.
- [x] Define the complete 8-Field Detailed Plan Specification Schema with architectural reasoning.
- [x] Formulate the interactive Review Docket template with Go/No-Go decision gates.
- [x] Author concrete production-grade plan card for Code Base archetype (`CB-DET-001` Session Validator refactor).
- [x] Author concrete production-grade plan card for System Architecture archetype (`SYS-DET-001` SQLite WAL & CAS subtask engine).
- [x] Author concrete production-grade plan card for UI/UX archetype (`UI-DET-001` 72px sticky header & 14px safe-region navigation).
- [x] Define symmetrical Antigravity (`.agents/skills/`) and Cursor (`.cursor/skills/`) skill specifications with YAML frontmatter.
- [x] Document platform-specific IDE skill link syntax invariants (`[skill](file;.agents/...)` vs `[skill](file;.cursor/...)`).
- [x] Author and persist `02-spec/21-app/02-detail-plans-and-improvements-for-review/02-component-spec.md`.
- [x] Author and persist `.ai-memory/plans/subtasks/02-detail-plans-and-improvements-for-review/02-subtask-component-and-examples.md`.
- [x] Verify relative path hygiene (zero `file:///` URIs, zero absolute paths, strictly lowercase naming).

---

## 4. Execution Trace & Deliverables Summary

### 4.1 Primary Deliverable: `02-spec/21-app/02-detail-plans-and-improvements-for-review/02-component-spec.md`
- **Section 1: Executive Summary & Review-First Foundation:** Defines the intent, systems reasoning, and activation triggers for multi-step plan detailing. Establishes the review-first gate prohibiting code mutation before human sign-off.
- **Section 2: The 8-Field Detailed Plan Specification Schema:**
  - Complete schema definition and Mermaid class diagram.
  - Detailed field requirements: Identifier/Slug, Strategic Intent, Current State Analysis (verbatim code & line numbers), Execution Blueprint (micro-subtasks, bounding boxes), Code Contracts (`*appfault.AppError`, `*Params`, affirmative booleans), Blast Radius & Concurrency Boundaries, Zero Harm & Backward Compatibility Shims, and Review Docket.
- **Section 3: The Review Docket & Human-in-the-Loop Protocol:**
  - Interactive Review Docket Markdown UI template.
  - Mermaid sequence diagram modeling human review, approval gate, and execution handoff.
  - Strict stop-turn invariant where the agent halts and awaits human direction.
- **Section 4: Concrete Production-Grade Detailed Plan Cards:**
  - `CB-DET-001`: Monolithic Session Token Validator refactoring to `result.Result[SessionClaims]`, `*appfault.AppError`, `ValidateTokenParams`, and backward-compatible alias forwarders.
  - `SYS-DET-001`: SQLite Split-DB Agent Task Queue with WAL mode, single writer / multi-reader pooling, and atomic CAS claiming (`WHERE status = 'PENDING'`).
  - `UI-DET-001`: 72px Glassmorphic Sticky Header with `window.scrollY > 12` scroll elevation, 14px safe-region hover corridor, and WCAG 2.2 AAA text contrast (`9.88:1`).
- **Section 5: Symmetrical Antigravity & Cursor Skills Architecture:**
  - Full directory layouts for `.agents/skills/detail-plans-and-improvements-for-review/` and `.cursor/skills/detail-plans-and-improvements-for-review/`.
  - Production-ready `skill.md` definitions with YAML frontmatter.
  - IDE skill link syntax invariants.
- **Section 6: Verification Protocol & Acceptance Criteria:** Comprehensive verification matrix with deterministic GitMap commands.

---

## 5. Binary Verification Criteria

| Check | Target | Command / Proof | Expected Outcome | Result |
|:---|:---|:---|:---|:---:|
| Spec 02 Exists | `02-spec/21-app/02-detail-plans-and-improvements-for-review/02-component-spec.md` | `gitmap lf 02-spec/21-app/02-detail-plans-and-improvements-for-review` | File present, >400 lines | PASS |
| Subtask 02 Exists | `.ai-memory/plans/subtasks/02-detail-plans-and-improvements-for-review/02-subtask-component-and-examples.md` | `gitmap lf .ai-memory/plans/subtasks/02-detail-plans-and-improvements-for-review` | File present, non-empty | PASS |
| Absolute Path Ban | Both files | `gitmap aum search "file:///" 02-spec/21-app/02-detail-plans-and-improvements-for-review` | 0 hits | PASS |
| Absolute Path Ban | Both files | `gitmap aum search "file:///" .ai-memory/plans/subtasks/02-detail-plans-and-improvements-for-review` | 0 hits | PASS |
| Windows Drive Ban | Both files | `gitmap aum search "C:\\" 02-spec/21-app/02-detail-plans-and-improvements-for-review` | 0 hits (except descriptive bans) | PASS |
| Lowercase Naming | File paths | Inspect filenames | Strictly lowercase | PASS |
| Zero Placeholders | Both files | Inspect content | 0 instances of `[N]`, `TODO`, or `...` | PASS |
