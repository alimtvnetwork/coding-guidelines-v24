# Subtask 29.2: Rename `mu` and `ReentrantMutex` in `pkg/streamwriter`

## Context
Standardize concurrency fields in `streamwriter` to use `lock` and `ReentrantLock`.

## Target Files
1. `04-code/golang/pkg/streamwriter/mutex.go`:
   - Rename `ReentrantMutex` -> `ReentrantLock` with `type ReentrantMutex = ReentrantLock` alias
   - Rename field `mu` -> `lock`
2. `04-code/golang/pkg/streamwriter/writer.go`:
   - Rename `Writer.mu` -> `Writer.lock`
   - Rename `Writer.configMu` -> `Writer.configLock`
3. `04-code/golang/pkg/streamwriter/async_writer.go`:
   - Rename `AsyncWriter.mu` -> `AsyncWriter.lock`
4. `04-code/golang/pkg/streamwriter/locked_streamer.go`:
   - Rename `LockedStreamer.mu` -> `LockedStreamer.lock`
5. `04-code/golang/pkg/streamwriter/logger.go`:
   - Rename `Logger.mu` -> `Logger.lock`
6. `04-code/golang/pkg/streamwriter/streamwriter_test.go` and `async_writer_test.go`:
   - Rename `SafeBuffer.mu` -> `SafeBuffer.lock`, `MockWriter.mu` -> `MockWriter.lock`

## Verification
- `go test -v -C 04-code/golang -count=1 ./pkg/streamwriter/...`
