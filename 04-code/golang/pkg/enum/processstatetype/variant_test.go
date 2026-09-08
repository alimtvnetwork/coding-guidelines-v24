package processstatetype_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/enum/processstatetype"
)

func TestInvalidZeroValue(t *testing.T) {
	var zero processstatetype.Variant
	if zero != processstatetype.Invalid {
		t.Fatalf("expected zero value to be Invalid, got %v", zero)
	}

	if zero != processstatetype.Unknown {
		t.Fatalf("expected Invalid to equal Unknown")
	}

	if zero.IsValid() {
		t.Fatalf("expected Invalid to not be valid")
	}

	if !zero.IsInvalid() {
		t.Fatalf("expected Invalid to be IsInvalid")
	}

	if zero.String() != "Unknown" {
		t.Fatalf("expected String 'Unknown', got %s", zero.String())
	}
}

func TestVariantsAndLabels(t *testing.T) {
	tests := []struct {
		variant      processstatetype.Variant
		expectedName string
	}{
		{processstatetype.Pending, "Pending"},
		{processstatetype.Running, "Running"},
		{processstatetype.Completed, "Completed"},
		{processstatetype.Failed, "Failed"},
		{processstatetype.Cancelled, "Cancelled"},
	}

	for _, tc := range tests {
		if tc.variant.Name() != tc.expectedName {
			t.Errorf("expected name %s, got %s", tc.expectedName, tc.variant.Name())
		}

		if tc.variant.Label() != tc.expectedName {
			t.Errorf("expected label %s, got %s", tc.expectedName, tc.variant.Label())
		}

		if !tc.variant.IsValid() {
			t.Errorf("expected %s to be valid", tc.expectedName)
		}
	}
}

func TestCheckers(t *testing.T) {
	p := processstatetype.Pending
	if !p.IsPending() {
		t.Fatalf("expected IsPending true")
	}

	r := processstatetype.Running
	if !r.IsRunning() {
		t.Fatalf("expected IsRunning true")
	}

	c := processstatetype.Completed
	if !c.IsCompleted() {
		t.Fatalf("expected IsCompleted true")
	}

	f := processstatetype.Failed
	if !f.IsFailed() {
		t.Fatalf("expected IsFailed true")
	}

	cn := processstatetype.Cancelled
	if !cn.IsCancelled() {
		t.Fatalf("expected IsCancelled true")
	}
}

func TestAllAndValues(t *testing.T) {
	all := processstatetype.All()
	if len(all) != 5 {
		t.Fatalf("expected 5 variants, got %d", len(all))
	}

	vals := processstatetype.Values()
	if len(vals) != 5 {
		t.Fatalf("expected 5 values, got %d", len(vals))
	}

	if vals[0] != "Pending" || vals[4] != "Cancelled" {
		t.Fatalf("unexpected values ordering: %v", vals)
	}
}

func TestParse_Success(t *testing.T) {
	cases := []struct {
		input    string
		expected processstatetype.Variant
	}{
		{"Pending", processstatetype.Pending},
		{"pending", processstatetype.Pending},
		{"PENDING", processstatetype.Pending},
		{"Running", processstatetype.Running},
		{"running", processstatetype.Running},
		{"Completed", processstatetype.Completed},
		{"Failed", processstatetype.Failed},
		{"Cancelled", processstatetype.Cancelled},
		{"1", processstatetype.Pending},
		{"2", processstatetype.Running},
	}

	for _, tc := range cases {
		v, isOk := processstatetype.Parse(tc.input)
		if !isOk {
			t.Fatalf("expected parse %q to succeed", tc.input)
		}

		if v != tc.expected {
			t.Fatalf("expected %v, got %v", tc.expected, v)
		}
	}
}

func TestParse_Failure(t *testing.T) {
	if _, isOk := processstatetype.Parse(""); isOk {
		t.Fatalf("expected empty string to fail")
	}

	if _, isOk := processstatetype.Parse("invalid_variant"); isOk {
		t.Fatalf("expected invalid variant to fail")
	}

	if v := processstatetype.ParseOrInvalid("invalid_variant"); v != processstatetype.Invalid {
		t.Fatalf("expected Invalid fallback, got %v", v)
	}

	if v := processstatetype.ParseOrUnknown("invalid_variant"); v != processstatetype.Unknown {
		t.Fatalf("expected Unknown fallback, got %v", v)
	}
}

func TestJSON_MarshalUnmarshal(t *testing.T) {
	orig := processstatetype.Running
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	if string(data) != "\"Running\"" {
		t.Fatalf("expected \"Running\", got %s", string(data))
	}

	var parsed processstatetype.Variant
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if parsed != orig {
		t.Fatalf("expected %v, got %v", orig, parsed)
	}
}

func TestJSON_UnmarshalNullAndNumeric(t *testing.T) {
	var v processstatetype.Variant
	if err := json.Unmarshal([]byte("null"), &v); err != nil {
		t.Fatalf("unmarshal null error: %v", err)
	}

	if v != processstatetype.Invalid {
		t.Fatalf("expected Invalid on null, got %v", v)
	}

	if err := json.Unmarshal([]byte("2"), &v); err != nil {
		t.Fatalf("unmarshal numeric 2 error: %v", err)
	}

	if v != processstatetype.Running {
		t.Fatalf("expected Running on numeric 2, got %v", v)
	}
}

func TestBoundaries(t *testing.T) {
	var _ baseenumer.BoundedEnumer[processstatetype.Variant] = processstatetype.Variant(0)

	minVal, maxVal := processstatetype.Min(), processstatetype.Max()
	if minVal != processstatetype.Invalid || maxVal != processstatetype.Cancelled {
		t.Fatalf("expected min %v, max %v", processstatetype.Invalid, processstatetype.Cancelled)
	}

	if minVal.Min() != minVal || maxVal.Max() != maxVal {
		t.Fatalf("receiver Min/Max mismatch")
	}
}

func TestBoundaryPredicates(t *testing.T) {
	minVal, maxVal := processstatetype.Min(), processstatetype.Max()
	if !minVal.IsMin() || !maxVal.IsMax() {
		t.Fatalf("boundary predicates failed")
	}

	if maxVal.IsMin() || minVal.IsMax() {
		t.Fatalf("inverse boundary predicates failed")
	}

	if !minVal.IsInRange(minVal, maxVal) || !processstatetype.Pending.IsInRange(minVal, maxVal) {
		t.Fatalf("IsInRange failed for valid range")
	}

	if processstatetype.Variant(99).IsInRange(minVal, maxVal) {
		t.Fatalf("IsInRange succeeded for out-of-range value")
	}
}
