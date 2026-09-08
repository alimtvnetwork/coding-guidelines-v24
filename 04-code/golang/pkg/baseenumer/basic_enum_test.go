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

func TestBasicIntegerEnum_MinMax(t *testing.T) {
	b := baseenumer.NewBasicInteger(sampleByteLabels[:], sampleByteInvalid)
	if b.Min() != sampleByteInvalid {
		t.Fatalf("expected Min %v, got %v", sampleByteInvalid, b.Min())
	}

	if b.Max() != sampleByteInactive {
		t.Fatalf("expected Max %v, got %v", sampleByteInactive, b.Max())
	}

	assertIntegerPredicates(t, b)
}

func assertIntegerPredicates(t *testing.T, b *baseenumer.BasicIntegerEnum[sampleByteEnum]) {
	if !b.IsMin(sampleByteInvalid) {
		t.Fatal("expected IsMin to be true for Min")
	}

	if b.IsMin(sampleByteActive) {
		t.Fatal("expected IsMin to be false for Active")
	}

	if !b.IsMax(sampleByteInactive) {
		t.Fatal("expected IsMax to be true for Max")
	}

	if b.IsMax(sampleByteActive) {
		t.Fatal("expected IsMax to be false for Active")
	}
}

func TestBasicIntegerEnum_IsInRange(t *testing.T) {
	b := baseenumer.NewBasicInteger(sampleByteLabels[:], sampleByteInvalid)
	if !b.IsInRange(sampleByteActive, b.Min(), b.Max()) {
		t.Fatal("expected Active to be in range [Min, Max]")
	}

	if b.IsInRange(sampleByteEnum(99), b.Min(), b.Max()) {
		t.Fatal("expected 99 to not be in range [Min, Max]")
	}

	if !b.IsInRange(sampleByteActive, sampleByteActive, sampleByteInactive) {
		t.Fatal("expected Active to be in range [Active, Inactive]")
	}

	if b.IsInRange(sampleByteInvalid, sampleByteActive, sampleByteInactive) {
		t.Fatal("expected Invalid to not be in range [Active, Inactive]")
	}
}

func TestBasicSparseIntegerEnum_MinMax(t *testing.T) {
	b := baseenumer.NewBasicSparseInteger(sampleSparseVariants, sampleSparseNames, sampleSparseZero, 50)
	if b.Min() != sampleSparseZero {
		t.Fatalf("expected Min %v, got %v", sampleSparseZero, b.Min())
	}

	if b.Max() != sampleSparseFifty {
		t.Fatalf("expected Max %v, got %v", sampleSparseFifty, b.Max())
	}

	assertSparsePredicates(t, b)
}

func assertSparsePredicates(t *testing.T, b *baseenumer.BasicIntegerEnum[sampleSparseEnum]) {
	if !b.IsMin(sampleSparseZero) {
		t.Fatal("expected IsMin to be true for Zero")
	}

	if b.IsMin(sampleSparseTen) {
		t.Fatal("expected IsMin to be false for Ten")
	}

	if !b.IsMax(sampleSparseFifty) {
		t.Fatal("expected IsMax to be true for Fifty")
	}

	if b.IsMax(sampleSparseTen) {
		t.Fatal("expected IsMax to be false for Ten")
	}
}

func TestBasicSparseIntegerEnum_IsInRange(t *testing.T) {
	b := baseenumer.NewBasicSparseInteger(sampleSparseVariants, sampleSparseNames, sampleSparseZero, 50)
	if !b.IsInRange(sampleSparseTen, b.Min(), b.Max()) {
		t.Fatal("expected Ten to be in range [Min, Max]")
	}

	if b.IsInRange(sampleSparseEnum(100), b.Min(), b.Max()) {
		t.Fatal("expected 100 to not be in range [Min, Max]")
	}

	if !b.IsInRange(sampleSparseTen, sampleSparseTen, sampleSparseFifty) {
		t.Fatal("expected Ten to be in range [Ten, Fifty]")
	}

	if b.IsInRange(sampleSparseZero, sampleSparseTen, sampleSparseFifty) {
		t.Fatal("expected Zero to not be in range [Ten, Fifty]")
	}
}

