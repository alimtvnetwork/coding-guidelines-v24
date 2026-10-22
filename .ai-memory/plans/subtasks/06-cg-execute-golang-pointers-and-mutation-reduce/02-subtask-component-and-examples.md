# Subtask 02: Component Specification, Schemas & Concrete Transformation Rules

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/06-cg-execute-golang-pointers-and-mutation-reduce/02-subtask-component-and-examples.md`  
> **Parent Plan:** `.ai-memory/plans/pending/06-cg-execute-golang-pointers-and-mutation-reduce.md`  
> **Specification:** `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/02-component-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Spec Author 02  

---

## 1. Objective & Scope

Author the authoritative, comprehensive component specification, schemas, parameter models, code transformation rules, V6 prompt Section 4 integration schema, and companion skill definitions for **Task 06: Golang Pointers Reduction and Mutation Reduction** (`06-cg-execute-golang-pointers-and-mutation-reduce`).

This subtask delivers `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/02-component-spec.md` covering:
1. **Header & Parameter Metadata:**
   - Standardized Header (`# High Priority Instruction` & `# Golang Pointers Reduction and Mutation Reduction`).
   - Canonical Slug (`## slug: cg-golang-pointers-and-mutation-reduction`), Version 1.0.0.
   - Slash command invoke syntax (`> Invoke: /golang-pointers-and-mutation-reduction <task>`).
   - Trigger keywords and aliases (`golang-pointers-and-mutation-reduction`, `cg-golang-pointers`, `reduce-pointers-and-mutation`, `golang-value-semantics-and-immutability`, `fix-pointer-mutation`).
   - Parameter matrix ($N = 400$, $A = 3$, $H = 2$, Capacity = 6 concurrent operations, Phase 1 = 100, Phase 2 = 150, Phase 3 = 150).
2. **The 5 Core Pointer Reduction Rules (with detailed Before/After Go code blocks):**
   - **Rule P1:** Small structs ($\le 64$ bytes) passed and returned by value `T`, not `*T`.
   - **Rule P2:** Read-only / Query methods use value receiver `func (s MyStruct) Query() Result[T]`, not `func (s *MyStruct)`.
   - **Rule P3:** Slices (`[]T`) and maps (`map[K]V`) passed by value (never `*[]T` or `*map[K]V`).
   - **Rule P4:** Small parameter structs passed by value `Params`, not `*Params`.
   - **Rule P5:** Pointer Justification Whitelist (`sync.Mutex`/`sync.RWMutex`, stateful mutation methods, tri-state nilability, `*appfault.AppError`).
3. **The 5 Core Mutation Reduction Rules (with detailed Before/After Go code blocks):**
   - **Rule M1:** Single-assignment immutability: declare and initialize in one line; ban subsequent field mutation.
   - **Rule M2:** No post-construction field mutation: ban `obj := New(); obj.Field = val`; use pure constructor `obj := New(val)` or direct struct literal.
   - **Rule M3:** No intermediate variable mutation across branch ladders: ban `var data []byte; switch ... case: data = x; case: data = y`; return immediately or extract pure helper.
   - **Rule M4:** Copy-on-Write / Functional update patterns: methods return a new copy `func (s Config) WithTimeout(t time.Duration) Config`.
   - **Rule M5:** Synchronized Mutex encapsulation for unavoidable mutable caches/buffers: provide paired locked methods (`GetLinesLock()`) and private unlocked methods.
4. **Section 4 Schema Integration into the V6 Execute Prompt Structure:**
   - Formal specification of Section 4 text inserting between Section 3 (GitMap Commands) and Section 5 (Step 0 Preflight).
5. **Symmetric Companion Skill Schemas:**
   - Antigravity Skill (`.agents/skills/cg-golang-pointers-and-mutation-reduction/skill.md`) with `[skill](file;.agents/skills/skill)` link syntax.
   - Cursor IDE Skill (`.cursor/skills/cg-golang-pointers-and-mutation-reduction/skill.md`) with `[skill](file;.cursor/skills/skill)` link syntax.
6. **Concrete Production Transformation Case Studies:**
   - Case Study 1: CLI Configuration & Options Refactoring.
   - Case Study 2: Repository Scanner & File Metadata Extraction.
7. **Verification & Acceptance Criteria:**
   - Deterministic 10-point binary verification matrix.

---

## 2. Guardrail Invariants

- **Strict Relative Git Paths:** All paths, citations, markdown links, and subtask references MUST strictly use relative paths starting from the repository root (`02-spec/...`, `.ai-memory/...`, `01-prompts/...`). TOTAL BAN on absolute filesystem paths and `file:///` URIs.
- **Strict Lowercase Naming:** All created and modified files, directories, and paths MUST adhere strictly to lowercase naming conventions.
- **Zero Builds & Tests (Rule R1):** No builds, compilation, or full test suites may be executed during specification authoring.
- **Worker Git Ban (Rule R7):** Subagents MUST NEVER execute git commands (`git add`, `git commit`, `git push`, `git status`).
- **Zero Placeholders:** Specifications and subtasks must contain fully resolved file paths, exact code samples, and deterministic criteria. Placeholders like `[N]`, `TODO`, `TBD`, or `...` are strictly prohibited.
- **Top-Instruction Priority Mandate:** Top-instruction precedence verified and active.

