# Subtask 03: Safe Atomic File Writing in fileutil

## 1. Goal
Implement `AtomicWriteFile` in `04-code/golang/pkg/fileutil/` to prevent corrupted, partial, or half-written files during unexpected power loss or process crashes.

**Status:** ✅ Completed

## 2. Target Files
- `04-code/golang/pkg/fileutil/atomic_write.go`
- `04-code/golang/pkg/fileutil/atomic_write_test.go`

## 3. Detailed Specifications
1. **Atomic Write Function**:
   - `AtomicWriteFile(filePath string, data []byte, perm filepermtype.Type) *appfault.AppError`
   - Algorithm:
     1. Ensure target directory exists via `fileutil.EnsureDir`.
     2. Create a temporary file in the same directory (e.g. `filepath.Join(dir, fmt.Sprintf(".tmp-%d-%d", os.Getpid(), time.Now().UnixNano())))`.
     3. Write `data` to temporary file.
     4. Call `file.Sync()` to flush OS page cache to disk.
     5. Close temporary file safely.
     6. Atomically replace target file using `os.Rename`. On Windows, if destination exists, handle safe atomic swap.
     7. If error occurs at any point, clean up temporary file via defer.
2. **Coding Guidelines Compliance**:
   - Functions <= 15 lines.
   - Clean Unix LF, blank lines before return.
