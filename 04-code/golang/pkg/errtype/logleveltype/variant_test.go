package logleveltype_test

import (
	"encoding/json"
	"testing"

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
