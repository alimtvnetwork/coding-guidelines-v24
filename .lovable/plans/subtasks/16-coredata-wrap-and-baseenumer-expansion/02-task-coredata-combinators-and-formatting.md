# Subtask 02: Coredata Combinators and Formatting

Target Directories: `04-code/golang/pkg/appfault/` and `04-code/golang/pkg/result/`
Affected Files:
- `04-code/golang/pkg/appfault/result_slice.go`
- `04-code/golang/pkg/appfault/result_map.go`
- `04-code/golang/pkg/appfault/result_methods.go`
- `04-code/golang/pkg/result/combinators.go`
- `04-code/golang/pkg/appfault/combinators_test.go`
- `04-code/golang/pkg/result/combinators_test.go`

## Instructions
1. In `04-code/golang/pkg/appfault/result_slice.go`:
   - Add `Filter(predicate func(item T) bool) ResultSlice[T]`.
   - Add `ForEach(fn func(index int, item T)) ResultSlice[T]`.
   - Add `ForEachBreak(fn func(index int, item T) bool) ResultSlice[T]`.
   - Add `FormatStruct() string`: formats slice items in aligned block or error banner.
2. In `04-code/golang/pkg/appfault/result_map.go`:
   - Add `Keys() []K`: returns sorted slice of keys (or slice if unsorted).
   - Add `Values() []V`: returns slice of values.
   - Add `Filter(predicate func(key K, val V) bool) ResultMap[K, V]`.
   - Add `ForEach(fn func(key K, val V)) ResultMap[K, V]`.
   - Add `FormatStruct() string`: formats map key-values in aligned block or error banner.
3. In `04-code/golang/pkg/appfault/result_methods.go`:
   - Add `FormatStruct() string`: formats struct or primitive payload or error banner.
   - Add `ToMap() map[string]any`: converts payload into map[string]any via JSON serialization.
4. In `04-code/golang/pkg/result/combinators.go`:
   - Add `MapSlice[T, U any](res Wrap[[]T], fn func(T) U) ResultSlice[U]`.
   - Add `FlatMapSlice[T, U any](res Wrap[[]T], fn func(T) []U) ResultSlice[U]`.
   - Add `MapMapValues[K comparable, V, U any](res ResultMap[K, V], fn func(V) U) ResultMap[K, U]`.
5. Tests:
   - In `04-code/golang/pkg/appfault/combinators_test.go`: test `Filter`, `ForEach`, `ForEachBreak`, `Keys`, `Values`, `FormatStruct`, `ToMap` on both success and failure cases.
   - In `04-code/golang/pkg/result/combinators_test.go`: test `MapSlice`, `FlatMapSlice`, `MapMapValues`.
   - All functions strictly <= 15 lines.
   - Blank lines after closing brace `}` if followed by code.
