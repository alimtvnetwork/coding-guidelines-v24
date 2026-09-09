package appfault

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// String returns a human-readable string representation of the value or error.
func (r Result[T]) String() string {
	if r.IsFailed() {
		return r.appError.FormatStdout()
	}

	return stringifyValue(r.value)
}

// PrettyJson returns the payload or error formatted as indented JSON string.
func (r Result[T]) PrettyJson() string {
	if r.IsFailed() {
		return r.appError.FormatJson()
	}

	bytes, err := json.MarshalIndent(r.value, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(bytes)
}

// ToJsonPretty is an alias for PrettyJson.
func (r Result[T]) ToJsonPretty() string {
	return r.PrettyJson()
}

// PrettyMap returns the payload converted to a map and formatted as indented JSON.
func (r Result[T]) PrettyMap() string {
	m := r.ToMap()
	if m == nil {
		return "{}"
	}

	bytes, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(bytes)
}

// ToYaml serializes the result value (or error) to YAML string.
func (r Result[T]) ToYaml() (string, error) {
	var target any = r.value
	if r.IsFailed() {
		target = r.appError
	}

	bytes, err := yaml.Marshal(target)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Yaml returns the YAML representation of the payload, falling back to empty on error.
func (r Result[T]) Yaml() string {
	res, err := r.ToYaml()
	if err != nil {
		return ""
	}

	return res
}

// PrintConsole prints the String() representation to stdout and returns r for chaining.
func (r Result[T]) PrintConsole() Result[T] {
	fmt.Println(r.String())

	return r
}
