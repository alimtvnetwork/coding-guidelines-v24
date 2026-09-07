package appfault

import (
	"testing"

	"coding-guidelines/common/pkg/errtype"
)

var (
	_ IsSuccessChecker = Result[string]{}
	_ IsFailureChecker = Result[string]{}
	_ IsInvalidChecker = Result[string]{}
	_ IsNullChecker    = Result[string]{}
	_ IsEmptyChecker   = Result[string]{}
	_ IsDefinedChecker = Result[string]{}
	_ DefinableChecker = Result[string]{}
	_ StatusChecker    = Result[string]{}

	_ IsSuccessChecker = (*AppError)(nil)
	_ IsFailureChecker = (*AppError)(nil)
	_ IsInvalidChecker = (*AppError)(nil)
	_ IsNullChecker    = (*AppError)(nil)
	_ IsEmptyChecker   = (*AppError)(nil)
	_ IsDefinedChecker = (*AppError)(nil)
	_ DefinableChecker = (*AppError)(nil)
	_ StatusChecker    = (*AppError)(nil)
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
