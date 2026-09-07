# Subtask 01: Checker Interfaces and Implementations

## Objective
Define the complete suite of boolean checker interfaces in `04-code/golang/pkg/appfault/interfaces.go` and implement all corresponding methods on `Result[T]` / `Wrap[T]` and `*AppError`.

## Target Files
- `04-code/golang/pkg/appfault/interfaces.go`
- `04-code/golang/pkg/appfault/result_methods.go`
- `04-code/golang/pkg/appfault/apperror.go`
- `04-code/golang/pkg/appfault/methods.go`

## Detailed Instructions
1. In `interfaces.go`, declare:
   - `IsSuccessChecker interface { IsSuccess() bool }`
   - `IsFailureChecker interface { IsFailure() bool }`
   - `IsInvalidChecker interface { IsInvalid() bool }`
   - `IsNullChecker interface { IsNull() bool }`
   - `IsEmptyChecker interface { IsEmpty() bool }`
   - `IsDefinedChecker interface { IsDefined() bool }`
   - `DefinableChecker interface { IsDefinedChecker; IsEmptyChecker }`
   - `StatusChecker interface { IsSuccessChecker; IsFailureChecker }`
2. In `result_methods.go`, ensure `Result[T]` implements:
   - `IsDefined() bool`: returns `r.IsSuccess()`
   - `IsSuccess() bool`, `IsFailure() bool`, `IsInvalid() bool`, `IsNull() bool`, `IsEmpty() bool`
3. In `apperror.go` and `methods.go`, ensure `*AppError` implements:
   - `IsFailure() bool`: returns `e.IsFailed()`
   - `IsSuccess() bool`, `IsInvalid() bool`, `IsNull() bool`, `IsEmpty() bool`, `IsDefined() bool`
4. Add interface verification assertions in test files:
   - `var _ IsSuccessChecker = Result[string]{}`
   - `var _ IsFailureChecker = Result[string]{}`
   - `var _ IsDefinedChecker = Result[string]{}`
   - `var _ IsNullChecker = Result[string]{}`
   - `var _ IsSuccessChecker = (*AppError)(nil)`
   - `var _ IsFailureChecker = (*AppError)(nil)`
   - `var _ IsDefinedChecker = (*AppError)(nil)`
