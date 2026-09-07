# Master Plan 23: Errtype Enum Folder Isolation (`logleveltype` & `processstatetype`)

## 1. Problem Statement & User Mandate
- **User Mandate:** "loglevel and process state should be on it's own folder fix it please" with screenshot `.lovable/assets/enum/07-errtype-folder-structure.png` showing `log_level_type.go` and `process_state_type.go` residing directly in the root of `04-code/golang/pkg/errtype/`.
- **Current State in `04-code/golang/pkg/errtype/`:**
  - Root directory has `log_level_type.go`, `log_level_type_test.go`, `process_state_type.go`, and `process_state_type_test.go`.
  - Non-error enum types (`LogLevelType` and `ProcessStateType`) clutter the `errtype` error variation root folder.
- **Architectural Solution:**
  1. Move `LogLevelType` and its test suite into dedicated subfolder `04-code/golang/pkg/errtype/logleveltype/` (`variant.go`, `variant_test.go`, `readme.md`) as package `logleveltype`.
  2. Move `ProcessStateType` and its test suite into dedicated subfolder `04-code/golang/pkg/errtype/processstatetype/` (`variant.go`, `variant_test.go`, `readme.md`) as package `processstatetype`.
  3. Remove `log_level_type.go`, `log_level_type_test.go`, `process_state_type.go`, and `process_state_type_test.go` from the root of `04-code/golang/pkg/errtype/`.
  4. Update callers in `04-code/golang/examples/converter_and_enum_examples.go` and `04-code/golang/examples/converter_and_enum_examples_test.go` to import the new subpackages.
  5. Update `04-code/golang/pkg/errtype/base_enumer_test.go` to test `ToEnum` using `Variation` and subpackages.
  6. Update `04-code/golang/pkg/errtype/readme.md` documenting the subfolder architecture.

---

## 2. Task-Specific Rule Set
1. **Rule 1 (Subfolder & Package Suffix):** Both new subfolders MUST end with `type` suffix: `logleveltype` and `processstatetype`. Package declarations MUST be `package logleveltype` and `package processstatetype`.
2. **Rule 2 (Clean Root):** Zero loglevel or processstate files in root `04-code/golang/pkg/errtype/`.
3. **Rule 3 (No Circular Imports):** Subpackages `pkg/errtype/logleveltype` and `pkg/errtype/processstatetype` MUST NOT import `result`, `appfault`, or `errtype`. They import only standard library and `baseenumer`.
4. **Rule 4 (Canonical Function Size):** Functions must strictly adhere to `<= 15 lines`. Exactly one empty line between functions.
5. **Rule 5 (Strict Relative Git Paths):** All markdown links and references must use strictly relative git paths.

---

## 3. Subtask Breakdown
- Subtask 01: `.lovable/plans/subtasks/23-errtype-enum-folder-isolation/01-task-isolate-logleveltype-and-processstatetype.md`
- Subtask 02: `.lovable/plans/subtasks/23-errtype-enum-folder-isolation/02-task-update-callers-and-docs.md`

---

## 4. Verification Plan
- Unit test verification: `go test -C 04-code/golang -v ./pkg/errtype/... ./examples/...`.
- All Go packages test verification: `go test -C 04-code/golang -count=1 ./...`.
- Local CI runner: `python 03-ai-scripts/06-cicd-local-runner.py --all` (all 31 gates green).
- Sync check: `npm run sync` and `node scripts/sync-check.mjs`.
