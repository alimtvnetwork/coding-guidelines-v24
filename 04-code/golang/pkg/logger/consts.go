package logger

// LogLevel enum constants defining severity ranking for structured log messages.
const (
	LevelUnknown LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// LogFormatterFunc formats a LogEntry into a serialized string representation.
type LogFormatterFunc func(entry LogEntry) string

// LogFilterFunc tests whether a LogEntry should be emitted or suppressed.
type LogFilterFunc func(entry LogEntry) bool
