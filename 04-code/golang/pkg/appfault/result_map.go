package appfault

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// ResultMap wraps a generic key-value map with monadic error state.
type ResultMap[K comparable, V any] struct {
	Data     map[K]V
	appError *AppError
}

type resultMapDTO[K comparable, V any] struct {
	Data     map[K]V   `json:"data,omitempty" yaml:"data,omitempty"`
	AppError *AppError `json:"appError,omitempty" yaml:"appError,omitempty"`
}

// MarshalJSON provides JSON serialization for ResultMap[K, V].
func (rm ResultMap[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(resultMapDTO[K, V]{
		Data:     rm.Data,
		AppError: rm.appError,
	})
}

// UnmarshalJSON provides JSON deserialization for ResultMap[K, V].
func (rm *ResultMap[K, V]) UnmarshalJSON(data []byte) error {
	var dto resultMapDTO[K, V]
	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	rm.Data = dto.Data
	rm.appError = dto.AppError

	return nil
}

// MarshalYAML provides YAML serialization for ResultMap[K, V].
func (rm ResultMap[K, V]) MarshalYAML() (any, error) {
	return resultMapDTO[K, V]{
		Data:     rm.Data,
		AppError: rm.appError,
	}, nil
}

// UnmarshalYAML provides YAML deserialization for ResultMap[K, V].
func (rm *ResultMap[K, V]) UnmarshalYAML(value *yaml.Node) error {
	var dto resultMapDTO[K, V]
	if err := value.Decode(&dto); err != nil {
		return err
	}

	rm.Data = dto.Data
	rm.appError = dto.AppError

	return nil
}

// OkMap creates a successful ResultMap.
func OkMap[K comparable, V any](data map[K]V) ResultMap[K, V] {
	return ResultMap[K, V]{
		Data: data,
	}
}

// FailMap creates a failed ResultMap from an AppError.
func FailMap[K comparable, V any](err *AppError) ResultMap[K, V] {
	return ResultMap[K, V]{
		appError: err,
	}
}

// IsSuccess returns true if no error is present.
func (rm ResultMap[K, V]) IsSuccess() bool {
	return rm.appError == nil
}

// IsFailed returns true if an error is present.
func (rm ResultMap[K, V]) IsFailed() bool {
	return rm.appError != nil
}

// IsFailure returns true if an error is present.
func (rm ResultMap[K, V]) IsFailure() bool {
	return rm.IsFailed()
}

// IsInvalid returns true if an error is present.
func (rm ResultMap[K, V]) IsInvalid() bool {
	return rm.IsFailed()
}

// IsNull returns true if no error is present.
func (rm ResultMap[K, V]) IsNull() bool {
	return rm.appError == nil
}

// IsEmpty returns true if no active error is present (or map is empty).
func (rm ResultMap[K, V]) IsEmpty() bool {
	if rm.appError == nil {
		return len(rm.Data) == 0
	}

	return rm.appError.IsEmpty()
}

// IsDefined returns true if operation succeeded.
func (rm ResultMap[K, V]) IsDefined() bool {
	return rm.IsSuccess()
}

// AsSimpleVerifier returns the ResultMap conforming to SimpleVerifier.
func (rm ResultMap[K, V]) AsSimpleVerifier() SimpleVerifier {
	return rm
}

// AsSimpleVerifyChecker returns the ResultMap conforming to SimpleVerifier.
func (rm ResultMap[K, V]) AsSimpleVerifyChecker() SimpleVerifier {
	return rm
}

// HasError returns true if an error is present.
func (rm ResultMap[K, V]) HasError() bool {
	return rm.IsFailed()
}

// Has returns true if the key exists in the map.
func (rm ResultMap[K, V]) Has(key K) bool {
	if rm.IsFailed() || rm.Data == nil {
		return false
	}

	_, ok := rm.Data[key]

	return ok
}

// Get retrieves the value associated with key.
func (rm ResultMap[K, V]) Get(key K) (V, bool) {
	if rm.IsFailed() || rm.Data == nil {
		var zero V

		return zero, false
	}

	val, ok := rm.Data[key]

	return val, ok
}

// Count returns the number of entries in the map or 0 if failed.
func (rm ResultMap[K, V]) Count() int {
	if rm.IsFailed() || rm.Data == nil {
		return 0
	}

	return len(rm.Data)
}

// AppError returns the underlying *AppError.
func (rm ResultMap[K, V]) AppError() *AppError {
	return rm.appError
}

// Fault returns the underlying *AppError (alias for AppError).
func (rm ResultMap[K, V]) Fault() *AppError {
	return rm.appError
}

// Error returns the underlying *AppError.
func (rm ResultMap[K, V]) Error() *AppError {
	return rm.appError
}

// ValueAny returns the data map as an any interface.
func (rm ResultMap[K, V]) ValueAny() any {
	return rm.Data
}

// Unwrap unpacks the (map[K]V, *AppError) tuple.
func (rm ResultMap[K, V]) Unwrap() (map[K]V, *AppError) {
	return rm.Data, rm.appError
}

func collectMapKeys[K comparable, V any](data map[K]V) []K {
	keys := make([]K, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	return keys
}

// Keys returns a slice of map keys, sorted deterministically by string representation.
func (rm ResultMap[K, V]) Keys() []K {
	if rm.IsFailed() || len(rm.Data) == 0 {
		return []K{}
	}

	keys := collectMapKeys(rm.Data)
	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprint(keys[i]) < fmt.Sprint(keys[j])
	})

	return keys
}

