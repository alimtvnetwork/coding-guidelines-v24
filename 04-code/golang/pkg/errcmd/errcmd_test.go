package errcmd_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/applogger/sqlitelogger"
	"coding-guidelines/common/pkg/errcmd"
)

// Mock SQL driver for errcmd integration tests
type mockStoreRow struct {
	taskId     string
	level      string
	message    string
	durationMs int64
	status     string
}

var (
	mockRowsLock sync.Mutex
	mockRowsList = make(map[string][]mockStoreRow)
	regMockOnce  sync.Once
)

type cmdMockDriver struct{}

func (d *cmdMockDriver) Open(name string) (driver.Conn, error) {
	return &cmdMockConn{dbName: name}, nil
}

type cmdMockConn struct{ dbName string }

func (c *cmdMockConn) Prepare(query string) (driver.Stmt, error) {
	return &cmdMockStmt{conn: c, query: query}, nil
}

func (c *cmdMockConn) Close() error              { return nil }
func (c *cmdMockConn) Begin() (driver.Tx, error) { return &cmdMockTx{}, nil }

type cmdMockTx struct{}

func (t *cmdMockTx) Commit() error   { return nil }
func (t *cmdMockTx) Rollback() error { return nil }

type cmdMockStmt struct {
	conn  *cmdMockConn
	query string
}

func (s *cmdMockStmt) Close() error  { return nil }
func (s *cmdMockStmt) NumInput() int { return -1 }
func (s *cmdMockStmt) Exec(args []driver.Value) (driver.Result, error) {
	mockRowsLock.Lock()
	defer mockRowsLock.Unlock()

	q := strings.ToUpper(s.query)
	if strings.Contains(q, "INSERT INTO LOGS") && len(args) >= 9 {
		row := mockStoreRow{
			taskId:     fmt.Sprint(args[0]),
			level:      fmt.Sprint(args[2]),
			message:    fmt.Sprint(args[3]),
			durationMs: args[7].(int64),
			status:     fmt.Sprint(args[8]),
		}

		mockRowsList[s.conn.dbName] = append(mockRowsList[s.conn.dbName], row)
	}

	return driver.RowsAffected(1), nil
}

func (s *cmdMockStmt) Query(args []driver.Value) (driver.Rows, error) {
	return &cmdEmptyRows{}, nil
}

type cmdEmptyRows struct{}

func (r *cmdEmptyRows) Columns() []string              { return []string{"col"} }
func (r *cmdEmptyRows) Close() error                   { return nil }
func (r *cmdEmptyRows) Next(dest []driver.Value) error { return io.EOF }

func getCmdMockOpener() sqlitelogger.DBOpenerFunc {
	regMockOnce.Do(func() {
		sql.Register("cmd_mock_sqlite", &cmdMockDriver{})
	})

	return func(dsn string) (*sql.DB, error) {
		return sql.Open("cmd_mock_sqlite", dsn)
	}
}

func TestScriptBuilder_ShellFormatting(t *testing.T) {
	psBuilder := errcmd.NewScriptBuilder().
		SetShell(errcmd.ShellPowerShell).
		AddLine("Write-Output 'hello'")

	cmd, fault := psBuilder.BuildCommand()
	if fault != nil {
		t.Fatalf("failed to build powershell command: %s", fault.Message())
	}

	if len(cmd.Args) < 2 || !strings.Contains(cmd.Args[0], "powershell") {
		t.Fatalf("expected powershell executable, got %v", cmd.Args)
	}

	bashBuilder := errcmd.NewScriptBuilder().
		SetShell(errcmd.ShellBash).
		AddLines("echo 'one'", "echo 'two'")

	bCmd, bFault := bashBuilder.BuildCommand()
	if bFault != nil {
		t.Fatalf("failed to build bash command: %s", bFault.Message())
	}

	if bCmd.Args[0] != "bash" {
		t.Fatalf("expected bash binary, got %s", bCmd.Args[0])
	}
}

func TestScriptBuilder_ValidationRejectsEmpty(t *testing.T) {
	emptyBuilder := errcmd.NewScriptBuilder()
	_, fault := emptyBuilder.BuildCommand()
	if fault == nil {
		t.Fatalf("expected validation fault on empty lines")
	}
}

type faultyCloser struct{ isFail bool }

func (f *faultyCloser) Close() error {
	if f.isFail {
		return fmt.Errorf("close broken")
	}

	return nil
}

func TestSafeDefer_ClosesAndRecovers(t *testing.T) {
	var faultHolder *appfault.AppError

	// Success closer
	errcmd.SafeClose(&faultyCloser{isFail: false}, &faultHolder)
	if faultHolder != nil {
		t.Fatalf("expected nil faultHolder, got %v", faultHolder)
	}

	// Faulty closer
	errcmd.SafeClose(&faultyCloser{isFail: true}, &faultHolder)
	if faultHolder == nil {
		t.Fatalf("expected error wrapped in faultHolder")
	}

	// SafeRecover
	var recoverFault *appfault.AppError
	func() {
		defer errcmd.SafeRecover(&recoverFault, "test context")
		panic("boom")
	}()

	if recoverFault == nil || !strings.Contains(recoverFault.Message(), "boom") {
		t.Fatalf("expected recovered panic fault, got %v", recoverFault)
	}
}

