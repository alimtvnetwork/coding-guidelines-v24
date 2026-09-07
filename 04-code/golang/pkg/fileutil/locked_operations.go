package fileutil

import (
	"coding-guidelines/common/pkg/result"
)

// ReadTextLocked acquires a read lock for the file before executing ReadText.
func ReadTextLocked(path string) StringResult {
	lock := GetFileLock(path)
	defer ReleaseFileLock(path)
	lock.RLock()
	defer lock.RUnlock()

	return ReadText(path)
}

// ReadLinesLocked acquires a read lock for the file before executing ReadLines.
func ReadLinesLocked(path string) LinesResult {
	lock := GetFileLock(path)
	defer ReleaseFileLock(path)
	lock.RLock()
	defer lock.RUnlock()

	return ReadLines(path)
}

// ReadJsonLocked acquires a read lock for the file before executing ReadJson.
func ReadJsonLocked[T any](path string) result.Wrap[T] {
	lock := GetFileLock(path)
	defer ReleaseFileLock(path)
	lock.RLock()
	defer lock.RUnlock()

	return ReadJson[T](path)
}

// ReadYamlLocked acquires a read lock for the file before executing ReadYaml.
func ReadYamlLocked[T any](path string) result.Wrap[T] {
	lock := GetFileLock(path)
	defer ReleaseFileLock(path)
	lock.RLock()
	defer lock.RUnlock()

	return ReadYaml[T](path)
}

// ReadJSONLocked is an alias for ReadJsonLocked.
func ReadJSONLocked[T any](path string) result.Wrap[T] {
	return ReadJsonLocked[T](path)
}

// ReadYAMLLocked is an alias for ReadYamlLocked.
func ReadYAMLLocked[T any](path string) result.Wrap[T] {
	return ReadYamlLocked[T](path)
}

// ExportTextLocked acquires an exclusive write lock for the file before executing ExportText.
func ExportTextLocked(path string, content string, perm FilePermType) BoolResult {
	lock := GetFileLock(path)
	defer ReleaseFileLock(path)
	lock.Lock()
	defer lock.Unlock()

	return ExportText(path, content, perm)
}

// ExportLinesLocked acquires an exclusive write lock for the file before executing ExportLines.
func ExportLinesLocked(path string, lines []string, perm FilePermType) BoolResult {
	lock := GetFileLock(path)
	defer ReleaseFileLock(path)
	lock.Lock()
	defer lock.Unlock()

	return ExportLines(path, lines, perm)
}

// ExportJsonLocked acquires an exclusive write lock for the file before executing ExportJson.
func ExportJsonLocked(path string, data any, perm FilePermType) BoolResult {
	lock := GetFileLock(path)
	defer ReleaseFileLock(path)
	lock.Lock()
	defer lock.Unlock()

	return ExportJson(path, data, perm)
}

// ExportYamlLocked acquires an exclusive write lock for the file before executing ExportYaml.
func ExportYamlLocked(path string, data any, perm FilePermType) BoolResult {
	lock := GetFileLock(path)
	defer ReleaseFileLock(path)
	lock.Lock()
	defer lock.Unlock()

	return ExportYaml(path, data, perm)
}

// ExportJSONLocked is an alias for ExportJsonLocked.
func ExportJSONLocked(path string, data any, perm FilePermType) BoolResult {
	return ExportJsonLocked(path, data, perm)
}

// ExportYAMLLocked is an alias for ExportYamlLocked.
func ExportYAMLLocked(path string, data any, perm FilePermType) BoolResult {
	return ExportYamlLocked(path, data, perm)
}
