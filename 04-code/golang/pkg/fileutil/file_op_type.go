package fileutil

import (
	"coding-guidelines/common/pkg/enum/fileoptype"
	"coding-guidelines/common/pkg/result"
)

type FileOpType = fileoptype.Variant

const (
	FileOpInvalid        = fileoptype.Invalid
	FileOpReadOnly       = fileoptype.ReadOnly
	FileOpWriteOnly      = fileoptype.WriteOnly
	FileOpReadWrite      = fileoptype.ReadWrite
	FileOpAppend         = fileoptype.Append
	FileOpCreate         = fileoptype.Create
	FileOpCreateAppend   = fileoptype.CreateAppend
	FileOpCreateTruncate = fileoptype.CreateTruncate
	FileOpDelete         = fileoptype.Delete
)

func ParseFileOp(s string) result.Wrap[FileOpType] {
	return fileoptype.Parse(s)
}
