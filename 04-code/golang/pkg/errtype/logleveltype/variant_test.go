package logleveltype_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/errtype/logleveltype"
)

func TestLogLevelType_Lifecycle(t *testing.T) {
	lvl := logleveltype.Warn
	if lvl.Name() != "Warn" {
		t.Fatalf("expected Name 'Warn', got %s", lvl.Name())
	}

	if lvl.String() != "Warn" || lvl.ValueString() != "3" {
		t.Fatalf("unexpected string/value string for Warn")
	}

	if lvl.Code() != 3 || lvl.Int() != 3 {
		t.Fatalf("unexpected code/int for Warn")
	}

	if !lvl.IsValid() || !lvl.IsEnum() {
		t.Fatal("expected Warn to be valid and enum")
	}

	if !lvl.IsCompare(logleveltype.Warn) {
		t.Fatal("expected IsCompare(Warn) to be true")
	}

	if lvl.IsCompare(logleveltype.Error) {
		t.Fatal("expected IsCompare(Error) to be false")
	}
}

func TestLogLevelType_JSON(t *testing.T) {
	lvl := logleveltype.Warn
	bytes, err := json.Marshal(lvl)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	if string(bytes) != `"Warn"` {
		t.Fatalf("expected '\"Warn\"', got %s", string(bytes))
	}

	var unmarshaled logleveltype.Variant
	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if unmarshaled != logleveltype.Warn {
		t.Fatalf("expected unmarshaled == Warn, got %v", unmarshaled)
	}

	parsed := logleveltype.Parse("warn")
	if parsed != logleveltype.Warn {
		t.Fatalf("expected Parse('warn') to match, got %v", parsed)
	}

	if unknown := logleveltype.Parse("non-existent"); unknown != 0 {
		t.Fatalf("expected 0 for unknown string, got %v", unknown)
	}
}

func TestLogLevelType_UnmarshalJSON_Invalid(t *testing.T) {
	var l logleveltype.Variant
	if err := json.Unmarshal([]byte(`"InvalidLevel"`), &l); err == nil {
		t.Fatal("expected error on unknown string level")
	}

	if err := json.Unmarshal([]byte(`999`), &l); err == nil {
		t.Fatal("expected error on invalid numeric level")
	}

	if err := json.Unmarshal([]byte(`null`), &l); err != nil || l != 0 {
		t.Fatalf("expected zero on null, got %v, err: %v", l, err)
	}

	if err := json.Unmarshal([]byte(`"null"`), &l); err != nil || l != 0 {
		t.Fatalf("expected zero on 'null', got %v, err: %v", l, err)
	}

	if err := json.Unmarshal([]byte(`2`), &l); err != nil {
		t.Fatalf("unexpected error unmarshaling numeric 2: %v", err)
	}

	if l != logleveltype.Info {
		t.Fatalf("expected Info on numeric 2, got %v", l)
	}
}

func TestLogLevelType_All(t *testing.T) {
	all := logleveltype.All()
	if len(all) != 5 {
		t.Fatalf("expected 5 levels, got %d", len(all))
	}

	allAlias := logleveltype.AllLogLevels()
	if len(allAlias) != 5 {
		t.Fatalf("expected 5 levels from AllLogLevels, got %d", len(allAlias))
	}
}

func TestLogLevelType_Values(t *testing.T) {
	vals := logleveltype.Values()
	if len(vals) != 5 {
		t.Fatalf("expected 5 values, got %d", len(vals))
	}
}

func TestLogLevelType_Boundaries(t *testing.T) {
	var _ baseenumer.BoundedEnumer[logleveltype.Variant] = logleveltype.Variant(0)

	minVal, maxVal := logleveltype.Min(), logleveltype.Max()
	if minVal != logleveltype.Variant(0) || maxVal != logleveltype.Fatal {
		t.Fatalf("expected min %v, max %v", logleveltype.Variant(0), logleveltype.Fatal)
	}

	if minVal.Min() != minVal || maxVal.Max() != maxVal {
		t.Fatalf("receiver Min/Max mismatch")
	}
}

func TestLogLevelType_BoundaryPredicates(t *testing.T) {
	minVal, maxVal := logleveltype.Min(), logleveltype.Max()
	if !minVal.IsMin() || !maxVal.IsMax() {
		t.Fatalf("boundary predicates failed")
	}

	if maxVal.IsMin() || minVal.IsMax() {
		t.Fatalf("inverse boundary predicates failed")
	}

	if !minVal.IsInRange(minVal, maxVal) || !logleveltype.Debug.IsInRange(minVal, maxVal) {
		t.Fatalf("IsInRange failed for valid range")
	}

	if logleveltype.Variant(99).IsInRange(minVal, maxVal) {
		t.Fatalf("IsInRange succeeded for out-of-range value")
	}
}

func TestLogLevelType_ValueAndCollections(t *testing.T) {
	d := logleveltype.Debug
	if d.Value() != 1 {
		t.Fatalf("expected 1 from Value(), got %d", d.Value())
	}

	all := logleveltype.All()
	if len(all) != 5 || len(d.All()) != 5 {
		t.Fatalf("expected 5 variants from All")
	}

	vals := logleveltype.Values()
	if len(vals) != 5 || len(d.Values()) != 5 {
		t.Fatalf("expected 5 values from Values")
	}
}
