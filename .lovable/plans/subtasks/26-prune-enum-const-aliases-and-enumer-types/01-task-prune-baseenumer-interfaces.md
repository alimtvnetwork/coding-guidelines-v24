# Subtask 26.1: Prune Over-Engineered Enumer Types

## Objective
Prune `04-code/golang/pkg/baseenumer/` and `04-code/golang/pkg/errtype/base_enumer.go` down to only the 4 essential interfaces:
1. `BaseEnumer` (alias `BaseEnum`)
2. `ByteEnumer` (alias `ByteEnum`)
3. `NumberEnumer` (alias `NumberEnum`)
4. `StringEnumer` (alias `StringEnum`)

## Changes
- Delete `04-code/golang/pkg/baseenumer/utf16_enumer.go`.
- Delete `04-code/golang/pkg/baseenumer/utf32_enumer.go`.
- Modify `04-code/golang/pkg/baseenumer/byte_enumer.go`: remove `Utf8Enumer`, `Utf8Enum`, `UTF8Enumer`, `UTF8Enum`.
- Modify `04-code/golang/pkg/baseenumer/number_enumer.go`: remove `IntEnumer`, `IntEnum`.
- Modify `04-code/golang/pkg/baseenumer/base_enumer_test.go`: remove mocks and tests for deleted interfaces; update `TestByteEnumer_Validity`.
- Modify `04-code/golang/pkg/baseenumer/readme.md`: update docs.
- Modify `04-code/golang/pkg/errtype/base_enumer.go`: remove aliases for deleted types.
- Modify `04-code/golang/pkg/errtype/methods.go`: remove `_ IntEnumer = Variation(0)`.
- Modify `04-code/golang/pkg/errtype/logleveltype/variant.go`: remove `_ baseenumer.IntEnumer = Variant(0)`.
- Modify `04-code/golang/pkg/errtype/logleveltype/readme.md`: remove `IntEnumer` references.

## Acceptance Criteria
- `go test -C 04-code/golang -v ./pkg/baseenumer ./pkg/errtype/...` passes with zero failures.
