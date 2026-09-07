package errtype

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// LogLevelType represents an integer-backed enum conforming to NumberEnum and BaseEnum.
type LogLevelType uint16

// LogLevelType constants conforming to NumberEnum and BaseEnum.
const (
	LogLevelDebug LogLevelType = 1
	LogLevelInfo  LogLevelType = 2
	LogLevelWarn  LogLevelType = 3
	LogLevelError LogLevelType = 4
	LogLevelFatal LogLevelType = 5
)

var (
	logLevelNames = map[LogLevelType]string{
		LogLevelDebug: "Debug",
		LogLevelInfo:  "Info",
		LogLevelWarn:  "Warn",
		LogLevelError: "Error",
		LogLevelFatal: "Fatal",
	}

	errtypeLogLevelMap = compileErrtypeLogLevelMap()
)

func compileErrtypeLogLevelMap() map[string]LogLevelType {
	m := make(map[string]LogLevelType, len(logLevelNames)*4)
	for lvl, name := range logLevelNames {
		m[name] = lvl
		m[strings.ToLower(name)] = lvl
		m[strings.ToUpper(name)] = lvl
		m[fmt.Sprintf("%d", uint16(lvl))] = lvl
	}

	return m
}

// Name returns the uppercase identifier.
func (l LogLevelType) Name() string {
	if name, ok := logLevelNames[l]; ok {
		return name
	}

	return fmt.Sprintf("LogLevel(%d)", uint16(l))
}

// String implements fmt.Stringer.
func (l LogLevelType) String() string {
	return l.Name()
}

// ValueString returns the integer code formatted as a string.
func (l LogLevelType) ValueString() string {
	return fmt.Sprintf("%d", uint16(l))
}

// Code returns the raw uint16 code value.
func (l LogLevelType) Code() uint16 {
	return uint16(l)
}

// Int returns the int representation.
func (l LogLevelType) Int() int {
	return int(l)
}

// IsValid returns true if this log level is known.
func (l LogLevelType) IsValid() bool {
	_, ok := logLevelNames[l]

	return ok
}

// IsEnum returns true if this log level exists in registry.
func (l LogLevelType) IsEnum() bool {
	_, ok := logLevelNames[l]

	return ok
}

// IsCompare checks equality against another LogLevelType.
func (l LogLevelType) IsCompare(target LogLevelType) bool {
	return l == target
}

// MarshalJSON implements json.Marshaler.
func (l LogLevelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.Name())
}

// UnmarshalJSON implements json.Unmarshaler.
func (l *LogLevelType) UnmarshalJSON(data []byte) error {
	trimmed := strings.TrimSpace(string(data))
	if len(trimmed) == 0 || trimmed == "null" {
		*l = LogLevelType(0)

		return nil
	}

	var raw string
	if err := json.Unmarshal(data, &raw); err == nil {
		parsed := ParseLogLevel(raw)
		if !parsed.IsValid() {
			names := make([]string, 0, len(logLevelNames))
			for _, name := range logLevelNames {
				names = append(names, name)
			}

			sort.Strings(names)

			return fmt.Errorf("unknown LogLevelType %q, supported: [%s]", raw, strings.Join(names, ", "))
		}

		*l = parsed

		return nil
	}

	var code uint16
	if err := json.Unmarshal(data, &code); err != nil {
		return err
	}

	candidate := LogLevelType(code)
	if !candidate.IsValid() {
		return fmt.Errorf("invalid LogLevelType numeric code %d", code)
	}

	*l = candidate

	return nil
}

// AllLogLevels returns all registered LogLevelType values.
func AllLogLevels() []LogLevelType {
	return []LogLevelType{
		LogLevelDebug,
		LogLevelInfo,
		LogLevelWarn,
		LogLevelError,
		LogLevelFatal,
	}
}

// ParseLogLevel parses a string into LogLevelType case-insensitively.
func ParseLogLevel(val string) LogLevelType {
	cleaned := strings.ToLower(strings.TrimSpace(val))
	if lvl, ok := errtypeLogLevelMap[cleaned]; ok {
		return lvl
	}

	return 0
}

var (
	_ BaseEnumer   = LogLevelType(0)
	_ NumberEnumer = LogLevelType(0)
	_ IntEnumer    = LogLevelType(0)
)
