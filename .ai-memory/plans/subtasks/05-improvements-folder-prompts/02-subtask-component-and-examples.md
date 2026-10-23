# Subtask 02: Component Specification, Schemas & Concrete Examples

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/05-improvements-folder-prompts/02-subtask-component-and-examples.md`  
> **Parent Plan:** `.ai-memory/plans/pending/05-improvements-folder-prompts.md`  
> **Specification:** `02-spec/21-app/05-improvements-folder-prompts/02-component-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Spec Author 02  

---

## 1. Objective & Scope

Author the comprehensive, grounded component specification, inspection algorithms, dashboard schemas, fast-forwarding protocols, coding guideline enforcement gates, V6 checkpoint enhancements, symmetric companion skills, and operational examples for the **Improvements Folder Prompts** task (`05-improvements-folder-prompts`).

This subtask delivers `02-spec/21-app/05-improvements-folder-prompts/02-component-spec.md` covering:
1. **Prompt 06: `01-prompts/30-improvements/06-check-pending-improvements.md`:**
   - Header formatting (`# High Priority Instruction` & `# Check Pending Improvements`).
   - Slug definition (`## slug: check-pending-improvements`), Version 1.0.0.
   - Slash command invoke syntax (`> Invoke: /check-pending-improvements [scope-or-filter]`).
   - Trigger keywords and aliases (`check-pending-improvements`, `check-improvements`, `pending-improvements`, `improvements-status`, `how many improvements are pending`, `list-pending-improvements`, `improvements-dashboard`).
   - High-speed inspection engine querying `.ai-memory/plans/pending/` and `.ai-memory/plans/subtasks/` using GitMap commands exclusively.
   - Executive Improvement Dashboard Schema: Total pending count, categorical breakdown (UI, Code Base, Whole System, Detailing, Ideas), priority tiers (P0 Critical, P1 High, P2 Normal), completion percentages, and the deterministic Next 10 Queue for handoff to Prompt 07.
2. **Prompt 07: `01-prompts/30-improvements/07-implement-next-10-improvements.md`:**
   - Header formatting (`# High Priority Instruction` & `# Implement Next 10 Improvements`).
   - Slug definition (`## slug: implement-next-10-improvements`), Version 1.0.0.
   - Slash command invoke syntax (`> Invoke: /implement-next-10-improvements [plan-or-subtask-id]`).
   - Trigger keywords and aliases (`implement-next-10-improvements`, `implement-improvements`, `implement-next-10`, `next-10-improvements`, `next`, `run-next-10`, `fast-forward-improvements`).
   - Fast-Forwarding Protocol: Strictly NO-BUILD and NO-TEST execution during implementation turns. Total ban on `go build`, `npm run build`, `pytest`, `go test`. Bounded micro-batch consuming the next 10 unexecuted items from backlog. Repeatable loop ergonomics where user can type `"next"` to advance.
   - End-to-End Testing Separation Gate: Formal architectural separation decoupling implementation from verification, deferring heavy builds and tests to a dedicated downstream phase.
3. **The 8-Pillar Strict Coding Guideline Enforcement Matrix:**
   - Pillar 1: Error Management (zero orphaned errors `_ = err`, `*appfault.AppError` standard).
   - Pillar 2: Monadic Result Envelopes (`Result[T]` or `*appfault.AppError`).
   - Pillar 3: Parameter Structs (functions with >2-3 parameters use dedicated `*Params` structs).
   - Pillar 4: String Constants & Enums (zero raw string literals, PascalCase enums with `Type` suffix).
   - Pillar 5: Function Sizing (functions bounded strictly between 8–15 lines, target 8–10 lines).
   - Pillar 6: File Sizing (files bounded within 150–300 lines, UI components <= 100 lines).
   - Pillar 7: Vertical Line Gaps (mandatory blank line before `if`, after `}`, before `return`, and around multiline structs).
   - Pillar 8: Positive Boolean Standards (affirmative `is*`/`has*` only, zero explicit `== true`, zero mixed polarity).
4. **V6 Parent Task Checkpoint Enhancements:**
   - Detailed specifications of checkpoints to integrate into `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` and companion skill `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`:
     - Fast-forwarding micro-batches in Phase 2.
     - Strict 8-Pillar quality gates during implementation turns.
     - Separated end-to-end testing gate at Phase 3.
5. **Symmetrical Companion Skills Architecture:**
   - Full specification for Antigravity skills (`.agents/skills/check-pending-improvements/skill.md` and `.agents/skills/implement-next-10-improvements/skill.md`) and Cursor skills (`.cursor/skills/check-pending-improvements/skill.md` and `.cursor/skills/implement-next-10-improvements/skill.md`) with platform-compliant link syntaxes (`[skill](file;.agents/skills/<name>)` vs `[skill](file;.cursor/skills/<name>)`).
