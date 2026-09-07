package fileutil

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"

	"gopkg.in/yaml.v3"

	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

// ReadText reads the entire file content as a string.
func ReadText(path string) StringResult {
	fRes := OpenFile(path, FileOpenReadOnly, FilePermStandard)
	if fRes.HasError() {
		return result.WrapFailure[string](fRes.Fault())
	}

	defer fRes.Data().Close()

	content, err := io.ReadAll(fRes.Data())
	if err != nil {
		return StringFailure(errtype.IO, err, path, "failed to read file content")
	}

	return StringSuccess(string(content))
}

func scanLines(r io.Reader, path string) LinesResult {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return LinesFailure(errtype.IO, err, path, "error scanning file lines")
	}

	return LinesSuccess(lines)
}

// ReadLines reads the file and splits it into a string array by lines.
func ReadLines(path string) LinesResult {
	fRes := OpenFile(path, FileOpenReadOnly, FilePermStandard)
	if fRes.HasError() {
		return result.WrapFailure[[]string](fRes.Fault())
	}

	defer fRes.Data().Close()

	return scanLines(fRes.Data(), path)
}

// ReadJson parses a JSON file into the specified type T.
func ReadJson[T any](path string) result.Wrap[T] {
	var val T
	fRes := OpenFile(path, FileOpenReadOnly, FilePermStandard)
	if fRes.HasError() {
		return result.WrapFailure[T](fRes.Fault())
	}

	defer fRes.Data().Close()

	if err := json.NewDecoder(fRes.Data()).Decode(&val); err != nil {
		return result.WrapFailureFile[T](errtype.Serialization, err, path, "failed to decode JSON")
	}

	return result.WrapSuccess(val)
}

// ReadJSON is an alias for ReadJson.
func ReadJSON[T any](path string) result.Wrap[T] {
	return ReadJson[T](path)
}

func unmarshalYaml[T any](f io.Reader, path string) result.Wrap[T] {
	var val T
	content, err := io.ReadAll(f)
	if err != nil {
		return result.WrapFailureFile[T](errtype.IO, err, path, "failed to read file for YAML decoding")
	}

	content = bytes.TrimPrefix(content, []byte("\xef\xbb\xbf"))
	if err := yaml.Unmarshal(content, &val); err != nil {
		return result.WrapFailureFile[T](errtype.Serialization, err, path, "failed to decode YAML")
	}

	return result.WrapSuccess(val)
}

// ReadYaml parses a YAML file into the specified type T.
func ReadYaml[T any](path string) result.Wrap[T] {
	fRes := OpenFile(path, FileOpenReadOnly, FilePermStandard)
	if fRes.HasError() {
		return result.WrapFailure[T](fRes.Fault())
	}

	defer fRes.Data().Close()

	return unmarshalYaml[T](fRes.Data(), path)
}

// ReadYAML is an alias for ReadYaml.
func ReadYAML[T any](path string) result.Wrap[T] {
	return ReadYaml[T](path)
}
