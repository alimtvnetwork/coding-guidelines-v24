package ostype_test

import (
	"runtime"
	"testing"

	"coding-guidelines/common/pkg/enum/ostype"
)

func TestGetCurrentOsDetail(t *testing.T) {
	detail := ostype.GetCurrentOsDetail()
	if detail == nil {
		t.Fatalf("expected non-nil detail from GetCurrentOsDetail()")
	}

	if !detail.IsDefined() {
		t.Fatalf("expected detail.IsDefined() to be true")
	}

	if detail.Architecture != runtime.GOARCH {
		t.Errorf("expected Architecture %s, got %s", runtime.GOARCH, detail.Architecture)
	}

	if !detail.Is64BitArch() && runtime.GOARCH == "amd64" {
		t.Errorf("expected 64-bit architecture for amd64")
	}
}

func TestCurrentOsType(t *testing.T) {
	curr := ostype.CurrentOsType()
	if !curr.IsValid() {
		t.Fatalf("expected valid CurrentOsType, got %v", curr)
	}

	if runtime.GOOS == "windows" && !curr.IsWindows() {
		t.Errorf("expected Windows on windows GOOS")
	}
}

func TestPureUbuntuParser(t *testing.T) {
	fixture := "NAME=\"Ubuntu\"\n" +
		"VERSION=\"22.04.1 LTS (Jammy Jellyfish)\"\n" +
		"ID=ubuntu\n" +
		"ID_LIKE=debian\n" +
		"PRETTY_NAME=\"Ubuntu 22.04.1 LTS\"\n" +
		"VERSION_ID=\"22.04\"\n"

	detail := ostype.ParseOSReleaseContent(fixture)
	if !detail.IsUbuntu() {
		t.Errorf("expected IsUbuntu() to be true")
	}

	if !detail.IsLinux {
		t.Errorf("expected IsLinux to be true")
	}

	if detail.Major != 22 || detail.Minor != 4 {
		t.Errorf("expected version 22.4, got %d.%d", detail.Major, detail.Minor)
	}

	if detail.Release != "22.04.1" {
		t.Errorf("expected release 22.04.1, got %s", detail.Release)
	}

	if !detail.IsMajorVersionAtLeast(20) {
		t.Errorf("expected major version at least 20")
	}

	if detail.IsMajorVersionAtLeast(24) {
		t.Errorf("did not expect major version at least 24")
	}
}

func TestPureDebianParser(t *testing.T) {
	fixture := "PRETTY_NAME=\"Debian GNU/Linux 11 (bullseye)\"\n" +
		"NAME=\"Debian GNU/Linux\"\n" +
		"VERSION_ID=\"11\"\n" +
		"VERSION=\"11 (bullseye)\"\n" +
		"ID=debian\n"

	detail := ostype.ParseOSReleaseContent(fixture)
	if !detail.IsDebian() {
		t.Errorf("expected IsDebian() to be true")
	}

	if detail.Major != 11 {
		t.Errorf("expected major version 11, got %d", detail.Major)
	}
}

func TestPureCentOSParser(t *testing.T) {
	releaseStr := ostype.ParseCentOSRelease("CentOS Linux release 7.2.1511 (Core)")
	if releaseStr != "7.2.1511" {
		t.Errorf("expected 7.2.1511, got %s", releaseStr)
	}

	fixture := "NAME=\"CentOS Linux\"\n" +
		"VERSION=\"7 (Core)\"\n" +
		"ID=\"centos\"\n" +
		"PRETTY_NAME=\"CentOS Linux 7 (Core)\"\n" +
		"VERSION_ID=\"7\"\n"

	detail := ostype.ParseOSReleaseContent(fixture)
	if !detail.IsCentos() {
		t.Errorf("expected IsCentos() to be true")
	}
}

