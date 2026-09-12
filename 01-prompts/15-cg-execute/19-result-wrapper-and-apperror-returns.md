# Result Wrapper Types, Collections & AppError Returns — Coding Guideline (must follow)

Trigger Keywords & Aliases: `cg-result-wrapper`, `cg-apperror-returns`, `cg-execute result-wrapper`, `audit result wrapper`, `fix map return error`, `fix slice return error`, `single return object audit`, `enforce apperror returns`, `enforce result map`, `fix multi-value returns`, `is-count-other-than`, `has-record`, `is-defined`, `result-wrapper-null-safety`, `pointer-null-safety`

> **Prompt Version:** 2.3.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

```text
N = 200
```

N = total self-loop steps budget that the agents will perform.

/goal Autonomously scan, discover, plan, refactor, and verify all Go functions returning multi-value error tuples (such as `(map[K]V, error)`, `([]T, error)`, or `(T, error)`), eliminating raw standard library error returns, replacing them with strongly-typed result wrappers (`ResultMap[K, V]`, `ResultSlice[T]`, `Result[T]`) and structured `*appfault.AppError` returns, guaranteeing a single return object, pointer-attached null safety (`*Result[T]`, `*ResultSlice[T]`, `*ResultMap[K, V]`) with methods attached to pointer receivers (`(r *Result[T])`, `(rs *ResultSlice[T])`, `(rm *ResultMap[K, V])`) that verify `if r == nil` before dereferencing any fields or checking errors, standardized outer-layer inspection predicates (`IsSuccess`, `IsFailure`, `HasError`, `IsEmptyError`, `IsEmpty`, `HasRecord`, `IsDefined`, `IsCountOtherThan`, `Data`, `Items`, `AppError`, `Fault`, `Get`, `Has`, `Count`), eliminating dual-handling, and replacing verbose `if err != nil || len(...) != N` or `IsFailure() || Count() != N` conditions with fluent `if res.IsCountOtherThan(N)` across the entire codebase until 100% green without stopping.

### Master Task Checklist (Atomic Numbered Steps)

