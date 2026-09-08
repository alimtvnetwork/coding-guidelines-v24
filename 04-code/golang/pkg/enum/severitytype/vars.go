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

	basicEnum = baseenumer.NewBasicInteger(variantLabels[:], Unknown)
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
	return basicEnum.Parse(s)
}

func ParseOrUnknown(s string) Variant {
	return basicEnum.ParseOrZero(s)
}

func ParseOrZero(s string) Variant {
	return basicEnum.ParseOrZero(s)
}

func ParseOrInvalid(s string) Variant {
	return basicEnum.ParseOrZero(s)
}
