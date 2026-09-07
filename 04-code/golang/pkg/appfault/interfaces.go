package appfault

import "io"

type (
	// IsDefiner checks whether an object represents a defined, non-empty, non-null state.
	IsDefiner interface {
		IsDefined() bool
	}

	// IsEmptyer checks whether an object represents an empty or zero state.
	IsEmptyer interface {
		IsEmpty() bool
	}

	// DefinableChecker combines both emptiness and definition predicates.
	DefinableChecker interface {
		IsDefiner
		IsEmptyer
	}

	// FaultWriter writes an AppError to an io.Writer output.
	FaultWriter interface {
		WriteFault(w io.Writer, e *AppError) error
	}

	// FaultRenderer formats an AppError into a string.
	FaultRenderer interface {
		RenderFault(e *AppError) string
	}
)
