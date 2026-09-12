package processstatetype

import (
	"coding-guidelines/common/pkg/baseenumer"
)

var (
	variantLabels = [...]string{
		Invalid:   "Unknown",
		Pending:   "Pending",
		Running:   "Running",
		Completed: "Completed",
		Failed:    "Failed",
		Canceled:  "Canceled",
	}

	basicEnum = baseenumer.NewBasicInteger(variantLabels[:], Invalid)
)

func init() {
	basicEnum.RegisterAlias("Cancelled", Canceled)
}

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
