# Master Plan 30: Boolean Naming Standardization (`is`/`has` Prefixes) & BoundFileWriter Parameter Actions

## Overview
Enforce repository-wide boolean naming conventions where every boolean field, variable, and parameter must strictly have an `is` or `has` prefix. Additionally, expand `04-code/golang/pkg/fileutil/bound_file_writer.go` so that the parameters shown in the user screenshot (`mode`, `perm`, `isSyncOnWrite`, `isAutoClose`) have dedicated constructor functions, query methods, and action mutators that accept and act on those parameters.

## Checklist

- [x] Step 30.1: Refactor `BoundFileWriter` in `04-code/golang/pkg/fileutil/bound_file_writer.go` to rename booleans to `isSyncOnWrite` / `isAutoClose`, add `IsSyncOnWrite()` / `IsAutoClose()`, add parameterized constructors (`NewBoundFileWriterConfig`, `NewBoundFileWriterWithMode`, `NewBoundFileWriterWithPerm`, `NewBoundFileWriterWithSync`, `NewBoundFileWriterWithAutoClose`), and action methods (`With*`, `Enable*`, `Disable*`).
- [x] Step 30.2: Refactor `04-code/golang/pkg/fileutil/writer_appender.go` and `results.go` for boolean prefixes (`IsSyncOnWrite`, `isSyncOnWrite`, `isAutoSync`, `isSuccess`).
- [x] Step 30.3: Remediate all remaining non-is/has booleans across Go packages (`applogger`, `logger`, `streamwriter`, `baseenumer`, `bytetype`, `typecast`).
- [x] Step 30.4: Run full verification suite (Go tests, code formatter, sequence integrity, CI/CD local runner, sync-check) and move plan to completed.

## Subtasks
- [01-task-bound-file-writer-params-and-booleans.md](../subtasks/30-boolean-prefix-and-bound-writer-params/01-task-bound-file-writer-params-and-booleans.md)
- [02-task-fileutil-writer-appender-booleans.md](../subtasks/30-boolean-prefix-and-bound-writer-params/02-task-fileutil-writer-appender-booleans.md)
- [03-task-repo-wide-booleans-remediation.md](../subtasks/30-boolean-prefix-and-bound-writer-params/03-task-repo-wide-booleans-remediation.md)
- [04-task-quality-gates-and-verification.md](../subtasks/30-boolean-prefix-and-bound-writer-params/04-task-quality-gates-and-verification.md)
