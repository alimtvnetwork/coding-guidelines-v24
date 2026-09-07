package fileutil

import (
	"path/filepath"
	"testing"

	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/enum/openfiletype"
)

func testOpenReadOnlyAndTruncate(t *testing.T, path string) {
	truncRes := File.Open.Truncate(path, filepermtype.Standard)
	if !truncRes.IsSuccess() {
		t.Fatalf("Open.Truncate failed: %v", truncRes.Fault())
	}

	defer truncRes.Data().Close()

	roRes := File.Open.ReadOnly(path)
	if !roRes.IsSuccess() {
		t.Fatalf("Open.ReadOnly failed: %v", roRes.Fault())
	}

	defer roRes.Data().Close()
}

func testOpenAppendAndRW(t *testing.T, path string) {
	appRes := File.Open.CreateAppend(path, filepermtype.Standard)
	if !appRes.IsSuccess() {
		t.Fatalf("Open.CreateAppend failed: %v", appRes.Fault())
	}

	defer appRes.Data().Close()

	rwRes := File.Open.ReadWrite(path, filepermtype.Standard)
	if !rwRes.IsSuccess() {
		t.Fatalf("Open.ReadWrite failed: %v", rwRes.Fault())
	}

	defer rwRes.Data().Close()
}

func TestFileOpenOperations(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "open_test.txt")

	testOpenReadOnlyAndTruncate(t, path)
	testOpenAppendAndRW(t, path)

	fileRes := File.Open.File(path, openfiletype.ReadOnly, filepermtype.Standard)
	if !fileRes.IsSuccess() {
		t.Fatalf("Open.File failed: %v", fileRes.Fault())
	}

	defer fileRes.Data().Close()
}

func testCreateFileAndDir(t *testing.T, dir string) {
	fRes := File.Create.File(filepath.Join(dir, "new.txt"), filepermtype.Standard)
	if !fRes.IsSuccess() {
		t.Fatalf("Create.File failed: %v", fRes.Fault())
	}

	defer fRes.Data().Close()

	dRes := File.Create.Dir(filepath.Join(dir, "sub"), filepermtype.Standard)
	if !dRes.IsSuccess() {
		t.Fatalf("Create.Dir failed: %v", dRes.Fault())
	}
}

func testCreateEnsureAndTemp(t *testing.T, dir string) {
	ensureRes := File.Create.EnsureDir(filepath.Join(dir, "ensured"), filepermtype.Standard)
	if !ensureRes.IsSuccess() {
		t.Fatalf("Create.EnsureDir failed: %v", ensureRes.Fault())
	}

	tmpFile := File.Create.Temp("pattern*")
	if !tmpFile.IsSuccess() {
		t.Fatalf("Create.Temp failed: %v", tmpFile.Fault())
	}

	defer tmpFile.Data().Close()
}

func TestFileCreateOperations(t *testing.T) {
	dir := t.TempDir()
	testCreateFileAndDir(t, dir)
	testCreateEnsureAndTemp(t, dir)

	tmpDir := File.Create.TempDir("pattern_dir*")
	if !tmpDir.IsSuccess() {
		t.Fatalf("Create.TempDir failed: %v", tmpDir.Fault())
	}
}

func testWriteBytesAndString(t *testing.T, dir string) {
	bPath := filepath.Join(dir, "bytes.txt")
	bRes := File.Write.Bytes(bPath, []byte("data"), filepermtype.Standard)
	if !bRes.IsSuccess() {
		t.Fatalf("Write.Bytes failed: %v", bRes.Fault())
	}

	sPath := filepath.Join(dir, "str.txt")
	sRes := File.Write.String(sPath, "string_content", filepermtype.Standard)
	if !sRes.IsSuccess() {
		t.Fatalf("Write.String failed: %v", sRes.Fault())
	}
}

func testWriteLinesAndAtomic(t *testing.T, dir string) {
	lPath := filepath.Join(dir, "lines.txt")
	lRes := File.Write.Lines(lPath, []string{"l1", "l2"}, filepermtype.Standard)
	if !lRes.IsSuccess() {
		t.Fatalf("Write.Lines failed: %v", lRes.Fault())
	}

	aPath := filepath.Join(dir, "atomic.txt")
	aRes := File.Write.Atomic(aPath, []byte("atomic_val"), filepermtype.Standard)
	if !aRes.IsSuccess() {
		t.Fatalf("Write.Atomic failed: %v", aRes.Fault())
	}
}

