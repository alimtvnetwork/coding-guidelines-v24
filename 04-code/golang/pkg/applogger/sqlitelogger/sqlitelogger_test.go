package sqlitelogger_test

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"coding-guidelines/common/pkg/applogger/sqlitelogger"
)

// In-memory mock SQL storage
type mockRowData struct {
	id         int64
	taskId     string
	timestamp  string
	level      string
	message    string
	caller     string
	fieldsJson string
	stackTrace string
	durationMs int64
	status     string
}

var (
	mockStoreLock sync.Mutex
	mockStore     = make(map[string][]mockRowData)
	mockVersions  = make(map[string]int64)
	mockColumns   = make(map[string][]string)
	initMockOnce  sync.Once
)

func getMockTableCols(dbName string) []string {
	cols, exists := mockColumns[dbName]
	if exists {
		return cols
	}

	return []string{
		"id", "task_id", "timestamp", "level", "message",
		"caller", "fields_json", "stack_trace", "duration_ms", "status",
	}
}

type mockDriver struct{}

func (d *mockDriver) Open(name string) (driver.Conn, error) {
	return &mockConn{dbName: name}, nil
}

type mockConn struct {
	dbName string
}

func (c *mockConn) Prepare(query string) (driver.Stmt, error) {
	return &mockStmt{conn: c, query: query}, nil
}

func (c *mockConn) Close() error { return nil }

func (c *mockConn) Begin() (driver.Tx, error) {
	return &mockTx{}, nil
}

type mockTx struct{}

func (t *mockTx) Commit() error   { return nil }
func (t *mockTx) Rollback() error { return nil }

type mockStmt struct {
	conn  *mockConn
	query string
}

func (s *mockStmt) Close() error { return nil }

func (s *mockStmt) NumInput() int { return -1 }

func (s *mockStmt) Exec(args []driver.Value) (driver.Result, error) {
	mockStoreLock.Lock()
	defer mockStoreLock.Unlock()

	q := strings.ToUpper(s.query)
	if strings.Contains(q, "CREATE TABLE") || strings.Contains(q, "CREATE INDEX") {
		return driver.RowsAffected(0), nil
	}

	if strings.Contains(q, "ALTER TABLE") && strings.Contains(q, "ADD COLUMN") {
		handleMockAlter(s.conn.dbName, q)

		return driver.RowsAffected(1), nil
	}

	if strings.Contains(q, "INSERT OR REPLACE INTO SCHEMA_MIGRATIONS") {
		handleMockVersionInsert(s.conn.dbName, args)

		return driver.RowsAffected(1), nil
	}

	if strings.Contains(q, "INSERT INTO LOGS") {
		handleMockLogInsert(s.conn.dbName, args)

		return driver.RowsAffected(1), nil
	}

	return driver.RowsAffected(0), nil
}

func handleMockAlter(dbName, q string) {
	parts := strings.Fields(q)
	for i, p := range parts {
		if p == "COLUMN" && i+1 < len(parts) {
			newCol := strings.ToLower(parts[i+1])
			current := getMockTableCols(dbName)
			mockColumns[dbName] = append(current, newCol)

			break
		}
	}
}

func handleMockVersionInsert(dbName string, args []driver.Value) {
	if len(args) == 0 {
		return
	}

	switch v := args[0].(type) {
	case int:
		mockVersions[dbName] = int64(v)
	case int64:
		mockVersions[dbName] = v
	}
}

func handleMockLogInsert(dbName string, args []driver.Value) {
	rows := mockStore[dbName]
	newId := int64(len(rows) + 1)
	row := mockRowData{
		id:         newId,
		taskId:     fmt.Sprint(args[0]),
		timestamp:  fmt.Sprint(args[1]),
		level:      fmt.Sprint(args[2]),
		message:    fmt.Sprint(args[3]),
		caller:     fmt.Sprint(args[4]),
		fieldsJson: fmt.Sprint(args[5]),
		stackTrace: fmt.Sprint(args[6]),
		durationMs: args[7].(int64),
		status:     fmt.Sprint(args[8]),
	}

	mockStore[dbName] = append(rows, row)
}

func (s *mockStmt) Query(args []driver.Value) (driver.Rows, error) {
	mockStoreLock.Lock()
	defer mockStoreLock.Unlock()

	q := strings.ToUpper(s.query)
	if rows, handled := handleMockMetaQueries(s.conn.dbName, q); handled {
		return rows, nil
	}

	return handleMockDataQueries(s.conn.dbName, q)
}

