package fileutil

type readOps struct{}

func ReadBytes(path string) BytesResult {
	return ReadAll(path)
}

func (readOps) Bytes(path string) BytesResult {
	return ReadBytes(path)
}

func (readOps) String(path string) StringResult {
	return ReadString(path)
}

func (readOps) Text(path string) StringResult {
	return ReadText(path)
}

func (readOps) Lines(path string) LinesResult {
	return ReadLines(path)
}

func (readOps) Chunked(path string, chunkSize int, onChunk ChunkCallbackFunc) Int64Result {
	return ReadChunked(path, chunkSize, onChunk)
}

func (readOps) TextLocked(path string) StringResult {
	return ReadTextLocked(path)
}

func (readOps) LinesLocked(path string) LinesResult {
	return ReadLinesLocked(path)
}
