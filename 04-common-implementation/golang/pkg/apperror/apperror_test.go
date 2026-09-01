package apperror_test

import (
	"errors"
	"testing"

	"coding-guidelines/common/pkg/apperror"
)

func TestAppErrorCreation(t *testing.T) {
	appErr := apperror.New(apperror.ErrValidation, "invalid input")
	if !appErr.IsCode(apperror.ErrValidation) {
		t.Fatalf("expected ErrValidation, got %s", appErr.Code())
	}

	if appErr.StackTrace().IsEmpty() {
		t.Fatal("expected non-empty stack trace")
	}
}

func TestAppErrorWrap(t *testing.T) {
	rawErr := errors.New("raw connection lost")
	appErr := apperror.Wrap(rawErr, apperror.ErrDatabaseQuery, "query failed").
		WithUrl("https://db.local:5432").
		WithStatusCode(500)
	if appErr.StatusCode() != 500 {
		t.Fatalf("expected status code 500, got %d", appErr.StatusCode())
	}
}

func TestResultSuccess(t *testing.T) {
	res := apperror.Ok("hello-world")
	if !res.IsSafe() || res.HasError() {
		t.Fatal("expected result to be safe without error")
	}

	if res.Value() != "hello-world" {
		t.Fatalf("expected 'hello-world', got %s", res.Value())
	}
}

func TestResultFailureAndWrap(t *testing.T) {
	rawErr := errors.New("file missing")
	res := apperror.FailWrap[string](rawErr, apperror.ErrFileNotFound, "read config failed")
	if !res.HasError() || res.IsSafe() {
		t.Fatal("expected result to have error and not be safe")
	}

	if res.AppError().Code() != apperror.ErrFileNotFound {
		t.Fatalf("expected ErrFileNotFound, got %s", res.AppError().Code())
	}
}

func TestResultSlice(t *testing.T) {
	sliceRes := apperror.OkSlice([]string{"alpha", "beta"})
	if sliceRes.Count() != 2 || !sliceRes.HasItems() {
		t.Fatalf("expected count 2, got %d", sliceRes.Count())
	}
}

func TestResultMap(t *testing.T) {
	mapRes := apperror.OkMap(map[string]int{"one": 1, "two": 2})
	if !mapRes.Has("one") {
		t.Fatal("expected key 'one' to exist")
	}

	getRes := mapRes.Get("one")
	if !getRes.IsSafe() || getRes.Value() != 1 {
		t.Fatalf("expected 1, got %v", getRes.ValueOr(0))
	}
}
