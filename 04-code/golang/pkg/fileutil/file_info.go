package fileutil

import (
	"os"
	"path/filepath"
	"strings"

	"coding-guidelines/common/pkg/errtype"
)

type FileInfo struct {
	path string
}

func NewFileInfo(path string) *FileInfo {
	return &FileInfo{path: filepath.Clean(path)}
}

func (f *FileInfo) Path() string {
	if f == nil {
		return ""
	}

	return f.path
}

func (f *FileInfo) Name() string {
	if f == nil || len(f.path) == 0 {
		return ""
	}

	return filepath.Base(f.path)
}

func (f *FileInfo) Extension() string {
	if f == nil || len(f.path) == 0 {
		return ""
	}

	return filepath.Ext(f.path)
}

func (f *FileInfo) Ext() string {
	return f.Extension()
}

func (f *FileInfo) ExtNoDot() string {
	return strings.TrimPrefix(f.Extension(), ExtDot)
}

func (f *FileInfo) Stem() string {
	if f == nil || len(f.path) == 0 {
		return ""
	}

	return Stem(f.path)
}

func (f *FileInfo) Size() int64 {
	if f == nil || len(f.path) == 0 {
		return 0
	}

	info, err := os.Stat(f.path)
	if err != nil {
		return 0
	}

	return info.Size()
}

func (f *FileInfo) Folder() *FolderInfo {
	if f == nil || len(f.path) == 0 {
		return NewFolderInfo(CurrentDir)
	}

	return NewFolderInfo(filepath.Dir(f.path))
}

func (f *FileInfo) Directory() *FolderInfo {
	return f.Folder()
}

func (f *FileInfo) ParentFolderName() string {
	return f.Folder().Name()
}

func (f *FileInfo) IsExists() bool {
	if f == nil || len(f.path) == 0 {
		return false
	}

	info, err := os.Stat(f.path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func (f *FileInfo) Exists() bool {
	return f.IsExists()
}

func (f *FileInfo) Stat() FileInfoResult {
	if f == nil || len(f.path) == 0 {
		return FileInfoFailureMsg(errtype.Validation, "", "nil or empty file path")
	}

	return Stat(f.path)
}

func (f *FileInfo) Abs() *FileInfo {
	if f == nil || len(f.path) == 0 {
		return NewFileInfo(CurrentDir)
	}

	abs, err := filepath.Abs(f.path)
	if err != nil {
		return f
	}

	return NewFileInfo(abs)
}

func (f *FileInfo) ReadBytes() BytesResult {
	if f == nil || len(f.path) == 0 {
		return BytesFailureMsg(errtype.Validation, "", "nil or empty file path")
	}

	return ReadAll(f.path)
}

func (f *FileInfo) ReadString() StringResult {
	if f == nil || len(f.path) == 0 {
		return StringFailureMsg(errtype.Validation, "", "nil or empty file path")
	}

	return ReadString(f.path)
}

func (f *FileInfo) ReadLines() LinesResult {
	if f == nil || len(f.path) == 0 {
		return LinesFailureMsg(errtype.Validation, "", "nil or empty file path")
	}

	return ReadLines(f.path)
}

func (f *FileInfo) WriteBytes(data []byte, perm FilePermType) BoolResult {
	if f == nil || len(f.path) == 0 {
		return BoolFailureMsg(errtype.Validation, "", "nil or empty file path")
	}

	return WriteFile(f.path, data, perm)
}

func (f *FileInfo) WriteString(content string, perm FilePermType) BoolResult {
	if f == nil || len(f.path) == 0 {
		return BoolFailureMsg(errtype.Validation, "", "nil or empty file path")
	}

	return WriteString(f.path, content, perm)
}

func (f *FileInfo) Delete() BoolResult {
	if f == nil || len(f.path) == 0 {
		return BoolFailureMsg(errtype.Validation, "", "nil or empty file path")
	}

	return DeleteFile(f.path)
}
