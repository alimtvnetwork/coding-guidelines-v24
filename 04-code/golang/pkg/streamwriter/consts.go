package streamwriter

import (
	"context"
	"io"

	"coding-guidelines/common/pkg/appfault"
)

// LogLevel enum constants defining standardized severity tiers.
const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// PayloadKind enum constants identifying classification of an incoming generic payload.
const (
	// PayloadNil indicates a nil or uninitialized payload.
	PayloadNil PayloadKind = iota
	// PayloadBytes indicates raw byte slice ([]byte).
	PayloadBytes
	// PayloadString indicates a string payload.
	PayloadString
	// PayloadError indicates a structured *appfault.AppError or standard error.
	PayloadError
	// PayloadMap indicates a key-value map.
	PayloadMap
	// PayloadStruct indicates a struct or pointer to struct.
	PayloadStruct
	// PayloadPrimitive indicates a scalar primitive (int, bool, float).
	PayloadPrimitive
)

// StreamFunc defines the swappable function signature returning *appfault.AppError.
type StreamFunc[T any] func(ctx context.Context, payload T, dest io.Writer) *appfault.AppError

// WriteFunc defines the swappable function signature returning *appfault.AppError.
// It receives the attached streamer as the first parameter, the active context, the current writer object, and the generic payload.
type WriteFunc[T any] func(streamer Streamer[T], ctx context.Context, writer *PluggableWriter[T], payload T) *appfault.AppError

// FormatFunc defines the serialization transformation returning Bytes[T].
type FormatFunc[T any] func(payload T) Bytes[T]

// ErrorHandlerFunc handles an AppError callback.
type ErrorHandlerFunc func(err *appfault.AppError)

// LogFormatterFunc formats a LogRecord into a string.
type LogFormatterFunc func(record LogRecord) string
