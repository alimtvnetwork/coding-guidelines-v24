package streamwriter_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
	"testing"

	"coding-guidelines/common/pkg/streamwriter"
)

// SafeBuffer wraps bytes.Buffer with mutex for test inspections
type SafeBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (s *SafeBuffer) Write(p []byte) (n int, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.buf.Write(p)
}

func (s *SafeBuffer) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.buf.String()
}

func (s *SafeBuffer) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.buf.Reset()
}

func TestLockedStreamer_ConcurrentSafe(t *testing.T) {
	buf := &SafeBuffer{}
	streamer := streamwriter.NewLockedStreamer(streamwriter.LockedOptions{
		Name:        "concurrent-test",
		Destination: buf,
	})

	if !streamer.IsLocked() {
		t.Fatalf("expected IsLocked() to be true")
	}

	var wg sync.WaitGroup
	ctx := context.Background()
	concurrency := 25

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			err := streamer.Stream(ctx, fmt.Sprintf("goroutine-%d", idx))
			if err != nil {
				t.Errorf("stream failed: %v", err)
			}
		}(i)
	}
	wg.Wait()

	out := buf.String()
	for i := 0; i < concurrency; i++ {
		expected := fmt.Sprintf("goroutine-%d", i)
		if !strings.Contains(out, expected) {
			t.Errorf("missing expected output: %s", expected)
		}
	}
}

