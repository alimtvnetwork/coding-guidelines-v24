# Subtask 02: Concrete Result Types & Clean Path Context in Fileutil

> **Parent Plan:** `17-structured-fileutil-context-and-concrete-results`  
> **Status:** Completed  
> **Bounded Files:**  
> - `04-code/golang/pkg/fileutil/types.go`  
> - `04-code/golang/pkg/fileutil/results.go`  
> - `04-code/golang/pkg/fileutil/fileutil.go`  
> - `04-code/golang/pkg/fileutil/create.go`  
> - `04-code/golang/pkg/fileutil/write.go`  
> - `04-code/golang/pkg/fileutil/parsers.go`  
> - `04-code/golang/pkg/fileutil/exporters.go`  
> - `04-code/golang/pkg/fileutil/stream.go`  
> - `04-code/golang/pkg/fileutil/advanced.go`  
> - `04-code/golang/pkg/fileutil/locked_operations.go`  
> - `04-code/golang/pkg/fileutil/fileutil_test.go`  
> - `04-code/golang/pkg/fileutil/create_test.go`  
> - `04-code/golang/pkg/fileutil/advanced_test.go`

---

## Instructions

1. **In `04-code/golang/pkg/fileutil/types.go`:**
   - Define concrete non-generic type aliases:
     ```go
     type (
         FileResult     = result.Wrap[*os.File]
         BytesResult    = result.Wrap[[]byte]
         StringResult   = result.Wrap[string]
         LinesResult    = result.Wrap[[]string]
         BoolResult     = result.Wrap[bool]
         FileInfoResult = result.Wrap[os.FileInfo]
         Int64Result    = result.Wrap[int64]
     )
     ```

2. **In `04-code/golang/pkg/fileutil/results.go`:**
   - Define specialized constructor helpers:
     - `FileSuccess(f *os.File) FileResult`
     - `FileFailure(variation errtype.Variation, err error, path string, msg string) FileResult`
     - `FileFailureMsg(variation errtype.Variation, path string, msg string) FileResult`
     - `BoolSuccess(val bool) BoolResult`
     - `BoolFailure(variation errtype.Variation, err error, path string, msg string) BoolResult`
     - `BoolFailureMsg(variation errtype.Variation, path string, msg string) BoolResult`
     - `BytesSuccess(data []byte) BytesResult`
     - `BytesFailure(variation errtype.Variation, err error, path string, msg string) BytesResult`
     - `BytesFailureMsg(variation errtype.Variation, path string, msg string) BytesResult`
     - `StringSuccess(s string) StringResult`
     - `StringFailure(variation errtype.Variation, err error, path string, msg string) StringResult`
     - `StringFailureMsg(variation errtype.Variation, path string, msg string) StringResult`
     - `LinesSuccess(lines []string) LinesResult`
     - `LinesFailure(variation errtype.Variation, err error, path string, msg string) LinesResult`
     - `LinesFailureMsg(variation errtype.Variation, path string, msg string) LinesResult`
     - `FileInfoSuccess(info os.FileInfo) FileInfoResult`
     - `FileInfoFailure(variation errtype.Variation, err error, path string, msg string) FileInfoResult`
     - `FileInfoFailureMsg(variation errtype.Variation, path string, msg string) FileInfoResult`
     - `Int64Success(n int64) Int64Result`
     - `Int64Failure(variation errtype.Variation, err error, path string, msg string) Int64Result`
     - `Int64FailureMsg(variation errtype.Variation, path string, msg string) Int64Result`
   - All failure constructors use `appfault.WrapFile` / `appfault.NewFile` to store `"Path"` in context and avoid string concatenation.

3. **In `04-code/golang/pkg/fileutil/fileutil.go`:**
   - Update return types to `FileResult`, `BytesResult`, `StringResult`, `BoolResult`, `FileInfoResult`.
   - Replace all 17 string concatenations with `FileFailure`, `BoolFailure`, `BytesFailure`, `FileInfoFailure` passing path as context.

4. **In `04-code/golang/pkg/fileutil/create.go`, `write.go`, `parsers.go`, `exporters.go`, `stream.go`, `advanced.go`, `locked_operations.go`:**
   - Update function signatures to return concrete aliases.
   - Replace remaining string concatenations with structured context constructors.

5. **Unit Tests:**
   - Verify all tests in `pkg/fileutil` pass.
   - Run `go test -C 04-code/golang -v ./pkg/fileutil`.
