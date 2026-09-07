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

### `Utf8Enumer` (Alias: `Utf8Enum`, `UTF8Enumer`, `UTF8Enum`)
Extends `ByteEnumer` for UTF-8 byte-backed enumerations.

### `Utf16Enumer` (Alias: `Utf16Enum`, `UTF16Enumer`, `UTF16Enum`)
Extends `BaseEnumer` for UTF-16 code-unit-backed enumerations:
- `Utf16() uint16`: Returns the uint16 code unit.
- `ValueUtf16() uint16`: Returns the uint16 code unit value.
- `Code() uint16`: Returns the uint16 code value.

### `Utf32Enumer` (Alias: `Utf32Enum`, `UTF32Enumer`, `UTF32Enum`)
Extends `BaseEnumer` for UTF-32 / rune-backed enumerations:
- `Rune() rune`: Returns the rune character.
- `ValueRune() rune`: Returns the rune value.
- `Int32() int32`: Returns the int32 code point.

### `RuneEnumer` (Alias: `RuneEnum`)
Extends `Utf32Enumer` for rune-backed enumerations.

### `StringEnumer` (Alias: `StringEnum`)
Extends `BaseEnumer` for string-backed enumerations.

### `NumberEnumer` (Alias: `NumberEnum`)
Extends `BaseEnumer` for numeric-backed enumerations:
- `Int() int`: Returns the integer value.
- `Code() uint16`: Returns the numeric unsigned 16-bit code.

### `IntEnumer` (Alias: `IntEnum`)
Extends `NumberEnumer` for integer-backed enumerations.

## Utilities

### `ToEnum[T BaseEnumer](val string, all []T) (T, bool)`
Case-insensitively searches a slice of `BaseEnumer` instances by `Name()` or `ValueString()`.
