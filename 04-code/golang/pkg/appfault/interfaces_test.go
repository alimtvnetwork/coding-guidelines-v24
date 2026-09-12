package appfault

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"

	"coding-guidelines/common/pkg/errtype"
)

var (
	_ SimpleVerifier = (*Result[string])(nil)
	_ SimpleVerifier = (*AppError)(nil)
	_ SimpleVerifier = (*ResultSlice[string])(nil)
	_ SimpleVerifier = (*ResultMap[string, int])(nil)

	_ SimpleVerifyChecker = (*Result[string])(nil)
	_ SimpleVerifyChecker = (*AppError)(nil)
	_ SimpleVerifyChecker = (*ResultSlice[string])(nil)
	_ SimpleVerifyChecker = (*ResultMap[string, int])(nil)

	_ SimpleVerifiable = (*Result[string])(nil)
	_ SimpleVerifiable = (*AppError)(nil)
	_ SimpleVerifiable = (*ResultSlice[string])(nil)
	_ SimpleVerifiable = (*ResultMap[string, int])(nil)

	_ SimpleVerifyCheckable = (*Result[string])(nil)
	_ SimpleVerifyCheckable = (*AppError)(nil)
	_ SimpleVerifyCheckable = (*ResultSlice[string])(nil)
	_ SimpleVerifyCheckable = (*ResultMap[string, int])(nil)

	_ ResultInspector = (*Result[string])(nil)
	_ ResultInspector = (*ResultSlice[string])(nil)
	_ ResultInspector = (*ResultMap[string, any])(nil)

	_ ResultUnwrapper = (*Result[string])(nil)
	_ ResultUnwrapper = (*ResultSlice[string])(nil)
	_ ResultUnwrapper = (*ResultMap[string, any])(nil)

	_ ResultCarrier = (*Result[string])(nil)
	_ ResultCarrier = (*ResultSlice[string])(nil)
	_ ResultCarrier = (*ResultMap[string, any])(nil)
)

func TestResultSuccessCheckers(t *testing.T) {
	okRes := SuccessResult("data")
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
	r := SuccessResult("data")
	var v SimpleVerifier = r.AsSimpleVerifier()
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
	s := OkSlice([]string{"a"})
	var v SimpleVerifier = s.AsSimpleVerifier()
	if !v.IsSuccess() {
		t.Fatal("expected ResultSlice AsSimpleVerifier to be success")
	}

	if v.IsEmpty() {
		t.Fatal("expected ResultSlice AsSimpleVerifier not to be empty")
	}
}

func TestAsSimpleVerifier_Map(t *testing.T) {
	m := OkMap(map[string]int{"k": 1})
	var v SimpleVerifier = m.AsSimpleVerifier()
	if !v.IsSuccess() {
		t.Fatal("expected ResultMap AsSimpleVerifier to be success")
	}

	if v.IsEmpty() {
		t.Fatal("expected ResultMap AsSimpleVerifier not to be empty")
	}
}

func TestAsSimpleVerifyChecker_Result(t *testing.T) {
	r := SuccessResult("data")
	var v SimpleVerifier = r.AsSimpleVerifyChecker()
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
	s := OkSlice([]string{"a"})
	var v SimpleVerifier = s.AsSimpleVerifyChecker()
	if !v.IsSuccess() {
		t.Fatal("expected ResultSlice AsSimpleVerifyChecker to be success")
	}

	if v.IsEmpty() {
		t.Fatal("expected ResultSlice AsSimpleVerifyChecker not to be empty")
	}
}

func TestAsSimpleVerifyChecker_Map(t *testing.T) {
	m := OkMap(map[string]int{"k": 1})
	var v SimpleVerifier = m.AsSimpleVerifyChecker()
	if !v.IsSuccess() {
		t.Fatal("expected ResultMap AsSimpleVerifyChecker to be success")
	}

	if v.IsEmpty() {
		t.Fatal("expected ResultMap AsSimpleVerifyChecker not to be empty")
	}
}

func assertInspectorSuccess(t *testing.T, inspector ResultInspector) {
	if inspector.IsFailed() {
		t.Fatal("expected success, got isFailed true")
	}

	if inspector.AppError() != nil {
		t.Fatal("expected nil AppError on success")
	}
}

