# Subtask 39.1: Result Unwrapping Interface & Container Methods

## Description
Define non-generic `ResultInspector` (`ValueAny() any`, `IsFailed() bool`, `AppError() *AppError`) and aliases `ResultUnwrapper` and `ResultCarrier` in `04-code/golang/pkg/appfault/interfaces.go`. Re-export them in `04-code/golang/pkg/result/result.go`. Implement `ValueAny() any` on `Result[T]`, `ResultSlice[T]`, and `ResultMap[K, V]`. Encapsulate `appError *AppError` in `ResultSlice[T]` and `ResultMap[K, V]` to provide `.AppError() *AppError` and `.Fault() *AppError` methods with JSON/YAML serialization DTOs.

## Target Files
- `04-code/golang/pkg/appfault/interfaces.go`
- `04-code/golang/pkg/appfault/result.go`
- `04-code/golang/pkg/appfault/result_slice.go`
- `04-code/golang/pkg/appfault/result_map.go`
- `04-code/golang/pkg/result/result.go`
- `04-code/golang/pkg/appfault/interfaces_test.go`

## Acceptance Criteria
- [x] `ResultInspector` interface defined with `ValueAny() any`, `IsFailed() bool`, and `AppError() *AppError`.
- [x] `ResultUnwrapper` and `ResultCarrier` aliases defined and exported.
- [x] `Result[T]`, `ResultSlice[T]`, and `ResultMap[K, V]` satisfy `ResultInspector`.
- [x] `ResultSlice[T]` and `ResultMap[K, V]` provide `.AppError() *AppError`, `.Fault() *AppError`, and `.ValueAny() any`.
- [x] Serialization DTOs ensure 100% JSON/YAML roundtrip fidelity.
- [x] Compile-time interface checks added in `interfaces_test.go`.
- [x] All functions <= 15 lines and implicit booleans only.
