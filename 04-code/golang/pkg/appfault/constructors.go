package appfault

// NewSimple creates an error with automatic caller and stack trace capture.
func NewSimple(op, code string) *AppError {
	return New(op, code, nil)
}

// New creates a standard AppError with caller, stack trace, and context map.
func New(op, code string, ctx map[string]any) *AppError {
	return NewWithDetails(op, code, "", "", ErrorTypeExecution, SeverityError, ctx)
}

// NewValidationError creates a specialized validation AppError.
func NewValidationError(msg string) *AppError {
	return NewWithDetails("validation", ErrValidation.String(), msg, "", ErrorTypeValidation, SeverityError, nil)
}

// createAppErrorInstance sets up the base AppError struct.
func createAppErrorInstance(op, code, msg, creator string, errType ErrorType, sev SeverityType) *AppError {
	trace := captureStackTrace(3)

	return &AppError{
		Op:       op,
		Code:     code,
		Type:     errType,
		Severity: sev,
		Creator:  creator,
		Message:  msg,
		Caller:   trace.CallerLine(),
		Stack:    trace.String(),
	}
}

// NewWithDetails provides full-fidelity AppError construction.
func NewWithDetails(op, code, msg, creator string, errType ErrorType, sev SeverityType, ctx map[string]any) *AppError {
	e := createAppErrorInstance(op, code, msg, creator, errType, sev)
	e.Ctx = ensureContext(ctx)

	return e
}

// WrapSimple wraps an existing error with default code and operation label.
func WrapSimple(err error, op string) *AppError {
	return Wrap(err, op, nil)
}

// Wrap wraps an existing error with operation label and context map.
func Wrap(err error, op string, ctx map[string]any) *AppError {
	if err == nil {
		return nil
	}

	return WrapWithDetails(err, op, ErrUnknown.String(), err.Error(), "", ErrorTypeExecution, SeverityError, ctx)
}

// WrapWithDetails wraps an error preserving the underlying root cause.
func WrapWithDetails(err error, op, code, msg, creator string, errType ErrorType, sev SeverityType, ctx map[string]any) *AppError {
	if err == nil {
		return nil
	}

	e := NewWithDetails(op, code, msg, creator, errType, sev, ctx)
	e.Cause = err

	return e
}

// ensureContext safely returns a non-nil context map.
func ensureContext(ctx map[string]any) map[string]any {
	if ctx == nil {
		return make(map[string]any)
	}

	return ctx
}
