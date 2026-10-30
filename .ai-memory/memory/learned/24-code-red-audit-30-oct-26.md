# Code Red audit 30-Oct-26

Plan 11 asked for inverted success checks, string unions, magic strings, and query wrappers.

## What the audit found

- `04-code` has no `isSuccess === false` and no `!isSuccess`.
- Go failures stay on `*appfault.AppError` and `Result[T]`.
- Query logging for this repo stays in `02-spec/03-error-manage`. Do not add a second wrapper package beside `04-code/golang/pkg/applogger/syncbridge`.
- Enums still end with `Type`. Comparisons use the enum symbol.

## What an agent must do

1. A new success check uses a positively named boolean. A failure check uses `isFail` or `isFailure`.
2. A new status value is an enum ending in `Type`, in its own file.
3. A caught error logs the operation name and the key inputs, then returns `*appfault.AppError` or rethrows.
