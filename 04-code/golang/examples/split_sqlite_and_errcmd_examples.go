package examples

import (
	"context"
	"path/filepath"
	"time"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/applogger"
	"coding-guidelines/common/pkg/applogger/sqlitelogger"
	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/errcmd"
	"coding-guidelines/common/pkg/fileutil"
	"coding-guidelines/common/pkg/lazyonce"
	"coding-guidelines/common/pkg/result"
)

// ExampleSplitSQLiteLogging demonstrates task-by-task isolated logging alongside global logs.
func ExampleSplitSQLiteLogging(
	workDir string,
	opener sqlitelogger.DBOpenerFunc,
) *appfault.AppError {
	mgr, fault := sqlitelogger.NewSplitDBManager(workDir, opener)
	if fault != nil {
		return fault
	}

	defer mgr.Close()

	// Global system log
	mainLog := sqlitelogger.TaskLogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     "INFO",
		Message:   "Application server initialized",
	}

	if err := mgr.WriteMain(mainLog); err != nil {
		return err
	}

	// Task-specific logging
	taskLogger, tlFault := sqlitelogger.NewTaskLogger("task-deploy-42", mgr)
	if tlFault != nil {
		return tlFault
	}

	if err := taskLogger.Log("INFO", "Beginning deployment step"); err != nil {
		return err
	}

	return taskLogger.LogExecution("Deployment completed successfully", 350, "OK")
}

// ExampleRotatingFileLogger demonstrates creating a rotating text logger with 2MB limit and archiving.
func ExampleRotatingFileLogger(logDir string) (applogger.Logger, error) {
	logFilePath := filepath.Join(logDir, "app-service.log")

	cfg := applogger.Config{
		MinLevel: applogger.LevelInfo,
		Driver:   applogger.DriverRotatingFile,
		Rotation: applogger.RotationConfig{
			FilePath:         logFilePath,
			MaxSizeBytes:     applogger.DefaultMaxSizeBytes, // 2 MB threshold
			MaxBackups:       applogger.DefaultMaxBackups,   // 20 logs retained
			IsArchiveEnabled: true,
			ArchiveDir:       filepath.Join(logDir, "archives"),
			IsCompress:       true,
		},
	}

	return applogger.New(cfg)
}

// ExampleLazyOnceUsage demonstrates 0-param, 1-param, and 2-param memoization.
func ExampleLazyOnceUsage() (string, int, *appfault.AppError) {
	// Zero-param lazy initialization
	lazyConfig := lazyonce.New(func() (string, *appfault.AppError) {
		return "loaded-db-connection-string", nil
	})

	cfgVal, cfgFault := lazyConfig.Value()
	if cfgFault != nil {
		return "", 0, cfgFault
	}

	// Single-param lazy initialization
	lazyWorker := lazyonce.New1(func(workerCount int) (int, *appfault.AppError) {
		return workerCount * 2, nil
	})

	workerTotal, _ := lazyWorker.Value(4)

	return cfgVal, workerTotal, nil
}

// ExampleCommandRunWithTelemetry demonstrates executing a command with task database logging.
func ExampleCommandRunWithTelemetry(
	ctx context.Context,
	taskId string,
	mgr *sqlitelogger.SplitDBManager,
) result.Result[*errcmd.CommandResult] {
	taskLogger, fault := sqlitelogger.NewTaskLogger(taskId, mgr)
	if fault != nil {
		return result.Failure[*errcmd.CommandResult](fault)
	}

	builder := errcmd.NewScriptBuilder().
		AddLine("echo 'deploy step'")

	runner := errcmd.NewRunner(builder).
		WithTaskLogger(taskLogger).
		WithSplitDB(mgr, taskId).
		WithTimeout(15 * time.Second)

	cmdRes, cmdFault := runner.Run(ctx)
	if cmdFault != nil {
		return result.Failure[*errcmd.CommandResult](cmdFault)
	}

	return result.Success(cmdRes)
}

// ExampleTaskRetentionAndFiltering demonstrates pruning old task databases and querying with filters.
func ExampleTaskRetentionAndFiltering(
	workDir string,
	opener sqlitelogger.DBOpenerFunc,
) (int, []sqlitelogger.TaskLogEntry, *appfault.AppError) {
	mgr, fault := sqlitelogger.NewSplitDBManager(workDir, opener)
	if fault != nil {
		return 0, nil, fault
	}

	defer mgr.Close()

	prunedCount, pFault := mgr.PruneTasks(48 * time.Hour)
	if pFault != nil {
		return 0, nil, pFault
	}

	filter := sqlitelogger.FilterOptions{Level: "ERROR", Limit: 25}
	logs, qFault := mgr.QueryTaskLogs("task-retention-demo", filter)

	return prunedCount, logs, qFault
}

// ExampleLiveStreamingCommand runs a command streaming stdout lines to a callback in real time.
func ExampleLiveStreamingCommand(
	ctx context.Context,
	workDir string,
) (*errcmd.CommandResult, []string, *appfault.AppError) {
	var captured []string
	builder := errcmd.NewScriptBuilder().
		AddLine("echo 'stream step 1'").
		AddLine("echo 'stream step 2'")

	runner := errcmd.NewRunner(builder).
		WithCwd(workDir).
		WithEnv(map[string]string{"MODE": "streaming"}).
		WithStdoutHandler(func(line string) {
			captured = append(captured, line)
		})

	res, fault := runner.Run(ctx)

	return res, captured, fault
}

// ExampleAtomicFileWrite safely writes data via temporary file and atomic swap.
func ExampleAtomicFileWrite(targetPath string, content []byte) *appfault.AppError {
	return fileutil.AtomicWriteFile(targetPath, content, filepermtype.Standard)
}

// ExampleLazyOnceContextAndReset demonstrates context-aware lazy evaluation and cache reset.
func ExampleLazyOnceContextAndReset(ctx context.Context) (string, *appfault.AppError) {
	lazyService := lazyonce.New(func() (string, *appfault.AppError) {
		return "initialized-service-instance", nil
	})

	val, fault := lazyService.ValueContext(ctx)
	if fault != nil {
		return "", fault
	}

	lazyService.Reset()

	return val, nil
}

// ExampleApiManagerRemoteLogging configures remote logging with custom headers and flush intervals.
func ExampleApiManagerRemoteLogging(endpoint string) (*applogger.ApiManager, error) {
	cfg := applogger.ApiConfig{
		Endpoint:      endpoint,
		Headers:       map[string]string{"Authorization": "Bearer secret-token"},
		BatchSize:     50,
		FlushInterval: 10 * time.Second,
	}

	mgr, err := applogger.NewApiManager(cfg)
	if err != nil {
		return nil, err
	}

	mgr.SetRotationPolicy(func(batch []applogger.LogEntry, elapsed time.Duration, c applogger.ApiConfig) bool {
		return len(batch) >= c.BatchSize || elapsed >= c.FlushInterval
	})

	return mgr, nil
}
