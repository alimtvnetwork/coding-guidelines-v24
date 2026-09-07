# Master Plan 27: Direct Concrete Result Types Consolidation Across Enum Packages and FileUtil

## 1. Problem Statement & User Mandate
- **User Mandate:**
  - *"The constant folder or like types, types.go file, try to have all these types defined and make it like a direct type, uh, rather than the generic type that would be returned. Okay? Do you understand? I think we have discussed this several times. So fix it everywhere inside the file and find everything, everything similar and try to fix it in all these, uh, small packages. I think I'm repeating several times"*
  - Screenshot [10-direct-result-types.png](.lovable/assets/enum/10-direct-result-types.png) specifically targets `04-code/golang/pkg/fileutil/file_write_mode_type.go`:
    ```go
    func ParseFileWriteMode(s string) result.Wrap[FileWriteModeType] { ... }
    ```
    where `result.Wrap[FileWriteModeType]` is underlined in red.
- **Architectural Analysis:**
  1. In each small enum package (`pkg/enum/bytetype`, `fileoptype`, `filepermtype`, `filewritemodetype`, `logleveltype`, `openfiletype`, `processstatetype`), the parser functions currently return generic `result.Wrap[Variant]`. They should define a direct concrete type:
     ```go
     type Result = result.Wrap[Variant]
     ```
     and return `Result` directly: `func Parse(s string) Result`.
  2. In `pkg/fileutil/`:
     - `types.go` defines `FileResult`, `BytesResult`, `StringResult`, `LinesResult`, `BoolResult`, `FileInfoResult`, `Int64Result`.
     - Missing direct types: `FilePermResult`, `FileOpResult`, `FileWriteModeResult`, `FileOpenModeResult`, and `FileWriterResult`.
     - `file_perm_type.go`: `ParsePerm(octalStr string)` should return `FilePermResult`.
     - `file_op_type.go`: `ParseFileOp(s string)` should return `FileOpResult`.
     - `file_write_mode_type.go`: `ParseFileWriteMode(s string)` should return `FileWriteModeResult`.
     - `advanced.go` and `file_namespace.go`: `NewFileWriter(...)` and `StreamWriterAny/Append/Truncate(...)` should return `FileWriterResult`.
  3. In `pkg/payloadconv/`:
     - `converter.go`: `func ToBytes(payload any) result.Wrap[[]byte]` should return direct type `BytesResult = result.Wrap[[]byte]`.

---

## 2. Task-Specific Rules & Invariants
1. **Rule 1 (Zero-Bracket Concrete Return Types):** Exported and domain parse/creator functions must return direct concrete result types rather than verbose bracket generics (`result.Wrap[...]`).
2. **Rule 2 (Canonical Enum Result Standard):** All canonical enum packages returning wrapped variants must define `type Result = result.Wrap[Variant]`.
3. **Rule 3 (Fileutil Types Centralization):** All fileutil concrete results must be defined in `types.go` and aliased in domain files where relevant.
4. **Rule 4 (Strict Coding Guidelines):** Functions <= 15 lines, implicit booleans only (`if isReady`), no mixed polarity.
5. **Rule 5 (Strict Relative Git Paths):** All paths and references must be strictly relative to the repository root.
6. **Rule 6 (WOR Policy Active):** No version bumping, no git tags.

---

## 3. Subtask Breakdown
- [x] Subtask 1: `.lovable/plans/subtasks/27-direct-result-types-consolidation/01-task-enum-packages-direct-results.md`
- [x] Subtask 2: `.lovable/plans/subtasks/27-direct-result-types-consolidation/02-task-fileutil-direct-results.md`
- [x] Subtask 3: `.lovable/plans/subtasks/27-direct-result-types-consolidation/03-task-payloadconv-direct-result.md`
- [x] Subtask 4: `.lovable/plans/subtasks/27-direct-result-types-consolidation/04-task-verification-and-ci-gates.md`

---

## 4. Verification Plan
- Targeted package test: `go test -C 04-code/golang -v ./pkg/enum/... ./pkg/fileutil/... ./pkg/payloadconv/...`
- Full test pass: `go test -C 04-code/golang -count=1 ./...` across all 25 packages.
- Go formatter: `python 03-ai-scripts/26-go-code-formatter.py`
- CI local runner: `python 03-ai-scripts/06-cicd-local-runner.py --all`
- Sync checks: `npm run sync` and `node scripts/sync-check.mjs`
