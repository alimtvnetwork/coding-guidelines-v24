package fileutil

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"coding-guidelines/common/pkg/errtype"
)

func verifyWriteAndRead(t *testing.T, filePath string, f *os.File) {
	defer f.Close()

	if _, err := f.WriteString("hello log\n"); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	readRes := ReadString(filePath)
	if readRes.IsFailed() {
		t.Fatalf("read failed: %v", readRes.Fault())
	}

	if readRes.Data() != "hello log\n" {
		t.Fatalf("unexpected content: %s", readRes.Data())
	}
}

func TestOpenFile_CreateAppend(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "nested", "test.log")
	res := OpenFile(filePath, FileOpenCreateAppend, FilePermStandard)
	if res.IsFailed() {
		t.Fatalf("expected success, got error: %v", res.Fault())
	}

	verifyWriteAndRead(t, filePath, res.Data())
}

func TestOpenFile_NotFound(t *testing.T) {
	nonExistent := filepath.Join(t.TempDir(), "missing.txt")
	res := OpenFile(nonExistent, FileOpenReadOnly, FilePermStandard)
	if res.IsSuccess() {
		t.Fatalf("expected failure for non-existent file")
	}

	if res.Fault() == nil {
		t.Fatalf("expected fault object")
	}

	if res.Fault().Context().GetString("Path") != nonExistent {
		t.Fatalf("expected path context to match: %s", res.Fault().Context().GetString("Path"))
	}
}

func checkOpenEnumNames(t *testing.T) {
	if FileOpenCreateAppend.Flags() == 0 {
		t.Fatalf("expected non-zero flags for FileOpenCreateAppend")
	}

	if FileOpenCreateAppend.Name() != "CreateAppend" {
		t.Fatalf("unexpected name: %s", FileOpenCreateAppend.Name())
	}

	if FileOpenReadOrCreateOnly.Name() != "ReadOrCreateOnly" {
		t.Fatalf("unexpected name: %s", FileOpenReadOrCreateOnly.Name())
	}
}

func checkPermEnumModes(t *testing.T) {
	if FileOpenWriteOrCreateOnly.Name() != "WriteOrCreateOnly" {
		t.Fatalf("unexpected name: %s", FileOpenWriteOrCreateOnly.Name())
	}

	if FileOpenReadWriteOrCreateOnly.Name() != "ReadWriteOrCreateOnly" {
		t.Fatalf("unexpected name: %s", FileOpenReadWriteOrCreateOnly.Name())
	}

	if FilePermStandard.Mode() != os.FileMode(0644) {
		t.Fatalf("unexpected mode: %v", FilePermStandard.Mode())
	}
}

func TestEnums_NamesAndFlags(t *testing.T) {
	checkOpenEnumNames(t)
	checkPermEnumModes(t)
}

func checkPermStringsStandard(t *testing.T) {
	if FilePermStandard.OctalString() != "0644" {
		t.Fatalf("expected 0644, got: %s", FilePermStandard.OctalString())
	}

	if FilePermStandard.PosixString() != "rw-r--r--" {
		t.Fatalf("expected rw-r--r--, got: %s", FilePermStandard.PosixString())
	}
}

func checkPermStringsSpecial(t *testing.T) {
	if FilePermExecutable.PosixString() != "rwxr-xr-x" {
		t.Fatalf("expected rwxr-xr-x, got: %s", FilePermExecutable.PosixString())
	}

	if FilePermPrivate.PosixString() != "rw-------" {
		t.Fatalf("expected rw-------, got: %s", FilePermPrivate.PosixString())
	}
}

func TestFilePermType_StringsAndOctal(t *testing.T) {
	checkPermStringsStandard(t)
	checkPermStringsSpecial(t)
}