1. [ ] /goal Phase 1 (Step A): Deeply scan the target codebase using ripgrep to inventory all functions returning multi-value tuples `(T, error)`, `(map[K]V, error)`, `([]T, error)`, raw stdlib `error` returns, and any Result methods declared with value receivers `func (r Result[...])` lacking pointer-attached null safety. Also inventory clumsy caller checks like `err != nil || len(...) != N` and `IsFailure() || Count() != N`.
2. [ ] /goal Phase 1 (Step B): Write the master audit specification in `.lovable/plans/pending/XX-result-wrapper-audit.md` with an exhaustive Violation Ledger table.
3. [ ] /goal Phase 1 (Step C): Decompose the master plan into granular, atomic subtasks in `.lovable/plans/subtasks/XX-result-wrapper/`.
4. [ ] /goal Phase 1 (Step D): Verify or create the automated quality linter and register in `03-ai-scripts/01-index.md`.
5. [ ] /goal Phase 2 (Step A): Open each target file and refactor function signatures from multi-value returns to single `ResultMap[K, V]`, `ResultSlice[T]`, or `Result[T]` envelopes.
6. [ ] /goal Phase 2 (Step B): Replace raw stdlib `error` returns with structured `*appfault.AppError` instances using `appfault.New()` or `appfault.Wrap()`.
7. [ ] /goal Phase 2 (Step C): Enforce pointer-attached null safety on all Result wrappers: attach all inspection methods (`IsSuccess`, `IsFailure`, `IsEmpty`, `HasRecord`, `IsDefined`, `IsCountOtherThan`, `Count`, `AppError`, `Data`, `Items`) to pointer receivers (`(r *Result[T])`, `(rs *ResultSlice[T])`, `(rm *ResultMap[K, V])`) with explicit `nil` checks (`if r == nil`) guarding against nil pointer panics and returning safe defaults.
8. [ ] /goal Phase 2 (Step D): Modernize all caller call sites to utilize outer-layer inspection methods (`res.IsSuccess()`, `res.IsFailure()`, `res.IsEmpty()`, `res.HasRecord()`, `res.IsDefined()`, `res.IsCountOtherThan(N)`, `res.Get()`, `res.AppError()`), eliminating manual `err != nil || len(...) != N` boilerplate.
9. [ ] /goal Phase 2 (Step E): Enforce <= 8–15 line function decomposition and clean blank-line spacing.
10. [ ] /goal Phase 2 (Step F): Execute targeted file-level linters (`python linter-scripts/check-function-lengths.py`, `check-mws-error-codes.py`, `check-newline-styling.py`) to verify 0 remaining violations. DO NOT run the full CI/CD pipeline runner (`06-cicd-local-runner.py`) during routine coding guideline execution turns.
11. [ ] /learn Ingest `.lovable/memory/01-index.md` for project memory index and past learnings.
12. [ ] /learn Ingest `.lovable/strictly-avoid.md` for banned anti-patterns and strict constraints.
13. [ ] /learn Ingest `02-spec/02-coding-guidelines/02-canonical-size-tier.md` for canonical file and function size tiers.
14. [ ] /learn Ingest `02-spec/02-coding-guidelines/01-cross-language/01-index.md` for single return type mandates and micro-tasking.
15. [ ] /learn Ingest `02-spec/02-coding-guidelines/01-cross-language/01-index.md` for strict relative path citation requirements.
16. [ ] /learn Ingest `02-spec/03-error-manage/01-index.md` for universal AppError wrapping and error envelopes.
17. [ ] /learn Ingest `02-spec/03-error-manage/02-error-architecture/02-error-handling-reference.md` for error handling architecture and Result wrappers.
18. [ ] /learn Ingest `02-spec/03-error-manage/03-error-code-registry/02-registry.md` for structured error code catalog.
19. [ ] /learn Ingest `02-spec/03-error-manage/02-error-architecture/05-response-envelope/05-response-envelope-reference.md` for response envelope schemas.
20. [ ] /learn Ingest `02-spec/03-error-manage/02-error-architecture/06-apperror-package/03-go-apperror-linter-spec.md` for Go AppError implementation specifications.
21. [ ] /learn Ingest `02-spec/03-error-manage/02-error-architecture/06-apperror-package/01-apperror-reference/04-result-types.md` for Result[T], ResultSlice[T], and ResultMap[K, V] method specifications and pointer null-safety rules.
22. [ ] /learn Ingest `.lovable/coding-guidelines.md` for master consolidated coding guidelines.
23. [ ] /goal Create or update agent rules in the repository if missing from agent memory.

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

Result envelopes (`ResultSlice[T]`, `ResultMap[K, V]`, `Result[T]`) provide four core predicate methods that eliminate call-site boilerplate, null pointer panics, and compound boolean conditions:

### 1. `res.IsCountOtherThan(number int) bool`

- **Exact Semantics:** Returns `true` if the operation failed (has error or nil receiver) **OR** if the record count is not equal to `number`.
- **Purpose:** Replaces the compound check `if res.IsFailure() || res.Count() != N` (and legacy `if err != nil || len(...) != N`) with a single, intention-revealing predicate.
- **Behavior:**
  - If `res == nil` or `res.IsFailure()`: returns `true`.
  - If `res.IsSuccess()`: returns `res.Count() != number`.
- **Example:**
  ```go
  // ✅ Clean single-condition guard replacing err != nil || len(users) != 1
  userRes := userRepo.FindById(userId)
  if userRes.IsCountOtherThan(1) {
      return appfault.New(appfault.ErrNotFound).WithMessage("expected exactly 1 user")
  }
  ```

### 2. `res.IsEmpty() bool`

- **Exact Semantics:** Returns `true` if the collection contains `0` items/entries, if the scalar payload `T` is empty/null/zero, or if the receiver is `nil`.
- **Purpose:** Cleanly checks for zero records after validating success, without fragile `len()` checks.
- **Behavior:**
  - If `res == nil`: returns `true`.
  - If `res.IsFailure()`: returns `true` (failed results contain zero valid records).
  - If `res.IsSuccess()`: returns `res.Count() == 0`.
- **Example:**
  ```go
  orderRes := orderService.ListPendingOrders(ctx)
  if orderRes.IsFailure() {
      return orderRes.AppError()
  }
  if orderRes.IsEmpty() {
      logger.Info("no pending orders to process")
      return nil
  }
  ```

