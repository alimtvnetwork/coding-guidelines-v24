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
	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/enum/filewritemodetype"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/streamwriter"
)

type (
	BoundFileWriterOptions struct {
		Path          string
		Mode          FileWriteModeType
		Perm          FilePermType
		IsSyncOnWrite bool
		IsAutoClose   bool
	}

	BoundFileWriter struct {
		lock          sync.Mutex
		path          string
		mode          FileWriteModeType
		perm          FilePermType
		isSyncOnWrite bool
		isAutoClose   bool
		file          *os.File
		bytesWritten  atomic.Int64
		bytesAppended atomic.Int64
		writeCount    atomic.Int64
	}
)

func NewBoundFileWriter(path string) *BoundFileWriter {
	return &BoundFileWriter{
		path:          path,
		mode:          filewritemodetype.Direct,
		perm:          filepermtype.Standard,
		isSyncOnWrite: false,
		isAutoClose:   false,
	}
}

func NewBoundFileWriterWithMode(path string, mode FileWriteModeType) *BoundFileWriter {
	w := NewBoundFileWriter(path)
	w.mode = mode

	return w
}

func NewBoundFileWriterWithPerm(path string, perm FilePermType) *BoundFileWriter {
	w := NewBoundFileWriter(path)
	w.perm = perm

	return w
}

func NewBoundFileWriterWithSync(path string, isSyncOnWrite bool) *BoundFileWriter {
	w := NewBoundFileWriter(path)
	w.isSyncOnWrite = isSyncOnWrite

	return w
}

func NewBoundFileWriterWithAutoClose(path string, isAutoClose bool) *BoundFileWriter {
	w := NewBoundFileWriter(path)
	w.isAutoClose = isAutoClose

	return w
}

func NewBoundFileWriterConfig(
	path string,
	mode FileWriteModeType,
	perm FilePermType,
	isSyncOnWrite bool,
	isAutoClose bool,
) *BoundFileWriter {
	return &BoundFileWriter{
		path:          path,
		mode:          mode,
		perm:          perm,
		isSyncOnWrite: isSyncOnWrite,
		isAutoClose:   isAutoClose,
	}
}

func NewSpecificFileWriter(path string) *BoundFileWriter {
	return NewBoundFileWriter(path)
}

func NewFileHandler(path string) *BoundFileWriter {
	return NewBoundFileWriter(path)
}

func NewBoundFileWriterWithOptions(opts BoundFileWriterOptions) *BoundFileWriter {
	perm := opts.Perm
	if perm == 0 {
		perm = filepermtype.Standard
	}

	mode := opts.Mode
	if mode == 0 {
		mode = filewritemodetype.Direct
	}

	return &BoundFileWriter{
		path:          opts.Path,
		mode:          mode,
		perm:          perm,
		isSyncOnWrite: opts.IsSyncOnWrite,
		isAutoClose:   opts.IsAutoClose,
	}
}

func (w *BoundFileWriter) Path() string {
	return w.path
}

func (w *BoundFileWriter) Mode() FileWriteModeType {
	w.lock.Lock()
	defer w.lock.Unlock()

	return w.mode
}

func (w *BoundFileWriter) Perm() FilePermType {
	w.lock.Lock()
	defer w.lock.Unlock()

	return w.perm
}

func (w *BoundFileWriter) SetMode(mode FileWriteModeType) *BoundFileWriter {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.mode = mode

	return w
}

func (w *BoundFileWriter) SetPerm(perm FilePermType) *BoundFileWriter {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.perm = perm

	return w
}

func (w *BoundFileWriter) IsSyncOnWrite() bool {
	w.lock.Lock()
	defer w.lock.Unlock()

	return w.isSyncOnWrite
}

func (w *BoundFileWriter) IsAutoClose() bool {
	w.lock.Lock()
	defer w.lock.Unlock()

	return w.isAutoClose
}

func (w *BoundFileWriter) SetSyncOnWrite(isSyncOnWrite bool) *BoundFileWriter {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.isSyncOnWrite = isSyncOnWrite

	return w
}

func (w *BoundFileWriter) SetAutoClose(isAutoClose bool) *BoundFileWriter {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.isAutoClose = isAutoClose

	return w
}

func (w *BoundFileWriter) WithMode(mode FileWriteModeType) *BoundFileWriter {
	return w.SetMode(mode)
}

func (w *BoundFileWriter) WithPerm(perm FilePermType) *BoundFileWriter {
	return w.SetPerm(perm)
}

func (w *BoundFileWriter) WithSyncOnWrite(isSyncOnWrite bool) *BoundFileWriter {
	return w.SetSyncOnWrite(isSyncOnWrite)
}

func (w *BoundFileWriter) WithAutoClose(isAutoClose bool) *BoundFileWriter {
	return w.SetAutoClose(isAutoClose)
}

func (w *BoundFileWriter) EnableSyncOnWrite() *BoundFileWriter {
	return w.SetSyncOnWrite(true)
}

func (w *BoundFileWriter) DisableSyncOnWrite() *BoundFileWriter {
	return w.SetSyncOnWrite(false)
}

func (w *BoundFileWriter) EnableAutoClose() *BoundFileWriter {
	return w.SetAutoClose(true)
}

func (w *BoundFileWriter) DisableAutoClose() *BoundFileWriter {
	return w.SetAutoClose(false)
}

func (w *BoundFileWriter) IsOpen() bool {
	w.lock.Lock()
	defer w.lock.Unlock()

	return w.file != nil
}

func (w *BoundFileWriter) BytesWritten() int64 {
	return w.bytesWritten.Load()
}

func (w *BoundFileWriter) BytesAppended() int64 {
	return w.bytesAppended.Load()
}

