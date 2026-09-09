package applogger

import (
	"encoding/json"
	"fmt"
	"strings"
)

type DriverType byte

const (
	DriverConsole DriverType = iota
	DriverFile
	DriverSQLite
	DriverZap
	DriverComposite
	DriverRotatingFile
	DriverApi
	DriverJsonWriterLogger
	DriverStreamer
)

const (
	DriverFileWriter        = DriverFile
	DriverFileWriterRotator = DriverRotatingFile
	DriverSqliteDbWriter    = DriverSQLite
	DriverAPI               = DriverApi
	DriverJson              = DriverJsonWriterLogger
)

var driverNames = [...]string{
	"Console",
	"File",
	"SQLite",
	"Zap",
	"Composite",
	"RotatingFile",
	"Api",
	"JsonWriterLogger",
	"Streamer",
}

var driverLookup = map[string]DriverType{
	"console":           DriverConsole,
	"0":                 DriverConsole,
	"file":              DriverFile,
	"filewriter":        DriverFile,
	"1":                 DriverFile,
	"sqlite":            DriverSQLite,
	"sqlitedbwriter":    DriverSQLite,
	"2":                 DriverSQLite,
	"zap":               DriverZap,
	"3":                 DriverZap,
	"composite":         DriverComposite,
	"4":                 DriverComposite,
	"rotatingfile":      DriverRotatingFile,
	"filewriterrotator": DriverRotatingFile,
	"5":                 DriverRotatingFile,
	"api":               DriverApi,
	"6":                 DriverApi,
	"jsonwriterlogger":  DriverJsonWriterLogger,
	"json":              DriverJsonWriterLogger,
	"7":                 DriverJsonWriterLogger,
	"streamer":          DriverStreamer,
	"8":                 DriverStreamer,
}

func (d DriverType) Name() string {
	if int(d) < len(driverNames) {
		return driverNames[d]
	}

	return fmt.Sprintf("Driver(%d)", byte(d))
}

func (d DriverType) String() string {
	return d.Name()
}

func (d DriverType) IsValid() bool {
	return int(d) < len(driverNames)
}

func normalizeDriverName(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "_", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, " ", "")
	if strings.HasPrefix(s, "driver(") && strings.HasSuffix(s, ")") {
		return s[7 : len(s)-1]
	}

	return s
}

func ParseDriverType(s string) (DriverType, bool) {
	norm := normalizeDriverName(s)
	if d, isOk := driverLookup[norm]; isOk {
		return d, true
	}

	d, isOk := driverLookup[strings.TrimPrefix(norm, "driver")]

	return d, isOk
}

func ParseDriverTypeOrZero(s string) DriverType {
	d, isOk := ParseDriverType(s)
	if isOk {
		return d
	}

	return DriverConsole
}

func (d DriverType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *DriverType) UnmarshalJSON(data []byte) error {
	trimmed := strings.TrimSpace(string(data))
	if trimmed == "null" || len(trimmed) == 0 {
		*d = DriverConsole

		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		return d.unmarshalString(str)
	}

	return d.unmarshalNumeric(data)
}

func (d *DriverType) unmarshalString(str string) error {
	parsed, isOk := ParseDriverType(str)
	if isOk {
		*d = parsed

		return nil
	}

	return fmt.Errorf("unknown DriverType %q", str)
}

func (d *DriverType) unmarshalNumeric(data []byte) error {
	var num int
	if err := json.Unmarshal(data, &num); err != nil {
		return err
	}

	if num < 0 || num >= len(driverNames) {
		return fmt.Errorf("invalid DriverType numeric value %d", num)
	}

	*d = DriverType(num)

	return nil
}
