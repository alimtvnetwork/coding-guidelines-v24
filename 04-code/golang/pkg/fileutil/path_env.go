package fileutil

import (
	"os"
	"path/filepath"
	"strings"

	"coding-guidelines/common/pkg/errtype"
)

func isLetterChar(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}

	return b >= 'A' && b <= 'Z'
}

func isEnvIdentifierChar(b byte) bool {
	if isLetterChar(b) {
		return true
	}

	if b >= '0' && b <= '9' {
		return true
	}

	return b == '_' || b == '(' || b == ')' || b == '-'
}

func isValidEnvName(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if !isEnvIdentifierChar(s[i]) {
			return false
		}
	}

	return true
}

func isPosixIdentStart(b byte) bool {
	if isLetterChar(b) {
		return true
	}

	return b == '_'
}

func isPosixIdentChar(b byte) bool {
	if isPosixIdentStart(b) {
		return true
	}

	return b >= '0' && b <= '9'
}

func lookupEnv(key string) (string, bool) {
	val, ok := os.LookupEnv(key)
	if ok {
		return val, true
	}

	lowerKey := strings.ToLower(key)

	for _, env := range os.Environ() {
		idx := strings.IndexByte(env, CharEqual)
		if idx > 0 && strings.ToLower(env[:idx]) == lowerKey {
			return env[idx+1:], true
		}
	}

	return "", false
}

func parseWindowsVar(s string, start int) (string, int, bool) {
	end := strings.IndexByte(s[start+1:], CharPercent)
	if end < 0 {
		return "", start, false
	}

	next := start + 1 + end
	name := s[start+1 : next]
	if !isValidEnvName(name) {
		return "", start, false
	}

	return name, next + 1, true
}

func parsePosixBraced(s string, start int) (string, int, bool) {
	end := strings.IndexByte(s[start+2:], CharRightBrace)
	if end < 0 {
		return "", start, false
	}

	next := start + 2 + end
	name := s[start+2 : next]
	if !isValidEnvName(name) {
		return "", start, false
	}

	return name, next + 1, true
}

func parsePosixIdent(s string, start int) (string, int, bool) {
	idx := start + 1
	if idx >= len(s) {
		return "", start, false
	}

	if !isPosixIdentStart(s[idx]) {
		return "", start, false
	}

	idx++
	for idx < len(s) && isPosixIdentChar(s[idx]) {
		idx++
	}

	return s[start+1 : idx], idx, true
}

func scanDollar(s string, idx int, b *strings.Builder) int {
	if idx+1 < len(s) && s[idx+1] == CharLeftBrace {
		name, next, ok := parsePosixBraced(s, idx)
		if ok {
			val, _ := lookupEnv(name)
			b.WriteString(val)

			return next
		}
	}

	name, next, ok := parsePosixIdent(s, idx)
	if ok {
		val, _ := lookupEnv(name)
		b.WriteString(val)

		return next
	}

	b.WriteByte(CharDollar)

	return idx + 1
}

func scanPercent(s string, idx int, b *strings.Builder) int {
	name, next, ok := parseWindowsVar(s, idx)
	if ok {
		val, _ := lookupEnv(name)
		b.WriteString(val)

		return next
	}

	b.WriteByte(CharPercent)

	return idx + 1
}

func expandEnvString(path string) string {
	var b strings.Builder
	b.Grow(len(path))

	idx := 0
	for idx < len(path) {
		switch path[idx] {
		case CharPercent:
			idx = scanPercent(path, idx, &b)
		case CharDollar:
			idx = scanDollar(path, idx, &b)
		default:
			b.WriteByte(path[idx])
			idx++
		}
	}

	return b.String()
}

func ExpandEnv(path string) StringResult {
	if len(path) == 0 {
		return StringSuccess("")
	}

	return StringSuccess(expandEnvString(path))
}

func hasTildePrefix(path string) bool {
	if strings.HasPrefix(path, PrefixTildeSlash) {
		return true
	}

	return strings.HasPrefix(path, PrefixTildeBackslash)
}

func resolveHomePath(raw string, sub string) StringResult {
	home, err := os.UserHomeDir()
	if err != nil {
		return StringFailure(errtype.IO, err, raw, "failed to resolve user home directory")
	}

	if len(sub) == 0 {
		return StringSuccess(home)
	}

	return StringSuccess(filepath.Join(home, sub))
}

func ExpandTilde(path string) StringResult {
	if path == HomeTilde {
		return resolveHomePath(path, "")
	}

	if hasTildePrefix(path) {
		return resolveHomePath(path, path[len(PrefixTildeSlash):])
	}

	return StringSuccess(path)
}

func Expand(path string) StringResult {
	tildeRes := ExpandTilde(path)
	if tildeRes.IsFailed() {
		return tildeRes
	}

	return ExpandEnv(tildeRes.Data())
}

type pathEnvNamespace struct{}

func (pathEnvNamespace) Expand(path string) StringResult {
	return Expand(path)
}

func (pathEnvNamespace) ExpandEnv(path string) StringResult {
	return ExpandEnv(path)
}

func (pathEnvNamespace) ExpandTilde(path string) StringResult {
	return ExpandTilde(path)
}
