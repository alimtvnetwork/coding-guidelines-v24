# Subtask 25.4: Create `pkg/enum/prioritytype` Package

## Objective
Create `04-code/golang/pkg/enum/prioritytype/` containing:
- `variant.go`: `type Variant byte`, constants (`Unknown = 0`, `Low = 1`, `Normal = 2`, `High = 3`, `Critical = 4`), predicates, JSON marshaling.
- `vars.go`: `variantLabels`, `variantMap`, `All()`, `Values()`, `Parse()`, `ParseOrUnknown()`.
- `variant_test.go`: table-driven tests.
- `readme.md`: package documentation.

## Acceptance Criteria
- Leaf imports only: standard library + `baseenumer`. Strictly NO `pkg/result` or `pkg/appfault`.
- Implements `baseenumer.BaseEnumer`, `baseenumer.ByteEnumer`, `baseenumer.NumberEnumer`, `json.Marshaler`, `json.Unmarshaler`.
- All functions <= 15 lines.
- Implicit boolean checks only.
- 100% test coverage.
