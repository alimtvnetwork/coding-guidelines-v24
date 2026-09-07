package typecast

import (
	"testing"
)

type MyInt int

func TestReflectSetTo(t *testing.T) {
	// 1. Value to Pointer (same base type)
	var intDest int
	err := ReflectSetTo(42, &intDest)
	if err != nil || intDest != 42 {
		t.Errorf("expected 42, got %v, err %v", intDest, err)
	}

	// 2. Pointer to Pointer (same type)
	fromPtr := 100
	ptrDest := new(int) // ptrDest is *int, passed as *int to match left type
	err = ReflectSetTo(&fromPtr, ptrDest)
	if err != nil || *ptrDest != 100 {
		t.Errorf("expected 100, got %v, err %v", ptrDest, err)
	}

	// 3. Bytes to Struct
	type Person struct {
		Name string `json:"name"`
	}
	var p Person
	err = ReflectSetTo([]byte(`{"name":"Alice"}`), &p)
	if err != nil || p.Name != "Alice" {
		t.Errorf("expected Alice, got %v, err %v", p.Name, err)
	}

	// 4. Struct to Bytes
	p2 := Person{Name: "Bob"}
	var b []byte
	err = ReflectSetTo(p2, &b)
	if err != nil || string(b) != `{"name":"Bob"}` {
		t.Errorf("expected JSON bytes, got %s, err %v", string(b), err)
	}

	// 5. Type mismatch (strict)
	var myIntDest MyInt
	err = ReflectSetTo(42, &myIntDest) // int != MyInt
	if err != ErrTypeMismatch {
		t.Errorf("expected ErrTypeMismatch, got %v", err)
	}
}

func TestReflectTo(t *testing.T) {
	v1, err := ReflectTo[int](42)
	if err != nil || v1 != 42 {
		t.Errorf("expected 42, got %v, err %v", v1, err)
	}

	_, err = ReflectTo[MyInt](42) // int != MyInt
	if err != ErrTypeMismatch {
		t.Errorf("expected ErrTypeMismatch, got %v", err)
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
