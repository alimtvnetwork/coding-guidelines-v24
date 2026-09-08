# Subtask 01: Go Enum Migrations and BaseEnumer Test Coverage

**Plan:** [32-codebase-and-prompt-improvements.md](.lovable/plans/completed/32-codebase-and-prompt-improvements.md)  
**Status:** Completed  
**Disjoint File Scope:**
- `04-code/golang/pkg/baseenumer/basic_enum.go`
- `04-code/golang/pkg/baseenumer/basic_enum_test.go`
- `04-code/golang/pkg/baseenumer/resolve_type_test.go`
- `04-code/golang/pkg/errtype/processstatetype/variant.go`
- `04-code/golang/pkg/errtype/processstatetype/variant_test.go`
- `04-code/golang/pkg/errtype/logleveltype/variant.go`
- `04-code/golang/pkg/errtype/logleveltype/variant_test.go`
- `04-code/golang/pkg/enum/bytetype/vars.go`
- `04-code/golang/pkg/enum/bytetype/variant.go`
- `04-code/golang/pkg/enum/bytetype/variant_test.go`

---

## Acceptance Criteria

- [x] 1. Add `NewBasicSparseInteger[V IntNumber](variants []V, names []string, zero V, maxValid int) *BasicIntegerEnum[V]` to `04-code/golang/pkg/baseenumer/basic_enum.go` to handle sparse non-contiguous enums without empty slice entries.
- [x] 2. Migrate `04-code/golang/pkg/errtype/processstatetype/variant.go` to `baseenumer.NewBasicString`, removing manual registry and map compilation.
- [x] 3. Migrate `04-code/golang/pkg/errtype/logleveltype/variant.go` to `baseenumer.NewBasicInteger`, standardizing `All()`, `Values()`, and `UnmarshalJSON()`.
- [x] 4. Migrate `04-code/golang/pkg/enum/bytetype/vars.go` to `baseenumer.NewBasicSparseInteger`.
- [x] 5. Create `04-code/golang/pkg/baseenumer/resolve_type_test.go` and achieve 100% statement and branch coverage on `ResolveTypeName` (specifically hitting line 28 `return t.Name()` for built-in primitive types, nil pointers, and package structs).
- [x] 6. Expand `04-code/golang/pkg/baseenumer/basic_enum_test.go` to test `Map()`, `TypeName()`, `Zero()`, `MaxValid()`, and `ParseLookup()` for integer, string, and sparse enums.
- [x] 7. Ensure all modified/created Go functions are $\le 15$ lines.
- [x] 8. Verify all Go package tests pass: `go test ./pkg/baseenumer/... ./pkg/errtype/... ./pkg/enum/...`

---

## Verification Commands

```powershell
cd 04-code/golang
go test ./pkg/baseenumer/... -v -cover
go test ./pkg/errtype/... -v
go test ./pkg/enum/bytetype/... -v
```
