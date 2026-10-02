# Specification: CG-Execute Prompt 37 & Antigravity Skill for String Normalization and EqualFoldAny

> **/goal** Establish the formal specification and execution contracts for Prompt 37 (`01-prompts/15-cg-execute/37-string-normalization-and-equalfoldany.md`) and native Antigravity skill `cg-string-normalization-and-equalfoldany`, enforcing automated AST discovery, code reduction, and surgical refactoring of repetitive string normalization and chained equality checks across polyglot codebases.  
> **/learn** Master the V6 execution parameters (`N = 300, A = 2, H = 2, C = 30`), semicolon slash commands (`[/goal]`, `[/learn]`, `[/plan]`), mandatory subagent spawning gates, the Search First Directive for `pkg/strutil/strutil.go`, AST discovery patterns via GitMap, polyglot code transformation rules (Go, TypeScript, Rust), and strict R1–R16 operational rule citations.

**Version:** 1.0.0  
**Updated:** 2026-10-02  
**Status:** Active  
**AI Confidence:** Production-Ready  
**Ambiguity:** None  

---

## 1. Executive Summary & Motivation

In large multi-language codebases (Go, TypeScript, Rust, Python), string comparisons, sanitization, and normalization logic are frequently implemented ad-hoc at individual call sites. Developers repeatedly write verbose, allocation-heavy, and error-prone patterns:

1. **Repetitive Case Normalization & Trimming:** Calling `strings.ToLower(strings.TrimSpace(val))` or `val.trim().toLowerCase()` repeatedly across functions before evaluation.
2. **Chained Case-Insensitive Disjunctions:** Writing long OR-chains of equality checks:
   ```go
   if strings.EqualFold(role, "admin") || strings.EqualFold(role, "superadmin") || strings.EqualFold(role, "owner") { ... }
   ```
   or with manual trimming:
   ```go
   trimmed := strings.TrimSpace(status)
   if strings.EqualFold(trimmed, "active") || strings.EqualFold(trimmed, "pending") || strings.EqualFold(trimmed, "verified") { ... }
   ```
3. **Redundant Allocations in Loops:** In Go, calling `strings.ToLower()` creates a heap allocation for the lowered string copy. When repeated inside tight loops or high-throughput routes, these allocations degrade performance.
4. **Fragile Negative and Complex Conditions:** Chaining manual trimming with inequality checks or mixing polarity across conditions, violating repository boolean guidelines.

### The Solution: Prompt 37 & Dedicated Antigravity Skill

To eliminate this boilerplate repository-wide, we define:
- **`01-prompts/15-cg-execute/37-string-normalization-and-equalfoldany.md`:** A fully autonomous V6 guideline execution prompt that discovers chained string comparisons, manual trim-and-compare operations, and case-conversion disjunctions, refactoring them into canonical utility calls like `strutil.EqualFoldAny` and `strutil.EqualFoldAnyTrim`.
- **`.agents/skills/cg-string-normalization-and-equalfoldany/skill.md`:** A native Antigravity skill exposing these audit, refactoring, and verification capabilities to IDE and CLI runtimes.

---

## 2. Specification for Prompt 37 (`37-string-normalization-and-equalfoldany.md`)

### 2.1 File Placement & Header Architecture

- **Path:** `01-prompts/15-cg-execute/37-string-normalization-and-equalfoldany.md`
- **Naming:** Strictly lowercase, hyphen-separated.
- **Top Parameter Header (Strict V6 Format):**

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

### 2.2 Semicolon Slash Commands

Prompt 37 must begin with three standardized semicolon slash commands:

