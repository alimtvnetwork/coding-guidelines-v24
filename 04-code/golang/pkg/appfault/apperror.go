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
// Note: If e is nil, it returns false (indicating no error / success).
func (e *AppError) HasError() bool {
	if e == nil {
		return false
	}

	return e.Type.HasError()
}

// IsSuccess returns true if no error is present (e is nil or Type is None).
// Note: When e is nil or Type == errtype.None, this represents a successful operation.
func (e *AppError) IsSuccess() bool {
	return !e.HasError()
}
