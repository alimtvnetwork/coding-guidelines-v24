package appfault

import "coding-guidelines/common/pkg/errtype"

// New creates an AppError for a given error type variation and message.
// If errType is errtype.None, it returns nil (no error allocated).
func New(errType errtype.Variation, message string) *AppError {
	if errType == errtype.None {
		return nil
	}

	return NewWithContext(errType, message, nil)
}

// NewValidationError creates a validation error.
func NewValidationError(message string) *AppError {
	return New(errtype.Validation, message)
}

// createAppErrorInstance constructs the AppError capturing stack trace.
func createAppErrorInstance(errType errtype.Variation, message string) *AppError {
	trace := CaptureStackTrace(3)

	return &AppError{
		Type:    errType,
		Message: message,
		Caller:  trace.CallerLine(),
		Stack:   trace.String(),
	}
}

// NewWithContext constructs an AppError with an initial context map.
func NewWithContext(errType errtype.Variation, message string, ctx map[string]any) *AppError {
	if errType == errtype.None {
		return nil
	}

	e := createAppErrorInstance(errType, message)
	e.Ctx = ensureContextMap(ctx)

	return e
}

// Wrap wraps an existing error with an error type variation and custom message.
// If err is nil, it returns nil (no allocation).
func Wrap(err error, errType errtype.Variation, message string) *AppError {
	if err == nil {
		return nil
	}

	e := New(errType, message)
	if e == nil {
		return nil
	}

	e.Cause = err

	return e
}

// WrapSimple wraps a raw error into an Execution error type.
func WrapSimple(err error) *AppError {
	if err == nil {
		return nil
	}

	return Wrap(err, errtype.Execution, err.Error())
}

// ensureContextMap safely converts a map[string]any to ContextMap.
func ensureContextMap(ctx map[string]any) ContextMap {
	if ctx == nil {
		return NewContextMap()
	}

	cm := NewContextMapWithCapacity(len(ctx))
	for k, v := range ctx {
		cm[k] = v
	}

	return cm
}
