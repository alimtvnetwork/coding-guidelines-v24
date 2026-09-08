package filepermtype_test

import (
	"encoding/json"
	"os"
	"testing"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/enum/filepermtype"
)

func TestFilePermType_Interfaces(t *testing.T) {
	var (
		_ baseenumer.BaseEnumer                          = filepermtype.Standard
		_ baseenumer.NumberEnumer                        = filepermtype.Standard
		_ baseenumer.MinMaxer[filepermtype.Variant]      = filepermtype.Standard
		_ baseenumer.BoundedEnumer[filepermtype.Variant] = filepermtype.Standard
		_ baseenumer.Bounder[filepermtype.Variant]       = filepermtype.Standard
		_ json.Marshaler                                 = filepermtype.Standard
		_ json.Unmarshaler                               = (*filepermtype.Variant)(nil)
	)
}

func TestFilePermType_Basics(t *testing.T) {
	p := filepermtype.Standard
	if p.Uint32() != 0644 {
		t.Fatalf("expected uint32 0644, got %o", p.Uint32())
	}

	if p.Int() != 0644 || p.Code() != 0644 {
		t.Fatalf("expected int/code 0644")
	}

	if p.OctalString() != "0644" {
		t.Fatalf("expected octal '0644', got %s", p.OctalString())
	}

	if p.PosixString() != "rw-r--r--" {
		t.Fatalf("expected posix 'rw-r--r--', got %s", p.PosixString())
	}

	if p.Mode() != 0644 {
		t.Fatalf("expected mode 0644, got %v", p.Mode())
	}

	var zero filepermtype.Variant
	if zero.Mode() != 0644 {
		t.Fatalf("expected zero mode to default to Standard")
	}
}

func TestFilePermType_Predicates(t *testing.T) {
	std := filepermtype.Standard
	if !std.IsPublic() || std.IsPrivate() || std.IsExecutable() {
		t.Fatalf("standard predicates failed")
	}

	if !std.IsOwnerReadable() || !std.IsOwnerWritable() || !std.IsGroupReadable() || !std.IsOtherReadable() {
		t.Fatalf("standard read/write flags failed")
	}

	if std.IsGroupWritable() || std.IsOtherWritable() {
		t.Fatalf("standard should not have group/other write")
	}

	exec := filepermtype.Executable
	if !exec.IsExecutable() || !exec.IsValid() || !exec.IsEnum() {
		t.Fatalf("executable predicates failed")
	}
}

func TestFilePermType_Mutators(t *testing.T) {
	std := filepermtype.Standard
	priv := std.WithPrivate()
	if !priv.IsPrivate() || priv.Uint32() != 0600 {
		t.Fatalf("expected WithPrivate 0600, got %o", priv.Uint32())
	}

	ro := std.WithReadOnly()
	if ro.Uint32() != 0444 {
		t.Fatalf("expected WithReadOnly 0444, got %o", ro.Uint32())
	}

	withExec := std.WithExecutable()
	if !withExec.IsExecutable() || withExec.Uint32() != 0755 {
		t.Fatalf("expected WithExecutable 0755, got %o", withExec.Uint32())
	}
}

func TestFilePermType_Naming(t *testing.T) {
	if filepermtype.Standard.Name() != "Standard(0644)" {
		t.Fatalf("Standard Name mismatch")
	}

	if filepermtype.Standard.String() != "Standard(0644)" {
		t.Fatalf("Standard String mismatch")
	}

	if filepermtype.Standard.ValueString() != "0644" {
		t.Fatalf("Standard ValueString mismatch")
	}

	custom := filepermtype.Variant(0611)
	if custom.Name() != "Perm(0611)" {
		t.Fatalf("custom Name mismatch: %s", custom.Name())
	}
}

func TestFilePermType_AllAndValues(t *testing.T) {
	all := filepermtype.All()
	if len(all) == 0 {
		t.Fatalf("expected non-empty All")
	}

	vals := filepermtype.Values()
	if len(vals) != len(all) {
		t.Fatalf("expected matching lengths for All and Values")
	}
}

