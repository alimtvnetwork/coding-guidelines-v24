package fileutil

import (
	"os"

	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/result"
)

type FilePermType = filepermtype.Variant

func ParsePerm(octalStr string) result.Wrap[FilePermType] {
	return filepermtype.ParsePerm(octalStr)
}

func FromFileMode(mode os.FileMode) FilePermType {
	return filepermtype.FromFileMode(mode)
}
