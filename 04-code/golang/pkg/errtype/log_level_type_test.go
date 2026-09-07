package errtype_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/errtype"
)

func TestLogLevelType_Lifecycle(t *testing.T) {
	lvl := errtype.LogLevelWarn

	if lvl.Name() != "Warn" {
		t.Fatalf("expected Name() == 'Warn', got %s", lvl.Name())
	}

	if lvl.String() != "Warn" || lvl.ValueString() != "3" {
		t.Fatalf("unexpected string/value string for LogLevelWarn")
	}

	if lvl.Code() != 3 || lvl.Int() != 3 {
		t.Fatalf("unexpected code/int: %d/%d", lvl.Code(), lvl.Int())
	}

	if !lvl.IsValid() || !lvl.IsEnum() {
		t.Fatal("expected LogLevelWarn to be valid and enum")
	}

	if !lvl.IsCompare(errtype.LogLevelWarn) {
		t.Fatal("expected IsCompare true for self")
	}

	if lvl.IsCompare(errtype.LogLevelError) {
		t.Fatal("expected IsCompare false for different level")
	}

	// JSON string marshal
	data, err := json.Marshal(lvl)
	if err != nil {
		t.Fatalf("json marshal failed: %v", err)
	}

	var unmarshaled errtype.LogLevelType
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("json unmarshal failed: %v", err)
	}

	if unmarshaled != errtype.LogLevelWarn {
		t.Fatalf("expected unmarshaled == LogLevelWarn, got %v", unmarshaled)
	}

	// Parsing
	parsed := errtype.ParseLogLevel("warn")
	if parsed != errtype.LogLevelWarn {
		t.Fatalf("expected ParseLogLevel('warn') to match, got %v", parsed)
	}

	if unknown := errtype.ParseLogLevel("non-existent"); unknown != 0 {
		t.Fatalf("expected 0 for non-existent log level, got %v", unknown)
	}
}

func TestLogLevelType_UnmarshalJSON_Invalid(t *testing.T) {
	var l errtype.LogLevelType

	if err := json.Unmarshal([]byte(`"invalid_level"`), &l); err == nil {
		t.Fatalf("expected error unmarshaling invalid log level, got nil")
	}

	if err := json.Unmarshal([]byte(`999`), &l); err == nil {
		t.Fatalf("expected error unmarshaling out-of-range numeric code 999, got nil")
	}

	if err := json.Unmarshal([]byte(`null`), &l); err != nil {
		t.Fatalf("expected nil error on null, got %v", err)
	}

	if err := json.Unmarshal([]byte(`2`), &l); err != nil {
		t.Fatalf("expected nil error on numeric 2, got %v", err)
	}

	if l != errtype.LogLevelInfo {
		t.Fatalf("expected LogLevelInfo on numeric 2, got %v", l)
	}
}
