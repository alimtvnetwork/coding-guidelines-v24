# Spec: 16-coredata-wrap-and-baseenumer-expansion

## 1. Overview & Objectives
This specification coordinates two major architectural expansions inspired by `D:\work\03-aukgo\core\coredata`, `D:\work\03-aukgo\core\bytetype`, and `D:\work\03-aukgo\enum`:
1. **Modular Base Enum Family in `pkg/baseenumer`**:
   Expand `04-code/golang/pkg/baseenumer/` from basic numeric enums into a complete, modular, zero-dependency family of typed enum interfaces and contracts:
   - Byte & UTF-8 enums (`byte_enumer.go`): `ByteEnumer`, `ByteEnum`, `UTF8Enumer`, `UTF8Enum`.
   - UTF-16 enums (`utf16_enumer.go`): `UTF16Enumer`, `UTF16Enum`.
   - UTF-32 & Rune enums (`utf32_enumer.go`): `UTF32Enumer`, `UTF32Enum`, `RuneEnumer`, `RuneEnum`.
   - String enums (`string_enumer.go`): `StringEnumer`, `StringEnum`.
   - Number & Int enums (`number_enumer.go`): `NumberEnumer`, `NumberEnum`, `IntEnumer`, `IntEnum`.
   - Base enum foundation (`base_enumer.go`): `BaseEnumer`, `BaseEnum`, `ToEnum[T BaseEnumer]`.
   - Re-exports in `04-code/golang/pkg/errtype/base_enum.go` for seamless backward compatibility.

2. **Coredata Collection Combinators & Dynamic Struct Formatting in `pkg/appfault` & `pkg/result`**:
   Incorporate collection combinators (`Filter`, `ForEach`, `ForEachBreak`, `Keys`, `Values`) and dynamic struct inspection/formatting methods (`FormatStruct()`, `ToMap()`) on monadic containers `Result[T]`, `ResultSlice[T]`, and `ResultMap[K, V]`.
   Provide package-level cross-type mappers in `pkg/result/combinators.go` (`MapSlice`, `FlatMapSlice`, `MapMapValues`).

3. **Package Release Architecture & Strategy**:
   Formulate a documented, reproducible release blueprint detailing package versioning, dependency hierarchy, automated gates, and zero-breaking-change migration.

---

## 2. Task-Specific Rules & Constraints
1. **Strict Interface Naming (`er` Suffix):** Every named Go interface MUST end with `er` (e.g. `ByteEnumer`, `UTF8Enumer`, `UTF16Enumer`, `UTF32Enumer`, `RuneEnumer`, `StringEnumer`, `NumberEnumer`, `IntEnumer`). Type aliases use the concise `*Enum = *Enumer` form.
2. **Zero-Dependency Guarantee:** `04-code/golang/pkg/baseenumer/` MUST have zero internal repository dependencies. It only imports stdlib (`strings`, `fmt`, `math`).
3. **Strict Function Size Limit:** No function or test may exceed 15 lines (strictly <= 15 lines, target <= 8 lines).
4. **Vertical Whitespace:** Blank line required after closing brace `}` if followed by more code. No double empty lines.
5. **Strict Relative Git Paths:** All file paths, markdown links, and citations MUST be strictly relative to the repository root.

---

## 3. Subtask Decomposition

### Subtask 01: Modular Base Enum Family (`01-task-baseenumer-modular-family.md`)
- Files:
  - `04-code/golang/pkg/baseenumer/byte_enumer.go`
  - `04-code/golang/pkg/baseenumer/utf16_enumer.go`
  - `04-code/golang/pkg/baseenumer/utf32_enumer.go`
  - `04-code/golang/pkg/baseenumer/string_enumer.go`
  - `04-code/golang/pkg/baseenumer/number_enumer.go`
  - `04-code/golang/pkg/baseenumer/base_enumer.go`
  - `04-code/golang/pkg/baseenumer/base_enumer_test.go`
  - `04-code/golang/pkg/errtype/base_enum.go`
- Description: Create modular enum interface contracts and aliases for byte/utf8, utf16, utf32/rune, string, and number/int; add unit tests and verify backwards compatibility.

### Subtask 02: Coredata Collection Combinators & Struct Formatting (`02-task-coredata-combinators-and-formatting.md`)
- Files:
  - `04-code/golang/pkg/appfault/result_slice.go`
  - `04-code/golang/pkg/appfault/result_map.go`
  - `04-code/golang/pkg/appfault/result_methods.go`
  - `04-code/golang/pkg/result/combinators.go`
  - `04-code/golang/pkg/appfault/combinators_test.go`
  - `04-code/golang/pkg/result/combinators_test.go`
- Description: Implement `Filter`, `ForEach`, `ForEachBreak`, `Keys`, `Values`, `FormatStruct`, and `ToMap` on `ResultSlice`, `ResultMap`, and `Result`; provide package-level functional combinators; add tests.

---

## 4. Package Release Architecture & Strategy
To safely and reliably release `baseenumer`, `appfault`, and `result`:
1. **Dependency Inversion & Zero-Cycle Layering:**
   - Layer 0 (Leaf): `pkg/baseenumer` (no dependencies).
   - Layer 1: `pkg/appfault` (depends only on `pkg/errtype` and stdlib).
   - Layer 2: `pkg/result` and `pkg/streamwriter` (depend on `pkg/appfault` and `pkg/typecast`).
2. **Release Lifecycle Steps:**
   - **Quality Gate Verification:** Run `python 03-ai-scripts/06-cicd-local-runner.py` (all 22 gates must exit 0).
   - **SemVer Tagging:** Increment patch/minor version in `version.json` via release runner scripts.
   - **Package Sync:** Synchronize package version across `package.json`, documentation, and Go modules.
   - **Changelog Generation:** Generate automated changelog detailing additions, backward-compatibility forwarders, and interface contracts.
