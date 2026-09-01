package appfault

// AppError is the universal structured error type carrying full diagnostics.
type AppError struct {
	Op         string       `json:"Op,omitempty" yaml:"Op,omitempty"`
	Code       string       `json:"Code,omitempty" yaml:"Code,omitempty"`
	Type       ErrorType    `json:"Type,omitempty" yaml:"Type,omitempty"`
	Severity   SeverityType `json:"Severity,omitempty" yaml:"Severity,omitempty"`
	Creator    string       `json:"Creator,omitempty" yaml:"Creator,omitempty"`
	Message    string       `json:"Message,omitempty" yaml:"Message,omitempty"`
	Caller     string       `json:"Caller,omitempty" yaml:"Caller,omitempty"`
	Stack      string       `json:"Stack,omitempty" yaml:"Stack,omitempty"`
	Ctx        ContextMap   `json:"Ctx,omitempty" yaml:"Ctx,omitempty"`
	Cause      error        `json:"Cause,omitempty" yaml:"Cause,omitempty"`
	StatusCode int          `json:"StatusCode,omitempty" yaml:"StatusCode,omitempty"`
}

// Fault is retained as a type alias for AppError for backward compatibility.
type Fault = AppError

// GetCode returns the error code string.
func (e *AppError) GetCode() string {
	if e == nil {
		return ""
	}

	return e.Code
}

// GetMessage returns the human-readable diagnostic message.
func (e *AppError) GetMessage() string {
	if e == nil {
		return ""
	}

	return e.Message
}

// GetOp returns the operation label.
func (e *AppError) GetOp() string {
	if e == nil {
		return ""
	}

	return e.Op
}

// GetStatusCode returns the attached HTTP status code or 0.
func (e *AppError) GetStatusCode() int {
	if e == nil {
		return 0
	}

	return e.StatusCode
}
