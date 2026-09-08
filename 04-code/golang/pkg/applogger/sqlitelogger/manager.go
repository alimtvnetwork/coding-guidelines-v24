package sqlitelogger

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

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
	lock                  sync.RWMutex
	workDir               string
	mainDbPath            string
	tasksDir              string
	resolver              TaskDbPathResolverFunc
	customPaths           map[string]string
	mainDb                *sql.DB
	taskDbs               map[string]*sql.DB
	opener                DBOpenerFunc
	isManualMigrationOnly bool
}

// NewSplitDBManager instantiates a SplitDBManager for a given work directory.
func NewSplitDBManager(
	workDir string,
	opener DBOpenerFunc,
) (*SplitDBManager, *appfault.AppError) {
	return NewSplitDBManagerWithConfig(SplitDBConfig{
		WorkDir: workDir,
		Opener:  opener,
	})
}

// NewSplitDBManagerWithConfig instantiates a SplitDBManager with configurable paths.
func NewSplitDBManagerWithConfig(cfg SplitDBConfig) (*SplitDBManager, *appfault.AppError) {
	if cfg.WorkDir == "" {
		return nil, appfault.New(errtype.Validation, "WorkDir cannot be empty")
	}

	mainDbPath := cfg.MainDbPath
	if mainDbPath == "" {
		mainDbPath = filepath.Join(cfg.WorkDir, defaultMainDbName)
	}

	tasksDir := cfg.TasksDir
	if tasksDir == "" {
		tasksDir = filepath.Join(cfg.WorkDir, tasksDirName)
	}

	mgr := &SplitDBManager{
		workDir:               cfg.WorkDir,
		mainDbPath:            mainDbPath,
		tasksDir:              tasksDir,
		resolver:              cfg.TaskDbPathResolver,
		customPaths:           make(map[string]string),
		taskDbs:               make(map[string]*sql.DB),
		opener:                cfg.Opener,
		isManualMigrationOnly: cfg.IsManualMigrationOnly,
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

// WorkDir returns the configured base directory for logs.
func (m *SplitDBManager) WorkDir() string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.workDir
}

// TasksDir returns the active directory holding task SQLite databases.
func (m *SplitDBManager) TasksDir() string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.tasksDir
}

// MainDbPath returns the absolute or relative path to the main logs database.
func (m *SplitDBManager) MainDbPath() string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.mainDbPath
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

// SetTasksDir overrides the directory holding task SQLite databases.
func (m *SplitDBManager) SetTasksDir(customDir string) *appfault.AppError {
	m.lock.Lock()
	defer m.lock.Unlock()

	if customDir == "" {
		return appfault.New(errtype.Validation, "customDir cannot be empty")
	}

	m.tasksDir = customDir
	dirRes := fileutil.EnsureDir(customDir, filepermtype.Standard)
	if dirRes.IsFailed() {
		return dirRes.Fault()
	}

	return nil
}

// SetTaskDbPath assigns an explicit custom database file path for a specific task.
func (m *SplitDBManager) SetTaskDbPath(taskId, customPath string) *appfault.AppError {
	if taskId == "" || customPath == "" {
		return appfault.New(errtype.Validation, "taskId and customPath cannot be empty")
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	m.customPaths[taskId] = customPath

	return nil
}

// SetTaskDbPathResolver configures a custom naming algorithm for task databases.
func (m *SplitDBManager) SetTaskDbPathResolver(resolver TaskDbPathResolverFunc) *appfault.AppError {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.resolver = resolver

	return nil
}

// ResolveTaskDbPath computes the target database path for a given task ID.
func (m *SplitDBManager) ResolveTaskDbPath(taskId string) string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.resolveTaskDbPathUnsafe(taskId)
}

// resolveTaskDbPathUnsafe resolves path without acquiring mutex (caller must hold lock).
func (m *SplitDBManager) resolveTaskDbPathUnsafe(taskId string) string {
	if custom, hasCustom := m.customPaths[taskId]; hasCustom && custom != "" {
		return custom
	}

	if m.resolver != nil {
		return m.resolver(m.tasksDir, taskId)
	}

	return filepath.Join(m.tasksDir, fmt.Sprintf("%s.db", taskId))
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

	taskPath := m.resolveTaskDbPathUnsafe(taskId)
	_ = fileutil.EnsureDir(filepath.Dir(taskPath), filepermtype.Standard)

	db, fault := m.openDbInternal(taskPath)
	if fault != nil {
		return nil, fault
	}

	m.taskDbs[taskId] = db

	return db, nil
}

