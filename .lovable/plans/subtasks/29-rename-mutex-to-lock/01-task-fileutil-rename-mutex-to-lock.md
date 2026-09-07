# Subtask 29.1: Rename `mu` to `lock` in `pkg/fileutil`

## Context
Standardize concurrency fields in `fileutil` to use `lock` instead of `mu`.

## Target Files
1. `04-code/golang/pkg/fileutil/bound_file_writer.go`:
   - Rename field `mu sync.Mutex` -> `lock sync.Mutex`
   - Update all `w.mu.Lock()` -> `w.lock.Lock()` and `w.mu.Unlock()` -> `w.lock.Unlock()`
2. `04-code/golang/pkg/fileutil/writer_appender.go`:
   - Rename `FileWriter.mu` -> `FileWriter.lock`
   - Rename `FileAppender.mu` -> `FileAppender.lock`
   - Update all method calls
3. `04-code/golang/pkg/fileutil/locker.go`:
   - Rename `lockEntry.mu` -> `lockEntry.lock`
   - Rename `fileLocksMu` -> `fileLocksLock`
   - Rename local variable `mu` -> `lock`
4. `04-code/golang/pkg/fileutil/append_ops.go`:
   - Rename local variable `mu := GetFileLock(path)` -> `lock := GetFileLock(path)`
5. `04-code/golang/pkg/fileutil/locked_operations.go`:
   - Rename local variables `mu` -> `lock`
6. `04-code/golang/pkg/fileutil/write.go`:
   - Rename local variables `mu` -> `lock`
7. `04-code/golang/pkg/fileutil/readme.md`:
   - Update documentation `mu` -> `lock`

## Verification
- `go test -v -C 04-code/golang -count=1 ./pkg/fileutil/...`
