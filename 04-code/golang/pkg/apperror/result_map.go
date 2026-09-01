package apperror

// ResultMap wraps a generic key-value map with explicit status and Fault.
type ResultMap[K comparable, V any] struct {
	isSuccess bool
	isFailed  bool
	data      map[K]V
	fault     *Fault
}

// OkMap creates a successful ResultMap holding map data.
func OkMap[K comparable, V any](data map[K]V) ResultMap[K, V] {
	return ResultMap[K, V]{
		isSuccess: true,
		isFailed:  false,
		data:      data,
	}
}

// FailMap creates a failed ResultMap from a Fault.
func FailMap[K comparable, V any](err *Fault) ResultMap[K, V] {
	return ResultMap[K, V]{
		isSuccess: false,
		isFailed:  true,
		fault:     err,
	}
}

// FailMapWrap wraps a raw error into a Fault and returns a failed ResultMap.
func FailMapWrap[K comparable, V any](cause error, code ErrorCodeType, message string) ResultMap[K, V] {
	return FailMap[K, V](Wrap(cause, code, message))
}

// HasError returns true if the map operation failed.
func (rm ResultMap[K, V]) HasError() bool {
	return rm.isFailed
}

// IsSafe returns true if the map operation succeeded with no error.
func (rm ResultMap[K, V]) IsSafe() bool {
	return rm.isSuccess
}

// HasItems returns true if the map contains at least one entry.
func (rm ResultMap[K, V]) HasItems() bool {
	return len(rm.data) > 0
}

// Count returns the number of entries in the map.
func (rm ResultMap[K, V]) Count() int {
	return len(rm.data)
}

// Items returns the raw map or nil if in error state.
func (rm ResultMap[K, V]) Items() map[K]V {
	if rm.isFailed {
		return nil
	}

	return rm.data
}

// Get retrieves a key's value as a Result[V].
func (rm ResultMap[K, V]) Get(key K) Result[V] {
	if rm.isFailed {
		return Fail[V](rm.fault)
	}

	val, ok := rm.data[key]
	if !ok {
		return FailNew[V](ErrDatabaseNotFound, "key not found in map")
	}

	return Ok(val)
}

// Has returns true if the specified key exists in the map.
func (rm ResultMap[K, V]) Has(key K) bool {
	if rm.isFailed {
		return false
	}

	_, ok := rm.data[key]

	return ok
}

// Fault returns the underlying *Fault or nil.
func (rm ResultMap[K, V]) Fault() *Fault {
	return rm.fault
}

// AppError returns the underlying *Fault (alias for Fault()).
func (rm ResultMap[K, V]) AppError() *Fault {
	return rm.fault
}
