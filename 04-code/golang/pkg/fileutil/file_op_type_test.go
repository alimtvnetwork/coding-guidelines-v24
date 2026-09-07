package fileutil_test

import (
	"testing"

	"coding-guidelines/common/pkg/fileutil"
)

func TestFileOpType_Properties(t *testing.T) {
	op := fileutil.FileOpDelete
	if !op.IsDelete() {
		t.Fatalf("expected FileOpDelete to have IsDelete true")
	}

	if op.IsReadOnly() {
		t.Fatalf("expected FileOpDelete to have IsReadOnly false")
	}

	if op.Name() != "Delete" || op.String() != "Delete" {
		t.Fatalf("expected Name/String 'Delete', got %s", op.Name())
	}

	readOp := fileutil.FileOpReadOnly
	if !readOp.IsReadOnly() {
		t.Fatalf("expected FileOpReadOnly to have IsReadOnly true")
	}

	if readOp.OpenMode() != fileutil.FileOpenReadOnly {
		t.Fatalf("expected FileOpenReadOnly, got %v", readOp.OpenMode())
	}

	appendOp := fileutil.FileOpAppend
	if !appendOp.IsAppend() {
		t.Fatalf("expected FileOpAppend to have IsAppend true")
	}

	createAppendOp := fileutil.FileOpCreateAppend
	if !createAppendOp.IsAppend() {
		t.Fatalf("expected FileOpCreateAppend to have IsAppend true")
	}
}
