package appfault

// ResultMap wraps a generic key-value map with monadic error state.
type ResultMap[K comparable, V any] struct {
	Data   map[K]V   `json:"data,omitempty"`
	Err    *AppError `json:"err,omitempty"`
	AppErr error     `json:"appError,omitempty"`
}

// OkMap creates a successful ResultMap.
func OkMap[K comparable, V any](data map[K]V) ResultMap[K, V] {
	return ResultMap[K, V]{
		Data: data,
	}
}

// FailMap creates a failed ResultMap from a AppError.
func FailMap[K comparable, V any](err *AppError) ResultMap[K, V] {
	return ResultMap[K, V]{
		Err:    err,
		AppErr: err,
	}
}

// IsSuccess returns true if no error is present.
func (rm ResultMap[K, V]) IsSuccess() bool {
	return rm.Err == nil && rm.AppErr == nil
}

// IsFailed returns true if an error is present.
func (rm ResultMap[K, V]) IsFailed() bool {
	return rm.Err != nil || rm.AppErr != nil
}

// HasError returns true if an error is present.
func (rm ResultMap[K, V]) HasError() bool {
	return rm.IsFailed()
}

// Has returns true if the key exists in the map.
func (rm ResultMap[K, V]) Has(key K) bool {
	if rm.IsFailed() || rm.Data == nil {
		return false
	}

	_, ok := rm.Data[key]

	return ok
}

// Get retrieves a key's value as a Result[V].
func (rm ResultMap[K, V]) Get(key K) Result[V] {
	if rm.IsFailed() {
		return Fail[V](rm.Err)
	}

	val, ok := rm.Data[key]
	if !ok {
		return FailNew[V]("map.get", ErrDatabaseNotFound.String(), "key not found in map")
	}

	return Ok(val)
}

// Count returns the number of entries in the map.
func (rm ResultMap[K, V]) Count() int {
	if rm.IsFailed() || rm.Data == nil {
		return 0
	}

	return len(rm.Data)
}

// AppError returns the underlying *AppError.
func (rm ResultMap[K, V]) AppError() *AppError {
	return rm.Err
}

// Fault returns the underlying *AppError (alias for AppError()).
func (rm ResultMap[K, V]) Fault() *AppError {
	return rm.Err
}
