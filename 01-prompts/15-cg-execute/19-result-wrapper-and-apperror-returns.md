# Result Wrapper Types, Collections & AppError Returns — Coding Guideline (must follow)

Trigger Keywords & Aliases: `cg-result-wrapper`, `cg-apperror-returns`, `cg-execute result-wrapper`, `audit result wrapper`, `fix map return error`, `fix slice return error`, `single return object audit`, `enforce apperror returns`, `enforce result map`, `fix multi-value returns`, `is-count-other-than`, `has-record`, `is-defined`

> **Prompt Version:** 2.2.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

```text
N = 200
```

N = total self-loop steps budget that the agents will perform.

/goal Autonomously scan, discover, plan, refactor, and verify all Go functions returning multi-value error tuples (such as `(map[K]V, error)`, `([]T, error)`, or `(T, error)`), eliminating raw standard library error returns, replacing them with strongly-typed result wrappers (`ResultMap[K, V]`, `ResultSlice[T]`, `Result[T]`) and structured `*appfault.AppError` returns, guaranteeing a single return object, standardized outer-layer inspection predicates (`IsSuccess`, `IsFailure`, `HasError`, `IsEmptyError`, `IsEmpty`, `HasRecord`, `IsDefined`, `IsCountOtherThan`, `Data`, `Items`, `AppError`, `Fault`, `Get`, `Has`, `Count`), eliminating dual-handling, and replacing verbose `if err != nil || len(...) != N` or `IsFailure() || Count() != N` conditions with fluent `if res.IsCountOtherThan(N)` across the entire codebase until 100% green without stopping.

### Master Task Checklist (Atomic Numbered Steps)

1. [ ] /goal Phase 1 (Step A): Deeply scan the target codebase using ripgrep to inventory all functions returning multi-value tuples `(T, error)`, `(map[K]V, error)`, `([]T, error)`, and raw stdlib `error` returns. Also inventory clumsy caller checks like `err != nil || len(...) != N` and `IsFailure() || Count() != N`.
2. [ ] /goal Phase 1 (Step B): Write the master audit specification in `.lovable/plans/pending/XX-result-wrapper-audit.md` with an exhaustive Violation Ledger table.
3. [ ] /goal Phase 1 (Step C): Decompose the master plan into granular, atomic subtasks in `.lovable/plans/subtasks/XX-result-wrapper/`.
4. [ ] /goal Phase 1 (Step D): Verify or create the automated quality linter and register in `03-ai-scripts/01-index.md`.
5. [ ] /goal Phase 2 (Step A): Open each target file and refactor function signatures from multi-value returns to single `ResultMap[K, V]`, `ResultSlice[T]`, or `Result[T]` envelopes.
6. [ ] /goal Phase 2 (Step B): Replace raw stdlib `error` returns with structured `*appfault.AppError` instances using `appfault.New()` or `appfault.Wrap()`.
7. [ ] /goal Phase 2 (Step C): Modernize all caller call sites to utilize outer-layer inspection methods (`res.IsSuccess()`, `res.IsFailure()`, `res.IsEmpty()`, `res.HasRecord()`, `res.IsDefined()`, `res.IsCountOtherThan(N)`, `res.Get()`, `res.AppError()`), eliminating manual `err != nil || len(...) != N` boilerplate.
8. [ ] /goal Phase 2 (Step D): Enforce <= 8–15 line function decomposition and clean blank-line spacing.
9. [ ] /goal Phase 2 (Step E): Execute targeted file-level linters (`python linter-scripts/check-function-lengths.py`, `check-mws-error-codes.py`, `check-newline-styling.py`) to verify 0 remaining violations. DO NOT run the full CI/CD pipeline runner (`06-cicd-local-runner.py`) during routine coding guideline execution turns.
10. [ ] /learn Ingest `.lovable/memory/01-index.md` for project memory index and past learnings.
11. [ ] /learn Ingest `.lovable/strictly-avoid.md` for banned anti-patterns and strict constraints.
12. [ ] /learn Ingest `02-spec/02-coding-guidelines/02-canonical-size-tier.md` for canonical file and function size tiers.
13. [ ] /learn Ingest `02-spec/02-coding-guidelines/01-cross-language/01-index.md` for single return type mandates and micro-tasking.
14. [ ] /learn Ingest `02-spec/02-coding-guidelines/01-cross-language/01-index.md` for strict relative path citation requirements.
15. [ ] /learn Ingest `02-spec/03-error-manage/01-index.md` for universal AppError wrapping and error envelopes.
16. [ ] /learn Ingest `02-spec/03-error-manage/02-error-architecture/02-error-handling-reference.md` for error handling architecture and Result wrappers.
17. [ ] /learn Ingest `02-spec/03-error-manage/03-error-code-registry/02-registry.md` for structured error code catalog.
18. [ ] /learn Ingest `02-spec/03-error-manage/02-error-architecture/05-response-envelope/05-response-envelope-reference.md` for response envelope schemas.
19. [ ] /learn Ingest `02-spec/03-error-manage/02-error-architecture/06-apperror-package/03-go-apperror-linter-spec.md` for Go AppError implementation specifications.
20. [ ] /learn Ingest `02-spec/03-error-manage/02-error-architecture/06-apperror-package/01-apperror-reference/04-result-types.md` for Result[T], ResultSlice[T], and ResultMap[K, V] method specifications.
21. [ ] /learn Ingest `.lovable/coding-guidelines.md` for master consolidated coding guidelines.
22. [ ] /goal Create or update agent rules in the repository if missing from agent memory.