func TestBasicStringEnum_MinMax(t *testing.T) {
	b := baseenumer.NewBasicString(sampleStringVariants, sampleStringUnknown)
	if b.Min() != sampleStringAlpha {
		t.Fatalf("expected Min %v, got %v", sampleStringAlpha, b.Min())
	}

	if b.Max() != sampleStringBeta {
		t.Fatalf("expected Max %v, got %v", sampleStringBeta, b.Max())
	}

	assertStringPredicates(t, b)
}

func assertStringPredicates(t *testing.T, b *baseenumer.BasicStringEnum[sampleStringEnum]) {
	if !b.IsMin(sampleStringAlpha) {
		t.Fatal("expected IsMin to be true for Alpha")
	}

	if b.IsMin(sampleStringBeta) {
		t.Fatal("expected IsMin to be false for Beta")
	}

	if !b.IsMax(sampleStringBeta) {
		t.Fatal("expected IsMax to be true for Beta")
	}

	if b.IsMax(sampleStringAlpha) {
		t.Fatal("expected IsMax to be false for Alpha")
	}
}

func TestBasicStringEnum_IsInRange(t *testing.T) {
	b := baseenumer.NewBasicString(sampleStringVariants, sampleStringUnknown)
	if !b.IsInRange(sampleStringAlpha, b.Min(), b.Max()) {
		t.Fatal("expected Alpha to be in range [Min, Max]")
	}

	if b.IsInRange(sampleStringEnum("Gamma"), b.Min(), b.Max()) {
		t.Fatal("expected Gamma to not be in range [Min, Max]")
	}

	if !b.IsInRange(sampleStringBeta, sampleStringAlpha, sampleStringBeta) {
		t.Fatal("expected Beta to be in range [Alpha, Beta]")
	}

	if b.IsInRange(sampleStringEnum("000"), sampleStringAlpha, sampleStringBeta) {
		t.Fatal("expected 000 to not be in range [Alpha, Beta]")
	}
}

func TestBasicStringEnum_WithMinMax(t *testing.T) {
	b := baseenumer.NewBasicString(sampleStringVariants, sampleStringUnknown)
	b.WithMinMax(sampleStringBeta, sampleStringBeta)
	if b.Min() != sampleStringBeta {
		t.Fatalf("expected Min to be overridden to Beta, got %v", b.Min())
	}

	if b.Max() != sampleStringBeta {
		t.Fatalf("expected Max to be overridden to Beta, got %v", b.Max())
	}
}

func TestBasicStringEnum_EmptyVariants(t *testing.T) {
	b := baseenumer.NewBasicString([]sampleStringEnum{}, sampleStringUnknown)
	if b.Min() != sampleStringUnknown {
		t.Fatalf("expected Min to be Unknown for empty, got %v", b.Min())
	}

	if b.Max() != sampleStringUnknown {
		t.Fatalf("expected Max to be Unknown for empty, got %v", b.Max())
	}
}

func TestMinMaxer_InterfaceConformance(t *testing.T) {
	var intEngine baseenumer.MinMaxer[sampleByteEnum] = baseenumer.NewBasicInteger(sampleByteLabels[:], sampleByteInvalid)
	if intEngine.Min() != sampleByteInvalid {
		t.Fatal("unexpected int MinMaxer Min")
	}

	if intEngine.Max() != sampleByteInactive {
		t.Fatal("unexpected int MinMaxer Max")
	}

	var strEngine baseenumer.MinMaxer[sampleStringEnum] = baseenumer.NewBasicString(sampleStringVariants, sampleStringUnknown)
	if strEngine.Min() != sampleStringAlpha {
		t.Fatal("unexpected str MinMaxer Min")
	}

	if strEngine.Max() != sampleStringBeta {
		t.Fatal("unexpected str MinMaxer Max")
	}
}
