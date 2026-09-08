package bytetype

import (
	"strconv"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type Result = result.Wrap[Variant]

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

func Parse(s string) Result {
	v, trimmed, isOk := basicEnum.ParseLookup(s)
	if isOk {
		return result.WrapSuccess(v)
	}

	return parseFallback(s, trimmed)
}

func parseFallback(original, trimmed string) Result {
	if len(trimmed) == 0 {
		return result.WrapFailureWithId[Variant](errtype.Validation, baseenumer.FormatEmptyParseError("bytetype"))
	}

	parsed, err := strconv.ParseUint(trimmed, 10, 8)
	if err != nil {
		return result.WrapFailureWithId[Variant](
			errtype.NotFound,
			baseenumer.FormatParseError("bytetype", original, Values()),
		)
	}

	return result.WrapSuccess(Variant(parsed))
}
