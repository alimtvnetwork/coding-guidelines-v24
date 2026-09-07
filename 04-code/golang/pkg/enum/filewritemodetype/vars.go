package filewritemodetype

import (
	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type Result = result.Wrap[Variant]

var (
	variantLabels = [...]string{
		Invalid:  "Invalid",
		Direct:   "Direct",
		Atomic:   "Atomic",
		Truncate: "Truncate",
	}

	variantMap = baseenumer.CompileMap(variantLabels[:], Invalid)
)

func All() []Variant {
	return baseenumer.SliceVariants[Variant](variantLabels[:])
}

func Values() []string {
	return baseenumer.SliceValues(variantLabels[:])
}

func Parse(s string) Result {
	v, trimmed, ok := baseenumer.ParseLookup(s, variantMap)
	if len(trimmed) == 0 {
		return result.WrapFailureWithId[Variant](errtype.Validation, baseenumer.FormatEmptyParseError("filewritemodetype"))
	}

	if ok {
		return result.WrapSuccess(v)
	}

	return result.WrapFailureWithId[Variant](
		errtype.NotFound,
		baseenumer.FormatParseError("filewritemodetype", s, Values()),
	)
}
