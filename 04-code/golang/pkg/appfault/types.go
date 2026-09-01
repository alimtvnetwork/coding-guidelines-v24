package appfault

import "coding-guidelines/common/pkg/errtype"

// Variation is an alias for errtype.Variation for convenience.
type Variation = errtype.Variation

// Standard severity level constants.
const (
	SeverityInfo     = "INFO"
	SeverityWarn     = "WARN"
	SeverityError    = "ERROR"
	SeverityCritical = "CRITICAL"
	SeverityFatal    = "FATAL"
)

// Standard priority level constants.
const (
	PriorityLow      = "LOW"
	PriorityNormal   = "NORMAL"
	PriorityHigh     = "HIGH"
	PriorityCritical = "CRITICAL"
)
