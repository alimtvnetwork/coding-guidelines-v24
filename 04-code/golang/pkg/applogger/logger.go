package applogger

import (
	"fmt"
	"time"

	"coding-guidelines/common/pkg/appfault"
)

type appLogger struct {
	minLevel     LogLevel
	sink         LogSink
	fields       appfault.ContextMap
	driverType   DriverType
	filePath     string
	endpointPath string
	streamer     any
}

func (l *appLogger) createEntry(lvl LogLevel, msg, stack string) LogEntry {
	return LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     lvl,
		Message:   msg,
		Fields:    l.fields.Clone(),
		Caller:    appfault.CaptureCaller(4),
		Stack:     stack,
	}
}

func (l *appLogger) write(lvl LogLevel, msg string, stack string) {
	if lvl.IsEnabled(l.minLevel) && l.sink != nil {
		_ = l.sink.WriteEntry(l.createEntry(lvl, msg, stack))
	}
}

func (l *appLogger) Debug(args ...any) Logger {
	l.write(LevelDebug, fmt.Sprint(args...), "")

	return l
}

func (l *appLogger) Info(args ...any) Logger {
	l.write(LevelInfo, fmt.Sprint(args...), "")

	return l
}

func (l *appLogger) Warn(args ...any) Logger {
	l.write(LevelWarn, fmt.Sprint(args...), "")

	return l
}

func (l *appLogger) Error(args ...any) Logger {
	l.write(LevelError, fmt.Sprint(args...), "")

	return l
}

func (l *appLogger) Fatal(args ...any) Logger {
	l.write(LevelFatal, fmt.Sprint(args...), "")

	return l
}

func (l *appLogger) Debugf(format string, args ...any) Logger {
	l.write(LevelDebug, fmt.Sprintf(format, args...), "")

	return l
}

func (l *appLogger) Infof(format string, args ...any) Logger {
	l.write(LevelInfo, fmt.Sprintf(format, args...), "")

	return l
}

func (l *appLogger) Warnf(format string, args ...any) Logger {
	l.write(LevelWarn, fmt.Sprintf(format, args...), "")

	return l
}

func (l *appLogger) Errorf(format string, args ...any) Logger {
	l.write(LevelError, fmt.Sprintf(format, args...), "")

	return l
}

func (l *appLogger) Fatalf(format string, args ...any) Logger {
	l.write(LevelFatal, fmt.Sprintf(format, args...), "")

	return l
}

func (l *appLogger) FilePath() string {
	if len(l.filePath) > 0 {
		return l.filePath
	}

	if acc, isOk := l.sink.(FilePathProvider); isOk {
		return acc.FilePath()
	}

	return ""
}

func (l *appLogger) EndpointPath() string {
	if len(l.endpointPath) > 0 {
		return l.endpointPath
	}

	if acc, isOk := l.sink.(EndpointPathProvider); isOk {
		return acc.EndpointPath()
	}

	return ""
}

func (l *appLogger) EndPointPath() string {
	return l.EndpointPath()
}

func (l *appLogger) Type() DriverType {
	if l.driverType.IsValid() && l.driverType != DriverConsole {
		return l.driverType
	}

	if acc, isOk := l.sink.(DriverTypeProvider); isOk {
		return acc.DriverType()
	}

	return l.driverType
}

func (l *appLogger) Clone() Logger {
	return &appLogger{
		minLevel:     l.minLevel,
		sink:         l.sink,
		fields:       l.fields.Clone(),
		driverType:   l.driverType,
		filePath:     l.filePath,
		endpointPath: l.endpointPath,
		streamer:     l.streamer,
	}
}

func (l *appLogger) AddWriters(writers ...LogSink) Logger {
	return &appLogger{
		minLevel:     l.minLevel,
		sink:         combineSinks(l.sink, writers...),
		fields:       l.fields.Clone(),
		driverType:   DriverComposite,
		filePath:     l.filePath,
		endpointPath: l.endpointPath,
		streamer:     l.streamer,
	}
}

func extractBaseSinks(base LogSink) []LogSink {
	if cs, isOk := base.(*CompositeSink); isOk {
		return cs.Sinks()
	}

	if base != nil {
		return []LogSink{base}
	}

	return nil
}

func combineSinks(base LogSink, additions ...LogSink) LogSink {
	if len(additions) == 0 {
		return base
	}

	all := extractBaseSinks(base)
	for _, w := range additions {
		if w != nil {
			all = append(all, w)
		}
	}

	return NewCompositeSink(all...)
}

func (l *appLogger) AddStreamer(streamer any) Logger {
	if streamer == nil {
		return l.Clone()
	}

	sink := NewStreamerSink(streamer)
	cloned := l.AddWriters(sink).(*appLogger)
	cloned.streamer = streamer

	return cloned
}

func (l *appLogger) Writers() []LogSink {
	return extractBaseSinks(l.sink)
}

func (l *appLogger) WriterNames() []string {
	writers := l.Writers()
	names := make([]string, 0, len(writers))
	for _, w := range writers {
		if w != nil {
			names = append(names, w.Name())
		}
	}

	return names
}

func extractSinkStreamers(writers []LogSink) []Streamer {
	var collected []Streamer
	for _, w := range writers {
		if ss, isOk := w.(*StreamerSink); isOk {
			collected = append(collected, ss)
		} else if st, isStreamer := w.(Streamer); isStreamer {
			collected = append(collected, st)
		}
	}

	return collected
}

func (l *appLogger) Streamers() []Streamer {
	collected := extractSinkStreamers(l.Writers())
	if len(collected) > 0 {
		return collected
	}

	if l.streamer != nil {
		if st, isStreamer := l.streamer.(Streamer); isStreamer {
			return []Streamer{st}
		}

		return []Streamer{NewStreamerSink(l.streamer)}
	}

	return nil
}
