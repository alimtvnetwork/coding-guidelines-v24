package fileutil

type openOps struct{}

func (openOps) File(path string, openMode FileOpenModeType, perm FilePermType) FileResult {
	return OpenFile(path, openMode, perm)
}

func (openOps) ReadOnly(path string) FileResult {
	return OpenFile(path, FileOpenReadOnly, FilePermStandard)
}

func (openOps) ReadWrite(path string, perm FilePermType) FileResult {
	return OpenFile(path, FileOpenReadWrite, perm)
}

func (openOps) Append(path string, perm FilePermType) FileResult {
	return OpenFile(path, FileOpenAppend, perm)
}

func (openOps) Truncate(path string, perm FilePermType) FileResult {
	return OpenFile(path, FileOpenCreateTruncate, perm)
}

func (openOps) CreateAppend(path string, perm FilePermType) FileResult {
	return OpenFile(path, FileOpenCreateAppend, perm)
}

func (openOps) Write(path string, perm FilePermType) FileResult {
	return OpenFile(path, FileOpenCreateTruncate, perm)
}

func OpenAppend(path string, perm FilePermType) FileResult {
	return OpenFile(path, FileOpenCreateAppend, perm)
}

func OpenWrite(path string, perm FilePermType) FileResult {
	return OpenFile(path, FileOpenCreateTruncate, perm)
}
