# Subtask 02: Author CG-Execute Prompt, Antigravity Skill, and Align Prompts 31 & 32 with V6 Parameter Headers

> **Parent Spec:** [`02-spec/21-app/01-variadic-spread-params-and-multi-repo/01-architecture-spec.md`](../../../../02-spec/21-app/01-variadic-spread-params-and-multi-repo/01-architecture-spec.md)  
> **Status:** `PENDING`  
> **Traceability ID:** `Task-02`  
> **Target Files:**
> - `01-prompts/15-cg-execute/36-variadic-and-spread-parameters.md` (to create)
> - `.agents/skills/cg-variadic-and-spread-parameters/skill.md` (to create)
> - `01-prompts/15-cg-execute/31-cg-execute-in-below-steps.md` (to update)
> - `01-prompts/15-cg-execute/32-cg-follow-other-prompts.md` (to update)
> - `01-prompts/15-cg-execute/readme.md` (to update index table)
> - `01-prompts/readme.md` (to update master prompt catalog)

---

## 1. Subtask Objective & Context

This subtask governs authoring the guideline execution tooling and prompt infrastructure for variadic and spread parameters, while standardizing prompt execution headers across the `15-cg-execute` family to comply with V6 specifications:

1. **Prompt 36 (`36-variadic-and-spread-parameters.md`):** Complete autonomous execution prompt for auditing and refactoring rigid slice/array parameters into variadic/spread parameters across polyglot codebases.
2. **Native Skill (`cg-variadic-and-spread-parameters`):** Corresponding Antigravity skill in `.agents/skills/` for IDE and CLI runtime invocation.
3. **V6 Header Standardization for Prompts 31 & 32:** Align `31-cg-execute-in-below-steps.md` and `32-cg-follow-other-prompts.md` with the V6 parameter header format (`N = 300, A = 2, H = 2, C = 30`, budgets, waves, hyphen GitMap commits).
4. **Registry Updates:** Update prompt catalogs in `01-prompts/15-cg-execute/readme.md` and `01-prompts/readme.md`.

---

## 2. Deliverable Specifications

### 2.1 Deliverable A: Prompt 36 Specification (`36-variadic-and-spread-parameters.md`)

- **Placement:** `01-prompts/15-cg-execute/36-variadic-and-spread-parameters.md`
- **Header Format:** Strict V6 Parameter Header:
  ```text
  N = 300 (Total self-loop steps budget — editable top-header parameter, default: 300)
  A = 2   (MANDATORY number of spawned autonomous subagents running concurrently via invoke_subagent, default: 2)
  H = 2   (Operational hands per agent: dual-task batch capacity & parallel tool dispatch, default: 2)
  C = 30  (Tool calls per worker before it must report, default: 30)

  System Concurrency Capacity = A × H = 2 agents × 2 hands = 4 concurrent subtask operations
  PHASE_1_BUDGET = N / 2   (Steps 1 .. 150: Planning, Parallel Discovery Subagents, Detailed Spec, and Lean Subtask Generation)
  PHASE_2_BUDGET = N / 2   (Steps 151 .. 300: Mandatory Parallel Subagent Execution, Self-Looping, Targeted Quality Linting)
  WAVES = ceil(subtasks / (A x H))
  ```
- **Slash Commands:** `[/goal]`, `[/learn]`, `[/plan]`
- **Mandatory Subagent Spawning Gate:** Mandatory `invoke_subagent` tool calls for both Phase 1 AST discovery and Phase 2 disjoint file refactoring; zero solo execution allowed.
- **3-Phase Lifecycle:**
  - *Phase 0:* Bootstrap `.agents/skills/cg-variadic-and-spread-parameters/skill.md`.
  - *Phase 1A:* Verbatim capture, task extraction, Turn 1 visible task breakdown.
  - *Phase 1B:* GitMap discovery, Violation Ledger in `.ai-memory/plans/pending/`, subtasks in `.ai-memory/plans/subtasks/`.
  - *Phase 2:* Parallel worker execution in 5–8 file micro-batches. Refactor rigid slice signatures to variadic (`...T` in Go, `...items` in TypeScript, `&[T]` in Rust), and refactor call sites to eliminate artificial slice wrappers (`[]string{x}`).
  - *Phase 3:* Subtask consolidation into `.ai-memory/plans/completed/`, atomic push with `gitmap cpf "<module> - <summary>"`.
- **Strict Anti-Hallucination & Quality Rules:**
  - Zero-build and zero-test mandate during routine turns (only targeted linters).
  - Positive booleans only (`is`/`has`), no `== true`, no mixed polarity.
  - Functions <= 8–15 lines, vertical blank lines, `*appfault.AppError`.

### 2.2 Deliverable B: Native Antigravity Skill (`cg-variadic-and-spread-parameters`)

