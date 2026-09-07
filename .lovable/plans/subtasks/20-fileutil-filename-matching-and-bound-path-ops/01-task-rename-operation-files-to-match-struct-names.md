# Subtask 01: Rename Operation Files to Match Struct Names

## Objective
Align file names with their struct definitions in `04-code/golang/pkg/fileutil/` so developers and AI agents can find structs instantly:
- `file_append.go` -> `append_ops.go`
- `file_append_test.go` -> `append_ops_test.go`
- `file_open.go` -> `open_ops.go`
- `file_create.go` -> `create_ops.go`
- `file_read.go` -> `read_ops.go`
- `file_write.go` -> `write_ops.go`

## Target Files
- `04-code/golang/pkg/fileutil/append_ops.go` (renamed from `file_append.go`)
- `04-code/golang/pkg/fileutil/append_ops_test.go` (renamed from `file_append_test.go`)
- `04-code/golang/pkg/fileutil/open_ops.go` (renamed from `file_open.go`)
- `04-code/golang/pkg/fileutil/create_ops.go` (renamed from `file_create.go`)
- `04-code/golang/pkg/fileutil/read_ops.go` (renamed from `file_read.go`)
- `04-code/golang/pkg/fileutil/write_ops.go` (renamed from `file_write.go`)

## Implementation Instructions
1. Use `git mv` to rename the files so git history is preserved cleanly.
2. Check for any references to the old file names in comments or imports.
3. Run `go test -C 04-code/golang -v ./pkg/fileutil` to verify all tests pass after renaming.