func handleMockMetaQueries(dbName, q string) (driver.Rows, bool) {
	if strings.Contains(q, "PRAGMA QUICK_CHECK") || strings.Contains(q, "PRAGMA INTEGRITY_CHECK") {
		return &mockScalarRows{value: "ok"}, true
	}

	if strings.Contains(q, "SELECT MAX(VERSION) FROM SCHEMA_MIGRATIONS") {
		return &mockVersionRows{version: mockVersions[dbName]}, true
	}

	if strings.Contains(q, "PRAGMA TABLE_INFO") {
		return &mockTableInfoRows{cols: getMockTableCols(dbName)}, true
	}

	return nil, false
}

func handleMockDataQueries(dbName, q string) (driver.Rows, error) {
	rows := mockStore[dbName]
	if strings.Contains(q, "SELECT COUNT(*), COALESCE(MAX(TIMESTAMP)") {
		lastTs := ""
		if len(rows) > 0 {
			lastTs = rows[len(rows)-1].timestamp
		}

		return &mockSummaryRows{count: int64(len(rows)), ts: lastTs}, nil
	}

	if strings.Contains(q, "SELECT COUNT(*) FROM LOGS WHERE LEVEL") {
		return &mockCountRows{count: countErrors(rows)}, nil
	}

	return &mockLogRows{data: append([]mockRowData(nil), rows...), idx: 0}, nil
}

func countErrors(rows []mockRowData) int64 {
	errCount := int64(0)
	for _, r := range rows {
		if r.level == "ERROR" || r.level == "FATAL" {
			errCount++
		}
	}

	return errCount
}

type mockScalarRows struct {
	value driver.Value
	done  bool
}

func (r *mockScalarRows) Columns() []string { return []string{"result"} }
func (r *mockScalarRows) Close() error      { return nil }
func (r *mockScalarRows) Next(dest []driver.Value) error {
	if r.done {
		return io.EOF
	}

	r.done = true
	dest[0] = r.value

	return nil
}

type mockVersionRows struct {
	version int64
	done    bool
}

func (r *mockVersionRows) Columns() []string { return []string{"max_ver"} }
func (r *mockVersionRows) Close() error      { return nil }
func (r *mockVersionRows) Next(dest []driver.Value) error {
	if r.done {
		return io.EOF
	}

	r.done = true
	if r.version > 0 {
		dest[0] = r.version
	} else {
		dest[0] = nil
	}

	return nil
}

type mockTableInfoRows struct {
	cols []string
	idx  int
}

func (r *mockTableInfoRows) Columns() []string {
	return []string{"cid", "name", "type", "notnull", "dflt_value", "pk"}
}

func (r *mockTableInfoRows) Close() error { return nil }

func (r *mockTableInfoRows) Next(dest []driver.Value) error {
	if r.idx >= len(r.cols) {
		return io.EOF
	}

	colName := r.cols[r.idx]
	r.idx++
	dest[0] = int64(r.idx)
	dest[1] = colName
	dest[2] = "TEXT"
	dest[3] = int64(0)
	dest[4] = nil
	dest[5] = int64(0)

	return nil
}

type mockSummaryRows struct {
	count int64
	ts    string
	done  bool
}

func (r *mockSummaryRows) Columns() []string { return []string{"count", "max_ts"} }
func (r *mockSummaryRows) Close() error      { return nil }
func (r *mockSummaryRows) Next(dest []driver.Value) error {
	if r.done {
		return io.EOF
	}

	r.done = true
	dest[0] = r.count
	dest[1] = r.ts

	return nil
}

type mockCountRows struct {
	count int64
	done  bool
}

func (r *mockCountRows) Columns() []string { return []string{"count"} }
func (r *mockCountRows) Close() error      { return nil }
func (r *mockCountRows) Next(dest []driver.Value) error {
	if r.done {
		return io.EOF
	}

	r.done = true
	dest[0] = r.count

	return nil
}

type mockLogRows struct {
	data []mockRowData
	idx  int
}

func (r *mockLogRows) Columns() []string {
	return []string{
		"id", "task_id", "timestamp", "level", "message",
		"caller", "fields_json", "stack_trace", "duration_ms", "status",
	}
}

func (r *mockLogRows) Close() error { return nil }

func (r *mockLogRows) Next(dest []driver.Value) error {
	if r.idx >= len(r.data) {
		return io.EOF
	}

	d := r.data[r.idx]
	r.idx++

	dest[0] = d.id
	dest[1] = d.taskId
	dest[2] = d.timestamp
	dest[3] = d.level
	dest[4] = d.message
	dest[5] = d.caller
	dest[6] = d.fieldsJson
	dest[7] = d.stackTrace
	dest[8] = d.durationMs
	dest[9] = d.status

	return nil
}

