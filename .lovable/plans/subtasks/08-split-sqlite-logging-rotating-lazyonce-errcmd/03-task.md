# Subtask 03: Generic Lazy Once Engine with 0, 1, and 2 Parameter Variants

## Parent Plan
`.lovable/plans/completed/08-split-sqlite-logging-rotating-lazyonce-errcmd.md`

## Target Files
- `04-code/golang/pkg/lazyonce/lazyonce.go`
- `04-code/golang/pkg/lazyonce/lazyonce1.go`
- `04-code/golang/pkg/lazyonce/lazyonce2.go`

## Instructions
1. Implement `LazyOnce[T]` for 0-parameter initializer:
   - Initializer: `func() (T, *appfault.AppError)`
   - `Value() (T, *appfault.AppError)`
   - `Result() result.Result[T]`
   - `IsEvaluated() bool`
2. Implement `LazyOnce1[TInput, TOutput]` for 1-parameter initializer:
   - Initializer: `func(TInput) (TOutput, *appfault.AppError)`
   - `Value(input TInput) (TOutput, *appfault.AppError)`
   - `Result(input TInput) result.Result[TOutput]`
   - `IsEvaluated() bool`
3. Implement `LazyOnce2[T1, T2, TOutput]` for 2-parameter initializer:
   - Initializer: `func(T1, T2) (TOutput, *appfault.AppError)`
   - `Value(arg1 T1, arg2 T2) (TOutput, *appfault.AppError)`
   - `Result(arg1 T1, arg2 T2) result.Result[TOutput]`
   - `IsEvaluated() bool`
4. Guarantee thread safety using double-checked locking or `sync.Once`.
5. Integrate with `result.Result[T]` from `pkg/result` and `*appfault.AppError`.
6. Enforce <= 15 lines per function and LF line endings.
