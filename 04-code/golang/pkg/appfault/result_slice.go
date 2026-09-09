package appfault

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// ResultSlice wraps a generic slice collection with monadic error state.
type ResultSlice[T any] struct {
	Items    []T
	appError *AppError
}

type resultSliceDTO[T any] struct {
	Items    []T       `json:"items,omitempty" yaml:"items,omitempty"`
	AppError *AppError `json:"appError,omitempty" yaml:"appError,omitempty"`
}

// MarshalJSON provides JSON serialization for ResultSlice[T].
func (rs ResultSlice[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(resultSliceDTO[T]{
		Items:    rs.Items,
		AppError: rs.appError,
	})
}

// UnmarshalJSON provides JSON deserialization for ResultSlice[T].
func (rs *ResultSlice[T]) UnmarshalJSON(data []byte) error {
	var dto resultSliceDTO[T]
	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	rs.Items = dto.Items
	rs.appError = dto.AppError

	return nil
}

// MarshalYAML provides YAML serialization for ResultSlice[T].
func (rs ResultSlice[T]) MarshalYAML() (any, error) {
	return resultSliceDTO[T]{
		Items:    rs.Items,
		AppError: rs.appError,
	}, nil
}

// UnmarshalYAML provides YAML deserialization for ResultSlice[T].
func (rs *ResultSlice[T]) UnmarshalYAML(value *yaml.Node) error {
	var dto resultSliceDTO[T]
	if err := value.Decode(&dto); err != nil {
		return err
	}

	rs.Items = dto.Items
	rs.appError = dto.AppError

	return nil
}

// OkSlice creates a successful ResultSlice.
func OkSlice[T any](items []T) ResultSlice[T] {
	return ResultSlice[T]{
		Items: items,
	}
}

// FailSlice creates a failed ResultSlice from an AppError.
func FailSlice[T any](err *AppError) ResultSlice[T] {
	return ResultSlice[T]{
		appError: err,
	}
}

// IsSuccess returns true if no error is present.
func (rs ResultSlice[T]) IsSuccess() bool {
	return rs.appError == nil
}

// IsFailed returns true if an error is present.
func (rs ResultSlice[T]) IsFailed() bool {
	return rs.appError != nil
}

// IsFailure returns true if an error is present.
func (rs ResultSlice[T]) IsFailure() bool {
	return rs.IsFailed()
}

// IsInvalid returns true if an error is present.
func (rs ResultSlice[T]) IsInvalid() bool {
	return rs.IsFailed()
}

// IsNull returns true if no error is present.
func (rs ResultSlice[T]) IsNull() bool {
	return rs.appError == nil
}

// IsEmpty returns true if no active error is present (or items are empty).
func (rs ResultSlice[T]) IsEmpty() bool {
	if rs.appError == nil {
		return len(rs.Items) == 0
	}

	return rs.appError.IsEmpty()
}

// IsDefined returns true if operation succeeded.
func (rs ResultSlice[T]) IsDefined() bool {
	return rs.IsSuccess()
}

// AsSimpleVerifier returns the ResultSlice conforming to SimpleVerifier.
func (rs ResultSlice[T]) AsSimpleVerifier() SimpleVerifier {
	return rs
}

// AsSimpleVerifyChecker returns the ResultSlice conforming to SimpleVerifier.
func (rs ResultSlice[T]) AsSimpleVerifyChecker() SimpleVerifier {
	return rs
}

// HasError returns true if an error is present.
func (rs ResultSlice[T]) HasError() bool {
	return rs.IsFailed()
}

// HasItems returns true if the slice contains elements and is safe.
func (rs ResultSlice[T]) HasItems() bool {
	if rs.IsFailed() {
		return false
	}

	return len(rs.Items) > 0
}

// Count returns the number of items or 0 if failed.
func (rs ResultSlice[T]) Count() int {
	if rs.IsFailed() {
		return 0
	}

	return len(rs.Items)
}

// Length is an alias for Count.
func (rs ResultSlice[T]) Length() int {
	return rs.Count()
}

// AppError returns the underlying *AppError.
func (rs ResultSlice[T]) AppError() *AppError {
	return rs.appError
}

// Fault returns the underlying *AppError (alias for AppError).
func (rs ResultSlice[T]) Fault() *AppError {
	return rs.appError
}

// Error returns the underlying *AppError.
func (rs ResultSlice[T]) Error() *AppError {
	return rs.appError
}

// ValueAny returns the items as an any interface.
func (rs ResultSlice[T]) ValueAny() any {
	return rs.Items
}

// Unwrap unpacks the ([]T, *AppError) tuple.
func (rs ResultSlice[T]) Unwrap() ([]T, *AppError) {
	return rs.Items, rs.appError
}

// Filter returns a new ResultSlice containing items that satisfy predicate.
func (rs ResultSlice[T]) Filter(predicate func(item T) bool) ResultSlice[T] {
	if rs.IsFailed() || predicate == nil {
		return rs
	}

	filtered := make([]T, 0, len(rs.Items))
	for _, item := range rs.Items {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}

	return OkSlice(filtered)
}

// ForEach iterates over all items passing index and item to fn.
func (rs ResultSlice[T]) ForEach(fn func(index int, item T)) ResultSlice[T] {
	if rs.IsFailed() || fn == nil {
		return rs
	}

	for i, item := range rs.Items {
		fn(i, item)
	}

	return rs
}

// ForEachBreak iterates over items passing index and item to fn, stopping early if fn returns true.
func (rs ResultSlice[T]) ForEachBreak(fn func(index int, item T) bool) ResultSlice[T] {
	if rs.IsFailed() || fn == nil {
		return rs
	}

	for i, item := range rs.Items {
		if fn(i, item) {
			break
		}
	}

	return rs
}

func buildSliceBlock[T any](items []T) string {
	var b strings.Builder
	b.WriteString("[\n")
	for i, item := range items {
		b.WriteString(fmt.Sprintf("  [%d] %+v\n", i, item))
	}

	b.WriteString("]")

	return b.String()
}

// FormatStruct formats slice items in aligned block or error banner.
func (rs ResultSlice[T]) FormatStruct() string {
	if rs.IsFailed() {
		return rs.appError.FormatStdout()
	}

	if len(rs.Items) == 0 {
		return "[]"
	}

	return buildSliceBlock(rs.Items)
}

// String returns a human-readable string representation of the slice or error.
func (rs ResultSlice[T]) String() string {
	if rs.IsFailed() {
		return rs.appError.FormatStdout()
	}

	return FormatValue(rs.Items)
}

// PrettyJson returns the items formatted as indented JSON with sorted keys.
func (rs ResultSlice[T]) PrettyJson() string {
	if rs.IsFailed() {
		return rs.appError.FormatJson()
	}

	return FormatSortedJson(rs.Items)
}
