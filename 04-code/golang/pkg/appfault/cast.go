package appfault

import (
	"fmt"

	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/typecast"
)

// CastTo casts source to target type T using type assertion.
func CastTo[T any](source any) (T, *AppError) {
	val, err := typecast.CastTo[T](source)
	if err != nil {
		return val, Wrap(errtype.TypeMismatch, err, err.Error())
	}

	return val, nil
}

// ReflectTo dynamically converts source to target type T using reflection.
func ReflectTo[T any](source any) (T, *AppError) {
	val, err := typecast.ReflectTo[T](source)
	if err != nil {
		return val, Wrap(errtype.TypeMismatch, err, err.Error())
	}

	return val, nil
}

// CastResult converts Result[T] to Result[U] using ReflectTo on data.
func CastResult[T any, U any](r Result[T]) Result[U] {
	if r.IsFailed() {
		return FailureResult[U](r.Fault())
	}

	val, appErr := ReflectTo[U](r.Data())
	if appErr != nil {
		return FailureResult[U](appErr)
	}

	return SuccessResult(val)
}

// CastContextPayload retrieves a key from an AppError context and casts to T.
func CastContextPayload[T any](e *AppError, key string) (T, *AppError) {
	var target T
	if e == nil {
		return target, New(errtype.NotFound, "cannot extract context from nil AppError")
	}

	val, exists := e.Context().Get(key)
	if !exists {
		return target, New(errtype.NotFound, fmt.Sprintf("context key %q not found", key))
	}

	return CastTo[T](val)
}

// ResultToBytes serializes result data or fault using typecast.ToBytes.
func ResultToBytes[T any](r Result[T]) Result[[]byte] {
	var payload any = r.Data()
	if r.IsFailed() {
		payload = r.Fault()
	}

	b, err := typecast.ToBytes(payload)
	if err != nil {
		return FailureResult[[]byte](Wrap(errtype.Serialization, err, err.Error()))
	}

	return SuccessResult(b)
}

// ResultToJSON serializes result data or fault using typecast.ToJSON.
func ResultToJSON[T any](r Result[T]) Result[[]byte] {
	var payload any = r.Data()
	if r.IsFailed() {
		payload = r.Fault()
	}

	b, err := typecast.ToJSON(payload)
	if err != nil {
		return FailureResult[[]byte](Wrap(errtype.Serialization, err, err.Error()))
	}

	return SuccessResult(b)
}