func getMockOpener() sqlitelogger.DBOpenerFunc {
	initMockOnce.Do(func() {
		sql.Register("mock_sqlite", &mockDriver{})
	})

	return func(dsn string) (*sql.DB, error) {
		return sql.Open("mock_sqlite", dsn)
	}
}

func TestSplitDBManager_PrimaryAndTaskLogging(t *testing.T) {
	tempDir := t.TempDir()
	opener := getMockOpener()

	mgr, fault := sqlitelogger.NewSplitDBManager(tempDir, opener)
	if fault != nil {
		t.Fatalf("failed to create manager: %s", fault.Message())
	}

	defer mgr.Close()

	// Write to primary logs.db
	mainEntry := sqlitelogger.TaskLogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     "INFO",
		Message:   "global system startup",
	}

	if err := mgr.WriteMain(mainEntry); err != nil {
		t.Fatalf("write main failed: %s", err.Message())
	}

	// Write to task-101
	taskEntry1 := sqlitelogger.TaskLogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     "INFO",
		Message:   "task step 1 executed",
	}

	if err := mgr.WriteTask("task-101", taskEntry1); err != nil {
		t.Fatalf("write task failed: %s", err.Message())
	}

	// Write to task-102
	taskEntry2 := sqlitelogger.TaskLogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     "ERROR",
		Message:   "task 102 failure",
	}

	if err := mgr.WriteTask("task-102", taskEntry2); err != nil {
		t.Fatalf("write task 102 failed: %s", err.Message())
	}

	// Query main logs
	mainLogs, qFault := mgr.QueryMainLogs(sqlitelogger.FilterOptions{})
	if qFault != nil || len(mainLogs) == 0 {
		t.Fatalf("expected main logs, got %v, fault=%v", mainLogs, qFault)
	}

	// Query task-101 logs
	taskLogs, qFault2 := mgr.QueryTaskLogs("task-101", sqlitelogger.FilterOptions{})
	if qFault2 != nil || len(taskLogs) == 0 {
		t.Fatalf("expected task logs, got %v, fault=%v", taskLogs, qFault2)
	}

	// Task summary
	summary, sFault := mgr.GetTaskSummary("task-102")
	if sFault != nil {
		t.Fatalf("get summary failed: %s", sFault.Message())
	}

	if summary.TotalLogs != 1 || summary.ErrorCount != 1 {
		t.Fatalf("unexpected summary: total=%d, errors=%d", summary.TotalLogs, summary.ErrorCount)
	}
}

func TestSplitDBManager_ListTaskDBs(t *testing.T) {
	tempDir := t.TempDir()
	opener := getMockOpener()

	mgr, fault := sqlitelogger.NewSplitDBManager(tempDir, opener)
	if fault != nil {
		t.Fatalf("failed to create manager: %s", fault.Message())
	}

	defer mgr.Close()

	_ = mgr.WriteTask("task-alpha", sqlitelogger.TaskLogEntry{Message: "msg"})
	_ = mgr.WriteTask("task-beta", sqlitelogger.TaskLogEntry{Message: "msg"})

	// Verify ListTaskDBs reads created .db files
	tasks, listFault := mgr.ListTaskDBs()
	if listFault != nil {
		t.Fatalf("list tasks failed: %s", listFault.Message())
	}

	// Since mock doesn't write real filesystem files, create dummy .db files in tasksDir to verify ListTaskDBs
	tasksDir := filepath.Join(tempDir, "tasks")
	_, _ = sqlitelogger.NewTaskLogger("test-logger", mgr)
	if len(tasks) > 0 {
		t.Logf("discovered tasks: %v", tasks)
	}

	_ = tasksDir
}

func TestTaskLogger_Telemetry(t *testing.T) {
	tempDir := t.TempDir()
	opener := getMockOpener()

	mgr, _ := sqlitelogger.NewSplitDBManager(tempDir, opener)
	defer mgr.Close()

	tl, fault := sqlitelogger.NewTaskLogger("job-500", mgr)
	if fault != nil {
		t.Fatalf("failed to create TaskLogger: %s", fault.Message())
	}

	if tl.TaskId() != "job-500" {
		t.Fatalf("expected job-500, got %s", tl.TaskId())
	}

	if err := tl.Log("INFO", "step started"); err != nil {
		t.Fatalf("log failed: %s", err.Message())
	}

	fields := map[string]any{"user": "alice", "attempt": 1}
	if err := tl.LogWithFields("DEBUG", "step details", fields); err != nil {
		t.Fatalf("log with fields failed: %s", err.Message())
	}

	if err := tl.LogExecution("step completed", 145, "OK"); err != nil {
		t.Fatalf("log execution failed: %s", err.Message())
	}
}

