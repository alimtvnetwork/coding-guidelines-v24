package appfault

import "coding-guidelines/common/pkg/enum/prioritytype"

// PriorityType represents an integer-backed priority level (byte).
type PriorityType = prioritytype.Variant

// ParsePriorityName looks up a PriorityType by name.
func ParsePriorityName(str string) (PriorityType, bool) {
	return prioritytype.Parse(str)
}

func parsePriorityName(str string) (PriorityType, bool) {
	return prioritytype.Parse(str)
}
