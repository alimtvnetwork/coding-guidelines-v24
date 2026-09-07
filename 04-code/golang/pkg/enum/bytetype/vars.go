package bytetype

import (
	"strconv"
	"strings"

	"coding-guidelines/common/pkg/baseenumer"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

var (
	variantLabels = [...]string{
		Zero:  "Zero",
		One:   "One",
		Two:   "Two",
		Three: "Three",
	}

	allVariants = []Variant{Zero, One, Two, Three, Max}

	allValues = []string{"Zero", "One", "Two", "Three", "Max"}

	variantMap = compileVariantMap()
)

func compileVariantMap() map[string]Variant {
	m := make(map[string]Variant, 32)
	populateStandardVariants(m)
	populateAliases(m)

	return m
}

func populateStandardVariants(m map[string]Variant) {
	for i, label := range variantLabels {
		v := Variant(i)
		m[label] = v
		m[strings.ToLower(label)] = v
		m[strings.ToUpper(label)] = v
		m[strconv.Itoa(i)] = v
	}
}

func populateAliases(m map[string]Variant) {
	m["max"] = Max
	m["MAX"] = Max
	m["Max"] = Max
	m["255"] = Max
	m["min"] = Min
	m["MIN"] = Min
	m["unknown"] = Unknown
	m["invalid"] = Invalid
}

func All() []Variant {
	return append([]Variant(nil), allVariants...)
}

func Values() []string {
	return append([]string(nil), allValues...)
}

func Parse(s string) result.Wrap[Variant] {
	v, trimmed, ok := baseenumer.ParseLookup(s, variantMap)
	if ok {
		return result.WrapSuccess(v)
	}

	if len(trimmed) == 0 {
		return result.WrapFailureWithId[Variant](errtype.Validation, baseenumer.FormatEmptyParseError("bytetype"))
	}

	parsed, err := strconv.ParseUint(trimmed, 10, 8)
	if err != nil {
		return result.WrapFailureWithId[Variant](
			errtype.NotFound,
			baseenumer.FormatParseError("bytetype", s, Values()),
		)
	}

	return result.WrapSuccess(Variant(parsed))
}
