package logger

import (
	"coding-guidelines/common/pkg/enum/logleveltype"
)

type LogLevel = logleveltype.Variant

func ParseLogLevel(s string) LogLevel {
	return logleveltype.ParseOrUnknown(s)
}