```text
PHASE_1_STEPS = N / 2   (Steps 1 .. N/2: Scan Multi-Value Returns, Build Violation Ledger in .lovable/plans/pending/, Subtasks, Linter Hook)
PHASE_2_STEPS = N / 2   (Steps N/2+1 .. N: Actively Edit Code, Refactor Signatures to ResultMap/Result, Modernize Call Sites, Verify Local Linters)
```

N, PHASE_1_STEPS, and PHASE_2_STEPS are read-only after initialization. Never modify them mid-execution.

---

## The Code Pattern & Anti-Pattern Analysis

In legacy Go codebases, developers frequently write functions that return multi-value tuples pairing collections with the standard library `error` interface, and call sites perform fragile, compound boolean assertions.

### 1. The Problematic Legacy Store Pattern

Consider this common database query implementation:

```go
// ❌ ANTI-PATTERN: Multi-value tuple return with raw standard library error
func (s *SQLiteStore) queryAllMacroSteps(db *sql.DB) (map[string][]MacroStep, error) {
    rows, err := db.Query("SELECT macro_id, step_name, action, payload FROM macro_steps ORDER BY macro_id, step_order")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return scanMacroStepsMap(rows)
}

// ❌ ANTI-PATTERN: Secondary scanner returning raw map and error tuple
func scanMacroStepsMap(rows *sql.Rows) (map[string][]MacroStep, error) {
    stepsMap := make(map[string][]MacroStep)
    for rows.Next() {
        var macroId, name, action, payload string
        if err := rows.Scan(&macroId, &name, &action, &payload); err != nil {
            return nil, err
        }
        stepsMap[macroId] = append(stepsMap[macroId], MacroStep{
            Name:    name,
            Action:  action,
            Payload: payload,
        })
    }

    return stepsMap, rows.Err()
}
```

### 2. The Problematic Legacy Caller & Test Assertion Pattern

At call sites and in test suites, multi-value returns force awkward tuple unpacking and compound checks:

```go
// ❌ ANTI-PATTERN: Tuple unpacking + compound condition mixing error check and length check
details, err := pipeDb.QueryDetailedErrorLogsByRunId(runId)
if err != nil || len(details) != 1 {
    t.Fatalf("expected 1 detail log, got %d (err: %v)", len(details), err)
}

if !strings.Contains(details[0].RawLogs, "PASS: Test0") {
    t.Errorf("detail log missing raw PASS line: %s", details[0].RawLogs)
}
```

### Why This Pattern Violates Repository Guidelines

