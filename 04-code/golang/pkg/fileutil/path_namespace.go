package fileutil

import (
	"os"
	"path/filepath"

	"coding-guidelines/common/pkg/errtype"
)

type (
	pathTempNamespace struct{}
	pathEnvNamespace  struct{}
	pathNormNamespace struct{}
	pathInfoNamespace struct{}

	pathNamespace struct {
		Temp pathTempNamespace
		Env  pathEnvNamespace
		Norm pathNormNamespace
		Info pathInfoNamespace
	}
)

var Path = pathNamespace{
	Temp: pathTempNamespace{},
	Env:  pathEnvNamespace{},
	Norm: pathNormNamespace{},
	Info: pathInfoNamespace{},
}

func (pathNamespace) Join(elem ...string) string {
	return filepath.Join(elem...)
}

func (pathTempNamespace) UserTempDir() StringResult {
	return UserTempDir()
}

func (pathTempNamespace) UserTempPath(subpath ...string) StringResult {
	return UserTempPath(subpath...)
}

func (pathTempNamespace) TempFile(pattern string) FileResult {
	return TempFile(pattern)
}

func (pathTempNamespace) TempDir(pattern string) StringResult {
	return TempDir(pattern)
}

func (pathTempNamespace) CreateTempFile(dir string, pattern string, perm FilePermType) FileResult {
	return CreateTempFile(dir, pattern, perm)
}

func (pathTempNamespace) CreateTempDir(dir string, pattern string, perm FilePermType) StringResult {
	return CreateTempDir(dir, pattern, perm)
}

func (pathEnvNamespace) Expand(path string) StringResult {
	return Expand(path)
}

func (pathEnvNamespace) ExpandEnv(path string) StringResult {
	return ExpandEnv(path)
}

func (pathEnvNamespace) ExpandTilde(path string) StringResult {
	return ExpandTilde(path)
}

func (pathNormNamespace) Clean(path string) string {
	return Clean(path)
}

func (pathNormNamespace) Normalize(path string) StringResult {
	return Normalize(path)
}

func (pathNormNamespace) NormalizeToSlash(path string) StringResult {
	return NormalizeToSlash(path)
}

func (pathNormNamespace) ToSlash(path string) string {
	return ToSlash(path)
}

func (pathNormNamespace) ToBackslash(path string) string {
	return ToBackslash(path)
}

func (pathNormNamespace) ToNative(path string) string {
	return ToNative(path)
}

func (pathNormNamespace) Deduplicate(path string) string {
	return DeduplicateSeparators(path)
}

func (pathInfoNamespace) Ext(path string) string {
	return Ext(path)
}

func (pathInfoNamespace) ExtNoDot(path string) string {
	return ExtNoDot(path)
}

func (pathInfoNamespace) HasExt(path string, ext string) bool {
	return HasExt(path, ext)
}

func (pathInfoNamespace) Base(path string) string {
	return Base(path)
}

func (pathInfoNamespace) Stem(path string) string {
	return Stem(path)
}

func (pathInfoNamespace) StemFull(path string) string {
	return StemFull(path)
}

func (pathInfoNamespace) Slug(path string) string {
	return Slug(path)
}

func (pathInfoNamespace) Dir(path string) string {
	return Dir(path)
}

func (pathInfoNamespace) Split(path string) (string, string) {
	return Split(path)
}

func (pathInfoNamespace) Parent(path string) string {
	return Parent(path)
}

func (pathInfoNamespace) ParentN(path string, levels int) string {
	return ParentN(path, levels)
}

func (pathInfoNamespace) IsAbs(path string) bool {
	return IsAbs(path)
}

func (pathInfoNamespace) IsRel(path string) bool {
	return IsRel(path)
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
