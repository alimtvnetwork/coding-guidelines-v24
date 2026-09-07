package appfault

import "coding-guidelines/common/pkg/errtype"

// DefaultFrameSkip defines standard runtime stack frame skip depth.
const DefaultFrameSkip = 3

// calculateFrameSkip calculates caller frame offset from default and optional user skip.
func calculateFrameSkip(skipFrames ...int) int {
	if len(skipFrames) == 0 {
		return DefaultFrameSkip
	}

	return DefaultFrameSkip + skipFrames[0]
}

// createAppErrorInstance constructs the AppError capturing stack trace at the calculated skip.
func createAppErrorInstance(errType errtype.Variation, message string, skip int) *AppError {
	return &AppError{
		errType: errType,
		message: message,
		stack:   CaptureStackTrace(skip),
		ctx:     nil,
	}
}

// New creates an AppError for a given error type variation and message.
// If errType is errtype.None, it returns nil (no error allocated).
func New(errType errtype.Variation, message string, skipFrames ...int) *AppError {
	if errType == errtype.None {
		return nil
	}

	return createAppErrorInstance(errType, message, calculateFrameSkip(skipFrames...))
}

// NewType creates an AppError using default type name as message.
func NewType(errType errtype.Variation, skipFrames ...int) *AppError {
	if errType == errtype.None {
		return nil
	}

	return createAppErrorInstance(errType, errType.Name(), calculateFrameSkip(skipFrames...))
}

// NewWithContext constructs an AppError with an initial context map.
func NewWithContext(errType errtype.Variation, message string, ctx map[string]any, skipFrames ...int) *AppError {
	if errType == errtype.None {
		return nil
	}

	e := createAppErrorInstance(errType, message, calculateFrameSkip(skipFrames...))
	if ctx != nil && len(ctx) > 0 {
		e.ctx = ensureContextMap(ctx)
	}

	return e
}

// Wrap wraps an existing cause with an explicit errtype and custom message.
// If cause is nil or errType is None, it returns nil (no allocation).
func Wrap(errType errtype.Variation, cause error, message string, skipFrames ...int) *AppError {
	if cause == nil || errType == errtype.None {
		return nil
	}

	e := createAppErrorInstance(errType, message, calculateFrameSkip(skipFrames...))
	e.cause = cause

	return e
}

// WrapType wraps an existing cause using cause.Error() as message.
func WrapType(errType errtype.Variation, cause error, skipFrames ...int) *AppError {
	if cause == nil || errType == errtype.None {
		return nil
	}

	e := createAppErrorInstance(errType, cause.Error(), calculateFrameSkip(skipFrames...))
	e.cause = cause

	return e
}

// WrapFile wraps an existing cause with file path context.
func WrapFile(variation errtype.Variation, cause error, path string, msg string, skipFrames ...int) *AppError {
	if cause == nil || variation == errtype.None {
		return nil
	}

	e := createAppErrorInstance(variation, msg, calculateFrameSkip(skipFrames...))
	e.cause = cause
	e.ctx = NewContextMapWithCapacity(1).Set("Path", path)

	return e
}

// NewFile creates an AppError with file path context.
func NewFile(variation errtype.Variation, path string, msg string, skipFrames ...int) *AppError {
	if variation == errtype.None {
		return nil
	}

	e := createAppErrorInstance(variation, msg, calculateFrameSkip(skipFrames...))
	e.ctx = NewContextMapWithCapacity(1).Set("Path", path)

	return e
}

// WrapPath wraps an existing cause with path context.
func WrapPath(variation errtype.Variation, cause error, path string, msg string, skipFrames ...int) *AppError {
	if cause == nil || variation == errtype.None {
		return nil
	}

	e := createAppErrorInstance(variation, msg, calculateFrameSkip(skipFrames...))
	e.cause = cause
	e.ctx = NewContextMapWithCapacity(1).Set("Path", path)

	return e
}

// NewPath creates an AppError with path context.
func NewPath(variation errtype.Variation, path string, msg string, skipFrames ...int) *AppError {
	if variation == errtype.None {
		return nil
	}

	e := createAppErrorInstance(variation, msg, calculateFrameSkip(skipFrames...))
	e.ctx = NewContextMapWithCapacity(1).Set("Path", path)

	return e
}

// WrapVar wraps an existing cause with variable context.
func WrapVar(variation errtype.Variation, cause error, key string, val any, msg string, skipFrames ...int) *AppError {
	if cause == nil || variation == errtype.None {
		return nil
	}

	e := createAppErrorInstance(variation, msg, calculateFrameSkip(skipFrames...))
	e.cause = cause
	e.ctx = NewContextMapWithCapacity(1).Set(key, val)

	return e
}

// NewVar creates an AppError with variable context.
func NewVar(variation errtype.Variation, key string, val any, msg string, skipFrames ...int) *AppError {
	if variation == errtype.None {
		return nil
	}

	e := createAppErrorInstance(variation, msg, calculateFrameSkip(skipFrames...))
	e.ctx = NewContextMapWithCapacity(1).Set(key, val)

	return e
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
