# Plan 37: Typed Streamer Objects for AppLogger Introspection

## 1. Executive Summary

This plan addresses the direct user requirement:
"streamers needs to return it's own objects please" (referencing `StreamersProvider interface { Streamers() []any }`).

Instead of returning untyped `[]any`, `StreamersProvider` and `Logger` will return a typed slice of streamer objects (`[]Streamer` / `[]LogStreamer`).

### Core Deliverables
1. **`LogStreamer` & `Streamer` Interface Family:**
   - Define `LogStreamer` interface in `04-code/golang/pkg/applogger/interfaces.go` ending in mandatory `er` suffix.
   - Provide aliases `Streamer = LogStreamer` and `LogStream = LogStreamer`.
   - Contract methods: `Name() string`, `StreamEntry(entry LogEntry) error`, `Stream(ctx context.Context, payload any) *appfault.AppError`, `Destination() io.Writer`, `Sync() error`, `Close() error`, `Streamer() any`.
2. **`StreamerSink` Enhancement:**
   - Implement `LogStreamer` methods on `*StreamerSink` in `04-code/golang/pkg/applogger/streamer_sink.go`.
   - Add compile-time interface assertions: `var _ LogStreamer = (*StreamerSink)(nil)` and `var _ Streamer = (*StreamerSink)(nil)`.
3. **`StreamersProvider` & `Logger` Update:**
   - Update `StreamersProvider` to `Streamers() []Streamer`.
   - Update `Logger` interface to `Streamers() []Streamer`.
   - Update `extractSinkStreamers` and `(l *appLogger) Streamers()` in `04-code/golang/pkg/applogger/logger.go` to collect and return `[]Streamer`.
4. **Verification & Quality Automation:**
   - Update unit tests in `04-code/golang/pkg/applogger/applogger_test.go` and `streamer_sink_test.go`.
   - Run `go test ./...` and `python 03-ai-scripts/06-cicd-local-runner.py --all` verifying all 36 quality gates pass.

---

## 2. Task-Specific Rules

1. **Strictly Relative Git Paths:** All documentation, plans, links, and code paths MUST be strictly relative to the repository root. Zero absolute paths or `file:///` URIs.
2. **Interface Naming Mandate:** Every Go interface MUST end with the `er` suffix (e.g. `LogStreamer`, `StreamersProvider`).
3. **Hard Cap $\le 15$ Lines:** Every function or method in Go and Python must stay strictly $\le 15$ lines.
4. **Implicit Booleans:** Implicit boolean evaluations only; fields/variables prefixed with `is` or `has`.

---

## 3. Subtask Decomposition

- [x] `01-task-streamer-interface-and-streamer-sink-enhancements.md`: Define `LogStreamer` interface and implement on `*StreamerSink`.
- [x] `02-task-logger-streamers-implementation-and-verification.md`: Update `StreamersProvider`, `Logger`, and `appLogger.Streamers()` and verify all test suites.
