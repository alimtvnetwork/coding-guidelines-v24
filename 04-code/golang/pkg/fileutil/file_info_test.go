package fileutil

import (
	"path/filepath"
	"testing"

	"coding-guidelines/common/pkg/enum/filepermtype"
)

func testFileInfoProperties(t *testing.T, filePath string) {
	fi := NewFileInfo(filePath)
	if fi.Path() != filepath.Clean(filePath) {
		t.Fatalf("expected path %s, got %s", filePath, fi.Path())
	}

	if fi.Name() != "test.txt" {
		t.Fatalf("expected test.txt, got %s", fi.Name())
	}

	if fi.Extension() != ".txt" || fi.Ext() != ".txt" || fi.ExtNoDot() != "txt" {
		t.Fatalf("unexpected ext: %s", fi.Ext())
	}

	if fi.Stem() != "test" {
		t.Fatalf("expected stem test, got %s", fi.Stem())
	}

	if fi.Size() != 12 {
		t.Fatalf("expected size 12, got %d", fi.Size())
	}
}

func testFileInfoFolder(t *testing.T, filePath, dir string) {
	fi := NewFileInfo(filePath)
	folder := fi.Folder()
	if folder == nil || folder.Path() != filepath.Clean(dir) {
		t.Fatalf("unexpected folder: %v", folder)
	}

	directory := fi.Directory()
	if directory == nil || directory.Path() != filepath.Clean(dir) {
		t.Fatalf("unexpected directory: %v", directory)
	}

	if fi.ParentFolderName() != filepath.Base(dir) {
		t.Fatalf("unexpected parent folder name: %s", fi.ParentFolderName())
	}
}

func TestFileInfo_Basics(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "test.txt")
	_ = WriteString(filePath, "hello world\n", filepermtype.Standard)

	testFileInfoProperties(t, filePath)
	testFileInfoFolder(t, filePath, dir)
}

func testFileInfoIORead(t *testing.T, fi *FileInfo) {
	bRes := fi.ReadBytes()
	if !bRes.IsSuccess() || len(bRes.Data()) == 0 {
		t.Fatalf("expected ReadBytes to succeed")
	}

	sRes := fi.ReadString()
	if !sRes.IsSuccess() || sRes.Data() != "initial" {
		t.Fatalf("expected ReadString initial, got %s", sRes.Data())
	}

	lRes := fi.ReadLines()
	if !lRes.IsSuccess() || len(lRes.Data()) == 0 {
		t.Fatalf("expected ReadLines to succeed")
	}
}

func testFileInfoIOWriteAndDelete(t *testing.T, fi *FileInfo) {
	wRes := fi.WriteString("updated", filepermtype.Standard)
	if !wRes.IsSuccess() {
		t.Fatalf("expected WriteString to succeed")
	}

	readRes := fi.ReadString()
	if readRes.Data() != "updated" {
		t.Fatalf("expected updated content")
	}

	dRes := fi.Delete()
	if !dRes.IsSuccess() || fi.IsExists() {
		t.Fatalf("expected Delete to succeed")
	}
}

func TestFileInfo_IO(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "io.txt")
	_ = WriteString(filePath, "initial", filepermtype.Standard)

	fi := NewFileInfo(filePath)
	testFileInfoIORead(t, fi)
	testFileInfoIOWriteAndDelete(t, fi)
}

func testPathInfoConversions(t *testing.T, filePath, dir string) {
	piFile := NewPathInfo(filePath)
	if !piFile.IsExists() || !piFile.IsFile() || piFile.IsDir() {
		t.Fatalf("expected piFile to be existing file")
	}

	asFile := piFile.AsFile()
	if asFile == nil || asFile.Name() != "path_file.txt" {
		t.Fatalf("expected AsFile to yield FileInfo")
	}

	piDir := NewPathInfo(dir)
	if !piDir.IsExists() || !piDir.IsDir() || piDir.IsFile() {
		t.Fatalf("expected piDir to be existing dir")
	}

	asFolder := piDir.AsFolder()
	if asFolder == nil || asFolder.Name() != filepath.Base(dir) {
		t.Fatalf("expected AsFolder to yield FolderInfo")
	}
}

func TestPathInfo_Object(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "path_file.txt")
	_ = WriteString(filePath, "data", filepermtype.Standard)

	testPathInfoConversions(t, filePath, dir)
}

func TestFileInfo_NavigationAndNormalize(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "my_report.final.pdf")
	_ = WriteString(filePath, "data", filepermtype.Standard)
	fi := NewFileInfo(filePath)

	if fi.Normalize().Name() != "my_report.final.pdf" {
		t.Fatalf("unexpected normalized name")
	}

	if fi.StemFull() != "my_report" || fi.Slug() != "my-report-final" {
		t.Fatalf("unexpected stemFull (%s) or slug (%s)", fi.StemFull(), fi.Slug())
	}

	if fi.Up().Path() != filepath.Clean(dir) || fi.Parent().Path() != filepath.Clean(dir) {
		t.Fatalf("unexpected parent folder from fi.Up()")
	}
}
