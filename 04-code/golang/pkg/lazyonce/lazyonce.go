package lazyonce

import (
	"context"
	"sync"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type evalResult[T any] struct {
	val   T
	fault *appfault.AppError
}

// LazyOnce provides thread-safe, lazy memoization for zero-parameter initializers.
type LazyOnce[T any] struct {
	lock        sync.Mutex
	isEvaluated bool
	cachedValue T
	cachedFault *appfault.AppError
	initFn      func() (T, *appfault.AppError)
}

// New creates a LazyOnce instance with the provided evaluation function.
func New[T any](initFn func() (T, *appfault.AppError)) *LazyOnce[T] {
	return &LazyOnce[T]{initFn: initFn}
}

// NewResult creates a LazyOnce instance from a Result-returning function.
func NewResult[T any](initFn func() result.Result[T]) *LazyOnce[T] {
	adapted := func() (T, *appfault.AppError) {
		res := initFn()

		return res.Data(), res.Fault()
	}

	return New[T](adapted)
}

// Value executes the initializer once and returns the memoized value and fault.
func (o *LazyOnce[T]) Value() (T, *appfault.AppError) {
	o.lock.Lock()
	defer o.lock.Unlock()

	if o.isEvaluated {
		return o.cachedValue, o.cachedFault
	}

	if o.initFn != nil {
		o.cachedValue, o.cachedFault = o.initFn()
	}

	o.isEvaluated = true

	return o.cachedValue, o.cachedFault
}

// ValueContext evaluates the initializer respecting context deadlines.
func (o *LazyOnce[T]) ValueContext(ctx context.Context) (T, *appfault.AppError) {
	if ctx == nil {
		return o.Value()
	}

	if ctx.Err() != nil {
		var zero T

		return zero, appfault.Wrap(errtype.Timeout, ctx.Err(), "context canceled before evaluation")
	}

	return o.awaitContext(ctx)
}

func (o *LazyOnce[T]) awaitContext(ctx context.Context) (T, *appfault.AppError) {
	resCh := make(chan evalResult[T], 1)
	go func() {
		v, f := o.Value()
		resCh <- evalResult[T]{val: v, fault: f}
	}()

	select {
	case <-ctx.Done():
		var zero T

		return zero, appfault.Wrap(errtype.Timeout, ctx.Err(), "context canceled during evaluation")
	case res := <-resCh:
		return res.val, res.fault
	}
}

// Result executes the initializer once and wraps the memoized output in a Result container.
func (o *LazyOnce[T]) Result() result.Result[T] {
	val, fault := o.Value()
	if fault != nil {
		return result.Failure[T](fault)
	}

	return result.Success[T](val)
}

// ResultContext executes the initializer respecting context cancellation.
func (o *LazyOnce[T]) ResultContext(ctx context.Context) result.Result[T] {
	val, fault := o.ValueContext(ctx)
	if fault != nil {
		return result.Failure[T](fault)
	}

	return result.Success[T](val)
}

// IsEvaluated reports whether the initializer has already been executed.
func (o *LazyOnce[T]) IsEvaluated() bool {
	o.lock.Lock()
	defer o.lock.Unlock()

	return o.isEvaluated
}

// Reset clears the cached value allowing re-execution (primarily for testing).
func (o *LazyOnce[T]) Reset() {
	o.lock.Lock()
	defer o.lock.Unlock()

	var zero T
	o.cachedValue = zero
	o.cachedFault = nil
	o.isEvaluated = false
}
