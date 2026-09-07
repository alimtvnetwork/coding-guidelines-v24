package fileutil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
	"coding-guidelines/common/pkg/streamwriter"
)

var chunkBufferPool = sync.Pool{
	New: func() any {
		b := make([]byte, DefaultBufferSize)

		return &b
	},
}

func createTempFile(dir string, filename string, perm FilePermType) (*os.File, string, *appfault.AppError) {
	tmpPattern := fmt.Sprintf(".%s.tmp.%d.%d", filename, os.Getpid(), time.Now().UnixNano())
	tmpPath := filepath.Join(dir, tmpPattern)
	tmpFile, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, perm.Mode())
	if err != nil {
		return nil, "", appfault.WrapFile(errtype.IO, err, tmpPath, "failed to create atomic temp file")
	}

	return tmpFile, tmpPath, nil
}

func writeTempData(tmpFile *os.File, tmpPath string, data []byte) *appfault.AppError {
	if len(data) == 0 {
		return nil
	}

	if _, err := tmpFile.Write(data); err != nil {
		_ = tmpFile.Close()
		_ = os.Remove(tmpPath)

		return appfault.WrapFile(errtype.IO, err, tmpPath, "failed to write atomic data")
	}

	return nil
}

func syncAndCloseTemp(tmpFile *os.File, tmpPath string) *appfault.AppError {
	if err := tmpFile.Sync(); err != nil {
		_ = tmpFile.Close()
		_ = os.Remove(tmpPath)

		return appfault.WrapFile(errtype.IO, err, tmpPath, "failed to sync atomic file")
	}

	if err := tmpFile.Close(); err != nil {
		_ = os.Remove(tmpPath)

		return appfault.WrapFile(errtype.IO, err, tmpPath, "failed to close atomic file")
	}

	return nil
}

func renameTemp(tmpPath string, path string) *appfault.AppError {
	_ = os.Remove(path)

	if err := os.Rename(tmpPath, path); err != nil {
		_ = os.Remove(tmpPath)

		return appfault.WrapFile(errtype.IO, err, path, "failed to atomically rename")
	}

	return nil
}

func commitAtomicWrite(tmpFile *os.File, tmpPath string, path string, data []byte) BoolResult {
	if err := writeTempData(tmpFile, tmpPath, data); err != nil {
		return result.WrapFailure[bool](err)
	}

	if err := syncAndCloseTemp(tmpFile, tmpPath); err != nil {
		return result.WrapFailure[bool](err)
	}

	if err := renameTemp(tmpPath, path); err != nil {
		return result.WrapFailure[bool](err)
	}

	return BoolSuccess(true)
}

func WriteAtomic(path string, data []byte, perm FilePermType) BoolResult {
	if len(path) == 0 {
		return BoolFailureMsg(errtype.Validation, path, "path cannot be empty")
	}

	dir := filepath.Dir(path)
	if ensureRes := EnsureDir(dir, FilePermStandard); ensureRes.IsFailed() {
		return result.WrapFailure[bool](ensureRes.Fault())
	}

	tmpFile, tmpPath, err := createTempFile(dir, filepath.Base(path), perm)
	if err != nil {
		return result.WrapFailure[bool](err)
	}

	return commitAtomicWrite(tmpFile, tmpPath, path, data)
}

func acquireChunkBuffer(size int) []byte {
	if size <= 0 {
		size = DefaultBufferSize
	}

	bufPtr := chunkBufferPool.Get().(*[]byte)
	buf := *bufPtr
	if len(buf) < size {
		return make([]byte, size)
	}

	return buf[:size]
}

func processChunk(n int, buf []byte, onChunk ChunkCallbackFunc) *appfault.AppError {
	if n <= 0 {
		return nil
	}

	if onChunk == nil {
		return nil
	}

	return onChunk(buf[:n])
}

func handleChunkStep(n int, buf []byte, err error, onChunk ChunkCallbackFunc, path string) *appfault.AppError {
	if fault := processChunk(n, buf, onChunk); fault != nil {
		return fault
	}

	if err == nil || err == io.EOF {
		return nil
	}

	return appfault.WrapFile(errtype.IO, err, path, "error reading chunk")
}

