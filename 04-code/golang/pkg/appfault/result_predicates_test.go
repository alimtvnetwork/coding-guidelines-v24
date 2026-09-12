package appfault

import (
	"testing"

	"coding-guidelines/common/pkg/errtype"
)

func TestResultSlice_PredicatesAndCount(t *testing.T) {
	items := []string{"log1", "log2"}
	okRes := OkSlice(items)

	if !okRes.IsSuccess() {
		t.Fatal("expected okRes to be success")
	}

	if !okRes.HasRecord() {
		t.Fatal("expected okRes to have record")
	}

	if !okRes.HasRecords() {
		t.Fatal("expected okRes to have records")
	}

	if !okRes.IsDefined() {
		t.Fatal("expected okRes to be defined")
	}

	if okRes.IsEmpty() {
		t.Fatal("expected okRes not to be empty")
	}

	if okRes.IsCountOtherThan(2) {
		t.Fatal("expected IsCountOtherThan(2) to be false on slice with 2 items")
	}

	if !okRes.IsCountOtherThan(1) {
		t.Fatal("expected IsCountOtherThan(1) to be true on slice with 2 items")
	}
}

func TestResultSlice_Empty(t *testing.T) {
	emptyRes := OkSlice([]string{})

	if !emptyRes.IsEmpty() {
		t.Fatal("expected emptyRes to be empty")
	}

	if emptyRes.HasRecord() {
		t.Fatal("expected emptyRes to have no record")
	}

	if emptyRes.IsDefined() {
		t.Fatal("expected emptyRes not to be defined")
	}

	if emptyRes.IsCountOtherThan(0) {
		t.Fatal("expected IsCountOtherThan(0) to be false on empty slice")
	}

	if !emptyRes.IsCountOtherThan(1) {
		t.Fatal("expected IsCountOtherThan(1) to be true on empty slice")
	}
}

func TestResultSlice_Failure(t *testing.T) {
	failRes := FailSlice[string](New(errtype.Database, "db failure"))

	if !failRes.IsFailure() {
		t.Fatal("expected failRes to be failure")
	}

	if failRes.HasRecord() {
		t.Fatal("expected failRes to have no record")
	}

	if failRes.IsDefined() {
		t.Fatal("expected failRes not to be defined")
	}

	if !failRes.IsCountOtherThan(1) {
		t.Fatal("expected IsCountOtherThan(1) to be true on failure")
	}

	if !failRes.IsCountOtherThan(0) {
		t.Fatal("expected IsCountOtherThan(0) to be true on failure")
	}
}

func TestResultMap_PredicatesAndCount(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2}
	okRes := OkMap(data)

	if !okRes.HasRecord() {
		t.Fatal("expected okRes to have record")
	}

	if !okRes.IsDefined() {
		t.Fatal("expected okRes to be defined")
	}

	if okRes.IsEmpty() {
		t.Fatal("expected okRes not to be empty")
	}

	if okRes.IsCountOtherThan(2) {
		t.Fatal("expected IsCountOtherThan(2) to be false on map with 2 entries")
	}

	if !okRes.IsCountOtherThan(1) {
		t.Fatal("expected IsCountOtherThan(1) to be true on map with 2 entries")
	}
}

func TestResultMap_EmptyAndFailure(t *testing.T) {
	emptyRes := OkMap(map[string]int{})
	if !emptyRes.IsEmpty() {
		t.Fatal("expected emptyRes to be empty")
	}

	if emptyRes.HasRecord() {
		t.Fatal("expected emptyRes to have no record")
	}

	if emptyRes.IsDefined() {
		t.Fatal("expected emptyRes not to be defined")
	}

	failRes := FailMap[string, int](New(errtype.Database, "query failed"))
	if !failRes.IsCountOtherThan(1) {
		t.Fatal("expected IsCountOtherThan(1) to be true on map failure")
	}
}

func TestResult_PredicatesAndCount(t *testing.T) {
	scalarRes := SuccessResult("content")

	if !scalarRes.HasRecord() {
		t.Fatal("expected scalarRes to have record")
	}

	if !scalarRes.IsDefined() {
		t.Fatal("expected scalarRes to be defined")
	}

	if scalarRes.IsEmpty() {
		t.Fatal("expected scalarRes not to be empty")
	}

	if scalarRes.IsCountOtherThan(1) {
		t.Fatal("expected IsCountOtherThan(1) to be false on scalar with value")
	}

	emptyStringRes := SuccessResult("")
	if !emptyStringRes.IsEmpty() {
		t.Fatal("expected emptyStringRes to be empty")
	}

	if emptyStringRes.HasRecord() {
		t.Fatal("expected emptyStringRes to have no record")
	}

	if emptyStringRes.IsDefined() {
		t.Fatal("expected emptyStringRes not to be defined")
	}
}

func TestResult_RecordCountChecker(t *testing.T) {
	var checker RecordCountChecker = OkSlice([]string{"one", "two"})

	if checker.Count() != 2 {
		t.Fatalf("expected count 2, got %d", checker.Count())
	}

	if !checker.HasRecord() {
		t.Fatal("expected checker to have record")
	}

	if checker.IsCountOtherThan(2) {
		t.Fatal("expected IsCountOtherThan(2) to be false")
	}
}
