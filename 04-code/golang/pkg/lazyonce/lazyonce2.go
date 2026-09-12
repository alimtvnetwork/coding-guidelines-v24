package lazyonce

import (
	"context"
	"sync"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
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

// ValueContext evaluates the initializer respecting context deadlines.
func (o *LazyOnce2[T1, T2, TOutput]) ValueContext(
	ctx context.Context,
	arg1 T1,
	arg2 T2,
) (TOutput, *appfault.AppError) {
	if ctx == nil {
		return o.Value(arg1, arg2)
	}

	if ctx.Err() != nil {
		var zero TOutput

		return zero, appfault.Wrap(errtype.Timeout, ctx.Err(), "context canceled before evaluation")
	}

	return o.awaitContext(ctx, arg1, arg2)
}

func (o *LazyOnce2[T1, T2, TOutput]) awaitContext(
	ctx context.Context,
	arg1 T1,
	arg2 T2,
) (TOutput, *appfault.AppError) {
	resCh := make(chan evalResult[TOutput], 1)
	go func() {
		v, f := o.Value(arg1, arg2)
		resCh <- evalResult[TOutput]{val: v, fault: f}
	}()

	select {
	case <-ctx.Done():
		var zero TOutput

		return zero, appfault.Wrap(errtype.Timeout, ctx.Err(), "context canceled during evaluation")
	case res := <-resCh:
		return res.val, res.fault
	}
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

// ResultContext executes the initializer respecting context cancellation.
func (o *LazyOnce2[T1, T2, TOutput]) ResultContext(
	ctx context.Context,
	arg1 T1,
	arg2 T2,
) result.Result[TOutput] {
	val, fault := o.ValueContext(ctx, arg1, arg2)
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
