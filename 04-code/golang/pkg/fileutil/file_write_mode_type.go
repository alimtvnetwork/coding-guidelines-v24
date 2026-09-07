package fileutil

import (
	"coding-guidelines/common/pkg/enum/filewritemodetype"
	"coding-guidelines/common/pkg/result"
)

type FileWriteModeType = filewritemodetype.Variant

const (
	FileWriteModeInvalid  = filewritemodetype.Invalid
	FileWriteModeDirect   = filewritemodetype.Direct
	FileWriteModeAtomic   = filewritemodetype.Atomic
	FileWriteModeTruncate = filewritemodetype.Truncate
)

func ParseFileWriteMode(s string) result.Wrap[FileWriteModeType] {
	return filewritemodetype.Parse(s)
}
