package baseenumer_test

import (
	"encoding/json"
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
)

type (
	testByteEnum   byte
	testStringEnum string
)

const (
	testByteInvalid testByteEnum = iota
	testByteActive
	testByteInactive
)

const (
	testStringUnknown testStringEnum = ""
	testStringPending testStringEnum = "Pending"
	testStringRunning testStringEnum = "Running"
)

var (
	testByteLabels = []string{"Invalid", "Active", "Inactive"}
	testByteMap    = baseenumer.CompileMap(testByteLabels, testByteInvalid)

	testStringMap = map[string]testStringEnum{
		"pending": testStringPending,
		"running": testStringRunning,
	}
)

func TestMarshalJSON(t *testing.T) {
	bytes, err := baseenumer.MarshalJSON("Active")
	if err != nil || string(bytes) != `"Active"` {
		t.Fatalf("MarshalJSON failed: %v, got %s", err, string(bytes))
	}
}

func TestUnmarshalIntegerJSON_ValidString(t *testing.T) {
	var target testByteEnum
	err := baseenumer.UnmarshalIntegerJSON([]byte(`"Active"`), &target, testByteMap, 2, testByteInvalid)
	if err != nil || target != testByteActive {
		t.Fatalf("failed to unmarshal string: %v, got %v", err, target)
	}

	err = baseenumer.UnmarshalIntegerJSON([]byte(`"inactive"`), &target, testByteMap, 2, testByteInvalid)
	if err != nil || target != testByteInactive {
		t.Fatalf("failed to unmarshal lower string: %v, got %v", err, target)
	}
}

func TestUnmarshalIntegerJSON_Numeric(t *testing.T) {
	var target testByteEnum
	err := baseenumer.UnmarshalIntegerJSON([]byte(`1`), &target, testByteMap, 2, testByteInvalid)
	if err != nil || target != testByteActive {
		t.Fatalf("failed to unmarshal numeric literal: %v, got %v", err, target)
	}

	err = baseenumer.UnmarshalIntegerJSON([]byte(`"2"`), &target, testByteMap, 2, testByteInvalid)
	if err != nil || target != testByteInactive {
		t.Fatalf("failed to unmarshal numeric string: %v, got %v", err, target)
	}
}

func TestUnmarshalIntegerJSON_NullAndEmpty(t *testing.T) {
	var target testByteEnum = testByteActive
	err := baseenumer.UnmarshalIntegerJSON([]byte(`null`), &target, testByteMap, 2, testByteInvalid)
	if err != nil || target != testByteInvalid {
		t.Fatalf("failed on null: %v, got %v", err, target)
	}

	target = testByteActive
	err = baseenumer.UnmarshalIntegerJSON([]byte(`"null"`), &target, testByteMap, 2, testByteInvalid)
	if err != nil || target != testByteInvalid {
		t.Fatalf("failed on 'null': %v, got %v", err, target)
	}
}

func TestUnmarshalIntegerJSON_OutOfRange(t *testing.T) {
	var target testByteEnum
	err := baseenumer.UnmarshalIntegerJSON([]byte(`99`), &target, testByteMap, 2, testByteInvalid)
	if err == nil {
		t.Fatalf("expected error on out of range numeric 99, got nil")
	}

	err = baseenumer.UnmarshalIntegerJSON([]byte(`"UnknownVariant"`), &target, testByteMap, 2, testByteInvalid)
	if err == nil {
		t.Fatalf("expected error on unknown string, got nil")
	}
}

func TestUnmarshalStringJSON_Valid(t *testing.T) {
	var target testStringEnum
	err := baseenumer.UnmarshalStringJSON([]byte(`"Pending"`), &target, testStringMap, testStringUnknown)
	if err != nil || target != testStringPending {
		t.Fatalf("failed to unmarshal valid string: %v, got %v", err, target)
	}

	err = baseenumer.UnmarshalStringJSON([]byte(`"running"`), &target, testStringMap, testStringUnknown)
	if err != nil || target != testStringRunning {
		t.Fatalf("failed to unmarshal lowercase string: %v, got %v", err, target)
	}
}

func TestUnmarshalStringJSON_NullAndErrors(t *testing.T) {
	var target testStringEnum = testStringPending
	err := baseenumer.UnmarshalStringJSON([]byte(`null`), &target, testStringMap, testStringUnknown)
	if err != nil || target != testStringUnknown {
		t.Fatalf("failed on null: %v, got %v", err, target)
	}

	err = baseenumer.UnmarshalStringJSON([]byte(`"NonExistent"`), &target, testStringMap, testStringUnknown)
	if err == nil {
		t.Fatalf("expected error on non-existent string variant, got nil")
	}
}

func TestResolveTypeName(t *testing.T) {
	var strTarget testStringEnum
	name := baseenumer.ResolveTypeName(&strTarget)
	if name == "" {
		t.Fatal("expected non-empty resolved type name")
	}

	var nilPtr *testStringEnum
	nilName := baseenumer.ResolveTypeName(nilPtr)
	if nilName != "Variant" {
		t.Fatalf("expected Variant for nil target, got %s", nilName)
	}
}

func TestUnmarshalJSON_InterfaceConformance(t *testing.T) {
	var _ json.Marshaler = testByteInvalid
	var _ json.Unmarshaler = (*testByteEnum)(nil)
}

func (b testByteEnum) MarshalJSON() ([]byte, error) {
	return baseenumer.MarshalJSON(testByteLabels[b])
}

func (b *testByteEnum) UnmarshalJSON(data []byte) error {
	return baseenumer.UnmarshalIntegerJSON(data, b, testByteMap, 2, testByteInvalid)
}
