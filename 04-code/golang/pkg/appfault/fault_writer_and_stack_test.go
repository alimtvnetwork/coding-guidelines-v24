package appfault

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"coding-guidelines/common/pkg/errtype"
)

type customTestWriter struct {
	captured string
}

func (w *customTestWriter) WriteFault(out io.Writer, e *AppError) error {
	w.captured = fmt.Sprintf("CUSTOM: %s - %s", e.errType.Name(), e.message)
	_, err := fmt.Fprint(out, w.captured)

	return err
}

func TestCustomFaultWriter(t *testing.T) {
	writer := &customTestWriter{}
	SetGlobalFaultWriter(writer)
	defer SetGlobalFaultWriter(&defaultFaultWriter{})

	err := New(errtype.Validation, "bad parameter")
	var buf bytes.Buffer
	writeErr := err.WriteTo(&buf)
	if writeErr != nil {
		t.Fatalf("WriteTo failed: %v", writeErr)
	}

	if !strings.Contains(buf.String(), "CUSTOM: Validation - bad parameter") {
		t.Errorf("Custom writer output unexpected: %s", buf.String())
	}
}

func TestStackFramesSelfFormat(t *testing.T) {
	frames := NewStackTrace(
		NewStackFrame("HandleRequest", "server/handler.go", 42),
		NewStackFrame("ExecuteQuery", "db/query.go", 108),
	)

	if !frames.IsDefined() {
		t.Errorf("expected frames to be defined")
	}

	if frames.IsEmpty() {
		t.Errorf("expected frames to not be empty")
	}

	formatted := frames.Format("  ")
	if !strings.Contains(formatted, "  #0 HandleRequest\n     server/handler.go:42") {
		t.Errorf("Format with indent unexpected:\n%s", formatted)
	}

	caller := frames.CallerLine()
	if caller != "server/handler.go:42" {
		t.Errorf("Expected caller line server/handler.go:42, got %s", caller)
	}
}

func TestIsDefinerAndIsEmptyer(t *testing.T) {
	var definer IsDefiner
	var emptyer IsEmptyer

	// 1. AppError
	err := New(errtype.Database, "connection timeout")
	definer = err
	emptyer = err
	if !definer.IsDefined() || emptyer.IsEmpty() {
		t.Errorf("AppError definer check failed")
	}

	var nilErr *AppError
	definer = nilErr
	emptyer = nilErr
	if definer.IsDefined() || !emptyer.IsEmpty() {
		t.Errorf("Nil AppError check failed")
	}

	// 2. ContextMap
	cm := NewContextMap().Set("user_id", "42")
	definer = cm
	emptyer = cm
	if !definer.IsDefined() || emptyer.IsEmpty() {
		t.Errorf("ContextMap definer check failed")
	}

	emptyCm := NewContextMap()
	definer = emptyCm
	emptyer = emptyCm
	if definer.IsDefined() || !emptyer.IsEmpty() {
		t.Errorf("Empty ContextMap check failed")
	}

	// 3. StackFrames
	sf := CaptureStackTrace(0)
	definer = sf
	emptyer = sf
	if !definer.IsDefined() || emptyer.IsEmpty() {
		t.Errorf("StackFrames definer check failed")
	}
}

func TestContextMapEnhancements(t *testing.T) {
	cm := NewContextMap().
		Set("b_key", "val_b").
		Set("a_key", "val_a").
		Set("empty_key", "").
		Set("nil_key", nil)

	// NonEmptyKeys
	nonEmpty := cm.NonEmptyKeys()
	if len(nonEmpty) != 2 || nonEmpty[0] != "a_key" || nonEmpty[1] != "b_key" {
		t.Errorf("NonEmptyKeys failed: %v", nonEmpty)
	}

	// Values
	vals := cm.Values()
	if len(vals) != 4 {
		t.Errorf("Values count unexpected: %d", len(vals))
	}

	// JSON roundtrip
	jsonBytes, jsonErr := cm.ToJSON()
	if jsonErr != nil {
		t.Fatalf("ToJSON failed: %v", jsonErr)
	}

	restored, restoreErr := ContextMapFromJSON(jsonBytes)
	if restoreErr != nil {
		t.Fatalf("FromJSON failed: %v", restoreErr)
	}

	if restored.GetString("a_key") != "val_a" {
		t.Errorf("Restored value mismatch: %s", restored.GetString("a_key"))
	}

	// YAML roundtrip
	yamlBytes, yamlErr := cm.ToYAML()
	if yamlErr != nil {
		t.Fatalf("ToYAML failed: %v", yamlErr)
	}

	restoredYaml, yamlRestoreErr := ContextMapFromYAML(yamlBytes)
	if yamlRestoreErr != nil {
		t.Fatalf("FromYAML failed: %v", yamlRestoreErr)
	}

	if restoredYaml.GetString("b_key") != "val_b" {
		t.Errorf("YAML Restored mismatch")
	}
}

func TestConstructorsFrameSkip(t *testing.T) {
	err := New(errtype.IO, "disk full")
	if !strings.Contains(err.stack.CallerLine(), "fault_writer_and_stack_test.go") {
		t.Errorf("New() caller line unexpected: %s", err.stack.CallerLine())
	}

	errType := NewType(errtype.NotFound)
	if !strings.Contains(errType.stack.CallerLine(), "fault_writer_and_stack_test.go") {
		t.Errorf("NewType() caller line unexpected: %s", errType.stack.CallerLine())
	}
}

func TestFinalErrorOutputVisualDemonstration(t *testing.T) {
	err := NewWithContext(errtype.DatabaseNotFound, "user record missing from database", map[string]any{
		"tenant_id": "corp-8812",
		"query":     "SELECT * FROM users WHERE id = 99",
	}).WithStatusCode(404)

	// Print visual formats to verify everything works end-to-end
	t.Logf("\n=== COMPILED DIAGNOSTIC REPORT ===\n%s", err.CompileWithStack())
	t.Logf("\n=== STDOUT BANNER ===\n%s", err.FormatStdout())
	t.Logf("\n=== TEXT LOG FORMAT ===\n%s", err.FormatTextLog())

	if !strings.Contains(err.CompileWithStack(), "Stack Trace:") {
		t.Errorf("CompileWithStack missing stack trace")
	}

	if !strings.Contains(err.FormatStdout(), "tenant_id=corp-8812") {
		t.Errorf("FormatStdout missing context")
	}
}
