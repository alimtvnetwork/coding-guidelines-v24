# Subtask 01: Baseenumer Extended Tests

> **Plan:** `15-comprehensive-tests-for-enums-and-baseenumer`  
> **Status:** Completed  
> **Target File:** `04-code/golang/pkg/baseenumer/basic_enum_test.go`, `04-code/golang/pkg/baseenumer/json_marshaling_test.go`  

---

## Intent
1. Add tests in `json_marshaling_test.go` for:
   - `parseAndAssignNumericString`: invalid numeric string ("abc"), out of range numeric string ("999"), and valid numeric strings ("1").
   - `isOutOfRange`: variants below min and variants above max.
   - `parseStringVariant`: empty variant and unknown string error reporting.
2. In `basic_enum_test.go`:
   - Test `isNonEmptyVariant` with custom sparse variants and empty values.
   - Test `WithMinMax` builder pattern on `BasicStringEnum`.
   - Test `ParseLookup` with empty and spaces.

## Verification
Run `go test ./pkg/baseenumer -v -count=1`.
