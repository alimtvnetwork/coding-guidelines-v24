package fileutil

import (
	"coding-guidelines/common/pkg/enum/fileoptype"
	"coding-guidelines/common/pkg/result"
)

type FileOpType = fileoptype.Variant

func ParseFileOp(s string) result.Wrap[FileOpType] {
	return fileoptype.Parse(s)
}
