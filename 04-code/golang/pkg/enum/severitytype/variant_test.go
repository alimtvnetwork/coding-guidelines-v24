package severitytype_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/enum/severitytype"
)

func TestSeverityType_Interfaces(t *testing.T) {
	var (
		_ baseenumer.BaseEnumer   = severitytype.Error
		_ baseenumer.ByteEnumer   = severitytype.Error
		_ baseenumer.NumberEnumer = severitytype.Error
		_ json.Marshaler          = severitytype.Error
		_ json.Unmarshaler        = (*severitytype.Variant)(nil)
	)
}

func TestSeverityType_Properties(t *testing.T) {
	if severitytype.Error.Byte() != 3 {
		t.Fatalf("expected Error.Byte() to be 3")
	}

	if severitytype.Error.ValueByte() != 3 {
		t.Fatalf("expected Error.ValueByte() to be 3")
	}

	if len(severitytype.Error.Bytes()) != 1 || severitytype.Error.Bytes()[0] != 3 {
		t.Fatalf("expected Error.Bytes() to be [3]")
	}

	if severitytype.Error.Int() != 3 || severitytype.Error.Code() != 3 {
		t.Fatalf("expected Error.Int() and Code() to be 3")
	}

	if !severitytype.Error.IsValid() || severitytype.Error.IsInvalid() {
		t.Fatalf("expected Error to be valid")
	}

	if !severitytype.Error.IsEnum() {
		t.Fatalf("expected Error.IsEnum() to be true")
	}

	if severitytype.Unknown.IsValid() || !severitytype.Unknown.IsInvalid() {
		t.Fatalf("expected Unknown to be invalid")
	}
}

func TestSeverityType_Predicates(t *testing.T) {
	if !severitytype.Unknown.IsUnknown() || severitytype.Error.IsUnknown() {
		t.Fatalf("unexpected IsUnknown result")
	}

	if !severitytype.Info.IsInfo() || severitytype.Error.IsInfo() {
		t.Fatalf("unexpected IsInfo result")
	}

	if !severitytype.Warn.IsWarn() || severitytype.Error.IsWarn() {
		t.Fatalf("unexpected IsWarn result")
	}

	if !severitytype.Error.IsError() || severitytype.Warn.IsError() {
		t.Fatalf("unexpected IsError result")
	}

	if !severitytype.Critical.IsCritical() || severitytype.Error.IsCritical() {
		t.Fatalf("unexpected IsCritical result")
	}

	if !severitytype.Fatal.IsFatal() || severitytype.Error.IsFatal() {
		t.Fatalf("unexpected IsFatal result")
	}
}

func TestSeverityType_Names(t *testing.T) {
	if severitytype.Error.Name() != "Error" || severitytype.Error.Label() != "Error" {
		t.Fatalf("unexpected name")
	}

	if severitytype.Error.String() != "Error" {
		t.Fatalf("unexpected stringer")
	}

	if severitytype.Error.ValueString() != "Error(3)" {
		t.Fatalf("unexpected ValueString: %s", severitytype.Error.ValueString())
	}

	unknown := severitytype.Variant(99)
	if unknown.Name() != "Severity(99)" {
		t.Fatalf("unexpected fallback name: %s", unknown.Name())
	}
}

func TestSeverityType_Vars(t *testing.T) {
	all := severitytype.All()
	if len(all) != 5 {
		t.Fatalf("expected 5 valid variants (excluding Unknown), got %d", len(all))
	}

	vals := severitytype.Values()
	if len(vals) != 5 {
		t.Fatalf("expected 5 valid values, got %d", len(vals))
	}

	v, ok := severitytype.Parse("Critical")
	if !ok || v != severitytype.Critical {
		t.Fatalf("expected Critical parse success")
	}

	vLower, okLower := severitytype.Parse("fatal")
	if !okLower || vLower != severitytype.Fatal {
		t.Fatalf("expected fatal parse success")
	}

	_, okEmpty := severitytype.Parse("   ")
	if okEmpty {
		t.Fatalf("expected empty parse to fail")
	}

	vUnknown := severitytype.ParseOrUnknown("NonExistent")
	if vUnknown != severitytype.Unknown {
		t.Fatalf("expected ParseOrUnknown to return Unknown")
	}
}

func TestSeverityType_JSON(t *testing.T) {
	data, err := json.Marshal(severitytype.Critical)
	if err != nil || string(data) != `"Critical"` {
		t.Fatalf("json marshal failed: %v", err)
	}

	var v severitytype.Variant
	if err := json.Unmarshal([]byte(`"Critical"`), &v); err != nil || v != severitytype.Critical {
		t.Fatalf("json unmarshal string failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`null`), &v); err != nil || v != severitytype.Unknown {
		t.Fatalf("json unmarshal null failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`4`), &v); err != nil || v != severitytype.Critical {
		t.Fatalf("json unmarshal numeric failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`99`), &v); err == nil {
		t.Fatalf("expected error for numeric out of bounds")
	}

	if err := json.Unmarshal([]byte(`"UnknownSeverity"`), &v); err == nil {
		t.Fatalf("expected error for invalid severity name")
	}
}