6. **3 Concrete Production Operational Examples:**
   - Example 1: Real-time Check Pending Output Dashboard.
   - Example 2: Fast-forwarding code transformation with before/after diffs adhering to the 8 pillars.
   - Example 3: Separated End-to-End Verification Gate report proving clean compilation and green tests.
7. **Binary Verification Matrix:** Deterministic verification commands and pass criteria.

---

## 2. Context & Guardrail Invariants

- **Strict Relative Git Paths:** All markdown links, citations, subtask paths, and references must strictly use relative paths starting from the repository root (`02-spec/...`, `.ai-memory/...`, `01-prompts/...`). TOTAL BAN on absolute filesystem paths and file URIs.
- **Strict Lowercase Naming:** All created and modified files, directories, and paths must adhere strictly to lowercase naming conventions.
- **Zero Builds & Tests (Rule R1):** No builds, compilation, or test suites may be executed during specification authoring.
- **No Git Commands:** Subagents must never run `git add`, `git commit`, `git push`, or `git status` to prevent working tree collisions and index locks.
- **Zero Placeholders:** Specifications and subtasks must contain fully resolved file paths, exact code samples, and deterministic criteria. Placeholders like `[N]`, `TODO`, `TBD`, or `...` are strictly prohibited.
- **Top-Instruction Priority Mandate:** Top-instruction precedence verified and active.

---

## 3. Action Checklist

- [x] Ingest master plan in `.ai-memory/plans/pending/05-improvements-folder-prompts.md`.
- [x] Review existing specifications in `02-spec/21-app/04-next-100-ideas-for-improvements/` and `01-prompts/30-improvements/readme.md`.
- [x] Define Prompt 06 component architecture, slug (`## slug: check-pending-improvements`), version 1.0.0, `/check-pending-improvements` invoke syntax, trigger keywords, and execution parameters ($N=50$, $A=3$, $H=2$, Capacity=6).
- [x] Formulate high-speed inspection engine algorithm utilizing GitMap streaming commands (`gitmap lf`, `gitmap aum search`, `gitmap cat`).
- [x] Define Executive Improvement Dashboard schema with categorical breakdown, priority tiers, and Next 10 Queue.
- [x] Define Prompt 07 component architecture, slug (`## slug: implement-next-10-improvements`), version 1.0.0, `/implement-next-10-improvements` invoke syntax, trigger keywords, and execution parameters ($N=200$, $A=3$, $H=2$, Capacity=6).
- [x] Specify Fast-Forwarding Protocol with explicit NO-BUILD and NO-TEST execution during implementation turns.
- [x] Detail the 8-Pillar Strict Coding Guideline Enforcement Matrix (`*appfault.AppError`, `Result[T]`, parameter structs, string enums, 8-15 line functions, 150-300 line files, vertical spacing, positive booleans).
- [x] Define End-to-End Testing Separation Gate decoupling editing velocity from terminal verification.
- [x] Formulate V6 parent task checkpoint enhancements for `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` and companion skills.
- [x] Specify symmetrical companion skills across `.agents/skills/` and `.cursor/skills/` with mandatory IDE skill link syntax.
- [x] Author 3 concrete production-grade examples (dashboard render, fast-forwarding code transformation, and E2E verification report).
- [x] Author and persist `02-spec/21-app/05-improvements-folder-prompts/02-component-spec.md`.
- [x] Author and persist `.ai-memory/plans/subtasks/05-improvements-folder-prompts/02-subtask-component-and-examples.md`.
- [x] Validate relative path hygiene (zero `file:///` URIs, zero absolute paths, strictly lowercase naming).

---

## 4. Execution Trace & Deliverables Summary

