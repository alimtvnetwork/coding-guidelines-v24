package baseenumer_test

import (
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
)

type mockStringEnum string

func (m mockStringEnum) Name() string {
	return string(m)
}

func (m mockStringEnum) String() string {
	return string(m)
}

func (m mockStringEnum) ValueString() string {
	return string(m)
}

func (m mockStringEnum) IsValid() bool {
	return len(m) > 0
}

func (m mockStringEnum) IsEnum() bool {
	return len(m) > 0
}

type mockByteEnum byte

func (m mockByteEnum) Name() string {
	return "MockByte"
}

func (m mockByteEnum) String() string {
	return "MockByte"
}

func (m mockByteEnum) ValueString() string {
	return "MockByte"
}

func (m mockByteEnum) IsValid() bool {
	return m > 0
}

func (m mockByteEnum) IsEnum() bool {
	return m > 0
}

func (m mockByteEnum) Byte() byte {
	return byte(m)
}

func (m mockByteEnum) ValueByte() byte {
	return byte(m)
}

func (m mockByteEnum) Bytes() []byte {
	return []byte{byte(m)}
}

type mockUtf16Enum uint16

type mockUTF16Enum = mockUtf16Enum

func (m mockUtf16Enum) Name() string {
	return "MockUtf16"
}

func (m mockUtf16Enum) String() string {
	return "MockUtf16"
}

func (m mockUtf16Enum) ValueString() string {
	return "65"
}

func (m mockUtf16Enum) IsValid() bool {
	return m > 0
}

func (m mockUtf16Enum) IsEnum() bool {
	return m > 0
}

func (m mockUtf16Enum) Utf16() uint16 {
	return uint16(m)
}

func (m mockUtf16Enum) ValueUtf16() uint16 {
	return uint16(m)
}

func (m mockUtf16Enum) UTF16() uint16 {
	return uint16(m)
}

func (m mockUtf16Enum) ValueUTF16() uint16 {
	return uint16(m)
}

func (m mockUtf16Enum) Code() uint16 {
	return uint16(m)
}

type mockUtf32Enum rune

type mockUTF32Enum = mockUtf32Enum

func (m mockUtf32Enum) Name() string {
	return "MockUtf32"
}

func (m mockUtf32Enum) String() string {
	return "MockUtf32"
}

func (m mockUtf32Enum) ValueString() string {
	return string(m)
}

func (m mockUtf32Enum) IsValid() bool {
	return m > 0
}

func (m mockUtf32Enum) IsEnum() bool {
	return m > 0
}

func (m mockUtf32Enum) Rune() rune {
	return rune(m)
}

func (m mockUtf32Enum) ValueRune() rune {
	return rune(m)
}

func (m mockUtf32Enum) Int32() int32 {
	return int32(m)
}

type mockNumberEnum uint16

func (m mockNumberEnum) Name() string {
	return "MockNumber"
}

func (m mockNumberEnum) String() string {
	return "MockNumber"
}

func (m mockNumberEnum) ValueString() string {
	return "10"
}

func (m mockNumberEnum) IsValid() bool {
	return m > 0
}

func (m mockNumberEnum) IsEnum() bool {
	return m > 0
}

func (m mockNumberEnum) Int() int {
	return int(m)
}

func (m mockNumberEnum) Code() uint16 {
	return uint16(m)
}

var (
	_ baseenumer.BaseEnumer   = mockStringEnum("")
	_ baseenumer.BaseEnum     = mockStringEnum("")
	_ baseenumer.StringEnumer = mockStringEnum("")
	_ baseenumer.StringEnum   = mockStringEnum("")

	_ baseenumer.BaseEnumer = mockByteEnum(0)
	_ baseenumer.ByteEnumer = mockByteEnum(0)
	_ baseenumer.ByteEnum   = mockByteEnum(0)
	_ baseenumer.Utf8Enumer = mockByteEnum(0)
	_ baseenumer.Utf8Enum   = mockByteEnum(0)
	_ baseenumer.UTF8Enumer = mockByteEnum(0)
	_ baseenumer.UTF8Enum   = mockByteEnum(0)

	_ baseenumer.BaseEnumer  = mockUtf16Enum(0)
	_ baseenumer.Utf16Enumer = mockUtf16Enum(0)
	_ baseenumer.Utf16Enum   = mockUtf16Enum(0)
	_ baseenumer.UTF16Enumer = mockUtf16Enum(0)
	_ baseenumer.UTF16Enum   = mockUtf16Enum(0)

	_ baseenumer.BaseEnumer  = mockUtf32Enum(0)
	_ baseenumer.Utf32Enumer = mockUtf32Enum(0)
	_ baseenumer.Utf32Enum   = mockUtf32Enum(0)
	_ baseenumer.UTF32Enumer = mockUtf32Enum(0)
	_ baseenumer.UTF32Enum   = mockUtf32Enum(0)
	_ baseenumer.RuneEnumer  = mockUtf32Enum(0)
	_ baseenumer.RuneEnum    = mockUtf32Enum(0)

	_ baseenumer.BaseEnumer   = mockNumberEnum(0)
	_ baseenumer.NumberEnumer = mockNumberEnum(0)
	_ baseenumer.NumberEnum   = mockNumberEnum(0)
	_ baseenumer.IntEnumer    = mockNumberEnum(0)
	_ baseenumer.IntEnum      = mockNumberEnum(0)
)

