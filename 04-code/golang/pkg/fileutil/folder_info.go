package fileutil

import (
	"os"
	"path/filepath"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type FolderInfo struct {
	path string
}

func NewFolderInfo(path string) *FolderInfo {
	if len(path) == 0 {
		return &FolderInfo{path: CurrentDir}
	}

	return &FolderInfo{path: filepath.Clean(path)}
}

func (f *FolderInfo) Path() string {
	if f == nil {
		return ""
	}

	return f.path
}

func (f *FolderInfo) Name() string {
	if f == nil || len(f.path) == 0 {
		return ""
	}

	return filepath.Base(f.path)
}

func (f *FolderInfo) Parent() *FolderInfo {
	if f == nil || len(f.path) == 0 {
		return nil
	}

	dir := filepath.Dir(f.path)
	if dir == f.path || f.path == CurrentDir {
		return nil
	}

	return NewFolderInfo(dir)
}

func (f *FolderInfo) ParentPath() string {
	parent := f.Parent()
	if parent == nil {
		return ""
	}

	return parent.Path()
}

func (f *FolderInfo) ParentFolderName() string {
	parent := f.Parent()
	if parent == nil {
		return ""
	}

	return parent.Name()
}

func (f *FolderInfo) IsExists() bool {
	if f == nil || len(f.path) == 0 {
		return false
	}

	info, err := os.Stat(f.path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func (f *FolderInfo) Exists() bool {
	return f.IsExists()
}

func (f *FolderInfo) Stat() FileInfoResult {
	if f == nil || len(f.path) == 0 {
		return FileInfoFailureMsg(errtype.Validation, "", "nil or empty folder path")
	}

	return Stat(f.path)
}

func (f *FolderInfo) Abs() *FolderInfo {
	if f == nil || len(f.path) == 0 {
		return NewFolderInfo(CurrentDir)
	}

	abs, err := filepath.Abs(f.path)
	if err != nil {
		return f
	}

	return NewFolderInfo(abs)
}

func (f *FolderInfo) Subfolder(name string) *FolderInfo {
	return NewFolderInfo(filepath.Join(f.path, name))
}

func (f *FolderInfo) File(name string) *FileInfo {
	return NewFileInfo(filepath.Join(f.path, name))
}

func readDirEntries(path string) ([]os.DirEntry, *appfault.AppError) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, appfault.WrapFile(errtype.IO, err, path, "failed to read directory")
	}

	return entries, nil
}

func filterSubfolders(base string, entries []os.DirEntry) []*FolderInfo {
	var subs []*FolderInfo
	for _, entry := range entries {
		if entry.IsDir() {
			subs = append(subs, NewFolderInfo(filepath.Join(base, entry.Name())))
		}
	}

	return subs
}

func (f *FolderInfo) Subfolders() ResultSlice[*FolderInfo] {
	if f == nil || len(f.path) == 0 {
		return result.FailSlice[*FolderInfo](appfault.NewFile(errtype.Validation, "", "nil folder"))
	}

	entries, fault := readDirEntries(f.path)
	if fault != nil {
		return result.FailSlice[*FolderInfo](fault)
	}

	return result.OkSlice(filterSubfolders(f.path, entries))
}

func (f *FolderInfo) Directories() ResultSlice[*FolderInfo] {
	return f.Subfolders()
}

func (f *FolderInfo) FolderNames() ResultSlice[string] {
	subs := f.Subfolders()
	if subs.IsFailed() {
		return result.FailSlice[string](subs.Fault())
	}

	names := make([]string, 0, subs.Count())
	for _, sub := range subs.Items {
		names = append(names, sub.Name())
	}

	return result.OkSlice(names)
}

func filterFiles(base string, entries []os.DirEntry) []*FileInfo {
	var files []*FileInfo
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, NewFileInfo(filepath.Join(base, entry.Name())))
		}
	}

	return files
}

func (f *FolderInfo) Files() ResultSlice[*FileInfo] {
	if f == nil || len(f.path) == 0 {
		return result.FailSlice[*FileInfo](appfault.NewFile(errtype.Validation, "", "nil folder"))
	}

	entries, fault := readDirEntries(f.path)
	if fault != nil {
		return result.FailSlice[*FileInfo](fault)
	}

	return result.OkSlice(filterFiles(f.path, entries))
}

