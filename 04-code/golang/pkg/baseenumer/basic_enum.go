package baseenumer

import (
	"errors"
	"fmt"
	"strings"
)

// BasicIntegerEnum encapsulates reusable enum operations for integer-backed enums.
type BasicIntegerEnum[V IntNumber] struct {
	labels     []string
	variants   []V
	variantMap map[string]V
	maxValid   int
	zero       V
	typeName   string
}

// NewBasicInteger initializes a BasicIntegerEnum from label slice and zero-value variant.
func NewBasicInteger[V IntNumber](labels []string, zero V) *BasicIntegerEnum[V] {
	vMap := CompileMap(labels, zero)
	maxVal := len(labels) - 1
	var dummy V
	tName := ResolveTypeName(&dummy)

	return &BasicIntegerEnum[V]{
		labels:     labels,
		variantMap: vMap,
		maxValid:   maxVal,
		zero:       zero,
		typeName:   tName,
	}
}

// NewBasicSparseInteger initializes a BasicIntegerEnum for sparse non-contiguous enums.
func NewBasicSparseInteger[V IntNumber](
	variants []V,
	names []string,
	zero V,
	maxValid int,
) *BasicIntegerEnum[V] {
	vMap := CompileSparseIntegerMap(variants, names, zero)
	var dummy V
	tName := ResolveTypeName(&dummy)

	return &BasicIntegerEnum[V]{
		labels:     names,
		variants:   variants,
		variantMap: vMap,
		maxValid:   maxValid,
		zero:       zero,
		typeName:   tName,
	}
}

// UnmarshalJSON unmarshals JSON data into target using encapsulated metadata.
func (b *BasicIntegerEnum[V]) UnmarshalJSON(data []byte, target *V) error {
	return UnmarshalIntegerJSONWithName(data, target, b.typeName, b.variantMap, b.maxValid, b.zero)
}

// All returns all valid enum variants.
func (b *BasicIntegerEnum[V]) All() []V {
	if len(b.variants) > 0 {
		res := make([]V, len(b.variants))
		copy(res, b.variants)

		return res
	}

	return SliceVariants[V](b.labels)
}

// Values returns all valid string labels (excluding index 0).
func (b *BasicIntegerEnum[V]) Values() []string {
	if len(b.variants) > 0 {
		res := make([]string, len(b.labels))
		copy(res, b.labels)

		return res
	}

	return SliceValues(b.labels)
}

// Parse parses a string into an enum variant, returning an error if unrecognized.
func (b *BasicIntegerEnum[V]) Parse(s string) (V, error) {
	v, trimmed, isOk := ParseLookup(s, b.variantMap)
	if len(trimmed) == 0 {
		return b.zero, errors.New(FormatEmptyParseError(b.typeName))
	}

	if isOk {
		return v, nil
	}

	return b.zero, errors.New(FormatParseError(b.typeName, s, b.Values()))
}

// ParseLookup looks up string in variant map.
func (b *BasicIntegerEnum[V]) ParseLookup(s string) (V, string, bool) {
	return ParseLookup(s, b.variantMap)
}

// Map returns the internal variant map.
func (b *BasicIntegerEnum[V]) Map() map[string]V {
	return b.variantMap
}

// TypeName returns the resolved type name.
func (b *BasicIntegerEnum[V]) TypeName() string {
	return b.typeName
}

// Zero returns the zero-value variant.
func (b *BasicIntegerEnum[V]) Zero() V {
	return b.zero
}

// MaxValid returns the maximum valid index.
func (b *BasicIntegerEnum[V]) MaxValid() int {
	return b.maxValid
}

// CompileSparseIntegerMap builds a fast lookup map for sparse integer enum variants.
func CompileSparseIntegerMap[V IntNumber](variants []V, names []string, zero V) map[string]V {
	count := len(variants)
	if len(names) < count {
		count = len(names)
	}

	m := make(map[string]V, (count*4)+6)
	populateSparseEntries(m, variants[:count], names[:count])
	populateSparseAliases(m, zero)

	return m
}

func populateSparseEntries[V IntNumber](m map[string]V, variants []V, names []string) {
	for i, v := range variants {
		name := names[i]
		m[name] = v
		m[strings.ToLower(name)] = v
		m[strings.ToUpper(name)] = v
		m[fmt.Sprintf("%d", v)] = v
	}
}

func populateSparseAliases[V any](m map[string]V, zero V) {
	m["unknown"] = zero
	m["invalid"] = zero
	m["UNKNOWN"] = zero
	m["INVALID"] = zero
	m["min"] = zero
	m["MIN"] = zero
}

// CompileStringMap builds a fast lookup map for string enum variants.
func CompileStringMap[V ~string](variants []V, zero V) map[string]V {
	m := make(map[string]V, len(variants)*3+4)
	for _, v := range variants {
		str := string(v)
		if len(str) > 0 {
			m[str] = v
			m[strings.ToLower(str)] = v
			m[strings.ToUpper(str)] = v
		}
	}

	m["unknown"] = zero
	m["invalid"] = zero
	m["UNKNOWN"] = zero
	m["INVALID"] = zero

	return m
}

// BasicStringEnum encapsulates reusable enum operations for string-backed enums.
type BasicStringEnum[V ~string] struct {
	variants   []V
	variantMap map[string]V
	zero       V
	typeName   string
}

// NewBasicString initializes a BasicStringEnum from variant slice and zero-value variant.
func NewBasicString[V ~string](variants []V, zero V) *BasicStringEnum[V] {
	vMap := CompileStringMap(variants, zero)
	var dummy V
	tName := ResolveTypeName(&dummy)

	return &BasicStringEnum[V]{
		variants:   variants,
		variantMap: vMap,
		zero:       zero,
		typeName:   tName,
	}
}

// UnmarshalJSON unmarshals JSON data into target using encapsulated metadata.
func (b *BasicStringEnum[V]) UnmarshalJSON(data []byte, target *V) error {
	return UnmarshalStringJSONWithName(data, target, b.typeName, b.variantMap, b.zero)
}

// All returns all valid enum variants.
func (b *BasicStringEnum[V]) All() []V {
	res := make([]V, len(b.variants))
	copy(res, b.variants)

	return res
}

// Values returns all valid string labels.
func (b *BasicStringEnum[V]) Values() []string {
	res := make([]string, 0, len(b.variants))
	for _, v := range b.variants {
		if string(v) != "" && v != b.zero {
			res = append(res, string(v))
		}
	}

	return res
}

// Parse parses a string into an enum variant, returning an error if unrecognized.
func (b *BasicStringEnum[V]) Parse(s string) (V, error) {
	v, trimmed, isOk := ParseLookup(s, b.variantMap)
	if len(trimmed) == 0 {
		return b.zero, errors.New(FormatEmptyParseError(b.typeName))
	}

	if isOk {
		return v, nil
	}

	return b.zero, errors.New(FormatParseError(b.typeName, s, b.Values()))
}

// ParseLookup looks up string in variant map.
func (b *BasicStringEnum[V]) ParseLookup(s string) (V, string, bool) {
	return ParseLookup(s, b.variantMap)
}

// Map returns the internal variant map.
func (b *BasicStringEnum[V]) Map() map[string]V {
	return b.variantMap
}

// TypeName returns the resolved type name.
func (b *BasicStringEnum[V]) TypeName() string {
	return b.typeName
}

// Zero returns the zero-value variant.
func (b *BasicStringEnum[V]) Zero() V {
	return b.zero
}
