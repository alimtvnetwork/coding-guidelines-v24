package logger

// LogLevel defines the severity ranking for structured log messages.
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var levelNames = [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

// String returns the string representation of the log level.
func (l LogLevel) String() string {
	if int(l) >= 0 && int(l) < len(levelNames) {
		return levelNames[l]
	}

	return "UNKNOWN"
}

// IsEnabled returns true if the current level meets or exceeds the target threshold.
func (l LogLevel) IsEnabled(threshold LogLevel) bool {
	return l >= threshold
}
