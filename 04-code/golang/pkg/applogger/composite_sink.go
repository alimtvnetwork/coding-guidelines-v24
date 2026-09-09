package applogger

import "sync"

// CompositeSink broadcasts log entries to multiple sinks.
type CompositeSink struct {
	lock  sync.RWMutex
	sinks []LogSink
}

// NewCompositeSink creates a multi-destination broadcaster sink.
func NewCompositeSink(sinks ...LogSink) *CompositeSink {
	return &CompositeSink{
		sinks: sinks,
	}
}

// AddSink appends an additional sink destination.
func (cs *CompositeSink) AddSink(sink LogSink) *CompositeSink {
	if sink == nil {
		return cs
	}

	cs.lock.Lock()
	defer cs.lock.Unlock()
	cs.sinks = append(cs.sinks, sink)

	return cs
}

// WriteEntry forwards the entry to all configured sinks.
func (cs *CompositeSink) WriteEntry(e LogEntry) error {
	cs.lock.RLock()
	defer cs.lock.RUnlock()

	for _, sink := range cs.sinks {
		_ = sink.WriteEntry(e)
	}

	return nil
}

// Sync flushes all inner sinks.
func (cs *CompositeSink) Sync() error {
	cs.lock.RLock()
	defer cs.lock.RUnlock()

	for _, sink := range cs.sinks {
		_ = sink.Sync()
	}

	return nil
}

// Close closes all inner sinks.
func (cs *CompositeSink) Close() error {
	cs.lock.Lock()
	defer cs.lock.Unlock()

	for _, sink := range cs.sinks {
		_ = sink.Close()
	}

	return nil
}

// DriverType returns the driver type.
func (cs *CompositeSink) DriverType() DriverType {
	return DriverComposite
}

// Sinks returns a safe cloned slice of registered sinks.
func (cs *CompositeSink) Sinks() []LogSink {
	cs.lock.RLock()
	defer cs.lock.RUnlock()

	cloned := make([]LogSink, len(cs.sinks))
	copy(cloned, cs.sinks)

	return cloned
}

// FilePath traverses sinks for FilePathProvider and returns the first non-empty path.
func (cs *CompositeSink) FilePath() string {
	cs.lock.RLock()
	defer cs.lock.RUnlock()

	for _, s := range cs.sinks {
		if acc, isOk := s.(FilePathProvider); isOk {
			if path := acc.FilePath(); len(path) > 0 {
				return path
			}
		}
	}

	return ""
}

// EndpointPath traverses sinks for EndpointPathProvider and returns the first non-empty endpoint.
func (cs *CompositeSink) EndpointPath() string {
	cs.lock.RLock()
	defer cs.lock.RUnlock()

	for _, s := range cs.sinks {
		if acc, isOk := s.(EndpointPathProvider); isOk {
			if ep := acc.EndpointPath(); len(ep) > 0 {
				return ep
			}
		}
	}

	return ""
}

// EndPointPath traverses sinks for EndpointPathAccessor and returns the first non-empty endpoint.
func (cs *CompositeSink) EndPointPath() string {
	return cs.EndpointPath()
}
