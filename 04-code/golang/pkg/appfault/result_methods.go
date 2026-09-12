package appfault

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func valueRecordCount(val any) int {
	if val == nil {
		return 0
	}

	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len()
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Pointer, reflect.UnsafePointer:
		if v.IsNil() {
			return 0
		}

		return 1
	case reflect.String:
		if v.Len() == 0 {
			return 0
		}

		return 1
	default:
		return 1
	}
}

// IsSuccess returns true if no error is present.
func (r Result[T]) IsSuccess() bool {
	return r.appError == nil
}

// IsFailed returns true if an error is present.
func (r Result[T]) IsFailed() bool {
	return r.appError != nil
}

// IsFailure is an alias for IsFailed.
func (r Result[T]) IsFailure() bool {
	return r.IsFailed()
}

// IsInvalid is an alias for IsFailed.
func (r Result[T]) IsInvalid() bool {
	return r.IsFailed()
}

// IsValid returns true if the operation succeeded with no error.
func (r Result[T]) IsValid() bool {
	return r.IsSuccess()
}

// IsDefined returns true if the operation succeeded (no error) and contains more than 0 records or non-null data T.
func (r Result[T]) IsDefined() bool {
	if r.IsFailed() {
		return false
	}

	return r.Count() > 0
}

// HasRecord returns true if the operation succeeded and contains more than 0 records.
func (r Result[T]) HasRecord() bool {
	if r.IsFailed() {
		return false
	}

	return r.Count() > 0
}

// HasRecords is an alias for HasRecord.
func (r Result[T]) HasRecords() bool {
	return r.HasRecord()
}

// Count returns the number of records in the payload (0 if failed or empty, or length if collection).
func (r Result[T]) Count() int {
	if r.IsFailed() {
		return 0
	}

	return valueRecordCount(r.value)
}

// IsCountOtherThan returns true if the operation failed or its record count differs from n.
func (r Result[T]) IsCountOtherThan(n int) bool {
	if r.IsFailed() {
		return true
	}

	return r.Count() != n
}

// AsSimpleVerifier returns the Result conforming to SimpleVerifier.
func (r Result[T]) AsSimpleVerifier() SimpleVerifier {
	return r
}

// AsSimpleVerifyChecker returns the Result conforming to SimpleVerifier.
func (r Result[T]) AsSimpleVerifyChecker() SimpleVerifier {
	return r
}

// HasError returns true if an error is present.
func (r Result[T]) HasError() bool {
	return r.IsFailed()
}

// HasNoError returns true if no error exists.
func (r Result[T]) HasNoError() bool {
	return r.IsSuccess()
}

// HasValidError returns true if the embedded error has a valid code.
func (r Result[T]) HasValidError() bool {
	if r.appError == nil {
		return false
	}

	return r.appError.HasValidError()
}

// IsSafe returns true if the operation succeeded with no error.
func (r Result[T]) IsSafe() bool {
	return r.IsSuccess()
}

// IsEmpty returns true if no active error is present and payload has 0 records, or error is empty.
func (r Result[T]) IsEmpty() bool {
	if r.appError == nil {
		return r.Count() == 0
	}

	return r.appError.IsEmpty()
}

// HasZero returns true if error is nil or represents a zero-value/None error state.
func (r Result[T]) HasZero() bool {
	return r.IsEmpty()
}

// IsZero returns true if error is nil or represents a zero-value/None error state.
func (r Result[T]) IsZero() bool {
	return r.IsEmpty()
}

// IsNull returns true if the embedded AppError pointer is nil.
func (r Result[T]) IsNull() bool {
	return r.appError == nil
}

// HasNull returns true if the embedded AppError is nil or represents no error.
func (r Result[T]) HasNull() bool {
	return r.IsEmpty()
}

// Clone returns a deep copy of Result with its AppError safely cloned.
func (r Result[T]) Clone() Result[T] {
	return Result[T]{
		value:    r.value,
		appError: r.appError.Clone(),
	}
}

