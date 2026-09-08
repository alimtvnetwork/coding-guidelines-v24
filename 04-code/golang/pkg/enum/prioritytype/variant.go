package prioritytype

import (
	"encoding/json"
	"fmt"

	"coding-guidelines/common/pkg/baseenumer"
)

type (
	Variant byte

	PriorityType = Variant

	VariantPredicate func(v Variant) bool
)

const (
	Unknown Variant = iota
	Low
	Normal
	High
	Critical
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
	return baseenumer.IsBetween(v, Low, Critical)
}

func (v Variant) IsInvalid() bool {
	return baseenumer.IsNotBetween(v, Low, Critical)
}

func (v Variant) IsEnum() bool {
	return v.IsValid()
}

func (v Variant) IsUnknown() bool {
	return v == Unknown
}

func (v Variant) IsLow() bool {
	return v == Low
}

func (v Variant) IsNormal() bool {
	return v == Normal
}

func (v Variant) IsHigh() bool {
	return v == High
}

func (v Variant) IsCritical() bool {
	return v == Critical
}

func (v Variant) Name() string {
	if int(v) < len(variantLabels) {
		return variantLabels[v]
	}

	return fmt.Sprintf("Priority(%d)", byte(v))
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
	return baseenumer.MarshalJSON(v.Name())
}

func (v *Variant) UnmarshalJSON(data []byte) error {
	return basicEnum.UnmarshalJSON(data, v)
}
