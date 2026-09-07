package fileutil

import (
	"path/filepath"

	"coding-guidelines/common/pkg/enum/filewritemodetype"
	"coding-guidelines/common/pkg/enum/openfiletype"
)

type fileNamespace struct {
	Open   openOps
	Create createOps
	Write  writeOps
	Append appendOps
	Read   readOps
	Path   pathNamespace
}

var File = &fileNamespace{
	Open:   openOps{},
	Create: createOps{},
	Write:  writeOps{},
	Append: appendOps{},
	Read:   readOps{},
	Path:   Path,
}

type (
	fileNewCreator struct {
		Writer       fileWriterCreator
		Appender     fileAppenderCreator
		BoundWriter  fileBoundWriterCreator
		Path         filePathCreator
		StreamWriter fileStreamWriterCreator
	}

	fileWriterCreator       struct{}
	fileAppenderCreator     struct{}
	fileBoundWriterCreator  struct{}
	filePathCreator         struct{}
	fileStreamWriterCreator struct{}
)

var New = &fileNewCreator{
	Writer:       fileWriterCreator{},
	Appender:     fileAppenderCreator{},
	BoundWriter:  fileBoundWriterCreator{},
	Path:         filePathCreator{},
	StreamWriter: fileStreamWriterCreator{},
}

// Writer creators

func (fileWriterCreator) Default(path string) *FileWriter {
	return NewFileWriterEngine(path)
}

func (fileWriterCreator) Engine(path string) *FileWriter {
	return NewFileWriterEngine(path)
}

func (fileWriterCreator) WithOptions(opts FileWriterOptions) *FileWriter {
	return NewFileWriterWithOptions(opts)
}

func (fileWriterCreator) Atomic(path string, perm FilePermType) *FileWriter {
	return NewFileWriterWithOptions(FileWriterOptions{
		Path: path,
		Mode: filewritemodetype.Atomic,
		Perm: perm,
	})
}

func (fileWriterCreator) Direct(path string, perm FilePermType) *FileWriter {
	return NewFileWriterWithOptions(FileWriterOptions{
		Path: path,
		Mode: filewritemodetype.Direct,
		Perm: perm,
	})
}

func (fileWriterCreator) Truncate(path string, perm FilePermType) *FileWriter {
	return NewFileWriterWithOptions(FileWriterOptions{
		Path: path,
		Mode: filewritemodetype.Truncate,
		Perm: perm,
	})
}

// Appender creators

func (fileAppenderCreator) Default(path string, perm FilePermType) *FileAppender {
	return NewFileAppender(path, perm)
}

func (fileAppenderCreator) AutoSync(path string, perm FilePermType) *FileAppender {
	appender := NewFileAppender(path, perm)
	appender.SetAutoSync(true)

	return appender
}

// BoundWriter creators

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

// Path creators

func (filePathCreator) Default(raw string) *PathWrapper {
	return NewPath(raw)
}

func (filePathCreator) FromParts(elem ...string) *PathWrapper {
	return NewPath(filepath.Join(elem...))
}

// StreamWriter creators

func (fileStreamWriterCreator) Any(path string, openMode FileOpenModeType, perm FilePermType) FileWriterResult {
	return NewFileWriter(path, openMode, perm)
}

func (fileStreamWriterCreator) Append(path string, perm FilePermType) FileWriterResult {
	return NewFileWriter(path, openfiletype.CreateAppend, perm)
}

func (fileStreamWriterCreator) Truncate(path string, perm FilePermType) FileWriterResult {
	return NewFileWriter(path, openfiletype.CreateTruncate, perm)
}

// Shortcut helpers on fileNewCreator

func (fileNewCreator) FileWriter(path string) *FileWriter {
	return NewFileWriterEngine(path)
}

func (fileNewCreator) FileAppender(path string, perm FilePermType) *FileAppender {
	return NewFileAppender(path, perm)
}

func (fileNewCreator) BoundFileWriter(path string) *BoundFileWriter {
	return NewBoundFileWriter(path)
}

func (fileNewCreator) PathWrapper(raw string) *PathWrapper {
	return NewPath(raw)
}

func (fileNewCreator) StreamWriterAny(path string, openMode FileOpenModeType, perm FilePermType) FileWriterResult {
	return NewFileWriter(path, openMode, perm)
}

func (fileNamespace) Target(path string) *FilePathOps {
	return NewFilePathOps(path)
}

func (fileNamespace) At(workDir, relPath string) *FilePathOps {
	return NewFilePathOpsAt(workDir, relPath)
}

func (fileNewCreator) Target(path string) *FilePathOps {
	return NewFilePathOps(path)
}

func (fileNewCreator) At(workDir, relPath string) *FilePathOps {
	return NewFilePathOpsAt(workDir, relPath)
}

func (fileNewCreator) FilePathOps(path string) *FilePathOps {
	return NewFilePathOps(path)
}

func (fileNewCreator) FilePathOpsAt(workDir, relPath string) *FilePathOps {
	return NewFilePathOpsAt(workDir, relPath)
}
