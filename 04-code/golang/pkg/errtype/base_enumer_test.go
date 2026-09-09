package errtype_test

import (
	"testing"

	"coding-guidelines/common/pkg/enum/logleveltype"
	"coding-guidelines/common/pkg/enum/processstatetype"
	"coding-guidelines/common/pkg/errtype"
)

func TestBaseEnum_VariationConforms(t *testing.T) {
	var e errtype.BaseEnumer = errtype.Validation
	if e.Name() != "Validation" {
		t.Fatalf("expected Name() == 'Validation', got %s", e.Name())
	}

	if e.ValueString() != "2" {
		t.Fatalf("expected ValueString() == '2', got %s", e.ValueString())
	}

	if !e.IsEnum() {
		t.Fatal("expected Validation to be registered enum")
	}

	var aliasE errtype.BaseEnum = e
	if aliasE.Name() != "Validation" {
		t.Fatalf("expected alias Name() == 'Validation', got %s", aliasE.Name())
	}

	var ne errtype.NumberEnumer = errtype.NotFound
	if ne.Int() != 3 || ne.Code() != 3 {
		t.Fatalf("unexpected number enum values: int=%d code=%d", ne.Int(), ne.Code())
	}

	var aliasNE errtype.NumberEnum = ne
	if aliasNE.Int() != 3 || aliasNE.Code() != 3 {
		t.Fatalf("unexpected alias number enum values: int=%d code=%d", aliasNE.Int(), aliasNE.Code())
	}
}

func TestToEnum_GenericHelper(t *testing.T) {
	foundState, ok := errtype.ToEnum("Completed", processstatetype.All())
	if !ok || foundState != processstatetype.Completed {
		t.Fatalf("expected ToEnum to find Completed, got %v, ok=%v", foundState, ok)
	}

	foundLvl, ok := errtype.ToEnum("4", logleveltype.All())
	if !ok || foundLvl != logleveltype.Error {
		t.Fatalf("expected ToEnum to find Error by '4', got %v, ok=%v", foundLvl, ok)
	}

	_, ok = errtype.ToEnum("non-existent", processstatetype.All())
	if ok {
		t.Fatal("expected non-existent enum to return ok=false")
	}
}
