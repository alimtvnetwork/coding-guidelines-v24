package openfiletype

import (
	"os"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type Result = result.Wrap[Variant]

var (
	variantLabels = [...]string{
		Invalid:               "Invalid",
		ReadOnly:              "ReadOnly",
		WriteOnly:             "WriteOnly",
		ReadWrite:             "ReadWrite",
		Append:                "Append",
		CreateAppend:          "CreateAppend",
		CreateTruncate:        "CreateTruncate",
		CreateNew:             "CreateNew",
		ReadOrCreateOnly:      "ReadOrCreateOnly",
		WriteOrCreateOnly:     "WriteOrCreateOnly",
		ReadWriteOrCreateOnly: "ReadWriteOrCreateOnly",
	}

	openFlags = [...]int{
		Invalid:               os.O_RDONLY,
		ReadOnly:              os.O_RDONLY,
		WriteOnly:             os.O_WRONLY,
		ReadWrite:             os.O_RDWR,
		Append:                os.O_WRONLY | os.O_APPEND,
		CreateAppend:          os.O_CREATE | os.O_WRONLY | os.O_APPEND,
		CreateTruncate:        os.O_CREATE | os.O_WRONLY | os.O_TRUNC,
		CreateNew:             os.O_CREATE | os.O_EXCL | os.O_WRONLY,
		ReadOrCreateOnly:      os.O_RDONLY | os.O_CREATE,
		WriteOrCreateOnly:     os.O_WRONLY | os.O_CREATE,
		ReadWriteOrCreateOnly: os.O_RDWR | os.O_CREATE,
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
		return result.WrapFailureWithId[Variant](errtype.Validation, baseenumer.FormatEmptyParseError("openfiletype"))
	}

	if ok {
		return result.WrapSuccess(v)
	}

	return result.WrapFailureWithId[Variant](
		errtype.NotFound,
		baseenumer.FormatParseError("openfiletype", s, Values()),
	)
}
