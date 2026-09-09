# Plan 39: Recursive Map Sorting & Result Printing

## Goal Description
Enhance `pkg/appfault` and `pkg/result` printing, stringification, and serialization engines with:
1. **Recursive Value Formatting (`FormatValue`):** Recursively inspect every element in any data structure (maps, slices, arrays, structs, interfaces).
2. **Deterministic Map Key Sorting:** Sort map keys lexicographically before printing, eliminating Go's randomized map iteration order in string and JSON representations.
3. **Monadic Result Unwrapping (`ResultInspector`):** Non-generic runtime interface (`ValueAny() any`, `IsFailed() bool`, `AppError() *AppError`) implemented by `Result[T]`, `ResultSlice[T]`, and `ResultMap[K, V]`. When an item in a map/slice/result is a Result:
   - If successful: recursively unwraps and prints its inner value.
   - If failed: formats the error cleanly (e.g. `[Error: <message>]`).
4. **Deep Recursive Unwrapping (`UnwrapRecursive`):** Converts nested Result monads into plain Go data structures for clean, unwrapped JSON/YAML output.
5. **Output Parity:** Update `Result[T].String()`, `PrettyJson()`, `PrettyMap()`, and add uniform `String()` and `PrettyJson()` to `ResultSlice[T]` and `ResultMap[K, V]`.

## Task-Specific Rules
1. **Function Sizing:** Every function MUST be <= 15 lines.
2. **Boolean Conventions:** Implicit boolean evaluation only (`if isOk`), zero explicit `== true`. Positive prefixes `is`/`has` only.
3. **Strict Relative Git Paths:** Zero absolute filesystem paths or `file:///` URIs anywhere in plans, code, or documentation.
4. **Recursion Guard:** Recursion depth capped at 32 to prevent infinite loops or stack overflow on cyclic structures.
5. **Deterministic Ordering:** All map outputs MUST be sorted by string key representation across all nesting levels.

## Subtasks Decomposition

1. `01-task-result-unwrapping-interface-and-container-methods.md`: Non-generic `ResultInspector` interface, `ValueAny()` on Result containers, and encapsulation of `ResultSlice[T]` & `ResultMap[K, V]`.
2. `02-task-recursive-formatter-with-sorted-map-keys.md`: `result_dynamic_formatter.go` with `FormatValue`, `sortMapKeys`, `formatMapValue`, `UnwrapRecursive`, and depth protection.
3. `03-task-output-integration-and-unit-tests.md`: Integration into `String()`, `PrettyJson()`, `PrettyMap()`, `result_dynamic_formatter_test.go`, and CI verification.

## Verification Plan
- `cd 04-code/golang && go test ./pkg/appfault/... -v`
- `cd 04-code/golang && go test ./pkg/result/... -v`
- `cd 04-code/golang && go test ./... -v`
- `python 03-ai-scripts/26-go-code-formatter.py`
- `python 03-ai-scripts/06-cicd-local-runner.py --all`
