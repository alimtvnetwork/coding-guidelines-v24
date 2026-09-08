package filepermtype

import (
	"encoding/json"
	"fmt"
	"os"

	"coding-guidelines/common/pkg/baseenumer"
)

type (
	Variant uint32

	FilePermType = Variant

	VariantPredicate func(v Variant) bool
)

const (
	None                 Variant = 0000
	OwnerReadOnly        Variant = 0400
	OwnerWriteOnly       Variant = 0200
	OwnerExecOnly        Variant = 0100
	OwnerReadWrite       Variant = 0600
	Private              Variant = 0600
	OwnerAll             Variant = 0700
	OwnerExec            Variant = 0700
	GroupReadOnly        Variant = 0440
	GroupWriteOnly       Variant = 0220
	GroupReadWrite       Variant = 0660
	GroupExec            Variant = 0750
	GroupAll             Variant = 0770
	ReadOnly             Variant = 0444
	PublicReadOnly       Variant = 0444
	PublicWriteOnly      Variant = 0222
	Standard             Variant = 0644
	GroupSharedOtherRead Variant = 0664
	PublicReadWrite      Variant = 0666
	Executable           Variant = 0755
	GroupSharedDir       Variant = 0775
	PublicAll            Variant = 0777
	StickyDir            Variant = 01777
	SetuidExec           Variant = 04755
	SetgidExec           Variant = 02755
)

var knownNames = map[Variant]string{
	None:                 "None(0000)",
	OwnerReadOnly:        "OwnerReadOnly(0400)",
	OwnerWriteOnly:       "OwnerWriteOnly(0200)",
	Private:              "Private(0600)",
	OwnerAll:             "OwnerAll(0700)",
	GroupReadOnly:        "GroupReadOnly(0440)",
	GroupWriteOnly:       "GroupWriteOnly(0220)",
	GroupReadWrite:       "GroupReadWrite(0660)",
	GroupExec:            "GroupExec(0750)",
	GroupAll:             "GroupAll(0770)",
	ReadOnly:             "ReadOnly(0444)",
	PublicWriteOnly:      "PublicWriteOnly(0222)",
	Standard:             "Standard(0644)",
	GroupSharedOtherRead: "GroupSharedOtherRead(0664)",
	PublicReadWrite:      "PublicReadWrite(0666)",
	Executable:           "Executable(0755)",
	GroupSharedDir:       "GroupSharedDir(0775)",
	PublicAll:            "PublicAll(0777)",
	StickyDir:            "StickyDir(01777)",
	SetuidExec:           "SetuidExec(04755)",
	SetgidExec:           "SetgidExec(02755)",
}

func (p Variant) Mode() os.FileMode {
	if p == 0 {
		return os.FileMode(Standard)
	}

	return os.FileMode(p)
}

func (p Variant) Uint32() uint32 {
	return uint32(p)
}

func (p Variant) Int() int {
	return int(p)
}

func (p Variant) Code() uint16 {
	return uint16(p)
}

func (p Variant) OctalString() string {
	return fmt.Sprintf("0%o", uint32(p))
}

func (p Variant) PosixString() string {
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

func (p Variant) IsPrivate() bool {
	return (uint32(p) & 0077) == 0
}

func (p Variant) IsPublic() bool {
	return (uint32(p) & 0007) != 0
}

func (p Variant) IsExecutable() bool {
	return (uint32(p) & 0111) != 0
}

func (p Variant) IsOwnerReadable() bool {
	return (uint32(p) & 0400) != 0
}

func (p Variant) IsOwnerWritable() bool {
	return (uint32(p) & 0200) != 0
}

func (p Variant) IsGroupReadable() bool {
	return (uint32(p) & 0040) != 0
}

func (p Variant) IsGroupWritable() bool {
	return (uint32(p) & 0020) != 0
}

func (p Variant) IsOtherReadable() bool {
	return (uint32(p) & 0004) != 0
}

func (p Variant) IsOtherWritable() bool {
	return (uint32(p) & 0002) != 0
}

func (p Variant) IsValid() bool {
	return p <= 07777
}

func (p Variant) IsEnum() bool {
	return p.IsValid()
}

func (p Variant) WithPrivate() Variant {
	return Variant(uint32(p) & 0700)
}

func (p Variant) WithReadOnly() Variant {
	return Variant(uint32(p) &^ 0222)
}

func (p Variant) WithExecutable() Variant {
	bits := uint32(p)
	bits = applyOwnerExec(bits)
	bits = applyGroupExec(bits)
	bits = applyOtherExec(bits)

	return Variant(bits)
}

func applyOwnerExec(bits uint32) uint32 {
	if (bits & 0400) != 0 {
		return bits | 0100
	}

	return bits
}

func applyGroupExec(bits uint32) uint32 {
	if (bits & 0040) != 0 {
		return bits | 0010
	}

	return bits
}

func applyOtherExec(bits uint32) uint32 {
	if (bits & 0004) != 0 {
		return bits | 0001
	}

	return bits
}

func (p Variant) Name() string {
	if name, ok := knownNames[p]; ok {
		return name
	}

	return fmt.Sprintf("Perm(%s)", p.OctalString())
}

func (p Variant) String() string {
	return p.Name()
}

func (p Variant) ValueString() string {
	return p.OctalString()
}

func (p Variant) MarshalJSON() ([]byte, error) {
	return baseenumer.MarshalJSON(p.OctalString())
}

func (p *Variant) UnmarshalJSON(data []byte) error {
	return baseenumer.UnmarshalIntegerJSON(data, p, nil, 07777, Standard)
}

var (
	_ baseenumer.BaseEnumer   = Variant(0)
	_ baseenumer.NumberEnumer = Variant(0)
	_ json.Marshaler          = Variant(0)
	_ json.Unmarshaler        = (*Variant)(nil)
)