// openConnection safely invokes the opener function and wraps errors.
func openConnection(opener DBOpenerFunc, dbPath string) (*sql.DB, *appfault.AppError) {
	if opener == nil {
		return nil, appfault.New(errtype.Internal, "database opener func is not registered")
	}

	db, err := opener(dbPath)
	if err != nil {
		return nil, appfault.Wrap(errtype.Database, err, "failed to open database at "+dbPath)
	}

	return db, nil
}

// applyAutoMigration executes schema repairs and migrations if not configured for manual only.
func applyAutoMigration(db *sql.DB, isManual bool) *appfault.AppError {
	isAuto := !isManual
	if isAuto {
		return MigrateAndRepairDatabase(db)
	}

	return nil
}

// openDbInternal opens a database connection and ensures tables, migrations, and repairs.
func (m *SplitDBManager) openDbInternal(dbPath string) (*sql.DB, *appfault.AppError) {
	db, fault := openConnection(m.opener, dbPath)
	if fault != nil {
		return nil, fault
	}

	if aFault := applyAutoMigration(db, m.isManualMigrationOnly); aFault != nil {
		_ = db.Close()

		return nil, aFault
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

// ListTaskDBs returns a slice of all active task IDs discovered in tasksDir and custom paths.
func (m *SplitDBManager) ListTaskDBs() ([]string, *appfault.AppError) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	taskIdsMap := make(map[string]bool)
	for id := range m.customPaths {
		taskIdsMap[id] = true
	}

	entries, err := os.ReadDir(m.tasksDir)
	if err == nil {
		for _, entry := range entries {
			name := entry.Name()
			if strings.HasSuffix(name, ".db") {
				taskIdsMap[strings.TrimSuffix(name, ".db")] = true
			}
		}
	}

	taskIds := make([]string, 0, len(taskIdsMap))
	for id := range taskIdsMap {
		taskIds = append(taskIds, id)
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

// queryLogs performs row scanning for log queries with dynamic filtering.
func (m *SplitDBManager) queryLogs(
	db *sql.DB,
	filter FilterOptions,
) ([]TaskLogEntry, *appfault.AppError) {
	query, args := buildLogsQuery(filter)
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, appfault.Wrap(errtype.Database, err, "failed to execute log query")
	}

	defer rows.Close()

	return scanLogRows(rows)
}

// buildLogsQuery constructs dynamic SQL and parameter bindings from filter options.
func buildLogsQuery(f FilterOptions) (string, []any) {
	base := "SELECT id, task_id, timestamp, level, message, caller, fields_json, stack_trace, duration_ms, status FROM logs WHERE 1=1"
	query, args := appendFilterConditions(base, f)
	limit := normalizeLimit(f.Limit)
	query += " ORDER BY id DESC LIMIT ? OFFSET ?"
	args = append(args, limit, f.Offset)

	return query, args
}

// appendFilterConditions appends WHERE clauses for level, startTime, and endTime.
func appendFilterConditions(base string, f FilterOptions) (string, []any) {
	q := base
	var args []any

	if f.Level != "" {
		q += " AND level = ?"
		args = append(args, f.Level)
	}

	if f.StartTime != "" {
		q += " AND timestamp >= ?"
		args = append(args, f.StartTime)
	}

	if f.EndTime != "" {
		q += " AND timestamp <= ?"
		args = append(args, f.EndTime)
	}

	return q, args
}

// normalizeLimit ensures query limit is positive with a standard default of 100.
func normalizeLimit(limit int) int {
	if limit <= 0 {
		return 100
	}

	return limit
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
		DbPath: m.ResolveTaskDbPath(taskId),
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

// IsManualMigrationOnly returns true if automatic migration on DB open is disabled.
func (m *SplitDBManager) IsManualMigrationOnly() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.isManualMigrationOnly
}

// SetManualMigrationOnly sets whether database auto-migration on open is bypassed.
func (m *SplitDBManager) SetManualMigrationOnly(isManual bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.isManualMigrationOnly = isManual
}

// MigrateTaskDb runs schema migrations and repairs on a specific task database.
func (m *SplitDBManager) MigrateTaskDb(taskId string) *appfault.AppError {
	db, fault := m.GetTaskDb(taskId)
	if fault != nil {
		return fault
	}

	return MigrateAndRepairDatabase(db)
}

// RepairTaskDb executes integrity verification and column repair on a task database.
func (m *SplitDBManager) RepairTaskDb(taskId string) *appfault.AppError {
	db, fault := m.GetTaskDb(taskId)
	if fault != nil {
		return fault
	}

	return RepairDatabase(db)
}

// MigrateAllTaskDbs discovers and runs migrations on all available task databases.
func (m *SplitDBManager) MigrateAllTaskDbs() *appfault.AppError {
	taskIds, fault := m.ListTaskDBs()
	if fault != nil {
		return fault
	}

	for _, id := range taskIds {
		if mFault := m.MigrateTaskDb(id); mFault != nil {
			return mFault
		}
	}

	return nil
}

// RepairAllTaskDbs discovers and runs repairs on all available task databases.
func (m *SplitDBManager) RepairAllTaskDbs() *appfault.AppError {
	taskIds, fault := m.ListTaskDBs()
	if fault != nil {
		return fault
	}

	for _, id := range taskIds {
		if rFault := m.RepairTaskDb(id); rFault != nil {
			return rFault
		}
	}

	return nil
}

// MigrateMainDb executes migrations on the global logs database.
func (m *SplitDBManager) MigrateMainDb() *appfault.AppError {
	db, fault := m.GetMainDb()
	if fault != nil {
		return fault
	}

	return MigrateAndRepairDatabase(db)
}

// RepairMainDb executes integrity checks and repairs on the global logs database.
func (m *SplitDBManager) RepairMainDb() *appfault.AppError {
	db, fault := m.GetMainDb()
	if fault != nil {
		return fault
	}

	return RepairDatabase(db)
}

// PruneTasks removes task databases whose files are older than maxAge.
func (m *SplitDBManager) PruneTasks(maxAge time.Duration) (int, *appfault.AppError) {
	m.lock.Lock()
	defer m.lock.Unlock()

	entries, err := os.ReadDir(m.tasksDir)
	if err != nil {
		return 0, appfault.Wrap(errtype.IO, err, "failed to read tasks dir")
	}

	cutoff := time.Now().Add(-maxAge)

	return m.pruneExpiredEntries(entries, cutoff)
}

// pruneExpiredEntries inspects directory entries and deletes files older than cutoff.
func (m *SplitDBManager) pruneExpiredEntries(entries []os.DirEntry, cutoff time.Time) (int, *appfault.AppError) {
	prunedCount := 0
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".db") {
			continue
		}

		info, err := entry.Info()
		if err == nil && info.ModTime().Before(cutoff) {
			taskId := strings.TrimSuffix(entry.Name(), ".db")
			m.removeTaskDatabaseFiles(taskId)
			prunedCount++
		}
	}

	return prunedCount, nil
}

