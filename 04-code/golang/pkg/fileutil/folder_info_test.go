package fileutil

import (
	"path/filepath"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/enum/filepermtype"
)

func testFolderInfoBasicsAndStat(t *testing.T, dir string) {
	fi := NewFolderInfo(dir)
	if fi.Path() != filepath.Clean(dir) {
		t.Fatalf("expected path %s, got %s", dir, fi.Path())
	}

	if fi.Name() != filepath.Base(dir) {
		t.Fatalf("expected name %s, got %s", filepath.Base(dir), fi.Name())
	}

	if !fi.IsExists() || !fi.Exists() {
		t.Fatalf("expected folder to exist")
	}

	statRes := fi.Stat()
	if !statRes.IsSuccess() || !statRes.Data().IsDir() {
		t.Fatalf("expected stat to succeed and be dir")
	}
}

func testFolderInfoAbsAndSubfolder(t *testing.T, dir string) {
	fi := NewFolderInfo(dir)
	absFi := fi.Abs()
	if !absFi.IsExists() {
		t.Fatalf("expected abs folder to exist")
	}

	childFi := fi.Subfolder("nested")
	if childFi.Name() != "nested" {
		t.Fatalf("expected nested subfolder name")
	}

	childFile := fi.File("child.txt")
	if childFile.Name() != "child.txt" {
		t.Fatalf("expected child file name")
	}
}

func TestFolderInfo_Basics(t *testing.T) {
	dir := t.TempDir()
	testFolderInfoBasicsAndStat(t, dir)
	testFolderInfoAbsAndSubfolder(t, dir)
}

func testFolderInfoParentValid(t *testing.T, parentDir, childDir string) {
	child := NewFolderInfo(childDir)
	parent := child.Parent()
	if parent == nil || parent.Path() != filepath.Clean(parentDir) {
		t.Fatalf("unexpected parent: %v", parent)
	}

	if child.ParentPath() != filepath.Clean(parentDir) {
		t.Fatalf("unexpected parent path: %s", child.ParentPath())
	}

	if child.ParentFolderName() != filepath.Base(parentDir) {
		t.Fatalf("unexpected parent folder name: %s", child.ParentFolderName())
	}
}

func testFolderInfoParentRoot(t *testing.T) {
	rootFi := NewFolderInfo(CurrentDir)
	if rootFi.Parent() != nil {
		t.Fatalf("expected root Parent to be nil")
	}

	if rootFi.ParentPath() != "" || rootFi.ParentFolderName() != "" {
		t.Fatalf("expected empty parent info on root")
	}
}

func TestFolderInfo_Parent(t *testing.T) {
	parentDir := t.TempDir()
	childDir := filepath.Join(parentDir, "sub")
	_ = EnsureDir(childDir, filepermtype.Standard)

	testFolderInfoParentValid(t, parentDir, childDir)
	testFolderInfoParentRoot(t)
}

func createTestStructure(t *testing.T, base string) {
	d1 := filepath.Join(base, "dir1")
	d2 := filepath.Join(base, "dir2")
	_ = EnsureDir(d1, filepermtype.Standard)
	_ = EnsureDir(d2, filepermtype.Standard)

	_ = WriteString(filepath.Join(base, "file1.txt"), "hello", filepermtype.Standard)
	_ = WriteString(filepath.Join(base, "file2.log"), "world", filepermtype.Standard)
}

func testFolderInfoSubfolders(t *testing.T, fi *FolderInfo) {
	subs := fi.Subfolders()
	if !subs.IsSuccess() || subs.Count() != 2 {
		t.Fatalf("expected 2 subfolders, got %d", subs.Count())
	}

	dirs := fi.Directories()
	if !dirs.IsSuccess() || dirs.Count() != 2 {
		t.Fatalf("expected 2 directories alias")
	}

	names := fi.FolderNames()
	if !names.IsSuccess() || names.Count() != 2 {
		t.Fatalf("expected 2 folder names")
	}
}

func testFolderInfoFiles(t *testing.T, fi *FolderInfo) {
	files := fi.Files()
	if !files.IsSuccess() || files.Count() != 2 {
		t.Fatalf("expected 2 files, got %d", files.Count())
	}

	names := fi.FileNames()
	if !names.IsSuccess() || names.Count() != 2 {
		t.Fatalf("expected 2 file names")
	}
}

func TestFolderInfo_SubfoldersAndFiles(t *testing.T) {
	dir := t.TempDir()
	createTestStructure(t, dir)

	fi := NewFolderInfo(dir)
	testFolderInfoSubfolders(t, fi)
	testFolderInfoFiles(t, fi)
}

