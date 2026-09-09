package applogger

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/enum/logleveltype"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/streamwriter"
)

func TestStreamerSink_Basics(t *testing.T) {
	sink := NewStreamerSink(nil)
	assertBasics(t, sink)
	assertNilOperations(t, sink)
}

func assertBasics(t *testing.T, sink *StreamerSink) {
	if sink.DriverType() != DriverStreamer {
		t.Fatalf("expected DriverStreamer, got %v", sink.DriverType())
	}

	if sink.Streamer() != nil {
		t.Fatalf("expected nil streamer, got %v", sink.Streamer())
	}
}

func assertNilOperations(t *testing.T, sink *StreamerSink) {
	if err := sink.Sync(); err != nil {
		t.Fatalf("unexpected sync error: %v", err)
	}

	if err := sink.Close(); err != nil {
		t.Fatalf("unexpected close error: %v", err)
	}

	if err := sink.WriteEntry(LogEntry{}); err == nil {
		t.Fatalf("expected error writing to nil streamer")
	}
}

func TestStreamerSink_Writer(t *testing.T) {
	buf := &bytes.Buffer{}
	sink := NewStreamerSink(buf)
	if sink.Streamer() != buf {
		t.Fatalf("expected buffer instance")
	}

	assertBufferWrite(t, sink, buf)
	assertBufferSyncClose(t, sink)
}

func assertBufferWrite(t *testing.T, sink *StreamerSink, buf *bytes.Buffer) {
	entry := LogEntry{Message: "hello buffer", Level: logleveltype.Info}
	if err := sink.WriteEntry(entry); err != nil {
		t.Fatalf("failed to write entry: %v", err)
	}

	out := buf.String()
	hasMessage := strings.Contains(out, "hello buffer")
	if !hasMessage {
		t.Fatalf("expected log entry in buffer, got: %s", out)
	}
}

func assertBufferSyncClose(t *testing.T, sink *StreamerSink) {
	if err := sink.Sync(); err != nil {
		t.Fatalf("unexpected sync error: %v", err)
	}

	if err := sink.Close(); err != nil {
		t.Fatalf("unexpected close error: %v", err)
	}
}

type mockFaultStreamer struct {
	payloads []string
	isSynced bool
	isClosed bool
	failOp   string
}

func (m *mockFaultStreamer) Stream(ctx context.Context, payload string) *appfault.AppError {
	if m.failOp == "stream" {
		return appfault.New(errtype.Execution, "mock stream error")
	}

	m.payloads = append(m.payloads, payload)

	return nil
}

func (m *mockFaultStreamer) Sync() *appfault.AppError {
	if m.failOp == "sync" {
		return appfault.New(errtype.Execution, "mock sync error")
	}

	m.isSynced = true

	return nil
}

func (m *mockFaultStreamer) Close() *appfault.AppError {
	if m.failOp == "close" {
		return appfault.New(errtype.Execution, "mock close error")
	}

	m.isClosed = true

	return nil
}

func TestStreamerSink_FaultStreamer(t *testing.T) {
	mock := &mockFaultStreamer{}
	sink := NewStreamerSink(mock)
	testFaultStreamerSuccess(t, sink, mock)
	testFaultStreamerFailure(t)
}

func testFaultStreamerSuccess(t *testing.T, sink *StreamerSink, mock *mockFaultStreamer) {
	entry := LogEntry{Message: "stream msg"}
	if err := sink.WriteEntry(entry); err != nil {
		t.Fatalf("unexpected write error: %v", err)
	}

	if len(mock.payloads) != 1 {
		t.Fatalf("expected 1 payload, got %d", len(mock.payloads))
	}

	assertFaultStreamerSyncClose(t, sink, mock)
}

func assertFaultStreamerSyncClose(t *testing.T, sink *StreamerSink, mock *mockFaultStreamer) {
	if err := sink.Sync(); err != nil {
		t.Fatalf("unexpected sync error: %v", err)
	}

	if !mock.isSynced {
		t.Fatalf("expected mock to be synced")
	}

	if err := sink.Close(); err != nil {
		t.Fatalf("unexpected close error: %v", err)
	}

	if !mock.isClosed {
		t.Fatalf("expected mock to be closed")
	}
}

func testFaultStreamerFailure(t *testing.T) {
	mock := &mockFaultStreamer{failOp: "stream"}
	sink := NewStreamerSink(mock)
	if err := sink.WriteEntry(LogEntry{}); err == nil {
		t.Fatalf("expected stream error")
	}

	mock.failOp = "sync"
	if err := sink.Sync(); err == nil {
		t.Fatalf("expected sync error")
	}

	mock.failOp = "close"
	if err := sink.Close(); err == nil {
		t.Fatalf("expected close error")
	}
}

type mockErrStreamer struct {
	payloads []any
	isSynced bool
	isClosed bool
	failOp   string
}

func (m *mockErrStreamer) Stream(ctx context.Context, payload any) error {
	if m.failOp == "stream" {
		return errors.New("err stream")
	}

	m.payloads = append(m.payloads, payload)

	return nil
}

