package applogger

import "coding-guidelines/common/pkg/logger"

// LogLevel constants mirrored from package logger.
const (
	LevelUnknown = logger.LevelUnknown
	LevelDebug   = logger.LevelDebug
	LevelInfo    = logger.LevelInfo
	LevelWarn    = logger.LevelWarn
	LevelError   = logger.LevelError
	LevelFatal   = logger.LevelFatal
)

// DriverType selects the backend implementation.
type DriverType byte

// DriverType constants for logger backends.
const (
	DriverConsole DriverType = iota
	DriverFile
	DriverSQLite
	DriverZap
	DriverComposite
)

// SQL table schema for SQLite logging sink.
const createLogsTableSQL = `CREATE TABLE IF NOT EXISTS app_logs (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	timestamp TEXT NOT NULL,
	level TEXT NOT NULL,
	message TEXT NOT NULL,
	caller TEXT,
	fields_json TEXT,
	stack_trace TEXT
);`

// EntryFormatter formats a LogEntry into a string.
type EntryFormatter func(entry LogEntry) string

// EntryFilterFunc filters LogEntry records based on criteria.
type EntryFilterFunc func(entry LogEntry) bool
