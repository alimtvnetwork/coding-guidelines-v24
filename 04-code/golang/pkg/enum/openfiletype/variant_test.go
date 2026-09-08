package openfiletype_test

import (
	"encoding/json"
	"os"
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/enum/openfiletype"
)

func TestInvalidZeroValue(t *testing.T) {
	var zero openfiletype.Variant
	if zero != openfiletype.Invalid {
		t.Fatalf("expected zero value to be Invalid, got %v", zero)
	}

	if zero.IsValid() {
		t.Fatalf("expected Invalid to not be valid")
	}

	if !zero.IsInvalid() {
		t.Fatalf("expected Invalid to be IsInvalid")
	}

	if zero.String() != "Invalid" {
		t.Fatalf("expected String 'Invalid', got %s", zero.String())
	}
}

func TestVariantsAndFlags(t *testing.T) {
	tests := []struct {
		variant       openfiletype.Variant
		expectedName  string
		expectedFlags int
	}{
		{openfiletype.ReadOnly, "ReadOnly", os.O_RDONLY},
		{openfiletype.WriteOnly, "WriteOnly", os.O_WRONLY},
		{openfiletype.ReadWrite, "ReadWrite", os.O_RDWR},
		{openfiletype.Append, "Append", os.O_WRONLY | os.O_APPEND},
		{openfiletype.CreateAppend, "CreateAppend", os.O_CREATE | os.O_WRONLY | os.O_APPEND},
		{openfiletype.CreateTruncate, "CreateTruncate", os.O_CREATE | os.O_WRONLY | os.O_TRUNC},
		{openfiletype.CreateNew, "CreateNew", os.O_CREATE | os.O_EXCL | os.O_WRONLY},
		{openfiletype.ReadOrCreateOnly, "ReadOrCreateOnly", os.O_RDONLY | os.O_CREATE},
		{openfiletype.WriteOrCreateOnly, "WriteOrCreateOnly", os.O_WRONLY | os.O_CREATE},
		{openfiletype.ReadWriteOrCreateOnly, "ReadWriteOrCreateOnly", os.O_RDWR | os.O_CREATE},
	}

	for _, tc := range tests {
		if tc.variant.Name() != tc.expectedName {
			t.Errorf("expected name %s, got %s", tc.expectedName, tc.variant.Name())
		}

		if tc.variant.Flags() != tc.expectedFlags {
			t.Errorf("expected flags %d, got %d for %s", tc.expectedFlags, tc.variant.Flags(), tc.expectedName)
		}

		if !tc.variant.IsValid() {
			t.Errorf("expected %s to be valid", tc.expectedName)
		}
	}
}

func TestCheckers(t *testing.T) {
	ro := openfiletype.ReadOnly
	if !ro.IsReadOnly() {
		t.Fatalf("expected IsReadOnly true")
	}

	if ro.IsWriteOnly() {
		t.Fatalf("expected IsWriteOnly false")
	}

	ca := openfiletype.CreateAppend
	if !ca.IsCreateAppend() {
		t.Fatalf("expected IsCreateAppend true")
	}

	ct := openfiletype.CreateTruncate
	if !ct.IsCreateTruncate() {
		t.Fatalf("expected IsCreateTruncate true")
	}

	cn := openfiletype.CreateNew
	if !cn.IsCreateNew() {
		t.Fatalf("expected IsCreateNew true")
	}

	roco := openfiletype.ReadOrCreateOnly
	if !roco.IsReadOrCreateOnly() {
		t.Fatalf("expected IsReadOrCreateOnly true")
	}

	woco := openfiletype.WriteOrCreateOnly
	if !woco.IsWriteOrCreateOnly() {
		t.Fatalf("expected IsWriteOrCreateOnly true")
	}

	rwoco := openfiletype.ReadWriteOrCreateOnly
	if !rwoco.IsReadWriteOrCreateOnly() {
		t.Fatalf("expected IsReadWriteOrCreateOnly true")
	}

	rw := openfiletype.ReadWrite
	if !rw.IsReadWrite() {
		t.Fatalf("expected IsReadWrite true")
	}

	app := openfiletype.Append
	if !app.IsAppend() {
		t.Fatalf("expected IsAppend true")
	}

	if ro.Label() != "ReadOnly" {
		t.Fatalf("expected Label ReadOnly, got %s", ro.Label())
	}
}

