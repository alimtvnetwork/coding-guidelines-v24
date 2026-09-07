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

	// StringEnumer defines the interface for string-backed enumerations (forwarded from baseenumer).
	StringEnumer = baseenumer.StringEnumer

	// StringEnum is an alias for StringEnumer.
	StringEnum = baseenumer.StringEnum

	// NumberEnumer defines the interface for numeric-backed enumerations (forwarded from baseenumer).
	NumberEnumer = baseenumer.NumberEnumer

	// NumberEnum is an alias for NumberEnumer.
	NumberEnum = baseenumer.NumberEnum
)

// ToEnum finds an enum by name in any slice of BaseEnumer (delegates to baseenumer.ToEnum).
func ToEnum[T BaseEnumer](val string, all []T) (T, bool) {
	return baseenumer.ToEnum(val, all)
}
