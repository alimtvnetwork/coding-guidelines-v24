package logleveltype

import (
	"encoding/json"
	"fmt"
)

// Variant represents the log severity level enum backed by byte.
type Variant byte

// Name returns the PascalCase label.
func (v Variant) Name() string {
	if int(v) < len(variantLabels) {
		return variantLabels[v]
	}

	return fmt.Sprintf("LogLevel(%d)", byte(v))
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

// IsEnabled returns true if the current level meets or exceeds the target threshold.
func (v Variant) IsEnabled(threshold Variant) bool {
	return v >= threshold
}

// IsDebug checks if level is Debug.
func (v Variant) IsDebug() bool {
	return v == Debug
}

// IsInfo checks if level is Info.
func (v Variant) IsInfo() bool {
	return v == Info
}

// IsWarn checks if level is Warn.
func (v Variant) IsWarn() bool {
	return v == Warn
}

// IsError checks if level is Error.
func (v Variant) IsError() bool {
	return v == Error
}

// IsFatal checks if level is Fatal.
func (v Variant) IsFatal() bool {
	return v == Fatal
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