func testWriteJsonAndYaml(t *testing.T, dir string) {
	payload := map[string]string{"key": "val"}
	jPath := filepath.Join(dir, "data.json")
	if !File.Write.Json(jPath, payload, filepermtype.Standard).IsSuccess() {
		t.Fatalf("Write.Json failed")
	}

	yPath := filepath.Join(dir, "data.yaml")
	if !File.Write.Yaml(yPath, payload, filepermtype.Standard).IsSuccess() {
		t.Fatalf("Write.Yaml failed")
	}
}

func TestFileWriteOperations(t *testing.T) {
	dir := t.TempDir()
	testWriteBytesAndString(t, dir)
	testWriteLinesAndAtomic(t, dir)
	testWriteJsonAndYaml(t, dir)

	anyPath := filepath.Join(dir, "any.txt")
	if !File.Write.Any(anyPath, "any_val", filepermtype.Standard).IsSuccess() {
		t.Fatalf("Write.Any failed")
	}
}

func TestFileAppendOperations(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "append_ns.txt")

	res1 := File.Append.String(path, "step1 ", filepermtype.Standard)
	res2 := File.Append.Bytes(path, []byte("step2 "), filepermtype.Standard)
	res3 := File.Append.Lines(path, []string{"step3"}, filepermtype.Standard)
	if !res1.IsSuccess() || !res2.IsSuccess() || !res3.IsSuccess() {
		t.Fatalf("File.Append operations failed")
	}

	if !File.Append.BytesLocked(path, []byte(" lock"), filepermtype.Standard).IsSuccess() {
		t.Fatalf("File.Append.BytesLocked failed")
	}
}

func verifyReadNsData(t *testing.T, lRes LinesResult) {
	if len(lRes.Data()) != 2 || lRes.Data()[0] != "first line" {
		t.Fatalf("unexpected read lines data: %v", lRes.Data())
	}
}

func TestFileReadOperations(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "read_ns.txt")
	_ = File.Write.Lines(path, []string{"first line", "second line"}, filepermtype.Standard)

	bRes := File.Read.Bytes(path)
	sRes := File.Read.String(path)
	tRes := File.Read.Text(path)
	lRes := File.Read.Lines(path)
	if !bRes.IsSuccess() || !sRes.IsSuccess() || !tRes.IsSuccess() || !lRes.IsSuccess() {
		t.Fatalf("File.Read operations failed")
	}

	verifyReadNsData(t, lRes)
}

func testNewFileWriterEngines(t *testing.T, dir string) {
	w := New.Writer.Default(filepath.Join(dir, "w.txt"))
	if w == nil {
		t.Fatalf("New.Writer.Default returned nil")
	}

	optsWriter := New.Writer.WithOptions(FileWriterOptions{Path: filepath.Join(dir, "w_opt.txt")})
	if optsWriter == nil {
		t.Fatalf("New.Writer.WithOptions returned nil")
	}
}

func testNewWriters(t *testing.T, dir string) {
	testNewFileWriterEngines(t, dir)

	bw := New.BoundWriter.Default(filepath.Join(dir, "bw.txt"))
	if bw == nil {
		t.Fatalf("New.BoundWriter.Default returned nil")
	}
}

func testNewAppender(t *testing.T, dir string) {
	app := New.Appender.Default(filepath.Join(dir, "app.txt"), filepermtype.Standard)
	if app == nil {
		t.Fatalf("New.Appender.Default returned nil")
	}

	autoApp := New.Appender.AutoSync(filepath.Join(dir, "app_auto.txt"), filepermtype.Standard)
	if autoApp == nil {
		t.Fatalf("New.Appender.AutoSync returned nil")
	}
}

func testNewPath(t *testing.T) {
	p := New.Path.Default("foo/bar")
	if p == nil || p.Raw() != "foo/bar" {
		t.Fatalf("New.Path.Default failed: %v", p)
	}

	fromParts := New.Path.FromParts("foo", "bar")
	if fromParts == nil {
		t.Fatalf("New.Path.FromParts returned nil")
	}
}

