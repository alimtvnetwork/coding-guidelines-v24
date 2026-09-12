---
name: cg-boolean-and-naming
description: Autonomously audits, refactors, and validates repository-wide boolean conventions, positive prefixes, implicit checks, enum Type suffixes, and nested if flattening against 02-spec/02-coding-guidelines/.
---

# Skill: Coding Guidelines — Booleans, Naming & Enums (`cg-boolean`)

This skill governs autonomous execution for boolean conventions, semantic naming, enum standardization, and conditional flattening across all source files.

## Mandatory Architectural Rules

1. **Implicit Boolean Checks Only:**
   - NEVER write `if isReady == true` or `if (isValid === true)`.
   - Positive booleans MUST ALWAYS be evaluated implicitly: `if isReady { ... }` or `if !isReady { ... }`.
   - Never compare against boolean literals (`== false`, `!= true`).

2. **Boolean Prefixes (`is`, `has`) & Affirmative Naming:**
   - `is`, `has` as prefix is only acceptable and nothing else acceptable including but not limited to `can`, `should`, `was`, `will`, `did`, `must`, etc.
   - Every boolean identifier must begin with `is` or `has` (e.g. `isValid`, `hasAccess`).
   - No negative boolean identifiers (`isNotValid`, `hasNoData` are banned).
   - **Total Ban on Single-Letter Parameters:** NEVER use single-letter boolean parameters (`v bool`, `b bool`, `val bool`, `flag bool`) in method and function signatures (e.g. setters).
   - **Total Ban on Bare Unprefixed Names:** NEVER use bare verbs, nouns, or adjectives (`stop bool`, `pause bool`, `force bool`, `dryRun bool`, `header bool`).
   - **Mandatory Affirmative Prefixes:** Every boolean parameter, struct field, property, and variable MUST carry an affirmative prefix (`is*` or `has*`):
     - `stop` -> `isStopped`
     - `stopOnFail` -> `isStopOnFail` (e.g. `SetStopOnFail(isStopOnFail bool)`)
     - `defined` -> `isDefined` (e.g. struct field `isDefined bool`, method `IsDefined() bool`)
     - `pause` / `paused` -> `isPaused`
     - `force` -> `isForced` or `isForce`
     - `enable` / `enabled` -> `isEnabled`
     - `dryRun` -> `isDryRun`
     - `debug` -> `isDebug`
     - `verbose` -> `isVerbose`
     - `header` -> `hasHeader`
     - `records` -> `hasRecords`

### Generic Code Patterns (Affirmative Naming)

#### Pattern A: Setter Method Parameter & Field Assignment (`v bool` -> `isStopOnFail bool`)

```go
// ❌ ANTI-PATTERN: Single-letter parameter `v bool` and un-prefixed field
func (p *BatchProgress) SetStopOnFail(v bool) {
    p.mu.Lock()
    defer p.mu.Unlock()
    p.stopOnFail = v
}

// ✅ REQUIRED: Meaningful, affirmative boolean parameter and property
func (p *BatchProgress) SetStopOnFail(isStopOnFail bool) {
    p.mu.Lock()
    defer p.mu.Unlock()
    p.stopOnFail = isStopOnFail
}
```

#### Pattern B: Generic State Flag & Struct Worker (`stop` -> `isStopped`)

```go
// ❌ ANTI-PATTERN: Bare verb `stop` and lazy `b bool` in stateful worker
type TaskWorker struct {
    stop bool
}

func (w *TaskWorker) SetStop(b bool) {
    w.stop = b
}

func (w *TaskWorker) Run() {
    for {
        if w.stop {
            break
        }
        processTask()
    }
}

// ✅ REQUIRED: Generic affirmative `isStopped` state and parameter
type TaskWorker struct {
    isStopped bool
}

func (w *TaskWorker) SetStopped(isStopped bool) {
    w.isStopped = isStopped
}

func (w *TaskWorker) Run() {
    for {
        if w.isStopped {
            break
        }
        processTask()
    }
}
```

