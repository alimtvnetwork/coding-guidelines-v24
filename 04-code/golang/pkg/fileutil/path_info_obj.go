package fileutil

import (
	"os"
	"path/filepath"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type PathInfo struct {
	path string
}

func NewPathInfo(path string) *PathInfo {
	if len(path) == 0 {
		return &PathInfo{path: CurrentDir}
	}

	return &PathInfo{path: filepath.Clean(path)}
}

func (p *PathInfo) Path() string {
	if p == nil {
		return ""
	}

	return p.path
}

func (p *PathInfo) String() string {
	return p.Path()
}

func (p *PathInfo) Name() string {
	if p == nil || len(p.path) == 0 {
		return ""
	}

	return filepath.Base(p.path)
}

func (p *PathInfo) Normalize() *PathInfo {
	if p == nil {
		return NewPathInfo(CurrentDir)
	}

	res := Normalize(p.path)
	if res.IsSuccess() {
		return NewPathInfo(res.Data())
	}

	return NewPathInfo(p.path)
}

func (p *PathInfo) ToSlash() string {
	if p == nil {
		return ""
	}

	return ToSlash(p.path)
}

func (p *PathInfo) ToNative() string {
	if p == nil {
		return ""
	}

	return ToNative(p.path)
}

func (p *PathInfo) Clean() *PathInfo {
	if p == nil {
		return NewPathInfo(CurrentDir)
	}

	return NewPathInfo(Clean(p.path))
}

func (p *PathInfo) Stem() string {
	if p == nil {
		return ""
	}

	return Stem(p.path)
}

func (p *PathInfo) StemFull() string {
	if p == nil {
		return ""
	}

	return StemFull(p.path)
}

func (p *PathInfo) Ext() string {
	if p == nil {
		return ""
	}

	return Ext(p.path)
}

func (p *PathInfo) Extension() string {
	return p.Ext()
}

func (p *PathInfo) ExtNoDot() string {
	if p == nil {
		return ""
	}

	return ExtNoDot(p.path)
}

func (p *PathInfo) HasExt(ext string) bool {
	if p == nil {
		return false
	}

	return HasExt(p.path, ext)
}

func (p *PathInfo) Slug() string {
	if p == nil {
		return ""
	}

	return Slug(p.path)
}

func (p *PathInfo) Base() string {
	if p == nil {
		return ""
	}

	return Base(p.path)
}

func (p *PathInfo) Dir() string {
	if p == nil {
		return ""
	}

	return Dir(p.path)
}

func (p *PathInfo) Split() (string, string) {
	if p == nil {
		return "", ""
	}

	return Split(p.path)
}

func (p *PathInfo) Parent() *FolderInfo {
	if p == nil || len(p.path) == 0 {
		return nil
	}

	return p.AsFolder().Parent()
}

func (p *PathInfo) ParentPath() string {
	parent := p.Parent()
	if parent == nil {
		return ""
	}

	return parent.Path()
}

func (p *PathInfo) ParentFolderName() string {
	parent := p.Parent()
	if parent == nil {
		return ""
	}

	return parent.Name()
}

func (p *PathInfo) Up() *PathInfo {
	return p.UpN(1)
}

func (p *PathInfo) UpN(levels int) *PathInfo {
	if p == nil {
		return NewPathInfo(CurrentDir)
	}

	return NewPathInfo(ParentN(p.path, levels))
}

func (p *PathInfo) Cd(relPath string) *PathInfo {
	if p == nil {
		return NewPathInfo(relPath)
	}

	return NewPathInfo(filepath.Join(p.path, relPath))
}

func (p *PathInfo) Sub(relPath string) *PathInfo {
	return p.Cd(relPath)
}

func (p *PathInfo) Join(elems ...string) *PathInfo {
	if p == nil {
		return NewPathInfo(filepath.Join(elems...))
	}

	all := append([]string{p.path}, elems...)

	return NewPathInfo(filepath.Join(all...))
}

func (p *PathInfo) IsExists() bool {
	if p == nil || len(p.path) == 0 {
		return false
	}

	_, err := os.Stat(p.path)

	return err == nil
}

func (p *PathInfo) Exists() bool {
	return p.IsExists()
}

func (p *PathInfo) IsDir() bool {
	if p == nil || len(p.path) == 0 {
		return false
	}

	info, err := os.Stat(p.path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func (p *PathInfo) IsDirectory() bool {
	return p.IsDir()
}

func (p *PathInfo) IsFile() bool {
	if p == nil || len(p.path) == 0 {
		return false
	}

	info, err := os.Stat(p.path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func (p *PathInfo) IsAbs() bool {
	if p == nil {
		return false
	}

	return IsAbs(p.path)
}

func (p *PathInfo) IsRel() bool {
	if p == nil {
		return true
	}

	return IsRel(p.path)
}

func (p *PathInfo) AsFolder() *FolderInfo {
	if p == nil {
		return NewFolderInfo(CurrentDir)
	}

	return NewFolderInfo(p.path)
}

func (p *PathInfo) AsFile() *FileInfo {
	if p == nil {
		return NewFileInfo(CurrentDir)
	}

	return NewFileInfo(p.path)
}

func (p *PathInfo) Walk(walkFn func(path string, isDir bool) *appfault.AppError) *appfault.AppError {
	if p == nil || len(p.path) == 0 {
		return appfault.NewFile(errtype.Validation, "", "nil or empty path")
	}

	if walkFn == nil {
		return appfault.NewFile(errtype.Validation, p.path, "walkFn cannot be nil")
	}

	return executeWalk(p.path, walkFn)
}

func (p *PathInfo) WalkFiles(walkFn func(filePath string) *appfault.AppError) *appfault.AppError {
	return p.AsFolder().WalkFiles(walkFn)
}

func (p *PathInfo) WalkFolders(walkFn func(dirPath string) *appfault.AppError) *appfault.AppError {
	return p.AsFolder().WalkDirectories(walkFn)
}

func (p *PathInfo) AllFiles() ResultSlice[string] {
	return p.AsFolder().AllFiles()
}

func (p *PathInfo) AllDirectories() ResultSlice[string] {
	return p.AsFolder().AllDirectories()
}

func (p *PathInfo) Subfolders() ResultSlice[*FolderInfo] {
	return p.AsFolder().Subfolders()
}

func (p *PathInfo) Files() ResultSlice[*FileInfo] {
	return p.AsFolder().Files()
}

func (p *PathInfo) FolderNames() ResultSlice[string] {
	return p.AsFolder().FolderNames()
}

func (p *PathInfo) FileNames() ResultSlice[string] {
	return p.AsFolder().FileNames()
}

func (p *PathInfo) Find(pattern string) ResultSlice[string] {
	if p == nil {
		return result.FailSlice[string](appfault.NewFile(errtype.Validation, "", "nil path info"))
	}

	return findMatchingPaths(p.path, pattern)
}

func (p *PathInfo) FindFiles(pattern string) ResultSlice[*FileInfo] {
	if p == nil {
		return result.FailSlice[*FileInfo](appfault.NewFile(errtype.Validation, "", "nil path info"))
	}

	return findMatchingFiles(p.path, pattern)
}

func (p *PathInfo) FindFolders(pattern string) ResultSlice[*FolderInfo] {
	if p == nil {
		return result.FailSlice[*FolderInfo](appfault.NewFile(errtype.Validation, "", "nil path info"))
	}

	return findMatchingFolders(p.path, pattern)
}

func (p *PathInfo) Filter(filterFn FileFilterFunc) ResultSlice[string] {
	if p == nil {
		return result.FailSlice[string](appfault.NewFile(errtype.Validation, "", "nil path info"))
	}

	return filterMatchingPaths(p.path, filterFn)
}
