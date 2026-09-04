package streamwriter

import (
	"context"
	"fmt"
	"io"
	"time"
)

// Interfacer represents the self-binding contract returning its own interface.
type Interfacer interface {
	AsInterfacer() Interfacer
}

// WriterInterface defines universal write operations over generic type T with self-binding.
type WriterInterface[T any] interface {
	Interfacer
	Name() string
	Write(ctx context.Context, payload T) error
	AsWriter() WriterInterface[T]
	Sync() error
	Close() error
}

// StreamerInterface defines streaming operations over generic type T with locking introspection.
type StreamerInterface[T any] interface {
	Interfacer
	Name() string
	Stream(ctx context.Context, payload T) error
	AsStreamer() StreamerInterface[T]
	AsWriter() WriterInterface[T]
	IsLocked() bool
	Destination() io.Writer
	Sync() error
	Close() error
}

// StreamFunc defines the swappable function signature for streaming data of type T.
type StreamFunc[T any] func(ctx context.Context, payload T, dest io.Writer) error

// WriteFunc defines the swappable function signature for write operations over type T.
type WriteFunc[T any] func(ctx context.Context, payload T) error

// FormatFunc defines the serialization transformation from payload T to bytes.
type FormatFunc[T any] func(payload T) ([]byte, error)

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

// Compile satisfies the Compilable interface for LogRecord with deterministic ordering.
func (r LogRecord) Compile() string {
	res := fmt.Sprintf("[%s] %-5s: %s", r.Timestamp.Format("15:04:05.000"), r.Level.String(), r.Message)
	if r.TraceID != "" {
		res += fmt.Sprintf(" [trace=%s]", r.TraceID)
	}
	if r.UserID != "" {
		res += fmt.Sprintf(" [user=%s]", r.UserID)
	}
	if len(r.Fields) > 0 {
		res += fmt.Sprintf(" fields=%s", Compile(r.Fields))
	}
	return res
}

// Ensure LogRecord implements Compilable at compile-time.
var _ Compilable = LogRecord{}
