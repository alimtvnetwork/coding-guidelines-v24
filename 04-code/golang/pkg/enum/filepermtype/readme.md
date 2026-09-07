# Package `filepermtype`: POSIX File Permission Bitmask Enum

`coding-guidelines/common/pkg/enum/filepermtype` provides a strongly-typed `uint32` POSIX file permission bitmask enumeration conforming to `baseenumer.NumberEnumer` and global repository enum standards.

---

## 1. Overview & Architecture

- **Dedicated Package:** `filepermtype` located in `04-code/golang/pkg/enum/filepermtype/`.
- **Underlying Type:** `type Variant uint32` (aliased to `FilePermType = Variant`).
- **Interface Conformance:** Conforms to `baseenumer.BaseEnumer`, `baseenumer.NumberEnumer`, `json.Marshaler`, `json.Unmarshaler`.
- **Constants:** `Standard` (0644), `Private` (0600), `Executable` (0755), `ReadOnly` (0444), `PublicAll` (0777), `StickyDir` (01777), etc.

---

## 2. Usage Example

```go
import "coding-guidelines/common/pkg/enum/filepermtype"

// Inspection
perm := filepermtype.Standard
fmt.Println(perm.OctalString()) // "0644"
fmt.Println(perm.PosixString()) // "rw-r--r--"
osMode := perm.Mode()           // os.FileMode(0644)

// Predicates
if perm.IsPublic() { ... }
if perm.IsExecutable() { ... }

// Mutators
priv := perm.WithPrivate()      // 0600
exec := perm.WithExecutable()   // 0755

// Parsing
res := filepermtype.Parse("0644")
if res.IsSuccess() {
    v := res.Data()
}
```
