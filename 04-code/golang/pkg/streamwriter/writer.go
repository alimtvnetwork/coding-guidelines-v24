package streamwriter

import (
	"context"
	"sync"
)

// WriterOptions configures the pluggable writer.
type WriterOptions struct {
	Name         string
	Streamer     StreamerInterface
	FormatMethod FormatFunc
	WriteMethod  WriteFunc
}

// PluggableWriter provides a composable write engine with swappable methods.
type PluggableWriter struct {
	mu           sync.RWMutex
	name         string
	streamer     StreamerInterface
	formatMethod FormatFunc
	writeMethod  WriteFunc
}

// NewPluggableWriter constructs a pluggable writer.
func NewPluggableWriter(opts WriterOptions) *PluggableWriter {
	name := opts.Name
	if name == "" {
		name = "pluggable-writer"
	}

	w := &PluggableWriter{
		name:         name,
		streamer:     opts.Streamer,
		formatMethod: opts.FormatMethod,
	}

	if opts.WriteMethod != nil {
		w.writeMethod = opts.WriteMethod
	} else {
		w.writeMethod = w.defaultWrite
	}
	return w
}

// Name returns the writer identifier.
func (w *PluggableWriter) Name() string {
	return w.name
}

// Write delegates to the active writeMethod function under read-lock.
func (w *PluggableWriter) Write(ctx context.Context, payload any) error {
	w.mu.RLock()
	fn := w.writeMethod
	w.mu.RUnlock()

	return fn(ctx, payload)
}

// SetWriteMethod hot-swaps the write method at runtime.
func (w *PluggableWriter) SetWriteMethod(fn WriteFunc) {
	if fn == nil {
		return
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	w.writeMethod = fn
}

// SetFormatMethod hot-swaps the formatter function at runtime.
func (w *PluggableWriter) SetFormatMethod(fn FormatFunc) {
	if fn == nil {
		return
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	w.formatMethod = fn
}

// SetStreamer hot-swaps the underlying streamer at runtime.
func (w *PluggableWriter) SetStreamer(s StreamerInterface) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.streamer = s
}

// Streamer returns the attached streamer under read-lock.
func (w *PluggableWriter) Streamer() StreamerInterface {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.streamer
}

// AsWriter returns the self-binding WriterInterface.
func (w *PluggableWriter) AsWriter() WriterInterface {
	return w
}

// AsInterfacer returns the self-binding Interfacer.
func (w *PluggableWriter) AsInterfacer() Interfacer {
	return w
}

// Sync flushes the underlying streamer if attached.
func (w *PluggableWriter) Sync() error {
	w.mu.RLock()
	s := w.streamer
	w.mu.RUnlock()

	if s != nil {
		return s.Sync()
	}
	return nil
}

// Close closes the underlying streamer if attached.
func (w *PluggableWriter) Close() error {
	w.mu.Lock()
	s := w.streamer
	w.mu.Unlock()

	if s != nil {
		return s.Close()
	}
	return nil
}

func (w *PluggableWriter) defaultWrite(ctx context.Context, payload any) error {
	w.mu.RLock()
	s := w.streamer
	formatter := w.formatMethod
	w.mu.RUnlock()

	// If a custom formatter is provided, transform the payload first
	if formatter != nil {
		formattedBytes, err := formatter(payload)
		if err != nil {
			return err
		}
		if s != nil {
			return s.Stream(ctx, formattedBytes)
		}
		return nil
	}

	// Forward directly to attached streamer
	if s != nil {
		return s.Stream(ctx, payload)
	}

	// No streamer attached: silent discard
	return nil
}

// Ensure interface satisfaction at compile-time
var _ WriterInterface = (*PluggableWriter)(nil)
