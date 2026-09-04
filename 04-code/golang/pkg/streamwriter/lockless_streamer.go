package streamwriter

import (
	"context"
	"fmt"
	"io"
	"os"
)

// LocklessOptions configures the zero-overhead lockless streamer.
type LocklessOptions struct {
	Name         string
	Destination  io.Writer
	StreamMethod StreamFunc
}

// LocklessStreamer implements StreamerInterface with zero lock overhead.
type LocklessStreamer struct {
	name         string
	destination  io.Writer
	streamMethod StreamFunc
}

// NewLocklessStreamer constructs a zero-lock streamer.
func NewLocklessStreamer(opts LocklessOptions) *LocklessStreamer {
	name := opts.Name
	if name == "" {
		name = "lockless-streamer"
	}
	dest := opts.Destination
	if dest == nil {
		dest = os.Stdout
	}

	s := &LocklessStreamer{
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
func (s *LocklessStreamer) Name() string {
	return s.name
}

// Stream executes directly with zero mutex operations.
func (s *LocklessStreamer) Stream(ctx context.Context, payload any) error {
	return s.streamMethod(ctx, payload, s.destination)
}

// Write satisfies WriterInterface by delegating to Stream.
func (s *LocklessStreamer) Write(ctx context.Context, payload any) error {
	return s.Stream(ctx, payload)
}

// SetStreamMethod swaps the streaming logic.
func (s *LocklessStreamer) SetStreamMethod(fn StreamFunc) {
	if fn != nil {
		s.streamMethod = fn
	}
}

// SetDestination swaps the output destination.
func (s *LocklessStreamer) SetDestination(dest io.Writer) {
	if dest != nil {
		s.destination = dest
	}
}

// IsLocked reports false for LocklessStreamer.
func (s *LocklessStreamer) IsLocked() bool {
	return false
}

// Destination returns the active destination.
func (s *LocklessStreamer) Destination() io.Writer {
	return s.destination
}

// AsStreamer returns the self-binding StreamerInterface.
func (s *LocklessStreamer) AsStreamer() StreamerInterface {
	return s
}

// AsWriter returns the self-binding WriterInterface.
func (s *LocklessStreamer) AsWriter() WriterInterface {
	return s
}

// AsInterfacer returns the self-binding Interfacer.
func (s *LocklessStreamer) AsInterfacer() Interfacer {
	return s
}

// Sync flushes the underlying destination if supported.
func (s *LocklessStreamer) Sync() error {
	if syncer, isOk := s.destination.(interface{ Sync() error }); isOk {
		return syncer.Sync()
	}
	return nil
}

// Close closes the underlying destination if it implements io.Closer.
func (s *LocklessStreamer) Close() error {
	if closer, isOk := s.destination.(io.Closer); isOk {
		return closer.Close()
	}
	return nil
}

func (s *LocklessStreamer) defaultStream(ctx context.Context, payload any, dest io.Writer) error {
	var line string
	if record, isOk := payload.(LogRecord); isOk {
		line = fmt.Sprintf("[%s][lockless] %s %s: %s\n",
			s.name,
			record.Timestamp.Format("15:04:05.000"),
			record.Level.String(),
			record.Message,
		)
	} else {
		line = fmt.Sprintf("[%s][lockless] %v\n", s.name, payload)
	}
	_, err := dest.Write([]byte(line))
	return err
}
