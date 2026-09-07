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

	// FaultWriter writes an AppError to an io.Writer output.
	FaultWriter interface {
		WriteFault(w io.Writer, e *AppError) *AppError
	}

	// FaultRenderer formats an AppError into a string.
	FaultRenderer interface {
		RenderFault(e *AppError) string
	}
)
