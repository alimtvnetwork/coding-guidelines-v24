package fileutil

type createOps struct{}

func (createOps) File(path string, perm FilePermType) FileResult {
	return CreateFile(path, perm)
}

func (createOps) Dir(path string, perm FilePermType) BoolResult {
	return CreateDir(path, perm)
}

func (createOps) EnsureDir(path string, perm FilePermType) BoolResult {
	return EnsureDir(path, perm)
}

func (createOps) Temp(pattern string) FileResult {
	return TempFile(pattern)
}

func (createOps) TempDir(pattern string) StringResult {
	return TempDir(pattern)
}

func (createOps) TempFileIn(dir string, pattern string, perm FilePermType) FileResult {
	return CreateTempFile(dir, pattern, perm)
}

func (createOps) TempDirIn(dir string, pattern string, perm FilePermType) StringResult {
	return CreateTempDir(dir, pattern, perm)
}

func Create(path string, perm FilePermType) FileResult {
	return CreateFile(path, perm)
}
