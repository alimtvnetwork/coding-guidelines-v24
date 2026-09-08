# Task 24: Task Retention, errcmd Streaming, Atomic Fileutil, lazyonce Context, and ApiManager

## 1. Header & Metadata
- **Date:** 2026-09-09
- **Author/Agent:** Antigravity Master Orchestrator
- **Status:** Completed
- **Affected Packages:**
  - `coding-guidelines/common/pkg/applogger/sqlitelogger`
  - `coding-guidelines/common/pkg/errcmd`
  - `coding-guidelines/common/pkg/fileutil`
  - `coding-guidelines/common/pkg/lazyonce`
  - `coding-guidelines/common/pkg/applogger`

---

## 2. Context & Goals
The user requested five key architectural enhancements and end-to-end integration tests:
1. **Option 1 (sqlitelogger):** Task database retention pruning (`PruneTasks(maxAge)`, `PruneTaskCount(maxDbs)`) and dynamic query filtering (`FilterOptions` with level, start/end time, limit, offset).
2. **Option 2 (errcmd):** Live line streaming callbacks (`WithStdoutHandler`, `WithStderrHandler`) and process context (`WithEnv`, `WithCwd`).
3. **Option 3 (fileutil):** Safe atomic file writing (`AtomicWriteFile`, `AtomicWrite`) via temp file, disk sync, and atomic rename.
4. **Option 4 (lazyonce):** Cache invalidation (`Reset()`) and context-aware evaluation (`ValueContext(ctx)`, `ResultContext(ctx)`) across zero, single, and two-parameter memoizers.
5. **Option 5 (applogger):** End-to-end integration test for `RotatingFileSink` (rotation, gzip compression, backup retention) and an extensible `ApiManager` / `ApiSink` for remote HTTP log transmission with pluggable sender and rotation policies.
6. **Specifications:** Formal architectural specifications authored in `02-spec/05-split-db-architecture/02-features/`.

---

## 3. Files Changed / Created

### New Source & Test Files
- `04-code/golang/pkg/fileutil/atomic_write.go`:
  - Implements `AtomicWriteFile(filePath, data, perm)` and `AtomicWrite(filePath, data, perm)`.
- `04-code/golang/pkg/fileutil/atomic_write_test.go`:
  - Unit tests verifying atomic write, overwrites, directory creation, and error handling.
- `04-code/golang/pkg/applogger/api_sink.go`:
  - Implements `ApiSink` / `ApiManager` (implements `LogSink`) with `ApiConfig`, `DefaultHttpSender`, and extensible `ApiRotationPolicyFunc`.
- `04-code/golang/pkg/applogger/api_sink_test.go`:
  - Unit tests covering batch rotation, custom rotation policies, sender overrides, HTTP sender, and aliases.
- `04-code/golang/pkg/applogger/rotating_file_sink_e2e_test.go`:
  - End-to-end integration test writing realistic multiline logs, verifying gzip compression and backup retention pruning.
- `02-spec/05-split-db-architecture/02-features/07-task-retention-and-query-filtering.md`:
  - Specification for task retention pruning and dynamic query filtering.
- `02-spec/05-split-db-architecture/02-features/08-api-manager-and-remote-logging.md`:
  - Specification for remote API logging and rotation policies.

### Modified Files
- `04-code/golang/pkg/applogger/sqlitelogger/manager.go`:
  - Added `PruneTasks(maxAge)` and `PruneTaskCount(maxDbs)`.
  - Added dynamic query filtering in `QueryMainLogs` and `QueryTaskLogs`.
- `04-code/golang/pkg/applogger/sqlitelogger/sqlitelogger_test.go`:
  - Added unit tests for pruning by age, pruning by count, and filtered queries.
- `04-code/golang/pkg/errcmd/command_runner.go`:
  - Added `WithStdoutHandler`, `WithStderrHandler`, `WithEnv`, and `WithCwd`.
  - Added real-time `lineStreamWriter` for synchronous line-by-line streaming.
- `04-code/golang/pkg/errcmd/errcmd_test.go`:
  - Added unit tests for stdout streaming, stderr streaming, environment injection, and working directory context.
- `04-code/golang/pkg/lazyonce/lazyonce.go`:
  - Added `ValueContext(ctx)` and `ResultContext(ctx)`.
- `04-code/golang/pkg/lazyonce/lazyonce1.go`:
  - Added `ValueContext(ctx, input)` and `ResultContext(ctx, input)`.
- `04-code/golang/pkg/lazyonce/lazyonce2.go`:
  - Added `ValueContext(ctx, arg1, arg2)` and `ResultContext(ctx, arg1, arg2)`.
- `04-code/golang/pkg/lazyonce/lazyonce_test.go`:
  - Added unit tests for context cancellation, timeout fault generation, and reset verification.

---

## 4. Architectural Decisions & Rationale
1. **Synchronous Real-Time Line Streaming in errcmd:**
   Implemented `lineStreamWriter` wrapping `bytes.Buffer` and line splitting. This avoids complex pipe lifecycle issues and goroutine race conditions while providing immediate line dispatch to handlers and task loggers.
2. **Safe Atomic Replacement in fileutil:**
   Writing to a sibling temporary file followed by `f.Sync()` guarantees data reaches persistent storage before `os.Rename` replaces the target file.
3. **Context-Aware Lazy Evaluation:**
   `awaitContext` combines non-blocking channel evaluation with `select` on `ctx.Done()`, ensuring callers immediately unblock if a deadline expires while keeping the memoized result intact upon eventual completion.
4. **Pluggable Remote Logging via ApiSink:**
   By parameterizing `ApiLogSenderFunc` and `ApiRotationPolicyFunc`, `ApiSink` supports testing, custom HTTP clients, gRPC bridges, and custom batch flush policies without modifying core sink logic.
5. **Guideline Adherence:**
   All function bodies remain <= 15 lines, booleans use positive prefixes, all fallible functions return `*appfault.AppError`, and all file links use strict relative git paths.
