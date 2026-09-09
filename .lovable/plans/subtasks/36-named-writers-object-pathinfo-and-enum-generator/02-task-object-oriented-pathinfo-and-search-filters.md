# Subtask 36.2: Object-Oriented PathInfo and Search Filters

## Status
- **State:** Complete
- **Assigned Files:**
  - `04-code/golang/pkg/fileutil/path_info_obj.go`
  - `04-code/golang/pkg/fileutil/path_info.go`
  - `04-code/golang/pkg/fileutil/folder_info.go`
  - `04-code/golang/pkg/fileutil/file_info.go`
  - `04-code/golang/pkg/fileutil/path_namespace.go`
  - `04-code/golang/pkg/fileutil/path_info_test.go`
  - `04-code/golang/pkg/fileutil/folder_info_test.go`
  - `04-code/golang/pkg/fileutil/file_info_test.go`

## Acceptance Criteria
1. Refactor `PathInfo` in `04-code/golang/pkg/fileutil/path_info_obj.go` into a comprehensive object created from any path (`fileutil.NewPathInfo(path)` or `fileutil.Path.Inspect(path)` or `fileutil.Path.Info(path)`):
   - Normalization & formatting: `Normalize() *PathInfo`, `ToSlash() string`, `ToNative() string`, `Clean() *PathInfo`, `Abs() result.Wrap[string]`, `Rel(base string) result.Wrap[string]`.
   - Dissection: `Stem() string`, `StemFull() string`, `Ext() string`, `ExtNoDot() string`, `HasExt(ext string) bool`, `Slug() string`, `Base() string`, `Dir() string`, `Split() (string, string)`.
   - Navigation: `Parent() *FolderInfo`, `ParentPath() string`, `ParentFolderName() string`, `Up() *PathInfo`, `UpN(levels int) *PathInfo`, `Cd(relPath string) *PathInfo`, `Sub(relPath string) *PathInfo`, `Join(elems ...string) *PathInfo`.
   - Finding & Filtering: `Find(pattern string) ResultSlice[string]`, `FindFiles(pattern string) ResultSlice[*FileInfo]`, `FindFolders(pattern string) ResultSlice[*FolderInfo]`, `Filter(filterFn FileFilterFunc) ResultSlice[string]`.
   - Looping & Walking: `Walk(...)`, `WalkFiles(...)`, `WalkFolders(...)`, `AllFiles() ResultSlice[string]`, `AllDirectories() ResultSlice[string]`, `Subfolders() ResultSlice[*FolderInfo]`, `Files() ResultSlice[*FileInfo]`, `FolderNames() ResultSlice[string]`, `FileNames() ResultSlice[string]`.
   - Status & Casting: `IsExists() bool`, `Exists() bool`, `IsDir() bool`, `IsDirectory() bool`, `IsFile() bool`, `IsAbs() bool`, `IsRel() bool`, `AsFolder() *FolderInfo`, `AsFile() *FileInfo`.
2. Enhance `FolderInfo` (`folder_info.go`) with:
   - `Normalize() *FolderInfo`
   - `Find(pattern string) ResultSlice[string]`
   - `FindFiles(pattern string) ResultSlice[*FileInfo]`
   - `FindFolders(pattern string) ResultSlice[*FolderInfo]`
   - `Filter(filterFn FileFilterFunc) ResultSlice[string]`
   - `Up() *FolderInfo`
   - `Cd(relPath string) *FolderInfo`
3. Enhance `FileInfo` (`file_info.go`) with:
   - `Normalize() *FileInfo`
   - `StemFull() string`
   - `Slug() string`
   - `Up() *FolderInfo`
4. Add unit test coverage across `path_info_test.go`, `folder_info_test.go`, and `file_info_test.go`.
5. Verify package compiles and tests pass: `go test ./pkg/fileutil/... -v`.
