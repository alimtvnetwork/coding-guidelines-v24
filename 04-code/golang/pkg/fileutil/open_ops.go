package fileutil

import (
	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/enum/openfiletype"
)

type openOps struct{}

func (openOps) File(path string, openMode FileOpenModeType, perm FilePermType) FileResult {
	return OpenFile(path, openMode, perm)
}

func (openOps) ReadOnly(path string) FileResult {
	return OpenFile(path, openfiletype.ReadOnly, filepermtype.Standard)
}

func (openOps) ReadWrite(path string, perm FilePermType) FileResult {
	return OpenFile(path, openfiletype.ReadWrite, perm)
}

func (openOps) Append(path string, perm FilePermType) FileResult {
	return OpenFile(path, openfiletype.Append, perm)
}

func (openOps) Truncate(path string, perm FilePermType) FileResult {
	return OpenFile(path, openfiletype.CreateTruncate, perm)
}

func (openOps) CreateAppend(path string, perm FilePermType) FileResult {
	return OpenFile(path, openfiletype.CreateAppend, perm)
}

func (openOps) Write(path string, perm FilePermType) FileResult {
	return OpenFile(path, openfiletype.CreateTruncate, perm)
}

func OpenAppend(path string, perm FilePermType) FileResult {
	return OpenFile(path, openfiletype.CreateAppend, perm)
}

func OpenWrite(path string, perm FilePermType) FileResult {
	return OpenFile(path, openfiletype.CreateTruncate, perm)
}
