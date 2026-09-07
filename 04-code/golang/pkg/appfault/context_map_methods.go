package appfault

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// IsDefined reports true if the ContextMap is non-nil and has at least one entry.
func (cm ContextMap) IsDefined() bool {
	return len(cm) > 0
}

// IsEmpty reports true if the ContextMap is nil or has zero entries.
func (cm ContextMap) IsEmpty() bool {
	return len(cm) == 0
}

// Clone creates a shallow copy of the ContextMap.
func (cm ContextMap) Clone() ContextMap {
	if cm == nil {
		return nil
	}

	cloned := make(ContextMap, len(cm))
	for k, v := range cm {
		cloned[k] = v
	}

	return cloned
}

// Merge copies all entries from other into cm.
func (cm ContextMap) Merge(other ContextMap) ContextMap {
	for k, v := range other {
		cm[k] = v
	}

	return cm
}

// Keys returns a sorted slice of keys.
func (cm ContextMap) Keys() []string {
	keys := make([]string, 0, len(cm))
	for k := range cm {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

// NonEmptyKeys returns a sorted slice of keys that have non-empty string or non-nil values.
func (cm ContextMap) NonEmptyKeys() []string {
	var keys []string
	for k, v := range cm {
		if k == "" || v == nil {
			continue
		}

		if str, ok := v.(string); ok && str == "" {
			continue
		}

		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

// Values returns all values from the ContextMap in sorted key order.
func (cm ContextMap) Values() []any {
	keys := cm.Keys()
	values := make([]any, 0, len(keys))
	for _, k := range keys {
		values = append(values, cm[k])
	}

	return values
}

// ToMap returns a standard Go map[string]any shallow copy.
func (cm ContextMap) ToMap() map[string]any {
	result := make(map[string]any, len(cm))
	for k, v := range cm {
		result[k] = v
	}

	return result
}

// ToJSON exports the ContextMap as indented JSON bytes.
func (cm ContextMap) ToJSON() ([]byte, error) {
	return json.MarshalIndent(cm, "", "  ")
}

// ToJSONString exports the ContextMap as a JSON string.
func (cm ContextMap) ToJSONString() string {
	b, err := cm.ToJSON()
	if err != nil {
		return "{}"
	}

	return string(b)
}

// ToYAML exports the ContextMap as YAML bytes.
func (cm ContextMap) ToYAML() ([]byte, error) {
	return yaml.Marshal(cm)
}

// ToYAMLString exports the ContextMap as a YAML string.
func (cm ContextMap) ToYAMLString() string {
	b, err := cm.ToYAML()
	if err != nil {
		return ""
	}

	return string(b)
}

// ContextMapFromJSON parses a ContextMap from JSON bytes.
func ContextMapFromJSON(data []byte) (ContextMap, error) {
	var cm ContextMap
	if err := json.Unmarshal(data, &cm); err != nil {
		return nil, err
	}

	return cm, nil
}

// ContextMapFromJSONString parses a ContextMap from a JSON string.
func ContextMapFromJSONString(s string) (ContextMap, error) {
	return ContextMapFromJSON([]byte(s))
}

// ContextMapFromYAML parses a ContextMap from YAML bytes.
func ContextMapFromYAML(data []byte) (ContextMap, error) {
	var cm ContextMap
	if err := yaml.Unmarshal(data, &cm); err != nil {
		return nil, err
	}

	return cm, nil
}

// Format formats the map as a human-readable comma-separated string.
func (cm ContextMap) Format() string {
	if cm.IsEmpty() {
		return "{}"
	}

	keys := cm.Keys()
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%v", k, cm[k]))
	}

	return "{" + strings.Join(parts, ", ") + "}"
}
