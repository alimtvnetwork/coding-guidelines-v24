package severitytype

import (
	"coding-guidelines/common/pkg/baseenumer"
)

var (
	variantLabels = [...]string{
		Unknown:  "Unknown",
		Info:     "Info",
		Warn:     "Warn",
		Error:    "Error",
		Critical: "Critical",
		Fatal:    "Fatal",
	}

	variantMap = baseenumer.CompileMap(variantLabels[:], Unknown)
)

func All() []Variant {
	return baseenumer.SliceVariants[Variant](variantLabels[:])
}

func Values() []string {
	return baseenumer.SliceValues(variantLabels[:])
}

func Parse(s string) (Variant, bool) {
	v, _, ok := baseenumer.ParseLookup(s, variantMap)

	return v, ok
}

func ParseOrUnknown(s string) Variant {
	v, ok := Parse(s)
	if ok {
		return v
	}

	return Unknown
}
