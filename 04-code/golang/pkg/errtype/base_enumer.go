package errtype

import (
	"coding-guidelines/common/pkg/baseenumer"
)

type (
	// BaseEnumer defines the interface for all type-safe enumerations (forwarded from baseenumer).
	BaseEnumer = baseenumer.BaseEnumer

	// BaseEnum is an alias for BaseEnumer.
	BaseEnum = baseenumer.BaseEnum

	// ByteEnumer defines the interface for byte-backed enumerations (forwarded from baseenumer).
	ByteEnumer = baseenumer.ByteEnumer

	// ByteEnum is an alias for ByteEnumer.
	ByteEnum = baseenumer.ByteEnum

	// Utf8Enumer defines the interface for UTF-8 byte-backed enumerations (forwarded from baseenumer).
	Utf8Enumer = baseenumer.Utf8Enumer

	// Utf8Enum is an alias for Utf8Enumer.
	Utf8Enum = baseenumer.Utf8Enum

	// UTF8Enumer is an alias for backwards compatibility.
	UTF8Enumer = baseenumer.UTF8Enumer

	// UTF8Enum is an alias for backwards compatibility.
	UTF8Enum = baseenumer.UTF8Enum

	// Utf16Enumer defines the interface for UTF-16 code-unit-backed enumerations (forwarded from baseenumer).
	Utf16Enumer = baseenumer.Utf16Enumer

	// Utf16Enum is an alias for Utf16Enumer.
	Utf16Enum = baseenumer.Utf16Enum

	// UTF16Enumer is an alias for backwards compatibility.
	UTF16Enumer = baseenumer.UTF16Enumer

	// UTF16Enum is an alias for backwards compatibility.
	UTF16Enum = baseenumer.UTF16Enum

	// Utf32Enumer defines the interface for UTF-32 / rune-backed enumerations (forwarded from baseenumer).
	Utf32Enumer = baseenumer.Utf32Enumer

	// Utf32Enum is an alias for Utf32Enumer.
	Utf32Enum = baseenumer.Utf32Enum

	// UTF32Enumer is an alias for backwards compatibility.
	UTF32Enumer = baseenumer.UTF32Enumer

	// UTF32Enum is an alias for backwards compatibility.
	UTF32Enum = baseenumer.UTF32Enum

	// RuneEnumer defines the interface for rune-backed enumerations (forwarded from baseenumer).
	RuneEnumer = baseenumer.RuneEnumer

	// RuneEnum is an alias for RuneEnumer.
	RuneEnum = baseenumer.RuneEnum

	// StringEnumer defines the interface for string-backed enumerations (forwarded from baseenumer).
	StringEnumer = baseenumer.StringEnumer

	// StringEnum is an alias for StringEnumer.
	StringEnum = baseenumer.StringEnum

	// NumberEnumer defines the interface for numeric-backed enumerations (forwarded from baseenumer).
	NumberEnumer = baseenumer.NumberEnumer

	// NumberEnum is an alias for NumberEnumer.
	NumberEnum = baseenumer.NumberEnum

	// IntEnumer defines the interface for integer-backed enumerations (forwarded from baseenumer).
	IntEnumer = baseenumer.IntEnumer

	// IntEnum is an alias for IntEnumer.
	IntEnum = baseenumer.IntEnum
)

// ToEnum finds an enum by name in any slice of BaseEnumer (delegates to baseenumer.ToEnum).
func ToEnum[T BaseEnumer](val string, all []T) (T, bool) {
	return baseenumer.ToEnum(val, all)
}
