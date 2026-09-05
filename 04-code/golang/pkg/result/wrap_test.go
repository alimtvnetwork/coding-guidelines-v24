package result_test

import (
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

func TestWrapSuccess(t *testing.T) {
	val := "hello-world"
	wrap := result.WrapSuccess(val)

	if wrap.IsFailed() {
		t.Fatalf("expected success, got failure")
	}

	if wrap.Data() != val {
		t.Fatalf("expected %s, got %s", val, wrap.Data())
	}
}

func TestWrapFailure(t *testing.T) {
	appErr := appfault.New(errtype.IO, "disk read error")
	wrap := result.WrapFailure[string](appErr)

	if wrap.IsSuccess() {
		t.Fatalf("expected failure, got success")
	}

	if wrap.Fault() == nil {
		t.Fatalf("expected non-nil fault")
	}

	if wrap.Fault().Message() != "disk read error" {
		t.Fatalf("unexpected message: %s", wrap.Fault().Message())
	}
}

func TestWrapFormat(t *testing.T) {
	appErr := appfault.New(errtype.Validation, "invalid configuration")
	wrap := result.WrapFailure[int](appErr)

	// Default formatting check
	formatted := wrap.Format(nil)
	if len(formatted) == 0 {
		t.Fatalf("expected non-empty formatted string")
	}

	// Custom formatting check
	custom := wrap.Format(func(r result.Wrap[int]) string {
		return "CUSTOM-ERROR: " + r.Fault().Message()
	})

	if custom != "CUSTOM-ERROR: invalid configuration" {
		t.Fatalf("unexpected custom format: %s", custom)
	}

	// Success formatting check
	successWrap := result.WrapSuccess(42)
	successFormatted := successWrap.Format(nil)
	if successFormatted != "✅ [OK] 42" {
		t.Fatalf("unexpected success format: %s", successFormatted)
	}
}

func TestWrapFailureWithId(t *testing.T) {
	w := result.WrapFailureWithId[string](errtype.Validation, "bad param")
	if w.IsSuccess() {
		t.Fatalf("expected failure")
	}

	if w.Fault().Type() != errtype.Validation {
		t.Fatalf("unexpected type: %v", w.Fault().Type())
	}

	if w.Fault().Message() != "bad param" {
		t.Fatalf("unexpected message: %s", w.Fault().Message())
	}
}

func TestWrapFailureFromWrap(t *testing.T) {
	first := result.WrapFailureWithId[int](errtype.IO, "io failed")
	second := result.WrapFailureFromWrap[string](first)
	if second.IsSuccess() {
		t.Fatalf("expected failure in propagated wrap")
	}

	if second.Fault().Message() != "io failed" {
		t.Fatalf("unexpected message: %s", second.Fault().Message())
	}
}
