package fileutil

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/streamwriter"
)

// FileWriteModeType specifies the behavior strategy for FileWriter.
type FileWriteModeType uint8

const (
	// FileWriteModeDirect writes directly in-place to the target file.
	FileWriteModeDirect FileWriteModeType = 1
	// FileWriteModeAtomic writes to a temp file and atomically renames.
	FileWriteModeAtomic FileWriteModeType = 2
	// FileWriteModeTruncate creates or truncates the file before writing.
	FileWriteModeTruncate FileWriteModeType = 3
)

var writeModeNames = map[FileWriteModeType]string{
	FileWriteModeDirect:   "Direct",
	FileWriteModeAtomic:   "Atomic",
	FileWriteModeTruncate: "Truncate",
}

// Name returns the PascalCase mode name.
func (m FileWriteModeType) Name() string {
	if name, ok := writeModeNames[m]; ok {
		return name
	}

	return fmt.Sprintf("FileWriteMode(%d)", uint8(m))
}

// String implements fmt.Stringer.
func (m FileWriteModeType) String() string {
	return m.Name()
}

// IsValid returns true if the mode is recognized.
func (m FileWriteModeType) IsValid() bool {
	_, ok := writeModeNames[m]

	return ok
}

// FileWriterOptions configures FileWriter behavior.
type FileWriterOptions struct {
	Path        string
	Mode        FileWriteModeType
	Perm        FilePermType
	SyncOnWrite bool
}

// FileWriter provides a configurable, behavior-shifting file output engine.
type FileWriter struct {
	mu          sync.RWMutex
	path        string
	mode        FileWriteModeType
	perm        FilePermType
	syncOnWrite bool
	file        *os.File
}

// NewFileWriterEngine constructs a new FileWriter with default FileWriteModeDirect and FilePermStandard.
func NewFileWriterEngine(path string) *FileWriter {
	return &FileWriter{
		path:        path,
		mode:        FileWriteModeDirect,
		perm:        FilePermStandard,
		syncOnWrite: false,
	}
}

// NewFileWriterWithOptions constructs a FileWriter with custom options.
func NewFileWriterWithOptions(opts FileWriterOptions) *FileWriter {
	perm := opts.Perm
	if perm == 0 {
		perm = FilePermStandard
	}

	mode := opts.Mode
	if mode == 0 {
		mode = FileWriteModeDirect
	}

	return &FileWriter{
		path:        opts.Path,
		mode:        mode,
		perm:        perm,
		syncOnWrite: opts.SyncOnWrite,
	}
}

// Name returns the writer name for streamwriter compatibility.
func (w *FileWriter) Name() string {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return fmt.Sprintf("file-writer[%s]", filepath.Base(w.path))
}

// Path returns the configured target file path.
func (w *FileWriter) Path() string {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.path
}

// Mode returns the active write strategy mode.
func (w *FileWriter) Mode() FileWriteModeType {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.mode
}

// SetMode dynamically shifts the file writing behavior strategy.
func (w *FileWriter) SetMode(mode FileWriteModeType) *FileWriter {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.mode = mode

	return w
}

// SetPerm dynamically shifts the file permission bitmask.
func (w *FileWriter) SetPerm(perm FilePermType) *FileWriter {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.perm = perm

	return w
}

// SetSyncOnWrite dynamically shifts the fsync behavior flag.
func (w *FileWriter) SetSyncOnWrite(isSync bool) *FileWriter {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.syncOnWrite = isSync

	return w
}

// Write executes a write using the active behavior mode, returning *appfault.AppError.
func (w *FileWriter) Write(ctx context.Context, payload []byte) *appfault.AppError {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.path == "" {
		return appfault.New(errtype.Precondition, "file path cannot be empty")
	}

	switch w.mode {
	case FileWriteModeAtomic:
		res := WriteAtomic(w.path, payload, w.perm)
		if res.IsFailed() {
			return res.Fault()
		}

		return nil

	case FileWriteModeTruncate:
		return w.writeDirect(payload, os.O_CREATE|os.O_TRUNC|os.O_WRONLY)

	default:
		return w.writeDirect(payload, os.O_CREATE|os.O_WRONLY)
	}
}

