package fileutil_test

import (
	"testing"

	"coding-guidelines/common/pkg/enum/filewritemodetype"
)

func TestFileWriteModeType_Basics(t *testing.T) {
	mode := filewritemodetype.Atomic
	if !mode.IsValid() {
		t.Fatalf("expected Atomic to be valid")
	}

	if mode.Name() != "Atomic" || mode.String() != "Atomic" {
		t.Fatalf("expected Name/String 'Atomic', got %s", mode.Name())
	}

	direct := filewritemodetype.Direct
	if direct.Name() != "Direct" {
		t.Fatalf("expected Direct, got %s", direct.Name())
	}

	trunc := filewritemodetype.Truncate
	if trunc.Name() != "Truncate" {
		t.Fatalf("expected Truncate, got %s", trunc.Name())
	}

	invalid := filewritemodetype.Variant(99)
	if invalid.IsValid() {
		t.Fatalf("expected invalid mode to return IsValid false")
	}
}
