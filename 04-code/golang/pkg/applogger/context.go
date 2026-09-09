package applogger

import (
	"context"
)

// ContextKey defines a custom type for context values to avoid collisions.
type ContextKey string

const (
	// RequestIDKey is the standard context key for request tracking.
	RequestIDKey ContextKey = "RequestId"
	// TraceIDKey is the standard context key for distributed trace tracking.
	TraceIDKey ContextKey = "TraceId"
	// UserIDKey is the standard context key for the authenticated user ID.
	UserIDKey ContextKey = "UserId"

	// RequestIdKey is a PascalCase alias for RequestIDKey.
	RequestIdKey = RequestIDKey
	// TraceIdKey is a PascalCase alias for TraceIDKey.
	TraceIdKey = TraceIDKey
	// UserIdKey is a PascalCase alias for UserIDKey.
	UserIdKey = UserIDKey
)

// ExtractContextFields retrieves standard observability fields from context.
func ExtractContextFields(ctx context.Context) map[string]any {
	fields := make(map[string]any)
	extractField(ctx, RequestIDKey, fields, "RequestID", "request_id")
	extractField(ctx, TraceIDKey, fields, "TraceID", "trace_id")
	extractField(ctx, UserIDKey, fields, "UserID", "user_id")

	return fields
}

// extractField sets the field value if found in context.
func extractField(ctx context.Context, key ContextKey, fields map[string]any, alts ...string) {
	val := resolveContextValue(ctx, key, alts...)
	if val != "" {
		fields[string(key)] = val
	}
}

// resolveContextValue searches for a string value by key and alternate names.
func resolveContextValue(ctx context.Context, key ContextKey, alts ...string) string {
	if val, ok := ctx.Value(key).(string); ok && val != "" {
		return val
	}

	if val, ok := ctx.Value(string(key)).(string); ok && val != "" {
		return val
	}

	return resolveFallbackValues(ctx, alts...)
}

// resolveFallbackValues checks alternate key variants in context.
func resolveFallbackValues(ctx context.Context, alts ...string) string {
	for _, alt := range alts {
		if val, ok := ctx.Value(alt).(string); ok && val != "" {
			return val
		}

		if val, ok := ctx.Value(ContextKey(alt)).(string); ok && val != "" {
			return val
		}
	}

	return ""
}

// FromContext creates a child Logger that inherits standard observability fields
// (like RequestId, TraceId, and UserId) directly from the provided context.Context.
// If the context is nil, it simply returns the unmodified logger.
func FromContext(ctx context.Context, logger Logger) Logger {
	if ctx == nil || logger == nil {
		return logger
	}

	fields := ExtractContextFields(ctx)
	if len(fields) == 0 {
		return logger
	}

	return logger.WithFields(fields)
}
