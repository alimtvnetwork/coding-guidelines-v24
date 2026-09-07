package fileutil

import (
	"path/filepath"
	"regexp"
	"strings"
)

var (
	slugNonAlphaNumRegex = regexp.MustCompile(`[^a-z0-9-]+`)
	slugDashesRegex      = regexp.MustCompile(`-{2,}`)
)

func Ext(path string) string {
	return filepath.Ext(path)
}

func ExtNoDot(path string) string {
	return strings.TrimPrefix(Ext(path), ".")
}

func HasExt(path string, ext string) bool {
	cleanExpected := strings.TrimPrefix(ext, ".")

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
	if strings.HasPrefix(base, ".") {
		idx := strings.Index(base[1:], ".")
		if idx >= 0 {
			return idx + 1
		}

		return -1
	}

	return strings.Index(base, ".")
}

func StemFull(path string) string {
	base := Base(path)
	if base == "." || base == ".." {
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
	slug = slugNonAlphaNumRegex.ReplaceAllString(slug, "-")
	slug = slugDashesRegex.ReplaceAllString(slug, "-")

	return strings.Trim(slug, "-")
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

	if strings.HasPrefix(path, "/") {
		return true
	}

	return HasLongPathPrefix(path)
}

func IsRel(path string) bool {
	if IsAbs(path) {
		return false
	}

	return true
}
