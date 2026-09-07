package appfault

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"coding-guidelines/common/pkg/errtype"

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
func (cm ContextMap) ToJSON() Result[[]byte] {
	b, err := json.MarshalIndent(cm, "", "  ")
	if err != nil {
		return FailureResult[[]byte](Wrap(errtype.Internal, err, "failed to serialize context map to JSON"))
	}

	return SuccessResult(b)
}

// ToJSONString exports the ContextMap as a JSON string.
func (cm ContextMap) ToJSONString() string {
	res := cm.ToJSON()
	if res.Fault() != nil {
		return "{}"
	}

	return string(res.Data())
}

// ToYaml exports the ContextMap as YAML bytes.
func (cm ContextMap) ToYaml() Result[[]byte] {
	b, err := yaml.Marshal(cm)
	if err != nil {
		return FailureResult[[]byte](Wrap(errtype.Internal, err, "failed to serialize context map to YAML"))
	}

	return SuccessResult(b)
}

// ToYAML is an alias for ToYaml.
func (cm ContextMap) ToYAML() Result[[]byte] {
	return cm.ToYaml()
}

// ToYamlString exports the ContextMap as a YAML string.
func (cm ContextMap) ToYamlString() string {
	res := cm.ToYaml()
	if res.Fault() != nil {
		return ""
	}

	return string(res.Data())
}

// ToYAMLString is an alias for ToYamlString.
func (cm ContextMap) ToYAMLString() string {
	return cm.ToYamlString()
}

// ContextMapFromJSON parses a ContextMap from JSON bytes.
func ContextMapFromJSON(data []byte) Result[ContextMap] {
	var cm ContextMap
	if err := json.Unmarshal(data, &cm); err != nil {
		return FailureResult[ContextMap](Wrap(errtype.Internal, err, "failed to parse context map from JSON"))
	}

	return SuccessResult(cm)
}

// ContextMapFromJSONString parses a ContextMap from a JSON string.
func ContextMapFromJSONString(s string) Result[ContextMap] {
	return ContextMapFromJSON([]byte(s))
}

// ContextMapFromYaml parses a ContextMap from YAML bytes.
func ContextMapFromYaml(data []byte) Result[ContextMap] {
	var cm ContextMap
	if err := yaml.Unmarshal(data, &cm); err != nil {
		return FailureResult[ContextMap](Wrap(errtype.Internal, err, "failed to parse context map from YAML"))
	}

	return SuccessResult(cm)
}

// ContextMapFromYAML is an alias for ContextMapFromYaml.
func ContextMapFromYAML(data []byte) Result[ContextMap] {
	return ContextMapFromYaml(data)
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
