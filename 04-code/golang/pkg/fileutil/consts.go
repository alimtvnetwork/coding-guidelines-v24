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

// FileOpenModeType enum constants aliased from openfiletype.
const (
	FileOpenInvalid        FileOpenModeType = openfiletype.Invalid
	FileOpenReadOnly       FileOpenModeType = openfiletype.ReadOnly
	FileOpenWriteOnly      FileOpenModeType = openfiletype.WriteOnly
	FileOpenReadWrite      FileOpenModeType = openfiletype.ReadWrite
	FileOpenAppend         FileOpenModeType = openfiletype.Append
	FileOpenCreateAppend   FileOpenModeType = openfiletype.CreateAppend
	FileOpenCreateTruncate FileOpenModeType = openfiletype.CreateTruncate
	FileOpenCreateNew      FileOpenModeType = openfiletype.CreateNew
)

// FileOpType enum constants representing file operations.
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

// FilePermType enum constants specifying POSIX filesystem permission bitmasks.
const (
	// FilePermNone represents no permissions (0000 / ---------).
	FilePermNone FilePermType = 0000

	// FilePermOwnerReadOnly represents owner read-only (0400 / r--------).
	FilePermOwnerReadOnly FilePermType = 0400

	// FilePermOwnerWriteOnly represents owner write-only (0200 / -w-------).
	FilePermOwnerWriteOnly FilePermType = 0200

	// FilePermOwnerExecOnly represents owner execute-only (0100 / --x------).
	FilePermOwnerExecOnly FilePermType = 0100

	// FilePermOwnerReadWrite represents owner read-write (0600 / rw-------).
	FilePermOwnerReadWrite FilePermType = 0600

	// FilePermPrivate is an alias for FilePermOwnerReadWrite (0600 / rw-------).
	FilePermPrivate FilePermType = 0600

	// FilePermOwnerAll represents owner full read, write, execute (0700 / rwx------).
	FilePermOwnerAll FilePermType = 0700

	// FilePermOwnerExec is an alias for FilePermOwnerAll (0700 / rwx------).
	FilePermOwnerExec FilePermType = 0700

	// FilePermGroupReadOnly represents owner & group read-only (0440 / r--r-----).
	FilePermGroupReadOnly FilePermType = 0440

	// FilePermGroupWriteOnly represents owner & group write-only (0220 / -w--w----).
	FilePermGroupWriteOnly FilePermType = 0220

	// FilePermGroupReadWrite represents owner & group read-write (0660 / rw-rw----).
	FilePermGroupReadWrite FilePermType = 0660

	// FilePermGroupExec represents owner full, group read-exec (0750 / rwxr-x---).
	FilePermGroupExec FilePermType = 0750

	// FilePermGroupAll represents owner & group full control (0770 / rwxrwx---).
	FilePermGroupAll FilePermType = 0770

	// FilePermReadOnly represents world read-only (0444 / r--r--r--).
	FilePermReadOnly FilePermType = 0444

	// FilePermPublicReadOnly is an alias for FilePermReadOnly (0444 / r--r--r--).
	FilePermPublicReadOnly FilePermType = 0444

	// FilePermPublicWriteOnly represents world write-only (0222 / -w--w--w-).
	FilePermPublicWriteOnly FilePermType = 0222

	// FilePermStandard represents standard file permissions (0644 / rw-r--r--).
	FilePermStandard FilePermType = 0644

	// FilePermGroupSharedOtherRead represents group rw, world r (0664 / rw-rw-r--).
	FilePermGroupSharedOtherRead FilePermType = 0664

	// FilePermPublicReadWrite represents world read-write (0666 / rw-rw-rw-).
	FilePermPublicReadWrite FilePermType = 0666

	// FilePermExecutable represents standard directory/executable (0755 / rwxr-xr-x).
	FilePermExecutable FilePermType = 0755

	// FilePermGroupSharedDir represents group collaborative dir (0775 / rwxrwxr-x).
	FilePermGroupSharedDir FilePermType = 0775

	// FilePermPublicAll represents full unrestricted world access (0777 / rwxrwxrwx).
	FilePermPublicAll FilePermType = 0777

	// FilePermStickyDir represents world-writable directory with sticky bit (01777 / rwxrwxrwt).
	FilePermStickyDir FilePermType = 01777

	// FilePermSetuidExec represents setuid executable (04755 / rwsr-xr-x).
	FilePermSetuidExec FilePermType = 04755

	// FilePermSetgidExec represents setgid executable/directory (02755 / rwxr-sr-x).
	FilePermSetgidExec FilePermType = 02755
)

// FileWriteModeType enum constants specifying behavior strategies for FileWriter.
const (
	// FileWriteModeDirect writes directly in-place to the target file.
	FileWriteModeDirect FileWriteModeType = 1
	// FileWriteModeAtomic writes to a temp file and atomically renames.
	FileWriteModeAtomic FileWriteModeType = 2
	// FileWriteModeTruncate creates or truncates the file before writing.
	FileWriteModeTruncate FileWriteModeType = 3
)

// ChunkHandlerFunc handles a slice of bytes during chunked file reading.
type ChunkHandlerFunc func(chunk []byte) error

// ChunkCallbackFunc handles a slice of bytes and returns *appfault.AppError during chunked file reading.
type ChunkCallbackFunc func(chunk []byte) *appfault.AppError

// BoundFileActionFunc defines an atomic transaction closure executed under a BoundFileWriter lock.
type BoundFileActionFunc func(w *BoundFileWriter) *appfault.AppError

// WithLockFunc is an alias for BoundFileActionFunc.
type WithLockFunc = BoundFileActionFunc

// FileFilterFunc filters file paths or info during directory inspection.
type FileFilterFunc func(path string, info os.FileInfo) bool
