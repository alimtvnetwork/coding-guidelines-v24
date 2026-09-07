package baseenumer

import "strings"

type (
	// BaseEnumer defines the interface for all type-safe enumerations.
	BaseEnumer interface {
		Name() string
		String() string
		ValueString() string
		IsValid() bool
		IsEnum() bool
	}

	// NumberEnumer defines the interface for numeric-backed enumerations.
	NumberEnumer interface {
		BaseEnumer
		Int() int
		Code() uint16
	}

	// BaseEnum is an alias for BaseEnumer.
	BaseEnum = BaseEnumer

	// NumberEnum is an alias for NumberEnumer.
	NumberEnum = NumberEnumer
)

// ToEnum finds an enum by name in any slice of BaseEnumer.
func ToEnum[T BaseEnumer](val string, all []T) (T, bool) {
	cleaned := strings.TrimSpace(val)
	for _, item := range all {
		if strings.EqualFold(item.Name(), cleaned) || strings.EqualFold(item.ValueString(), cleaned) {
			return item, true
		}
	}

	var zero T

	return zero, false
}
