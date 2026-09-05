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

// BoundFileWriterOptions configures a file-specific BoundFileWriter instance.
type BoundFileWriterOptions struct {
	Path        string
	Mode        FileWriteModeType
	Perm        FilePermType
	SyncOnWrite bool
	AutoClose   bool
}

// BoundFileWriter provides a dedicated, file-bound writer and appender supporting
// automatic locking/unlocking, manual transactional locks, persistent reuse,
// and auto-closing immediately after writing or appending.
type BoundFileWriter struct {
	mu            sync.Mutex
	path          string
	mode          FileWriteModeType
	perm          FilePermType
	syncOnWrite   bool
	autoClose     bool
	file          *os.File
	bytesWritten  atomic.Int64
	bytesAppended atomic.Int64
	writeCount    atomic.Int64
}

// NewBoundFileWriter constructs a BoundFileWriter bound to a specific file path.
// By default, it operates in FileWriteModeDirect with FilePermStandard.
func NewBoundFileWriter(path string) *BoundFileWriter {
	return &BoundFileWriter{
		path:        path,
		mode:        FileWriteModeDirect,
		perm:        FilePermStandard,
		syncOnWrite: false,
		autoClose:   false,
	}
}

// NewSpecificFileWriter is an alias for NewBoundFileWriter.
func NewSpecificFileWriter(path string) *BoundFileWriter {
	return NewBoundFileWriter(path)
}

// NewFileHandler is an alias for NewBoundFileWriter.
func NewFileHandler(path string) *BoundFileWriter {
	return NewBoundFileWriter(path)
}

// NewBoundFileWriterWithOptions constructs a BoundFileWriter with custom settings.
func NewBoundFileWriterWithOptions(opts BoundFileWriterOptions) *BoundFileWriter {
	perm := opts.Perm
	if perm == 0 {
		perm = FilePermStandard
	}

	mode := opts.Mode
	if mode == 0 {
		mode = FileWriteModeDirect
	}

	return &BoundFileWriter{
		path:        opts.Path,
		mode:        mode,
		perm:        perm,
		syncOnWrite: opts.SyncOnWrite,
		autoClose:   opts.AutoClose,
	}
}

// Path returns the bound target file path.
func (w *BoundFileWriter) Path() string {
	return w.path
}

// Mode returns the active write strategy mode.
func (w *BoundFileWriter) Mode() FileWriteModeType {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.mode
}

// Perm returns the active permission bitmask.
func (w *BoundFileWriter) Perm() FilePermType {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.perm
}

// SetMode dynamically shifts the writing strategy (Direct, Atomic, Truncate).
func (w *BoundFileWriter) SetMode(mode FileWriteModeType) *BoundFileWriter {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.mode = mode

	return w
}

// SetPerm dynamically shifts the file permissions.
func (w *BoundFileWriter) SetPerm(perm FilePermType) *BoundFileWriter {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.perm = perm

	return w
}

// SetSyncOnWrite dynamically shifts the fsync behavior flag.
func (w *BoundFileWriter) SetSyncOnWrite(isSync bool) *BoundFileWriter {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.syncOnWrite = isSync

	return w
}

// SetAutoClose configures whether the file descriptor is closed immediately
// after every write or append operation.
func (w *BoundFileWriter) SetAutoClose(isAuto bool) *BoundFileWriter {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.autoClose = isAuto

	return w
}

// IsOpen checks if an active file descriptor is currently held open.
func (w *BoundFileWriter) IsOpen() bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.file != nil
}

// BytesWritten returns total bytes written via Write operations.
func (w *BoundFileWriter) BytesWritten() int64 {
	return w.bytesWritten.Load()
}

// BytesAppended returns total bytes written via Append operations.
func (w *BoundFileWriter) BytesAppended() int64 {
	return w.bytesAppended.Load()
}

// WriteCount returns total count of write/append operations executed.
func (w *BoundFileWriter) WriteCount() int64 {
	return w.writeCount.Load()
}

// ResetCounters resets the diagnostic byte and operation counters to zero.
func (w *BoundFileWriter) ResetCounters() {
	w.bytesWritten.Store(0)
	w.bytesAppended.Store(0)
	w.writeCount.Store(0)
}

