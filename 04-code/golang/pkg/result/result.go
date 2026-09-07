package result

import (
	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

type (
	Wrap[T any] = appfault.Result[T]

	Result[T any] = Wrap[T]

	ResultSlice[T any] = appfault.ResultSlice[T]

	ResultMap[K comparable, V any] = appfault.ResultMap[K, V]

	// SimpleVerifier combines all core state and status checkers into a single verification contract.
	SimpleVerifier = appfault.SimpleVerifier

	// SimpleVerifyChecker is an alias for SimpleVerifier adhering to the Checker convention.
	SimpleVerifyChecker = appfault.SimpleVerifyChecker

	// SimpleVerifierProvider provides AsSimpleVerifier.
	SimpleVerifierProvider = appfault.SimpleVerifierProvider

	// SimpleVerifyCheckerProvider provides AsSimpleVerifyChecker.
	SimpleVerifyCheckerProvider = appfault.SimpleVerifyCheckerProvider

	// SimpleVerifiable provides AsSimpleVerifier.
	SimpleVerifiable = appfault.SimpleVerifiable

	// SimpleVerifyCheckable provides AsSimpleVerifyChecker.
	SimpleVerifyCheckable = appfault.SimpleVerifyCheckable
)

func WrapSuccess[T any](data T) Wrap[T] {
	return appfault.NewSuccess(data)
}

func WrapFailure[T any](err *appfault.AppError) Wrap[T] {
	return appfault.FailureResult[T](err)
}

func Success[T any](data T) Wrap[T] {
	return appfault.NewSuccess(data)
}

func Failure[T any](err *appfault.AppError) Wrap[T] {
	return appfault.FailureResult[T](err)
}

func WrapFailureFromError[T any](err *appfault.AppError) Wrap[T] {
	return appfault.FailureResult[T](err)
}

func WrapFailureWithId[T any](errType errtype.Variation, msg string) Wrap[T] {
	return appfault.NewFailureWithId[T](errType, msg)
}

func WrapFailureWithCause[T any](errType errtype.Variation, cause error, msg string) Wrap[T] {
	return appfault.NewFailureWithCause[T](errType, cause, msg)
}

func WrapFailureFromWrap[T any, U any](failed Wrap[U]) Wrap[T] {
	return appfault.FailureFromWrap[T](failed)
}

func FailureFromWrap[T any, U any](failed Wrap[U]) Wrap[T] {
	return appfault.FailureFromWrap[T](failed)
}

func FailureWithId[T any](errType errtype.Variation, msg string) Wrap[T] {
	return appfault.NewFailureWithId[T](errType, msg)
}

func FailureWithCause[T any](errType errtype.Variation, cause error, msg string) Wrap[T] {
	return appfault.NewFailureWithCause[T](errType, cause, msg)
}

// WrapFailurePath wraps a failure with path context and root cause error.
func WrapFailurePath[T any](variation errtype.Variation, cause error, path string, msg string) Wrap[T] {
	return appfault.NewFailureWithPath[T](variation, cause, path, msg)
}

// WrapFailureFile wraps a failure with file context and root cause error.
func WrapFailureFile[T any](variation errtype.Variation, cause error, path string, msg string) Wrap[T] {
	return appfault.NewFailureWithFile[T](variation, cause, path, msg)
}

// WrapFailureVar wraps a failure with a single variable context and root cause error.
func WrapFailureVar[T any](variation errtype.Variation, cause error, key string, val any, msg string) Wrap[T] {
	return appfault.NewFailureWithVar[T](variation, cause, key, val, msg)
}

// FailurePath creates a failed Result with path context.
func FailurePath[T any](variation errtype.Variation, path string, msg string) Wrap[T] {
	return appfault.FailurePath[T](variation, path, msg)
}

// FailureFile creates a failed Result with file context.
func FailureFile[T any](variation errtype.Variation, path string, msg string) Wrap[T] {
	return appfault.FailureFile[T](variation, path, msg)
}

// FailureVar creates a failed Result with variable context.
func FailureVar[T any](variation errtype.Variation, key string, val any, msg string) Wrap[T] {
	return appfault.FailureVar[T](variation, key, val, msg)
}

func SuccessResult[T any](val T) Result[T] {
	return appfault.SuccessResult(val)
}

func NewSuccess[T any](data T) Result[T] {
	return appfault.NewSuccess(data)
}

func FailureResult[T any](err *appfault.AppError) Result[T] {
	return appfault.FailureResult[T](err)
}

func NewFailure[T any](errType errtype.Variation, cause error) Result[T] {
	return appfault.NewFailure[T](errType, cause)
}

func NewFailureWithType[T any](errType errtype.Variation, msg string, caller string) Result[T] {
	return appfault.NewFailureWithType[T](errType, msg, caller)
}

// NewFailureWithPath creates a failed Result with path context and root cause error.
func NewFailureWithPath[T any](variation errtype.Variation, cause error, path string, msg string) Result[T] {
	return appfault.NewFailureWithPath[T](variation, cause, path, msg)
}

// NewFailureWithFile creates a failed Result with file context and root cause error.
func NewFailureWithFile[T any](variation errtype.Variation, cause error, path string, msg string) Result[T] {
	return appfault.NewFailureWithFile[T](variation, cause, path, msg)
}

// NewFailureWithVar creates a failed Result with variable context and root cause error.
func NewFailureWithVar[T any](variation errtype.Variation, cause error, key string, val any, msg string) Result[T] {
	return appfault.NewFailureWithVar[T](variation, cause, key, val, msg)
}

// CastTo safely casts a generic payload into target type T.
func CastTo[T any](source any) (T, *appfault.AppError) {
	return appfault.CastTo[T](source)
}

// ReflectTo dynamically casts a generic payload into target type T.
func ReflectTo[T any](source any) (T, *appfault.AppError) {
	return appfault.ReflectTo[T](source)
}

// CastResult converts Result[T] to Result[U] via reflection.
func CastResult[T any, U any](r Result[T]) Result[U] {
	return appfault.CastResult[T, U](r)
}

// CastContextPayload retrieves a key from an AppError context and casts to T.
func CastContextPayload[T any](e *appfault.AppError, key string) (T, *appfault.AppError) {
	return appfault.CastContextPayload[T](e, key)
}

// ResultToBytes serializes result data or fault using typecast.ToBytes.
func ResultToBytes[T any](r Result[T]) Result[[]byte] {
	return appfault.ResultToBytes[T](r)
}

// ResultToJson serializes result data or fault using typecast.ToJson.
func ResultToJson[T any](r Result[T]) Result[[]byte] {
	return appfault.ResultToJson[T](r)
}

// ResultToJSON is an alias for ResultToJson.
func ResultToJSON[T any](r Result[T]) Result[[]byte] {
	return appfault.ResultToJson[T](r)
}

// OkSlice creates a successful ResultSlice.
func OkSlice[T any](items []T) ResultSlice[T] {
	return appfault.OkSlice(items)
}

// FailSlice creates a failed ResultSlice from an AppError.
func FailSlice[T any](err *appfault.AppError) ResultSlice[T] {
	return appfault.FailSlice[T](err)
}

// OkMap creates a successful ResultMap.
func OkMap[K comparable, V any](data map[K]V) ResultMap[K, V] {
	return appfault.OkMap(data)
}

// FailMap creates a failed ResultMap from an AppError.
func FailMap[K comparable, V any](err *appfault.AppError) ResultMap[K, V] {
	return appfault.FailMap[K, V](err)
}