### 3. `res.HasRecord() bool` (and alias `res.HasRecords() bool`)

- **Exact Semantics:** Returns `true` if the operation succeeded (no error) **AND** contains **more than 0 records** (`res.Count() > 0 && !res.IsFailure()`).
- **Purpose:** Direct positive check when business logic requires at least one record before continuing, avoiding inverted `!IsEmpty()` logic.
- **Behavior:**
  - If `res == nil` or `res.IsFailure()`: returns `false`.
  - If `res.IsSuccess()`: returns `res.Count() > 0`.
- **Example:**
  ```go
  itemRes := catalog.QueryItemsByCategory(catId)
  if itemRes.HasRecord() {
      dispatcher.EnqueueBatch(itemRes.Items)
  }
  ```

### 4. `res.IsDefined() bool`

- **Exact Semantics:** Returns `true` if the operation succeeded (no error) **AND** has `recordCount > 0` (or the underlying data `T` is non-null/non-empty).
- **Distinction from `IsSuccess()`:**
  - `IsSuccess()` means "no error occurred" (an empty query returning 0 items succeeds without error).
  - `IsDefined()` means "no error occurred AND actual data exists" (`recordCount > 0` or payload not null/empty).
- **Behavior:**
  - If `res == nil` or `res.IsFailure()`: returns `false`.
  - For `Result[T]`: returns `r.isDefined && r.IsSuccess() && !isValueEmpty(r.value)` (delegates error check to `r.IsSuccess()`).
  - For `ResultSlice[T]`: returns `rs.Count() > 0 && rs.IsSuccess()` (delegates to `rs.Count()` and `rs.IsSuccess()`).
  - For `ResultMap[K, V]`: returns `rm.Count() > 0 && rm.IsSuccess()` (delegates to `rm.Count()` and `rm.IsSuccess()`).
- **Example:**
  ```go
  profileRes := userProfileService.GetProfile(userId)
  if profileRes.IsDefined() {
      displayProfileBadge(profileRes.Value())
  }
  ```

---

## Pointer-Attached Null Safety (*Result[T], *ResultSlice[T], *ResultMap[K, V])

### The Fatal Flaw of Value Receivers in Go

In Go, declaring methods with a **value receiver** (`func (r Result[T]) Method()`) creates an inescapable runtime crash vulnerability:
If a caller has a `nil` pointer to a result (`var res *Result[T] = nil`), calling `res.IsFailure()` or `res.Count()` **panics immediately** with:
```text
panic: runtime error: invalid memory address or nil pointer dereference
```
This panic happens **before the method body even begins executing**, because the Go runtime must evaluate `*res` to create a value copy for the receiver.

By contrast, declaring methods with a **pointer receiver** (`func (r *Result[T]) Method()`) passes the pointer itself directly. If `r == nil`, the method body executes normally and can guard itself on line 1:
```go
if r == nil {
    return true // Safe default, ZERO panic!
}
```

Furthermore, in Go, methods declared on a pointer receiver `(r *T)` can still be called directly on an addressable value `T` (`res := Ok(val); res.IsSuccess()`) because the Go compiler automatically passes `&res`. Therefore, pointer receivers provide 100% backward compatibility while providing complete immunity against nil pointer crashes!

### Canonical Nil Receiver Defaults Table

When any inspection method is invoked on a `nil` pointer (`(*Result[T])(nil)`, `(*ResultSlice[T])(nil)`, or `(*ResultMap[K, V])(nil)`), it MUST never panic and MUST return these canonical safe defaults:

| Method | Return on `nil` Pointer | Rationale & Semantic Behavior |
|---|---|---|
| `r.IsFailure()` / `r.IsFailed()` | `true` | An uninitialized/missing result is an error/failure state. |
| `r.IsSuccess()` / `r.IsSafe()` | `false` | A nil pointer cannot represent a successful operation. |
| `r.HasError()` | `true` | Alias for `IsFailed()`. |
| `r.IsEmptyError()` / `r.HasNoError()` | `false` | A nil result is not error-free. |
| `r.Count()` | `0` | A nil result contains zero records. |
| `r.IsEmpty()` | `true` | A nil result contains no elements. |
| `r.HasRecord()` / `r.HasRecords()` | `false` | A nil container has 0 records, never > 0. |
| `r.IsDefined()` | `false` | A nil container has no defined data payload. |
| `r.IsCountOtherThan(number)` | `true` | A nil/failed result differs from any expected record count. |
| `r.AppError()` / `r.Fault()` | `nil` | Safely returns nil error without crashing. |
| `r.Value()` / `r.Data()` | `zero value of T` | Safely returns zero value of type `T`. |
| `rs.Items()` / `rs.Data` | `nil` | Safely returns nil slice. |
| `rm.Get(key)` | `zero, false` | Safely returns zero value and `false` indicating missing key. |
| `rm.Has(key)` | `false` | Key cannot exist in a nil map. |

