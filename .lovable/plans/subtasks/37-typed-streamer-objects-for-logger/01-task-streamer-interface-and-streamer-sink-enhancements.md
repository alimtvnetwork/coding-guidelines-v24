# Subtask 37.1: Streamer Interface and StreamerSink Enhancements

## Status
- **State:** Complete
- **Assigned Files:**
  - `04-code/golang/pkg/applogger/interfaces.go`
  - `04-code/golang/pkg/applogger/streamer_sink.go`

## Acceptance Criteria
1. Define `LogStreamer` interface in `04-code/golang/pkg/applogger/interfaces.go`:
   - `Name() string`
   - `StreamEntry(entry LogEntry) error`
   - `Stream(ctx context.Context, payload any) *appfault.AppError`
   - `Destination() io.Writer`
   - `Sync() error`
   - `Close() error`
   - `Streamer() any`
2. Add type aliases:
   - `Streamer = LogStreamer`
   - `LogStream = LogStreamer`
3. Update `StreamersProvider` in `interfaces.go`:
   - `Streamers() []Streamer`
4. Update `Logger` in `interfaces.go`:
   - `Streamers() []Streamer`
5. Implement `LogStreamer` methods on `*StreamerSink` in `04-code/golang/pkg/applogger/streamer_sink.go`:
   - `StreamEntry(e LogEntry) error`
   - `Stream(ctx context.Context, payload any) *appfault.AppError`
   - `Destination() io.Writer`
   - `Unwrap() any`
6. Add compile-time interface assertions:
   - `var _ LogStreamer = (*StreamerSink)(nil)`
   - `var _ Streamer = (*StreamerSink)(nil)`
7. Keep all functions $\le 15$ lines and implicit booleans.
