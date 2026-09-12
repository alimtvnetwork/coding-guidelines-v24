package fileutil

import (
	"os"
	"path/filepath"
	"testing"

	"coding-guidelines/common/pkg/enum/filepermtype"
)

func TestUserTempDir(t *testing.T) {
	res := UserTempDir()
	if res.IsFailed() {
		t.Fatalf("UserTempDir failed: %v", res.Fault())
	}

	info, err := os.Stat(res.Data())
	if err != nil {
		t.Fatalf("UserTempDir stat failed: %v", err)
	}

	if !info.IsDir() {
		t.Fatalf("UserTempDir is not a directory: %s", res.Data())
	}
}

func TestUserTempPath(t *testing.T) {
	res := UserTempPath("sub1", "sub2")
	if res.IsFailed() {
		t.Fatalf("UserTempPath failed: %v", res.Fault())
	}

	uDirRes := UserTempDir()
	tempDir := uDirRes.Data()
	expected := filepath.Join(tempDir, "sub1", "sub2")
	if res.Data() != expected {
		t.Fatalf("expected %s, got %s", expected, res.Data())
	}
}

func TestCreateTempFile(t *testing.T) {
	res := CreateTempFile("", "test-file-*.txt", filepermtype.Standard)
	if res.IsFailed() {
		t.Fatalf("CreateTempFile failed: %v", res.Fault())
	}

	f := res.Data()
	defer os.Remove(f.Name())
	defer f.Close()

	if _, err := f.WriteString("content"); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
}

func TestTempFile(t *testing.T) {
	res := TempFile("quick-file-*.txt")
	if res.IsFailed() {
		t.Fatalf("TempFile failed: %v", res.Fault())
	}

	f := res.Data()
	defer os.Remove(f.Name())
	defer f.Close()

	if _, err := os.Stat(f.Name()); err != nil {
		t.Fatalf("TempFile stat failed: %v", err)
	}
}

func TestCreateTempDir(t *testing.T) {
	res := CreateTempDir("", "test-dir-*", filepermtype.Standard)
	if res.IsFailed() {
		t.Fatalf("CreateTempDir failed: %v", res.Fault())
	}

	dir := res.Data()
	defer os.RemoveAll(dir)

	info, err := os.Stat(dir)
	if err != nil {
		t.Fatalf("CreateTempDir stat error: %v", err)
	}

	if !info.IsDir() {
		t.Fatalf("CreateTempDir is not directory: %s", dir)
	}
}

func TestTempDir(t *testing.T) {
	res := TempDir("quick-dir-*")
	if res.IsFailed() {
		t.Fatalf("TempDir failed: %v", res.Fault())
	}

	dir := res.Data()
	defer os.RemoveAll(dir)

	info, err := os.Stat(dir)
	if err != nil {
		t.Fatalf("TempDir stat error: %v", err)
	}

	if !info.IsDir() {
		t.Fatalf("TempDir is not directory: %s", dir)
	}
}

func mockWindowsEnv(k string) string {
	switch k {
	case "TEMP":
		return `C:\MockTemp`
	case "TMP":
		return `C:\MockTmp`
	case "LOCALAPPDATA":
		return `C:\Users\Mock\AppData\Local`
	case "USERPROFILE":
		return `C:\Users\Mock`
	default:
		return ""
	}
}

func TestWindowsTempCandidates(t *testing.T) {
	candidates := windowsTempCandidates(mockWindowsEnv)
	if len(candidates) != 5 {
		t.Fatalf("expected 5 candidates, got %d", len(candidates))
	}

	if candidates[0] != `C:\MockTemp` {
		t.Fatalf("expected C:\\MockTemp, got %s", candidates[0])
	}

	expectedLocal := filepath.Join(`C:\Users\Mock\AppData\Local`, "Temp")
	if candidates[2] != expectedLocal {
		t.Fatalf("expected %s, got %s", expectedLocal, candidates[2])
	}
}

func mockDarwinEnv(k string) string {
	if k == "TMPDIR" {
		return "/var/folders/mock"
	}

	return ""
}

func TestDarwinTempCandidates(t *testing.T) {
	candidates := darwinTempCandidates(mockDarwinEnv)
	if len(candidates) != 3 {
		t.Fatalf("expected 3 candidates, got %d", len(candidates))
	}

	if candidates[0] != "/var/folders/mock" {
		t.Fatalf("expected /var/folders/mock, got %s", candidates[0])
	}

	if candidates[2] != "/tmp" {
		t.Fatalf("expected /tmp, got %s", candidates[2])
	}
}

func mockLinuxEnv(k string) string {
	switch k {
	case "TMPDIR":
		return "/mock/tmpdir"
	case "XDG_RUNTIME_DIR":
		return "/run/user/1000"
	default:
		return ""
	}
}

func TestLinuxTempCandidates_WithCacheTmp(t *testing.T) {
	alwaysExists := func(p string) bool { return true }
	candidates := linuxTempCandidates(mockLinuxEnv, "/home/mock", alwaysExists)

	if len(candidates) != 4 {
		t.Fatalf("expected 4 candidates, got %d", len(candidates))
	}

	expectedCache := filepath.Join("/home/mock", ".cache", "tmp")
	if candidates[2] != expectedCache {
		t.Fatalf("expected %s, got %s", expectedCache, candidates[2])
	}
}

func TestLinuxTempCandidates_WithoutCacheTmp(t *testing.T) {
	neverExists := func(p string) bool { return false }
	candidates := linuxTempCandidates(mockLinuxEnv, "/home/mock", neverExists)

	if len(candidates) != 3 {
		t.Fatalf("expected 3 candidates, got %d", len(candidates))
	}

	if candidates[2] != "/tmp" {
		t.Fatalf("expected /tmp as 3rd candidate, got %s", candidates[2])
	}
}

func TestSelectTempDir_Fallback(t *testing.T) {
	tmpDir := os.TempDir()
	candidates := []string{"", "/path/that/cannot/exist/and/cannot/be/created/??!!", tmpDir}

	dir, err := selectTempDir(candidates)
	if err != nil {
		t.Fatalf("expected fallback to succeed, got error: %v", err)
	}

	if dir != tmpDir {
		t.Fatalf("expected fallback to %s, got %s", tmpDir, dir)
	}
}
