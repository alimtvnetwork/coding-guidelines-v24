package filewritemodetype_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/enum/filewritemodetype"
)

func TestFileWriteModeType_Interfaces(t *testing.T) {
	var (
		_ baseenumer.BaseEnumer                               = filewritemodetype.Direct
		_ baseenumer.ByteEnumer                               = filewritemodetype.Direct
		_ baseenumer.NumberEnumer                             = filewritemodetype.Direct
		_ baseenumer.BoundedEnumer[filewritemodetype.Variant] = filewritemodetype.Direct
		_ json.Marshaler                                      = filewritemodetype.Direct
		_ json.Unmarshaler                                    = (*filewritemodetype.Variant)(nil)
	)
}

func TestFileWriteModeType_Properties(t *testing.T) {
	if filewritemodetype.Direct.Byte() != 1 {
		t.Fatalf("expected Direct.Byte() to be 1")
	}

	if filewritemodetype.Direct.ValueByte() != 1 {
		t.Fatalf("expected Direct.ValueByte() to be 1")
	}

	if len(filewritemodetype.Direct.Bytes()) != 1 || filewritemodetype.Direct.Bytes()[0] != 1 {
		t.Fatalf("expected Direct.Bytes() to be [1]")
	}

	if filewritemodetype.Direct.Int() != 1 || filewritemodetype.Direct.Code() != 1 {
		t.Fatalf("expected Direct.Int() and Code() to be 1")
	}

	if !filewritemodetype.Direct.IsValid() || filewritemodetype.Direct.IsInvalid() {
		t.Fatalf("expected Direct to be valid")
	}

	if !filewritemodetype.Direct.IsEnum() {
		t.Fatalf("expected Direct.IsEnum() to be true")
	}

	if filewritemodetype.Invalid.IsValid() || !filewritemodetype.Invalid.IsInvalid() {
		t.Fatalf("expected Invalid to be invalid")
	}
}

func TestFileWriteModeType_Predicates(t *testing.T) {
	if !filewritemodetype.Direct.IsDirect() || filewritemodetype.Atomic.IsDirect() {
		t.Fatalf("unexpected IsDirect result")
	}

	if !filewritemodetype.Atomic.IsAtomic() || filewritemodetype.Direct.IsAtomic() {
		t.Fatalf("unexpected IsAtomic result")
	}

	if !filewritemodetype.Truncate.IsTruncate() || filewritemodetype.Direct.IsTruncate() {
		t.Fatalf("unexpected IsTruncate result")
	}
}

func TestFileWriteModeType_Names(t *testing.T) {
	if filewritemodetype.Direct.Name() != "Direct" || filewritemodetype.Direct.Label() != "Direct" {
		t.Fatalf("unexpected name")
	}

	if filewritemodetype.Direct.String() != "Direct" {
		t.Fatalf("unexpected stringer")
	}

	if filewritemodetype.Direct.ValueString() != "Direct(1)" {
		t.Fatalf("unexpected ValueString: %s", filewritemodetype.Direct.ValueString())
	}

	unknown := filewritemodetype.Variant(99)
	if unknown.Name() != "FileWriteMode(99)" {
		t.Fatalf("unexpected fallback name: %s", unknown.Name())
	}
}

func TestFileWriteModeType_Vars(t *testing.T) {
	all := filewritemodetype.All()
	if len(all) != 3 {
		t.Fatalf("expected 3 valid variants, got %d", len(all))
	}

	vals := filewritemodetype.Values()
	if len(vals) != 3 {
		t.Fatalf("expected 3 valid values, got %d", len(vals))
	}

	parsed, isOk := filewritemodetype.Parse("Atomic")
	if !isOk || parsed != filewritemodetype.Atomic {
		t.Fatalf("expected Atomic parse success")
	}

	parsedLower, isLowerOk := filewritemodetype.Parse("truncate")
	if !isLowerOk || parsedLower != filewritemodetype.Truncate {
		t.Fatalf("expected truncate parse success")
	}

	_, isEmptyOk := filewritemodetype.Parse("   ")
	if isEmptyOk {
		t.Fatalf("expected empty parse error")
	}

	_, isInvalidOk := filewritemodetype.Parse("NonExistent")
	if isInvalidOk {
		t.Fatalf("expected invalid parse error")
	}

	if filewritemodetype.ParseOrZero("Atomic") != filewritemodetype.Atomic {
		t.Fatalf("ParseOrZero Atomic failed")
	}

	if filewritemodetype.ParseOrInvalid("bogus") != filewritemodetype.Invalid {
		t.Fatalf("ParseOrInvalid bogus failed")
	}

	if filewritemodetype.ParseOrUnknown("bogus") != filewritemodetype.Invalid {
		t.Fatalf("ParseOrUnknown bogus failed")
	}
}

func TestFileWriteModeType_JSON(t *testing.T) {
	data, err := json.Marshal(filewritemodetype.Atomic)
	if err != nil || string(data) != `"Atomic"` {
		t.Fatalf("json marshal failed: %v", err)
	}

	var v filewritemodetype.Variant
	if err := json.Unmarshal([]byte(`"Atomic"`), &v); err != nil || v != filewritemodetype.Atomic {
		t.Fatalf("json unmarshal string failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`null`), &v); err != nil || v != filewritemodetype.Invalid {
		t.Fatalf("json unmarshal null failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`2`), &v); err != nil || v != filewritemodetype.Atomic {
		t.Fatalf("json unmarshal numeric failed: %v", err)
	}

	if err := json.Unmarshal([]byte(`99`), &v); err == nil {
		t.Fatalf("expected error for numeric out of bounds")
	}

	if err := json.Unmarshal([]byte(`"UnknownMode"`), &v); err == nil {
		t.Fatalf("expected error for invalid mode name")
	}
}

func TestFileWriteModeType_Boundaries(t *testing.T) {
	var _ baseenumer.BoundedEnumer[filewritemodetype.Variant] = filewritemodetype.Variant(0)

	minVal, maxVal := filewritemodetype.Min(), filewritemodetype.Max()
	if minVal != filewritemodetype.Invalid || maxVal != filewritemodetype.Truncate {
		t.Fatalf("expected min %v, max %v", filewritemodetype.Invalid, filewritemodetype.Truncate)
	}

	if minVal.Min() != minVal || maxVal.Max() != maxVal {
		t.Fatalf("receiver Min/Max mismatch")
	}
}

func TestFileWriteModeType_BoundaryPredicates(t *testing.T) {
	minVal, maxVal := filewritemodetype.Min(), filewritemodetype.Max()
	if !minVal.IsMin() || !maxVal.IsMax() {
		t.Fatalf("boundary predicates failed")
	}

	if maxVal.IsMin() || minVal.IsMax() {
		t.Fatalf("inverse boundary predicates failed")
	}

	if !minVal.IsInRange(minVal, maxVal) || !filewritemodetype.Direct.IsInRange(minVal, maxVal) {
		t.Fatalf("IsInRange failed for valid range")
	}

	if filewritemodetype.Variant(99).IsInRange(minVal, maxVal) {
		t.Fatalf("IsInRange succeeded for out-of-range value")
	}
}
