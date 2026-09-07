package baseenumer_test

import (
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
)

type testByteVariant byte

const (
	testInvalid testByteVariant = iota
	testAlpha
	testBeta
	testGamma
)

var testLabels = []string{
	"Unknown",
	"Alpha",
	"Beta",
	"Gamma",
}

func TestCompileMap_Lookups(t *testing.T) {
	m := baseenumer.CompileMap(testLabels, testInvalid)

	if m["Alpha"] != testAlpha || m["alpha"] != testAlpha || m["ALPHA"] != testAlpha {
		t.Fatalf("expected Alpha lookup to resolve")
	}

	if m["1"] != testAlpha || m["2"] != testBeta {
		t.Fatalf("expected numeric string lookups to resolve")
	}

	if m["unknown"] != testInvalid || m["invalid"] != testInvalid {
		t.Fatalf("expected fallback aliases to resolve to invalid")
	}
}

func TestSliceValuesAndVariants(t *testing.T) {
	vals := baseenumer.SliceValues(testLabels)
	if len(vals) != 3 || vals[0] != "Alpha" || vals[2] != "Gamma" {
		t.Fatalf("unexpected SliceValues result: %v", vals)
	}

	vars := baseenumer.SliceVariants[testByteVariant](testLabels)
	if len(vars) != 3 || vars[0] != testAlpha || vars[2] != testGamma {
		t.Fatalf("unexpected SliceVariants result: %v", vars)
	}

	emptyVals := baseenumer.SliceValues([]string{"OnlyOne"})
	if len(emptyVals) != 0 {
		t.Fatalf("expected empty slice for single element")
	}

	emptyVars := baseenumer.SliceVariants[testByteVariant]([]string{"OnlyOne"})
	if len(emptyVars) != 0 {
		t.Fatalf("expected empty variants for single element")
	}
}

func TestFormatNameValue(t *testing.T) {
	formatted := baseenumer.FormatNameValue("LogLevel", 3)
	if formatted != "LogLevel(3)" {
		t.Fatalf("expected 'LogLevel(3)', got %s", formatted)
	}
}

func TestIsBetweenAndIsNotBetween(t *testing.T) {
	if !baseenumer.IsBetween(5, 1, 10) {
		t.Fatalf("expected 5 to be between 1 and 10")
	}

	if baseenumer.IsBetween(0, 1, 10) {
		t.Fatalf("expected 0 not to be between 1 and 10")
	}

	if !baseenumer.IsNotBetween(0, 1, 10) {
		t.Fatalf("expected 0 to be NotBetween 1 and 10")
	}

	if baseenumer.IsNotBetween(5, 1, 10) {
		t.Fatalf("expected 5 not to be NotBetween 1 and 10")
	}
}

func TestParseLookup(t *testing.T) {
	m := baseenumer.CompileMap(testLabels, testInvalid)

	v, trimmed, ok := baseenumer.ParseLookup("  Beta  ", m)
	if !ok || v != testBeta || trimmed != "Beta" {
		t.Fatalf("expected successful ParseLookup for Beta")
	}

	_, emptyTrimmed, emptyOk := baseenumer.ParseLookup("   ", m)
	if emptyOk || len(emptyTrimmed) != 0 {
		t.Fatalf("expected empty string lookup to fail")
	}

	_, _, missingOk := baseenumer.ParseLookup("non-existent", m)
	if missingOk {
		t.Fatalf("expected missing lookup to fail")
	}
}

func TestFormatErrorHelpers(t *testing.T) {
	parseErr := baseenumer.FormatParseError("testType", "foo", []string{"Alpha", "Beta"})
	if parseErr != "unknown testType variant \"foo\", supported variants: [Alpha, Beta]" {
		t.Fatalf("unexpected parse error string: %s", parseErr)
	}

	emptyErr := baseenumer.FormatEmptyParseError("testType")
	if emptyErr != "cannot parse empty string as testType" {
		t.Fatalf("unexpected empty error string: %s", emptyErr)
	}

	rangeErr := baseenumer.FormatNumericRangeError("testType", 99, 5)
	if rangeErr.Error() != "invalid testType numeric value 99, supported range: 0..5" {
		t.Fatalf("unexpected range error string: %s", rangeErr.Error())
	}
}