func TestFilePermType_Inspections(t *testing.T) {
	if !FilePermPrivate.IsPrivate() {
		t.Fatalf("expected FilePermPrivate to be private")
	}

	if !FilePermStandard.IsPublic() {
		t.Fatalf("expected FilePermStandard to be public")
	}

	if !FilePermExecutable.IsExecutable() {
		t.Fatalf("expected FilePermExecutable to be executable")
	}
}

func TestFilePermType_Modifiers(t *testing.T) {
	standard := FilePermStandard
	if standard.WithPrivate().OctalString() != "0600" {
		t.Fatalf("expected 0600 after WithPrivate")
	}

	if standard.WithReadOnly().OctalString() != "0444" {
		t.Fatalf("expected 0444 after WithReadOnly")
	}

	if standard.WithExecutable().OctalString() != "0755" {
		t.Fatalf("expected 0755 after WithExecutable")
	}
}

func checkParsePermValid(t *testing.T) {
	wrap := ParsePerm("0644")
	if wrap.IsFailed() || wrap.Data() != FilePermStandard {
		t.Fatalf("unexpected result for 0644: %v", wrap)
	}

	fromMode := FromFileMode(os.FileMode(0755))
	if fromMode != FilePermExecutable {
		t.Fatalf("expected FilePermExecutable, got: %v", fromMode)
	}
}

func TestFilePermType_ParseAndConvert(t *testing.T) {
	checkParsePermValid(t)

	invalidWrap := ParsePerm("invalid")
	if invalidWrap.IsSuccess() {
		t.Fatalf("expected parse failure on invalid string")
	}
}

func checkFileOpTypeNames(t *testing.T) {
	if FileOpReadOnly.Name() != "ReadOnly" {
		t.Fatalf("unexpected name: %s", FileOpReadOnly.Name())
	}

	if FileOpDelete.Name() != "Delete" {
		t.Fatalf("unexpected name: %s", FileOpDelete.Name())
	}
}

func checkFileOpTypePredicates(t *testing.T) {
	if !FileOpDelete.IsDelete() || !FileOpReadOnly.IsReadOnly() {
		t.Fatalf("unexpected predicate result")
	}

	if !FileOpAppend.IsAppend() || FileOpCreate.OpenMode() != FileOpenCreateNew {
		t.Fatalf("unexpected append or openmode")
	}
}

func TestFileOpType_Enums(t *testing.T) {
	checkFileOpTypeNames(t)
	checkFileOpTypePredicates(t)
}

func checkWriteAndRead(t *testing.T, path string) {
	writeRes := WriteFile(path, []byte("hello world"), FilePermStandard)
	if writeRes.IsFailed() {
		t.Fatalf("write failed: %v", writeRes.Fault())
	}

	readRes := ReadFile(path)
	if readRes.IsFailed() || string(readRes.Data()) != "hello world" {
		t.Fatalf("unexpected read: %v", readRes)
	}
}

func checkStatAndSize(t *testing.T, path string) {
	statRes := Stat(path)
	if statRes.IsFailed() {
		t.Fatalf("stat failed: %v", statRes.Fault())
	}

	sizeRes := FileSize(path)
	if sizeRes.IsFailed() || sizeRes.Data() != int64(11) {
		t.Fatalf("unexpected size: %v", sizeRes)
	}
}

func checkDeleteAndVerify(t *testing.T, path string) {
	delRes := DeleteFile(path)
	if delRes.IsFailed() {
		t.Fatalf("delete failed: %v", delRes.Fault())
	}

	if Stat(path).IsSuccess() {
		t.Fatalf("expected failure after delete")
	}
}

func TestFileutil_DeleteAndStat(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "sample.txt")
	checkWriteAndRead(t, filePath)
	checkStatAndSize(t, filePath)
	checkDeleteAndVerify(t, filePath)
}

