package appfault

// Unwrap returns the underlying root cause error.
func (e *AppError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Cause
}

// HasError returns true if the receiver is non-nil.
func (e *AppError) HasError() bool {
	return e != nil
}

// HasNoError returns true if the receiver is nil.
func (e *AppError) HasNoError() bool {
	return e == nil
}

// HasValidError returns true if the receiver is non-nil and has an error code.
func (e *AppError) HasValidError() bool {
	if e == nil {
		return false
	}

	return len(e.Code) > 0
}

// IsValid returns true if the error is populated with an error code.
func (e *AppError) IsValid() bool {
	return e.HasValidError()
}

// IsErrorCode checks if the error code matches the target string.
func (e *AppError) IsErrorCode(code string) bool {
	if e == nil {
		return false
	}

	return e.Code == code
}

// IsCode is an alias for IsErrorCode.
func (e *AppError) IsCode(code string) bool {
	return e.IsErrorCode(code)
}

// IsCodeType checks if the error code matches the target ErrorCodeType.
func (e *AppError) IsCodeType(code ErrorCodeType) bool {
	return e.IsErrorCode(code.String())
}
