package appfault_test

import (
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

func TestAppError_NilReceiverSafety(t *testing.T) {
	var nilErr *appfault.AppError

	// 1. Null, zero, and empty checks on nil receiver MUST NEVER PANIC
	if !nilErr.IsNull() {
		t.Fatal("expected IsNull() to be true on nil *AppError")
	}

	if !nilErr.IsEmpty() {
		t.Fatal("expected IsEmpty() to be true on nil *AppError")
	}

	if !nilErr.HasZero() {
		t.Fatal("expected HasZero() to be true on nil *AppError")
	}

	if !nilErr.IsZero() {
		t.Fatal("expected IsZero() to be true on nil *AppError")
	}

	if !nilErr.HasNull() {
		t.Fatal("expected HasNull() to be true on nil *AppError")
	}

	if !nilErr.HasNullError() {
		t.Fatal("expected HasNullError() to be true on nil *AppError")
	}

	if nilErr.HasError() {
		t.Fatal("expected HasError() to be false on nil *AppError")
	}

	if !nilErr.IsSuccess() {
		t.Fatal("expected IsSuccess() to be true on nil *AppError")
	}

	if nilErr.IsFailed() {
		t.Fatal("expected IsFailed() to be false on nil *AppError")
	}

	if !nilErr.IsValid() {
		t.Fatal("expected IsValid() to be true on nil *AppError")
	}

	if nilErr.IsInvalid() {
		t.Fatal("expected IsInvalid() to be false on nil *AppError")
	}

	// 2. Value getters on nil receiver MUST NEVER PANIC
	if nilErr.Code() != 0 {
		t.Fatalf("expected Code() == 0 on nil, got %d", nilErr.Code())
	}

	if nilErr.Type() != errtype.None {
		t.Fatalf("expected Type() == None on nil, got %v", nilErr.Type())
	}

	if nilErr.Message() != "" {
		t.Fatalf("expected Message() == '' on nil, got %q", nilErr.Message())
	}

	if nilErr.StatusCode() != 0 {
		t.Fatalf("expected StatusCode() == 0 on nil, got %d", nilErr.StatusCode())
	}

	if nilErr.StackTrace() != nil {
		t.Fatal("expected StackTrace() == nil on nil")
	}

	if nilErr.Context().Count() != 0 {
		t.Fatal("expected Context().Count() == 0 on nil")
	}

	if nilErr.LoopCount() != 1 {
		t.Fatalf("expected LoopCount() == 1 on nil, got %d", nilErr.LoopCount())
	}

	// 3. Formatting and printing on nil receiver MUST NEVER PANIC
	if nilErr.FormatStdout() != "" {
		t.Fatalf("expected FormatStdout() == '' on nil, got %q", nilErr.FormatStdout())
	}

	if nilErr.FormatJson() != "{}" {
		t.Fatalf("expected FormatJson() == '{}' on nil, got %q", nilErr.FormatJson())
	}

	if nilErr.FormatTextLog() != "" {
		t.Fatalf("expected FormatTextLog() == '' on nil, got %q", nilErr.FormatTextLog())
	}

	nilErr.Print()
	nilErr.PrintStdout()
	nilErr.PrintJson()
	nilErr.PrintLog()

	// 4. Clone and Concat on nil receiver MUST NEVER PANIC
	cloned := nilErr.Clone()
	if cloned != nil {
		t.Fatal("expected nilErr.Clone() == nil")
	}

	otherErr := appfault.New(errtype.Validation, "input invalid")
	concatenated := nilErr.Concat(otherErr)
	if concatenated != otherErr {
		t.Fatal("expected nilErr.Concat(otherErr) to return otherErr")
	}

	concatNil := otherErr.Concat(nilErr)
	if concatNil != otherErr {
		t.Fatal("expected otherErr.Concat(nilErr) to return otherErr")
	}
}

func TestResult_NullSafety(t *testing.T) {
	// Zero-value result
	var r appfault.Result[string]

	if !r.IsNull() {
		t.Fatal("expected r.IsNull() to be true on zero Result")
	}

	if !r.IsEmpty() {
		t.Fatal("expected r.IsEmpty() to be true on zero Result")
	}

	if !r.HasZero() {
		t.Fatal("expected r.HasZero() to be true on zero Result")
	}

	if !r.IsZero() {
		t.Fatal("expected r.IsZero() to be true on zero Result")
	}

	if !r.HasNull() {
		t.Fatal("expected r.HasNull() to be true on zero Result")
	}

	if !r.IsSuccess() {
		t.Fatal("expected r.IsSuccess() to be true on zero Result")
	}

	if r.IsFailed() {
		t.Fatal("expected r.IsFailed() to be false on zero Result")
	}

	// Clone on zero-value Result
	cloned := r.Clone()
	if !cloned.IsNull() || !cloned.IsSuccess() {
		t.Fatal("cloned zero Result should remain zero")
	}

	// Concat on results
	rFail := appfault.FailureResult[string](appfault.New(errtype.NotFound, "record not found"))
	merged := r.Concat(rFail)
	if !merged.IsFailed() {
		t.Fatal("expected merged result to be failed")
	}
}

func TestResult_NilPointerSafety(t *testing.T) {
	var nilRes *appfault.Result[string]

	// 1. Status checks on nil *Result[T] MUST NEVER PANIC
	if !nilRes.IsFailure() {
		t.Fatal("expected IsFailure() to be true on nil *Result")
	}

	if !nilRes.IsFailed() {
		t.Fatal("expected IsFailed() to be true on nil *Result")
	}

	if !nilRes.IsInvalid() {
		t.Fatal("expected IsInvalid() to be true on nil *Result")
	}

	if !nilRes.HasError() {
		t.Fatal("expected HasError() to be true on nil *Result")
	}

	if nilRes.IsSuccess() {
		t.Fatal("expected IsSuccess() to be false on nil *Result")
	}

	if nilRes.IsValid() {
		t.Fatal("expected IsValid() to be false on nil *Result")
	}

	if nilRes.HasNoError() {
		t.Fatal("expected HasNoError() to be false on nil *Result")
	}

	if nilRes.IsSafe() {
		t.Fatal("expected IsSafe() to be false on nil *Result")
	}

	// 2. Cardinality and definition predicates on nil *Result[T] MUST NEVER PANIC
	if nilRes.Count() != 0 {
		t.Fatalf("expected Count() == 0 on nil *Result, got %d", nilRes.Count())
	}

	if !nilRes.IsEmpty() {
		t.Fatal("expected IsEmpty() to be true on nil *Result")
	}

	if nilRes.HasRecord() {
		t.Fatal("expected HasRecord() to be false on nil *Result")
	}

	if nilRes.HasRecords() {
		t.Fatal("expected HasRecords() to be false on nil *Result")
	}

	if nilRes.IsDefined() {
		t.Fatal("expected IsDefined() to be false on nil *Result")
	}

	if !nilRes.IsCountOtherThan(1) {
		t.Fatal("expected IsCountOtherThan(1) to be true on nil *Result")
	}

	if !nilRes.IsCountOtherThan(0) {
		t.Fatal("expected IsCountOtherThan(0) to be true on nil *Result")
	}

	// 3. Payload and error access on nil *Result[T] MUST NEVER PANIC
	if nilRes.AppError() != nil {
		t.Fatal("expected AppError() == nil on nil *Result")
	}

	if nilRes.Fault() != nil {
		t.Fatal("expected Fault() == nil on nil *Result")
	}

	if nilRes.Error() != nil {
		t.Fatal("expected Error() == nil on nil *Result")
	}

	if nilRes.Data() != "" {
		t.Fatalf("expected Data() == '' on nil *Result, got %q", nilRes.Data())
	}

	if nilRes.Value() != "" {
		t.Fatalf("expected Value() == '' on nil *Result, got %q", nilRes.Value())
	}

	if nilRes.ValueAny() != nil {
		t.Fatal("expected ValueAny() == nil on nil *Result")
	}

	if nilRes.ToMap() != nil {
		t.Fatal("expected ToMap() == nil on nil *Result")
	}

	// 4. Safe operations, formatters, and printers
	nilRes.Print()
	nilRes.PrintFault()
	nilRes.PrintStdout()
	nilRes.PrintJson()
	nilRes.PrintLog()
	nilRes.HandleError()
}

func TestResultSlice_NilPointerSafety(t *testing.T) {
	var nilSlice *appfault.ResultSlice[string]

	// 1. Status checks on nil *ResultSlice[T] MUST NEVER PANIC
	if !nilSlice.IsFailure() {
		t.Fatal("expected IsFailure() to be true on nil *ResultSlice")
	}

	if !nilSlice.IsFailed() {
		t.Fatal("expected IsFailed() to be true on nil *ResultSlice")
	}

	if !nilSlice.IsInvalid() {
		t.Fatal("expected IsInvalid() to be true on nil *ResultSlice")
	}

	if !nilSlice.HasError() {
		t.Fatal("expected HasError() to be true on nil *ResultSlice")
	}

	if nilSlice.IsSuccess() {
		t.Fatal("expected IsSuccess() to be false on nil *ResultSlice")
	}

	// 2. Cardinality and definition predicates on nil *ResultSlice[T] MUST NEVER PANIC
	if nilSlice.Count() != 0 {
		t.Fatalf("expected Count() == 0 on nil *ResultSlice, got %d", nilSlice.Count())
	}

	if nilSlice.Length() != 0 {
		t.Fatalf("expected Length() == 0 on nil *ResultSlice, got %d", nilSlice.Length())
	}

	if !nilSlice.IsEmpty() {
		t.Fatal("expected IsEmpty() to be true on nil *ResultSlice")
	}

	if nilSlice.HasRecord() {
		t.Fatal("expected HasRecord() to be false on nil *ResultSlice")
	}

	if nilSlice.HasRecords() {
		t.Fatal("expected HasRecords() to be false on nil *ResultSlice")
	}

	if nilSlice.HasItems() {
		t.Fatal("expected HasItems() to be false on nil *ResultSlice")
	}

	if nilSlice.IsDefined() {
		t.Fatal("expected IsDefined() to be false on nil *ResultSlice")
	}

	if !nilSlice.IsCountOtherThan(1) {
		t.Fatal("expected IsCountOtherThan(1) to be true on nil *ResultSlice")
	}

	if !nilSlice.IsCountOtherThan(0) {
		t.Fatal("expected IsCountOtherThan(0) to be true on nil *ResultSlice")
	}

	// 3. Payload and error access on nil *ResultSlice[T] MUST NEVER PANIC
	if nilSlice.AppError() != nil {
		t.Fatal("expected AppError() == nil on nil *ResultSlice")
	}

	if nilSlice.Fault() != nil {
		t.Fatal("expected Fault() == nil on nil *ResultSlice")
	}

	if nilSlice.ValueAny() != nil {
		t.Fatal("expected ValueAny() == nil on nil *ResultSlice")
	}

	// 4. Safe combinators on nil *ResultSlice[T] MUST NEVER PANIC
	filtered := nilSlice.Filter(nil)
	if filtered.Count() != 0 {
		t.Fatal("expected filtered nilSlice to have count 0")
	}

	nilSlice.ForEach(nil)
	nilSlice.ForEachBreak(nil)
}

func TestResultMap_NilPointerSafety(t *testing.T) {
	var nilMap *appfault.ResultMap[string, int]

	// 1. Status checks on nil *ResultMap[K, V] MUST NEVER PANIC
	if !nilMap.IsFailure() {
		t.Fatal("expected IsFailure() to be true on nil *ResultMap")
	}

	if !nilMap.IsFailed() {
		t.Fatal("expected IsFailed() to be true on nil *ResultMap")
	}

	if !nilMap.IsInvalid() {
		t.Fatal("expected IsInvalid() to be true on nil *ResultMap")
	}

	if !nilMap.HasError() {
		t.Fatal("expected HasError() to be true on nil *ResultMap")
	}

	if nilMap.IsSuccess() {
		t.Fatal("expected IsSuccess() to be false on nil *ResultMap")
	}

	// 2. Cardinality and definition predicates on nil *ResultMap[K, V] MUST NEVER PANIC
	if nilMap.Count() != 0 {
		t.Fatalf("expected Count() == 0 on nil *ResultMap, got %d", nilMap.Count())
	}

	if !nilMap.IsEmpty() {
		t.Fatal("expected IsEmpty() to be true on nil *ResultMap")
	}

	if nilMap.HasRecord() {
		t.Fatal("expected HasRecord() to be false on nil *ResultMap")
	}

	if nilMap.HasRecords() {
		t.Fatal("expected HasRecords() to be false on nil *ResultMap")
	}

	if nilMap.IsDefined() {
		t.Fatal("expected IsDefined() to be false on nil *ResultMap")
	}

	if !nilMap.IsCountOtherThan(1) {
		t.Fatal("expected IsCountOtherThan(1) to be true on nil *ResultMap")
	}

	if !nilMap.IsCountOtherThan(0) {
		t.Fatal("expected IsCountOtherThan(0) to be true on nil *ResultMap")
	}

	// 3. Map lookups and getters on nil *ResultMap[K, V] MUST NEVER PANIC
	if nilMap.Has("key") {
		t.Fatal("expected Has('key') to be false on nil *ResultMap")
	}

	val, ok := nilMap.Get("key")
	if ok || val != 0 {
		t.Fatalf("expected Get('key') to return (0, false) on nil *ResultMap, got (%d, %v)", val, ok)
	}

	if len(nilMap.Keys()) != 0 {
		t.Fatal("expected Keys() to be empty on nil *ResultMap")
	}

	if len(nilMap.Values()) != 0 {
		t.Fatal("expected Values() to be empty on nil *ResultMap")
	}

	if nilMap.AppError() != nil {
		t.Fatal("expected AppError() == nil on nil *ResultMap")
	}

	if nilMap.Fault() != nil {
		t.Fatal("expected Fault() == nil on nil *ResultMap")
	}

	if nilMap.ValueAny() != nil {
		t.Fatal("expected ValueAny() == nil on nil *ResultMap")
	}

	// 4. Safe combinators on nil *ResultMap[K, V] MUST NEVER PANIC
	filtered := nilMap.Filter(nil)
	if filtered.Count() != 0 {
		t.Fatal("expected filtered nilMap to have count 0")
	}

	nilMap.ForEach(nil)
}
