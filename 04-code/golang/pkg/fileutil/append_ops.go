package fileutil

import (
	"os"
	"strings"

	"coding-guidelines/common/pkg/enum/openfiletype"
	"coding-guidelines/common/pkg/errtype"
)

type appendOps struct{}

func appendBytesToFile(f *os.File, data []byte, path string) BoolResult {
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return BoolFailure(errtype.IO, err, path, "failed to append bytes")
	}

	return BoolSuccess(true)
}

func AppendBytes(path string, data []byte, perm FilePermType) BoolResult {
	fRes := OpenFile(path, openfiletype.CreateAppend, perm)
	if fRes.HasError() {
		return BoolFailureFault(fRes.Fault())
	}

	return appendBytesToFile(fRes.Data(), data, path)
}

func AppendBytesLocked(path string, data []byte, perm FilePermType) BoolResult {
	lock := GetFileLock(path)
	defer ReleaseFileLock(path)
	lock.Lock()
	defer lock.Unlock()

	return AppendBytes(path, data, perm)
}

func AppendString(path string, content string, perm FilePermType) BoolResult {
	return AppendBytes(path, []byte(content), perm)
}

func AppendStringLocked(path string, content string, perm FilePermType) BoolResult {
	lock := GetFileLock(path)
	defer ReleaseFileLock(path)
	lock.Lock()
	defer lock.Unlock()

	return AppendString(path, content, perm)
}

func AppendLines(path string, lines []string, perm FilePermType) BoolResult {
	if len(lines) == 0 {
		return AppendString(path, "", perm)
	}

	return AppendString(path, strings.Join(lines, "\n")+"\n", perm)
}

func AppendLinesLocked(path string, lines []string, perm FilePermType) BoolResult {
	lock := GetFileLock(path)
	defer ReleaseFileLock(path)
	lock.Lock()
	defer lock.Unlock()

	return AppendLines(path, lines, perm)
}

func (appendOps) Bytes(path string, data []byte, perm FilePermType) BoolResult {
	return AppendBytes(path, data, perm)
}

func (appendOps) String(path string, content string, perm FilePermType) BoolResult {
	return AppendString(path, content, perm)
}

func (appendOps) Lines(path string, lines []string, perm FilePermType) BoolResult {
	return AppendLines(path, lines, perm)
}

func (appendOps) BytesLocked(path string, data []byte, perm FilePermType) BoolResult {
	return AppendBytesLocked(path, data, perm)
}

func (appendOps) StringLocked(path string, content string, perm FilePermType) BoolResult {
	return AppendStringLocked(path, content, perm)
}

func (appendOps) LinesLocked(path string, lines []string, perm FilePermType) BoolResult {
	return AppendLinesLocked(path, lines, perm)
}

func (appendOps) NewAppender(path string, perm FilePermType) *FileAppender {
	return NewFileAppender(path, perm)
}
