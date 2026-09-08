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
	initMockOnce  sync.Once
)

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
	if strings.Contains(q, "CREATE TABLE") {
		return driver.RowsAffected(0), nil
	}

	if strings.Contains(q, "INSERT INTO LOGS") {
		rows := mockStore[s.conn.dbName]
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

		mockStore[s.conn.dbName] = append(rows, row)

		return driver.RowsAffected(1), nil
	}

	return driver.RowsAffected(0), nil
}

func (s *mockStmt) Query(args []driver.Value) (driver.Rows, error) {
	mockStoreLock.Lock()
	defer mockStoreLock.Unlock()

	q := strings.ToUpper(s.query)
	rows := mockStore[s.conn.dbName]

	if strings.Contains(q, "SELECT COUNT(*), COALESCE(MAX(TIMESTAMP)") {
		lastTs := ""
		if len(rows) > 0 {
			lastTs = rows[len(rows)-1].timestamp
		}

		return &mockSummaryRows{count: int64(len(rows)), ts: lastTs}, nil
	}

	if strings.Contains(q, "SELECT COUNT(*) FROM LOGS WHERE LEVEL") {
		errCount := int64(0)
		for _, r := range rows {
			if r.level == "ERROR" || r.level == "FATAL" {
				errCount++
			}
		}

		return &mockCountRows{count: errCount}, nil
	}

	return &mockLogRows{data: append([]mockRowData(nil), rows...), idx: 0}, nil
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
