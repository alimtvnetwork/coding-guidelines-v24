# Subtask 39.2: Recursive Formatter with Sorted Map Keys

## Description
Create `04-code/golang/pkg/appfault/result_dynamic_formatter.go` implementing `FormatValue(val any) string` with recursive unwrapping of `ResultInspector` instances, deterministic map formatting with sorted keys (`map[k1:v1 k2:v2]`), slice/array formatting, error formatting `[Error: <message>]`, cycle protection (`maxDepth = 32`), and `UnwrapRecursive(val any) any` for deep structure simplification. Re-export in `04-code/golang/pkg/result/result.go`.

## Target Files
- `04-code/golang/pkg/appfault/result_dynamic_formatter.go` (new file)
- `04-code/golang/pkg/result/result.go`

## Acceptance Criteria
- [x] `FormatValue(val any) string` unwraps `ResultInspector` instances recursively.
- [x] Failed Result monads format cleanly as `[Error: <message>]` (or `[Error]` if empty).
- [x] Map entries are sorted alphabetically by string key representation (`map[k1:v1 k2:v2]`).
- [x] Slices and arrays format as `[item1 item2 ...]`.
- [x] `UnwrapRecursive(val any) any` simplifies arbitrary structures into standard Go primitives/maps/slices.
- [x] Recursion depth capped at 32 to prevent cyclical reference panics.
- [x] Every function is <= 15 lines.
- [x] All booleans use `is`/`has` prefixes and implicit evaluations.