### Pointer-Attached Implementation Standard (`pkg/appfault/`)

Every result wrapper struct and inspection method in `pkg/appfault` MUST follow these two architectural rules:
1. **Affirmative Boolean Field Naming:** Boolean fields MUST use affirmative prefixes (e.g. `isDefined bool`, NEVER bare `defined bool`).
2. **Method Composition & Reuse:** Methods MUST delegate to and compose existing inspection methods (`r.IsFailure()`, `r.IsSuccess()`, `r.Count()`) rather than repeating raw pointer and error checks (`r == nil || r.err != nil`).

```go
type Result[T any] struct {
    value     T
    err       *AppError
    isDefined bool // ✅ REQUIRED: Affirmative boolean prefix (TOTAL BAN on bare 'defined')
}

// ✅ POINTER-ATTACHED & NULL-SAFE: Inspecting nil pointer returns false, never panics!
func (r *Result[T]) IsSuccess() bool {
    if r == nil {
        return false
    }

    return r.err == nil
}

// ✅ POINTER-ATTACHED & NULL-SAFE: Inspecting nil pointer returns true, never panics!
func (r *Result[T]) IsFailed() bool {
    if r == nil {
        return true
    }

    return r.err != nil
}

// ✅ METHOD COMPOSITION: Delegates alias directly to IsFailed()
func (r *Result[T]) IsFailure() bool {
    return r.IsFailed()
}

// ✅ METHOD COMPOSITION: Delegates alias directly to IsSuccess()
func (r *Result[T]) IsValid() bool {
    return r.IsSuccess()
}

// ✅ METHOD COMPOSITION: Reuses r.IsFailure() to eliminate redundant null checks
func (r *Result[T]) Count() int {
    if r.IsFailure() {
        return 0
    }

    if r.isDefined {
        return 1
    }

    return 0
}

// ✅ METHOD COMPOSITION: Reuses r.IsFailure() and r.Count()
func (r *Result[T]) IsCountOtherThan(expected int) bool {
    if r.IsFailure() {
        return true
    }

    return r.Count() != expected
}

// ✅ METHOD COMPOSITION: Reuses r.IsFailure() and checks affirmative isDefined field
func (r *Result[T]) IsEmpty() bool {
    if r.IsFailure() {
        return true
    }

    return !r.isDefined || isValueEmpty(r.value)
}

// ✅ METHOD COMPOSITION: Reuses r.IsFailure() and r.Count()
func (r *Result[T]) HasRecord() bool {
    if r.IsFailure() {
        return false
    }

    return r.Count() > 0
}

// ✅ METHOD COMPOSITION: Delegates alias directly to HasRecord()
func (r *Result[T]) HasRecords() bool {
    return r.HasRecord()
}

// ✅ METHOD COMPOSITION: Reuses r.IsFailure() and checks affirmative isDefined field
func (r *Result[T]) IsDefined() bool {
    if r.IsFailure() {
        return false
    }

    return r.isDefined && !isValueEmpty(r.value)
}
```

### Dedicated Rule: Affirmative Field Naming (`isDefined bool`) & Method Composition

- **TOTAL BAN on bare `defined bool`:** Struct fields, properties, local variables, and parameters MUST ALWAYS carry an affirmative prefix (`is*` or `has*`). In Result wrappers, the definition state MUST be named `isDefined bool` (NEVER `defined bool`).
- **TOTAL BAN on bare `Defined()` method:** Predicate methods MUST be named `IsDefined() bool` (NEVER `Defined() bool`).
- **Method Composition Mandate:** Never write duplicate expressions like `if r == nil || r.err != nil` across 10 different methods. Always call `if r.IsFailure()` or `if !r.IsSuccess()`. Composing methods ensures single-point maintenance, prevents cognitive drift, and enforces consistent null-safety semantics.

