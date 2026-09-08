# Plan 14: Reduce Baseenumer Code and Remove Enum Result Wrap

> **Status:** Completed  
> **Slug:** 14-reduce-baseenumer-and-remove-enum-result-wrap  
> **Target Module:** `04-code/golang/pkg/baseenumer`, `04-code/golang/pkg/enum/**`, `04-code/golang/pkg/fileutil`, `04-code/golang/pkg/logger`  
> **Date:** 2026-09-09  

---

## 1. Problem Statement & Root Cause

1. **Import Cycle Risk:**
   Seven packages in `04-code/golang/pkg/enum/` (`processstatetype`, `bytetype`, `fileoptype`, `filepermtype`, `filewritemodetype`, `logleveltype`, `openfiletype`) imported `coding-guidelines/common/pkg/errtype` and `coding-guidelines/common/pkg/result`.
   Because `pkg/result` imports `pkg/appfault`, and `pkg/appfault` imports enums (`severitytype`, `prioritytype`), any enum package importing `result` risked creating an unresolvable Go import cycle (`appfault` -> `enum` -> `result` -> `appfault`).
   In idiomatic Go, enum packages are foundational leaf packages and must never import high-level domain error/result wrappers.

2. **Redundant Boilerplate in `vars.go`:**
   Each enum package manually exported `variantMap = basicEnum.Map()` and wrote 15 lines of manual map lookup and error formatting boilerplate inside `func Parse(s string) Result`.
   `baseenumer.BasicIntegerEnum` and `baseenumer.BasicStringEnum` already encapsulated the variant map internally, but lacked direct `(V, bool)` and `ParseOrZero` convenience methods.

---

## 2. Task-Specific Rules & Constraints

1. **Zero High-Level Imports in Enums:** No enum package under `04-code/golang/pkg/enum/` may import `pkg/result`, `pkg/errtype`, or `pkg/appfault`.
2. **Standard Parse Signature `(Variant, bool)`:** Every enum package `Parse(s string)` returns `(Variant, bool)`, matching `severitytype` and `prioritytype`.
3. **ParseOrZero / ParseOrInvalid Fallback:** Every enum package provides `ParseOrInvalid(s string) Variant` (or `ParseOrZero` / `ParseOrUnknown`) returning the zero variant if parsing fails.
4. **Baseenumer Encapsulation:** `basicEnum.Parse(s)` and `basicEnum.ParseOrZero(s)` perform the lookup internally without requiring `variantMap = basicEnum.Map()` in each `vars.go`.
5. **Downstream Non-Breaking Migration:** Callers of enum `Parse` (`pkg/logger/level.go`, `pkg/fileutil/file_op_type.go`, unit tests) are updated to handle `(Variant, bool)` cleanly.

---

## 3. Subtask Decomposition & Execution Record

1. [Subtask 01](.lovable/plans/subtasks/14-reduce-baseenumer-and-remove-enum-result-wrap/01-baseenumer-parse-helpers.md): Enhanced `BasicIntegerEnum` and `BasicStringEnum` with `Parse(s string) (V, bool)`, `ParseOrZero(s string) V`, and `ParseErr(s string) (V, error)`. (✅ Completed)
2. [Subtask 02](.lovable/plans/subtasks/14-reduce-baseenumer-and-remove-enum-result-wrap/02-refactor-processstatetype.md): Refactored `processstatetype` removing `result.Wrap`, updating `Parse` and tests. (✅ Completed)
3. [Subtask 03](.lovable/plans/subtasks/14-reduce-baseenumer-and-remove-enum-result-wrap/03-refactor-system-enums.md): Refactored `bytetype`, `filepermtype`, and `openfiletype` into leaf packages. (✅ Completed)
4. [Subtask 04](.lovable/plans/subtasks/14-reduce-baseenumer-and-remove-enum-result-wrap/04-refactor-operation-and-logger-enums.md): Refactored `fileoptype`, `filewritemodetype`, `logleveltype`, and downstream packages (`pkg/logger`, `pkg/fileutil`). (✅ Completed)
5. [Subtask 05](.lovable/plans/subtasks/14-reduce-baseenumer-and-remove-enum-result-wrap/05-clean-priority-severity-and-verify.md): Cleaned `prioritytype` and `severitytype` boilerplate, verified `go test ./...` and `06-cicd-local-runner.py` (36/36 gates green). (✅ Completed)

---

## 4. Quality Verification
- `go test ./pkg/... -count=1`: PASS (24 packages)
- `go test ./examples/... -count=1`: PASS
- `python 03-ai-scripts/06-cicd-local-runner.py`: PASS (36/36 gates)
