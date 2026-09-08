package errcmd

import (
	"context"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/applogger/sqlitelogger"
)

// RunPowerShellWithTaskLog executes PowerShell commands with automatic task database logging.
func RunPowerShellWithTaskLog(
	ctx context.Context,
	script string,
	logger *sqlitelogger.TaskLogger,
) (*CommandResult, *appfault.AppError) {
	builder := NewScriptBuilder().
		SetShell(ShellPowerShell).
		AddLine(script)

	return NewRunner(builder).
		WithTaskLogger(logger).
		Run(ctx)
}

// RunBashWithTaskLog executes Bash commands with automatic task database logging.
func RunBashWithTaskLog(
	ctx context.Context,
	script string,
	logger *sqlitelogger.TaskLogger,
) (*CommandResult, *appfault.AppError) {
	builder := NewScriptBuilder().
		SetShell(ShellBash).
		AddLine(script)

	return NewRunner(builder).
		WithTaskLogger(logger).
		Run(ctx)
}

// RunAutoWithTaskLog executes scripts on the host's native shell with task database logging.
func RunAutoWithTaskLog(
	ctx context.Context,
	script string,
	logger *sqlitelogger.TaskLogger,
) (*CommandResult, *appfault.AppError) {
	builder := NewScriptBuilder().
		AddLine(script)

	return NewRunner(builder).
		WithTaskLogger(logger).
		Run(ctx)
}
