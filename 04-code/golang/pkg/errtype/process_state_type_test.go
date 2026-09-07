package errtype_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/errtype"
)

func TestProcessStateType_Lifecycle(t *testing.T) {
	state := errtype.ProcessStateRunning

	if state.Name() != "Running" {
		t.Fatalf("expected Name() == 'Running', got %s", state.Name())
	}

	if state.String() != "Running" || state.Value() != "Running" || state.ValueString() != "Running" {
		t.Fatalf("unexpected string/value getters for ProcessStateRunning")
	}

	if !state.IsValid() {
		t.Fatal("expected Running to be valid")
	}

	if !state.IsEnum() {
		t.Fatal("expected Running to be registered enum")
	}

	if !state.IsCompare(errtype.ProcessStateRunning) {
		t.Fatal("expected IsCompare true for self")
	}

	if state.IsCompare(errtype.ProcessStatePending) {
		t.Fatal("expected IsCompare false for different state")
	}

	// JSON roundtrip
	data, err := json.Marshal(state)
	if err != nil {
		t.Fatalf("json marshal failed: %v", err)
	}

	var unmarshaled errtype.ProcessStateType
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("json unmarshal failed: %v", err)
	}

	if unmarshaled != errtype.ProcessStateRunning {
		t.Fatalf("expected unmarshaled == Running, got %s", unmarshaled)
	}

	// Case-insensitive parsing
	parsed := errtype.ParseProcessState("running")
	if parsed != errtype.ProcessStateRunning {
		t.Fatalf("expected ParseProcessState('running') to match, got %s", parsed)
	}

	if unknown := errtype.ParseProcessState("non-existent"); unknown != errtype.ProcessStateUnknown {
		t.Fatalf("expected unknown state, got %s", unknown)
	}
}

func TestProcessStateType_UnmarshalJSON_Invalid(t *testing.T) {
	var s errtype.ProcessStateType

	if err := json.Unmarshal([]byte(`"invalid_state"`), &s); err == nil {
		t.Fatalf("expected error unmarshaling invalid process state, got nil")
	}

	if err := json.Unmarshal([]byte(`null`), &s); err != nil {
		t.Fatalf("expected nil error on null, got %v", err)
	}

	if s != errtype.ProcessStateUnknown {
		t.Fatalf("expected ProcessStateUnknown on null, got %v", s)
	}
}