func (m *mockErrStreamer) Sync() error {
	if m.failOp == "sync" {
		return errors.New("err sync")
	}

	m.isSynced = true

	return nil
}

func (m *mockErrStreamer) Close() error {
	if m.failOp == "close" {
		return errors.New("err close")
	}

	m.isClosed = true

	return nil
}

func TestStreamerSink_ErrStreamer(t *testing.T) {
	mock := &mockErrStreamer{}
	sink := NewStreamerSink(mock)
	testErrStreamerSuccess(t, sink, mock)
	testErrStreamerFailure(t)
}

func testErrStreamerSuccess(t *testing.T, sink *StreamerSink, mock *mockErrStreamer) {
	if err := sink.WriteEntry(LogEntry{Message: "any msg"}); err != nil {
		t.Fatalf("unexpected write error: %v", err)
	}

	if len(mock.payloads) != 1 {
		t.Fatalf("expected 1 payload, got %d", len(mock.payloads))
	}

	if err := sink.Sync(); err != nil {
		t.Fatalf("unexpected sync error: %v", err)
	}

	if err := sink.Close(); err != nil {
		t.Fatalf("unexpected close error: %v", err)
	}
}

func testErrStreamerFailure(t *testing.T) {
	mock := &mockErrStreamer{failOp: "stream"}
	sink := NewStreamerSink(mock)
	if err := sink.WriteEntry(LogEntry{}); err == nil {
		t.Fatalf("expected stream error")
	}

	mock.failOp = "sync"
	if err := sink.Sync(); err == nil {
		t.Fatalf("expected sync error")
	}

	mock.failOp = "close"
	if err := sink.Close(); err == nil {
		t.Fatalf("expected close error")
	}
}

type duckStreamer struct {
	received string
}

func (d *duckStreamer) Stream(payload string) error {
	d.received = payload

	return nil
}

func TestStreamerSink_DuckTyped(t *testing.T) {
	duck := &duckStreamer{}
	sink := NewStreamerSink(duck)
	if err := sink.WriteEntry(LogEntry{Message: "duck test"}); err != nil {
		t.Fatalf("unexpected write error: %v", err)
	}

	hasMsg := strings.Contains(duck.received, "duck test")
	if !hasMsg {
		t.Fatalf("expected duck to receive payload")
	}
}

func TestStreamerSink_Unsupported(t *testing.T) {
	unsupportedSink := NewStreamerSink(12345)
	if err := unsupportedSink.WriteEntry(LogEntry{}); err == nil {
		t.Fatalf("expected error for unsupported streamer type")
	}
}

func TestStreamerSink_StreamwriterLockedStreamer(t *testing.T) {
	buf := &bytes.Buffer{}
	sw := streamwriter.NewLockedStreamer[string](streamwriter.LockedOptions[string]{
		Destination: buf,
	})
	sink := NewStreamerSink(sw)
	testStreamwriterSink(t, sink, buf)
}

func testStreamwriterSink(t *testing.T, sink *StreamerSink, buf *bytes.Buffer) {
	entry := LogEntry{Message: "real streamwriter test"}
	if err := sink.WriteEntry(entry); err != nil {
		t.Fatalf("unexpected write error: %v", err)
	}

	hasMsg := strings.Contains(buf.String(), "real streamwriter test")
	if !hasMsg {
		t.Fatalf("expected streamwriter to write to buffer")
	}

	if err := sink.Sync(); err != nil {
		t.Fatalf("unexpected sync error: %v", err)
	}

	if err := sink.Close(); err != nil {
		t.Fatalf("unexpected close error: %v", err)
	}
}

func TestStreamerSink_LogStreamerInterface(t *testing.T) {
	buf := &bytes.Buffer{}
	sink := NewStreamerSink(buf)

	var streamer LogStreamer = sink
	if streamer.Name() != "streamer" {
		t.Fatalf("expected name streamer, got %s", streamer.Name())
	}

	if streamer.Destination() != buf {
		t.Fatalf("expected destination to match buffer")
	}

	if sink.Unwrap() != buf {
		t.Fatalf("expected unwrap to match buffer")
	}

	testStreamerMethods(t, streamer, buf)
}

func testStreamerMethods(t *testing.T, s LogStreamer, buf *bytes.Buffer) {
	entry := LogEntry{Message: "stream-entry-direct"}
	if err := s.StreamEntry(entry); err != nil {
		t.Fatalf("unexpected stream entry error: %v", err)
	}

	fault := s.Stream(context.Background(), "raw-stream-payload")
	if fault != nil {
		t.Fatalf("unexpected stream fault: %v", fault)
	}

	if !strings.Contains(buf.String(), "stream-entry-direct") {
		t.Fatalf("expected buffer to contain stream entry")
	}

	if !strings.Contains(buf.String(), "raw-stream-payload") {
		t.Fatalf("expected buffer to contain raw stream payload")
	}
}
