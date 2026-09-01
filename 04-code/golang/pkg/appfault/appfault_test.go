package appfault_test

import (
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

func TestContextMapOperations(t *testing.T) {
	cm := appfault.NewContextMap().Set("siteId", 101).Set("slug", "test-plugin")
	if !cm.Has("siteId") || cm.GetString("siteId") != "101" || cm.Count() != 2 {
		t.Fatalf("expected siteId=101 and count 2, got %s (count=%d)", cm.GetString("siteId"), cm.Count())
	}

	cm.Remove("slug")
	if cm.Has("slug") {
		t.Fatal("expected slug to be removed")
	}
}

func TestStackFrameAndCaller(t *testing.T) {
	frame := appfault.NewStackFrame("main.run", "main.go", 42)
	if frame.Function != "main.run" || frame.File != "main.go" || frame.Line != 42 {
		t.Fatalf("unexpected frame: %+v", frame)
	}

	caller := appfault.CaptureCaller(0)
	if len(caller) == 0 {
		t.Fatal("expected non-empty caller")
	}
}

func TestResultMonadicOperations(t *testing.T) {
	res := result.SuccessResult("data-payload")
	if !res.IsSuccess() || res.IsFailed() || res.Data() != "data-payload" {
		t.Fatal("expected success result")
	}

	failRes := result.NewFailureWithType[string]("E1001", "bad input", "validator")
	if !failRes.IsFailed() || failRes.UnwrapOr("fallback") != "fallback" {
		t.Fatal("expected failed result with fallback")
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
