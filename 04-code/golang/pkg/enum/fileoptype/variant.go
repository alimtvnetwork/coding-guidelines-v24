package fileoptype

import (
	"encoding/json"
	"fmt"
	"strings"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/enum/openfiletype"
)

type (
	Variant byte

	FileOpType = Variant

	VariantPredicate func(v Variant) bool
)

const (
	Invalid Variant = iota
	ReadOnly
	WriteOnly
	ReadWrite
	Append
	Create
	CreateAppend
	CreateTruncate
	Delete
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
	return baseenumer.IsBetween(v, ReadOnly, Delete)
}

func (v Variant) IsInvalid() bool {
	return baseenumer.IsNotBetween(v, ReadOnly, Delete)
}

func (v Variant) IsEnum() bool {
	return v.IsValid()
}

func (v Variant) IsDelete() bool {
	return v == Delete
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
	if v == Append {
		return true
	}

	return v == CreateAppend
}

func (v Variant) IsCreate() bool {
	return v == Create
}

func (v Variant) IsCreateAppend() bool {
	return v == CreateAppend
}

func (v Variant) IsCreateTruncate() bool {
	return v == CreateTruncate
}

func (v Variant) createOpenMode() openfiletype.Variant {
	switch v {
	case Create:
		return openfiletype.CreateNew
	case CreateAppend:
		return openfiletype.CreateAppend
	case CreateTruncate:
		return openfiletype.CreateTruncate
	default:
		return openfiletype.ReadOnly
	}
}

func (v Variant) OpenMode() openfiletype.Variant {
	switch v {
	case WriteOnly:
		return openfiletype.WriteOnly
	case ReadWrite:
		return openfiletype.ReadWrite
	case Append:
		return openfiletype.Append
	default:
		return v.createOpenMode()
	}
}

func (v Variant) Name() string {
	if int(v) < len(variantLabels) {
		return variantLabels[v]
	}

	return fmt.Sprintf("FileOp(%d)", byte(v))
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
		return baseenumer.FormatNumericRangeError("fileoptype", raw, len(variantLabels)-1)
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
