package apperror_test

import (
	"errors"
	"testing"

	"coding-guidelines/common/pkg/apperror"
	"coding-guidelines/common/pkg/result"
)

func TestFaultCreationAndDetails(t *testing.T) {
	fault := apperror.NewWithDetails("repo.find", "E2004", "not found", "repo", apperror.ErrorTypeNotFound, apperror.SeverityWarn, nil)
	if !fault.IsErrorCode("E2004") || !fault.HasValidError() {
		t.Fatalf("expected code E2004, got %s", fault.GetCode())
	}

	if fault.GetOp() != "repo.find" {
		t.Fatalf("expected op repo.find, got %s", fault.GetOp())
	}
}

func TestFaultWrapSimple(t *testing.T) {
	rawErr := errors.New("raw disk failure")
	fault := apperror.WrapSimple(rawErr, "fs.write").WithStatusCode(500)
	if fault.StatusCode() != 500 {
		t.Fatalf("expected status 500, got %d", fault.StatusCode())
	}

	if fault.Unwrap() != rawErr {
		t.Fatalf("expected unwrapped error to match rawErr")
	}
}

func TestResultMonadicSuccess(t *testing.T) {
	res := result.SuccessResult("data-payload")
	if !res.IsSuccess() || res.IsFailed() || !res.HasNoError() {
		t.Fatal("expected success status")
	}

	val, err := res.Unwrap()
	if err != nil || val != "data-payload" {
		t.Fatalf("expected data-payload, got %s", val)
	}
}

func TestResultMonadicFailure(t *testing.T) {
	res := result.NewFailureWithType[string]("E1001", "missing field", "config.load")
	if !res.IsFailed() || !res.IsFailure() || !res.HasValidError() {
		t.Fatal("expected failure status with valid error")
	}

	if res.UnwrapOr("fallback") != "fallback" {
		t.Fatal("expected fallback value")
	}
}

func TestResultSliceAndMap(t *testing.T) {
	sliceRes := apperror.OkSlice([]string{"alpha", "beta"})
	if sliceRes.Count() != 2 || !sliceRes.HasItems() {
		t.Fatalf("expected count 2, got %d", sliceRes.Count())
	}

	mapRes := apperror.OkMap(map[string]int{"key1": 100})
	if !mapRes.Has("key1") || mapRes.Count() != 1 {
		t.Fatal("expected key1 to exist")
	}
}
