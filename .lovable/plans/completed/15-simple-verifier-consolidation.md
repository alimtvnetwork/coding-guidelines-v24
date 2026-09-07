# Spec: 15-simple-verifier-consolidation

## 1. Overview & Objective
Consolidate all verification and checker interfaces across the Go codebase into unified `SimpleVerifier` and `SimpleVerifyChecker` contracts with zero cycles. Provide `.AsSimpleVerifier()` and `.AsSimpleVerifyChecker()` conversion methods and compile-time type assertions across all result, error, and streaming envelope types (`Result[T]`, `*AppError`, `ResultSlice[T]`, `ResultMap[K, V]`, `Bytes[T]`, `JsonResult`, `JsonPayloadResult[T]`).

## 2. Task-Specific Rules & Constraints
1. **Zero-Cycle Dependency Direction:** `pkg/streamwriter` and `pkg/result` may depend on `pkg/appfault`. `pkg/appfault` MUST NEVER import `pkg/streamwriter` or `pkg/result`. All core contracts remain in `pkg/appfault`.
2. **Double Alias Suffix Parity:** Every verifier contract MUST support both the canonical name (`SimpleVerifier`, `.AsSimpleVerifier()`) and the `Checker` suffix standard (`SimpleVerifyChecker`, `.AsSimpleVerifyChecker()`).
3. **Compile-Time Static Assertions:** All conforming types MUST have explicit compile-time static type assertion checks (`var _ appfault.SimpleVerifier = ...`) to prevent future contract breakage.
4. **Function Size Bound:** No function or test may exceed 15 lines (strictly <= 15 lines, target <= 8 lines).
5. **Strict Relative Git Paths:** All file paths and markdown citations MUST be strictly relative to the repository root.

## 3. Architectural Design

### 3.1 Interface Contracts (`04-code/golang/pkg/appfault/interfaces.go`)
- `SimpleVerifier`: Embeds `IsSuccessChecker`, `IsFailureChecker`, `IsInvalidChecker`, `IsNullChecker`, `IsEmptyChecker`, `IsDefinedChecker`, `DefinableChecker`, `StatusChecker`.
- `SimpleVerifyChecker`: Type alias to `SimpleVerifier`.
- `SimpleVerifiable`: Requires `AsSimpleVerifier() SimpleVerifier`.
- `SimpleVerifyCheckable`: Requires `AsSimpleVerifyChecker() SimpleVerifier`.

### 3.2 Conforming Types
1. `Result[T]` (`04-code/golang/pkg/appfault/result_methods.go`)
2. `*AppError` (`04-code/golang/pkg/appfault/methods.go`)
3. `ResultSlice[T]` (`04-code/golang/pkg/appfault/result_slice.go`)
4. `ResultMap[K, V]` (`04-code/golang/pkg/appfault/result_map.go`)
5. `Bytes[T]` (`04-code/golang/pkg/streamwriter/bytes.go`)
6. `JsonResult` (`04-code/golang/pkg/streamwriter/json_result.go`)
7. `JsonPayloadResult[T]` (`04-code/golang/pkg/streamwriter/json_result.go`)

## 4. Subtasks Decomposition
- `01-task-appfault-verifier-parity.md`: Add `SimpleVerifiable`, `SimpleVerifyCheckable`, and `.AsSimpleVerifyChecker()` on all appfault result/error types; update compile-time assertions and unit tests.
- `02-task-streamwriter-verifier-conformance.md`: Add missing checker methods (`IsFailure`, `IsInvalid`, `IsDefined`) and `.AsSimpleVerifier()`, `.AsSimpleVerifyChecker()` to `Bytes[T]`, `JsonResult`, `JsonPayloadResult[T]`; add compile-time assertions and unit tests.
