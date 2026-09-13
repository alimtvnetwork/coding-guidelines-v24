package ostype_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/enum/ostype"
)

func TestOSTypeType_Interfaces(t *testing.T) {
	var _ baseenumer.BaseEnumer = ostype.AnyOs
	var _ baseenumer.ByteEnumer = ostype.AnyOs
	var _ baseenumer.NumberEnumer = ostype.AnyOs
	var _ baseenumer.MinMaxer[ostype.Variant] = ostype.AnyOs
	var _ baseenumer.BoundedEnumer[ostype.Variant] = ostype.AnyOs
	var _ json.Marshaler = ostype.AnyOs
	var _ json.Unmarshaler = (*ostype.Variant)(nil)
}

func TestOSTypeType_Properties(t *testing.T) {
	if ostype.AnyOs.Byte() != 1 || ostype.AnyOs.Value() != 1 {
		t.Fatalf("expected 1 from Byte/Value()")
	}

	if ostype.AnyOs.Int() != 1 || ostype.AnyOs.Code() != 1 {
		t.Fatalf("expected 1 from Int/Code")
	}

	if !ostype.AnyOs.IsValid() {
		t.Fatalf("expected AnyOs to be valid")
	}
}

func TestOSTypeType_Boundaries(t *testing.T) {
	minVal, maxVal := ostype.Min(), ostype.Max()
	if minVal.Min() != minVal || maxVal.Max() != maxVal {
		t.Fatalf("receiver Min/Max mismatch")
	}

	if !minVal.IsMin() || !maxVal.IsMax() {
		t.Fatalf("boundary predicates failed")
	}

	if !minVal.IsInRange(minVal, maxVal) {
		t.Fatalf("IsInRange failed")
	}
}

func TestOSTypeType_Predicates(t *testing.T) {
	if !ostype.AnyOs.IsAnyOs() {
		t.Fatalf("expected IsAnyOs() to be true")
	}

	if ostype.AnyOs.IsInvalid() {
		t.Fatalf("expected IsInvalid() to be false for AnyOs")
	}
}

func TestOSTypeType_VarsAndParse(t *testing.T) {
	all := ostype.All()
	if len(all) != 13 || len(ostype.AnyOs.All()) != 13 {
		t.Fatalf("expected 13 variants, got %d", len(all))
	}

	vals := ostype.Values()
	if len(vals) != 13 || len(ostype.AnyOs.Values()) != 13 {
		t.Fatalf("expected 13 values, got %d", len(vals))
	}

	v, isOk := ostype.Parse("AnyOs")
	if !isOk || v != ostype.AnyOs {
		t.Fatalf("expected successful Parse for AnyOs")
	}

	if _, isOk := ostype.Parse(""); isOk {
		t.Fatalf("expected failure for empty string")
	}

	if _, isOk := ostype.Parse("invalid_variant_value"); isOk {
		t.Fatalf("expected failure for bad variant")
	}

	if ostype.ParseOrZero("AnyOs") != ostype.AnyOs {
		t.Fatalf("ParseOrZero failed")
	}

	if ostype.ParseOrInvalid("invalid_variant_value") != ostype.Invalid {
		t.Fatalf("ParseOrInvalid fallback failed")
	}
}

func TestOSTypeType_JSON(t *testing.T) {
	data, err := json.Marshal(ostype.AnyOs)
	if err != nil || string(data) != `"AnyOs"` {
		t.Fatalf("marshal failed: %v", err)
	}

	var v ostype.Variant
	if err := json.Unmarshal([]byte(`"AnyOs"`), &v); err != nil || v != ostype.AnyOs {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`null`), &v); err != nil || v != ostype.Invalid {
		t.Fatalf("unmarshal null failed: %v", err)
	}
}
