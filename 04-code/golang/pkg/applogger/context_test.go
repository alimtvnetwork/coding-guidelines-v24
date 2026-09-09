package applogger

import (
	"context"
	"testing"
)

func TestFromContext(t *testing.T) {
	loggerRes := Default()
	if loggerRes.IsFailed() {
		t.Fatalf("Default() failed: %v", loggerRes.Fault())
	}

	ctx := context.WithValue(context.Background(), RequestIDKey, "req-1234")
	ctx = context.WithValue(ctx, UserIDKey, "usr-999")

	childLogger := FromContext(ctx, loggerRes.Data())
	if childLogger == nil {
		t.Fatalf("Expected non-nil child logger")
	}

	childLogger.Info("Context-aware log message")
}
