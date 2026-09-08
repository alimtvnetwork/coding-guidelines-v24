package processstatetype_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/errtype/processstatetype"
)

func TestProcessStateType_Lifecycle(t *testing.T) {
	state := processstatetype.Running
	if state.Name() != "Running" {
		t.Fatalf("expected Name 'Running', got %s", state.Name())
	}

	if state.String() != "Running" || state.ValueString() != "Running" {
		t.Fatalf("unexpected string/value string for Running")
	}

	if state.Value() != "Running" {
		t.Fatalf("expected Value 'Running', got %s", state.Value())
	}

	if !state.IsValid() || !state.IsEnum() {
		t.Fatal("expected Running to be valid and enum")
	}

	if !state.IsCompare(processstatetype.Running) {
		t.Fatal("expected IsCompare(Running) to be true")
	}

	if state.IsCompare(processstatetype.Pending) {
		t.Fatal("expected IsCompare(Pending) to be false")
	}
}

func TestProcessStateType_JSON(t *testing.T) {
	state := processstatetype.Running
	bytes, err := json.Marshal(state)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	if string(bytes) != `"Running"` {
		t.Fatalf("expected '\"Running\"', got %s", string(bytes))
	}

	var unmarshaled processstatetype.Variant
	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if unmarshaled != processstatetype.Running {
		t.Fatalf("expected unmarshaled == Running, got %v", unmarshaled)
	}

	parsed := processstatetype.Parse("running")
	if parsed != processstatetype.Running {
		t.Fatalf("expected Parse('running') to match, got %v", parsed)
	}

	if unknown := processstatetype.Parse("non-existent"); unknown != processstatetype.Unknown {
		t.Fatalf("expected Unknown for unknown string, got %v", unknown)
	}
}

func TestProcessStateType_UnmarshalNull(t *testing.T) {
	var s processstatetype.Variant
	if err := json.Unmarshal([]byte(`null`), &s); err != nil {
		t.Fatalf("expected nil error on null, got %v", err)
	}

	if s != processstatetype.Unknown {
		t.Fatalf("expected Unknown on null, got %v", s)
	}

	if err := json.Unmarshal([]byte(`"null"`), &s); err != nil {
		t.Fatalf("expected nil error on 'null', got %v", err)
	}

	if s != processstatetype.Unknown {
		t.Fatalf("expected Unknown on 'null', got %v", s)
	}
}

func TestProcessStateType_All(t *testing.T) {
	all := processstatetype.All()
	if len(all) != 5 {
		t.Fatalf("expected 5 states, got %d", len(all))
	}

	allAlias := processstatetype.AllProcessStates()
	if len(allAlias) != 5 {
		t.Fatalf("expected 5 states from AllProcessStates, got %d", len(allAlias))
	}
}

func TestProcessStateType_Values(t *testing.T) {
	vals := processstatetype.Values()
	if len(vals) != 5 {
		t.Fatalf("expected 5 values, got %d", len(vals))
	}
}
