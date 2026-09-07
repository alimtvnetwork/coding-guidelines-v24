package filewritemodetype

import (
	"encoding/json"
	"fmt"
	"strings"

	"coding-guidelines/common/pkg/baseenumer"
)

type (
	Variant uint8

	FileWriteModeType = Variant

	VariantPredicate func(v Variant) bool
)

const (
	Invalid Variant = iota
	Direct
	Atomic
	Truncate
)

const (
	FileWriteModeInvalid  = Invalid
	FileWriteModeDirect   = Direct
	FileWriteModeAtomic   = Atomic
	FileWriteModeTruncate = Truncate
)

var (
	_ baseenumer.BaseEnumer   = Variant(0)
	_ baseenumer.ByteEnumer   = Variant(0)
	_ baseenumer.NumberEnumer = Variant(0)
	_ json.Marshaler          = Variant(0)
	_ json.Unmarshaler        = (*Variant)(nil)
)

func (v Variant) Byte() byte {
	return byte(v)
}

func (v Variant) ValueByte() byte {
	return byte(v)
}

func (v Variant) Bytes() []byte {
	return []byte{byte(v)}
}

func (v Variant) Int() int {
	return int(v)
}

func (v Variant) Code() uint16 {
	return uint16(v)
}

func (v Variant) IsValid() bool {
	return baseenumer.IsBetween(v, Direct, Truncate)
}

func (v Variant) IsInvalid() bool {
	return baseenumer.IsNotBetween(v, Direct, Truncate)
}

func (v Variant) IsEnum() bool {
	return v.IsValid()
}

func (v Variant) IsDirect() bool {
	return v == Direct
}

func (v Variant) IsAtomic() bool {
	return v == Atomic
}

func (v Variant) IsTruncate() bool {
	return v == Truncate
}

func (v Variant) Name() string {
	if int(v) < len(variantLabels) {
		return variantLabels[v]
	}

	return fmt.Sprintf("FileWriteMode(%d)", uint8(v))
}

func (v Variant) Label() string {
	return v.Name()
}

func (v Variant) String() string {
	return v.Name()
}

func (v Variant) ValueString() string {
	return baseenumer.FormatNameValue(v.Name(), byte(v))
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

	return v.unmarshalData(data)
}

func (v *Variant) unmarshalData(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		return v.unmarshalString(str)
	}

	var raw byte
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if int(raw) >= len(variantLabels) {
		return baseenumer.FormatNumericRangeError("filewritemodetype", raw, len(variantLabels)-1)
	}

	*v = Variant(raw)

	return nil
}

func (v *Variant) unmarshalString(str string) error {
	res := Parse(str)
	if res.IsSuccess() {
		*v = res.Data()

		return nil
	}

	return res.Fault()
}
