# Subtask 02: Fileutil, Appfault & Applogger Enum Isolation

## Objective
Isolate enums across `pkg/fileutil`, `pkg/appfault`, and `pkg/applogger` into dedicated files matching their exact snake_case type names.

## Target Files
- `04-code/golang/pkg/fileutil/file_perm_type.go` (renamed from `perm_types.go`)
- `04-code/golang/pkg/fileutil/file_perm_type_test.go`
- `04-code/golang/pkg/fileutil/file_op_type.go` (extracted from `types.go`)
- `04-code/golang/pkg/fileutil/file_op_type_test.go`
- `04-code/golang/pkg/fileutil/file_write_mode_type.go` (extracted from `writer_appender.go`)
- `04-code/golang/pkg/fileutil/types.go` (cleaned up)
- `04-code/golang/pkg/fileutil/writer_appender.go` (cleaned up)
- `04-code/golang/pkg/appfault/severity_type.go` (extracted from `types.go`)
- `04-code/golang/pkg/appfault/severity_type_test.go`
- `04-code/golang/pkg/appfault/types.go` (cleaned up)
- `04-code/golang/pkg/appfault/priority_type.go` (renamed from `priority.go`)
- `04-code/golang/pkg/appfault/priority_type_test.go` (renamed from `priority_test.go`)
- `04-code/golang/pkg/applogger/driver_type.go` (extracted from `consts.go`)
- `04-code/golang/pkg/applogger/consts.go` (cleaned up)

## Implementation Steps
1. In `pkg/fileutil`:
   - Rename `perm_types.go` to `file_perm_type.go`.
   - Extract `FileOpType`, its constants, names, methods, and flags into `file_op_type.go`.
   - Extract `FileWriteModeType` into `file_write_mode_type.go`.
2. In `pkg/appfault`:
   - Extract `SeverityType`, its constants, names, methods, and JSON marshalers into `severity_type.go`.
   - Rename `priority.go` to `priority_type.go` and `priority_test.go` to `priority_type_test.go`.
3. In `pkg/applogger`:
   - Extract `DriverType` and its constants into `driver_type.go`.
4. Run Go tests across all affected packages: `go test -C 04-code/golang -v ./pkg/fileutil ./pkg/appfault ./pkg/applogger`.
