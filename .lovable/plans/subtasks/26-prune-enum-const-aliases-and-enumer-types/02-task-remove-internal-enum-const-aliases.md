# Subtask 26.2: Remove Internal Enum Const Aliases

## Objective
Remove redundant `const` alias blocks that duplicate enum variant values inside canonical enum packages.

## Changes
- `04-code/golang/pkg/enum/filepermtype/variant.go`: delete lines 48–74 (`FilePermNone = None ...`).
- `04-code/golang/pkg/enum/fileoptype/variant.go`: delete lines 32–42 (`FileOpInvalid = Invalid ...`).
- `04-code/golang/pkg/enum/filewritemodetype/variant.go`: delete lines 26–31 (`FileWriteModeInvalid = Invalid ...`).
- `04-code/golang/pkg/enum/severitytype/variant.go`: delete lines 28–35 (`SeverityUnknown = Unknown ...`).
- `04-code/golang/pkg/enum/prioritytype/variant.go`: delete lines 27–33 (`PriorityUnknown = Unknown ...`).
- `04-code/golang/pkg/errtype/logleveltype/variant.go`: delete lines 26–32 (`LogLevelDebug = Debug ...`).

## Acceptance Criteria
- `go test -C 04-code/golang -v ./pkg/enum/... ./pkg/errtype/logleveltype` passes.