1. `[/goal](slashCommand;goal)`:
   > Autonomously scan, plan, refactor, and fix all repetitive string normalization, manual trimming, and chained equality comparisons across Go, TypeScript, and Rust codebases. Modifying source files directly, enforce the Search First Directive for `pkg/strutil/strutil.go` (or repo string utility), replace manual trimming and `EqualFold` / `ToLower` OR-chains with canonical helpers (`strutil.EqualFoldAny`, `strutil.EqualFoldAnyTrim`, `strutil.NormalizeLowerTrim`), eliminate redundant allocations, enforce positive boolean conventions (`is`/`has` only), maintain functions <= 8-15 lines, and defer verification strictly to targeted linters without running intermediate tests or builds: FIRST showcase and list out the given task in visible chat during Turn 1, capture it verbatim, plan it in the repo, spawn autonomous subagents via `invoke_subagent` (A = 2, H = 2; solo execution without calling `invoke_subagent` is an auto-reject failure) in disjoint file boxes using GitMap high-speed commands as primary, prove every single claim with concrete evidence, enforce coding guidelines to 100%, and finish with one atomic GitMap commit that holds strictly this task's files.
2. `[/learn](slashCommand;learn)`:
   > Enforce the Top-Instruction Priority Mandate: whatever directives, custom requirements, checklists, or user instructions are provided ABOVE this prompt outrank everything below. Turn 1 MUST showcase the given task list in visible chat before any background execution. Master the string normalization conventions: `02-spec/02-coding-guidelines/01-cross-language/34-string-normalization-and-equalfoldany.md` and `02-spec/21-app/02-string-normalization-and-equalfoldany/01-core-architecture-spec.md`. Each rule is stated once (R1 to R16) and cited by ID. Progress lives in the ledger and in `.ai-memory/plans/`, never only in chat.
3. `[/plan](slashCommand;plan)`:
   > Execute thorough step-by-step planning in the repository before execution. Ensure all deliverables, architecture boundaries, and requirements are clearly defined in the audit ledger and subtask plans before dispatching worker waves.

### 2.3 Mandatory Subagent Spawning Gate (`A = 2, H = 2`)

- **Tool Call Enforcement:** Lead orchestrator MUST execute actual `invoke_subagent` tool calls. Merely outputting markdown text claiming workers are dispatched without making the tool call is an auto-reject failure.
- **3-Stage Dispatch:**
  1. *Planning Step:* Spawn 2 subagents (`research`) to scan the codebase using GitMap and author the plan.
  2. *Spec Step:* Spawn 2 subagents (`research` or `self`) to write specifications and decompose into subtasks.
  3. *Execution Step:* Spawn 2 worker subagents (`self`) per wave, each handling up to H subtasks in disjoint file boxes.
- **Solo Execution Ban:** The lead orchestrator is strictly banned from executing file modifications solo.

### 2.4 Search First Directive for `pkg/strutil/strutil.go`

Before writing local helper functions, modifying signatures, or introducing duplicate logic, agents must strictly follow the **Search First Directive**:

1. **Discovery Order:**
   - Go: Scan for `pkg/strutil/strutil.go`, `04-code/golang/pkg/strutil/strutil.go`, or existing repository string utility packages using `gitmap find-files-any "strutil"`.
   - TypeScript: Scan for `src/utils/strutil.ts`, `lib/strutil.ts`, or `@/shared/utils/string`.
   - Rust: Scan for `crate::strutil`, `crate::util::string`, or equivalent modules.
2. **Re-use Mandate:** If the canonical string utility exists, all refactored code must import and utilize it directly. Creating local duplicate helper functions (e.g. `isOneOf`, `matchesAny`, `cleanEqual`) inside implementation packages is strictly banned.
3. **Absence Protocol:** If the canonical utility is absent in the target repository, agents must create or register it according to the canonical specification (`02-spec/02-coding-guidelines/01-cross-language/34-string-normalization-and-equalfoldany.md`) before refactoring call sites.

### 2.5 Code Reduction Transformation Rules

Prompt 37 must specify exact before/after code refactoring transformations:

#### Rule T1: Chained `strings.EqualFold` OR-Chains
- **Anti-Pattern:**
  ```go
  // ❌ ANTI-PATTERN: Verbose, repetitive EqualFold chaining
  if strings.EqualFold(role, "admin") || strings.EqualFold(role, "root") || strings.EqualFold(role, "superuser") {
      return allowAccess()
  }
  ```