// Lock implements sync.Locker for manual transactional locking across multiple writes.
func (w *BoundFileWriter) Lock() {
	w.mu.Lock()
}

// Unlock implements sync.Locker.
func (w *BoundFileWriter) Unlock() {
	w.mu.Unlock()
}

// WithLock executes a batch of operations under a single mutex lock.
// This allows callers to perform multiple writes/appends atomically without interleaving.
func (w *BoundFileWriter) WithLock(ctx context.Context, fn BoundFileActionFunc) *appfault.AppError {
	w.mu.Lock()
	defer w.mu.Unlock()

	if fn == nil {
		return nil
	}

	return fn(w)
}

// Write writes payload to the bound file with automatic locking and unlocking.
// If autoClose is enabled, the file descriptor is closed immediately after the write.
func (w *BoundFileWriter) Write(ctx context.Context, payload []byte) *appfault.AppError {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.writeInternal(payload, w.autoClose)
}

// WriteString writes a string payload with automatic locking and unlocking.
func (w *BoundFileWriter) WriteString(ctx context.Context, text string) *appfault.AppError {
	return w.Write(ctx, []byte(text))
}

// WriteLocked writes payload assuming the caller already holds w.Lock().
// It does not acquire or release the lock, enabling compound batches.
func (w *BoundFileWriter) WriteLocked(ctx context.Context, payload []byte) *appfault.AppError {
	return w.writeInternal(payload, false)
}

// WriteAndClose writes the payload and guarantees the file handle is closed immediately,
// regardless of the default autoClose setting.
func (w *BoundFileWriter) WriteAndClose(ctx context.Context, payload []byte) *appfault.AppError {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.writeInternal(payload, true)
}

// Append appends payload to the end of the file with automatic locking and unlocking.
// If autoClose is enabled, the file descriptor is closed immediately after appending.
func (w *BoundFileWriter) Append(ctx context.Context, payload []byte) *appfault.AppError {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.appendInternal(payload, w.autoClose)
}

// AppendString appends a string with automatic locking and unlocking.
func (w *BoundFileWriter) AppendString(ctx context.Context, text string) *appfault.AppError {
	return w.Append(ctx, []byte(text))
}

// AppendLocked appends payload assuming the caller already holds w.Lock().
func (w *BoundFileWriter) AppendLocked(ctx context.Context, payload []byte) *appfault.AppError {
	return w.appendInternal(payload, false)
}

// AppendAndClose appends payload and guarantees the file handle is closed immediately,
// regardless of the default autoClose setting.
func (w *BoundFileWriter) AppendAndClose(ctx context.Context, payload []byte) *appfault.AppError {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.appendInternal(payload, true)
}

// writeInternal executes the write logic under an existing lock.
func (w *BoundFileWriter) writeInternal(payload []byte, closeAfter bool) *appfault.AppError {
	if w.path == "" {
		return appfault.New(errtype.Precondition, "file path cannot be empty")
	}

	if w.mode == FileWriteModeAtomic {
		res := WriteAtomic(w.path, payload, w.perm)
		if res.IsFailed() {
			return res.Fault()
		}

		w.bytesWritten.Add(int64(len(payload)))
		w.writeCount.Add(1)

		return nil
	}

	flags := os.O_CREATE | os.O_WRONLY
	if w.mode == FileWriteModeTruncate {
		flags |= os.O_TRUNC
	}

	dir := filepath.Dir(w.path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return appfault.Wrap(errtype.IO, err, "failed to create parent directories")
	}

	f, err := os.OpenFile(w.path, flags, w.perm.Mode())
	if err != nil {
		return appfault.Wrap(errtype.IO, err, "failed to open bound target file for writing")
	}

	if _, err := f.Write(payload); err != nil {
		f.Close()

		return appfault.Wrap(errtype.IO, err, "failed to write payload to bound file")
	}

	if w.syncOnWrite {
		if err := f.Sync(); err != nil {
			f.Close()

			return appfault.Wrap(errtype.IO, err, "failed to sync bound file to storage")
		}
	}

	w.bytesWritten.Add(int64(len(payload)))
	w.writeCount.Add(1)

	if closeAfter {
		if err := f.Close(); err != nil {
			return appfault.Wrap(errtype.IO, err, "failed to close bound file after write")
		}

		if w.file != nil && w.file == f {
			w.file = nil
		}

		return nil
	}

	// Retain open file handle for subsequent reuse
	if w.file != nil && w.file != f {
		w.file.Close()
	}

	w.file = f

	return nil
}

