# Master Plan: Task Retention, errcmd Streaming, Atomic File Writes, LazyOnce Context, and API Manager

## 1. Header & Metadata
- **Plan ID:** 12-task-retention-errcmd-streaming-atomic-fileutil-lazyonce-apimanager
- **Date:** 2026-09-09
- **Author/Agent:** Antigravity Master Orchestrator
- **Status:** ✅ Completed
- **Affected Packages:**
  - `04-code/golang/pkg/applogger`
  - `04-code/golang/pkg/applogger/sqlitelogger`
  - `04-code/golang/pkg/errcmd`
  - `04-code/golang/pkg/fileutil`
  - `04-code/golang/pkg/lazyonce`
  - `04-code/golang/examples`

---

## 2. Architectural Objectives

1. **Option 1: Task Database Retention & Dynamic SQL Query Filtering (`sqlitelogger`)**:
   - Add `PruneTasks(maxAge time.Duration) (int, *appfault.AppError)` to prune task database files older than `maxAge`.
   - Add `PruneTaskCount(maxDbs int) (int, *appfault.AppError)` to limit total accumulated task databases, removing oldest first.
   - Implement dynamic `WHERE` filtering in `queryLogs` matching `FilterOptions` fields: `level`, `startTime`, `endTime`, `limit`, and `offset`.
2. **Option 2: Real-time Live Line Streaming & Process Context (`errcmd`)**:
   - Add `WithStdoutHandler(func(line string))` and `WithStderrHandler(func(line string))` on `CommandRunner`.
   - Add `WithEnv(env map[string]string)` and `WithCwd(dir string)` on `CommandRunner`.
   - As commands execute, emit lines in real-time to handlers while simultaneously accumulating output for `CommandResult` and task log telemetry.
3. **Option 3: Atomic File Writes (`fileutil`)**:
   - Implement `AtomicWriteFile(path string, data []byte, perm filepermtype.Type) *appfault.AppError`.
   - Write to a sibling temporary file (`.<filename>.<pid>.<timestamp>.tmp`), flush/sync to disk, and execute atomic rename (`os.Rename`), safely handling Windows overwrite semantics.
4. **Option 4: Cache Invalidation & Context Cancellation (`lazyonce`)**:
   - Add `Reset()` to `LazyOnce[T]`, `LazyOnce1`, and `LazyOnce2` to clear memoized state and permit re-initialization.
   - Add `ValueContext(ctx context.Context)` on `LazyOnce[T]`, respecting context cancellation/timeout during initialization or lock contention.
5. **Option 5: Rotating Logger E2E Tests & Extensible ApiManager for Remote API Logging**:
   - Add end-to-end integration tests verifying file rotation, threshold rollover, backup retention limits, archive directory placement, and gzip compression.
   - Implement `ApiManager` / `ApiSink` implementing `LogSink`, supporting configurable endpoint URL, HTTP headers, batch sizing, flush intervals, pluggable `ApiLogSenderFunc`, and full `*appfault.AppError` integration.

---

## 3. Task-Specific Rules & Constraints

1. **Strict Function Length:** Every function body MUST remain <= 15 lines. Extract modular helpers for query construction, file cleanup, and stream parsing.
2. **Strict Relative Git Paths:** No absolute paths or `file:///` URIs anywhere in plans, code, or documentation.
3. **Positive Boolean Naming:** All booleans use `is` or `has` prefixes (e.g. `isFlushed`, `hasFilters`). Positive implicit checks only; zero `== true` comparisons.
4. **Error Return Standard:** All fallible operations return structured `*appfault.AppError`.
5. **Zero CI/CD Disabling:** Never disable or comment out quality gates. All 36 gates must pass on every loop.

---

## 4. Subtask Decomposition

- [01-task-retention-and-query-filtering.md](.lovable/plans/subtasks/12-task-retention-errcmd-streaming-atomic-fileutil-lazyonce-apimanager/01-task-retention-and-query-filtering.md): Task retention pruning and dynamic query filtering in `sqlitelogger`.
- [02-errcmd-streaming-and-process-context.md](.lovable/plans/subtasks/12-task-retention-errcmd-streaming-atomic-fileutil-lazyonce-apimanager/02-errcmd-streaming-and-process-context.md): Live stdout/stderr line streaming and env/cwd configuration in `errcmd`.
- [03-fileutil-atomic-write.md](.lovable/plans/subtasks/12-task-retention-errcmd-streaming-atomic-fileutil-lazyonce-apimanager/03-fileutil-atomic-write.md): Safe atomic file writing in `fileutil`.
- [04-lazyonce-reset-and-context.md](.lovable/plans/subtasks/12-task-retention-errcmd-streaming-atomic-fileutil-lazyonce-apimanager/04-lazyonce-reset-and-context.md): Cache reset and context-aware evaluation in `lazyonce`.
- [05-logger-e2e-and-api-manager.md](.lovable/plans/subtasks/12-task-retention-errcmd-streaming-atomic-fileutil-lazyonce-apimanager/05-logger-e2e-and-api-manager.md): E2E rotation testing and `ApiManager` for remote endpoint logging.
- [06-verification-and-ci-runner.md](.lovable/plans/subtasks/12-task-retention-errcmd-streaming-atomic-fileutil-lazyonce-apimanager/06-verification-and-ci-runner.md): Quality gate verification, formatter, and transaction logging.
