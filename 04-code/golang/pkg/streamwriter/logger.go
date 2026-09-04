package streamwriter

import (
	"context"
	"sync"
	"time"
)

// Logger coordinates multiple writers and streamers with fluent chaining.
type Logger struct {
	mu      sync.RWMutex
	writers []WriterInterface
}

// NewLogger creates an empty Logger in silent mode (0 writers, 0 allocations).
func NewLogger() *Logger {
	return &Logger{
		writers: make([]WriterInterface, 0),
	}
}

// AddWriter fluently registers a single writer.
func (l *Logger) AddWriter(w WriterInterface) *Logger {
	if w == nil {
		return l
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.writers = append(l.writers, w.AsWriter())
	return l
}

// AddWriters fluently registers multiple writers in one call.
func (l *Logger) AddWriters(ws ...WriterInterface) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, w := range ws {
		if w != nil {
			l.writers = append(l.writers, w.AsWriter())
		}
	}
	return l
}

// AddStreamer fluently registers a streamer (adapting it via AsWriter()).
func (l *Logger) AddStreamer(s StreamerInterface) *Logger {
	if s == nil {
		return l
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.writers = append(l.writers, s.AsWriter())
	return l
}

// ClearWriters removes all registered writers (switches to silent mode).
func (l *Logger) ClearWriters() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.writers = l.writers[:0]
	return l
}

// RemoveWriter removes a registered writer by name.
func (l *Logger) RemoveWriter(name string) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	filtered := make([]WriterInterface, 0, len(l.writers))
	for _, w := range l.writers {
		if w.Name() != name {
			filtered = append(filtered, w)
		}
	}
	l.writers = filtered
	return l
}

// WriterCount returns the number of active writers.
func (l *Logger) WriterCount() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return len(l.writers)
}

// Emit sends an arbitrary payload (log-based or non-log-based) to all active writers.
func (l *Logger) Emit(ctx context.Context, payload any) error {
	l.mu.RLock()
	if len(l.writers) == 0 {
		l.mu.RUnlock()
		return nil
	}
	active := make([]WriterInterface, len(l.writers))
	copy(active, l.writers)
	l.mu.RUnlock()

	var firstErr error
	for _, w := range active {
		err := w.Write(ctx, payload)
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

// Info emits a structured LevelInfo log record to all writers.
func (l *Logger) Info(ctx context.Context, msg string, fields ...map[string]any) error {
	return l.dispatch(ctx, LevelInfo, msg, fields...)
}

// Error emits a structured LevelError log record to all writers.
func (l *Logger) Error(ctx context.Context, msg string, fields ...map[string]any) error {
	return l.dispatch(ctx, LevelError, msg, fields...)
}

// Debug emits a structured LevelDebug log record to all writers.
func (l *Logger) Debug(ctx context.Context, msg string, fields ...map[string]any) error {
	return l.dispatch(ctx, LevelDebug, msg, fields...)
}

// Warn emits a structured LevelWarn log record to all writers.
func (l *Logger) Warn(ctx context.Context, msg string, fields ...map[string]any) error {
	return l.dispatch(ctx, LevelWarn, msg, fields...)
}

// Sync flushes all active writers.
func (l *Logger) Sync() error {
	l.mu.RLock()
	active := make([]WriterInterface, len(l.writers))
	copy(active, l.writers)
	l.mu.RUnlock()

	for _, w := range active {
		_ = w.Sync()
	}
	return nil
}

// Close closes all active writers.
func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, w := range l.writers {
		_ = w.Close()
	}
	l.writers = l.writers[:0]
	return nil
}

func (l *Logger) dispatch(ctx context.Context, lvl LogLevel, msg string, fields ...map[string]any) error {
	l.mu.RLock()
	// Zero-allocation silent guard: if no writers, return immediately
	if len(l.writers) == 0 {
		l.mu.RUnlock()
		return nil
	}
	l.mu.RUnlock()

	traceID := ""
	userID := ""
	if ctx != nil {
		if tid, isOk := ctx.Value("traceId").(string); isOk {
			traceID = tid
		}
		if uid, isOk := ctx.Value("userId").(string); isOk {
			userID = uid
		}
	}

	merged := make(map[string]any)
	for _, f := range fields {
		for k, v := range f {
			merged[k] = v
		}
	}

	record := LogRecord{
		Timestamp: time.Now().UTC(),
		Level:     lvl,
		Message:   msg,
		Context:   ctx,
		Fields:    merged,
		TraceID:   traceID,
		UserID:    userID,
	}

	return l.Emit(ctx, record)
}
