# Task 22: Split SQLite Logging, Rotating File Logger, LazyOnce & errcmd Integration

## 1. Header & Metadata

- **Date:** 2026-09-09
- **Author/Agent:** Antigravity Master Orchestrator
- **Status:** Completed
- **Affected Packages:**
  - `coding-guidelines/common/pkg/applogger`
  - `coding-guidelines/common/pkg/applogger/sqlitelogger`
  - `coding-guidelines/common/pkg/lazyonce`
  - `coding-guidelines/common/pkg/errcmd`
  - `coding-guidelines/common/examples`

---

## 2. Context & Goals

The user required a comprehensive logging and execution architecture addressing:
1. **Split SQLite DB Architecture:** Work directory (`workDir` / `DBDir`) hosting a primary `logs.db` for global system logs and a `tasks/` directory containing task-isolated SQLite databases (`tasks/<task-id>.db`). This architecture eliminates database file locking contention between parallel tasks and allows independent retrieval, pruning, and archiving.
2. **Text-Based File Logging with Configurable Rotation & Archiving:** A robust `RotatingFileSink` conforming to `LogSink`, supporting configurable maximum file size (default 2 MB), retention count (default 20 logs), directory archiving, and optional gzip compression.
3. **Generic Lazy Once Engine (`lazyonce`):** Inspired by `errwonce` but modern, generic, and flexible, providing 0-param, 1-param, and 2-param evaluation variants (`LazyOnce[T]`, `LazyOnce1[TInput, TOutput]`, `LazyOnce2[T1, T2, TOutput]`) with thread-safe memoization and full `result.Result[T]` / `*appfault.AppError` integration.
4. **Cross-Platform OS Script Execution & Command Logger Integration (`errcmd`):** Inspired by `03-aukgo/errorwrapper/errcmd`, implementing PowerShell and Bash script builders, safe deferral (`safedefer`), and direct automated telemetry recording into task SQLite databases and rotating text loggers.

---

## 3. Files Changed / Created

### New Packages & Source Files

- `04-code/golang/pkg/applogger/sqlitelogger/models.go`: Data models (`TaskLogEntry`, `FilterOptions`, `TaskSummary`, `DBOpenerFunc`).
- `04-code/golang/pkg/applogger/sqlitelogger/manager.go`: `SplitDBManager` managing global and task SQLite databases.
- `04-code/golang/pkg/applogger/sqlitelogger/sqlitelogger.go`: `TaskLogger` binding task-scoped logging to task DBs.
- `04-code/golang/pkg/applogger/sqlitelogger/sqlitelogger_test.go`: Unit tests for `SplitDBManager` and `TaskLogger`.
- `04-code/golang/pkg/applogger/rotation_config.go`: `RotationConfig` defining size limits, retention, and archive paths.
- `04-code/golang/pkg/applogger/rotating_file_sink.go`: `RotatingFileSink` implementing size-based rotation and archiving.
- `04-code/golang/pkg/applogger/rotating_file_sink_test.go`: Unit tests verifying rotation, retention pruning, and gzip archiving.
- `04-code/golang/pkg/lazyonce/lazyonce.go`: `LazyOnce[T]` for 0-parameter initializers.
- `04-code/golang/pkg/lazyonce/lazyonce1.go`: `LazyOnce1[TInput, TOutput]` for 1-parameter initializers.
- `04-code/golang/pkg/lazyonce/lazyonce2.go`: `LazyOnce2[T1, T2, TOutput]` for 2-parameter initializers.
- `04-code/golang/pkg/lazyonce/lazyonce_test.go`: Concurrency and memoization unit tests.
- `04-code/golang/pkg/errcmd/script_builder.go`: Cross-platform `ScriptBuilder` for shell script compilation.
- `04-code/golang/pkg/errcmd/powershell.go`: Native PowerShell command builder.
- `04-code/golang/pkg/errcmd/bash.go`: Native POSIX Bash command builder.
- `04-code/golang/pkg/errcmd/safedefer.go`: Safe deferral utilities (`SafeClose`, `SafeRecover`).
- `04-code/golang/pkg/errcmd/command_runner.go`: `CommandRunner` with integrated task DB & rotating file logging.
- `04-code/golang/pkg/errcmd/command_logger.go`: High-level convenience functions for automated command execution.
- `04-code/golang/pkg/errcmd/errcmd_test.go`: Unit tests for `ScriptBuilder`, `CommandRunner`, and safe defer.

### Modifications & Examples

- `04-code/golang/pkg/applogger/config.go`: Added `Rotation` config and `DriverRotatingFile` support.
- `04-code/golang/pkg/applogger/driver_type.go`: Added `DriverRotatingFile` enum variant.
- `04-code/golang/examples/split_sqlite_and_errcmd_examples.go`: Production-grade usage examples for all new systems.
- `04-code/golang/examples/split_sqlite_and_errcmd_examples_test.go`: Automated tests for all example flows.

### Plans & Tracking

- `.lovable/plans/completed/08-split-sqlite-logging-rotating-lazyonce-errcmd.md`: Master architectural plan.
- `.lovable/plans/subtasks/08-split-sqlite-logging-rotating-lazyonce-errcmd/01..06`: Granular subtasks.
- `.lovable/plans/01-index.md`: Registered completed plan 08.

---

## 4. Architectural Decisions & Rationale

- **Pluggable `DBOpenerFunc`:** Allows runtime injection of pure Go SQLite drivers (`modernc.org/sqlite`), CGo drivers (`mattn/go-sqlite3`), or mock in-memory connections, making tests 100% deterministic and portable.
- **Split DB Isolation:** Dividing logging into global `logs.db` and per-task `tasks/<task-id>.db` isolates write contention, drastically improving throughput and allowing clean per-task queries and archiving.
- **Atomic File Rotation:** Rotation checks size before write and renames/archives the file while holding a lock, preventing log message interleaving or file corruption.
- **Parametric Lazy Once:** Extending `errwonce` with generic type parameters (`T`, `TInput`, `T1`, `T2`) allows arbitrary return types wrapped in `result.Result[T]` or `*appfault.AppError` without type assertions or cycle issues.

---

## 5. Verification & Quality Gate Results

- `go test ./pkg/... ./examples/... -count=1`: **PASS** (28/28 packages green).
- `python linter-scripts/check-relative-paths.py`: **PASS** (0 absolute paths across tracked files).
- `python linter-scripts/check-sequence-integrity.py`: **PASS** (142 documents audited, 0 broken references).
- `node scripts/sync-check.mjs`: **PASS** (all 4 sync-managed files up to date).
- `node linter-scripts/check-newline-styling.mjs`: **PASS** (all files compliant).
- `python 03-ai-scripts/26-go-code-formatter.py`: **PASS** (260 Go files verified/formatted).
- `python 03-ai-scripts/31-md-gap-fixer.py`: **PASS** (1,163 files clean).
- `python 03-ai-scripts/06-cicd-local-runner.py`: **PASS** (36/36 quality gates green in 18.18s).
