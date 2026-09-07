# Master Plan: Generic Typecast, Result Checkers, and Reflection Optimization

## Executive Summary
This architectural plan establishes a unified, bulletproof casting, conversion, and validation framework across Go packages (`pkg/typecast`, `pkg/appfault`, `pkg/result`, `pkg/payloadconv`, and `pkg/fileutil`). It introduces a complete suite of `Checker` interfaces, optimizes `ReflectSetTo` to eliminate reflection overhead on hot paths, equips `Result[T]` and `*AppError` with generic casting capabilities, and verifies thread-safe file handling without panics.

## Task-Specific Rule Set
1. **Rule C1 — Checker Interface Suffix:** All validation interfaces that check boolean states MUST end with the `Checker` suffix (`IsSuccessChecker`, `IsFailureChecker`, `IsInvalidChecker`, `IsNullChecker`, `IsEmptyChecker`, `IsDefinedChecker`).
2. **Rule C2 — Standard Error Boundary in Internal Typecast:** Package `pkg/typecast` MUST return standard Go `error` internally to avoid cyclic import dependencies with `appfault`. The outer consumption layers (`appfault`, `result`, `payloadconv`) wrap these errors into `*appfault.AppError`.
3. **Rule C3 — Fast-Path Reflection Optimization:** `ReflectSetTo` MUST check primitive types (`string`, `int`, `int64`, `bool`, `float64`, `[]byte`) and byte slice bridges using direct type-switch assertions before falling back to `reflect.ValueOf()`/`reflect.TypeOf()`.
4. **Rule C4 — Zero Panic Policy:** Conversion, serialization, and file I/O operations MUST never panic. All failures must propagate as error returns (`*AppError` or `Result[T]`).
5. **Rule C5 — Strict Relative Paths:** All file paths, markdown links, and citations in plans and specs must be strictly relative to the repository root.

## Architecture Breakdown

### 1. Checker Interface Family (`pkg/appfault/interfaces.go`)
Define the complete suite of boolean state checkers:
- `IsSuccessChecker`: `IsSuccess() bool`
- `IsFailureChecker`: `IsFailure() bool`
- `IsInvalidChecker`: `IsInvalid() bool`
- `IsNullChecker`: `IsNull() bool`
- `IsEmptyChecker`: `IsEmpty() bool`
- `IsDefinedChecker`: `IsDefined() bool`
- `DefinableChecker`: `IsDefinedChecker` + `IsEmptyChecker`
- `StatusChecker`: `IsSuccessChecker` + `IsFailureChecker`

Implement missing methods:
- On `Result[T]` / `Wrap[T]`: Add `IsDefined() bool` (`return r.IsSuccess()`).
- On `*AppError`: Add `IsFailure() bool` (`return e.IsFailed()`).

### 2. Bulletproof Typecast Suite (`pkg/typecast/cast.go`)
Expand `pkg/typecast` with high-performance conversion methods:
- `ReflectSetTo(from, toPointer any) error`:
  - Fast-path for nil combinations
  - Fast-path type switch for primitive pointer targets (`*string`, `*int`, `*int64`, `*bool`, `*float64`, `*[]byte`)
  - Direct type assertions for `[]byte` unmarshal and `*[]byte` marshal
  - Pre-allocated singleton reflect types for byte slices
  - Fallback reflection for complex structs and pointer-to-pointer matching
- `ReflectTo[T any](payload any) (T, error)`
- `CastTo[T any](payload any) (T, error)`
- `ToBytes(payload any) ([]byte, error)`
- `ToJSON(payload any) ([]byte, error)`
- `ToJSONString(payload any) (string, error)`

### 3. Generic Casting & Serialization in Result and AppError
- In `pkg/appfault`:
  - `CastTo[T any](source any) (T, *AppError)`
  - `CastContextPayload[T any](e *AppError, key string) (T, *AppError)`
  - `CastResultData[T any, U any](r Result[T]) Result[U]`
  - `ResultToJSON[T any](r Result[T]) Result[[]byte]`
  - `ResultToBytes[T any](r Result[T]) Result[[]byte]`

### 4. Fileutil Verification
- Confirm path-level mutex with reference-counted eviction (`GetFileLock`/`ReleaseFileLock`).
- Confirm non-panicking file writers (`Write`, `WriteLocked`, `WriteJSON`, etc.).

### 5. Performance Pointers for `ReflectSetTo`
Document key performance optimizations for the user's legacy `ReflectSetFromTo` design:
1. Fast-path type switches for primitive types bypass `reflect.ValueOf` heap allocations.
2. Package-level singleton reflect types eliminate per-call `reflect.TypeOf` overhead.
3. Pointer validation using interface type assertions avoids reflect kind inspections.
4. Direct JSON unmarshaling via type assertions bypasses reflection for byte payloads.

## Subtask Mapping
- `01-task-checker-interfaces-and-methods.md`: Checker interfaces and Result/AppError implementations.
- `02-task-typecast-bulletproof-conversions.md`: Optimized `ReflectSetTo`, `ToBytes`, `ToJSON`, and unit tests.
- `03-task-result-and-apperror-casting-helpers.md`: Generic casting helpers across `Result[T]` and `*AppError`.
- `04-task-fileutil-and-appwriter-hardening.md`: Verification of file tactics and lock hygiene.
- `05-task-guidelines-sync-and-ci-verification.md`: Guidelines documentation, CI runner verification, and performance pointers.
