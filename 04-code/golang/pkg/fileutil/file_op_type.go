package fileutil

import (
	"coding-guidelines/common/pkg/enum/fileoptype"
)

type FileOpType = fileoptype.Variant

func ParseFileOp(s string) FileOpResult {
	return fileoptype.Parse(s)
}