func testParentFolderContext(t *testing.T, subDir string) {
	subFi := NewFolderInfo(subDir)
	pFiles := subFi.ParentFolderFiles()
	if !pFiles.IsSuccess() || pFiles.Count() != 2 {
		t.Fatalf("expected 2 parent folder files, got %d", pFiles.Count())
	}

	pDirs := subFi.ParentFolderDirectories()
	if !pDirs.IsSuccess() || pDirs.Count() != 2 {
		t.Fatalf("expected 2 parent folder directories, got %d", pDirs.Count())
	}
}

func testParentFolderContextRoot(t *testing.T) {
	rootFi := NewFolderInfo(CurrentDir)
	if !rootFi.ParentFolderFiles().IsFailed() {
		t.Fatalf("expected ParentFolderFiles to fail on root")
	}

	if !rootFi.ParentFolderDirectories().IsFailed() {
		t.Fatalf("expected ParentFolderDirectories to fail on root")
	}
}

func TestFolderInfo_ParentFolderContext(t *testing.T) {
	dir := t.TempDir()
	createTestStructure(t, dir)

	subDir := filepath.Join(dir, "dir1")
	testParentFolderContext(t, subDir)
	testParentFolderContextRoot(t)
}

func testWalkFilesAndDirs(t *testing.T, fi *FolderInfo) {
	var filesCount int
	walkErr := fi.WalkFiles(func(filePath string) *appfault.AppError {
		filesCount++

		return nil
	})
	if walkErr != nil || filesCount < 2 {
		t.Fatalf("expected WalkFiles to find files, got %d", filesCount)
	}

	var dirsCount int
	walkDirErr := fi.WalkDirectories(func(dirPath string) *appfault.AppError {
		dirsCount++

		return nil
	})
	if walkDirErr != nil || dirsCount < 2 {
		t.Fatalf("expected WalkDirectories to find dirs, got %d", dirsCount)
	}
}

func testAllFilesAndDirs(t *testing.T, fi *FolderInfo) {
	allFiles := fi.AllFiles()
	if !allFiles.IsSuccess() || allFiles.Count() < 2 {
		t.Fatalf("expected AllFiles >= 2, got %d", allFiles.Count())
	}

	allDirs := fi.AllDirectories()
	if !allDirs.IsSuccess() || allDirs.Count() < 2 {
		t.Fatalf("expected AllDirectories >= 2, got %d", allDirs.Count())
	}
}

func TestFolderInfo_Walking(t *testing.T) {
	dir := t.TempDir()
	createTestStructure(t, dir)

	fi := NewFolderInfo(dir)
	testWalkFilesAndDirs(t, fi)
	testAllFilesAndDirs(t, fi)
}

func TestFolderInfo_Actions(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "action_folder")

	fi := NewFolderInfo(target)
	ensureRes := fi.EnsureDir(filepermtype.Standard)
	if !ensureRes.IsSuccess() || !fi.IsExists() {
		t.Fatalf("expected EnsureDir to create folder")
	}

	delRes := fi.Delete()
	if !delRes.IsSuccess() || fi.IsExists() {
		t.Fatalf("expected Delete to delete empty folder")
	}

	_ = fi.EnsureDir(filepermtype.Standard)
	_ = WriteString(filepath.Join(target, "temp.txt"), "data", filepermtype.Standard)
	remAllRes := fi.RemoveAll()
	if !remAllRes.IsSuccess() || fi.IsExists() {
		t.Fatalf("expected RemoveAll to delete folder and contents")
	}
}

func TestFolderInfo_FindAndFilter(t *testing.T) {
	dir := t.TempDir()
	createTestStructure(t, dir)
	fi := NewFolderInfo(dir)

	files := fi.FindFiles("*.txt")
	if !files.IsSuccess() || files.Count() != 1 {
		t.Fatalf("expected 1 txt file, got %d", files.Count())
	}

	dirs := fi.FindFolders("dir*")
	if !dirs.IsSuccess() || dirs.Count() != 2 {
		t.Fatalf("expected 2 folders, got %d", dirs.Count())
	}

	paths := fi.Find("*.log")
	if !paths.IsSuccess() || paths.Count() != 1 {
		t.Fatalf("expected 1 log file path, got %d", paths.Count())
	}
}

func TestFolderInfo_NormalizeAndNav(t *testing.T) {
	dir := t.TempDir()
	child := filepath.Join(dir, "nested", "leaf")
	_ = EnsureDir(child, filepermtype.Standard)
	fi := NewFolderInfo(child)

	if fi.Normalize().Name() != "leaf" {
		t.Fatalf("unexpected normalized name")
	}

	if fi.Up().Name() != "nested" {
		t.Fatalf("unexpected Up name: %s", fi.Up().Name())
	}

	if fi.UpN(2).Name() != filepath.Base(dir) {
		t.Fatalf("unexpected UpN(2) name")
	}
}
