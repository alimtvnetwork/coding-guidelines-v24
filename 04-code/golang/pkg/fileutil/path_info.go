package fileutil

import (
	"path/filepath"
	"regexp"
	"strings"
)

var (
	slugNonAlphaNumRegex = regexp.MustCompile(PatternSlugNonAlphaNum)
	slugDashesRegex      = regexp.MustCompile(PatternSlugDashes)
)

func Ext(path string) string {
	return filepath.Ext(path)
}

func ExtNoDot(path string) string {
	return strings.TrimPrefix(Ext(path), ExtDot)
}

func HasExt(path string, ext string) bool {
	cleanExpected := strings.TrimPrefix(ext, ExtDot)

	return strings.EqualFold(ExtNoDot(path), cleanExpected)
}

func Base(path string) string {
	return filepath.Base(ToNative(path))
}

func Stem(path string) string {
	base := Base(path)
	ext := filepath.Ext(base)
	if base == ext {
		return base
	}

	return strings.TrimSuffix(base, ext)
}

func findFirstExtDot(base string) int {
	if strings.HasPrefix(base, ExtDot) {
		idx := strings.Index(base[1:], ExtDot)
		if idx >= 0 {
			return idx + 1
		}

		return -1
	}

	return strings.Index(base, ExtDot)
}

func StemFull(path string) string {
	base := Base(path)
	if base == CurrentDir || base == ParentDir {
		return base
	}

	dotIdx := findFirstExtDot(base)
	if dotIdx >= 0 {
		return base[:dotIdx]
	}

	return base
}

func Slug(path string) string {
	name := Stem(path)
	slug := strings.ToLower(name)
	slug = slugNonAlphaNumRegex.ReplaceAllString(slug, SlugHyphen)
	slug = slugDashesRegex.ReplaceAllString(slug, SlugHyphen)

	return strings.Trim(slug, SlugHyphen)
}

func Dir(path string) string {
	return filepath.Dir(ToNative(path))
}

func Split(path string) (string, string) {
	return filepath.Split(ToNative(path))
}

func Parent(path string) string {
	return filepath.Dir(Clean(path))
}

func ParentN(path string, levels int) string {
	cur := Clean(path)
	if levels <= 0 {
		return cur
	}

	for i := 0; i < levels; i++ {
		cur = filepath.Dir(cur)
	}

	return cur
}

func IsAbs(path string) bool {
	if filepath.IsAbs(path) {
		return true
	}

	if isWindowsDriveAbs(path) {
		return true
	}

	if strings.HasPrefix(path, SepSlash) {
		return true
	}

	return HasLongPathPrefix(path)
}

func IsRel(path string) bool {
	return !IsAbs(path)
}

type pathInfoNamespace struct{}

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

func (pathInfoNamespace) Folder(path string) *FolderInfo {
	return NewFolderInfo(path)
}

func (pathInfoNamespace) FolderInfo(path string) *FolderInfo {
	return NewFolderInfo(path)
}

func (pathInfoNamespace) File(path string) *FileInfo {
	return NewFileInfo(path)
}

func (pathInfoNamespace) FileInfo(path string) *FileInfo {
	return NewFileInfo(path)
}

func (pathInfoNamespace) Inspect(path string) *PathInfo {
	return NewPathInfo(path)
}
