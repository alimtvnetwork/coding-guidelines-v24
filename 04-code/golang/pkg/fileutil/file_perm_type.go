package fileutil

import (
	"os"

	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/result"
)

type FilePermType = filepermtype.Variant

const (
	FilePermNone                 = filepermtype.None
	FilePermOwnerReadOnly        = filepermtype.OwnerReadOnly
	FilePermOwnerWriteOnly       = filepermtype.OwnerWriteOnly
	FilePermOwnerExecOnly        = filepermtype.OwnerExecOnly
	FilePermOwnerReadWrite       = filepermtype.OwnerReadWrite
	FilePermPrivate              = filepermtype.Private
	FilePermOwnerAll             = filepermtype.OwnerAll
	FilePermOwnerExec            = filepermtype.OwnerExec
	FilePermGroupReadOnly        = filepermtype.GroupReadOnly
	FilePermGroupWriteOnly       = filepermtype.GroupWriteOnly
	FilePermGroupReadWrite       = filepermtype.GroupReadWrite
	FilePermGroupExec            = filepermtype.GroupExec
	FilePermGroupAll             = filepermtype.GroupAll
	FilePermReadOnly             = filepermtype.ReadOnly
	FilePermPublicReadOnly       = filepermtype.PublicReadOnly
	FilePermPublicWriteOnly      = filepermtype.PublicWriteOnly
	FilePermStandard             = filepermtype.Standard
	FilePermGroupSharedOtherRead = filepermtype.GroupSharedOtherRead
	FilePermPublicReadWrite      = filepermtype.PublicReadWrite
	FilePermExecutable           = filepermtype.Executable
	FilePermGroupSharedDir       = filepermtype.GroupSharedDir
	FilePermPublicAll            = filepermtype.PublicAll
	FilePermStickyDir            = filepermtype.StickyDir
	FilePermSetuidExec           = filepermtype.SetuidExec
	FilePermSetgidExec           = filepermtype.SetgidExec
)

func ParsePerm(octalStr string) result.Wrap[FilePermType] {
	return filepermtype.ParsePerm(octalStr)
}

func FromFileMode(mode os.FileMode) FilePermType {
	return filepermtype.FromFileMode(mode)
}
