# Subtask 04: Fileutil and Appwriter Verification & Hardening

## Objective
Verify all requested file tactics in `pkg/fileutil` and `pkg/appwriter`, ensuring zero panics, path-level mutex concurrency with automatic reference-counted eviction, and error wrapping.

## Target Files
- `04-code/golang/pkg/fileutil/write.go`
- `04-code/golang/pkg/fileutil/locker.go`
- `04-code/golang/pkg/appwriter/file_writer.go`

## Detailed Instructions
1. Review `pkg/fileutil/locker.go`:
   - Verify `GetFileLock(path string) *sync.RWMutex` increments `refCount`.
   - Verify `ReleaseFileLock(path string)` decrements `refCount` and removes idle locks from `fileLocksMap` when `refCount <= 0`.
2. Review `pkg/fileutil/write.go`:
   - Verify all file writing functions (`Write`, `WriteLocked`, `WriteBytes`, `WriteJSON`, etc.) return `result.Wrap[bool]` and properly release locks.
   - Verify zero panics exist in file operations.
3. Review `pkg/appwriter/file_writer.go`:
   - Verify `fileWriteFunc` converts payloads via `payloadconv.ToBytes` and returns `*appfault.AppError` on failure without panicking.
4. Add or verify test coverage in `pkg/fileutil/fileutil_test.go` and `pkg/appwriter/writer_test.go`.
