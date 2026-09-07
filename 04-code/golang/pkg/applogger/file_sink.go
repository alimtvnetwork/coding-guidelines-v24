package applogger

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/enum/openfiletype"
	"coding-guidelines/common/pkg/fileutil"
)

// FileSink writes log entries to a file path.
type FileSink struct {
	lock     sync.Mutex
	filePath string
	file     *os.File
}

// openLogFile opens file using enum-driven fileutil utility wrapper.
func openLogFile(filePath string) (*os.File, error) {
	wrap := fileutil.OpenFile(filePath, openfiletype.CreateAppend, filepermtype.Standard)
	if wrap.IsFailed() {
		return nil, wrap.Fault()
	}

	return wrap.Data(), nil
}

// NewFileSink creates and opens a log file sink.
func NewFileSink(filePath string) (*FileSink, error) {
	f, err := openLogFile(filePath)
	if err != nil {
		return nil, err
	}

	return &FileSink{filePath: filePath, file: f}, nil
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
