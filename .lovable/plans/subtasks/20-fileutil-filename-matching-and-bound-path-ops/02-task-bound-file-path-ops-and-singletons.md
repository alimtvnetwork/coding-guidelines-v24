# Subtask 02: Bound File Path Operations & Singletons

## Objective
Implement `FilePathOps` in `04-code/golang/pkg/fileutil/file_path_ops.go` to provide a bound, path-encapsulated file operation API with immutable cloning, pre-flight safety methods, and zero-path argument methods. Integrate into `File` and `New` singletons.

## Target Files
- `04-code/golang/pkg/fileutil/file_path_ops.go`
- `04-code/golang/pkg/fileutil/file_path_ops_test.go`
- `04-code/golang/pkg/fileutil/file_namespace.go`
- `04-code/golang/pkg/fileutil/file_namespace_test.go`

## Implementation Details

1. **`FilePathOps` Struct Definition:**
   ```go
   type FilePathOps struct {
       workDir string
       relPath string
       absPath string
   }
   ```

2. **Constructors:**
   - `NewFilePathOps(path string) *FilePathOps`
   - `NewFilePathOpsAt(workDir, relPath string) *FilePathOps`

3. **Immutability & Modifiers:**
   - `Clone() *FilePathOps`
   - `WithWorkDir(dir string) *FilePathOps`
   - `WithRelPath(rel string) *FilePathOps`
   - `Join(elem ...string) *FilePathOps`

4. **Accessors:**
   - `WorkDir() string`, `RelPath() string`, `AbsPath() string`, `String() string`, `Dir() string`, `Base() string`, `Ext() string`, `Exists() bool`, `Stat() FileInfoResult`

5. **Pre-flight & Safety Operations:**
   - `EnsureParentDir() BoolResult`
   - `EnsureFile(perm FilePermType) BoolResult`
   - `CreateIfNotExist(perm FilePermType) FileResult`

6. **Zero-Path Bound Operations:**
   - `ReadBytes() BytesResult`, `ReadString() StringResult`, `ReadLines() LinesResult`
   - `WriteBytes(data []byte, perm FilePermType) BoolResult`, `WriteString(content string, perm FilePermType) BoolResult`, `WriteLines(lines []string, perm FilePermType) BoolResult`, `WriteAtomic(data []byte, perm FilePermType) BoolResult`
   - `AppendBytes(data []byte, perm FilePermType) BoolResult`, `AppendString(content string, perm FilePermType) BoolResult`, `AppendLines(lines []string, perm FilePermType) BoolResult`
   - `Open(openMode FileOpenModeType, perm FilePermType) FileResult`, `Create(perm FilePermType) FileResult`, `Delete() BoolResult`

7. **Singleton Integration (`file_namespace.go`):**
   - Add to `fileNamespace`:
     - `Target(path string) *FilePathOps`
     - `At(workDir, relPath string) *FilePathOps`
   - Add to `fileNewCreator`:
     - `Target(path string) *FilePathOps`
     - `At(workDir, relPath string) *FilePathOps`
     - `FilePathOps(path string) *FilePathOps`

8. **Coding Rules:**
   - Function length <= 15 lines.
   - Blank line after closing brace if followed by code.
   - Blank line before return unless sole statement.
   - Zero panic policy.
