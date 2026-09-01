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

// NewFailure creates a failed Result from a raw error.
func NewFailure[T any](err error) Result[T] {
	if err == nil {
		return Result[T]{}
	}

	e := WrapSimple(err)

	return FailureResult[T](e)
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
func FailWrap[T any](cause error, errType errtype.Variation, msg string) Result[T] {
	return FailureResult[T](Wrap(cause, errType, msg))
}

// FailNew creates a new AppError and returns a failed Result.
func FailNew[T any](errType errtype.Variation, msg string) Result[T] {
	return FailureResult[T](New(errType, msg))
}
