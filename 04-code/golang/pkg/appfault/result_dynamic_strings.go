package appfault

import (
	"fmt"
	"strings"
)

func stringifyDirect(val any) (string, bool) {
	switch v := val.(type) {
	case string:
		return v, true
	case []byte:
		return string(v), true
	case fmt.Stringer:
		return v.String(), true
	default:
		return "", false
	}
}

func stringifyValue(val any) string {
	if val == nil {
		return ""
	}

	if s, isOk := stringifyDirect(val); isOk {
		return s
	}

	return fmt.Sprintf("%v", val)
}

func splitNormalizedLines(raw string) []string {
	if len(raw) == 0 {
		return []string{}
	}

	norm := strings.ReplaceAll(raw, "\r\n", "\n")

	return strings.Split(norm, "\n")
}

func extractStringLines(val any) []string {
	if val == nil {
		return []string{}
	}

	if slice, isStrSlice := val.([]string); isStrSlice {
		return slice
	}

	return splitNormalizedLines(stringifyValue(val))
}

// Lines returns the payload as an array/slice of string lines.
func (r Result[T]) Lines() []string {
	if r.IsFailed() {
		return []string{}
	}

	return extractStringLines(r.value)
}

// LinesResult returns the payload as a monadic ResultSlice of string lines.
func (r Result[T]) LinesResult() ResultSlice[string] {
	if r.IsFailed() {
		return FailSlice[string](r.appError)
	}

	return OkSlice(r.Lines())
}

// Split divides a string payload into substrings separated by sep.
func (r Result[T]) Split(sep string) []string {
	if r.IsFailed() {
		return []string{}
	}

	return strings.Split(stringifyValue(r.value), sep)
}

// SplitResult divides a string payload returning a monadic ResultSlice of substrings.
func (r Result[T]) SplitResult(sep string) ResultSlice[string] {
	if r.IsFailed() {
		return FailSlice[string](r.appError)
	}

	return OkSlice(r.Split(sep))
}

func splitStringAt(str string, index int) (string, string) {
	runes := []rune(str)
	if index <= 0 {
		return "", str
	}

	if index >= len(runes) {
		return str, ""
	}

	return string(runes[:index]), string(runes[index:])
}

// SplitAt splits the string representation at the given character index.
func (r Result[T]) SplitAt(index int) (string, string) {
	if r.IsFailed() {
		return "", ""
	}

	return splitStringAt(stringifyValue(r.value), index)
}

// SplitByRune divides a string payload into substrings separated by a single rune character.
func (r Result[T]) SplitByRune(char rune) []string {
	if r.IsFailed() {
		return []string{}
	}

	return strings.Split(stringifyValue(r.value), string(char))
}
