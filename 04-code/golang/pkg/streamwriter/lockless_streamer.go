package streamwriter

import (
	"context"
	"fmt"
	"io"
	"os"
)

// LocklessOptions configures the zero-overhead lockless streamer for payload type T.
type LocklessOptions[T any] struct {
	Name         string
	Destination  io.Writer
	StreamMethod StreamFunc[T]
}

// LocklessStreamer implements StreamerInterface[T] with zero lock overhead.
type LocklessStreamer[T any] struct {
	name         string
	destination  io.Writer
	streamMethod StreamFunc[T]
}

// NewLocklessStreamer constructs a zero-lock streamer over generic type T.
func NewLocklessStreamer[T any](opts LocklessOptions[T]) *LocklessStreamer[T] {
	name := opts.Name
	if name == "" {
		name = "lockless-streamer"
	}
	dest := opts.Destination
	if dest == nil {
		dest = os.Stdout
	}

	s := &LocklessStreamer[T]{
		name:        name,
		destination: dest,
	}

	if opts.StreamMethod != nil {
		s.streamMethod = opts.StreamMethod
	} else {
		s.streamMethod = s.defaultStream
	}
	return s
}

// Name returns the streamer identifier.
func (s *LocklessStreamer[T]) Name() string {
	return s.name
}

// Stream executes directly with zero mutex operations.
func (s *LocklessStreamer[T]) Stream(ctx context.Context, payload T) error {
	return s.streamMethod(ctx, payload, s.destination)
}

// Write satisfies WriterInterface[T] by delegating to Stream.
func (s *LocklessStreamer[T]) Write(ctx context.Context, payload T) error {
	return s.Stream(ctx, payload)
}

// SetStreamMethod swaps the streaming logic.
func (s *LocklessStreamer[T]) SetStreamMethod(fn StreamFunc[T]) {
	if fn != nil {
		s.streamMethod = fn
	}
}

// SetDestination swaps the output destination.
func (s *LocklessStreamer[T]) SetDestination(dest io.Writer) {
	if dest != nil {
		s.destination = dest
	}
}

// IsLocked reports false for LocklessStreamer.
func (s *LocklessStreamer[T]) IsLocked() bool {
	return false
}

// Destination returns the active destination.
func (s *LocklessStreamer[T]) Destination() io.Writer {
	return s.destination
}

// AsStreamer returns the self-binding StreamerInterface[T].
func (s *LocklessStreamer[T]) AsStreamer() StreamerInterface[T] {
	return s
}

// AsWriter returns the self-binding WriterInterface[T].
func (s *LocklessStreamer[T]) AsWriter() WriterInterface[T] {
	return s
}

// AsInterfacer returns the self-binding Interfacer.
func (s *LocklessStreamer[T]) AsInterfacer() Interfacer {
	return s
}

// Sync flushes the underlying destination if supported.
func (s *LocklessStreamer[T]) Sync() error {
	if syncer, isOk := s.destination.(interface{ Sync() error }); isOk {
		return syncer.Sync()
	}
	return nil
}

// Close closes the underlying destination if it implements io.Closer.
func (s *LocklessStreamer[T]) Close() error {
	if closer, isOk := s.destination.(io.Closer); isOk {
		return closer.Close()
	}
	return nil
}

func (s *LocklessStreamer[T]) defaultStream(ctx context.Context, payload T, dest io.Writer) error {
	compiled := Compile(payload)
	line := fmt.Sprintf("[%s][lockless] %s\n", s.name, compiled)
	_, err := dest.Write([]byte(line))
	return err
}

var _ StreamerInterface[any] = (*LocklessStreamer[any])(nil)
