package processstatetype

import (
	"encoding/json"
	"strings"

	"coding-guidelines/common/pkg/baseenumer"
)

type (
	Variant string

	ProcessStateType = Variant
)

const (
	Pending   Variant = "Pending"
	Running   Variant = "Running"
	Completed Variant = "Completed"
	Failed    Variant = "Failed"
	Cancelled Variant = "Cancelled"
	Unknown   Variant = "Unknown"
)

const (
	ProcessStatePending   = Pending
	ProcessStateRunning   = Running
	ProcessStateCompleted = Completed
	ProcessStateFailed    = Failed
	ProcessStateCancelled = Cancelled
	ProcessStateUnknown   = Unknown
)

var (
	processStateRegistry = map[Variant]bool{
		Pending:   true,
		Running:   true,
		Completed: true,
		Failed:    true,
		Cancelled: true,
	}

	processStateMap = compileProcessStateMap()
)

func compileProcessStateMap() map[string]Variant {
	states := All()
	m := make(map[string]Variant, len(states)*3)
	for _, state := range states {
		str := string(state)
		m[str] = state
		m[strings.ToLower(str)] = state
		m[strings.ToUpper(str)] = state
	}

	return m
}

func (s Variant) Name() string {
	return string(s)
}

func (s Variant) String() string {
	return string(s)
}

func (s Variant) ValueString() string {
	return string(s)
}

func (s Variant) Value() string {
	return string(s)
}

func (s Variant) IsValid() bool {
	return processStateRegistry[s]
}

func (s Variant) IsEnum() bool {
	return processStateRegistry[s]
}

func (s Variant) IsCompare(target Variant) bool {
	return s == target
}

func (s Variant) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

func (s *Variant) UnmarshalJSON(data []byte) error {
	trimmed := strings.TrimSpace(string(data))
	if len(trimmed) == 0 || trimmed == "null" {
		*s = Unknown

		return nil
	}

	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	parsed := Parse(raw)
	*s = parsed

	return nil
}

func All() []Variant {
	return []Variant{
		Pending,
		Running,
		Completed,
		Failed,
		Cancelled,
	}
}

func AllProcessStates() []Variant {
	return All()
}

func Parse(val string) Variant {
	cleaned := strings.ToLower(strings.TrimSpace(val))
	if state, ok := processStateMap[cleaned]; ok {
		return state
	}

	return Unknown
}

func ParseProcessState(val string) Variant {
	return Parse(val)
}

var (
	_ baseenumer.BaseEnumer   = Variant("")
	_ baseenumer.StringEnumer = Variant("")
	_ json.Marshaler          = Variant("")
	_ json.Unmarshaler        = (*Variant)(nil)
)
