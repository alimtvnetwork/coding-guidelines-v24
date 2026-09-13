package ostype

import (
	"coding-guidelines/common/pkg/baseenumer"
)

var (
	variantLabels = [...]string{
		Invalid:               "Invalid",
		AnyOs:                 "AnyOs",
		Windows:               "Windows",
		Unix:                  "Unix",
		Linux:                 "Linux",
		MacOs:                 "MacOs",
		Ubuntu:                "Ubuntu",
		Debian:                "Debian",
		ArchLinux:             "ArchLinux",
		FreeBsd:               "FreeBsd",
		Centos:                "Centos",
		RedHatEnterpriseLinux: "RedHatEnterpriseLinux",
		Docker:                "Docker",
		Android:               "Android",
	}

	basicEnum = baseenumer.NewBasicInteger(variantLabels[:], Invalid)
)

func init() {
	basicEnum.RegisterAlias("macos", MacOs)
	basicEnum.RegisterAlias("MacOS", MacOs)
	basicEnum.RegisterAlias("darwin", MacOs)
	basicEnum.RegisterAlias("Darwin", MacOs)
	basicEnum.RegisterAlias("mac", MacOs)
	basicEnum.RegisterAlias("win", Windows)
	basicEnum.RegisterAlias("windows", Windows)
	basicEnum.RegisterAlias("unix", Unix)
	basicEnum.RegisterAlias("linux", Linux)
	basicEnum.RegisterAlias("centos", Centos)
	basicEnum.RegisterAlias("CentOS", Centos)
	basicEnum.RegisterAlias("rhel", RedHatEnterpriseLinux)
	basicEnum.RegisterAlias("RHEL", RedHatEnterpriseLinux)
	basicEnum.RegisterAlias("RedHat", RedHatEnterpriseLinux)
	basicEnum.RegisterAlias("freebsd", FreeBsd)
	basicEnum.RegisterAlias("FreeBSD", FreeBsd)
	basicEnum.RegisterAlias("docker", Docker)
	basicEnum.RegisterAlias("all", AnyOs)
	basicEnum.RegisterAlias("any", AnyOs)
	basicEnum.RegisterAlias("default", AnyOs)
}

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

func Parse(s string) (Variant, bool) {
	return basicEnum.Parse(s)
}

func ParseOrZero(s string) Variant {
	return basicEnum.ParseOrZero(s)
}

func ParseOrInvalid(s string) Variant {
	return basicEnum.ParseOrZero(s)
}
