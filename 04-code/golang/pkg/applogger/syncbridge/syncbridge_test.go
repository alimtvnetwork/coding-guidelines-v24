package syncbridge

import (
	"context"
	"sync"
	"testing"
	"time"

	"coding-guidelines/common/pkg/applogger"
	"coding-guidelines/common/pkg/enum/logleveltype"
)

type mockTestSink struct {
	lock      sync.Mutex
	name      string
	entries   []applogger.LogEntry
	syncCount int
	isClosed  bool
	syncDelay time.Duration
}

func newMockTestSink(name string, syncDelay time.Duration) *mockTestSink {
	return &mockTestSink{
		name:      name,
		entries:   make([]applogger.LogEntry, 0),
		syncDelay: syncDelay,
	}
}

func (m *mockTestSink) Name() string {
	return m.name
}

func (m *mockTestSink) WriteEntry(entry applogger.LogEntry) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.entries = append(m.entries, entry)

	return nil
}

func (m *mockTestSink) Sync() error {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.syncCount++
	if m.syncDelay > 0 {
		time.Sleep(m.syncDelay)
	}

	return nil
}

func (m *mockTestSink) Close() error {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.isClosed = true

	return nil
}

func (m *mockTestSink) EntriesCount() int {
	m.lock.Lock()
	defer m.lock.Unlock()

	return len(m.entries)
}

func (m *mockTestSink) GetSyncCount() int {
	m.lock.Lock()
	defer m.lock.Unlock()

	return m.syncCount
}

func (m *mockTestSink) IsClosedState() bool {
	m.lock.Lock()
	defer m.lock.Unlock()

	return m.isClosed
}

func (m *mockTestSink) FirstEntry() applogger.LogEntry {
	m.lock.Lock()
	defer m.lock.Unlock()

	if len(m.entries) > 0 {
		return m.entries[0]
	}

	return applogger.LogEntry{}
}

func TestMultiSinkBridge_NonBlockingQueueing(t *testing.T) {
	sink := newMockTestSink("test-sink", 0)
	cfg := BridgeConfig{
		QueueCapacity: 50,
		WorkerCount:   2,
		FlushTimeout:  2 * time.Second,
		IsDropOnFull:  false,
	}
	bridge := NewMultiSinkBridge(cfg, sink)
	defer bridge.Close()

	ctx := context.Background()
	for i := 0; i < 20; i++ {
		entry := applogger.LogEntry{
			Message: "test message",
			Level:   logleveltype.Info,
		}
		if err := bridge.DispatchContext(ctx, entry); err != nil {
			t.Fatalf("unexpected dispatch error: %v", err)
		}
	}

	if err := bridge.Flush(ctx); err != nil {
		t.Fatalf("flush failed: %v", err)
	}

	if count := sink.EntriesCount(); count != 20 {
		t.Fatalf("expected 20 entries, got %d", count)
	}
}

func TestMultiSinkBridge_DropOnFull(t *testing.T) {
	sink := newMockTestSink("drop-sink", 0)
	cfg := BridgeConfig{
		QueueCapacity: 2,
		WorkerCount:   1,
		FlushTimeout:  2 * time.Second,
		IsDropOnFull:  true,
	}
	bridge := NewMultiSinkBridge(cfg, sink)
	defer bridge.Close()

	ctx := context.Background()
	for i := 0; i < 10; i++ {
		entry := applogger.LogEntry{
			Message: "drop message",
			Level:   logleveltype.Debug,
		}
		_ = bridge.DispatchContext(ctx, entry)
	}

	_ = bridge.Flush(ctx)
}

func TestMultiSinkBridge_TraceEnrichment(t *testing.T) {
	sink := newMockTestSink("trace-sink", 0)
	cfg := DefaultBridgeConfig()
	bridge := NewMultiSinkBridge(cfg, sink)
	defer bridge.Close()

	ctx := context.WithValue(context.Background(), "traceId", "trace-xyz-789")
	ctx = context.WithValue(ctx, "spanId", "span-abc-123")

	entry := applogger.LogEntry{
		Message: "traced operation",
		Level:   logleveltype.Info,
	}

	if err := bridge.DispatchContext(ctx, entry); err != nil {
		t.Fatalf("dispatch error: %v", err)
	}

	if err := bridge.Flush(context.Background()); err != nil {
		t.Fatalf("flush error: %v", err)
	}

	first := sink.FirstEntry()
	if first.Fields == nil {
		t.Fatalf("expected fields to be populated")
	}

	if val := first.Fields.GetString("traceId"); val != "trace-xyz-789" {
		t.Errorf("expected traceId 'trace-xyz-789', got %q", val)
	}

	if val := first.Fields.GetString("spanId"); val != "span-abc-123" {
		t.Errorf("expected spanId 'span-abc-123', got %q", val)
	}
}

func TestMultiSinkBridge_ConcurrentSinkFlushing(t *testing.T) {
	delay := 30 * time.Millisecond
	s1 := newMockTestSink("sink-1", delay)
	s2 := newMockTestSink("sink-2", delay)
	s3 := newMockTestSink("sink-3", delay)

	cfg := DefaultBridgeConfig()
	bridge := NewMultiSinkBridge(cfg, s1, s2, s3)
	defer bridge.Close()

	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := bridge.Flush(ctx); err != nil {
		t.Fatalf("flush failed: %v", err)
	}

	elapsed := time.Since(start)
	if elapsed > 75*time.Millisecond {
		t.Errorf("expected concurrent flush under 75ms, took %v", elapsed)
	}

	if s1.GetSyncCount() != 1 || s2.GetSyncCount() != 1 || s3.GetSyncCount() != 1 {
		t.Errorf("expected all sinks to sync once")
	}
}

func TestMultiSinkBridge_CloseGraceful(t *testing.T) {
	s1 := newMockTestSink("sink-1", 0)
	cfg := DefaultBridgeConfig()
	bridge := NewMultiSinkBridge(cfg, s1)

	entry := applogger.LogEntry{
		Message: "pre-close message",
		Level:   logleveltype.Info,
	}

	_ = bridge.DispatchContext(context.Background(), entry)

	if err := bridge.Close(); err != nil {
		t.Fatalf("close error: %v", err)
	}

	if isClosed := s1.IsClosedState(); isClosed {
		// correctly closed
	} else {
		t.Errorf("expected sink to be closed")
	}

	if isClosed := bridge.IsClosed(); isClosed {
		// correctly closed
	} else {
		t.Errorf("expected bridge to be closed")
	}

	if err := bridge.DispatchContext(context.Background(), entry); err != nil {
		// expected error
	} else {
		t.Errorf("expected error on dispatching to closed bridge")
	}
}
