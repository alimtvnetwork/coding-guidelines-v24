# Subtask 02: Openfiletype and System Enums Tests

> **Plan:** `15-comprehensive-tests-for-enums-and-baseenumer`  
> **Status:** Completed  
> **Target File:** `04-code/golang/pkg/enum/openfiletype/variant_test.go`, `04-code/golang/pkg/enum/bytetype/variant_test.go`, `04-code/golang/pkg/enum/filepermtype/variant_test.go`  

---

## Intent
1. `openfiletype/variant_test.go`:
   - Test `IsReadWrite()` and `IsAppend()` predicates.
   - Test `Label()`.
   - Test `ParseOrUnknown()`.
   - Test out of bounds variant fallback: `Variant(99).Name()` -> `"OpenFile(99)"` and `Variant(99).Flags()` -> `os.O_RDONLY`.
   - Achieve 100% statement coverage.
2. `bytetype/variant_test.go`:
   - Test custom variant out of bounds `Variant(99).Name()` -> `"Byte(99)"`.
   - Test `Variant(0).Name()` -> `"Zero"`.
3. `filepermtype/variant_test.go`:
   - Test `WithExecutable` on existing executable mode to exercise already-set branches.
   - Test `ParseOrInvalid` and `ParseOrUnknown`.

## Verification
Run `go test ./pkg/enum/openfiletype ./pkg/enum/bytetype ./pkg/enum/filepermtype -v -count=1`.
