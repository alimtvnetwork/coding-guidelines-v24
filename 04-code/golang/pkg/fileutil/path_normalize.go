package fileutil

import (
	"os"
	"path/filepath"
	"strings"
)

func ToSlash(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}

func ToBackslash(path string) string {
	return strings.ReplaceAll(path, "/", "\\")
}

func ToNative(path string) string {
	if os.PathSeparator == '/' {
		return ToSlash(path)
	}

	return ToBackslash(path)
}

func HasLongPathPrefix(path string) bool {
	return strings.HasPrefix(path, `\\?\`)
}

func TrimLongPathPrefix(path string) string {
	if HasLongPathPrefix(path) {
		return strings.TrimPrefix(path, `\\?\`)
	}

	return path
}

func isWindowsDriveLetter(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}

	return b >= 'A' && b <= 'Z'
}

func isWindowsDriveAbs(p string) bool {
	if len(p) < 3 {
		return false
	}

	if !isWindowsDriveLetter(p[0]) {
		return false
	}

	return p[1] == ':' && (p[2] == '\\' || p[2] == '/')
}

func ToLongPath(path string) string {
	if HasLongPathPrefix(path) {
		return path
	}

	if strings.HasPrefix(path, `\\`) {
		return `\\?\UNC\` + strings.TrimPrefix(path, `\\`)
	}

	if isWindowsDriveAbs(path) {
		return `\\?\` + ToBackslash(path)
	}

	return path
}

func findDedupePrefix(path string) (string, int) {
	if strings.HasPrefix(path, `\\?\`) {
		return `\\?\`, 4
	}

	if strings.HasPrefix(path, `\\`) {
		return `\\`, 2
	}

	return "", 0
}

func isSeparatorByte(c byte) bool {
	return c == '/' || c == '\\'
}

func appendNonDuplicate(sb *strings.Builder, path string, i int) {
	c := path[i]
	if isSeparatorByte(c) {
		if i > 0 && isSeparatorByte(path[i-1]) {
			return
		}
	}

	sb.WriteByte(c)
}

func DeduplicateSeparators(path string) string {
	if len(path) <= 1 {
		return path
	}

	prefix, startIdx := findDedupePrefix(path)
	var sb strings.Builder
	sb.WriteString(prefix)
	for i := startIdx; i < len(path); i++ {
		appendNonDuplicate(&sb, path, i)
	}

	return sb.String()
}

func cleanLongPath(path string) string {
	trimmed := TrimLongPathPrefix(path)

	return `\\?\` + filepath.Clean(trimmed)
}

func Clean(path string) string {
	if len(path) == 0 {
		return "."
	}

	if HasLongPathPrefix(path) {
		return cleanLongPath(path)
	}

	return filepath.Clean(path)
}

func Normalize(path string) StringResult {
	cleaned := Clean(path)
	native := ToNative(cleaned)
	deduped := DeduplicateSeparators(native)

	return StringSuccess(deduped)
}

func NormalizeToSlash(path string) StringResult {
	cleaned := Clean(path)
	slashed := ToSlash(cleaned)
	deduped := DeduplicateSeparators(slashed)

	return StringSuccess(deduped)
}

type pathNormNamespace struct{}

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
