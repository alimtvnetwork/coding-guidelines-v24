# Subtask 31.1: Reflection-Based Type Name Resolution & Signature Simplification

## Context
Eliminate manual `typeName string` arguments from `UnmarshalIntegerJSON` and `UnmarshalStringJSON` in `04-code/golang/pkg/baseenumer/json_marshaling.go` by deriving the type name via reflection.

## Target Files
1. `04-code/golang/pkg/baseenumer/resolve_type.go` [NEW]:
   - `func ResolveTypeName[V any](target *V) string`
2. `04-code/golang/pkg/baseenumer/json_marshaling.go` [MODIFY]:
   - `func UnmarshalIntegerJSON[V IntNumber](data []byte, target *V, variantMap map[string]V, maxValid int, zero V) error`
   - `func UnmarshalStringJSON[V ~string](data []byte, target *V, variantMap map[string]V, zero V) error`
   - `func UnmarshalIntegerJSONWithName[V IntNumber](data []byte, target *V, typeName string, variantMap map[string]V, maxValid int, zero V) error`
   - `func UnmarshalStringJSONWithName[V ~string](data []byte, target *V, typeName string, variantMap map[string]V, zero V) error`
3. `04-code/golang/pkg/baseenumer/json_marshaling_test.go` [MODIFY]:
   - Update and add test cases verifying reflection type name resolution and error message formatting.

## Verification
- `go test -v -C 04-code/golang -count=1 ./pkg/baseenumer/...`
