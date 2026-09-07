# Master Architectural Specification: Fileutil Struct Grouping & Creator Modular Decomposition

> **Plan Identifier:** `19-fileutil-struct-grouping-and-modular-decomposition`  
> **Status:** Completed  
> **Created:** 2026-09-07  
> **Scope:** `04-code/golang/pkg/fileutil/`

---

## 1. Problem Statement & User Rationale

1. **Modular Struct Grouping of Operations:**
   - Functions in `04-code/golang/pkg/fileutil/` are currently dispersed across monolithic files (`fileutil.go`, `create.go`, `write.go`, `parsers.go`, etc.).
   - The user mandated:
     > *"for file util let's try to keep each func or struct in each file , keep struct with it;s methods but I do think all the open, create, write, append , read methods needs to be grouped together like struct using what is done for core package create concepts... D:\work\03-aukgo\core\coredata"*
2. **Adopting the Coredata Creator Concept:**
   - In `core/coredata/corestr/` (`newCreator.go`, `newCollectionCreator.go`, `vars.go`), operations and constructors are grouped into zero-allocation sub-structs and exposed via clean namespace singletons: `var New = &newCreator{}` and structured sub-namespaces.
   - We will establish two cohesive namespaces:
     - `var File = &fileNamespace{}`: Operational namespace grouping `.Open.*`, `.Create.*`, `.Write.*`, `.Append.*`, `.Read.*`, and `.Path.*`.
     - `var New = &fileNewCreator{}`: Creation namespace grouping constructors for `Writer`, `Appender`, `BoundWriter`, `Path`, and `StreamWriter`.
3. **Eradicating Remaining Bracket Generics:**
   - The user highlighted a lingering bracket generic in `fileutil.go:55`: `return result.WrapFailure[*os.File](err)`!
   - We will replace all residual `result.WrapFailure[T]` calls in non-generic functions with concrete constructors (`FileFailure`, `BoolFailure`, `StringFailure`, and companion fault forwarders `FileFailureFault`, `BoolFailureFault`, `BytesFailureFault`, etc.).
4. **Function Size Compliance:**
   - Split all functions exceeding 15 lines into focused, clean helpers (<= 15 lines per function).
5. **100% Backward Compatibility:**
   - All existing package-level functions (`OpenFile`, `Open`, `CreateFile`, `CreateDir`, `EnsureDir`, `WriteBytes`, `WriteString`, `ReadAll`, `ReadString`, etc.) are retained as forwarders to preserve zero breaking changes for existing consumers and tests.

---

## 2. Task-Specific Rules & Invariants (Auto-Reject on Violation)

1. **Rule 1 (Zero-Bracket Concrete Types in Fileutil Operations):** No non-generic function in `pkg/fileutil/` may call bracketed generic constructors like `result.WrapFailure[*os.File](...)`. All errors must return concrete aliases (`FileResult`, `BytesResult`, `StringResult`, `LinesResult`, `BoolResult`, `FileInfoResult`, `Int64Result`) using `FileFailure`, `FileFailureFault`, `BoolFailureFault`, etc.
2. **Rule 2 (Single-Responsibility Struct Grouping):** Keep each functional domain in its own dedicated source file:
   - `file_open.go` (`openOps`)
   - `file_create.go` (`createOps`)
   - `file_write.go` (`writeOps`)
   - `file_append.go` (`appendOps`)
   - `file_read.go` (`readOps`)
   - `file_namespace.go` (`var File`, `var New`)
3. **Rule 3 (Coredata Creator Conformance):** Sub-operation structs (`openOps`, `createOps`, etc.) and sub-creators (`fileWriterCreator`, etc.) must be zero-allocation empty structs matching the `coredata` creator pattern.
4. **Rule 4 (Strict 15-Line Function Limit & Blank Lines):** Every function and helper must be strictly <= 15 lines. Blank line after closing brace `}` if followed by code. Blank line before `return` unless sole statement.
5. **Rule 5 (Strict Zero-Panic & Relative Git Paths):** Zero panics, zero `log.Fatal()`. All links and paths in documentation and plans must be strictly relative to the git repository root.

---

## 3. Subtask Decomposition

| Subtask | Title | Bounded Files |
|---|---|---|
| `01-task-fileutil-operations-decomposition-and-concrete-results.md` | Operation Structs Decomposition & Concrete Results | `04-code/golang/pkg/fileutil/file_open.go`, `file_create.go`, `file_append.go`, `file_read.go`, `file_write.go`, `results.go`, `fileutil.go` |
| `02-task-fileutil-namespace-singletons-creator-and-tests.md` | File & New Namespace Singletons, Tests & Docs | `04-code/golang/pkg/fileutil/file_namespace.go`, `file_namespace_test.go`, `file_append_test.go`, `readme.md` |

---

## 4. Verification Plan

1. `go test -C 04-code/golang -v ./pkg/fileutil`
2. `go test -C 04-code/golang -count=1 ./...` across all 17 Go packages
3. `node scripts/sync-check.mjs`
4. `python 03-ai-scripts/06-cicd-local-runner.py --all` (all 31 gates green)
