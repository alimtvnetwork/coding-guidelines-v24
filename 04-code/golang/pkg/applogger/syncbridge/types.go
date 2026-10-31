package syncbridge

import (
	"context"
	"sync"
	"time"

	"coding-guidelines/common/pkg/applogger"
)

// BridgeConfig configures the asynchronous multi-sink buffer and worker pool.
type BridgeConfig struct {
	QueueCapacity int
	WorkerCount   int
	FlushTimeout  time.Duration
	IsDropOnFull  bool
}

// EntryWithContext binds a LogEntry with its originating context and optional done notification.
type EntryWithContext struct {
	Context context.Context
	Entry   applogger.LogEntry
	Done    chan struct{}
}

// MultiSinkBridge multiplexes and asynchronously dispatches log entries across sinks.
type MultiSinkBridge struct {
	cfg      BridgeConfig
	sinks    []applogger.LogSink
	queue    chan EntryWithContext
	done     chan struct{}
	wg       sync.WaitGroup
	lock     sync.RWMutex
	isClosed bool
}

// DefaultBridgeConfig returns production defaults for BridgeConfig.
func DefaultBridgeConfig() BridgeConfig {
	return BridgeConfig{
		QueueCapacity: 256,
		WorkerCount:   2,
		FlushTimeout:  3 * time.Second,
		IsDropOnFull:  false,
	}
}

// Config returns the configuration of the bridge.
func (b *MultiSinkBridge) Config() BridgeConfig {
	return b.cfg
}

// Sinks returns a cloned slice of registered sinks.
func (b *MultiSinkBridge) Sinks() []applogger.LogSink {
	b.lock.RLock()
	defer b.lock.RUnlock()

	cloned := make([]applogger.LogSink, len(b.sinks))
	copy(cloned, b.sinks)

	return cloned
}

// IsClosed returns true if the bridge has been closed.
func (b *MultiSinkBridge) IsClosed() bool {
	b.lock.RLock()
	defer b.lock.RUnlock()

	return b.isClosed
}

// QueueLen returns the current number of buffered items in the queue.
func (b *MultiSinkBridge) QueueLen() int {
	return len(b.queue)
}
