package appfault

// AppError is the universal structured error type carrying full diagnostics.
type AppError struct {
	Op         string         `json:"op,omitempty"`
	Code       string         `json:"code,omitempty"`
	Type       ErrorType      `json:"type,omitempty"`
	Severity   SeverityType   `json:"severity,omitempty"`
	Creator    string         `json:"creator,omitempty"`
	Message    string         `json:"message,omitempty"`
	Caller     string         `json:"caller,omitempty"`
	Stack      string         `json:"stack,omitempty"`
	Ctx        map[string]any `json:"ctx,omitempty"`
	Cause      error          `json:"cause,omitempty"`
	statusCode int
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

// StatusCode returns the attached HTTP status code or 0.
func (e *AppError) StatusCode() int {
	if e == nil {
		return 0
	}

	return e.statusCode
}
