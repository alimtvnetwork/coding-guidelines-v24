# Subtask 35.4: Namespace Integration & Path Bridges

## Status: Complete

## Goal
Connect `FolderInfo` and `FileInfo` models to existing root namespaces `Path`, `File`, `New`, `pathInfoNamespace`, and `PathWrapper`.

## Target Files
- `04-code/golang/pkg/fileutil/path_namespace.go`
- `04-code/golang/pkg/fileutil/file_namespace.go`
- `04-code/golang/pkg/fileutil/path_info.go`

## Acceptance Criteria
1. In `path_namespace.go`:
   - Add methods to `pathNamespace`:
     - `Folder(path string) *FolderInfo`
     - `FolderInfo(path string) *FolderInfo`
     - `File(path string) *FileInfo`
     - `FileInfo(path string) *FileInfo`
     - `Inspect(path string) *PathInfo`
   - Add methods to `PathWrapper`:
     - `Folder() *FolderInfo`
     - `File() *FileInfo`
2. In `path_info.go`:
   - Add methods to `pathInfoNamespace`:
     - `Folder(path string) *FolderInfo`
     - `File(path string) *FileInfo`
3. In `file_namespace.go`:
   - Add methods to `fileNamespace`:
     - `Folder(path string) *FolderInfo`
     - `FolderInfo(path string) *FolderInfo`
     - `FileInfo(path string) *FileInfo`
   - Add methods to `fileNewCreator`:
     - `Folder(path string) *FolderInfo`
     - `FolderInfo(path string) *FolderInfo`
     - `File(path string) *FileInfo`
     - `FileInfo(path string) *FileInfo`
4. Functions $\le 15$ lines.
5. `cd 04-code/golang && go test ./pkg/fileutil/...` passes cleanly.
