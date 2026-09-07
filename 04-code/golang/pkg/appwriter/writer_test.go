package appwriter_test

import (
	"context"
	"path/filepath"
	"sync"
	"testing"

	"coding-guidelines/common/pkg/appwriter"
	"coding-guidelines/common/pkg/errtype"
)

func TestFileWriter_ClosedDestinationReturnsFault(t *testing.T) {
	wrap := appwriter.NewFileWriter(appwriter.FileWriterOptions{
		FilePath: filepath.Join(t.TempDir(), "closed.log"),
	})
	if wrap.IsFailed() {
		t.Fatalf("failed to create writer: %v", wrap.Fault())
	}

	w := wrap.Data()
	if closeErr := w.Close(); closeErr != nil {
		t.Fatalf("failed to close: %v", closeErr)
	}

	verifyClosedWriteFault(t, w)
}

func verifyClosedWriteFault(t *testing.T, w *appwriter.BaseWriter) {
	fault := w.Write(context.Background(), "after close")
	if fault == nil {
		t.Fatalf("expected fault writing to closed destination")
	}

	if fault.Type() != errtype.IO {
		t.Fatalf("expected errtype.IO, got: %v", fault.Type())
	}
}

func TestFileWriteFunc_PayloadTypes(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "payloads.log")
	wrap := appwriter.NewFileWriter(appwriter.FileWriterOptions{
		FilePath: logPath,
		IsLocked: true,
	})
	if wrap.IsFailed() {
		t.Fatalf("init failed: %v", wrap.Fault())
	}

	defer wrap.Data().Close()

	verifyPayloadWrites(t, wrap.Data(), logPath)
}

func verifyPayloadWrites(t *testing.T, w *appwriter.BaseWriter, logPath string) {
	ctx := context.Background()
	if err := w.Write(ctx, []byte("raw-bytes\n")); err != nil {
		t.Fatalf("bytes write failed: %v", err)
	}

	payload := map[string]string{"key": "value"}
	if err := w.Write(ctx, payload); err != nil {
		t.Fatalf("map write failed: %v", err)
	}

	if syncErr := w.Sync(); syncErr != nil {
		t.Fatalf("sync failed: %v", syncErr)
	}
}

func TestFileWriter_ConcurrentLockedWrites(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "concurrent.log")
	wrap := appwriter.NewFileWriter(appwriter.FileWriterOptions{
		FilePath: logPath,
		IsLocked: true,
	})
	if wrap.IsFailed() {
		t.Fatalf("init failed: %v", wrap.Fault())
	}

	defer wrap.Data().Close()

	runConcurrentWriterLoops(wrap.Data(), 20)
}

func runConcurrentWriterLoops(w *appwriter.BaseWriter, routines int) {
	var wg sync.WaitGroup
	ctx := context.Background()
	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = w.Write(ctx, "concurrent line\n")
		}()
	}

	wg.Wait()
}
