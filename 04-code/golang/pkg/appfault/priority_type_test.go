package appfault_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/appfault"
)

func TestPriorityTypeEnumAndJSON(t *testing.T) {
	pri := appfault.PriorityHigh
	data, err := json.Marshal(pri)
	if err != nil || string(data) != "\"High\"" || pri.Name() != "High" {
		t.Fatalf("expected \"High\" JSON, got %s", string(data))
	}

	var parsed appfault.PriorityType
	if err := json.Unmarshal([]byte("\"Low\""), &parsed); err != nil || parsed != appfault.PriorityLow {
		t.Fatalf("expected PriorityLow, got %v", parsed)
	}
}

func TestPriorityType_UnmarshalJSON_Invalid(t *testing.T) {
	var pri appfault.PriorityType

	if err := json.Unmarshal([]byte(`"NonExistentPriority"`), &pri); err == nil {
		t.Fatalf("expected error unmarshaling invalid priority string, got nil")
	}

	if err := json.Unmarshal([]byte(`99`), &pri); err == nil {
		t.Fatalf("expected error unmarshaling invalid numeric priority 99, got nil")
	}

	if err := json.Unmarshal([]byte(`null`), &pri); err != nil {
		t.Fatalf("expected nil error on null, got %v", err)
	}

	if pri != appfault.PriorityUnknown {
		t.Fatalf("expected PriorityUnknown on null, got %v", pri)
	}

	if err := json.Unmarshal([]byte(`1`), &pri); err != nil {
		t.Fatalf("expected nil error on numeric 1, got %v", err)
	}

	if pri != appfault.PriorityLow {
		t.Fatalf("expected PriorityLow on numeric 1, got %v", pri)
	}
}
