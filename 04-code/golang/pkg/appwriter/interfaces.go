package appwriter

import (
	"context"
	"io"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/result"
)

// Writer defines the unified writer interface returning structured AppError.
type Writer interface {
	Write(ctx context.Context, payload any) *appfault.AppError
	AsStreamer() Streamer[any]
	AsWriter() Writer
	Destination() io.Writer
	IsLocked() bool
	Lock()
	Unlock()
	RLock()
	RUnlock()
	Sync() *appfault.AppError
	Close() *appfault.AppError
}

// Streamer defines the typed streaming interface returning structured AppError.
type Streamer[T any] interface {
	Stream(ctx context.Context, payload T) *appfault.AppError
	AsStreamer() Streamer[T]
	AsWriter() Writer
	Destination() io.Writer
	IsLocked() bool
	Lock()
	Unlock()
	RLock()
	RUnlock()
	Sync() *appfault.AppError
	Close() *appfault.AppError
}

// WriteMethodFunc defines an injected write method that accepts self as the first parameter.
type WriteMethodFunc func(ctx context.Context, self Writer, payload any) *appfault.AppError

// StreamMethodFunc defines an injected streamer method that accepts self as the first parameter.
type StreamMethodFunc[T any] func(ctx context.Context, self Streamer[T], payload T) *appfault.AppError

// BaseWriterWrap encapsulates Result/Wrap containing *BaseWriter.
type BaseWriterWrap = result.Wrap[*BaseWriter]
