package applogger_test

import (
	"bytes"
	"strings"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/appfaults"
	"coding-guidelines/common/pkg/applogger"
	"coding-guidelines/common/pkg/errtype"
)

func newTestConsoleLogger(buf *bytes.Buffer, isUseJSON bool) applogger.Logger {
	sink := applogger.NewConsoleSink(buf, isUseJSON)
	cfg := applogger.Config{
		MinLevel: applogger.LevelDebug,
		Driver:   applogger.DriverComposite,
		Sinks:    []applogger.LogSink{sink},
	}

	res := applogger.New(cfg)
	if res.IsFailed() {
		panic(res.Fault())
	}

	return res.Data()
}

func TestConsoleSinkLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	l := newTestConsoleLogger(buf, false)
	l.Infof("hello %s", "world")

	if !strings.Contains(buf.String(), "hello world") {
		t.Fatalf("expected log output, got %s", buf.String())
	}
}

func TestLogErrorAndLogFaults(t *testing.T) {
	buf := &bytes.Buffer{}
	l := newTestConsoleLogger(buf, true)
	err := appfault.New(errtype.Database, "connection reset").WithOp("db.connect")

	l.LogError(err)
	l.LogFaults(appfaults.New().Add(err))

	if !strings.Contains(buf.String(), "connection reset") {
		t.Fatalf("expected error message in log output, got %s", buf.String())
	}
}

func TestLoggerChaining(t *testing.T) {
	buf := &bytes.Buffer{}
	l := newTestConsoleLogger(buf, false)
	err := appfault.New(errtype.Database, "connection reset")

	ret := l.Debug("chain debug").
		Info("chain info").
		Warn("chain warn").
		Error("chain error").
		Debugf("chain debugf %d", 1).
		Infof("chain infof %d", 2).
		Warnf("chain warnf %d", 3).
		Errorf("chain errorf %d", 4).
		WithContext("key", "val").
		WithFields(map[string]any{"env": "test"}).
		LogError(err).
		LogFaults(appfaults.New().Add(err))

	if ret == nil {
		t.Fatalf("expected non-nil logger from chain")
	}

	out := buf.String()
	if !strings.Contains(out, "chain info") || !strings.Contains(out, "chain errorf 4") {
		t.Fatalf("expected chained log output, got %s", out)
	}
}

func TestLogger_SinkIntrospection(t *testing.T) {
	cfg := applogger.Config{
		MinLevel:  applogger.LevelInfo,
		Driver:    applogger.DriverApi,
		Endpoint:  "https://api.example.com/logs",
		IsUseJSON: true,
	}

	logRes := applogger.New(cfg)
	if logRes.IsFailed() {
		t.Fatalf("expected logger success, got fault: %v", logRes.Fault())
	}

	l := logRes.Data()
	if l.Type() != applogger.DriverApi {
		t.Errorf("expected DriverApi, got %v", l.Type())
	}

	if l.EndpointPath() != "https://api.example.com/logs" {
		t.Errorf("expected endpoint, got %s", l.EndpointPath())
	}

	if l.EndPointPath() != l.EndpointPath() {
		t.Errorf("expected EndPointPath alias match")
	}
}

func TestLogger_Clone(t *testing.T) {
	buf := &bytes.Buffer{}
	l := newTestConsoleLogger(buf, false)
	l2 := l.Clone()
	if l2 == nil {
		t.Fatalf("expected non-nil cloned logger")
	}

	child := l2.WithContext("worker", "1")
	child.Infof("worker log")
	if !strings.Contains(buf.String(), "worker log") {
		t.Errorf("expected log output from cloned logger")
	}
}

func TestLogger_AddWriters(t *testing.T) {
	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}
	l := newTestConsoleLogger(buf1, false)
	sink2 := applogger.NewConsoleSink(buf2, false)

	multi := l.AddWriters(sink2)
	multi.Infof("multicast message")

	hasBuf1 := strings.Contains(buf1.String(), "multicast message")
	hasBuf2 := strings.Contains(buf2.String(), "multicast message")
	if !hasBuf1 || !hasBuf2 {
		t.Fatalf("expected multicast write to both buffers")
	}
}

func TestLogger_AddStreamer(t *testing.T) {
	buf1 := &bytes.Buffer{}
	streamBuf := &bytes.Buffer{}
	l := newTestConsoleLogger(buf1, false)

	streamLogger := l.AddStreamer(streamBuf)
	streamLogger.Infof("stream hello")

	if !strings.Contains(streamBuf.String(), "stream hello") {
		t.Fatalf("expected stream buffer to receive log entry, got %s", streamBuf.String())
	}

	nilStreamer := l.AddStreamer(nil)
	if nilStreamer == nil {
		t.Fatalf("expected non-nil logger from nil streamer add")
	}
}

func TestLogger_WritersAndWriterNames(t *testing.T) {
	buf1, buf2 := &bytes.Buffer{}, &bytes.Buffer{}
	l := newTestConsoleLogger(buf1, false)
	sink2 := applogger.NewConsoleSink(buf2, false)
	multi := l.AddWriters(sink2)

	writers := multi.Writers()
	if len(writers) != 2 {
		t.Fatalf("expected 2 writers, got %d", len(writers))
	}

	names := multi.WriterNames()
	if len(names) != 2 || names[0] != "console" || names[1] != "console" {
		t.Fatalf("unexpected writer names: %v", names)
	}
}

func TestLogger_StreamersIntrospection(t *testing.T) {
	buf, streamBuf := &bytes.Buffer{}, &bytes.Buffer{}
	l := newTestConsoleLogger(buf, false)
	streamLogger := l.AddStreamer(streamBuf)

	streamers := streamLogger.Streamers()
	if len(streamers) != 1 {
		t.Fatalf("expected 1 streamer, got %d", len(streamers))
	}

	if streamers[0] != streamBuf {
		t.Fatalf("expected stream buffer instance in streamers list")
	}
}

func TestSinks_NameMethod(t *testing.T) {
	cs := applogger.NewConsoleSink(nil, false)
	if cs.Name() != "console" {
		t.Errorf("expected console, got %s", cs.Name())
	}

	ss := applogger.NewStreamerSink(&bytes.Buffer{})
	if ss.Name() != "streamer" {
		t.Errorf("expected streamer, got %s", ss.Name())
	}

	comp := applogger.NewCompositeSink(cs, ss)
	if comp.Name() != "composite" {
		t.Errorf("expected composite, got %s", comp.Name())
	}
}
