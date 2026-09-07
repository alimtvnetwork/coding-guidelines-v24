package fileutil

import (
	"os"

	"coding-guidelines/common/pkg/appfault"
)

// Default I/O chunk and buffer size constants (64 KB).
const (
	DefaultChunkSize  = 64 * 1024
	DefaultBufferSize = DefaultChunkSize
)

type (
	ChunkHandlerFunc func(chunk []byte) error

	ChunkCallbackFunc func(chunk []byte) *appfault.AppError

	BoundFileActionFunc func(w *BoundFileWriter) *appfault.AppError

	WithLockFunc = BoundFileActionFunc

	FileFilterFunc func(path string, info os.FileInfo) bool
)
