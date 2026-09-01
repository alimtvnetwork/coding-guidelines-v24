package result

import "coding-guidelines/common/pkg/apperror"

// Result re-exports apperror.Result for direct import compatibility.
type Result[T any] = apperror.Result[T]

// SuccessResult creates a successful Result.
func SuccessResult[T any](val T) Result[T] {
	return apperror.SuccessResult(val)
}

// NewSuccess creates a successful Result.
func NewSuccess[T any](data T) Result[T] {
	return apperror.NewSuccess(data)
}

// FailureResult creates a failed Result.
func FailureResult[T any](err *apperror.Fault) Result[T] {
	return apperror.FailureResult[T](err)
}

// NewFailure creates a failed Result from a raw error.
func NewFailure[T any](err error) Result[T] {
	return apperror.NewFailure[T](err)
}

// NewFailureWithType creates a failed Result with explicit error code.
func NewFailureWithType[T any](errCode string, msg string, caller string) Result[T] {
	return apperror.NewFailureWithType[T](errCode, msg, caller)
}
