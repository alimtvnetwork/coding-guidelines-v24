package baseenumer_test

import (
	"testing"
	"time"

	"coding-guidelines/common/pkg/baseenumer"
)

func TestResolveTypeName_NilPointer(t *testing.T) {
	name := baseenumer.ResolveTypeName[int](nil)
	if name != "Variant" {
		t.Fatalf("expected Variant for nil, got %q", name)
	}
}

func TestResolveTypeName_BuiltinPrimitives(t *testing.T) {
	var intVal int
	if got := baseenumer.ResolveTypeName(&intVal); got != "int" {
		t.Fatalf("expected 'int', got %q", got)
	}

	var strVal string
	if got := baseenumer.ResolveTypeName(&strVal); got != "string" {
		t.Fatalf("expected 'string', got %q", got)
	}

	var boolVal bool
	if got := baseenumer.ResolveTypeName(&boolVal); got != "bool" {
		t.Fatalf("expected 'bool', got %q", got)
	}
}

func TestResolveTypeName_PackageStruct(t *testing.T) {
	var enumObj baseenumer.BasicIntegerEnum[int]
	if got := baseenumer.ResolveTypeName(&enumObj); got != "baseenumer" {
		t.Fatalf("expected 'baseenumer', got %q", got)
	}
}

func TestResolveTypeName_StdlibPackageStruct(t *testing.T) {
	var tm time.Time
	if got := baseenumer.ResolveTypeName(&tm); got != "time" {
		t.Fatalf("expected 'time', got %q", got)
	}
}
