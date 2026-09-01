package apperror

// ErrorType categorizes error domains.
type ErrorType string

const (
	ErrorTypeValidation   ErrorType = "VALIDATION"
	ErrorTypePrecondition ErrorType = "PRECONDITION"
	ErrorTypeNotFound     ErrorType = "NOT_FOUND"
	ErrorTypeExecution    ErrorType = "EXECUTION"
	ErrorTypeAbort        ErrorType = "ABORT"
	ErrorTypeInternal     ErrorType = "INTERNAL"
)

// SeverityType indicates the severity level of an error.
type SeverityType string

const (
	SeverityInfo  SeverityType = "INFO"
	SeverityWarn  SeverityType = "WARN"
	SeverityError SeverityType = "ERROR"
	SeverityFatal SeverityType = "FATAL"
)

// ErrorCodeType represents standard system error codes.
type ErrorCodeType string

const (
	ErrUnknown           ErrorCodeType = "E9999"
	ErrValidation        ErrorCodeType = "E1001"
	ErrConfiguration     ErrorCodeType = "E1002"
	ErrDatabaseQuery     ErrorCodeType = "E2001"
	ErrDatabaseExec      ErrorCodeType = "E2002"
	ErrDatabaseNotFound  ErrorCodeType = "E2004"
	ErrRemoteRequest     ErrorCodeType = "E3001"
	ErrRemoteAuth        ErrorCodeType = "E3002"
	ErrRemoteServerError ErrorCodeType = "E3003"
	ErrFileNotFound      ErrorCodeType = "E4001"
	ErrFileRead          ErrorCodeType = "E4002"
	ErrFileWrite         ErrorCodeType = "E4003"
	ErrSyncFailed        ErrorCodeType = "E5001"
	ErrSyncConflict      ErrorCodeType = "E5002"
	ErrBackupFailed      ErrorCodeType = "E6001"
	ErrGitExec           ErrorCodeType = "E7001"
)

// String returns the string representation of ErrorCodeType.
func (c ErrorCodeType) String() string {
	return string(c)
}

// IsValid returns true if the error code is non-empty.
func (c ErrorCodeType) IsValid() bool {
	return len(c) > 0
}
