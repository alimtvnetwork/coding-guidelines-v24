# Master Architectural Specification: Structured Fileutil Context & Concrete Results

> **Plan Identifier:** `17-structured-fileutil-context-and-concrete-results`  
> **Status:** Completed  
> **Created:** 2026-09-07  
> **Scope:** `04-code/golang/pkg/appfault/`, `04-code/golang/pkg/result/`, `04-code/golang/pkg/fileutil/`

---

## 1. Problem Statement & User Rationale

1. **Path Concatenation Anti-Pattern in Error Messages:**
   - In `04-code/golang/pkg/fileutil/`, 35 locations concatenate filesystem paths directly into error message strings (e.g. `"failed to create parent directory: " + dir`, `"failed to open file: " + path`, `"failed to stat file: " + path`).
   - The user mandated:
     > *"Remember, the file path is not something that is going to be exposed as a message string combination. It should be, uh, a variable, okay, inside the context map. So this is how we need to add. So based on that, we have to change everywhere that we are dealing with the file system."*
2. **First-Class Single-Variable & Path Injection on AppError & Result:**
   - Instead of manual string formatting, errors must provide direct single-variable injection: `WithVar(key, val)`, `WithPath(path)`, `WithFilePath(path)`.
   - Direct constructors on `appfault`: `WrapFile`, `NewFile`, `WrapVar`, `NewVar`.
   - Result wrappers on `result`: `WrapFailurePath[T]`, `WrapFailureFile[T]`, `WrapFailureVar[T]`, `FailurePath[T]`.
3. **Verbose Bracket Generics (`result.Wrap[*os.File]`):**
   - Writing generic types with brackets (`result.Wrap[*os.File]`, `result.WrapFailure[*os.File]`) across the codebase is repetitive and error-prone for AI agents and human developers alike.
   - Dedicated concrete type aliases (`FileResult`, `BytesResult`, `StringResult`, `LinesResult`, `BoolResult`, `FileInfoResult`, `Int64Result`) and typed constructor helpers (`FileSuccess`, `FileFailure`, `BoolSuccess`, `BoolFailure`) allow ergonomic, type-safe error propagation without bracket generics.

---

## 2. Task-Specific Rules & Invariants (Auto-Reject on Violation)

1. **Rule 1 (No String Concatenation for Paths in Error Messages):** File paths MUST NEVER be appended to error message strings via `+ path` or `fmt.Sprintf("%s: %s", msg, path)`. All paths MUST be stored exclusively inside the structured context map with key `"Path"` using `WithContext("Path", path)`, `WithPath(path)`, `WrapFile(...)`, or `NewFile(...)`.
2. **Rule 2 (Zero-Bracket Concrete Types in Fileutil):** All exported and internal functions in `pkg/fileutil/` operating on standard types MUST return concrete result type aliases (`FileResult`, `BytesResult`, `StringResult`, `LinesResult`, `BoolResult`, `FileInfoResult`, `Int64Result`) instead of generic parameter declarations.
3. **Rule 3 (Strict Immutability & Nil Safety):** All `*AppError` context methods (`WithVar`, `WithPath`, `WithFilePath`) MUST safely handle nil receivers and return a cloned immutable instance.
4. **Rule 4 (Coding Guidelines Enforcement):** All functions MUST be strictly <= 15 lines. Blank line before every `return` (unless sole statement). Blank line after closing brace `}` if followed by code. No explicit true checks (`if ok == true`).
5. **Rule 5 (Strict Relative Git Paths):** All file references and links in documentation and plans MUST be strictly relative to the repository root.

---

## 3. Subtask Decomposition

| Subtask | Title | Bounded Files |
|---|---|---|
| `01-task-appfault-single-variable-and-path-constructors.md` | Single-Variable Context & Path Constructors | `04-code/golang/pkg/appfault/context.go`, `builder.go`, `constructors.go`, `result_constructors.go`, `pkg/result/result.go` |
| `02-task-fileutil-concrete-results-and-clean-context-refactor.md` | Concrete Result Types & Clean Path Context | `04-code/golang/pkg/fileutil/types.go`, `results.go`, `fileutil.go`, `create.go`, `write.go`, `parsers.go`, `exporters.go`, `stream.go`, `advanced.go`, `locked_operations.go` |

---

## 4. Verification Plan

1. `go test -C 04-code/golang -v ./pkg/appfault`
2. `go test -C 04-code/golang -v ./pkg/result`
3. `go test -C 04-code/golang -v ./pkg/fileutil`
4. `go test -C 04-code/golang -count=1 ./...` across all 17 Go packages
5. `python 03-ai-scripts/06-cicd-local-runner.py --all` (all 31 gates green)
