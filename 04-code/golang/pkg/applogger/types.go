package applogger

import "coding-guidelines/common/pkg/result"

type (
	// Result and Wrap aliases for generic usage in applogger
	Result[T any] = result.Wrap[T]
	Wrap[T any]   = result.Wrap[T]

	// FileSink direct result and wrap types
	FileSinkResult = result.Wrap[*FileSink]
	FileSinkWrap   = FileSinkResult

	// RotatingFileSink direct result and wrap types
	RotatingFileSinkResult = result.Wrap[*RotatingFileSink]
	RotatingFileSinkWrap   = RotatingFileSinkResult

	// SQLiteSink direct result and wrap types
	SQLiteSinkResult = result.Wrap[*SQLiteSink]
	SQLiteSinkWrap   = SQLiteSinkResult

	// ApiSink and ApiManager direct result and wrap types
	ApiSinkResult    = result.Wrap[*ApiSink]
	ApiSinkWrap      = ApiSinkResult
	ApiManagerResult = result.Wrap[*ApiManager]
	ApiManagerWrap   = ApiManagerResult

	// Logger and LogSink direct result and wrap types
	LoggerResult  = result.Wrap[Logger]
	LoggerWrap    = LoggerResult
	LogSinkResult = result.Wrap[LogSink]
	LogSinkWrap   = LogSinkResult
)
