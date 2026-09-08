package baseenumer_test

import (
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
)

type sampleByteEnum byte

const (
	sampleByteInvalid sampleByteEnum = iota
	sampleByteActive
	sampleByteInactive
)

var sampleByteLabels = [...]string{
	sampleByteInvalid:  "Unknown",
	sampleByteActive:   "Active",
	sampleByteInactive: "Inactive",
}

type sampleStringEnum string

const (
	sampleStringUnknown sampleStringEnum = ""
	sampleStringAlpha   sampleStringEnum = "Alpha"
	sampleStringBeta    sampleStringEnum = "Beta"
)

var sampleStringVariants = []sampleStringEnum{
	sampleStringUnknown,
	sampleStringAlpha,
	sampleStringBeta,
}

func TestBasicIntegerEnum_Operations(t *testing.T) {
	b := baseenumer.NewBasicInteger(sampleByteLabels[:], sampleByteInvalid)

	all := b.All()
	if len(all) != 2 {
		t.Fatalf("expected 2 items in All(), got %d", len(all))
	}

	vals := b.Values()
	if len(vals) != 2 {
		t.Fatalf("expected 2 values, got %d", len(vals))
	}

	val, err := b.Parse("Active")
	if err != nil || val != sampleByteActive {
		t.Fatalf("expected Active, got %v, err: %v", val, err)
	}

	_, err = b.Parse("")
	if err == nil {
		t.Fatal("expected error on empty parse")
	}

	_, err = b.Parse("NonExistent")
	if err == nil {
		t.Fatal("expected error on unknown parse")
	}
}

func TestBasicIntegerEnum_UnmarshalJSON(t *testing.T) {
	b := baseenumer.NewBasicInteger(sampleByteLabels[:], sampleByteInvalid)
	var target sampleByteEnum

	err := b.UnmarshalJSON([]byte(`"Inactive"`), &target)
	if err != nil || target != sampleByteInactive {
		t.Fatalf("failed unmarshaling string: %v, got %v", err, target)
	}

	err = b.UnmarshalJSON([]byte(`1`), &target)
	if err != nil || target != sampleByteActive {
		t.Fatalf("failed unmarshaling numeric: %v, got %v", err, target)
	}

	err = b.UnmarshalJSON([]byte(`null`), &target)
	if err != nil || target != sampleByteInvalid {
		t.Fatalf("failed unmarshaling null: %v, got %v", err, target)
	}
}

func TestBasicStringEnum_Operations(t *testing.T) {
	b := baseenumer.NewBasicString(sampleStringVariants, sampleStringUnknown)

	all := b.All()
	if len(all) != 3 {
		t.Fatalf("expected 3 items in All(), got %d", len(all))
	}

	vals := b.Values()
	if len(vals) != 2 {
		t.Fatalf("expected 2 values, got %d", len(vals))
	}

	val, err := b.Parse("Alpha")
	if err != nil || val != sampleStringAlpha {
		t.Fatalf("expected Alpha, got %v, err: %v", val, err)
	}

	_, err = b.Parse("")
	if err == nil {
		t.Fatal("expected error on empty parse")
	}

	_, err = b.Parse("NonExistent")
	if err == nil {
		t.Fatal("expected error on unknown parse")
	}
}

func TestBasicStringEnum_UnmarshalJSON(t *testing.T) {
	b := baseenumer.NewBasicString(sampleStringVariants, sampleStringUnknown)
	var target sampleStringEnum

	err := b.UnmarshalJSON([]byte(`"Beta"`), &target)
	if err != nil || target != sampleStringBeta {
		t.Fatalf("failed unmarshaling string: %v, got %v", err, target)
	}

	err = b.UnmarshalJSON([]byte(`null`), &target)
	if err != nil || target != sampleStringUnknown {
		t.Fatalf("failed unmarshaling null: %v, got %v", err, target)
	}
}
