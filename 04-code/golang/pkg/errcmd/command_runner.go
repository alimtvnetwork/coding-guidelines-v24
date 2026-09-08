package errcmd

import (
	"bytes"
	"context"
	"os/exec"
	"time"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/applogger"
	"coding-guidelines/common/pkg/applogger/sqlitelogger"
	"coding-guidelines/common/pkg/errtype"
)

// CommandResult encapsulates the execution metrics and outputs of a command.
type CommandResult struct {
	ExitCode   int                `json:"exitCode"`
	Stdout     string             `json:"stdout"`
	Stderr     string             `json:"stderr"`
	DurationMs int64              `json:"durationMs"`
	IsSuccess  bool               `json:"isSuccess"`
	Fault      *appfault.AppError `json:"fault,omitempty"`
}

// CommandRunner orchestrates command execution with integrated SQLite and file logging.
type CommandRunner struct {
	builder    *ScriptBuilder
	taskId     string
	taskLogger *sqlitelogger.TaskLogger
	splitMgr   *sqlitelogger.SplitDBManager
	fileSink   applogger.LogSink
	timeout    time.Duration
}

// NewRunner initializes a CommandRunner with the specified script builder.
func NewRunner(builder *ScriptBuilder) *CommandRunner {
	return &CommandRunner{
		builder: builder,
		timeout: 30 * time.Second,
	}
}

// WithTaskId assigns a task identifier to the runner.
func (r *CommandRunner) WithTaskId(taskId string) *CommandRunner {
	r.taskId = taskId

	return r
}

// WithTaskLogger attaches a TaskLogger for automated execution telemetry.
func (r *CommandRunner) WithTaskLogger(l *sqlitelogger.TaskLogger) *CommandRunner {
	r.taskLogger = l
	if l != nil {
		r.taskId = l.TaskId()
	}

	return r
}

// WithSplitDB binds a SplitDBManager for direct task database recording.
func (r *CommandRunner) WithSplitDB(
	mgr *sqlitelogger.SplitDBManager,
	taskId string,
) *CommandRunner {
	r.splitMgr = mgr
	r.taskId = taskId

	return r
}

// WithFileSink attaches an active file sink for text logging.
func (r *CommandRunner) WithFileSink(sink applogger.LogSink) *CommandRunner {
	r.fileSink = sink

	return r
}

// WithTimeout sets a maximum execution duration before cancellation.
func (r *CommandRunner) WithTimeout(d time.Duration) *CommandRunner {
	r.timeout = d

	return r
}

// Run executes the script, captures output streams, and logs results to databases and sinks.
func (r *CommandRunner) Run(ctx context.Context) (*CommandResult, *appfault.AppError) {
	cmd, fault := r.builder.BuildCommand()
	if fault != nil {
		return nil, fault
	}

	execCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	cmdWithCtx := exec.CommandContext(execCtx, cmd.Path, cmd.Args[1:]...)
	cmdWithCtx.Dir = cmd.Dir
	cmdWithCtx.Env = cmd.Env

	var stdoutBuf, stderrBuf bytes.Buffer
	cmdWithCtx.Stdout = &stdoutBuf
	cmdWithCtx.Stderr = &stderrBuf

	r.logStart(r.builder.ScriptText())
	startTime := time.Now()
	err := cmdWithCtx.Run()
	durationMs := time.Since(startTime).Milliseconds()

	res := r.buildResult(err, &stdoutBuf, &stderrBuf, durationMs)
	r.logCompletion(res)

	return res, res.Fault
}

// buildResult constructs the CommandResult from outputs and execution status.
func (r *CommandRunner) buildResult(
	err error,
	stdoutBuf, stderrBuf *bytes.Buffer,
	durationMs int64,
) *CommandResult {
	stdout := stdoutBuf.String()
	stderr := stderrBuf.String()
	exitCode := 0

	if err != nil {
		if exitErr, isExit := err.(*exec.ExitError); isExit {
			exitCode = exitErr.ExitCode()
		} else {
			exitCode = 1
		}

		return &CommandResult{
			ExitCode:   exitCode,
			Stdout:     stdout,
			Stderr:     stderr,
			DurationMs: durationMs,
			IsSuccess:  false,
			Fault:      appfault.Wrap(errtype.Execution, err, "command execution failed"),
		}
	}

	return &CommandResult{
		ExitCode:   0,
		Stdout:     stdout,
		Stderr:     stderr,
		DurationMs: durationMs,
		IsSuccess:  true,
		Fault:      nil,
	}
}

// logStart emits start notifications to attached loggers and sinks.
func (r *CommandRunner) logStart(script string) {
	msg := "starting command: " + script
	if r.taskLogger != nil {
		_ = r.taskLogger.Log("INFO", msg)
	}

	if r.fileSink != nil {
		_ = r.fileSink.WriteEntry(applogger.LogEntry{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Level:     applogger.LevelInfo,
			Message:   msg,
		})
	}
}

// logCompletion records execution metrics and exit status into attached databases and sinks.
func (r *CommandRunner) logCompletion(res *CommandResult) {
	status := "OK"
	level := applogger.LevelInfo
	if !res.IsSuccess {
		status = "ERROR"
		level = applogger.LevelError
	}

	if r.taskLogger != nil {
		_ = r.taskLogger.LogExecution(r.builder.ScriptText(), res.DurationMs, status)
	}

	if r.splitMgr != nil && r.taskId != "" {
		_ = r.splitMgr.WriteTask(r.taskId, sqlitelogger.TaskLogEntry{
			TaskId:     r.taskId,
			Timestamp:  time.Now().UTC().Format(time.RFC3339),
			Level:      level.Name(),
			Message:    r.builder.ScriptText(),
			DurationMs: res.DurationMs,
			Status:     status,
		})
	}

	if r.fileSink != nil {
		_ = r.fileSink.WriteEntry(applogger.LogEntry{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Level:     level,
			Message:   r.builder.ScriptText(),
		})
	}
}
