package appfault_test

import (
	"errors"
	"strings"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

// createSampleAppError helper for serialization testing.
func createSampleAppError() (*appfault.AppError, error) {
	rawErr := errors.New("underlying socket closed")
	orig := appfault.Wrap(errtype.Network, rawErr, "dial timeout").
		WithOp("net.dial").
		WithStatusCode(504).
		WithContext("port", 8080)

	return orig, rawErr
}

func TestAppErrorSerializationRoundtrip(t *testing.T) {
	orig, rawErr := createSampleAppError()
	restored, err := appfault.FromJSON([]byte(orig.ToJSONString()))
	if err != nil || restored.Type != orig.Type || restored.StatusCode != 504 {
		t.Fatalf("JSON restore mismatch or error: %v, %+v", err, restored)
	}

	if restored.Cause == nil || restored.Cause.Error() != rawErr.Error() {
		t.Fatalf("expected cause '%s', got '%v'", rawErr.Error(), restored.Cause)
	}
}

func TestAppErrorYAMLSerialization(t *testing.T) {
	appErr := appfault.New(errtype.Unauthorized, "invalid credentials").WithStatusCode(401)
	yamlStr := appErr.ToYAMLString()

	if !strings.Contains(yamlStr, "Type: 10") || !strings.Contains(yamlStr, "StatusCode: 401") {
		t.Fatalf("unexpected YAML string: %s", yamlStr)
	}
}
