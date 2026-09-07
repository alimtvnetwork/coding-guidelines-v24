# Subtask 25.1: Create `pkg/enum/fileoptype` Package

## Objective
Create `04-code/golang/pkg/enum/fileoptype/` containing:
- `variant.go`: `type Variant byte`, constants, predicates, `OpenMode() openfiletype.Variant`, JSON marshaling.
- `vars.go`: `variantLabels`, `variantMap`, `All()`, `Values()`, `Parse()`.
- `variant_test.go`: table-driven tests.
- `readme.md`: package documentation.

## Acceptance Criteria
- Implements `baseenumer.BaseEnumer`, `baseenumer.ByteEnumer`, `baseenumer.NumberEnumer`, `json.Marshaler`, `json.Unmarshaler`.
- All functions <= 15 lines.
- Implicit boolean checks only.
- 100% test coverage.
