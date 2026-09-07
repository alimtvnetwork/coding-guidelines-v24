package appfault

import "coding-guidelines/common/pkg/enum/prioritytype"

// PriorityType represents an integer-backed priority level (byte).
type PriorityType = prioritytype.Variant

// Priority level constants.
const (
	PriorityUnknown  = prioritytype.Unknown
	PriorityLow      = prioritytype.Low
	PriorityNormal   = prioritytype.Normal
	PriorityHigh     = prioritytype.High
	PriorityCritical = prioritytype.Critical
)

// ParsePriorityName looks up a PriorityType by name.
func ParsePriorityName(str string) (PriorityType, bool) {
	return prioritytype.Parse(str)
}

func parsePriorityName(str string) (PriorityType, bool) {
	return prioritytype.Parse(str)
}
