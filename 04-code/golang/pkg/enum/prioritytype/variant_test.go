package prioritytype_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/enum/prioritytype"
)

func TestPriorityType_Interfaces(t *testing.T) {
	var (
		_ baseenumer.BaseEnumer                          = prioritytype.High
		_ baseenumer.ByteEnumer                          = prioritytype.High
		_ baseenumer.NumberEnumer                        = prioritytype.High
		_ baseenumer.MinMaxer[prioritytype.Variant]      = prioritytype.High
		_ baseenumer.BoundedEnumer[prioritytype.Variant] = prioritytype.High
		_ baseenumer.Bounder[prioritytype.Variant]       = prioritytype.High
		_ json.Marshaler                                 = prioritytype.High
		_ json.Unmarshaler                               = (*prioritytype.Variant)(nil)
	)
}

func TestPriorityType_Properties(t *testing.T) {
	if prioritytype.High.Byte() != 3 {
		t.Fatalf("expected High.Byte() to be 3")
	}

	if prioritytype.High.ValueByte() != 3 {
		t.Fatalf("expected High.ValueByte() to be 3")
	}

	if len(prioritytype.High.Bytes()) != 1 || prioritytype.High.Bytes()[0] != 3 {
		t.Fatalf("expected High.Bytes() to be [3]")
	}

	if prioritytype.High.Int() != 3 || prioritytype.High.Code() != 3 {
		t.Fatalf("expected High.Int() and Code() to be 3")
	}

	if !prioritytype.High.IsValid() || prioritytype.High.IsInvalid() {
		t.Fatalf("expected High to be valid")
	}

	if !prioritytype.High.IsEnum() {
		t.Fatalf("expected High.IsEnum() to be true")
	}

	if prioritytype.Unknown.IsValid() || !prioritytype.Unknown.IsInvalid() {
		t.Fatalf("expected Unknown to be invalid")
	}
}

func TestPriorityType_Predicates(t *testing.T) {
	if !prioritytype.Unknown.IsUnknown() || prioritytype.High.IsUnknown() {
		t.Fatalf("unexpected IsUnknown result")
	}

	if !prioritytype.Low.IsLow() || prioritytype.High.IsLow() {
		t.Fatalf("unexpected IsLow result")
	}

	if !prioritytype.Normal.IsNormal() || prioritytype.High.IsNormal() {
		t.Fatalf("unexpected IsNormal result")
	}

	if !prioritytype.High.IsHigh() || prioritytype.Normal.IsHigh() {
		t.Fatalf("unexpected IsHigh result")
	}

	if !prioritytype.Critical.IsCritical() || prioritytype.High.IsCritical() {
		t.Fatalf("unexpected IsCritical result")
	}
}

func TestPriorityType_Names(t *testing.T) {
	if prioritytype.High.Name() != "High" || prioritytype.High.Label() != "High" {
		t.Fatalf("unexpected name")
	}

	if prioritytype.High.String() != "High" {
		t.Fatalf("unexpected stringer")
	}

	if prioritytype.High.ValueString() != "High(3)" {
		t.Fatalf("unexpected ValueString: %s", prioritytype.High.ValueString())
	}

	unknown := prioritytype.Variant(99)
	if unknown.Name() != "Priority(99)" {
		t.Fatalf("unexpected fallback name: %s", unknown.Name())
	}
}

