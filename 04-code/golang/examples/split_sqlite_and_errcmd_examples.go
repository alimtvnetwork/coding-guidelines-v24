package examples

import (
	"context"
	"path/filepath"
	"time"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/applogger"
	"coding-guidelines/common/pkg/applogger/sqlitelogger"
	"coding-guidelines/common/pkg/errcmd"
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
