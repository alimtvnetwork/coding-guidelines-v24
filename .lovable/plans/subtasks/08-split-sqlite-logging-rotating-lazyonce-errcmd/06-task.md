# Subtask 06: Comprehensive Tests, Verification & CI/CD Gate Enforcement

## Parent Plan
`.lovable/plans/completed/08-split-sqlite-logging-rotating-lazyonce-errcmd.md`

## Target Files
- `04-code/golang/pkg/applogger/sqlitelogger/sqlitelogger_test.go`
- `04-code/golang/pkg/applogger/rotating_file_sink_test.go`
- `04-code/golang/pkg/lazyonce/lazyonce_test.go`
- `04-code/golang/pkg/errcmd/errcmd_test.go`

## Instructions
1. Write unit tests for `SplitDBManager` testing main log writes, task-specific log writes, task listing, and querying.
2. Write unit tests for `RotatingFileSink` testing rotation threshold, max backup pruning, and archiving.
3. Write unit tests for `LazyOnce`, `LazyOnce1`, and `LazyOnce2` testing concurrency, single execution guarantee, and error caching.
4. Write unit tests for `errcmd` testing PowerShell/Bash script generation, runner execution, safe defer, and command logging.
5. Run `go test ./pkg/...` across all packages.
6. Run `python 03-ai-scripts/06-cicd-local-runner.py` and verify all 36 quality gates exit 0.
