# Package `errtype/logleveltype`: Number-Backed Log Severity Enum

`coding-guidelines/common/pkg/errtype/logleveltype` provides a strongly-typed, `uint16`-backed log level enumeration conforming to `baseenumer.NumberEnumer`.

---

## 1. Overview & Architecture

- **Dedicated Package:** `logleveltype` located in `04-code/golang/pkg/errtype/logleveltype/`.
- **Underlying Type:** `type Variant uint16` (aliased to `LogLevelType`).
- **Interface Conformance:** Conforms to `baseenumer.BaseEnumer`, `baseenumer.NumberEnumer`, `json.Marshaler`, `json.Unmarshaler`.
- **Constants:** `Debug` (1), `Info` (2), `Warn` (3), `Error` (4), `Fatal` (5).

---

## 2. Usage Example

```go
import "coding-guidelines/common/pkg/errtype/logleveltype"

lvl := logleveltype.Warn
fmt.Println(lvl.Name())        // "Warn"
fmt.Println(lvl.Code())        // 3
fmt.Println(lvl.ValueString()) // "3"

parsed := logleveltype.Parse("debug")
if parsed.IsValid() {
    // parsed == logleveltype.Debug
}
```