// Values returns a slice of map values ordered according to Keys().
func (rm ResultMap[K, V]) Values() []V {
	if rm.IsFailed() || len(rm.Data) == 0 {
		return []V{}
	}

	vals := make([]V, 0, len(rm.Data))
	for _, k := range rm.Keys() {
		vals = append(vals, rm.Data[k])
	}

	return vals
}

// Filter returns a new ResultMap containing entries that satisfy predicate.
func (rm ResultMap[K, V]) Filter(predicate func(key K, val V) bool) ResultMap[K, V] {
	if rm.IsFailed() || predicate == nil {
		return rm
	}

	filtered := make(map[K]V)
	for k, v := range rm.Data {
		if predicate(k, v) {
			filtered[k] = v
		}
	}

	return OkMap(filtered)
}

// ForEach iterates over map entries passing key and value to fn.
func (rm ResultMap[K, V]) ForEach(fn func(key K, val V)) ResultMap[K, V] {
	if rm.IsFailed() || fn == nil {
		return rm
	}

	for _, k := range rm.Keys() {
		fn(k, rm.Data[k])
	}

	return rm
}

func buildMapBlock[K comparable, V any](rm ResultMap[K, V]) string {
	var b strings.Builder
	b.WriteString("{\n")
	for _, k := range rm.Keys() {
		b.WriteString(fmt.Sprintf("  %v: %+v\n", k, rm.Data[k]))
	}

	b.WriteString("}")

	return b.String()
}

// FormatStruct formats map key-values in aligned block or error banner.
func (rm ResultMap[K, V]) FormatStruct() string {
	if rm.IsFailed() {
		return rm.appError.FormatStdout()
	}

	if len(rm.Data) == 0 {
		return "{}"
	}

	return buildMapBlock(rm)
}

// String returns a human-readable string representation of the map or error.
func (rm ResultMap[K, V]) String() string {
	if rm.IsFailed() {
		return rm.appError.FormatStdout()
	}

	return FormatValue(rm.Data)
}

// PrettyJson returns the map formatted as indented JSON with sorted keys.
func (rm ResultMap[K, V]) PrettyJson() string {
	if rm.IsFailed() {
		return rm.appError.FormatJson()
	}

	return FormatSortedJson(rm.Data)
}
