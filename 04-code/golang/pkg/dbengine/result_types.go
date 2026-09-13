package dbengine

import (
	"coding-guidelines/common/pkg/appfault"
)

// Exact typed result envelopes wrapping appfault.Result[T].
type Uint64Result = appfault.Result[uint64]
type Int64Result = appfault.Result[int64]
type StringResult = appfault.Result[string]
type BoolResult = appfault.Result[bool]
type RowsAffectedResult = appfault.Result[int64]
type EntityResult[T any] = appfault.Result[*T]
type ListResult[T any] = appfault.Result[[]T]

// SuccessUint64 wraps a uint64 value in a successful result envelope.
func SuccessUint64(val uint64) Uint64Result {
	return appfault.NewSuccess(val)
}

// FailureUint64 wraps an AppError in a failed Uint64Result envelope.
func FailureUint64(err *appfault.AppError) Uint64Result {
	return appfault.Fail[uint64](err)
}

// SuccessInt64 wraps an int64 value in a successful result envelope.
func SuccessInt64(val int64) Int64Result {
	return appfault.NewSuccess(val)
}

// FailureInt64 wraps an AppError in a failed Int64Result envelope.
func FailureInt64(err *appfault.AppError) Int64Result {
	return appfault.Fail[int64](err)
}

// SuccessString wraps a string value in a successful result envelope.
func SuccessString(val string) StringResult {
	return appfault.NewSuccess(val)
}

// FailureString wraps an AppError in a failed StringResult envelope.
func FailureString(err *appfault.AppError) StringResult {
	return appfault.Fail[string](err)
}

// SuccessBool wraps a boolean value in a successful result envelope.
func SuccessBool(val bool) BoolResult {
	return appfault.NewSuccess(val)
}

// FailureBool wraps an AppError in a failed BoolResult envelope.
func FailureBool(err *appfault.AppError) BoolResult {
	return appfault.Fail[bool](err)
}

// SuccessRowsAffected wraps the number of rows affected in a successful result envelope.
func SuccessRowsAffected(val int64) RowsAffectedResult {
	return appfault.NewSuccess(val)
}

// FailureRowsAffected wraps an AppError in a failed RowsAffectedResult envelope.
func FailureRowsAffected(err *appfault.AppError) RowsAffectedResult {
	return appfault.Fail[int64](err)
}

// SuccessEntity wraps an entity pointer in a successful result envelope.
func SuccessEntity[T any](val *T) EntityResult[T] {
	return appfault.NewSuccess(val)
}

// FailureEntity wraps an AppError in a failed EntityResult envelope.
func FailureEntity[T any](err *appfault.AppError) EntityResult[T] {
	return appfault.Fail[*T](err)
}

// SuccessList wraps an entity slice in a successful result envelope.
func SuccessList[T any](val []T) ListResult[T] {
	return appfault.NewSuccess(val)
}

// FailureList wraps an AppError in a failed ListResult envelope.
func FailureList[T any](err *appfault.AppError) ListResult[T] {
	return appfault.Fail[[]T](err)
}

// CompiledQuery encapsulates compiled SQL text, bound parameters, and deterministic hash.
type CompiledQuery struct {
	SQL       string
	Args      []any
	QueryHash string
}

// CompiledQueryResult wraps a CompiledQuery in a result envelope.
type CompiledQueryResult = appfault.Result[CompiledQuery]

// SuccessCompiledQuery wraps a CompiledQuery in a successful result envelope.
func SuccessCompiledQuery(val CompiledQuery) CompiledQueryResult {
	return appfault.NewSuccess(val)
}

// FailureCompiledQuery wraps an AppError in a failed CompiledQueryResult envelope.
func FailureCompiledQuery(err *appfault.AppError) CompiledQueryResult {
	return appfault.Fail[CompiledQuery](err)
}
