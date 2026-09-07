package streamwriter

import (
	"context"
	"sync"
	"time"

	"coding-guidelines/common/pkg/appfault"
)

// Logger coordinates multiple generic writers and streamers over type T with AppError returns.
type Logger[T any] struct {
	lock    sync.RWMutex
	writers []Writer[T]
}

// NewLogger creates an empty generic Logger in silent mode (0 writers, 0 allocations).
func NewLogger[T any]() *Logger[T] {
	return &Logger[T]{
		writers: make([]Writer[T], 0),
	}
}

// NewAnyLogger creates a universal AnyLogger (*Logger[any]) in silent mode.
func NewAnyLogger() *AnyLogger {
	return NewLogger[any]()
}

// AddWriter fluently registers a single writer.
func (l *Logger[T]) AddWriter(w Writer[T]) *Logger[T] {
	if w == nil {
		return l
	}

	l.lock.Lock()
	defer l.lock.Unlock()
	l.writers = append(l.writers, w.AsWriter())

	return l
}

// AddWriters fluently registers multiple writers in one call.
func (l *Logger[T]) AddWriters(ws ...Writer[T]) *Logger[T] {
	l.lock.Lock()
	defer l.lock.Unlock()
	for _, w := range ws {
		if w != nil {
			l.writers = append(l.writers, w.AsWriter())
		}
	}

	return l
}

// AddStreamer fluently registers a streamer (adapting it via AsWriter()).
func (l *Logger[T]) AddStreamer(s Streamer[T]) *Logger[T] {
	if s == nil {
		return l
	}

	l.lock.Lock()
	defer l.lock.Unlock()
	l.writers = append(l.writers, s.AsWriter())

	return l
}

// ClearWriters removes all registered writers (switches to silent mode).
func (l *Logger[T]) ClearWriters() *Logger[T] {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.writers = l.writers[:0]

	return l
}

// RemoveWriter removes a registered writer by name.
func (l *Logger[T]) RemoveWriter(name string) *Logger[T] {
	l.lock.Lock()
	defer l.lock.Unlock()
	filtered := make([]Writer[T], 0, len(l.writers))
	for _, w := range l.writers {
		if w.Name() != name {
			filtered = append(filtered, w)
		}
	}

	l.writers = filtered

	return l
}

// WriterCount returns the number of active writers.
func (l *Logger[T]) WriterCount() int {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return len(l.writers)
}

// Emit sends a generic payload T to all active writers, returning *appfault.AppError.
func (l *Logger[T]) Emit(ctx context.Context, payload T) *appfault.AppError {
	l.lock.RLock()
	// Zero-allocation silent guard
	if len(l.writers) == 0 {
		l.lock.RUnlock()

		return nil
	}

	active := make([]Writer[T], len(l.writers))
	copy(active, l.writers)
	l.lock.RUnlock()

	var firstErr *appfault.AppError
	for _, w := range active {
		err := w.Write(ctx, payload)
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}

	return firstErr
}

// Info emits a LevelInfo structured record if T can represent it, returning *appfault.AppError.
func (l *Logger[T]) Info(ctx context.Context, msg string, fields ...map[string]any) *appfault.AppError {
	return l.dispatchRecord(ctx, LevelInfo, msg, fields...)
}

// Error emits a LevelError structured record if T can represent it, returning *appfault.AppError.
func (l *Logger[T]) Error(ctx context.Context, msg string, fields ...map[string]any) *appfault.AppError {
	return l.dispatchRecord(ctx, LevelError, msg, fields...)
}

// Debug emits a LevelDebug structured record if T can represent it, returning *appfault.AppError.
func (l *Logger[T]) Debug(ctx context.Context, msg string, fields ...map[string]any) *appfault.AppError {
	return l.dispatchRecord(ctx, LevelDebug, msg, fields...)
}

// Warn emits a LevelWarn structured record if T can represent it, returning *appfault.AppError.
func (l *Logger[T]) Warn(ctx context.Context, msg string, fields ...map[string]any) *appfault.AppError {
	return l.dispatchRecord(ctx, LevelWarn, msg, fields...)
}

// Sync flushes all active writers, returning *appfault.AppError.
func (l *Logger[T]) Sync() *appfault.AppError {
	l.lock.RLock()
	active := make([]Writer[T], len(l.writers))
	copy(active, l.writers)
	l.lock.RUnlock()

	var firstErr *appfault.AppError
	for _, w := range active {
		if err := w.Sync(); err != nil && firstErr == nil {
			firstErr = err
		}
	}

	return firstErr
}

// Close closes all active writers, returning *appfault.AppError.
func (l *Logger[T]) Close() *appfault.AppError {
	l.lock.Lock()
	defer l.lock.Unlock()

	var firstErr *appfault.AppError
	for _, w := range l.writers {
		if err := w.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}

	l.writers = l.writers[:0]

	return firstErr
}

func (l *Logger[T]) dispatchRecord(ctx context.Context, lvl LogLevel, msg string, fields ...map[string]any) *appfault.AppError {
	l.lock.RLock()
	if len(l.writers) == 0 {
		l.lock.RUnlock()

		return nil
	}

	l.lock.RUnlock()

	traceId := ""
	userId := ""
	if ctx != nil {
		if tid, isOk := ctx.Value("traceId").(string); isOk {
			traceId = tid
		}

		if uid, isOk := ctx.Value("userId").(string); isOk {
			userId = uid
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
		TraceId:   traceId,
		UserId:    userId,
	}

	// Case 1: T is any or LogRecord
	if payload, isOk := any(record).(T); isOk {
		return l.Emit(ctx, payload)
	}

	// Case 2: T is string (compiled representation)
	if strPayload, isOk := any(record.Compile()).(T); isOk {
		return l.Emit(ctx, strPayload)
	}

	return nil
}
