# Subtask 04: Cross-Platform Command Execution, PowerShell/Bash Integration & Safe Defer

## Parent Plan
`.lovable/plans/completed/08-split-sqlite-logging-rotating-lazyonce-errcmd.md`

## Target Files
- `04-code/golang/pkg/errcmd/script_builder.go`
- `04-code/golang/pkg/errcmd/powershell.go`
- `04-code/golang/pkg/errcmd/bash.go`
- `04-code/golang/pkg/errcmd/safedefer.go`

## Instructions
1. Implement PowerShell script builder with proper flags (`-NoProfile`, `-NonInteractive`, `-ExecutionPolicy`, `Bypass`).
2. Implement Bash/Shell script builder (`-c`).
3. Implement host OS script detection using `runtime.GOOS`.
4. Implement safe defer cleanup utilities inspired by `errdefer` (`SafeClose`, `SafeCloseWithErr`, `CapturePanic`).
5. Ensure strict adherence to coding guidelines: <= 15 lines per function, `*appfault.AppError` return type.
