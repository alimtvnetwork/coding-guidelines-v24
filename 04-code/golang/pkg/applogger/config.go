package applogger

import "os"

// Config configures the logger instance.
type Config struct {
	MinLevel     LogLevel
	Driver       DriverType
	FilePath     string
	ZapLogger    ZapLoggerInterface
	Sinks        []LogSink
	IsStackTrace bool
	IsUseJSON    bool
}

// createSinkFromDriver instantiates the requested driver sink.
func createSinkFromDriver(cfg Config) (LogSink, error) {
	switch cfg.Driver {
	case DriverFile:
		return NewFileSink(cfg.FilePath)
	case DriverZap:
		return NewZapAdapter(cfg.ZapLogger), nil
	case DriverComposite:
		return NewCompositeSink(cfg.Sinks...), nil
	case DriverConsole:
		return NewConsoleSink(os.Stdout, cfg.IsUseJSON), nil
	default:
		return NewConsoleSink(os.Stdout, cfg.IsUseJSON), nil
	}
}

// New constructs a Logger using the requested configuration and sink driver.
func New(cfg Config) (Logger, error) {
	sink, err := createSinkFromDriver(cfg)
	if err != nil {
		return nil, err
	}

	return &appLogger{
		minLevel: cfg.MinLevel,
		sink:     sink,
		fields:   nil,
	}, nil
}

// Default returns a standard Console logger at Info level.
func Default() Logger {
	l, _ := New(Config{
		MinLevel:  LevelInfo,
		Driver:    DriverConsole,
		IsUseJSON: false,
	})

	return l
}
