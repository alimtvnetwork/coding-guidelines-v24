package applogger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sync"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

type (
	contextStringFaultStreamer interface {
		Stream(ctx context.Context, payload string) *appfault.AppError
	}

	contextStringErrStreamer interface {
		Stream(ctx context.Context, payload string) error
	}

	contextAnyFaultStreamer interface {
		Stream(ctx context.Context, payload any) *appfault.AppError
	}

	contextAnyErrStreamer interface {
		Stream(ctx context.Context, payload any) error
	}

	faultSyncer interface {
		Sync() *appfault.AppError
	}

	errSyncer interface {
		Sync() error
	}

	faultCloser interface {
		Close() *appfault.AppError
	}

	errCloser interface {
		Close() error
	}
)

// StreamerSink bridges an arbitrary streamer or io.Writer into the LogSink interface.
type StreamerSink struct {
	lock     sync.Mutex
	streamer any
}

// Compile-time check that StreamerSink implements LogSink, LogStreamer, and Streamer.
var (
	_ LogSink     = (*StreamerSink)(nil)
	_ LogStreamer = (*StreamerSink)(nil)
	_ Streamer    = (*StreamerSink)(nil)
)

// NewStreamerSink creates a new StreamerSink instance.
func NewStreamerSink(streamer any) *StreamerSink {
	return &StreamerSink{
		streamer: streamer,
	}
}

// Streamer returns the underlying streamer or writer.
func (s *StreamerSink) Streamer() any {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.streamer
}

// Unwrap returns the underlying raw streamer object.
func (s *StreamerSink) Unwrap() any {
	return s.Streamer()
}

// Destination returns the underlying destination io.Writer if available.
func (s *StreamerSink) Destination() io.Writer {
	s.lock.Lock()
	defer s.lock.Unlock()
	if w, isWriter := s.streamer.(io.Writer); isWriter {
		return w
	}

	if dp, isDest := s.streamer.(interface{ Destination() io.Writer }); isDest {
		return dp.Destination()
	}

	return nil
}

// DriverType returns DriverStreamer for introspection.
func (s *StreamerSink) DriverType() DriverType {
	return DriverStreamer
}

// Name returns the sink identifier name, preferring the inner streamer name if present.
func (s *StreamerSink) Name() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	if np, isNamed := s.streamer.(interface{ Name() string }); isNamed {
		if name := np.Name(); name != "" {
			return name
		}
	}

	return "streamer"
}

// StreamEntry satisfies the LogStreamer interface.
func (s *StreamerSink) StreamEntry(e LogEntry) error {
	return s.WriteEntry(e)
}

// Stream sends an arbitrary payload to the underlying streamer.
func (s *StreamerSink) Stream(ctx context.Context, payload any) *appfault.AppError {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.streamer == nil {
		return appfault.New(errtype.Validation, "streamer is nil")
	}

	return s.dispatchPayload(ctx, payload)
}

func (s *StreamerSink) dispatchPayload(ctx context.Context, payload any) *appfault.AppError {
	if str, isStr := payload.(string); isStr {
		if err := s.dispatchStreamWithContext(ctx, []byte(str)); err != nil {
			return appfault.Wrap(errtype.IO, err, "failed to stream payload")
		}

		return nil
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return appfault.Wrap(errtype.Validation, err, "failed to marshal payload")
	}

	if err := s.dispatchStreamWithContext(ctx, b); err != nil {
		return appfault.Wrap(errtype.IO, err, "failed to stream payload")
	}

	return nil
}

// WriteEntry serializes entry to JSON and routes to the underlying streamer.
func (s *StreamerSink) WriteEntry(e LogEntry) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.streamer == nil {
		return fmt.Errorf("streamer is nil")
	}

	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	return s.dispatchStream(b)
}

func (s *StreamerSink) dispatchStream(b []byte) error {
	return s.dispatchStreamWithContext(context.Background(), b)
}

func (s *StreamerSink) dispatchStreamWithContext(ctx context.Context, b []byte) error {
	str := string(b)
	if err, hasHandled := s.streamStringFault(ctx, str); hasHandled {
		return err
	}

	if err, hasHandled := s.streamAnyFault(ctx, str); hasHandled {
		return err
	}

	return s.dispatchRemainingStream(ctx, str, b)
}

func (s *StreamerSink) dispatchRemainingStream(ctx context.Context, str string, b []byte) error {
	if err, hasHandled := s.streamStringErr(ctx, str); hasHandled {
		return err
	}

	if err, hasHandled := s.streamAnyErr(ctx, str); hasHandled {
		return err
	}

	if err, hasHandled := s.streamWriter(b); hasHandled {
		return err
	}

	return s.dispatchFallback(ctx, str)
}

