package fileutil

import (
	"coding-guidelines/common/pkg/enum/filewritemodetype"
)

type FileWriteModeType = filewritemodetype.Variant

func ParseFileWriteMode(s string) FileWriteModeResult {
	return filewritemodetype.Parse(s)
}
