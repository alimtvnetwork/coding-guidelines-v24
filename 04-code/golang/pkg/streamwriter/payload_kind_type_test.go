package streamwriter_test

import (
	"testing"

	"coding-guidelines/common/pkg/streamwriter"
)

func TestPayloadKind_Basics(t *testing.T) {
	pk := streamwriter.PayloadBytes
	if pk.Name() != "Bytes" || pk.String() != "Bytes" {
		t.Fatalf("expected Bytes, got %s", pk.Name())
	}

	if !pk.IsValid() {
		t.Fatalf("expected PayloadBytes to be valid")
	}

	kinds := []struct {
		kind         streamwriter.PayloadKind
		expectedName string
	}{
		{streamwriter.PayloadNil, "Nil"},
		{streamwriter.PayloadBytes, "Bytes"},
		{streamwriter.PayloadString, "String"},
		{streamwriter.PayloadError, "Error"},
		{streamwriter.PayloadMap, "Map"},
		{streamwriter.PayloadStruct, "Struct"},
		{streamwriter.PayloadPrimitive, "Primitive"},
	}

	for _, tc := range kinds {
		if tc.kind.Name() != tc.expectedName {
			t.Errorf("expected %s, got %s", tc.expectedName, tc.kind.Name())
		}
	}

	invalid := streamwriter.PayloadKind(99)
	if invalid.IsValid() {
		t.Fatalf("expected invalid payload kind to return IsValid false")
	}
}
