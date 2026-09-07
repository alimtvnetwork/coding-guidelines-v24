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
