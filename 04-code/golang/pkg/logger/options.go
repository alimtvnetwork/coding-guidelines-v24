package logger

import (
	"io"
	"os"
)

// LoggerOptions configures the behavior and formatting of Logger instances.
type LoggerOptions struct {
	Level           LogLevel
	IsJson          bool
	IsStackTrace    bool
	Output          io.Writer
	CallerSkipDepth int
}

// DefaultOptions returns standard sensible defaults for logging.
func DefaultOptions() LoggerOptions {
	return LoggerOptions{
		Level:           LevelInfo,
		IsJson:          false,
		IsStackTrace:    true,
		Output:          os.Stdout,
		CallerSkipDepth: 2,
	}
}

// WithLevel sets the minimum log level filter.
func (o LoggerOptions) WithLevel(level LogLevel) LoggerOptions {
	o.Level = level

	return o
}

// WithJson enables or disables JSON output formatting.
func (o LoggerOptions) WithJson(isEnabled bool) LoggerOptions {
	o.IsJson = isEnabled

	return o
}

// WithStackTrace enables or disables stack trace inclusion on errors.
func (o LoggerOptions) WithStackTrace(isEnabled bool) LoggerOptions {
	o.IsStackTrace = isEnabled

	return o
}

// WithOutput sets the destination writer for log outputs.
func (o LoggerOptions) WithOutput(out io.Writer) LoggerOptions {
	o.Output = out

	return o
}
