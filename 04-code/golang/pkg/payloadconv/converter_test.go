package payloadconv

import (
	"strings"
	"testing"
)

func TestToBytes(t *testing.T) {
	// Byte case
	resB := ToBytes([]byte("hello"))
	b := resB.Data()
	if string(b) != "hello" {
		t.Errorf("byte conversion failed")
	}

	// String case
	resS := ToBytes("world")
	s := resS.Data()
	if string(s) != "world" {
		t.Errorf("string conversion failed")
	}

	// String array case
	resArr := ToBytes([]string{"line1", "line2"})
	arr := resArr.Data()
	if string(arr) != "line1\nline2\n" {
		t.Errorf("array conversion failed, got %s", string(arr))
	}

	// Struct case (JSON)
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "Alice", Age: 30}
	resP := ToBytes(p)
	j := resP.Data()
	if !strings.Contains(string(j), `"name": "Alice"`) {
		t.Errorf("JSON struct conversion failed, got %s", string(j))
	}
}
