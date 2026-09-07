package openfiletype

import (
	"encoding/json"
	"os"
	"strings"

	"coding-guidelines/common/pkg/baseenumer"
)

type (
	Variant byte

	VariantPredicate func(v Variant) bool
)

const (
	Invalid Variant = iota
	ReadOnly
	WriteOnly
	ReadWrite
	Append
	CreateAppend
	CreateTruncate
	CreateNew
	ReadOrCreateOnly
	WriteOrCreateOnly
	ReadWriteOrCreateOnly
)

func (v Variant) Flags() int {
	if int(v) < len(openFlags) {
		return openFlags[v]
	}

	return os.O_RDONLY
}

func (v Variant) Name() string {
	if int(v) < len(variantLabels) {
		return variantLabels[v]
	}

	return baseenumer.FormatNameValue("OpenFile", byte(v))
}

func (v Variant) Label() string {
	return v.Name()
}

func (v Variant) String() string {
	return v.Name()
}

func (v Variant) IsValid() bool {
	return baseenumer.IsBetween(v, ReadOnly, ReadWriteOrCreateOnly)
}

func (v Variant) IsInvalid() bool {
	return baseenumer.IsNotBetween(v, ReadOnly, ReadWriteOrCreateOnly)
}

func (v Variant) IsReadOnly() bool {
	return v == ReadOnly
}

func (v Variant) IsWriteOnly() bool {
	return v == WriteOnly
}

func (v Variant) IsReadWrite() bool {
	return v == ReadWrite
}

func (v Variant) IsAppend() bool {
	return v == Append
}

func (v Variant) IsCreateAppend() bool {
	return v == CreateAppend
}

func (v Variant) IsCreateTruncate() bool {
	return v == CreateTruncate
}

func (v Variant) IsCreateNew() bool {
	return v == CreateNew
}

func (v Variant) IsReadOrCreateOnly() bool {
	return v == ReadOrCreateOnly
}

func (v Variant) IsWriteOrCreateOnly() bool {
	return v == WriteOrCreateOnly
}

func (v Variant) IsReadWriteOrCreateOnly() bool {
	return v == ReadWriteOrCreateOnly
}

func (v Variant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Name())
}

func (v *Variant) UnmarshalJSON(data []byte) error {
	trimmed := strings.TrimSpace(string(data))
	if len(trimmed) == 0 || trimmed == "null" {
		*v = Invalid

		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		res := Parse(str)
		if res.IsSuccess() {
			*v = res.Data()

			return nil
		}

		return res.Fault()
	}

	var raw byte
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if int(raw) >= len(variantLabels) {
		return baseenumer.FormatNumericRangeError("openfiletype", raw, len(variantLabels)-1)
	}

	*v = Variant(raw)

	return nil
}