func (s *StreamerSink) streamStringFault(ctx context.Context, str string) (error, bool) {
	st, isMatch := s.streamer.(contextStringFaultStreamer)
	if isMatch {
		if fault := st.Stream(ctx, str); fault != nil {
			return fault, true
		}

		return nil, true
	}

	return nil, false
}

func (s *StreamerSink) streamAnyFault(ctx context.Context, str string) (error, bool) {
	st, isMatch := s.streamer.(contextAnyFaultStreamer)
	if isMatch {
		if fault := st.Stream(ctx, str); fault != nil {
			return fault, true
		}

		return nil, true
	}

	return nil, false
}

func (s *StreamerSink) streamStringErr(ctx context.Context, str string) (error, bool) {
	st, isMatch := s.streamer.(contextStringErrStreamer)
	if isMatch {
		return st.Stream(ctx, str), true
	}

	return nil, false
}

func (s *StreamerSink) streamAnyErr(ctx context.Context, str string) (error, bool) {
	st, isMatch := s.streamer.(contextAnyErrStreamer)
	if isMatch {
		return st.Stream(ctx, str), true
	}

	return nil, false
}

func (s *StreamerSink) streamWriter(b []byte) (error, bool) {
	w, isMatch := s.streamer.(io.Writer)
	if isMatch {
		_, err := fmt.Fprintln(w, string(b))

		return err, true
	}

	return nil, false
}

func (s *StreamerSink) dispatchFallback(ctx context.Context, str string) error {
	if err, hasHandled := s.invokeReflectStream(ctx, str); hasHandled {
		return err
	}

	return fmt.Errorf("streamer does not support Write, Stream, or io.Writer")
}

func (s *StreamerSink) invokeReflectStream(ctx context.Context, str string) (error, bool) {
	val := reflect.ValueOf(s.streamer)
	method := val.MethodByName("Stream")
	if method.IsValid() {
		return s.callReflectStream(method, ctx, str)
	}

	return nil, false
}

func (s *StreamerSink) callReflectStream(method reflect.Value, ctx context.Context, str string) (error, bool) {
	numArgs := method.Type().NumIn()
	var in []reflect.Value
	if numArgs == 2 {
		in = []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(str)}
	} else if numArgs == 1 {
		in = []reflect.Value{reflect.ValueOf(str)}
	}

	if len(in) > 0 {
		return s.evalReflectResults(method.Call(in))
	}

	return nil, false
}

func (s *StreamerSink) evalReflectResults(out []reflect.Value) (error, bool) {
	if len(out) == 0 {
		return nil, true
	}

	if first := out[0]; first.IsValid() {
		if errVal, isErr := first.Interface().(error); isErr {
			return errVal, true
		}
	}

	return nil, true
}

func (s *StreamerSink) syncFault() (error, bool) {
	syncer, isMatch := s.streamer.(faultSyncer)
	if isMatch {
		if fault := syncer.Sync(); fault != nil {
			return fault, true
		}

		return nil, true
	}

	return nil, false
}

func (s *StreamerSink) syncErr() (error, bool) {
	syncer, isMatch := s.streamer.(errSyncer)
	if isMatch {
		return syncer.Sync(), true
	}

	return nil, false
}

// Sync flushes the underlying streamer if it supports Sync.
func (s *StreamerSink) Sync() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.streamer == nil {
		return nil
	}

	if err, hasHandled := s.syncFault(); hasHandled {
		return err
	}

	if err, hasHandled := s.syncErr(); hasHandled {
		return err
	}

	return nil
}

func (s *StreamerSink) closeFault() (error, bool) {
	closer, isMatch := s.streamer.(faultCloser)
	if isMatch {
		if fault := closer.Close(); fault != nil {
			return fault, true
		}

		return nil, true
	}

	return nil, false
}

func (s *StreamerSink) closeErr() (error, bool) {
	closer, isMatch := s.streamer.(errCloser)
	if isMatch {
		return closer.Close(), true
	}

	return nil, false
}

// Close closes the underlying streamer if it supports Close.
func (s *StreamerSink) Close() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.streamer == nil {
		return nil
	}

	if err, hasHandled := s.closeFault(); hasHandled {
		return err
	}

	if err, hasHandled := s.closeErr(); hasHandled {
		return err
	}

	return nil
}
