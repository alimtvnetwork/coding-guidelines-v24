package fileutil

import (
	"coding-guidelines/common/pkg/enum/fileoptype"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type FileOpType = fileoptype.Variant

func ParseFileOp(s string) FileOpResult {
	val, isOk := fileoptype.Parse(s)
	if !isOk {
		return result.WrapFailureWithId[FileOpType](errtype.NotFound, "invalid file op: "+s)
	}

	return result.WrapSuccess(val)
}
