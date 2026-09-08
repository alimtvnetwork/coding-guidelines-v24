# Subtask 02: errcmd Live Line Streaming & Process Context

## 1. Goal
Enhance `CommandRunner` in `04-code/golang/pkg/errcmd/` to support real-time line-by-line streaming callbacks (`WithStdoutHandler`, `WithStderrHandler`) and environment/directory context (`WithEnv`, `WithCwd`).

**Status:** ✅ Completed

## 2. Target Files
- `04-code/golang/pkg/errcmd/command_runner.go`
- `04-code/golang/pkg/errcmd/errcmd_test.go`

## 3. Detailed Specifications
1. **Streaming Handlers**:
   - `WithStdoutHandler(handler func(line string)) *CommandRunner`
   - `WithStderrHandler(handler func(line string)) *CommandRunner`
   - During `Run(ctx context.Context)`:
     - Pipes stdout and stderr through `bufio.Scanner` routines.
     - Calls handlers concurrently as lines arrive.
     - Concurrently records lines into `TaskLogger` if attached.
     - Preserves full output buffer for `CommandResult`.
2. **Process Context**:
   - `WithEnv(env map[string]string) *CommandRunner`: merges environment variables into `cmd.Env`.
   - `WithCwd(dir string) *CommandRunner`: sets `cmd.Dir`.
3. **Coding Guidelines Compliance**:
   - Functions <= 15 lines. Extract stream processing helpers.
   - All booleans positive (`hasHandler`, `isDone`).
