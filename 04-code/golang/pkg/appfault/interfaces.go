package appfault

import "io"

type (
	// IsDefinedChecker checks whether an object represents a defined, non-empty, non-null state.
	IsDefinedChecker interface {
		IsDefined() bool
	}

	// IsEmptyChecker checks whether an object represents an empty or zero state.
	IsEmptyChecker interface {
		IsEmpty() bool
	}

	// DefinableChecker combines both emptiness and definition predicates.
	DefinableChecker interface {
		IsDefinedChecker
		IsEmptyChecker
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
