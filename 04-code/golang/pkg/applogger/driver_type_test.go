package applogger_test

import (
	"encoding/json"
	"strings"
	"testing"

	"coding-guidelines/common/pkg/applogger"
)

var allTestDrivers = []struct {
	driver       applogger.DriverType
	expectedName string
}{
	{applogger.DriverConsole, "Console"},
	{applogger.DriverFile, "File"},
	{applogger.DriverSQLite, "SQLite"},
	{applogger.DriverZap, "Zap"},
	{applogger.DriverComposite, "Composite"},
	{applogger.DriverRotatingFile, "RotatingFile"},
	{applogger.DriverApi, "Api"},
	{applogger.DriverJsonWriterLogger, "JsonWriterLogger"},
	{applogger.DriverStreamer, "Streamer"},
}

var aliasParseCases = map[string]applogger.DriverType{
	"FileWriter":        applogger.DriverFile,
	"file_writer":       applogger.DriverFile,
	"FileWriterRotator": applogger.DriverRotatingFile,
	"SqliteDbWriter":    applogger.DriverSQLite,
	"API":               applogger.DriverApi,
	"Json":              applogger.DriverJsonWriterLogger,
	"driver(6)":         applogger.DriverApi,
	"Driver(8)":         applogger.DriverStreamer,
	"0":                 applogger.DriverConsole,
}

func TestDriverType_Basics(t *testing.T) {
	d := applogger.DriverConsole
	if d.Name() != "Console" || d.String() != "Console" {
		t.Fatalf("expected Console name, got %s", d.Name())
	}

	if !d.IsValid() {
		t.Fatalf("expected DriverConsole to be valid")
	}

	invalid := applogger.DriverType(99)
	if invalid.IsValid() {
		t.Fatalf("expected invalid driver to return IsValid false")
	}

	if invalid.Name() != "Driver(99)" {
		t.Fatalf("expected Driver(99), got %s", invalid.Name())
	}
}

func TestDriverType_AllNames(t *testing.T) {
	for _, tc := range allTestDrivers {
		if tc.driver.Name() != tc.expectedName {
			t.Errorf("expected %s, got %s", tc.expectedName, tc.driver.Name())
		}

		if !tc.driver.IsValid() {
			t.Errorf("expected %s to be valid", tc.expectedName)
		}
	}
}

func TestDriverType_Aliases(t *testing.T) {
	if applogger.DriverFileWriter != applogger.DriverFile {
		t.Errorf("DriverFileWriter alias mismatch")
	}

	if applogger.DriverFileWriterRotator != applogger.DriverRotatingFile {
		t.Errorf("DriverFileWriterRotator alias mismatch")
	}

	if applogger.DriverSqliteDbWriter != applogger.DriverSQLite {
		t.Errorf("DriverSqliteDbWriter alias mismatch")
	}

	if applogger.DriverAPI != applogger.DriverApi {
		t.Errorf("DriverAPI alias mismatch")
	}

	if applogger.DriverJson != applogger.DriverJsonWriterLogger {
		t.Errorf("DriverJson alias mismatch")
	}
}

func TestDriverType_ParseCanonical(t *testing.T) {
	for _, tc := range allTestDrivers {
		d, isOk := applogger.ParseDriverType(tc.expectedName)
		if !isOk || d != tc.driver {
			t.Errorf("failed parsing canonical %s", tc.expectedName)
		}

		dLower, isOkLower := applogger.ParseDriverType(strings.ToLower(tc.expectedName))
		if !isOkLower || dLower != tc.driver {
			t.Errorf("failed parsing lowercase %s", tc.expectedName)
		}
	}
}

func TestDriverType_ParseAliasesAndNumeric(t *testing.T) {
	for input, expected := range aliasParseCases {
		d, isOk := applogger.ParseDriverType(input)
		if !isOk || d != expected {
			t.Errorf("ParseDriverType(%q) = %v, %v; expected %v", input, d, isOk, expected)
		}
	}
}

func TestDriverType_ParseFallbackAndZero(t *testing.T) {
	_, isOk := applogger.ParseDriverType("unknown_invalid_driver")
	if isOk {
		t.Errorf("expected isOk=false for unknown driver")
	}

	if d := applogger.ParseDriverTypeOrZero("unknown_driver"); d != applogger.DriverConsole {
		t.Errorf("expected zero to be DriverConsole, got %v", d)
	}

	if d := applogger.ParseDriverTypeOrZero("Api"); d != applogger.DriverApi {
		t.Errorf("expected DriverApi, got %v", d)
	}
}

func TestDriverType_JSONMarshal(t *testing.T) {
	data, err := json.Marshal(applogger.DriverApi)
	if err != nil {
		t.Fatalf("failed to marshal DriverApi: %v", err)
	}

	if string(data) != `"Api"` {
		t.Errorf("expected `\"Api\"`, got %s", string(data))
	}

	invData, err := json.Marshal(applogger.DriverType(99))
	if err != nil {
		t.Fatalf("failed to marshal invalid driver: %v", err)
	}

	if string(invData) != `"Driver(99)"` {
		t.Errorf("expected `\"Driver(99)\"`, got %s", string(invData))
	}
}

func TestDriverType_JSONUnmarshalValid(t *testing.T) {
	var d1, d2, d3, d4 applogger.DriverType
	if err := json.Unmarshal([]byte(`"Api"`), &d1); err != nil || d1 != applogger.DriverApi {
		t.Errorf("failed unmarshaling `\"Api\"`: %v", err)
	}

	if err := json.Unmarshal([]byte(`"file"`), &d2); err != nil || d2 != applogger.DriverFile {
		t.Errorf("failed unmarshaling `\"file\"`: %v", err)
	}

	if err := json.Unmarshal([]byte(`6`), &d3); err != nil || d3 != applogger.DriverApi {
		t.Errorf("failed unmarshaling numeric 6: %v", err)
	}

	if err := json.Unmarshal([]byte(`null`), &d4); err != nil || d4 != applogger.DriverConsole {
		t.Errorf("failed unmarshaling null: %v", err)
	}
}

func TestDriverType_JSONUnmarshalInvalid(t *testing.T) {
	var d applogger.DriverType
	if err := json.Unmarshal([]byte(`"unknown_type"`), &d); err == nil {
		t.Errorf("expected error unmarshaling unknown string")
	}

	if err := json.Unmarshal([]byte(`99`), &d); err == nil {
		t.Errorf("expected error unmarshaling out-of-range number")
	}

	if err := json.Unmarshal([]byte(`{}`), &d); err == nil {
		t.Errorf("expected error unmarshaling invalid json object")
	}
}
