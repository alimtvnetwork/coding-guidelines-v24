# `prioritytype` Package

`prioritytype` provides an integer-backed (byte) enumeration for error and fault priority levels (`Unknown`, `Low`, `Normal`, `High`, `Critical`) within the `coding-guidelines` Go runtime.

## Key Features
- **Zero Circular Dependencies:** Imports only `baseenumer` and standard libraries.
- **Interface Conformance:** Implements `baseenumer.BaseEnumer`, `baseenumer.ByteEnumer`, `baseenumer.NumberEnumer`, `json.Marshaler`, and `json.Unmarshaler`.
- **Parsing & JSON:** Case-insensitive parsing and standard JSON serialization.