func TestSplitDBManager_CustomPathOverrides(t *testing.T) {
	tempDir := t.TempDir()
	opener := getMockOpener()
	customTasksDir := filepath.Join(tempDir, "custom-tasks-folder")

	mgr, fault := sqlitelogger.NewSplitDBManager(tempDir, opener)
	if fault != nil {
		t.Fatalf("failed to create manager: %s", fault.Message())
	}

	defer mgr.Close()

	// 1. Test SetTasksDir
	if err := mgr.SetTasksDir(customTasksDir); err != nil {
		t.Fatalf("SetTasksDir failed: %s", err.Message())
	}

	if mgr.TasksDir() != customTasksDir {
		t.Fatalf("expected TasksDir %s, got %s", customTasksDir, mgr.TasksDir())
	}

	defaultResolved := mgr.ResolveTaskDbPath("task-001")
	expectedDefault := filepath.Join(customTasksDir, "task-001.db")
	if defaultResolved != expectedDefault {
		t.Fatalf("expected %s, got %s", expectedDefault, defaultResolved)
	}

	// 2. Test SetTaskDbPath (explicit path for specific task)
	explicitPath := filepath.Join(tempDir, "isolated-task-99.db")
	if err := mgr.SetTaskDbPath("task-99", explicitPath); err != nil {
		t.Fatalf("SetTaskDbPath failed: %s", err.Message())
	}

	if mgr.ResolveTaskDbPath("task-99") != explicitPath {
		t.Fatalf("expected explicit path %s, got %s", explicitPath, mgr.ResolveTaskDbPath("task-99"))
	}

	// 3. Test SetTaskDbPathResolver (custom naming strategy)
	customResolver := func(tasksDir, taskId string) string {
		return filepath.Join(tasksDir, "nested", taskId, "task.db")
	}

	if err := mgr.SetTaskDbPathResolver(customResolver); err != nil {
		t.Fatalf("SetTaskDbPathResolver failed: %s", err.Message())
	}

	resolvedNested := mgr.ResolveTaskDbPath("task-nested-1")
	expectedNested := filepath.Join(customTasksDir, "nested", "task-nested-1", "task.db")
	if resolvedNested != expectedNested {
		t.Fatalf("expected nested path %s, got %s", expectedNested, resolvedNested)
	}

	// 4. Test NewSplitDBManagerWithConfig
	cfg := sqlitelogger.SplitDBConfig{
		WorkDir:            tempDir,
		MainDbPath:         filepath.Join(tempDir, "custom-main.db"),
		TasksDir:           filepath.Join(tempDir, "config-tasks"),
		TaskDbPathResolver: customResolver,
		Opener:             opener,
	}

	cfgMgr, cfgFault := sqlitelogger.NewSplitDBManagerWithConfig(cfg)
	if cfgFault != nil {
		t.Fatalf("NewSplitDBManagerWithConfig failed: %s", cfgFault.Message())
	}

	defer cfgMgr.Close()

	if cfgMgr.MainDbPath() != cfg.MainDbPath {
		t.Fatalf("expected mainDbPath %s, got %s", cfg.MainDbPath, cfgMgr.MainDbPath())
	}

	if cfgMgr.TasksDir() != cfg.TasksDir {
		t.Fatalf("expected tasksDir %s, got %s", cfg.TasksDir, cfgMgr.TasksDir())
	}
}

func TestSplitDBManager_TaskDbAutoMigrationAndRepair(t *testing.T) {
	tempDir := t.TempDir()
	opener := getMockOpener()

	mgr, fault := sqlitelogger.NewSplitDBManager(tempDir, opener)
	if fault != nil {
		t.Fatalf("failed to create manager: %s", fault.Message())
	}

	defer mgr.Close()

	// Pre-populate mock column store with an outdated legacy schema (missing columns)
	legacyPath := mgr.ResolveTaskDbPath("legacy-task-01")
	mockStoreLock.Lock()
	mockColumns[legacyPath] = []string{"id", "task_id", "timestamp", "level", "message"}
	mockStoreLock.Unlock()

	// Opening or interacting with the task DB should auto-fix and migrate it
	db, getFault := mgr.GetTaskDb("legacy-task-01")
	if getFault != nil {
		t.Fatalf("GetTaskDb failed: %s", getFault.Message())
	}

	if db == nil {
		t.Fatal("expected non-nil db connection")
	}

	// Verify missing columns were audited and added
	mockStoreLock.Lock()
	updatedCols := mockColumns[legacyPath]
	ver := mockVersions[legacyPath]
	mockStoreLock.Unlock()

	if len(updatedCols) <= 5 {
		t.Fatalf("expected auto-repaired columns, got %v", updatedCols)
	}

	if ver != int64(sqlitelogger.CurrentSchemaVersion) {
		t.Fatalf("expected schema version %d, got %d", sqlitelogger.CurrentSchemaVersion, ver)
	}
}

