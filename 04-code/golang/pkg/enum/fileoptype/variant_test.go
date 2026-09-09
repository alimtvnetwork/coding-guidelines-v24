package fileoptype_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/enum/fileoptype"
	"coding-guidelines/common/pkg/enum/openfiletype"
)

func TestFileOpType_Interfaces(t *testing.T) {
	var (
		_ baseenumer.BaseEnumer                        = fileoptype.ReadOnly
		_ baseenumer.ByteEnumer                        = fileoptype.ReadOnly
		_ baseenumer.NumberEnumer                      = fileoptype.ReadOnly
		_ baseenumer.BoundedEnumer[fileoptype.Variant] = fileoptype.ReadOnly
		_ json.Marshaler                               = fileoptype.ReadOnly
		_ json.Unmarshaler                             = (*fileoptype.Variant)(nil)
	)
}

func TestFileOpType_Properties(t *testing.T) {
	if fileoptype.ReadOnly.Byte() != 1 {
		t.Fatalf("expected ReadOnly.Byte() to be 1")
	}

	if fileoptype.ReadOnly.ValueByte() != 1 {
		t.Fatalf("expected ReadOnly.ValueByte() to be 1")
	}

	if len(fileoptype.ReadOnly.Bytes()) != 1 || fileoptype.ReadOnly.Bytes()[0] != 1 {
		t.Fatalf("expected ReadOnly.Bytes() to be [1]")
	}

	if fileoptype.ReadOnly.Int() != 1 || fileoptype.ReadOnly.Code() != 1 {
		t.Fatalf("expected ReadOnly.Int() and Code() to be 1")
	}

	if !fileoptype.ReadOnly.IsValid() || fileoptype.ReadOnly.IsInvalid() {
		t.Fatalf("expected ReadOnly to be valid")
	}

	if !fileoptype.ReadOnly.IsEnum() {
		t.Fatalf("expected ReadOnly.IsEnum() to be true")
	}

	if fileoptype.Invalid.IsValid() || !fileoptype.Invalid.IsInvalid() {
		t.Fatalf("expected Invalid to be invalid")
	}
}

func TestFileOpType_Predicates(t *testing.T) {
	if !fileoptype.ReadOnly.IsReadOnly() || fileoptype.WriteOnly.IsReadOnly() {
		t.Fatalf("unexpected IsReadOnly result")
	}

	if !fileoptype.WriteOnly.IsWriteOnly() || fileoptype.ReadOnly.IsWriteOnly() {
		t.Fatalf("unexpected IsWriteOnly result")
	}

	if !fileoptype.ReadWrite.IsReadWrite() || fileoptype.ReadOnly.IsReadWrite() {
		t.Fatalf("unexpected IsReadWrite result")
	}

	if !fileoptype.Append.IsAppend() || !fileoptype.CreateAppend.IsAppend() {
		t.Fatalf("expected Append and CreateAppend to be append")
	}

	if !fileoptype.Create.IsCreate() || !fileoptype.CreateAppend.IsCreateAppend() {
		t.Fatalf("unexpected Create or CreateAppend predicate")
	}

	if !fileoptype.CreateTruncate.IsCreateTruncate() || !fileoptype.Delete.IsDelete() {
		t.Fatalf("unexpected CreateTruncate or Delete predicate")
	}
}

func TestFileOpType_OpenMode(t *testing.T) {
	tests := []struct {
		op       fileoptype.Variant
		expected openfiletype.Variant
	}{
		{fileoptype.ReadOnly, openfiletype.ReadOnly},
		{fileoptype.WriteOnly, openfiletype.WriteOnly},
		{fileoptype.ReadWrite, openfiletype.ReadWrite},
		{fileoptype.Append, openfiletype.Append},
		{fileoptype.Create, openfiletype.CreateNew},
		{fileoptype.CreateAppend, openfiletype.CreateAppend},
		{fileoptype.CreateTruncate, openfiletype.CreateTruncate},
		{fileoptype.Delete, openfiletype.ReadOnly},
	}

	for _, tt := range tests {
		if tt.op.OpenMode() != tt.expected {
			t.Fatalf("for %v expected OpenMode %v, got %v", tt.op, tt.expected, tt.op.OpenMode())
		}
	}
}

func TestFileOpType_Names(t *testing.T) {
	if fileoptype.ReadOnly.Name() != "ReadOnly" || fileoptype.ReadOnly.Label() != "ReadOnly" {
		t.Fatalf("unexpected name")
	}

	if fileoptype.ReadOnly.String() != "ReadOnly" {
		t.Fatalf("unexpected stringer")
	}

	if fileoptype.ReadOnly.ValueString() != "ReadOnly(1)" {
		t.Fatalf("unexpected ValueString: %s", fileoptype.ReadOnly.ValueString())
	}

	unknown := fileoptype.Variant(99)
	if unknown.Name() != "FileOp(99)" {
		t.Fatalf("unexpected fallback name: %s", unknown.Name())
	}
}

