package appfault_test

import (
	"bytes"
	"strings"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

type castTestPayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestCastTo_Success(t *testing.T) {
	val, err := appfault.CastTo[string]("hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if val != "hello" {
		t.Fatalf("expected hello, got %s", val)
	}
}

func TestCastTo_Failure(t *testing.T) {
	_, err := appfault.CastTo[int]("not an int")
	if err == nil {
		t.Fatal("expected type mismatch error")
	}

	if err.Type() != errtype.TypeMismatch {
		t.Fatalf("expected TypeMismatch variation, got %v", err.Type())
	}
}

func TestReflectTo_Success(t *testing.T) {
	jsonBytes := []byte(`{"name":"Alice","age":30}`)
	val, err := appfault.ReflectTo[castTestPayload](jsonBytes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if val.Name != "Alice" {
		t.Fatalf("expected Alice, got %s", val.Name)
	}
}

func TestReflectTo_Failure(t *testing.T) {
	_, err := appfault.ReflectTo[int]("not a number")
	if err == nil {
		t.Fatal("expected reflection error")
	}

	if err.Type() != errtype.TypeMismatch {
		t.Fatalf("expected TypeMismatch variation, got %v", err.Type())
	}
}

func TestCastResult_Success(t *testing.T) {
	r := appfault.SuccessResult([]byte(`{"name":"Bob","age":25}`))
	res := appfault.CastResult[[]byte, castTestPayload](r)
	if res.IsFailed() {
		t.Fatalf("expected success, got fault: %v", res.Fault())
	}

	if res.Data().Age != 25 {
		t.Fatalf("expected age 25, got %d", res.Data().Age)
	}
}

func TestCastResult_PropagatesExistingFault(t *testing.T) {
	appErr := appfault.New(errtype.Validation, "input invalid")
	r := appfault.FailureResult[string](appErr)
	res := appfault.CastResult[string, int](r)
	if !res.IsFailed() {
		t.Fatal("expected failed result")
	}

	if res.Fault().Type() != errtype.Validation {
		t.Fatalf("expected Validation, got %v", res.Fault().Type())
	}
}

func TestCastResult_CastFailure(t *testing.T) {
	r := appfault.SuccessResult("invalid-payload")
	res := appfault.CastResult[string, int](r)
	if !res.IsFailed() {
		t.Fatal("expected failed result from cast error")
	}

	if res.Fault().Type() != errtype.TypeMismatch {
		t.Fatalf("expected TypeMismatch, got %v", res.Fault().Type())
	}
}

func TestCastContextPayload_Success(t *testing.T) {
	e := appfault.New(errtype.Generic, "test").
		WithContext("tag", "alpha").
		WithContext("count", 42)
	tag, err := appfault.CastContextPayload[string](e, "tag")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if tag != "alpha" {
		t.Fatalf("expected alpha, got %s", tag)
	}

	count, err2 := appfault.CastContextPayload[int](e, "count")
	if err2 != nil {
		t.Fatalf("unexpected error: %v", err2)
	}

	if count != 42 {
		t.Fatalf("expected 42, got %d", count)
	}
}

func TestCastContextPayload_MissingKey(t *testing.T) {
	e := appfault.New(errtype.Generic, "test")
	_, err := appfault.CastContextPayload[string](e, "missing")
	if err == nil {
		t.Fatal("expected error for missing key")
	}

	if err.Type() != errtype.NotFound {
		t.Fatalf("expected NotFound, got %v", err.Type())
	}
}

func TestCastContextPayload_TypeMismatch(t *testing.T) {
	e := appfault.New(errtype.Generic, "test").WithContext("key", "not-an-int")
	_, err := appfault.CastContextPayload[int](e, "key")
	if err == nil {
		t.Fatal("expected type mismatch error")
	}

	if err.Type() != errtype.TypeMismatch {
		t.Fatalf("expected TypeMismatch, got %v", err.Type())
	}
}

func TestCastContextPayload_NilError(t *testing.T) {
	_, err := appfault.CastContextPayload[string](nil, "key")
	if err == nil {
		t.Fatal("expected error for nil AppError")
	}

	if err.Type() != errtype.NotFound {
		t.Fatalf("expected NotFound, got %v", err.Type())
	}
}

func TestResultToBytes_Success(t *testing.T) {
	r := appfault.SuccessResult("sample text")
	res := appfault.ResultToBytes(r)
	if res.IsFailed() {
		t.Fatalf("unexpected failure: %v", res.Fault())
	}

	if !bytes.Equal(res.Data(), []byte("sample text")) {
		t.Fatalf("expected 'sample text', got %s", string(res.Data()))
	}
}

func TestResultToBytes_Fault(t *testing.T) {
	appErr := appfault.New(errtype.Timeout, "operation timed out")
	r := appfault.FailureResult[string](appErr)
	res := appfault.ResultToBytes(r)
	if res.IsFailed() {
		t.Fatalf("unexpected failure: %v", res.Fault())
	}

	if !strings.Contains(string(res.Data()), "operation timed out") {
		t.Fatalf("expected error message in bytes, got %s", string(res.Data()))
	}
}

func TestResultToJson_Success(t *testing.T) {
	payload := castTestPayload{Name: "David", Age: 40}
	r := appfault.SuccessResult(payload)
	res := appfault.ResultToJson(r)
	if res.IsFailed() {
		t.Fatalf("unexpected failure: %v", res.Fault())
	}

	if !strings.Contains(string(res.Data()), `"name": "David"`) {
		t.Fatalf("expected David in JSON, got %s", string(res.Data()))
	}

	resAlias := appfault.ResultToJSON(r)
	if !strings.Contains(string(resAlias.Data()), `"name": "David"`) {
		t.Fatalf("expected David in alias ResultToJSON, got %s", string(resAlias.Data()))
	}
}

func TestResultToJson_Fault(t *testing.T) {
	appErr := appfault.New(errtype.Forbidden, "permission denied")
	r := appfault.FailureResult[castTestPayload](appErr)
	res := appfault.ResultToJson(r)
	if res.IsFailed() {
		t.Fatalf("unexpected failure: %v", res.Fault())
	}

	if !strings.Contains(string(res.Data()), "permission denied") {
		t.Fatalf("expected permission denied in JSON, got %s", string(res.Data()))
	}
}
