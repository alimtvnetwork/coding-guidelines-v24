package streamwriter

import "fmt"

// PayloadKind identifies the classification of an incoming generic payload.
type PayloadKind byte

const (
	PayloadNil PayloadKind = iota
	PayloadBytes
	PayloadString
	PayloadError
	PayloadMap
	PayloadStruct
	PayloadPrimitive
)

var payloadKindNames = [...]string{
	"Nil",
	"Bytes",
	"String",
	"Error",
	"Map",
	"Struct",
	"Primitive",
}

// Name returns the identifier name for PayloadKind.
func (k PayloadKind) Name() string {
	idx := int(k)
	if idx < len(payloadKindNames) {
		return payloadKindNames[idx]
	}

	return fmt.Sprintf("PayloadKind(%d)", idx)
}

// String implements fmt.Stringer for PayloadKind.
func (k PayloadKind) String() string {
	return k.Name()
}

// IsValid reports whether this is a known payload kind.
func (k PayloadKind) IsValid() bool {
	return int(k) < len(payloadKindNames)
}
