# Unified Pluggable Logger Architecture and Uber Zap Migration

Slug: pluggable-logger-backend-and-uber-zap-migration
Status: resolved
Raised: 2026-09-03
Resolved: 2026-10-27
Resolved By: Worker 04: AppLogger SyncBridge & Ambiguity Resolution
Related Package: `04-code/golang/pkg/applogger/syncbridge`

## Question

How should `04-code/golang/pkg/applogger` be refactored to support seamless hot-swapping between internal Formatter+Writer pipelines and external enterprise loggers like Uber Zap (`go.uber.org/zap`), while maintaining native Go `context.Context` tracing and `*appfault.AppError` integration?

## Problem & Missing Capabilities in Prior Code

1. **Coupled Sinks:** In `pkg/applogger`, each sink (`ConsoleSink`, `FileSink`, `SQLiteSink`) coupled formatting (JSON vs text string interpolation) directly with destination I/O.
2. **Missing Uber Zap Interoperability:** The existing `ZapAdapter` only supported untyped `fmt.Sprint` logging without passing through structured fields, Zap typed cores, or Go `context.Context`.
3. **No Dynamic Swapper (Log Changer):** No unified manager existed to switch the active logging engine or redirect outputs at runtime without rebuilding the application.
4. **Proprietary Context:** Used custom `appfault.ContextMap` rather than standard Go `context.Context`, making distributed tracing (`traceId`, `spanId`) difficult to propagate across HTTP handlers and gRPC boundaries.

## Options Considered

### Option A: Standard `log/slog` Handler Adapter (Idiomatic Go 1.21+)
- **Architecture:** Adopt Go's native `slog.Logger` as the foundation. Implement custom `slog.Handler` wrappers for context extraction, and use `zapslog` (or `zapcore`) as a swappable handler.
- **Pros:** Zero third-party dependencies for standard use; native Go standard library compatibility; ecosystem tools natively integrate.
- **Cons:** Uber Zap features (like zero-alloc field encoders) are adapted to `slog.Record` instead of native `zap.Field`.

### Option B: Unified Dual-Mode `LogBackend` Engine with SyncBridge (Selected)
- **Architecture:** Define a universal `LogBackend` interface and provide dual-mode execution:
  1. Lightweight zero-dependency multi-sink pipeline (`ConsoleSink`, `FileSink`, `RotatingFileSink`) coordinated via `04-code/golang/pkg/applogger/syncbridge`.
  2. Enterprise Zap adapter (`ZapAdapter`) for zero-allocation high-throughput environments.
  3. `MultiSinkBridge` handles asynchronous queue buffering, context trace extraction (`traceId`, `spanId`), concurrent sink flushing (`Flush(ctx)`), and graceful drainage (`Close()`).
- **Pros:** Full access to native Uber Zap performance when needed; lightweight zero-dependency pipeline for CLI and embedded modes; distributed tracing natively supported via `context.Context`.
- **Cons:** None in practice; dual-mode encapsulates both CLI and enterprise service tiers cleanly.

### Option C: Uber Zap as Exclusive Engine
- **Architecture:** Make `go.uber.org/zap` the sole logging engine across all services, wrapping it in an `AppError` bridge.
- **Pros:** One single logging engine; highest possible throughput.
- **Cons:** Heavy external dependency for lightweight tools, CLI scripts, and embedded packages.

## Concrete Resolution & Implementation Details

Option B has been selected and implemented in `04-code/golang/pkg/applogger/syncbridge/`:

1. **Asynchronous Multi-Sink Multiplexing (`MultiSinkBridge`):**
   - Configured via `BridgeConfig` with queue capacity, worker pool sizing, flush timeouts, and drop-on-full backpressure flags.
   - Decouples log producers from slower file or remote I/O sinks.

2. **Distributed Trace Context Propagation:**
   - In `enrichTraceContext(ctx, entry)`, standard observability fields (`traceId`, `spanId`, `requestId`) are extracted from `context.Context` and injected into `entry.Fields` (`appfault.ContextMap`).

3. **Parallel Multi-Sink Flushing:**
   - `Flush(ctx)` drains pending buffered entries and concurrently synchronizes all attached sinks using a `sync.WaitGroup` to eliminate sequential I/O latency bottlenecks.

4. **Zero-Void Error Management:**
   - Returns `*appfault.AppError` across all dispatch and lifecycle methods.

## Artifacts Created

- Types: `04-code/golang/pkg/applogger/syncbridge/types.go`
- Engine: `04-code/golang/pkg/applogger/syncbridge/syncbridge.go`
- Unit Tests: `04-code/golang/pkg/applogger/syncbridge/syncbridge_test.go`
- Architecture Documentation: `04-code/golang/pkg/applogger/syncbridge/readme.md`
