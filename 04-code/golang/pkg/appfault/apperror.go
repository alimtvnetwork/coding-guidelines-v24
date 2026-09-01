package appfault

import "coding-guidelines/common/pkg/errtype"

// AppError is the universal structured error type carrying full diagnostics.
type AppError struct {
	Type       errtype.Variation `json:"Type,omitempty" yaml:"Type,omitempty"`
	Message    string            `json:"Message,omitempty" yaml:"Message,omitempty"`
	Caller     string            `json:"Caller,omitempty" yaml:"Caller,omitempty"`
	Stack      string            `json:"Stack,omitempty" yaml:"Stack,omitempty"`
	Ctx        ContextMap        `json:"Ctx,omitempty" yaml:"Ctx,omitempty"`
	Cause      error             `json:"Cause,omitempty" yaml:"Cause,omitempty"`
	StatusCode int               `json:"StatusCode,omitempty" yaml:"StatusCode,omitempty"`
}

// Fault is retained as a type alias for AppError for backward compatibility.
type Fault = AppError

// GetMessage returns the human-readable diagnostic message.
func (e *AppError) GetMessage() string {
	if e == nil {
		return ""
	}

	return e.Message
}

// GetStatusCode returns the attached HTTP status code or 0.
func (e *AppError) GetStatusCode() int {
	if e == nil {
		return 0
	}

	return e.StatusCode
}

// GetType returns the error type variation.
func (e *AppError) GetType() errtype.Variation {
	if e == nil {
		return errtype.None
	}

	return e.Type
}

// HasError returns true if the AppError exists and is not errtype.None.
func (e *AppError) HasError() bool {
	if e == nil {
		return false
	}

	return e.Type.HasError()
}

// HasNullError returns true if e is nil or represents no error.
func (e *AppError) HasNullError() bool {
	if e == nil {
		return true
	}

	return e.Type.IsNone()
}

// IsNull returns true if e is nil.
func (e *AppError) IsNull() bool {
	return e == nil
}

// IsEmpty returns true if e is nil or Type is None.
func (e *AppError) IsEmpty() bool {
	return e.HasNullError()
}

// IsSuccess returns true if no error is present (e is nil or Type is None).
func (e *AppError) IsSuccess() bool {
	return e.HasNullError()
}
