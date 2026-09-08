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

type sampleSparseEnum int

const (
	sampleSparseZero  sampleSparseEnum = 0
	sampleSparseTen   sampleSparseEnum = 10
	sampleSparseFifty sampleSparseEnum = 50
)

var (
	sampleSparseVariants = []sampleSparseEnum{sampleSparseZero, sampleSparseTen, sampleSparseFifty}
	sampleSparseNames    = []string{"Zero", "Ten", "Fifty"}
)

func TestBasicIntegerEnum_AllAndValues(t *testing.T) {
	b := baseenumer.NewBasicInteger(sampleByteLabels[:], sampleByteInvalid)
	if all := b.All(); len(all) != 2 {
		t.Fatalf("expected 2 items in All(), got %d", len(all))
	}

	if vals := b.Values(); len(vals) != 2 {
		t.Fatalf("expected 2 values, got %d", len(vals))
	}
}

func TestBasicIntegerEnum_Parse(t *testing.T) {
	b := baseenumer.NewBasicInteger(sampleByteLabels[:], sampleByteInvalid)
	if val, err := b.Parse("Active"); err != nil || val != sampleByteActive {
		t.Fatalf("expected Active, got %v, err: %v", val, err)
	}

	if _, err := b.Parse(""); err == nil {
		t.Fatal("expected error on empty parse")
	}

	if _, err := b.Parse("NonExistent"); err == nil {
		t.Fatal("expected error on unknown parse")
	}
}

func TestBasicIntegerEnum_Accessors(t *testing.T) {
	b := baseenumer.NewBasicInteger(sampleByteLabels[:], sampleByteInvalid)
	if b.Zero() != sampleByteInvalid {
		t.Fatalf("expected Zero %v, got %v", sampleByteInvalid, b.Zero())
	}

	if b.MaxValid() != 2 {
		t.Fatalf("expected MaxValid 2, got %d", b.MaxValid())
	}

	if b.TypeName() != "baseenumer_test" {
		t.Fatalf("expected TypeName baseenumer_test, got %s", b.TypeName())
	}

	if len(b.Map()) == 0 {
		t.Fatal("expected non-empty variant map")
	}
}

func TestBasicIntegerEnum_ParseLookup(t *testing.T) {
	b := baseenumer.NewBasicInteger(sampleByteLabels[:], sampleByteInvalid)
	v, trimmed, isOk := b.ParseLookup("  Active ")
	if !isOk || v != sampleByteActive || trimmed != "Active" {
		t.Fatalf("unexpected ParseLookup result: %v, %s, %v", v, trimmed, isOk)
	}

	if _, _, isOk = b.ParseLookup(""); isOk {
		t.Fatal("expected isOk=false for empty string")
	}

	if _, _, isOk = b.ParseLookup("non-existent"); isOk {
		t.Fatal("expected isOk=false for non-existent")
	}
}

func TestBasicIntegerEnum_UnmarshalJSON(t *testing.T) {
	b := baseenumer.NewBasicInteger(sampleByteLabels[:], sampleByteInvalid)
	var target sampleByteEnum
	if err := b.UnmarshalJSON([]byte(`"Inactive"`), &target); err != nil || target != sampleByteInactive {
		t.Fatalf("failed unmarshaling string: %v, got %v", err, target)
	}

	if err := b.UnmarshalJSON([]byte(`1`), &target); err != nil || target != sampleByteActive {
		t.Fatalf("failed unmarshaling numeric: %v, got %v", err, target)
	}

	if err := b.UnmarshalJSON([]byte(`null`), &target); err != nil || target != sampleByteInvalid {
		t.Fatalf("failed unmarshaling null: %v, got %v", err, target)
	}
}

func TestBasicStringEnum_AllAndValues(t *testing.T) {
	b := baseenumer.NewBasicString(sampleStringVariants, sampleStringUnknown)
	if all := b.All(); len(all) != 3 {
		t.Fatalf("expected 3 items in All(), got %d", len(all))
	}

	if vals := b.Values(); len(vals) != 2 {
		t.Fatalf("expected 2 values, got %d", len(vals))
	}
}

func TestBasicStringEnum_Parse(t *testing.T) {
	b := baseenumer.NewBasicString(sampleStringVariants, sampleStringUnknown)
	if val, err := b.Parse("Alpha"); err != nil || val != sampleStringAlpha {
		t.Fatalf("expected Alpha, got %v, err: %v", val, err)
	}

	if _, err := b.Parse(""); err == nil {
		t.Fatal("expected error on empty parse")
	}

	if _, err := b.Parse("NonExistent"); err == nil {
		t.Fatal("expected error on unknown parse")
	}
}

func TestBasicStringEnum_Accessors(t *testing.T) {
	b := baseenumer.NewBasicString(sampleStringVariants, sampleStringUnknown)
	if b.Zero() != sampleStringUnknown {
		t.Fatalf("expected Zero %v, got %v", sampleStringUnknown, b.Zero())
	}

	if b.TypeName() != "baseenumer_test" {
		t.Fatalf("expected TypeName baseenumer_test, got %s", b.TypeName())
	}

	if len(b.Map()) == 0 {
		t.Fatal("expected non-empty variant map")
	}
}

