package ostype

import (
	"fmt"
	"runtime"
	"strings"
)

// OperatingSystemDetail captures rich runtime identity and versioning metadata.
type OperatingSystemDetail struct {
	OsMixType     Variant
	Name          string
	ProductName   string
	Vendor        string
	Version       string
	Release       string
	Architecture  string
	WindowsDetail *WindowsSystemDetail
	IsLinux       bool
	IsMacOs       bool
	IsDocker      bool
	Major         int
	Minor         int
	Patch         int
}

// WindowsSystemDetail captures Windows specific installation and versioning info.
type WindowsSystemDetail struct {
	WindowsVersion       int
	ServerVersion        int
	ReleaseId            string
	CurrentBuildId       int
	BuildBranch          string
	InstallType          string
	SystemRoot           string
	Edition              string
	CompositionEdition   string
	GeneratedWindowsName string
	RegisteredOwner      string
	IsServer             bool
	IsClient             bool
}

func (d *OperatingSystemDetail) IsDefined() bool {
	return d != nil
}

func (d *OperatingSystemDetail) IsWindows() bool {
	if d == nil {
		return false
	}

	return d.OsMixType.IsWindows()
}

func (d *OperatingSystemDetail) IsUbuntu() bool {
	if d == nil {
		return false
	}

	return d.OsMixType.IsUbuntu()
}

func (d *OperatingSystemDetail) IsDebian() bool {
	if d == nil {
		return false
	}

	return d.OsMixType.IsDebian()
}

func (d *OperatingSystemDetail) IsCentos() bool {
	if d == nil {
		return false
	}

	return d.OsMixType.IsCentos()
}

func (d *OperatingSystemDetail) IsRedhat() bool {
	if d == nil {
		return false
	}

	return d.OsMixType.IsRedHatEnterpriseLinux()
}

func (d *OperatingSystemDetail) IsWindows11() bool {
	if d == nil || d.WindowsDetail == nil {
		return false
	}

	return d.WindowsDetail.IsWindows11()
}

func (d *OperatingSystemDetail) IsWindows10() bool {
	if d == nil || d.WindowsDetail == nil {
		return false
	}

	return d.WindowsDetail.IsWindows10()
}

func (d *OperatingSystemDetail) IsWindowsServer() bool {
	if d == nil || d.WindowsDetail == nil {
		return false
	}

	return d.WindowsDetail.IsWindowsServer()
}

func (d *OperatingSystemDetail) IsMajorVersionAtLeast(targetMajor int) bool {
	if d == nil {
		return false
	}

	return d.Major >= targetMajor
}

func (d *OperatingSystemDetail) Is64BitArch() bool {
	arch := runtime.GOARCH
	if d != nil && d.Architecture != "" {
		arch = d.Architecture
	}

	return strings.Contains(arch, "64")
}

func (w *WindowsSystemDetail) IsWindows11() bool {
	if w == nil {
		return false
	}

	return w.CurrentBuildId >= 22000
}

func (w *WindowsSystemDetail) IsWindows10() bool {
	if w == nil || w.IsServer {
		return false
	}

	isWin10 := w.WindowsVersion == 10
	isNotWin11 := w.CurrentBuildId < 22000

	return isWin10 && isNotWin11
}

func (w *WindowsSystemDetail) IsWindows8() bool {
	if w == nil || w.IsServer {
		return false
	}

	return w.WindowsVersion == 8
}

func (w *WindowsSystemDetail) IsWindows7() bool {
	if w == nil || w.IsServer {
		return false
	}

	return w.WindowsVersion == 7
}

func (w *WindowsSystemDetail) IsWindowsServer() bool {
	if w == nil {
		return false
	}

	return w.IsServer
}

func (w *WindowsSystemDetail) IsWindowsServerEqual(year int) bool {
	if w == nil || !w.IsServer {
		return false
	}

	return w.ServerVersion == year
}

func (w *WindowsSystemDetail) FormatWindowsName(productName string) string {
	if w == nil {
		return productName
	}

	if w.IsServer {
		return fmt.Sprintf("Windows Server %d %s %d", w.ServerVersion, w.Edition, w.CurrentBuildId)
	}

	if w.IsWindows11() {
		return fmt.Sprintf("Windows 11.%d %s", w.CurrentBuildId, w.Edition)
	}

	if w.IsWindows10() {
		return fmt.Sprintf("Windows 10.%d %s", w.CurrentBuildId, w.Edition)
	}

	return productName
}
