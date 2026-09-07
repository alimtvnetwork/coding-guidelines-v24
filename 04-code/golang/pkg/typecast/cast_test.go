package typecast

import (
	"errors"
	"testing"
)

type myIntType int

type samplePerson struct {
	Name string `json:"name"`
}

func TestReflectSetTo_PrimitiveInt(t *testing.T) {
	var dest int
	err := ReflectSetTo(42, &dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dest != 42 {
		t.Fatalf("expected 42, got %d", dest)
	}
}

func TestReflectSetTo_PrimitiveIntPtr(t *testing.T) {
	src := 42
	var dest int
	err := ReflectSetTo(&src, &dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dest != 42 {
		t.Fatalf("expected 42, got %d", dest)
	}
}

func TestReflectSetTo_PrimitiveInt64(t *testing.T) {
	var dest int64
	err := ReflectSetTo(int64(999), &dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dest != int64(999) {
		t.Fatalf("expected 999, got %d", dest)
	}
}

func TestReflectSetTo_PrimitiveString(t *testing.T) {
	var dest string
	err := ReflectSetTo("hello", &dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dest != "hello" {
		t.Fatalf("expected hello, got %s", dest)
	}
}

func TestReflectSetTo_PrimitiveBool(t *testing.T) {
	var dest bool
	err := ReflectSetTo(true, &dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !dest {
		t.Fatalf("expected true, got %v", dest)
	}
}

func TestReflectSetTo_PrimitiveFloat64(t *testing.T) {
	var dest float64
	err := ReflectSetTo(3.14, &dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dest != 3.14 {
		t.Fatalf("expected 3.14, got %f", dest)
	}
}

func TestReflectSetTo_PointerToPointer(t *testing.T) {
	src := 100
	dest := new(int)
	err := ReflectSetTo(&src, dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if *dest != 100 {
		t.Fatalf("expected 100, got %d", *dest)
	}
}

func TestReflectSetTo_StructPointerToPointer(t *testing.T) {
	src := samplePerson{Name: "Alice"}
	dest := new(samplePerson)
	err := ReflectSetTo(&src, dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dest.Name != "Alice" {
		t.Fatalf("expected Alice, got %s", dest.Name)
	}
}

func TestReflectSetTo_BytesToStruct(t *testing.T) {
	var dest samplePerson
	err := ReflectSetTo([]byte(`{"name":"Alice"}`), &dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dest.Name != "Alice" {
		t.Fatalf("expected Alice, got %s", dest.Name)
	}
}

func TestReflectSetTo_StructToBytes(t *testing.T) {
	src := samplePerson{Name: "Bob"}
	var dest []byte
	err := ReflectSetTo(src, &dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(dest) != `{"name":"Bob"}` {
		t.Fatalf("expected json, got %s", string(dest))
	}
}

func TestReflectSetTo_BytesToPrimitive(t *testing.T) {
	var dest int
	err := ReflectSetTo([]byte("42"), &dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dest != 42 {
		t.Fatalf("expected 42, got %d", dest)
	}
}

func TestReflectSetTo_ErrorNilDest(t *testing.T) {
	err := ReflectSetTo(42, nil)
	if !errors.Is(err, ErrInvalidNullPointer) {
		t.Fatalf("expected ErrInvalidNullPointer, got %v", err)
	}

	var typedNil *int
	err = ReflectSetTo(42, typedNil)
	if !errors.Is(err, ErrInvalidNullPointer) {
		t.Fatalf("expected ErrInvalidNullPointer, got %v", err)
	}
}

func TestReflectSetTo_ErrorNonPointerDest(t *testing.T) {
	err := ReflectSetTo(42, 10)
	if !errors.Is(err, ErrDestinationNotPointer) {
		t.Fatalf("expected ErrDestinationNotPointer, got %v", err)
	}

	err = ReflectSetTo([]byte("foo"), "bar")
	if !errors.Is(err, ErrDestinationNotPointer) {
		t.Fatalf("expected ErrDestinationNotPointer, got %v", err)
	}
}

func TestReflectSetTo_ErrorIncompatible(t *testing.T) {
	var myIntDest myIntType
	err := ReflectSetTo(42, &myIntDest)
	if !errors.Is(err, ErrTypeMismatch) {
		t.Fatalf("expected ErrTypeMismatch, got %v", err)
	}

	var strDest string
	err = ReflectSetTo(42, &strDest)
	if !errors.Is(err, ErrTypeMismatch) {
		t.Fatalf("expected ErrTypeMismatch, got %v", err)
	}
}

func TestReflectSetTo_ErrorNilSource(t *testing.T) {
	var dest int
	err := ReflectSetTo(nil, &dest)
	if !errors.Is(err, ErrInvalidValueType) {
		t.Fatalf("expected ErrInvalidValueType, got %v", err)
	}

	var nilPtr *int
	err = ReflectSetTo(nilPtr, &dest)
	if !errors.Is(err, ErrInvalidValueType) {
		t.Fatalf("expected ErrInvalidValueType, got %v", err)
	}
}

func TestReflectSetTo_BothNil(t *testing.T) {
	err := ReflectSetTo(nil, nil)
	if err != nil {
		t.Fatalf("expected nil error for both nil, got %v", err)
	}
}

func TestToBytes_Basic(t *testing.T) {
	b, err := ToBytes([]byte("hello"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(b) != "hello" {
		t.Fatalf("expected hello, got %s", string(b))
	}
}

func TestToBytes_String(t *testing.T) {
	b, err := ToBytes("world")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(b) != "world" {
		t.Fatalf("expected world, got %s", string(b))
	}
}

func TestToBytes_StringSlice(t *testing.T) {
	b, err := ToBytes([]string{"line1", "line2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(b) != "line1\nline2" {
		t.Fatalf("expected newline separated, got %s", string(b))
	}
}

func TestToBytes_Error(t *testing.T) {
	sampleErr := errors.New("boom")
	b, err := ToBytes(sampleErr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(b) != "boom" {
		t.Fatalf("expected boom, got %s", string(b))
	}
}

func TestToBytes_JSONFallback(t *testing.T) {
	p := samplePerson{Name: "Charlie"}
	b, err := ToBytes(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(b) != `{"name":"Charlie"}` {
		t.Fatalf("expected json, got %s", string(b))
	}
}

func TestToBytes_Nil(t *testing.T) {
	b, err := ToBytes(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(b) != 0 {
		t.Fatalf("expected empty bytes, got %d", len(b))
	}
}

func TestToJson_Success(t *testing.T) {
	p := samplePerson{Name: "Alice"}
	b, err := ToJson(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "{\n  \"name\": \"Alice\"\n}\n"
	if string(b) != expected {
		t.Fatalf("expected %q, got %q", expected, string(b))
	}

	bAlias, _ := ToJSON(p)
	if string(bAlias) != expected {
		t.Fatalf("expected alias ToJSON %q, got %q", expected, string(bAlias))
	}
}

func TestToJson_Error(t *testing.T) {
	ch := make(chan int)
	_, err := ToJson(ch)
	if err == nil {
		t.Fatal("expected error for unmarshalable type")
	}
}

func TestToJsonString_Success(t *testing.T) {
	p := samplePerson{Name: "Alice"}
	s, err := ToJsonString(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "{\n  \"name\": \"Alice\"\n}"
	if s != expected {
		t.Fatalf("expected %q, got %q", expected, s)
	}

	sAlias, _ := ToJSONString(p)
	if sAlias != expected {
		t.Fatalf("expected alias ToJSONString %q, got %q", expected, sAlias)
	}
}

func TestToJsonString_Error(t *testing.T) {
	ch := make(chan int)
	_, err := ToJsonString(ch)
	if err == nil {
		t.Fatal("expected error for unmarshalable type")
	}
}

func TestReflectTo(t *testing.T) {
	v1, err := ReflectTo[int](42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if v1 != 42 {
		t.Fatalf("expected 42, got %v", v1)
	}

	_, err = ReflectTo[myIntType](42)
	if !errors.Is(err, ErrTypeMismatch) {
		t.Fatalf("expected ErrTypeMismatch, got %v", err)
	}
}

func TestCastTo(t *testing.T) {
	v1, err := CastTo[int](42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if v1 != 42 {
		t.Fatalf("expected 42, got %v", v1)
	}

	_, err = CastTo[myIntType](42)
	if err == nil {
		t.Fatal("expected error casting int to myIntType")
	}

	_, err = CastTo[int](nil)
	if err == nil {
		t.Fatal("expected error casting nil payload")
	}
}

func BenchmarkReflectSetTo_Primitive(b *testing.B) {
	var dest int
	val := 42
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = ReflectSetTo(val, &dest)
	}
}

func BenchmarkReflectSetTo_BytesToStruct(b *testing.B) {
	raw := []byte(`{"name":"Alice"}`)
	var dest samplePerson
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = ReflectSetTo(raw, &dest)
	}
}

func BenchmarkReflectSetTo_FallbackStruct(b *testing.B) {
	src := samplePerson{Name: "Alice"}
	var dest samplePerson
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = ReflectSetTo(src, &dest)
	}
}
