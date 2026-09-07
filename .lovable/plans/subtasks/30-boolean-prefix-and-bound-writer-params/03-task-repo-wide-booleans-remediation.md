# Subtask 30.3: Repo-Wide Boolean Prefix Standardization

## Context
Enforce `is`/`has` prefixes across all remaining Go packages.

## Target Files
1. `04-code/golang/pkg/applogger/config.go`, `console_sink.go`, `applogger_test.go`:
   - Rename `UseJSON` / `useJSON` to `IsUseJSON` / `isUseJSON`.
2. `04-code/golang/pkg/logger/options.go`:
   - Rename parameter `enabled bool` to `isEnabled bool`.
3. `04-code/golang/pkg/streamwriter/async_writer.go`:
   - Rename `DropOnFull bool` to `IsDropOnFull bool`.
4. `04-code/golang/pkg/streamwriter/async_writer_test.go`:
   - Rename `closed bool` to `isClosed bool`.
   - Rename `failNext bool` to `isFailNext bool`.
5. `04-code/golang/pkg/streamwriter/bytes.go`:
   - Rename `status bool` to `isStatus bool`.
6. `04-code/golang/pkg/baseenumer/helpers.go`:
   - Rename return param `ok bool` to `isOk bool`.
7. `04-code/golang/pkg/enum/bytetype/variant_test.go`:
   - Rename `success bool` to `isSuccess bool`.
8. `04-code/golang/pkg/typecast/cast_test.go`:
   - Rename `dest bool` to `isDest bool`.

## Verification
- `go test -v -C 04-code/golang -count=1 ./...`
