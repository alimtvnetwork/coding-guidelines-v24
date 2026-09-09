package filepermtype

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"coding-guidelines/common/pkg/baseenumer"
)

const (
	DefaultPosixTemplate = "---------"
	PosixRwxPattern      = "rwxrwxrwx"

	MaskOwnerRead  uint32 = 0400
	MaskOwnerWrite uint32 = 0200
	MaskOwnerExec  uint32 = 0100
	MaskGroupRead  uint32 = 0040
	MaskGroupWrite uint32 = 0020
	MaskGroupExec  uint32 = 0010
	MaskOtherRead  uint32 = 0004
	MaskOtherWrite uint32 = 0002
	MaskOtherExec  uint32 = 0001

	MaskExecAll     uint32 = 0111
	MaskGroupOther  uint32 = 0077
	MaskOtherOnly   uint32 = 0007
	MaskStandard    uint32 = 0777
	MaskOwnerAll    uint32 = 0700
	MaskWriteAll    uint32 = 0222
	MaskAllPermBits uint32 = 07777
	MaxStandardPerm        = 512
)

var (
	posixFlags = [9]uint32{0400, 0200, 0100, 0040, 0020, 0010, 0004, 0002, 0001}

	posixStringTable = initPosixStringTable()
	octalStringTable = initOctalStringTable()
)

func initPosixStringTable() [MaxStandardPerm]string {
	var table [MaxStandardPerm]string
	for perm := 0; perm < MaxStandardPerm; perm++ {
		table[perm] = computePosixString(uint32(perm))
	}

	return table
}

func computePosixString(perm uint32) string {
	chars := []byte(DefaultPosixTemplate)
	for idx, flag := range posixFlags {
		if (perm & flag) != 0 {
			chars[idx] = PosixRwxPattern[idx]
		}
	}

	return string(chars)
}

func initOctalStringTable() [MaxStandardPerm]string {
	var table [MaxStandardPerm]string
	for perm := 0; perm < MaxStandardPerm; perm++ {
		table[perm] = fmt.Sprintf("0%o", perm)
	}

	return table
}

var (
	allVariants = []Variant{
		None,
		OwnerReadOnly,
		OwnerWriteOnly,
		OwnerExecOnly,
		OwnerReadWrite,
		OwnerAll,
		GroupReadOnly,
		GroupWriteOnly,
		GroupReadWrite,
		GroupExec,
		GroupAll,
		ReadOnly,
		PublicWriteOnly,
		Standard,
		GroupSharedOtherRead,
		PublicReadWrite,
		Executable,
		GroupSharedDir,
		PublicAll,
		StickyDir,
		SetuidExec,
		SetgidExec,
	}

	allValues = []string{
		"None",
		"OwnerReadOnly",
		"OwnerWriteOnly",
		"OwnerExecOnly",
		"OwnerReadWrite",
		"OwnerAll",
		"GroupReadOnly",
		"GroupWriteOnly",
		"GroupReadWrite",
		"GroupExec",
		"GroupAll",
		"ReadOnly",
		"PublicWriteOnly",
		"Standard",
		"GroupSharedOtherRead",
		"PublicReadWrite",
		"Executable",
		"GroupSharedDir",
		"PublicAll",
		"StickyDir",
		"SetuidExec",
		"SetgidExec",
	}

	basicEnum = baseenumer.NewBasicSparseInteger(allVariants, allValues, None, 07777)
)

func All() []Variant {
	return basicEnum.All()
}

func Values() []string {
	return basicEnum.Values()
}

func Min() Variant {
	return basicEnum.Min()
}

func Max() Variant {
	return basicEnum.Max()
}

func Parse(octalStr string) (Variant, bool) {
	trimmed := strings.TrimSpace(octalStr)
	if len(trimmed) == 0 {
		return None, false
	}

	val, err := strconv.ParseUint(trimmed, 8, 32)
	if err != nil {
		return None, false
	}

	return Variant(val), true
}

func ParsePerm(octalStr string) (Variant, bool) {
	return Parse(octalStr)
}

func ParseOrZero(octalStr string) Variant {
	v, isOk := Parse(octalStr)
	if isOk {
		return v
	}

	return None
}

func ParseOrInvalid(octalStr string) Variant {
	return ParseOrZero(octalStr)
}

func ParseOrUnknown(octalStr string) Variant {
	return ParseOrZero(octalStr)
}

func FromFileMode(mode os.FileMode) Variant {
	return Variant(mode.Perm())
}

func PosixString(p Variant) string {
	return p.PosixString()
}

func OctalString(p Variant) string {
	return p.OctalString()
}
