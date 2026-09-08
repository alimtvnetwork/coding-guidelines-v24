package filepermtype

import (
	"os"
	"strconv"
	"strings"

	"coding-guidelines/common/pkg/baseenumer"
)

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