func TestBaseEnumer_StringMethods(t *testing.T) {
	var e baseenumer.BaseEnumer = mockStringEnum("Active")
	if e.Name() != "Active" {
		t.Fatalf("expected Name() == 'Active', got %s", e.Name())
	}

	if e.String() != "Active" {
		t.Fatalf("expected String() == 'Active', got %s", e.String())
	}

	if e.ValueString() != "Active" {
		t.Fatalf("expected ValueString() == 'Active', got %s", e.ValueString())
	}
}

func TestBaseEnumer_Validity(t *testing.T) {
	var e baseenumer.BaseEnumer = mockStringEnum("Active")
	if !e.IsValid() {
		t.Fatal("expected IsValid() to be true")
	}

	if !e.IsEnum() {
		t.Fatal("expected IsEnum() to be true")
	}
}

func TestStringEnumer_Methods(t *testing.T) {
	var se baseenumer.StringEnumer = mockStringEnum("Hello")
	if se.Name() != "Hello" {
		t.Fatalf("expected Name() == 'Hello', got %s", se.Name())
	}

	if se.String() != "Hello" {
		t.Fatalf("expected String() == 'Hello', got %s", se.String())
	}

	if se.ValueString() != "Hello" {
		t.Fatalf("expected ValueString() == 'Hello', got %s", se.ValueString())
	}
}

func TestByteEnumer_Methods(t *testing.T) {
	var be baseenumer.ByteEnumer = mockByteEnum(65)
	if be.Byte() != 65 {
		t.Fatalf("expected byte 65, got %d", be.Byte())
	}

	if be.ValueByte() != 65 {
		t.Fatalf("expected ValueByte 65, got %d", be.ValueByte())
	}

	if len(be.Bytes()) != 1 {
		t.Fatalf("expected 1 byte, got %d", len(be.Bytes()))
	}
}

func TestByteEnumer_Validity(t *testing.T) {
	var be baseenumer.Utf8Enumer = mockByteEnum(65)
	if !be.IsValid() {
		t.Fatal("expected IsValid() to be true")
	}

	if !be.IsEnum() {
		t.Fatal("expected IsEnum() to be true")
	}
}

func TestUtf16Enumer_Methods(t *testing.T) {
	var ue baseenumer.Utf16Enumer = mockUtf16Enum(123)
	if ue.Utf16() != 123 {
		t.Fatalf("expected Utf16() == 123, got %d", ue.Utf16())
	}

	if ue.ValueUtf16() != 123 {
		t.Fatalf("expected ValueUtf16() == 123, got %d", ue.ValueUtf16())
	}

	if ue.Code() != 123 {
		t.Fatalf("expected Code() == 123, got %d", ue.Code())
	}
}

func TestUtf16Enumer_Validity(t *testing.T) {
	var ue baseenumer.Utf16Enum = mockUtf16Enum(123)
	if !ue.IsValid() {
		t.Fatal("expected IsValid() to be true")
	}

	if !ue.IsEnum() {
		t.Fatal("expected IsEnum() to be true")
	}
}

func TestUtf32Enumer_Methods(t *testing.T) {
	var ue baseenumer.Utf32Enumer = mockUtf32Enum('A')
	if ue.Rune() != 'A' {
		t.Fatalf("expected Rune() == 'A', got %c", ue.Rune())
	}

	if ue.ValueRune() != 'A' {
		t.Fatalf("expected ValueRune() == 'A', got %c", ue.ValueRune())
	}

	if ue.Int32() != 65 {
		t.Fatalf("expected Int32() == 65, got %d", ue.Int32())
	}
}

func TestRuneEnumer_Validity(t *testing.T) {
	var re baseenumer.RuneEnumer = mockUtf32Enum('A')
	if !re.IsValid() {
		t.Fatal("expected IsValid() to be true")
	}

	if !re.IsEnum() {
		t.Fatal("expected IsEnum() to be true")
	}
}

func TestNumberEnumer_Methods(t *testing.T) {
	var ne baseenumer.NumberEnumer = mockNumberEnum(10)
	if ne.Int() != 10 {
		t.Fatalf("expected Int() == 10, got %d", ne.Int())
	}

	if ne.Code() != 10 {
		t.Fatalf("expected Code() == 10, got %d", ne.Code())
	}

	if !ne.IsValid() {
		t.Fatal("expected IsValid() to be true")
	}
}

func TestIntEnumer_Methods(t *testing.T) {
	var ie baseenumer.IntEnumer = mockNumberEnum(20)
	if ie.Int() != 20 {
		t.Fatalf("expected Int() == 20, got %d", ie.Int())
	}

	if !ie.IsEnum() {
		t.Fatal("expected IsEnum() to be true")
	}
}

func TestToEnum_FoundByName(t *testing.T) {
	all := []mockStringEnum{mockStringEnum("First"), mockStringEnum("Second")}
	found, ok := baseenumer.ToEnum("first", all)
	if !ok {
		t.Fatal("expected ToEnum to find item")
	}

	if found != mockStringEnum("First") {
		t.Fatalf("expected First, got %s", found)
	}
}

func TestToEnum_NotFound(t *testing.T) {
	all := []mockStringEnum{mockStringEnum("First")}
	_, ok := baseenumer.ToEnum("missing", all)
	if ok {
		t.Fatal("expected ToEnum to return false for missing")
	}
}
