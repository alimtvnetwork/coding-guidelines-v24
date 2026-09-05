package openfiletype

import (
	"encoding/json"
	"fmt"
	"os"
)

// Variant represents the file open mode enum backed by byte.
type Variant byte

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

// Flags converts the variant to standard os.OpenFile integer flags.
func (v Variant) Flags() int {
	if int(v) < len(openFlags) {
		return openFlags[v]
	}

	return os.O_RDONLY
}

// Name returns the PascalCase label.
func (v Variant) Name() string {
	if int(v) < len(variantLabels) {
		return variantLabels[v]
	}

	return fmt.Sprintf("OpenFile(%d)", byte(v))
}

// Label delegates to Name.
func (v Variant) Label() string {
	return v.Name()
}

// String implements fmt.Stringer returning the PascalCase label.
func (v Variant) String() string {
	return v.Name()
}

// IsValid returns true if the variant is within defined valid bounds (non-zero).
func (v Variant) IsValid() bool {
	return v > Invalid && int(v) < len(variantLabels)
}

// IsInvalid returns true if the variant is the zero value or undefined.
func (v Variant) IsInvalid() bool {
	return v <= Invalid || int(v) >= len(variantLabels)
}

// IsReadOnly checks if mode is ReadOnly.
func (v Variant) IsReadOnly() bool {
	return v == ReadOnly
}

// IsWriteOnly checks if mode is WriteOnly.
func (v Variant) IsWriteOnly() bool {
	return v == WriteOnly
}

// IsReadWrite checks if mode is ReadWrite.
func (v Variant) IsReadWrite() bool {
	return v == ReadWrite
}

// IsAppend checks if mode is Append.
func (v Variant) IsAppend() bool {
	return v == Append
}

// IsCreateAppend checks if mode is CreateAppend.
func (v Variant) IsCreateAppend() bool {
	return v == CreateAppend
}

// IsCreateTruncate checks if mode is CreateTruncate.
func (v Variant) IsCreateTruncate() bool {
	return v == CreateTruncate
}

// IsCreateNew checks if mode is CreateNew.
func (v Variant) IsCreateNew() bool {
	return v == CreateNew
}

// IsReadOrCreateOnly checks if mode is ReadOrCreateOnly.
func (v Variant) IsReadOrCreateOnly() bool {
	return v == ReadOrCreateOnly
}

// IsWriteOrCreateOnly checks if mode is WriteOrCreateOnly.
func (v Variant) IsWriteOrCreateOnly() bool {
	return v == WriteOrCreateOnly
}

// IsReadWriteOrCreateOnly checks if mode is ReadWriteOrCreateOnly.
func (v Variant) IsReadWriteOrCreateOnly() bool {
	return v == ReadWriteOrCreateOnly
}

// MarshalJSON implements json.Marshaler serializing the variant as a PascalCase string.
func (v Variant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Name())
}

// UnmarshalJSON implements json.Unmarshaler supporting string or byte unmarshaling.
func (v *Variant) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		res := Parse(str)
		if res.IsSuccess() {
			*v = res.Data()

			return nil
		}
	}

	var raw byte
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*v = Variant(raw)

	return nil
}
