package fileutil_test

import (
	"testing"

	"coding-guidelines/common/pkg/enum/fileoptype"
	"coding-guidelines/common/pkg/enum/openfiletype"
)

func TestFileOpType_Properties(t *testing.T) {
	op := fileoptype.Delete
	if !op.IsDelete() {
		t.Fatalf("expected Delete to have IsDelete true")
	}

	if op.IsReadOnly() {
		t.Fatalf("expected Delete to have IsReadOnly false")
	}

	if op.Name() != "Delete" || op.String() != "Delete" {
		t.Fatalf("expected Name/String 'Delete', got %s", op.Name())
	}

	readOp := fileoptype.ReadOnly
	if !readOp.IsReadOnly() {
		t.Fatalf("expected ReadOnly to have IsReadOnly true")
	}

	if readOp.OpenMode() != openfiletype.ReadOnly {
		t.Fatalf("expected openfiletype.ReadOnly, got %v", readOp.OpenMode())
	}

	appendOp := fileoptype.Append
	if !appendOp.IsAppend() {
		t.Fatalf("expected Append to have IsAppend true")
	}

	createAppendOp := fileoptype.CreateAppend
	if !createAppendOp.IsAppend() {
		t.Fatalf("expected CreateAppend to have IsAppend true")
	}
}