### 4.1 Primary Deliverable: `02-spec/21-app/05-improvements-folder-prompts/02-component-spec.md`
- **Section 1: Executive Summary & Ecosystem Alignment:** Establishes the 7-Stage Continuous Improvement Flywheel, framing Prompts 06 and 07 as the inspection, triage, and fast-forward implementation pillars.
- **Section 2: Prompt 06: `06-check-pending-improvements.md` Component Specification:** Complete header block, slug, trigger aliases, execution parameters matrix, inspection engine architecture, and executive dashboard schema.
- **Section 3: Prompt 07: `07-implement-next-10-improvements.md` Component Specification:** Complete header block, slug, trigger aliases, execution parameters matrix, fast-forwarding protocol mechanics (NO-BUILD / NO-TEST invariant), and continuous `"next"` loop.
- **Section 4: The 8-Pillar Strict Coding Guideline Enforcement Matrix:** Exhaustive rules and code patterns for `*appfault.AppError`, monadic Result returns, parameter structs, string constants/enums, function sizing (8–15 lines), file sizing (150–300 lines), vertical line gaps, and implicit positive booleans.
- **Section 5: End-to-End Testing Separation Gate:** Architectural decoupling between fast-forward editing and dedicated downstream verification, detailing the verification suite sequence.
- **Section 6: V6 Parent Task Checkpoint Enhancements:** Exact checkpoint text insertions for Section 7.2.D (Fast-Forwarding Micro-Batches), Section 1.1.E (Quality Gates), and Section 7.3.A (Separated E2E Testing Gate) for `02-execute-parent-task-with-n-steps-v6.md` and companion skills.
- **Section 7: Symmetric Companion Skills Specifications:** Complete YAML frontmatter, headers, directives, and platform-compliant link syntaxes for `.agents/skills/` and `.cursor/skills/`.
- **Section 8: Concrete Production Operational Examples:** 3 detailed, production-grade examples:
  - Example 1: Authentic Continuous Improvement Status Dashboard with Next 10 Queue.
  - Example 2: Fast-forward code refactoring (`pkg/appwriter/writer.go`) with before/after diffs satisfying all 8 pillars.
  - Example 3: Separated End-to-End Verification Gate Report proving green compilation and zero test failures.
- **Section 9: Binary Verification Matrix & Acceptance Criteria:** 10 deterministic verification assertions with commands and pass criteria.

### 4.2 Secondary Deliverable: `.ai-memory/plans/subtasks/05-improvements-folder-prompts/02-subtask-component-and-examples.md`
- Persisted subtask tracking ledger with metadata table, objective & scope, invariants, action checklist, execution trace, and verification criteria.

---

## 5. Binary Verification Criteria

| Check | Target | Command / Proof | Expected Outcome | Result |
|:---|:---|:---|:---|:---:|
| **Spec 02 Exists** | `02-spec/21-app/05-improvements-folder-prompts/02-component-spec.md` | `gitmap lf 02-spec/21-app/05-improvements-folder-prompts` | File present, >600 lines | `PASS` |
| **Subtask 02 Exists** | `.ai-memory/plans/subtasks/05-improvements-folder-prompts/02-subtask-component-and-examples.md` | `gitmap lf .ai-memory/plans/subtasks/05-improvements-folder-prompts` | File present, non-empty | `PASS` |
| **Prompt 06 Schema Defined** | `02-component-spec.md` | `gitmap aum search "slug: check-pending-improvements" 02-component-spec.md` | Documented in Section 2 | `PASS` |
| **Prompt 07 Schema Defined** | `02-component-spec.md` | `gitmap aum search "slug: implement-next-10-improvements" 02-component-spec.md` | Documented in Section 3 | `PASS` |
| **8-Pillar Matrix Defined** | `02-component-spec.md` | `gitmap aum search "8-Pillar" 02-component-spec.md` | Documented in Section 4 | `PASS` |
| **Fast-Forward NO-BUILD Ban** | `02-component-spec.md` | `gitmap aum search "NO-BUILD" 02-component-spec.md` | Total ban documented | `PASS` |
| **Separation Gate Defined** | `02-component-spec.md` | `gitmap aum search "Separation Gate" 02-component-spec.md` | Documented in Section 5 | `PASS` |
| **V6 Checkpoints Specified** | `02-component-spec.md` | `gitmap aum search "02-execute-parent-task-with-n-steps-v6" 02-component-spec.md` | Checkpoints documented in Section 6 | `PASS` |
| **Symmetric Skills Specified** | `02-component-spec.md` | `gitmap aum search ".cursor/skills/implement-next-10-improvements" 02-component-spec.md` | Both environments specified | `PASS` |
| **Production Examples Included** | `02-component-spec.md` | `gitmap aum search "pkg/appwriter/writer.go" 02-component-spec.md` | 3 production examples included | `PASS` |
| **Absolute Path Ban** | Both files | `gitmap aum search "file:/" 02-spec/21-app/05-improvements-folder-prompts/` | 0 raw URIs (only rule citations) | `PASS` |
| **Windows Drive Ban** | Both files | Inspect file paths | 0 raw drive letters | `PASS` |
| **Lowercase Naming** | File paths | Inspect filenames | Strictly lowercase | `PASS` |
| **Zero Placeholders** | Both files | Inspect content | 0 instances of `[N]`, `TODO`, `TBD`, or `...` | `PASS` |
