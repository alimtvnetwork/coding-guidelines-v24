package sqlitelogger

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/fileutil"
)

const (
	defaultMainDbName = "logs.db"
	tasksDirName      = "tasks"
	createTableSql    = `CREATE TABLE IF NOT EXISTS logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_id TEXT,
		timestamp TEXT,
		level TEXT,
		message TEXT,
		caller TEXT,
		fields_json TEXT,
		stack_trace TEXT,
		duration_ms INTEGER,
		status TEXT
	);`
)

// SplitDBManager orchestrates primary and task-isolated SQLite log databases.
type SplitDBManager struct {
	lock       sync.RWMutex
	workDir    string
	mainDbPath string
	tasksDir   string
	mainDb     *sql.DB
	taskDbs    map[string]*sql.DB
	opener     DBOpenerFunc
}

// NewSplitDBManager instantiates a SplitDBManager for a given work directory.
func NewSplitDBManager(
	workDir string,
	opener DBOpenerFunc,
) (*SplitDBManager, *appfault.AppError) {
	if workDir == "" {
		return nil, appfault.New(errtype.Validation, "workDir cannot be empty")
	}

	mgr := &SplitDBManager{
		workDir:    workDir,
		mainDbPath: filepath.Join(workDir, defaultMainDbName),
		tasksDir:   filepath.Join(workDir, tasksDirName),
		taskDbs:    make(map[string]*sql.DB),
		opener:     opener,
	}

	return mgr, mgr.Init()
}

// Init creates necessary directories and ensures database files are accessible.
func (m *SplitDBManager) Init() *appfault.AppError {
	m.lock.Lock()
	defer m.lock.Unlock()

	dirRes := fileutil.EnsureDir(m.tasksDir, filepermtype.Standard)
	if dirRes.IsFailed() {
		return dirRes.Fault()
	}

	return nil
}

// SetMainDbPath overrides the default main database location.
func (m *SplitDBManager) SetMainDbPath(customPath string) *appfault.AppError {
	m.lock.Lock()
	defer m.lock.Unlock()

	if customPath == "" {
		return appfault.New(errtype.Validation, "customPath cannot be empty")
	}

	m.mainDbPath = customPath

	return nil
}

// GetMainDb lazily opens and returns the global logs database connection.
func (m *SplitDBManager) GetMainDb() (*sql.DB, *appfault.AppError) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.mainDb != nil {
		return m.mainDb, nil
	}

	db, fault := m.openDbInternal(m.mainDbPath)
	if fault != nil {
		return nil, fault
	}

	m.mainDb = db

	return m.mainDb, nil
}

