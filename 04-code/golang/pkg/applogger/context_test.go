package applogger

import (
	"context"
	"testing"
)

func TestContextKey_PascalCaseValues(t *testing.T) {
	if string(RequestIDKey) != "RequestId" {
		t.Fatalf("Expected RequestId, got %s", RequestIDKey)
	}

	if string(TraceIDKey) != "TraceId" {
		t.Fatalf("Expected TraceId, got %s", TraceIDKey)
	}

	if string(UserIDKey) != "UserId" {
		t.Fatalf("Expected UserId, got %s", UserIDKey)
	}

	assertAliases(t)
}

func assertAliases(t *testing.T) {
	if RequestIdKey != RequestIDKey {
		t.Fatalf("Expected alias match for RequestIdKey")
	}

	if TraceIdKey != TraceIDKey {
		t.Fatalf("Expected alias match for TraceIdKey")
	}

	if UserIdKey != UserIDKey {
		t.Fatalf("Expected alias match for UserIdKey")
	}
}

func TestExtractContextFields_TypedKeys(t *testing.T) {
	ctx := context.WithValue(context.Background(), RequestIDKey, "req-1")
	ctx = context.WithValue(ctx, TraceIDKey, "trc-1")
	ctx = context.WithValue(ctx, UserIDKey, "usr-1")

	fields := ExtractContextFields(ctx)
	if fields["RequestId"] != "req-1" {
		t.Fatalf("Expected req-1, got %v", fields["RequestId"])
	}

	if fields["TraceId"] != "trc-1" {
		t.Fatalf("Expected trc-1, got %v", fields["TraceId"])
	}

	if fields["UserId"] != "usr-1" {
		t.Fatalf("Expected usr-1, got %v", fields["UserId"])
	}
}

func TestExtractContextFields_StringAndFallbackKeys(t *testing.T) {
	ctx := context.WithValue(context.Background(), "request_id", "req-fallback")
	ctx = context.WithValue(ctx, "TraceID", "trc-fallback")
	ctx = context.WithValue(ctx, "user_id", "usr-fallback")

	fields := ExtractContextFields(ctx)
	if fields["RequestId"] != "req-fallback" {
		t.Fatalf("Expected req-fallback, got %v", fields["RequestId"])
	}

	if fields["TraceId"] != "trc-fallback" {
		t.Fatalf("Expected trc-fallback, got %v", fields["TraceId"])
	}

	if fields["UserId"] != "usr-fallback" {
		t.Fatalf("Expected usr-fallback, got %v", fields["UserId"])
	}
}

func TestFromContext_NilAndEmpty(t *testing.T) {
	logger := MustDefault()
	if FromContext(nil, logger) != logger {
		t.Fatalf("Expected same logger on nil context")
	}

	if FromContext(context.Background(), nil) != nil {
		t.Fatalf("Expected nil on nil logger")
	}

	emptyLogger := FromContext(context.Background(), logger)
	if emptyLogger != logger {
		t.Fatalf("Expected same logger on empty context")
	}
}

func TestFromContext_WithFields(t *testing.T) {
	logger := MustDefault()
	ctx := context.WithValue(context.Background(), RequestIdKey, "req-1234")
	ctx = context.WithValue(ctx, UserIdKey, "usr-999")

	childLogger := FromContext(ctx, logger)
	if childLogger == nil {
		t.Fatalf("Expected non-nil child logger")
	}

	childLogger.Info("Context-aware log message")
}
