package applogger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

// ApiLogSenderFunc defines pluggable network transmission for serialized log batches.
type ApiLogSenderFunc func(
	ctx context.Context,
	endpoint string,
	headers map[string]string,
	payload []byte,
) *appfault.AppError

// ApiRotationPolicyFunc evaluates whether the current in-memory batch should be flushed/rotated.
type ApiRotationPolicyFunc func(
	batch []LogEntry,
	elapsed time.Duration,
	cfg ApiConfig,
) bool

// ApiConfig configures destination, batching, intervals, and extension policies for remote logging.
type ApiConfig struct {
	Endpoint       string
	Headers        map[string]string
	BatchSize      int
	FlushInterval  time.Duration
	Timeout        time.Duration
	Sender         ApiLogSenderFunc
	RotationPolicy ApiRotationPolicyFunc
}

func (cfg *ApiConfig) normalize() {
	if cfg.BatchSize <= 0 {
		cfg.BatchSize = 50
	}

	if cfg.Timeout <= 0 {
		cfg.Timeout = 10 * time.Second
	}

	if cfg.Sender == nil {
		cfg.Sender = DefaultHttpSender
	}

	if cfg.RotationPolicy == nil {
		cfg.RotationPolicy = defaultRotationPolicy
	}
}

// ApiSink implements LogSink, buffering and dispatching structured logs to an HTTP/API endpoint.
type ApiSink struct {
	lock           sync.Mutex
	cfg            ApiConfig
	buffer         []LogEntry
	lastFlush      time.Time
	sender         ApiLogSenderFunc
	rotationPolicy ApiRotationPolicyFunc
	stopCh         chan struct{}
	wg             sync.WaitGroup
}

// ApiManager is an alias for ApiSink providing unified terminology across SDK and specs.
type ApiManager = ApiSink

// NewApiSink constructs an active remote API logging sink.
func NewApiSink(cfg ApiConfig) (*ApiSink, error) {
	cfg.normalize()
	sink := &ApiSink{
		cfg:            cfg,
		buffer:         make([]LogEntry, 0, cfg.BatchSize),
		lastFlush:      time.Now(),
		sender:         cfg.Sender,
		rotationPolicy: cfg.RotationPolicy,
		stopCh:         make(chan struct{}),
	}

	sink.startBackgroundFlusher()

	return sink, nil
}

// NewApiManager instantiates an ApiManager instance.
func NewApiManager(cfg ApiConfig) (*ApiManager, error) {
	return NewApiSink(cfg)
}

func (s *ApiSink) startBackgroundFlusher() {
	if s.cfg.FlushInterval <= 0 {
		return
	}

	s.wg.Add(1)
	go s.runFlushLoop()
}

func (s *ApiSink) runFlushLoop() {
	defer s.wg.Done()
	ticker := time.NewTicker(s.cfg.FlushInterval)
	defer ticker.Stop()

	for {
		select {
		case <-s.stopCh:
			return
		case <-ticker.C:
			_ = s.Flush()
		}
	}
}

// WriteEntry persists an entry into the batch buffer, flushing if rotation policy triggers.
func (s *ApiSink) WriteEntry(e LogEntry) error {
	s.lock.Lock()
	s.buffer = append(s.buffer, e)
	elapsed := time.Since(s.lastFlush)
	shouldFlush := s.rotationPolicy(s.buffer, elapsed, s.cfg)
	s.lock.Unlock()

	if shouldFlush {
		if fault := s.Flush(); fault != nil {
			return fault
		}
	}

	return nil
}

// Flush immediately transmits all buffered log entries to the configured endpoint.
func (s *ApiSink) Flush() *appfault.AppError {
	s.lock.Lock()
	if len(s.buffer) == 0 {
		s.lock.Unlock()

		return nil
	}

	items := s.buffer
	s.buffer = make([]LogEntry, 0, s.cfg.BatchSize)
	s.lastFlush = time.Now()
	s.lock.Unlock()

	return s.sendBatch(items)
}

func (s *ApiSink) sendBatch(entries []LogEntry) *appfault.AppError {
	payload, err := json.Marshal(entries)
	if err != nil {
		return appfault.Wrap(errtype.Serialization, err, "failed to serialize log batch")
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.Timeout)
	defer cancel()

	return s.sender(ctx, s.cfg.Endpoint, s.cfg.Headers, payload)
}

// Sync forces pending logs to be pushed.
func (s *ApiSink) Sync() error {
	if fault := s.Flush(); fault != nil {
		return fault
	}

	return nil
}

// Close flushes remaining entries and shuts down background flush routines.
func (s *ApiSink) Close() error {
	close(s.stopCh)
	s.wg.Wait()

	return s.Sync()
}

// SetSender dynamically replaces the transport sender function.
func (s *ApiSink) SetSender(sender ApiLogSenderFunc) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if sender != nil {
		s.sender = sender
	}
}

// SetRotationPolicy configures a custom rule for batch flushing and rotation.
func (s *ApiSink) SetRotationPolicy(policy ApiRotationPolicyFunc) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if policy != nil {
		s.rotationPolicy = policy
	}
}

// BufferedCount returns the count of entries currently held in memory.
func (s *ApiSink) BufferedCount() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.buffer)
}

func defaultRotationPolicy(batch []LogEntry, elapsed time.Duration, cfg ApiConfig) bool {
	if len(batch) >= cfg.BatchSize {
		return true
	}

	if cfg.FlushInterval > 0 && elapsed >= cfg.FlushInterval {
		return true
	}

	return false
}

// DefaultHttpSender delivers log payloads over HTTP POST.
func DefaultHttpSender(
	ctx context.Context,
	endpoint string,
	headers map[string]string,
	payload []byte,
) *appfault.AppError {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(payload))
	if err != nil {
		return appfault.Wrap(errtype.Network, err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return executeHttpRequest(req)
}

func executeHttpRequest(req *http.Request) *appfault.AppError {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return appfault.Wrap(errtype.Network, err, "http log delivery failed")
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		msg := fmt.Sprintf("remote endpoint rejected logs: status %d", resp.StatusCode)

		return appfault.New(errtype.Network, msg)
	}

	return nil
}
