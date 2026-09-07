package appwriter

import (
	"context"
	"io"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/fileutil"
	"coding-guidelines/common/pkg/payloadconv"
)

// FileWriterOptions specifies configuration for file-backed writers.
type FileWriterOptions struct {
	Name     string
	FilePath string
	OpenMode fileutil.FileOpenModeType
	PermMode fileutil.FilePermType
	IsLocked bool
}

func resolveFileModes(openMode fileutil.FileOpenModeType, permMode fileutil.FilePermType) (fileutil.FileOpenModeType, fileutil.FilePermType) {
	if openMode == 0 {
		openMode = fileutil.FileOpenCreateAppend
	}

	if permMode == 0 {
		permMode = fileutil.FilePermStandard
	}

	return openMode, permMode
}

func resolveWriterName(name, filePath string) string {
	if len(name) == 0 {
		return filePath
	}

	return name
}

// NewFileWriter creates a file writer using fileutil enums and wrap constructors.
func NewFileWriter(opts FileWriterOptions) BaseWriterWrap {
	if len(opts.FilePath) == 0 {
		return WrapWriter.FailureWithId(errtype.Validation, "file path cannot be empty")
	}

	openMode, permMode := resolveFileModes(opts.OpenMode, opts.PermMode)
	fileWrap := fileutil.OpenFile(opts.FilePath, openMode, permMode)
	if fileWrap.IsFailed() {
		return WrapWriter.FailureFromWrap(fileWrap)
	}

	name := resolveWriterName(opts.Name, opts.FilePath)
	writer := NewBaseWriter(name, fileWrap.Data(), opts.IsLocked, fileWriteFunc)

	return WrapWriter.Success(writer)
}

func validateWriterDest(self Writer) (io.Writer, *appfault.AppError) {
	if self == nil {
		return nil, appfault.New(errtype.Precondition, "writer is unassigned")
	}

	dest := self.Destination()
	if dest == nil {
		return nil, appfault.New(errtype.Precondition, "writer destination is unassigned")
	}

	return dest, nil
}

func writeBytesToDest(dest io.Writer, data []byte) *appfault.AppError {
	if _, err := dest.Write(data); err != nil {
		return appfault.Wrap(errtype.IO, err, "failed to write payload to file destination")
	}

	return nil
}

func fileWriteFunc(ctx context.Context, self Writer, payload any) *appfault.AppError {
	dest, destFault := validateWriterDest(self)
	if destFault != nil {
		return destFault
	}

	res := payloadconv.ToBytes(payload)
	if res.IsFailure() {
		return res.Fault()
	}

	return writeBytesToDest(dest, res.Data())
}
