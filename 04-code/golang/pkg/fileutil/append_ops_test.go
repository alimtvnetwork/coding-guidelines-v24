package fileutil

import (
	"path/filepath"
	"testing"

	"coding-guidelines/common/pkg/enum/filepermtype"
)

func verifyAppendedText(t *testing.T, path string, expected string) {
	readRes := ReadText(path)
	if !readRes.IsSuccess() {
		t.Fatalf("read text failed: %v", readRes.Fault())
	}

	if readRes.Data() != expected {
		t.Fatalf("expected %q, got %q", expected, readRes.Data())
	}
}

func verifyAppendedLines(t *testing.T, path string, count int, lastElem string) {
	readRes := ReadLines(path)
	if !readRes.IsSuccess() {
		t.Fatalf("read lines failed: %v", readRes.Fault())
	}

	lines := readRes.Data()
	if len(lines) != count {
		t.Fatalf("expected %d lines, got %d", count, len(lines))
	}

	if lines[count-1] != lastElem {
		t.Fatalf("expected last line %q, got %q", lastElem, lines[count-1])
	}
}

func TestAppendBytes(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "append_bytes.txt")

	res1 := AppendBytes(path, []byte("hello "), filepermtype.Standard)
	res2 := AppendBytes(path, []byte("world"), filepermtype.Standard)
	if !res1.IsSuccess() || !res2.IsSuccess() {
		t.Fatalf("append bytes failed: %v / %v", res1.Fault(), res2.Fault())
	}

	verifyAppendedText(t, path, "hello world")
}

func TestAppendBytesLocked(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "append_bytes_locked.txt")

	res1 := AppendBytesLocked(path, []byte("part1 "), filepermtype.Standard)
	res2 := AppendBytesLocked(path, []byte("part2"), filepermtype.Standard)
	if !res1.IsSuccess() || !res2.IsSuccess() {
		t.Fatalf("append bytes locked failed: %v / %v", res1.Fault(), res2.Fault())
	}

	verifyAppendedText(t, path, "part1 part2")
}

func TestAppendString(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "append_string.txt")

	res1 := AppendString(path, "first ", filepermtype.Standard)
	res2 := AppendString(path, "second", filepermtype.Standard)
	if !res1.IsSuccess() || !res2.IsSuccess() {
		t.Fatalf("append string failed: %v / %v", res1.Fault(), res2.Fault())
	}

	verifyAppendedText(t, path, "first second")
}

func TestAppendStringLocked(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "append_string_locked.txt")

	res1 := AppendStringLocked(path, "lock_a ", filepermtype.Standard)
	res2 := AppendStringLocked(path, "lock_b", filepermtype.Standard)
	if !res1.IsSuccess() || !res2.IsSuccess() {
		t.Fatalf("append string locked failed: %v / %v", res1.Fault(), res2.Fault())
	}

	verifyAppendedText(t, path, "lock_a lock_b")
}

func TestAppendLines(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "append_lines.txt")

	res1 := AppendLines(path, []string{"row1", "row2"}, filepermtype.Standard)
	res2 := AppendLines(path, []string{"row3"}, filepermtype.Standard)
	if !res1.IsSuccess() || !res2.IsSuccess() {
		t.Fatalf("append lines failed: %v / %v", res1.Fault(), res2.Fault())
	}

	verifyAppendedLines(t, path, 3, "row3")
}

func TestAppendLinesLocked(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "append_lines_locked.txt")

	res1 := AppendLinesLocked(path, []string{"alpha", "beta"}, filepermtype.Standard)
	res2 := AppendLinesLocked(path, []string{"gamma"}, filepermtype.Standard)
	if !res1.IsSuccess() || !res2.IsSuccess() {
		t.Fatalf("append lines locked failed: %v / %v", res1.Fault(), res2.Fault())
	}

	verifyAppendedLines(t, path, 3, "gamma")
}

func TestAppendEmptyPathFailures(t *testing.T) {
	resBytes := AppendBytes("", []byte("x"), filepermtype.Standard)
	resStr := AppendString("", "x", filepermtype.Standard)
	resLines := AppendLines("", []string{"x"}, filepermtype.Standard)

	if !resBytes.IsFailed() || !resStr.IsFailed() || !resLines.IsFailed() {
		t.Fatalf("expected empty path failure across all append variants")
	}
}
