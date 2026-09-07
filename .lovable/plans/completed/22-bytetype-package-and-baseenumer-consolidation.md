# Master Plan 22: ByteType Package Creation & BaseEnumer Helpers Consolidation

## 1. Problem Statement & User Mandate
- **User Mandate:**
  1. Port and create the canonical `bytetype` package inside `04-code/golang/pkg/enum/bytetype/` from `D:\work\03-aukgo\core\bytetype`.
  2. Consolidate repeated, common enum functions across all enum packages into `04-code/golang/pkg/baseenumer/` so enum packages don't have to duplicate lookup map compilation, label extraction, parsing, range validation, and formatting logic.
- **Identified Redundancies:**
  - `compileVariantMap`: identical quad-key registration and fallback map construction duplicated in `logleveltype`, `openfiletype`, `processstatetype`.
  - `All()` / `Values()`: identical index-1 slicing logic duplicated across all enum packages.
  - `Parse()`: whitespace trimming, empty validation error, lowercase map lookup, and error message formatting duplicated across all enum packages.
  - `FormatNameValue`: fallback string formatting `"Name(Value)"` duplicated across all enum packages.
  - Range validation: manual boundary comparisons instead of unified `IsBetween` / `IsNotBetween`.
- **Target Deliverables:**
  1. `04-code/golang/pkg/baseenumer/helpers.go` & `helpers_test.go`: Generic, zero-dependency helper functions (`CompileMap`, `SliceValues`, `SliceVariants`, `FormatNameValue`, `IsBetween`, `IsNotBetween`, `ParseLookup`, `FormatParseError`, `FormatEmptyParseError`, `FormatNumericRangeError`).
  2. Refactor `pkg/enum/logleveltype/`, `pkg/enum/openfiletype/`, `pkg/enum/processstatetype/` to use the shared base helpers.
  3. Create `04-code/golang/pkg/enum/bytetype/` with complete file structure:
     - `variant.go`: `Variant byte`, constants (`Zero`, `Min`, `One`, `Two`, `Three`, `Max`, `Invalid`, `Unknown`), predicates, comparisons, arithmetic (`Add`, `Subtract`), JSON handlers, and full conformance to `baseenumer.ByteEnumer` and `baseenumer.NumberEnumer`.
     - `vars.go`: label mappings, `All()`, `Values()`, `Parse()`.
     - `variant_test.go`: 100% test coverage.
     - `readme.md`: package architecture specification.

---

## 2. Task-Specific Rule Set
1. **Rule 1 (Zero-Dependency Base Layer):** `pkg/baseenumer` must remain a zero-dependency package importing only standard library packages (`cmp`, `fmt`, `strconv`, `strings`). It cannot import `errtype` or `result`.
2. **Rule 2 (Dual Interface Conformance):** `bytetype.Variant` must conform to both `baseenumer.ByteEnumer` and `baseenumer.NumberEnumer`, guarded by compile-time interface assertions.
3. **Rule 3 (Zero Breaking Changes):** All existing public function signatures and return types (`Parse() result.Wrap[Variant]`, `All() []Variant`, `Values() []string`) must remain 100% identical.
4. **Rule 4 (Canonical Function Size):** Functions must strictly respect `<= 15 lines`, blank line after closing brace if followed by code.
5. **Rule 5 (Strict Relative Git Paths):** All markdown links and plans must use strictly relative git paths.

---

## 3. Subtask Breakdown
- `.lovable/plans/subtasks/22-bytetype-package-and-baseenumer-consolidation/01-task-baseenumer-helpers-and-consolidation.md`
- `.lovable/plans/subtasks/22-bytetype-package-and-baseenumer-consolidation/02-task-bytetype-package-implementation.md`

---

## 4. Verification Plan
- Unit tests for all modified and created packages: `go test -C 04-code/golang -count=1 -v ./pkg/baseenumer ./pkg/enum/...`.
- Local CI runner: `python 03-ai-scripts/06-cicd-local-runner.py --all` (all 31 gates green).
- Metadata sync: `npm run sync` and `node scripts/sync-check.mjs`.
