package fileutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/fileutil"
)

func TestAtomicWriteFile_Success(t *testing.T) {
	tempDir := t.TempDir()
	target := filepath.Join(tempDir, "atomic_test.txt")
	content := []byte("hello atomic world")

	fault := fileutil.AtomicWriteFile(target, content, filepermtype.Standard)
	if fault != nil {
		t.Fatalf("atomic write failed: %s", fault.Message())
	}

	readBack, err := os.ReadFile(target)
	if err != nil {
		t.Fatalf("failed reading target file: %v", err)
	}

	if string(readBack) != string(content) {
		t.Fatalf("content mismatch: got %s, want %s", string(readBack), string(content))
	}
}

func TestAtomicWriteFile_OverwritesExisting(t *testing.T) {
	tempDir := t.TempDir()
	target := filepath.Join(tempDir, "overwrite_test.txt")
	_ = os.WriteFile(target, []byte("initial content"), 0o600)

	newContent := []byte("replaced content successfully")
	fault := fileutil.AtomicWriteFile(target, newContent, filepermtype.Standard)
	if fault != nil {
		t.Fatalf("atomic overwrite failed: %s", fault.Message())
	}

	readBack, _ := os.ReadFile(target)
	if string(readBack) != string(newContent) {
		t.Fatalf("overwrite failed: got %s, want %s", string(readBack), string(newContent))
	}

	verifyNoTempFilesLeft(t, tempDir)
}

func verifyNoTempFilesLeft(t *testing.T, dir string) {
	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		if filepath.Ext(e.Name()) == ".tmp" || len(e.Name()) > 4 && e.Name()[:4] == ".tmp" {
			t.Fatalf("orphan temp file found: %s", e.Name())
		}
	}
}

func TestAtomicWriteFile_CreatesMissingDirectories(t *testing.T) {
	tempDir := t.TempDir()
	nested := filepath.Join(tempDir, "a", "b", "c", "nested.txt")
	data := []byte("nested atomic data")

	fault := fileutil.AtomicWriteFile(nested, data, filepermtype.Standard)
	if fault != nil {
		t.Fatalf("atomic write to nested path failed: %s", fault.Message())
	}

	readBack, err := os.ReadFile(nested)
	if err != nil {
		t.Fatalf("failed to read nested file: %v", err)
	}

	if string(readBack) != string(data) {
		t.Fatalf("nested data mismatch")
	}
}

func TestAtomicWriteFile_Validation(t *testing.T) {
	fault := fileutil.AtomicWriteFile("", []byte("test"), filepermtype.Standard)
	if fault == nil {
		t.Fatal("expected error on empty path")
	}
}

func TestAtomicWrite_BoolResult(t *testing.T) {
	tempDir := t.TempDir()
	target := filepath.Join(tempDir, "bool_result_test.txt")
	res := fileutil.AtomicWrite(target, []byte("bool test"), filepermtype.Standard)

	if res.IsFailure() {
		t.Fatalf("expected success, got fault: %v", res.Fault())
	}

	if !res.Data() {
		t.Fatal("expected true data in BoolResult")
	}
}
