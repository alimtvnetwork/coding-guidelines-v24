package appfault

import "coding-guidelines/common/pkg/enum/severitytype"

// SeverityType represents an integer-backed severity level (byte).
type SeverityType = severitytype.Variant

// Severity level constants.
const (
	SeverityUnknown  = severitytype.Unknown
	SeverityInfo     = severitytype.Info
	SeverityWarn     = severitytype.Warn
	SeverityError    = severitytype.Error
	SeverityCritical = severitytype.Critical
	SeverityFatal    = severitytype.Fatal
)

// ParseSeverityName looks up a SeverityType by name.
func ParseSeverityName(str string) (SeverityType, bool) {
	return severitytype.Parse(str)
}

func parseSeverityName(str string) (SeverityType, bool) {
	return severitytype.Parse(str)
}
