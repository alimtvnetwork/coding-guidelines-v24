package result

import (
	"coding-guidelines/common/pkg/appfault"
)

type (
	// Wrap wraps a typed value bundled with monadic error state.
	Wrap[T any] = appfault.Result[T]

	// Result is the canonical alias for Wrap[T].
	Result[T any] = Wrap[T]

	// ResultSlice wraps a slice of values bundled with monadic error state.
	ResultSlice[T any] = appfault.ResultSlice[T]

	// ResultMap wraps a key-value map bundled with monadic error state.
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

	// ResultInspecter inspects a result container's value, failure status, and error payload.
	ResultInspecter = appfault.ResultInspecter

	// ResultInspector inspects a result container's value, failure status, and error payload.
	ResultInspector = appfault.ResultInspector

	// ResultUnwrapper is an alias for ResultInspecter.
	ResultUnwrapper = appfault.ResultUnwrapper

	// ResultCarrier is an alias for ResultInspecter.
	ResultCarrier = appfault.ResultCarrier
)
