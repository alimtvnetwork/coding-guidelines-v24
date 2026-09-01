package apperror

// Result wraps a single generic value bundled with explicit status and Fault.
type Result[T any] struct {
	isSuccess bool
	isFailed  bool
	value     T
	fault     *Fault
}

// Ok creates a successful Result holding a valid value.
func Ok[T any](value T) Result[T] {
	return Result[T]{
		isSuccess: true,
		isFailed:  false,
		value:     value,
	}
}

// Fail creates a failed Result from an existing Fault.
func Fail[T any](err *Fault) Result[T] {
	return Result[T]{
		isSuccess: false,
		isFailed:  true,
		fault:     err,
	}
}

// FailWrap wraps a raw error into a Fault and returns a failed Result.
func FailWrap[T any](cause error, code ErrorCodeType, message string) Result[T] {
	return Fail[T](Wrap(cause, code, message))
}

// FailNew creates a new Fault and returns a failed Result.
func FailNew[T any](code ErrorCodeType, message string) Result[T] {
	return Fail[T](New(code, message))
}

// HasError returns true if the operation failed.
func (r Result[T]) HasError() bool {
	return r.isFailed
}

// IsSafe returns true if the operation succeeded with no error.
func (r Result[T]) IsSafe() bool {
	return r.isSuccess
}

// Value returns the inner value or panics if the result is in an error state.
func (r Result[T]) Value() T {
	if r.isFailed {
		panic("called Value() on failed Result: " + r.fault.Error())
	}

	return r.value
}

// ValueOr returns the inner value if safe, or the provided fallback.
func (r Result[T]) ValueOr(fallback T) T {
	if r.isFailed {
		return fallback
	}

	return r.value
}

// Fault returns the underlying *Fault or nil.
func (r Result[T]) Fault() *Fault {
	return r.fault
}

// AppError returns the underlying *Fault (alias for Fault()).
func (r Result[T]) AppError() *Fault {
	return r.fault
}

// Unwrap bridges Result[T] back to Go's standard (T, error) tuple at framework boundaries.
func (r Result[T]) Unwrap() (T, error) {
	if r.isFailed {
		return r.value, r.fault
	}

	return r.value, nil
}
