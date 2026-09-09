package fileutil

import (
	"os"
	"path/filepath"

	"coding-guidelines/common/pkg/errtype"
)

type pathNamespace struct {
	Temp pathTempNamespace
	Env  pathEnvNamespace
	Norm pathNormNamespace
	Info pathInfoNamespace
}

var Path = pathNamespace{
	Temp: pathTempNamespace{},
	Env:  pathEnvNamespace{},
	Norm: pathNormNamespace{},
	Info: pathInfoNamespace{},
}

func (pathNamespace) Join(elem ...string) string {
	return filepath.Join(elem...)
}

func (pathNamespace) Folder(path string) *FolderInfo {
	return NewFolderInfo(path)
}

func (pathNamespace) FolderInfo(path string) *FolderInfo {
	return NewFolderInfo(path)
}

func (pathNamespace) File(path string) *FileInfo {
	return NewFileInfo(path)
}

func (pathNamespace) FileInfo(path string) *FileInfo {
	return NewFileInfo(path)
}

func (pathNamespace) Inspect(path string) *PathInfo {
	return NewPathInfo(path)
}

type PathWrapper struct {
	raw  string
	path string
}

func NewPath(raw string) *PathWrapper {
	return &PathWrapper{
		raw:  raw,
		path: raw,
	}
}

func (p *PathWrapper) Raw() string {
	return p.raw
}

func (p *PathWrapper) String() string {
	return p.path
}

func (p *PathWrapper) Clean() *PathWrapper {
	p.path = Clean(p.path)

	return p
}

func (p *PathWrapper) Normalize() *PathWrapper {
	res := Normalize(p.path)
	if res.IsSuccess() {
		p.path = res.Data()
	}

	return p
}

func (p *PathWrapper) ToSlash() *PathWrapper {
	p.path = ToSlash(p.path)

	return p
}

func (p *PathWrapper) ToBackslash() *PathWrapper {
	p.path = ToBackslash(p.path)

	return p
}

func (p *PathWrapper) ToNative() *PathWrapper {
	p.path = ToNative(p.path)

	return p
}

func (p *PathWrapper) Expand() *PathWrapper {
	res := Expand(p.path)
	if res.IsSuccess() {
		p.path = res.Data()
	}

	return p
}

func (p *PathWrapper) ExpandEnv() *PathWrapper {
	res := ExpandEnv(p.path)
	if res.IsSuccess() {
		p.path = res.Data()
	}

	return p
}

func (p *PathWrapper) ExpandTilde() *PathWrapper {
	res := ExpandTilde(p.path)
	if res.IsSuccess() {
		p.path = res.Data()
	}

	return p
}

func (p *PathWrapper) Join(elem ...string) *PathWrapper {
	parts := append([]string{p.path}, elem...)
	p.path = filepath.Join(parts...)

	return p
}

func (p *PathWrapper) Parent() *PathWrapper {
	p.path = Parent(p.path)

	return p
}

func (p *PathWrapper) ParentN(levels int) *PathWrapper {
	p.path = ParentN(p.path, levels)

	return p
}

func (p *PathWrapper) Base() string {
	return Base(p.path)
}

func (p *PathWrapper) Stem() string {
	return Stem(p.path)
}

func (p *PathWrapper) StemFull() string {
	return StemFull(p.path)
}

func (p *PathWrapper) Ext() string {
	return Ext(p.path)
}

func (p *PathWrapper) ExtNoDot() string {
	return ExtNoDot(p.path)
}

func (p *PathWrapper) HasExt(ext string) bool {
	return HasExt(p.path, ext)
}

func (p *PathWrapper) Dir() string {
	return Dir(p.path)
}

func (p *PathWrapper) Slug() string {
	return Slug(p.path)
}

func (p *PathWrapper) IsAbs() bool {
	return IsAbs(p.path)
}

func (p *PathWrapper) IsRel() bool {
	return IsRel(p.path)
}

func (p *PathWrapper) Exists() BoolResult {
	if _, err := os.Stat(p.path); err != nil {
		if os.IsNotExist(err) {
			return BoolSuccess(false)
		}

		return BoolFailure(errtype.IO, err, p.path, "failed to check existence")
	}

	return BoolSuccess(true)
}

func (p *PathWrapper) Stat() FileInfoResult {
	return Stat(p.path)
}

func (p *PathWrapper) Read() BytesResult {
	return ReadAll(p.path)
}

func (p *PathWrapper) ReadString() StringResult {
	return ReadString(p.path)
}

func (p *PathWrapper) Write(data any, perm FilePermType) BoolResult {
	return Write(p.path, data, perm)
}

func (p *PathWrapper) WriteString(content string, perm FilePermType) BoolResult {
	return WriteString(p.path, content, perm)
}

func (p *PathWrapper) Folder() *FolderInfo {
	return NewFolderInfo(p.path)
}

func (p *PathWrapper) File() *FileInfo {
	return NewFileInfo(p.path)
}

func (p *PathWrapper) Inspect() *PathInfo {
	return NewPathInfo(p.path)
}

type filePathCreator struct{}

func (filePathCreator) Default(raw string) *PathWrapper {
	return NewPath(raw)
}

func (filePathCreator) FromParts(elem ...string) *PathWrapper {
	return NewPath(filepath.Join(elem...))
}

func (filePathCreator) Folder(path string) *FolderInfo {
	return NewFolderInfo(path)
}

func (filePathCreator) File(path string) *FileInfo {
	return NewFileInfo(path)
}

func (fileNewCreator) PathWrapper(raw string) *PathWrapper {
	return NewPath(raw)
}

func (fileNewCreator) Folder(path string) *FolderInfo {
	return NewFolderInfo(path)
}

func (fileNewCreator) FolderInfo(path string) *FolderInfo {
	return NewFolderInfo(path)
}

func (fileNewCreator) FileInfo(path string) *FileInfo {
	return NewFileInfo(path)
}
