package errtype

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"coding-guidelines/common/pkg/baseenumer"
)

type (
	// BaseEnumer defines the interface for all type-safe enumerations (forwarded from baseenumer).
	BaseEnumer = baseenumer.BaseEnumer

	// BaseEnum is an alias for BaseEnumer.
	BaseEnum = baseenumer.BaseEnum

	// ByteEnumer defines the interface for byte-backed enumerations (forwarded from baseenumer).
	ByteEnumer = baseenumer.ByteEnumer

	// ByteEnum is an alias for ByteEnumer.
	ByteEnum = baseenumer.ByteEnum

	// UTF8Enumer defines the interface for UTF-8 byte-backed enumerations (forwarded from baseenumer).
	UTF8Enumer = baseenumer.UTF8Enumer

	// UTF8Enum is an alias for UTF8Enumer.
	UTF8Enum = baseenumer.UTF8Enum

	// UTF16Enumer defines the interface for UTF-16 code-unit-backed enumerations (forwarded from baseenumer).
	UTF16Enumer = baseenumer.UTF16Enumer

	// UTF16Enum is an alias for UTF16Enumer.
	UTF16Enum = baseenumer.UTF16Enum

	// UTF32Enumer defines the interface for UTF-32 / rune-backed enumerations (forwarded from baseenumer).
	UTF32Enumer = baseenumer.UTF32Enumer

	// UTF32Enum is an alias for UTF32Enumer.
	UTF32Enum = baseenumer.UTF32Enum

	// RuneEnumer defines the interface for rune-backed enumerations (forwarded from baseenumer).
	RuneEnumer = baseenumer.RuneEnumer

	// RuneEnum is an alias for RuneEnumer.
	RuneEnum = baseenumer.RuneEnum

	// StringEnumer defines the interface for string-backed enumerations (forwarded from baseenumer).
	StringEnumer = baseenumer.StringEnumer

	// StringEnum is an alias for StringEnumer.
	StringEnum = baseenumer.StringEnum

	// NumberEnumer defines the interface for numeric-backed enumerations (forwarded from baseenumer).
	NumberEnumer = baseenumer.NumberEnumer

	// NumberEnum is an alias for NumberEnumer.
	NumberEnum = baseenumer.NumberEnum

	// IntEnumer defines the interface for integer-backed enumerations (forwarded from baseenumer).
	IntEnumer = baseenumer.IntEnumer

	// IntEnum is an alias for IntEnumer.
	IntEnum = baseenumer.IntEnum
)

// ValueString returns the string representation of the variation code.
func (v Variation) ValueString() string {
	return fmt.Sprintf("%d", uint16(v))
}

// IsEnum returns true if the Variation is one of the recognized standard variations.
func (v Variation) IsEnum() bool {
	_, ok := variationNames[v]

	return ok
}

// ProcessStateType represents a string-backed enum conforming to BaseEnum.
type ProcessStateType string

var (
	processStateRegistry = map[ProcessStateType]bool{
		ProcessStatePending:   true,
		ProcessStateRunning:   true,
		ProcessStateCompleted: true,
		ProcessStateFailed:    true,
		ProcessStateCancelled: true,
	}

	processStateMap = compileProcessStateMap()
)

func compileProcessStateMap() map[string]ProcessStateType {
	states := AllProcessStates()
	m := make(map[string]ProcessStateType, len(states)*3)
	for _, state := range states {
		str := string(state)
		m[str] = state
		m[strings.ToLower(str)] = state
		m[strings.ToUpper(str)] = state
	}

	return m
}

// Name returns the identifier name.
func (s ProcessStateType) Name() string {
	return string(s)
}

// String returns the string representation.
func (s ProcessStateType) String() string {
	return string(s)
}

// ValueString returns the string representation of value.
func (s ProcessStateType) ValueString() string {
	return string(s)
}

// Value returns the raw string value.
func (s ProcessStateType) Value() string {
	return string(s)
}

// IsValid returns true if this state is non-empty and known.
func (s ProcessStateType) IsValid() bool {
	return processStateRegistry[s]
}

// IsEnum returns true if this state exists in the registry.
func (s ProcessStateType) IsEnum() bool {
	return processStateRegistry[s]
}

// IsCompare checks equality against another ProcessStateType.
func (s ProcessStateType) IsCompare(target ProcessStateType) bool {
	return s == target
}

// MarshalJSON implements json.Marshaler.
func (s ProcessStateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *ProcessStateType) UnmarshalJSON(data []byte) error {
	trimmed := strings.TrimSpace(string(data))
	if len(trimmed) == 0 || trimmed == "null" {
		*s = ProcessStateUnknown

		return nil
	}

	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	parsed := ParseProcessState(raw)
	if !parsed.IsValid() {
		names := make([]string, 0, len(AllProcessStates()))
		for _, st := range AllProcessStates() {
			names = append(names, string(st))
		}

		return fmt.Errorf("unknown ProcessStateType %q, supported: [%s]", raw, strings.Join(names, ", "))
	}

	*s = parsed

	return nil
}

// AllProcessStates returns all registered ProcessStateType values.
func AllProcessStates() []ProcessStateType {
	return []ProcessStateType{
		ProcessStatePending,
		ProcessStateRunning,
		ProcessStateCompleted,
		ProcessStateFailed,
		ProcessStateCancelled,
	}
}

// ParseProcessState parses a string into ProcessStateType case-insensitively.
func ParseProcessState(val string) ProcessStateType {
	if s, ok := processStateMap[strings.ToLower(strings.TrimSpace(val))]; ok {
		return s
	}

	return ProcessStateUnknown
}

// LogLevelType represents an integer-backed enum conforming to NumberEnum and BaseEnum.
type LogLevelType uint16

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

// ToEnum finds an enum by name in any slice of BaseEnumer (delegates to baseenumer.ToEnum).
func ToEnum[T BaseEnumer](val string, all []T) (T, bool) {
	return baseenumer.ToEnum(val, all)
}

var _ BaseEnumer = Variation(0)
var _ NumberEnumer = Variation(0)
var _ IntEnumer = Variation(0)
var _ BaseEnumer = ProcessStateType("")
var _ StringEnumer = ProcessStateType("")
var _ BaseEnumer = LogLevelType(0)
var _ NumberEnumer = LogLevelType(0)
var _ IntEnumer = LogLevelType(0)
