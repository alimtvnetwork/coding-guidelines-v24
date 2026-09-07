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

func cleanNative(p string) string {
	if HasLongPathPrefix(p) {
		trimmed := TrimLongPathPrefix(p)

		return `\\?\` + filepath.Clean(trimmed)
	}

	return filepath.Clean(p)
}

func Clean(path string) string {
	if len(path) == 0 {
		return "."
	}

	return cleanNative(ToNative(path))
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
