package result

import "coding-guidelines/common/pkg/appfault"

// Result re-exports appfault.Result for direct import compatibility.
type Result[T any] = appfault.Result[T]

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

// NewFailure creates a failed Result from a raw error.
func NewFailure[T any](err error) Result[T] {
	return appfault.NewFailure[T](err)
}

// NewFailureWithType creates a failed Result with explicit error code.
func NewFailureWithType[T any](errCode string, msg string, caller string) Result[T] {
	return appfault.NewFailureWithType[T](errCode, msg, caller)
}
