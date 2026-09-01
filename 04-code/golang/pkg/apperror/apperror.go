package apperror

// Fault is the universal structured error type carrying code, context, and stack trace.
type Fault struct {
	code       ErrorCodeType
	message    string
	cause      error
	stackTrace StackTrace
	statusCode int
	context    map[string]any
}

// AppError is retained as a type alias for Fault to ease backwards compatibility.
type AppError = Fault

// New creates a new Fault with code, message, and automatic stack trace.
func New(code ErrorCodeType, message string) *Fault {
	return &Fault{
		code:       code,
		message:    message,
		stackTrace: captureStackTrace(2),
		context:    make(map[string]any),
	}
}

// NewValidationError creates a Fault specialized for validation failures.
func NewValidationError(message string) *Fault {
	return New(ErrValidation, message)
}

// Wrap wraps a raw error with an ErrorCodeType, message, and captured stack trace.
func Wrap(cause error, code ErrorCodeType, message string) *Fault {
	if cause == nil {
		return nil
	}

	return &Fault{
		code:       code,
		message:    message,
		cause:      cause,
		stackTrace: captureStackTrace(2),
		context:    make(map[string]any),
	}
}

// Code returns the standardized ErrorCodeType.
func (f *Fault) Code() ErrorCodeType {
	return f.code
}

// Message returns the human-readable error message.
func (f *Fault) Message() string {
	return f.message
}

// Cause returns the underlying wrapped error or nil.
func (f *Fault) Cause() error {
	return f.cause
}

// StackTrace returns the structured call stack frames.
func (f *Fault) StackTrace() StackTrace {
	return f.stackTrace
}

// StatusCode returns the attached HTTP status code or 0.
func (f *Fault) StatusCode() int {
	return f.statusCode
}

// IsCode returns true if the error code matches target code.
func (f *Fault) IsCode(target ErrorCodeType) bool {
	return f.code == target
}
