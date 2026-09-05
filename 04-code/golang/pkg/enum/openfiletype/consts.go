package openfiletype

// Enum constants for openfiletype conforming to aukgo and global standards.
const (
	// Invalid represents an uninitialized or invalid file open mode (iota zero-value).
	Invalid Variant = iota

	// ReadOnly opens the file in read-only mode (os.O_RDONLY).
	ReadOnly

	// WriteOnly opens the file in write-only mode (os.O_WRONLY).
	WriteOnly

	// ReadWrite opens the file for reading and writing (os.O_RDWR).
	ReadWrite

	// Append opens the file in write-only append mode (os.O_WRONLY | os.O_APPEND).
	Append

	// CreateAppend creates the file if missing and appends writes (os.O_CREATE | os.O_WRONLY | os.O_APPEND).
	CreateAppend

	// CreateTruncate creates or truncates the file for writing (os.O_CREATE | os.O_WRONLY | os.O_TRUNC).
	CreateTruncate

	// CreateNew creates a new file atomically, failing if it already exists (os.O_CREATE | os.O_EXCL | os.O_WRONLY).
	CreateNew

	// ReadOrCreateOnly opens the file in read-only mode, creating it if it does not exist (os.O_RDONLY | os.O_CREATE).
	ReadOrCreateOnly

	// WriteOrCreateOnly opens the file in write-only mode, creating it if it does not exist (os.O_WRONLY | os.O_CREATE).
	WriteOrCreateOnly

	// ReadWriteOrCreateOnly opens the file in read-write mode, creating it if it does not exist (os.O_RDWR | os.O_CREATE).
	ReadWriteOrCreateOnly
)

// VariantPredicate defines a filter or condition check over a Variant.
type VariantPredicate func(v Variant) bool