// appendInternal executes append logic under an existing lock.
func (w *BoundFileWriter) appendInternal(payload []byte, closeAfter bool) *appfault.AppError {
	if w.path == "" {
		return appfault.New(errtype.Precondition, "file path cannot be empty")
	}

	dir := filepath.Dir(w.path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return appfault.Wrap(errtype.IO, err, "failed to create parent directories")
	}

	f := w.file
	var openErr error

	if f == nil {
		f, openErr = os.OpenFile(w.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, w.perm.Mode())
		if openErr != nil {
			return appfault.Wrap(errtype.IO, openErr, "failed to open bound file for appending")
		}
	}

	n, writeErr := f.Write(payload)
	if writeErr != nil {
		if closeAfter {
			f.Close()
			w.file = nil
		}

		return appfault.Wrap(errtype.IO, writeErr, "failed to append payload to bound file")
	}

	if w.syncOnWrite {
		if err := f.Sync(); err != nil {
			if closeAfter {
				f.Close()
				w.file = nil
			}

			return appfault.Wrap(errtype.IO, err, "failed to sync bound file after append")
		}
	}

	w.bytesAppended.Add(int64(n))
	w.writeCount.Add(1)

	if closeAfter {
		closeErr := f.Close()
		w.file = nil
		if closeErr != nil {
			return appfault.Wrap(errtype.IO, closeErr, "failed to close bound file after append")
		}

		return nil
	}

	w.file = f

	return nil
}

// Sync forces pending OS buffers to persistent storage if a file handle is open.
func (w *BoundFileWriter) Sync() *appfault.AppError {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.file != nil {
		if err := w.file.Sync(); err != nil {
			return appfault.Wrap(errtype.IO, err, "failed to sync open bound file descriptor")
		}
	}

	return nil
}

// Close safely flushes and closes the active file descriptor if currently open.
func (w *BoundFileWriter) Close() *appfault.AppError {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.file != nil {
		err := w.file.Close()
		w.file = nil
		if err != nil {
			return appfault.Wrap(errtype.IO, err, "failed to close bound file descriptor")
		}
	}

	return nil
}

// Name returns the writer name for streamwriter compatibility.
func (w *BoundFileWriter) Name() string {
	return fmt.Sprintf("bound-file-writer[%s]", filepath.Base(w.path))
}

// AsWriter returns streamwriter.Writer[[]byte] representation.
func (w *BoundFileWriter) AsWriter() streamwriter.Writer[[]byte] {
	return w
}

// StdWriter returns an io.WriteCloser adapter that writes to the bound file.
func (w *BoundFileWriter) StdWriter() io.WriteCloser {
	return &boundWriterStdAdapter{writer: w, isAppend: false}
}

// StdAppender returns an io.WriteCloser adapter that appends to the bound file.
func (w *BoundFileWriter) StdAppender() io.WriteCloser {
	return &boundWriterStdAdapter{writer: w, isAppend: true}
}

type boundWriterStdAdapter struct {
	writer   *BoundFileWriter
	isAppend bool
}

func (s *boundWriterStdAdapter) Write(p []byte) (n int, err error) {
	ctx := context.Background()
	var appErr *appfault.AppError

	if s.isAppend {
		appErr = s.writer.Append(ctx, p)
	} else {
		appErr = s.writer.Write(ctx, p)
	}

	if appErr != nil {
		return 0, appErr
	}

	return len(p), nil
}

func (s *boundWriterStdAdapter) Close() error {
	appErr := s.writer.Close()
	if appErr != nil {
		return appErr
	}

	return nil
}

var _ streamwriter.Writer[[]byte] = (*BoundFileWriter)(nil)
var _ sync.Locker = (*BoundFileWriter)(nil)
var _ io.WriteCloser = (*boundWriterStdAdapter)(nil)
