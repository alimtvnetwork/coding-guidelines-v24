# Subtask 01: BaseEnumer MinMax Interfaces and Engine

**Plan:** [33-enum-min-max-methods-and-baseenumer-enhancement.md](.lovable/plans/completed/33-enum-min-max-methods-and-baseenumer-enhancement.md)  
**Status:** Complete  
**Disjoint File Scope:**
- `04-code/golang/pkg/baseenumer/min_maxer.go`
- `04-code/golang/pkg/baseenumer/basic_enum.go`
- `04-code/golang/pkg/baseenumer/basic_enum_test.go`
- `04-code/golang/pkg/baseenumer/readme.md`
- `04-code/golang/pkg/errtype/base_enumer.go`

---

## Acceptance Criteria

- [x] 1. Create `04-code/golang/pkg/baseenumer/min_maxer.go` with:
  - `MinMaxer[V any]` interface (`Min() V`, `Max() V`) and alias `MinMax[V any]`.
  - `BoundedEnumer[V any]` interface (`MinMaxer[V]`, `IsMin() bool`, `IsMax() bool`) and alias `BoundedEnum[V any]`.
  - `Bounder[V any]` interface (`BoundedEnumer[V]`, `IsInRange(min, max V) bool`).
- [x] 2. Update `04-code/golang/pkg/baseenumer/basic_enum.go`:
  - Add `min` and `max` fields to `BasicIntegerEnum[V]`.
  - In `NewBasicInteger`, set `min: zero` (always starts from zero) and `max: V(maxVal)`.
  - In `NewBasicSparseInteger`, set `min: zero` (starts from zero) and `max: V(maxValid)`.
  - Implement `Min() V`, `Max() V`, `IsMin(v V) bool`, `IsMax(v V) bool`, `IsInRange(v, min, max V) bool` on `BasicIntegerEnum[V]`.
  - Add `min` and `max` fields to `BasicStringEnum[V]`.
  - In `NewBasicString`, compute `min` (first valid variant) and `max` (last valid variant).
  - Implement `Min() V`, `Max() V`, `IsMin(v V) bool`, `IsMax(v V) bool`, `IsInRange(v, min, max V) bool`, and `WithMinMax(min, max V)` on `BasicStringEnum[V]`.
  - Ensure static type assertions: `var _ MinMaxer[int] = (*BasicIntegerEnum[int])(nil)` and `var _ MinMaxer[string] = (*BasicStringEnum[string])(nil)`.
- [x] 3. Update `04-code/golang/pkg/errtype/base_enumer.go` with alias forwarders for `MinMaxer`, `MinMax`, `BoundedEnumer`, and `BoundedEnum`.
- [x] 4. Add comprehensive unit tests in `04-code/golang/pkg/baseenumer/basic_enum_test.go` verifying `Min()`, `Max()`, `IsMin()`, `IsMax()`, and `IsInRange()` across contiguous integers, sparse integers, and string enums.
- [x] 5. Update `04-code/golang/pkg/baseenumer/readme.md` documenting the min/max interfaces and `basicEnum` methods.
- [x] 6. Ensure all Go functions are $\le 15$ lines and follow implicit booleans.

---

## Verification Commands

```powershell
cd 04-code/golang
go test ./pkg/baseenumer/... -v -cover
go test ./pkg/errtype/... -v
```
