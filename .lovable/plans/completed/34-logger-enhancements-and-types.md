# Plan 34: Logger Type Taxonomy & Extended Interface Methods

> **Status:** Complete  
> **Type:** Feature & Architecture Enhancement  
> **Related Issue / User Request:** Add logger types (Api, FileWriter, FileWriterRotator, JsonWriterLogger, SqliteDbWriter), FilePath, EndPointPath, Clone(), AddWriters(), AddStreamer() to Logger interface and implementations.  
> **Workflow:** Parent Task N-Step Continuous Loop & Multi-Agent Orchestration ($N = 100$)  

---

## 1. Architectural Overview & Objectives

This plan enhances the Go structured logging engine in `04-code/golang/pkg/applogger/` to provide:
1. **Extended Driver Taxonomy (`driver_type.go`):** Add `DriverApi`, `DriverJsonWriterLogger`, `DriverStreamer`, and comprehensive aliases (`DriverFileWriter`, `DriverFileWriterRotator`, `DriverSqliteDbWriter`, `DriverAPI`, `DriverJson`). Implement string normalization, parsing (`ParseDriverType`), and JSON serialization.
2. **Sink Introspection Protocol:** Equip sinks (`FileSink`, `RotatingFileSink`, `ApiSink`, `CompositeSink`, `ConsoleSink`, `SQLiteSink`, `ZapAdapter`) with accessor methods for active paths (`FilePath()`, `EndpointPath()`, `EndPointPath()`) and driver identification (`DriverType()`).
3. **Universal Streamer Adapter (`streamer_sink.go`):** Create `StreamerSink` bridging `streamwriter.Streamer`, `io.Writer`, and arbitrary stream receivers directly into the `LogSink` interface.
4. **Fluent Logger Chaining & Immutability (`logger.go`, `interfaces.go`, `config.go`):** Update `Logger` interface and `appLogger` with `FilePath()`, `EndpointPath()`, `EndPointPath()`, `Clone()`, `AddWriters()`, `AddStreamer()`, and `Type()`. Maintain immutability so child loggers and goroutines remain safe from mutations.

---

## 2. Task-Specific Non-Negotiable Rules

1. **Immutable Chaining:** `Clone()`, `AddWriters()`, `AddStreamer()`, `WithContext()`, and `WithFields()` MUST return fresh copies of `*appLogger` without mutating the parent instance in-place.
2. **Accessor Duck-Typing & Aliasing:** Provide both `EndpointPath() string` and `EndPointPath() string` on `ApiSink`, `CompositeSink`, and `appLogger` to satisfy both camelCase and PascalCase/mixed naming preferences seamlessly.
3. **Streamer Invariance Safety:** `AddStreamer(streamer any)` must accept `any` and use duck typing to support invariant Go generic types (`Streamer[string]`, `Streamer[LogRecord]`, `AnyStreamer`, `io.Writer`) without breaking compile-time type safety.
4. **Function Line Limit:** Every new or edited function must strictly contain 15 or fewer lines of code.
5. **Strict Relative Git Paths:** No absolute paths or `file:///` URIs may appear in code, comments, or documentation.

---

## 3. Disjoint Subtask Decomposition

```
.lovable/plans/subtasks/34-logger-enhancements-and-types/
├── 01-task-driver-type-taxonomy-expansion.md
├── 02-task-concrete-sink-introspection.md
├── 03-task-streamer-sink-bridge-implementation.md
└── 04-task-logger-core-fluent-chaining-and-verification.md
```

### Subtask Summary & File Ownership Matrix

| Subtask | Focus Area | Exclusive Assigned Files |
|---|---|---|
| **34.1** | Driver Enum Expansion & Results | `04-code/golang/pkg/applogger/driver_type.go`<br>`04-code/golang/pkg/applogger/interfaces.go`<br>`04-code/golang/pkg/applogger/types.go`<br>`04-code/golang/pkg/applogger/results.go`<br>`04-code/golang/pkg/applogger/driver_type_test.go`<br>`04-code/golang/pkg/applogger/results_test.go` |
| **34.2** | Concrete Sink Introspection | `04-code/golang/pkg/applogger/file_sink.go`<br>`04-code/golang/pkg/applogger/rotating_file_sink.go`<br>`04-code/golang/pkg/applogger/api_sink.go`<br>`04-code/golang/pkg/applogger/console_sink.go`<br>`04-code/golang/pkg/applogger/sqlite_sink.go`<br>`04-code/golang/pkg/applogger/zap_adapter.go`<br>`04-code/golang/pkg/applogger/composite_sink.go` |
| **34.3** | Universal Streamer Sink Bridge | `04-code/golang/pkg/applogger/streamer_sink.go`<br>`04-code/golang/pkg/applogger/streamer_sink_test.go` |
| **34.4** | Logger Core, Config, Examples & Gates | `04-code/golang/pkg/applogger/logger.go`<br>`04-code/golang/pkg/applogger/logger_enrichment.go`<br>`04-code/golang/pkg/applogger/config.go`<br>`04-code/golang/pkg/applogger/applogger_test.go`<br>`04-code/golang/examples/split_sqlite_and_errcmd_examples.go` |

---

## 4. Verification & Quality Gates

1. Package-level test suite: `go test -v ./04-code/golang/pkg/applogger/...`
2. Full Go test suite: `go test ./04-code/golang/...`
3. Go code formatter: `python 03-ai-scripts/26-go-code-formatter.py`
4. Full CI/CD runner: `python 03-ai-scripts/06-cicd-local-runner.py --all`
