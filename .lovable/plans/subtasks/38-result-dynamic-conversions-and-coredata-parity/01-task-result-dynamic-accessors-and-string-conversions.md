# Subtask 38.1: Result Dynamic Accessors & String Conversions

## Description
Encapsulate `Result[T]` struct fields as `value T` and `appError *AppError` to eliminate the Go compiler identifier collision between field and method `Value`. Implement accessors `Value() T`, `Data() T`, `Payload() T`, `Result() T`, `AppError() *AppError`, along with JSON and YAML marshaling/unmarshaling. Implement dynamic string line splitting and delimiters: `Lines()`, `LinesResult()`, `Split()`, `SplitResult()`, `SplitAt()`, and `SplitByRune()`.

## Target Files
- `04-code/golang/pkg/appfault/result.go`
- `04-code/golang/pkg/appfault/result_methods.go`
- `04-code/golang/pkg/appfault/result_constructors.go`
- `04-code/golang/pkg/appfault/result_dynamic_strings.go`

## Acceptance Criteria
1. `Result[T]` struct contains unexported `value T` and `appError *AppError`.
2. Accessors `Value() T`, `Data() T`, `Payload() T`, `Result() T`, and `AppError() *AppError` return the underlying fields.
3. `MarshalJSON()` and `UnmarshalJSON()` correctly serialize and deserialize `Result[T]` transparently.
4. `MarshalYAML()` and `UnmarshalYAML()` correctly serialize and deserialize `Result[T]`.
5. `Lines() []string` splits any string/bytes/fmt.Stringer/`[]string` payload into individual lines, respecting both `\r\n` and `\n`.
6. `LinesResult() ResultSlice[string]` returns lines wrapped in a monadic `ResultSlice[string]`.
7. `Split(sep string) []string` splits string payload by delimiter.
8. `SplitResult(sep string) ResultSlice[string]` returns split results in a monadic `ResultSlice[string]`.
9. `SplitAt(index int) (string, string)` safely splits a string at character position without out-of-bounds panics.
10. `SplitByRune(r rune) []string` splits string payload by single rune.
11. Every function is $\le 15$ lines.
12. All booleans use `is`/`has` prefixes and implicit evaluations.

## Status
- [x] Complete
