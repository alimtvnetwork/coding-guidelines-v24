package appfault

import (
	"encoding/json"
	"strings"
	"testing"

	"coding-guidelines/common/pkg/errtype"
	"gopkg.in/yaml.v3"
)

func TestResult_LinesAndSplitting(t *testing.T) {
	text := "line1\r\nline2\nline3"
	res := SuccessResult(text)

	lines := res.Lines()
	if len(lines) != 3 || lines[0] != "line1" || lines[1] != "line2" || lines[2] != "line3" {
		t.Fatalf("unexpected lines: %v", lines)
	}

	linesRes := res.LinesResult()
	if linesRes.IsFailed() || linesRes.Count() != 3 {
		t.Fatalf("expected 3 lines in LinesResult")
	}

	parts := res.Split("line")
	if len(parts) < 4 {
		t.Fatalf("unexpected split parts: %v", parts)
	}

	splitRes := res.SplitResult("\n")
	if splitRes.IsFailed() {
		t.Fatalf("expected success in SplitResult")
	}

	left, right := res.SplitAt(5)
	if left != "line1" {
		t.Fatalf("expected left 'line1', got %q", left)
	}

	if len(right) == 0 {
		t.Fatalf("expected non-empty right")
	}

	byRune := SuccessResult("a:b:c").SplitByRune(':')
	if len(byRune) != 3 || byRune[0] != "a" {
		t.Fatalf("unexpected byRune: %v", byRune)
	}
}

func TestResult_LinesFailure(t *testing.T) {
	fail := FailureResult[string](New(errtype.Validation, "err"))
	if len(fail.Lines()) != 0 {
		t.Fatalf("expected empty lines on failure")
	}

	if fail.LinesResult().IsSuccess() {
		t.Fatalf("expected failed LinesResult")
	}

	if len(fail.Split(",")) != 0 {
		t.Fatalf("expected empty split on failure")
	}

	l, r := fail.SplitAt(2)
	if l != "" || r != "" {
		t.Fatalf("expected empty split at on failure")
	}
}

func TestResult_NumberConversions(t *testing.T) {
	intRes := SuccessResult(42)
	v, ok := intRes.Int()
	if !ok || v != 42 {
		t.Fatalf("expected int 42, got %d", v)
	}

	if intRes.IntDefault(0) != 42 {
		t.Fatalf("expected 42 default")
	}

	strRes := SuccessResult("100")
	if strRes.IntDefault(0) != 100 {
		t.Fatalf("expected parsed 100")
	}

	v64, ok64 := strRes.Int64()
	if !ok64 || v64 != 100 {
		t.Fatalf("expected int64 100")
	}

	if strRes.Int64Default(0) != 100 {
		t.Fatalf("expected 100 int64 default")
	}

	fltRes := SuccessResult("3.14")
	f, fOk := fltRes.Float64()
	if !fOk || f < 3.13 || f > 3.15 {
		t.Fatalf("expected float 3.14, got %f", f)
	}

	if fltRes.Float64Default(0) < 3.13 {
		t.Fatalf("expected 3.14 float default")
	}

	if d, dOk := fltRes.Double(); !dOk || d < 3.13 {
		t.Fatalf("expected double 3.14")
	}

	if fltRes.DoubleDefault(0) < 3.13 {
		t.Fatalf("expected double default")
	}

	byteRes := SuccessResult(255)
	b, bOk := byteRes.Byte()
	if !bOk || b != 255 {
		t.Fatalf("expected byte 255")
	}

	if byteRes.ByteDefault(0) != 255 {
		t.Fatalf("expected byte default 255")
	}
}