func TestLocklessStreamer_Direct(t *testing.T) {
	buf := &bytes.Buffer{}
	streamer := streamwriter.NewLocklessStreamer(streamwriter.LocklessOptions{
		Name:        "cli-test",
		Destination: buf,
	})

	if streamer.IsLocked() {
		t.Fatalf("expected IsLocked() to be false")
	}

	ctx := context.Background()
	err := streamer.Stream(ctx, "single-thread event")
	if err != nil {
		t.Fatalf("stream failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "single-thread event") {
		t.Fatalf("expected output to contain payload, got: %s", out)
	}
	if !strings.Contains(out, "[lockless]") {
		t.Fatalf("expected [lockless] tag in output, got: %s", out)
	}
}

func TestSelfBinding_Contracts(t *testing.T) {
	locked := streamwriter.NewLockedStreamer(streamwriter.LockedOptions{Name: "test-locked"})
	lockless := streamwriter.NewLocklessStreamer(streamwriter.LocklessOptions{Name: "test-lockless"})
	writer := streamwriter.NewPluggableWriter(streamwriter.WriterOptions{Name: "test-writer", Streamer: locked})

	// Verify LockedStreamer self-binding
	var s1 streamwriter.StreamerInterface = locked.AsStreamer()
	var w1 streamwriter.WriterInterface = locked.AsWriter()
	var i1 streamwriter.Interfacer = locked.AsInterfacer()
	if s1 == nil || w1 == nil || i1 == nil {
		t.Fatal("locked streamer self-binding failed")
	}

	// Verify LocklessStreamer self-binding
	var s2 streamwriter.StreamerInterface = lockless.AsStreamer()
	var w2 streamwriter.WriterInterface = lockless.AsWriter()
	var i2 streamwriter.Interfacer = lockless.AsInterfacer()
	if s2 == nil || w2 == nil || i2 == nil {
		t.Fatal("lockless streamer self-binding failed")
	}

	// Verify PluggableWriter self-binding
	var w3 streamwriter.WriterInterface = writer.AsWriter()
	var i3 streamwriter.Interfacer = writer.AsInterfacer()
	if w3 == nil || i3 == nil {
		t.Fatal("pluggable writer self-binding failed")
	}
}

func TestSwappableMethods_Runtime(t *testing.T) {
	buf := &SafeBuffer{}
	streamer := streamwriter.NewLockedStreamer(streamwriter.LockedOptions{
		Name:        "swappable-test",
		Destination: buf,
	})

	ctx := context.Background()

	// Initial default stream
	_ = streamer.Stream(ctx, "initial")
	if !strings.Contains(buf.String(), "[swappable-test][locked] initial") {
		t.Fatalf("unexpected initial output: %s", buf.String())
	}

	buf.Reset()

	// Hot-swap stream method to JSON
	streamer.SetStreamMethod(func(ctx context.Context, payload any, dest io.Writer) error {
		b, _ := json.Marshal(map[string]any{"data": payload, "custom": "swapped"})
		_, err := dest.Write(append(b, '\n'))
		return err
	})

	_ = streamer.Stream(ctx, "hello-json")
	if !strings.Contains(buf.String(), `{"custom":"swapped","data":"hello-json"}`) {
		t.Fatalf("unexpected swapped output: %s", buf.String())
	}
}

func TestCompositeLogger_FluentChaining(t *testing.T) {
	buf1 := &SafeBuffer{}
	buf2 := &SafeBuffer{}
	buf3 := &SafeBuffer{}

	w1 := streamwriter.NewLockedStreamer(streamwriter.LockedOptions{Name: "w1", Destination: buf1})
	w2 := streamwriter.NewLocklessStreamer(streamwriter.LocklessOptions{Name: "w2", Destination: buf2})

	customWriter := streamwriter.NewPluggableWriter(streamwriter.WriterOptions{
		Name: "custom-api",
		WriteMethod: func(ctx context.Context, payload any) error {
			_, err := fmt.Fprintf(buf3, "CUSTOM-API: %v\n", payload)
			return err
		},
	})

	// FLUENT REGISTRATION
	log := streamwriter.NewLogger().
		AddWriters(w1, w2).
		AddWriter(customWriter)

	if log.WriterCount() != 3 {
		t.Fatalf("expected 3 writers, got %d", log.WriterCount())
	}

	ctx := context.WithValue(context.Background(), "traceId", "trace-999")
	err := log.Info(ctx, "Order placed successfully")
	if err != nil {
		t.Fatalf("log.Info failed: %v", err)
	}

	// Verify emission across all 3 destinations
	if !strings.Contains(buf1.String(), "Order placed successfully") {
		t.Errorf("w1 did not receive log")
	}
	if !strings.Contains(buf2.String(), "Order placed successfully") {
		t.Errorf("w2 did not receive log")
	}
	if !strings.Contains(buf3.String(), "CUSTOM-API: ") {
		t.Errorf("customWriter did not receive log")
	}

	// Dynamic removal
	log.RemoveWriter("custom-api")
	if log.WriterCount() != 2 {
		t.Fatalf("expected 2 writers after removal, got %d", log.WriterCount())
	}

	// Clear to silent mode
	log.ClearWriters()
	if log.WriterCount() != 0 {
		t.Fatalf("expected 0 writers after clear, got %d", log.WriterCount())
	}

	buf1.Reset()
	_ = log.Info(ctx, "Silent message")
	if buf1.String() != "" {
		t.Fatalf("expected empty buffer in silent mode, got: %s", buf1.String())
	}
}

func TestLogAndNonLogPayloads(t *testing.T) {
	buf := &SafeBuffer{}
	streamer := streamwriter.NewLockedStreamer(streamwriter.LockedOptions{
		Name:        "payload-test",
		Destination: buf,
	})

	log := streamwriter.NewLogger().AddStreamer(streamer)
	ctx := context.Background()

	// 1. Structured log record
	_ = log.Info(ctx, "Structured message")
	if !strings.Contains(buf.String(), "INFO: Structured message") {
		t.Errorf("expected structured log, got: %s", buf.String())
	}

	buf.Reset()

	// 2. Non-log arbitrary payload (raw string, map, or custom struct)
	type DomainEvent struct {
		EventID string
		Amount  float64
	}
	_ = log.Emit(ctx, DomainEvent{EventID: "evt-123", Amount: 99.50})
	if !strings.Contains(buf.String(), "evt-123") {
		t.Errorf("expected domain event payload, got: %s", buf.String())
	}
}
