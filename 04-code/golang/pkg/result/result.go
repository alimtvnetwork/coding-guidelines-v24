package result

import (
	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

// Result re-exports appfault.Result for direct import compatibility.
type Result[T any] = appfault.Result[T]

// Wrap aliases Result[T] to eliminate package-type stutter (result.Wrap instead of result.Result).
type Wrap[T any] = appfault.Result[T]

// WrapSuccess creates a successful Wrap container.
func WrapSuccess[T any](data T) Wrap[T] {
	return appfault.NewSuccess(data)
}

// WrapFailure creates a failed Wrap container.
func WrapFailure[T any](err *appfault.AppError) Wrap[T] {
	return appfault.FailureResult[T](err)
}

// SuccessResult creates a successful Result.
func SuccessResult[T any](val T) Result[T] {
	return appfault.SuccessResult(val)
}

// NewSuccess creates a successful Result.
func NewSuccess[T any](data T) Result[T] {
	return appfault.NewSuccess(data)
}

// FailureResult creates a failed Result.
func FailureResult[T any](err *appfault.AppError) Result[T] {
	return appfault.FailureResult[T](err)
}

// NewFailure creates a failed Result from an explicit error type and cause.
func NewFailure[T any](errType errtype.Variation, cause error) Result[T] {
	return appfault.NewFailure[T](errType, cause)
}

// NewFailureWithType creates a failed Result with explicit error type.
func NewFailureWithType[T any](errType errtype.Variation, msg string, caller string) Result[T] {
	return appfault.NewFailureWithType[T](errType, msg, caller)
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
