# AppLogger SyncBridge

The `syncbridge` package provides an asynchronous, non-blocking multi-sink buffer and worker dispatcher for `pkg/applogger`. It connects distinct log destinations such as `ConsoleSink` and `FileSink` through a bounded memory queue, concurrent background worker routines, trace context enrichment, and parallel sink flushing.

## Architectural Overview

```
                      ┌────────────────────────────────────────┐
                      │          DispatchContext(ctx)         │
                      └───────────────────┬────────────────────┘
                                          │
                            [Trace Enrichment: traceId/spanId]
                                          │
                                          ▼
                      ┌────────────────────────────────────────┐
                      │      Bounded Channel Queue (chan)      │
                      │       (QueueCapacity / DropOnFull)     │
                      └───────────────────┬────────────────────┘
                                          │
                         ┌────────────────┴────────────────┐
                         ▼                                 ▼
                 ┌───────────────┐                 ┌───────────────┐
                 │ Worker Routine│                 │ Worker Routine│
                 └───────┬───────┘                 └───────┬───────┘
                         │                                 │
                         └────────────────┬────────────────┘
                                          │
                                   [broadcastEntry]
                                          │
                         ┌────────────────┴────────────────┐
                         ▼                                 ▼
                 ┌───────────────┐                 ┌───────────────┐
                 │  ConsoleSink  │                 │   FileSink    │
                 └───────────────┘                 └───────────────┘
```

## Key Capabilities

1. **Non-Blocking Queueing & Backpressure Control:**
   - Buffers incoming entries inside a bounded FIFO queue (`BridgeConfig.QueueCapacity`).
   - Supports `IsDropOnFull: true` for zero-blocking operations during traffic spikes, or context-bounded waiting when drops are disallowed.

2. **Context Trace Enrichment:**
   - Automatically extracts distributed tracing IDs (`traceId`, `spanId`) and request IDs from `context.Context`.
   - Injects them directly into the `applogger.LogEntry.Fields` (`appfault.ContextMap`) without manual field mapping.

3. **Concurrent Multi-Sink Synchronization:**
   - Synchronizes multiple sink destinations (`ConsoleSink`, `FileSink`, `RotatingFileSink`) in parallel using a `sync.WaitGroup` to minimize flush latency.
   - Graceful queue drainage and sink termination upon `Close()`.

4. **Zero-Void Error Management:**
   - All dispatch, flush, and shutdown methods return structured `*appfault.AppError` failures.

## Configuration & Usage

```go
package main

import (
	"context"
	"os"
	"time"

	"coding-guidelines/common/pkg/applogger"
	"coding-guidelines/common/pkg/applogger/syncbridge"
)

func main() {
	console := applogger.NewConsoleSink(os.Stdout, false)
	file := applogger.NewFileSink("logs/app.log", false)

	cfg := syncbridge.BridgeConfig{
		QueueCapacity: 512,
		WorkerCount:   2,
		FlushTimeout:  3 * time.Second,
		IsDropOnFull:  false,
	}

	bridge := syncbridge.NewMultiSinkBridge(cfg, console, file)
	defer bridge.Close()

	ctx := context.WithValue(context.Background(), "traceId", "req-trace-456")
	ctx = context.WithValue(ctx, "spanId", "span-789")

	entry := applogger.LogEntry{
		Message: "Order processing initiated",
	}

	if err := bridge.DispatchContext(ctx, entry); err != nil {
		// Handle AppError
	}

	_ = bridge.Flush(ctx)
}
```

## Related Components

- Sinks: `04-code/golang/pkg/applogger/console_sink.go`, `04-code/golang/pkg/applogger/file_sink.go`
- Interfaces: `04-code/golang/pkg/applogger/interfaces.go`
- Types: `04-code/golang/pkg/applogger/syncbridge/types.go`
- Core Engine: `04-code/golang/pkg/applogger/syncbridge/syncbridge.go`
- Ambiguity Resolution: `.ai-memory/ambiguous-questions/02-ambiguity-resolved/01-pluggable-logger-backend-and-uber-zap-migration.md`
