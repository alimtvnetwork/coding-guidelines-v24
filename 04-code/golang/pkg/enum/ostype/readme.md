# `ostype` Package

`ostype` provides a type-safe byte-backed enumeration for OSType variants (`AnyOs`, `Windows`, `Unix`, `Linux`, `MacOs`, `Ubuntu`, `Debian`, `ArchLinux`, `FreeBsd`, `Centos`, `RedHatEnterpriseLinux`, `Docker`, `Android`).

## Key Features

- **Zero Circular Dependencies:** Foundational leaf enum importing only `baseenumer` and standard libraries.
- **Zero-Allocation Parsing:** `Parse(s string) (Variant, bool)` directly delegates to `baseenumer.BasicInteger`.
- **Safe Fallback Parsers:** `ParseOrZero(s)`, `ParseOrInvalid(s)`, and `ParseOrUnknown(s)` helpers.
- **DRY JSON Marshaling:** Implements `json.Marshaler` and `json.Unmarshaler`.
- **Boundary Operations:** First-class `Min()`, `Max()`, `IsMin()`, `IsMax()`, and `IsInRange(min, max Variant) bool` conforming to `baseenumer.BoundedEnumer[Variant]`.
