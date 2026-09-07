package appwriter_test

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/appwriter"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/fileutil"
)

func initTestFileWriter(t *testing.T, logPath string) *appwriter.BaseWriter {
	wrap := appwriter.NewFileWriter(appwriter.FileWriterOptions{
		Name:     "test-logger",
		FilePath: logPath,
		OpenMode: fileutil.FileOpenCreateAppend,
		PermMode: fileutil.FilePermStandard,
		IsLocked: true,
	})
	if wrap.IsFailed() {
		t.Fatalf("expected success, got fault: %v", wrap.Fault())
	}

	return wrap.Data()
}

func verifyWriterBasics(t *testing.T, w *appwriter.BaseWriter) {
	if w.Name() != "test-logger" {
		t.Fatalf("expected name 'test-logger', got: %s", w.Name())
	}

	if !w.IsLocked() {
		t.Fatalf("expected writer to be locked")
	}
}

func verifyWriterOperations(t *testing.T, w *appwriter.BaseWriter, logPath string) {
	ctx := context.Background()
	if fault := w.Write(ctx, "event line 1\n"); fault != nil {
		t.Fatalf("write failed: %v", fault)
	}

	if syncFault := w.Sync(); syncFault != nil {
		t.Fatalf("sync failed: %v", syncFault)
	}

	verifyLogContent(t, logPath)
}

func verifyLogContent(t *testing.T, logPath string) {
	readRes := fileutil.ReadString(logPath)
	if readRes.IsFailed() {
		t.Fatalf("read failed: %v", readRes.Fault())
	}

	if readRes.Data() != "event line 1\n" {
		t.Fatalf("unexpected content: %s", readRes.Data())
	}
}

func TestNewFileWriter_Success(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "logs", "app.log")
	w := initTestFileWriter(t, logPath)
	defer w.Close()

	verifyWriterBasics(t, w)
	verifyWriterOperations(t, w, logPath)
}

func TestNewFileWriter_EmptyPathFailureWithId(t *testing.T) {
	wrap := appwriter.NewFileWriter(appwriter.FileWriterOptions{FilePath: ""})
	if wrap.IsSuccess() {
		t.Fatalf("expected failure for empty file path")
	}

	verifyEmptyPathFault(t, wrap.Fault())
}

func verifyEmptyPathFault(t *testing.T, fault *appfault.AppError) {
	if fault == nil {
		t.Fatalf("expected non-nil fault")
	}

	if fault.Type() != errtype.Validation {
		t.Fatalf("expected errtype.Validation, got: %v", fault.Type())
	}
}

func TestWrapWriterFailureFromWrap(t *testing.T) {
	failedFile := fileutil.OpenFile("", fileutil.FileOpenReadOnly, fileutil.FilePermStandard)
	writerWrap := appwriter.WrapWriterFailureFromWrap(failedFile)
	if writerWrap.IsSuccess() {
		t.Fatalf("expected failed writer wrap")
	}

	if writerWrap.Fault().Type() != errtype.Validation {
		t.Fatalf("expected validation error type in propagated fault")
	}
}

func initSharedLockerWriter(t *testing.T, logPath string) *appwriter.BaseWriter {
	wrap := appwriter.NewFileWriter(appwriter.FileWriterOptions{
		Name:     "shared-locker-test",
		FilePath: logPath,
		OpenMode: fileutil.FileOpenCreateAppend,
		PermMode: fileutil.FilePermStandard,
		IsLocked: true,
	})
	if wrap.IsFailed() {
		t.Fatalf("expected success, got: %v", wrap.Fault())
	}

	return wrap.Data()
}

func TestBaseWriter_SharedLocker(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "shared_locker.log")
	w := initSharedLockerWriter(t, logPath)
	defer w.Close()

	w.SharedLockerLock()
	w.SharedLockerUnlock()
	w.RLock()
	w.RUnlock()
}

func verifyWrapFailuresPart1(t *testing.T, appErr *appfault.AppError) {
	if appwriter.WrapWriter.Failure(appErr).IsSuccess() {
		t.Fatalf("expected failure")
	}

	if appwriter.WrapWriter.FailureFromError(appErr).IsSuccess() {
		t.Fatalf("expected failure")
	}

	if appwriter.WrapWriter.FailureWithId(errtype.IO, "io error").IsSuccess() {
		t.Fatalf("expected failure")
	}
}

func verifyWrapFailuresPart2(t *testing.T, appErr *appfault.AppError) {
	if appwriter.WrapWriter.FailureWithCause(errtype.IO, errors.New("underlying"), "cause error").IsSuccess() {
		t.Fatalf("expected failure")
	}

	fail3 := appwriter.WrapWriter.FailureWithId(errtype.IO, "io error")
	if appwriter.WrapWriter.FailureFromWrap(fail3).IsSuccess() {
		t.Fatalf("expected failure")
	}

	if appwriter.WriterWrap.Failure(appErr).IsSuccess() {
		t.Fatalf("expected failure")
	}

	if appwriter.Wrap.Failure(appErr).IsSuccess() {
		t.Fatalf("expected failure")
	}
}

func TestWriterWrapConstructor(t *testing.T) {
	if appwriter.WrapWriter.Success(nil).IsFailed() {
		t.Fatalf("expected success")
	}

	appErr := appfault.New(errtype.Validation, "validation error")
	verifyWrapFailuresPart1(t, appErr)
	verifyWrapFailuresPart2(t, appErr)
}
