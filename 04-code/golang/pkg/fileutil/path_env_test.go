package fileutil

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExpandTilde_Home(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("skipping test: UserHomeDir unavailable")
	}

	res := ExpandTilde("~")
	if res.IsFailed() {
		t.Fatalf("ExpandTilde(~) failed: %v", res.Fault())
	}

	if res.Data() != home {
		t.Fatalf("expected %s, got %s", home, res.Data())
	}
}

func TestExpandTilde_SubpathSlash(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("skipping test: UserHomeDir unavailable")
	}

	res := ExpandTilde("~/foo/bar")
	if res.IsFailed() {
		t.Fatalf("ExpandTilde failed: %v", res.Fault())
	}

	expected := filepath.Join(home, "foo/bar")
	if res.Data() != expected {
		t.Fatalf("expected %s, got %s", expected, res.Data())
	}
}

func TestExpandTilde_SubpathBackslash(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("skipping test: UserHomeDir unavailable")
	}

	res := ExpandTilde("~\\foo\\bar")
	if res.IsFailed() {
		t.Fatalf("ExpandTilde failed: %v", res.Fault())
	}

	expected := filepath.Join(home, "foo\\bar")
	if res.Data() != expected {
		t.Fatalf("expected %s, got %s", expected, res.Data())
	}
}

func TestExpandTilde_NonTildePaths(t *testing.T) {
	cases := []string{"", "/var/log", "C:\\Windows", "plain/rel", "~other/sub"}
	for _, p := range cases {
		res := ExpandTilde(p)
		if res.IsFailed() {
			t.Fatalf("ExpandTilde(%s) failed: %v", p, res.Fault())
		}

		if res.Data() != p {
			t.Fatalf("expected %s unchanged, got %s", p, res.Data())
		}
	}
}

func TestExpandEnv_PosixDollar(t *testing.T) {
	t.Setenv("FILEUTIL_ENV_TEST_A", "valA")

	res := ExpandEnv("data/$FILEUTIL_ENV_TEST_A/sub")
	if res.IsFailed() {
		t.Fatalf("ExpandEnv failed: %v", res.Fault())
	}

	if res.Data() != "data/valA/sub" {
		t.Fatalf("expected data/valA/sub, got %s", res.Data())
	}
}

func TestExpandEnv_PosixBraced(t *testing.T) {
	t.Setenv("FILEUTIL_ENV_TEST_B", "valB")

	res := ExpandEnv("prefix_${FILEUTIL_ENV_TEST_B}_suffix")
	if res.IsFailed() {
		t.Fatalf("ExpandEnv failed: %v", res.Fault())
	}

	if res.Data() != "prefix_valB_suffix" {
		t.Fatalf("expected prefix_valB_suffix, got %s", res.Data())
	}
}

func TestExpandEnv_WindowsPercent(t *testing.T) {
	t.Setenv("FILEUTIL_ENV_TEST_C", "valC")

	res := ExpandEnv(`C:\data\%FILEUTIL_ENV_TEST_C%\sub`)
	if res.IsFailed() {
		t.Fatalf("ExpandEnv failed: %v", res.Fault())
	}

	expected := `C:\data\valC\sub`
	if res.Data() != expected {
		t.Fatalf("expected %s, got %s", expected, res.Data())
	}
}

func TestExpandEnv_WindowsCaseInsensitive(t *testing.T) {
	t.Setenv("FILEUTIL_ENV_UPPER", "upperVal")

	res := ExpandEnv("%fileutil_env_upper%/dir")
	if res.IsFailed() {
		t.Fatalf("ExpandEnv failed: %v", res.Fault())
	}

	if res.Data() != "upperVal/dir" {
		t.Fatalf("expected upperVal/dir, got %s", res.Data())
	}
}

func TestExpandEnv_WindowsSpecialChars(t *testing.T) {
	t.Setenv("MY_APP(X86)", "x86_path")

	res := ExpandEnv("%MY_APP(X86)%/bin")
	if res.IsFailed() {
		t.Fatalf("ExpandEnv failed: %v", res.Fault())
	}

	if res.Data() != "x86_path/bin" {
		t.Fatalf("expected x86_path/bin, got %s", res.Data())
	}
}

func TestExpandEnv_UndefinedSafe(t *testing.T) {
	resPosix := ExpandEnv("start/$DEFINITELY_NOT_EXISTING_ENV_VAR/end")
	if resPosix.IsFailed() {
		t.Fatalf("unexpected failure: %v", resPosix.Fault())
	}

	if resPosix.Data() != "start//end" {
		t.Fatalf("expected empty expansion for undefined var, got %s", resPosix.Data())
	}

	resWin := ExpandEnv("start/%DEFINITELY_NOT_EXISTING_ENV_VAR%/end")
	if resWin.IsFailed() {
		t.Fatalf("unexpected failure: %v", resWin.Fault())
	}

	if resWin.Data() != "start//end" {
		t.Fatalf("expected empty expansion for undefined win var, got %s", resWin.Data())
	}
}

func TestExpandEnv_Literals(t *testing.T) {
	res1 := ExpandEnv("100% discount")
	if res1.Data() != "100% discount" {
		t.Fatalf("expected 100%% discount, got %s", res1.Data())
	}

	res2 := ExpandEnv("cost is $5")
	if res2.Data() != "cost is $5" {
		t.Fatalf("expected cost is $5, got %s", res2.Data())
	}

	res3 := ExpandEnv("")
	if res3.Data() != "" {
		t.Fatalf("expected empty string, got %s", res3.Data())
	}
}

func TestExpand_Combined(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("skipping test: UserHomeDir unavailable")
	}

	t.Setenv("FILEUTIL_EXPAND_SUB", "my_project")

	res := Expand("~/projects/$FILEUTIL_EXPAND_SUB")
	if res.IsFailed() {
		t.Fatalf("Expand failed: %v", res.Fault())
	}

	expected := filepath.Join(home, "projects", "my_project")
	if res.Data() != expected {
		t.Fatalf("expected %s, got %s", expected, res.Data())
	}
}