1. **Violates the Single Return Object Mandate:**
   - Multi-value returns like `(map[string][]MacroStep, error)` or `([]DetailedErrorLog, error)` violate the core architectural standard requiring functions to return a single strongly-typed envelope object.
   - Returning tuples forces dual assignment unpacking (`val, err := ...`).
2. **Raw Standard Library `error` Anti-Pattern:**
   - Returning the bare standard library `error` interface strips domain context, structured error codes, file/line tracing, and machine-readable metadata.
   - All errors in the repository MUST use structured `*appfault.AppError`.
3. **Dual-Handling Risk and Ambiguous Empty Returns:**
   - Does returning `(nil, nil)` represent an empty database table or an uninitialized store?
   - If an error occurs midway through row scanning, returning `(nil, err)` drops partially collected records, while returning `(stepsMap, err)` tempts callers into dual handling (processing data *and* logging error).
4. **Call-Site Clutter & Compound Disjunctions:**
   - The check `if err != nil || len(details) != 1` combines two disparate concepts (execution failure vs. cardinality mismatch) into an unseparated condition.
   - If `details` is nil upon failure, downstream slice index operations (`details[0]`) risk index panics if the guard block is refactored improperly.

---

## The Modern Refactored Architecture

Under Prompt Architect coding guidelines, all multi-value returns are refactored into dedicated result container types from package `pkg/appfault`:
- Key-Value Maps: `appfault.ResultMap[K, V]`
- Lists & Slices: `appfault.ResultSlice[T]`
- Scalar Values: `appfault.Result[T]`
- Pure Side-Effects: `*appfault.AppError` (zero bare `void` / empty returns)

### Modern Refactored Store Implementation

```go
// ✅ MODERN PATTERN: Single ResultMap return envelope with structured AppError
func (s *SQLiteStore) queryAllMacroSteps(db *sql.DB) appfault.ResultMap[string, []MacroStep] {
    rows, err := db.Query("SELECT macro_id, step_name, action, payload FROM macro_steps ORDER BY macro_id, step_order")
    if err != nil {
        return appfault.FailMap[string, []MacroStep](
            appfault.New(appfault.ErrDatabaseQuery).
                WithCause(err).
                WithMessage("failed to query macro steps from database"),
        )
    }
    defer rows.Close()

    return scanMacroStepsMap(rows)
}

// ✅ MODERN PATTERN: Scanner returning strongly-typed ResultMap
func scanMacroStepsMap(rows *sql.Rows) appfault.ResultMap[string, []MacroStep] {
    stepsMap := make(map[string][]MacroStep)
    for rows.Next() {
        var macroId, name, action, payload string
        if err := rows.Scan(&macroId, &name, &action, &payload); err != nil {
            return appfault.FailMap[string, []MacroStep](
                appfault.New(appfault.ErrDatabaseScan).
                    WithCause(err).
                    WithMessage("failed to scan macro step row"),
            )
        }

        stepsMap[macroId] = append(stepsMap[macroId], MacroStep{
            Name:    name,
            Action:  action,
            Payload: payload,
        })
    }

    if err := rows.Err(); err != nil {
        return appfault.FailMap[string, []MacroStep](
            appfault.New(appfault.ErrDatabaseIteration).
                WithCause(err).
                WithMessage("row iteration failed for macro steps"),
        )
    }

    return appfault.OkMap(stepsMap)
}
```

---

## The 4 Core Predicate Methods (Must Enforce)

Result envelopes (`ResultSlice[T]`, `ResultMap[K, V]`, `Result[T]`) provide four core predicate methods that eliminate call-site boilerplate and compound boolean conditions:

### 1. `res.IsCountOtherThan(number int) bool`

- **Exact Semantics:** Returns `true` if the operation failed (has error) **OR** if the record count is not equal to `number`.
- **Purpose:** Replaces the compound check `if res.IsFailure() || res.Count() != N` with a single, intention-revealing predicate.
- **Behavior:**
  - If `res.IsFailure()`: returns `true`.
  - If `res.IsSuccess()`: returns `res.Count() != number`.

### 2. `res.IsEmpty() bool`

