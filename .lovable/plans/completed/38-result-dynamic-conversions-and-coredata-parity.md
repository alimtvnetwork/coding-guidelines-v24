# Plan 38: Result[T] Rich Conversions & CoreDynamic Parity

## Goal
Enrich `Result[T]` (and its alias `result.Wrap[T]`) with comprehensive dynamic accessors, string line splitting, number conversions, dynamic reflection casting (`ReflectTo`), type inspection, pretty JSON/YAML serialization, and full parity with `coredynamic` from `03-aukgo/core/coredata/coredynamic`.

## Task-Specific Rules
1. **Function Size:** Every function and method MUST be $\le 15$ lines.
2. **Boolean Principles:** Implicit boolean evaluations only (`if isOk`), zero explicit `== true`. All booleans prefixed with `is` or `has`.
3. **Strict Relative Git Paths:** Zero absolute paths or `file:///` URIs in any repo files, plans, or docs.
4. **Zero Circular Imports:** Maintain clean unidirectional layering (`pkg/typecast` -> `pkg/appfault` -> `pkg/result`).
5. **CoreDynamic Parity:** Align method naming and behaviors with `03-aukgo/core/coredata/coredynamic` (`Data`, `Value`, `Lines`, `Split`, `SplitAt`, `Int`, `Int64`, `Float64`, `ReflectTo`, `Type`, `TypeName`, `Kind`, `IsNumber`, `PrettyJson`, `ToYaml`).

## Subtasks Breakdown
- [x] `01-task-result-dynamic-accessors-and-string-conversions.md`: Encapsulate fields, add `Value()`, `Data()`, `Lines()`, `Split()`, `SplitAt()`, and string converters.
- [x] `02-task-result-numbers-and-reflection-conversions.md`: Add number conversions (`Int`, `Int64`, `Float64`), `ReflectTo()`, type inspection, and length getters.
- [x] `03-task-result-pretty-yaml-and-test-verification.md`: Add `PrettyJson()`, `ToYaml()`, caller updates, comprehensive unit tests, and full CI/CD validation.