func (w *FileWriter) writeDirect(payload []byte, flags int) *appfault.AppError {
	dir := filepath.Dir(w.path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return appfault.Wrap(errtype.IO, err, "failed to create parent directories")
	}

	f, err := os.OpenFile(w.path, flags, w.perm.Mode())
	if err != nil {
		return appfault.Wrap(errtype.IO, err, "failed to open target file")
	}

	defer f.Close()

	if _, err := f.Write(payload); err != nil {
		return appfault.Wrap(errtype.IO, err, "failed to write payload to file")
	}

	if w.syncOnWrite {
		if err := f.Sync(); err != nil {
			return appfault.Wrap(errtype.IO, err, "failed to sync file to storage")
		}
	}

	return nil
}

// WriteString writes a string payload to the target file.
func (w *FileWriter) WriteString(ctx context.Context, text string) *appfault.AppError {
	return w.Write(ctx, []byte(text))
}

// WriteStd satisfies io.Writer compatibility.
func (w *FileWriter) WriteStd(p []byte) (n int, err error) {
	appErr := w.Write(context.Background(), p)
	if appErr != nil {
		return 0, appErr
	}

	return len(p), nil
}

// StdWriter returns an io.WriteCloser adapter for standard library stream compatibility.
func (w *FileWriter) StdWriter() io.WriteCloser {
	return &fileWriterStdAdapter{writer: w}
}

type fileWriterStdAdapter struct {
	writer *FileWriter
}

func (s *fileWriterStdAdapter) Write(p []byte) (n int, err error) {
	return s.writer.WriteStd(p)
}

func (s *fileWriterStdAdapter) Close() error {
	appErr := s.writer.Close()
	if appErr != nil {
		return appErr
	}

	return nil
}

// Sync flushes pending OS buffers if an active file descriptor is open.
func (w *FileWriter) Sync() *appfault.AppError {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if w.file != nil {
		if err := w.file.Sync(); err != nil {
			return appfault.Wrap(errtype.IO, err, "sync failed")
		}
	}

	return nil
}

// Close closes open file handles if active.
func (w *FileWriter) Close() *appfault.AppError {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.file != nil {
		err := w.file.Close()
		w.file = nil
		if err != nil {
			return appfault.Wrap(errtype.IO, err, "close failed")
		}
	}

	return nil
}

// Lock satisfies sync.Locker.
func (w *FileWriter) Lock() {
	w.mu.Lock()
}

// Unlock satisfies sync.Locker.
func (w *FileWriter) Unlock() {
	w.mu.Unlock()
}

// AsWriter returns streamwriter.Writer[[]byte] representation.
func (w *FileWriter) AsWriter() streamwriter.Writer[[]byte] {
	return w
}

// FileAppender provides a dedicated, continuous appending engine with auto-create.
type FileAppender struct {
	mu            sync.Mutex
	path          string
	perm          FilePermType
	autoSync      bool
	file          *os.File
	bytesAppended atomic.Int64
}

// NewFileAppender constructs a FileAppender for continuous journal or log appending.
func NewFileAppender(path string, perm FilePermType) *FileAppender {
	if perm == 0 {
		perm = FilePermStandard
	}

	return &FileAppender{
		path:     path,
		perm:     perm,
		autoSync: false,
	}
}

// Path returns target file path.
func (a *FileAppender) Path() string {
	return a.path
}

// SetAutoSync dynamically shifts auto-sync behavior.
func (a *FileAppender) SetAutoSync(isAuto bool) *FileAppender {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.autoSync = isAuto

	return a
}

