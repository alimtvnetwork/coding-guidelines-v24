package processstatetype

import (
	"encoding/json"
	"fmt"
	"strings"
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

	return fmt.Sprintf("ProcessState(%d)", byte(v))
}

func (v Variant) Label() string {
	return v.Name()
}

func (v Variant) String() string {
	return v.Name()
}

func (v Variant) IsValid() bool {
	return v > Invalid && int(v) < len(variantLabels)
}

func (v Variant) IsInvalid() bool {
	return v <= Invalid || int(v) >= len(variantLabels)
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
		return fmt.Errorf("invalid processstatetype numeric value %d, supported range: 0..%d", raw, len(variantLabels)-1)
	}

	*v = Variant(raw)

	return nil
}