func assertInspectorFailure(t *testing.T, inspector ResultInspector) {
	if !inspector.IsFailed() {
		t.Fatal("expected failure, got isFailed false")
	}

	if inspector.AppError() == nil {
		t.Fatal("expected non-nil AppError on failure")
	}
}

func testInspectorResult(t *testing.T) {
	okRes := SuccessResult("test-result")
	assertInspectorSuccess(t, &okRes)
	if okRes.ValueAny() != "test-result" {
		t.Fatalf("unexpected ValueAny: %v", okRes.ValueAny())
	}

	failRes := FailureResult[string](New(errtype.Validation, "bad result"))
	assertInspectorFailure(t, &failRes)
}

func testInspectorSlice(t *testing.T) {
	okSlice := OkSlice([]string{"alpha", "beta"})
	assertInspectorSuccess(t, &okSlice)
	if okSlice.ValueAny() == nil {
		t.Fatal("expected non-nil ValueAny")
	}

	failSlice := FailSlice[string](New(errtype.Validation, "bad slice"))
	assertInspectorFailure(t, &failSlice)
	if failSlice.Fault() == nil {
		t.Fatal("expected non-nil Fault")
	}
}

func testInspectorMap(t *testing.T) {
	okMap := OkMap(map[string]any{"k": "v"})
	assertInspectorSuccess(t, &okMap)
	if okMap.ValueAny() == nil {
		t.Fatal("expected non-nil ValueAny")
	}

	failMap := FailMap[string, any](New(errtype.Validation, "bad map"))
	assertInspectorFailure(t, &failMap)
	if failMap.Fault() == nil {
		t.Fatal("expected non-nil Fault")
	}
}

func TestResultInspectorConformance(t *testing.T) {
	testInspectorResult(t)
	testInspectorSlice(t)
	testInspectorMap(t)
}

func TestResultSlice_JsonRoundtrip(t *testing.T) {
	orig := OkSlice([]string{"x", "y"})
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var res ResultSlice[string]
	if err := json.Unmarshal(data, &res); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if res.Count() != 2 {
		t.Fatalf("expected count 2, got %d", res.Count())
	}
}

func TestResultSlice_YamlRoundtrip(t *testing.T) {
	orig := OkSlice([]string{"x", "y"})
	data, err := yaml.Marshal(orig)
	if err != nil {
		t.Fatalf("yaml marshal failed: %v", err)
	}

	var res ResultSlice[string]
	if err := yaml.Unmarshal(data, &res); err != nil {
		t.Fatalf("yaml unmarshal failed: %v", err)
	}

	if res.Count() != 2 {
		t.Fatalf("expected count 2, got %d", res.Count())
	}
}

func TestResultSlice_FailureSerializationRoundtrip(t *testing.T) {
	orig := FailSlice[string](New(errtype.Validation, "slice fail"))
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var res ResultSlice[string]
	if err := json.Unmarshal(data, &res); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if !res.IsFailed() {
		t.Fatal("expected failure on unmarshaled failed slice")
	}
}

func TestResultMap_JsonRoundtrip(t *testing.T) {
	orig := OkMap(map[string]int{"one": 1, "two": 2})
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var res ResultMap[string, int]
	if err := json.Unmarshal(data, &res); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if res.Count() != 2 {
		t.Fatalf("expected count 2, got %d", res.Count())
	}
}

func TestResultMap_YamlRoundtrip(t *testing.T) {
	orig := OkMap(map[string]int{"one": 1, "two": 2})
	data, err := yaml.Marshal(orig)
	if err != nil {
		t.Fatalf("yaml marshal failed: %v", err)
	}

	var res ResultMap[string, int]
	if err := yaml.Unmarshal(data, &res); err != nil {
		t.Fatalf("yaml unmarshal failed: %v", err)
	}

	if res.Count() != 2 {
		t.Fatalf("expected count 2, got %d", res.Count())
	}
}

func TestResultMap_FailureSerializationRoundtrip(t *testing.T) {
	orig := FailMap[string, int](New(errtype.NotFound, "map fail"))
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var res ResultMap[string, int]
	if err := json.Unmarshal(data, &res); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if !res.IsFailed() {
		t.Fatal("expected failure on unmarshaled failed map")
	}
}
