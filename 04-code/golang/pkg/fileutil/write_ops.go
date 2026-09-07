package fileutil

import (
	"io"
)

type writeOps struct{}

func (writeOps) Any(path string, payload any, perm FilePermType) BoolResult {
	return Write(path, payload, perm)
}

func (writeOps) AnyLocked(path string, payload any, perm FilePermType) BoolResult {
	return WriteLocked(path, payload, perm)
}

func (writeOps) File(path string, data []byte, perm FilePermType) BoolResult {
	return WriteFile(path, data, perm)
}

func (writeOps) Bytes(path string, data []byte, perm FilePermType) BoolResult {
	return WriteBytes(path, data, perm)
}

func (writeOps) BytesLocked(path string, data []byte, perm FilePermType) BoolResult {
	return WriteBytesLocked(path, data, perm)
}

func (writeOps) String(path string, content string, perm FilePermType) BoolResult {
	return WriteString(path, content, perm)
}

func (writeOps) StringLocked(path string, content string, perm FilePermType) BoolResult {
	return WriteStringLocked(path, content, perm)
}

func (writeOps) Lines(path string, lines []string, perm FilePermType) BoolResult {
	return WriteLines(path, lines, perm)
}

func (writeOps) LinesLocked(path string, lines []string, perm FilePermType) BoolResult {
	return WriteLinesLocked(path, lines, perm)
}

func (writeOps) Json(path string, data any, perm FilePermType) BoolResult {
	return WriteJson(path, data, perm)
}

func (writeOps) JsonLocked(path string, data any, perm FilePermType) BoolResult {
	return WriteJsonLocked(path, data, perm)
}

func (writeOps) Yaml(path string, data any, perm FilePermType) BoolResult {
	return WriteYaml(path, data, perm)
}

func (writeOps) YamlLocked(path string, data any, perm FilePermType) BoolResult {
	return WriteYamlLocked(path, data, perm)
}

func (writeOps) JSON(path string, data any, perm FilePermType) BoolResult {
	return WriteJson(path, data, perm)
}

func (writeOps) JSONLocked(path string, data any, perm FilePermType) BoolResult {
	return WriteJsonLocked(path, data, perm)
}

func (writeOps) YAML(path string, data any, perm FilePermType) BoolResult {
	return WriteYaml(path, data, perm)
}

func (writeOps) YAMLLocked(path string, data any, perm FilePermType) BoolResult {
	return WriteYamlLocked(path, data, perm)
}

func (writeOps) Atomic(path string, data []byte, perm FilePermType) BoolResult {
	return WriteAtomic(path, data, perm)
}

func (writeOps) Chunked(path string, perm FilePermType, reader io.Reader, bufferSize int) Int64Result {
	return WriteChunked(path, perm, reader, bufferSize)
}
