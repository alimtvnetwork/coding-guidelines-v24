# Subtask 02: Normalization, Inspection & Namespace Fluent Builder

> **Parent Plan:** `18-modular-filepath-util-and-cross-platform-temp`  
> **Status:** Completed  
> **Bounded Files:**  
> - `04-code/golang/pkg/fileutil/path_normalize.go`  
> - `04-code/golang/pkg/fileutil/path_normalize_test.go`  
> - `04-code/golang/pkg/fileutil/path_info.go`  
> - `04-code/golang/pkg/fileutil/path_info_test.go`  
> - `04-code/golang/pkg/fileutil/path_namespace.go`  
> - `04-code/golang/pkg/fileutil/path_namespace_test.go`  
> - `04-code/golang/pkg/fileutil/readme.md`

---

## Instructions

1. **In `04-code/golang/pkg/fileutil/path_normalize.go`:**
   - `ToSlash(path string) string`: Converts separators to `/`.
   - `ToBackslash(path string) string`: Converts separators to `\`.
   - `ToNative(path string) string`: Converts separators to host `os.PathSeparator`.
   - `Clean(path string) string`: `filepath.Clean` preserving UNC and long path prefixes (`\\?\`).
   - `DeduplicateSeparators(path string) string`: Collapses redundant consecutive slashes/backslashes.
   - `HasLongPathPrefix(path string) bool`: Checks for `\\?\` prefix.
   - `TrimLongPathPrefix(path string) string`: Strips `\\?\` prefix.
   - `ToLongPath(path string) string`: Prepends `\\?\` to absolute Windows paths when needed.
   - `Normalize(path string) StringResult`: Clean + ToNative + Deduplicate.
   - `NormalizeToSlash(path string) StringResult`: Clean + ToSlash + Deduplicate.

2. **In `04-code/golang/pkg/fileutil/path_info.go`:**
   - `Ext(path string) string`: Extension with leading dot (e.g. `.json`).
   - `ExtNoDot(path string) string`: Extension without leading dot (e.g. `json`).
   - `HasExt(path string, ext string) bool`: Case-insensitive extension comparison.
   - `Base(path string) string`: Final element of path.
   - `Stem(path string) string`: Filename without final extension (`file.tar.gz` -> `file.tar`).
   - `StemFull(path string) string`: Filename without any extensions (`file.tar.gz` -> `file`).
   - `Slug(path string) string`: Safe lowercase filename slug.
   - `Dir(path string) string`: Directory containing the path.
   - `Split(path string) (dir string, file string)`: Splits directory and file.
   - `Parent(path string) string`: Immediate parent directory.
   - `ParentN(path string, levels int) string`: Traverses N levels up.
   - `IsAbs(path string) bool`: Reports if path is absolute.
   - `IsRel(path string) bool`: Reports if path is relative.

3. **In `04-code/golang/pkg/fileutil/path_namespace.go`:**
   - Define `var Path = pathNamespace{}` singleton:
     - `Path.Temp.*`: Sub-namespace forwarding to `UserTempDir()`, `UserTempPath()`, `TempFile()`, `TempDir()`, `CreateTempFile()`, `CreateTempDir()`.
     - `Path.Env.*`: Sub-namespace forwarding to `Expand()`, `ExpandEnv()`, `ExpandTilde()`.
     - `Path.Norm.*`: Sub-namespace forwarding to `Clean()`, `Normalize()`, `NormalizeToSlash()`, `ToSlash()`, `ToBackslash()`, `ToNative()`, `Deduplicate()`.
     - `Path.Info.*`: Sub-namespace forwarding to `Ext()`, `ExtNoDot()`, `Base()`, `Stem()`, `StemFull()`, `Slug()`, `Dir()`, `Parent()`, `ParentN()`, `IsAbs()`, `IsRel()`.
     - `Path.Join(elem ...string) string`: `filepath.Join` shortcut.
   - Define fluent constructor `NewPath(raw string) *PathWrapper`:
     - Chaining methods: `Raw()`, `String()`, `Clean()`, `Normalize()`, `ToSlash()`, `ToBackslash()`, `ToNative()`, `Expand()`, `ExpandEnv()`, `ExpandTilde()`, `Join(elem...)`, `Parent()`, `ParentN(n)`.
     - Inspection methods: `Base()`, `Stem()`, `Ext()`, `ExtNoDot()`, `Dir()`, `IsAbs()`, `IsRel()`.
     - Direct file operations: `Exists() BoolResult`, `Stat() FileInfoResult`, `Read() BytesResult`, `ReadString() StringResult`, `Write(data, perm) BoolResult`, `WriteString(content, perm) BoolResult`.

4. **In `04-code/golang/pkg/fileutil/readme.md`:**
   - Document new modular filepath APIs and `Path.*` namespace.

5. **Unit Tests:**
   - Add unit tests in `path_normalize_test.go`, `path_info_test.go`, `path_namespace_test.go`.
   - Verify `go test -C 04-code/golang -v ./pkg/fileutil` passes.
