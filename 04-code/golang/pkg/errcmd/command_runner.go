package errcmd

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
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
	builder       *ScriptBuilder
	taskId        string
	taskLogger    *sqlitelogger.TaskLogger
	splitMgr      *sqlitelogger.SplitDBManager
	fileSink      applogger.LogSink
	timeout       time.Duration
	stdoutHandler func(line string)
	stderrHandler func(line string)
	env           map[string]string
	cwd           string
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

// WithStdoutHandler binds a streaming callback for each standard output line.
func (r *CommandRunner) WithStdoutHandler(handler func(line string)) *CommandRunner {
	r.stdoutHandler = handler

	return r
}

// WithStderrHandler binds a streaming callback for each standard error line.
func (r *CommandRunner) WithStderrHandler(handler func(line string)) *CommandRunner {
	r.stderrHandler = handler

	return r
}

// WithEnv configures environment variables to be injected into the process.
func (r *CommandRunner) WithEnv(env map[string]string) *CommandRunner {
	r.env = env

	return r
}

// WithCwd sets the working directory for command execution.
func (r *CommandRunner) WithCwd(dir string) *CommandRunner {
	r.cwd = dir

	return r
}

// Run executes the script, captures output streams, and logs results to databases and sinks.
func (r *CommandRunner) Run(ctx context.Context) (*CommandResult, *appfault.AppError) {
	cmd, cancel, fault := r.prepareCommand(ctx)
	if fault != nil {
		return nil, fault
	}

	defer cancel()

	var stdoutBuf, stderrBuf bytes.Buffer
	outWriter := newLineStreamWriter(&stdoutBuf, r.createStdoutHandler())
	errWriter := newLineStreamWriter(&stderrBuf, r.createStderrHandler())
	durationMs, err := r.executeWithStreams(cmd, outWriter, errWriter)
	res := r.buildResult(err, &stdoutBuf, &stderrBuf, durationMs)
	r.logCompletion(res)

	return res, res.Fault
}

func (r *CommandRunner) prepareCommand(
	ctx context.Context,
) (*exec.Cmd, context.CancelFunc, *appfault.AppError) {
	cmd, fault := r.builder.BuildCommand()
	if fault != nil {
		return nil, nil, fault
	}

	execCtx, cancel := context.WithTimeout(ctx, r.timeout)
	cmdWithCtx := exec.CommandContext(execCtx, cmd.Path, cmd.Args[1:]...)
	cmdWithCtx.Dir = cmd.Dir
	cmdWithCtx.Env = cmd.Env
	r.applyProcessContext(cmdWithCtx)

	return cmdWithCtx, cancel, nil
}

func (r *CommandRunner) executeWithStreams(
	cmd *exec.Cmd,
	outWriter, errWriter *lineStreamWriter,
) (int64, error) {
	cmd.Stdout = outWriter
	cmd.Stderr = errWriter
	r.logStart(r.builder.ScriptText())
	startTime := time.Now()
	err := cmd.Run()
	durationMs := time.Since(startTime).Milliseconds()
	outWriter.Flush()
	errWriter.Flush()

	return durationMs, err
}

func (r *CommandRunner) createStdoutHandler() func(string) {
	hasReceiver := r.stdoutHandler != nil || r.taskLogger != nil
	if !hasReceiver {
		return nil
	}

	return func(line string) {
		if r.stdoutHandler != nil {
			r.stdoutHandler(line)
		}

		if r.taskLogger != nil {
			_ = r.taskLogger.Log("INFO", line)
		}
	}
}

func (r *CommandRunner) createStderrHandler() func(string) {
	hasReceiver := r.stderrHandler != nil || r.taskLogger != nil
	if !hasReceiver {
		return nil
	}

	return func(line string) {
		if r.stderrHandler != nil {
			r.stderrHandler(line)
		}

		if r.taskLogger != nil {
			_ = r.taskLogger.Log("ERROR", line)
		}
	}
}

func (r *CommandRunner) applyProcessContext(cmd *exec.Cmd) {
	if r.cwd != "" {
		cmd.Dir = r.cwd
	}

	if len(r.env) > 0 {
		cmd.Env = mergeEnvironment(cmd.Env, r.env)
	}
}

func mergeEnvironment(base []string, overrides map[string]string) []string {
	envList := base
	if len(envList) == 0 {
		envList = os.Environ()
	}

	for k, v := range overrides {
		envList = append(envList, fmt.Sprintf("%s=%s", k, v))
	}

	return envList
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

// lineStreamWriter buffers raw bytes and passes complete lines to a handler.
type lineStreamWriter struct {
	buf     *bytes.Buffer
	handler func(string)
	pending bytes.Buffer
	lock    sync.Mutex
}

func newLineStreamWriter(buf *bytes.Buffer, handler func(string)) *lineStreamWriter {
	return &lineStreamWriter{
		buf:     buf,
		handler: handler,
	}
}

func (w *lineStreamWriter) Write(p []byte) (int, error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.buf != nil {
		w.buf.Write(p)
	}

	if w.handler == nil {
		return len(p), nil
	}

	return w.processLines(p)
}

func (w *lineStreamWriter) processLines(p []byte) (int, error) {
	w.pending.Write(p)
	for {
		line, err := w.pending.ReadString('\n')
		if err != nil {
			w.pending.WriteString(line)
			break
		}

		cleanLine := strings.TrimRight(line, "\r\n")
		w.handler(cleanLine)
	}

	return len(p), nil
}

func (w *lineStreamWriter) Flush() {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.handler != nil && w.pending.Len() > 0 {
		w.handler(w.pending.String())
		w.pending.Reset()
	}
}