func TestResult_TypeInspectionAndPredicates(t *testing.T) {
	numRes := SuccessResult(123)
	if !numRes.IsNumber() || !numRes.IsPrimitive() {
		t.Fatalf("expected 123 to be number and primitive")
	}

	if numRes.TypeName() != "int" {
		t.Fatalf("expected type name int, got %s", numRes.TypeName())
	}

	strRes := SuccessResult("hello")
	if !strRes.IsStringType() || strRes.Length() != 5 {
		t.Fatalf("expected string type and length 5")
	}

	sliceRes := SuccessResult([]int{1, 2, 3})
	if !sliceRes.IsSliceOrArray() || sliceRes.Length() != 3 {
		t.Fatalf("expected slice and length 3")
	}

	mapRes := SuccessResult(map[string]int{"a": 1})
	if !mapRes.IsMap() || mapRes.Length() != 1 {
		t.Fatalf("expected map and length 1")
	}

	type sampleStruct struct{ Name string }
	structRes := SuccessResult(sampleStruct{Name: "test"})
	if !structRes.IsStruct() {
		t.Fatalf("expected struct")
	}

	ptrRes := SuccessResult(&sampleStruct{Name: "ptr"})
	if !ptrRes.IsPointer() {
		t.Fatalf("expected pointer")
	}
}

func TestResult_ReflectTo(t *testing.T) {
	src := SuccessResult("hello world")
	var target string
	err := src.ReflectTo(&target)
	if err != nil || target != "hello world" {
		t.Fatalf("expected target 'hello world', got %q, err: %v", target, err)
	}

	fail := FailureResult[string](New(errtype.Validation, "bad"))
	var failTarget string
	failErr := fail.ReflectTo(&failTarget)
	if failErr == nil {
		t.Fatalf("expected error from failed result ReflectTo")
	}

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	personMap := map[string]any{"Name": "Bob", "Age": 25}
	mapRes := SuccessResult(personMap)
	m := mapRes.Map()
	if m["Name"] != "Bob" {
		t.Fatalf("expected Bob in Map()")
	}

	mapResult := mapRes.ToMapResult()
	if mapResult.IsFailed() || mapResult.Count() != 2 {
		t.Fatalf("expected ToMapResult success with 2 entries")
	}
}

func TestResult_OutputAndSerialization(t *testing.T) {
	res := SuccessResult(map[string]string{"greeting": "hello"})
	if !strings.Contains(res.String(), "hello") {
		t.Fatalf("expected String() to contain hello")
	}

	pretty := res.PrettyJson()
	if !strings.Contains(pretty, "greeting") || !strings.Contains(pretty, "\n") {
		t.Fatalf("expected indented JSON in PrettyJson: %s", pretty)
	}

	if res.ToJsonPretty() != pretty {
		t.Fatalf("expected ToJsonPretty to match PrettyJson")
	}

	prettyMap := res.PrettyMap()
	if !strings.Contains(prettyMap, "greeting") {
		t.Fatalf("expected greeting in PrettyMap")
	}

	yamlStr, err := res.ToYaml()
	if err != nil || !strings.Contains(yamlStr, "greeting") {
		t.Fatalf("expected valid YAML: %s, err: %v", yamlStr, err)
	}

	if res.Yaml() != yamlStr {
		t.Fatalf("expected Yaml() to match ToYaml()")
	}

	chain := res.PrintConsole()
	if chain.Value()["greeting"] != "hello" {
		t.Fatalf("expected PrintConsole to return same Result for chaining")
	}
}

func TestResult_JsonAndYamlRoundtrip(t *testing.T) {
	orig := SuccessResult("roundtrip-payload")
	jsonBytes, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("failed to marshal Result to JSON: %v", err)
	}

	var unmarshaled Result[string]
	err = json.Unmarshal(jsonBytes, &unmarshaled)
	if err != nil || unmarshaled.Value() != "roundtrip-payload" {
		t.Fatalf("failed to unmarshal JSON: %v, val: %s", err, unmarshaled.Value())
	}

	yamlBytes, err := yaml.Marshal(orig)
	if err != nil {
		t.Fatalf("failed to marshal Result to YAML: %v", err)
	}

	var unmarshaledYaml Result[string]
	err = yaml.Unmarshal(yamlBytes, &unmarshaledYaml)
	if err != nil || unmarshaledYaml.Value() != "roundtrip-payload" {
		t.Fatalf("failed to unmarshal YAML: %v, val: %s", err, unmarshaledYaml.Value())
	}
}