- **Exact Semantics:** Returns `true` if the collection contains `0` items/entries, or if scalar payload `T` is null/empty/zero.
- **Purpose:** Cleanly checks for zero records after validating success.

### 3. `res.HasRecord() bool` (and alias `res.HasRecords() bool`)

- **Exact Semantics:** Returns `true` if the operation succeeded (no error) **AND** contains **more than 0 records** (`res.Count() > 0`).
- **Purpose:** Direct check when logic requires at least one record before continuing.

### 4. `res.IsDefined() bool`

- **Exact Semantics:** Returns `true` if the operation succeeded (no error) **AND** has `recordCount > 0` (or data `T` is non-null/non-empty).
- **Distinction from `IsSuccess()`:**
  - `IsSuccess()` means "no error occurred" (an empty query returning 0 items succeeds without error).
  - `IsDefined()` means "no error occurred AND actual data exists" (recordCount > 0 or not null).

---

## Comparison of Caller Patterns

### Evolution: Legacy vs. Transitional vs. Modern

```go
// ------------------------------------------------------------
// ❌ 1. LEGACY PATTERN: Raw multi-value unpacking + compound boolean
// ------------------------------------------------------------
details, err := pipeDb.QueryDetailedErrorLogsByRunId(runId)
if err != nil || len(details) != 1 {
    t.Fatalf("expected 1 detail log, got %d (err: %v)", len(details), err)
}

// ------------------------------------------------------------
// ⚠️ 2. TRANSITIONAL PATTERN: Single Result envelope, but verbose manual check
// ------------------------------------------------------------
detailRes := pipeDb.QueryDetailedErrorLogsByRunId(runId)
if detailRes.IsFailure() || detailRes.Count() != 1 {
    t.Fatalf("expected 1 detail log, got %d (err: %v)", detailRes.Count(), detailRes.AppError())
}

// ------------------------------------------------------------
// ✅ 3. CANONICAL MODERN PATTERN: Single expressive guard
// ------------------------------------------------------------
detailRes := pipeDb.QueryDetailedErrorLogsByRunId(runId)
if detailRes.IsCountOtherThan(1) {
    t.Fatalf("expected 1 detail log, got %d (err: %v)", detailRes.Count(), detailRes.AppError())
}

details := detailRes.Data
if !strings.Contains(details[0].RawLogs, "PASS: Test0") {
    t.Errorf("detail log missing raw PASS line: %s", details[0].RawLogs)
}
```

### Production Service Caller Example

```go
// Loading workflow steps:
stepRes := store.QueryMacroSteps(db, macroId)
if stepRes.IsFailure() {
    return stepRes.AppError().WithContext("caller", "executeWorkflow")
}

// Check if no records were found:
if stepRes.IsEmpty() {
    log.Println("No macro steps configured for macroId")
    return nil
}

// Check if at least one record exists:
if stepRes.HasRecord() {
    log.Printf("Loaded %d macro step(s)", stepRes.Count())
}

// Process single value safely:
if stepRes.IsDefined() {
    executeSteps(stepRes.Items)
}
```

---

## Standardized Outer-Layer Inspection Methods Table

