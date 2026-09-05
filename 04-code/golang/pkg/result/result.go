package result

import (
	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

// Wrap represents the canonical monadic result container wrapping either a typed value
// or an *appfault.AppError. Callers write result.Wrap[T] to eliminate package-type stutter
// (preventing result.Result[T]).
type Wrap[T any] = appfault.Result[T]

// Result is maintained as an alias to Wrap[T] for backward compatibility.
type Result[T any] = Wrap[T]

// WrapSuccess creates a successful Wrap container wrapping data.
func WrapSuccess[T any](data T) Wrap[T] {
	return appfault.NewSuccess(data)
}

// WrapFailure creates a failed Wrap container with a structured AppError.
func WrapFailure[T any](err *appfault.AppError) Wrap[T] {
	return appfault.FailureResult[T](err)
}

// Success creates a successful Wrap container wrapping data (short form).
func Success[T any](data T) Wrap[T] {
	return appfault.NewSuccess(data)
}

// Failure creates a failed Wrap container with a structured AppError (short form).
func Failure[T any](err *appfault.AppError) Wrap[T] {
	return appfault.FailureResult[T](err)
}

// WrapFailureFromError creates a failed Wrap container directly from an AppError object.
func WrapFailureFromError[T any](err *appfault.AppError) Wrap[T] {
	return appfault.FailureResult[T](err)
}

// WrapFailureWithId creates a failed Wrap using an error ID (errtype.Variation) and message.
func WrapFailureWithId[T any](errType errtype.Variation, msg string) Wrap[T] {
	return appfault.NewFailureWithId[T](errType, msg)
}

// WrapFailureWithCause creates a failed Wrap using an error ID, cause error, and message.
func WrapFailureWithCause[T any](errType errtype.Variation, cause error, msg string) Wrap[T] {
	return appfault.NewFailureWithCause[T](errType, cause, msg)
}

// WrapFailureFromWrap propagates an error from one failed Wrap to another.
func WrapFailureFromWrap[T any, U any](failed Wrap[U]) Wrap[T] {
	return appfault.FailureFromWrap[T](failed)
}

// FailureFromWrap propagates an error from one failed Wrap to another (short form).
func FailureFromWrap[T any, U any](failed Wrap[U]) Wrap[T] {
	return appfault.FailureFromWrap[T](failed)
}

// FailureWithId creates a failed Wrap with an error ID and message (short form).
func FailureWithId[T any](errType errtype.Variation, msg string) Wrap[T] {
	return appfault.NewFailureWithId[T](errType, msg)
}

// FailureWithCause creates a failed Wrap with an error ID, cause error, and message (short form).
func FailureWithCause[T any](errType errtype.Variation, cause error, msg string) Wrap[T] {
	return appfault.NewFailureWithCause[T](errType, cause, msg)
}

// SuccessResult creates a successful Result (legacy compatibility).
func SuccessResult[T any](val T) Result[T] {
	return appfault.SuccessResult(val)
}

// NewSuccess creates a successful Result (legacy compatibility).
func NewSuccess[T any](data T) Result[T] {
	return appfault.NewSuccess(data)
}

// FailureResult creates a failed Result (legacy compatibility).
func FailureResult[T any](err *appfault.AppError) Result[T] {
	return appfault.FailureResult[T](err)
}

// NewFailure creates a failed Result from an explicit error type and cause (legacy compatibility).
func NewFailure[T any](errType errtype.Variation, cause error) Result[T] {
	return appfault.NewFailure[T](errType, cause)
}

// NewFailureWithType creates a failed Result with explicit error type (legacy compatibility).
func NewFailureWithType[T any](errType errtype.Variation, msg string, caller string) Result[T] {
	return appfault.NewFailureWithType[T](errType, msg, caller)
}
