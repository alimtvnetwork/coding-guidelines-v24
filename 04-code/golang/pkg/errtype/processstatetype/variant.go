package processstatetype

import (
	"encoding/json"

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
	allVariants = []Variant{
		Pending,
		Running,
		Completed,
		Failed,
		Cancelled,
	}

	basicEnum = baseenumer.NewBasicString(allVariants, Unknown)
)

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
	switch s {
	case Pending, Running, Completed, Failed, Cancelled:
		return true
	default:
		return false
	}
}

func (s Variant) IsEnum() bool {
	return s.IsValid()
}

func (s Variant) IsCompare(target Variant) bool {
	return s == target
}

func (s Variant) MarshalJSON() ([]byte, error) {
	return baseenumer.MarshalJSON(string(s))
}

func (s *Variant) UnmarshalJSON(data []byte) error {
	return basicEnum.UnmarshalJSON(data, s)
}

func All() []Variant {
	return basicEnum.All()
}

func AllProcessStates() []Variant {
	return All()
}

func Values() []string {
	return basicEnum.Values()
}

func Parse(val string) Variant {
	v, _, isOk := basicEnum.ParseLookup(val)
	if isOk {
		return v
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
