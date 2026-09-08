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

	// MinMaxer defines boundary value retrieval for enums (forwarded from baseenumer).
	MinMaxer[V any] = baseenumer.MinMaxer[V]

	// MinMax is an alias for MinMaxer (forwarded from baseenumer).
	MinMax[V any] = baseenumer.MinMax[V]

	// BoundedEnumer defines boundary checks for enum instances (forwarded from baseenumer).
	BoundedEnumer[V any] = baseenumer.BoundedEnumer[V]

	// BoundedEnum is an alias for BoundedEnumer (forwarded from baseenumer).
	BoundedEnum[V any] = baseenumer.BoundedEnum[V]

	// Bounder defines range checking for enum instances (forwarded from baseenumer).
	Bounder[V any] = baseenumer.Bounder[V]
)

// ToEnum finds an enum by name in any slice of BaseEnumer (delegates to baseenumer.ToEnum).
func ToEnum[T BaseEnumer](val string, all []T) (T, bool) {
	return baseenumer.ToEnum(val, all)
}
