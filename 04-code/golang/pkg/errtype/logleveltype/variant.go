package logleveltype

import (
	"encoding/json"
	"fmt"
	"strings"

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
	logLevelNames = map[Variant]string{
		Debug: "Debug",
		Info:  "Info",
		Warn:  "Warn",
		Error: "Error",
		Fatal: "Fatal",
	}

	variantMap = compileVariantMap()
)

func compileVariantMap() map[string]Variant {
	m := make(map[string]Variant, len(logLevelNames)*4)
	for lvl, name := range logLevelNames {
		m[name] = lvl
		m[strings.ToLower(name)] = lvl
		m[strings.ToUpper(name)] = lvl
		m[fmt.Sprintf("%d", uint16(lvl))] = lvl
	}

	return m
}

func (l Variant) Name() string {
	if name, ok := logLevelNames[l]; ok {
		return name
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

func (l Variant) IsValid() bool {
	_, ok := logLevelNames[l]

	return ok
}

func (l Variant) IsEnum() bool {
	_, ok := logLevelNames[l]

	return ok
}

func (l Variant) IsCompare(target Variant) bool {
	return l == target
}

func (l Variant) MarshalJSON() ([]byte, error) {
	return baseenumer.MarshalJSON(l.Name())
}

func (l *Variant) UnmarshalJSON(data []byte) error {
	return baseenumer.UnmarshalIntegerJSON(data, l, "logleveltype", variantMap, 5, 0)
}

func All() []Variant {
	return []Variant{Debug, Info, Warn, Error, Fatal}
}

func AllLogLevels() []Variant {
	return All()
}

func Parse(val string) Variant {
	cleaned := strings.ToLower(strings.TrimSpace(val))
	if lvl, ok := variantMap[cleaned]; ok {
		return lvl
	}

	return 0
}

func ParseLogLevel(val string) Variant {
	return Parse(val)
}

var (
	_ baseenumer.BaseEnumer   = Variant(0)
	_ baseenumer.NumberEnumer = Variant(0)
	_ json.Marshaler          = Variant(0)
	_ json.Unmarshaler        = (*Variant)(nil)
)
