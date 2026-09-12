package appfault

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// Result wraps a typed value bundled with monadic error state.
type Result[T any] struct {
	value    T
	appError *AppError
}

type resultDTO[T any] struct {
	Value    T         `json:"value,omitempty" yaml:"value,omitempty"`
	AppError *AppError `json:"appError,omitempty" yaml:"appError,omitempty"`
}

// MarshalJSON provides JSON serialization for Result[T].
func (r Result[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(resultDTO[T]{
		Value:    r.value,
		AppError: r.appError,
	})
}

// UnmarshalJSON provides JSON deserialization for Result[T].
func (r *Result[T]) UnmarshalJSON(data []byte) error {
	var dto resultDTO[T]
	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	r.value = dto.Value
	r.appError = dto.AppError

	return nil
}

// MarshalYAML provides YAML serialization for Result[T].
func (r Result[T]) MarshalYAML() (any, error) {
	return resultDTO[T]{
		Value:    r.value,
		AppError: r.appError,
	}, nil
}

// UnmarshalYAML provides YAML deserialization for Result[T].
func (r *Result[T]) UnmarshalYAML(value *yaml.Node) error {
	var dto resultDTO[T]
	if err := value.Decode(&dto); err != nil {
		return err
	}

	r.value = dto.Value
	r.appError = dto.AppError

	return nil
}

// Value returns the underlying value payload.
func (r *Result[T]) Value() T {
	if r == nil {
		var zero T
		return zero
	}

	return r.value
}

// ValueAny returns the underlying value payload as an any interface.
func (r *Result[T]) ValueAny() any {
	if r == nil {
		return nil
	}

	return r.value
}

// Data returns the underlying value payload for API envelope compatibility.
func (r *Result[T]) Data() T {
	if r == nil {
		var zero T
		return zero
	}

	return r.value
}

// Payload returns the underlying value payload (mirroring streamwriter.Bytes).
func (r *Result[T]) Payload() T {
	if r == nil {
		var zero T
		return zero
	}

	return r.value
}

// Result returns the underlying value payload (mirroring coredynamic.TypedSimpleResult).
func (r *Result[T]) Result() T {
	if r == nil {
		var zero T
		return zero
	}

	return r.value
}

// AppError returns the underlying *AppError.
func (r *Result[T]) AppError() *AppError {
	if r == nil {
		return nil
	}

	return r.appError
}

// Fault returns the underlying *AppError (alias for AppError).
func (r *Result[T]) Fault() *AppError {
	if r == nil {
		return nil
	}

	return r.appError
}

// Error returns the underlying *AppError.
func (r *Result[T]) Error() *AppError {
	if r == nil {
		return nil
	}

	return r.appError
}

// AppErrorOrNil returns *AppError if present, otherwise nil.
func (r *Result[T]) AppErrorOrNil() *AppError {
	if r == nil {
		return nil
	}

	return r.appError
}

// ToJson exports Result as indented JSON bytes.
func (r *Result[T]) ToJson() ([]byte, error) {
	if r == nil {
		return []byte("{}"), nil
	}

	return json.MarshalIndent(r, "", "  ")
}

// ToJsonString exports Result as a JSON string.
func (r *Result[T]) ToJsonString() string {
	if r == nil {
		return "{}"
	}

	b, err := r.ToJson()
	if err != nil {
		return "{}"
	}

	return string(b)
}

// HandleError processes the underlying AppError if it exists.
// It defers to the AppError's internal null-check to proceed safely.
func (r *Result[T]) HandleError() {
	if r != nil && r.appError != nil {
		r.appError.HandleError()
	}
}
