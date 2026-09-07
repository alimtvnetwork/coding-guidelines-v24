package appfault

import "coding-guidelines/common/pkg/enum/severitytype"

// SeverityType represents an integer-backed severity level (byte).
type SeverityType = severitytype.Variant

// ParseSeverityName looks up a SeverityType by name.
func ParseSeverityName(str string) (SeverityType, bool) {
	return severitytype.Parse(str)
}

func parseSeverityName(str string) (SeverityType, bool) {
	return severitytype.Parse(str)
}
