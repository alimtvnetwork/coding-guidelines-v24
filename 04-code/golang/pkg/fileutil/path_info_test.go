package fileutil

import (
	"os"
	"strings"
	"testing"
)

func TestExtAndExtNoDot(t *testing.T) {
	p := "config.json"
	if got := Ext(p); got != ".json" {
		t.Errorf("expected .json, got %s", got)
	}

	if got := ExtNoDot(p); got != "json" {
		t.Errorf("expected json, got %s", got)
	}
}

func TestHasExt(t *testing.T) {
	p := "data.archive.TAR.GZ"
	if !HasExt(p, "gz") {
		t.Error("expected true for HasExt gz")
	}

	if !HasExt(p, ".GZ") {
		t.Error("expected true for HasExt .GZ")
	}

	if HasExt(p, "tar") {
		t.Error("expected false for intermediate ext tar")
	}
}

func TestBaseAndDir(t *testing.T) {
	p := "foo/bar/baz.txt"
	if got := Base(p); got != "baz.txt" {
		t.Errorf("expected baz.txt, got %s", got)
	}

	if got := ToSlash(Dir(p)); got != "foo/bar" {
		t.Errorf("expected foo/bar, got %s", got)
	}
}

func TestStem(t *testing.T) {
	p := "archive.tar.gz"
	if got := Stem(p); got != "archive.tar" {
		t.Errorf("expected archive.tar, got %s", got)
	}

	p2 := "simple.txt"
	if got := Stem(p2); got != "simple" {
		t.Errorf("expected simple, got %s", got)
	}
}

func TestStemFull(t *testing.T) {
	p := "archive.tar.gz"
	if got := StemFull(p); got != "archive" {
		t.Errorf("expected archive, got %s", got)
	}

	hidden := ".gitignore"
	if got := StemFull(hidden); got != ".gitignore" {
		t.Errorf("expected .gitignore, got %s", got)
	}
}

func TestSlug(t *testing.T) {
	p := "My Document (Draft #2)!.md"
	expected := "my-document-draft-2"
	if got := Slug(p); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestSplit(t *testing.T) {
	dir, file := Split("foo/bar/test.go")
	if file != "test.go" {
		t.Errorf("expected test.go, got %s", file)
	}

	if ToSlash(dir) != "foo/bar/" {
		t.Errorf("expected foo/bar/, got %s", dir)
	}
}

func TestParent(t *testing.T) {
	p := "a/b/c"
	if got := ToSlash(Parent(p)); got != "a/b" {
		t.Errorf("expected a/b, got %s", got)
	}
}

func TestParentN(t *testing.T) {
	p := "a/b/c/d"
	if got := ToSlash(ParentN(p, 2)); got != "a/b" {
		t.Errorf("expected a/b, got %s", got)
	}

	if got := ToSlash(ParentN(p, 0)); got != "a/b/c/d" {
		t.Errorf("expected a/b/c/d, got %s", got)
	}
}

func TestIsAbsAndIsRel(t *testing.T) {
	abs1 := "/usr/local/bin"
	if !IsAbs(abs1) {
		t.Error("expected /usr/local/bin to be absolute")
	}

	if IsRel(abs1) {
		t.Error("expected /usr/local/bin not to be relative")
	}

	rel1 := "foo/bar"
	if !IsRel(rel1) {
		t.Error("expected foo/bar to be relative")
	}
}

func TestIsAbsWindowsDrive(t *testing.T) {
	absWin := `C:\Windows\System32`
	if !IsAbs(absWin) {
		t.Error("expected C:\\Windows\\System32 to be absolute")
	}

	if IsRel(absWin) {
		t.Error("expected C:\\Windows\\System32 not to be relative")
	}
}

func TestPathInfo_ObjectBasicsAndTransforms(t *testing.T) {
	pi := NewPathInfo("nested/dir/my-report.final.pdf")
	if pi.Name() != "my-report.final.pdf" || pi.Base() != "my-report.final.pdf" {
		t.Fatalf("unexpected name: %s", pi.Name())
	}

	if pi.Stem() != "my-report.final" || pi.StemFull() != "my-report" {
		t.Fatalf("unexpected stem: %s, stemFull: %s", pi.Stem(), pi.StemFull())
	}

	if pi.Ext() != ".pdf" || pi.ExtNoDot() != "pdf" || !pi.HasExt("pdf") {
		t.Fatalf("unexpected ext: %s", pi.Ext())
	}

	if ToSlash(pi.Clean().Path()) != "nested/dir/my-report.final.pdf" {
		t.Fatalf("unexpected clean path: %s", pi.Clean().Path())
	}
}

func TestPathInfo_Navigation(t *testing.T) {
	pi := NewPathInfo("a/b/c/d")
	up := pi.Up()
	if ToSlash(up.Path()) != "a/b/c" {
		t.Fatalf("unexpected Up: %s", up.Path())
	}

	up2 := pi.UpN(2)
	if ToSlash(up2.Path()) != "a/b" {
		t.Fatalf("unexpected UpN(2): %s", up2.Path())
	}

	sub := up.Cd("sub").Join("file.txt")
	if ToSlash(sub.Path()) != "a/b/c/sub/file.txt" {
		t.Fatalf("unexpected Cd/Join: %s", sub.Path())
	}
}

func TestPathInfo_FindAndFilter(t *testing.T) {
	dir := t.TempDir()
	createTestStructure(t, dir)
	pi := NewPathInfo(dir)

	files := pi.FindFiles("*.txt")
	if !files.IsSuccess() || files.Count() < 1 {
		t.Fatalf("expected FindFiles to find txt, got count %d", files.Count())
	}

	folders := pi.FindFolders("dir*")
	if !folders.IsSuccess() || folders.Count() < 2 {
		t.Fatalf("expected FindFolders to find 2 dirs, got %d", folders.Count())
	}

	filtered := pi.Filter(func(path string, info os.FileInfo) bool {
		return strings.HasSuffix(path, ".log")
	})
	if !filtered.IsSuccess() || filtered.Count() < 1 {
		t.Fatalf("expected Filter to find .log, got %d", filtered.Count())
	}
}
