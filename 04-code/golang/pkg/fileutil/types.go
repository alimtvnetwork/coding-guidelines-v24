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

// FileOpType represents standard file operation choices (read, write, append, create, delete).
type FileOpType byte

const (
	// FileOpReadOnly represents reading an existing file.
	FileOpReadOnly FileOpType = iota
	// FileOpWriteOnly represents writing to a file.
	FileOpWriteOnly
	// FileOpReadWrite represents reading and writing to a file.
	FileOpReadWrite
	// FileOpAppend represents appending data to an existing or new file.
	FileOpAppend
	// FileOpCreate represents creating a new file atomically.
	FileOpCreate
	// FileOpCreateAppend represents creating a file if missing and appending writes.
	FileOpCreateAppend
	// FileOpCreateTruncate represents creating or truncating a file for writing.
	FileOpCreateTruncate
	// FileOpDelete represents deleting/removing a file.
	FileOpDelete
)

var fileOpNames = [...]string{
	"ReadOnly",
	"WriteOnly",
	"ReadWrite",
	"Append",
	"Create",
	"CreateAppend",
	"CreateTruncate",
	"Delete",
}

// Name returns the PascalCase representation of the file operation.
func (o FileOpType) Name() string {
	if int(o) < len(fileOpNames) {
		return fileOpNames[o]
	}

	return fmt.Sprintf("FileOp(%d)", byte(o))
}

// String implements fmt.Stringer.
func (o FileOpType) String() string {
	return o.Name()
}

// IsDelete returns true if the operation is a deletion.
func (o FileOpType) IsDelete() bool {
	return o == FileOpDelete
}

// IsReadOnly returns true if the operation is read-only.
func (o FileOpType) IsReadOnly() bool {
	return o == FileOpReadOnly
}

// IsAppend returns true if the operation is append or create-append.
func (o FileOpType) IsAppend() bool {
	if o == FileOpAppend {
		return true
	}

	return o == FileOpCreateAppend
}

// OpenMode returns the corresponding FileOpenModeType for the operation.
func (o FileOpType) OpenMode() FileOpenModeType {
	switch o {
	case FileOpWriteOnly:
		return FileOpenWriteOnly
	case FileOpReadWrite:
		return FileOpenReadWrite
	case FileOpAppend:
		return FileOpenAppend
	case FileOpCreate:
		return FileOpenCreateNew
	case FileOpCreateAppend:
		return FileOpenCreateAppend
	case FileOpCreateTruncate:
		return FileOpenCreateTruncate
	default:
		return FileOpenReadOnly
	}
}