func TestFileOpType_Vars(t *testing.T) {
	all := fileoptype.All()
	if len(all) != 8 {
		t.Fatalf("expected 8 valid variants, got %d", len(all))
	}

	vals := fileoptype.Values()
	if len(vals) != 8 {
		t.Fatalf("expected 8 valid values, got %d", len(vals))
	}

	parsed, isOk := fileoptype.Parse("Delete")
	if !isOk || parsed != fileoptype.Delete {
		t.Fatalf("expected Delete parse success")
	}

	parsedLower, isLowerOk := fileoptype.Parse("createappend")
	if !isLowerOk || parsedLower != fileoptype.CreateAppend {
		t.Fatalf("expected createappend parse success")
	}

	_, isEmptyOk := fileoptype.Parse("   ")
	if isEmptyOk {
		t.Fatalf("expected empty parse error")
	}

	_, isInvalidOk := fileoptype.Parse("NonExistent")
	if isInvalidOk {
		t.Fatalf("expected invalid parse error")
	}

	if fileoptype.ParseOrZero("Delete") != fileoptype.Delete {
		t.Fatalf("ParseOrZero Delete failed")
	}

	if fileoptype.ParseOrInvalid("bogus") != fileoptype.Invalid {
		t.Fatalf("ParseOrInvalid bogus failed")
	}

	if fileoptype.ParseOrUnknown("bogus") != fileoptype.Invalid {
		t.Fatalf("ParseOrUnknown bogus failed")
	}
}

func TestFileOpType_JSON(t *testing.T) {
	data, err := json.Marshal(fileoptype.CreateTruncate)
	if err != nil || string(data) != `"CreateTruncate"` {
		t.Fatalf("json marshal failed: %v", err)
	}

	var v fileoptype.Variant
	if err := json.Unmarshal([]byte(`"CreateTruncate"`), &v); err != nil || v != fileoptype.CreateTruncate {
		t.Fatalf("json unmarshal string failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`null`), &v); err != nil || v != fileoptype.Invalid {
		t.Fatalf("json unmarshal null failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`7`), &v); err != nil || v != fileoptype.CreateTruncate {
		t.Fatalf("json unmarshal numeric failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`99`), &v); err == nil {
		t.Fatalf("expected error for numeric out of bounds")
	}

	if err := json.Unmarshal([]byte(`"UnknownOp"`), &v); err == nil {
		t.Fatalf("expected error for invalid op name")
	}
}

func TestFileOpType_Boundaries(t *testing.T) {
	var _ baseenumer.BoundedEnumer[fileoptype.Variant] = fileoptype.Variant(0)

	minVal, maxVal := fileoptype.Min(), fileoptype.Max()
	if minVal != fileoptype.Invalid || maxVal != fileoptype.Delete {
		t.Fatalf("expected min %v, max %v", fileoptype.Invalid, fileoptype.Delete)
	}

	if minVal.Min() != minVal || maxVal.Max() != maxVal {
		t.Fatalf("receiver Min/Max mismatch")
	}
}

func TestFileOpType_BoundaryPredicates(t *testing.T) {
	minVal, maxVal := fileoptype.Min(), fileoptype.Max()
	if !minVal.IsMin() || !maxVal.IsMax() {
		t.Fatalf("boundary predicates failed")
	}

	if maxVal.IsMin() || minVal.IsMax() {
		t.Fatalf("inverse boundary predicates failed")
	}

	if !minVal.IsInRange(minVal, maxVal) || !fileoptype.ReadOnly.IsInRange(minVal, maxVal) {
		t.Fatalf("IsInRange failed for valid range")
	}

	if fileoptype.Variant(99).IsInRange(minVal, maxVal) {
		t.Fatalf("IsInRange succeeded for out-of-range value")
	}
}

func TestFileOpType_ValueAndCollections(t *testing.T) {
	ro := fileoptype.ReadOnly
	if ro.Value() != 1 {
		t.Fatalf("expected 1 from Value(), got %d", ro.Value())
	}

	all := fileoptype.All()
	if len(all) != 8 || len(ro.All()) != 8 {
		t.Fatalf("expected 8 variants, got %d", len(all))
	}

	vals := fileoptype.Values()
	if len(vals) != 8 || len(ro.Values()) != 8 {
		t.Fatalf("expected 8 values, got %d", len(vals))
	}
}
