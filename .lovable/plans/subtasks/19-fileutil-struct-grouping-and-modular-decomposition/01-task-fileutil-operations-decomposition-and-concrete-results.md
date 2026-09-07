# Subtask 01: Operation Structs Decomposition & Concrete Results

> **Parent Plan:** `19-fileutil-struct-grouping-and-modular-decomposition`  
> **Status:** Completed  
> **Bounded Files:**  
> - `04-code/golang/pkg/fileutil/file_open.go`  
> - `04-code/golang/pkg/fileutil/file_create.go`  
> - `04-code/golang/pkg/fileutil/file_append.go`  
> - `04-code/golang/pkg/fileutil/file_read.go`  
> - `04-code/golang/pkg/fileutil/file_write.go`  
> - `04-code/golang/pkg/fileutil/results.go`  
> - `04-code/golang/pkg/fileutil/fileutil.go`

---

## Instructions

1. **In `04-code/golang/pkg/fileutil/results.go`:**
   - Add companion fault forwarders to eradicate all bracket generics:
     ```go
     func FileFailureFault(fault *appfault.AppError) FileResult       { return result.WrapFailure[*os.File](fault) }
     func BoolFailureFault(fault *appfault.AppError) BoolResult       { return result.WrapFailure[bool](fault) }
     func BytesFailureFault(fault *appfault.AppError) BytesResult     { return result.WrapFailure[[]byte](fault) }
     func StringFailureFault(fault *appfault.AppError) StringResult   { return result.WrapFailure[string](fault) }
     func LinesFailureFault(fault *appfault.AppError) LinesResult     { return result.WrapFailure[[]string](fault) }
     func Int64FailureFault(fault *appfault.AppError) Int64Result     { return result.WrapFailure[int64](fault) }
     ```

2. **In `04-code/golang/pkg/fileutil/fileutil.go`:**
   - Fix line 55: replace `result.WrapFailure[*os.File]` with:
     ```go
     if err := ensureParentDir(path, openMode.Flags()); err != nil {
         return FileFailure(errtype.IO, err, path, "failed to create parent directory")
     }
     ```
   - Replace any remaining `result.WrapFailure[T]` calls with `StringFailureFault`, `BoolFailureFault`, `BytesFailureFault`, etc.
   - Refactor any function > 15 lines (`ensureParentDir`, `OpenFile`, `ExecuteOp`) into <= 15 line helpers.

3. **In `04-code/golang/pkg/fileutil/file_open.go`:**
   - Define `type openOps struct{}`
   - Implement methods on `openOps`:
     - `(openOps) File(path string, openMode FileOpenModeType, perm FilePermType) FileResult`
     - `(openOps) ReadOnly(path string) FileResult`
     - `(openOps) ReadWrite(path string, perm FilePermType) FileResult`
     - `(openOps) Append(path string, perm FilePermType) FileResult`
     - `(openOps) Truncate(path string, perm FilePermType) FileResult`
     - `(openOps) CreateAppend(path string, perm FilePermType) FileResult`
   - Add top-level package functions `OpenAppend` and `OpenWrite`.

4. **In `04-code/golang/pkg/fileutil/file_create.go`:**
   - Define `type createOps struct{}`
   - Implement methods on `createOps`:
     - `(createOps) File(path string, perm FilePermType) FileResult`
     - `(createOps) Dir(path string, perm FilePermType) BoolResult`
     - `(createOps) EnsureDir(path string, perm FilePermType) BoolResult`
     - `(createOps) Temp(pattern string) FileResult`
     - `(createOps) TempDir(pattern string) StringResult`
     - `(createOps) TempFileIn(dir string, pattern string, perm FilePermType) FileResult`
     - `(createOps) TempDirIn(dir string, pattern string, perm FilePermType) StringResult`
   - Add top-level package function `Create(path string, perm FilePermType) FileResult`.

5. **In `04-code/golang/pkg/fileutil/file_append.go`:**
   - Define `type appendOps struct{}`
   - Implement methods on `appendOps`:
     - `(appendOps) Bytes(path string, data []byte, perm FilePermType) BoolResult`
     - `(appendOps) String(path string, content string, perm FilePermType) BoolResult`
     - `(appendOps) Lines(path string, lines []string, perm FilePermType) BoolResult`
     - `(appendOps) BytesLocked(path string, data []byte, perm FilePermType) BoolResult`
     - `(appendOps) StringLocked(path string, content string, perm FilePermType) BoolResult`
     - `(appendOps) LinesLocked(path string, lines []string, perm FilePermType) BoolResult`
     - `(appendOps) NewAppender(path string, perm FilePermType) *FileAppender`
   - Add top-level package functions `AppendBytes`, `AppendBytesLocked`, `AppendString`, `AppendStringLocked`, `AppendLines`, `AppendLinesLocked`.

6. **In `04-code/golang/pkg/fileutil/file_read.go`:**
   - Define `type readOps struct{}`
   - Implement methods on `readOps`:
     - `(readOps) Bytes(path string) BytesResult`
     - `(readOps) String(path string) StringResult`
     - `(readOps) Text(path string) StringResult`
     - `(readOps) Lines(path string) LinesResult`
     - `(readOps) Chunked(path string, chunkSize int, onChunk ChunkCallbackFunc) Int64Result`
     - `(readOps) TextLocked(path string) StringResult`
     - `(readOps) LinesLocked(path string) LinesResult`
   - Add top-level package function `ReadBytes(path string) BytesResult`.

7. **In `04-code/golang/pkg/fileutil/file_write.go`:**
   - Define `type writeOps struct{}`
   - Implement methods on `writeOps`:
     - `(writeOps) Any(path string, payload any, perm FilePermType) BoolResult`
     - `(writeOps) Bytes(path string, data []byte, perm FilePermType) BoolResult`
     - `(writeOps) String(path string, content string, perm FilePermType) BoolResult`
     - `(writeOps) Lines(path string, lines []string, perm FilePermType) BoolResult`
     - `(writeOps) Json(path string, data any, perm FilePermType) BoolResult`
     - `(writeOps) Yaml(path string, data any, perm FilePermType) BoolResult`
     - `(writeOps) Atomic(path string, data []byte, perm FilePermType) BoolResult`
     - `(writeOps) Chunked(path string, perm FilePermType, reader io.Reader, bufferSize int) Int64Result`
     - And locked variants `BytesLocked`, `StringLocked`, `LinesLocked`, `JsonLocked`, `YamlLocked`.

8. **Strict Guidelines Compliance:**
   - All functions <= 15 lines.
   - Blank line after closing brace `}` if followed by code.
   - Blank line before `return` unless sole statement in block.
   - Run `go test -C 04-code/golang -v ./pkg/fileutil` to verify.