// removeTaskDatabaseFiles closes connections and removes primary and temporary SQLite files.
func (m *SplitDBManager) removeTaskDatabaseFiles(taskId string) {
	if db, exists := m.taskDbs[taskId]; exists {
		_ = db.Close()
		delete(m.taskDbs, taskId)
	}

	basePath := m.resolveTaskDbPathUnsafe(taskId)
	_ = os.Remove(basePath)
	_ = os.Remove(basePath + "-wal")
	_ = os.Remove(basePath + "-shm")
}

// PruneTaskCount enforces an upper bound on task databases, deleting oldest first.
func (m *SplitDBManager) PruneTaskCount(maxDbs int) (int, *appfault.AppError) {
	m.lock.Lock()
	defer m.lock.Unlock()

	files, fault := m.collectTaskFilesByAge()
	if fault != nil {
		return 0, fault
	}

	excessCount := len(files) - maxDbs
	if excessCount <= 0 {
		return 0, nil
	}

	return m.pruneExcessFiles(files[:excessCount]), nil
}

type taskFileInfo struct {
	taskId  string
	modTime time.Time
}

// collectTaskFilesByAge discovers and sorts task files by modification timestamp.
func (m *SplitDBManager) collectTaskFilesByAge() ([]taskFileInfo, *appfault.AppError) {
	entries, err := os.ReadDir(m.tasksDir)
	if err != nil {
		return nil, appfault.Wrap(errtype.IO, err, "failed to read tasks dir")
	}

	var files []taskFileInfo
	for _, e := range entries {
		if item, ok := extractTaskFileInfo(e); ok {
			files = append(files, item)
		}
	}

	sort.Slice(files, func(i, j int) bool { return files[i].modTime.Before(files[j].modTime) })

	return files, nil
}

// extractTaskFileInfo parses metadata from a database directory entry.
func extractTaskFileInfo(e os.DirEntry) (taskFileInfo, bool) {
	if !strings.HasSuffix(e.Name(), ".db") {
		return taskFileInfo{}, false
	}

	info, err := e.Info()
	if err != nil {
		return taskFileInfo{}, false
	}

	return taskFileInfo{
		taskId:  strings.TrimSuffix(e.Name(), ".db"),
		modTime: info.ModTime(),
	}, true
}

// pruneExcessFiles deletes the specified excess task files and cleans up memory state.
func (m *SplitDBManager) pruneExcessFiles(excess []taskFileInfo) int {
	for _, f := range excess {
		m.removeTaskDatabaseFiles(f.taskId)
	}

	return len(excess)
}