// ensureOpen opens the target file in append mode if not already open.
func (a *FileAppender) ensureOpen() error {
	if a.file != nil {
		return nil
	}

	dir := filepath.Dir(a.path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	f, err := os.OpenFile(a.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, a.perm.Mode())
	if err != nil {
		return err
	}

	a.file = f

	return nil
}

// Append writes data to the end of the file, returning *appfault.AppError.
func (a *FileAppender) Append(ctx context.Context, data []byte) *appfault.AppError {
	a.mu.Lock()
	defer a.mu.Unlock()

	if err := a.ensureOpen(); err != nil {
		return appfault.Wrap(errtype.IO, err, "failed to open appender target file")
	}

	n, err := a.file.Write(data)
	if err != nil {
		return appfault.Wrap(errtype.IO, err, "failed to append bytes to file")
	}

	a.bytesAppended.Add(int64(n))

	if a.autoSync {
		if err := a.file.Sync(); err != nil {
			return appfault.Wrap(errtype.IO, err, "failed to sync appender file")
		}
	}

	return nil
}

// AppendString appends a string to the file.
func (a *FileAppender) AppendString(ctx context.Context, text string) *appfault.AppError {
	return a.Append(ctx, []byte(text))
}

// WriteStd satisfies io.Writer compatibility.
func (a *FileAppender) WriteStd(p []byte) (n int, err error) {
	appErr := a.Append(context.Background(), p)
	if appErr != nil {
		return 0, appErr
	}

	return len(p), nil
}

// StdWriter returns an io.WriteCloser adapter for standard library stream compatibility.
func (a *FileAppender) StdWriter() io.WriteCloser {
	return &appenderStdAdapter{appender: a}
}

type appenderStdAdapter struct {
	appender *FileAppender
}

func (s *appenderStdAdapter) Write(p []byte) (n int, err error) {
	return s.appender.WriteStd(p)
}

func (s *appenderStdAdapter) Close() error {
	appErr := s.appender.Close()
	if appErr != nil {
		return appErr
	}

	return nil
}

// BytesAppended returns the total number of bytes written via this appender.
func (a *FileAppender) BytesAppended() int64 {
	return a.bytesAppended.Load()
}

// Sync forces all pending writes to disk.
func (a *FileAppender) Sync() *appfault.AppError {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.file != nil {
		if err := a.file.Sync(); err != nil {
			return appfault.Wrap(errtype.IO, err, "appender sync failed")
		}
	}

	return nil
}

// Close flushes and closes the active file descriptor.
func (a *FileAppender) Close() *appfault.AppError {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.file != nil {
		err := a.file.Close()
		a.file = nil
		if err != nil {
			return appfault.Wrap(errtype.IO, err, "appender close failed")
		}
	}

	return nil
}

// Lock satisfies sync.Locker.
func (a *FileAppender) Lock() {
	a.mu.Lock()
}

// Unlock satisfies sync.Locker.
func (a *FileAppender) Unlock() {
	a.mu.Unlock()
}

// Name returns streamwriter identifier.
func (a *FileAppender) Name() string {
	return fmt.Sprintf("file-appender[%s]", filepath.Base(a.path))
}

// Write satisfies streamwriter.Writer[[]byte].
func (a *FileAppender) Write(ctx context.Context, payload []byte) *appfault.AppError {
	return a.Append(ctx, payload)
}

// AsWriter returns streamwriter.Writer[[]byte] representation.
func (a *FileAppender) AsWriter() streamwriter.Writer[[]byte] {
	return a
}

var _ streamwriter.Writer[[]byte] = (*FileWriter)(nil)
var _ sync.Locker = (*FileWriter)(nil)
var _ streamwriter.Writer[[]byte] = (*FileAppender)(nil)
var _ sync.Locker = (*FileAppender)(nil)
var _ io.WriteCloser = (*fileWriterStdAdapter)(nil)
var _ io.WriteCloser = (*appenderStdAdapter)(nil)
