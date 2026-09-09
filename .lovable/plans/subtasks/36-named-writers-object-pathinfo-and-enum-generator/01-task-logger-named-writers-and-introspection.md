# Subtask 36.1: Logger Named Writers and Introspection

## Status
- **State:** Complete
- **Assigned Files:**
  - `04-code/golang/pkg/applogger/interfaces.go`
  - `04-code/golang/pkg/applogger/file_sink.go`
  - `04-code/golang/pkg/applogger/rotating_file_sink.go`
  - `04-code/golang/pkg/applogger/api_sink.go`
  - `04-code/golang/pkg/applogger/console_sink.go`
  - `04-code/golang/pkg/applogger/sqlite_sink.go`
  - `04-code/golang/pkg/applogger/zap_adapter.go`
  - `04-code/golang/pkg/applogger/composite_sink.go`
  - `04-code/golang/pkg/applogger/streamer_sink.go`
  - `04-code/golang/pkg/applogger/logger.go`
  - `04-code/golang/pkg/applogger/applogger_test.go`
  - `04-code/golang/examples/split_sqlite_and_errcmd_examples.go`

## Acceptance Criteria
1. Add `Name() string` to `LogSinker` interface in `04-code/golang/pkg/applogger/interfaces.go`.
2. Implement `Name() string` on every concrete sink:
   - `ConsoleSink`: returns `"console"`
   - `FileSink`: returns `"file"` (or `"file:" + s.filePath`)
   - `RotatingFileSink`: returns `"rotating_file"` (or `"rotating_file:" + s.cfg.FilePath`)
   - `ApiSink`: returns `"api"` (or `"api:" + s.cfg.Endpoint`)
   - `SQLiteSink`: returns `"sqlite"` (or `"sqlite:" + s.dbPath`)
   - `ZapAdapter`: returns `"zap"`
   - `CompositeSink`: returns `"composite"`
   - `StreamerSink`: returns `"streamer"`
3. Add `Writers() []LogSink`, `WriterNames() []string`, and `Streamers() []any` to `Logger` interface.
4. Add interface definitions ending in `er`: `WritersProvider`, `WriterNamesProvider`, `StreamersProvider`.
5. In `logger.go`, implement:
   - `Writers() []LogSink`: returns flat slice of sinks using `extractBaseSinks(l.sink)`
   - `WriterNames() []string`: returns string slice of `Name()` from all writers
   - `Streamers() []any`: returns slice of streamer objects attached to logger
6. Update `applogger_test.go` and `examples/split_sqlite_and_errcmd_examples.go` with complete test coverage.
7. Verify package compiles and tests pass: `go test ./pkg/applogger/... -v`.
