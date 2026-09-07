# Subtask 01: Single-Variable Context & Path Constructors

> **Parent Plan:** `17-structured-fileutil-context-and-concrete-results`  
> **Status:** Completed  
> **Bounded Files:**  
> - `04-code/golang/pkg/appfault/context.go`  
> - `04-code/golang/pkg/appfault/builder.go`  
> - `04-code/golang/pkg/appfault/constructors.go`  
> - `04-code/golang/pkg/appfault/result_constructors.go`  
> - `04-code/golang/pkg/result/result.go`  
> - `04-code/golang/pkg/appfault/context_test.go`  
> - `04-code/golang/pkg/result/wrap_test.go`

---

## Instructions

1. **In `04-code/golang/pkg/appfault/context.go`:**
   - Add `WithVar(key string, value any) *AppError` (delegates to `WithContext`).
   - Add `WithPath(path string) *AppError` (delegates to `WithContext("Path", path)`).
   - Add `WithFilePath(path string) *AppError` (delegates to `WithPath(path)`).
   - Add `WithField(key string, value any) *AppError` (delegates to `WithContext(key, value)`).
   - Verify nil receiver safety and immutability.

2. **In `04-code/golang/pkg/appfault/builder.go`:**
   - Add `WithVar(key string, value any) *AppErrorBuilder` (delegates to `SetContext`).
   - Add `WithPath(path string) *AppErrorBuilder` (delegates to `SetContext("Path", path)`).
   - Add `WithFilePath(path string) *AppErrorBuilder` (delegates to `WithPath(path)`).
   - Add `WithField(key string, value any) *AppErrorBuilder` (delegates to `SetContext(key, value)`).

3. **In `04-code/golang/pkg/appfault/constructors.go`:**
   - Add `WrapFile(variation errtype.Variation, cause error, path string, msg string, skipFrames ...int) *AppError`.
   - Add `NewFile(variation errtype.Variation, path string, msg string, skipFrames ...int) *AppError`.
   - Add `WrapPath(variation errtype.Variation, cause error, path string, msg string, skipFrames ...int) *AppError`.
   - Add `NewPath(variation errtype.Variation, path string, msg string, skipFrames ...int) *AppError`.
   - Add `WrapVar(variation errtype.Variation, cause error, key string, val any, msg string, skipFrames ...int) *AppError`.
   - Add `NewVar(variation errtype.Variation, key string, val any, msg string, skipFrames ...int) *AppError`.

4. **In `04-code/golang/pkg/appfault/result_constructors.go`:**
   - Add `NewFailureWithFile[T any](variation errtype.Variation, cause error, path string, msg string) Result[T]`.
   - Add `NewFailureWithPath[T any](variation errtype.Variation, cause error, path string, msg string) Result[T]`.
   - Add `NewFailureWithVar[T any](variation errtype.Variation, cause error, key string, val any, msg string) Result[T]`.
   - Add `FailurePath[T any](variation errtype.Variation, path string, msg string) Result[T]`.
   - Add `FailureFile[T any](variation errtype.Variation, path string, msg string) Result[T]`.

5. **In `04-code/golang/pkg/result/result.go`:**
   - Re-export `WrapFailurePath[T any]`, `WrapFailureFile[T any]`, `WrapFailureVar[T any]`, `FailurePath[T any]`, `FailureFile[T any]`.

6. **Unit Tests:**
   - Add comprehensive tests in `04-code/golang/pkg/appfault/context_test.go` and `04-code/golang/pkg/result/wrap_test.go`.
   - Run `go test -C 04-code/golang -v ./pkg/appfault ./pkg/result`.
