package lazyonce

import (
	"sync"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/result"
)

// LazyOnce2 provides thread-safe, lazy memoization for two-parameter initializers.
type LazyOnce2[T1, T2, TOutput any] struct {
	lock        sync.Mutex
	isEvaluated bool
	cachedValue TOutput
	cachedFault *appfault.AppError
	initFn      func(T1, T2) (TOutput, *appfault.AppError)
}

// New2 creates a LazyOnce2 instance with the provided two-parameter evaluation function.
func New2[T1, T2, TOutput any](
	initFn func(T1, T2) (TOutput, *appfault.AppError),
) *LazyOnce2[T1, T2, TOutput] {
	return &LazyOnce2[T1, T2, TOutput]{initFn: initFn}
}

// Value executes the initializer on the first call and returns the memoized result.
func (o *LazyOnce2[T1, T2, TOutput]) Value(
	arg1 T1,
	arg2 T2,
) (TOutput, *appfault.AppError) {
	o.lock.Lock()
	defer o.lock.Unlock()

	if o.isEvaluated {
		return o.cachedValue, o.cachedFault
	}

	if o.initFn != nil {
		o.cachedValue, o.cachedFault = o.initFn(arg1, arg2)
	}

	o.isEvaluated = true

	return o.cachedValue, o.cachedFault
}

// Result executes the initializer on the first call and returns a Result container.
func (o *LazyOnce2[T1, T2, TOutput]) Result(
	arg1 T1,
	arg2 T2,
) result.Result[TOutput] {
	val, fault := o.Value(arg1, arg2)
	if fault != nil {
		return result.Failure[TOutput](fault)
	}

	return result.Success[TOutput](val)
}

// IsEvaluated reports whether the initializer has already been executed.
func (o *LazyOnce2[T1, T2, TOutput]) IsEvaluated() bool {
	o.lock.Lock()
	defer o.lock.Unlock()

	return o.isEvaluated
}

// Reset clears the cached value allowing re-execution.
func (o *LazyOnce2[T1, T2, TOutput]) Reset() {
	o.lock.Lock()
	defer o.lock.Unlock()

	var zero TOutput
	o.cachedValue = zero
	o.cachedFault = nil
	o.isEvaluated = false
}
