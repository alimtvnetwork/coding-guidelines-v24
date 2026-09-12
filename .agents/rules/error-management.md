# Rule: Error Management Architecture

1. **Go Error Type:** All Go packages returning structured error metadata MUST use `*appfault.AppError` from `04-code/golang/pkg/appfault`.
2. **Never Swallow Errors:** Every catch block or error branch must log the operation name, key inputs, and preserve the original error cause.
3. **Universal Envelope:** API endpoints and service interfaces return standard envelopes (`{ data, errors[], meta }`).
4. **Single Result Container Return Types:** In Go, functions returning collections/values alongside errors MUST return `appfault.ResultMap[K, V]`, `appfault.ResultSlice[T]`, or `appfault.Result[T]` from `04-code/golang/pkg/appfault` instead of multi-value error tuples. Functions with side-effects only return `*appfault.AppError`.
5. **Pointer-Attached Null Safety & Core Predicates:** All Result inspection methods MUST attach to pointer receivers (`(r *Result[T])`, `(rs *ResultSlice[T])`, `(rm *ResultMap[K, V])`) with line-1 `if r == nil` guards returning safe defaults. Enforce core predicates `IsCountOtherThan(N)`, `IsEmpty()`, `HasRecord()`, `IsDefined()`.
