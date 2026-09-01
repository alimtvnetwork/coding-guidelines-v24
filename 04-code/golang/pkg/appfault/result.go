package appfault

// Result wraps a typed value bundled with monadic error state.
type Result[T any] struct {
	Value  T         `json:"value,omitempty"`
	Err    *AppError `json:"err,omitempty"`
	AppErr error     `json:"appError,omitempty"`
}

// Data returns the underlying Value payload for API envelope compatibility.
func (r Result[T]) Data() T {
	return r.Value
}

// AppError returns the underlying *AppError.
func (r Result[T]) AppError() *AppError {
	return r.Err
}

// Fault returns the underlying *AppError (alias for AppError()).
func (r Result[T]) Fault() *AppError {
	return r.Err
}

// AppErrorOrNil returns *AppError if present, otherwise nil.
func (r Result[T]) AppErrorOrNil() *AppError {
	if r.Err != nil {
		return r.Err
	}

	return nil
}
