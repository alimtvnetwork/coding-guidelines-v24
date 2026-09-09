package applogger

import (
	"database/sql"
	"encoding/json"
	"sync"

	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

// SQLiteSink writes structured log entries to an SQLite table.
type SQLiteSink struct {
	lock sync.Mutex
	db   *sql.DB
}

// NewSQLiteSink creates and initializes the SQLite logging table.
func NewSQLiteSink(db *sql.DB) SQLiteSinkResult {
	sink := &SQLiteSink{db: db}
	if err := sink.initTable(); err != nil {
		return result.WrapFailureWithCause[*SQLiteSink](errtype.Database, err, "failed to initialize SQLite logs table")
	}

	return SQLiteSinkSuccess(sink)
}

// initTable ensures the log table exists.
func (ss *SQLiteSink) initTable() error {
	if ss.db == nil {
		return nil
	}

	_, err := ss.db.Exec(createLogsTableSQL)

	return err
}

// WriteEntry persists a structured log entry to the SQLite database.
func (ss *SQLiteSink) WriteEntry(e LogEntry) error {
	if ss.db == nil {
		return nil
	}

	ss.lock.Lock()
	defer ss.lock.Unlock()

	fieldsJSON, err := json.Marshal(e.Fields)
	if err != nil {
		return err
	}

	query := `INSERT INTO app_logs (timestamp, level, message, caller, fields_json, stack_trace) VALUES (?, ?, ?, ?, ?, ?)`
	_, execErr := ss.db.Exec(query, e.Timestamp, e.Level.Name(), e.Message, e.Caller, string(fieldsJSON), e.Stack)

	return execErr
}

// Sync is a no-op for SQLite transactions.
func (ss *SQLiteSink) Sync() error { return nil }

// Close closes the underlying db connection if needed.
func (ss *SQLiteSink) Close() error { return nil }

// DriverType returns the driver type.
func (ss *SQLiteSink) DriverType() DriverType {
	return DriverSQLite
}
