package appfault_test

import (
	"errors"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/result"
)

func TestAppErrorCreationAndDetails(t *testing.T) {
	appErr := appfault.NewWithDetails("repo.find", "E2004", "not found", "repo", appfault.ErrorTypeNotFound, appfault.SeverityWarn, nil)
	if !appErr.IsErrorCode("E2004") || !appErr.HasValidError() {
		t.Fatalf("expected code E2004, got %s", appErr.GetCode())
	}

	if appErr.GetOp() != "repo.find" {
		t.Fatalf("expected op repo.find, got %s", appErr.GetOp())
	}
}

func TestAppErrorWrapSimple(t *testing.T) {
	rawErr := errors.New("raw disk failure")
	appErr := appfault.WrapSimple(rawErr, "fs.write").WithStatusCode(500)
	if appErr.StatusCode() != 500 {
		t.Fatalf("expected status 500, got %d", appErr.StatusCode())
	}

	if appErr.Unwrap() != rawErr {
		t.Fatalf("expected unwrapped error to match rawErr")
	}
}

func TestResultMonadicSuccess(t *testing.T) {
	res := result.SuccessResult("data-payload")
	if !res.IsSuccess() || res.IsFailed() || res.Data() != "data-payload" {
		t.Fatal("expected valid success result and matching payload")
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
	sliceRes := appfault.OkSlice([]string{"alpha", "beta"})
	if sliceRes.Count() != 2 || !sliceRes.HasItems() {
		t.Fatalf("expected count 2, got %d", sliceRes.Count())
	}

	mapRes := appfault.OkMap(map[string]int{"key1": 100})
	if !mapRes.Has("key1") || mapRes.Count() != 1 {
		t.Fatal("expected key1 to exist")
	}
}
