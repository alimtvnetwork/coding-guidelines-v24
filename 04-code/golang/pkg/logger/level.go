package logger

import (
	"coding-guidelines/common/pkg/enum/logleveltype"
)

// LogLevel defines the byte-backed severity ranking for structured log messages.
// Aliased to logleveltype.Variant for modular enum architecture.
type LogLevel = logleveltype.Variant

// ParseLogLevel parses a string into a LogLevel.
func ParseLogLevel(s string) LogLevel {
	res := logleveltype.Parse(s)
	if res.IsSuccess() {
		return res.Data()
	}

	return LevelUnknown
}
