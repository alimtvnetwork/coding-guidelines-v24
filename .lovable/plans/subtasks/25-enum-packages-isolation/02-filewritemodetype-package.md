# Subtask 25.2: Create `pkg/enum/filewritemodetype` Package

## Objective
Create `04-code/golang/pkg/enum/filewritemodetype/` containing:
- `variant.go`: `type Variant uint8`, constants (`Direct = 1`, `Atomic = 2`, `Truncate = 3`), predicates, JSON marshaling.
- `vars.go`: `variantLabels`, `variantMap`, `All()`, `Values()`, `Parse()`.
- `variant_test.go`: table-driven tests.
- `readme.md`: package documentation.

## Acceptance Criteria
- Implements `baseenumer.BaseEnumer`, `baseenumer.ByteEnumer`, `baseenumer.NumberEnumer`, `json.Marshaler`, `json.Unmarshaler`.
- All functions <= 15 lines.
- Implicit boolean checks only.
- 100% test coverage.
