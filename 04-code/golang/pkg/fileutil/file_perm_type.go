package fileutil

import (
	"os"

	"coding-guidelines/common/pkg/enum/filepermtype"
)

type FilePermType = filepermtype.Variant

func ParsePerm(octalStr string) FilePermResult {
	return filepermtype.ParsePerm(octalStr)
}

func FromFileMode(mode os.FileMode) FilePermType {
	return filepermtype.FromFileMode(mode)
}
