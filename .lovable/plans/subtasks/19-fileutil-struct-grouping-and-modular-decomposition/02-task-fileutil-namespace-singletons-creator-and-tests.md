# Subtask 02: File & New Namespace Singletons, Tests & Docs

> **Parent Plan:** `19-fileutil-struct-grouping-and-modular-decomposition`  
> **Status:** Completed  
> **Bounded Files:**  
> - `04-code/golang/pkg/fileutil/file_namespace.go`  
> - `04-code/golang/pkg/fileutil/file_namespace_test.go`  
> - `04-code/golang/pkg/fileutil/file_append_test.go`  
> - `04-code/golang/pkg/fileutil/readme.md`

---

## Instructions

1. **In `04-code/golang/pkg/fileutil/file_namespace.go`:**
   - Define the `File` operational namespace:
     ```go
     type fileNamespace struct {
         Open   openOps
         Create createOps
         Write  writeOps
         Append appendOps
         Read   readOps
         Path   pathNamespace
     }

     var File = &fileNamespace{
         Open:   openOps{},
         Create: createOps{},
         Write:  writeOps{},
         Append: appendOps{},
         Read:   readOps{},
         Path:   Path,
     }
     ```
   - Define the `New` creator namespace (matching the `coredata` creator concept):
     ```go
     type (
         fileNewCreator struct {
             Writer       fileWriterCreator
             Appender     fileAppenderCreator
             BoundWriter  fileBoundWriterCreator
             Path         filePathCreator
             StreamWriter fileStreamWriterCreator
         }

         fileWriterCreator       struct{}
         fileAppenderCreator     struct{}
         fileBoundWriterCreator  struct{}
         filePathCreator         struct{}
         fileStreamWriterCreator struct{}
     )

     var New = &fileNewCreator{
         Writer:       fileWriterCreator{},
         Appender:     fileAppenderCreator{},
         BoundWriter:  fileBoundWriterCreator{},
         Path:         filePathCreator{},
         StreamWriter: fileStreamWriterCreator{},
     }
     ```
   - Implement creator methods on `New.Writer`, `New.Appender`, `New.BoundWriter`, `New.Path`, `New.StreamWriter`, plus shortcut helpers on `fileNewCreator`.

2. **In `04-code/golang/pkg/fileutil/file_namespace_test.go`:**
   - Comprehensive unit tests covering:
     - `File.Open.*` (`File`, `ReadOnly`, `ReadWrite`, `Append`, `Truncate`, `CreateAppend`)
     - `File.Create.*` (`File`, `Dir`, `EnsureDir`, `Temp`, `TempDir`)
     - `File.Write.*` (`Any`, `Bytes`, `String`, `Lines`, `Json`, `Yaml`, `Atomic`)
     - `File.Append.*` (`Bytes`, `String`, `Lines`, `BytesLocked`)
     - `File.Read.*` (`Bytes`, `String`, `Text`, `Lines`)
     - `New.*` creators (`New.Writer`, `New.Appender`, `New.BoundWriter`, `New.Path`, `New.StreamWriter`)
   - Test backward compatibility: all package-level functions match `File.*` operations.

3. **In `04-code/golang/pkg/fileutil/file_append_test.go`:**
   - Unit tests covering `AppendBytes`, `AppendString`, `AppendLines`, and their locked variants.

4. **In `04-code/golang/pkg/fileutil/readme.md`:**
   - Document `File.*` and `New.*` namespaces, sub-operation structs, creator pattern, and usage examples.

5. **Strict Guidelines Compliance:**
   - All functions <= 15 lines.
   - Blank line after closing brace `}` if followed by code.
   - Blank line before `return` unless sole statement in block.
   - Run `go test -C 04-code/golang -v ./pkg/fileutil` to verify.
