package applogger

import "os"

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
func createSinkFromDriver(cfg Config) LogSinkResult {
	switch cfg.Driver {
	case DriverFile:
		res := NewFileSink(cfg.FilePath)
		if res.IsFailed() {
			return LogSinkFailure(res)
		}

		return LogSinkSuccess(res.Data())
	case DriverRotatingFile:
		res := NewRotatingFileSink(cfg.Rotation)
		if res.IsFailed() {
			return LogSinkFailure(res)
		}

		return LogSinkSuccess(res.Data())
	case DriverZap:
		return LogSinkSuccess(NewZapAdapter(cfg.ZapLogger))
	case DriverComposite:
		return LogSinkSuccess(NewCompositeSink(cfg.Sinks...))
	default:
		return LogSinkSuccess(NewConsoleSink(os.Stdout, cfg.IsUseJSON))
	}
}

// New constructs a Logger using the requested configuration and sink driver.
func New(cfg Config) LoggerResult {
	sinkRes := createSinkFromDriver(cfg)
	if sinkRes.IsFailed() {
		return LoggerFailure(sinkRes)
	}

	return LoggerSuccess(&appLogger{
		minLevel: cfg.MinLevel,
		sink:     sinkRes.Data(),
		fields:   nil,
	})
}

// Default returns a standard Console logger at Info level wrapped in a Result.
func Default() LoggerResult {
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