### Go Addressability Rule for Callers

In Go, methods declared on pointer receivers `*T` can be called on:
1. Pointers directly: `ptr := &res; ptr.IsSuccess()` or `var ptr *Result[T]; ptr.IsFailure()`
2. Addressable value variables: `res := store.Query(); if res.IsSuccess() { ... }` (compiler passes `&res`)
3. Slice elements and struct fields: `results[0].IsSuccess()`, `s.result.IsSuccess()`

**Important Caller Rule:** In Go, you cannot call a pointer method on an *unaddressable temporary expression* directly (e.g. `store.Query().IsSuccess()` will not compile if `Query()` returns by value). Callers MUST assign the result to a variable first:
```go
// ❌ COMPILE ERROR (if Query returns Result[T] by value):
if store.Query().IsSuccess() { ... }

// ✅ CORRECT: Assign to variable first (variable is addressable)
res := store.Query()
if res.IsSuccess() { ... }
```

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
// ✅ 3. CANONICAL MODERN PATTERN: Pointer-safe single expressive guard
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

All methods below are declared on **pointer receivers** (`*Result[T]`, `*ResultSlice[T]`, `*ResultMap[K, V]`) and guarantee complete **null safety** (zero runtime panics when invoked on `nil` pointers):

| Method | Return Type | Receiver | Purpose & Exact Behavior (Safe on `nil`) |
|---|---|---|---|
| `res.IsSuccess()` | `bool` | `*Result`, `*ResultSlice`, `*ResultMap` | Returns `true` if operation succeeded with no error. (`false` on `nil`). |
| `res.IsFailure()` / `res.IsFailed()` | `bool` | `*Result`, `*ResultSlice`, `*ResultMap` | Returns `true` if operation encountered an error or receiver is `nil`. |
| `res.HasError()` | `bool` | `*Result`, `*ResultSlice`, `*ResultMap` | Alias for `IsFailed()`. Returns `true` if error is present or `nil`. |
| `res.IsEmptyError()` / `res.HasNoError()` | `bool` | `*Result`, `*ResultSlice`, `*ResultMap` | Returns `true` if receiver is non-nil and has no active error. |
| `res.IsCountOtherThan(number)` | `bool` | `*Result`, `*ResultSlice`, `*ResultMap` | Returns `true` if operation failed (or `nil`) OR count != number. |
| `res.IsEmpty()` | `bool` | `*Result`, `*ResultSlice`, `*ResultMap` | Returns `true` if collection has 0 items, payload is empty/null, or `nil`. |
| `res.HasRecord()` / `res.HasRecords()` | `bool` | `*Result`, `*ResultSlice`, `*ResultMap` | Returns `true` if succeeded AND count > 0 (more than 0 records). |
| `res.IsDefined()` | `bool` | `*Result`, `*ResultSlice`, `*ResultMap` | Returns `true` if succeeded AND recordCount > 0 / non-null payload. |
| `res.Count()` | `int` | `*Result`, `*ResultSlice`, `*ResultMap` | Returns total number of records/entries (0 if failed or `nil`). |
| `res.Data` / `res.Items` / `res.Value()` | `T` / `[]T` / `map[K]V` | `*Result`, `*ResultSlice`, `*ResultMap` | Accesses the underlying data payload directly (zero value on `nil`). |
| `res.AppError()` / `res.Fault()` | `*appfault.AppError` | `*Result`, `*ResultSlice`, `*ResultMap` | Retrieves structured error context (`nil` on `nil` receiver). |
| `res.Get(key)` | `(V, bool)` | `*ResultMap[K, V]` | Safely retrieves map entry by key without nil-map panics. |
| `res.Has(key)` | `bool` | `*ResultMap[K, V]` | Checks whether a key exists within the result map (`false` on `nil`). |
| `res.Keys()` | `[]K` | `*ResultMap[K, V]` | Returns deterministically sorted slice of all map keys (`nil` on `nil`). |
| `res.Values()` | `[]V` | `*ResultMap[K, V]` | Returns slice of map values ordered according to `Keys()`. |
| `res.Filter(predicate)` | `ResultSlice[T]` | `*ResultSlice[T]` | Returns filtered slice matching predicate (or self if failed/nil). |
| `res.ForEach(fn)` | `ResultSlice[T]` | `*ResultSlice[T]` | Iterates over elements with early exit via `ForEachBreak`. |

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
  - Mandatory implementation of pointer-attached null safety and the 4 core predicates (`IsCountOtherThan`, `IsEmpty`, `HasRecord`, `IsDefined`) on all result containers.

