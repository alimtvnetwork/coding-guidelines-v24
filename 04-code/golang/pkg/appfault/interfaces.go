package appfault

import "io"

type (
	// IsSuccessChecker checks whether an operation succeeded.
	IsSuccessChecker interface {
		IsSuccess() bool
	}

	// IsFailureChecker checks whether an operation failed.
	IsFailureChecker interface {
		IsFailure() bool
	}

	// IsInvalidChecker checks whether an object represents an invalid state.
	IsInvalidChecker interface {
		IsInvalid() bool
	}

	// IsNullChecker checks whether an object represents a null/nil state.
	IsNullChecker interface {
		IsNull() bool
	}

	// IsEmptyChecker checks whether an object represents an empty or zero state.
	IsEmptyChecker interface {
		IsEmpty() bool
	}

	// IsDefinedChecker checks whether an object represents a defined, non-empty, non-null state.
	IsDefinedChecker interface {
		IsDefined() bool
	}

	// DefinableChecker combines both emptiness and definition predicates.
	DefinableChecker interface {
		IsDefinedChecker
		IsEmptyChecker
	}

	// StatusChecker combines success and failure predicates.
	StatusChecker interface {
		IsSuccessChecker
		IsFailureChecker
	}

	// SimpleVerifier combines all core state and status checkers into a single verification contract.
	SimpleVerifier interface {
		IsSuccessChecker
		IsFailureChecker
		IsInvalidChecker
		IsNullChecker
		IsEmptyChecker
		IsDefinedChecker
		DefinableChecker
		StatusChecker
	}

	// SimpleVerifyChecker is an alias for SimpleVerifier adhering to the Checker suffix convention.
	SimpleVerifyChecker = SimpleVerifier

	// SimpleVerifierProvider produces a SimpleVerifier instance.
	SimpleVerifierProvider interface {
		AsSimpleVerifier() SimpleVerifier
	}

	// SimpleVerifyCheckerProvider produces a SimpleVerifier instance adhering to the Checker suffix convention.
	SimpleVerifyCheckerProvider interface {
		AsSimpleVerifyChecker() SimpleVerifier
	}

	// SimpleVerifiable is an alias for SimpleVerifierProvider.
	SimpleVerifiable = SimpleVerifierProvider

	// SimpleVerifyCheckable is an alias for SimpleVerifyCheckerProvider.
	SimpleVerifyCheckable = SimpleVerifyCheckerProvider

	// FaultWriter writes an AppError to an io.Writer output.
	FaultWriter interface {
		WriteFault(w io.Writer, e *AppError) *AppError
	}

	// FaultRenderer formats an AppError into a string.
	FaultRenderer interface {
		RenderFault(e *AppError) string
	}

	// ResultInspecter inspects a result container's value, failure status, and error payload.
	ResultInspecter interface {
		ValueAny() any
		IsFailed() bool
		AppError() *AppError
	}

	// ResultInspector is an alias for ResultInspecter.
	ResultInspector = ResultInspecter

	// ResultUnwrapper is an alias for ResultInspecter.
	ResultUnwrapper = ResultInspecter

	// ResultCarrier is an alias for ResultInspecter.
	ResultCarrier = ResultInspecter
)
