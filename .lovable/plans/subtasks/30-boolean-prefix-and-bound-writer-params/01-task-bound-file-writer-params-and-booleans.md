# Subtask 30.1: BoundFileWriter Parameter Actions and Boolean Standardization

## Context
In `04-code/golang/pkg/fileutil/bound_file_writer.go`, rename booleans to `isSyncOnWrite` and `isAutoClose`, add query getters `IsSyncOnWrite()` / `IsAutoClose()`, and add constructors and action methods for all parameters shown in the user screenshot (`mode`, `perm`, `isSyncOnWrite`, `isAutoClose`).

## Target File
`04-code/golang/pkg/fileutil/bound_file_writer.go`

## Requirements
1. Struct fields:
   - `BoundFileWriterOptions.IsSyncOnWrite bool`
   - `BoundFileWriterOptions.IsAutoClose bool`
   - `BoundFileWriter.isSyncOnWrite bool`
   - `BoundFileWriter.isAutoClose bool`
2. Constructors acting on params:
   - `NewBoundFileWriter(path string) *BoundFileWriter` (defaults `isSyncOnWrite: false, isAutoClose: false`)
   - `NewBoundFileWriterWithMode(path string, mode FileWriteModeType) *BoundFileWriter`
   - `NewBoundFileWriterWithPerm(path string, perm FilePermType) *BoundFileWriter`
   - `NewBoundFileWriterWithSync(path string, isSyncOnWrite bool) *BoundFileWriter`
   - `NewBoundFileWriterWithAutoClose(path string, isAutoClose bool) *BoundFileWriter`
   - `NewBoundFileWriterConfig(path string, mode FileWriteModeType, perm FilePermType, isSyncOnWrite bool, isAutoClose bool) *BoundFileWriter`
3. Query methods:
   - `func (w *BoundFileWriter) IsSyncOnWrite() bool`
   - `func (w *BoundFileWriter) IsAutoClose() bool`
4. Action mutators:
   - `func (w *BoundFileWriter) SetSyncOnWrite(isSyncOnWrite bool) *BoundFileWriter`
   - `func (w *BoundFileWriter) SetAutoClose(isAutoClose bool) *BoundFileWriter`
   - `func (w *BoundFileWriter) WithMode(mode FileWriteModeType) *BoundFileWriter`
   - `func (w *BoundFileWriter) WithPerm(perm FilePermType) *BoundFileWriter`
   - `func (w *BoundFileWriter) WithSyncOnWrite(isSyncOnWrite bool) *BoundFileWriter`
   - `func (w *BoundFileWriter) WithAutoClose(isAutoClose bool) *BoundFileWriter`
   - `func (w *BoundFileWriter) EnableSyncOnWrite() *BoundFileWriter`
   - `func (w *BoundFileWriter) DisableSyncOnWrite() *BoundFileWriter`
   - `func (w *BoundFileWriter) EnableAutoClose() *BoundFileWriter`
   - `func (w *BoundFileWriter) DisableAutoClose() *BoundFileWriter`
5. Internal helpers:
   - Rename `closeAfter bool` in `writeInternal` and `appendInternal` to `isCloseAfter bool`.

## Verification
- `go test -v -C 04-code/golang -count=1 ./pkg/fileutil/...`
