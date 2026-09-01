package apperror

// ResultSlice wraps a generic slice collection with explicit status and AppError.
type ResultSlice[T any] struct {
	isSuccess bool
	isFailed  bool
	items     []T
	appErr    *AppError
}

// OkSlice creates a successful ResultSlice from items.
func OkSlice[T any](items []T) ResultSlice[T] {
	return ResultSlice[T]{
		isSuccess: true,
		isFailed:  false,
		items:     items,
	}
}

// FailSlice creates a failed ResultSlice from an AppError.
func FailSlice[T any](err *AppError) ResultSlice[T] {
	return ResultSlice[T]{
		isSuccess: false,
		isFailed:  true,
		appErr:    err,
	}
}

// FailSliceWrap wraps a raw error into an AppError and returns a failed ResultSlice.
func FailSliceWrap[T any](cause error, code ErrorCodeType, message string) ResultSlice[T] {
	return FailSlice[T](Wrap(cause, code, message))
}

// HasError returns true if the collection query failed.
func (rs ResultSlice[T]) HasError() bool {
	return rs.isFailed
}

// IsSafe returns true if the operation succeeded with no error.
func (rs ResultSlice[T]) IsSafe() bool {
	return rs.isSuccess
}

// HasItems returns true if the slice contains at least one element.
func (rs ResultSlice[T]) HasItems() bool {
	return len(rs.items) > 0
}

// IsEmpty returns true if the slice contains zero elements.
func (rs ResultSlice[T]) IsEmpty() bool {
	return len(rs.items) == 0
}

// Count returns the total number of items in the collection.
func (rs ResultSlice[T]) Count() int {
	return len(rs.items)
}

// Items returns the raw slice of items or nil if in error state.
func (rs ResultSlice[T]) Items() []T {
	if rs.isFailed {
		return nil
	}

	return rs.items
}

// First returns a Result containing the first item or empty.
func (rs ResultSlice[T]) First() Result[T] {
	if rs.isFailed {
		return Fail[T](rs.appErr)
	}

	if len(rs.items) == 0 {
		return FailNew[T](ErrDatabaseNotFound, "slice is empty")
	}

	return Ok(rs.items[0])
}

// AppError returns the underlying *AppError or nil.
func (rs ResultSlice[T]) AppError() *AppError {
	return rs.appErr
}
