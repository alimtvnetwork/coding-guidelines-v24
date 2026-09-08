package applogger_test

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"coding-guidelines/common/pkg/applogger"
	"coding-guidelines/common/pkg/enum/logleveltype"
)

func TestRotatingFileSink_E2E_RotationAndCompression(t *testing.T) {
	tempDir := t.TempDir()
	archiveDir := filepath.Join(tempDir, "archives")
	logPath := filepath.Join(tempDir, "app.log")

	cfg := applogger.RotationConfig{
		FilePath:         logPath,
		MaxSizeBytes:     250,
		MaxBackups:       2,
		IsArchiveEnabled: true,
		ArchiveDir:       archiveDir,
		IsCompress:       true,
	}

	sink, err := applogger.NewRotatingFileSink(cfg)
	if err != nil {
		t.Fatalf("failed to instantiate rotating sink: %v", err)
	}

	writeBatchOfEntries(t, sink, 8)
	_ = sink.Sync()
	_ = sink.Close()

	verifyArchiveDirectoryContents(t, archiveDir, 2)
	verifyActiveLogState(t, logPath)
}

func writeBatchOfEntries(t *testing.T, sink *applogger.RotatingFileSink, count int) {
	for i := 0; i < count; i++ {
		writeSingleEntry(t, sink, i)
		time.Sleep(5 * time.Millisecond)
	}
}

func writeSingleEntry(t *testing.T, sink *applogger.RotatingFileSink, idx int) {
	entry := applogger.LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		Level:     logleveltype.Info,
		Message:   fmt.Sprintf("e2e log entry message number %d with verbose content", idx),
	}

	if err := sink.WriteEntry(entry); err != nil {
		t.Fatalf("write failed on entry %d: %v", idx, err)
	}
}

func verifyArchiveDirectoryContents(t *testing.T, archiveDir string, maxExpected int) {
	entries, err := os.ReadDir(archiveDir)
	if err != nil {
		t.Fatalf("failed to read archive directory: %v", err)
	}

	if len(entries) == 0 {
		t.Fatal("expected rotated archive files, found 0")
	}

	if len(entries) > maxExpected {
		t.Fatalf("expected at most %d backups, got %d", maxExpected, len(entries))
	}

	for _, e := range entries {
		verifyGzipArchiveEntry(t, filepath.Join(archiveDir, e.Name()))
	}
}

func verifyGzipArchiveEntry(t *testing.T, archivePath string) {
	if !strings.HasSuffix(archivePath, ".gz") {
		t.Fatalf("expected .gz extension on archive file: %s", archivePath)
	}

	f, err := os.Open(archivePath)
	if err != nil {
		t.Fatalf("failed to open gzip archive: %v", err)
	}

	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		t.Fatalf("invalid gzip reader: %v", err)
	}

	defer gz.Close()

	content, err := io.ReadAll(gz)
	if err != nil || len(content) == 0 {
		t.Fatalf("failed reading decompressed archive: %v", err)
	}
}

func verifyActiveLogState(t *testing.T, logPath string) {
	info, err := os.Stat(logPath)
	if err != nil {
		t.Fatalf("active log file does not exist: %v", err)
	}

	if info.Size() == 0 {
		t.Fatal("expected active log file to contain current entries")
	}
}
