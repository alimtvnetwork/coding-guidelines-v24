package ostype

import (
	"coding-guidelines/common/pkg/baseenumer"
)

type (
	Variant byte

	OSTypeType = Variant

	VariantPredicate func(v Variant) bool
)

const (
	Invalid Variant = iota
	AnyOs
	Windows
	Unix
	Linux
	MacOs
	Ubuntu
	Debian
	ArchLinux
	FreeBsd
	Centos
	RedHatEnterpriseLinux
	Docker
	Android
)

const (
	Unknown = Invalid
	MacOS   = MacOs
	Darwin  = MacOs
	CentOS  = Centos
	RHEL    = RedHatEnterpriseLinux
	FreeBSD = FreeBsd
)

var (
	_ baseenumer.BaseEnumer             = Variant(0)
	_ baseenumer.ByteEnumer             = Variant(0)
	_ baseenumer.NumberEnumer           = Variant(0)
	_ baseenumer.BoundedEnumer[Variant] = Variant(0)
)

func (v Variant) Value() byte {
	return byte(v)
}

func (v Variant) Byte() byte {
	return byte(v)
}

func (v Variant) ValueByte() byte {
	return byte(v)
}

func (v Variant) Bytes() []byte {
	return []byte{byte(v)}
}

func (v Variant) Int() int {
	return int(v)
}

func (v Variant) Code() uint16 {
	return uint16(v)
}

func (v Variant) All() []Variant {
	return All()
}

func (v Variant) Values() []string {
	return Values()
}

func (v Variant) IsValid() bool {
	return baseenumer.IsBetween(v, AnyOs, Android)
}

func (v Variant) IsInvalid() bool {
	return baseenumer.IsNotBetween(v, AnyOs, Android)
}

func (v Variant) IsEnum() bool {
	return v.IsValid()
}

func (v Variant) Min() Variant {
	return basicEnum.Min()
}

func (v Variant) Max() Variant {
	return basicEnum.Max()
}

func (v Variant) IsMin() bool {
	return basicEnum.IsMin(v)
}

func (v Variant) IsMax() bool {
	return basicEnum.IsMax(v)
}

func (v Variant) IsInRange(min, max Variant) bool {
	return baseenumer.IsBetween(v, min, max)
}

func (v Variant) IsAnyOs() bool {
	return v == AnyOs
}

func (v Variant) IsWindows() bool {
	return v == Windows
}

func (v Variant) IsUnix() bool {
	return v == Unix
}

func (v Variant) IsLinux() bool {
	return v == Linux
}

func (v Variant) IsMacOs() bool {
	return v == MacOs
}

func (v Variant) IsUbuntu() bool {
	return v == Ubuntu
}

func (v Variant) IsDebian() bool {
	return v == Debian
}

func (v Variant) IsArchLinux() bool {
	return v == ArchLinux
}

func (v Variant) IsFreeBsd() bool {
	return v == FreeBsd
}

func (v Variant) IsCentos() bool {
	return v == Centos
}

func (v Variant) IsRedHatEnterpriseLinux() bool {
	return v == RedHatEnterpriseLinux
}

func (v Variant) IsDocker() bool {
	return v == Docker
}

func (v Variant) IsAndroid() bool {
	return v == Android
}

func (v Variant) IsUnixLogically() bool {
	return v != Windows && v != Invalid
}

func (v Variant) IsLinuxLogically() bool {
	return v == Linux || v == Ubuntu || v == Debian || v == Centos || v == RedHatEnterpriseLinux || v == ArchLinux
}

func (v Variant) IsAnyOsLogically() bool {
	return v != Invalid
}

func (v Variant) DefaultCmdProcessName() string {
	if v.IsWindows() {
		return "powershell"
	}

	return "bash"
}

func (v Variant) Name() string {
	if int(v) < len(variantLabels) {
		return variantLabels[v]
	}

	return baseenumer.FormatNameValue("OSType", byte(v))
}

func (v Variant) Label() string {
	return v.Name()
}

func (v Variant) String() string {
	return v.Name()
}

func (v Variant) ValueString() string {
	return baseenumer.FormatNameValue(v.Name(), byte(v))
}

func (v Variant) MarshalJSON() ([]byte, error) {
	return baseenumer.MarshalJSON(v.Name())
}

func (v *Variant) UnmarshalJSON(data []byte) error {
	return basicEnum.UnmarshalJSON(data, v)
}
