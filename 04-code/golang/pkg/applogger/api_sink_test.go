package applogger_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/applogger"
	"coding-guidelines/common/pkg/enum/logleveltype"
)

type mockSenderRecord struct {
	endpoint string
	headers  map[string]string
	payload  []byte
}

func TestApiSink_BatchRotation(t *testing.T) {
	var sentBatches [][]applogger.LogEntry
	var mu sync.Mutex

	mockSender := func(
		ctx context.Context,
		endpoint string,
		headers map[string]string,
		payload []byte,
	) *appfault.AppError {
		var entries []applogger.LogEntry
		_ = json.Unmarshal(payload, &entries)
		mu.Lock()
		sentBatches = append(sentBatches, entries)
		mu.Unlock()

		return nil
	}

	sink, _ := applogger.NewApiSink(applogger.ApiConfig{
		Endpoint:      "https://example.com/logs",
		BatchSize:     3,
		FlushInterval: 0,
		Sender:        mockSender,
	})
	defer sink.Close()

	for i := 0; i < 5; i++ {
		_ = sink.WriteEntry(applogger.LogEntry{Message: "msg"})
	}

	if len(sentBatches) != 1 || len(sentBatches[0]) != 3 {
		t.Fatalf("expected 1 batch of 3 entries, got %v", sentBatches)
	}

	_ = sink.Sync()
	if len(sentBatches) != 2 || len(sentBatches[1]) != 2 {
		t.Fatalf("expected 2nd batch with 2 entries after sync, got %v", sentBatches)
	}
}

func TestApiSink_CustomRotationPolicy(t *testing.T) {
	flushCount := 0
	policy := func(batch []applogger.LogEntry, elapsed time.Duration, cfg applogger.ApiConfig) bool {
		for _, e := range batch {
			if e.Level == logleveltype.Fatal {
				return true
			}
		}

		return false
	}

	sink, _ := applogger.NewApiSink(applogger.ApiConfig{
		BatchSize:      100,
		RotationPolicy: policy,
		Sender: func(ctx context.Context, endpoint string, h map[string]string, p []byte) *appfault.AppError {
			flushCount++

			return nil
		},
	})
	defer sink.Close()

	_ = sink.WriteEntry(applogger.LogEntry{Level: logleveltype.Info, Message: "normal"})
	if flushCount != 0 {
		t.Fatalf("expected 0 flushes, got %d", flushCount)
	}

	_ = sink.WriteEntry(applogger.LogEntry{Level: logleveltype.Fatal, Message: "emergency"})
	if flushCount != 1 {
		t.Fatalf("expected 1 flush on fatal log, got %d", flushCount)
	}
}

func TestApiSink_SenderOverrideAndHeaders(t *testing.T) {
	sink, _ := applogger.NewApiSink(applogger.ApiConfig{
		Endpoint: "http://api.local/logs",
		Headers:  map[string]string{"X-Auth": "token-123"},
	})
	defer sink.Close()

	capturedHeader := ""
	sink.SetSender(func(ctx context.Context, ep string, h map[string]string, p []byte) *appfault.AppError {
		capturedHeader = h["X-Auth"]

		return nil
	})

	_ = sink.WriteEntry(applogger.LogEntry{Message: "override-test"})
	_ = sink.Flush()

	if capturedHeader != "token-123" {
		t.Fatalf("expected X-Auth header token-123, got %s", capturedHeader)
	}
}

func TestApiSink_DefaultHttpSender(t *testing.T) {
	received := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received = true
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	ctx := context.Background()
	fault := applogger.DefaultHttpSender(ctx, server.URL, map[string]string{"K": "V"}, []byte(`[]`))
	if fault != nil {
		t.Fatalf("default http sender failed: %s", fault.Message())
	}

	if !received {
		t.Fatal("server did not receive payload")
	}
}

func TestApiManager_AliasConstructor(t *testing.T) {
	mgr, err := applogger.NewApiManager(applogger.ApiConfig{
		Endpoint: "http://localhost:8080",
	})
	if err != nil || mgr == nil {
		t.Fatalf("failed to create api manager: %v", err)
	}

	defer mgr.Close()
	if mgr.BufferedCount() != 0 {
		t.Fatalf("expected 0 buffered entries")
	}
}
