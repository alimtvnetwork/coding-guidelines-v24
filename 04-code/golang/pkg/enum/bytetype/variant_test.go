package bytetype_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/enum/bytetype"
)

func TestConstantsAndAliases(t *testing.T) {
	if bytetype.Zero != 0 || bytetype.Min != 0 {
		t.Fatalf("expected Zero and Min to be 0")
	}

	if bytetype.One != 1 || bytetype.Two != 2 || bytetype.Three != 3 {
		t.Fatalf("expected One=1, Two=2, Three=3")
	}

	if bytetype.Max != 255 {
		t.Fatalf("expected Max to be 255")
	}

	if bytetype.Invalid != bytetype.Zero || bytetype.Unknown != bytetype.Zero {
		t.Fatalf("expected Invalid and Unknown to equal Zero")
	}
}

func TestExtractors(t *testing.T) {
	v := bytetype.Two
	if v.Byte() != 2 || v.ValueByte() != 2 || v.Value() != 2 {
		t.Fatalf("byte extraction mismatch")
	}

	if v.Int() != 2 || v.ValueInt() != 2 || v.ValueInt8() != 2 {
		t.Fatalf("int extraction mismatch")
	}

	if v.ValueInt16() != 2 || v.ValueInt32() != 2 || v.ValueUInt16() != 2 || v.Code() != 2 {
		t.Fatalf("extended int extraction mismatch")
	}

	bytes := v.Bytes()
	if len(bytes) != 1 || bytes[0] != 2 {
		t.Fatalf("Bytes() mismatch")
	}

	if *v.ToPtr() != 2 {
		t.Fatalf("ToPtr() mismatch")
	}
}

func TestPredicates(t *testing.T) {
	if !bytetype.Zero.IsZero() || !bytetype.Min.IsMin() {
		t.Fatalf("IsZero/IsMin failed")
	}

	if !bytetype.One.IsOne() || !bytetype.Two.IsTwo() || !bytetype.Three.IsThree() {
		t.Fatalf("IsOne/IsTwo/IsThree failed")
	}

	if !bytetype.Max.IsMax() {
		t.Fatalf("IsMax failed")
	}

	if bytetype.Zero.IsValid() || !bytetype.Zero.IsInvalid() {
		t.Fatalf("Zero validity failed")
	}

	if !bytetype.One.IsValid() || !bytetype.One.IsEnum() {
		t.Fatalf("One validity failed")
	}

	if !bytetype.Two.Is(bytetype.Two) {
		t.Fatalf("Is() failed")
	}
}

func TestComparisons(t *testing.T) {
	v := bytetype.Two
	if !v.IsEqual(2) || !v.IsEqualInt(2) || !v.IsValueEqual(2) {
		t.Fatalf("equality failed")
	}

	if !v.IsGreater(1) || !v.IsGreaterInt(1) || !v.IsGreaterEqual(2) || !v.IsGreaterEqualInt(2) {
		t.Fatalf("greater failed")
	}

	if !v.IsLess(3) || !v.IsLessInt(3) || !v.IsLessEqual(2) || !v.IsLessEqualInt(2) {
		t.Fatalf("less failed")
	}

	if !v.IsBetween(1, 3) || !v.IsBetweenInt(1, 3) {
		t.Fatalf("between failed")
	}

	if !v.IsNameEqual("Two") || !v.IsAnyNamesOf("Zero", "Two", "Three") {
		t.Fatalf("name equal failed")
	}

	if v.IsAnyNamesOf("Zero", "One") {
		t.Fatalf("unexpected name match")
	}
}

func TestArithmetic(t *testing.T) {
	v := bytetype.One
	if v.Add(2) != bytetype.Three {
		t.Fatalf("Add failed")
	}

	if bytetype.Three.Subtract(1) != bytetype.Two {
		t.Fatalf("Subtract failed")
	}
}

func TestHasIndexInStrings(t *testing.T) {
	items := []string{"zero", "one", "two"}
	str, ok := bytetype.Two.HasIndexInStrings(items...)
	if !ok || str != "two" {
		t.Fatalf("HasIndexInStrings valid failed")
	}

	_, ok = bytetype.Three.HasIndexInStrings(items...)
	if ok {
		t.Fatalf("HasIndexInStrings out of bounds should fail")
	}

	_, ok = bytetype.Zero.HasIndexInStrings()
	if ok {
		t.Fatalf("HasIndexInStrings empty should fail")
	}
}

