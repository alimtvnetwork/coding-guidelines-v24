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
	_ baseenumer.NumberEnumer = mockNumberEnum(0)
	_ baseenumer.NumberEnum   = mockNumberEnum(0)
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