// GetTaskDb lazily opens and returns an isolated task database connection.
func (m *SplitDBManager) GetTaskDb(taskId string) (*sql.DB, *appfault.AppError) {
	if taskId == "" {
		return nil, appfault.New(errtype.Validation, "taskId cannot be empty")
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	if existing, hasDb := m.taskDbs[taskId]; hasDb {
		return existing, nil
	}

	taskPath := filepath.Join(m.tasksDir, fmt.Sprintf("%s.db", taskId))
	db, fault := m.openDbInternal(taskPath)
	if fault != nil {
		return nil, fault
	}

	m.taskDbs[taskId] = db

	return db, nil
}

// openDbInternal opens a database connection and creates tables if needed.
func (m *SplitDBManager) openDbInternal(dbPath string) (*sql.DB, *appfault.AppError) {
	if m.opener == nil {
		return nil, appfault.New(errtype.Internal, "database opener func is not registered")
	}

	db, err := m.opener(dbPath)
	if err != nil {
		return nil, appfault.Wrap(errtype.Database, err, "failed to open database at "+dbPath)
	}

	_, execErr := db.Exec(createTableSql)
	if execErr != nil {
		_ = db.Close()

		return nil, appfault.Wrap(errtype.Database, execErr, "failed to initialize tables at "+dbPath)
	}

	return db, nil
}

// WriteMain persists a log entry to the primary logs database.
func (m *SplitDBManager) WriteMain(entry TaskLogEntry) *appfault.AppError {
	db, fault := m.GetMainDb()
	if fault != nil {
		return fault
	}

	return m.insertLogEntry(db, entry)
}

// WriteTask persists a log entry to the task's dedicated SQLite database.
func (m *SplitDBManager) WriteTask(
	taskId string,
	entry TaskLogEntry,
) *appfault.AppError {
	entry.TaskId = taskId
	db, fault := m.GetTaskDb(taskId)
	if fault != nil {
		return fault
	}

	return m.insertLogEntry(db, entry)
}

// insertLogEntry inserts a row into the logs table.
func (m *SplitDBManager) insertLogEntry(
	db *sql.DB,
	e TaskLogEntry,
) *appfault.AppError {
	stmt := `INSERT INTO logs (task_id, timestamp, level, message, caller, fields_json, stack_trace, duration_ms, status)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(
		stmt,
		e.TaskId,
		e.Timestamp,
		e.Level,
		e.Message,
		e.Caller,
		e.FieldsJson,
		e.StackTrace,
		e.DurationMs,
		e.Status,
	)
	if err != nil {
		return appfault.Wrap(errtype.Database, err, "failed to insert log entry")
	}

	return nil
}

// ListTaskDBs returns a slice of all active task IDs discovered in tasksDir.
func (m *SplitDBManager) ListTaskDBs() ([]string, *appfault.AppError) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	entries, err := os.ReadDir(m.tasksDir)
	if err != nil {
		return nil, appfault.Wrap(errtype.IO, err, "failed to read tasks directory")
	}

	taskIds := make([]string, 0, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasSuffix(name, ".db") {
			taskIds = append(taskIds, strings.TrimSuffix(name, ".db"))
		}
	}

	return taskIds, nil
}

// QueryTaskLogs executes filtered queries against a task database.
func (m *SplitDBManager) QueryTaskLogs(
	taskId string,
	filter FilterOptions,
) ([]TaskLogEntry, *appfault.AppError) {
	db, fault := m.GetTaskDb(taskId)
	if fault != nil {
		return nil, fault
	}

	return m.queryLogs(db, filter)
}

// QueryMainLogs executes filtered queries against the primary logs database.
func (m *SplitDBManager) QueryMainLogs(filter FilterOptions) ([]TaskLogEntry, *appfault.AppError) {
	db, fault := m.GetMainDb()
	if fault != nil {
		return nil, fault
	}

	return m.queryLogs(db, filter)
}

// queryLogs performs row scanning for log queries.
func (m *SplitDBManager) queryLogs(
	db *sql.DB,
	filter FilterOptions,
) ([]TaskLogEntry, *appfault.AppError) {
	limit := filter.Limit
	if limit <= 0 {
		limit = 100
	}

	query := "SELECT id, task_id, timestamp, level, message, caller, fields_json, stack_trace, duration_ms, status FROM logs ORDER BY id DESC LIMIT ?"
	rows, err := db.Query(query, limit)
	if err != nil {
		return nil, appfault.Wrap(errtype.Database, err, "failed to execute log query")
	}

	defer rows.Close()

	return scanLogRows(rows)
}

// scanLogRows parses rows into TaskLogEntry slice.
func scanLogRows(rows *sql.Rows) ([]TaskLogEntry, *appfault.AppError) {
	var results []TaskLogEntry
	for rows.Next() {
		var e TaskLogEntry
		err := rows.Scan(
			&e.Id,
			&e.TaskId,
			&e.Timestamp,
			&e.Level,
			&e.Message,
			&e.Caller,
			&e.FieldsJson,
			&e.StackTrace,
			&e.DurationMs,
			&e.Status,
		)
		if err != nil {
			return nil, appfault.Wrap(errtype.Database, err, "failed to scan row")
		}

		results = append(results, e)
	}

	return results, nil
}

// GetTaskSummary produces count and error statistics for a task database.
func (m *SplitDBManager) GetTaskSummary(taskId string) (*TaskSummary, *appfault.AppError) {
	db, fault := m.GetTaskDb(taskId)
	if fault != nil {
		return nil, fault
	}

	summary := &TaskSummary{
		TaskId: taskId,
		DbPath: filepath.Join(m.tasksDir, fmt.Sprintf("%s.db", taskId)),
	}

	row := db.QueryRow("SELECT COUNT(*), COALESCE(MAX(timestamp), '') FROM logs")
	if err := row.Scan(&summary.TotalLogs, &summary.LastTimestamp); err != nil {
		return nil, appfault.Wrap(errtype.Database, err, "failed to scan summary")
	}

	errRow := db.QueryRow("SELECT COUNT(*) FROM logs WHERE level IN ('ERROR', 'FATAL')")
	if err := errRow.Scan(&summary.ErrorCount); err != nil {
		return nil, appfault.Wrap(errtype.Database, err, "failed to scan error count")
	}

	return summary, nil
}

// Close closes all open database connections.
func (m *SplitDBManager) Close() *appfault.AppError {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.mainDb != nil {
		_ = m.mainDb.Close()
		m.mainDb = nil
	}

	for id, db := range m.taskDbs {
		_ = db.Close()
		delete(m.taskDbs, id)
	}

	return nil
}