- **Canonical Refactoring:**
  ```go
  // ✅ CANONICAL: Concise, zero-allocation variadic helper
  isAuthorized := strutil.EqualFoldAny(role, "admin", "root", "superuser")
  if isAuthorized {
      return allowAccess()
  }
  ```

#### Rule T2: Manual Trimming + Chained `EqualFold`
- **Anti-Pattern:**
  ```go
  // ❌ ANTI-PATTERN: Multi-step trimming, intermediate variable, chained comparison
  trimmed := strings.TrimSpace(status)
  if strings.EqualFold(trimmed, "active") || strings.EqualFold(trimmed, "enabled") {
      proceed()
  }
  ```
- **Canonical Refactoring:**
  ```go
  // ✅ CANONICAL: Single-expression atomic trim-and-match
  isActive := strutil.EqualFoldAnyTrim(status, "active", "enabled")
  if isActive {
      proceed()
  }
  ```

#### Rule T3: `strings.ToLower(strings.TrimSpace(s)) == "val"`
- **Anti-Pattern:**
  ```go
  // ❌ ANTI-PATTERN: Heap allocation for ToLower and manual comparison
  if strings.ToLower(strings.TrimSpace(env)) == "prod" || strings.ToLower(strings.TrimSpace(env)) == "production" {
      enableHardening()
  }
  ```
- **Canonical Refactoring:**
  ```go
  // ✅ CANONICAL: Normalized case-insensitive matching without allocations
  isProduction := strutil.EqualFoldAnyTrim(env, "prod", "production")
  if isProduction {
      enableHardening()
  }
  ```

#### Rule T4: Single Value Trimmed Comparison
- **Anti-Pattern:**
  ```go
  if strings.EqualFold(strings.TrimSpace(input), "yes") { ... }
  ```
- **Canonical Refactoring:**
  ```go
  isAffirmative := strutil.EqualFoldTrim(input, "yes")
  if isAffirmative { ... }
  ```

#### Rule T5: Boolean Polarity & Guard Clause Integrity
- Positive naming mandatory: `isValidRole := strutil.EqualFoldAny(...)`, never `!isInvalid`.
- No explicit `== true`: `if isValidRole { ... }`, never `if isValidRole == true`.
- Discrete conditions: Never mix positive and negative checks in the same if statement.
- Vertical line gaps: Mandatory blank line before `if`, after `}`, and before `return`.

### 2.6 AST Discovery Patterns via GitMap

Prompt 37 must enforce multi-core GitMap search commands and ban generic shell search tools (`Select-String`, `git grep`):

| Pattern Category | Target Construct | High-Speed GitMap Search Command |
| :--- | :--- | :--- |
| **EqualFold Chains (Go)** | `EqualFold(...) \|\| EqualFold(...)` | `gitmap aum search -r "EqualFold\(.*\|\|.*EqualFold\(" [dir] -e .go` |
| **Trim + EqualFold (Go)** | `TrimSpace(...) ... EqualFold(...)` | `gitmap aum search -r "TrimSpace\(.*EqualFold" [dir] -e .go` |
| **ToLower Equality (Go)** | `strings.ToLower\(.*==` | `gitmap aum search -r "strings\.ToLower\(.*==.*\|\|" [dir] -e .go` |
| **Generic EqualFold (Go)** | All `EqualFold` call sites | `gitmap aum search "strings.EqualFold" [dir] -e .go` |
| **TypeScript toLowerCase** | `.toLowerCase() === ... \|\| ...` | `gitmap aum search -r "\.toLowerCase\(\)\s*===\s*.*\|\|" [dir] -e .ts` |
| **TypeScript Trim + Lower** | `.trim().toLowerCase()` | `gitmap aum search -r "\.trim\(\)\.toLowerCase\(\)" [dir] -e .ts` |
| **Rust eq_ignore_ascii** | `eq_ignore_ascii_case(...) \|\|` | `gitmap aum search -r "eq_ignore_ascii_case\(.*\|\|" [dir] -e .rs` |

### 2.7 3-Phase Pipeline Architecture

