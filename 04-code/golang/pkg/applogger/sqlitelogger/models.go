package sqlitelogger

import "database/sql"

// DBOpenerFunc defines the pluggable database connection factory.
type DBOpenerFunc func(dsn string) (*sql.DB, error)

// TaskLogEntry represents a structured log entry within SQLite databases.
type TaskLogEntry struct {
	Id         int64  `json:"id"`
	TaskId     string `json:"taskId"`
	Timestamp  string `json:"timestamp"`
	Level      string `json:"level"`
	Message    string `json:"message"`
	Caller     string `json:"caller"`
	FieldsJson string `json:"fieldsJson"`
	StackTrace string `json:"stackTrace"`
	DurationMs int64  `json:"durationMs"`
	Status     string `json:"status"`
}

// FilterOptions specifies query constraints for log retrieval.
type FilterOptions struct {
	Level     string
	Limit     int
	Offset    int
	StartTime string
	EndTime   string
}

// TaskSummary contains diagnostic metadata for a task database.
type TaskSummary struct {
	TaskId        string `json:"taskId"`
	DbPath        string `json:"dbPath"`
	TotalLogs     int    `json:"totalLogs"`
	ErrorCount    int    `json:"errorCount"`
	LastTimestamp string `json:"lastTimestamp"`
}