func testNewStreamWriter(t *testing.T, dir string) {
	path := filepath.Join(dir, "sw.txt")
	swRes := New.StreamWriter.Append(path, filepermtype.Standard)
	if !swRes.IsSuccess() {
		t.Fatalf("New.StreamWriter.Append failed: %v", swRes.Fault())
	}

	if closer, ok := swRes.Data().Destination().(interface{ Close() error }); ok {
		_ = closer.Close()
	}
}

func testNewShortcuts(t *testing.T, path string) {
	if New.FileWriter(path) == nil || New.FileAppender(path, filepermtype.Standard) == nil {
		t.Fatalf("New writer/appender shortcuts failed")
	}

	if New.BoundFileWriter(path) == nil || New.PathWrapper(path) == nil {
		t.Fatalf("New bound/path shortcuts failed")
	}
}

func TestNewCreatorNamespace(t *testing.T) {
	dir := t.TempDir()
	testNewWriters(t, dir)
	testNewAppender(t, dir)
	testNewPath(t)
	testNewStreamWriter(t, dir)
	testNewShortcuts(t, filepath.Join(dir, "short.txt"))
}

func verifyCompatText(t *testing.T, p1 string, p2 string) {
	r1 := ReadText(p1)
	r2 := File.Read.Text(p2)
	if !r1.IsSuccess() || !r2.IsSuccess() {
		t.Fatalf("read failed during compat check")
	}

	if r1.Data() != r2.Data() {
		t.Fatalf("mismatch between top-level and File.*: %q != %q", r1.Data(), r2.Data())
	}
}

func TestBackwardCompatibility(t *testing.T) {
	dir := t.TempDir()
	p1 := filepath.Join(dir, "compat1.txt")
	p2 := filepath.Join(dir, "compat2.txt")

	_ = WriteBytes(p1, []byte("hello"), filepermtype.Standard)
	_ = File.Write.Bytes(p2, []byte("hello"), filepermtype.Standard)
	verifyCompatText(t, p1, p2)
}

func testFileTarget(t *testing.T, dir string) {
	path := filepath.Join(dir, "target.txt")
	target := File.Target(path)
	if target == nil {
		t.Fatalf("File.Target returned nil")
	}

	if target.AbsPath() != path {
		t.Fatalf("unexpected target AbsPath: %q", target.AbsPath())
	}
}

func testFileAt(t *testing.T, dir string) {
	at := File.At(dir, "sub/at.txt")
	if at == nil {
		t.Fatalf("File.At returned nil")
	}

	if at.RelPath() != filepath.Clean("sub/at.txt") {
		t.Fatalf("unexpected at RelPath: %q", at.RelPath())
	}
}

func testNewTarget(t *testing.T, dir string) {
	path := filepath.Join(dir, "new_target.txt")
	target := New.Target(path)
	if target == nil {
		t.Fatalf("New.Target returned nil")
	}

	if target.AbsPath() != path {
		t.Fatalf("unexpected target AbsPath: %q", target.AbsPath())
	}
}

func testNewAt(t *testing.T, dir string) {
	at := New.At(dir, "sub/new_at.txt")
	if at == nil {
		t.Fatalf("New.At returned nil")
	}

	if at.RelPath() != filepath.Clean("sub/new_at.txt") {
		t.Fatalf("unexpected at RelPath: %q", at.RelPath())
	}
}

func testNewFilePathOps(t *testing.T, dir string) {
	path := filepath.Join(dir, "fp_ops.txt")
	ops := New.FilePathOps(path)
	if ops == nil {
		t.Fatalf("New.FilePathOps returned nil")
	}

	if ops.AbsPath() != path {
		t.Fatalf("unexpected ops AbsPath: %q", ops.AbsPath())
	}
}

func testNewFilePathOpsAt(t *testing.T, dir string) {
	opsAt := New.FilePathOpsAt(dir, "sub.txt")
	if opsAt == nil {
		t.Fatalf("New.FilePathOpsAt returned nil")
	}

	if opsAt.WorkDir() != filepath.Clean(dir) {
		t.Fatalf("unexpected opsAt WorkDir: %q", opsAt.WorkDir())
	}
}

func TestFileAndNewFilePathOps(t *testing.T) {
	dir := t.TempDir()
	testFileTarget(t, dir)
	testFileAt(t, dir)
	testNewTarget(t, dir)
	testNewAt(t, dir)
	testNewFilePathOps(t, dir)
	testNewFilePathOpsAt(t, dir)
}
