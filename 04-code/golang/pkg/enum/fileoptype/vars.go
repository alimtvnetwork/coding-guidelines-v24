package fileoptype

import (
	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

var (
	variantLabels = [...]string{
		Invalid:        "Invalid",
		ReadOnly:       "ReadOnly",
		WriteOnly:      "WriteOnly",
		ReadWrite:      "ReadWrite",
		Append:         "Append",
		Create:         "Create",
		CreateAppend:   "CreateAppend",
		CreateTruncate: "CreateTruncate",
		Delete:         "Delete",
	}

	variantMap = baseenumer.CompileMap(variantLabels[:], Invalid)
)

func All() []Variant {
	return baseenumer.SliceVariants[Variant](variantLabels[:])
}

func Values() []string {
	return baseenumer.SliceValues(variantLabels[:])
}

func Parse(s string) result.Wrap[Variant] {
	v, trimmed, ok := baseenumer.ParseLookup(s, variantMap)
	if len(trimmed) == 0 {
		return result.WrapFailureWithId[Variant](errtype.Validation, baseenumer.FormatEmptyParseError("fileoptype"))
	}

	if ok {
		return result.WrapSuccess(v)
	}

	return result.WrapFailureWithId[Variant](
		errtype.NotFound,
		baseenumer.FormatParseError("fileoptype", s, Values()),
	)
}
