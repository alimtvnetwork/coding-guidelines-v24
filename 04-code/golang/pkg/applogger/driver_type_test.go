package applogger_test

import (
	"testing"

	"coding-guidelines/common/pkg/applogger"
)

func TestDriverType_Basics(t *testing.T) {
	d := applogger.DriverConsole
	if d.Name() != "Console" || d.String() != "Console" {
		t.Fatalf("expected Console name, got %s", d.Name())
	}

	if !d.IsValid() {
		t.Fatalf("expected DriverConsole to be valid")
	}

	drivers := []struct {
		driver       applogger.DriverType
		expectedName string
	}{
		{applogger.DriverConsole, "Console"},
		{applogger.DriverFile, "File"},
		{applogger.DriverSQLite, "SQLite"},
		{applogger.DriverZap, "Zap"},
		{applogger.DriverComposite, "Composite"},
	}

	for _, tc := range drivers {
		if tc.driver.Name() != tc.expectedName {
			t.Errorf("expected %s, got %s", tc.expectedName, tc.driver.Name())
		}
	}

	invalid := applogger.DriverType(99)
	if invalid.IsValid() {
		t.Fatalf("expected invalid driver to return IsValid false")
	}
}
