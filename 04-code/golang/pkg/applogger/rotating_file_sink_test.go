package applogger_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"coding-guidelines/common/pkg/applogger"
)

func TestRotatingFileSink_RotatesOnSizeThreshold(t *testing.T) {
	tempDir := t.TempDir()
	logFile := filepath.Join(tempDir, "app.log")

	cfg := applogger.RotationConfig{
		FilePath:         logFile,
		MaxSizeBytes:     200, // Small threshold for testing
		MaxBackups:       5,
		IsArchiveEnabled: false,
	}

	sinkRes := applogger.NewRotatingFileSink(cfg)
	if sinkRes.IsFailed() {
		t.Fatalf("failed to create sink: %v", sinkRes.Fault())
	}

	sink := sinkRes.Data()
	defer sink.Close()

	entry := applogger.LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     applogger.LevelInfo,
		Message:   "A payload string designed to exceed 200 bytes across iterations",
	}

	// Write 5 entries, which should cause multiple rotations
	for i := 0; i < 5; i++ {
		if writeErr := sink.WriteEntry(entry); writeErr != nil {
			t.Fatalf("write failed: %v", writeErr)
		}
	}

	_ = sink.Sync()

	files, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("failed to read dir: %v", err)
	}

	if len(files) < 2 {
		t.Fatalf("expected at least 2 files (active + backup), got %d", len(files))
	}
}

func TestRotatingFileSink_ArchivesAndCompresses(t *testing.T) {
	tempDir := t.TempDir()
	logFile := filepath.Join(tempDir, "app.log")
	archiveDir := filepath.Join(tempDir, "archives")

	cfg := applogger.RotationConfig{
		FilePath:         logFile,
		MaxSizeBytes:     150,
		MaxBackups:       3,
		IsArchiveEnabled: true,
		ArchiveDir:       archiveDir,
		IsCompress:       true,
	}

	sinkRes := applogger.NewRotatingFileSink(cfg)
	if sinkRes.IsFailed() {
		t.Fatalf("failed to create sink: %v", sinkRes.Fault())
	}

	sink := sinkRes.Data()
	defer sink.Close()

	entry := applogger.LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     applogger.LevelInfo,
		Message:   "Archive and compression test message with enough length",
	}

	for i := 0; i < 4; i++ {
		if writeErr := sink.WriteEntry(entry); writeErr != nil {
			t.Fatalf("write failed: %v", writeErr)
		}
	}

	archivedEntries, err := os.ReadDir(archiveDir)
	if err != nil {
		t.Fatalf("failed to read archive dir: %v", err)
	}

	if len(archivedEntries) == 0 {
		t.Fatalf("expected at least 1 archived file, got 0")
	}

	hasGz := false
	for _, f := range archivedEntries {
		if filepath.Ext(f.Name()) == ".gz" {
			hasGz = true
			break
		}
	}

	if !hasGz {
		t.Fatalf("expected compressed .gz file in archive")
	}
}

func TestRotatingFileSink_PrunesExcessBackups(t *testing.T) {
	tempDir := t.TempDir()
	logFile := filepath.Join(tempDir, "app.log")

	cfg := applogger.RotationConfig{
		FilePath:         logFile,
		MaxSizeBytes:     100,
		MaxBackups:       2,
		IsArchiveEnabled: false,
	}

	sinkRes := applogger.NewRotatingFileSink(cfg)
	if sinkRes.IsFailed() {
		t.Fatalf("failed to create sink: %v", sinkRes.Fault())
	}

	sink := sinkRes.Data()
	defer sink.Close()

	entry := applogger.LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     applogger.LevelInfo,
		Message:   "Pruning test message payload",
	}

	for i := 0; i < 10; i++ {
		_ = sink.WriteEntry(entry)
	}

	files, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("failed to read dir: %v", err)
	}

	// Active file + at most MaxBackups
	if len(files) > 3 {
		t.Fatalf("expected at most 3 files (1 active + 2 backups), got %d", len(files))
	}
}

func TestRotatingFileSink_DriverCreation(t *testing.T) {
	tempDir := t.TempDir()
	logFile := filepath.Join(tempDir, "driver.log")

	cfg := applogger.Config{
		MinLevel: applogger.LevelInfo,
		Driver:   applogger.DriverRotatingFile,
		Rotation: applogger.RotationConfig{
			FilePath:     logFile,
			MaxSizeBytes: 1024,
			MaxBackups:   5,
		},
	}

	loggerRes := applogger.New(cfg)
	if loggerRes.IsFailed() {
		t.Fatalf("failed to create logger with DriverRotatingFile: %v", loggerRes.Fault())
	}

	l := loggerRes.Data()

	l.Info("testing rotating file driver")
	_ = l.Close()

	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Fatalf("expected log file to exist at %s", logFile)
	}
}