func TestFilePermType_Parse(t *testing.T) {
	val, isOk := filepermtype.Parse("0644")
	if !isOk || val != filepermtype.Standard {
		t.Fatalf("parse 0644 failed")
	}

	valAlias, isOkAlias := filepermtype.ParsePerm("0755")
	if !isOkAlias || valAlias != filepermtype.Executable {
		t.Fatalf("parse 0755 failed")
	}

	_, isEmptyOk := filepermtype.Parse("")
	if isEmptyOk {
		t.Fatalf("expected failure on empty string")
	}

	_, isInvalidOk := filepermtype.Parse("invalid-octal")
	if isInvalidOk {
		t.Fatalf("expected failure on invalid octal")
	}

	if filepermtype.ParseOrZero("0644") != filepermtype.Standard {
		t.Fatalf("ParseOrZero 0644 failed")
	}

	if filepermtype.ParseOrInvalid("bogus") != filepermtype.None {
		t.Fatalf("ParseOrInvalid bogus failed")
	}

	if filepermtype.ParseOrUnknown("bogus") != filepermtype.None {
		t.Fatalf("ParseOrUnknown bogus failed")
	}
}

func TestFilePermType_FromFileMode(t *testing.T) {
	mode := os.FileMode(0755)
	p := filepermtype.FromFileMode(mode)
	if p != filepermtype.Executable {
		t.Fatalf("expected Executable, got %o", p.Uint32())
	}
}

func TestFilePermType_JSON(t *testing.T) {
	p := filepermtype.Standard
	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	if string(data) != `"0644"` {
		t.Fatalf("expected '\"0644\"', got %s", string(data))
	}

	var decoded filepermtype.Variant
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if decoded != p {
		t.Fatalf("expected %v, got %v", p, decoded)
	}

	var numDecoded filepermtype.Variant
	if err := json.Unmarshal([]byte(`420`), &numDecoded); err != nil {
		t.Fatalf("unmarshal numeric failed: %v", err)
	}

	if numDecoded != filepermtype.Standard {
		t.Fatalf("expected Standard, got %v", numDecoded)
	}

	var nullDecoded filepermtype.Variant
	if err := json.Unmarshal([]byte(`null`), &nullDecoded); err != nil {
		t.Fatalf("unmarshal null failed: %v", err)
	}

	if nullDecoded != filepermtype.Standard {
		t.Fatalf("expected Standard on null, got %v", nullDecoded)
	}
}

func TestFilePermType_Boundary(t *testing.T) {
	if filepermtype.Min() != filepermtype.None {
		t.Fatalf("expected Min to be None (0000)")
	}

	if filepermtype.Max() != filepermtype.Variant(07777) {
		t.Fatalf("expected Max to be 07777")
	}

	if !filepermtype.None.IsMin() {
		t.Fatalf("expected None.IsMin() to be true")
	}

	if filepermtype.Standard.IsMin() {
		t.Fatalf("expected Standard.IsMin() to be false")
	}
}

func TestFilePermType_BoundaryMax(t *testing.T) {
	maxVal := filepermtype.Variant(07777)
	if !maxVal.IsMax() {
		t.Fatalf("expected 07777.IsMax() to be true")
	}

	if filepermtype.None.IsMax() {
		t.Fatalf("expected None.IsMax() to be false")
	}

	if filepermtype.Standard.Min() != filepermtype.None {
		t.Fatalf("expected Standard.Min() to be None")
	}

	if filepermtype.Standard.Max() != maxVal {
		t.Fatalf("expected Standard.Max() to be 07777")
	}
}

func TestFilePermType_IsInRange(t *testing.T) {
	if !filepermtype.Standard.IsInRange(filepermtype.None, filepermtype.Variant(07777)) {
		t.Fatalf("expected Standard to be in range [None, 07777]")
	}

	if filepermtype.Variant(010000).IsInRange(filepermtype.None, filepermtype.Variant(07777)) {
		t.Fatalf("expected 010000 to not be in range [None, 07777]")
	}

	if !filepermtype.None.IsInRange(filepermtype.None, filepermtype.Standard) {
		t.Fatalf("expected None to be in range [None, Standard]")
	}
}
