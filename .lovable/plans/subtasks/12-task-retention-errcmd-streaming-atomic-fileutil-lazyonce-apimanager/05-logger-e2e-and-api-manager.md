# Subtask 05: Logger E2E Tests & Extensible ApiManager

## 1. Goal
Provide robust end-to-end integration testing for log rotation and implement `ApiManager` / `ApiSink` to send logs to remote HTTP/API endpoints with pluggable sending, batching, and rotation extensions.

**Status:** ✅ Completed

## 2. Target Files
- `04-code/golang/pkg/applogger/api_sink.go`
- `04-code/golang/pkg/applogger/api_sink_test.go`
- `04-code/golang/pkg/applogger/rotating_file_sink_e2e_test.go`

## 3. Detailed Specifications
1. **Rotating File Sink E2E Test**:
   - Write realistic multiline logs exceeding threshold.
   - Verify file is rotated into archives folder, compressed with gzip (`.gz`), backup retention limit keeps at most N files, and active log file resets to 0 bytes.
2. **ApiManager / ApiSink**:
   - Implements `LogSink`.
   - Fields:
     - `Endpoint string`
     - `Headers map[string]string`
     - `BatchSize int`
     - `FlushInterval time.Duration`
     - `Sender ApiLogSenderFunc` (allows overriding transport or mock testing)
   - Methods:
     - `WriteLog(entry LogEntry) *appfault.AppError`
     - `Flush() *appfault.AppError`
     - `Close() error`
     - `SetSender(sender ApiLogSenderFunc)`
3. **Coding Guidelines Compliance**:
   - Functions <= 15 lines.
