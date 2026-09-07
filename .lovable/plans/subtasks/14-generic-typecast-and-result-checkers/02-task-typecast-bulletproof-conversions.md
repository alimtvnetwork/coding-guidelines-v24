# Subtask 02: Bulletproof Typecast Suite and Reflection Optimization

## Objective
Optimize `ReflectSetTo` for high-throughput performance with fast-paths, add column/serialization methods (`ToBytes`, `ToJSON`, `ToJSONString`), and provide exhaustive test coverage.

## Target Files
- `04-code/golang/pkg/typecast/cast.go`
- `04-code/golang/pkg/typecast/cast_test.go`

## Detailed Instructions
1. In `cast.go`:
   - Declare package-level singleton reflect types:
     ```go
     var (
         emptyBytesType        = reflect.TypeOf([]byte(nil))
         emptyBytesPointerType = reflect.TypeOf((*[]byte)(nil))
     )
     ```
   - Implement fast-path type switch in `ReflectSetTo` for common destination pointer types (`*string`, `*int`, `*int64`, `*bool`, `*float64`, `*[]byte`) when source is direct value or pointer, completely skipping `reflect.ValueOf()`/`reflect.TypeOf()`.
   - Implement fast-path for JSON unmarshaling when source is `[]byte`.
   - Implement fast-path for JSON marshaling when destination is `*[]byte`.
   - Preserve robust reflection fallback for arbitrary struct types and matching pointers.
   - Implement `ToBytes(payload any) ([]byte, error)`:
     - Handles `[]byte`, `string`, `[]string` (newline separated), `error`, and fallback to JSON.
   - Implement `ToJSON(payload any) ([]byte, error)`:
     - Formats payload as indented JSON with trailing newline.
   - Implement `ToJSONString(payload any) (string, error)`:
     - Formats payload as indented JSON string.
2. In `cast_test.go`:
   - Add comprehensive tests covering all conversion paths:
     - Primitive conversions (`int`, `string`, `bool`, `float64`)
     - Pointer to pointer conversions
     - Byte slice to struct unmarshal
     - Struct to byte slice marshal
     - `ToBytes`, `ToJSON`, and `ToJSONString` with various payload types
     - Error conditions (nil destination, non-pointer destination, incompatible types)
   - Add benchmark tests comparing fast-path vs reflection fallback (`BenchmarkReflectSetTo_Primitive`, `BenchmarkReflectSetTo_BytesToStruct`).
