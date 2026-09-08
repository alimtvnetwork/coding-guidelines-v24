package bytetype

import (
	"strconv"

	"coding-guidelines/common/pkg/baseenumer"
)

var (
	allVariants = []Variant{Zero, One, Two, Three, Max}

	allValues = []string{"Zero", "One", "Two", "Three", "Max"}

	basicEnum = baseenumer.NewBasicSparseInteger(allVariants, allValues, Zero, 255)
)

func All() []Variant {
	return basicEnum.All()
}

func Values() []string {
	return basicEnum.Values()
}

func Parse(s string) (Variant, bool) {
	v, trimmed, isOk := basicEnum.ParseLookup(s)
	if isOk {
		return v, true
	}

	return parseFallback(trimmed)
}

func parseFallback(trimmed string) (Variant, bool) {
	if len(trimmed) == 0 {
		return Zero, false
	}

	parsed, err := strconv.ParseUint(trimmed, 10, 8)
	if err != nil {
		return Zero, false
	}

	return Variant(parsed), true
}

func ParseOrZero(s string) Variant {
	v, isOk := Parse(s)
	if isOk {
		return v
	}

	return Zero
}

func ParseOrInvalid(s string) Variant {
	return ParseOrZero(s)
}

func ParseOrUnknown(s string) Variant {
	return ParseOrZero(s)
}
