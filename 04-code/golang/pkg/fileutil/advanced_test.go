package fileutil_test

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/enum/openfiletype"
	"coding-guidelines/common/pkg/fileutil"
)

func verifyInitialAtomicWrite(t *testing.T, targetFile string) {
	initialData := []byte("first atomic revision")
	res := fileutil.WriteAtomic(targetFile, initialData, filepermtype.Standard)
	if res.IsFailed() {
		t.Fatalf("WriteAtomic failed: %v", res.Fault())
	}

	readRes := fileutil.ReadString(targetFile)
	if readRes.IsFailed() || readRes.Data() != "first atomic revision" {
		t.Fatalf("unexpected read data: %v", readRes)
	}
}

func verifyUpdatedAtomicWrite(t *testing.T, targetFile string) {
	updatedData := []byte("second atomic revision - updated cleanly")
	res := fileutil.WriteAtomic(targetFile, updatedData, filepermtype.Standard)
	if res.IsFailed() {
		t.Fatalf("WriteAtomic overwrite failed: %v", res.Fault())
	}

	readRes := fileutil.ReadString(targetFile)
	if readRes.IsFailed() || readRes.Data() != "second atomic revision - updated cleanly" {
		t.Fatalf("unexpected updated data: %v", readRes)
	}
}

func TestWriteAtomic(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "atomic-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	defer os.RemoveAll(tempDir)

	targetFile := filepath.Join(tempDir, "atomic-test.txt")
	verifyInitialAtomicWrite(t, targetFile)
	verifyUpdatedAtomicWrite(t, targetFile)
}

func performChunkedRead(t *testing.T, targetFile string, payload []byte) {
	var accumulated bytes.Buffer
	chunkCount := 0

	readRes := fileutil.ReadChunked(targetFile, 4096, func(chunk []byte) *appfault.AppError {
		chunkCount++
		accumulated.Write(chunk)

		return nil
	})

	if readRes.IsFailed() || readRes.Data() != int64(len(payload)) || chunkCount != 4 {
		t.Fatalf("chunked read failed or size mismatch: %v", readRes)
	}

	if !bytes.Equal(accumulated.Bytes(), payload) {
		t.Fatalf("accumulated bytes do not match original payload")
	}
}

func TestReadChunked(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "chunked-read-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	defer os.RemoveAll(tempDir)

	targetFile := filepath.Join(tempDir, "chunked.bin")
	payload := bytes.Repeat([]byte("0123456789ABCDEF"), 1024)
	if res := fileutil.WriteFile(targetFile, payload, filepermtype.Standard); res.IsFailed() {
		t.Fatalf("WriteFile failed: %v", res.Fault())
	}

	performChunkedRead(t, targetFile, payload)
}

func performWriteChunked(t *testing.T, targetFile string, payload []byte) {
	reader := bytes.NewReader(payload)
	res := fileutil.WriteChunked(targetFile, filepermtype.Standard, reader, 2048)
	if res.IsFailed() || res.Data() != int64(len(payload)) {
		t.Fatalf("WriteChunked failed: %v", res)
	}

	readRes := fileutil.ReadAll(targetFile)
	if readRes.IsFailed() || !bytes.Equal(readRes.Data(), payload) {
		t.Fatalf("read bytes do not match written payload")
	}
}

func TestWriteChunked(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "chunked-write-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	defer os.RemoveAll(tempDir)

	targetFile := filepath.Join(tempDir, "write-chunked.bin")
	payload := bytes.Repeat([]byte("CHUNKED-DATA-BLOCK"), 512)
	performWriteChunked(t, targetFile, payload)
}

func verifyFileWriterOutput(t *testing.T, targetFile string) {
	readRes := fileutil.ReadString(targetFile)
	if readRes.IsFailed() {
		t.Fatalf("ReadString failed: %v", readRes.Fault())
	}

	if len(readRes.Data()) == 0 {
		t.Fatalf("expected non-empty written file")
	}
}

func performFileWriterOperations(t *testing.T, targetFile string) {
	writerRes := fileutil.NewFileWriter(targetFile, openfiletype.CreateAppend, filepermtype.Standard)
	if writerRes.IsFailed() {
		t.Fatalf("NewFileWriter failed: %v", writerRes.Fault())
	}

	writer := writerRes.Data()
	defer writer.Close()

	if appErr := writer.Write(context.Background(), "log event 1"); appErr != nil {
		t.Fatalf("writer.Write failed: %v", appErr)
	}
}

func TestNewFileWriter(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "file-writer-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	defer os.RemoveAll(tempDir)

	targetFile := filepath.Join(tempDir, "writer-output.txt")
	performFileWriterOperations(t, targetFile)
	verifyFileWriterOutput(t, targetFile)
}
