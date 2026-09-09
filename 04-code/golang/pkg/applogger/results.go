package applogger

import (
	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/result"
)

// FileSinkSuccess constructs a successful FileSinkResult.
func FileSinkSuccess(sink *FileSink) FileSinkResult {
	return result.WrapSuccess(sink)
}

// FileSinkFailure propagates an existing failed Result into a FileSinkResult.
func FileSinkFailure[U any](failed result.Wrap[U]) FileSinkResult {
	return result.FailureFromWrap[*FileSink](failed)
}

// FileSinkFailureFault creates a failed FileSinkResult from an AppError.
func FileSinkFailureFault(fault *appfault.AppError) FileSinkResult {
	return result.WrapFailure[*FileSink](fault)
}

// RotatingFileSinkSuccess constructs a successful RotatingFileSinkResult.
func RotatingFileSinkSuccess(sink *RotatingFileSink) RotatingFileSinkResult {
	return result.WrapSuccess(sink)
}

// RotatingFileSinkFailure propagates an existing failed Result into a RotatingFileSinkResult.
func RotatingFileSinkFailure[U any](failed result.Wrap[U]) RotatingFileSinkResult {
	return result.FailureFromWrap[*RotatingFileSink](failed)
}

// SQLiteSinkSuccess constructs a successful SQLiteSinkResult.
func SQLiteSinkSuccess(sink *SQLiteSink) SQLiteSinkResult {
	return result.WrapSuccess(sink)
}

// SQLiteSinkFailure propagates an existing failed Result into a SQLiteSinkResult.
func SQLiteSinkFailure[U any](failed result.Wrap[U]) SQLiteSinkResult {
	return result.FailureFromWrap[*SQLiteSink](failed)
}

// ApiSinkSuccess constructs a successful ApiSinkResult.
func ApiSinkSuccess(sink *ApiSink) ApiSinkResult {
	return result.WrapSuccess(sink)
}

// ApiSinkFailure propagates an existing failed Result into an ApiSinkResult.
func ApiSinkFailure[U any](failed result.Wrap[U]) ApiSinkResult {
	return result.FailureFromWrap[*ApiSink](failed)
}

// LoggerSuccess constructs a successful LoggerResult.
func LoggerSuccess(l Logger) LoggerResult {
	return result.WrapSuccess(l)
}

// LoggerFailure propagates an existing failed Result into a LoggerResult.
func LoggerFailure[U any](failed result.Wrap[U]) LoggerResult {
	return result.FailureFromWrap[Logger](failed)
}

// LoggerFailureFault creates a failed LoggerResult from an AppError.
func LoggerFailureFault(fault *appfault.AppError) LoggerResult {
	return result.WrapFailure[Logger](fault)
}

// LogSinkSuccess constructs a successful LogSinkResult.
func LogSinkSuccess(sink LogSink) LogSinkResult {
	return result.WrapSuccess(sink)
}

// LogSinkFailure propagates an existing failed Result into a LogSinkResult.
func LogSinkFailure[U any](failed result.Wrap[U]) LogSinkResult {
	return result.FailureFromWrap[LogSink](failed)
}
