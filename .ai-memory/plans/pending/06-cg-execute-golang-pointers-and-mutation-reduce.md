# Execution Plan: 06-cg-execute-golang-pointers-and-mutation-reduce

## User Request (Verbatim)
```text
in cg execute add a prompt for golang pointers reduce and mutation reduce following patterns of excute in n steps v6 prompt, can you please do it, sync with all repos
```

## Plan Metadata
- **Slug:** `06-cg-execute-golang-pointers-and-mutation-reduce`
- **Category:** `15-cg-execute`
- **Prompt Sequence:** `39` (`01-prompts/15-cg-execute/39-golang-pointers-and-mutation-reduction.md`)
- **Skill Name:** `cg-golang-pointers-and-mutation-reduction`
- **Pattern:** `execute-parent-task-with-n-steps-v6` ($N = 400$, $A = 3$, $H = 2$, Capacity = 6)
- **Status:** In-repo work is done (prompt 39, both skills, spec 06, catalog row). The 43-repository sync script also bumps and tags every connected repo, so it was not started from this pass.
- **Spec Directory:** `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/`
- **Subtasks Directory:** `.ai-memory/plans/subtasks/06-cg-execute-golang-pointers-and-mutation-reduce/`

---

## Strategic Intent & Architectural Continuum
Author and deploy canonical prompt 39 in `01-prompts/15-cg-execute/` dedicated to **Golang Pointer Reduction and Mutation Reduction** (`39-golang-pointers-and-mutation-reduction.md`), alongside symmetric companion skills for Antigravity Agents (`.agents/skills/cg-golang-pointers-and-mutation-reduction/skill.md`) and Cursor IDE (`.cursor/skills/cg-golang-pointers-and-mutation-reduction/skill.md`), fully registered in all catalogs and synchronized across all 43 fleet repositories.

The prompt implements the canonical V6 Tree Swarm pattern ($N = 400$, $A = 3$, $H = 2$, Capacity = 6, Immediate Dynamic Respawn & Join, Fast-Forwarding Micro-Batches, Decoupled Terminal E2E Verification Gate, and Cognitive Heatmap Telemetry). It enforces:
1. **Golang Pointer Reduction & Value Semantics:**
   - Value semantics for small structs (<= 64 bytes).
   - Value receivers on read-only / query methods (`func (s Struct) Method()`).
   - Pure slice (`[]T`) and map (`map[K]V`) parameter passing without pointer indirection.
   - Elimination of heap allocations, pointer dereferences, and nil-pointer panic surfaces.
   - Pointer justification whitelist: structs with `sync.Mutex`/`sync.RWMutex`, methods mutating internal struct state, tri-state nilability, `*appfault.AppError`.
2. **Golang Mutation Avoidance & Immutability:**
   - Single-assignment immutability: variables initialized once upon declaration.
   - Total ban on post-construction field mutation (`s.Field = x` after constructor); pass all values to pure constructors or struct literals.
   - Total ban on intermediate variable mutation across `switch` / `if-else if` branch ladders; return directly or extract helper functions.
   - Pure functional transformations: copy-on-write methods returning new instances.
   - Unavoidable state (lazy caches, buffers) protected strictly behind mutex locks (`GetLinesLock()`).

---

## Decomposed Subtasks Breakdown

### Wave 1: Specification Authoring (Phase 1)
- **Task-01:** Spec Author 01: Architecture & Continuum Spec
  - File: `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/01-architecture-spec.md`
  - Subtask: `.ai-memory/plans/subtasks/06-cg-execute-golang-pointers-and-mutation-reduce/01-subtask-architecture-and-contracts.md`
- **Task-02:** Spec Author 02: Component Spec, Pointer Reduction & Mutation Safety
  - File: `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/02-component-spec.md`
  - Subtask: `.ai-memory/plans/subtasks/06-cg-execute-golang-pointers-and-mutation-reduce/02-subtask-component-and-examples.md`
- **Task-03:** Spec Author 03: Migration, Fleet Sync & Verification Spec
  - File: `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/03-migration-and-verification-spec.md`
  - Subtask: `.ai-memory/plans/subtasks/06-cg-execute-golang-pointers-and-mutation-reduce/03-subtask-migration-and-verification.md`

### Wave 2: Execution & Implementation (Phase 2)
- **Task-04:** Worker 01: Canonical Prompt 39 Authoring (`39-golang-pointers-and-mutation-reduction.md` + `01-prompts/15-cg-execute/readme.md`)
- **Task-05:** Worker 02: Companion Skills Parity Authoring (`.agents/skills/cg-golang-pointers-and-mutation-reduction/skill.md` + `.cursor/skills/cg-golang-pointers-and-mutation-reduction/skill.md`)
- **Task-06:** Worker 03: Master Catalogs, Registries, Quality Linters & Fleet Synchronization

---

## Non-Negotiable Invariants & Guardrails
- **GitMap Search Primacy:** Live searches exclusively via `gitmap aum search`, `gitmap find`, `gitmap cat`, `gitmap lf`. Total ban on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`.
- **Strict Relative Git Paths:** Only relative paths starting from git root. Total ban on absolute filesystem paths and `file:///` URIs.
- **Strict Lowercase Naming:** All paths and files lowercase.
- **Worker Git Ban:** Subagents never execute git commands.
- **No Push = Not Done:** Task is incomplete until staged, committed, and pushed to remote `main` via `gitmap cpf`.
