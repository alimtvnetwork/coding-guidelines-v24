package typecast

import (
	"testing"
)

type MyInt int

func TestReflectTo(t *testing.T) {
	// Int to Int conversion (identity)
	v1, err := ReflectTo[int](42)
	if err != nil || v1 != 42 {
		t.Errorf("expected 42, got %v, err %v", v1, err)
	}

	// Int to MyInt conversion
	v2, err := ReflectTo[MyInt](42)
	if err != nil || v2 != MyInt(42) {
		t.Errorf("expected MyInt(42), got %v, err %v", v2, err)
	}

	// MyInt to Int conversion
	v3, err := ReflectTo[int](MyInt(100))
	if err != nil || v3 != 100 {
		t.Errorf("expected 100, got %v, err %v", v3, err)
	}

	// Float to Int
	v4, err := ReflectTo[int](3.14)
	if err != nil || v4 != 3 {
		t.Errorf("expected 3, got %v, err %v", v4, err)
	}

	// String to Int (should fail)
	_, err = ReflectTo[int]("hello")
	if err == nil {
		t.Errorf("expected error converting string to int")
	}
}

func TestCastTo(t *testing.T) {
	v1, err := CastTo[int](42)
	if err != nil || v1 != 42 {
		t.Errorf("expected 42, got %v, err %v", v1, err)
	}

	// Int to MyInt (should fail type assertion)
	_, err = CastTo[MyInt](42)
	if err == nil {
		t.Errorf("expected error casting int to MyInt")
	}
}
