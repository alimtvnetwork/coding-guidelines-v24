package processstatetype

import (
	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type Result = result.Wrap[Variant]

var (
	variantLabels = [...]string{
		Invalid:   "Unknown",
		Pending:   "Pending",
		Running:   "Running",
		Completed: "Completed",
		Failed:    "Failed",
		Cancelled: "Cancelled",
	}

	basicEnum  = baseenumer.NewBasicInteger(variantLabels[:], Invalid)
	variantMap = basicEnum.Map()
)

func All() []Variant {
	return basicEnum.All()
}

func Values() []string {
	return basicEnum.Values()
}

func Parse(s string) Result {
	v, trimmed, ok := baseenumer.ParseLookup(s, variantMap)
	if !ok {
		if len(trimmed) == 0 {
			return result.WrapFailureWithId[Variant](errtype.Validation, baseenumer.FormatEmptyParseError("processstatetype"))
		}

		return result.WrapFailureWithId[Variant](
			errtype.NotFound,
			baseenumer.FormatParseError("processstatetype", s, Values()),
		)
	}

	return result.WrapSuccess(v)
}
