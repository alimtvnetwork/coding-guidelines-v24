package apperror

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

// FailureResult creates a failed Result from a Fault.
func FailureResult[T any](err *Fault) Result[T] {
	return Result[T]{
		Err:      err,
		AppError: err,
	}
}

// NewFailure creates a failed Result from a raw error.
func NewFailure[T any](err error) Result[T] {
	if err == nil {
		return Result[T]{}
	}

	f := WrapSimple(err, "operation")

	return FailureResult[T](f)
}

// NewFailureWithType creates a failed Result with explicit code and caller.
func NewFailureWithType[T any](errCode string, msg string, caller string) Result[T] {
	f := &Fault{
		Code:     errCode,
		Message:  msg,
		Caller:   caller,
		Type:     ErrorTypeExecution,
		Severity: SeverityError,
		Ctx:      make(map[string]any),
	}

	return FailureResult[T](f)
}

// Fail creates a failed Result from a Fault.
func Fail[T any](err *Fault) Result[T] {
	return FailureResult[T](err)
}

// FailWrap wraps a raw error into a failed Result.
func FailWrap[T any](cause error, op string, ctx map[string]any) Result[T] {
	return FailureResult[T](Wrap(cause, op, ctx))
}

// FailNew creates a new Fault and returns a failed Result.
func FailNew[T any](op, code, msg string) Result[T] {
	f := NewWithDetails(op, code, msg, "", ErrorTypeExecution, SeverityError, nil)

	return FailureResult[T](f)
}
