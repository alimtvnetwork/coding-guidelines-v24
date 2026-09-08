package severitytype

import (
	"encoding/json"
	"fmt"

	"coding-guidelines/common/pkg/baseenumer"
)

type (
	Variant byte

	SeverityType = Variant

	VariantPredicate func(v Variant) bool
)

const (
	Unknown Variant = iota
	Info
	Warn
	Error
	Critical
	Fatal
)

var (
	_ baseenumer.BaseEnumer             = Variant(0)
	_ baseenumer.ByteEnumer             = Variant(0)
	_ baseenumer.NumberEnumer           = Variant(0)
	_ baseenumer.MinMaxer[Variant]      = Variant(0)
	_ baseenumer.BoundedEnumer[Variant] = Variant(0)
	_ baseenumer.Bounder[Variant]       = Variant(0)
	_ json.Marshaler                    = Variant(0)
	_ json.Unmarshaler                  = (*Variant)(nil)
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
	return baseenumer.IsBetween(v, Info, Fatal)
}

func (v Variant) IsInvalid() bool {
	return baseenumer.IsNotBetween(v, Info, Fatal)
}

func (v Variant) IsEnum() bool {
	return v.IsValid()
}

func (v Variant) IsUnknown() bool {
	return v == Unknown
}

func (v Variant) IsInfo() bool {
	return v == Info
}

func (v Variant) IsWarn() bool {
	return v == Warn
}

func (v Variant) IsError() bool {
	return v == Error
}

func (v Variant) IsCritical() bool {
	return v == Critical
}

func (v Variant) IsFatal() bool {
	return v == Fatal
}

func (v Variant) Min() Variant {
	return basicEnum.Min()
}

func (v Variant) Max() Variant {
	return basicEnum.Max()
}

func (v Variant) IsMin() bool {
	return basicEnum.IsMin(v)
}

func (v Variant) IsMax() bool {
	return basicEnum.IsMax(v)
}

func (v Variant) IsInRange(min, max Variant) bool {
	return baseenumer.IsBetween(v, min, max)
}

func (v Variant) Name() string {
	if int(v) < len(variantLabels) {
		return variantLabels[v]
	}

	return fmt.Sprintf("Severity(%d)", byte(v))
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
