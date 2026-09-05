package fileutil

import (
	"os"
	"path/filepath"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

// OpenFile opens or creates a file using explicit mode and permission enums.
func OpenFile(path string, openMode FileOpenModeType, perm FilePermType) result.Wrap[*os.File] {
	if len(path) == 0 {
		return result.WrapFailure[*os.File](appfault.New(errtype.Validation, "path cannot be empty"))
	}

	flags := openMode.Flags()
	isCreateMode := (flags & os.O_CREATE) != 0
	if isCreateMode {
		dir := filepath.Dir(path)
		if len(dir) > 0 {
			if dir != "." {
				if err := os.MkdirAll(dir, 0755); err != nil {
					return result.WrapFailure[*os.File](appfault.Wrap(errtype.IO, err, "failed to create parent directory: "+dir))
				}
			}
		}
	}

	f, err := os.OpenFile(path, flags, perm.Mode())
	if err != nil {
		if os.IsNotExist(err) {
			return result.WrapFailure[*os.File](appfault.Wrap(errtype.NotFound, err, "file not found: "+path))
		}

		if os.IsPermission(err) {
			return result.WrapFailure[*os.File](appfault.Wrap(errtype.Forbidden, err, "permission denied: "+path))
		}

		return result.WrapFailure[*os.File](appfault.Wrap(errtype.IO, err, "failed to open file: "+path))
	}

	return result.WrapSuccess(f)
}

// Open opens a file in read-only mode with standard permissions.
func Open(path string) result.Wrap[*os.File] {
	return OpenFile(path, FileOpenReadOnly, FilePermStandard)
}

// EnsureDir recursively creates a directory path if missing.
func EnsureDir(path string, perm FilePermType) result.Wrap[bool] {
	if len(path) == 0 {
		return result.WrapFailure[bool](appfault.New(errtype.Validation, "directory path cannot be empty"))
	}

	err := os.MkdirAll(path, perm.Mode())
	if err != nil {
		return result.WrapFailure[bool](appfault.Wrap(errtype.IO, err, "failed to create directory: "+path))
	}

	return result.WrapSuccess(true)
}

// ReadAll reads entire file content into byte slice.
func ReadAll(path string) result.Wrap[[]byte] {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return result.WrapFailure[[]byte](appfault.Wrap(errtype.NotFound, err, "file not found: "+path))
		}

		return result.WrapFailure[[]byte](appfault.Wrap(errtype.IO, err, "failed to read file: "+path))
	}

	return result.WrapSuccess(data)
}

// ReadString reads entire file content as a string.
func ReadString(path string) result.Wrap[string] {
	res := ReadAll(path)
	if res.IsFailed() {
		return result.WrapFailure[string](res.Fault())
	}

	return result.WrapSuccess(string(res.Data()))
}

// WriteFile writes data to a file replacing its content.
func WriteFile(path string, data []byte, perm FilePermType) result.Wrap[bool] {
	wrap := OpenFile(path, FileOpenCreateTruncate, perm)
	if wrap.IsFailed() {
		return result.WrapFailure[bool](wrap.Fault())
	}

	f := wrap.Data()
	defer f.Close()

	_, err := f.Write(data)
	if err != nil {
		return result.WrapFailure[bool](appfault.Wrap(errtype.IO, err, "failed to write data: "+path))
	}

	return result.WrapSuccess(true)
}