| Method | Return Type | Applicable To | Purpose & Exact Behavior |
|---|---|---|---|
| `res.IsSuccess()` | `bool` | All Wrappers | Returns `true` if the operation succeeded with no error. |
| `res.IsFailure()` / `res.IsFailed()` | `bool` | All Wrappers | Returns `true` if the operation encountered an error. |
| `res.HasError()` | `bool` | All Wrappers | Alias for `IsFailed()`. Returns `true` if error is present. |
| `res.IsEmptyError()` / `res.HasNoError()` | `bool` | All Wrappers | Returns `true` if no active error exists. |
| `res.IsCountOtherThan(number)` | `bool` | All Wrappers | Returns `true` if operation failed OR count != number. |
| `res.IsEmpty()` | `bool` | All Wrappers | Returns `true` if underlying collection has 0 items or payload is empty/null. |
| `res.HasRecord()` / `res.HasRecords()` | `bool` | All Wrappers | Returns `true` if operation succeeded AND count > 0 (more than 0 records). |
| `res.IsDefined()` | `bool` | All Wrappers | Returns `true` if operation succeeded (no error) AND recordCount > 0 / non-null. |
| `res.Count()` | `int` | All Wrappers | Returns total number of records/entries (or 0 if failed). |
| `res.Data` / `res.Items` / `res.Value()` | `T` / `[]T` / `map[K]V` | All Wrappers | Accesses the underlying data payload directly. |
| `res.AppError()` / `res.Fault()` | `*appfault.AppError` | All Wrappers | Retrieves structured error context for logging, propagation, or HTTP responses. |
| `res.Get(key)` | `(V, bool)` | `ResultMap[K, V]` | Safely retrieves map entry by key without nil-map panics. |
| `res.Has(key)` | `bool` | `ResultMap[K, V]` | Checks whether a key exists within the result map. |
| `res.Keys()` | `[]K` | `ResultMap[K, V]` | Returns deterministically sorted slice of all map keys. |
| `res.Values()` | `[]V` | `ResultMap[K, V]` | Returns slice of map values ordered according to `Keys()`. |
| `res.Filter(predicate)` | `ResultSlice[T]` | `ResultSlice[T]` | Returns filtered slice matching predicate (or self if failed). |
| `res.ForEach(fn)` | `ResultSlice[T]` | `ResultSlice[T]` | Iterates over elements with early exit via `ForEachBreak`. |

---

## Error Management Learning Checklist (`02-spec/03-error-manage/`)

Before refactoring error handling in any package, the agent must study and enforce the repository error management specifications:

- [ ] **Universal `*appfault.AppError` Standard (`02-spec/03-error-manage/01-index.md`):**
  - Never return bare `error` from domain services, repositories, or business logic.
  - Wrap third-party and standard library errors with `appfault.New()` or `appfault.Wrap()`.
- [ ] **Structured Error Codes (`02-spec/03-error-manage/03-error-code-registry/02-registry.md`):**
  - All errors must carry a typed `ErrorCode` string identifying the fault category (e.g. `ErrDatabaseQuery`, `ErrValidationFailed`, `ErrNotFound`).
- [ ] **Deterministic Error Handling & Envelopes (`02-spec/03-error-manage/02-error-architecture/02-error-handling-reference.md`):**
  - Use `appfault.Ok()`, `appfault.OkMap()`, and `appfault.OkSlice()` for successful results.
  - Use `appfault.Fail()`, `appfault.FailMap()`, and `appfault.FailSlice()` for failed results.
- [ ] **Universal Response Envelopes (`02-spec/03-error-manage/02-error-architecture/05-response-envelope/05-response-envelope-reference.md`):**
  - HTTP handlers and JSON serializers marshal `Result` and `ResultMap` into universal JSON response envelopes `{ "data": ..., "appError": ... }`.
- [ ] **Go AppError Architecture (`02-spec/03-error-manage/02-error-architecture/06-apperror-package/03-go-apperror-linter-spec.md`):**
  - Strict enforcement of `*appfault.AppError` return types and monadic helper methods across Go packages.
- [ ] **Result Types Specification (`02-spec/03-error-manage/02-error-architecture/06-apperror-package/01-apperror-reference/04-result-types.md`):**
  - Mandatory implementation of `IsCountOtherThan`, `IsEmpty`, `HasRecord`, and `IsDefined` on all result containers.

---

## Automated Codebase Scanning Guide

Use these exact `ripgrep` regex commands to discover legacy multi-value return patterns and clumsy caller checks across the codebase:

