package processstatetype

import (
	"coding-guidelines/common/pkg/baseenumer"
)

type (
	Variant byte

	VariantPredicate func(v Variant) bool
)

const (
	Invalid Variant = iota
	Pending
	Running
	Completed
	Failed
	Cancelled
)

const Unknown = Invalid

func (v Variant) Name() string {
	if int(v) < len(variantLabels) {
		return variantLabels[v]
	}

	return baseenumer.FormatNameValue("ProcessState", byte(v))
}

func (v Variant) Label() string {
	return v.Name()
}

func (v Variant) String() string {
	return v.Name()
}

func (v Variant) IsValid() bool {
	return baseenumer.IsBetween(v, Pending, Cancelled)
}

func (v Variant) IsInvalid() bool {
	return baseenumer.IsNotBetween(v, Pending, Cancelled)
}

func (v Variant) IsPending() bool {
	return v == Pending
}

func (v Variant) IsRunning() bool {
	return v == Running
}

func (v Variant) IsCompleted() bool {
	return v == Completed
}

func (v Variant) IsFailed() bool {
	return v == Failed
}

func (v Variant) IsCancelled() bool {
	return v == Cancelled
}

func (v Variant) MarshalJSON() ([]byte, error) {
	return baseenumer.MarshalJSON(v.Name())
}

func (v *Variant) UnmarshalJSON(data []byte) error {
	return baseenumer.UnmarshalIntegerJSON(data, v, "processstatetype", variantMap, len(variantLabels)-1, Invalid)
}