func executeChunkedRead(r io.Reader, path string, buf []byte, onChunk ChunkCallbackFunc) Int64Result {
	var total int64
	for {
		n, err := r.Read(buf)
		total += int64(n)
		if fault := handleChunkStep(n, buf, err, onChunk, path); fault != nil {
			return result.WrapFailure[int64](fault)
		}

		if err == io.EOF {
			return Int64Success(total)
		}
	}
}

func ReadChunked(path string, chunkSize int, onChunk ChunkCallbackFunc) Int64Result {
	if len(path) == 0 {
		return Int64FailureMsg(errtype.Validation, path, "path cannot be empty")
	}

	openRes := Open(path)
	if openRes.IsFailed() {
		return result.WrapFailure[int64](openRes.Fault())
	}

	defer openRes.Data().Close()

	buf := acquireChunkBuffer(chunkSize)
	defer chunkBufferPool.Put(&buf)

	return executeChunkedRead(openRes.Data(), path, buf, onChunk)
}

func writeChunkToFile(f *os.File, path string, buf []byte) *appfault.AppError {
	if len(buf) == 0 {
		return nil
	}

	if _, err := f.Write(buf); err != nil {
		return appfault.WrapFile(errtype.IO, err, path, "failed writing chunk")
	}

	return nil
}

func readAndWriteChunk(f *os.File, path string, reader io.Reader, buf []byte) (int, bool, *appfault.AppError) {
	n, err := reader.Read(buf)
	if writeFault := writeChunkToFile(f, path, buf[:n]); writeFault != nil {
		return n, false, writeFault
	}

	if err == io.EOF {
		return n, true, nil
	}

	if err != nil {
		return n, false, appfault.WrapFile(errtype.IO, err, path, "error reading source for write")
	}

	return n, false, nil
}

func syncChunkedFile(f *os.File, path string, total int64) Int64Result {
	if err := f.Sync(); err != nil {
		return Int64Failure(errtype.IO, err, path, "failed syncing file")
	}

	return Int64Success(total)
}

func writeAllChunks(f *os.File, path string, reader io.Reader, buf []byte) Int64Result {
	var total int64
	for {
		n, isEof, fault := readAndWriteChunk(f, path, reader, buf)
		total += int64(n)
		if fault != nil {
			return result.WrapFailure[int64](fault)
		}

		if isEof {
			return syncChunkedFile(f, path, total)
		}
	}
}

func executeWriteChunked(path string, perm FilePermType, reader io.Reader, bufferSize int) Int64Result {
	openRes := OpenFile(path, FileOpenCreateTruncate, perm)
	if openRes.IsFailed() {
		return result.WrapFailure[int64](openRes.Fault())
	}

	defer openRes.Data().Close()

	buf := acquireChunkBuffer(bufferSize)
	defer chunkBufferPool.Put(&buf)

	return writeAllChunks(openRes.Data(), path, reader, buf)
}

func WriteChunked(path string, perm FilePermType, reader io.Reader, bufferSize int) Int64Result {
	if len(path) == 0 {
		return Int64FailureMsg(errtype.Validation, path, "path cannot be empty")
	}

	if reader == nil {
		return Int64FailureMsg(errtype.Validation, path, "reader cannot be nil")
	}

	return executeWriteChunked(path, perm, reader, bufferSize)
}

func NewFileWriter(path string, openMode FileOpenModeType, perm FilePermType) result.Wrap[*streamwriter.PluggableWriter[any]] {
	openRes := OpenFile(path, openMode, perm)
	if openRes.IsFailed() {
		return result.WrapFailure[*streamwriter.PluggableWriter[any]](openRes.Fault())
	}

	file := openRes.Data()
	writer := streamwriter.NewAnyWriter(streamwriter.WriterOptions[any]{
		Name:        "file-writer:" + filepath.Base(path),
		Destination: file,
	})

	return result.WrapSuccess(writer)
}
