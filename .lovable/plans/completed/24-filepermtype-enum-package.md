# Master Plan 24: FilePermType Enum Package Isolation (`pkg/enum/filepermtype`)

## 1. Problem Statement & User Mandate
- **User Mandate:** "Should be placed in its own enum package" with screenshot `.lovable/assets/enum/08-file-perm-type.png` showing `FilePermType` in `04-code/golang/pkg/fileutil/file_perm_type.go`.
- **Current State:**
  - `FilePermType` is defined as a loose enum directly inside package `fileutil` (`04-code/golang/pkg/fileutil/file_perm_type.go`).
  - All other enums in the repository (`bytetype`, `logleveltype`, `openfiletype`, `processstatetype`) have their own dedicated packages under `04-code/golang/pkg/enum/`.
- **Architectural Solution:**
  1. Create canonical enum package `04-code/golang/pkg/enum/filepermtype/`:
     - `variant.go`: `type Variant uint32`, `type FilePermType = Variant`, constants (`None`, `Standard`, `Private`, `Executable`, etc.), methods (`Mode`, `Uint32`, `OctalString`, `PosixString`, predicates, mutators, JSON marshaling).
     - `vars.go`: `All() []Variant`, `Values() []string`, `Parse() result.Wrap[Variant]`, `FromFileMode() Variant`.
     - `variant_test.go`: 100% unit test coverage.
     - `readme.md`: package documentation.
  2. Update `04-code/golang/pkg/fileutil/file_perm_type.go` to be a clean forwarder re-exporting `filepermtype.Variant`, aliases, and constants so that all existing callers throughout `fileutil` and `appwriter` continue working seamlessly without breaking changes.
  3. Ensure all quality gates pass (formatting, unit tests, 31 local CI gates).

---

## 2. Task-Specific Rule Set
1. **Rule 1 (Package Name Suffix):** Package MUST be named `filepermtype` in directory `04-code/golang/pkg/enum/filepermtype/`.
2. **Rule 2 (Zero Breaking Changes):** `pkg/fileutil` must re-export `FilePermType = filepermtype.Variant` and all `FilePerm*` constants so callers remain 100% functional.
3. **Rule 3 (No Mixed Polarity / Boolean Purity):** All predicates must use positive implicit conditions (`if hasRead { ... }`).
4. **Rule 4 (Canonical Function Size):** Functions must strictly respect `<= 15 lines`.
5. **Rule 5 (Strict Relative Git Paths):** All markdown links and references must use strictly relative git paths.

---

## 3. Subtask Breakdown
- Subtask 01: `.lovable/plans/subtasks/24-filepermtype-enum-package/01-task-create-filepermtype-package.md`
- Subtask 02: `.lovable/plans/subtasks/24-filepermtype-enum-package/02-task-fileutil-forwarders-and-verification.md`

---

## 4. Verification Plan
- Unit test verification: `go test -C 04-code/golang -v ./pkg/enum/filepermtype ./pkg/fileutil`.
- All Go packages test verification: `go test -C 04-code/golang -count=1 ./...`.
- Local CI runner: `python 03-ai-scripts/06-cicd-local-runner.py --all` (all 31 gates green).
- Sync check: `npm run sync` and `node scripts/sync-check.mjs`.
