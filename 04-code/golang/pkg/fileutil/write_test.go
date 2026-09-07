package fileutil

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestUnifiedWrite(t *testing.T) {
	tmp := t.TempDir()

	t.Run("Write Any Struct as JSON", func(t *testing.T) {
		path := filepath.Join(tmp, "person.json")
		type Person struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		p := Person{Name: "Charlie", Age: 25}

		res := Write(path, p, FilePermStandard)
		if res.IsFailure() {
			t.Fatalf("Write struct failed: %v", res.Fault().Error())
		}

		txtRes := ReadText(path)
		if txtRes.IsFailure() || !strings.Contains(txtRes.Data(), `"name": "Charlie"`) {
			t.Fatalf("Expected JSON output in file, got %s", txtRes.Data())
		}
	})

	t.Run("Write Any Array as Lines", func(t *testing.T) {
		path := filepath.Join(tmp, "lines.txt")
		lines := []string{"first line", "second line"}

		res := Write(path, lines, FilePermStandard)
		if res.IsFailure() {
			t.Fatalf("Write lines failed: %v", res.Fault().Error())
		}

		readRes := ReadLines(path)
		if readRes.IsFailure() || len(readRes.Data()) != 2 {
			t.Fatalf("Expected 2 lines, got %v", readRes.Data())
		}
	})

	t.Run("Write String and Bytes", func(t *testing.T) {
		pathStr := filepath.Join(tmp, "text.txt")
		resStr := WriteString(pathStr, "pure string", FilePermStandard)
		if resStr.IsFailure() {
			t.Fatalf("WriteString failed: %v", resStr.Fault().Error())
		}

		pathByte := filepath.Join(tmp, "bytes.bin")
		resByte := WriteBytesLocked(pathByte, []byte("byte content"), FilePermStandard)
		if resByte.IsFailure() {
			t.Fatalf("WriteBytesLocked failed: %v", resByte.Fault().Error())
		}
	})
}
