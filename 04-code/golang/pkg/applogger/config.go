package applogger

import "os"

// Config configures the logger instance.
type Config struct {
	MinLevel     LogLevel
	Driver       DriverType
	FilePath     string
	Endpoint     string
	ApiConfig    ApiConfig
	ZapLogger    ZapLoggerInterface
	Sinks        []LogSink
	Rotation     RotationConfig
	IsStackTrace bool
	IsUseJSON    bool
}

func makeFileSink(path string) LogSinkResult {
	res := NewFileSink(path)
	if res.IsFailed() {
		return LogSinkFailure(res)
	}

	return LogSinkSuccess(res.Data())
}

func makeRotatingSink(rot RotationConfig) LogSinkResult {
	res := NewRotatingFileSink(rot)
	if res.IsFailed() {
		return LogSinkFailure(res)
	}

	return LogSinkSuccess(res.Data())
}

func makeApiSink(cfg Config) LogSinkResult {
	apiCfg := cfg.ApiConfig
	if apiCfg.Endpoint == "" && cfg.Endpoint != "" {
		apiCfg.Endpoint = cfg.Endpoint
	}

	res := NewApiSink(apiCfg)
	if res.IsFailed() {
		return LogSinkFailure(res)
	}

	return LogSinkSuccess(res.Data())
}

// createSinkFromDriver instantiates the requested driver sink.
func createSinkFromDriver(cfg Config) LogSinkResult {
	switch cfg.Driver {
	case DriverFile:
		return makeFileSink(cfg.FilePath)
	case DriverRotatingFile:
		return makeRotatingSink(cfg.Rotation)
	case DriverZap:
		return LogSinkSuccess(NewZapAdapter(cfg.ZapLogger))
	case DriverComposite:
		return LogSinkSuccess(NewCompositeSink(cfg.Sinks...))
	case DriverApi:
		return makeApiSink(cfg)
	case DriverJsonWriterLogger:
		return LogSinkSuccess(NewConsoleSink(os.Stdout, true))
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
		minLevel:     cfg.MinLevel,
		sink:         sinkRes.Data(),
		fields:       nil,
		driverType:   cfg.Driver,
		filePath:     cfg.FilePath,
		endpointPath: cfg.Endpoint,
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
	res := Default()

	return res.Data()
}
