# Architectural Specification: Split SQLite Logging, Rotating Logger, LazyOnce & errcmd Integration

## Executive Summary

This specification establishes an enterprise-grade logging, command orchestration, and lazy evaluation architecture within the repository. The implementation integrates four core pillars:
1. **Split SQLite DB Logging Architecture (`applogger/sqlitelogger`):** Dedicated work/DB directory managing a primary `logs.db` and isolated, task-by-task SQLite databases in a `tasks/` directory, mitigating database locking and enabling independent task auditing, query APIs, and lifecycle management.
2. **Text-Based File Logging with Configurable Rotation & Archiving (`applogger`):** A robust rotating file sink with configurable maximum file sizes (default 2 MB), configurable retention limits (default 20 logs), directory archiving, and optional compression.
3. **Generic Lazy Once Engine (`lazyonce`):** A thread-safe, generic once-execution framework offering 0-param, 1-param, and 2-param evaluation variants (`LazyOnce[T]`, `LazyOnce1[TInput, TOutput]`, `LazyOnce2[T1, T2, TOutput]`), fully integrated with `result.Result[T]` and `*appfault.AppError`.
4. **Cross-Platform OS Script Execution & Command Logger Integration (`errcmd`):** Safe, high-performance PowerShell and Bash script builders inspired by `03-aukgo/errorwrapper/errcmd`, integrated with `errdefer` safe cleanup and direct task-by-task SQLite & rotating file logger recording.

---

## Task-Specific Rule Set (Strict Enforcement)

1. **Rule 1 (SQLite Split Isolation):** Task-specific logs MUST never contend on the global `logs.db`. Every task execution MUST write to its dedicated database at `<DBDir>/tasks/<task-id>.db`. All database operations must handle connection lifecycle and directory creation safely.
2. **Rule 2 (Rotation Predictability):** File rotation MUST trigger strictly when current file size + new entry size exceeds `MaxSizeBytes` (default 2 MB). The total number of preserved backup logs MUST never exceed `MaxBackups` (default 20). Archiving MUST move rotated files to the designated `ArchiveDir`.
3. **Rule 3 (Thread-Safe Lazy Once Invariant):** The `LazyOnce` initializer function MUST execute exactly once across concurrent callers. Subsequent invocations MUST return the memoized `Result[T]` and `*appfault.AppError` with zero re-evaluation.
4. **Rule 4 (Universal AppError & Single Return Types):** All public methods returning errors MUST return `*appfault.AppError`. Never return raw standard `error` or mixed polarity tuples.
5. **Rule 5 (Function Sizing & LF Line Endings):** All function bodies MUST remain <= 15 lines. Code must strictly use Unix LF (`\n`) and lowercase filenames.

---

## Architectural Breakdown

### 1. Split SQLite DB Architecture (`04-code/golang/pkg/applogger/sqlitelogger`)

```
<WorkDir / DBDir>/
├── logs.db                # Default main database for global / application logs
└── tasks/                 # Task databases directory
    ├── task-1001.db       # Isolated SQLite DB for task 1001
    ├── task-1002.db       # Isolated SQLite DB for task 1002
    └── ...
```

- **Manager Struct:** `SplitDBManager`
  - `workDir string`: Root directory for databases.
  - `mainDBPath string`: Path to global `logs.db` (customizable).
  - `tasksDir string`: Path to `<workDir>/tasks/`.
  - `opener DBOpenerFunc`: Pluggable `func(dsn string) (*sql.DB, error)` allowing testing with mock/in-memory drivers as well as production SQLite drivers (`sqlite`, `sqlite3`).
- **Log Entry Schema:**
  - Table: `app_logs` and `task_logs`
  - Columns: `id INTEGER PRIMARY KEY AUTOINCREMENT`, `task_id TEXT`, `timestamp TEXT`, `level TEXT`, `message TEXT`, `caller TEXT`, `fields_json TEXT`, `stack_trace TEXT`, `duration_ms INTEGER`, `status TEXT`.
