package apperror

// Result wraps a typed value bundled with monadic error state.
type Result[T any] struct {
	Value    T      `json:"value,omitempty"`
	Err      *Fault `json:"err,omitempty"`
	AppError error  `json:"appError,omitempty"`
}

// Data returns the underlying Value payload for API envelope compatibility.
func (r Result[T]) Data() T {
	return r.Value
}

// Fault returns the underlying *Fault.
func (r Result[T]) Fault() *Fault {
	return r.Err
}

// FaultOrNil returns *Fault if present, otherwise nil.
func (r Result[T]) FaultOrNil() *Fault {
	if r.Err != nil {
		return r.Err
	}

	return nil
}