// Concat combines errors from two results into an immutable Result.
func (r Result[T]) Concat(other Result[T]) Result[T] {
	mergedErr := Merge(r.appError, other.appError)
	val := other.value
	if r.IsSuccess() && other.IsFailed() {
		val = r.value
	}

	return Result[T]{
		value:    val,
		appError: mergedErr,
	}
}

// Unwrap unpacks the (Value, *AppError) tuple.
func (r Result[T]) Unwrap() (T, *AppError) {
	return r.value, r.appError
}

// UnwrapOr returns the inner value if successful, or defaultVal on failure.
func (r Result[T]) UnwrapOr(defaultVal T) T {
	if r.IsFailed() {
		return defaultVal
	}

	return r.value
}

// DefaultResultFormatter formats Result[T]: error banner if failed, or data value if success.
func DefaultResultFormatter[T any](r Result[T]) string {
	if r.IsFailed() {
		return r.appError.Format(DefaultFaultFormatter)
	}

	return fmt.Sprintf("✅ [OK] %v", r.value)
}

// Print outputs the result representation to standard output.
func (r Result[T]) Print() {
	fmt.Println(r.Format(nil))
}

// PrintFault outputs the fault representation to standard output if failed.
func (r Result[T]) PrintFault() {
	if r.IsFailed() {
		r.appError.Print()
	}
}

// Format formats the Result using a custom or default formatter.
func (r Result[T]) Format(formatter ResultFormatter[T]) string {
	if formatter != nil {
		return formatter(r)
	}

	return DefaultResultFormatter(r)
}

// PrintWith outputs the result using a custom formatter.
func (r Result[T]) PrintWith(formatter ResultFormatter[T]) {
	fmt.Println(r.Format(formatter))
}

// FormatStdout formats the Result: rich error banner if failed, or success message if ok.
func (r Result[T]) FormatStdout() string {
	if r.IsFailed() {
		return r.appError.FormatStdout()
	}

	return fmt.Sprintf("✅ SUCCESS: %v", r.value)
}

// FormatJson formats the Result: JSON error if failed, or marshaled value JSON.
func (r Result[T]) FormatJson() string {
	if r.IsFailed() {
		return r.appError.FormatJson()
	}

	return fmt.Sprintf(`{"success":true,"data":%v}`, r.value)
}

// FormatJSON is an alias for FormatJson.
func (r Result[T]) FormatJSON() string {
	return r.FormatJson()
}

// FormatTextLog formats the Result: structured log error if failed, or log info if ok.
func (r Result[T]) FormatTextLog() string {
	if r.IsFailed() {
		return r.appError.FormatTextLog()
	}

	return fmt.Sprintf("[INFO] status=200 msg=%q", fmt.Sprintf("%v", r.value))
}

// PrintStdout prints the result formatted for stdout.
func (r Result[T]) PrintStdout() {
	fmt.Println(r.FormatStdout())
}

// PrintJson prints the result formatted as JSON.
func (r Result[T]) PrintJson() {
	fmt.Println(r.FormatJson())
}

// PrintJSON is an alias for PrintJson.
func (r Result[T]) PrintJSON() {
	r.PrintJson()
}

// PrintLog prints the result formatted as a structured log line.
func (r Result[T]) PrintLog() {
	fmt.Println(r.FormatTextLog())
}

// FormatStruct formats struct or primitive payload or error banner.
func (r Result[T]) FormatStruct() string {
	if r.IsFailed() {
		return r.appError.FormatStdout()
	}

	return fmt.Sprintf("%+v", r.value)
}

func payloadToMap(val any) map[string]any {
	bytes, err := json.Marshal(val)
	if err != nil {
		return nil
	}

	var result map[string]any
	_ = json.Unmarshal(bytes, &result)

	return result
}

// ToMap converts payload into map[string]any via JSON serialization.
func (r Result[T]) ToMap() map[string]any {
	if r.IsFailed() {
		return nil
	}

	return payloadToMap(r.value)
}
