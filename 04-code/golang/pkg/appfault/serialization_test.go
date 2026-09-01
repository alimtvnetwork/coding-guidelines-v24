package appfault_test

import (
	"errors"
	"strings"
	"testing"

	"coding-guidelines/common/pkg/appfault"
)

// createSampleAppError helper for serialization testing.
func createSampleAppError() (*appfault.AppError, error) {
	rawErr := errors.New("underlying socket closed")
	orig := appfault.WrapWithDetails(rawErr, "net.dial", "E3001", "dial timeout", "network", appfault.ErrorTypeExecution, appfault.SeverityError, map[string]any{"port": 8080})
	orig.WithStatusCode(504)

	return orig, rawErr
}

func TestAppErrorSerializationRoundtrip(t *testing.T) {
	orig, rawErr := createSampleAppError()
	restored, err := appfault.FromJSON([]byte(orig.ToJSONString()))
	if err != nil || restored.Op != orig.Op || restored.StatusCode != 504 {
		t.Fatalf("JSON restore mismatch or error: %v, %+v", err, restored)
	}

	if restored.Cause == nil || restored.Cause.Error() != rawErr.Error() {
		t.Fatalf("expected cause '%s', got '%v'", rawErr.Error(), restored.Cause)
	}
}

func TestAppErrorYAMLSerialization(t *testing.T) {
	appErr := appfault.NewSimple("auth.login", "E1001").WithStatusCode(401)
	yamlStr := appErr.ToYAMLString()

	if !strings.Contains(yamlStr, "Op: \"auth.login\"") || !strings.Contains(yamlStr, "StatusCode: 401") {
		t.Fatalf("unexpected YAML string: %s", yamlStr)
	}
}
