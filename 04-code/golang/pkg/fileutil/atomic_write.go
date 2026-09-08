package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/errtype"
)

// AtomicWriteFile writes data to a temporary file, syncs to disk, and replaces target file.
func AtomicWriteFile(
	filePath string,
	data []byte,
	perm filepermtype.Variant,
) *appfault.AppError {
	if filePath == "" {
		return appfault.New(errtype.Validation, "filePath cannot be empty")
	}

	dir := filepath.Dir(filePath)
	dirRes := EnsureDir(dir, filepermtype.Standard)
	if dirRes.IsFailure() {
		return dirRes.Fault()
	}

	tmpPath := generateTempPath(dir)

	return writeAndCommitTemp(filePath, tmpPath, data, perm)
}

func generateTempPath(dir string) string {
	name := fmt.Sprintf(".tmp-%d-%d", os.Getpid(), time.Now().UnixNano())

	return filepath.Join(dir, name)
}

func writeAndCommitTemp(
	targetPath, tmpPath string,
	data []byte,
	perm filepermtype.Variant,
) *appfault.AppError {
	f, err := os.OpenFile(tmpPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm.Mode())
	if err != nil {
		return appfault.WrapFile(errtype.IO, err, tmpPath, "failed to create atomic temp file")
	}

	if fault := writeAndSync(f, data, tmpPath); fault != nil {
		_ = os.Remove(tmpPath)

		return fault
	}

	return commitTempFile(tmpPath, targetPath)
}

func writeAndSync(f *os.File, data []byte, path string) *appfault.AppError {
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return appfault.WrapFile(errtype.IO, err, path, "failed writing atomic temp file")
	}

	if err := f.Sync(); err != nil {
		return appfault.WrapFile(errtype.IO, err, path, "failed syncing atomic temp file")
	}

	return nil
}

func commitTempFile(tmpPath, targetPath string) *appfault.AppError {
	err := replaceFile(tmpPath, targetPath)
	if err != nil {
		_ = os.Remove(tmpPath)

		return appfault.WrapFile(errtype.IO, err, targetPath, "failed to atomically commit file")
	}

	return nil
}

func replaceFile(tmpPath, targetPath string) error {
	err := os.Rename(tmpPath, targetPath)
	if err == nil {
		return nil
	}

	_ = os.Remove(targetPath)

	return os.Rename(tmpPath, targetPath)
}

// AtomicWrite writes data atomically and returns a BoolResult container.
func AtomicWrite(
	filePath string,
	data []byte,
	perm filepermtype.Variant,
) BoolResult {
	if fault := AtomicWriteFile(filePath, data, perm); fault != nil {
		return BoolFailureFault(fault)
	}

	return BoolSuccess(true)
}
