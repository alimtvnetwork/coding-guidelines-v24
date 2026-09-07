package fileutil_test

import (
	"testing"

	"coding-guidelines/common/pkg/fileutil"
)

func TestFilePermType_Basics(t *testing.T) {
	p := fileutil.FilePermStandard
	if p.Uint32() != 0644 {
		t.Fatalf("expected uint32 0644, got %o", p.Uint32())
	}

	if p.OctalString() != "0644" {
		t.Fatalf("expected octal '0644', got %s", p.OctalString())
	}

	if p.PosixString() != "rw-r--r--" {
		t.Fatalf("expected posix 'rw-r--r--', got %s", p.PosixString())
	}

	if p.Mode() != 0644 {
		t.Fatalf("expected mode 0644, got %v", p.Mode())
	}
}

func TestFilePermType_Predicates(t *testing.T) {
	std := fileutil.FilePermStandard
	if !std.IsPublic() {
		t.Fatalf("expected Standard perm to be public")
	}

	if std.IsPrivate() {
		t.Fatalf("expected Standard perm to not be private")
	}

	if std.IsExecutable() {
		t.Fatalf("expected Standard perm to not be executable")
	}

	exec := fileutil.FilePermExecutable
	if !exec.IsExecutable() {
		t.Fatalf("expected Executable perm to be executable")
	}
}

func TestFilePermType_Mutators(t *testing.T) {
	std := fileutil.FilePermStandard
	priv := std.WithPrivate()
	if !priv.IsPrivate() {
		t.Fatalf("expected WithPrivate to result in private perm")
	}

	ro := std.WithReadOnly()
	if ro.Uint32() != 0444 {
		t.Fatalf("expected WithReadOnly 0444, got %o", ro.Uint32())
	}
}
