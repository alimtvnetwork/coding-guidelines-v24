package ostype

func CurrentOsType() Variant {
	detail := GetCurrentOsDetail()
	if detail != nil && detail.OsMixType.IsValid() {
		return detail.OsMixType
	}

	return AnyOs
}

func IsWindows() bool {
	return GetCurrentOsDetail().IsWindows()
}

func IsWindows11() bool {
	return GetCurrentOsDetail().IsWindows11()
}

func IsWindows10() bool {
	return GetCurrentOsDetail().IsWindows10()
}

func IsWindowsServer() bool {
	return GetCurrentOsDetail().IsWindowsServer()
}

func IsUbuntu() bool {
	return GetCurrentOsDetail().IsUbuntu()
}

func IsCentos() bool {
	return GetCurrentOsDetail().IsCentos()
}

func IsDebian() bool {
	return GetCurrentOsDetail().IsDebian()
}

func IsRedhat() bool {
	return GetCurrentOsDetail().IsRedhat()
}

func IsMacOs() bool {
	detail := GetCurrentOsDetail()
	if detail == nil {
		return false
	}

	return detail.IsMacOs
}

func IsDocker() bool {
	detail := GetCurrentOsDetail()
	if detail == nil {
		return false
	}

	return detail.IsDocker
}
