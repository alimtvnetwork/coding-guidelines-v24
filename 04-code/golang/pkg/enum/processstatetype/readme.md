# Package `processstatetype`: Standard Process State Enum

`coding-guidelines/common/pkg/enum/processstatetype` provides the canonical byte-backed enum for lifecycle states of processes, tasks, and jobs across all Go applications, conforming to Aukgo (`03-aukgo`) and global enum guidelines.

---

## 1. Overview & Architecture

Following the repository's modular enum design, `processstatetype` provides:
- **Dedicated Package:** `processstatetype` (ends with `type` suffix).
- **Core Type:** `type Variant byte` in `variant.go`.
- **Zero-Value:** `Invalid Variant = iota` (aliased to `Unknown`).
- **Zero Stutter:** Callers write `processstatetype.Variant` and `processstatetype.Pending`, `processstatetype.Running`, etc.

---

## 2. Variant Catalog

| Variant | Value (`byte`) | Name / String | Description |
| :--- | :--- | :--- | :--- |
| `Invalid` / `Unknown` | `0` | `"Unknown"` | Uninitialized or invalid process state (zero value) |
| `Pending` | `1` | `"Pending"` | Queued or scheduled, waiting to run |
| `Running` | `2` | `"Running"` | Actively executing |
| `Completed` | `3` | `"Completed"` | Successfully completed execution |
| `Failed` | `4` | `"Failed"` | Execution terminated with errors |
| `Cancelled` | `5` | `"Cancelled"` | Execution aborted or cancelled |

---

## 3. Usage Example

```go
import "coding-guidelines/common/pkg/enum/processstatetype"

func CanExecute(state processstatetype.Variant) bool {
    return state.IsPending()
}
```

---

## 4. Helper Methods

- `.Name() string` / `.Label() string` / `.String() string` - PascalCase name.
- `.IsValid() bool` - Checks if variant is within valid defined range (`> Invalid`).
- `.IsInvalid() bool` - Checks if variant is zero-value or undefined.
- `.IsPending() bool`, `.IsRunning() bool`, etc. - Positive boolean predicates.
- `All() []Variant` - Returns slice of all valid variants (`Pending` through `Cancelled`).
- `Values() []string` - Returns slice of valid variant names.
- `Parse(s string) (Variant, bool)` - Case-insensitive string parser returning the typed variant and success boolean.
- `ParseOrInvalid(s string) Variant` / `ParseOrUnknown(s string) Variant` - String parser returning `Invalid` (0) on failure.
- `MarshalJSON()` / `UnmarshalJSON()` - PascalCase JSON serialization with string and byte fallback.

