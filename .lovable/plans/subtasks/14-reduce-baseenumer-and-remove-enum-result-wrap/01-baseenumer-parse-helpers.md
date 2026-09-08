# Subtask 01: Baseenumer Parse Helpers

> **Plan:** `14-reduce-baseenumer-and-remove-enum-result-wrap`  
> **Status:** Completed  
> **Target File:** `04-code/golang/pkg/baseenumer/basic_enum.go`, `04-code/golang/pkg/baseenumer/basic_enum_test.go`  

---

## Intent
Add lightweight, cycle-free parse methods to `BasicIntegerEnum[V]` and `BasicStringEnum[V]` in `04-code/golang/pkg/baseenumer/basic_enum.go`:
1. `Parse(s string) (V, bool)`: Performs trimmed lowercase lookup against internal `variantMap` and returns `(val, isOk)`.
2. `ParseOrZero(s string) V`: Returns the variant if found, or `b.zero` fallback.
3. `ParseErr(s string) (V, error)`: Preserves error-returning parse behavior for callers that need formatted errors.

## Verification
Run `go test ./pkg/baseenumer -v -count=1` to ensure all tests pass.