1. **Phase 1A: Verbatim Capture, Task Extraction & Turn 1 Chat Gate:**
   - Output confirmed task breakdown in visible chat on Turn 1.
   - Capture prompt losslessly under `## User Request (Verbatim)` in spec and plan.
   - Extract discrete deliverables with ordered IDs (`Task-01`, `Task-02`, ...).
   - Chain first tool call in the same turn without stopping.
2. **Phase 1B: Planning & AST Discovery (Steps 1 .. 150):**
   - Preflight tool check (`gitmap lf readme.md`, `python --version`).
   - Initialize Ledger (`.ai-memory/temp-agents/NN-<slug>/ledger.md`).
   - Spawn `A = 2` research subagents via `invoke_subagent` to discover string manipulation patterns using GitMap.
   - Author Canonical Spec in `02-spec/21-app/NN-<slug>/` and Lean Subtask Plans in `.ai-memory/plans/subtasks/NN-<slug>/`.
3. **Phase 2: Parallel Worker Execution (Steps 151 .. 300):**
   - Dispatch `A = 2` worker subagents (`TypeName: "self"`, `H = 2` tasks) per wave.
   - Enforce disjoint bounding boxes (one owner per file).
   - Strict micro-batching: 5–8 files per subtask.
   - Workers report back via structured JSON contract within `C = 30` tool calls.
   - Yield turn immediately after dispatching workers; no busy-waiting.
4. **Phase 3: Consolidation, Evidence Verification & Atomic GitMap Push:**
   - Consolidate subtasks into `.ai-memory/plans/completed/NN-<slug>.md`.
   - Update registries (`.ai-memory/plans/readme.md`, `02-spec/21-app/readme.md`, `01-prompts/readme.md`).
   - Verify push gate (targeted linters exit 0, secrets gate clean, `.gitignore` synced).
   - Atomic commit & push: `gitmap cpf "<module> - <summary>"` or `gitmap cpb "<module> - <summary>"`.

### 2.8 Core Operational Rules (Cited R1 to R16 by ID)

Prompt 37 must explicitly cite and enforce every rule:
- **R1:** Zero builds or test suites (TOTAL BAN on `go build`, `npm run build`, `go test ./...`, `pytest`).
- **R2:** Targeted checks only (run fast file-scoped linters; check scanning 0 files is a FAIL).
- **R3:** Evidence or it did not happen (must cite concrete file path, diffstat, or exit 0).
- **R4:** Never invent commands, flags, or paths.
- **R5:** Mandatory subagents (`invoke_subagent`, `A = 2, H = 2`). Solo execution is an auto-reject failure.
- **R6:** One owner per file (disjoint bounding boxes across workers).
- **R7:** Git safety & isolation (subagents never run git commands).
- **R8/R9:** Atomic commit & push via GitMap (`gitmap cpf "<module> - <summary>"`).
- **R10:** Zero unauthorized releases (no version bumps or changelog edits without command).
- **R11:** Strict relative git paths & lowercase hygiene (zero absolute paths, zero `file:///` URIs).
- **R12:** No polling / immediate turn yielding after dispatch.
- **R13:** Two-strike retry cap & anti-looping.
- **R14:** 100% ambiguity & decision boundaries.
- **R15:** Zero generated artifacts committed.
- **R16:** Zero secrets in standard repos (`repo-secrets` offload via `gitmap rs`).

---

## 3. Specification for Native Antigravity Skill (`cg-string-normalization-and-equalfoldany`)

### 3.1 Skill Placement & Frontmatter

- **Path:** `.agents/skills/cg-string-normalization-and-equalfoldany/skill.md`
- **YAML Frontmatter:**

```yaml
---
name: cg-string-normalization-and-equalfoldany
description: Autonomously audits, refactors, and standardizes string normalization, case-insensitive comparisons, and chained equality into reusable strutil helpers (EqualFoldAny, EqualFoldAnyTrim) across Go, TypeScript, and Rust.
---
```

### 3.2 Trigger Keywords & Aliases

The skill must declare primary trigger keywords and aliases:
- `cg-string-normalization`
- `cg-equalfoldany`
- `strutil-normalization`
- `audit-string-comparisons`
- `cg-execute equalfoldany`

