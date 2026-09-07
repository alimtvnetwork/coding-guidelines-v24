# Subtask 29.3: Rename `mu` in Logger, Writer, Fault, and Regex Packages

## Context
Standardize concurrency fields in `applogger`, `appwriter`, `appfault`, `appfaults`, and `regexnew`.

## Target Files
1. `04-code/golang/pkg/applogger/composite_sink.go`: `CompositeSink.mu` -> `CompositeSink.lock`
2. `04-code/golang/pkg/applogger/console_sink.go`: `ConsoleSink.mu` -> `ConsoleSink.lock`
3. `04-code/golang/pkg/applogger/file_sink.go`: `FileSink.mu` -> `FileSink.lock`
4. `04-code/golang/pkg/applogger/sqlite_sink.go`: `SqliteSink.mu` -> `SqliteSink.lock`
5. `04-code/golang/pkg/appwriter/base_writer.go`: `BaseWriter.mu` -> `BaseWriter.lock`
6. `04-code/golang/pkg/appfault/display.go`: `globalFaultWriterMu` -> `globalFaultWriterLock`
7. `04-code/golang/pkg/appfaults/mutex_collection.go`: Add `LockCollection` and alias `MutexCollection = LockCollection`
8. `04-code/golang/pkg/regexnew/vars.go` & `compile_helpers.go`: `regexMutex` -> `regexLock`

## Verification
- `go test -v -C 04-code/golang -count=1 ./pkg/applogger/... ./pkg/appwriter/... ./pkg/appfault/... ./pkg/appfaults/... ./pkg/regexnew/...`
