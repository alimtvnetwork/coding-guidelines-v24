# Subtask 26.3: Remove Domain Const Aliases and Update Callers

## Objective
Remove redundant `const` alias blocks in domain packages (`pkg/fileutil`, `pkg/appfault`) and update all call sites to reference canonical enum packages directly.

## Changes
- In `04-code/golang/pkg/fileutil/`:
  - `file_perm_type.go`: delete `FilePermNone` ... `FilePermSetgidExec` const block.
  - `file_op_type.go`: delete `FileOpInvalid` ... `FileOpDelete` const block.
  - `file_write_mode_type.go`: delete `FileWriteModeInvalid` ... `FileWriteModeTruncate` const block.
  - `consts.go`: delete `FileOpenInvalid` ... `FileOpenReadWriteOrCreateOnly` const block.
  - Update internal callers: `parsers.go`, `stream.go`, `open_ops.go`, `append_ops.go`, `file_path_ops.go`, `advanced.go`, `file_namespace.go`.
  - Update tests: `write_test.go`, `append_ops_test.go`, `advanced_test.go`, `create_test.go`, `file_namespace_test.go`, `file_path_ops_test.go`, `file_perm_type_test.go`, `fileutil_test.go`, `locker_test.go`, `parsers_test.go`, `path_temp_test.go`, `path_namespace_test.go`, `stream_test.go`, `writer_appender_test.go`, `bound_file_writer_test.go`, `file_write_mode_type_test.go`, `file_op_type_test.go`.
- In `04-code/golang/pkg/appfault/`:
  - `severity_type.go`: delete `SeverityUnknown` ... `SeverityFatal` const block.
  - `priority_type.go`: delete `PriorityUnknown` ... `PriorityCritical` const block.
  - Update tests: `appfault_test.go`, `severity_type_test.go`, `priority_type_test.go`.
- In `04-code/golang/pkg/appwriter/`:
  - `file_writer.go`, `appwriter_test.go`: update to `filepermtype.Standard`, `openfiletype.ReadOnly`.
- In `04-code/golang/examples/`:
  - `converter_and_enum_examples.go`, `streamwriter_examples.go`: update to `filepermtype.Standard`, `filewritemodetype.*`.

## Acceptance Criteria
- Zero compilation errors.
- `go test -C 04-code/golang -v ./pkg/fileutil/... ./pkg/appfault/... ./pkg/appwriter/... ./examples/...` passes.
