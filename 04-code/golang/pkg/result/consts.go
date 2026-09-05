package result

// Display banner constants for result formatting and presentation.
const (
	DefaultSuccessBanner = "✅ [OK]"
	DefaultFailureBanner = "❌ [FAIL]"
)

// ResultFormatter formats a generic Wrap[T] container into a serialized string.
type ResultFormatter[T any] func(r Wrap[T]) string

// ResultPredicate tests a generic Wrap[T] container against a condition.
type ResultPredicate[T any] func(r Wrap[T]) bool

// ResultMapper maps a successful payload of type T to a new type U.
type ResultMapper[T any, U any] func(val T) U
