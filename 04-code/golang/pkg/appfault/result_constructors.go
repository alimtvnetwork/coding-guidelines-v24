package appfault

import "coding-guidelines/common/pkg/errtype"

// SuccessResult creates a successful Result holding a valid payload.
func SuccessResult[T any](val T) Result[T] {
	return Result[T]{
		Value: val,
	}
}

// NewSuccess creates a successful Result (alias for SuccessResult).
func NewSuccess[T any](data T) Result[T] {
	return SuccessResult(data)
}

// Ok creates a successful Result (standard alias).
func Ok[T any](val T) Result[T] {
	return SuccessResult(val)
}

// FailureResult creates a failed Result from an AppError.
func FailureResult[T any](err *AppError) Result[T] {
	return Result[T]{
		AppError: err,
	}
}

// NewFailure creates a failed Result from an explicit type and cause error.
func NewFailure[T any](errType errtype.Variation, cause error) Result[T] {
	if cause == nil || errType == errtype.None {
		return Result[T]{}
	}

	return FailureResult[T](WrapType(errType, cause))
}

// NewFailureWithType creates a failed Result with explicit type and caller.
func NewFailureWithType[T any](errType errtype.Variation, msg string, caller string) Result[T] {
	e := &AppError{
		Type:    errType,
		Message: msg,
		Caller:  caller,
		Ctx:     NewContextMap(),
	}

	return FailureResult[T](e)
}

// Fail creates a failed Result from an AppError.
func Fail[T any](err *AppError) Result[T] {
	return FailureResult[T](err)
}

// FailWrap wraps a raw error into a failed Result.
func FailWrap[T any](errType errtype.Variation, cause error, msg string) Result[T] {
	return FailureResult[T](Wrap(errType, cause, msg))
}

// FailNew creates a new AppError and returns a failed Result.
func FailNew[T any](errType errtype.Variation, msg string) Result[T] {
	return FailureResult[T](New(errType, msg))
}
