# Subtask 01: Errtype Enum Isolation & ProcessStateType Package

## Objective
Isolate `ProcessStateType` and `LogLevelType` from `base_enum.go` in `04-code/golang/pkg/errtype/` into dedicated files, create `pkg/enum/processstatetype/`, and clean up `base_enum.go` into `base_enumer.go`.

## Target Files
- `04-code/golang/pkg/errtype/process_state_type.go`
- `04-code/golang/pkg/errtype/process_state_type_test.go`
- `04-code/golang/pkg/errtype/log_level_type.go`
- `04-code/golang/pkg/errtype/log_level_type_test.go`
- `04-code/golang/pkg/errtype/base_enumer.go` (renamed from `base_enum.go`)
- `04-code/golang/pkg/errtype/base_enumer_test.go` (renamed from `base_enum_test.go`)
- `04-code/golang/pkg/errtype/consts.go`
- `04-code/golang/pkg/enum/processstatetype/variant.go`
- `04-code/golang/pkg/enum/processstatetype/vars.go`
- `04-code/golang/pkg/enum/processstatetype/variant_test.go`
- `04-code/golang/pkg/enum/processstatetype/readme.md`

## Implementation Steps
1. Create `04-code/golang/pkg/enum/processstatetype/` with standard enum architecture (`variant.go`, `vars.go`, `variant_test.go`, `readme.md`).
2. In `04-code/golang/pkg/errtype/process_state_type.go`, define `ProcessStateType` with its constants, methods, JSON marshaling, registry, and parser.
3. In `04-code/golang/pkg/errtype/log_level_type.go`, define `LogLevelType` with its constants, methods, JSON marshaling, registry, and parser.
4. Clean `04-code/golang/pkg/errtype/consts.go` removing `ProcessState*` and `LogLevel*` constants so they live exclusively in their respective enum files.
5. In `04-code/golang/pkg/errtype/base_enumer.go`, retain only `baseenumer` interface aliases and `ToEnum`.
6. Split tests into `process_state_type_test.go`, `log_level_type_test.go`, and `base_enumer_test.go`.
7. Verify all tests pass: `go test -C 04-code/golang -v ./pkg/errtype ./pkg/enum/processstatetype`.
