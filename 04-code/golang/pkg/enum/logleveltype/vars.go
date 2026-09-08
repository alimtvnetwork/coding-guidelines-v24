package logleveltype

import (
	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type Result = result.Wrap[Variant]

var (
	variantLabels = [...]string{
		Invalid: "Unknown",
		Debug:   "Debug",
		Info:    "Info",
		Warn:    "Warn",
		Error:   "Error",
		Fatal:   "Fatal",
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
	if len(trimmed) == 0 {
		return result.WrapFailureWithId[Variant](errtype.Validation, baseenumer.FormatEmptyParseError("logleveltype"))
	}

	if ok {
		return result.WrapSuccess(v)
	}

	return result.WrapFailureWithId[Variant](
		errtype.NotFound,
		baseenumer.FormatParseError("logleveltype", s, Values()),
	)
}
