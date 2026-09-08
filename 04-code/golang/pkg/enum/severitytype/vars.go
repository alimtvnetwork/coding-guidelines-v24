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

	basicEnum  = baseenumer.NewBasicInteger(variantLabels[:], Unknown)
	variantMap = basicEnum.Map()
)

func All() []Variant {
	return basicEnum.All()
}

func Values() []string {
	return basicEnum.Values()
}

func Min() Variant {
	return basicEnum.Min()
}

func Max() Variant {
	return basicEnum.Max()
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