#### Pattern C: Struct Field Definition State (`defined bool` -> `isDefined bool`)

```go
// -----------------------------------------------------------------------------
// ❌ ANTI-PATTERN: Bare field name `defined bool` in wrapper struct
// -----------------------------------------------------------------------------
type Result[T any] struct {
    value   T
    err     *AppError
    defined bool // VIOLATION: Bare boolean without is/has prefix
}

// -----------------------------------------------------------------------------
// ✅ REQUIRED: Meaningful, affirmative boolean struct field `isDefined bool`
// -----------------------------------------------------------------------------
type Result[T any] struct {
    value     T
    err       *AppError
    isDefined bool // REQUIRED: Explicit affirmative boolean prefix
}
```

#### Pattern D: Generic Transformation Reference Table

| Target Category | ❌ Anti-Pattern (Lazy / Bare) | ✅ Required Affirmative Identifier | Context / Description |
|---|---|---|---|
| Setter Parameter | `SetStopOnFail(v bool)` | `SetStopOnFail(isStopOnFail bool)` | Early termination flag parameter |
| State Variable | `stop := false` | `isStopped := false` | Process / loop cancellation state |
| Method Parameter | `Stop(stop bool)` | `SetStopped(isStopped bool)` | State toggle parameter |
| Struct Field | `defined bool` | `isDefined bool` | Value/record definition presence indicator |
| Method Name | `Defined() bool` | `IsDefined() bool` | Definition verification predicate |
| Struct Field | `pause bool` | `isPaused bool` | Pause / suspend indicator |
| CLI / Config Flag | `force bool` | `isForced bool` | Force override flag |
| Struct Field | `dryRun bool` | `isDryRun bool` | Dry run simulation flag |
| Option Parameter | `debug bool` | `isDebug bool` | Debug mode toggle |
| Struct Field | `header bool` | `hasHeader bool` | Header presence indicator |
| Option Parameter | `records bool` | `hasRecords bool` | Records presence requirement |

3. **No Inverted Success Checks:**
   - Never invert positive success checks (e.g. `!response.isSuccess`).
   - Use explicit failure states (e.g. `response.isFail`, `isError`).

4. **Zero Tolerance for Nested `if` (Nesting Depth <= 1):**
   - No `if` statements inside another `if` block.
   - Flatten all conditionals using guard clauses and early returns.
   - Never combine mixed polarity (`if isA && !isB` -> split into separate guard clauses).

5. **Enum Suffix `Type`:**
   - All enum declarations across TypeScript, Go, PHP must end with `Type` (e.g. `UserRoleType`, `CommandStatusType`).

6. **Function and File Size Caps:**
   - Functions: <= 8 lines preferred, <= 15 lines maximum.
   - Files: <= 100 lines coding maximum (recommended <= 80 lines).
   - Zero line compression (no single-line `if/else`, no deleted blank lines).

## Validation Linters & Execution Policies

- **No Releases:** Strictly forbidden from bumping versions or cutting releases.
- **No Test Execution:** Test execution is disabled unless explicitly commanded by the repository owner.
- **Atomic Change Tracking:** Append all modified files to `.lovable/temp/recent-file-changes.json` under lock (`python 03-ai-scripts/33-test-inventory-generator.py --record <files...>`), mapping to associated tests in `.lovable/test-inventory.json`.
- **Linter:** `python linter-scripts/check-enum-and-boolean.py`
- **Local Runner:** `python 03-ai-scripts/06-cicd-local-runner.py --no-tests`

## Routine Execution Policy

- **NO FULL CI/CD RUNNER (Strict Policy):** DO NOT run `python 03-ai-scripts/06-cicd-local-runner.py` during routine coding guideline execution turns or micro-batch loops. Running the heavy 28-38 gate pipeline across the entire repository wastes massive amounts of time. Verify code strictly using targeted file-level linters / autofixers on the specific modified files.
