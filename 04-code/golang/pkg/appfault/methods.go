package appfault

import "coding-guidelines/common/pkg/errtype"

// Unwrap returns the underlying root cause error.
func (e *AppError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Cause
}

// HasNoError returns true if the receiver is nil or Type is None.
func (e *AppError) HasNoError() bool {
	if e == nil {
		return true
	}

	return e.Type.IsNone()
}

// HasValidError returns true if the receiver is non-nil and has a valid error type.
func (e *AppError) HasValidError() bool {
	if e == nil {
		return false
	}

	return e.Type.HasError()
}

// IsValid returns true if no error is present (valid / healthy state).
func (e *AppError) IsValid() bool {
	return e.IsSuccess()
}

// IsInvalid returns true if an active error is present.
func (e *AppError) IsInvalid() bool {
	return e.HasError()
}

// IsFailed returns true if an active error is present.
func (e *AppError) IsFailed() bool {
	return e.HasError()
}

// Is checks if the error type matches the target Variation.
func (e *AppError) Is(target errtype.Variation) bool {
	if e == nil {
		return target == errtype.None
	}

	return e.Type == target
}

// IsType is an alias for Is.
func (e *AppError) IsType(target errtype.Variation) bool {
	return e.Is(target)
}
