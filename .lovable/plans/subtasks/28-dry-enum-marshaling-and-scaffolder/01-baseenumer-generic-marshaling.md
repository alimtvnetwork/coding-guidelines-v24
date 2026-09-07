# Subtask 28.1: Generic JSON Marshaling in `pkg/baseenumer`

## Context
Eliminate boilerplate JSON marshaling and unmarshaling logic across all enums in the repository by providing generic, type-safe helpers in `04-code/golang/pkg/baseenumer`.

## Requirements
1. Implement `04-code/golang/pkg/baseenumer/types.go`:
   - Define type constraint `IntNumber` (~int, ~uint, etc.) in a centralized types file to eliminate spell check warnings and standardize naming.
   - Provide `IntegerVarianter = IntNumber` backward-compatible alias.
2. Implement `04-code/golang/pkg/baseenumer/json_marshaling.go`:
   - `MarshalJSON(name string) ([]byte, error)`
   - `UnmarshalIntegerJSON[V IntNumber](data []byte, target *V, typeName string, variantMap map[string]V, maxValid int, zero V) error`
   - `UnmarshalStringJSON[V ~string](data []byte, target *V, typeName string, variantMap map[string]V, zero V) error`
   - Helper functions decomposed so every function is <= 15 lines.
   - Implicit booleans only, no mixed polarity.
2. Implement `04-code/golang/pkg/baseenumer/json_marshaling_test.go`:
   - Test integer and byte enum unmarshaling (by name, case-insensitive, number string, raw number, null, empty).
   - Test out-of-bounds number rejection.
   - Test string enum unmarshaling (by string value, case-insensitive, null, empty).
   - Test unknown variant rejection.
   - 100% test pass.

## Verification
- `go test -v -C 04-code/golang -count=1 ./pkg/baseenumer/...`
