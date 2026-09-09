package processstatetype_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/errtype/processstatetype"
)

func TestProcessStateType_Interfaces(t *testing.T) {
	var (
		_ baseenumer.BaseEnumer                              = processstatetype.Running
		_ baseenumer.StringEnumer                            = processstatetype.Running
		_ baseenumer.MinMaxer[processstatetype.Variant]      = processstatetype.Running
		_ baseenumer.BoundedEnumer[processstatetype.Variant] = processstatetype.Running
		_ baseenumer.Bounder[processstatetype.Variant]       = processstatetype.Running
		_ json.Marshaler                                     = processstatetype.Running
		_ json.Unmarshaler                                   = (*processstatetype.Variant)(nil)
	)
}

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

func TestProcessStateType_Boundary(t *testing.T) {
	if processstatetype.Min() != processstatetype.Pending {
		t.Fatalf("expected Min to be Pending")
	}

	if processstatetype.Max() != processstatetype.Cancelled {
		t.Fatalf("expected Max to be Cancelled")
	}

	if !processstatetype.Pending.IsMin() {
		t.Fatalf("expected Pending.IsMin() to be true")
	}

	if processstatetype.Running.IsMin() {
		t.Fatalf("expected Running.IsMin() to be false")
	}
}

func TestProcessStateType_BoundaryMax(t *testing.T) {
	if !processstatetype.Cancelled.IsMax() {
		t.Fatalf("expected Cancelled.IsMax() to be true")
	}

	if processstatetype.Failed.IsMax() {
		t.Fatalf("expected Failed.IsMax() to be false")
	}

	if processstatetype.Running.Min() != processstatetype.Pending {
		t.Fatalf("expected Running.Min() to be Pending")
	}

	if processstatetype.Running.Max() != processstatetype.Cancelled {
		t.Fatalf("expected Running.Max() to be Cancelled")
	}
}

func TestProcessStateType_IsInRange(t *testing.T) {
	if !processstatetype.Completed.IsInRange(processstatetype.Cancelled, processstatetype.Running) {
		t.Fatalf("expected Completed to be in range [Cancelled, Running]")
	}

	if processstatetype.Unknown.IsInRange(processstatetype.Cancelled, processstatetype.Running) {
		t.Fatalf("expected Unknown to not be in range [Cancelled, Running]")
	}

	if !processstatetype.Running.IsInRange(processstatetype.Pending, processstatetype.Running) {
		t.Fatalf("expected Running to be in range [Pending, Running]")
	}
}

func TestProcessStateType_ValueAndCollections(t *testing.T) {
	p := processstatetype.Pending
	if p.Value() != "Pending" {
		t.Fatalf("expected 'Pending' from Value(), got %s", p.Value())
	}

	all := processstatetype.All()
	if len(all) != 5 || len(p.All()) != 5 {
		t.Fatalf("expected 5 variants from All")
	}

	vals := processstatetype.Values()
	if len(vals) != 5 || len(p.Values()) != 5 {
		t.Fatalf("expected 5 values from Values")
	}
}
