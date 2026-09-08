# Subtask 05: Command Logger Integration with Task-by-Task SQLite DB & Rotating File Sink

## Parent Plan
`.lovable/plans/completed/08-split-sqlite-logging-rotating-lazyonce-errcmd.md`

## Target Files
- `04-code/golang/pkg/errcmd/command_runner.go`
- `04-code/golang/pkg/errcmd/command_logger.go`

## Instructions
1. Implement `CommandRunner` executing scripts with context, capturing execution time, exit code, stdout, and stderr.
2. Implement `WithTaskLogging` attaching a `SplitDBManager` and/or `Logger` to command execution.
3. Automatically record command invocation, arguments, duration, exit code, stdout, and stderr into the task-specific SQLite DB (`tasks/<task-id>.db`).
4. Ensure error propagation wraps failures into `*appfault.AppError`.
5. Functions <= 15 lines.
