# Package `errtype/processstatetype`: String-Backed Lifecycle State Enum

`coding-guidelines/common/pkg/errtype/processstatetype` provides a strongly-typed, `string`-backed process execution state enumeration conforming to `baseenumer.StringEnumer`.

---

## 1. Overview & Architecture

- **Dedicated Package:** `processstatetype` located in `04-code/golang/pkg/errtype/processstatetype/`.
- **Underlying Type:** `type Variant string` (aliased to `ProcessStateType`).
- **Interface Conformance:** Conforms to `baseenumer.BaseEnumer`, `baseenumer.StringEnumer`, `json.Marshaler`, `json.Unmarshaler`.
- **Constants:** `Pending`, `Running`, `Completed`, `Failed`, `Cancelled`, `Unknown`.

---

## 2. Usage Example

```go
import "coding-guidelines/common/pkg/errtype/processstatetype"

state := processstatetype.Running
fmt.Println(state.Name())        // "Running"
fmt.Println(state.Value())       // "Running"
fmt.Println(state.ValueString()) // "Running"

parsed := processstatetype.Parse("completed")
if parsed.IsValid() {
    // parsed == processstatetype.Completed
}
```
