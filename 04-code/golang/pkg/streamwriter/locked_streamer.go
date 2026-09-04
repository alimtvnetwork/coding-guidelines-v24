package streamwriter

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
)

// LockedOptions configures the thread-safe locked streamer.
type LockedOptions struct {
	Name         string
	Destination  io.Writer
	StreamMethod StreamFunc
}

// LockedStreamer implements StreamerInterface with mutex synchronization.
type LockedStreamer struct {
	mu           sync.RWMutex
	name         string
	destination  io.Writer
	streamMethod StreamFunc
}

// NewLockedStreamer constructs a thread-safe streamer.
func NewLockedStreamer(opts LockedOptions) *LockedStreamer {
	name := opts.Name
	if name == "" {
		name = "locked-streamer"
	}
	dest := opts.Destination
	if dest == nil {
		dest = os.Stdout
	}

	s := &LockedStreamer{
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
func (s *LockedStreamer) Name() string {
	return s.name
}

// Stream executes the swappable stream method under mutex lock.
func (s *LockedStreamer) Stream(ctx context.Context, payload any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.streamMethod(ctx, payload, s.destination)
}

// Write satisfies WriterInterface by delegating to Stream.
func (s *LockedStreamer) Write(ctx context.Context, payload any) error {
	return s.Stream(ctx, payload)
}

// SetStreamMethod hot-swaps the streaming logic at runtime.
func (s *LockedStreamer) SetStreamMethod(fn StreamFunc) {
	if fn == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.streamMethod = fn
}

// SetDestination hot-swaps the output destination at runtime.
func (s *LockedStreamer) SetDestination(dest io.Writer) {
	if dest == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.destination = dest
}

// IsLocked reports true for LockedStreamer.
func (s *LockedStreamer) IsLocked() bool {
	return true
}

// Destination returns the active destination under read-lock.
func (s *LockedStreamer) Destination() io.Writer {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.destination
}

// AsStreamer returns the self-binding StreamerInterface.
func (s *LockedStreamer) AsStreamer() StreamerInterface {
	return s
}

// AsWriter returns the self-binding WriterInterface.
func (s *LockedStreamer) AsWriter() WriterInterface {
	return s
}

// AsInterfacer returns the self-binding Interfacer.
func (s *LockedStreamer) AsInterfacer() Interfacer {
	return s
}

// Sync flushes the underlying destination if supported.
func (s *LockedStreamer) Sync() error {
	s.mu.RLock()
	dest := s.destination
	s.mu.RUnlock()

	if syncer, isOk := dest.(interface{ Sync() error }); isOk {
		return syncer.Sync()
	}
	return nil
}

// Close closes the underlying destination if it implements io.Closer.
func (s *LockedStreamer) Close() error {
	s.mu.Lock()
	dest := s.destination
	s.mu.Unlock()

	if closer, isOk := dest.(io.Closer); isOk {
		return closer.Close()
	}
	return nil
}

func (s *LockedStreamer) defaultStream(ctx context.Context, payload any, dest io.Writer) error {
	var line string
	if record, isOk := payload.(LogRecord); isOk {
		line = fmt.Sprintf("[%s][locked] %s %s: %s\n",
			s.name,
			record.Timestamp.Format("15:04:05.000"),
			record.Level.String(),
			record.Message,
		)
	} else {
		line = fmt.Sprintf("[%s][locked] %v\n", s.name, payload)
	}
	_, err := dest.Write([]byte(line))
	return err
}
