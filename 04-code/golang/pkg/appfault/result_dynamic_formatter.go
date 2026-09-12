package appfault

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

const maxFormatDepth = 32

// FormatValue recursively formats any value with sorted map keys and unwrapped Result monads.
func FormatValue(val any) string {
	return formatValueWithDepth(val, 0)
}

func isNilPointer(val any) bool {
	if val == nil {
		return false
	}

	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Pointer {
		return false
	}

	return rv.IsNil()
}

func isNilOrNilPointer(val any) bool {
	if val == nil {
		return true
	}

	return isNilPointer(val)
}

func formatError(appErr *AppError) string {
	if appErr == nil {
		return "[Error]"
	}

	msg := appErr.GetMessage()
	if len(msg) == 0 {
		return "[Error]"
	}

	return fmt.Sprintf("[Error: %s]", msg)
}

func formatResultInspector(res ResultInspector, depth int) string {
	if res.IsFailed() {
		return formatError(res.AppError())
	}

	return formatValueWithDepth(res.ValueAny(), depth+1)
}

func formatDirect(val any) (string, bool) {
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

func sortMapKeys(keys []reflect.Value) {
	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprint(keys[i].Interface()) < fmt.Sprint(keys[j].Interface())
	})
}

func formatMapEntries(rv reflect.Value, keys []reflect.Value, depth int) []string {
	entries := make([]string, 0, len(keys))
	for _, k := range keys {
		keyStr := fmt.Sprint(k.Interface())
		valStr := formatValueWithDepth(rv.MapIndex(k).Interface(), depth+1)
		entries = append(entries, keyStr+":"+valStr)
	}

	return entries
}

func formatMapValue(rv reflect.Value, depth int) string {
	keys := rv.MapKeys()
	sortMapKeys(keys)
	entries := formatMapEntries(rv, keys, depth)

	return "map[" + strings.Join(entries, " ") + "]"
}

func formatSliceElements(rv reflect.Value, depth int) []string {
	n := rv.Len()
	items := make([]string, 0, n)
	for i := 0; i < n; i++ {
		elemStr := formatValueWithDepth(rv.Index(i).Interface(), depth+1)
		items = append(items, elemStr)
	}

	return items
}

func formatSliceValue(rv reflect.Value, depth int) string {
	items := formatSliceElements(rv, depth)

	return "[" + strings.Join(items, " ") + "]"
}

func formatReflectValue(rv reflect.Value, depth int) (string, bool) {
	switch rv.Kind() {
	case reflect.Map:
		return formatMapValue(rv, depth), true
	case reflect.Slice, reflect.Array:
		return formatSliceValue(rv, depth), true
	default:
		return "", false
	}
}

func formatTypedValue(val any, depth int) string {
	if s, isDirect := formatDirect(val); isDirect {
		return s
	}

	rv := reflect.ValueOf(val)
	if formatted, isReflect := formatReflectValue(rv, depth); isReflect {
		return formatted
	}

	return fmt.Sprintf("%v", val)
}

func tryAsResultInspector(val any) (ResultInspector, bool) {
	if val == nil {
		return nil, false
	}

	if ri, ok := val.(ResultInspector); ok {
		return ri, true
	}

	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Struct {
		return nil, false
	}

	ptr := reflect.New(rv.Type())
	ptr.Elem().Set(rv)
	if ri, ok := ptr.Interface().(ResultInspector); ok {
		return ri, true
	}

	return nil, false
}

func formatValueWithDepth(val any, depth int) string {
	if depth > maxFormatDepth {
		return "..."
	}

	if isNilOrNilPointer(val) {
		return "<nil>"
	}

	if res, isRes := tryAsResultInspector(val); isRes {
		return formatResultInspector(res, depth)
	}

	return formatTypedValue(val, depth)
}

// UnwrapRecursive deeply unwraps ResultInspector instances, maps, and slices into plain Go types.
func UnwrapRecursive(val any) any {
	return unwrapWithDepth(val, 0)
}

func unwrapFailedResult(appErr *AppError) map[string]any {
	msg := ""
	if appErr != nil {
		msg = appErr.GetMessage()
	}

	return map[string]any{"error": msg}
}

func unwrapResultInspector(res ResultInspector, depth int) any {
	if res.IsFailed() {
		return unwrapFailedResult(res.AppError())
	}

	return unwrapWithDepth(res.ValueAny(), depth+1)
}

func unwrapMapValue(rv reflect.Value, depth int) map[string]any {
	out := make(map[string]any, rv.Len())
	for _, k := range rv.MapKeys() {
		keyStr := fmt.Sprint(k.Interface())
		valAny := rv.MapIndex(k).Interface()
		out[keyStr] = unwrapWithDepth(valAny, depth+1)
	}

	return out
}

func unwrapSliceValue(rv reflect.Value, depth int) []any {
	n := rv.Len()
	out := make([]any, n)
	for i := 0; i < n; i++ {
		out[i] = unwrapWithDepth(rv.Index(i).Interface(), depth+1)
	}

	return out
}

func unwrapReflectKind(rv reflect.Value, depth int, fallback any) any {
	switch rv.Kind() {
	case reflect.Map:
		return unwrapMapValue(rv, depth)
	case reflect.Slice, reflect.Array:
		return unwrapSliceValue(rv, depth)
	default:
		return fallback
	}
}

func unwrapReflectValue(val any, depth int) any {
	if b, isBytes := val.([]byte); isBytes {
		return string(b)
	}

	return unwrapReflectKind(reflect.ValueOf(val), depth, val)
}

func unwrapWithDepth(val any, depth int) any {
	if depth > maxFormatDepth {
		return "..."
	}

	if isNilOrNilPointer(val) {
		return nil
	}

	if res, isRes := tryAsResultInspector(val); isRes {
		return unwrapResultInspector(res, depth)
	}

	return unwrapReflectValue(val, depth)
}

// FormatSortedJson formats any value to indented JSON with unwrapped Results and sorted map keys.
func FormatSortedJson(val any) string {
	unwrapped := UnwrapRecursive(val)
	bytes, err := json.MarshalIndent(unwrapped, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(bytes)
}

// FormatSortedJSON is an alias for FormatSortedJson.
func FormatSortedJSON(val any) string {
	return FormatSortedJson(val)
}

// FormatSortedCompactJson formats any value to compact JSON with unwrapped Results and sorted map keys.
func FormatSortedCompactJson(val any) string {
	unwrapped := UnwrapRecursive(val)
	bytes, err := json.Marshal(unwrapped)
	if err != nil {
		return "{}"
	}

	return string(bytes)
}

// FormatSortedCompactJSON is an alias for FormatSortedCompactJson.
func FormatSortedCompactJSON(val any) string {
	return FormatSortedCompactJson(val)
}
