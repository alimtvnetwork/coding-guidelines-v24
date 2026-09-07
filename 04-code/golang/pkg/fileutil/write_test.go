package fileutil

import (
	"path/filepath"
	"strings"
	"testing"
)

type samplePerson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestWrite_StructAsJSON(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "person.json")
	p := samplePerson{Name: "Charlie", Age: 25}

	res := Write(path, p, FilePermStandard)
	if res.IsFailure() {
		t.Fatalf("Write struct failed: %v", res.Fault().Error())
	}

	verifyJSONOutput(t, path)
}

func verifyJSONOutput(t *testing.T, path string) {
	txtRes := ReadText(path)
	if txtRes.IsFailure() {
		t.Fatalf("ReadText failed: %v", txtRes.Fault().Error())
	}

	if !strings.Contains(txtRes.Data(), `"name": "Charlie"`) {
		t.Fatalf("Expected JSON output in file, got %s", txtRes.Data())
	}
}

func TestWrite_ArrayAsLines(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "lines.txt")
	lines := []string{"first line", "second line"}

	res := Write(path, lines, FilePermStandard)
	if res.IsFailure() {
		t.Fatalf("Write lines failed: %v", res.Fault().Error())
	}

	verifyLinesOutput(t, path)
}

func verifyLinesOutput(t *testing.T, path string) {
	readRes := ReadLines(path)
	if readRes.IsFailure() {
		t.Fatalf("ReadLines failed: %v", readRes.Fault().Error())
	}

	if len(readRes.Data()) != 2 {
		t.Fatalf("Expected 2 lines, got %v", readRes.Data())
	}
}

func TestWrite_StringAndBytes(t *testing.T) {
	tmp := t.TempDir()
	pathStr := filepath.Join(tmp, "text.txt")
	resStr := WriteString(pathStr, "pure string", FilePermStandard)
	if resStr.IsFailure() {
		t.Fatalf("WriteString failed: %v", resStr.Fault().Error())
	}

	testWriteBytesLocked(t, tmp)
}

func testWriteBytesLocked(t *testing.T, tmp string) {
	pathByte := filepath.Join(tmp, "bytes.bin")
	resByte := WriteBytesLocked(pathByte, []byte("byte content"), FilePermStandard)
	if resByte.IsFailure() {
		t.Fatalf("WriteBytesLocked failed: %v", resByte.Fault().Error())
	}
}

func TestWrite_YAMLAndYAMLLocked(t *testing.T) {
	tmp := t.TempDir()
	pathYaml := filepath.Join(tmp, "config.yaml")
	data := map[string]string{"env": "production"}

	res := WriteYAML(pathYaml, data, FilePermStandard)
	if res.IsFailure() {
		t.Fatalf("WriteYAML failed: %v", res.Fault().Error())
	}

	resLocked := WriteYAMLLocked(pathYaml, data, FilePermStandard)
	if resLocked.IsFailure() {
		t.Fatalf("WriteYAMLLocked failed: %v", resLocked.Fault().Error())
	}
}