### 3.3 Skill Execution Workflow

1. **Phase 1: Search First & Inventory:**
   - Verify presence of canonical `pkg/strutil/strutil.go` (or language equivalent).
   - If missing, consult `02-spec/02-coding-guidelines/01-cross-language/34-string-normalization-and-equalfoldany.md` and generate core utility functions first.
   - Run GitMap AST searches to locate all occurrences of chained equality and manual trim-before-compare operations.
2. **Phase 2: Micro-Batched Refactoring (5–8 Files per Batch):**
   - Replace verbose chains with `strutil.EqualFoldAny` or `strutil.EqualFoldAnyTrim`.
   - Hoist repeated comparisons out of loops.
   - Enforce positive boolean variables (`isMatch`, `hasTargetRole`).
   - Enforce function size limits (<= 8–15 lines).
3. **Phase 3: Targeted Verification & Quality Gates:**
   - Run guideline autofixer in check-only mode: `python 03-ai-scripts/05-guideline-autofixer.py <folder> --check-only`.
   - Run relative path and forbidden strings linters.
   - Do NOT run full compilation or heavy unit test suites.

### 3.4 Polyglot Code Transformation Catalog

The skill must provide complete, production-ready code examples across Go, TypeScript, and Rust:

#### Go (`pkg/strutil`)

```go
// ❌ BAD: Redundant trimming, intermediate variables, chained EqualFold
trimmedRole := strings.TrimSpace(userRole)
if strings.EqualFold(trimmedRole, "admin") || strings.EqualFold(trimmedRole, "owner") || strings.EqualFold(trimmedRole, "moderator") {
    grantPermission()
}

// ✅ GOOD: Clean atomic evaluation via strutil helper
isElevated := strutil.EqualFoldAnyTrim(userRole, "admin", "owner", "moderator")
if isElevated {
    grantPermission()
}
```

```go
// ❌ BAD: Allocating ToLower in loop
for _, item := range rawItems {
    if strings.ToLower(item) == "active" || strings.ToLower(item) == "pending" {
        results = append(results, item)
    }
}

// ✅ GOOD: Zero-allocation EqualFoldAny
for _, item := range rawItems {
    isTargetStatus := strutil.EqualFoldAny(item, "active", "pending")
    if isTargetStatus {
        results = append(results, item)
    }
}
```

#### TypeScript (`strutil.ts`)

```typescript
// ❌ BAD: Chained .toLowerCase() and .trim()
const clean = input ? input.trim().toLowerCase() : '';
if (clean === 'yes' || clean === 'y' || clean === 'true' || clean === '1') {
  enableFeature();
}

// ✅ GOOD: Reusable helper
const isAffirmative = strutil.equalFoldAnyTrim(input, 'yes', 'y', 'true', '1');
if (isAffirmative) {
  enableFeature();
}
```

#### Rust (`strutil.rs`)

```rust
// ❌ BAD: Repetitive trimming and case-insensitive check chaining
let val = input.trim();
if val.eq_ignore_ascii_case("draft") || val.eq_ignore_ascii_case("review") {
    process_document();
}

// ✅ GOOD: Slice-based case-folding helper
let is_pending = strutil::equal_fold_any_trim(input, &["draft", "review"]);
if is_pending {
    process_document();
}
```

---

## 4. Traceability & Dependencies

- **Parent Spec:** `02-spec/21-app/02-string-normalization-and-equalfoldany/01-core-architecture-spec.md`
- **Coding Guideline Spec:** `02-spec/02-coding-guidelines/01-cross-language/34-string-normalization-and-equalfoldany.md`
- **Subtask Plan:** `.ai-memory/plans/subtasks/02-string-normalization-and-equalfoldany/02-cg-execute-prompt-and-skill.md`
- **Master Plan Register:** `.ai-memory/plans/readme.md`
- **Prompt Catalog Registries:**
  - `01-prompts/15-cg-execute/readme.md` (Sequence #37)
  - `01-prompts/readme.md` (Section 15)
  - `.ai-memory/prompts.md`
