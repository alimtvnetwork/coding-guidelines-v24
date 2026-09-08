package lazyonce

import (
	"sync"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/result"
)

// LazyOnce1 provides thread-safe, lazy memoization for single-parameter initializers.
type LazyOnce1[TInput, TOutput any] struct {
	lock        sync.Mutex
	isEvaluated bool
	cachedValue TOutput
	cachedFault *appfault.AppError
	initFn      func(TInput) (TOutput, *appfault.AppError)
}

// New1 creates a LazyOnce1 instance with the provided evaluation function.
func New1[TInput, TOutput any](
	initFn func(TInput) (TOutput, *appfault.AppError),
) *LazyOnce1[TInput, TOutput] {
	return &LazyOnce1[TInput, TOutput]{initFn: initFn}
}

// Value executes the initializer on the first call and returns the memoized result.
func (o *LazyOnce1[TInput, TOutput]) Value(input TInput) (TOutput, *appfault.AppError) {
	o.lock.Lock()
	defer o.lock.Unlock()

	if o.isEvaluated {
		return o.cachedValue, o.cachedFault
	}

	if o.initFn != nil {
		o.cachedValue, o.cachedFault = o.initFn(input)
	}

	o.isEvaluated = true

	return o.cachedValue, o.cachedFault
}

// Result executes the initializer on the first call and returns a Result container.
func (o *LazyOnce1[TInput, TOutput]) Result(input TInput) result.Result[TOutput] {
	val, fault := o.Value(input)
	if fault != nil {
		return result.Failure[TOutput](fault)
	}

	return result.Success[TOutput](val)
}

// IsEvaluated reports whether the initializer has already been executed.
func (o *LazyOnce1[TInput, TOutput]) IsEvaluated() bool {
	o.lock.Lock()
	defer o.lock.Unlock()

	return o.isEvaluated
}

// Reset clears the cached value allowing re-execution.
func (o *LazyOnce1[TInput, TOutput]) Reset() {
	o.lock.Lock()
	defer o.lock.Unlock()

	var zero TOutput
	o.cachedValue = zero
	o.cachedFault = nil
	o.isEvaluated = false
}
