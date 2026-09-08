package fileutil

import (
	"os"

	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type FilePermType = filepermtype.Variant

func ParsePerm(octalStr string) FilePermResult {
	val, isOk := filepermtype.ParsePerm(octalStr)
	if !isOk {
		return result.WrapFailureWithId[FilePermType](errtype.Validation, "invalid octal permission: "+octalStr)
	}

	return result.WrapSuccess(val)
}

func FromFileMode(mode os.FileMode) FilePermType {
	return filepermtype.FromFileMode(mode)
}
