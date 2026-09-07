package fileutil

import (
	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/payloadconv"
	"coding-guidelines/common/pkg/result"
)

// Write writes any generic payload (struct, map, array, string, bytes) to a file.
// It automatically converts the payload via payloadconv.ToBytes.
func Write(path string, payload any, perm FilePermType) result.Wrap[bool] {
	convRes := payloadconv.ToBytes(payload)
	if convRes.IsFailure() {
		return result.WrapFailure[bool](convRes.Fault())
	}

	return WriteBytes(path, convRes.Data(), perm)
}

// WriteLocked writes any generic payload with an exclusive file-path lock.
func WriteLocked(path string, payload any, perm FilePermType) result.Wrap[bool] {
	mu := GetFileLock(path)
	defer ReleaseFileLock(path)
	mu.Lock()
	defer mu.Unlock()

	return Write(path, payload, perm)
}

// WriteBytes writes raw byte slice to the specified path.
func WriteBytes(path string, data []byte, perm FilePermType) result.Wrap[bool] {
	fRes := CreateFile(path, perm)
	if fRes.HasError() {
		return result.WrapFailure[bool](fRes.Fault())
	}

	f := fRes.Data()
	defer f.Close()

	_, err := f.Write(data)
	if err != nil {
		return result.WrapFailure[bool](appfault.Wrap(errtype.IO, err, "failed to write bytes to: "+path))
	}

	return result.WrapSuccess(true)
}

// WriteBytesLocked writes raw byte slice with an exclusive file-path lock.
func WriteBytesLocked(path string, data []byte, perm FilePermType) result.Wrap[bool] {
	mu := GetFileLock(path)
	defer ReleaseFileLock(path)
	mu.Lock()
	defer mu.Unlock()

	return WriteBytes(path, data, perm)
}

// WriteString writes a string to the specified path.
func WriteString(path string, content string, perm FilePermType) result.Wrap[bool] {
	return ExportText(path, content, perm)
}

// WriteStringLocked writes a string with an exclusive file-path lock.
func WriteStringLocked(path string, content string, perm FilePermType) result.Wrap[bool] {
	return ExportTextLocked(path, content, perm)
}

// WriteLines writes an array of strings line-by-line to the specified path.
func WriteLines(path string, lines []string, perm FilePermType) result.Wrap[bool] {
	return ExportLines(path, lines, perm)
}

// WriteLinesLocked writes an array of strings line-by-line with an exclusive file-path lock.
func WriteLinesLocked(path string, lines []string, perm FilePermType) result.Wrap[bool] {
	return ExportLinesLocked(path, lines, perm)
}

// WriteJSON serializes and writes any struct or map as indented JSON.
func WriteJSON(path string, data any, perm FilePermType) result.Wrap[bool] {
	return ExportJSON(path, data, perm)
}

// WriteJSONLocked serializes and writes any struct or map as indented JSON with an exclusive lock.
func WriteJSONLocked(path string, data any, perm FilePermType) result.Wrap[bool] {
	return ExportJSONLocked(path, data, perm)
}
