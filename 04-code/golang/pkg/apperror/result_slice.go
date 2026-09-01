package apperror

// ResultSlice wraps a generic slice collection with monadic error state.
type ResultSlice[T any] struct {
	Items    []T    `json:"items,omitempty"`
	Err      *Fault `json:"err,omitempty"`
	AppError error  `json:"appError,omitempty"`
}

// OkSlice creates a successful ResultSlice.
func OkSlice[T any](items []T) ResultSlice[T] {
	return ResultSlice[T]{
		Items: items,
	}
}

// FailSlice creates a failed ResultSlice from a Fault.
func FailSlice[T any](err *Fault) ResultSlice[T] {
	return ResultSlice[T]{
		Err:      err,
		AppError: err,
	}
}

// IsSuccess returns true if no error is present.
func (rs ResultSlice[T]) IsSuccess() bool {
	return rs.Err == nil && rs.AppError == nil
}

// IsFailed returns true if an error is present.
func (rs ResultSlice[T]) IsFailed() bool {
	return rs.Err != nil || rs.AppError != nil
}

// HasError returns true if an error is present.
func (rs ResultSlice[T]) HasError() bool {
	return rs.IsFailed()
}

// HasItems returns true if the slice contains elements and is safe.
func (rs ResultSlice[T]) HasItems() bool {
	if rs.IsFailed() {
		return false
	}

	return len(rs.Items) > 0
}

// Count returns the number of items.
func (rs ResultSlice[T]) Count() int {
	if rs.IsFailed() {
		return 0
	}

	return len(rs.Items)
}

// First returns the first element or empty.
func (rs ResultSlice[T]) First() Result[T] {
	if rs.IsFailed() {
		return Fail[T](rs.Err)
	}

	if len(rs.Items) == 0 {
		return FailNew[T]("slice.first", ErrDatabaseNotFound.String(), "slice is empty")
	}

	return Ok(rs.Items[0])
}

// Fault returns the underlying *Fault.
func (rs ResultSlice[T]) Fault() *Fault {
	return rs.Err
}
