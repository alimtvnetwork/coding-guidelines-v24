package syncbridge

import (
	"context"
	"sync"
	"time"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/applogger"
	"coding-guidelines/common/pkg/errtype"
)

// NewMultiSinkBridge initializes the asynchronous bridge with configured worker pool.
func NewMultiSinkBridge(cfg BridgeConfig, sinks ...applogger.LogSink) *MultiSinkBridge {
	normCfg := normalizeBridgeConfig(cfg)
	bridge := &MultiSinkBridge{
		cfg:      normCfg,
		sinks:    sinks,
		queue:    make(chan EntryWithContext, normCfg.QueueCapacity),
		done:     make(chan struct{}),
		isClosed: false,
	}

	bridge.startWorkers()

	return bridge
}

// normalizeBridgeConfig ensures queue capacity and worker count have valid minimums.
func normalizeBridgeConfig(cfg BridgeConfig) BridgeConfig {
	if cfg.QueueCapacity <= 0 {
		cfg.QueueCapacity = 256
	}

	if cfg.WorkerCount <= 0 {
		cfg.WorkerCount = 1
	}

	if cfg.FlushTimeout <= 0 {
		cfg.FlushTimeout = 3 * time.Second
	}

	return cfg
}

// startWorkers spawns background worker routines.
func (b *MultiSinkBridge) startWorkers() {
	b.wg.Add(b.cfg.WorkerCount)
	for i := 0; i < b.cfg.WorkerCount; i++ {
		go b.workerLoop()
	}
}

// workerLoop continuously consumes entries from the queue until closed.
func (b *MultiSinkBridge) workerLoop() {
	defer b.wg.Done()

	for {
		select {
		case <-b.done:
			b.drainQueue()
			return
		case item, isOk := <-b.queue:
			if isOk {
				b.processItem(item)
			}
		}
	}
}

// drainQueue flushes remaining entries from the queue channel.
func (b *MultiSinkBridge) drainQueue() {
	for {
		select {
		case item, isOk := <-b.queue:
			if isOk {
				b.processItem(item)
				continue
			}

			return
		default:
			return
		}
	}
}

// processItem broadcasts an entry to sinks and signals completion if requested.
func (b *MultiSinkBridge) processItem(item EntryWithContext) {
	if b.hasEntryPayload(item.Entry) {
		b.broadcastEntry(item.Entry)
	}

	if item.Done != nil {
		close(item.Done)
	}
}

// hasEntryPayload determines whether a log entry contains writeable data.
func (b *MultiSinkBridge) hasEntryPayload(entry applogger.LogEntry) bool {
	if entry.Message != "" {
		return true
	}

	if entry.Timestamp != "" {
		return true
	}

	return false
}

// broadcastEntry distributes a single LogEntry across all registered sinks.
func (b *MultiSinkBridge) broadcastEntry(entry applogger.LogEntry) {
	b.lock.RLock()
	defer b.lock.RUnlock()

	for _, sink := range b.sinks {
		if sink != nil {
			_ = sink.WriteEntry(entry)
		}
	}
}

// DispatchContext asynchronously queues a log entry with enriched trace metadata.
func (b *MultiSinkBridge) DispatchContext(ctx context.Context, entry applogger.LogEntry) *appfault.AppError {
	if b.isBridgeClosed() {
		return appfault.New(errtype.Precondition, "syncbridge is closed")
	}

	safeCtx := ensureValidContext(ctx)
	enriched := enrichTraceContext(safeCtx, entry)
	item := EntryWithContext{Context: safeCtx, Entry: enriched}

	return b.enqueueItem(safeCtx, item)
}

// isBridgeClosed checks whether the bridge has been shut down.
func (b *MultiSinkBridge) isBridgeClosed() bool {
	b.lock.RLock()
	defer b.lock.RUnlock()

	return b.isClosed
}

// ensureValidContext guarantees a non-nil context.
func ensureValidContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}

	return ctx
}

// enqueueItem handles queuing with drop-on-full or context cancellation support.
func (b *MultiSinkBridge) enqueueItem(ctx context.Context, item EntryWithContext) *appfault.AppError {
	if b.cfg.IsDropOnFull {
		return b.tryEnqueueDrop(item)
	}

	select {
	case b.queue <- item:
		return nil
	case <-ctx.Done():
		return appfault.New(errtype.Timeout, "context expired while queueing log entry")
	}
}

// tryEnqueueDrop attempts a non-blocking push, dropping the entry if full.
func (b *MultiSinkBridge) tryEnqueueDrop(item EntryWithContext) *appfault.AppError {
	select {
	case b.queue <- item:
		return nil
	default:
		return nil
	}
}

