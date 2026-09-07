# Package `bytetype`: Canonical Byte Enumeration & Numeric Wrapper

`coding-guidelines/common/pkg/enum/bytetype` provides the canonical byte-backed enum and value wrapper conforming to Aukgo (`03-aukgo/core/bytetype`) and repository enum standards (`baseenumer.ByteEnumer`, `baseenumer.NumberEnumer`).

---

## 1. Overview & Architecture

Following the repository's modular enum design, `bytetype` provides:
- **Dedicated Package:** `bytetype` (ends with `type` suffix).
- **Core Type:** `type Variant byte` in `variant.go`.
- **Dual Interface Compliance:** Implements both `baseenumer.ByteEnumer` (`Byte()`, `ValueByte()`, `Bytes()`) and `baseenumer.NumberEnumer` (`Int()`, `Code()`).
- **Zero Stutter:** Callers write `bytetype.Variant`, `bytetype.Zero`, `bytetype.One`, `bytetype.Max`, etc.

---

## 2. Variant Catalog

| Variant | Value (`byte`) | Name / String | Description |
| :--- | :--- | :--- | :--- |
| `Zero` / `Min` | `0` | `"Zero"` | Minimum byte value, aliased to `Invalid` and `Unknown` |
| `One` | `1` | `"One"` | Byte value 1 |
| `Two` | `2` | `"Two"` | Byte value 2 |
| `Three` | `3` | `"Three"` | Byte value 3 |
| `Max` | `255` | `"Max"` | Maximum byte value (`math.MaxUint8`) |

---

## 3. Usage Example

```go
import "coding-guidelines/common/pkg/enum/bytetype"

// Value extraction
b := bytetype.Two.Byte()       // byte(2)
num := bytetype.Two.Int()      // int(2)
str := bytetype.Two.Name()     // "Two"

// Arithmetic & comparisons
three := bytetype.One.Add(2)   // bytetype.Three
isGt := three.IsGreater(1)     // true
isBt := three.IsBetween(1, 5)  // true

// Parsing
res := bytetype.Parse("two")
if res.IsSuccess() {
    v := res.Data()            // bytetype.Two
}
```

---

## 4. Helper Methods & Functions

- **Byte Extraction:** `Byte()`, `ValueByte()`, `Value()`, `Bytes() []byte`
- **Integer Extraction:** `Int()`, `ValueInt()`, `ValueInt8()`, `ValueInt16()`, `ValueInt32()`, `ValueUInt16()`, `Code()`
- **Identity & Formatting:** `Name()`, `Label()`, `String()`, `ValueString()`, `StringValue()`, `ToNumberString()`, `NameValue()`, `JsonString()`
- **Predicates:** `IsValid()`, `IsInvalid()`, `IsEnum()`, `IsZero()`, `IsMin()`, `IsOne()`, `IsTwo()`, `IsThree()`, `IsMax()`, `Is(Variant)`
- **Comparisons:** `IsEqual`, `IsEqualInt`, `IsGreater`, `IsGreaterInt`, `IsGreaterEqual`, `IsGreaterEqualInt`, `IsLess`, `IsLessInt`, `IsLessEqual`, `IsLessEqualInt`, `IsBetween`, `IsBetweenInt`, `IsValueEqual`, `IsNameEqual`, `IsAnyNamesOf`
- **Arithmetic:** `Add(byte) Variant`, `Subtract(byte) Variant`
- **Package Functions:** `New(byte) Variant`, `GetSet(bool, Variant, Variant) Variant`, `GetSetVariant(bool, byte, byte) Variant`, `String([]byte) string`
- **Catalog & Parse:** `All() []Variant`, `Values() []string`, `Parse(string) Result`
- **JSON Serialization:** `MarshalJSON() ([]byte, error)`, `UnmarshalJSON([]byte) error`
