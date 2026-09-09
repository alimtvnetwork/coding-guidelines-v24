package applogger

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/enum/openfiletype"
	"coding-guidelines/common/pkg/fileutil"
	"coding-guidelines/common/pkg/result"
)

// FileSink writes log entries to a file path.
type FileSink struct {
	lock     sync.Mutex
	filePath string
	file     *os.File
}

// openLogFile opens file using enum-driven fileutil utility wrapper.
func openLogFile(filePath string) result.Wrap[*os.File] {
	return fileutil.OpenFile(filePath, openfiletype.CreateAppend, filepermtype.Standard)
}

// NewFileSink creates and opens a log file sink.
func NewFileSink(filePath string) result.Wrap[*FileSink] {
	wrap := openLogFile(filePath)
	if wrap.IsFailed() {
		return result.FailureFromWrap[*FileSink](wrap)
	}

	return result.WrapSuccess(&FileSink{filePath: filePath, file: wrap.Data()})
}

// WriteEntry serializes and writes entry to file.
func (fs *FileSink) WriteEntry(e LogEntry) error {
	fs.lock.Lock()
	defer fs.lock.Unlock()

	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(fs.file, string(b))

	return err
}

// Sync flushes the file buffer to disk.
func (fs *FileSink) Sync() error {
	fs.lock.Lock()
	defer fs.lock.Unlock()
	if fs.file != nil {
		return fs.file.Sync()
	}

	return nil
}

// Close closes the file descriptor.
func (fs *FileSink) Close() error {
	fs.lock.Lock()
	defer fs.lock.Unlock()
	if fs.file != nil {
		err := fs.file.Close()
		fs.file = nil

		return err
	}

	return nil
}
