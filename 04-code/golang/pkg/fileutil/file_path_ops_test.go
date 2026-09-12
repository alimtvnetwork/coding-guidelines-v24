package fileutil

import (
	"os"
	"path/filepath"
	"testing"

	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/enum/openfiletype"
)

func testFilePathOpsConstructorsRel(t *testing.T) {
	ops := NewFilePathOps("sub/doc.txt")
	if ops.RelPath() != filepath.Clean("sub/doc.txt") {
		t.Fatalf("unexpected relPath: %q", ops.RelPath())
	}

	if ops.WorkDir() != "." {
		t.Fatalf("unexpected workDir: %q", ops.WorkDir())
	}
}

func testFilePathOpsConstructorsAbs(t *testing.T, dir string) {
	absFile := filepath.Join(dir, "root.txt")
	ops := NewFilePathOps(absFile)
	if ops.AbsPath() != absFile {
		t.Fatalf("unexpected absPath: %q", ops.AbsPath())
	}

	if ops.Base() != "root.txt" {
		t.Fatalf("unexpected base: %q", ops.Base())
	}
}

func testFilePathOpsConstructorsEmpty(t *testing.T) {
	empty := NewFilePathOps("")
	if empty.WorkDir() != "." {
		t.Fatalf("unexpected empty workDir: %q", empty.WorkDir())
	}

	if empty.RelPath() != "." {
		t.Fatalf("unexpected empty relPath: %q", empty.RelPath())
	}
}

func TestFilePathOps_Constructors(t *testing.T) {
	dir := t.TempDir()
	testFilePathOpsConstructorsRel(t)
	testFilePathOpsConstructorsAbs(t, dir)
	testFilePathOpsConstructorsEmpty(t)
}

func verifyWithWorkDir(t *testing.T, orig *FilePathOps) {
	modified := orig.WithWorkDir("/custom")
	if modified == orig {
		t.Fatalf("expected different pointer")
	}

	if orig.WorkDir() != "." {
		t.Fatalf("original workDir mutated")
	}

	if modified.WorkDir() != filepath.Clean("/custom") {
		t.Fatalf("unexpected modified workDir: %q", modified.WorkDir())
	}
}

func verifyWithRelPath(t *testing.T, orig *FilePathOps) {
	modified := orig.WithRelPath("new/doc.md")
	if modified == orig {
		t.Fatalf("expected different pointer")
	}

	if orig.RelPath() != "doc.txt" {
		t.Fatalf("original relPath mutated")
	}

	if modified.RelPath() != filepath.Clean("new/doc.md") {
		t.Fatalf("unexpected modified relPath: %q", modified.RelPath())
	}
}

func verifyJoinAndClone(t *testing.T, orig *FilePathOps) {
	joined := orig.Join("part1", "part2.txt")
	if orig.RelPath() != "doc.txt" {
		t.Fatalf("original relPath mutated")
	}

	if joined.Ext() != ".txt" {
		t.Fatalf("unexpected joined ext: %q", joined.Ext())
	}

	clone := orig.Clone()
	if clone == orig || clone.AbsPath() != orig.AbsPath() {
		t.Fatalf("unexpected clone behavior")
	}
}

func TestFilePathOps_Modifiers_Immutability(t *testing.T) {
	orig := NewFilePathOps("doc.txt")
	verifyWithWorkDir(t, orig)
	verifyWithRelPath(t, orig)
	verifyJoinAndClone(t, orig)
}

func testInspectionMethods(t *testing.T, dir string) {
	path := filepath.Join(dir, "archive.tar.gz")
	ops := NewFilePathOps(path)
	if ops.Ext() != ".gz" {
		t.Fatalf("unexpected ext: %q", ops.Ext())
	}

	if ops.Base() != "archive.tar.gz" {
		t.Fatalf("unexpected base: %q", ops.Base())
	}

	if ops.Dir() != filepath.Clean(dir) {
		t.Fatalf("unexpected dir: %q", ops.Dir())
	}
}

func testStringAndAbsMethods(t *testing.T, dir string) {
	path := filepath.Join(dir, "test.txt")
	ops := NewFilePathOps(path)
	if ops.String() != ops.AbsPath() {
		t.Fatalf("String() != AbsPath()")
	}

	if ops.Exists() {
		t.Fatalf("expected non-existing file to return false")
	}
}

func TestFilePathOps_Accessors(t *testing.T) {
	dir := t.TempDir()
	testInspectionMethods(t, dir)
	testStringAndAbsMethods(t, dir)
}

