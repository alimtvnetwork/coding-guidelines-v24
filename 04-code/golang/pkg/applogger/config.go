package applogger

import (
	"os"

	"coding-guidelines/common/pkg/result"
)

// Config configures the logger instance.
type Config struct {
	MinLevel     LogLevel
	Driver       DriverType
	FilePath     string
	ZapLogger    ZapLoggerInterface
	Sinks        []LogSink
	Rotation     RotationConfig
	IsStackTrace bool
	IsUseJSON    bool
}

// createSinkFromDriver instantiates the requested driver sink.
func createSinkFromDriver(cfg Config) result.Wrap[LogSink] {
	switch cfg.Driver {
	case DriverFile:
		res := NewFileSink(cfg.FilePath)
		if res.IsFailed() {
			return result.FailureFromWrap[LogSink](res)
		}

		return result.WrapSuccess[LogSink](res.Data())
	case DriverRotatingFile:
		res := NewRotatingFileSink(cfg.Rotation)
		if res.IsFailed() {
			return result.FailureFromWrap[LogSink](res)
		}

		return result.WrapSuccess[LogSink](res.Data())
	case DriverZap:
		return result.WrapSuccess[LogSink](NewZapAdapter(cfg.ZapLogger))
	case DriverComposite:
		return result.WrapSuccess[LogSink](NewCompositeSink(cfg.Sinks...))
	default:
		return result.WrapSuccess[LogSink](NewConsoleSink(os.Stdout, cfg.IsUseJSON))
	}
}

// New constructs a Logger using the requested configuration and sink driver.
func New(cfg Config) result.Wrap[Logger] {
	sinkRes := createSinkFromDriver(cfg)
	if sinkRes.IsFailed() {
		return result.FailureFromWrap[Logger](sinkRes)
	}

	return result.WrapSuccess[Logger](&appLogger{
		minLevel: cfg.MinLevel,
		sink:     sinkRes.Data(),
		fields:   nil,
	})
}

// Default returns a standard Console logger at Info level wrapped in a Result.
func Default() result.Wrap[Logger] {
	return New(Config{
		MinLevel:  LevelInfo,
		Driver:    DriverConsole,
		IsUseJSON: false,
	})
}

// MustDefault returns a standard Console logger at Info level.
func MustDefault() Logger {
	return Default().Data()
}