func TestNamingAndFormatting(t *testing.T) {
	if bytetype.Zero.Name() != "Zero" || bytetype.One.Label() != "One" || bytetype.Two.String() != "Two" {
		t.Fatalf("naming failed")
	}

	if bytetype.Max.Name() != "Max" {
		t.Fatalf("Max name failed")
	}

	custom := bytetype.Variant(42)
	if custom.Name() != "Byte(42)" {
		t.Fatalf("custom Name() expected 'Byte(42)', got %s", custom.Name())
	}

	if custom.ValueString() != "42" || custom.StringValue() != "42" || custom.ToNumberString() != "42" {
		t.Fatalf("numeric strings mismatch")
	}

	if bytetype.One.NameValue() != "One(1)" {
		t.Fatalf("NameValue mismatch")
	}

	if bytetype.One.JsonString() != `"One"` {
		t.Fatalf("JsonString mismatch")
	}
}

func TestFunctions(t *testing.T) {
	if bytetype.New(5) != bytetype.Variant(5) {
		t.Fatalf("New failed")
	}

	if bytetype.GetSet(true, bytetype.One, bytetype.Two) != bytetype.One {
		t.Fatalf("GetSet true failed")
	}

	if bytetype.GetSet(false, bytetype.One, bytetype.Two) != bytetype.Two {
		t.Fatalf("GetSet false failed")
	}

	if bytetype.GetSetVariant(true, 1, 2) != bytetype.One {
		t.Fatalf("GetSetVariant true failed")
	}

	if bytetype.GetSetVariant(false, 1, 2) != bytetype.Two {
		t.Fatalf("GetSetVariant false failed")
	}

	if bytetype.String([]byte("hello")) != "hello" || bytetype.String(nil) != "" {
		t.Fatalf("String helper failed")
	}
}

func TestAllAndValues(t *testing.T) {
	all := bytetype.All()
	if len(all) != 5 {
		t.Fatalf("expected 5 variants, got %d", len(all))
	}

	vals := bytetype.Values()
	if len(vals) != 5 {
		t.Fatalf("expected 5 values, got %d", len(vals))
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected bytetype.Variant
		success  bool
	}{
		{"Zero", bytetype.Zero, true},
		{"zero", bytetype.Zero, true},
		{"ONE", bytetype.One, true},
		{"Two", bytetype.Two, true},
		{"three", bytetype.Three, true},
		{"Max", bytetype.Max, true},
		{"min", bytetype.Min, true},
		{"42", bytetype.Variant(42), true},
		{"255", bytetype.Max, true},
		{"", bytetype.Zero, false},
		{"invalid-name", bytetype.Zero, false},
		{"256", bytetype.Zero, false},
	}

	for _, tc := range tests {
		res := bytetype.Parse(tc.input)
		if tc.success {
			if !res.IsSuccess() || res.Data() != tc.expected {
				t.Errorf("parse %q failed: %v", tc.input, res.Fault())
			}
		} else {
			if res.IsSuccess() {
				t.Errorf("parse %q expected failure", tc.input)
			}
		}
	}
}

func TestJSONRoundtrip(t *testing.T) {
	v := bytetype.Two
	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var decoded bytetype.Variant
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if decoded != v {
		t.Fatalf("expected %v, got %v", v, decoded)
	}

	var numDecoded bytetype.Variant
	if err := json.Unmarshal([]byte("2"), &numDecoded); err != nil {
		t.Fatalf("unmarshal numeric failed: %v", err)
	}

	if numDecoded != bytetype.Two {
		t.Fatalf("expected Two, got %v", numDecoded)
	}

	var nullDecoded bytetype.Variant
	if err := json.Unmarshal([]byte("null"), &nullDecoded); err != nil {
		t.Fatalf("unmarshal null failed: %v", err)
	}

	if nullDecoded != bytetype.Zero {
		t.Fatalf("expected Zero for null, got %v", nullDecoded)
	}
}
