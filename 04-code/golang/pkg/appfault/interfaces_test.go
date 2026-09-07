package appfault

import (
	"testing"

	"coding-guidelines/common/pkg/errtype"
)

var (
	_ SimpleVerifier = Result[string]{}
	_ SimpleVerifier = (*AppError)(nil)
	_ SimpleVerifier = ResultSlice[string]{}
	_ SimpleVerifier = ResultMap[string, int]{}

	_ SimpleVerifyChecker = Result[string]{}
	_ SimpleVerifyChecker = (*AppError)(nil)
	_ SimpleVerifyChecker = ResultSlice[string]{}
	_ SimpleVerifyChecker = ResultMap[string, int]{}

	_ SimpleVerifiable = Result[string]{}
	_ SimpleVerifiable = (*AppError)(nil)
	_ SimpleVerifiable = ResultSlice[string]{}
	_ SimpleVerifiable = ResultMap[string, int]{}

	_ SimpleVerifyCheckable = Result[string]{}
	_ SimpleVerifyCheckable = (*AppError)(nil)
	_ SimpleVerifyCheckable = ResultSlice[string]{}
	_ SimpleVerifyCheckable = ResultMap[string, int]{}
)

func TestResultSuccessCheckers(t *testing.T) {
	okRes := Result[string]{Value: "data"}
	if !okRes.IsSuccess() {
		t.Fatal("expected success")
	}

	if okRes.IsFailure() {
		t.Fatal("expected not failure")
	}

	if !okRes.IsDefined() {
		t.Fatal("expected defined")
	}

	if !okRes.IsNull() {
		t.Fatal("expected null")
	}
}

func TestResultFailureCheckers(t *testing.T) {
	failRes := FailureResult[string](New(errtype.Validation, "invalid input"))
	if failRes.IsSuccess() {
		t.Fatal("expected not success")
	}

	if !failRes.IsFailure() {
		t.Fatal("expected failure")
	}

	if !failRes.IsInvalid() {
		t.Fatal("expected invalid")
	}

	if failRes.IsDefined() {
		t.Fatal("expected not defined")
	}
}

func TestAppErrorCheckers(t *testing.T) {
	err := New(errtype.Database, "connection failed")
	if err.IsSuccess() {
		t.Fatal("expected not success")
	}

	if !err.IsFailure() {
		t.Fatal("expected failure")
	}

	if !err.IsInvalid() {
		t.Fatal("expected invalid")
	}

	if !err.IsDefined() {
		t.Fatal("expected defined")
	}
}

func TestNilAppErrorCheckers(t *testing.T) {
	var nilErr *AppError
	if !nilErr.IsSuccess() {
		t.Fatal("expected success")
	}

	if nilErr.IsFailure() {
		t.Fatal("expected not failure")
	}

	if nilErr.IsDefined() {
		t.Fatal("expected not defined")
	}

	if !nilErr.IsNull() {
		t.Fatal("expected null")
	}
}

func TestAsSimpleVerifier_Result(t *testing.T) {
	var v SimpleVerifier = Result[string]{Value: "data"}.AsSimpleVerifier()
	if !v.IsSuccess() {
		t.Fatal("expected Result AsSimpleVerifier to be success")
	}

	if !v.IsDefined() {
		t.Fatal("expected Result AsSimpleVerifier to be defined")
	}
}

func TestAsSimpleVerifier_Error(t *testing.T) {
	var v SimpleVerifier = New(errtype.Validation, "err").AsSimpleVerifier()
	if !v.IsFailure() {
		t.Fatal("expected AppError AsSimpleVerifier to be failure")
	}

	if !v.IsInvalid() {
		t.Fatal("expected AppError AsSimpleVerifier to be invalid")
	}
}

func TestAsSimpleVerifier_Slice(t *testing.T) {
	var v SimpleVerifier = OkSlice([]string{"a"}).AsSimpleVerifier()
	if !v.IsSuccess() {
		t.Fatal("expected ResultSlice AsSimpleVerifier to be success")
	}

	if v.IsEmpty() {
		t.Fatal("expected ResultSlice AsSimpleVerifier not to be empty")
	}
}

func TestAsSimpleVerifier_Map(t *testing.T) {
	var v SimpleVerifier = OkMap(map[string]int{"k": 1}).AsSimpleVerifier()
	if !v.IsSuccess() {
		t.Fatal("expected ResultMap AsSimpleVerifier to be success")
	}

	if v.IsEmpty() {
		t.Fatal("expected ResultMap AsSimpleVerifier not to be empty")
	}
}

func TestAsSimpleVerifyChecker_Result(t *testing.T) {
	var v SimpleVerifier = Result[string]{Value: "data"}.AsSimpleVerifyChecker()
	if !v.IsSuccess() {
		t.Fatal("expected Result AsSimpleVerifyChecker to be success")
	}

	if !v.IsDefined() {
		t.Fatal("expected Result AsSimpleVerifyChecker to be defined")
	}
}

func TestAsSimpleVerifyChecker_Error(t *testing.T) {
	var v SimpleVerifier = New(errtype.Validation, "err").AsSimpleVerifyChecker()
	if !v.IsFailure() {
		t.Fatal("expected AppError AsSimpleVerifyChecker to be failure")
	}

	if !v.IsInvalid() {
		t.Fatal("expected AppError AsSimpleVerifyChecker to be invalid")
	}
}

func TestAsSimpleVerifyChecker_Slice(t *testing.T) {
	var v SimpleVerifier = OkSlice([]string{"a"}).AsSimpleVerifyChecker()
	if !v.IsSuccess() {
		t.Fatal("expected ResultSlice AsSimpleVerifyChecker to be success")
	}

	if v.IsEmpty() {
		t.Fatal("expected ResultSlice AsSimpleVerifyChecker not to be empty")
	}
}

func TestAsSimpleVerifyChecker_Map(t *testing.T) {
	var v SimpleVerifier = OkMap(map[string]int{"k": 1}).AsSimpleVerifyChecker()
	if !v.IsSuccess() {
		t.Fatal("expected ResultMap AsSimpleVerifyChecker to be success")
	}

	if v.IsEmpty() {
		t.Fatal("expected ResultMap AsSimpleVerifyChecker not to be empty")
	}
}
