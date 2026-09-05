package openfiletype

import (
	"encoding/json"
	"fmt"
	"os"
)

// Variant represents the file open mode enum backed by byte.
type Variant byte

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
