package fileutil

import (
	"os"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/enum/openfiletype"
)

// Default I/O chunk and buffer size constants (64 KB).
const (
	DefaultChunkSize  = 64 * 1024
	DefaultBufferSize = DefaultChunkSize
)

const (
	FileOpenInvalid               FileOpenModeType = openfiletype.Invalid
	FileOpenReadOnly              FileOpenModeType = openfiletype.ReadOnly
	FileOpenWriteOnly             FileOpenModeType = openfiletype.WriteOnly
	FileOpenReadWrite             FileOpenModeType = openfiletype.ReadWrite
	FileOpenAppend                FileOpenModeType = openfiletype.Append
	FileOpenCreateAppend          FileOpenModeType = openfiletype.CreateAppend
	FileOpenCreateTruncate        FileOpenModeType = openfiletype.CreateTruncate
	FileOpenCreateNew             FileOpenModeType = openfiletype.CreateNew
	FileOpenReadOrCreateOnly      FileOpenModeType = openfiletype.ReadOrCreateOnly
	FileOpenWriteOrCreateOnly     FileOpenModeType = openfiletype.WriteOrCreateOnly
	FileOpenReadWriteOrCreateOnly FileOpenModeType = openfiletype.ReadWriteOrCreateOnly
)

type (
	ChunkHandlerFunc func(chunk []byte) error

	ChunkCallbackFunc func(chunk []byte) *appfault.AppError

	BoundFileActionFunc func(w *BoundFileWriter) *appfault.AppError

	WithLockFunc = BoundFileActionFunc

	FileFilterFunc func(path string, info os.FileInfo) bool
)
