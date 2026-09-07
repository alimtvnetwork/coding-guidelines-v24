package fileutil

import (
	"os"
	"strings"
	"testing"
)

func TestToSlashAndBackslash(t *testing.T) {
	p := `a\b\c/d`
	slashed := ToSlash(p)
	if strings.Contains(slashed, `\`) {
		t.Errorf("expected no backslashes in ToSlash, got %s", slashed)
	}

	backslashed := ToBackslash(p)
	if strings.Contains(backslashed, `/`) {
		t.Errorf("expected no slashes in ToBackslash, got %s", backslashed)
	}
}

func TestToNative(t *testing.T) {
	p := `a/b\c`
	native := ToNative(p)
	if os.PathSeparator == '/' {
		if strings.Contains(native, `\`) {
			t.Errorf("expected slash on Unix, got %s", native)
		}
	} else {
		if strings.Contains(native, `/`) {
			t.Errorf("expected backslash on Windows, got %s", native)
		}
	}
}

func TestLongPathPrefix(t *testing.T) {
	prefixed := `\\?\C:\foo\bar`
	if !HasLongPathPrefix(prefixed) {
		t.Error("expected true for HasLongPathPrefix")
	}

	trimmed := TrimLongPathPrefix(prefixed)
	if trimmed != `C:\foo\bar` {
		t.Errorf("expected C:\\foo\\bar, got %s", trimmed)
	}

	normal := `C:\foo\bar`
	if HasLongPathPrefix(normal) {
		t.Error("expected false for normal path")
	}
}

func TestToLongPathDrive(t *testing.T) {
	winPath := `C:\Windows\System32`
	longWin := ToLongPath(winPath)
	if !strings.HasPrefix(longWin, `\\?\`) {
		t.Errorf("expected long path prefix, got %s", longWin)
	}

	alreadyLong := `\\?\C:\already`
	if ToLongPath(alreadyLong) != alreadyLong {
		t.Errorf("expected identical return for already long path")
	}
}

func TestToLongPathUNC(t *testing.T) {
	uncPath := `\\server\share\file`
	longUnc := ToLongPath(uncPath)
	if !strings.HasPrefix(longUnc, `\\?\UNC\`) {
		t.Errorf("expected UNC prefix, got %s", longUnc)
	}
}

func TestDeduplicateSeparators(t *testing.T) {
	in1 := `a//b///c////d`
	expected1 := `a/b/c/d`
	if got := DeduplicateSeparators(in1); got != expected1 {
		t.Errorf("expected %s, got %s", expected1, got)
	}

	in2 := `a\\\\b\\\c`
	expected2 := `a\b\c`
	if got := DeduplicateSeparators(in2); got != expected2 {
		t.Errorf("expected %s, got %s", expected2, got)
	}
}

func TestDeduplicatePreservesUNC(t *testing.T) {
	unc := `\\server\\share\\\file`
	expected := `\\server\share\file`
	if got := DeduplicateSeparators(unc); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}

	longPref := `\\?\C:\\foo\\\bar`
	expectedLong := `\\?\C:\foo\bar`
	if got := DeduplicateSeparators(longPref); got != expectedLong {
		t.Errorf("expected %s, got %s", expectedLong, got)
	}
}

func TestClean(t *testing.T) {
	p := `foo/bar/../baz`
	cleaned := Clean(p)
	if !strings.HasSuffix(ToSlash(cleaned), "foo/baz") {
		t.Errorf("expected cleaned to end with foo/baz, got %s", cleaned)
	}

	longP := `\\?\C:\foo\bar\..\baz`
	cleanedLong := Clean(longP)
	if !HasLongPathPrefix(cleanedLong) {
		t.Errorf("expected Clean to preserve long path prefix, got %s", cleanedLong)
	}
}

func TestNormalize(t *testing.T) {
	raw := `foo//bar/../baz`
	res := Normalize(raw)
	if res.IsFailure() {
		t.Fatalf("Normalize failed: %v", res.Fault())
	}

	if strings.Contains(res.Data(), "//") {
		t.Errorf("expected no duplicate separators, got %s", res.Data())
	}
}

func TestNormalizeToSlash(t *testing.T) {
	resSlash := NormalizeToSlash(`foo\\bar//baz`)
	if resSlash.IsFailure() {
		t.Fatalf("NormalizeToSlash failed: %v", resSlash.Fault())
	}

	if strings.Contains(resSlash.Data(), `\`) {
		t.Errorf("expected no backslash in NormalizeToSlash, got %s", resSlash.Data())
	}
}
