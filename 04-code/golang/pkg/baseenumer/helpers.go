package baseenumer

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// CompileMap builds a fast lookup map for an enumeration from a slice of variant labels.
// It maps the raw label, lowercase label, uppercase label, and numeric string index to the typed variant,
// as well as standard aliases ("unknown", "invalid", "UNKNOWN", "INVALID") to the invalid zero-variant.
func CompileMap[V IntNumber](labels []string, invalid V) map[string]V {
	m := make(map[string]V, (len(labels)*4)+4)
	for i, label := range labels {
		v := V(i)
		m[label] = v
		m[strings.ToLower(label)] = v
		m[strings.ToUpper(label)] = v
		m[strconv.Itoa(i)] = v
	}

	m["unknown"] = invalid
	m["invalid"] = invalid
	m["UNKNOWN"] = invalid
	m["INVALID"] = invalid

	return m
}

// SliceValues extracts all valid label strings excluding the 0-index invalid/unknown label.
func SliceValues(labels []string) []string {
	if len(labels) <= 1 {
		return []string{}
	}

	res := make([]string, 0, len(labels)-1)

	return append(res, labels[1:]...)
}

// SliceVariants extracts all valid typed variant values excluding index 0.
func SliceVariants[V IntNumber](labels []string) []V {
	if len(labels) <= 1 {
		return []V{}
	}

	items := make([]V, 0, len(labels)-1)
	for i := 1; i < len(labels); i++ {
		items = append(items, V(i))
	}

	return items
}

// FormatNameValue formats a name and an underlying value into the canonical format: "Name(Value)".
func FormatNameValue(name string, val any) string {
	return fmt.Sprintf("%s(%v)", name, val)
}

// IsBetween reports whether val satisfies min <= val && val <= max.
func IsBetween[T cmp.Ordered](val, min, max T) bool {
	return val >= min && val <= max
}

// IsNotBetween reports whether val is strictly less than min or strictly greater than max.
func IsNotBetween[T cmp.Ordered](val, min, max T) bool {
	return val < min || val > max
}

// ParseLookup looks up a string in a variant map after trimming whitespace and lowercasing.
// It returns the resolved variant, trimmed string, and whether the lookup succeeded.
func ParseLookup[V any](s string, variantMap map[string]V) (val V, trimmed string, ok bool) {
	trimmed = strings.TrimSpace(s)
	if len(trimmed) == 0 {
		var zero V

		return zero, "", false
	}

	v, found := variantMap[strings.ToLower(trimmed)]

	return v, trimmed, found
}

// FormatParseError creates the standardized error message when a variant is not recognized.
func FormatParseError(typeName, raw string, supportedVariants []string) string {
	return fmt.Sprintf("unknown %s variant %q, supported variants: [%s]", typeName, raw, strings.Join(supportedVariants, ", "))
}

// FormatEmptyParseError creates the standardized error message when input string is empty.
func FormatEmptyParseError(typeName string) string {
	return fmt.Sprintf("cannot parse empty string as %s", typeName)
}

// FormatNumericRangeError creates the standardized error message when a numeric enum value exceeds bounds.
func FormatNumericRangeError(typeName string, raw any, max int) error {
	return fmt.Errorf("invalid %s numeric value %v, supported range: 0..%d", typeName, raw, max)
}
