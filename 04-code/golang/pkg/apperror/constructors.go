package apperror

// NewSimple creates an error with automatic caller and stack trace capture.
func NewSimple(op, code string) *Fault {
	return New(op, code, nil)
}

// New creates a standard Fault with caller, stack trace, and context map.
func New(op, code string, ctx map[string]any) *Fault {
	return NewWithDetails(op, code, "", "", ErrorTypeExecution, SeverityError, ctx)
}

// NewValidationError creates a specialized validation Fault.
func NewValidationError(msg string) *Fault {
	return NewWithDetails("validation", ErrValidation.String(), msg, "", ErrorTypeValidation, SeverityError, nil)
}

// createFaultInstance sets up the base Fault struct.
func createFaultInstance(op, code, msg, creator string, errType ErrorType, sev SeverityType) *Fault {
	trace := captureStackTrace(3)

	return &Fault{
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

// NewWithDetails provides full-fidelity Fault construction.
func NewWithDetails(op, code, msg, creator string, errType ErrorType, sev SeverityType, ctx map[string]any) *Fault {
	f := createFaultInstance(op, code, msg, creator, errType, sev)
	f.Ctx = ensureContext(ctx)

	return f
}

// WrapSimple wraps an existing error with default code and operation label.
func WrapSimple(err error, op string) *Fault {
	return Wrap(err, op, nil)
}

// Wrap wraps an existing error with operation label and context map.
func Wrap(err error, op string, ctx map[string]any) *Fault {
	if err == nil {
		return nil
	}

	return WrapWithDetails(err, op, ErrUnknown.String(), err.Error(), "", ErrorTypeExecution, SeverityError, ctx)
}

// WrapWithDetails wraps an error preserving the underlying root cause.
func WrapWithDetails(err error, op, code, msg, creator string, errType ErrorType, sev SeverityType, ctx map[string]any) *Fault {
	if err == nil {
		return nil
	}

	f := NewWithDetails(op, code, msg, creator, errType, sev, ctx)
	f.Cause = err

	return f
}

// ensureContext safely returns a non-nil context map.
func ensureContext(ctx map[string]any) map[string]any {
	if ctx == nil {
		return make(map[string]any)
	}

	return ctx
}
