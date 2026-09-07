package fileutil

import (
	"os"
	"path/filepath"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

func isValidParentDir(dir string) bool {
	if len(dir) == 0 {
		return false
	}

	return dir != "."
}

func ensureParentDir(path string, flags int) *appfault.AppError {
	if (flags & os.O_CREATE) == 0 {
		return nil
	}

	dir := filepath.Dir(path)
	if !isValidParentDir(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return appfault.WrapFile(errtype.IO, err, dir, "failed to create parent directory")
	}

	return nil
}

func openFileError(err error, path string) FileResult {
	if os.IsNotExist(err) {
		return FileFailure(errtype.NotFound, err, path, "file not found")
	}

	if os.IsPermission(err) {
		return FileFailure(errtype.Forbidden, err, path, "permission denied")
	}

	return FileFailure(errtype.IO, err, path, "failed to open file")
}

func OpenFile(path string, openMode FileOpenModeType, perm FilePermType) FileResult {
	if len(path) == 0 {
		return FileFailureMsg(errtype.Validation, path, "path cannot be empty")
	}

	if err := ensureParentDir(path, openMode.Flags()); err != nil {
		return result.WrapFailure[*os.File](err)
	}

	f, err := os.OpenFile(path, openMode.Flags(), perm.Mode())
	if err != nil {
		return openFileError(err, path)
	}

	return FileSuccess(f)
}

func Open(path string) FileResult {
	return OpenFile(path, FileOpenReadOnly, FilePermStandard)
}

func EnsureDir(path string, perm FilePermType) BoolResult {
	if len(path) == 0 {
		return BoolFailureMsg(errtype.Validation, path, "directory path cannot be empty")
	}

	if err := os.MkdirAll(path, perm.Mode()); err != nil {
		return BoolFailure(errtype.IO, err, path, "failed to create directory")
	}

	return BoolSuccess(true)
}

func readAllError(err error, path string) BytesResult {
	if os.IsNotExist(err) {
		return BytesFailure(errtype.NotFound, err, path, "file not found")
	}

	return BytesFailure(errtype.IO, err, path, "failed to read file")
}

func ReadAll(path string) BytesResult {
	data, err := os.ReadFile(path)
	if err != nil {
		return readAllError(err, path)
	}

	return BytesSuccess(data)
}

func ReadString(path string) StringResult {
	res := ReadAll(path)
	if res.IsFailed() {
		return result.WrapFailure[string](res.Fault())
	}

	return StringSuccess(string(res.Data()))
}

func WriteFile(path string, data []byte, perm FilePermType) BoolResult {
	wrap := OpenFile(path, FileOpenCreateTruncate, perm)
	if wrap.IsFailed() {
		return result.WrapFailure[bool](wrap.Fault())
	}

	defer wrap.Data().Close()

	if _, err := wrap.Data().Write(data); err != nil {
		return BoolFailure(errtype.IO, err, path, "failed to write data")
	}

	return BoolSuccess(true)
}

func deleteFileError(err error, path string) BoolResult {
	if os.IsNotExist(err) {
		return BoolFailure(errtype.NotFound, err, path, "file not found")
	}

	if os.IsPermission(err) {
		return BoolFailure(errtype.Forbidden, err, path, "permission denied")
	}

	return BoolFailure(errtype.IO, err, path, "failed to delete file")
}

func DeleteFile(path string) BoolResult {
	if len(path) == 0 {
		return BoolFailureMsg(errtype.Validation, path, "path cannot be empty")
	}

	if err := os.Remove(path); err != nil {
		return deleteFileError(err, path)
	}

	return BoolSuccess(true)
}

func Remove(path string) BoolResult {
	return DeleteFile(path)
}

func removeAllError(err error, path string) BoolResult {
	if os.IsPermission(err) {
		return BoolFailure(errtype.Forbidden, err, path, "permission denied")
	}

	return BoolFailure(errtype.IO, err, path, "failed to remove path")
}

func RemoveAll(path string) BoolResult {
	if len(path) == 0 {
		return BoolFailureMsg(errtype.Validation, path, "path cannot be empty")
	}

	if err := os.RemoveAll(path); err != nil {
		return removeAllError(err, path)
	}

	return BoolSuccess(true)
}

func ReadFile(path string) BytesResult {
	return ReadAll(path)
}

func statError(err error, path string) FileInfoResult {
	if os.IsNotExist(err) {
		return FileInfoFailure(errtype.NotFound, err, path, "file not found")
	}

	if os.IsPermission(err) {
		return FileInfoFailure(errtype.Forbidden, err, path, "permission denied")
	}

	return FileInfoFailure(errtype.IO, err, path, "failed to stat file")
}

func Stat(path string) FileInfoResult {
	if len(path) == 0 {
		return FileInfoFailureMsg(errtype.Validation, path, "path cannot be empty")
	}

	info, err := os.Stat(path)
	if err != nil {
		return statError(err, path)
	}

	return FileInfoSuccess(info)
}

func FileSize(path string) Int64Result {
	statRes := Stat(path)
	if statRes.IsFailed() {
		return result.WrapFailure[int64](statRes.Fault())
	}

	return Int64Success(statRes.Data().Size())
}

func writeOpData(f *os.File, path string, data []byte) BytesResult {
	defer f.Close()

	if len(data) > 0 {
		if _, err := f.Write(data); err != nil {
			return BytesFailure(errtype.IO, err, path, "failed to write during op")
		}
	}

	return BytesSuccess(data)
}

func executeWriteOp(path string, op FileOpType, perm FilePermType, data []byte) BytesResult {
	openRes := OpenFile(path, op.OpenMode(), perm)
	if openRes.IsFailed() {
		return result.WrapFailure[[]byte](openRes.Fault())
	}

	return writeOpData(openRes.Data(), path, data)
}

func ExecuteOp(path string, op FileOpType, perm FilePermType, data []byte) BytesResult {
	if op.IsDelete() {
		delRes := DeleteFile(path)
		if delRes.IsFailed() {
			return result.WrapFailure[[]byte](delRes.Fault())
		}

		return BytesSuccess(nil)
	}

	if op.IsReadOnly() {
		return ReadAll(path)
	}

	return executeWriteOp(path, op, perm, data)
}