---

## Automated Codebase Scanning Guide

Use these exact `ripgrep` regex commands to discover legacy multi-value return patterns, clumsy caller checks, and value-receiver anti-patterns across the codebase:

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

# 8. Find value receiver declarations on Result types (violates pointer null safety):
rg --pcre2 "func\s+\([a-zA-Z0-9_]+\s+Result(?:Slice|Map)?\["
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
| | - Detects value receiver declarations lacking pointer null safety   | |
| | - Authors master audit plan in .lovable/plans/pending/             | |
| | - Generates granular subtasks in .lovable/plans/subtasks/           | |
| +---------------------------------------------------------------------+ |
|                                                                         |
| Phase 2 (Steps 101..200): SURGICAL REFACTORING                          |
| +---------------------------------------------------------------------+ |
| | Sub-Agent 2: Code Refactorer & Outer-Layer Modernizer                | |
| | - Refactors store/repo signatures to ResultMap/ResultSlice/Result   | |
| | - Attaches methods to pointer receivers with nil-safety guards      | |
| | - Updates scanner functions to use appfault.OkMap / FailMap         | |
| | - Modernizes callers with IsCountOtherThan / HasRecord / IsDefined   | |
| | - Verifies zero regressions with targeted file linters              | |
| +---------------------------------------------------------------------+ |
+-------------------------------------------------------------------------+
```

---

## Strictly Avoid: Anti-Patterns & Prohibitions

- **NO VALUE RECEIVERS FOR RESULT INSPECTION METHODS:** NEVER define inspection methods on value receivers `func (r Result[T])`. ALL methods checking status, error, count, or data MUST be attached to pointer receivers `(r *Result[T])`, `(rs *ResultSlice[T])`, `(rm *ResultMap[K, V])` with mandatory `if r == nil` guards to eliminate nil-pointer dereference panics.
- **NO UNGUARDED FIELD ACCESS ON NIL POINTERS:** Never access `.Data`, `.items`, or `.err` directly on a pointer without verifying `r == nil` or calling pointer-safe inspection methods (`res.IsFailure()`, `res.Count()`, `res.IsDefined()`).
- **NO PIECEMEAL COMMITS:** NEVER commit 1 or 2 files in isolation. Consolidate all related changes across specs, code, and indices into a single atomic commit followed immediately by `git push origin main`.
- **NO ROUTINE FULL CI/CD RUNS:** DO NOT run `06-cicd-local-runner.py` during normal turns. It executes 28-38 heavy validation gates across unrelated packages and wastes minutes. Run targeted linters only on modified files.
- **NO UNIT TEST EXECUTION OF UNRELATED PACKAGES:** Only run tests for packages directly modified (e.g. `go test ./pkg/appfault/...`).
- **NO RAW `error` RETURNS:** Never leave bare `error` as a return type on domain or store functions; always use `*appfault.AppError` or `Result[T]`.
- **NO COMPOUND CARDINALITY DISJUNCTIONS:** Never write `if res.IsFailure() || res.Count() != N` when `res.IsCountOtherThan(N)` can express the guard directly.
- **NO CONFUSING `IsSuccess()` WITH `IsDefined()`:** Do not use `IsSuccess()` when you require actual data records to be present. Use `res.IsDefined()` or `res.HasRecord()`.
- **NO ABSOLUTE PATHS:** Never write absolute filesystem paths (`C:\...`, `/home/...`) or `file:///` URIs. Use strict relative Git paths starting from the repository root.
- **NO UPPERCASE FILENAMES:** Every file created or edited must be strictly lowercase.
- **NO MULTI-VALUE TUPLES:** Eliminate `(T, error)` in favor of `Result[T]`, `ResultMap[K, V]`, or `ResultSlice[T]`.

