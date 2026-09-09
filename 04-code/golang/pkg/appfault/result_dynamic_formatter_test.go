package appfault

import (
	"strings"
	"testing"

	"coding-guidelines/common/pkg/errtype"
)

func TestFormatValue_NilAndPrimitives(t *testing.T) {
	if got := FormatValue(nil); got != "<nil>" {
		t.Fatalf("expected <nil>, got %q", got)
	}

	if got := FormatValue("hello"); got != "hello" {
		t.Fatalf("expected hello, got %q", got)
	}

	if got := FormatValue([]byte("world")); got != "world" {
		t.Fatalf("expected world, got %q", got)
	}

	if got := FormatValue(42); got != "42" {
		t.Fatalf("expected 42, got %q", got)
	}
}

func TestFormatValue_ResultSuccess(t *testing.T) {
	res := SuccessResult("inner-value")
	if got := FormatValue(res); got != "inner-value" {
		t.Fatalf("expected inner-value, got %q", got)
	}

	nested := SuccessResult(SuccessResult(99))
	if got := FormatValue(nested); got != "99" {
		t.Fatalf("expected 99, got %q", got)
	}
}

func TestFormatValue_ResultFailure(t *testing.T) {
	fail := FailureResult[string](New(errtype.Validation, "invalid input"))
	if got := FormatValue(fail); got != "[Error: invalid input]" {
		t.Fatalf("expected [Error: invalid input], got %q", got)
	}

	emptyErr := FailureResult[string](New(errtype.Validation, ""))
	if got := FormatValue(emptyErr); got != "[Error]" {
		t.Fatalf("expected [Error], got %q", got)
	}
}

func TestFormatValue_SortedMapKeys(t *testing.T) {
	data := map[string]int{"z": 3, "a": 1, "m": 2}
	got := FormatValue(data)
	if got != "map[a:1 m:2 z:3]" {
		t.Fatalf("expected map[a:1 m:2 z:3], got %q", got)
	}

	emptyMap := map[string]int{}
	if got := FormatValue(emptyMap); got != "map[]" {
		t.Fatalf("expected map[], got %q", got)
	}
}

func TestFormatValue_NestedMapWithResults(t *testing.T) {
	m := map[string]any{
		"b": SuccessResult("ok"),
		"a": FailureResult[string](New(errtype.NotFound, "missing")),
	}

	got := FormatValue(m)
	expected := "map[a:[Error: missing] b:ok]"
	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}

func TestFormatValue_SlicesAndArrays(t *testing.T) {
	slice := []int{3, 1, 2}
	if got := FormatValue(slice); got != "[3 1 2]" {
		t.Fatalf("expected [3 1 2], got %q", got)
	}

	arr := [2]string{"hello", "world"}
	if got := FormatValue(arr); got != "[hello world]" {
		t.Fatalf("expected [hello world], got %q", got)
	}

	emptySlice := []string{}
	if got := FormatValue(emptySlice); got != "[]" {
		t.Fatalf("expected [], got %q", got)
	}
}

func TestFormatValue_SliceWithResults(t *testing.T) {
	items := []any{SuccessResult(1), FailureResult[int](New(errtype.IO, "io err"))}
	got := FormatValue(items)
	expected := "[1 [Error: io err]]"
	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}

func TestFormatValue_CycleProtection(t *testing.T) {
	m := make(map[string]any)
	m["self"] = m
	got := FormatValue(m)
	if !strings.Contains(got, "...") {
		t.Fatalf("expected depth cutoff '...' in cycle, got %q", got)
	}
}

func TestFormatValue_ResultContainers(t *testing.T) {
	rs := OkSlice([]string{"x", "y"})
	if got := FormatValue(rs); got != "[x y]" {
		t.Fatalf("expected [x y], got %q", got)
	}

	rm := OkMap(map[string]int{"two": 2, "one": 1})
	if got := FormatValue(rm); got != "map[one:1 two:2]" {
		t.Fatalf("expected map[one:1 two:2], got %q", got)
	}

	failMap := FailMap[string, int](New(errtype.Database, "db down"))
	if got := FormatValue(failMap); got != "[Error: db down]" {
		t.Fatalf("expected [Error: db down], got %q", got)
	}
}

func TestUnwrapRecursive_Primitives(t *testing.T) {
	if got := UnwrapRecursive(123); got != 123 {
		t.Fatalf("expected 123, got %v", got)
	}

	if got := UnwrapRecursive(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	if got := UnwrapRecursive([]byte("abc")); got != "abc" {
		t.Fatalf("expected 'abc', got %v", got)
	}
}

func TestUnwrapRecursive_Results(t *testing.T) {
	okRes := SuccessResult("value")
	if got := UnwrapRecursive(okRes); got != "value" {
		t.Fatalf("expected 'value', got %v", got)
	}

	failRes := FailureResult[string](New(errtype.Network, "timed out"))
	unwrapped, isMap := UnwrapRecursive(failRes).(map[string]any)
	if !isMap {
		t.Fatalf("expected map[string]any, got %v", UnwrapRecursive(failRes))
	}

	if unwrapped["error"] != "timed out" {
		t.Fatalf("expected error key timed out, got %v", unwrapped["error"])
	}
}

func TestUnwrapRecursive_Maps(t *testing.T) {
	m := map[string]any{"a": SuccessResult(10)}
	unwrapped, isMap := UnwrapRecursive(m).(map[string]any)
	if !isMap {
		t.Fatalf("expected map[string]any, got %v", unwrapped)
	}

	if unwrapped["a"] != 10 {
		t.Fatalf("expected unwrapped map with a: 10, got %v", unwrapped)
	}
}

func TestUnwrapRecursive_Slices(t *testing.T) {
	s := []any{SuccessResult("item")}
	unwrapped, isSlice := UnwrapRecursive(s).([]any)
	if !isSlice {
		t.Fatalf("expected []any, got %v", unwrapped)
	}

	if len(unwrapped) != 1 || unwrapped[0] != "item" {
		t.Fatalf("expected [item], got %v", unwrapped)
	}
}

func TestFormatSortedJson(t *testing.T) {
	m := map[string]any{"b": 2, "a": SuccessResult(1)}
	jsonStr := FormatSortedJson(m)
	if !strings.Contains(jsonStr, "\"a\": 1") || !strings.Contains(jsonStr, "\"b\": 2") {
		t.Fatalf("unexpected JSON: %s", jsonStr)
	}

	if FormatSortedJSON(m) != jsonStr {
		t.Fatalf("expected FormatSortedJSON to match FormatSortedJson")
	}
}

func TestFormatSortedCompactJson(t *testing.T) {
	m := map[string]any{"b": 2, "a": 1}
	compact := FormatSortedCompactJson(m)
	expected := "{\"a\":1,\"b\":2}"
	if compact != expected {
		t.Fatalf("expected %s, got %s", expected, compact)
	}

	if FormatSortedCompactJSON(m) != compact {
		t.Fatalf("expected FormatSortedCompactJSON to match FormatSortedCompactJson")
	}
}