func TestBasicStringEnum_ParseLookup(t *testing.T) {
	b := baseenumer.NewBasicString(sampleStringVariants, sampleStringUnknown)
	v, trimmed, isOk := b.ParseLookup("  Alpha ")
	if !isOk || v != sampleStringAlpha || trimmed != "Alpha" {
		t.Fatalf("unexpected ParseLookup result: %v, %s, %v", v, trimmed, isOk)
	}

	if _, _, isOk = b.ParseLookup(""); isOk {
		t.Fatal("expected isOk=false for empty string")
	}

	if _, _, isOk = b.ParseLookup("non-existent"); isOk {
		t.Fatal("expected isOk=false for non-existent")
	}
}

func TestBasicStringEnum_UnmarshalJSON(t *testing.T) {
	b := baseenumer.NewBasicString(sampleStringVariants, sampleStringUnknown)
	var target sampleStringEnum
	if err := b.UnmarshalJSON([]byte(`"Beta"`), &target); err != nil || target != sampleStringBeta {
		t.Fatalf("failed unmarshaling string: %v, got %v", err, target)
	}

	if err := b.UnmarshalJSON([]byte(`null`), &target); err != nil || target != sampleStringUnknown {
		t.Fatalf("failed unmarshaling null: %v, got %v", err, target)
	}
}

func TestBasicSparseIntegerEnum_Accessors(t *testing.T) {
	b := baseenumer.NewBasicSparseInteger(sampleSparseVariants, sampleSparseNames, sampleSparseZero, 50)
	if b.Zero() != sampleSparseZero {
		t.Fatalf("expected Zero %v, got %v", sampleSparseZero, b.Zero())
	}

	if b.MaxValid() != 50 {
		t.Fatalf("expected MaxValid 50, got %d", b.MaxValid())
	}

	if b.TypeName() != "baseenumer_test" {
		t.Fatalf("expected TypeName baseenumer_test, got %s", b.TypeName())
	}

	if len(b.Map()) == 0 {
		t.Fatal("expected non-empty variant map")
	}
}

func TestBasicSparseIntegerEnum_AllAndValues(t *testing.T) {
	b := baseenumer.NewBasicSparseInteger(sampleSparseVariants, sampleSparseNames, sampleSparseZero, 50)
	all := b.All()
	if len(all) != 3 || all[1] != sampleSparseTen {
		t.Fatalf("unexpected All() result: %v", all)
	}

	vals := b.Values()
	if len(vals) != 3 || vals[1] != "Ten" {
		t.Fatalf("unexpected Values() result: %v", vals)
	}
}

func TestBasicSparseIntegerEnum_Parse(t *testing.T) {
	b := baseenumer.NewBasicSparseInteger(sampleSparseVariants, sampleSparseNames, sampleSparseZero, 50)
	if val, err := b.Parse("Ten"); err != nil || val != sampleSparseTen {
		t.Fatalf("expected Ten, got %v, err: %v", val, err)
	}

	if val, err := b.Parse("50"); err != nil || val != sampleSparseFifty {
		t.Fatalf("expected Fifty from numeric string, got %v, err: %v", val, err)
	}

	if _, err := b.Parse(""); err == nil {
		t.Fatal("expected error on empty parse")
	}

	if _, err := b.Parse("UnknownVariant"); err == nil {
		t.Fatal("expected error on unknown parse")
	}
}

func TestBasicSparseIntegerEnum_ParseLookup(t *testing.T) {
	b := baseenumer.NewBasicSparseInteger(sampleSparseVariants, sampleSparseNames, sampleSparseZero, 50)
	v, trimmed, isOk := b.ParseLookup("  Ten ")
	if !isOk || v != sampleSparseTen || trimmed != "Ten" {
		t.Fatalf("unexpected ParseLookup result: %v, %s, %v", v, trimmed, isOk)
	}

	if v, _, isOk = b.ParseLookup("min"); !isOk || v != sampleSparseZero {
		t.Fatalf("expected min alias to map to Zero, got %v", v)
	}

	if _, _, isOk = b.ParseLookup(""); isOk {
		t.Fatal("expected isOk=false for empty string")
	}
}

func TestBasicSparseIntegerEnum_UnmarshalJSON(t *testing.T) {
	b := baseenumer.NewBasicSparseInteger(sampleSparseVariants, sampleSparseNames, sampleSparseZero, 50)
	var target sampleSparseEnum
	if err := b.UnmarshalJSON([]byte(`"Fifty"`), &target); err != nil || target != sampleSparseFifty {
		t.Fatalf("unmarshal Fifty failed: %v, got %v", err, target)
	}

	if err := b.UnmarshalJSON([]byte(`10`), &target); err != nil || target != sampleSparseTen {
		t.Fatalf("unmarshal numeric 10 failed: %v, got %v", err, target)
	}

	if err := b.UnmarshalJSON([]byte(`null`), &target); err != nil || target != sampleSparseZero {
		t.Fatalf("unmarshal null failed: %v, got %v", err, target)
	}

	if err := b.UnmarshalJSON([]byte(`999`), &target); err == nil {
		t.Fatal("expected error on out-of-range numeric unmarshal")
	}
}

func TestCompileSparseIntegerMap_MismatchedLengths(t *testing.T) {
	m := baseenumer.CompileSparseIntegerMap(sampleSparseVariants, sampleSparseNames[:2], sampleSparseZero)
	if len(m) == 0 {
		t.Fatal("expected non-empty map")
	}
}