```bash
# 1. Find functions returning multi-value map tuples: (map[...], error)
rg --pcre2 "func\s+\w+\([^\)]*\)\s*\(\s*map\[[^\]]+\][^,]+,\s*error\)"

# 2. Find functions returning multi-value slice tuples: ([]..., error)
rg --pcre2 "func\s+\w+\([^\)]*\)\s*\(\s*\[\][^,]+,\s*error\)"

# 3. Find any function returning a tuple ending in standard library error
rg --pcre2 "func\s+\w+\([^\)]*\)\s*\([^\)]*,\s*error\)"

# 4. Find functions returning bare standard library error
rg --pcre2 "func\s+\w+\([^\)]*\)\s+error\s*\{"

# 5. Find dual-assignment caller unpacking: val, err := ...
rg --pcre2 "\b(\w+),\s*err\s*:=\s*"

# 6. Find compound caller checks: err != nil || len(...) != N
rg --pcre2 "(err\s*!=\s*nil\s*\|\|\s*len\([^\)]+\)\s*!=\s*\d+|len\([^\)]+\)\s*!=\s*\d+\s*\|\|\s*err\s*!=\s*nil)"

# 7. Find transitional checks: IsFailure() || ... Count() != N
rg --pcre2 "(IsFailure\(\)\s*\|\|\s*\w+\.Count\(\)\s*!=\s*\d+|\w+\.Count\(\)\s*!=\s*\d+\s*\|\|\s*\w+\.IsFailure\(\))"
```

---

## 2-Agent Parallel Orchestration

To survive large codebases without hitting step limits or context loss, execute this prompt using a strict 2-agent parallel split:

```text
+-------------------------------------------------------------------------+
| MASTER ORCHESTRATOR (Budget: N = 200)                                   |
|                                                                         |
| Phase 1 (Steps 1..100): DISCOVERY & PLANNING                            |
| +---------------------------------------------------------------------+ |
| | Sub-Agent 1: Codebase Scanner & Spec Architect                      | |
| | - Runs ripgrep queries to catalog all multi-value error returns     | |
| | - Inventories compound caller assertions (err != nil || len != N)   | |
| | - Authors master audit plan in .lovable/plans/pending/             | |
| | - Generates granular subtasks in .lovable/plans/subtasks/           | |
| +---------------------------------------------------------------------+ |
|                                                                         |
| Phase 2 (Steps 101..200): SURGICAL REFACTORING                          |
| +---------------------------------------------------------------------+ |
| | Sub-Agent 2: Code Refactorer & Outer-Layer Modernizer                | |
| | - Refactors store/repo signatures to ResultMap/ResultSlice/Result   | |
| | - Updates scanner functions to use appfault.OkMap / FailMap         | |
| | - Modernizes callers with IsCountOtherThan / HasRecord / IsDefined   | |
| | - Verifies zero regressions with targeted file linters              | |
| +---------------------------------------------------------------------+ |
+-------------------------------------------------------------------------+
```

---

## Strictly Avoid: Anti-Patterns & Prohibitions

- **NO PIECEMEAL COMMITS:** NEVER commit 1 or 2 files in isolation. Consolidate all related changes across specs, code, and indices into a single atomic commit followed immediately by `git push origin main`.
- **NO ROUTINE FULL CI/CD RUNS:** DO NOT run `06-cicd-local-runner.py` during normal turns. It executes 28-38 heavy validation gates across unrelated packages and wastes minutes. Run targeted linters only on modified files.
- **NO UNIT TEST EXECUTION OF UNRELATED PACKAGES:** Only run tests for packages directly modified (e.g. `go test ./pkg/appfault/...`).
- **NO RAW `error` RETURNS:** Never leave bare `error` as a return type on domain or store functions; always use `*appfault.AppError` or `Result[T]`.
- **NO COMPOUND CARDINALITY DISJUNCTIONS:** Never write `if res.IsFailure() || res.Count() != N` when `res.IsCountOtherThan(N)` can express the guard directly.
- **NO CONFUSING `IsSuccess()` WITH `IsDefined()`:** Do not use `IsSuccess()` when you require actual data records to be present. Use `res.IsDefined()` or `res.HasRecord()`.
- **NO ABSOLUTE PATHS:** Never write absolute filesystem paths (`C:\...`, `/home/...`) or `file:///` URIs. Use strict relative Git paths starting from the repository root.
- **NO UPPERCASE FILENAMES:** Every file created or edited must be strictly lowercase.
- **NO MULTI-VALUE TUPLES:** Eliminate `(T, error)` in favor of `Result[T]`, `ResultMap[K, V]`, or `ResultSlice[T]`.
