package result

import (
	"strconv"
	"testing"

	"coding-guidelines/common/pkg/errtype"
)

func TestMap_Success(t *testing.T) {
	succ := Success(42)
	mapped := Map(succ, strconv.Itoa)

	if mapped.IsFailure() {
		t.Fatalf("Map failed on success")
	}

	if mapped.Data() != "42" {
		t.Fatalf("Map returned unexpected value %s", mapped.Data())
	}
}

func TestMap_Failure(t *testing.T) {
	fail := FailureWithId[int](errtype.Validation, "bad int")
	mappedFail := Map(fail, strconv.Itoa)

	if mappedFail.IsSuccess() {
		t.Fatalf("expected failure to persist")
	}

	if mappedFail.Fault().GetMessage() != "bad int" {
		t.Fatalf("unexpected message: %s", mappedFail.Fault().GetMessage())
	}
}

func TestFlatMap_Success(t *testing.T) {
	succ := Success(10)
	flatMapped := FlatMap(succ, func(v int) Result[string] {
		return Success(strconv.Itoa(v * 2))
	})

	if flatMapped.IsFailure() {
		t.Fatalf("expected success")
	}

	if flatMapped.Data() != "20" {
		t.Fatalf("expected 20, got %s", flatMapped.Data())
	}
}

func TestFlatMap_Failure(t *testing.T) {
	succ := Success(10)
	flatMappedFail := FlatMap(succ, func(v int) Result[string] {
		return FailureWithId[string](errtype.Execution, "exec fail")
	})

	if flatMappedFail.IsSuccess() {
		t.Fatalf("expected inner failure")
	}

	if flatMappedFail.Fault().GetMessage() != "exec fail" {
		t.Fatalf("unexpected error message: %s", flatMappedFail.Fault().GetMessage())
	}
}

func TestFlatMap_SkipInitialError(t *testing.T) {
	fail := FailureWithId[int](errtype.Validation, "initial fail")
	flatMappedSkip := FlatMap(fail, func(v int) Result[string] {
		return Success("should not run")
	})

	if flatMappedSkip.IsSuccess() {
		t.Fatalf("expected skip on error")
	}

	if flatMappedSkip.Fault().GetMessage() != "initial fail" {
		t.Fatalf("unexpected message: %s", flatMappedSkip.Fault().GetMessage())
	}
}

func TestTap_Success(t *testing.T) {
	succ := Success("hello")
	tapped := false
	res := Tap(succ, func(v string) {
		tapped = true
	})

	if !tapped {
		t.Fatalf("Tap failed to execute side effect")
	}

	if res.Data() != "hello" {
		t.Fatalf("Tap failed to return original result")
	}
}

func TestTap_Failure(t *testing.T) {
	fail := FailureWithId[string](errtype.Validation, "fail")
	tappedFail := false
	resFail := Tap(fail, func(v string) {
		tappedFail = true
	})

	if tappedFail {
		t.Fatalf("Tap executed side effect on failure")
	}

	if resFail.IsSuccess() {
		t.Fatalf("Tap result should be failure")
	}
}

func TestMapSlice_Success(t *testing.T) {
	res := MapSlice(Success([]int{1, 2}), strconv.Itoa)
	if res.IsFailure() {
		t.Fatalf("expected success, got %v", res.Fault())
	}

	if res.Count() != 2 {
		t.Fatalf("expected 2 items, got %d", res.Count())
	}

	if res.Items[0] != "1" {
		t.Fatalf("unexpected item: %s", res.Items[0])
	}
}

func TestMapSlice_Failure(t *testing.T) {
	input := FailureWithId[[]int](errtype.Validation, "slice err")
	res := MapSlice(input, strconv.Itoa)
	if res.IsSuccess() {
		t.Fatalf("expected failure, got success")
	}

	if res.Fault().GetMessage() != "slice err" {
		t.Fatalf("unexpected message: %s", res.Fault().GetMessage())
	}
}

func TestFlatMapSlice_Success(t *testing.T) {
	input := Success([]int{1, 2})
	res := FlatMapSlice(input, func(v int) []int {
		return []int{v, v * 10}
	})

	if res.IsFailure() {
		t.Fatalf("expected success, got %v", res.Fault())
	}

	if res.Count() != 4 {
		t.Fatalf("expected 4 items, got %d", res.Count())
	}
}

func TestFlatMapSlice_Failure(t *testing.T) {
	input := FailureWithId[[]int](errtype.Execution, "flat err")
	res := FlatMapSlice(input, func(v int) []int {
		return []int{v}
	})

	if res.IsSuccess() {
		t.Fatalf("expected failure, got success")
	}

	if res.Fault().GetMessage() != "flat err" {
		t.Fatalf("unexpected message: %s", res.Fault().GetMessage())
	}
}

func TestMapMapValues_Success(t *testing.T) {
	input := OkMap(map[string]int{"a": 1, "b": 2})
	res := MapMapValues(input, strconv.Itoa)
	if res.IsFailure() {
		t.Fatalf("expected success, got %v", res.Fault())
	}

	val, ok := res.Get("b")
	if !ok {
		t.Fatalf("expected key 'b' to exist")
	}

	if val != "2" {
		t.Fatalf("unexpected value: %s", val)
	}
}

func TestMapMapValues_Failure(t *testing.T) {
	failWrap := FailureWithId[int](errtype.NotFound, "map missing")
	err := failWrap.Fault()
	res := MapMapValues(FailMap[string, int](err), strconv.Itoa)
	if res.IsSuccess() {
		t.Fatalf("expected failure, got success")
	}

	if res.Fault().GetMessage() != "map missing" {
		t.Fatalf("unexpected message: %s", res.Fault().GetMessage())
	}
}
