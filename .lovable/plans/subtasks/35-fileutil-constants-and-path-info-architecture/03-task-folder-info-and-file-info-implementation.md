# Subtask 35.3: FolderInfo, FileInfo, and PathInfo Implementation

## Status: Complete

## Goal
Implement stateful .NET-style filesystem inspection domain models in `04-code/golang/pkg/fileutil/`:
- `folder_info.go`: `FolderInfo`
- `file_info.go`: `FileInfo`
- `path_info_obj.go`: `PathInfo`

## Target Files
- `04-code/golang/pkg/fileutil/folder_info.go`
- `04-code/golang/pkg/fileutil/file_info.go`
- `04-code/golang/pkg/fileutil/path_info_obj.go`

## Acceptance Criteria
1. `FolderInfo` in `folder_info.go`:
   - `NewFolderInfo(path string) *FolderInfo`
   - `Path() string`, `Name() string`, `IsExists() bool`, `Exists() bool`, `Stat() FileInfoResult`, `Abs() *FolderInfo`
   - `Parent() *FolderInfo`, `ParentPath() string`, `ParentFolderName() string`
   - `Subfolders() ResultSlice[*FolderInfo]`, `Directories() ResultSlice[*FolderInfo]`, `FolderNames() ResultSlice[string]`
   - `Files() ResultSlice[*FileInfo]`, `FileNames() ResultSlice[string]`
   - `Subfolder(name string) *FolderInfo`, `File(name string) *FileInfo`
   - `ParentFolderFiles() ResultSlice[*FileInfo]`, `ParentFolderDirectories() ResultSlice[*FolderInfo]`
   - `Walk(fn func(path string, isDir bool) *appfault.AppError) *appfault.AppError`
   - `WalkFiles(fn func(filePath string) *appfault.AppError) *appfault.AppError`
   - `WalkDirectories(fn func(dirPath string) *appfault.AppError) *appfault.AppError`
   - `AllFiles() ResultSlice[string]`, `AllDirectories() ResultSlice[string]`
   - `EnsureDir(perm FilePermType) BoolResult`, `Delete() BoolResult`, `RemoveAll() BoolResult`
2. `FileInfo` in `file_info.go`:
   - `NewFileInfo(path string) *FileInfo`
   - `Path() string`, `Name() string`, `Extension() string`, `Ext() string`, `ExtNoDot() string`, `Stem() string`, `Size() int64`
   - `Folder() *FolderInfo`, `Directory() *FolderInfo`, `ParentFolderName() string`
   - `IsExists() bool`, `Exists() bool`, `Stat() FileInfoResult`, `Abs() *FileInfo`
   - `ReadBytes() BytesResult`, `ReadString() StringResult`, `WriteString(content string, perm FilePermType) BoolResult`, `Delete() BoolResult`
3. `PathInfo` in `path_info_obj.go`:
   - `NewPathInfo(path string) *PathInfo`
   - `Path() string`, `Name() string`, `IsExists() bool`, `IsDir() bool`, `IsFile() bool`
   - `AsFolder() *FolderInfo`, `AsFile() *FileInfo`
4. Functions strictly $\le 15$ lines.
5. `cd 04-code/golang && go test ./pkg/fileutil/...` passes cleanly.