// enrichTraceContext extracts trace metadata from context and injects them into Fields.
func enrichTraceContext(ctx context.Context, entry applogger.LogEntry) applogger.LogEntry {
	if ctx == nil {
		return entry
	}

	if entry.Fields == nil {
		entry.Fields = appfault.NewContextMap()
	}

	injectTraceField(ctx, entry.Fields, "traceId", "traceId", "TraceId", "trace_id", "TraceID", string(applogger.TraceIDKey))
	injectTraceField(ctx, entry.Fields, "spanId", "spanId", "SpanId", "span_id", "SpanID", "spanID")

	return entry
}

// injectTraceField retrieves string from context keys and sets target key on fields.
func injectTraceField(ctx context.Context, fields appfault.ContextMap, targetKey string, keys ...string) {
	val := extractTraceValue(ctx, keys...)
	if val == "" {
		return
	}

	fields.Set(targetKey, val)
}

// extractTraceValue searches through possible context keys for a non-empty string.
func extractTraceValue(ctx context.Context, keys ...string) string {
	for _, key := range keys {
		if val, isOk := ctx.Value(key).(string); isOk && val != "" {
			return val
		}

		if val, isOk := ctx.Value(applogger.ContextKey(key)).(string); isOk && val != "" {
			return val
		}
	}

	return ""
}

// Flush flushes pending queue items and concurrently synchronizes all sinks.
func (b *MultiSinkBridge) Flush(ctx context.Context) *appfault.AppError {
	if b.isBridgeClosed() {
		return appfault.New(errtype.Precondition, "syncbridge is closed")
	}

	safeCtx := ensureValidContext(ctx)
	if err := b.flushQueue(safeCtx); err != nil {
		return err
	}

	return b.flushAllSinks(safeCtx)
}

// flushQueue submits a synchronization marker and waits for workers to reach it.
func (b *MultiSinkBridge) flushQueue(ctx context.Context) *appfault.AppError {
	marker := EntryWithContext{Done: make(chan struct{})}
	if err := b.enqueueMarker(ctx, marker); err != nil {
		return err
	}

	return b.waitForMarker(ctx, marker.Done)
}

// enqueueMarker pushes a flush marker into the queue.
func (b *MultiSinkBridge) enqueueMarker(ctx context.Context, marker EntryWithContext) *appfault.AppError {
	select {
	case b.queue <- marker:
		return nil
	case <-ctx.Done():
		return appfault.New(errtype.Timeout, "flush queue timed out on enqueue")
	}
}

// waitForMarker blocks until marker is closed or context expires.
func (b *MultiSinkBridge) waitForMarker(ctx context.Context, done chan struct{}) *appfault.AppError {
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return appfault.New(errtype.Timeout, "flush queue timed out on processing")
	}
}

// flushAllSinks concurrently calls Sync() on all registered sinks.
func (b *MultiSinkBridge) flushAllSinks(ctx context.Context) *appfault.AppError {
	b.lock.RLock()
	sinks := b.sinks
	b.lock.RUnlock()

	var wg sync.WaitGroup
	for _, s := range sinks {
		if s != nil {
			wg.Add(1)
			go b.syncSinkConcurrently(s, &wg)
		}
	}

	return b.waitForFlush(ctx, &wg)
}

// syncSinkConcurrently executes Sync on a single sink within a WaitGroup.
func (b *MultiSinkBridge) syncSinkConcurrently(s applogger.LogSink, wg *sync.WaitGroup) {
	defer wg.Done()
	_ = s.Sync()
}

// waitForFlush waits for the WaitGroup to complete or the context to expire.
func (b *MultiSinkBridge) waitForFlush(ctx context.Context, wg *sync.WaitGroup) *appfault.AppError {
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return appfault.New(errtype.Timeout, "flush timed out waiting for sinks")
	}
}

// Close gracefully drains the queue, terminates workers, and closes all sinks.
func (b *MultiSinkBridge) Close() *appfault.AppError {
	b.lock.Lock()
	if b.isClosed {
		b.lock.Unlock()
		return nil
	}

	b.isClosed = true
	close(b.done)
	b.lock.Unlock()

	b.wg.Wait()

	return b.closeSinks()
}

// closeSinks closes all registered sinks.
func (b *MultiSinkBridge) closeSinks() *appfault.AppError {
	b.lock.Lock()
	defer b.lock.Unlock()

	for _, s := range b.sinks {
		if s != nil {
			_ = s.Close()
		}
	}

	return nil
}
