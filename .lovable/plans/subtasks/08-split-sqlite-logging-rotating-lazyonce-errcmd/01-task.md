# Subtask 01: Split SQLite Logging Architecture & DB Directory Manager

## Parent Plan
`.lovable/plans/completed/08-split-sqlite-logging-rotating-lazyonce-errcmd.md`

## Target Files
- `04-code/golang/pkg/applogger/sqlitelogger/sqlitelogger.go`
- `04-code/golang/pkg/applogger/sqlitelogger/manager.go`
- `04-code/golang/pkg/applogger/sqlitelogger/models.go`

## Instructions
1. Implement the `SplitDBManager` managing `WorkDir` / `DBDir`.
2. Implement main database `logs.db` and task databases directory `<WorkDir>/tasks/`.
3. Support pluggable `DBOpenerFunc` so standard drivers or memory/mock drivers can be injected.
4. Implement table creation (`app_logs`, `task_logs`).
5. Implement `WriteMain(entry LogEntry) *appfault.AppError` and `WriteTask(taskId string, entry LogEntry) *appfault.AppError`.
6. Implement `QueryMainLogs(filter FilterOptions) ([]LogEntry, *appfault.AppError)` and `QueryTaskLogs(taskId string, filter FilterOptions) ([]LogEntry, *appfault.AppError)`.
7. Implement `ListTaskDBs() ([]string, *appfault.AppError)`.
8. Enforce all coding guidelines: functions <= 15 lines, `*appfault.AppError` return type, strict relative paths.
