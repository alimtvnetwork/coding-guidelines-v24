package fileutil

import (
	"path/filepath"
	"testing"
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

	res1 := AppendBytes(path, []byte("hello "), FilePermStandard)
	res2 := AppendBytes(path, []byte("world"), FilePermStandard)
	if !res1.IsSuccess() || !res2.IsSuccess() {
		t.Fatalf("append bytes failed: %v / %v", res1.Fault(), res2.Fault())
	}

	verifyAppendedText(t, path, "hello world")
}

func TestAppendBytesLocked(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "append_bytes_locked.txt")

	res1 := AppendBytesLocked(path, []byte("part1 "), FilePermStandard)
	res2 := AppendBytesLocked(path, []byte("part2"), FilePermStandard)
	if !res1.IsSuccess() || !res2.IsSuccess() {
		t.Fatalf("append bytes locked failed: %v / %v", res1.Fault(), res2.Fault())
	}

	verifyAppendedText(t, path, "part1 part2")
}

func TestAppendString(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "append_string.txt")

	res1 := AppendString(path, "first ", FilePermStandard)
	res2 := AppendString(path, "second", FilePermStandard)
	if !res1.IsSuccess() || !res2.IsSuccess() {
		t.Fatalf("append string failed: %v / %v", res1.Fault(), res2.Fault())
	}

	verifyAppendedText(t, path, "first second")
}

func TestAppendStringLocked(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "append_string_locked.txt")

	res1 := AppendStringLocked(path, "lock_a ", FilePermStandard)
	res2 := AppendStringLocked(path, "lock_b", FilePermStandard)
	if !res1.IsSuccess() || !res2.IsSuccess() {
		t.Fatalf("append string locked failed: %v / %v", res1.Fault(), res2.Fault())
	}

	verifyAppendedText(t, path, "lock_a lock_b")
}

func TestAppendLines(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "append_lines.txt")

	res1 := AppendLines(path, []string{"row1", "row2"}, FilePermStandard)
	res2 := AppendLines(path, []string{"row3"}, FilePermStandard)
	if !res1.IsSuccess() || !res2.IsSuccess() {
		t.Fatalf("append lines failed: %v / %v", res1.Fault(), res2.Fault())
	}

	verifyAppendedLines(t, path, 3, "row3")
}

func TestAppendLinesLocked(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "append_lines_locked.txt")

	res1 := AppendLinesLocked(path, []string{"alpha", "beta"}, FilePermStandard)
	res2 := AppendLinesLocked(path, []string{"gamma"}, FilePermStandard)
	if !res1.IsSuccess() || !res2.IsSuccess() {
		t.Fatalf("append lines locked failed: %v / %v", res1.Fault(), res2.Fault())
	}

	verifyAppendedLines(t, path, 3, "gamma")
}

func TestAppendEmptyPathFailures(t *testing.T) {
	resBytes := AppendBytes("", []byte("x"), FilePermStandard)
	resStr := AppendString("", "x", FilePermStandard)
	resLines := AppendLines("", []string{"x"}, FilePermStandard)

	if !resBytes.IsFailed() || !resStr.IsFailed() || !resLines.IsFailed() {
		t.Fatalf("expected empty path failure across all append variants")
	}
}