- **API Methods:**
  - `WriteMain(entry LogEntry) *appfault.AppError`
  - `WriteTask(taskId string, entry LogEntry) *appfault.AppError`
  - `QueryMainLogs(filter FilterOptions) ([]LogEntry, *appfault.AppError)`
  - `QueryTaskLogs(taskId string, filter FilterOptions) ([]LogEntry, *appfault.AppError)`
  - `ListTaskDBs() ([]string, *appfault.AppError)`
  - `GetTaskSummary(taskId string) (*TaskSummary, *appfault.AppError)`
  - `Close() *appfault.AppError`

### 2. Rotating File Sink & Archiving (`04-code/golang/pkg/applogger`)

- **Configuration:** `RotationConfig`
  - `FilePath string`: Path to active log file.
  - `MaxSizeBytes int64`: Maximum file size before rotation (default: `2 * 1024 * 1024` = 2 MB).
  - `MaxBackups int`: Maximum number of rotated log files to keep (default: 20).
  - `IsArchiveEnabled bool`: Whether rotated logs are archived into an archive folder.
  - `ArchiveDir string`: Directory where archived logs are moved (e.g. `<log_dir>/archives/`).
  - `IsCompress bool`: Whether to compress archived files (`.gz`).
- **Sink Struct:** `RotatingFileSink`
  - Implements `LogSink`.
  - Thread-safe write lock with atomic file size tracking.
  - Automatic rotation when write breaches `MaxSizeBytes`.
  - Pruning of oldest rotated logs beyond `MaxBackups`.

### 3. Generic Lazy Once Engine (`04-code/golang/pkg/lazyonce`)

- **Type `LazyOnce[T]` (0 parameters):**
  - Constructor: `New[T](fn func() (T, *appfault.AppError)) *LazyOnce[T]`
  - `Value() (T, *appfault.AppError)`
  - `Result() result.Result[T]`
  - `IsEvaluated() bool`
  - `Reset()` (for testing / reload)
- **Type `LazyOnce1[TInput, TOutput]` (1 parameter):**
  - Constructor: `New1[TInput, TOutput](fn func(TInput) (TOutput, *appfault.AppError)) *LazyOnce1[TInput, TOutput]`
  - `Value(input TInput) (TOutput, *appfault.AppError)`
  - `Result(input TInput) result.Result[TOutput]`
  - `IsEvaluated() bool`
- **Type `LazyOnce2[T1, T2, TOutput]` (2 parameters):**
  - Constructor: `New2[T1, T2, TOutput](fn func(T1, T2) (TOutput, *appfault.AppError)) *LazyOnce2[T1, T2, TOutput]`
  - `Value(arg1 T1, arg2 T2) (TOutput, *appfault.AppError)`
  - `Result(arg1 T1, arg2 T2) result.Result[TOutput]`
  - `IsEvaluated() bool`

### 4. Cross-Platform OS Command Engine & Task Logging Integration (`04-code/golang/pkg/errcmd`)

- **Cross-Platform Script Builders:**
  - `PowerShellBuilder`: Windows native script invocation (`powershell.exe -NoProfile -NonInteractive -ExecutionPolicy Bypass -Command ...`).
  - `BashBuilder`: Linux / Unix native script invocation (`bash -c ...` or `sh -c ...`).
  - `AutoScriptBuilder`: Inspects `runtime.GOOS` and selects the appropriate shell script engine.
- **Task Logger Integration:**
  - `CmdRunner`: Executes scripts, capturing execution time, exit code, stdout, and stderr.
  - Attaches directly to `SplitDBManager`: automatically writes command start, command completion, exit codes, and outputs into the task's SQLite database (`tasks/<task-id>.db`) and/or the rotating file logger.
- **Safe Defer Utilities (`errdefer`):**
  - Thread-safe deferred cleanup handlers capturing panics, closing files, and preserving `*appfault.AppError` chains.

---

## Verification Plan

1. **Unit & Integration Tests:**
   - Test Split DB operations: writing to main, writing to task DBs, listing task DBs, querying logs.
   - Test file rotation: writing logs > 2 MB, verifying rotation, verifying max 20 backups retention, verifying archive creation.
   - Test LazyOnce: 0-param, 1-param, 2-param variants, verifying single execution under high concurrency.
   - Test errcmd: script generation, command execution with task logger recording, cross-platform OS handling.
2. **Quality Gates:**
   - Pass all 36 repository quality gates via `python 03-ai-scripts/06-cicd-local-runner.py`.
