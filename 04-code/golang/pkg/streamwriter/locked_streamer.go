package streamwriter

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
)

// LockedOptions configures the thread-safe locked streamer for payload type T.
type LockedOptions[T any] struct {
	Name         string
	Destination  io.Writer
	StreamMethod StreamFunc[T]
}

// LockedStreamer implements StreamerInterface[T] with mutex synchronization.
type LockedStreamer[T any] struct {
	mu           sync.RWMutex
	name         string
	destination  io.Writer
	streamMethod StreamFunc[T]
}

// NewLockedStreamer constructs a thread-safe streamer over generic type T.
func NewLockedStreamer[T any](opts LockedOptions[T]) *LockedStreamer[T] {
	name := opts.Name
	if name == "" {
		name = "locked-streamer"
	}
	dest := opts.Destination
	if dest == nil {
		dest = os.Stdout
	}

	s := &LockedStreamer[T]{
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
func (s *LockedStreamer[T]) Name() string {
	return s.name
}

// Stream executes the swappable stream method under mutex lock.
func (s *LockedStreamer[T]) Stream(ctx context.Context, payload T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.streamMethod(ctx, payload, s.destination)
}

// Write satisfies WriterInterface[T] by delegating to Stream.
func (s *LockedStreamer[T]) Write(ctx context.Context, payload T) error {
	return s.Stream(ctx, payload)
}

// SetStreamMethod hot-swaps the streaming logic at runtime.
func (s *LockedStreamer[T]) SetStreamMethod(fn StreamFunc[T]) {
	if fn == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.streamMethod = fn
}

// SetDestination hot-swaps the output destination at runtime.
func (s *LockedStreamer[T]) SetDestination(dest io.Writer) {
	if dest == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.destination = dest
}

// IsLocked reports true for LockedStreamer.
func (s *LockedStreamer[T]) IsLocked() bool {
	return true
}

// Destination returns the active destination under read-lock.
func (s *LockedStreamer[T]) Destination() io.Writer {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.destination
}

// AsStreamer returns the self-binding StreamerInterface[T].
func (s *LockedStreamer[T]) AsStreamer() StreamerInterface[T] {
	return s
}

// AsWriter returns the self-binding WriterInterface[T].
func (s *LockedStreamer[T]) AsWriter() WriterInterface[T] {
	return s
}

// AsInterfacer returns the self-binding Interfacer.
func (s *LockedStreamer[T]) AsInterfacer() Interfacer {
	return s
}

// Sync flushes the underlying destination if supported.
func (s *LockedStreamer[T]) Sync() error {
	s.mu.RLock()
	dest := s.destination
	s.mu.RUnlock()

	if syncer, isOk := dest.(interface{ Sync() error }); isOk {
		return syncer.Sync()
	}
	return nil
}

// Close closes the underlying destination if it implements io.Closer.
func (s *LockedStreamer[T]) Close() error {
	s.mu.Lock()
	dest := s.destination
	s.mu.Unlock()

	if closer, isOk := dest.(io.Closer); isOk {
		return closer.Close()
	}
	return nil
}

func (s *LockedStreamer[T]) defaultStream(ctx context.Context, payload T, dest io.Writer) error {
	compiled := Compile(payload)
	line := fmt.Sprintf("[%s][locked] %s\n", s.name, compiled)
	_, err := dest.Write([]byte(line))
	return err
}

var _ StreamerInterface[any] = (*LockedStreamer[any])(nil)
