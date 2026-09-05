package fileutil

import (
	"fmt"
	"os"
)

// FileOpenModeType specifies the filesystem open flags for file descriptors.
type FileOpenModeType byte

const (
	// FileOpenReadOnly opens the file in read-only mode.
	FileOpenReadOnly FileOpenModeType = iota
	// FileOpenWriteOnly opens the file in write-only mode.
	FileOpenWriteOnly
	// FileOpenReadWrite opens the file for reading and writing.
	FileOpenReadWrite
	// FileOpenAppend opens the file for appending data.
	FileOpenAppend
	// FileOpenCreateAppend creates the file if missing and appends writes.
	FileOpenCreateAppend
	// FileOpenCreateTruncate creates or truncates the file for writing.
	FileOpenCreateTruncate
	// FileOpenCreateNew creates a new file atomically, failing if it exists.
	FileOpenCreateNew
)

var openModeNames = [...]string{
	"ReadOnly",
	"WriteOnly",
	"ReadWrite",
	"Append",
	"CreateAppend",
	"CreateTruncate",
	"CreateNew",
}

// Flags maps the enum to standard os.OpenFile integer flags.
func (m FileOpenModeType) Flags() int {
	switch m {
	case FileOpenWriteOnly:
		return os.O_WRONLY
	case FileOpenReadWrite:
		return os.O_RDWR
	case FileOpenAppend:
		return os.O_WRONLY | os.O_APPEND
	case FileOpenCreateAppend:
		return os.O_CREATE | os.O_WRONLY | os.O_APPEND
	case FileOpenCreateTruncate:
		return os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	case FileOpenCreateNew:
		return os.O_CREATE | os.O_EXCL | os.O_WRONLY
	case FileOpenReadOnly:
		return os.O_RDONLY
	default:
		return os.O_RDONLY
	}
}

// Name returns the PascalCase representation of the open mode.
func (m FileOpenModeType) Name() string {
	if int(m) < len(openModeNames) {
		return openModeNames[m]
	}

	return fmt.Sprintf("FileOpenMode(%d)", byte(m))
}

// String implements fmt.Stringer.
func (m FileOpenModeType) String() string {
	return m.Name()
}
