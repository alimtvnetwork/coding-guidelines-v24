package fileutil

import (
	"strings"
	"testing"

	"coding-guidelines/common/pkg/enum/filepermtype"
)

func TestPathNamespace_Temp(t *testing.T) {
	tmpRes := Path.Temp.UserTempDir()
	if tmpRes.IsFailure() {
		t.Fatalf("Path.Temp.UserTempDir failed: %v", tmpRes.Fault())
	}

	pRes := Path.Temp.UserTempPath("sub", "file.tmp")
	if pRes.IsFailure() {
		t.Fatalf("Path.Temp.UserTempPath failed: %v", pRes.Fault())
	}
}

func TestPathNamespace_Env(t *testing.T) {
	res := Path.Env.Expand("~")
	if res.IsFailure() {
		t.Fatalf("Path.Env.Expand failed: %v", res.Fault())
	}

	if res.Data() == "~" {
		t.Error("expected tilde expansion")
	}
}

func TestPathNamespace_Norm(t *testing.T) {
	slashed := Path.Norm.ToSlash(`a\b\c`)
	if strings.Contains(slashed, `\`) {
		t.Errorf("expected no backslash, got %s", slashed)
	}

	cleaned := Path.Norm.Clean("a/b/../c")
	if !strings.HasSuffix(Path.Norm.ToSlash(cleaned), "a/c") {
		t.Errorf("expected a/c, got %s", cleaned)
	}
}

func TestPathNamespace_Info(t *testing.T) {
	if got := Path.Info.Ext("doc.pdf"); got != ".pdf" {
		t.Errorf("expected .pdf, got %s", got)
	}

	if got := Path.Info.Stem("doc.pdf"); got != "doc" {
		t.Errorf("expected doc, got %s", got)
	}

	if got := Path.Info.Slug("Hello World!.md"); got != "hello-world" {
		t.Errorf("expected hello-world, got %s", got)
	}
}

func TestPathNamespace_Join(t *testing.T) {
	joined := Path.Join("a", "b", "c")
	if !strings.HasSuffix(Path.Norm.ToSlash(joined), "a/b/c") {
		t.Errorf("expected a/b/c, got %s", joined)
	}
}

func TestPathWrapper_Chaining(t *testing.T) {
	pw := NewPath("foo/bar/../baz").Clean().ToSlash()
	if pw.Raw() != "foo/bar/../baz" {
		t.Errorf("expected raw string preserved, got %s", pw.Raw())
	}

	if !strings.HasSuffix(pw.String(), "foo/baz") {
		t.Errorf("expected foo/baz, got %s", pw.String())
	}
}

func TestPathWrapper_Inspection(t *testing.T) {
	pw := NewPath("dir/archive.tar.gz")
	if pw.Base() != "archive.tar.gz" {
		t.Errorf("expected archive.tar.gz, got %s", pw.Base())
	}

	if pw.Stem() != "archive.tar" {
		t.Errorf("expected archive.tar, got %s", pw.Stem())
	}

	if pw.Ext() != ".gz" {
		t.Errorf("expected .gz, got %s", pw.Ext())
	}
}

func TestPathWrapper_Parent(t *testing.T) {
	pw := NewPath("a/b/c/d").Parent()
	if !strings.HasSuffix(pw.ToSlash().String(), "a/b/c") {
		t.Errorf("expected a/b/c, got %s", pw.String())
	}

	pw2 := NewPath("a/b/c/d").ParentN(2)
	if !strings.HasSuffix(pw2.ToSlash().String(), "a/b") {
		t.Errorf("expected a/b, got %s", pw2.String())
	}
}

func TestPathWrapper_FileOps_Exists(t *testing.T) {
	tmpDir := t.TempDir()
	p := NewPath(tmpDir).Join("missing.txt")
	existsRes := p.Exists()
	if existsRes.IsFailure() {
		t.Fatalf("Exists failed: %v", existsRes.Fault())
	}

	if existsRes.Data() {
		t.Error("expected missing file not to exist")
	}
}

func TestPathWrapper_FileOps_WriteAndRead(t *testing.T) {
	tmpDir := t.TempDir()
	p := NewPath(tmpDir).Join("testfile.txt")
	wRes := p.WriteString("hello pathwrapper", filepermtype.Standard)
	if wRes.IsFailure() {
		t.Fatalf("WriteString failed: %v", wRes.Fault())
	}

	rRes := p.ReadString()
	if rRes.IsFailure() || rRes.Data() != "hello pathwrapper" {
		t.Errorf("unexpected read result: %v", rRes.Data())
	}
}

func TestFilePathCreator(t *testing.T) {
	p1 := New.Path.Default("foo/bar")
	if p1 == nil || p1.String() != "foo/bar" {
		t.Errorf("unexpected p1: %v", p1)
	}

	p2 := New.Path.FromParts("a", "b", "c")
	if p2 == nil || !strings.HasSuffix(p2.ToSlash().String(), "a/b/c") {
		t.Errorf("unexpected p2: %v", p2)
	}

	p3 := New.PathWrapper("direct")
	if p3 == nil || p3.String() != "direct" {
		t.Errorf("unexpected p3: %v", p3)
	}
}