func TestPriorityType_Vars(t *testing.T) {
	all := prioritytype.All()
	if len(all) != 4 {
		t.Fatalf("expected 4 valid variants (excluding Unknown), got %d", len(all))
	}

	vals := prioritytype.Values()
	if len(vals) != 4 {
		t.Fatalf("expected 4 valid values, got %d", len(vals))
	}

	v, ok := prioritytype.Parse("Critical")
	if !ok || v != prioritytype.Critical {
		t.Fatalf("expected Critical parse success")
	}

	vLower, okLower := prioritytype.Parse("high")
	if !okLower || vLower != prioritytype.High {
		t.Fatalf("expected high parse success")
	}

	_, okEmpty := prioritytype.Parse("   ")
	if okEmpty {
		t.Fatalf("expected empty parse to fail")
	}

	vUnknown := prioritytype.ParseOrUnknown("NonExistent")
	if vUnknown != prioritytype.Unknown {
		t.Fatalf("expected ParseOrUnknown to return Unknown")
	}

	if prioritytype.ParseOrZero("Critical") != prioritytype.Critical {
		t.Fatalf("expected ParseOrZero Critical success")
	}

	if prioritytype.ParseOrInvalid("NonExistent") != prioritytype.Unknown {
		t.Fatalf("expected ParseOrInvalid to return Unknown")
	}
}

func TestPriorityType_JSON(t *testing.T) {
	data, err := json.Marshal(prioritytype.High)
	if err != nil || string(data) != `"High"` {
		t.Fatalf("json marshal failed: %v", err)
	}

	var v prioritytype.Variant
	if err := json.Unmarshal([]byte(`"High"`), &v); err != nil || v != prioritytype.High {
		t.Fatalf("json unmarshal string failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`null`), &v); err != nil || v != prioritytype.Unknown {
		t.Fatalf("json unmarshal null failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`3`), &v); err != nil || v != prioritytype.High {
		t.Fatalf("json unmarshal numeric failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`99`), &v); err == nil {
		t.Fatalf("expected error for numeric out of bounds")
	}

	if err := json.Unmarshal([]byte(`"UnknownPriority"`), &v); err == nil {
		t.Fatalf("expected error for invalid priority name")
	}
}

func TestPriorityType_Boundary(t *testing.T) {
	if prioritytype.Min() != prioritytype.Unknown {
		t.Fatalf("expected Min to be Unknown")
	}

	if prioritytype.Max() != prioritytype.Critical {
		t.Fatalf("expected Max to be Critical")
	}

	if !prioritytype.Unknown.IsMin() {
		t.Fatalf("expected Unknown.IsMin() to be true")
	}

	if prioritytype.Critical.IsMin() {
		t.Fatalf("expected Critical.IsMin() to be false")
	}
}

func TestPriorityType_BoundaryMax(t *testing.T) {
	if !prioritytype.Critical.IsMax() {
		t.Fatalf("expected Critical.IsMax() to be true")
	}

	if prioritytype.Unknown.IsMax() {
		t.Fatalf("expected Unknown.IsMax() to be false")
	}

	if prioritytype.High.Min() != prioritytype.Unknown {
		t.Fatalf("expected High.Min() to be Unknown")
	}

	if prioritytype.High.Max() != prioritytype.Critical {
		t.Fatalf("expected High.Max() to be Critical")
	}
}

func TestPriorityType_IsInRange(t *testing.T) {
	if !prioritytype.Normal.IsInRange(prioritytype.Low, prioritytype.Critical) {
		t.Fatalf("expected Normal to be in range [Low, Critical]")
	}

	if prioritytype.Unknown.IsInRange(prioritytype.Low, prioritytype.Critical) {
		t.Fatalf("expected Unknown to not be in range [Low, Critical]")
	}

	if !prioritytype.Critical.IsInRange(prioritytype.Unknown, prioritytype.Critical) {
		t.Fatalf("expected Critical to be in range [Unknown, Critical]")
	}
}

func TestPriorityType_ValueAndCollections(t *testing.T) {
	h := prioritytype.High
	if h.Value() != 3 {
		t.Fatalf("expected 3 from Value(), got %d", h.Value())
	}

	all := prioritytype.All()
	if len(all) != 4 || len(h.All()) != 4 {
		t.Fatalf("expected 4 variants from All")
	}

	vals := prioritytype.Values()
	if len(vals) != 4 || len(h.Values()) != 4 {
		t.Fatalf("expected 4 values from Values")
	}
}
