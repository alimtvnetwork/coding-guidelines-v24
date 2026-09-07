package appfault

import (
	"strings"
	"testing"

	"coding-guidelines/common/pkg/errtype"
)

func TestResultSlice_Filter_Success(t *testing.T) {
	s := OkSlice([]int{1, 2, 3, 4})
	filtered := s.Filter(func(item int) bool {
		return item%2 == 0
	})

	if filtered.Count() != 2 {
		t.Fatalf("expected 2 items, got %d", filtered.Count())
	}

	if filtered.Items[0] != 2 || filtered.Items[1] != 4 {
		t.Fatalf("unexpected items: %v", filtered.Items)
	}
}

func TestResultSlice_Filter_Failure(t *testing.T) {
	err := New(errtype.Validation, "slice fail")
	s := FailSlice[int](err)
	filtered := s.Filter(func(item int) bool {
		return true
	})

	if filtered.IsSuccess() {
		t.Fatalf("expected filtered failure to remain failed")
	}

	if filtered.Fault().GetMessage() != "slice fail" {
		t.Fatalf("unexpected error message: %s", filtered.Fault().GetMessage())
	}
}

func TestResultSlice_ForEach_Success(t *testing.T) {
	s := OkSlice([]string{"a", "b"})
	var collected []string
	res := s.ForEach(func(_ int, item string) {
		collected = append(collected, item)
	})

	if res.Count() != 2 {
		t.Fatalf("expected returned slice count 2")
	}

	if len(collected) != 2 || collected[0] != "a" {
		t.Fatalf("unexpected collected: %v", collected)
	}
}

func TestResultSlice_ForEach_Failure(t *testing.T) {
	s := FailSlice[int](New(errtype.Validation, "fail"))
	called := false
	s.ForEach(func(_ int, _ int) {
		called = true
	})

	if called {
		t.Fatalf("ForEach should not execute on failed ResultSlice")
	}
}

func TestResultSlice_ForEachBreak_Success(t *testing.T) {
	s := OkSlice([]int{10, 20, 30, 40})
	var visited []int
	s.ForEachBreak(func(_ int, item int) bool {
		visited = append(visited, item)

		return item == 20
	})

	if len(visited) != 2 {
		t.Fatalf("expected visited len 2, got %d", len(visited))
	}
}

func TestResultSlice_ForEachBreak_Failure(t *testing.T) {
	s := FailSlice[int](New(errtype.Validation, "fail"))
	called := false
	s.ForEachBreak(func(_ int, _ int) bool {
		called = true

		return true
	})

	if called {
		t.Fatalf("ForEachBreak should not execute on failed ResultSlice")
	}
}

func TestResultSlice_FormatStruct_Success(t *testing.T) {
	s := OkSlice([]int{1, 2})
	out := s.FormatStruct()
	if !strings.Contains(out, "[0] 1") {
		t.Fatalf("expected format to contain [0] 1, got %q", out)
	}

	empty := OkSlice([]int{})
	if empty.FormatStruct() != "[]" {
		t.Fatalf("expected [], got %q", empty.FormatStruct())
	}
}

func TestResultSlice_FormatStruct_Failure(t *testing.T) {
	s := FailSlice[int](New(errtype.NotFound, "not found item"))
	out := s.FormatStruct()
	if !strings.Contains(out, "not found item") {
		t.Fatalf("expected error banner to contain 'not found item', got %q", out)
	}
}

func TestResultMap_Keys_Success(t *testing.T) {
	m := OkMap(map[string]int{"b": 2, "a": 1})
	keys := m.Keys()
	if len(keys) != 2 {
		t.Fatalf("expected 2 keys, got %d", len(keys))
	}

	if keys[0] != "a" || keys[1] != "b" {
		t.Fatalf("expected sorted keys [a b], got %v", keys)
	}
}

func TestResultMap_Keys_Failure(t *testing.T) {
	m := FailMap[string, int](New(errtype.Validation, "map error"))
	keys := m.Keys()
	if len(keys) != 0 {
		t.Fatalf("expected empty keys on failure, got %v", keys)
	}
}

