# Subtask 01: Appfault Verifier Parity

Target Directory: `04-code/golang/pkg/appfault/`
Affected Files:
- `04-code/golang/pkg/appfault/interfaces.go`
- `04-code/golang/pkg/appfault/methods.go`
- `04-code/golang/pkg/appfault/result_methods.go`
- `04-code/golang/pkg/appfault/result_slice.go`
- `04-code/golang/pkg/appfault/result_map.go`
- `04-code/golang/pkg/result/result.go`
- `04-code/golang/pkg/appfault/interfaces_test.go`

## Instructions
1. In `interfaces.go`:
   - Define `SimpleVerifiable interface { AsSimpleVerifier() SimpleVerifier }`.
   - Define `SimpleVerifyCheckable interface { AsSimpleVerifyChecker() SimpleVerifier }`.
2. In `methods.go` (`*AppError`):
   - Add `func (e *AppError) AsSimpleVerifyChecker() SimpleVerifier { return e }`.
3. In `result_methods.go` (`Result[T]`):
   - Add `func (r Result[T]) AsSimpleVerifyChecker() SimpleVerifier { return r }`.
4. In `result_slice.go` (`ResultSlice[T]`):
   - Add `func (rs ResultSlice[T]) AsSimpleVerifyChecker() SimpleVerifier { return rs }`.
5. In `result_map.go` (`ResultMap[K, V]`):
   - Add `func (rm ResultMap[K, V]) AsSimpleVerifyChecker() SimpleVerifier { return rm }`.
6. In `04-code/golang/pkg/result/result.go`:
   - Re-export `SimpleVerifiable = appfault.SimpleVerifiable` and `SimpleVerifyCheckable = appfault.SimpleVerifyCheckable`.
7. In `interfaces_test.go`:
   - Add static compile assertions for `SimpleVerifiable` and `SimpleVerifyCheckable`.
   - Add unit tests for `AsSimpleVerifyChecker()`.
   - Ensure each function is <= 15 lines.
   - Blank line after each `}` if followed by code.
