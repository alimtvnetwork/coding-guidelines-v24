# Master Plan 21: Repo-Wide Enum Isolation — Dedicated Files and Packages

## 1. Problem Statement & User Mandate
- **User Mandate:** "each enum should be in it;s own file or package"
- **Identified Violation:** In `04-code/golang/pkg/errtype/base_enum.go`, multiple enums (`ProcessStateType` and `LogLevelType`) were co-located in a single 372-line file, and their constants were scattered into `consts.go`.
- **Repo-wide Scope:** Several other enums across `pkg/fileutil`, `pkg/appfault`, and `pkg/applogger` share files with other types or constants.
- **Goal:**
  1. In `pkg/errtype/`:
     - Move `ProcessStateType` and its constants to `process_state_type.go`.
     - Move `LogLevelType` and its constants to `log_level_type.go`.
     - Rename `base_enum.go` to `base_enumer.go` containing only interface definitions and forwarders.
     - Split tests into `process_state_type_test.go`, `log_level_type_test.go`, and `base_enumer_test.go`.
  2. In `pkg/enum/`:
     - Create standalone package `04-code/golang/pkg/enum/processstatetype/` matching the architecture of `pkg/enum/openfiletype/` and `pkg/enum/logleveltype/`.
  3. In `pkg/fileutil/`:
     - Rename `perm_types.go` -> `file_perm_type.go` (and test file to `file_perm_type_test.go`).
     - Extract `FileOpType` from `types.go` into its own file `file_op_type.go` (and `file_op_type_test.go`).
     - Extract `FileWriteModeType` from `writer_appender.go` into `file_write_mode_type.go`.
  4. In `pkg/appfault/`:
     - Extract `SeverityType` from `types.go` into `severity_type.go` (and `severity_type_test.go`).
     - Rename `priority.go` to `priority_type.go`.
  5. In `pkg/applogger/`:
     - Extract `DriverType` from `consts.go` into `driver_type.go`.

---

## 2. Task-Specific Rule Set
1. **Rule 1 (1:1 Enum File / Package Isolation):** Every enum type must reside in its own dedicated file matching its snake_case type name (e.g. `process_state_type.go`, `log_level_type.go`, `file_op_type.go`, `severity_type.go`).
2. **Rule 2 (Self-Contained Constants):** The enum's constants, lookup maps, parsers, and marshalers must be co-located with the enum in that file.
3. **Rule 3 (Zero Breaking Changes):** All public symbols, type aliases, and constants must remain exported and identical so all callers continue compiling without modification.
4. **Rule 4 (Canonical Function Size):** Functions must strictly respect `<= 15 lines`, blank line after closing brace if followed by code.
5. **Rule 5 (Strict Relative Git Paths):** All markdown links and plans must use strictly relative git paths.

---

## 3. Subtask Breakdown

- `.lovable/plans/subtasks/21-repo-wide-enum-isolation/01-task-errtype-enum-isolation-and-processstatetype-package.md`
- `.lovable/plans/subtasks/21-repo-wide-enum-isolation/02-task-fileutil-appfault-applogger-enum-isolation.md`

---

## 4. Verification Plan
- Unit tests for all modified and created packages: `go test -C 04-code/golang -count=1 -v ./...`.
- Local CI runner: `python 03-ai-scripts/06-cicd-local-runner.py --all` (all 31 gates).
- Metadata sync: `npm run sync` and `node scripts/sync-check.mjs`.