func TestCommandRunner_ExecutionAndTelemetry(t *testing.T) {
	tempDir := t.TempDir()
	mgr, _ := sqlitelogger.NewSplitDBManager(tempDir, getCmdMockOpener())
	defer mgr.Close()

	taskLogger, _ := sqlitelogger.NewTaskLogger("task-runner-01", mgr)

	// Build a cross-platform command
	builder := errcmd.NewScriptBuilder().
		AddLine("echo 'telemetry-ok'")

	runner := errcmd.NewRunner(builder).
		WithTaskLogger(taskLogger).
		WithSplitDB(mgr, "task-runner-01").
		WithTimeout(10 * time.Second)

	ctx := context.Background()
	res, fault := runner.Run(ctx)
	if fault != nil {
		t.Fatalf("command failed unexpectedly: %s", fault.Message())
	}

	if !res.IsSuccess || res.ExitCode != 0 {
		t.Fatalf("expected success, got exit=%d, out=%s, err=%s", res.ExitCode, res.Stdout, res.Stderr)
	}

	if !strings.Contains(res.Stdout, "telemetry-ok") {
		t.Fatalf("expected output to contain telemetry-ok, got %s", res.Stdout)
	}
}

func TestCommandLogger_Helpers(t *testing.T) {
	tempDir := t.TempDir()
	mgr, _ := sqlitelogger.NewSplitDBManager(tempDir, getCmdMockOpener())
	defer mgr.Close()

	taskLogger, _ := sqlitelogger.NewTaskLogger("helper-task", mgr)
	ctx := context.Background()

	res, fault := errcmd.RunAutoWithTaskLog(ctx, "echo 'auto-test'", taskLogger)
	if fault != nil {
		t.Fatalf("auto run failed: %s", fault.Message())
	}

	if !res.IsSuccess {
		t.Fatalf("expected success on auto run")
	}

	_ = os.Getenv
}

func TestCommandRunner_StreamingStdout(t *testing.T) {
	builder := errcmd.NewScriptBuilder().
		AddLine("echo 'stream-line-1'").
		AddLine("echo 'stream-line-2'")

	var lines []string
	var mu sync.Mutex
	runner := errcmd.NewRunner(builder).
		WithStdoutHandler(func(line string) {
			mu.Lock()
			lines = append(lines, line)
			mu.Unlock()
		}).
		WithTimeout(10 * time.Second)

	res, fault := runner.Run(context.Background())
	if fault != nil {
		t.Fatalf("run failed: %s", fault.Message())
	}

	verifyStreamingLines(t, lines, res)
}

func verifyStreamingLines(t *testing.T, lines []string, res *errcmd.CommandResult) {
	if len(lines) == 0 {
		t.Fatal("expected streamed stdout lines, got none")
	}

	foundLine := false
	for _, l := range lines {
		if strings.Contains(l, "stream-line-1") {
			foundLine = true
			break
		}
	}

	if !foundLine {
		t.Fatalf("expected stream-line-1 in lines: %v", lines)
	}

	_ = res
}

func TestCommandRunner_WithEnvAndCwd(t *testing.T) {
	tempDir := t.TempDir()
	builder := errcmd.NewScriptBuilder()
	if runtime.GOOS == "windows" {
		builder.AddLine("echo $env:STREAM_VAR")
	} else {
		builder.AddLine("echo $STREAM_VAR")
	}

	runner := errcmd.NewRunner(builder).
		WithCwd(tempDir).
		WithEnv(map[string]string{"STREAM_VAR": "custom_value_42"}).
		WithTimeout(10 * time.Second)

	res, fault := runner.Run(context.Background())
	if fault != nil {
		t.Fatalf("run with env/cwd failed: %s", fault.Message())
	}

	if !strings.Contains(res.Stdout, "custom_value_42") {
		t.Fatalf("expected custom_value_42 in stdout, got %s", res.Stdout)
	}
}

func TestCommandRunner_WithStderrHandler(t *testing.T) {
	builder := errcmd.NewScriptBuilder()
	if runtime.GOOS == "windows" {
		builder.AddLine("[Console]::Error.WriteLine('custom-stderr-line')")
	} else {
		builder.AddLine("echo 'custom-stderr-line' >&2")
	}

	var errLines []string
	var mu sync.Mutex
	runner := errcmd.NewRunner(builder).
		WithStderrHandler(func(line string) {
			mu.Lock()
			errLines = append(errLines, line)
			mu.Unlock()
		}).
		WithTimeout(10 * time.Second)

	res, _ := runner.Run(context.Background())
	verifyStderrLines(t, errLines, res)
}

func verifyStderrLines(t *testing.T, errLines []string, res *errcmd.CommandResult) {
	if len(errLines) == 0 {
		t.Fatal("expected streamed stderr lines, got none")
	}

	if !strings.Contains(errLines[0], "custom-stderr-line") {
		t.Fatalf("expected custom-stderr-line, got %v", errLines)
	}

	_ = res
}