func checkExecuteOpWriteRead(t *testing.T, path string) {
	if res := ExecuteOp(path, FileOpCreateAppend, FilePermStandard, []byte("step1\n")); res.IsFailed() {
		t.Fatalf("create failed: %v", res.Fault())
	}

	if res := ExecuteOp(path, FileOpAppend, FilePermStandard, []byte("step2\n")); res.IsFailed() {
		t.Fatalf("append failed: %v", res.Fault())
	}

	readRes := ExecuteOp(path, FileOpReadOnly, FilePermStandard, nil)
	if readRes.IsFailed() || string(readRes.Data()) != "step1\nstep2\n" {
		t.Fatalf("unexpected read: %v", readRes)
	}
}

func TestFileutil_ExecuteOp(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "op_test.txt")
	checkExecuteOpWriteRead(t, filePath)

	delRes := ExecuteOp(filePath, FileOpDelete, FilePermStandard, nil)
	if delRes.IsFailed() {
		t.Fatalf("delete op failed: %v", delRes.Fault())
	}
}

func checkFileAndBoolResults(t *testing.T, path string) {
	fRes := FileFailure(errtype.IO, io.ErrUnexpectedEOF, path, "read err")
	if fRes.IsSuccess() || fRes.Fault().Context().GetString("Path") != path {
		t.Fatalf("FileFailure context mismatch")
	}

	bRes := BoolFailureMsg(errtype.Validation, path, "invalid")
	if bRes.IsSuccess() || bRes.Fault().Context().GetString("Path") != path {
		t.Fatalf("BoolFailureMsg context mismatch")
	}
}

func checkBytesAndStringResults(t *testing.T, path string) {
	byRes := BytesFailure(errtype.NotFound, os.ErrNotExist, path, "not found")
	if byRes.IsSuccess() || byRes.Fault().Context().GetString("Path") != path {
		t.Fatalf("BytesFailure context mismatch")
	}

	sRes := StringFailureMsg(errtype.Validation, path, "invalid str")
	if sRes.IsSuccess() || sRes.Fault().Context().GetString("Path") != path {
		t.Fatalf("StringFailureMsg context mismatch")
	}
}

func checkLinesAndInt64Results(t *testing.T, path string) {
	lRes := LinesFailure(errtype.IO, io.ErrClosedPipe, path, "closed")
	if lRes.IsSuccess() || lRes.Fault().Context().GetString("Path") != path {
		t.Fatalf("LinesFailure context mismatch")
	}

	iRes := Int64FailureMsg(errtype.Validation, path, "bad size")
	if iRes.IsSuccess() || iRes.Fault().Context().GetString("Path") != path {
		t.Fatalf("Int64FailureMsg context mismatch")
	}
}

func checkFileInfoResults(t *testing.T, path string) {
	fiRes := FileInfoFailure(errtype.Forbidden, os.ErrPermission, path, "no access")
	if fiRes.IsSuccess() || fiRes.Fault().Context().GetString("Path") != path {
		t.Fatalf("FileInfoFailure context mismatch")
	}

	fiResMsg := FileInfoFailureMsg(errtype.Validation, path, "invalid fi")
	if fiResMsg.IsSuccess() || fiResMsg.Fault().Context().GetString("Path") != path {
		t.Fatalf("FileInfoFailureMsg context mismatch")
	}
}

func checkSuccessConstructors(t *testing.T) {
	if !BoolSuccess(true).Data() || string(BytesSuccess([]byte("ok")).Data()) != "ok" {
		t.Fatalf("BoolSuccess or BytesSuccess failed")
	}

	if StringSuccess("test").Data() != "test" || len(LinesSuccess([]string{"a"}).Data()) != 1 {
		t.Fatalf("StringSuccess or LinesSuccess failed")
	}

	if Int64Success(42).Data() != 42 {
		t.Fatalf("Int64Success failed")
	}
}

func TestConcreteResults_ConstructorsAndPathContext(t *testing.T) {
	testPath := "/var/log/app.log"
	checkFileAndBoolResults(t, testPath)
	checkBytesAndStringResults(t, testPath)
	checkLinesAndInt64Results(t, testPath)
	checkFileInfoResults(t, testPath)
	checkSuccessConstructors(t)
}