---

## 3. Action Checklist

- [x] Ingest master plan from `.ai-memory/plans/pending/06-cg-execute-golang-pointers-and-mutation-reduce.md`.
- [x] Ingest architecture spec from `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/01-architecture-spec.md`.
- [x] Define Prompt 39 component metadata: Header, Slug (`cg-golang-pointers-and-mutation-reduction`), Version 1.0.0, invoke syntax, and trigger keywords.
- [x] Define execution parameter matrix ($N = 400$, $A = 3$, $H = 2$, Capacity = 6).
- [x] Formulate the 5 Core Pointer Reduction Rules (P1–P5) with complete before/after Go code blocks.
- [x] Formulate the 5 Core Mutation Reduction Rules (M1–M5) with complete before/after Go code blocks.
- [x] Specify Section 4 Schema integration into the V6 execute prompt structure.
- [x] Detail symmetric companion skills for Antigravity (`.agents/skills/`) and Cursor (`.cursor/skills/`).
- [x] Author two production-grade before/after refactoring case studies.
- [x] Construct binary verification matrix (AC-01 through AC-10).
- [x] Author and persist `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/02-component-spec.md`.
- [x] Author and persist `.ai-memory/plans/subtasks/06-cg-execute-golang-pointers-and-mutation-reduce/02-subtask-component-and-examples.md`.
- [x] Validate relative path hygiene (zero `file:///` URIs, zero absolute paths, strictly lowercase naming).

---

## 4. Execution Trace & Deliverables Summary

### 4.1 Primary Deliverable: `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/02-component-spec.md`
- **Section 1: Executive Summary & Component Scope:** Defines the mission of Prompt 39 within Category 15.
- **Section 2: Canonical Header, Slug & Parameter Metadata:** Defines the exact header block, slug, trigger aliases, and the V6 parameter matrix ($N=400, A=3, H=2, \text{Capacity}=6$).
- **Section 3: The 5 Core Pointer Reduction Rules:**
  - Rule P1: Value Semantics for Small Structs ($\le 64$ Bytes).
  - Rule P2: Value Receivers on Read-Only / Query Methods.
  - Rule P3: Pure Slice (`[]T`) and Map (`map[K]V`) Parameter Passing.
  - Rule P4: Value-Based Small Parameter Structs (`Params` vs `*Params`).
  - Rule P5: Pointer Justification Whitelist (`sync.Mutex`, Stateful Mutators, Tri-State Nilability, `*appfault.AppError`).
- **Section 4: The 5 Core Mutation Reduction Rules:**
  - Rule M1: Single-Assignment Immutability (`:=` once).
  - Rule M2: Pure Construction & Zero Post-Construction Field Mutation (`obj := New(val)` or struct literal).
  - Rule M3: Zero Intermediate Variable Mutation Across Branch Ladders (direct returns or pure helpers).
  - Rule M4: Copy-on-Write / Functional Transformation Patterns (`WithXxx()`).
  - Rule M5: Synchronized Mutex Encapsulation for Unavoidable Mutable State (`GetLinesLock()`).
- **Section 5: V6 Execute Prompt Structure: Section 4 Integration Schema:** Detailed markdown text for Section 4 insertion between Section 3 and Section 5 in `39-golang-pointers-and-mutation-reduction.md`.
- **Section 6: Symmetric Companion Skill Schemas:** Exact YAML frontmatter and IDE link formatting for Antigravity (`.agents/skills/`) and Cursor (`.cursor/skills/`).
- **Section 7: Concrete Production Transformation Case Studies:**
  - Case Study 1: CLI Configuration & Options Refactoring (`ClientOptions`).
  - Case Study 2: Repository Scanner & File Metadata Extraction (`FileMeta`).
- **Section 8: Verification & Acceptance Criteria:** 10 deterministic verification assertions.

### 4.2 Secondary Deliverable: `.ai-memory/plans/subtasks/06-cg-execute-golang-pointers-and-mutation-reduce/02-subtask-component-and-examples.md`
- Tracks subtask metadata, invariants, checklist completion, execution trace, and verification criteria for downstream workers.

---

## 5. Verification & Quality Gate

```text
Check 1: File existence of 02-component-spec.md -> PASS
Check 2: File existence of 02-subtask-component-and-examples.md -> PASS
Check 3: Total ban on file:/// URIs and absolute paths -> PASS (0 hits)
Check 4: Strictly lowercase filenames and paths -> PASS
Check 5: Full coverage of 5 Pointer Rules and 5 Mutation Rules -> PASS
Check 6: Section 4 V6 prompt schema integration -> PASS
Check 7: Companion skills schemas for Antigravity and Cursor -> PASS
```
