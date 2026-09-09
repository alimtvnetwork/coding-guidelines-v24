# Subtask 38.2: Result Numbers, Reflection & Type Conversions

## Description
Implement number conversions (`Int`, `Int64`, `Float64`, `Double`, `Byte` with fallback defaults), reflection casting (`ReflectTo`), type inspection (`Type`, `TypeName`, `Kind`, `Length`), and type predicates (`IsNumber`, `IsStringType`, `IsSliceOrArray`, `IsMap`, `IsStruct`, `IsPointer`, `IsPrimitive`) directly on `Result[T]`, achieving parity with `03-aukgo/core/coredata/coredynamic`.

## Target Files
- `04-code/golang/pkg/appfault/result_dynamic_numbers.go`
- `04-code/golang/pkg/appfault/result_dynamic_types.go`
- `04-code/golang/pkg/appfault/result_dynamic_reflect.go`

## Acceptance Criteria
1. `Int() (int, bool)` converts integer, float, string, or boolean payloads to `int`.
2. `IntDefault(defaultVal int) int` returns the converted `int` or `defaultVal` on failure.
3. `Int64() (int64, bool)` and `Int64Default(defaultVal int64) int64` perform 64-bit integer conversion.
4. `Float64() (float64, bool)` and `Float64Default(defaultVal float64) float64` perform 64-bit float conversion.
5. `Double() (float64, bool)` and `DoubleDefault(defaultVal float64) float64` provide aliases for float conversion.
6. `Byte() (byte, bool)` and `ByteDefault(defaultVal byte) byte` provide byte conversion.
7. `Type() reflect.Type`, `TypeName() string`, and `Kind() reflect.Kind` provide accurate reflection metadata.
8. `Length() int` returns the length of slice, array, map, string, or channel, safely handling pointers and returning 0 on nil.
9. Type predicates `IsNumber()`, `IsStringType()`, `IsSliceOrArray()`, `IsMap()`, `IsStruct()`, `IsPointer()`, and `IsPrimitive()` accurately classify the underlying payload.
10. `ReflectTo(targetPointer any) *AppError` transfers the payload into `targetPointer` using `typecast.ReflectSetTo`, returning an `*AppError` if invalid or incompatible.
11. `ToMapResult() ResultMap[string, any]` and `Map() map[string]any` provide map conversions.
12. Every function is $\le 15$ lines.
13. All booleans use `is`/`has` prefixes and implicit evaluations.

## Status
- [x] Complete
