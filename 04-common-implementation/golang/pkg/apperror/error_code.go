package apperror

// ErrorCodeType represents standardized error codes across the platform.
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

// String returns the string representation of the ErrorCodeType.
func (c ErrorCodeType) String() string {
	return string(c)
}

// IsValid returns true when the error code string is non-empty.
func (c ErrorCodeType) IsValid() bool {
	return len(c) > 0
}
