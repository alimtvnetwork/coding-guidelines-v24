# Subtask 04: Cache Invalidation & Context Cancellation in lazyonce

## 1. Goal
Add thread-safe cache invalidation (`Reset()`) and context cancellation support (`ValueContext(ctx)`) to `lazyonce`.

**Status:** ✅ Completed

## 2. Target Files
- `04-code/golang/pkg/lazyonce/lazyonce.go`
- `04-code/golang/pkg/lazyonce/lazyonce1.go`
- `04-code/golang/pkg/lazyonce/lazyonce2.go`
- `04-code/golang/pkg/lazyonce/lazyonce_test.go`

## 3. Detailed Specifications
1. **Reset Method**:
   - `func (l *LazyOnce[T]) Reset()`
   - `func (l *LazyOnce1[TInput, TOutput]) Reset()`
   - `func (l *LazyOnce2[T1, T2, TOutput]) Reset()`
   - Acquires mutex, resets `isInitialized = false`, clears cached value and fault.
2. **Context-Aware Evaluation**:
   - `func (l *LazyOnce[T]) ValueContext(ctx context.Context) (T, *appfault.AppError)`
   - Checks `ctx.Err()`; if context is cancelled, returns `appfault.Wrap(errtype.Timeout, ctx.Err(), "context cancelled during lazyonce evaluation")`.
   - Runs initializer while respecting context completion.
3. **Coding Guidelines Compliance**:
   - Functions <= 15 lines.
