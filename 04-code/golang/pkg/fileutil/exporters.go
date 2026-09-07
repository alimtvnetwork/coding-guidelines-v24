package fileutil

import (
	"encoding/json"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

// ExportText writes a string to a file, overwriting if it exists.
func ExportText(path string, content string, perm FilePermType) BoolResult {
	fRes := CreateFile(path, perm)
	if fRes.HasError() {
		return result.WrapFailure[bool](fRes.Fault())
	}

	defer fRes.Data().Close()

	if _, err := fRes.Data().WriteString(content); err != nil {
		return BoolFailure(errtype.IO, err, path, "failed to write text to file")
	}

	return BoolSuccess(true)
}

// ExportLines writes an array of strings to a file, separated by newlines.
func ExportLines(path string, lines []string, perm FilePermType) BoolResult {
	if len(lines) == 0 {
		return ExportText(path, "", perm)
	}

	return ExportText(path, strings.Join(lines, "\n")+"\n", perm)
}

func encodeJsonToFile(f *os.File, data any, path string) BoolResult {
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return BoolFailure(errtype.Serialization, err, path, "failed to encode JSON")
	}

	return BoolSuccess(true)
}

// ExportJson writes a data structure to a file as formatted JSON.
func ExportJson(path string, data any, perm FilePermType) BoolResult {
	fRes := CreateFile(path, perm)
	if fRes.HasError() {
		return result.WrapFailure[bool](fRes.Fault())
	}

	return encodeJsonToFile(fRes.Data(), data, path)
}

// ExportJSON is an alias for ExportJson.
func ExportJSON(path string, data any, perm FilePermType) BoolResult {
	return ExportJson(path, data, perm)
}

func encodeYamlToFile(f *os.File, data any, path string) BoolResult {
	defer f.Close()

	encoder := yaml.NewEncoder(f)
	defer encoder.Close()

	encoder.SetIndent(2)
	if err := encoder.Encode(data); err != nil {
		return BoolFailure(errtype.Serialization, err, path, "failed to encode YAML")
	}

	return BoolSuccess(true)
}

// ExportYaml writes a data structure to a file as YAML.
func ExportYaml(path string, data any, perm FilePermType) BoolResult {
	fRes := CreateFile(path, perm)
	if fRes.HasError() {
		return result.WrapFailure[bool](fRes.Fault())
	}

	return encodeYamlToFile(fRes.Data(), data, path)
}

// ExportYAML is an alias for ExportYaml.
func ExportYAML(path string, data any, perm FilePermType) BoolResult {
	return ExportYaml(path, data, perm)
}
