# Subtask 03: Generic Casting & Serialization in Result and AppError

## Objective
Provide generic casting helpers and serialization methods in `pkg/appfault` and `pkg/result` so that `Result[T]`, `Wrap[T]`, and `*AppError` can cast payloads and convert objects cleanly.

## Target Files
- `04-code/golang/pkg/appfault/cast.go` (new helper file)
- `04-code/golang/pkg/appfault/cast_test.go` (new test file)
- `04-code/golang/pkg/result/result.go`

## Detailed Instructions
1. In `04-code/golang/pkg/appfault/cast.go`:
   - Implement `CastTo[T any](source any) (T, *AppError)`: uses `typecast.CastTo[T]` and wraps any failure in `*AppError` with `errtype.TypeMismatch`.
   - Implement `ReflectTo[T any](source any) (T, *AppError)`: uses `typecast.ReflectTo[T]` and wraps any failure in `*AppError`.
   - Implement `CastResult[T any, U any](r Result[T]) Result[U]`: converts `Result[T]` to `Result[U]` via `ReflectTo` on data if successful, or propagates existing fault.
   - Implement `CastContextPayload[T any](e *AppError, key string) (T, *AppError)`: retrieves key from `e.Context()` and casts to `T`.
   - Implement `ResultToBytes[T any](r Result[T]) Result[[]byte]`: serializes result data or fault using `typecast.ToBytes`.
   - Implement `ResultToJSON[T any](r Result[T]) Result[[]byte]`: serializes result data or fault using `typecast.ToJSON`.
2. In `04-code/golang/pkg/result/result.go`:
   - Re-export `CastTo`, `ReflectTo`, `CastResult`, `ResultToBytes`, and `ResultToJSON` for consumers of package `result`.
3. In `04-code/golang/pkg/appfault/cast_test.go`:
   - Unit tests for casting successful and failed results, context payloads, and byte/JSON conversions.
