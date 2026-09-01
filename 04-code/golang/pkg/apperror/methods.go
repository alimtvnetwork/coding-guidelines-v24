package apperror

// Unwrap returns the underlying root cause error.
func (f *Fault) Unwrap() error {
	if f == nil {
		return nil
	}

	return f.Cause
}

// HasError returns true if the receiver is non-nil.
func (f *Fault) HasError() bool {
	return f != nil
}

// HasNoError returns true if the receiver is nil.
func (f *Fault) HasNoError() bool {
	return f == nil
}

// HasValidError returns true if the receiver is non-nil and has an error code.
func (f *Fault) HasValidError() bool {
	if f == nil {
		return false
	}

	return len(f.Code) > 0
}

// IsValid returns true if the error is populated with an error code.
func (f *Fault) IsValid() bool {
	return f.HasValidError()
}

// IsErrorCode checks if the error code matches the target string.
func (f *Fault) IsErrorCode(code string) bool {
	if f == nil {
		return false
	}

	return f.Code == code
}

// IsCode checks if the error code matches the target ErrorCodeType.
func (f *Fault) IsCode(code ErrorCodeType) bool {
	return f.IsErrorCode(code.String())
}
