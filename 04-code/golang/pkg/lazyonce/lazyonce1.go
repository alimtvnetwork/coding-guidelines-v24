package lazyonce

import (
	"context"
	"sync"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
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

// ValueContext evaluates the initializer respecting context deadlines.
func (o *LazyOnce1[TInput, TOutput]) ValueContext(
	ctx context.Context,
	input TInput,
) (TOutput, *appfault.AppError) {
	if ctx == nil {
		return o.Value(input)
	}

	if ctx.Err() != nil {
		var zero TOutput

		return zero, appfault.Wrap(errtype.Timeout, ctx.Err(), "context cancelled before evaluation")
	}

	return o.awaitContext(ctx, input)
}

func (o *LazyOnce1[TInput, TOutput]) awaitContext(
	ctx context.Context,
	input TInput,
) (TOutput, *appfault.AppError) {
	resCh := make(chan evalResult[TOutput], 1)
	go func() {
		v, f := o.Value(input)
		resCh <- evalResult[TOutput]{val: v, fault: f}
	}()

	select {
	case <-ctx.Done():
		var zero TOutput

		return zero, appfault.Wrap(errtype.Timeout, ctx.Err(), "context cancelled during evaluation")
	case res := <-resCh:
		return res.val, res.fault
	}
}

// Result executes the initializer on the first call and returns a Result container.
func (o *LazyOnce1[TInput, TOutput]) Result(input TInput) result.Result[TOutput] {
	val, fault := o.Value(input)
	if fault != nil {
		return result.Failure[TOutput](fault)
	}

	return result.Success[TOutput](val)
}

// ResultContext executes the initializer respecting context cancellation.
func (o *LazyOnce1[TInput, TOutput]) ResultContext(
	ctx context.Context,
	input TInput,
) result.Result[TOutput] {
	val, fault := o.ValueContext(ctx, input)
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
