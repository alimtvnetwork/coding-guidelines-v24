package examples_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"io"
	"path/filepath"
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
	l, err := examples.ExampleRotatingFileLogger(tempDir)
	if err != nil {
		t.Fatalf("ExampleRotatingFileLogger failed: %v", err)
	}

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
	tempDir := t.TempDir()
	mgr, _ := sqlitelogger.NewSplitDBManager(tempDir, getExOpener())
	defer mgr.Close()

	ctx := context.Background()
	res := examples.ExampleCommandRunWithTelemetry(ctx, "deploy-task-01", mgr)
	if res.IsFailed() {
		t.Fatalf("ExampleCommandRunWithTelemetry failed: %s", res.Fault().Message())
	}

	if !res.Data().IsSuccess {
		t.Fatalf("expected command success")
	}

	_ = filepath.Join
}
