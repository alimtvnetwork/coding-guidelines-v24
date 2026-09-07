package fileutil

import (
	"coding-guidelines/common/pkg/enum/filewritemodetype"
	"coding-guidelines/common/pkg/result"
)

type FileWriteModeType = filewritemodetype.Variant

func ParseFileWriteMode(s string) result.Wrap[FileWriteModeType] {
	return filewritemodetype.Parse(s)
}
