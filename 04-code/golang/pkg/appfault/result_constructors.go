package appfault

import "coding-guidelines/common/pkg/errtype"

// SuccessResult creates a successful Result wrapping a value.
func SuccessResult[T any](val T) Result[T] {
	return Result[T]{value: val, appError: nil}
}

// NewResult creates a Result with value and appError.
func NewResult[T any](val T, appErr *AppError) Result[T] {
	return Result[T]{value: val, appError: appErr}
}

// ResultOf creates a Result with value and appError.
func ResultOf[T any](val T, appErr *AppError) Result[T] {
	return Result[T]{value: val, appError: appErr}
}

// NewSuccess creates a successful Result wrapping a value.
func NewSuccess[T any](data T) Result[T] {
	return SuccessResult(data)
}

// FailureResult creates a failed Result with a structured AppError.
func FailureResult[T any](err *AppError) Result[T] {
	return Result[T]{appError: err}
}

// Ok is an alias for SuccessResult.
func Ok[T any](val T) Result[T] {
	return SuccessResult(val)
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
	e := New(errType, msg)
	if len(caller) > 0 && e != nil {
		if len(e.stack) > 0 {
			e.stack[0].Function = caller
		} else {
			e.stack = NewStackTrace(NewStackFrame(caller, "", 0))
		}
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

// NewFailureWithId creates a failed Result with an error ID (errtype.Variation) and message.
func NewFailureWithId[T any](errType errtype.Variation, msg string) Result[T] {
	return FailureResult[T](New(errType, msg))
}

// NewFailureFromError creates a failed Result from an AppError object.
func NewFailureFromError[T any](err *AppError) Result[T] {
	return FailureResult[T](err)
}

// FailureWithId creates a failed Result with an error ID (errtype.Variation) and message.
func FailureWithId[T any](errType errtype.Variation, msg string) Result[T] {
	return NewFailureWithId[T](errType, msg)
}

// NewFailureWithCause creates a failed Result with an error ID, cause error, and message.
func NewFailureWithCause[T any](errType errtype.Variation, cause error, msg string) Result[T] {
	return FailureResult[T](Wrap(errType, cause, msg))
}

// FailureFromWrap creates a failed Result propagating the AppError from another Result.
func FailureFromWrap[T any, U any](failed Result[U]) Result[T] {
	return FailureResult[T](failed.appError)
}

// NewFailureWithFile creates a failed Result with file context and root cause error.
func NewFailureWithFile[T any](variation errtype.Variation, cause error, path string, msg string) Result[T] {
	return FailureResult[T](WrapFile(variation, cause, path, msg))
}

// NewFailureWithPath creates a failed Result with path context and root cause error.
func NewFailureWithPath[T any](variation errtype.Variation, cause error, path string, msg string) Result[T] {
	return FailureResult[T](WrapPath(variation, cause, path, msg))
}

// NewFailureWithVar creates a failed Result with variable context and root cause error.
func NewFailureWithVar[T any](variation errtype.Variation, cause error, key string, val any, msg string) Result[T] {
	return FailureResult[T](WrapVar(variation, cause, key, val, msg))
}

// FailurePath creates a failed Result with path context.
func FailurePath[T any](variation errtype.Variation, path string, msg string) Result[T] {
	return FailureResult[T](NewPath(variation, path, msg))
}

// FailureFile creates a failed Result with file context.
func FailureFile[T any](variation errtype.Variation, path string, msg string) Result[T] {
	return FailureResult[T](NewFile(variation, path, msg))
}

// FailureVar creates a failed Result with variable context.
func FailureVar[T any](variation errtype.Variation, key string, val any, msg string) Result[T] {
	return FailureResult[T](NewVar(variation, key, val, msg))
}