func TestPureRedHatParser(t *testing.T) {
	releaseStr := ostype.ParseRedHatRelease("Red Hat Enterprise Linux Server release 8.4 (Ootpa)")
	if releaseStr != "8.4" {
		t.Errorf("expected 8.4, got %s", releaseStr)
	}

	fixture := "NAME=\"Red Hat Enterprise Linux\"\n" +
		"VERSION=\"8.4 (Ootpa)\"\n" +
		"ID=\"rhel\"\n" +
		"PRETTY_NAME=\"Red Hat Enterprise Linux 8.4 (Ootpa)\"\n" +
		"VERSION_ID=\"8.4\"\n"

	detail := ostype.ParseOSReleaseContent(fixture)
	if !detail.IsRedhat() {
		t.Errorf("expected IsRedhat() to be true")
	}

	if detail.Major != 8 || detail.Minor != 4 {
		t.Errorf("expected 8.4, got %d.%d", detail.Major, detail.Minor)
	}
}

func TestPureMacOsParser(t *testing.T) {
	sampleOutput := "ProductName:\tMac OS X\nProductVersion:\t10.15.7\nBuildVersion:\t19H524\n"
	detail := ostype.ParseMacOsOutput(sampleOutput)

	if !detail.IsMacOs {
		t.Errorf("expected IsMacOs to be true")
	}

	if detail.Major != 10 || detail.Minor != 15 || detail.Patch != 7 {
		t.Errorf("expected 10.15.7, got %d.%d.%d", detail.Major, detail.Minor, detail.Patch)
	}

	if detail.Release != "19H524" {
		t.Errorf("expected build 19H524, got %s", detail.Release)
	}
}

func TestPureWindows11Parser(t *testing.T) {
	detail := ostype.ParseWindowsDetail("Windows 10 Pro", "Client", "Professional", 22631)
	if !detail.IsWindows11() {
		t.Errorf("expected Windows 11 for build 22631")
	}

	if detail.IsWindows10() {
		t.Errorf("expected IsWindows10() to be false for Windows 11")
	}

	if detail.IsServer {
		t.Errorf("expected client, got server")
	}

	expectedName := "Windows 11.22631 Professional"
	if detail.GeneratedWindowsName != expectedName {
		t.Errorf("expected %s, got %s", expectedName, detail.GeneratedWindowsName)
	}
}

func TestPureWindows10Parser(t *testing.T) {
	detail := ostype.ParseWindowsDetail("Windows 10 Enterprise", "Client", "Enterprise", 19045)
	if !detail.IsWindows10() {
		t.Errorf("expected Windows 10 for build 19045")
	}

	if detail.IsWindows11() {
		t.Errorf("expected IsWindows11() to be false for build 19045")
	}

	expectedName := "Windows 10.19045 Enterprise"
	if detail.GeneratedWindowsName != expectedName {
		t.Errorf("expected %s, got %s", expectedName, detail.GeneratedWindowsName)
	}
}

func TestPureWindowsServerParser(t *testing.T) {
	detail := ostype.ParseWindowsDetail("Windows Server 2019 Standard", "Server", "ServerStandard", 17763)
	if !detail.IsWindowsServer() {
		t.Errorf("expected IsWindowsServer() to be true")
	}

	if !detail.IsWindowsServerEqual(2019) {
		t.Errorf("expected ServerVersion 2019")
	}

	if detail.IsClient {
		t.Errorf("expected IsClient to be false")
	}

	expectedName := "Windows Server 2019 ServerStandard 17763"
	if detail.GeneratedWindowsName != expectedName {
		t.Errorf("expected %s, got %s", expectedName, detail.GeneratedWindowsName)
	}
}

func TestQuickHelpers(t *testing.T) {
	_ = ostype.IsWindows()
	_ = ostype.IsWindows11()
	_ = ostype.IsWindows10()
	_ = ostype.IsWindowsServer()
	_ = ostype.IsUbuntu()
	_ = ostype.IsCentos()
	_ = ostype.IsDebian()
	_ = ostype.IsRedhat()
	_ = ostype.IsMacOs()
	_ = ostype.IsDocker()
}
