package examples_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"coding-guidelines/common/examples"
	"coding-guidelines/common/pkg/applogger/sqlitelogger"
)

type exMockDriver struct{}

func (d *exMockDriver) Open(name string) (driver.Conn, error) {
	return &exMockConn{}, nil
}

type exMockConn struct{}

func (c *exMockConn) Prepare(q string) (driver.Stmt, error) { return &exMockStmt{}, nil }
func (c *exMockConn) Close() error                          { return nil }
func (c *exMockConn) Begin() (driver.Tx, error)             { return &exMockTx{}, nil }

type exMockTx struct{}

func (t *exMockTx) Commit() error   { return nil }
func (t *exMockTx) Rollback() error { return nil }

type exMockStmt struct{}

func (s *exMockStmt) Close() error  { return nil }
func (s *exMockStmt) NumInput() int { return -1 }
func (s *exMockStmt) Exec(args []driver.Value) (driver.Result, error) {
	return driver.RowsAffected(1), nil
}

func (s *exMockStmt) Query(args []driver.Value) (driver.Rows, error) { return &exEmptyRows{}, nil }

type exEmptyRows struct{}

func (r *exEmptyRows) Columns() []string              { return []string{"id"} }
func (r *exEmptyRows) Close() error                   { return nil }
func (r *exEmptyRows) Next(dest []driver.Value) error { return io.EOF }

var regExMockOnce sync.Once

func getExOpener() sqlitelogger.DBOpenerFunc {
	regExMockOnce.Do(func() {
		sql.Register("ex_mock_driver", &exMockDriver{})
	})

	return func(dsn string) (*sql.DB, error) {
		return sql.Open("ex_mock_driver", dsn)
	}
}

func TestExamples_SplitSQLiteLogging(t *testing.T) {
	tempDir := t.TempDir()
	fault := examples.ExampleSplitSQLiteLogging(tempDir, getExOpener())
	if fault != nil {
		t.Fatalf("ExampleSplitSQLiteLogging failed: %s", fault.Message())
	}
}

func TestExamples_RotatingFileLogger(t *testing.T) {
	tempDir := t.TempDir()
	res := examples.ExampleRotatingFileLogger(tempDir)
	if res.IsFailed() {
		t.Fatalf("ExampleRotatingFileLogger failed: %v", res.Fault())
	}

	l := res.Data()
	l.Info("example log message")
	_ = l.Close()
}

func TestExamples_LazyOnceUsage(t *testing.T) {
	cfg, total, fault := examples.ExampleLazyOnceUsage()
	if fault != nil {
		t.Fatalf("ExampleLazyOnceUsage failed: %s", fault.Message())
	}

	if cfg != "loaded-db-connection-string" || total != 8 {
		t.Fatalf("unexpected values: cfg=%s, total=%d", cfg, total)
	}
}

func TestExamples_CommandRunWithTelemetry(t *testing.T) {
	mgr, fault := sqlitelogger.NewSplitDBManager(t.TempDir(), getExOpener())
	if fault != nil {
		t.Fatalf("failed to create db mgr: %v", fault)
	}

	defer mgr.Close()

	res := examples.ExampleCommandRunWithTelemetry(context.Background(), "deploy-task-01", mgr)
	if res.IsFailed() || !res.Data().IsSuccess {
		t.Fatalf("expected command success, fault: %v", res.Fault())
	}
}

func TestExamples_TaskRetentionAndFiltering(t *testing.T) {
	tempDir := t.TempDir()
	pruned, logs, fault := examples.ExampleTaskRetentionAndFiltering(tempDir, getExOpener())
	if fault != nil {
		t.Fatalf("ExampleTaskRetentionAndFiltering failed: %s", fault.Message())
	}

	if pruned != 0 {
		t.Fatalf("expected 0 pruned, got %d", pruned)
	}

	_ = logs
}

func TestExamples_LiveStreamingCommand(t *testing.T) {
	tempDir := t.TempDir()
	ctx := context.Background()
	res, captured, fault := examples.ExampleLiveStreamingCommand(ctx, tempDir)
	if fault != nil {
		t.Fatalf("ExampleLiveStreamingCommand failed: %s", fault.Message())
	}

	if !res.IsSuccess {
		t.Fatal("expected command success")
	}

	verifyCapturedLines(t, captured)
}

func verifyCapturedLines(t *testing.T, lines []string) {
	if len(lines) == 0 {
		t.Fatal("expected captured lines")
	}

	if !strings.Contains(lines[0], "stream step") {
		t.Fatalf("unexpected line: %s", lines[0])
	}
}

func TestExamples_AtomicFileWrite(t *testing.T) {
	tempDir := t.TempDir()
	target := filepath.Join(tempDir, "atomic_sample.txt")
	content := []byte("atomic sample data")

	fault := examples.ExampleAtomicFileWrite(target, content)
	if fault != nil {
		t.Fatalf("ExampleAtomicFileWrite failed: %s", fault.Message())
	}

	data, err := os.ReadFile(target)
	if err != nil || string(data) != string(content) {
		t.Fatalf("atomic file write content mismatch")
	}
}

func TestExamples_LazyOnceContextAndReset(t *testing.T) {
	ctx := context.Background()
	val, fault := examples.ExampleLazyOnceContextAndReset(ctx)
	if fault != nil {
		t.Fatalf("ExampleLazyOnceContextAndReset failed: %s", fault.Message())
	}

	if val != "initialized-service-instance" {
		t.Fatalf("unexpected val: %s", val)
	}
}

func TestExamples_ApiManagerRemoteLogging(t *testing.T) {
	res := examples.ExampleApiManagerRemoteLogging("https://logs.example.internal")
	if res.IsFailed() {
		t.Fatalf("ExampleApiManagerRemoteLogging failed: %v", res.Fault())
	}

	mgr := res.Data()
	defer mgr.Close()
	if mgr.BufferedCount() != 0 {
		t.Fatalf("expected 0 initial buffered logs")
	}
}
