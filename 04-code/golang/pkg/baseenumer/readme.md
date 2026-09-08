# `baseenumer` Package

The `baseenumer` package provides universal, zero-dependency interfaces and contracts for strongly-typed Go enumerations across the repository.

## Interfaces

### `BaseEnumer` (Alias: `BaseEnum`)
Defines the base contract for all type-safe enumerations:
- `Name() string`: Returns the canonical string identifier of the enum.
- `String() string`: Implements `fmt.Stringer`.
- `ValueString() string`: Returns the string representation of the underlying value.
- `IsValid() bool`: Reports whether the enum instance is a valid recognized value.
- `IsEnum() bool`: Reports whether the value exists in the registered enum set.

### `ByteEnumer` (Alias: `ByteEnum`)
Extends `BaseEnumer` for byte-backed enumerations:
- `Byte() byte`: Returns the byte value.
- `ValueByte() byte`: Returns the byte value.
- `Bytes() []byte`: Returns the byte slice representation.

### `StringEnumer` (Alias: `StringEnum`)
Extends `BaseEnumer` for string-backed enumerations.

### `NumberEnumer` (Alias: `NumberEnum`)
Extends `BaseEnumer` for numeric-backed enumerations:
- `Int() int`: Returns the integer value.
- `Code() uint16`: Returns the numeric unsigned 16-bit code.

### `MinMaxer[V any]` (Alias: `MinMax[V any]`)
Defines boundary value retrieval for enumerations:
- `Min() V`: Returns the minimum valid variant.
- `Max() V`: Returns the maximum valid variant.

### `BoundedEnumer[V any]` (Alias: `BoundedEnum[V any]`)
Extends `MinMaxer[V]` for boundary checks on enum instances:
- `IsMin() bool`: Reports whether the instance equals the minimum variant.
- `IsMax() bool`: Reports whether the instance equals the maximum variant.

### `Bounder[V any]`
Extends `BoundedEnumer[V]` with range checking:
- `IsInRange(min, max V) bool`: Reports whether the instance falls within `[min, max]`.

## Utilities

### `ToEnum[T BaseEnumer](val string, all []T) (T, bool)`
Case-insensitively searches a slice of `BaseEnumer` instances by `Name()` or `ValueString()`.

## Type Constraints (`types.go`)

- `IntNumber`: Type constraint matching all signed and unsigned integer types (`~int`, `~int8`, `~int16`, `~int32`, `~int64`, `~uint`, `~uint8`, `~uint16`, `~uint32`, `~uint64`, `~uintptr`).
- `IntegerVarianter`: Backward-compatible alias for `IntNumber`.

## Generic Enum Helpers (`helpers.go`)

Reusable, zero-dependency generic utilities to eliminate boilerplate across concrete enum packages:

- `CompileMap[V IntNumber](labels []string, invalid V) map[string]V`: Generates lookup maps mapping name, uppercase, lowercase, numeric string, and `unknown`/`invalid` to typed variants.
- `SliceValues(labels []string) []string`: Extracts all valid variant label strings excluding index 0 (`Invalid`/`Unknown`).
- `SliceVariants[V IntNumber](labels []string) []V`: Extracts all valid typed variant values excluding index 0.
- `FormatNameValue(name string, val any) string`: Produces canonical `"Name(Value)"` formatting.
- `IsBetween[T cmp.Ordered](val, min, max T) bool`: Closed range boundary check `min <= val && val <= max`.
- `IsNotBetween[T cmp.Ordered](val, min, max T) bool`: Closed range boundary check `val < min || val > max`.
- `ParseLookup[V any](s string, variantMap map[string]V) (val V, trimmed string, ok bool)`: Fast trimmed lowercase map lookup.
- `FormatParseError(typeName, raw string, supportedVariants []string) string`: Canonical error string generator.
- `FormatEmptyParseError(typeName string) string`: Canonical empty input error string generator.
- `ResolveTypeName[V any](target *V) string`: Derives clean type/package names using reflection, eliminating manual string arguments.

## Generic JSON Marshaling Helpers (`json_marshaling.go`)

Reusable generic JSON serialization and deserialization functions:

- `MarshalJSON(name string) ([]byte, error)`: Standard JSON string marshaler for enum names.
- `UnmarshalIntegerJSON[V IntNumber](data []byte, target *V, variantMap map[string]V, maxValid int, zero V) error`: Unified integer- and byte-backed JSON unmarshaler with auto-resolved type name.
- `UnmarshalStringJSON[V ~string](data []byte, target *V, variantMap map[string]V, zero V) error`: Unified string-backed JSON unmarshaler with auto-resolved type name.
- `UnmarshalIntegerJSONWithName` / `UnmarshalStringJSONWithName`: Explicit type name variants for custom overrides.

## Basic Enum Engine (`basic_enum.go`)

Universal generic enum manager structs that encapsulate labels, lookup maps, bounds, zero values, and reflection-derived type names:

- `BasicIntegerEnum[V IntNumber]`: Reusable manager for integer- and byte-backed enums (`NewBasicInteger(labels, zero)`, `NewBasicSparseInteger(variants, names, zero, maxValid)`). Provides 2-parameter `UnmarshalJSON(data, target)`, `All()`, `Values()`, `Parse(s)`, and boundary methods:
  - `Min() V`: Returns the zero-value lower bound.
  - `Max() V`: Returns the upper valid bound.
  - `IsMin(v V) bool`: Reports whether variant equals `Min()`.
  - `IsMax(v V) bool`: Reports whether variant equals `Max()`.
  - `IsInRange(v, min, max V) bool`: Reports whether variant is within `[min, max]`.
- `BasicStringEnum[V ~string]`: Reusable manager for string-backed enums (`NewBasicString(variants, zero)`). Provides 2-parameter `UnmarshalJSON(data, target)`, `All()`, `Values()`, `Parse(s)`, and boundary methods:
  - `Min() V`: Returns the first valid variant bound.
  - `Max() V`: Returns the last valid variant bound.
  - `IsMin(v V) bool`: Reports whether variant equals `Min()`.
  - `IsMax(v V) bool`: Reports whether variant equals `Max()`.
  - `IsInRange(v, min, max V) bool`: Reports whether variant is within `[min, max]`.
  - `WithMinMax(min, max V) *BasicStringEnum[V]`: Overrides the computed min/max bounds.