- **Placement:** `.agents/skills/cg-variadic-and-spread-parameters/skill.md`
- **YAML Frontmatter:**
  ```yaml
  ---
  name: cg-variadic-and-spread-parameters
  description: Autonomously audits, refactors, and enforces variadic and spread parameters (...T, ...items, &[T]) across polyglot codebases to eliminate artificial slice wrappers.
  ---
  ```
- **Contents:**
  - Skill description and trigger keywords (`cg-variadic`, `cg-spread-params`, `cg-execute variadic`, `audit variadic`).
  - Core rules: Variadic parameter primacy, spread forwarding, zero-wrapper single-item calls, parameter limits, safe empty/nil handling.
  - Concrete Good vs Bad code snippets for Go, TypeScript, and Rust.
  - Verification guidelines.

### 2.3 Deliverable C: V6 Header Alignment for Prompts 31 & 32

#### Alignment of `31-cg-execute-in-below-steps.md`:
- Replace legacy `N = 200` parameter block with standardized V6 block:
  ```text
  N = 300 (Total self-loop steps budget — editable top-header parameter, default: 300)
  A = 2   (MANDATORY number of spawned autonomous subagents running concurrently via invoke_subagent, default: 2)
  H = 2   (Operational hands per agent: dual-task batch capacity & parallel tool dispatch, default: 2)
  C = 30  (Tool calls per worker before it must report, default: 30)

  System Concurrency Capacity = A × H = 2 agents × 2 hands = 4 concurrent subtask operations
  PHASE_1_BUDGET = N / 2   (Steps 1 .. 150: Planning, Parallel Discovery Subagents, Detailed Spec, and Lean Subtask Generation)
  PHASE_2_BUDGET = N / 2   (Steps 151 .. 300: Mandatory Parallel Subagent Execution, Self-Looping, Targeted Quality Linting)
  WAVES = ceil(subtasks / (A x H))
  ```
- Ensure GitMap atomic commit commands strictly use hyphen separator format:
  `gitmap cpf "<module> - <summary>"` / `gitmap cpb "<module> - <summary>"` (avoiding colon arguments).
- Preserve the Bottom-Instruction Priority Mandate (Below Precedence / Suffix Precedence).

#### Alignment of `32-cg-follow-other-prompts.md`:
- Verify V6 header parameter block consistency (`N = 300, A = 2, H = 2, C = 30`, `PHASE_1_BUDGET`, `PHASE_2_BUDGET`, `WAVES`).
- Ensure all GitMap atomic commit commands use hyphen format (`gitmap cpf "<module> - <summary>"`).
- Ensure skill copy `.agents/skills/cg-follow-other-prompts/skill.md` matches.

### 2.4 Deliverable D: Registries & Catalog Updates

- Update `01-prompts/15-cg-execute/readme.md`: Add row 36 to the Prompts Catalog & Execution Order table.
- Update `01-prompts/readme.md`: Add reference to prompt 36 under section 15.
- Update `.ai-memory/prompts.md`: Register prompt 36 with alias and description.

---

## 3. Execution Step-by-Step Checklist

- [ ] **Step 1:** Author `01-prompts/15-cg-execute/36-variadic-and-spread-parameters.md` with complete V6 header, 3-Phase lifecycle, rules R1–R16, and AST refactoring instructions.
- [ ] **Step 2:** Author `.agents/skills/cg-variadic-and-spread-parameters/skill.md` with valid YAML frontmatter.
- [ ] **Step 3:** Update `01-prompts/15-cg-execute/31-cg-execute-in-below-steps.md` to use V6 parameters (`N = 300, A = 2, H = 2, C = 30`) and hyphen GitMap commit formatting.
- [ ] **Step 4:** Verify and refine `01-prompts/15-cg-execute/32-cg-follow-other-prompts.md` commit commands and header formatting.
- [ ] **Step 5:** Update catalog tables in `01-prompts/15-cg-execute/readme.md`, `01-prompts/readme.md`, and `.ai-memory/prompts.md`.
- [ ] **Step 6:** Run sequence integrity and relative path linters to verify clean zero-drift integration.

---

## 4. Verification & Acceptance Criteria

- **AC-CG-036-A:** Prompt `36-variadic-and-spread-parameters.md` exists, defines `N = 300`, `A = 2`, `H = 2`, `C = 30`, and contains zero broken links.
- **AC-CG-036-B:** Skill `.agents/skills/cg-variadic-and-spread-parameters/skill.md` exists with valid YAML metadata.
- **AC-CG-036-C:** Prompt `31-cg-execute-in-below-steps.md` header reflects `N = 300` and V6 budget calculations.
- **AC-CG-036-D:** Catalog files (`01-prompts/15-cg-execute/readme.md`, `01-prompts/readme.md`) register sequence #36 cleanly.

**Verification commands:**

```bash
python linter-scripts/check-prompts-loaded.py
python linter-scripts/check-relative-paths.py
python 03-ai-scripts/21-sequence-integrity-linter.py 01-prompts/15-cg-execute
```

**Expected Result:** Exit code 0 across all verification linters.
