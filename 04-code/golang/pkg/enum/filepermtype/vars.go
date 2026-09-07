package filepermtype

import (
	"os"
	"strconv"
	"strings"

	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type Result = result.Wrap[Variant]

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
)

func All() []Variant {
	return append([]Variant(nil), allVariants...)
}

func Values() []string {
	return append([]string(nil), allValues...)
}

func Parse(octalStr string) Result {
	trimmed := strings.TrimSpace(octalStr)
	if len(trimmed) == 0 {
		return result.WrapFailureWithId[Variant](errtype.Validation, "octal string cannot be empty")
	}

	val, err := strconv.ParseUint(trimmed, 8, 32)
	if err != nil {
		return result.WrapFailureWithCause[Variant](errtype.Validation, err, "invalid octal permission: "+octalStr)
	}

	return result.WrapSuccess(Variant(val))
}

func ParsePerm(octalStr string) Result {
	return Parse(octalStr)
}

func FromFileMode(mode os.FileMode) Variant {
	return Variant(mode.Perm())
}
