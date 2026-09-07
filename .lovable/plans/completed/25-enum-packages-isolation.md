# Plan 25: Dedicated Enum Packages Isolation (`fileoptype`, `filewritemodetype`, `severitytype`, `prioritytype`)

## Context & Motivation
Following the successful creation of canonical enum packages `bytetype`, `filepermtype`, `logleveltype`, `openfiletype`, and `processstatetype`, four additional enums remain embedded or loose within domain packages:
1. `FileOpType` in `04-code/golang/pkg/fileutil/file_op_type.go`
2. `FileWriteModeType` in `04-code/golang/pkg/fileutil/file_write_mode_type.go`
3. `SeverityType` in `04-code/golang/pkg/appfault/severity_type.go`
4. `PriorityType` in `04-code/golang/pkg/appfault/priority_type.go`

This plan migrates each enum into its own dedicated package under `04-code/golang/pkg/enum/` with complete interfaces (`BaseEnumer`, `ByteEnumer`, `NumberEnumer`, `json.Marshaler`, `json.Unmarshaler`), table-driven tests, clean package documentation, zero circular dependencies, and seamless backward-compatibility forwarders.

---

## Architectural Constraints
1. **Zero Circular Imports:**
   - `pkg/enum/fileoptype`: imports `baseenumer`, `openfiletype`, `errtype`, `result`.
   - `pkg/enum/filewritemodetype`: imports `baseenumer`, `errtype`, `result`.
   - `pkg/enum/severitytype`: imports `baseenumer` ONLY. Must NOT import `pkg/result` or `pkg/appfault` to avoid cycle with `appfault`.
   - `pkg/enum/prioritytype`: imports `baseenumer` ONLY. Must NOT import `pkg/result` or `pkg/appfault` to avoid cycle with `appfault`.
2. **Function Sizing:** Strict `<= 15 lines` per function.
3. **Boolean Principles:** Implicit booleans only (`if isCondition`); no explicit `== true`; no mixed polarity.
4. **Strict Relative Git Paths:** Zero absolute filesystem paths or `file:///` URIs.

---

## Subtask Breakdown
- [x] Subtask 1: Create `04-code/golang/pkg/enum/fileoptype` (`variant.go`, `vars.go`, `variant_test.go`, `readme.md`)
- [x] Subtask 2: Create `04-code/golang/pkg/enum/filewritemodetype` (`variant.go`, `vars.go`, `variant_test.go`, `readme.md`)
- [x] Subtask 3: Create `04-code/golang/pkg/enum/severitytype` (`variant.go`, `vars.go`, `variant_test.go`, `readme.md`)
- [x] Subtask 4: Create `04-code/golang/pkg/enum/prioritytype` (`variant.go`, `vars.go`, `variant_test.go`, `readme.md`)
- [x] Subtask 5: Configure forwarders in `pkg/fileutil` and `pkg/appfault`, run tests, run CI/CD gates.
