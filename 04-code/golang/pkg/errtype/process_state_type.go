package errtype

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ProcessStateType represents a string-backed enum conforming to BaseEnum.
type ProcessStateType string

// ProcessStateType constants conforming to BaseEnum.
const (
	ProcessStatePending   ProcessStateType = "Pending"
	ProcessStateRunning   ProcessStateType = "Running"
	ProcessStateCompleted ProcessStateType = "Completed"
	ProcessStateFailed    ProcessStateType = "Failed"
	ProcessStateCancelled ProcessStateType = "Cancelled"
	ProcessStateUnknown   ProcessStateType = "Unknown"
)

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

var (
	_ BaseEnumer   = ProcessStateType("")
	_ StringEnumer = ProcessStateType("")
)
