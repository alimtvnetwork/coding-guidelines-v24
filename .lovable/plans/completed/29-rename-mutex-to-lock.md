# Master Plan 29: Rename `mu` and `mutex` to `lock` Repo-Wide

## Overview
Standardize concurrency naming by renaming all internal struct fields named `mu` and all identifier suffixes named `mutex`/`Mu` to `lock` across all 7 affected Go packages in `04-code/golang/` (`fileutil`, `streamwriter`, `applogger`, `appwriter`, `appfault`, `appfaults`, `regexnew`).

## Checklist

- [x] Step 29.1: Refactor `pkg/fileutil` (`bound_file_writer.go`, `writer_appender.go`, `locker.go`, `append_ops.go`, `locked_operations.go`, `write.go`, `readme.md`).
- [x] Step 29.2: Refactor `pkg/streamwriter` (`writer.go`, `async_writer.go`, `locked_streamer.go`, `logger.go`, `mutex.go`, `streamwriter_test.go`, `async_writer_test.go`).
- [x] Step 29.3: Refactor logging, writer, fault, and regex packages (`pkg/applogger`, `pkg/appwriter`, `pkg/appfault`, `pkg/appfaults`, `pkg/regexnew`).
- [x] Step 29.4: Full verification suite (Go tests, code formatter, CI/CD local runner, sync-check) and move plan to completed.

## Subtasks
- [01-task-fileutil-rename-mutex-to-lock.md](../subtasks/29-rename-mutex-to-lock/01-task-fileutil-rename-mutex-to-lock.md)
- [02-task-streamwriter-rename-mutex-to-lock.md](../subtasks/29-rename-mutex-to-lock/02-task-streamwriter-rename-mutex-to-lock.md)
- [03-task-logging-writer-fault-rename-mutex-to-lock.md](../subtasks/29-rename-mutex-to-lock/03-task-logging-writer-fault-rename-mutex-to-lock.md)
- [04-task-quality-gates-and-verification.md](../subtasks/29-rename-mutex-to-lock/04-task-quality-gates-and-verification.md)
