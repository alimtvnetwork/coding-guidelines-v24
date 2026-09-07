# Subtask 01: Baseenumer Modular Family

Target Directory: `04-code/golang/pkg/baseenumer/`
Affected Files:
- `04-code/golang/pkg/baseenumer/byte_enumer.go`
- `04-code/golang/pkg/baseenumer/utf16_enumer.go`
- `04-code/golang/pkg/baseenumer/utf32_enumer.go`
- `04-code/golang/pkg/baseenumer/string_enumer.go`
- `04-code/golang/pkg/baseenumer/number_enumer.go`
- `04-code/golang/pkg/baseenumer/base_enumer.go`
- `04-code/golang/pkg/baseenumer/base_enumer_test.go`
- `04-code/golang/pkg/errtype/base_enum.go`

## Instructions
1. In `04-code/golang/pkg/baseenumer/base_enumer.go`:
   - Keep `BaseEnumer` (`Name() string`, `String() string`, `ValueString() string`, `IsValid() bool`, `IsEnum() bool`).
   - Keep `BaseEnum = BaseEnumer`.
   - Keep `ToEnum[T BaseEnumer](val string, all []T) (T, bool)`.
2. Create `04-code/golang/pkg/baseenumer/byte_enumer.go`:
   - Define `ByteEnumer` (`BaseEnumer`, `Byte() byte`, `ValueByte() byte`, `Bytes() []byte`).
   - Define `ByteEnum = ByteEnumer`.
   - Define `UTF8Enumer` embedding `ByteEnumer`.
   - Define `UTF8Enum = UTF8Enumer`.
3. Create `04-code/golang/pkg/baseenumer/utf16_enumer.go`:
   - Define `UTF16Enumer` (`BaseEnumer`, `UTF16() uint16`, `ValueUTF16() uint16`, `Code() uint16`).
   - Define `UTF16Enum = UTF16Enumer`.
4. Create `04-code/golang/pkg/baseenumer/utf32_enumer.go`:
   - Define `UTF32Enumer` (`BaseEnumer`, `Rune() rune`, `ValueRune() rune`, `Int32() int32`).
   - Define `UTF32Enum = UTF32Enumer`.
   - Define `RuneEnumer` embedding `UTF32Enumer`.
   - Define `RuneEnum = RuneEnumer`.
5. Create `04-code/golang/pkg/baseenumer/string_enumer.go`:
   - Define `StringEnumer` (`BaseEnumer`, `ValueString() string`, `String() string`, `Name() string`).
   - Define `StringEnum = StringEnumer`.
6. Create `04-code/golang/pkg/baseenumer/number_enumer.go`:
   - Define `NumberEnumer` (`BaseEnumer`, `Int() int`, `Code() uint16`).
   - Define `NumberEnum = NumberEnumer`.
   - Define `IntEnumer` embedding `NumberEnumer`.
   - Define `IntEnum = IntEnumer`.
7. In `04-code/golang/pkg/errtype/base_enum.go`:
   - Forward all new interfaces and aliases from `baseenumer`.
8. In `04-code/golang/pkg/baseenumer/base_enumer_test.go`:
   - Add mock implementations and tests for each enum type (`mockByteEnum`, `mockUTF16Enum`, `mockUTF32Enum`, `mockStringEnum`, `mockNumberEnum`).
   - Ensure all functions <= 15 lines.
   - Blank line after each `}` if followed by code.
