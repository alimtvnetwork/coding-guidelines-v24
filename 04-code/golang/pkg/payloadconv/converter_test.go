package payloadconv

import (
	"strings"
	"testing"
)

func TestToBytes(t *testing.T) {
	// Byte case
	b := ToBytes([]byte("hello")).Data()
	if string(b) != "hello" {
		t.Errorf("byte conversion failed")
	}

	// String case
	s := ToBytes("world").Data()
	if string(s) != "world" {
		t.Errorf("string conversion failed")
	}

	// String array case
	arr := ToBytes([]string{"line1", "line2"}).Data()
	if string(arr) != "line1\nline2\n" {
		t.Errorf("array conversion failed, got %s", string(arr))
	}

	// Struct case (JSON)
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "Alice", Age: 30}
	j := ToBytes(p).Data()
	if !strings.Contains(string(j), `"name": "Alice"`) {
		t.Errorf("JSON struct conversion failed, got %s", string(j))
	}
}
