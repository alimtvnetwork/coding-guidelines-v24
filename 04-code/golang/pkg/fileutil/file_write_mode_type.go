package fileutil

import (
	"coding-guidelines/common/pkg/enum/filewritemodetype"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type FileWriteModeType = filewritemodetype.Variant

func ParseFileWriteMode(s string) FileWriteModeResult {
	val, isOk := filewritemodetype.Parse(s)
	if !isOk {
		return result.WrapFailureWithId[FileWriteModeType](errtype.NotFound, "invalid file write mode: "+s)
	}

	return result.WrapSuccess(val)
}
