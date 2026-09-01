package apperror

// AppError is the universal structured error type carrying code, context, and stack trace.
type AppError struct {
	code       ErrorCodeType
	message    string
	cause      error
	stackTrace StackTrace
	statusCode int
	context    map[string]any
}

// New creates a new AppError with code, message, and automatic stack trace.
func New(code ErrorCodeType, message string) *AppError {
	return &AppError{
		code:       code,
		message:    message,
		stackTrace: captureStackTrace(2),
		context:    make(map[string]any),
	}
}

// NewValidationError creates an AppError specialized for validation failures.
func NewValidationError(message string) *AppError {
	return New(ErrValidation, message)
}

// Wrap wraps a raw error with an ErrorCodeType, message, and captured stack trace.
func Wrap(cause error, code ErrorCodeType, message string) *AppError {
	if cause == nil {
		return nil
	}

	return &AppError{
		code:       code,
		message:    message,
		cause:      cause,
		stackTrace: captureStackTrace(2),
		context:    make(map[string]any),
	}
}

// Code returns the standardized ErrorCodeType.
func (e *AppError) Code() ErrorCodeType {
	return e.code
}

// Message returns the human-readable error message.
func (e *AppError) Message() string {
	return e.message
}

// Cause returns the underlying wrapped error or nil.
func (e *AppError) Cause() error {
	return e.cause
}

// StackTrace returns the structured call stack frames.
func (e *AppError) StackTrace() StackTrace {
	return e.stackTrace
}

// StatusCode returns the attached HTTP status code or 0.
func (e *AppError) StatusCode() int {
	return e.statusCode
}

// IsCode returns true if the error code matches target code.
func (e *AppError) IsCode(target ErrorCodeType) bool {
	return e.code == target
}
