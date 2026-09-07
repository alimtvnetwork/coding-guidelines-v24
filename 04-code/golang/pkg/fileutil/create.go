package fileutil

// CreateDir is a direct alias to EnsureDir that explicitly implies creating a folder.
func CreateDir(path string, perm FilePermType) BoolResult {
	return EnsureDir(path, perm)
}

// CreateFile creates a file at the given path with write-only and truncate flags, creating it if it doesn't exist.
func CreateFile(path string, perm FilePermType) FileResult {
	return OpenFile(path, FileOpenCreateTruncate, perm)
}
