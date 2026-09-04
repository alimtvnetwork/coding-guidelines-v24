package streamwriter

import (
	"context"
	"io"
	"time"
)

// Interfacer represents the self-binding contract returning its own interface.
type Interfacer interface {
	AsInterfacer() Interfacer
}

// WriterInterface defines universal write operations with self-binding.
type WriterInterface interface {
	Interfacer
	Name() string
	Write(ctx context.Context, payload any) error
	AsWriter() WriterInterface
	Sync() error
	Close() error
}

// StreamerInterface defines streaming operations with locking introspection and self-binding.
type StreamerInterface interface {
	Interfacer
	Name() string
	Stream(ctx context.Context, payload any) error
	AsStreamer() StreamerInterface
	AsWriter() WriterInterface
	IsLocked() bool
	Destination() io.Writer
	Sync() error
	Close() error
}

// StreamFunc defines the swappable signature for writing/streaming to a destination.
type StreamFunc func(ctx context.Context, payload any, dest io.Writer) error

// WriteFunc defines the swappable function signature for write operations.
type WriteFunc func(ctx context.Context, payload any) error

// FormatFunc defines the serialization transformation from payload to bytes.
type FormatFunc func(payload any) ([]byte, error)

// LogLevel defines standardized severity tiers.
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l LogLevel) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// LogRecord carries normalized event data for log-based flows.
type LogRecord struct {
	Timestamp time.Time       `json:"timestamp"`
	Level     LogLevel        `json:"level"`
	Message   string          `json:"message"`
	Context   context.Context `json:"-"`
	Fields    map[string]any  `json:"fields,omitempty"`
	TraceID   string          `json:"traceId,omitempty"`
	UserID    string          `json:"userId,omitempty"`
}