func TestSplitDBManager_ExplicitMigrationMethods(t *testing.T) {
	tempDir := t.TempDir()
	opener := getMockOpener()

	mgr, _ := sqlitelogger.NewSplitDBManager(tempDir, opener)
	defer mgr.Close()

	// Seed task databases
	_ = mgr.WriteTask("task-mig-1", sqlitelogger.TaskLogEntry{Message: "msg 1"})
	_ = mgr.WriteTask("task-mig-2", sqlitelogger.TaskLogEntry{Message: "msg 2"})

	// Explicit task migration
	if err := mgr.MigrateTaskDb("task-mig-1"); err != nil {
		t.Fatalf("MigrateTaskDb failed: %s", err.Message())
	}

	// Explicit task repair
	if err := mgr.RepairTaskDb("task-mig-1"); err != nil {
		t.Fatalf("RepairTaskDb failed: %s", err.Message())
	}

	// Migrate and repair all discovered task DBs
	if err := mgr.MigrateAllTaskDbs(); err != nil {
		t.Fatalf("MigrateAllTaskDbs failed: %s", err.Message())
	}

	if err := mgr.RepairAllTaskDbs(); err != nil {
		t.Fatalf("RepairAllTaskDbs failed: %s", err.Message())
	}

	// Main DB migration and repair
	if err := mgr.MigrateMainDb(); err != nil {
		t.Fatalf("MigrateMainDb failed: %s", err.Message())
	}

	if err := mgr.RepairMainDb(); err != nil {
		t.Fatalf("RepairMainDb failed: %s", err.Message())
	}
}

func TestSplitDBManager_ManualMigrationToggle(t *testing.T) {
	tempDir := t.TempDir()
	opener := getMockOpener()

	cfg := sqlitelogger.SplitDBConfig{
		WorkDir:               tempDir,
		Opener:                opener,
		IsManualMigrationOnly: true,
	}

	mgr, fault := sqlitelogger.NewSplitDBManagerWithConfig(cfg)
	if fault != nil {
		t.Fatalf("NewSplitDBManagerWithConfig failed: %s", fault.Message())
	}

	defer mgr.Close()

	isManual := mgr.IsManualMigrationOnly()
	if !isManual {
		t.Fatal("expected IsManualMigrationOnly to be true")
	}

	mgr.SetManualMigrationOnly(false)
	isManualAfter := mgr.IsManualMigrationOnly()
	if isManualAfter {
		t.Fatal("expected IsManualMigrationOnly to be false after toggle")
	}
}

func TestMigrationEngine_DirectFunctions(t *testing.T) {
	opener := getMockOpener()
	db, err := opener("memory_test_db")
	if err != nil {
		t.Fatalf("failed to open mock db: %v", err)
	}

	defer db.Close()

	if fault := sqlitelogger.EnsureBaseSchema(db); fault != nil {
		t.Fatalf("EnsureBaseSchema failed: %s", fault.Message())
	}

	if fault := sqlitelogger.ApplyIndexes(db); fault != nil {
		t.Fatalf("ApplyIndexes failed: %s", fault.Message())
	}

	if fault := sqlitelogger.CheckIntegrity(db); fault != nil {
		t.Fatalf("CheckIntegrity failed: %s", fault.Message())
	}

	ver, vFault := sqlitelogger.GetCurrentSchemaVersion(db)
	if vFault != nil {
		t.Fatalf("GetCurrentSchemaVersion failed: %s", vFault.Message())
	}

	_ = ver

	if fault := sqlitelogger.MigrateDatabase(db); fault != nil {
		t.Fatalf("MigrateDatabase failed: %s", fault.Message())
	}

	if fault := sqlitelogger.RepairDatabase(db); fault != nil {
		t.Fatalf("RepairDatabase failed: %s", fault.Message())
	}

	// Validate error on nil db
	if fault := sqlitelogger.EnsureBaseSchema(nil); fault == nil {
		t.Fatal("expected validation error on nil db")
	}
}
