package logger

import (
	"coding-guidelines/common/pkg/enum/logleveltype"
)

// LogLevel enum constants aliased from logleveltype.
const (
	LevelUnknown LogLevel = logleveltype.Unknown
	LevelInvalid LogLevel = logleveltype.Invalid
	LevelDebug   LogLevel = logleveltype.Debug
	LevelInfo    LogLevel = logleveltype.Info
	LevelWarn    LogLevel = logleveltype.Warn
	LevelError   LogLevel = logleveltype.Error
	LevelFatal   LogLevel = logleveltype.Fatal
)

// LogFormatterFunc formats a LogEntry into a serialized string representation.
type LogFormatterFunc func(entry LogEntry) string

// LogFilterFunc tests whether a LogEntry should be emitted or suppressed.
type LogFilterFunc func(entry LogEntry) bool
