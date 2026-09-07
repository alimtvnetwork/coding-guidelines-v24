# `severitytype` Package

`severitytype` provides an integer-backed (byte) enumeration for error and fault severity levels (`Unknown`, `Info`, `Warn`, `Error`, `Critical`, `Fatal`) within the `coding-guidelines` Go runtime.

## Key Features
- **Zero Circular Dependencies:** Imports only `baseenumer` and standard libraries.
- **Interface Conformance:** Implements `baseenumer.BaseEnumer`, `baseenumer.ByteEnumer`, `baseenumer.NumberEnumer`, `json.Marshaler`, and `json.Unmarshaler`.
- **Parsing & JSON:** Case-insensitive parsing and standard JSON serialization.
