package fileutil

import "fmt"

type FileOpType byte

const (
	FileOpReadOnly FileOpType = iota
	FileOpWriteOnly
	FileOpReadWrite
	FileOpAppend
	FileOpCreate
	FileOpCreateAppend
	FileOpCreateTruncate
	FileOpDelete
)

var fileOpNames = [...]string{
	"ReadOnly",
	"WriteOnly",
	"ReadWrite",
	"Append",
	"Create",
	"CreateAppend",
	"CreateTruncate",
	"Delete",
}

func (o FileOpType) Name() string {
	if int(o) < len(fileOpNames) {
		return fileOpNames[o]
	}

	return fmt.Sprintf("FileOp(%d)", byte(o))
}

func (o FileOpType) String() string {
	return o.Name()
}

func (o FileOpType) IsDelete() bool {
	return o == FileOpDelete
}

func (o FileOpType) IsReadOnly() bool {
	return o == FileOpReadOnly
}

func (o FileOpType) IsAppend() bool {
	if o == FileOpAppend {
		return true
	}

	return o == FileOpCreateAppend
}

func (o FileOpType) createOpenMode() FileOpenModeType {
	switch o {
	case FileOpCreate:
		return FileOpenCreateNew
	case FileOpCreateAppend:
		return FileOpenCreateAppend
	case FileOpCreateTruncate:
		return FileOpenCreateTruncate
	default:
		return FileOpenReadOnly
	}
}

func (o FileOpType) OpenMode() FileOpenModeType {
	switch o {
	case FileOpWriteOnly:
		return FileOpenWriteOnly
	case FileOpReadWrite:
		return FileOpenReadWrite
	case FileOpAppend:
		return FileOpenAppend
	default:
		return o.createOpenMode()
	}
}
