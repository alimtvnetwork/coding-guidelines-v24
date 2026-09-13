package ostype

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"coding-guidelines/common/pkg/regexnew"
)

var (
	detailOnce       sync.Once
	cachedHostDetail *OperatingSystemDetail
)

func GetCurrentOsDetail() *OperatingSystemDetail {
	detailOnce.Do(func() {
		cachedHostDetail = detectCurrentOS()
	})

	return cachedHostDetail
}

func detectCurrentOS() *OperatingSystemDetail {
	switch runtime.GOOS {
	case "windows":
		return detectWindowsHost()
	case "darwin":
		return detectDarwinHost()
	default:
		return detectLinuxHost()
	}
}

func detectWindowsHost() *OperatingSystemDetail {
	winDetail, _ := queryWindowsRegistry()
	name := "Windows"
	if winDetail != nil && winDetail.GeneratedWindowsName != "" {
		name = winDetail.GeneratedWindowsName
	}

	return &OperatingSystemDetail{
		OsMixType:     Windows,
		Name:          name,
		ProductName:   name,
		Vendor:        "microsoft",
		WindowsDetail: winDetail,
		Architecture:  runtime.GOARCH,
		IsDocker:      checkIsDocker(),
	}
}

func detectDarwinHost() *OperatingSystemDetail {
	out, err := exec.Command("sw_vers").CombinedOutput()
	if err == nil && len(out) > 0 {
		d := ParseMacOsOutput(string(out))
		d.IsDocker = checkIsDocker()

		return d
	}

	return &OperatingSystemDetail{
		OsMixType:    MacOs,
		Name:         "macOS",
		Vendor:       "macos",
		Architecture: runtime.GOARCH,
		IsMacOs:      true,
		IsDocker:     checkIsDocker(),
	}
}

func detectLinuxHost() *OperatingSystemDetail {
	data, err := os.ReadFile("/etc/os-release")
	if err == nil && len(data) > 0 {
		d := ParseOSReleaseContent(string(data))
		enrichLinuxReleaseFiles(d)
		d.IsDocker = checkIsDocker()

		return d
	}

	return &OperatingSystemDetail{
		OsMixType:    Linux,
		Name:         "Linux",
		Vendor:       "linux",
		Architecture: runtime.GOARCH,
		IsLinux:      true,
		IsDocker:     checkIsDocker(),
	}
}

func enrichLinuxReleaseFiles(d *OperatingSystemDetail) {
	if d.OsMixType == Debian {
		if content, err := os.ReadFile("/etc/debian_version"); err == nil {
			d.Release = strings.TrimSpace(string(content))
		}
	} else if d.OsMixType == Centos {
		if content, err := os.ReadFile("/etc/centos-release"); err == nil {
			d.Release = ParseCentOSRelease(string(content))
		}
	} else if d.OsMixType == RedHatEnterpriseLinux {
		if content, err := os.ReadFile("/etc/redhat-release"); err == nil {
			d.Release = ParseRedHatRelease(string(content))
		}
	}
}

func checkIsDocker() bool {
	_, err := os.Stat("/.dockerenv")

	return err == nil
}

func ParseWindowsDetail(productName, installType, edition string, buildNum int) *WindowsSystemDetail {
	isServer := strings.EqualFold(installType, "Server")
	winVer, srvVer := extractWinVersions(productName, isServer)

	detail := &WindowsSystemDetail{
		WindowsVersion: winVer,
		ServerVersion:  srvVer,
		CurrentBuildId: buildNum,
		InstallType:    installType,
		Edition:        edition,
		IsServer:       isServer,
		IsClient:       !isServer,
	}

	detail.GeneratedWindowsName = detail.FormatWindowsName(productName)

	return detail
}

func extractWinVersions(productName string, isServer bool) (int, int) {
	re := regexnew.WindowsVersionNumberCheckerRegex.CompileMust()
	m := re.FindString(productName)
	ver, _ := strconv.Atoi(m)
	if isServer {
		return 0, ver
	}

	return ver, 0
}

func ParseOSReleaseContent(content string) *OperatingSystemDetail {
	lines := strings.Split(content, "\n")
	fields := parseKeyValueLines(lines)
	name := fields["PRETTY_NAME"]
	vendor := strings.ToLower(fields["ID"])
	version := fields["VERSION_ID"]

	detail := &OperatingSystemDetail{
		OsMixType:    Linux,
		Name:         name,
		ProductName:  name,
		Vendor:       vendor,
		Version:      version,
		Release:      version,
		Architecture: runtime.GOARCH,
		IsLinux:      true,
	}

	enrichLinuxDetail(detail, name, vendor)
	parseVersionNumbers(detail, detail.Release)

	return detail
}

func parseKeyValueLines(lines []string) map[string]string {
	result := make(map[string]string, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		parts := strings.SplitN(trimmed, "=", 2)
		if len(parts) == 2 {
			k := strings.TrimSpace(parts[0])
			v := strings.Trim(strings.TrimSpace(parts[1]), "\"'")
			result[k] = v
		}
	}

	return result
}

func enrichLinuxDetail(detail *OperatingSystemDetail, name, vendor string) {
	if vendor == "ubuntu" || strings.Contains(strings.ToLower(name), "ubuntu") {
		detail.OsMixType = Ubuntu
		re := regexnew.UbuntuNameCheckerRegex.CompileMust()
		if m := re.FindStringSubmatch(name); len(m) > 1 {
			detail.Release = m[1]
		}
	} else if vendor == "debian" {
		detail.OsMixType = Debian
	} else if vendor == "centos" {
		detail.OsMixType = Centos
	} else if vendor == "rhel" {
		detail.OsMixType = RedHatEnterpriseLinux
	}
}

func parseVersionNumbers(detail *OperatingSystemDetail, rel string) {
	parts := strings.Split(rel, ".")
	if len(parts) > 0 {
		detail.Major, _ = strconv.Atoi(parts[0])
	}

	if len(parts) > 1 {
		detail.Minor, _ = strconv.Atoi(parts[1])
	}

	if len(parts) > 2 {
		detail.Patch, _ = strconv.Atoi(parts[2])
	}
}

func ParseCentOSRelease(content string) string {
	re := regexnew.CentOsNameCheckerRegex.CompileMust()
	m := re.FindStringSubmatch(content)
	if len(m) > 2 {
		return m[2]
	}

	return strings.TrimSpace(content)
}

func ParseRedHatRelease(content string) string {
	re := regexnew.RedHatNameCheckerRegex.CompileMust()
	m := re.FindStringSubmatch(content)
	if len(m) > 1 {
		return m[1]
	}

	return strings.TrimSpace(content)
}

func ParseMacOsOutput(output string) *OperatingSystemDetail {
	lines := strings.Split(output, "\n")
	data := make(map[string]string)
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			data[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	name := data["ProductName"]
	ver := data["ProductVersion"]
	build := data["BuildVersion"]

	detail := &OperatingSystemDetail{
		OsMixType:    MacOs,
		Name:         name,
		ProductName:  name,
		Vendor:       "macos",
		Version:      ver,
		Release:      build,
		Architecture: runtime.GOARCH,
		IsMacOs:      true,
	}

	parseVersionNumbers(detail, ver)

	return detail
}
