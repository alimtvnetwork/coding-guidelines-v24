package fileutil

import (
	"os"

	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/payloadconv"
	"coding-guidelines/common/pkg/result"
)

// Write writes any generic payload (struct, map, array, string, bytes) to a file.
// It automatically converts the payload via payloadconv.ToBytes.
func Write(path string, payload any, perm FilePermType) BoolResult {
	convRes := payloadconv.ToBytes(payload)
	if convRes.IsFailure() {
		return result.WrapFailure[bool](convRes.Fault())
	}

	return WriteBytes(path, convRes.Data(), perm)
}

// WriteLocked writes any generic payload with an exclusive file-path lock.
func WriteLocked(path string, payload any, perm FilePermType) BoolResult {
	mu := GetFileLock(path)
	defer ReleaseFileLock(path)
	mu.Lock()
	defer mu.Unlock()

	return Write(path, payload, perm)
}

func writeBytesToFile(f *os.File, data []byte, path string) BoolResult {
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return BoolFailure(errtype.IO, err, path, "failed to write bytes")
	}

	return BoolSuccess(true)
}

// WriteBytes writes raw byte slice to the specified path.
func WriteBytes(path string, data []byte, perm FilePermType) BoolResult {
	fRes := CreateFile(path, perm)
	if fRes.HasError() {
		return result.WrapFailure[bool](fRes.Fault())
	}

	return writeBytesToFile(fRes.Data(), data, path)
}

// WriteBytesLocked writes raw byte slice with an exclusive file-path lock.
func WriteBytesLocked(path string, data []byte, perm FilePermType) BoolResult {
	mu := GetFileLock(path)
	defer ReleaseFileLock(path)
	mu.Lock()
	defer mu.Unlock()

	return WriteBytes(path, data, perm)
}

// WriteString writes a string to the specified path.
func WriteString(path string, content string, perm FilePermType) BoolResult {
	return ExportText(path, content, perm)
}

// WriteStringLocked writes a string with an exclusive file-path lock.
func WriteStringLocked(path string, content string, perm FilePermType) BoolResult {
	return ExportTextLocked(path, content, perm)
}

// WriteLines writes an array of strings line-by-line to the specified path.
func WriteLines(path string, lines []string, perm FilePermType) BoolResult {
	return ExportLines(path, lines, perm)
}

// WriteLinesLocked writes an array of strings line-by-line with an exclusive file-path lock.
func WriteLinesLocked(path string, lines []string, perm FilePermType) BoolResult {
	return ExportLinesLocked(path, lines, perm)
}

// WriteJson serializes and writes any struct or map as indented JSON.
func WriteJson(path string, data any, perm FilePermType) BoolResult {
	return ExportJson(path, data, perm)
}

// WriteJsonLocked serializes and writes any struct or map as indented JSON with an exclusive lock.
func WriteJsonLocked(path string, data any, perm FilePermType) BoolResult {
	return ExportJsonLocked(path, data, perm)
}

// WriteYaml serializes and writes any struct or map as YAML.
func WriteYaml(path string, data any, perm FilePermType) BoolResult {
	return ExportYaml(path, data, perm)
}

// WriteYamlLocked serializes and writes any struct or map as YAML with an exclusive lock.
func WriteYamlLocked(path string, data any, perm FilePermType) BoolResult {
	return ExportYamlLocked(path, data, perm)
}

// WriteJSON is an alias for WriteJson.
func WriteJSON(path string, data any, perm FilePermType) BoolResult {
	return WriteJson(path, data, perm)
}

// WriteJSONLocked is an alias for WriteJsonLocked.
func WriteJSONLocked(path string, data any, perm FilePermType) BoolResult {
	return WriteJsonLocked(path, data, perm)
}

// WriteYAML is an alias for WriteYaml.
func WriteYAML(path string, data any, perm FilePermType) BoolResult {
	return WriteYaml(path, data, perm)
}

// WriteYAMLLocked is an alias for WriteYamlLocked.
func WriteYAMLLocked(path string, data any, perm FilePermType) BoolResult {
	return WriteYamlLocked(path, data, perm)
}