func (f *FolderInfo) FileNames() ResultSlice[string] {
	files := f.Files()
	if files.IsFailed() {
		return result.FailSlice[string](files.Fault())
	}

	names := make([]string, 0, files.Count())
	for _, file := range files.Items {
		names = append(names, file.Name())
	}

	return result.OkSlice(names)
}

func (f *FolderInfo) ParentFolderFiles() ResultSlice[*FileInfo] {
	parent := f.Parent()
	if parent == nil {
		return result.FailSlice[*FileInfo](appfault.NewFile(errtype.NotFound, f.path, "no parent folder"))
	}

	return parent.Files()
}

func (f *FolderInfo) ParentFolderDirectories() ResultSlice[*FolderInfo] {
	parent := f.Parent()
	if parent == nil {
		return result.FailSlice[*FolderInfo](appfault.NewFile(errtype.NotFound, f.path, "no parent folder"))
	}

	return parent.Subfolders()
}

func walkEntry(root string, p string, d os.DirEntry, fn func(string, bool) *appfault.AppError) error {
	if p == root {
		return nil
	}

	fault := fn(p, d.IsDir())
	if fault != nil {
		return fault
	}

	return nil
}

func executeWalk(root string, fn func(path string, isDir bool) *appfault.AppError) *appfault.AppError {
	err := filepath.WalkDir(root, func(p string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return appfault.WrapFile(errtype.IO, walkErr, p, "failed accessing path during walk")
		}

		return walkEntry(root, p, d, fn)
	})
	if err == nil {
		return nil
	}

	if appErr, ok := err.(*appfault.AppError); ok {
		return appErr
	}

	return appfault.WrapFile(errtype.IO, err, root, "walk traversal failed")
}

func (f *FolderInfo) Walk(walkFn func(path string, isDir bool) *appfault.AppError) *appfault.AppError {
	if f == nil || len(f.path) == 0 {
		return appfault.NewFile(errtype.Validation, "", "nil or empty folder")
	}

	if walkFn == nil {
		return appfault.NewFile(errtype.Validation, f.path, "walkFn cannot be nil")
	}

	return executeWalk(f.path, walkFn)
}

func (f *FolderInfo) WalkFiles(walkFn func(filePath string) *appfault.AppError) *appfault.AppError {
	if walkFn == nil {
		return appfault.NewFile(errtype.Validation, f.path, "walkFn cannot be nil")
	}

	return f.Walk(func(path string, isDir bool) *appfault.AppError {
		if isDir {
			return nil
		}

		return walkFn(path)
	})
}

func (f *FolderInfo) WalkDirectories(walkFn func(dirPath string) *appfault.AppError) *appfault.AppError {
	if walkFn == nil {
		return appfault.NewFile(errtype.Validation, f.path, "walkFn cannot be nil")
	}

	return f.Walk(func(path string, isDir bool) *appfault.AppError {
		if !isDir {
			return nil
		}

		return walkFn(path)
	})
}

func (f *FolderInfo) AllFiles() ResultSlice[string] {
	var files []string
	fault := f.WalkFiles(func(filePath string) *appfault.AppError {
		files = append(files, filePath)

		return nil
	})
	if fault != nil {
		return result.FailSlice[string](fault)
	}

	return result.OkSlice(files)
}

func (f *FolderInfo) AllDirectories() ResultSlice[string] {
	var dirs []string
	fault := f.WalkDirectories(func(dirPath string) *appfault.AppError {
		dirs = append(dirs, dirPath)

		return nil
	})
	if fault != nil {
		return result.FailSlice[string](fault)
	}

	return result.OkSlice(dirs)
}

func (f *FolderInfo) EnsureDir(perm FilePermType) BoolResult {
	if f == nil || len(f.path) == 0 {
		return BoolFailureMsg(errtype.Validation, "", "nil or empty folder path")
	}

	return EnsureDir(f.path, perm)
}

func (f *FolderInfo) Delete() BoolResult {
	if f == nil || len(f.path) == 0 {
		return BoolFailureMsg(errtype.Validation, "", "nil or empty folder path")
	}

	return DeleteFile(f.path)
}

func (f *FolderInfo) RemoveAll() BoolResult {
	if f == nil || len(f.path) == 0 {
		return BoolFailureMsg(errtype.Validation, "", "nil or empty folder path")
	}

	return RemoveAll(f.path)
}
