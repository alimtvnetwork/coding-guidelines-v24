# Master Plan 26: Eliminate Redundant Enum Const Aliases & Prune Over-Engineered Enumer Types

## 1. Problem Statement & User Mandate
- **User Mandate:**
  - *"Why we are again creating const from ENUM values WHY?? FIx it everywhere please"*: Remove redundant `const` alias blocks that duplicate enum values from their canonical enum packages. Screenshot `.lovable/assets/enum/09-const-enum-redundancy.png` highlighted `04-code/golang/pkg/fileutil/file_perm_type.go` re-declaring `FilePermNone = filepermtype.None`, `FilePermStandard = filepermtype.Standard`, etc. Callers must import and use the canonical enum packages directly (`filepermtype.Standard`, `fileoptype.ReadOnly`, `filewritemodetype.Direct`, `openfiletype.ReadOnly`, `severitytype.Error`, `prioritytype.High`, `logleveltype.Debug`).
  - *"and why we have too many Enumer type defined, can you please fix and keep ones we need"*: Prune bloated and unused `Enumer` types in `04-code/golang/pkg/baseenumer/` and `04-code/golang/pkg/errtype/base_enumer.go` down to only the 4 essential types actually needed by the runtime (`BaseEnumer`, `ByteEnumer`, `NumberEnumer`, `StringEnumer`).
- **Current State:**
  - `pkg/enum/filepermtype/variant.go`, `fileoptype/variant.go`, `filewritemodetype/variant.go`, `severitytype/variant.go`, `prioritytype/variant.go`, and `pkg/errtype/logleveltype/variant.go` contain redundant internal const aliases.
  - `pkg/fileutil/file_perm_type.go`, `file_op_type.go`, `file_write_mode_type.go`, `consts.go`, and `pkg/appfault/severity_type.go`, `priority_type.go` mirror enum values as domain constants.
  - `pkg/baseenumer/` defines unnecessary variants: `Utf8Enumer`, `Utf16Enumer`, `Utf32Enumer`, `RuneEnumer`, `IntEnumer`.
- **Architectural Solution:**
  1. Prune `pkg/baseenumer/` and `pkg/errtype/base_enumer.go` to retain only 4 core interfaces: `BaseEnumer`, `ByteEnumer`, `NumberEnumer`, `StringEnumer`. Delete `utf16_enumer.go` and `utf32_enumer.go`.
  2. Eliminate all internal const aliases inside canonical enum packages.
  3. Remove const alias blocks in domain packages (`pkg/fileutil`, `pkg/appfault`).
  4. Update all call sites across `pkg/fileutil`, `pkg/appwriter`, `pkg/appfault`, `examples/`, and tests to use enum packages directly.
  5. Validate 100% test pass and all 31 CI/CD quality gates.

---

## 2. Task-Specific Rule Set
1. **Rule 1 (Direct Enum Usage):** Callers must use `filepermtype.Standard`, `fileoptype.ReadOnly`, `openfiletype.ReadOnly`, etc. Zero redundant const alias blocks.
2. **Rule 2 (Keep 4 Canonical Enumer Interfaces):** `BaseEnumer`, `ByteEnumer`, `NumberEnumer`, `StringEnumer` (with their `*Enum` aliases).
3. **Rule 3 (No Mixed Polarity / Boolean Purity):** Implicit boolean checks only (`if isReady`).
4. **Rule 4 (Canonical Function Size):** Functions must strictly respect `<= 15 lines`.
5. **Rule 5 (Strict Relative Git Paths):** All paths and links must be relative to repository root.
6. **Rule 6 (WOR Policy Active):** No version bumping, no git tags.

---

## 3. Subtask Breakdown
- [x] Subtask 1: `.lovable/plans/subtasks/26-prune-enum-const-aliases-and-enumer-types/01-task-prune-baseenumer-interfaces.md`
- [x] Subtask 2: `.lovable/plans/subtasks/26-prune-enum-const-aliases-and-enumer-types/02-task-remove-internal-enum-const-aliases.md`
- [x] Subtask 3: `.lovable/plans/subtasks/26-prune-enum-const-aliases-and-enumer-types/03-task-remove-domain-const-aliases-and-update-callers.md`
- [x] Subtask 4: `.lovable/plans/subtasks/26-prune-enum-const-aliases-and-enumer-types/04-task-verification-and-ci-gates.md`

---

## 4. Verification Plan
- Targeted package test: `go test -C 04-code/golang -v ./pkg/baseenumer ./pkg/errtype/... ./pkg/enum/... ./pkg/fileutil/... ./pkg/appfault/... ./pkg/appwriter/...`
- Full test pass: `go test -C 04-code/golang -count=1 ./...`
- Go formatter: `python 03-ai-scripts/26-go-code-formatter.py`
- Local CI runner: `python 03-ai-scripts/06-cicd-local-runner.py --all`
- Sync check: `npm run sync` and `node scripts/sync-check.mjs`