func TestParse(t *testing.T) {
	v, isOk := openfiletype.Parse("createappend")
	if !isOk {
		t.Fatalf("expected parse to succeed")
	}

	if v != openfiletype.CreateAppend {
		t.Fatalf("expected CreateAppend, got %v", v)
	}

	vRO, isOkRO := openfiletype.Parse("readorcreateonly")
	if !isOkRO || vRO != openfiletype.ReadOrCreateOnly {
		t.Fatalf("expected ReadOrCreateOnly, got %v", vRO)
	}

	vWO, isOkWO := openfiletype.Parse("WriteOrCreateOnly")
	if !isOkWO || vWO != openfiletype.WriteOrCreateOnly {
		t.Fatalf("expected WriteOrCreateOnly, got %v", vWO)
	}

	vRWO, isOkRWO := openfiletype.Parse("READWRITEORCREATEONLY")
	if !isOkRWO || vRWO != openfiletype.ReadWriteOrCreateOnly {
		t.Fatalf("expected ReadWriteOrCreateOnly, got %v", vRWO)
	}

	if _, isBadOk := openfiletype.Parse("invalid-mode-string"); isBadOk {
		t.Fatalf("expected failure on bad string")
	}

	if fallback := openfiletype.ParseOrInvalid("invalid-mode-string"); fallback != openfiletype.Invalid {
		t.Fatalf("expected Invalid fallback, got %v", fallback)
	}

	if fallback := openfiletype.ParseOrUnknown("invalid-mode-string"); fallback != openfiletype.Invalid {
		t.Fatalf("expected Invalid fallback on unknown, got %v", fallback)
	}
}

func TestOutOfBounds(t *testing.T) {
	bogus := openfiletype.Variant(99)
	if bogus.Name() != "OpenFile(99)" {
		t.Fatalf("expected OpenFile(99), got %s", bogus.Name())
	}

	if bogus.Flags() != 0 {
		t.Fatalf("expected 0 flags for out of bounds, got %d", bogus.Flags())
	}
}

func TestAllAndValues(t *testing.T) {
	all := openfiletype.All()
	if len(all) != 10 {
		t.Fatalf("expected 10 valid variants, got %d", len(all))
	}

	values := openfiletype.Values()
	if len(values) != 10 {
		t.Fatalf("expected 10 string values, got %d", len(values))
	}
}

func TestJSONRoundtrip(t *testing.T) {
	val := openfiletype.CreateTruncate

	data, err := json.Marshal(val)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	if string(data) != `"CreateTruncate"` {
		t.Fatalf("unexpected json: %s", string(data))
	}

	var parsed openfiletype.Variant
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if parsed != openfiletype.CreateTruncate {
		t.Fatalf("expected CreateTruncate, got %v", parsed)
	}
}

func TestUnmarshalJSON_InvalidCases(t *testing.T) {
	var v openfiletype.Variant

	// Invalid string must return error
	if err := json.Unmarshal([]byte(`"NonExistentMode"`), &v); err == nil {
		t.Fatalf("expected error unmarshaling invalid string, got nil")
	}

	// Invalid numeric byte must return error
	if err := json.Unmarshal([]byte(`99`), &v); err == nil {
		t.Fatalf("expected error unmarshaling invalid numeric byte 99, got nil")
	}

	// Null should unmarshal to Invalid without error
	if err := json.Unmarshal([]byte(`null`), &v); err != nil {
		t.Fatalf("expected nil error on null, got %v", err)
	}

	if v != openfiletype.Invalid {
		t.Fatalf("expected Invalid on null, got %v", v)
	}

	// Valid numeric byte
	if err := json.Unmarshal([]byte(`1`), &v); err != nil {
		t.Fatalf("expected nil error on numeric 1, got %v", err)
	}

	if v != openfiletype.ReadOnly {
		t.Fatalf("expected ReadOnly on numeric 1, got %v", v)
	}
}

func TestBoundaries(t *testing.T) {
	var _ baseenumer.BoundedEnumer[openfiletype.Variant] = openfiletype.Variant(0)

	minVal, maxVal := openfiletype.Min(), openfiletype.Max()
	if minVal != openfiletype.Invalid || maxVal != openfiletype.ReadWriteOrCreateOnly {
		t.Fatalf("expected min %v, max %v", openfiletype.Invalid, openfiletype.ReadWriteOrCreateOnly)
	}

	if minVal.Min() != minVal || maxVal.Max() != maxVal {
		t.Fatalf("receiver Min/Max mismatch")
	}
}

func TestBoundaryPredicates(t *testing.T) {
	minVal, maxVal := openfiletype.Min(), openfiletype.Max()
	if !minVal.IsMin() || !maxVal.IsMax() {
		t.Fatalf("boundary predicates failed")
	}

	if maxVal.IsMin() || minVal.IsMax() {
		t.Fatalf("inverse boundary predicates failed")
	}

	if !minVal.IsInRange(minVal, maxVal) || !openfiletype.ReadOnly.IsInRange(minVal, maxVal) {
		t.Fatalf("IsInRange failed for valid range")
	}

	if openfiletype.Variant(99).IsInRange(minVal, maxVal) {
		t.Fatalf("IsInRange succeeded for out-of-range value")
	}
}
