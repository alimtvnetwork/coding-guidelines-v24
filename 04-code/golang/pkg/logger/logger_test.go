package logger_test

import (
	"bytes"
	"strings"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/logger"
)

func TestLoggerConsoleOutput(t *testing.T) {
	buf := &bytes.Buffer{}
	opts := logger.DefaultOptions().WithOutput(buf).WithLevel(logger.LevelDebug)
	log := logger.New(opts)

	log.Info("test info message")
	if !strings.Contains(buf.String(), "test info message") {
		t.Fatalf("expected log output to contain message, got %s", buf.String())
	}
}

func TestLoggerAppErrorLogging(t *testing.T) {
	buf := &bytes.Buffer{}
	opts := logger.DefaultOptions().WithOutput(buf).WithJson(true)
	log := logger.New(opts)

	appErr := appfault.NewWithDetails("db.find", "E2004", "record missing", "repo", appfault.ErrorTypeNotFound, appfault.SeverityError, nil).
		WithSiteId(101)
	log.LogError(appErr)

	output := buf.String()
	if !strings.Contains(output, "record missing") || !strings.Contains(output, "E2004") {
		t.Fatalf("expected JSON log to contain error details, got %s", output)
	}
}

func TestLoggerLevelFilterIgnored(t *testing.T) {
	buf := &bytes.Buffer{}
	log := logger.New(logger.DefaultOptions().WithOutput(buf).WithLevel(logger.LevelWarn))
	log.Debug("debug message")
	log.Info("info message")
	if buf.Len() > 0 {
		t.Fatalf("expected no output for levels below WARN, got %s", buf.String())
	}
}

func TestLoggerLevelFilterMatched(t *testing.T) {
	buf := &bytes.Buffer{}
	log := logger.New(logger.DefaultOptions().WithOutput(buf).WithLevel(logger.LevelWarn))
	log.Warn("warn message")
	if !strings.Contains(buf.String(), "warn message") {
		t.Fatalf("expected output for WARN level, got %s", buf.String())
	}
}
