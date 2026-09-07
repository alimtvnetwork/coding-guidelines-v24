# Subtask 27.2: FileUtil Direct Result Types and Domain Functions

## Objective
Define missing direct concrete result types in `pkg/fileutil/types.go` and update domain functions in `file_perm_type.go`, `file_op_type.go`, `file_write_mode_type.go`, `advanced.go`, and `file_namespace.go` to return these direct types instead of generic `result.Wrap[...]`.

## Target Files
1. `04-code/golang/pkg/fileutil/types.go`: Define `FilePermResult`, `FileOpResult`, `FileWriteModeResult`, `FileOpenModeResult`, and `FileWriterResult`.
2. `04-code/golang/pkg/fileutil/file_perm_type.go`: Define `FilePermResult = filepermtype.Result` and update `ParsePerm(...) FilePermResult`.
3. `04-code/golang/pkg/fileutil/file_op_type.go`: Define `FileOpResult = fileoptype.Result` and update `ParseFileOp(...) FileOpResult`.
4. `04-code/golang/pkg/fileutil/file_write_mode_type.go`: Define `FileWriteModeResult = filewritemodetype.Result` and update `ParseFileWriteMode(...) FileWriteModeResult`.
5. `04-code/golang/pkg/fileutil/advanced.go`: Update `NewFileWriter(...) FileWriterResult`.
6. `04-code/golang/pkg/fileutil/file_namespace.go`: Update `StreamWriterAny`, `Append`, `Truncate` to return `FileWriterResult`.

## Acceptance Criteria
- Zero occurrences of `result.Wrap[FileWriteModeType]`, `result.Wrap[FilePermType]`, `result.Wrap[FileOpType]`, or `result.Wrap[*streamwriter.PluggableWriter[any]]` in `pkg/fileutil/`.
- All unit tests in `pkg/fileutil/...` pass without errors.
