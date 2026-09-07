# Subtask 02: Streamwriter Verifier Conformance

Target Directory: `04-code/golang/pkg/streamwriter/`
Affected Files:
- `04-code/golang/pkg/streamwriter/bytes.go`
- `04-code/golang/pkg/streamwriter/json_result.go`
- `04-code/golang/pkg/streamwriter/streamwriter_test.go`

## Instructions
1. In `bytes.go` (`Bytes[T]`):
   - Add `IsFailure() bool`: returns `!b.IsSuccess()`.
   - Add `IsInvalid() bool`: returns `b.appError != nil`.
   - Add `IsDefined() bool`: returns `b.IsSuccess()`.
   - Add `AsSimpleVerifier() appfault.SimpleVerifier`: returns `b`.
   - Add `AsSimpleVerifyChecker() appfault.SimpleVerifier`: returns `b`.
   - Add static compile assertions:
     `var _ appfault.SimpleVerifier = Bytes[any]{}`
     `var _ appfault.SimpleVerifyChecker = Bytes[any]{}`
     `var _ appfault.SimpleVerifiable = Bytes[any]{}`
     `var _ appfault.SimpleVerifyCheckable = Bytes[any]{}`
2. In `json_result.go` (`JsonResult`):
   - Add `IsFailure() bool`: returns `j.appError != nil`.
   - Add `IsInvalid() bool`: returns `j.appError != nil`.
   - Add `IsDefined() bool`: returns `j.IsSuccess()`.
   - Add `AsSimpleVerifier() appfault.SimpleVerifier`: returns `j`.
   - Add `AsSimpleVerifyChecker() appfault.SimpleVerifier`: returns `j`.
   - Add static compile assertions for `JsonResult` and `JsonPayloadResult[any]`.
3. In `streamwriter_test.go`:
   - Add test functions asserting `AsSimpleVerifier` and `AsSimpleVerifyChecker` behavior on `Bytes[T]` and `JsonResult`.
   - Ensure each function is <= 15 lines.
   - Blank line after each `}` if followed by code.
