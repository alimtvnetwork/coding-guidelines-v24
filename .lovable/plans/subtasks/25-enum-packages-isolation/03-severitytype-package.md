# Subtask 25.3: Create `pkg/enum/severitytype` Package

## Objective
Create `04-code/golang/pkg/enum/severitytype/` containing:
- `variant.go`: `type Variant byte`, constants (`Unknown = 0`, `Info = 1`, `Warn = 2`, `Error = 3`, `Critical = 4`, `Fatal = 5`), predicates, JSON marshaling.
- `vars.go`: `variantLabels`, `variantMap`, `All()`, `Values()`, `Parse()`, `ParseOrUnknown()`.
- `variant_test.go`: table-driven tests.
- `readme.md`: package documentation.

## Acceptance Criteria
- Leaf imports only: standard library + `baseenumer`. Strictly NO `pkg/result` or `pkg/appfault`.
- Implements `baseenumer.BaseEnumer`, `baseenumer.ByteEnumer`, `baseenumer.NumberEnumer`, `json.Marshaler`, `json.Unmarshaler`.
- All functions <= 15 lines.
- Implicit boolean checks only.
- 100% test coverage.
