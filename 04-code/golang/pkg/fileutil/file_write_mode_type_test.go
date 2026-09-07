package fileutil_test

import (
	"testing"

	"coding-guidelines/common/pkg/fileutil"
)

func TestFileWriteModeType_Basics(t *testing.T) {
	mode := fileutil.FileWriteModeAtomic
	if !mode.IsValid() {
		t.Fatalf("expected FileWriteModeAtomic to be valid")
	}

	if mode.Name() != "Atomic" || mode.String() != "Atomic" {
		t.Fatalf("expected Name/String 'Atomic', got %s", mode.Name())
	}

	direct := fileutil.FileWriteModeDirect
	if direct.Name() != "Direct" {
		t.Fatalf("expected Direct, got %s", direct.Name())
	}

	trunc := fileutil.FileWriteModeTruncate
	if trunc.Name() != "Truncate" {
		t.Fatalf("expected Truncate, got %s", trunc.Name())
	}

	invalid := fileutil.FileWriteModeType(99)
	if invalid.IsValid() {
		t.Fatalf("expected invalid mode to return IsValid false")
	}
}