func testEnsureParentDirSuccess(t *testing.T, dir string) {
	nested := filepath.Join(dir, "a", "b", "c", "file.txt")
	ops := NewFilePathOps(nested)
	res := ops.EnsureParentDir()
	if !res.IsSuccess() {
		t.Fatalf("EnsureParentDir failed: %v", res.Fault())
	}

	parentDir := filepath.Dir(nested)
	if _, err := os.Stat(parentDir); err != nil {
		t.Fatalf("parent dir not created: %v", err)
	}
}

func testEnsureFileCreation(t *testing.T, dir string) {
	nestedFile := filepath.Join(dir, "x", "y", "z.txt")
	ops := NewFilePathOps(nestedFile)
	res := ops.EnsureFile(filepermtype.Standard)
	if !res.IsSuccess() {
		t.Fatalf("EnsureFile failed: %v", res.Fault())
	}

	if !ops.Exists() {
		t.Fatalf("EnsureFile did not create file")
	}
}

func testEnsureFileNoTruncate(t *testing.T, dir string) {
	filePath := filepath.Join(dir, "exist.txt")
	ops := NewFilePathOps(filePath)
	_ = ops.WriteString("initial data", filepermtype.Standard)

	res := ops.EnsureFile(filepermtype.Standard)
	if !res.IsSuccess() {
		t.Fatalf("EnsureFile on existing file failed: %v", res.Fault())
	}

	content := ops.ReadString()
	if content.Data() != "initial data" {
		t.Fatalf("EnsureFile truncated existing file: %q", content.Data())
	}
}

func testCreateIfNotExist(t *testing.T, dir string) {
	filePath := filepath.Join(dir, "new_or_exist.txt")
	ops := NewFilePathOps(filePath)
	res := ops.CreateIfNotExist(filepermtype.Standard)
	if !res.IsSuccess() {
		t.Fatalf("CreateIfNotExist failed: %v", res.Fault())
	}

	_ = res.Data().Close()

	if !ops.Exists() {
		t.Fatalf("expected file to exist")
	}
}

func TestFilePathOps_SafetyOps(t *testing.T) {
	dir := t.TempDir()
	testEnsureParentDirSuccess(t, dir)
	testEnsureFileCreation(t, dir)
	testEnsureFileNoTruncate(t, dir)
	testCreateIfNotExist(t, dir)
}

func testWriteAndReadBytes(t *testing.T, ops *FilePathOps) {
	data := []byte("hello bytes")
	wRes := ops.WriteBytes(data, filepermtype.Standard)
	if !wRes.IsSuccess() {
		t.Fatalf("WriteBytes failed: %v", wRes.Fault())
	}

	rRes := ops.ReadBytes()
	if !rRes.IsSuccess() || string(rRes.Data()) != "hello bytes" {
		t.Fatalf("ReadBytes failed: %v", rRes.Fault())
	}
}

func testWriteAndReadString(t *testing.T, ops *FilePathOps) {
	wRes := ops.WriteString("hello string", filepermtype.Standard)
	if !wRes.IsSuccess() {
		t.Fatalf("WriteString failed: %v", wRes.Fault())
	}

	rRes := ops.ReadString()
	if !rRes.IsSuccess() || rRes.Data() != "hello string" {
		t.Fatalf("ReadString failed: %v", rRes.Fault())
	}
}

func testWriteAndReadLines(t *testing.T, ops *FilePathOps) {
	lines := []string{"line1", "line2"}
	wRes := ops.WriteLines(lines, filepermtype.Standard)
	if !wRes.IsSuccess() {
		t.Fatalf("WriteLines failed: %v", wRes.Fault())
	}

	rRes := ops.ReadLines()
	if !rRes.IsSuccess() || len(rRes.Data()) != 2 {
		t.Fatalf("ReadLines failed: %v", rRes.Fault())
	}
}

func testWriteAtomic(t *testing.T, ops *FilePathOps) {
	wRes := ops.WriteAtomic([]byte("atomic payload"), filepermtype.Standard)
	if !wRes.IsSuccess() {
		t.Fatalf("WriteAtomic failed: %v", wRes.Fault())
	}

	rRes := ops.ReadString()
	if !rRes.IsSuccess() || rRes.Data() != "atomic payload" {
		t.Fatalf("Read after WriteAtomic failed: %v", rRes.Fault())
	}
}

