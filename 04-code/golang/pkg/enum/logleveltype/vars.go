package logleveltype

import (
	"coding-guidelines/common/pkg/baseenumer"
)

var (
	variantLabels = [...]string{
		Invalid: "Unknown",
		Debug:   "Debug",
		Info:    "Info",
		Warn:    "Warn",
		Error:   "Error",
		Fatal:   "Fatal",
	}

	basicEnum = baseenumer.NewBasicInteger(variantLabels[:], Invalid)
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

func ParseOrInvalid(s string) Variant {
	return basicEnum.ParseOrZero(s)
}

func ParseOrUnknown(s string) Variant {
	return basicEnum.ParseOrZero(s)
}

func ParseOrZero(s string) Variant {
	return basicEnum.ParseOrZero(s)
}
