package baseenumer

import (
	"strconv"
	"strings"
)

type (
	// BaseEnumer defines the interface for all type-safe enumerations.
	BaseEnumer interface {
		Name() string
		String() string
		ValueString() string
		IsValid() bool
		IsEnum() bool
	}

	// BaseEnum is an alias for BaseEnumer.
	BaseEnum = BaseEnumer
)

// ToEnum finds an enum by name or value in any slice of BaseEnumer.
func ToEnum[T BaseEnumer](val string, all []T) (T, bool) {
	cleaned := strings.TrimSpace(val)
	for _, item := range all {
		if isEnumMatch(item, cleaned) {
			return item, true
		}
	}

	var zero T

	return zero, false
}

func isEnumMatch[T BaseEnumer](item T, cleaned string) bool {
	if strings.EqualFold(item.Name(), cleaned) || strings.EqualFold(item.ValueString(), cleaned) {
		return true
	}

	if ne, isNum := any(item).(NumberEnumer); isNum && strconv.Itoa(ne.Int()) == cleaned {
		return true
	}

	return false
}
