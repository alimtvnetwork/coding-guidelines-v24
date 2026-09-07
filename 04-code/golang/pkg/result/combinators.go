package result

// Map transforms a successful Result[T] into a Result[U] using the provided function.
// If the input Result is a failure, it safely passes the error forward without calling the function.
func Map[T, U any](res Result[T], fn func(T) U) Result[U] {
	if res.IsFailure() {
		return FailureFromWrap[U](res)
	}

	return Success(fn(res.Data()))
}

// FlatMap binds a successful Result[T] to a function that returns a new Result[U].
// This allows chaining multiple operations that can fail, short-circuiting on the first error.
func FlatMap[T, U any](res Result[T], fn func(T) Result[U]) Result[U] {
	if res.IsFailure() {
		return FailureFromWrap[U](res)
	}

	return fn(res.Data())
}

// Tap executes a side-effect function if the Result is successful, then returns the original Result unmodified.
// This is useful for logging, metrics, or auditing without breaking the pipeline.
func Tap[T any](res Result[T], fn func(T)) Result[T] {
	if res.IsSuccess() {
		fn(res.Data())
	}

	return res
}

// MapSlice transforms each item of a successful Wrap[[]T] using fn.
func MapSlice[T, U any](res Wrap[[]T], fn func(T) U) ResultSlice[U] {
	if res.IsFailure() {
		return FailSlice[U](res.Fault())
	}

	items := res.Data()
	mapped := make([]U, len(items))
	for i, item := range items {
		mapped[i] = fn(item)
	}

	return OkSlice(mapped)
}

// FlatMapSlice transforms each item into a slice of U and flattens into a ResultSlice[U].
func FlatMapSlice[T, U any](res Wrap[[]T], fn func(T) []U) ResultSlice[U] {
	if res.IsFailure() {
		return FailSlice[U](res.Fault())
	}

	var flat []U
	for _, item := range res.Data() {
		flat = append(flat, fn(item)...)
	}

	return OkSlice(flat)
}

// MapMapValues transforms values in ResultMap using fn preserving keys.
func MapMapValues[K comparable, V, U any](res ResultMap[K, V], fn func(V) U) ResultMap[K, U] {
	if res.IsFailure() {
		return FailMap[K, U](res.Fault())
	}

	mapped := make(map[K]U, len(res.Data))
	for k, v := range res.Data {
		mapped[k] = fn(v)
	}

	return OkMap(mapped)
}
