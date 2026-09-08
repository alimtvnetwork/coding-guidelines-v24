package logleveltype

import (
	"coding-guidelines/common/pkg/baseenumer"
)

type (
	Variant byte

	VariantPredicate func(v Variant) bool
)

const (
	Invalid Variant = iota
	Debug
	Info
	Warn
	Error
	Fatal
)

const Unknown = Invalid

func (v Variant) Name() string {
	if int(v) < len(variantLabels) {
		return variantLabels[v]
	}

	return baseenumer.FormatNameValue("LogLevel", byte(v))
}

func (v Variant) Label() string {
	return v.Name()
}

func (v Variant) String() string {
	return v.Name()
}

func (v Variant) IsValid() bool {
	return baseenumer.IsBetween(v, Debug, Fatal)
}

func (v Variant) IsInvalid() bool {
	return baseenumer.IsNotBetween(v, Debug, Fatal)
}

func (v Variant) IsEnabled(threshold Variant) bool {
	return v >= threshold
}

func (v Variant) IsDebug() bool {
	return v == Debug
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

func (v Variant) IsFatal() bool {
	return v == Fatal
}

func (v Variant) MarshalJSON() ([]byte, error) {
	return baseenumer.MarshalJSON(v.Name())
}

func (v *Variant) UnmarshalJSON(data []byte) error {
	return basicEnum.UnmarshalJSON(data, v)
}
