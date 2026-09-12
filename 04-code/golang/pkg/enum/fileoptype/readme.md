# `fileoptype` Package

`fileoptype` provides a type-safe byte-backed enumeration for high-level file system operations (`ReadOnly`, `WriteOnly`, `ReadWrite`, `Append`, `Create`, `CreateAppend`, `CreateTruncate`, `Delete`) within the `coding-guidelines` Go runtime.

## Key Features

- **Interface Conformance:** Implements `baseenumer.BaseEnumer`, `baseenumer.ByteEnumer`, `baseenumer.NumberEnumer`, `json.Marshaler`, and `json.Unmarshaler`.
- **OpenMode Mapping:** Direct zero-allocation mapping to `openfiletype.Variant` via `.OpenMode()`.
- **Parsing & JSON:** High-performance case-insensitive parsing and standard JSON serialization.