func (w *BoundFileWriter) WriteCount() int64 {
	return w.writeCount.Load()
}

func (w *BoundFileWriter) ResetCounters() {
	w.bytesWritten.Store(0)
	w.bytesAppended.Store(0)
	w.writeCount.Store(0)
}

func (w *BoundFileWriter) Lock() {
	w.lock.Lock()
}

func (w *BoundFileWriter) Unlock() {
	w.lock.Unlock()
}

func (w *BoundFileWriter) WithLock(ctx context.Context, fn BoundFileActionFunc) *appfault.AppError {
	w.lock.Lock()
	defer w.lock.Unlock()

	if fn == nil {
		return nil
	}

	return fn(w)
}

func (w *BoundFileWriter) Write(ctx context.Context, payload []byte) *appfault.AppError {
	w.lock.Lock()
	defer w.lock.Unlock()

	return w.writeInternal(payload, w.isAutoClose)
}

func (w *BoundFileWriter) WriteString(ctx context.Context, text string) *appfault.AppError {
	return w.Write(ctx, []byte(text))
}

func (w *BoundFileWriter) WriteLocked(ctx context.Context, payload []byte) *appfault.AppError {
	return w.writeInternal(payload, false)
}

func (w *BoundFileWriter) WriteAndClose(ctx context.Context, payload []byte) *appfault.AppError {
	w.lock.Lock()
	defer w.lock.Unlock()

	return w.writeInternal(payload, true)
}

func (w *BoundFileWriter) Append(ctx context.Context, payload []byte) *appfault.AppError {
	w.lock.Lock()
	defer w.lock.Unlock()

	return w.appendInternal(payload, w.isAutoClose)
}

func (w *BoundFileWriter) AppendString(ctx context.Context, text string) *appfault.AppError {
	return w.Append(ctx, []byte(text))
}

func (w *BoundFileWriter) AppendLocked(ctx context.Context, payload []byte) *appfault.AppError {
	return w.appendInternal(payload, false)
}

func (w *BoundFileWriter) AppendAndClose(ctx context.Context, payload []byte) *appfault.AppError {
	w.lock.Lock()
	defer w.lock.Unlock()

	return w.appendInternal(payload, true)
}

// writeInternal executes the write logic under an existing lock.
func (w *BoundFileWriter) writeInternal(payload []byte, isCloseAfter bool) *appfault.AppError {
	if w.path == "" {
		return appfault.New(errtype.Precondition, "file path cannot be empty")
	}

	if w.mode == filewritemodetype.Atomic {
		res := WriteAtomic(w.path, payload, w.perm)
		if res.IsFailed() {
			return res.Fault()
		}

		w.bytesWritten.Add(int64(len(payload)))
		w.writeCount.Add(1)

		return nil
	}

	flags := os.O_CREATE | os.O_WRONLY
	if w.mode == filewritemodetype.Truncate {
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

	if w.isSyncOnWrite {
		if err := f.Sync(); err != nil {
			f.Close()

			return appfault.Wrap(errtype.IO, err, "failed to sync bound file to storage")
		}
	}

	w.bytesWritten.Add(int64(len(payload)))
	w.writeCount.Add(1)

	if isCloseAfter {
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

func (w *BoundFileWriter) appendInternal(payload []byte, isCloseAfter bool) *appfault.AppError {
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
		if isCloseAfter {
			f.Close()
			w.file = nil
		}

		return appfault.Wrap(errtype.IO, writeErr, "failed to append payload to bound file")
	}

	if w.isSyncOnWrite {
		if err := f.Sync(); err != nil {
			if isCloseAfter {
				f.Close()
				w.file = nil
			}

			return appfault.Wrap(errtype.IO, err, "failed to sync bound file after append")
		}
	}

	w.bytesAppended.Add(int64(n))
	w.writeCount.Add(1)

	if isCloseAfter {
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

func (w *BoundFileWriter) Sync() *appfault.AppError {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.file != nil {
		if err := w.file.Sync(); err != nil {
			return appfault.Wrap(errtype.IO, err, "failed to sync open bound file descriptor")
		}
	}

	return nil
}

func (w *BoundFileWriter) Close() *appfault.AppError {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.file != nil {
		err := w.file.Close()
		w.file = nil
		if err != nil {
			return appfault.Wrap(errtype.IO, err, "failed to close bound file descriptor")
		}
	}

	return nil
}

func (w *BoundFileWriter) Name() string {
	return fmt.Sprintf("bound-file-writer[%s]", filepath.Base(w.path))
}

func (w *BoundFileWriter) AsWriter() streamwriter.Writer[[]byte] {
	return w
}

func (w *BoundFileWriter) StdWriter() io.WriteCloser {
	return &boundWriterStdAdapter{writer: w, isAppend: false}
}

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

type fileBoundWriterCreator struct{}

func (fileBoundWriterCreator) Default(path string) *BoundFileWriter {
	return NewBoundFileWriter(path)
}

func (fileBoundWriterCreator) WithOptions(opts BoundFileWriterOptions) *BoundFileWriter {
	return NewBoundFileWriterWithOptions(opts)
}

func (fileBoundWriterCreator) Specific(path string) *BoundFileWriter {
	return NewSpecificFileWriter(path)
}

func (fileBoundWriterCreator) Handler(path string) *BoundFileWriter {
	return NewFileHandler(path)
}

func (fileBoundWriterCreator) AutoClose(path string, perm FilePermType) *BoundFileWriter {
	return NewBoundFileWriterWithOptions(BoundFileWriterOptions{
		Path:        path,
		Perm:        perm,
		IsAutoClose: true,
	})
}

func (fileNewCreator) BoundFileWriter(path string) *BoundFileWriter {
	return NewBoundFileWriter(path)
}