func testAppendBytesAndString(t *testing.T, ops *FilePathOps) {
	_ = ops.WriteBytes([]byte("base"), filepermtype.Standard)
	resB := ops.AppendBytes([]byte("_byte"), filepermtype.Standard)
	if !resB.IsSuccess() {
		t.Fatalf("AppendBytes failed")
	}

	resS := ops.AppendString("_str", filepermtype.Standard)
	if !resS.IsSuccess() {
		t.Fatalf("AppendString failed")
	}
}

func testAppendLinesAndVerify(t *testing.T, ops *FilePathOps) {
	resL := ops.AppendLines([]string{"_line"}, filepermtype.Standard)
	if !resL.IsSuccess() {
		t.Fatalf("AppendLines failed")
	}

	readRes := ops.ReadString()
	if readRes.Data() != "base_byte_str_line\n" {
		t.Fatalf("unexpected appended content: %q", readRes.Data())
	}
}

func testAppendOps(t *testing.T, ops *FilePathOps) {
	testAppendBytesAndString(t, ops)
	testAppendLinesAndVerify(t, ops)
}

func testStatAndOpen(t *testing.T, ops *FilePathOps) {
	sRes := ops.Stat()
	if !sRes.IsSuccess() {
		t.Fatalf("Stat failed: %v", sRes.Fault())
	}

	openRes := ops.Open(openfiletype.ReadOnly, filepermtype.Standard)
	if !openRes.IsSuccess() {
		t.Fatalf("Open failed: %v", openRes.Fault())
	}

	_ = openRes.Data().Close()
}

func testDeleteAndExists(t *testing.T, ops *FilePathOps) {
	delRes := ops.Delete()
	if !delRes.IsSuccess() {
		t.Fatalf("Delete failed")
	}

	if ops.Exists() {
		t.Fatalf("file should have been deleted")
	}
}

func testOpenCreateDeleteStat(t *testing.T, ops *FilePathOps) {
	testStatAndOpen(t, ops)
	testDeleteAndExists(t, ops)
}

func TestFilePathOps_BoundFileOps(t *testing.T) {
	dir := t.TempDir()
	ops := NewFilePathOps(filepath.Join(dir, "bound.txt"))
	testWriteAndReadBytes(t, ops)
	testWriteAndReadString(t, ops)
	testWriteAndReadLines(t, ops)
	testWriteAtomic(t, ops)
	testAppendOps(t, ops)
	testOpenCreateDeleteStat(t, ops)
}

func testNilSafetyAccessors(t *testing.T) {
	var nilOps *FilePathOps
	if nilOps.Exists() {
		t.Fatalf("nil.Exists() should be false")
	}

	if nilOps.WorkDir() != "" || nilOps.RelPath() != "" || nilOps.AbsPath() != "" {
		t.Fatalf("nil accessors should return empty")
	}

	if nilOps.String() != "" || nilOps.Dir() != "" || nilOps.Base() != "" || nilOps.Ext() != "" {
		t.Fatalf("nil inspections should return empty")
	}
}

func testNilSafetyModifiers(t *testing.T) {
	var nilOps *FilePathOps
	if nilOps.Clone() != nil {
		t.Fatalf("nil.Clone() should be nil")
	}

	if nilOps.WithWorkDir("dir") == nil || nilOps.WithRelPath("rel") == nil || nilOps.Join("j") == nil {
		t.Fatalf("nil modifiers should safely produce new FilePathOps")
	}
}

func testNilSafetyOperations(t *testing.T) {
	var nilOps *FilePathOps
	pRes := nilOps.EnsureParentDir()
	fRes := nilOps.EnsureFile(filepermtype.Standard)
	if pRes.IsSuccess() || fRes.IsSuccess() {
		t.Fatalf("nil safety operations should fail safely")
	}

	rRes := nilOps.ReadBytes()
	wRes := nilOps.WriteBytes([]byte("x"), filepermtype.Standard)
	if rRes.IsSuccess() || wRes.IsSuccess() {
		t.Fatalf("nil read/write should fail safely")
	}

	sRes := nilOps.Stat()
	dRes := nilOps.Delete()
	if sRes.IsSuccess() || dRes.IsSuccess() {
		t.Fatalf("nil stat/delete should fail safely")
	}
}

func TestFilePathOps_NilSafety(t *testing.T) {
	testNilSafetyAccessors(t)
	testNilSafetyModifiers(t)
	testNilSafetyOperations(t)
}