func TestResultMap_Values_Success(t *testing.T) {
	m := OkMap(map[string]int{"a": 1, "b": 2})
	vals := m.Values()
	if len(vals) != 2 {
		t.Fatalf("expected 2 values, got %d", len(vals))
	}

	if vals[0] != 1 || vals[1] != 2 {
		t.Fatalf("expected [1 2], got %v", vals)
	}
}

func TestResultMap_Values_Failure(t *testing.T) {
	m := FailMap[string, int](New(errtype.Validation, "map error"))
	vals := m.Values()
	if len(vals) != 0 {
		t.Fatalf("expected empty values on failure, got %v", vals)
	}
}

func TestResultMap_Filter_Success(t *testing.T) {
	m := OkMap(map[string]int{"alpha": 1, "beta": 2, "gamma": 3})
	filtered := m.Filter(func(k string, v int) bool {
		return v >= 2
	})

	if filtered.Count() != 2 {
		t.Fatalf("expected count 2, got %d", filtered.Count())
	}

	if filtered.Has("alpha") {
		t.Fatalf("alpha should be filtered out")
	}
}

func TestResultMap_Filter_Failure(t *testing.T) {
	m := FailMap[string, int](New(errtype.Validation, "filter fail"))
	filtered := m.Filter(func(k string, v int) bool {
		return true
	})

	if filtered.IsSuccess() {
		t.Fatalf("expected failure to remain")
	}
}

func TestResultMap_ForEach_Success(t *testing.T) {
	m := OkMap(map[string]int{"x": 10, "y": 20})
	total := 0
	res := m.ForEach(func(_ string, v int) {
		total += v
	})

	if total != 30 || res.Count() != 2 {
		t.Fatalf("unexpected total %d or count %d", total, res.Count())
	}
}

func TestResultMap_ForEach_Failure(t *testing.T) {
	m := FailMap[string, int](New(errtype.Execution, "failed"))
	called := false
	m.ForEach(func(_ string, _ int) {
		called = true
	})

	if called {
		t.Fatalf("ForEach should not execute on failed ResultMap")
	}
}

func TestResultMap_FormatStruct_Success(t *testing.T) {
	m := OkMap(map[string]int{"k": 1})
	out := m.FormatStruct()
	if !strings.Contains(out, "k: 1") {
		t.Fatalf("expected k: 1, got %q", out)
	}

	empty := OkMap(map[string]int{})
	if empty.FormatStruct() != "{}" {
		t.Fatalf("expected {}, got %q", empty.FormatStruct())
	}
}

func TestResultMap_FormatStruct_Failure(t *testing.T) {
	m := FailMap[string, int](New(errtype.NotFound, "missing map"))
	out := m.FormatStruct()
	if !strings.Contains(out, "missing map") {
		t.Fatalf("expected error banner to contain 'missing map', got %q", out)
	}
}

func TestResult_FormatStruct_Success(t *testing.T) {
	res := NewSuccess(map[string]string{"foo": "bar"})
	out := res.FormatStruct()
	if !strings.Contains(out, "foo:bar") {
		t.Fatalf("expected formatted string containing foo:bar, got %q", out)
	}
}

func TestResult_FormatStruct_Failure(t *testing.T) {
	res := FailureResult[string](New(errtype.Internal, "boom"))
	out := res.FormatStruct()
	if !strings.Contains(out, "boom") {
		t.Fatalf("expected formatted error banner containing boom, got %q", out)
	}
}

type testUserPayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestResult_ToMap_Success(t *testing.T) {
	u := testUserPayload{Name: "Alice", Age: 30}
	res := NewSuccess(u)
	m := res.ToMap()
	if m == nil || m["name"] != "Alice" || m["age"] != float64(30) {
		t.Fatalf("unexpected map result: %v", m)
	}
}

func TestResult_ToMap_Failure(t *testing.T) {
	res := FailureResult[testUserPayload](New(errtype.Validation, "invalid user"))
	m := res.ToMap()
	if m != nil {
		t.Fatalf("expected nil map on failed Result, got %v", m)
	}
}
