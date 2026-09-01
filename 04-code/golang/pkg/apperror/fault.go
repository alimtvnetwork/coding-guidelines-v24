package apperror

// Fault is the universal structured error type carrying full diagnostics.
type Fault struct {
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

// AppError is retained as a type alias for Fault for backward compatibility.
type AppError = Fault

// GetCode returns the error code string.
func (f *Fault) GetCode() string {
	if f == nil {
		return ""
	}

	return f.Code
}

// GetMessage returns the human-readable diagnostic message.
func (f *Fault) GetMessage() string {
	if f == nil {
		return ""
	}

	return f.Message
}

// GetOp returns the operation label.
func (f *Fault) GetOp() string {
	if f == nil {
		return ""
	}

	return f.Op
}

// StatusCode returns the attached HTTP status code or 0.
func (f *Fault) StatusCode() int {
	if f == nil {
		return 0
	}

	return f.statusCode
}
