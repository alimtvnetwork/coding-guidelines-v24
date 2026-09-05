package fileutil

import (
	"fmt"
	"os"
	"strconv"

	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

// FilePermType specifies POSIX filesystem permission bitmasks.
type FilePermType uint32

// Mode converts the enum to standard os.FileMode.
func (p FilePermType) Mode() os.FileMode {
	if p == 0 {
		return os.FileMode(FilePermStandard)
	}

	return os.FileMode(p)
}

// Uint32 returns the raw uint32 permission representation.
func (p FilePermType) Uint32() uint32 {
	return uint32(p)
}

// OctalString returns the 4-digit octal string representation (e.g. "0644").
func (p FilePermType) OctalString() string {
	return fmt.Sprintf("0%o", uint32(p))
}

// PosixString returns the 9-character POSIX notation string (e.g. "rwxr-xr-x").
func (p FilePermType) PosixString() string {
	chars := []byte("---------")
	flags := []uint32{0400, 0200, 0100, 0040, 0020, 0010, 0004, 0002, 0001}
	rwx := "rwxrwxrwx"

	for idx, flag := range flags {
		if (uint32(p) & flag) != 0 {
			chars[idx] = rwx[idx]
		}
	}

	return string(chars)
}

// IsPrivate returns true if group and others have zero permissions.
func (p FilePermType) IsPrivate() bool {
	return (uint32(p) & 0077) == 0
}

// IsPublic returns true if others have read or write permissions.
func (p FilePermType) IsPublic() bool {
	return (uint32(p) & 0007) != 0
}

// IsExecutable returns true if owner, group, or others have execute permissions.
func (p FilePermType) IsExecutable() bool {
	return (uint32(p) & 0111) != 0
}

// IsOwnerReadable returns true if owner has read permissions.
func (p FilePermType) IsOwnerReadable() bool {
	return (uint32(p) & 0400) != 0
}

// IsOwnerWritable returns true if owner has write permissions.
func (p FilePermType) IsOwnerWritable() bool {
	return (uint32(p) & 0200) != 0
}

// IsGroupReadable returns true if group has read permissions.
func (p FilePermType) IsGroupReadable() bool {
	return (uint32(p) & 0040) != 0
}

// IsGroupWritable returns true if group has write permissions.
func (p FilePermType) IsGroupWritable() bool {
	return (uint32(p) & 0020) != 0
}

// IsOtherReadable returns true if others have read permissions.
func (p FilePermType) IsOtherReadable() bool {
	return (uint32(p) & 0004) != 0
}

// IsOtherWritable returns true if others have write permissions.
func (p FilePermType) IsOtherWritable() bool {
	return (uint32(p) & 0002) != 0
}

// WithPrivate returns permissions with group and other access stripped.
func (p FilePermType) WithPrivate() FilePermType {
	return FilePermType(uint32(p) & 0700)
}

// WithReadOnly returns permissions with all write bits removed.
func (p FilePermType) WithReadOnly() FilePermType {
	return FilePermType(uint32(p) &^ 0222)
}

// WithExecutable returns permissions with execute bit added where read is granted.
func (p FilePermType) WithExecutable() FilePermType {
	bits := uint32(p)
	hasOwnerRead := (bits & 0400) != 0
	if hasOwnerRead {
		bits |= 0100
	}

	hasGroupRead := (bits & 0040) != 0
	if hasGroupRead {
		bits |= 0010
	}

	hasOtherRead := (bits & 0004) != 0
	if hasOtherRead {
		bits |= 0001
	}

	return FilePermType(bits)
}

// Name returns the descriptive name and octal representation.
func (p FilePermType) Name() string {
	switch p {
	case FilePermNone:
		return "None(0000)"
	case FilePermOwnerReadOnly:
		return "OwnerReadOnly(0400)"
	case FilePermOwnerWriteOnly:
		return "OwnerWriteOnly(0200)"
	case FilePermOwnerReadWrite:
		return "Private(0600)"
	case FilePermOwnerAll:
		return "OwnerAll(0700)"
	case FilePermGroupReadOnly:
		return "GroupReadOnly(0440)"
	case FilePermGroupWriteOnly:
		return "GroupWriteOnly(0220)"
	case FilePermGroupReadWrite:
		return "GroupReadWrite(0660)"
	case FilePermGroupExec:
		return "GroupExec(0750)"
	case FilePermGroupAll:
		return "GroupAll(0770)"
	case FilePermReadOnly:
		return "ReadOnly(0444)"
	case FilePermPublicWriteOnly:
		return "PublicWriteOnly(0222)"
	case FilePermStandard:
		return "Standard(0644)"
	case FilePermGroupSharedOtherRead:
		return "GroupSharedOtherRead(0664)"
	case FilePermPublicReadWrite:
		return "PublicReadWrite(0666)"
	case FilePermExecutable:
		return "Executable(0755)"
	case FilePermGroupSharedDir:
		return "GroupSharedDir(0775)"
	case FilePermPublicAll:
		return "PublicAll(0777)"
	case FilePermStickyDir:
		return "StickyDir(01777)"
	case FilePermSetuidExec:
		return "SetuidExec(04755)"
	case FilePermSetgidExec:
		return "SetgidExec(02755)"
	default:
		return fmt.Sprintf("Perm(%s)", p.OctalString())
	}
}

// String implements fmt.Stringer returning the Name.
func (p FilePermType) String() string {
	return p.Name()
}

// ParsePerm parses an octal string into FilePermType.
func ParsePerm(octalStr string) result.Wrap[FilePermType] {
	if len(octalStr) == 0 {
		return result.WrapFailureWithId[FilePermType](errtype.Validation, "octal string cannot be empty")
	}

	val, err := strconv.ParseUint(octalStr, 8, 32)
	if err != nil {
		return result.WrapFailureWithCause[FilePermType](errtype.Validation, err, "invalid octal permission: "+octalStr)
	}

	return result.WrapSuccess(FilePermType(val))
}

// FromFileMode converts standard os.FileMode into FilePermType.
func FromFileMode(mode os.FileMode) FilePermType {
	return FilePermType(mode.Perm())
}
