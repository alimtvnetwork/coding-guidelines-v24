package logleveltype

import (
	"encoding/json"
	"fmt"

	"coding-guidelines/common/pkg/baseenumer"
)

type (
	Variant uint16

	LogLevelType = Variant
)

const (
	Debug Variant = 1
	Info  Variant = 2
	Warn  Variant = 3
	Error Variant = 4
	Fatal Variant = 5
)

var (
	variantLabels = [...]string{
		0:     "Unknown",
		Debug: "Debug",
		Info:  "Info",
		Warn:  "Warn",
		Error: "Error",
		Fatal: "Fatal",
	}

	basicEnum = baseenumer.NewBasicInteger(variantLabels[:], Variant(0))
)

func (l Variant) Name() string {
	if int(l) >= int(Debug) && int(l) <= int(Fatal) {
		return variantLabels[l]
	}

	return fmt.Sprintf("LogLevel(%d)", uint16(l))
}

func (l Variant) String() string {
	return l.Name()
}

func (l Variant) ValueString() string {
	return fmt.Sprintf("%d", uint16(l))
}

func (l Variant) Code() uint16 {
	return uint16(l)
}

func (l Variant) Int() int {
	return int(l)
}

func (l Variant) Value() uint16 {
	return uint16(l)
}

func (l Variant) All() []Variant {
	return All()
}

func (l Variant) Values() []string {
	return Values()
}

func (l Variant) IsValid() bool {
	return l >= Debug && l <= Fatal
}

func (l Variant) IsEnum() bool {
	return l.IsValid()
}

func (l Variant) IsCompare(target Variant) bool {
	return l == target
}

func (l Variant) Min() Variant {
	return basicEnum.Min()
}

func (l Variant) Max() Variant {
	return basicEnum.Max()
}

func (l Variant) IsMin() bool {
	return basicEnum.IsMin(l)
}

func (l Variant) IsMax() bool {
	return basicEnum.IsMax(l)
}

func (l Variant) IsInRange(min, max Variant) bool {
	return baseenumer.IsBetween(l, min, max)
}

func (l Variant) MarshalJSON() ([]byte, error) {
	return baseenumer.MarshalJSON(l.Name())
}

func (l *Variant) UnmarshalJSON(data []byte) error {
	return basicEnum.UnmarshalJSON(data, l)
}

func All() []Variant {
	return basicEnum.All()
}

func AllLogLevels() []Variant {
	return All()
}

func Values() []string {
	return basicEnum.Values()
}

func Min() Variant {
	return basicEnum.Min()
}

func Max() Variant {
	return basicEnum.Max()
}

func Parse(val string) Variant {
	v, _, isOk := basicEnum.ParseLookup(val)
	if isOk {
		return v
	}

	return 0
}

func ParseLogLevel(val string) Variant {
	return Parse(val)
}

var (
	_ baseenumer.BaseEnumer             = Variant(0)
	_ baseenumer.NumberEnumer           = Variant(0)
	_ baseenumer.BoundedEnumer[Variant] = Variant(0)
	_ json.Marshaler                    = Variant(0)
	_ json.Unmarshaler                  = (*Variant)(nil)
)
