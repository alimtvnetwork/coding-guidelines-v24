package appfault_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/appfault"
)

func TestSeverityTypeEnumAndJSON(t *testing.T) {
	sev := appfault.SeverityError
	data, err := json.Marshal(sev)
	if err != nil || string(data) != "\"Error\"" || sev.Name() != "Error" {
		t.Fatalf("expected \"Error\" JSON, got %s", string(data))
	}

	var parsed appfault.SeverityType
	if err := json.Unmarshal([]byte("\"Critical\""), &parsed); err != nil || parsed != appfault.SeverityCritical {
		t.Fatalf("expected SeverityCritical, got %v", parsed)
	}
}

func TestSeverityType_UnmarshalJSON_Invalid(t *testing.T) {
	var sev appfault.SeverityType

	if err := json.Unmarshal([]byte(`"NonExistentSeverity"`), &sev); err == nil {
		t.Fatalf("expected error unmarshaling invalid severity string, got nil")
	}

	if err := json.Unmarshal([]byte(`99`), &sev); err == nil {
		t.Fatalf("expected error unmarshaling invalid numeric severity 99, got nil")
	}

	if err := json.Unmarshal([]byte(`null`), &sev); err != nil {
		t.Fatalf("expected nil error on null, got %v", err)
	}

	if sev != appfault.SeverityUnknown {
		t.Fatalf("expected SeverityUnknown on null, got %v", sev)
	}

	if err := json.Unmarshal([]byte(`2`), &sev); err != nil {
		t.Fatalf("expected nil error on numeric 2, got %v", err)
	}

	if sev != appfault.SeverityWarn {
		t.Fatalf("expected SeverityWarn on numeric 2, got %v", sev)
	}
}
